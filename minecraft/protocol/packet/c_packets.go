package packet

import (
	"github.com/git-fal7/sealantern/minecraft/protocol/stream"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/metadata"
	"github.com/git-fal7/sealantern/pkg/component"
	"github.com/git-fal7/sealantern/pkg/slot"
	"github.com/google/uuid"
	"github.com/seebs/nbt"
)

type PacketStatusResponse struct {
	Response string
}

func (packet *PacketStatusResponse) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteString(packet.Response)
	if err != nil {
		return
	}
	return
}

type PacketStatusPing struct {
	Time uint64
}

func (packet *PacketStatusPing) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.Time, err = r.ReadUInt64()
	if err != nil {
		return
	}
	return
}
func (packet *PacketStatusPing) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteUInt64(packet.Time)
	if err != nil {
		return
	}
	return
}

type PacketLoginDisconnect struct {
	Component component.IChatComponent
}

func (packet *PacketLoginDisconnect) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteChatComponent(packet.Component)
	if err != nil {
		return
	}
	return
}

type PacketLoginSuccess struct {
	UUID     uuid.UUID
	Username string
}

func (packet *PacketLoginSuccess) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteString(packet.UUID.String())
	if err != nil {
		return
	}
	err = w.WriteString(packet.Username)
	if err != nil {
		return
	}
	return
}

type PacketLoginSetCompression struct {
	Threshold int
}

func (packet *PacketLoginSetCompression) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(packet.Threshold)
	if err != nil {
		return
	}
	return
}

type PacketPlayTabComplete struct {
	Matches []string
}

func (packet *PacketPlayTabComplete) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(len(packet.Matches))
	if err != nil {
		return
	}
	for _, s := range packet.Matches {
		err = w.WriteString(s)
		if err != nil {
			return
		}
	}
	return
}

type PacketPlayMessage struct {
	Component component.IChatComponent
	Position  types.ChatPosition
}

func (packet *PacketPlayMessage) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteChatComponent(packet.Component)
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.Position))
	if err != nil {
		return
	}
	return
}

type PacketPlayServerDifficulty struct {
	Difficulty world.Difficulty
}

func (packet *PacketPlayServerDifficulty) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteUInt8(uint8(packet.Difficulty))
	if err != nil {
		return
	}
	return
}

type PacketPlayPluginMessage struct {
	Channel string
	Data    []byte
}

func (packet *PacketPlayPluginMessage) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.Channel, err = r.ReadString()
	if err != nil {
		return
	}
	packet.Data, err = r.ReadByteArray(length - len(packet.Channel) - 2)
	if err != nil {
		return
	}
	return
}
func (packet *PacketPlayPluginMessage) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteString(packet.Channel)
	if err != nil {
		return
	}
	err = w.WriteByteArray(packet.Data)
	if err != nil {
		return
	}
	return
}

type PacketPlayDisconnect struct {
	Component component.IChatComponent
}

func (packet *PacketPlayDisconnect) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.Component, err = r.ReadChatComponent()
	if err != nil {
		return
	}
	return
}
func (packet *PacketPlayDisconnect) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteChatComponent(packet.Component)
	if err != nil {
		return
	}
	return
}

type PacketPlayKeepAlive struct {
	Identifier int
}

func (packet *PacketPlayKeepAlive) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.Identifier, err = r.ReadVarInt()
	if err != nil {
		return
	}
	return
}
func (packet *PacketPlayKeepAlive) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(packet.Identifier)
	if err != nil {
		return
	}
	return
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

func (packet *PacketPlayParticle) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteUInt32(uint32(packet.Type))
	if err != nil {
		return
	}
	err = w.WriteBool(packet.LongDistance)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.X)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.Y)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.Z)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.OffsetX)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.OffsetY)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.OffsetZ)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.ParticleData)
	if err != nil {
		return
	}
	err = w.WriteUInt32(uint32(packet.Count))
	if err != nil {
		return
	}
	return
}

type PacketPlayJoinGame struct {
	Gamemode     types.Gamemode
	Dimension    world.Dimension
	Difficulty   world.Difficulty
	MaxPlayers   uint8
	LevelType    world.LevelType
	ReducedDebug bool
}

