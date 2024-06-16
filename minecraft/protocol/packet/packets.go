package packet

import (
	"log"
	"math"

	"github.com/git-fal7/sealantern/config"
	"github.com/git-fal7/sealantern/minecraft/player/clientsettings"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/metadata"
	"github.com/git-fal7/sealantern/pkg/readerwriter"
	"github.com/git-fal7/sealantern/pkg/slot"

	"github.com/google/uuid"
)

type PacketHandshake struct {
	Protocol int
	Address  string
	Port     uint16
	State    types.State
}

func (packet *PacketHandshake) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.Protocol, err = r.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Address, err = r.ReadStringLimited(config.LanternConfig.BufferConfig.HandshakeAddress)
	if err != nil {
		log.Print(err)
		return
	}
	packet.Port, err = r.ReadUInt16()
	if err != nil {
		log.Print(err)
		return
	}
	state, err := r.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	packet.State = types.State(state)
	return
}

func (packet *PacketHandshake) Id() int32 {
	return 0x00
}

type PacketStatusRequest struct{}

func (packet *PacketStatusRequest) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	return
}

func (packet *PacketStatusRequest) Id() int32 {
	return 0x00
}

type PacketStatusResponse struct {
	Response string
}

func (packet *PacketStatusResponse) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	return
}
func (packet *PacketStatusResponse) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteString(packet.Response)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketStatusResponse) Id() int32 {
	return 0x00
}

type PacketStatusPing struct {
	Time uint64
}

func (packet *PacketStatusPing) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.Time, err = r.ReadUInt64()
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketStatusPing) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteUInt64(packet.Time)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketStatusPing) Id() int32 {
	return 0x01
}

type PacketLoginStart struct {
	Username string
}

func (packet *PacketLoginStart) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.Username, err = r.ReadStringLimited(config.LanternConfig.BufferConfig.PlayerName)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketLoginStart) Id() int32 {
	return 0x00
}

type PacketLoginDisconnect struct {
	Component string
}

func (packet *PacketLoginDisconnect) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteString(packet.Component)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketLoginDisconnect) Id() int32 {
	return 0x00
}

type PacketLoginSuccess struct {
	UUID     uuid.UUID
	Username string
}

func (packet *PacketLoginSuccess) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteString(packet.UUID.String())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteString(packet.Username)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketLoginSuccess) Id() int32 {
	return 0x02
}

type PacketLoginSetCompression struct {
	Threshold int
}

func (packet *PacketLoginSetCompression) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(packet.Threshold)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketLoginSetCompression) Id() int32 {
	return 0x03
}

type PacketPlayChat struct {
	Message string
}

func (packet *PacketPlayChat) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.Message, err = r.ReadString()
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayChat) Id() int32 {
	return 0x01
}

type PacketPlayTabComplete struct {
	Matches []string
}

func (packet *PacketPlayTabComplete) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(len(packet.Matches))
	if err != nil {
		log.Print(err)
		return
	}
	for _, s := range packet.Matches {
		err = w.WriteString(s)
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}

func (packet *PacketPlayTabComplete) Id() int32 {
	return 0x3A
}

type PacketPlayClientStatus struct {
	Action types.ClientStatusAction
}

func (packet *PacketPlayClientStatus) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	act, err := r.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Action = types.ClientStatusAction(act)
	return
}

func (packet *PacketPlayClientStatus) Id() int32 {
	return 0x16
}

type PacketPlayMessage struct {
	Component string
	Position  types.ChatPosition
}

func (packet *PacketPlayMessage) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteString(packet.Component)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(uint8(packet.Position))
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayMessage) Id() int32 {
	return 0x02
}

type PacketPlayServerDifficulty struct {
	Difficulty world.Difficulty
}

func (packet *PacketPlayServerDifficulty) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteUInt8(uint8(packet.Difficulty))
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayServerDifficulty) Id() int32 {
	return 0x41
}

type PacketPlayPluginMessage struct {
	Channel string
	Data    []byte
}

func (packet *PacketPlayPluginMessage) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.Channel, err = r.ReadString()
	if err != nil {
		log.Print(err)
		return
	}

	// This should fix the invalid packet issues.
	dataLength := length - 10

	packet.Data, err = r.ReadByteArray(dataLength)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayPluginMessage) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteString(packet.Channel)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteByteArray(packet.Data)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayPluginMessage) Id() int32 {
	return 0x3F
}

