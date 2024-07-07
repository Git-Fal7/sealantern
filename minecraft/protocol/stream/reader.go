package stream

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"io"
	"math"

	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/pkg/component"
	"github.com/git-fal7/sealantern/pkg/material"
	"github.com/git-fal7/sealantern/pkg/slot"
	"github.com/google/uuid"
	"github.com/seebs/nbt"
)

type ProtocolReader struct {
	*bufio.Reader
	Buffer [16]byte
}

func (r *ProtocolReader) ReadVarInt() (i int, err error) {
	v, err := binary.ReadUvarint(r)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func (r *ProtocolReader) ReadBool() (b bool, err error) {
	data, err := r.ReadByte()
	if err != nil {
		return false, err
	}
	return data == 0x01, nil
}

func (r *ProtocolReader) ReadUInt8() (i uint8, err error) {
	v, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (r *ProtocolReader) ReadUInt16() (i uint16, err error) {
	buff := r.Buffer[:2]
	_, err = io.ReadFull(r, buff)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(buff), nil
}

func (r *ProtocolReader) ReadUInt32() (i uint32, err error) {
	buff := r.Buffer[:4]
	_, err = io.ReadFull(r, buff)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(buff), nil
}

func (r *ProtocolReader) ReadUInt64() (i uint64, err error) {
	buff := r.Buffer[:8]
	_, err = io.ReadFull(r, buff)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(buff), nil
}

func (r *ProtocolReader) ReadBlockPosition() (i world.BlockPosition, err error) {
	val, err := r.ReadUInt64()
	if err != nil {
		return world.BlockPosition{}, err
	}
	pos := world.BlockPosition{}
	pos.X = int(val >> 38)
	pos.Y = int((val >> 26) & 0xFFF)
	pos.Z = int(val << 38 >> 38)
	return pos, nil
}

func (r *ProtocolReader) ReadFloat32() (i float32, err error) {
	buff := r.Buffer[:4]
	_, err = io.ReadFull(r, buff)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(binary.BigEndian.Uint32(buff)), nil
}

func (r *ProtocolReader) ReadFloat64() (i float64, err error) {
	buff := r.Buffer[:8]
	_, err = io.ReadFull(r, buff)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.BigEndian.Uint64(buff)), nil
}

func (r *ProtocolReader) ReadString() (s string, err error) {
	length, err := r.ReadVarInt()
	if err != nil {
		return "", err
	}
	buffer := make([]byte, length)
	_, err = io.ReadFull(r, buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

func (r *ProtocolReader) ReadByteArray(length int) (data []byte, err error) {
	data = make([]byte, length)
	_, err = r.Read(data)
	return data, err
}

func (r *ProtocolReader) ReadNBTCompound() (tag nbt.Compound, err error) {
	nbtTag, _, err := nbt.LoadUncompressed(r)
	tag, ok := nbtTag.(nbt.Compound)
	if !ok {
		return nil, nil
	}
	return tag, err
}

func (r *ProtocolReader) ReadSlotItem() (slotItem slot.SlotItem, err error) {
	itemID, err := r.ReadUInt16()
	if err != nil {
		return slot.SlotItem{}, err
	}
	if int16(itemID) == -1 {
		return slot.SlotItem{}, err
	}
	amount, err := r.ReadByte()
	if err != nil {
		return slot.SlotItem{}, err
	}
	durability, err := r.ReadUInt16()
	if err != nil {
		return slot.SlotItem{}, err
	}
	compoundTag, err := r.ReadNBTCompound()
	if err != nil {
		return slot.SlotItem{}, err
	}
	slotItem = slot.SlotItem{
		Material:   material.FindMaterialByID(itemID),
		Amount:     amount,
		Durability: durability,
		NBT:        compoundTag,
	}
	return slotItem, err
}

func (r *ProtocolReader) ReadUUID() (uid uuid.UUID, err error) {
	uuidArray, err := r.ReadByteArray(16)
	if err != nil {
		return uuid.Nil, err
	}
	uid, err = uuid.FromBytes(uuidArray)
	if err != nil {
		return uuid.Nil, err
	}
	return uid, err
}

func (r *ProtocolReader) ReadChatComponent() (chat *component.ChatComponent, err error) {
	componentJSON, err := r.ReadString()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(componentJSON), chat)
	if err != nil {
		return nil, err
	}
	return chat, err
}
