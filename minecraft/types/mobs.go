package types

type MobType uint8

const (
	MobTypeMob MobType = iota + 48
	MobTypeMonster
	// Subclasses
	MobTypeCreeper
	MobTypeSkeleton
	MobTypeSpider
	MobTypeGiantZombie
	MobTypeZombie
	MobTypeSlime
	MobTypeGhast
	MobTypeZombiePigman
	MobTypeEnderman
	MobTypeCaveSpider
	MobTypeSilverFish
	MobTypeBlaze
	MobTypeMagmaCube
	MobTypeEnderDragon
	MobTypeWither
	MobTypeBat
	MobTypeWitch
	MobTypeEndermite
	MobTypeGuardian
	MobTypePig       MobType = 90
	MobTypeSheep     MobType = 91
	MobTypeCow       MobType = 92
	MobTypeChicken   MobType = 93
	MobTypeSquid     MobType = 94
	MobTypeWolf      MobType = 95
	MobTypeMooshroom MobType = 96
	MobTypeSnowman   MobType = 97
	MobTypeOcelot    MobType = 98
	MobTypeIronGolem MobType = 99
	MobTypeHorse     MobType = 100
	MobTypeRabbit    MobType = 101
	MobTypeVillager  MobType = 120
)