type PacketPlayDisconnect struct {
	Component string
}

func (packet *PacketPlayDisconnect) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.Component, err = r.ReadString()
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayDisconnect) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteString(packet.Component)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayDisconnect) Id() int32 {
	return 0x40
}

type PacketPlayKeepAlive struct {
	Identifier int
}

func (packet *PacketPlayKeepAlive) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.Identifier, err = r.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayKeepAlive) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(packet.Identifier)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayKeepAlive) Id() int32 {
	return 0x00
}

type PacketPlayParticle struct {
	Type         int
	LongDistance bool
	X            float32
	Y            float32
	Z            float32
	OffsetX      float32
	OffsetY      float32
	OffsetZ      float32
	ParticleData float32
	Count        int
	Data         []int
}

func (packet *PacketPlayParticle) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteUInt32(uint32(packet.Type))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteBool(packet.LongDistance)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat32(packet.X)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat32(packet.Y)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat32(packet.Z)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat32(packet.OffsetX)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat32(packet.OffsetY)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat32(packet.OffsetZ)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat32(packet.ParticleData)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt32(uint32(packet.Count))
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayParticle) Id() int32 {
	return 0x2A
}

type PacketPlayJoinGame struct {
	//EntityId     uint8
	Gamemode     types.Gamemode
	Dimension    world.Dimension
	Difficulty   world.Difficulty
	MaxPlayers   uint8
	LevelType    world.LevelType
	ReducedDebug bool
}

func (packet *PacketPlayJoinGame) Write(w *readerwriter.ConnReadWrite) (err error) {
	// self-id
	err = w.WriteUInt8(0) // entity ids are more than 255,
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(uint8(packet.Gamemode))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt32(uint32(packet.Dimension))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(uint8(packet.Difficulty))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.MaxPlayers)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteString(string(packet.LevelType))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteBool(packet.ReducedDebug)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayJoinGame) Id() int32 {
	return 0x01
}

type PacketPlayPlayerAbilities struct {
	Invulnerable bool
	Fly          bool
	CanFly       bool
	Creative     bool
	FlyingSpeed  float32
	FieldOfView  float32
}

func (packet *PacketPlayPlayerAbilities) Write(w *readerwriter.ConnReadWrite) (err error) {
	var flags uint8 = 0
	if packet.Invulnerable {
		flags |= 0x01
	}
	if packet.Fly {
		flags |= 0x02
	}
	if packet.CanFly {
		flags |= 0x04
	}
	if packet.Creative {
		flags |= 0x08
	}

	err = w.WriteUInt8(flags)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat32(packet.FlyingSpeed)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat32(packet.FieldOfView)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayPlayerAbilities) Id() int32 {
	return 0x39
}

type PacketPlayPlayerPositionAndLookClient struct {
	Position world.Position
	Flags    uint8
}

func (packet *PacketPlayPlayerPositionAndLookClient) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteFloat64(packet.Position.X)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat64(packet.Position.Y)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat64(packet.Position.Z)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat32(packet.Position.Yaw)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat32(packet.Position.Pitch)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.Flags)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayPlayerPositionAndLookClient) Id() int32 {
	return 0x08
}

type PacketPlayUpdateHealth struct {
	Health         float32
	Food           int
	FoodSaturation float32
}

func (packet *PacketPlayUpdateHealth) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteFloat32(packet.Health)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteVarInt(packet.Food)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteFloat32(packet.FoodSaturation)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayUpdateHealth) Id() int32 {
	return 0x06
}

type PacketPlaySpawnPosition struct {
	Position world.Position
}

func (packet *PacketPlaySpawnPosition) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteBlockPosition(packet.Position)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlaySpawnPosition) Id() int32 {
	return 0x05
}

type PacketPlayerListHeaderFooter struct {
	Header *string
	Footer *string
}

func (packet *PacketPlayerListHeaderFooter) Write(w *readerwriter.ConnReadWrite) (err error) {
	var str string
	if packet.Header == nil {
		str = `{"translate":""}`
	} else {
		str = *packet.Header
	}
	err = w.WriteString(str)
	if err != nil {
		log.Print(err)
		return
	}
	if packet.Footer == nil {
		str = `{"translate":""}`
	} else {
		str = *packet.Footer
	}
	err = w.WriteString(str)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayerListHeaderFooter) Id() int32 {
	return 0x47
}

