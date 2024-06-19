package packet

import (
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
		return
	}
	packet.Address, err = r.ReadStringLimited(config.LanternConfig.BufferConfig.HandshakeAddress)
	if err != nil {
		return
	}
	packet.Port, err = r.ReadUInt16()
	if err != nil {
		return
	}
	state, err := r.ReadVarInt()
	if err != nil {
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
		return
	}
	packet.Position.Y, err = r.ReadFloat64()
	if err != nil {
		return
	}
	packet.Position.Z, err = r.ReadFloat64()
	if err != nil {
		return
	}
	yaw, err := r.ReadFloat32()
	if err != nil {
		return
	}
	packet.Position.Yaw = float32(math.Mod((math.Mod(float64(yaw), 360) + 360), 360))

	packet.Position.Pitch, err = r.ReadFloat32()
	if err != nil {
		return
	}
	packet.OnGround, err = r.ReadBool()
	if err != nil {
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
		return
	}
	packet.Yaw = float32(math.Mod((math.Mod(float64(yaw), 360) + 360), 360))
	packet.Pitch, err = r.ReadFloat32()
	if err != nil {
		return
	}
	packet.OnGround, err = r.ReadBool()
	if err != nil {
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
		return
	}
	packet.FeetY, err = r.ReadFloat64()
	if err != nil {
		return
	}
	packet.Z, err = r.ReadFloat64()
	if err != nil {
		return
	}
	packet.OnGround, err = r.ReadBool()
	if err != nil {
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
		return
	}
	action, err := r.ReadVarInt()
	if err != nil {
		return
	}
	packet.ActionID = types.EntityAction(action)
	packet.ActionParameter, err = r.ReadVarInt()
	return
}

type PacketPlaySwingArm struct {
}

func (packet *PacketPlaySwingArm) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
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
		return
	}
	packet.Status = types.DiggingStatus(diggingStatus)
	packet.Location, err = r.ReadBlockPosition()
	if err != nil {
		return
	}
	blockFace, err := r.ReadByte()
	if err != nil {
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
		return
	}
	packet.ClientSettings.ViewDistance, err = r.ReadByte()
	if err != nil {
		return
	}
	chatMode, err := r.ReadByte()
	if err != nil {
		return
	}
	packet.ClientSettings.ChatMode = types.ChatMode(chatMode)
	packet.ClientSettings.ChatColors, err = r.ReadBool()
	if err != nil {
		return
	}
	displaySkinParts, err := r.ReadByte()
	if err != nil {
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
		return
	}
	entityType, err := r.ReadVarInt()
	if err != nil {
		return
	}
	packet.Type = types.UseEntityType(entityType)
	if packet.Type == types.UseEntityInteractAt {
		packet.TargetX, err = r.ReadFloat32()
		if err != nil {
			return
		}
		packet.TargetY, err = r.ReadFloat32()
		if err != nil {
			return
		}
		packet.TargetZ, err = r.ReadFloat32()
		if err != nil {
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
		return
	}
	packet.Slot, err = r.ReadUInt16()
	if err != nil {
		return
	}
	packet.Button, err = r.ReadUInt8()
	if err != nil {
		return
	}
	packet.ActionNumber, err = r.ReadUInt16()
	if err != nil {
		return
	}
	packet.Mode, err = r.ReadUInt8()
	if err != nil {
		return
	}
	packet.ClickedItem, err = r.ReadSlotItem()
	if err != nil {
		return
	}
	return
}

type PacketPlayBlockPlacement struct {
	Location   world.BlockPosition
	Face       types.BlockFace
	HeldItem   slot.SlotItem
	CursorPosX uint8
	CursorPosY uint8
	CursorPosZ uint8
}

func (packet *PacketPlayBlockPlacement) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.Location, err = r.ReadBlockPosition()
	if err != nil {
		return
	}
	blockFace, err := r.ReadUInt8()
	if err != nil {
		return
	}
	packet.Face = types.BlockFace(blockFace)
	packet.HeldItem, err = r.ReadSlotItem()
	if err != nil {
		return
	}
	packet.CursorPosX, err = r.ReadUInt8()
	if err != nil {
		return
	}
	packet.CursorPosY, err = r.ReadUInt8()
	if err != nil {
		return
	}
	packet.CursorPosZ, err = r.ReadUInt8()
	if err != nil {
		return
	}
	return
}

type PacketPlayPlayerAbilitiesServer struct {
	Flags        uint8
	FlyingSpeed  float32
	WalkingSpeed float32
}

func (packet *PacketPlayPlayerAbilitiesServer) Read(r *readerwriter.ConnReadWrite, length int) (err error) {
	packet.Flags, err = r.ReadUInt8()
	if err != nil {
		return
	}
	packet.FlyingSpeed, err = r.ReadFloat32()
	if err != nil {
		return
	}
	packet.WalkingSpeed, err = r.ReadFloat32()
	if err != nil {
		return
	}
	return
}
