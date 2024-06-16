package readerwriter

import (
	"encoding/binary"
	"math"

	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/metadata"
	"github.com/git-fal7/sealantern/pkg/slot"

	"github.com/google/uuid"
	"github.com/seebs/nbt"
)

func (w *ConnReadWrite) WriteVarInt(i int) (err error) {
	buff := w.Buffer[:]
	length := binary.PutUvarint(buff, uint64(i))
	_, err = w.Wtr.Write(buff[:length])
	return err
}

func (w *ConnReadWrite) WriteBool(b bool) (err error) {
	buff := w.Buffer[:1]
	if b {
		buff[0] = 0x01
	} else {
		buff[0] = 0x00
	}
	_, err = w.Wtr.Write(buff)
	if err != nil {
		return err
	}
	return
}

func (w *ConnReadWrite) WriteNBTCompound(tag nbt.Compound) (err error) {
	err = nbt.StoreUncompressed(w.Wtr, tag, "")
	if err != nil {
		return err
	}
	return
}

func (w *ConnReadWrite) WriteUInt8(i uint8) (err error) {
	buff := w.Buffer[:1]
	buff[0] = i
	_, err = w.Wtr.Write(buff)
	if err != nil {
		return err
	}
	return
}

func (w *ConnReadWrite) WriteUInt16(i uint16) (err error) {
	buff := w.Buffer[:2]
	binary.BigEndian.PutUint16(buff, i)
	_, err = w.Wtr.Write(buff)
	if err != nil {
		return err
	}
	return
}

func (w *ConnReadWrite) WriteLittleEndianUInt16(i uint16) (err error) {
	buff := w.Buffer[:2]
	binary.LittleEndian.PutUint16(buff, i)
	_, err = w.Wtr.Write(buff)
	if err != nil {
		return err
	}
	return
}

func (r *ConnReadWrite) WriteUInt32(i uint32) (err error) {
	buff := r.Buffer[:4]
	binary.BigEndian.PutUint32(buff, i)
	_, err = r.Wtr.Write(buff)
	if err != nil {
		return err
	}
	return
}

func (w *ConnReadWrite) WriteUInt64(i uint64) (err error) {
	buff := w.Buffer[:8]
	binary.BigEndian.PutUint64(buff, i)
	_, err = w.Wtr.Write(buff)
	if err != nil {
		return err
	}
	return
}

func (w *ConnReadWrite) WriteBlockPosition(i world.Position) (err error) {
	return w.WriteUInt64(
		((uint64(i.X) & 0x3FFFFFF) << 38) |
			((uint64(i.Y) & 0xFFF) << 26) |
			(uint64(i.Z) & 0x3FFFFFF))
}

func (w *ConnReadWrite) WriteFloat32(i float32) (err error) {
	buff := w.Buffer[:4]
	binary.BigEndian.PutUint32(buff, math.Float32bits(i))
	_, err = w.Wtr.Write(buff)
	if err != nil {
		return err
	}
	return
}

func (w *ConnReadWrite) WriteFloat64(i float64) (err error) {
	buff := w.Buffer[:8]
	binary.BigEndian.PutUint64(buff, math.Float64bits(i))
	_, err = w.Wtr.Write(buff)
	if err != nil {
		return err
	}
	return
}

func (w *ConnReadWrite) WriteByteArray(data []byte) (err error) {
	_, err = w.Wtr.Write(data)
	return err
}

func (w *ConnReadWrite) WriteString(s string) (err error) {
	buff := []byte(s)
	err = w.WriteVarInt(len(buff))
	if err != nil {
		return err
	}
	_, err = w.Wtr.Write(buff)
	return err
}

func (w *ConnReadWrite) WriteStringRestricted(s string, max int) (err error) {
	buff := []byte(s)
	if len(buff) > max {
		buff = buff[:max]
	}
	err = w.WriteVarInt(len(buff))
	if err != nil {
		return err
	}
	_, err = w.Wtr.Write(buff)
	return err
}

func (w *ConnReadWrite) WriteUUID(u uuid.UUID) (err error) {
	b, err := u.MarshalBinary()
	if err != nil {
		return err
	}
	err = w.WriteByteArray(b)
	return err
}

func (r *ConnReadWrite) WriteInt8(i int8) (err error) {
	err = binary.Write(r.Wtr, binary.BigEndian, i)
	return
}

func (r *ConnReadWrite) WriteInt32(i int32) (err error) {
	err = binary.Write(r.Wtr, binary.BigEndian, i)
	return
}

func (r *ConnReadWrite) WriteInt16(i int16) (err error) {
	err = binary.Write(r.Wtr, binary.BigEndian, i)
	return
}

func (w *ConnReadWrite) WriteMetadata(entries metadata.MetadataMap) (err error) {
	for index, value := range entries {
		if value == nil {
			continue
		}
		indexType := index.Type
		indexId := index.Index
		err = w.WriteUInt8((uint8(indexType) << 5) | indexId)
		if err != nil {
			return
		}
		switch indexType {
		case metadata.MetadataTypeByte:
			{
				err = w.WriteUInt8(value.(uint8))
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
	err = w.WriteUInt8(127)
	return
}

func (w *ConnReadWrite) WriteSlotItem(item slot.SlotItem) (err error) {
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
		err = w.WriteUInt8(0)
		if err != nil {
			return
		}
	}
	return
}