type PacketPlayPlayerPositionAndLook struct {
	Position world.Position
	OnGround bool
}

func (packet *PacketPlayPlayerPositionAndLook) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.Position.X, err = r.ReadFloat64()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Position.Y, err = r.ReadFloat64()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Position.Z, err = r.ReadFloat64()
	if err != nil {
		log.Print(err)
		return
	}
	yaw, err := r.ReadFloat32()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Position.Yaw = float32(math.Mod((math.Mod(float64(yaw), 360) + 360), 360))

	packet.Position.Pitch, err = r.ReadFloat32()
	if err != nil {
		log.Print(err)
		return
	}
	packet.OnGround, err = r.ReadBool()
	if err != nil {
		log.Print(err)
		return
	}

	return
}

func (packet *PacketPlayPlayerPositionAndLook) Id() int32 {
	return 0x06
}

type PacketPlayPlayerLook struct {
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (packet *PacketPlayPlayerLook) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	yaw, err := r.ReadFloat32()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Yaw = float32(math.Mod((math.Mod(float64(yaw), 360) + 360), 360))
	packet.Pitch, err = r.ReadFloat32()
	if err != nil {
		log.Print(err)
		return
	}
	packet.OnGround, err = r.ReadBool()
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayPlayerLook) Id() int32 {
	return 0x05
}

type PacketPlayPlayerPosition struct {
	X        float64
	FeetY    float64
	Z        float64
	OnGround bool
}

func (packet *PacketPlayPlayerPosition) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.X, err = r.ReadFloat64()
	if err != nil {
		log.Print(err)
		return
	}
	packet.FeetY, err = r.ReadFloat64()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Z, err = r.ReadFloat64()
	if err != nil {
		log.Print(err)
		return
	}
	packet.OnGround, err = r.ReadBool()
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayPlayerPosition) Id() int32 {
	return 0x04
}

type PacketPlayPlayer struct {
	OnGround bool
}

func (packet *PacketPlayPlayer) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.OnGround, err = r.ReadBool()
	if err != nil {
		log.Print(err)
		return
	}

	return
}

func (packet *PacketPlayPlayer) Id() int32 {
	return 0x03
}

type PacketPlayEntityAction struct {
	EntityID        int
	ActionID        types.EntityAction
	ActionParameter int
}

func (packet *PacketPlayEntityAction) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.EntityID, err = r.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	action, err := r.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	packet.ActionID = types.EntityAction(action)
	packet.ActionParameter, err = r.ReadVarInt()
	return
}

func (packet *PacketPlayEntityAction) Id() int32 {
	return 0x0B
}

type PacketPlayAnimationServer struct {
}

func (packet *PacketPlayAnimationServer) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	return
}

func (packet *PacketPlayAnimationServer) Id() int32 {
	return 0x0A
}

type PacketPlayPlayerDigging struct {
	Status   types.DiggingStatus
	Location world.BlockPosition
	Face     types.BlockFace
}

func (packet *PacketPlayPlayerDigging) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	diggingStatus, err := r.ReadByte()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Status = types.DiggingStatus(diggingStatus)
	packet.Location, err = r.ReadBlockPosition()
	if err != nil {
		log.Print(err)
		return
	}
	blockFace, err := r.ReadByte()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Face = types.BlockFace(blockFace)
	return
}

func (packet *PacketPlayPlayerDigging) Id() int32 {
	return 0x07
}

type PacketPlayClientSettings struct {
	ClientSettings clientsettings.ClientSettings
}

func (packet *PacketPlayClientSettings) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.ClientSettings.Locale, err = r.ReadString()
	if err != nil {
		log.Print(err)
		return
	}
	packet.ClientSettings.ViewDistance, err = r.ReadByte()
	if err != nil {
		log.Print(err)
		return
	}
	chatMode, err := r.ReadByte()
	if err != nil {
		log.Print(err)
		return
	}
	packet.ClientSettings.ChatMode = types.ChatMode(chatMode)
	packet.ClientSettings.ChatColors, err = r.ReadBool()
	if err != nil {
		log.Print(err)
		return
	}
	displaySkinParts, err := r.ReadByte()
	if err != nil {
		log.Print(err)
		return
	}
	packet.ClientSettings.DisplayedSkinParts = types.DisplayedSkinParts(displaySkinParts)
	return
}

