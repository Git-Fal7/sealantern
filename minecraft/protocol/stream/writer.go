package stream

import (
	"bytes"
	"encoding/binary"

	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/metadata"
	"github.com/git-fal7/sealantern/pkg/slot"
	"github.com/google/uuid"
	"github.com/seebs/nbt"
)

type ProtocolWriter struct {
	bytes.Buffer
}

func (w *ProtocolWriter) WriteVarInt(i int) (err error) {
	varIntBuf := make([]byte, binary.MaxVarintLen64)
	encodedLen := binary.PutUvarint(varIntBuf, uint64(i))
	_, err = w.Write(varIntBuf[:encodedLen])
	return
}

func (w *ProtocolWriter) WriteBool(b bool) (err error) {
	if b {
		err = w.WriteByte(0x01)
	} else {
		err = w.WriteByte(0x00)
	}
	return
}

func (w *ProtocolWriter) WriteUInt8(b uint8) (err error) {
	err = w.WriteByte(b)
	return
}

func (w *ProtocolWriter) WriteUInt16(i uint16) (err error) {
	err = binary.Write(&w.Buffer, binary.BigEndian, i)
	return
}

func (w *ProtocolWriter) WriteLittleEndianUInt16(i uint16) (err error) {
	err = binary.Write(&w.Buffer, binary.LittleEndian, i)
	return
}

func (w *ProtocolWriter) WriteUInt32(i uint32) (err error) {
	err = binary.Write(&w.Buffer, binary.BigEndian, i)
	return
}

func (w *ProtocolWriter) WriteUInt64(i uint64) (err error) {
	err = binary.Write(&w.Buffer, binary.BigEndian, i)
	return
}

func (w *ProtocolWriter) WriteInt8(i int8) (err error) {
	err = w.WriteByte(byte(i))
	return
}

func (w *ProtocolWriter) WriteInt16(i int16) (err error) {
	err = w.WriteUInt16(uint16(i))
	return
}

func (w *ProtocolWriter) WriteInt32(i int32) (err error) {
	err = w.WriteUInt32(uint32(i))
	return
}

func (w *ProtocolWriter) WriteFloat32(i float32) (err error) {
	err = binary.Write(&w.Buffer, binary.BigEndian, i)
	return
}

func (w *ProtocolWriter) WriteFloat64(i float64) (err error) {
	err = binary.Write(&w.Buffer, binary.BigEndian, i)
	return
}

func (w *ProtocolWriter) WriteByteArray(data []byte) (err error) {
	_, err = w.Write(data)
	return
}

func (w *ProtocolWriter) WriteString(s string) (err error) {
	err = w.WriteVarInt(len(s))
	if err != nil {
		return err
	}
	_, err = w.Buffer.WriteString(s)
	return
}

func (w *ProtocolWriter) WriteUUID(u uuid.UUID) (err error) {
	b := u[:]
	err = w.WriteByteArray(b)
	return err
}

func (w *ProtocolWriter) WriteNBTCompound(tag nbt.Compound) (err error) {
	err = nbt.StoreUncompressed(w, tag, "")
	if err != nil {
		return err
	}
	return
}

func (w *ProtocolWriter) WriteBlockPosition(i world.BlockPosition) (err error) {
	return w.WriteUInt64(
		((uint64(i.X) & 0x3FFFFFF) << 38) |
			((uint64(i.Y) & 0xFFF) << 26) |
			(uint64(i.Z) & 0x3FFFFFF))
}

func (w *ProtocolWriter) WriteSlotItem(item slot.SlotItem) (err error) {
	if item.ID == 0 {
		err = w.WriteInt16(-1)
		return
	}
	err = w.WriteUInt16(item.ID)
	if err != nil {
		return
	}
	err = w.WriteUInt8(item.Amount)
	if err != nil {
		return
	}
	err = w.WriteUInt16(item.Durability)
	if err != nil {
		return
	}
	if item.NBT != nil || len(item.NBT) > 0 {
		err = w.WriteNBTCompound(item.NBT)
		if err != nil {
			return
		}
	} else {
		err = w.WriteByte(0)
		if err != nil {
			return
		}
	}
	return
}

func (w *ProtocolWriter) WriteMetadata(entries metadata.MetadataMap) (err error) {
	for index, value := range entries {
		if value == nil {
			continue
		}
		indexType := index.Type
		indexId := index.Index
		err = w.WriteByte((uint8(indexType) << 5) | indexId)
		if err != nil {
			return
		}
		switch indexType {
		case metadata.MetadataTypeByte:
			{
				err = w.WriteByte(value.(uint8))
				if err != nil {
					return
				}
			}
		case metadata.MetadataTypeShort:
			{
				err = w.WriteUInt16(value.(uint16))
				if err != nil {
					return
				}
			}
		case metadata.MetadataTypeInt:
			{
				err = w.WriteInt32(value.(int32))
				if err != nil {
					return
				}
			}
		case metadata.MetadataTypeFloat:
			{
				err = w.WriteFloat64(value.(float64))
				if err != nil {
					return
				}
			}
		case metadata.MetadataTypeString:
			{
				err = w.WriteString(value.(string))
				if err != nil {
					return
				}
			}
		case metadata.MetadataTypeSlot:
			{
				// Write slot
			}
		case metadata.MetadataTypeVector:
			{
				// Write vector
			}
		case metadata.MetadataTypeEulerAngle:
			{
				// Write euler angle
			}
		}
	}
	err = w.WriteByte(127)
	return
}
