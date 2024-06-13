package metadata

type MetadataIndex struct {
	Index     uint8
	Type      MetadataType
	AppliesTo string
}

var (
	MetadataIndexStatus = MetadataIndex{
		Index:     0,
		Type:      MetadataTypeByte,
		AppliesTo: "ENTITY",
	}
	MetadataAirTime = MetadataIndex{
		Index:     1,
		Type:      MetadataTypeShort,
		AppliesTo: "ENTITY",
	}
	MetadataSilent = MetadataIndex{
		Index:     4,
		Type:      MetadataTypeByte,
		AppliesTo: "ENTITY",
	}

	MetadataNameTag = MetadataIndex{
		Index:     2,
		Type:      MetadataTypeString,
		AppliesTo: "ENTITY",
	}
	MetadataShowNameTag = MetadataIndex{
		Index:     3,
		Type:      MetadataTypeByte,
		AppliesTo: "ENTITY",
	}

	MetadataHealth = MetadataIndex{
		Index:     6,
		Type:      MetadataTypeFloat,
		AppliesTo: "LIVINGENTITY",
	}
	MetadataPotionCOlor = MetadataIndex{
		Index:     7,
		Type:      MetadataTypeInt,
		AppliesTo: "LIVINGENTITY",
	}
	MetadataPotionAmbient = MetadataIndex{
		Index:     8,
		Type:      MetadataTypeByte,
		AppliesTo: "LIVINGENTITY",
	}
	MetadataArrowCount = MetadataIndex{
		Index:     9,
		Type:      MetadataTypeByte,
		AppliesTo: "LIVINGENTITY",
	}
	MetadataNoAI = MetadataIndex{
		Index:     15,
		Type:      MetadataTypeByte,
		AppliesTo: "LIVINGENTITY",
	}

	MetadataAge = MetadataIndex{
		Index:     12,
		Type:      MetadataTypeByte,
		AppliesTo: "Ageable",
	}
)