func (packet *PacketPlayJoinGame) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteInt32(0) // entity ids are more than 255,
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.Gamemode))
	if err != nil {
		return
	}
	err = w.WriteInt8(int8(packet.Dimension))
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.Difficulty))
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.MaxPlayers)
	if err != nil {
		return
	}
	err = w.WriteString(string(packet.LevelType))
	if err != nil {
		return
	}
	err = w.WriteBool(packet.ReducedDebug)
	if err != nil {
		return
	}
	return
}

type PacketPlayPlayerAbilities struct {
	Invulnerable bool
	Fly          bool
	CanFly       bool
	Creative     bool
	FlyingSpeed  float32
	FieldOfView  float32
}

func (packet *PacketPlayPlayerAbilities) Write(w *stream.ProtocolWriter) (err error) {
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
		return
	}
	err = w.WriteFloat32(packet.FlyingSpeed)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.FieldOfView)
	if err != nil {
		return
	}
	return
}

type PacketPlayPlayerPositionAndLookClient struct {
	Position world.Position
	Flags    uint8
}

func (packet *PacketPlayPlayerPositionAndLookClient) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteFloat64(packet.Position.X)
	if err != nil {
		return
	}
	err = w.WriteFloat64(packet.Position.Y)
	if err != nil {
		return
	}
	err = w.WriteFloat64(packet.Position.Z)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.Position.Yaw)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.Position.Pitch)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Flags)
	if err != nil {
		return
	}
	return
}

type PacketPlayUpdateHealth struct {
	Health         float32
	Food           int
	FoodSaturation float32
}

func (packet *PacketPlayUpdateHealth) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteFloat32(packet.Health)
	if err != nil {
		return
	}
	err = w.WriteVarInt(packet.Food)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.FoodSaturation)
	if err != nil {
		return
	}
	return
}

type PacketPlaySpawnPosition struct {
	Position world.Position
}

func (packet *PacketPlaySpawnPosition) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteBlockPosition(packet.Position.ToBlockPosition())
	if err != nil {
		return
	}
	return
}

type PacketPlayerListHeaderFooter struct {
	Header component.IChatComponent
	Footer component.IChatComponent
}

func (packet *PacketPlayerListHeaderFooter) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteChatComponent(packet.Header)
	if err != nil {
		return
	}
	err = w.WriteChatComponent(packet.Footer)
	if err != nil {
		return
	}
	return
}

type PacketPlayPlayerListItem struct {
	Action  types.PlayerListAction
	Entries []types.PlayerListEntry
}

func (packet *PacketPlayPlayerListItem) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.Action))
	if err != nil {
		return
	}
	err = w.WriteVarInt(len(packet.Entries))
	if err != nil {
		return
	}

	for _, entry := range packet.Entries {
		err = w.WriteUUID(entry.Profile.UUID)
		if err != nil {
			return
		}
		switch packet.Action {
		case types.PlayerListActionAddPlayer:
			{
				err = w.WriteString(entry.Profile.Name)
				if err != nil {
					return
				}
				err = w.WriteVarInt(len(entry.Profile.Properties))
				if err != nil {
					return
				}
				for _, property := range entry.Profile.Properties {
					err = w.WriteString(property.Name)
					if err != nil {
						return
					}
					err = w.WriteString(property.Value)
					if err != nil {
						return
					}
					if property.Signature != "" {
						err = w.WriteBool(true)
						if err != nil {
							return
						}
						err = w.WriteString(property.Signature)
						if err != nil {
							return
						}
					} else {
						err = w.WriteBool(false)
						if err != nil {
							return
						}
					}
				}
				err = w.WriteVarInt(int(entry.GameMode))
				if err != nil {
					return
				}
				err = w.WriteVarInt(entry.Ping)
				if err != nil {
					return
				}
				if entry.DisplayName != nil {
					err = w.WriteBool(true)
					if err != nil {
						return
					}
					err = w.WriteChatComponent(entry.DisplayName)
					if err != nil {
						return
					}

				} else {
					err = w.WriteBool(false)
					if err != nil {
						return
					}
				}

			}
		case types.PlayerListActionUpdateGamemode:
			{
				err = w.WriteVarInt(int(entry.GameMode))
				if err != nil {
					return
				}
			}
		case types.PlayerListActionUpdateLatency:
			{
				err = w.WriteVarInt(entry.Ping)
				if err != nil {
					return
				}
			}
		case types.PlayerListActionUpdateDisplayName:
			{
				if entry.DisplayName != nil {
					err = w.WriteBool(true)
					if err != nil {
						return
					}
					err = w.WriteChatComponent(entry.DisplayName)
					if err != nil {
						return
					}
				} else {
					err = w.WriteBool(false)
					if err != nil {
						return
					}
				}
			}
		}
		// Remove player doesn't have any fields. so we dont even need to put it in the switch case
	}
	return
}

