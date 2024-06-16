package readerwriter

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/pkg/slot"
	"github.com/seebs/nbt"
)

func (r *ConnReadWrite) ReadByte() (b byte, err error) {
	buff := r.Buffer[:1]
	if _, err = r.Rdr.Read(buff); err != nil {
		return 0, err
	}
	return buff[0], nil
}

func (r *ConnReadWrite) ReadVarInt() (i int, err error) {
	v, err := binary.ReadUvarint(r)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func (r *ConnReadWrite) ReadBool() (b bool, err error) {
	buff := r.Buffer[:1]
	_, err = io.ReadFull(r.Rdr, buff)
	if err != nil {
		return false, err
	}
	return buff[0] == 0x01, nil
}

func (r *ConnReadWrite) ReadUInt8() (i uint8, err error) {
	buff := r.Buffer[:1]
	_, err = io.ReadFull(r.Rdr, buff)
	if err != nil {
		return 0, err
	}
	return buff[0], nil
}

func (r *ConnReadWrite) ReadUInt16() (i uint16, err error) {
	buff := r.Buffer[:2]
	_, err = io.ReadFull(r.Rdr, buff)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(buff), nil
}

func (r *ConnReadWrite) ReadUInt32() (i uint32, err error) {
	buff := r.Buffer[:4]
	_, err = io.ReadFull(r.Rdr, buff)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(buff), nil
}

func (r *ConnReadWrite) ReadUInt64() (i uint64, err error) {
	buff := r.Buffer[:8]
	_, err = io.ReadFull(r.Rdr, buff)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(buff), nil
}

func (r *ConnReadWrite) ReadBlockPosition() (i world.BlockPosition, err error) {
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

func (r *ConnReadWrite) ReadFloat32() (i float32, err error) {
	buff := r.Buffer[:4]
	_, err = io.ReadFull(r.Rdr, buff)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(binary.BigEndian.Uint32(buff)), nil
}

func (r *ConnReadWrite) ReadFloat64() (i float64, err error) {
	buff := r.Buffer[:8]
	_, err = io.ReadFull(r.Rdr, buff)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.BigEndian.Uint64(buff)), nil
}

func (r *ConnReadWrite) ReadString() (s string, err error) {
	length, err := r.ReadVarInt()
	if err != nil {
		return "", err
	}
	buffer := make([]byte, length)
	_, err = io.ReadFull(r.Rdr, buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

func (r *ConnReadWrite) ReadStringLimited(max int) (s string, err error) {
	max = (max * 4) + 3

	length, err := r.ReadVarInt()
	if err != nil {
		return "", err
	}
	if length > max {
		return "", fmt.Errorf("invalid packet")
	}
	buffer := make([]byte, length)
	_, err = io.ReadFull(r.Rdr, buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

func (r *ConnReadWrite) ReadNStringLimited(max int) (s string, read int, err error) {
	max = (max * 4) + 3

	length, err := r.ReadVarInt()
	buff := make([]byte, 8)
	read = binary.PutUvarint(buff, uint64(length))
	if err != nil {
		return "", read, err
	}
	if length > max {
		return "", read, fmt.Errorf("invalid packet")
	}
	buffer := make([]byte, length)
	_, err = io.ReadFull(r.Rdr, buffer)
	if err != nil {
		return "", read + length, err
	}
	return string(buffer), read + length, nil
}

func (r *ConnReadWrite) ReadByteArray(length int) (data []byte, err error) {
	data = make([]byte, length)
	_, err = r.Rdr.Read(data)
	return data, err
}

func (r *ConnReadWrite) ReadNBTCompound() (tag nbt.Compound, err error) {
	nbtTag, _, err := nbt.LoadUncompressed(r.Rdr)
	tag, ok := nbtTag.(nbt.Compound)
	if !ok {
		return nil, nil
	}
	return tag, err
}

func (r *ConnReadWrite) ReadSlotItem() (slotItem slot.SlotItem, err error) {
	slotType, err := r.ReadUInt16()
	if err != nil {
		return slot.SlotItem{}, err
	}
	if int16(slotType) == -1 {
		return slot.SlotItem{}, err
	}
	amount, err := r.ReadUInt8()
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
		ID:         slotType,
		Amount:     amount,
		Durability: durability,
		NBT:        compoundTag,
	}
	return slotItem, err
}
