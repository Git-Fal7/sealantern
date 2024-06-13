package types

type DisplayedSkinParts uint8

const (
	SkinPartCape        DisplayedSkinParts = 0x01
	SkinPartJacket      DisplayedSkinParts = 0x02
	SkinPartLeftSleeve  DisplayedSkinParts = 0x04
	SkinPartRightSleeve DisplayedSkinParts = 0x08
	SkinPartLeftPants   DisplayedSkinParts = 0x10
	SkinPartRightPants  DisplayedSkinParts = 0x20
	SkinPartHat         DisplayedSkinParts = 0x40
)