type PacketPlaySpawnPlayer struct {
	EntityID       uint16
	PlayerUUID     uuid.UUID
	PlayerPosition world.Position
	CurrentItem    uint16
	Metadata       metadata.MetadataMap
}

func (packet *PacketPlaySpawnPlayer) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteUUID(packet.PlayerUUID)
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.PlayerPosition.IntX())
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.PlayerPosition.IntY())
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.PlayerPosition.IntZ())
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.PlayerPosition.IntYaw())
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.PlayerPosition.IntPitch())
	if err != nil {
		return
	}
	err = w.WriteUInt16(packet.CurrentItem)
	if err != nil {
		return
	}
	err = w.WriteMetadata(packet.Metadata)
	if err != nil {
		return
	}
	return
}

type PacketPlayDestroyEntities struct {
	EntityIDs []uint16
}

func (packet *PacketPlayDestroyEntities) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(len(packet.EntityIDs))
	if err != nil {
		return
	}
	for _, enitiyId := range packet.EntityIDs {
		err = w.WriteVarInt(int(enitiyId))
		if err != nil {
			return
		}
	}
	return
}

type PacketPlayEntityRelativeMove struct {
	EntityID uint16
	DeltaX   int8
	DeltaY   int8
	DeltaZ   int8
	OnGround bool
}

func (packet *PacketPlayEntityRelativeMove) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.DeltaX))
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.DeltaY))
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.DeltaZ))
	if err != nil {
		return
	}
	err = w.WriteBool(packet.OnGround)
	if err != nil {
		return
	}
	return
}

type PacketPlayEntityLook struct {
	EntityID uint16
	Yaw      uint8
	Pitch    uint8
	OnGround bool
}

func (packet *PacketPlayEntityLook) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Yaw)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Pitch)
	if err != nil {
		return
	}
	err = w.WriteBool(packet.OnGround)
	if err != nil {
		return
	}
	return
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

func (packet *PacketPlayEntityLookAndRelativeMove) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.DeltaX))
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.DeltaY))
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.DeltaZ))
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Yaw)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Pitch)
	if err != nil {
		return
	}
	err = w.WriteBool(packet.OnGround)
	if err != nil {
		return
	}
	return
}

type PacketPlayEntityHeadLook struct {
	EntityID uint16
	HeadYaw  uint8
}

func (packet *PacketPlayEntityHeadLook) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.HeadYaw)
	if err != nil {
		return
	}
	return
}

type PacketPlayEntityMetadata struct {
	EntityID uint16
	Metadata metadata.MetadataMap
}

func (packet *PacketPlayEntityMetadata) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteMetadata(packet.Metadata)
	if err != nil {
		return
	}
	return
}

type PacketPlayAnimationClient struct {
	EntityID  uint16
	Animation types.Animation
}

func (packet *PacketPlayAnimationClient) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.Animation))
	if err != nil {
		return
	}
	return
}

type PacketPlayEntityVelocity struct {
	EntityID uint16
	Velocity world.Vector
}

func (packet *PacketPlayEntityVelocity) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteInt16(int16(packet.Velocity.X * 8000))
	if err != nil {
		return
	}
	err = w.WriteInt16(int16(packet.Velocity.Y * 8000))
	if err != nil {
		return
	}
	err = w.WriteInt16(int16(packet.Velocity.Z * 8000))
	if err != nil {
		return
	}
	return
}

type PacketPlayMapChunkBulk struct {
	Packets  []PacketPlayChunkData
	Skylight bool
}