func (packet *PacketPlayClientSettings) Id() int32 {
	return 0x15
}

type PacketPlayPlayerListItem struct {
	Action  types.PlayerListAction
	Entries []types.PlayerListEntry
}

func (packet *PacketPlayPlayerListItem) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(int(packet.Action))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteVarInt(len(packet.Entries))
	if err != nil {
		log.Print(err)
		return
	}

	for _, entry := range packet.Entries {
		err = w.WriteUUID(entry.Profile.UUID)
		if err != nil {
			log.Print(err)
			return
		}
		switch packet.Action {
		case types.PlayerListActionAddPlayer:
			{
				err = w.WriteString(entry.Profile.Name)
				if err != nil {
					log.Print(err)
					return
				}
				err = w.WriteVarInt(len(entry.Profile.Properties))
				if err != nil {
					log.Print(err)
					return
				}
				for _, property := range entry.Profile.Properties {
					err = w.WriteString(property.Name)
					if err != nil {
						log.Print(err)
						return
					}
					err = w.WriteString(property.Value)
					if err != nil {
						log.Print(err)
						return
					}
					if property.Signature != "" {
						err = w.WriteBool(true)
						if err != nil {
							log.Print(err)
							return
						}
						err = w.WriteString(property.Signature)
						if err != nil {
							log.Print(err)
							return
						}
					} else {
						err = w.WriteBool(false)
						if err != nil {
							log.Print(err)
							return
						}
					}
				}
				err = w.WriteVarInt(int(entry.GameMode))
				if err != nil {
					log.Print(err)
					return
				}
				err = w.WriteVarInt(entry.Ping)
				if err != nil {
					log.Print(err)
					return
				}
				if entry.DisplayName != nil {
					err = w.WriteBool(true)
					if err != nil {
						log.Print(err)
						return
					}
					var json string
					json, err = entry.DisplayName.JSON()
					if err != nil {
						log.Print(err)
						return
					}
					err = w.WriteString(json)
					if err != nil {
						log.Print(err)
						return
					}

				} else {
					err = w.WriteBool(false)
					if err != nil {
						log.Print(err)
						return
					}
				}

			}
		case types.PlayerListActionUpdateGamemode:
			{
				err = w.WriteVarInt(int(entry.GameMode))
				if err != nil {
					log.Print(err)
					return
				}
			}
		case types.PlayerListActionUpdateLatency:
			{
				err = w.WriteVarInt(entry.Ping)
				if err != nil {
					log.Print(err)
					return
				}
			}
		case types.PlayerListActionUpdateDisplayName:
			{
				if entry.DisplayName != nil {
					err = w.WriteBool(true)
					if err != nil {
						log.Print(err)
						return
					}
					var json string
					json, err = entry.DisplayName.JSON()
					if err != nil {
						log.Print(err)
						return
					}
					err = w.WriteString(json)
					if err != nil {
						log.Print(err)
						return
					}
				} else {
					err = w.WriteBool(false)
					if err != nil {
						log.Print(err)
						return
					}
				}
			}
		}
		// Remove player doesn't have any fields. so we dont even need to put it in the switch case
	}
	return
}

func (packet *PacketPlayPlayerListItem) Id() int32 {
	return 0x38
}

type PacketPlaySpawnPlayer struct {
	EntityID       uint16
	PlayerUUID     uuid.UUID
	PlayerPosition world.Position
	CurrentItem    uint16
}

func (packet *PacketPlaySpawnPlayer) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUUID(packet.PlayerUUID)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt32(packet.PlayerPosition.IntX())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt32(packet.PlayerPosition.IntY())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt32(packet.PlayerPosition.IntZ())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.PlayerPosition.IntYaw())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.PlayerPosition.IntPitch())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt16(packet.CurrentItem)
	if err != nil {
		log.Print(err)
		return
	}

	// metadata end
	err = w.WriteUInt8(127)
	return
}

func (packet *PacketPlaySpawnPlayer) Id() int32 {
	return 0x0C
}

type PacketPlayDestroyEntities struct {
	EntityIDs []uint16
}

func (packet *PacketPlayDestroyEntities) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(len(packet.EntityIDs))
	if err != nil {
		log.Print(err)
		return
	}
	for _, enitiyId := range packet.EntityIDs {
		err = w.WriteVarInt(int(enitiyId))
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}

