package packet

import (
	"reflect"

	"github.com/git-fal7/sealantern/minecraft/types"
)

var (
	serverBoundPackets map[int16]reflect.Type = make(map[int16]reflect.Type)
	clientBoundPackets map[reflect.Type]int16 = make(map[reflect.Type]int16)
)

func packetTypeHash(state types.State, id int) int16 {
	return int16(id) + (int16(state) << 8)
}

func InitRegistry() {
	serverBoundPackets[packetTypeHash(types.HANDSHAKING, 0x00)] = reflect.TypeOf((*PacketHandshake)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.STATUS, 0x00)] = reflect.TypeOf((*PacketStatusRequest)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.STATUS, 0x01)] = reflect.TypeOf((*PacketStatusPing)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.LOGIN, 0x00)] = reflect.TypeOf((*PacketLoginStart)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x00)] = reflect.TypeOf((*PacketPlayKeepAlive)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x01)] = reflect.TypeOf((*PacketPlayChat)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x02)] = reflect.TypeOf((*PacketPlayUseEntity)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x03)] = reflect.TypeOf((*PacketPlayPlayer)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x04)] = reflect.TypeOf((*PacketPlayPlayerPosition)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x05)] = reflect.TypeOf((*PacketPlayPlayerLook)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x06)] = reflect.TypeOf((*PacketPlayPlayerPositionAndLook)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x07)] = reflect.TypeOf((*PacketPlayPlayerDigging)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x08)] = reflect.TypeOf((*PacketPlayBlockPlacement)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x09)] = reflect.TypeOf((*PacketPlayHeldItemChange)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x0A)] = reflect.TypeOf((*PacketPlaySwingArm)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x0B)] = reflect.TypeOf((*PacketPlayEntityAction)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x0C)] = reflect.TypeOf((*PacketPlaySteerVehicle)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x0D)] = reflect.TypeOf((*PacketPlayCloseWindow)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x0E)] = reflect.TypeOf((*PacketPlayClickWindow)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x0F)] = reflect.TypeOf((*PacketPlayConfirmTransaction)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x10)] = reflect.TypeOf((*PacketPlayCreativeInventoryAction)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x11)] = reflect.TypeOf((*PacketPlayEnchantItem)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x12)] = reflect.TypeOf((*PacketPlayUpdateSign)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x13)] = reflect.TypeOf((*PacketPlayPlayerAbilitiesServer)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x14)] = reflect.TypeOf((*PacketPlayTabCompleteServer)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x15)] = reflect.TypeOf((*PacketPlayClientSettings)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x16)] = reflect.TypeOf((*PacketPlayClientStatus)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x17)] = reflect.TypeOf((*PacketPlayPluginMessage)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x18)] = reflect.TypeOf((*PacketPlaySpectate)(nil)).Elem()
	serverBoundPackets[packetTypeHash(types.PLAY, 0x19)] = reflect.TypeOf((*PacketPlayResourcePackStatus)(nil)).Elem()

	// Status
	clientBoundPackets[reflect.TypeOf((*PacketStatusResponse)(nil)).Elem()] = 0x00
	clientBoundPackets[reflect.TypeOf((*PacketStatusPing)(nil)).Elem()] = 0x01

	// Login
	clientBoundPackets[reflect.TypeOf((*PacketLoginDisconnect)(nil)).Elem()] = 0x00
	clientBoundPackets[reflect.TypeOf((*PacketLoginSuccess)(nil)).Elem()] = 0x02
	clientBoundPackets[reflect.TypeOf((*PacketLoginSetCompression)(nil)).Elem()] = 0x03

	// Play
	clientBoundPackets[reflect.TypeOf((*PacketPlayKeepAlive)(nil)).Elem()] = 0x00
	clientBoundPackets[reflect.TypeOf((*PacketPlayJoinGame)(nil)).Elem()] = 0x01
	clientBoundPackets[reflect.TypeOf((*PacketPlayMessage)(nil)).Elem()] = 0x02
	clientBoundPackets[reflect.TypeOf((*PacketPlayTimeUpdate)(nil)).Elem()] = 0x03
	clientBoundPackets[reflect.TypeOf((*PacketPlayEntityEquipment)(nil)).Elem()] = 0x04
	clientBoundPackets[reflect.TypeOf((*PacketPlaySpawnPosition)(nil)).Elem()] = 0x05
	clientBoundPackets[reflect.TypeOf((*PacketPlayUpdateHealth)(nil)).Elem()] = 0x06
	clientBoundPackets[reflect.TypeOf((*PacketPlayRespawn)(nil)).Elem()] = 0x07
	clientBoundPackets[reflect.TypeOf((*PacketPlayPlayerPositionAndLookClient)(nil)).Elem()] = 0x08
	clientBoundPackets[reflect.TypeOf((*PacketPlayHeldItemChange)(nil)).Elem()] = 0x09
	clientBoundPackets[reflect.TypeOf((*PacketPlayUseBed)(nil)).Elem()] = 0x0A
	clientBoundPackets[reflect.TypeOf((*PacketPlayAnimationClient)(nil)).Elem()] = 0x0B
	clientBoundPackets[reflect.TypeOf((*PacketPlaySpawnPlayer)(nil)).Elem()] = 0x0C
	clientBoundPackets[reflect.TypeOf((*PacketPlayCollectItem)(nil)).Elem()] = 0x0D
	clientBoundPackets[reflect.TypeOf((*PacketPlaySpawnObject)(nil)).Elem()] = 0x0E
	clientBoundPackets[reflect.TypeOf((*PacketPlaySpawnMob)(nil)).Elem()] = 0x0F
	clientBoundPackets[reflect.TypeOf((*PacketPlaySpawnPainting)(nil)).Elem()] = 0x10
	clientBoundPackets[reflect.TypeOf((*PacketPlaySpawnEXPOrb)(nil)).Elem()] = 0x11
	clientBoundPackets[reflect.TypeOf((*PacketPlayEntityVelocity)(nil)).Elem()] = 0x12
	clientBoundPackets[reflect.TypeOf((*PacketPlayDestroyEntities)(nil)).Elem()] = 0x13
	clientBoundPackets[reflect.TypeOf((*PacketPlayEntity)(nil)).Elem()] = 0x14
	clientBoundPackets[reflect.TypeOf((*PacketPlayEntityRelativeMove)(nil)).Elem()] = 0x15
	clientBoundPackets[reflect.TypeOf((*PacketPlayEntityLook)(nil)).Elem()] = 0x16
	clientBoundPackets[reflect.TypeOf((*PacketPlayEntityLookAndRelativeMove)(nil)).Elem()] = 0x17
	clientBoundPackets[reflect.TypeOf((*PacketPlayEntityTeleport)(nil)).Elem()] = 0x18
	clientBoundPackets[reflect.TypeOf((*PacketPlayEntityHeadLook)(nil)).Elem()] = 0x19
	clientBoundPackets[reflect.TypeOf((*PacketPlayEntityStatus)(nil)).Elem()] = 0x1A
	clientBoundPackets[reflect.TypeOf((*PacketPlayAttachEntity)(nil)).Elem()] = 0x1B
	clientBoundPackets[reflect.TypeOf((*PacketPlayEntityMetadata)(nil)).Elem()] = 0x1C
	clientBoundPackets[reflect.TypeOf((*PacketPlayEntityEffect)(nil)).Elem()] = 0x1D
	clientBoundPackets[reflect.TypeOf((*PacketPlayRemoveEntityEffect)(nil)).Elem()] = 0x1E
	clientBoundPackets[reflect.TypeOf((*PacketPlaySetExperience)(nil)).Elem()] = 0x1F
	clientBoundPackets[reflect.TypeOf((*PacketPlayEntityProperties)(nil)).Elem()] = 0x20
	clientBoundPackets[reflect.TypeOf((*PacketPlayChunkData)(nil)).Elem()] = 0x21
	clientBoundPackets[reflect.TypeOf((*PacketPlayMultiBlockChange)(nil)).Elem()] = 0x22
	clientBoundPackets[reflect.TypeOf((*PacketPlayBlockChange)(nil)).Elem()] = 0x23
	clientBoundPackets[reflect.TypeOf((*PacketPlayBlockAction)(nil)).Elem()] = 0x24
	clientBoundPackets[reflect.TypeOf((*PacketPlayBlockBreakAnimation)(nil)).Elem()] = 0x25
	clientBoundPackets[reflect.TypeOf((*PacketPlayMapChunkBulk)(nil)).Elem()] = 0x26
	clientBoundPackets[reflect.TypeOf((*PacketPlayExplosion)(nil)).Elem()] = 0x27
	clientBoundPackets[reflect.TypeOf((*PacketPlayEffect)(nil)).Elem()] = 0x28
	clientBoundPackets[reflect.TypeOf((*PacketPlaySoundEffect)(nil)).Elem()] = 0x29
	clientBoundPackets[reflect.TypeOf((*PacketPlayParticle)(nil)).Elem()] = 0x2A
	clientBoundPackets[reflect.TypeOf((*PacketPlayChangeGameState)(nil)).Elem()] = 0x2B
	clientBoundPackets[reflect.TypeOf((*PacketPlayOpenWindow)(nil)).Elem()] = 0x2D
	clientBoundPackets[reflect.TypeOf((*PacketPlayCloseWindow)(nil)).Elem()] = 0x2E
	clientBoundPackets[reflect.TypeOf((*PacketPlaySetSlot)(nil)).Elem()] = 0x2F
	clientBoundPackets[reflect.TypeOf((*PacketPlayWindowItems)(nil)).Elem()] = 0x30
	clientBoundPackets[reflect.TypeOf((*PacketPlayWindowProperty)(nil)).Elem()] = 0x31
	clientBoundPackets[reflect.TypeOf((*PacketPlayConfirmTransaction)(nil)).Elem()] = 0x32
	clientBoundPackets[reflect.TypeOf((*PacketPlayUpdateSign)(nil)).Elem()] = 0x33
	clientBoundPackets[reflect.TypeOf((*PacketPlayMap)(nil)).Elem()] = 0x34
	clientBoundPackets[reflect.TypeOf((*PacketPlayUpdateBlockEntity)(nil)).Elem()] = 0x35
	clientBoundPackets[reflect.TypeOf((*PacketPlayOpenSignEditor)(nil)).Elem()] = 0x36
	clientBoundPackets[reflect.TypeOf((*PacketPlayStatistics)(nil)).Elem()] = 0x37
	clientBoundPackets[reflect.TypeOf((*PacketPlayPlayerListItem)(nil)).Elem()] = 0x38
	clientBoundPackets[reflect.TypeOf((*PacketPlayPlayerAbilities)(nil)).Elem()] = 0x39
	clientBoundPackets[reflect.TypeOf((*PacketPlayTabComplete)(nil)).Elem()] = 0x3A
	clientBoundPackets[reflect.TypeOf((*PacketPlayScoreboardObjective)(nil)).Elem()] = 0x3B
	clientBoundPackets[reflect.TypeOf((*PacketPlayUpdateScore)(nil)).Elem()] = 0x3C
	clientBoundPackets[reflect.TypeOf((*PacketPlayDisplayScoreboard)(nil)).Elem()] = 0x3D
	clientBoundPackets[reflect.TypeOf((*PacketPlayTeams)(nil)).Elem()] = 0x3E
	clientBoundPackets[reflect.TypeOf((*PacketPlayPluginMessage)(nil)).Elem()] = 0x3F
	clientBoundPackets[reflect.TypeOf((*PacketPlayDisconnect)(nil)).Elem()] = 0x40
	clientBoundPackets[reflect.TypeOf((*PacketPlayServerDifficulty)(nil)).Elem()] = 0x41
	clientBoundPackets[reflect.TypeOf((*PacketPlayCombatEvent)(nil)).Elem()] = 0x42
	clientBoundPackets[reflect.TypeOf((*PacketPlayCamera)(nil)).Elem()] = 0x43
	clientBoundPackets[reflect.TypeOf((*PacketPlayWorldBorder)(nil)).Elem()] = 0x44
	clientBoundPackets[reflect.TypeOf((*PacketPlayTitle)(nil)).Elem()] = 0x45
	// 0x46 Doesnt work
	clientBoundPackets[reflect.TypeOf((*PacketPlayerListHeaderFooter)(nil)).Elem()] = 0x47
	clientBoundPackets[reflect.TypeOf((*PacketPlayResourcePackSend)(nil)).Elem()] = 0x48
	clientBoundPackets[reflect.TypeOf((*PacketPlayUpdateEntityNBT)(nil)).Elem()] = 0x49
}

func GetPacketTypeFromRegistry(state types.State, id int) reflect.Type {
	return serverBoundPackets[packetTypeHash(state, id)]
}

func GetPacketIDFromClientPacket(packetType reflect.Type) int16 {
	id, ok := clientBoundPackets[packetType]
	if !ok {
		return -1
	}
	return id
}