func (packet *PacketPlayMapChunkBulk) Write(w *stream.ProtocolWriter) (err error) {
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

type PacketPlayChunkData struct {
	X              int32
	Z              int32
	GroundUp       bool
	SectionBitMask uint16
	Data           []byte
}

func (packet *PacketPlayChunkData) Write(w *stream.ProtocolWriter) (err error) {
	w.WriteInt32(packet.X)
	w.WriteInt32(packet.Z)
	w.WriteBool(packet.GroundUp)
	w.WriteUInt16(packet.SectionBitMask)
	w.WriteVarInt(len(packet.Data))
	w.WriteByteArray(packet.Data)
	return
}

type PacketPlayScoreboardObjective struct {
	ObjectiveName string
	DisplayName   string
	Mode          types.ScoreboardMode
	RenderType    types.ObjectiveRenderType
}

func (packet *PacketPlayScoreboardObjective) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteString(packet.ObjectiveName)
	if err != nil {
		return
	}
	err = w.WriteInt8(int8(packet.Mode))
	if err != nil {
		return
	}
	if packet.Mode == types.CreateScoreboard || packet.Mode == types.UpdateScoreboard {
		err = w.WriteString(packet.DisplayName)
		if err != nil {
			return
		}
		err = w.WriteString(string(packet.RenderType))
		if err != nil {
			return
		}
	}
	return
}

type PacketPlayUpdateScore struct {
	ScoreName     string
	Action        types.UpdateScoreAction
	ObjectiveName string
	Value         int
}

func (packet *PacketPlayUpdateScore) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteString(packet.ScoreName)
	if err != nil {
		return
	}
	err = w.WriteInt8(int8(packet.Action))
	if err != nil {
		return
	}

	err = w.WriteString(packet.ObjectiveName)
	if err != nil {
		return
	}
	if packet.Action != types.RemoveScoreItem {
		err = w.WriteVarInt(packet.Value)
		if err != nil {
			return
		}
	}
	return
}

type PacketPlayDisplayScoreboard struct {
	Position  types.ObjectiveDisplaySlot
	ScoreName string
}

func (packet *PacketPlayDisplayScoreboard) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteInt8(int8(packet.Position))
	if err != nil {
		return
	}
	err = w.WriteString(packet.ScoreName)
	if err != nil {
		return
	}
	return
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

func (packet *PacketPlayTeams) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteString(packet.TeamName)
	if err != nil {
		return
	}

	err = w.WriteInt8(int8(packet.Mode))
	if err != nil {
		return
	}
	if packet.Mode == types.TeamModeCreate || packet.Mode == types.TeamModeUpdate {
		err = w.WriteString(packet.TeamDisplayName)
		if err != nil {
			return
		}
		err = w.WriteString(packet.TeamPrefix)
		if err != nil {
			return
		}
		err = w.WriteString(packet.TeamSuffix)
		if err != nil {
			return
		}
		err = w.WriteInt8(int8(packet.FriendlyFire))
		if err != nil {
			return
		}
		err = w.WriteString(string(packet.NameTagVisibility))
		if err != nil {
			return
		}
		err = w.WriteInt8(int8(packet.Color))
		if err != nil {
			return
		}
	}
	if packet.Mode == types.TeamModeCreate || packet.Mode == types.TeamModeAddPlayer || packet.Mode == types.TeamModeRemovePlayer {
		if packet.Players != nil {
			err = w.WriteVarInt(len(packet.Players))
			if err != nil {
				return
			}
			for _, player := range packet.Players {
				err = w.WriteString(player)
				if err != nil {
					return
				}
			}
		}
	}
	return
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

func (packet *PacketPlaySpawnMob) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.MobType))
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Position.IntX())
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Position.IntY())
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Position.IntZ())
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Position.IntYaw())
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Position.IntPitch())
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.HeadPitch)
	if err != nil {
		return
	}
	err = w.WriteUInt16(packet.VelocityX)
	if err != nil {
		return
	}
	err = w.WriteUInt16(packet.VelocityY)
	if err != nil {
		return
	}
	err = w.WriteUInt16(packet.VelocityZ)
	if err != nil {
		return
	}
	err = w.WriteMetadata(packet.Metadata)
	if err != nil {
		return
	}
	return
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

func (packet *PacketPlaySpawnObject) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.ObjectType))
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Position.IntX())
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Position.IntY())
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Position.IntZ())
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Position.IntPitch())
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Position.IntYaw())
	if err != nil {
		return
	}
	err = w.WriteInt32(int32(packet.Data))
	if err != nil {
		return
	}
	if packet.Data != 0 {
		err = w.WriteUInt16(packet.VelocityX)
		if err != nil {
			return
		}
		err = w.WriteUInt16(packet.VelocityY)
		if err != nil {
			return
		}
		err = w.WriteUInt16(packet.VelocityZ)
		if err != nil {
			return
		}
	}
	return
}