func (packet *PacketPlayDestroyEntities) Id() int32 {
	return 0x13
}

type PacketPlayEntityRelativeMove struct {
	EntityID uint16
	DeltaX   int8
	DeltaY   int8
	DeltaZ   int8
	OnGround bool
}

func (packet *PacketPlayEntityRelativeMove) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(uint8(packet.DeltaX))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(uint8(packet.DeltaY))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(uint8(packet.DeltaZ))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteBool(packet.OnGround)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayEntityRelativeMove) Id() int32 {
	return 0x15
}

type PacketPlayEntityLook struct {
	EntityID uint16
	Yaw      uint8
	Pitch    uint8
	OnGround bool
}

func (packet *PacketPlayEntityLook) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.Yaw)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.Pitch)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteBool(packet.OnGround)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayEntityLook) Id() int32 {
	return 0x16
}

type PacketPlayEntityLookAndRelativeMove struct {
	EntityID uint16
	DeltaX   int8
	DeltaY   int8
	DeltaZ   int8
	Yaw      uint8
	Pitch    uint8
	OnGround bool
}

func (packet *PacketPlayEntityLookAndRelativeMove) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(uint8(packet.DeltaX))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(uint8(packet.DeltaY))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(uint8(packet.DeltaZ))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.Yaw)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.Pitch)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteBool(packet.OnGround)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayEntityLookAndRelativeMove) Id() int32 {
	return 0x17
}

type PacketPlayEntityHeadLook struct {
	EntityID uint16
	HeadYaw  uint8
}

func (packet *PacketPlayEntityHeadLook) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.HeadYaw)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayEntityHeadLook) Id() int32 {
	return 0x19
}

type PacketPlayEntityMetadata struct {
	EntityID uint16
	Metadata metadata.MetadataMap
}

func (packet *PacketPlayEntityMetadata) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteMetadata(packet.Metadata)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayEntityMetadata) Id() int32 {
	return 0x1C
}

type PacketPlayAnimationClient struct {
	EntityID  uint16
	Animation types.Animation
}

func (packet *PacketPlayAnimationClient) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(uint8(packet.Animation))
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayAnimationClient) Id() int32 {
	return 0x0B
}

type PacketPlayUseEntity struct {
	TargetID int
	Type     types.UseEntityType
	TargetX  float32
	TargetY  float32
	TargetZ  float32
}

func (packet *PacketPlayUseEntity) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.TargetID, err = r.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	entityType, err := r.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Type = types.UseEntityType(entityType)
	if packet.Type == types.UseEntityInteractAt {
		packet.TargetX, err = r.ReadFloat32()
		if err != nil {
			log.Print(err)
			return
		}
		packet.TargetY, err = r.ReadFloat32()
		if err != nil {
			log.Print(err)
			return
		}
		packet.TargetZ, err = r.ReadFloat32()
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}
func (packet *PacketPlayUseEntity) Id() int32 {
	return 0x02
}

type PacketPlayEntityVelocity struct {
	EntityID uint16
	Velocity world.Vector
}

func (packet *PacketPlayEntityVelocity) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt16(int16(packet.Velocity.X * 8000))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt16(int16(packet.Velocity.Y * 8000))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt16(int16(packet.Velocity.Z * 8000))
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayEntityVelocity) Id() int32 {
	return 0x12
}

type PacketPlayMapChunkBulk struct {
	Packets  []PacketPlayChunkData
	Skylight bool
}

func (packet *PacketPlayMapChunkBulk) Write(w *readerwriter.ConnReadWrite) (err error) {
	w.WriteBool(packet.Skylight)
	w.WriteVarInt(len(packet.Packets))
	for _, entry := range packet.Packets {
		w.WriteInt32(entry.X)
		w.WriteInt32(entry.Z)
		w.WriteUInt16(entry.SectionBitMask)
	}
	for _, entry := range packet.Packets {
		w.WriteByteArray(entry.Data)
	}
	return
}

func (packet *PacketPlayMapChunkBulk) Id() int32 {
	return 0x26
}

type PacketPlayChunkData struct {
	X              int32
	Z              int32
	GroundUp       bool
	SectionBitMask uint16
	Data           []byte
}

