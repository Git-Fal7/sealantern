package packet

import (
	"math"

	"github.com/git-fal7/sealantern/minecraft/player/clientsettings"
	"github.com/git-fal7/sealantern/minecraft/protocol/stream"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/pkg/slot"
	"github.com/google/uuid"
)

type PacketHandshake struct {
	Protocol int
	Address  string
	Port     uint16
	State    types.State
}

func (packet *PacketHandshake) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.Protocol, err = r.ReadVarInt()
	if err != nil {
		return
	}
	packet.Address, err = r.ReadString()
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

func (packet *PacketStatusRequest) Read(r *stream.ProtocolReader, length int) (err error) {
	return
}

type PacketLoginStart struct {
	Username string
}

func (packet *PacketLoginStart) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.Username, err = r.ReadString()
	if err != nil {
		return
	}
	return
}

type PacketPlayChat struct {
	Message string
}

func (packet *PacketPlayChat) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.Message, err = r.ReadString()
	if err != nil {
		return
	}
	return
}

type PacketPlayClientStatus struct {
	Action types.ClientStatusAction
}

func (packet *PacketPlayClientStatus) Read(r *stream.ProtocolReader, length int) (err error) {
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

func (packet *PacketPlayPlayerPositionAndLook) Read(r *stream.ProtocolReader, length int) (err error) {
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

func (packet *PacketPlayPlayerLook) Read(r *stream.ProtocolReader, length int) (err error) {
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

func (packet *PacketPlayPlayerPosition) Read(r *stream.ProtocolReader, length int) (err error) {
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

func (packet *PacketPlayPlayer) Read(r *stream.ProtocolReader, length int) (err error) {
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

func (packet *PacketPlayEntityAction) Read(r *stream.ProtocolReader, length int) (err error) {
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

func (packet *PacketPlaySwingArm) Read(r *stream.ProtocolReader, length int) (err error) {
	return
}

type PacketPlayPlayerDigging struct {
	Status   types.DiggingStatus
	Location world.BlockPosition
	Face     types.BlockFace
}

func (packet *PacketPlayPlayerDigging) Read(r *stream.ProtocolReader, length int) (err error) {
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

func (packet *PacketPlayClientSettings) Read(r *stream.ProtocolReader, length int) (err error) {
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

func (packet *PacketPlayUseEntity) Read(r *stream.ProtocolReader, length int) (err error) {
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

func (packet *PacketPlayClickWindow) Read(r *stream.ProtocolReader, length int) (err error) {
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

func (packet *PacketPlayBlockPlacement) Read(r *stream.ProtocolReader, length int) (err error) {
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

func (packet *PacketPlayPlayerAbilitiesServer) Read(r *stream.ProtocolReader, length int) (err error) {
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

type PacketPlaySteerVehicle struct {
	Sideways float32 // Positive to the left of the player
	Foward   float32 // Positive forward
	Flags    uint8   // Bit mask. 0x1: jump, 0x2: unmount
}

func (packet *PacketPlaySteerVehicle) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.Sideways, err = r.ReadFloat32()
	if err != nil {
		return
	}
	packet.Foward, err = r.ReadFloat32()
	if err != nil {
		return
	}
	packet.Flags, err = r.ReadUInt8()
	if err != nil {
		return
	}
	return
}

type PacketPlayCreativeInventoryAction struct {
	Slot        uint16 // Inventory Slot
	ClickedItem slot.SlotItem
}

func (packet *PacketPlayCreativeInventoryAction) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.Slot, err = r.ReadUInt16()
	if err != nil {
		return
	}
	packet.ClickedItem, err = r.ReadSlotItem()
	if err != nil {
		return
	}
	return
}

type PacketPlayEnchantItem struct {
	WindowID    uint8 // The ID of the enchantment table window sent by Open Window
	Enchantment uint8 // The position of the enchantment on the enchantment table window, starting with 0 as the topmost one
}

func (packet *PacketPlayEnchantItem) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.WindowID, err = r.ReadUInt8()
	if err != nil {
		return
	}
	packet.Enchantment, err = r.ReadUInt8()
	if err != nil {
		return
	}
	return
}

type PacketPlayTabCompleteServer struct {
	Text          string // All text behind the cursor
	HasPosition   bool
	LookedAtBlock world.BlockPosition // The position of the block being looked at. Only sent if Has Position is true.
}

func (packet *PacketPlayTabCompleteServer) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.Text, err = r.ReadString()
	if err != nil {
		return
	}
	packet.HasPosition, err = r.ReadBool()
	if err != nil {
		return
	}
	if packet.HasPosition {
		packet.LookedAtBlock, err = r.ReadBlockPosition()
		if err != nil {
			return
		}
	}
	return
}

type PacketPlaySpectate struct {
	TargetPlayer uuid.UUID
}

func (packet *PacketPlaySpectate) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.TargetPlayer, err = r.ReadUUID()
	if err != nil {
		return
	}
	return
}

type PacketPlayResourcePackStatus struct {
	Hash   string
	Result types.ResourcePackResult // 0: successfully loaded, 1: declined, 2: failed download, 3: accepted
}

func (packet *PacketPlayResourcePackStatus) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.Hash, err = r.ReadString()
	if err != nil {
		return
	}
	res, err := r.ReadVarInt()
	if err != nil {
		return
	}
	packet.Result = types.ResourcePackResult(res)
	return
}