type PacketPlayOpenWindow struct {
	WindowID      uint8
	WindowType    types.WindowType
	WindowTitle   string
	NumberOfSlots uint8
	HorseEntityID uint16
}

func (packet *PacketPlayOpenWindow) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteUInt8(packet.WindowID)
	if err != nil {
		return
	}
	err = w.WriteString(string(packet.WindowType))
	if err != nil {
		return
	}
	err = w.WriteString(packet.WindowTitle)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.NumberOfSlots)
	if err != nil {
		return
	}
	if packet.WindowType == types.WindowTypeEntityHorse {
		err = w.WriteVarInt(int(packet.HorseEntityID))
		if err != nil {
			return
		}
	}
	return
}

type PacketPlayWindowItems struct {
	WindowID uint8
	SlotData []slot.SlotItem
}

func (packet *PacketPlayWindowItems) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteUInt8(packet.WindowID)
	if err != nil {
		return
	}
	err = w.WriteUInt16(uint16(len(packet.SlotData)))
	if err != nil {
		return
	}
	for _, slot := range packet.SlotData {
		err = w.WriteSlotItem(slot)
		if err != nil {
			return
		}
	}
	return
}

type PacketPlayConfirmTransaction struct {
	WindowID     uint8
	ActionNumber uint16
	Accepted     bool
}

func (packet *PacketPlayConfirmTransaction) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.WindowID, err = r.ReadUInt8()
	if err != nil {
		return
	}
	packet.ActionNumber, err = r.ReadUInt16()
	if err != nil {
		return
	}
	packet.Accepted, err = r.ReadBool()
	if err != nil {
		return
	}
	return
}

func (packet *PacketPlayConfirmTransaction) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteUInt8(packet.WindowID)
	if err != nil {
		return
	}
	err = w.WriteUInt16(packet.ActionNumber)
	if err != nil {
		return
	}
	err = w.WriteBool(packet.Accepted)
	if err != nil {
		return
	}
	return
}

type PacketPlaySetSlot struct {
	WindowID uint8
	Slot     int16
	SlotData slot.SlotItem
}

func (packet *PacketPlaySetSlot) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteUInt8(packet.WindowID)
	if err != nil {
		return
	}
	err = w.WriteInt16(packet.Slot)
	if err != nil {
		return
	}
	err = w.WriteSlotItem(packet.SlotData)
	if err != nil {
		return
	}
	return
}

type PacketPlayCloseWindow struct {
	WindowID uint8
}

func (packet *PacketPlayCloseWindow) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.WindowID, err = r.ReadUInt8()
	if err != nil {
		return
	}
	return
}

func (packet *PacketPlayCloseWindow) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteUInt8(packet.WindowID)
	if err != nil {
		return
	}
	return
}

type PacketPlayBlockChange struct {
	Location world.BlockPosition
	Type     int
}

func (packet *PacketPlayBlockChange) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteBlockPosition(packet.Location)
	if err != nil {
		return
	}
	err = w.WriteVarInt(packet.Type)
	if err != nil {
		return
	}
	return
}

type PacketPlayHeldItemChange struct {
	Slot uint8
}

func (packet *PacketPlayHeldItemChange) Read(r *stream.ProtocolReader, length int) (err error) {
	slot, err := r.ReadUInt16()
	if err != nil {
		return
	}
	packet.Slot = uint8(slot)
	return
}

func (packet *PacketPlayHeldItemChange) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteUInt8(packet.Slot)
	if err != nil {
		return
	}
	return
}

type PacketPlayEntityEquipment struct {
	EntityID uint16
	Slot     types.EquipmentSlot
	Item     slot.SlotItem
}

func (packet *PacketPlayEntityEquipment) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteUInt16(uint16(packet.Slot))
	if err != nil {
		return
	}
	err = w.WriteSlotItem(packet.Item)
	if err != nil {
		return
	}
	return
}

type PacketPlayRespawn struct {
	Dimension  world.Dimension
	Difficulty world.Difficulty
	Gamemode   types.Gamemode
	LevelType  world.LevelType
}