func (packet *PacketPlayChunkData) Write(w *readerwriter.ConnReadWrite) (err error) {
	w.WriteInt32(packet.X)
	w.WriteInt32(packet.Z)
	w.WriteBool(packet.GroundUp)
	w.WriteUInt16(packet.SectionBitMask)
	w.WriteVarInt(len(packet.Data))
	w.WriteByteArray(packet.Data)
	return
}

func (packet *PacketPlayChunkData) Id() int32 {
	return 0x21
}

type PacketPlayScoreboardObjective struct {
	ObjectiveName string
	DisplayName   string
	Mode          types.ScoreboardMode
	RenderType    types.ObjectiveRenderType
}

func (packet *PacketPlayScoreboardObjective) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteString(packet.ObjectiveName)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt8(int8(packet.Mode))
	if err != nil {
		log.Print(err)
		return
	}
	if packet.Mode == types.CreateScoreboard || packet.Mode == types.UpdateScoreboard {
		err = w.WriteString(packet.DisplayName)
		if err != nil {
			log.Print(err)
			return
		}
		err = w.WriteString(string(packet.RenderType))
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}

func (packet *PacketPlayScoreboardObjective) Id() int32 {
	return 0x3B
}

type PacketPlayUpdateScore struct {
	ScoreName     string
	Action        types.UpdateScoreAction
	ObjectiveName string
	Value         int
}

func (packet *PacketPlayUpdateScore) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteString(packet.ScoreName)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt8(int8(packet.Action))
	if err != nil {
		log.Print(err)
		return
	}

	err = w.WriteString(packet.ObjectiveName)
	if err != nil {
		log.Print(err)
		return
	}
	if packet.Action != types.RemoveScoreItem {
		err = w.WriteVarInt(packet.Value)
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}

func (packet *PacketPlayUpdateScore) Id() int32 {
	return 0x3C
}

type PacketPlayDisplayScoreboard struct {
	Position  types.ObjectiveDisplaySlot
	ScoreName string
}

func (packet *PacketPlayDisplayScoreboard) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteInt8(int8(packet.Position))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteString(packet.ScoreName)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayDisplayScoreboard) Id() int32 {
	return 0x3D
}

type PacketPlayTeams struct {
	TeamName          string
	Mode              types.TeamMode
	TeamDisplayName   string
	TeamPrefix        string
	TeamSuffix        string
	FriendlyFire      types.TeamFriendlyFire
	NameTagVisibility types.TeamNameTagVisibility
	Color             int8
	Players           []string
}

func (packet *PacketPlayTeams) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteString(packet.TeamName)
	if err != nil {
		log.Print(err)
		return
	}

	err = w.WriteInt8(int8(packet.Mode))
	if err != nil {
		log.Print(err)
		return
	}
	if packet.Mode == types.TeamModeCreate || packet.Mode == types.TeamModeUpdate {
		err = w.WriteString(packet.TeamDisplayName)
		if err != nil {
			log.Print(err)
			return
		}
		err = w.WriteString(packet.TeamPrefix)
		if err != nil {
			log.Print(err)
			return
		}
		err = w.WriteString(packet.TeamSuffix)
		if err != nil {
			log.Print(err)
			return
		}
		err = w.WriteInt8(int8(packet.FriendlyFire))
		if err != nil {
			log.Print(err)
			return
		}
		err = w.WriteString(string(packet.NameTagVisibility))
		if err != nil {
			log.Print(err)
			return
		}
		err = w.WriteInt8(int8(packet.Color))
		if err != nil {
			log.Print(err)
			return
		}
	}
	if packet.Mode == types.TeamModeCreate || packet.Mode == types.TeamModeAddPlayer || packet.Mode == types.TeamModeRemovePlayer {
		if packet.Players != nil {
			err = w.WriteVarInt(len(packet.Players))
			if err != nil {
				log.Print(err)
				return
			}
			for _, player := range packet.Players {
				err = w.WriteString(player)
				if err != nil {
					log.Print(err)
					return
				}
			}
		}
	}
	return
}

func (packet *PacketPlayTeams) Id() int32 {
	return 0x3E
}

type PacketPlaySpawnMob struct {
	EntityID  uint16
	MobType   types.MobType
	Position  world.Position
	HeadPitch uint8
	VelocityX uint16
	VelocityY uint16
	VelocityZ uint16
	Metadata  metadata.MetadataMap
}

