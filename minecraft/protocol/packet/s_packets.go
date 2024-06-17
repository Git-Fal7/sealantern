package packet

import (
	"log"
	"math"

	"github.com/git-fal7/sealantern/config"
	"github.com/git-fal7/sealantern/minecraft/player/clientsettings"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/pkg/readerwriter"
	"github.com/git-fal7/sealantern/pkg/slot"
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

type PacketStatusRequest struct{}

func (packet *PacketStatusRequest) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	return
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

type PacketLoginDisconnect struct {
	Component string
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

type PacketPlayAnimationServer struct {
}

func (packet *PacketPlayAnimationServer) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	return
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