func (packet *PacketPlayRespawn) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteInt32(int32(packet.Dimension))
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.Gamemode))
	if err != nil {
		return
	}
	err = w.WriteUInt8(uint8(packet.Difficulty))
	if err != nil {
		return
	}
	err = w.WriteString(string(packet.LevelType))
	if err != nil {
		return
	}
	return
}

type PacketPlayChangeGameState struct {
	Reason types.GameStateReason
	Value  float32
}

func (packet *PacketPlayChangeGameState) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteUInt8(uint8(packet.Reason))
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.Value)
	if err != nil {
		return
	}
	return
}

type PacketPlayTimeUpdate struct {
	WorldAge  float32
	TimeOfDay float32
}

func (packet *PacketPlayTimeUpdate) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteFloat32(packet.WorldAge)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.TimeOfDay)
	if err != nil {
		return
	}
	return
}

type PacketPlayUseBed struct {
	EntityID uint16
	Location world.BlockPosition
}

func (packet *PacketPlayUseBed) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteBlockPosition(packet.Location)
	if err != nil {
		return
	}
	return
}

type PacketPlayCollectItem struct {
	CollectedItemEntityID int
	CollectorEntityID     uint16
}

func (packet *PacketPlayCollectItem) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(packet.CollectedItemEntityID)
	if err != nil {
		return
	}
	err = w.WriteVarInt(int(packet.CollectorEntityID))
	if err != nil {
		return
	}
	return
}

type PacketPlaySpawnPainting struct {
	EntityID  int
	Title     string
	Location  world.BlockPosition
	Direction uint8 // 0: north (-z), 1: west (-x), 2: south (+z), 3: east (+x)
}

func (packet *PacketPlaySpawnPainting) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(packet.EntityID)
	if err != nil {
		return
	}
	err = w.WriteString(packet.Title)
	if err != nil {
		return
	}
	err = w.WriteBlockPosition(packet.Location)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Direction)
	if err != nil {
		return
	}
	return
}

type PacketPlaySpawnEXPOrb struct {
	EntityID int
	Position world.Position
	Count    uint16
}

func (packet *PacketPlaySpawnEXPOrb) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(packet.EntityID)
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Position.IntX())
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Position.IntY())
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Position.IntZ())
	if err != nil {
		return
	}
	err = w.WriteUInt16(packet.Count)
	if err != nil {
		return
	}
	return
}

type PacketPlayEntity struct {
	EntityID int
}

func (packet *PacketPlayEntity) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(packet.EntityID)
	if err != nil {
		return
	}
	return
}

type PacketPlayEntityTeleport struct {
	EntityID uint16
	Position world.Position
	OnGround bool
}

func (packet *PacketPlayEntityTeleport) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Position.IntX())
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Position.IntY())
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Position.IntZ())
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Position.IntYaw())
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Position.IntPitch())
	if err != nil {
		return
	}
	err = w.WriteBool(packet.OnGround)
	if err != nil {
		return
	}
	return
}

type PacketPlayEntityStatus struct {
	EntityID     uint16
	EntityStatus uint8
}

func (packet *PacketPlayEntityStatus) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.EntityStatus)
	if err != nil {
		return
	}
	return
}

type PacketPlayAttachEntity struct {
	EntityID  uint16
	VehicleID int // -1 to detach
	Leash     bool
}

func (packet *PacketPlayAttachEntity) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteVarInt(packet.VehicleID)
	if err != nil {
		return
	}
	err = w.WriteBool(packet.Leash)
	if err != nil {
		return
	}
	return
}

type PacketPlayEntityEffect struct {
	EntityID      int
	EffectID      uint8 // See https://minecraft.wiki/w/Status_effect%23List_of_effects
	Amplifier     uint8 // Notchian client displays effect level as Amplifier + 1
	Duration      int   // Seconds
	HideParticles bool
}

func (packet *PacketPlayEntityEffect) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(packet.EntityID)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.EffectID)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Amplifier)
	if err != nil {
		return
	}
	err = w.WriteVarInt(packet.Duration)
	if err != nil {
		return
	}
	err = w.WriteBool(packet.HideParticles)
	if err != nil {
		return
	}
	return
}

type PacketPlayRemoveEntityEffect struct {
	EntityID int
	EffectID uint8 // See https://minecraft.wiki/w/Status_effect%23List_of_effects
}

func (packet *PacketPlayRemoveEntityEffect) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(packet.EntityID)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.EffectID)
	if err != nil {
		return
	}
	return
}