func (packet *PacketPlaySpawnMob) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(uint8(packet.MobType))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt32(packet.Position.IntX())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt32(packet.Position.IntY())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt32(packet.Position.IntZ())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.Position.IntYaw())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.Position.IntPitch())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.HeadPitch)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt16(packet.VelocityX)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt16(packet.VelocityY)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt16(packet.VelocityZ)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteMetadata(packet.Metadata)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlaySpawnMob) Id() int32 {
	return 0x0F
}

type PacketPlaySpawnObject struct {
	EntityID   uint16
	ObjectType types.ObjectType
	Position   world.Position
	Data       int
	VelocityX  uint16
	VelocityY  uint16
	VelocityZ  uint16
}

func (packet *PacketPlaySpawnObject) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(uint8(packet.ObjectType))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt32(packet.Position.IntX())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt32(packet.Position.IntY())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt32(packet.Position.IntZ())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.Position.IntPitch())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.Position.IntYaw())
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt32(int32(packet.Data))
	if err != nil {
		log.Print(err)
		return
	}
	if packet.Data != 0 {
		err = w.WriteUInt16(packet.VelocityX)
		if err != nil {
			log.Print(err)
			return
		}
		err = w.WriteUInt16(packet.VelocityY)
		if err != nil {
			log.Print(err)
			return
		}
		err = w.WriteUInt16(packet.VelocityZ)
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}

func (packet *PacketPlaySpawnObject) Id() int32 {
	return 0x0E
}

type PacketPlayOpenWindow struct {
	WindowID      uint8
	WindowType    types.WindowType
	WindowTitle   string
	NumberOfSlots uint8
	HorseEntityID uint16
}

func (packet *PacketPlayOpenWindow) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteUInt8(packet.WindowID)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteString(string(packet.WindowType))
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteString(packet.WindowTitle)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt8(packet.NumberOfSlots)
	if err != nil {
		log.Print(err)
		return
	}
	if packet.WindowType == types.WindowTypeEntityHorse {
		err = w.WriteVarInt(int(packet.HorseEntityID))
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}

func (packet *PacketPlayOpenWindow) Id() int32 {
	return 0x2D
}

type PacketPlayWindowItems struct {
	WindowID uint8
	SlotData []slot.SlotItem
}

func (packet *PacketPlayWindowItems) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteUInt8(packet.WindowID)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt16(uint16(len(packet.SlotData)))
	if err != nil {
		log.Print(err)
		return
	}
	for _, slot := range packet.SlotData {
		err = w.WriteSlotItem(slot)
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}

func (packet *PacketPlayWindowItems) Id() int32 {
	return 0x30
}

type PacketPlayClickWindow struct {
	WindowID     uint8
	Slot         uint16
	Button       uint8
	ActionNumber uint16
	Mode         uint8
	ClickedItem  slot.SlotItem
}

func (packet *PacketPlayClickWindow) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.WindowID, err = r.ReadUInt8()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Slot, err = r.ReadUInt16()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Button, err = r.ReadUInt8()
	if err != nil {
		log.Print(err)
		return
	}
	packet.ActionNumber, err = r.ReadUInt16()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Mode, err = r.ReadUInt8()
	if err != nil {
		log.Print(err)
		return
	}
	packet.ClickedItem, err = r.ReadSlotItem()
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayClickWindow) Id() int32 {
	return 0x0E
}

type PacketPlayConfirmTransaction struct {
	WindowID     uint8
	ActionNumber uint16
	Accepted     bool
}

func (packet *PacketPlayConfirmTransaction) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.WindowID, err = r.ReadUInt8()
	if err != nil {
		log.Print(err)
		return
	}
	packet.ActionNumber, err = r.ReadUInt16()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Accepted, err = r.ReadBool()
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayConfirmTransaction) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteUInt8(packet.WindowID)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteUInt16(packet.ActionNumber)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteBool(packet.Accepted)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlayConfirmTransaction) Id() int32 {
	return 0x32
}

type PacketPlaySetSlot struct {
	WindowID uint8
	Slot     int16
	SlotData slot.SlotItem
}

func (packet *PacketPlaySetSlot) Write(w *readerwriter.ConnReadWrite) (err error) {
	err = w.WriteUInt8(packet.WindowID)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteInt16(packet.Slot)
	if err != nil {
		log.Print(err)
		return
	}
	err = w.WriteSlotItem(packet.SlotData)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (packet *PacketPlaySetSlot) Id() int32 {
	return 0x2F
}
