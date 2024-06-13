package metadata

type MetadataType uint8

const (
	MetadataTypeByte MetadataType = iota
	MetadataTypeShort
	MetadataTypeInt
	MetadataTypeFloat
	MetadataTypeString
	MetadataTypeSlot
	MetadataTypeVector
	MetadataTypeEulerAngle
)