type PacketPlaySetExperience struct {
	ExperienceBar   float32 // Between 0 and 1
	Level           int
	TotalExperience int // See https://minecraft.wiki/w/Experience%23Leveling_up
}

func (packet *PacketPlaySetExperience) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteFloat32(packet.ExperienceBar)
	if err != nil {
		return
	}
	err = w.WriteVarInt(packet.Level)
	if err != nil {
		return
	}
	err = w.WriteVarInt(packet.TotalExperience)
	if err != nil {
		return
	}
	return
}

// TODO: Attributes
type PacketPlayEntityProperties struct {
	EntityID uint16
}

func (packet *PacketPlayEntityProperties) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteVarInt(0) // TODO: Properties
	if err != nil {
		return
	}
	return
}

type PacketPlayMultiBlockChange struct {
	ChunkX int32
	ChunkZ int32
	// TODO: Records
}

func (packet *PacketPlayMultiBlockChange) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteInt32(packet.ChunkX)
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.ChunkZ)
	if err != nil {
		return
	}
	err = w.WriteVarInt(0)
	if err != nil {
		return
	}
	return
}

type PacketPlayBlockAction struct {
	Location  world.BlockPosition
	Data1     uint8 // See https://wiki.vg/Block_Actions
	Data2     uint8 // See https://wiki.vg/Block_Actions
	BlockType int
}

func (packet *PacketPlayBlockAction) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteBlockPosition(packet.Location)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Data1)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Data2)
	if err != nil {
		return
	}
	err = w.WriteVarInt(packet.BlockType)
	if err != nil {
		return
	}
	return
}

type PacketPlayBlockBreakAnimation struct {
	EntityID     int
	Location     world.BlockPosition
	DestroyStage uint8 // 0–9 to set it, any other value to remove it
}

func (packet *PacketPlayBlockBreakAnimation) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(packet.EntityID)
	if err != nil {
		return
	}
	err = w.WriteBlockPosition(packet.Location)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.DestroyStage)
	if err != nil {
		return
	}
	return
}

type PacketPlayExplosion struct {
	X      float32
	Y      float32
	Z      float32
	Radius float32
	// TODO Record
	PlayerMotionX float32
	PlayerMotionY float32
	PlayerMotionZ float32
}

func (packet *PacketPlayExplosion) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteFloat32(packet.X)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.Y)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.Z)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.Radius)
	if err != nil {
		return
	}
	err = w.WriteVarInt(0)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.PlayerMotionX)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.PlayerMotionY)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.PlayerMotionZ)
	if err != nil {
		return
	}
	return
}

type PacketPlayEffect struct {
	EffectID              int32
	Location              world.BlockPosition
	Data                  int32
	DisableRelativeVolume bool
}

func (packet *PacketPlayEffect) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteInt32(packet.EffectID)
	if err != nil {
		return
	}
	err = w.WriteBlockPosition(packet.Location)
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.Data)
	if err != nil {
		return
	}
	err = w.WriteBool(packet.DisableRelativeVolume)
	if err != nil {
		return
	}
	return
}

type PacketPlaySoundEffect struct {
	SoundName       string  // https://github.com/SirCmpwn/Craft.Net/blob/master/source/Craft.Net.Common/SoundEffect.cs
	EffectPositionX int32   // Effect X multiplied by 8
	EffectPositionY int32   // Effect Y multiplied by 8
	EffectPositionZ int32   // Effect Z multiplied by 8
	Volume          float32 // 1 is 100%, can be more
	Pitch           uint8   // 63 is 100%, can be more
}

func (packet *PacketPlaySoundEffect) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteString(packet.SoundName)
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.EffectPositionX)
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.EffectPositionY)
	if err != nil {
		return
	}
	err = w.WriteInt32(packet.EffectPositionZ)
	if err != nil {
		return
	}
	err = w.WriteFloat32(packet.Volume)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Pitch)
	if err != nil {
		return
	}
	return
}

type PacketPlayWindowProperty struct {
	WindowID uint8
	Property uint16
	Value    uint16
}

func (packet *PacketPlayWindowProperty) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteUInt8(packet.WindowID)
	if err != nil {
		return
	}
	err = w.WriteUInt16(packet.Property)
	if err != nil {
		return
	}
	err = w.WriteUInt16(packet.Value)
	if err != nil {
		return
	}
	return
}

