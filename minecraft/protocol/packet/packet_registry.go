package packet

import (
	"reflect"

	"github.com/git-fal7/sealantern/minecraft/types"
)

var (
	packets map[int64]reflect.Type = make(map[int64]reflect.Type)
)

func packetTypeHash(state types.State, id int) int64 {
	return int64(id) ^ (int64(state) << 32)
}

func InitRegistry() {
	packets[packetTypeHash(types.HANDSHAKING, 0x00)] = reflect.TypeOf((*PacketHandshake)(nil)).Elem()
	packets[packetTypeHash(types.STATUS, 0x00)] = reflect.TypeOf((*PacketStatusRequest)(nil)).Elem()
	packets[packetTypeHash(types.STATUS, 0x01)] = reflect.TypeOf((*PacketStatusPing)(nil)).Elem()
	packets[packetTypeHash(types.LOGIN, 0x00)] = reflect.TypeOf((*PacketLoginStart)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x00)] = reflect.TypeOf((*PacketPlayKeepAlive)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x01)] = reflect.TypeOf((*PacketPlayChat)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x02)] = reflect.TypeOf((*PacketPlayUseEntity)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x03)] = reflect.TypeOf((*PacketPlayPlayer)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x04)] = reflect.TypeOf((*PacketPlayPlayerPosition)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x05)] = reflect.TypeOf((*PacketPlayPlayerLook)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x06)] = reflect.TypeOf((*PacketPlayPlayerPositionAndLook)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x07)] = reflect.TypeOf((*PacketPlayPlayerDigging)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x08)] = reflect.TypeOf((*PacketPlayBlockPlacement)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x09)] = reflect.TypeOf((*PacketPlayHeldItemChange)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x0A)] = reflect.TypeOf((*PacketPlaySwingArm)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x0B)] = reflect.TypeOf((*PacketPlayEntityAction)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x0E)] = reflect.TypeOf((*PacketPlayClickWindow)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x0D)] = reflect.TypeOf((*PacketPlayCloseWindow)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x0F)] = reflect.TypeOf((*PacketPlayConfirmTransaction)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x15)] = reflect.TypeOf((*PacketPlayClientSettings)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x16)] = reflect.TypeOf((*PacketPlayClientStatus)(nil)).Elem()
	packets[packetTypeHash(types.PLAY, 0x17)] = reflect.TypeOf((*PacketPlayPluginMessage)(nil)).Elem()
}

func GetPacketTypeFromRegistry(state types.State, id int) reflect.Type {
	return packets[packetTypeHash(state, id)]
}
