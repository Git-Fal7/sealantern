package sealantern

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"reflect"
	"time"

	"github.com/git-fal7/sealantern/config"
	"github.com/git-fal7/sealantern/minecraft/player"
	"github.com/git-fal7/sealantern/minecraft/player/connplayer"
	"github.com/git-fal7/sealantern/minecraft/player/socket"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/protocol/packethandler"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/chunk"
	"github.com/git-fal7/sealantern/pkg/command"
	"github.com/git-fal7/sealantern/pkg/events"
	"github.com/git-fal7/sealantern/pkg/gameinstance"

	"github.com/google/uuid"
	"github.com/robinbraemer/event"
)

type Core struct {
	connCounter    int
	brand          string
	playerRegistry *player.PlayerRegistry
	instances      map[string]*gameinstance.GameInstance
	commandMgr     *command.Manager
	eventMgr       event.Manager
	eid            uint16
}

func Init() *Core {
	config.InitConfig()
	c := &Core{
		connCounter:    0,
		brand:          "SeaLantern",
		playerRegistry: player.NewPlayerRegistry(),
		instances:      make(map[string]*gameinstance.GameInstance),
		commandMgr:     command.NewManager(),
		eventMgr:       event.New(),
		eid:            0,
	}
	packethandler.InitRegistry(c)
	packet.InitRegistry()
	return c
}

func (c *Core) Start() {
	ln, err := net.Listen("tcp", config.LanternConfig.ListenAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server launched on port", config.LanternConfig.ListenAddress)
	go c.keepAlive()
	go c.tick()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
		} else {
			c.connCounter += 1
			socketConn := socket.NewConn(conn)
			go c.handleConnection(socketConn, c.connCounter)
		}
	}
}

func (c *Core) SetBrand(brand string) {
	br := make([]byte, len(brand)+1)
	copy(br[:len(brand)], []byte(brand))
	c.brand = string(br)
}

func (c *Core) Brand() string {
	return c.brand
}

func (c *Core) GetPlayerRegistry() *player.PlayerRegistry {
	return c.playerRegistry
}

func (c *Core) Event() event.Manager {
	return c.eventMgr
}

func (c *Core) Command() *command.Manager {
	return c.commandMgr
}

// "tick" is a wrong term, but whatever.
// basically tells all instances to "tick", which basically load all chunks for players, for now
func (c *Core) tick() {
	for {
		for _, instance := range c.instances {
			instance.Tick()
		}
		time.Sleep(500 * time.Millisecond) // 2 ticks / s
	}
}

func (c *Core) keepAlive() {
	r := rand.New(rand.NewSource(15768735131534))
	keepalive := &packet.PacketPlayKeepAlive{}
	for {
		id := int(r.Int31())
		keepalive.Identifier = id
		for _, player := range c.playerRegistry.GetPlayers() {
			if player.Conn.State == types.PLAY {
				player.Conn.KeepAlive = id
				player.WritePacket(keepalive)
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func (c *Core) handleConnection(conn *socket.Conn, id int) {
	defer conn.Close()

	log.Printf("%s(#%d) connected.", conn.RemoteAddr().String(), id)

	for {
		var err error
		_, err = c.readPacket(conn)
		if err != nil {
			break
		}
	}

	if conn.State == types.PLAY {
		registeredPlayer := c.playerRegistry.GetPlayerFromUUID(conn.UUID)
		if registeredPlayer != nil {
			c.playerRegistry.UnregisterPlayer(registeredPlayer)
			c.Event().Fire(&events.PlayerQuitEvent{
				Player: registeredPlayer,
			})
			for _, instance := range c.instances {
				instance.QuitPlayer(registeredPlayer)
			}
			if registeredPlayer.Team() != nil {
				registeredPlayer.Team().RemovePlayer(registeredPlayer.Username())
			}
		}
	}
	conn.Disconnected = true
	log.Printf("%s(#%d) disconnected.", conn.RemoteAddr().String(), id)
}

func (c *Core) GetConfig(config interface{}) {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, config); err != nil {
		panic(err)
	}
}

func (c *Core) AddGameInstance(name string, gameinstance *gameinstance.GameInstance) error {
	if _, ok := c.instances[name]; ok {
		return fmt.Errorf("already has an instance named %s", name)
	}

	c.instances[name] = gameinstance
	return nil
}

func (c *Core) GetGameInstance(name string) *gameinstance.GameInstance {
	if instance, ok := c.instances[name]; ok {
		return instance
	}
	return nil
}

func (c *Core) NextEID() uint16 {
	c.eid = c.eid + 1
	return c.eid
}

func (c *Core) GetInstanceFromUUID(uuid uuid.UUID) *gameinstance.GameInstance {
	for _, instance := range c.instances {
		if instance.HasPlayerFromUUID(uuid) {
			return instance
		}
	}
	return nil
}

func (c *Core) readPacket(conn *socket.Conn) (packet protocol.PacketIn, err error) {
	return c.readPacketWithoutCompression(conn)
}

func (c *Core) readPacketWithoutCompression(conn *socket.Conn) (packet protocol.PacketIn, err error) {
	length, err := conn.Reader.ReadVarInt()
	if err != nil {
		return
	}

	id, err := conn.Reader.ReadVarInt()
	if err != nil {
		return
	}
	packet, err = conn.HandlePacket(id, length)
	if err != nil {
		return
	} else if packet != nil {
		if config.LanternConfig.Logs {
			if id != 0x03 {
				log.Printf("# -> %d %s %s", id, reflect.TypeOf(packet), fmt.Sprint(packet))
			}
		}
		packethandler.ExecutePacketHandler(conn, packet, id, c.playerRegistry)
	}
	return
}

func (c *Core) SwitchToInstance(p *connplayer.ConnectedPlayer, newInstance *gameinstance.GameInstance) {
	if p == nil {
		return
	}
	instance := c.GetInstanceFromUUID(p.UUID())
	if instance == newInstance {
		return
	}
	if !instance.HasPlayerFromUUID(p.UUID()) {
		return
	}
	instance.QuitPlayer(p)
	entries := make([]types.PlayerListEntry, 0)
	for _, oldInstancePlayers := range instance.Players.GetPlayers() {
		entries = append(entries, types.PlayerListEntry{
			Profile: *oldInstancePlayers.Profile(),
		})
	}
	p.WritePacket(&packet.PacketPlayPlayerListItem{
		Action:  types.PlayerListActionRemovePlayer,
		Entries: entries,
	})
	if instance.World.Dimension == newInstance.World.Dimension {
		p.WritePacket(&packet.PacketPlayRespawn{
			Dimension:  (instance.World.Dimension + 1) % 2,
			Difficulty: newInstance.Difficulty,
			Gamemode:   types.SURVIVAL,
			LevelType:  world.DEFAULT,
		})
	}
	p.WritePacket(&packet.PacketPlayRespawn{
		Dimension:  newInstance.World.Dimension,
		Difficulty: newInstance.Difficulty,
		Gamemode:   types.SURVIVAL,
		LevelType:  world.DEFAULT,
	})
	p.SetGamemode(newInstance.Gamemode)
	p.KnownChunkKeys = make(map[chunk.ChunkKey]bool)
	newInstance.JoinPlayer(p)
	c.Event().Fire(&events.PlayerSwitchInstanceEvent{
		Player:          p,
		CurrentInstance: newInstance,
	})
}