type PacketPlayUpdateSign struct {
	Location   world.BlockPosition
	Components [4]component.IChatComponent
}

func (packet *PacketPlayUpdateSign) Read(r *stream.ProtocolReader, length int) (err error) {
	packet.Location, err = r.ReadBlockPosition()
	if err != nil {
		return
	}
	for i := 0; i < 4; i++ {
		packet.Components[i], err = r.ReadChatComponent()
		if err != nil {
			return
		}
	}
	return
}

func (packet *PacketPlayUpdateSign) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteBlockPosition(packet.Location)
	if err != nil {
		return
	}
	for _, component := range packet.Components {
		err = w.WriteChatComponent(component)
		if err != nil {
			return
		}
	}
	return
}

type PacketPlayMap struct {
	// TODO:
}

func (packet *PacketPlayMap) Write(w *stream.ProtocolWriter) (err error) {
	return
}

type PacketPlayUpdateBlockEntity struct {
	Location world.BlockPosition
	Action   uint8
	NBTData  nbt.Compound
}

func (packet *PacketPlayUpdateBlockEntity) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteBlockPosition(packet.Location)
	if err != nil {
		return
	}
	err = w.WriteUInt8(packet.Action)
	if err != nil {
		return
	}
	err = w.WriteNBTCompound(packet.NBTData)
	if err != nil {
		return
	}
	return
}

type PacketPlayOpenSignEditor struct {
	Location world.BlockPosition
}

func (packet *PacketPlayOpenSignEditor) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteBlockPosition(packet.Location)
	if err != nil {
		return
	}
	return
}

type PacketPlayStatistics struct {
	// TODO
}

func (packet *PacketPlayStatistics) Write(w *stream.ProtocolWriter) (err error) {
	return
}

type PacketPlayCombatEvent struct {
	Event    int    // 0: enter combat, 1: end combat, 2: entity dead
	Duration int    // Only for end combat
	PlayerID int    // Only for entity dead
	EntityID int32  // Only for end combat and entity dead
	Message  string // Only for entity dead
}

func (packet *PacketPlayCombatEvent) Write(w *stream.ProtocolWriter) (err error) {
	// TODO
	err = w.WriteVarInt(packet.Event)
	if err != nil {
		return
	}
	return
}

type PacketPlayCamera struct {
	ClientID int
}

func (packet *PacketPlayCamera) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(packet.ClientID)
	if err != nil {
		return
	}
	return
}

type PacketPlayWorldBorder struct {
	Action int
}

func (packet *PacketPlayWorldBorder) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(packet.Action)
	if err != nil {
		return
	}
	return
}

type PacketPlayTitle struct {
	Action   types.TitleAction
	Title    component.IChatComponent
	Subtitle component.IChatComponent
	FadeIn   int32 // ticks
	Stay     int32 // ticks
	FadeOut  int32 // ticks
}

func (packet *PacketPlayTitle) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.Action))
	if err != nil {
		return
	}
	switch packet.Action {
	case types.TitleActionSetTitle:
		{
			err = w.WriteChatComponent(packet.Title)
			if err != nil {
				return
			}
		}
	case types.TitleActionSetSubtitle:
		{
			err = w.WriteChatComponent(packet.Subtitle)
			if err != nil {
				return
			}
		}
	case types.TitleActionSetTimesAndDisplay:
		{
			err = w.WriteInt32(packet.FadeIn)
			if err != nil {
				return
			}
			err = w.WriteInt32(packet.Stay)
			if err != nil {
				return
			}
			err = w.WriteInt32(packet.FadeOut)
			if err != nil {
				return
			}
		}
	}
	return
}

type PacketPlayResourcePackSend struct {
	URL  string
	Hash string
}

func (packet *PacketPlayResourcePackSend) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteString(packet.URL)
	if err != nil {
		return
	}
	err = w.WriteString(packet.Hash)
	if err != nil {
		return
	}
	return
}

type PacketPlayUpdateEntityNBT struct {
	EntityID uint16
	Tag      nbt.Compound
}

func (packet *PacketPlayUpdateEntityNBT) Write(w *stream.ProtocolWriter) (err error) {
	err = w.WriteVarInt(int(packet.EntityID))
	if err != nil {
		return
	}
	err = w.WriteNBTCompound(packet.Tag)
	if err != nil {
		return
	}
	return
}
