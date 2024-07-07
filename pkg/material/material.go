package material

type Material struct {
	id            uint16
	maxStack      uint8
	maxDurability uint16
}

func (m Material) GetID() uint16 {
	return m.id
}
func (m Material) GetStack() uint8 {
	return m.maxStack
}
func (m Material) GetMaxDurability() uint16 {
	return m.maxDurability
}

var (
	Air                    Material = newMaterialWithStack(0, 0)
	Stone                  Material = newMaterial(1)
	Grass                  Material = newMaterial(2)
	Dirt                   Material = newMaterial(3)
	Cobblestone            Material = newMaterial(4)
	Wood                   Material = newMaterial(5)
	Sapling                Material = newMaterial(6)
	Bedrock                Material = newMaterial(7)
	Water                  Material = newMaterial(8)
	StationaryWater        Material = newMaterial(9)
	Lava                   Material = newMaterial(10)
	StationaryLava         Material = newMaterial(11)
	Sand                   Material = newMaterial(12)
	Gravel                 Material = newMaterial(13)
	GoldOre                Material = newMaterial(14)
	IronOre                Material = newMaterial(15)
	CoalOre                Material = newMaterial(16)
	Log                    Material = newMaterial(17) // Oak, Spruce, Birch and Jungle
	Leaves                 Material = newMaterial(18) // Oak, Spruce, Birch and Jungle
	Sponge                 Material = newMaterial(19)
	Glass                  Material = newMaterial(20)
	LapisOre               Material = newMaterial(21)
	LapisBlock             Material = newMaterial(22)
	Dispenser              Material = newMaterial(23)
	Sandstone              Material = newMaterial(24)
	NoteBlock              Material = newMaterial(25)
	BedBlock               Material = newMaterial(26)
	PoweredRail            Material = newMaterial(27)
	DetectorRail           Material = newMaterial(28)
	PistonStickyBase       Material = newMaterial(29)
	Web                    Material = newMaterial(30)
	LongGrass              Material = newMaterial(31)
	DeadBush               Material = newMaterial(32)
	PistonBase             Material = newMaterial(33)
	PistonExtension        Material = newMaterial(34)
	Wool                   Material = newMaterial(35)
	PistonMovingPiece      Material = newMaterial(36)
	YellowFlower           Material = newMaterial(37)
	RedRose                Material = newMaterial(38)
	BrownMushroom          Material = newMaterial(39)
	RedMushroom            Material = newMaterial(40)
	GoldBlock              Material = newMaterial(41)
	IronBlock              Material = newMaterial(42)
	DoubleStep             Material = newMaterial(43)
	Step                   Material = newMaterial(44)
	BrickBlock             Material = newMaterial(45)
	TNT                    Material = newMaterial(46)
	Bookshelf              Material = newMaterial(47)
	MossyCobbleston        Material = newMaterial(48)
	Obsidian               Material = newMaterial(49)
	Torch                  Material = newMaterial(50)
	Fire                   Material = newMaterial(51)
	MobSpawner             Material = newMaterial(52)
	WoodStairs             Material = newMaterial(53)
	Chest                  Material = newMaterial(54)
	RedstoneWire           Material = newMaterial(55)
	DiamondOre             Material = newMaterial(56)
	DiamondBlock           Material = newMaterial(57)
	WorkBench              Material = newMaterial(58)
	Crops                  Material = newMaterial(59)
	Soil                   Material = newMaterial(60)
	Furnance               Material = newMaterial(61)
	BurningFurnance        Material = newMaterial(62)
	SignPost               Material = newMaterial(63)
	WoodenDoor             Material = newMaterial(64)
	Ladder                 Material = newMaterial(65)
	Rails                  Material = newMaterial(66)
	CobblestoneStairs      Material = newMaterial(67)
	WallSign               Material = newMaterial(68)
	Lever                  Material = newMaterial(69)
	StonePlate             Material = newMaterial(70)
	IronDoorBlock          Material = newMaterial(71)
	WoodPlate              Material = newMaterial(72)
	RedstoneOre            Material = newMaterial(73)
	GlowingRedstoneOre     Material = newMaterial(74)
	RedstoneTorchOff       Material = newMaterial(75)
	RedstoneTorchOn        Material = newMaterial(76)
	StoneButton            Material = newMaterial(77)
	Snow                   Material = newMaterial(78)
	Ice                    Material = newMaterial(79)
	SnowBlock              Material = newMaterial(80)
	Cactus                 Material = newMaterial(81)
	Clay                   Material = newMaterial(82)
	SugarCaneBlock         Material = newMaterial(83)
	JukeBox                Material = newMaterial(84)
	Fence                  Material = newMaterial(85)
	Pumpkin                Material = newMaterial(86)
	Netherrack             Material = newMaterial(87)
	SoulSand               Material = newMaterial(88)
	Glowstone              Material = newMaterial(89)
	Portal                 Material = newMaterial(90)
	JackOLantern           Material = newMaterial(91)
	CakeBlock              Material = newMaterial(92)
	RedstoneRepeaterOff    Material = newMaterial(93)
	RedstoneRepeaterOn     Material = newMaterial(94)
	StainedGlass           Material = newMaterial(95)
	TrapDoor               Material = newMaterial(96)
	StoneMonsterEgg        Material = newMaterial(97)
	StoneBrick             Material = newMaterial(98)
	BrownMushroomBlock     Material = newMaterial(99)
	RedMushroomBlock       Material = newMaterial(100)
	IronBars               Material = newMaterial(101)
	GlassPane              Material = newMaterial(102)
	MelonBlock             Material = newMaterial(103)
	PumpkinStem            Material = newMaterial(104)
	MelonStem              Material = newMaterial(105)
	Vine                   Material = newMaterial(106)
	OakFenceGate           Material = newMaterial(107)
	BrickStairs            Material = newMaterial(108)
	SmoothStairs           Material = newMaterial(109)
	Mycel                  Material = newMaterial(110)
	WaterLily              Material = newMaterial(111)
	NetherBrickBlock       Material = newMaterial(112)
	NetherFence            Material = newMaterial(113)
	NetherBrickStairs      Material = newMaterial(114)
	NetherWartBlock        Material = newMaterial(115)
	EnchantmentTable       Material = newMaterial(116)
	BrewingStandBlock      Material = newMaterial(117)
	CauldronBlock          Material = newMaterial(118)
	EnderPortal            Material = newMaterial(119)
	EnderPortalFrame       Material = newMaterial(120)
	EnderStone             Material = newMaterial(121)
	DragonEgg              Material = newMaterial(122)
	RedstoneLampOff        Material = newMaterial(123)
	RedstoneLampOn         Material = newMaterial(124)
	WoodDoubleStep         Material = newMaterial(125)
	WoodStep               Material = newMaterial(126)
	CocoaBlock             Material = newMaterial(127)
	SandstoneStairs        Material = newMaterial(128)
	EmeraldOre             Material = newMaterial(129)
	EnderChest             Material = newMaterial(130)
	TripwireHook           Material = newMaterial(131)
	TripwireString         Material = newMaterial(132)
	EmeraldBlock           Material = newMaterial(133)
	SpruceWoodStairs       Material = newMaterial(134)
	BirchWoodStairs        Material = newMaterial(135)
	JungleWoodStairs       Material = newMaterial(136)
	CommandBlock           Material = newMaterial(137)
	Beacon                 Material = newMaterial(138)
	CobblestoneWall        Material = newMaterial(139)
	FlowerPotBlock         Material = newMaterial(140)
	CarrotBlock            Material = newMaterial(141)
	PotatoBlock            Material = newMaterial(142)
	WoodButton             Material = newMaterial(143)
	SkullBlock             Material = newMaterial(144)
	Anvil                  Material = newMaterial(145)
	TrappedChest           Material = newMaterial(146)
	GoldPlate              Material = newMaterial(147)
	IronPlate              Material = newMaterial(148)
	RedstoneComparatorOff  Material = newMaterial(149)
	RedstoneComparatorOn   Material = newMaterial(150)
	DaylightSensor         Material = newMaterial(151)
	RedstoneBlock          Material = newMaterial(152)
	QuartzOre              Material = newMaterial(153)
	Hopper                 Material = newMaterial(154)
	QuartzBlock            Material = newMaterial(155)
	QuartzStairs           Material = newMaterial(156)
	ActivatorRail          Material = newMaterial(157)
	Dropper                Material = newMaterial(158)
	StainedClay            Material = newMaterial(159)
	StainedGlassPane       Material = newMaterial(160)
	Leaves2                Material = newMaterial(161) // Acacia and Dark oak
	Log2                   Material = newMaterial(162) // Acacia and Dark oak
	AcaciaStairs           Material = newMaterial(163)
	DarkOakStairs          Material = newMaterial(164)
	SlimeBlock             Material = newMaterial(165)
	Barrier                Material = newMaterial(166)
	IronTrapdoor           Material = newMaterial(167)
	Prismarine             Material = newMaterial(168)
	SeaLantern             Material = newMaterial(169)
	HayBlock               Material = newMaterial(170)
	Carpet                 Material = newMaterial(171)
	HardClay               Material = newMaterial(172)
	CoalBlock              Material = newMaterial(173)
	PackedIce              Material = newMaterial(174)
	DoublePlant            Material = newMaterial(175)
	StandingBanner         Material = newMaterial(176)
	WallBanner             Material = newMaterial(177)
	DaylightSensorInverted Material = newMaterial(178)
	RedSandstone           Material = newMaterial(179)
	RedSandstoneStairs     Material = newMaterial(180)
	DoubleStoneSlab2       Material = newMaterial(181)
	StoneSlab2             Material = newMaterial(182)
	SpruceFenceGate        Material = newMaterial(183)
	BirchFenceGate         Material = newMaterial(184)
	JungleFenceGate        Material = newMaterial(185)
	DarkOakFenceGate       Material = newMaterial(186)
	AcaciaFenceGate        Material = newMaterial(187)
	SpruceFence            Material = newMaterial(188)
	BirchFence             Material = newMaterial(189)
	JungleFence            Material = newMaterial(190)
	DarkOakFence           Material = newMaterial(191)
	AcaciaFence            Material = newMaterial(192)
	SpruceDoorBlock        Material = newMaterial(193)
	BirchDoorBlock         Material = newMaterial(194)
	JungleDoorBlock        Material = newMaterial(195)
	AcaciaDoorBlock        Material = newMaterial(196)
	DarkOakDoorBlock       Material = newMaterial(197)
	//
	IronShovel          Material   = newMaterialWithStackDurability(256, 1, 250)
	IronPickaxe         Material   = newMaterialWithStackDurability(257, 1, 250)
	IronAxe             Material   = newMaterialWithStackDurability(258, 1, 250)
	FlintAndSteel       Material   = newMaterialWithStackDurability(259, 1, 64)
	Apple               Material   = newMaterial(260)
	Bow                 Material   = newMaterialWithStackDurability(261, 1, 384)
	Arrow               Material   = newMaterial(262)
	Coal                Material   = newMaterial(263)
	Diamond             Material   = newMaterial(264)
	IronIngot           Material   = newMaterial(265)
	GoldIngot           Material   = newMaterial(266)
	IronSword           Material   = newMaterialWithStackDurability(267, 1, 250)
	WoodSword           Material   = newMaterialWithStackDurability(268, 1, 59)
	WoodShovel          Material   = newMaterialWithStackDurability(269, 1, 59)
	WoodPickaxe         Material   = newMaterialWithStackDurability(270, 1, 59)
	WoodAxe             Material   = newMaterialWithStackDurability(271, 1, 59)
	StoneSword          Material   = newMaterialWithStackDurability(272, 1, 131)
	StoneShovel         Material   = newMaterialWithStackDurability(273, 1, 131)
	StonePickaxe        Material   = newMaterialWithStackDurability(274, 1, 131)
	StoneAxe            Material   = newMaterialWithStackDurability(275, 1, 131)
	DiamondSword        Material   = newMaterialWithStackDurability(276, 1, 1561)
	DiamondShovel       Material   = newMaterialWithStackDurability(277, 1, 1561)
	DiamondPickaxe      Material   = newMaterialWithStackDurability(278, 1, 1561)
	DiamondAxe          Material   = newMaterialWithStackDurability(279, 1, 1561)
	Stick               Material   = newMaterial(280)
	Bowl                Material   = newMaterial(281)
	MushroomSoup        Material   = newMaterialWithStack(282, 1)
	GoldSword           Material   = newMaterialWithStackDurability(283, 1, 32)
	GoldShovel          Material   = newMaterialWithStackDurability(284, 1, 32)
	GoldPickaxe         Material   = newMaterialWithStackDurability(285, 1, 32)
	GoldAxe             Material   = newMaterialWithStackDurability(286, 1, 32)
	String              Material   = newMaterial(287)
	Feather             Material   = newMaterial(288)
	Gunpowder           Material   = newMaterial(289)
	WoodHoe             Material   = newMaterialWithStackDurability(290, 1, 59)
	StoneHoe            Material   = newMaterialWithStackDurability(291, 1, 131)
	IronHoe             Material   = newMaterialWithStackDurability(292, 1, 250)
	DiamondHoe          Material   = newMaterialWithStackDurability(293, 1, 1561)
	GoldHoe             Material   = newMaterialWithStackDurability(294, 1, 32)
	WheatSeeds          Material   = newMaterial(295)
	Wheat               Material   = newMaterial(296)
	Bread               Material   = newMaterial(297)
	LeatherHelmet       Material   = newMaterialWithStackDurability(298, 1, 55)
	LeatherChestplate   Material   = newMaterialWithStackDurability(299, 1, 80)
	LeatherLeggings     Material   = newMaterialWithStackDurability(300, 1, 75)
	LeatherBoots        Material   = newMaterialWithStackDurability(301, 1, 65)
	ChainmailHelmet     Material   = newMaterialWithStackDurability(302, 1, 165)
	ChainmailChestplate Material   = newMaterialWithStackDurability(303, 1, 240)
	ChainmailLeggings   Material   = newMaterialWithStackDurability(304, 1, 225)
	ChainmailBoots      Material   = newMaterialWithStackDurability(305, 1, 195)
	IronHelmet          Material   = newMaterialWithStackDurability(306, 1, 165)
	IronChestplate      Material   = newMaterialWithStackDurability(307, 1, 240)
	IronLeggings        Material   = newMaterialWithStackDurability(308, 1, 225)
	IronBoots           Material   = newMaterialWithStackDurability(309, 1, 195)
	DiamondHelmet       Material   = newMaterialWithStackDurability(310, 1, 363)
	DiamondChestplate   Material   = newMaterialWithStackDurability(311, 1, 528)
	DiamondLeggings     Material   = newMaterialWithStackDurability(312, 1, 495)
	DiamondBoots        Material   = newMaterialWithStackDurability(313, 1, 429)
	GoldHelmet          Material   = newMaterialWithStackDurability(314, 1, 77)
	GoldChestplate      Material   = newMaterialWithStackDurability(315, 1, 112)
	GoldLeggings        Material   = newMaterialWithStackDurability(316, 1, 105)
	GoldBoots           Material   = newMaterialWithStackDurability(317, 1, 91)
	Flint               Material   = newMaterial(318)
	RawPorkchop         Material   = newMaterial(319)
	CookedPorkchop      Material   = newMaterial(320)
	Painting            Material   = newMaterial(321)
	GoldenApple         Material   = newMaterial(322)
	Sign                Material   = newMaterialWithStack(323, 16)
	WoodDoor            Material   = newMaterial(324)
	Bucket              Material   = newMaterialWithStack(325, 16)
	WaterBucket         Material   = newMaterialWithStack(326, 1)
	LavaBucket          Material   = newMaterialWithStack(327, 1)
	Minecart            Material   = newMaterialWithStack(328, 1)
	Saddle              Material   = newMaterialWithStack(329, 1)
	IronDoor            Material   = newMaterial(330)
	Redstone            Material   = newMaterial(331)
	Snowball            Material   = newMaterialWithStack(332, 16)
	Boat                Material   = newMaterialWithStack(333, 1)
	Leather             Material   = newMaterial(334)
	MilkBucket          Material   = newMaterialWithStack(335, 1)
	Brick               Material   = newMaterial(336)
	ClayBall            Material   = newMaterial(337)
	SugarCane           Material   = newMaterial(338)
	Paper               Material   = newMaterial(339)
	Book                Material   = newMaterial(340)
	SlimeBall           Material   = newMaterial(341)
	StorageMinecart     Material   = newMaterial(342)
	PoweredMinecart     Material   = newMaterial(343)
	Egg                 Material   = newMaterialWithStack(344, 16)
	Compass             Material   = newMaterial(345)
	FishingRod          Material   = newMaterialWithStackDurability(346, 1, 64)
	Clock               Material   = newMaterial(347)
	GlowstoneDust       Material   = newMaterial(348)
	RawFish             Material   = newMaterial(349)
	CookedFish          Material   = newMaterial(350)
	InkSack             Material   = newMaterial(351)
	Bone                Material   = newMaterial(352)
	Sugar               Material   = newMaterial(353)
	Cake                Material   = newMaterialWithStack(354, 1)
	Bed                 Material   = newMaterialWithStack(355, 1)
	RedstoneRepeater    Material   = newMaterial(356)
	Cookie              Material   = newMaterial(357)
	FilledMap           Material   = newMaterial(358)
	Shears              Material   = newMaterialWithStackDurability(359, 1, 238)
	Melon               Material   = newMaterial(360)
	PumpkinSeeds        Material   = newMaterial(361)
	MelonSeeds          Material   = newMaterial(362)
	RawBeef             Material   = newMaterial(363)
	CookedBeef          Material   = newMaterial(364)
	RawChicken          Material   = newMaterial(365)
	CookedChicken       Material   = newMaterial(366)
	RottenFlesh         Material   = newMaterial(367)
	EnderPearl          Material   = newMaterialWithStack(368, 16)
	BlazeRod            Material   = newMaterial(369)
	GhastTear           Material   = newMaterial(370)
	GoldNugget          Material   = newMaterial(371)
	NetherWart          Material   = newMaterial(372)
	Potion              Material   = newMaterialWithStack(373, 1)
	GlassBottle         Material   = newMaterial(374)
	SpiderEye           Material   = newMaterial(375)
	FermentedSpiderEye  Material   = newMaterial(376)
	BlazePowder         Material   = newMaterial(377)
	MagmaCream          Material   = newMaterial(378)
	BrewingStand        Material   = newMaterial(379)
	Cauldron            Material   = newMaterial(380)
	EyeOfEnder          Material   = newMaterial(381)
	GlisteringMelon     Material   = newMaterial(382)
	MonsterEgg          Material   = newMaterial(383)
	ExperienceBottle    Material   = newMaterial(384)
	Fireball            Material   = newMaterial(385)
	BookAndQuill        Material   = newMaterialWithStack(386, 1)
	WrittenBook         Material   = newMaterialWithStack(387, 1)
	Emerald             Material   = newMaterial(388)
	ItemFrame           Material   = newMaterial(389)
	FlowerPot           Material   = newMaterial(390)
	Carrot              Material   = newMaterial(391)
	Potato              Material   = newMaterial(392)
	BakedPotato         Material   = newMaterial(393)
	PoisonousPotato     Material   = newMaterial(394)
	EmptyMap            Material   = newMaterial(395)
	GoldenCarrot        Material   = newMaterial(396)
	Skull               Material   = newMaterial(397)
	CarrotOnAStick      Material   = newMaterialWithStackDurability(398, 1, 25)
	NetherStar          Material   = newMaterial(399)
	PumpkinPie          Material   = newMaterial(400)
	Firework            Material   = newMaterial(401)
	FireworkCharge      Material   = newMaterial(402)
	EnchantedBook       Material   = newMaterialWithStack(403, 1)
	RedstoneComparator  Material   = newMaterial(404)
	NetherBrick         Material   = newMaterial(405)
	Quartz              Material   = newMaterial(406)
	ExplosiveMinecart   Material   = newMaterialWithStack(407, 1)
	HopperMinecart      Material   = newMaterialWithStack(408, 1)
	PrismarineShard     Material   = newMaterial(409)
	PrismarineCrystals  Material   = newMaterial(410)
	RawRabbit           Material   = newMaterial(411)
	CookedRabbit        Material   = newMaterial(412)
	RabbitStew          Material   = newMaterialWithStack(413, 1)
	RabbitFoot          Material   = newMaterial(414)
	RabbitHide          Material   = newMaterial(415)
	ArmorStand          Material   = newMaterialWithStack(416, 16)
	IronHorseArmor      Material   = newMaterialWithStack(417, 1)
	GoldHorseArmor      Material   = newMaterialWithStack(418, 1)
	DiamondHorseArmor   Material   = newMaterialWithStack(419, 1)
	Leash               Material   = newMaterial(420)
	NameTag             Material   = newMaterial(421)
	CommandMinecart     Material   = newMaterialWithStack(422, 1)
	Mutton              Material   = newMaterial(423)
	CookedMutton        Material   = newMaterial(424)
	Banner              Material   = newMaterialWithStack(425, 16)
	SpruceDoor          Material   = newMaterial(427)
	BirchDoor           Material   = newMaterial(428)
	JungleDoor          Material   = newMaterial(429)
	AcaciaDoor          Material   = newMaterial(430)
	DarkOakDoor         Material   = newMaterial(431)
	Record13            Material   = newMaterialWithStack(2256, 1)
	RecordCat           Material   = newMaterialWithStack(2257, 1)
	RecordBlocks        Material   = newMaterialWithStack(2258, 1)
	RecordChirp         Material   = newMaterialWithStack(2259, 1)
	RecordFar           Material   = newMaterialWithStack(2260, 1)
	RecordMall          Material   = newMaterialWithStack(2261, 1)
	RecordMellohi       Material   = newMaterialWithStack(2262, 1)
	RecordStall         Material   = newMaterialWithStack(2263, 1)
	RecordStrad         Material   = newMaterialWithStack(2264, 1)
	RecordWard          Material   = newMaterialWithStack(2265, 1)
	Record11            Material   = newMaterialWithStack(2266, 1)
	RecordWait          Material   = newMaterialWithStack(2267, 1)
	materialArray       []Material = []Material{
		Air,
		Stone,
		Grass,
		Dirt,
		Cobblestone,
		Wood,
		Sapling,
		Bedrock,
		Water,
		StationaryWater,
		Lava,
		StationaryLava,
		Sand,
		Gravel,
		GoldOre,
		IronOre,
		CoalOre,
		Log,
		Leaves,
		Sponge,
		Glass,
		LapisOre,
		LapisBlock,
		Dispenser,
		Sandstone,
		NoteBlock,
		BedBlock,
		PoweredRail,
		DetectorRail,
		PistonStickyBase,
		Web,
		LongGrass,
		DeadBush,
		PistonBase,
		PistonExtension,
		Wool,
		PistonMovingPiece,
		YellowFlower,
		RedRose,
		BrownMushroom,
		RedMushroom,
		GoldBlock,
		IronBlock,
		DoubleStep,
		Step,
		BrickBlock,
		TNT,
		Bookshelf,
		MossyCobbleston,
		Obsidian,
		Torch,
		Fire,
		MobSpawner,
		WoodStairs,
		Chest,
		RedstoneWire,
		DiamondOre,
		DiamondBlock,
		WorkBench,
		Crops,
		Soil,
		Furnance,
		BurningFurnance,
		SignPost,
		WoodenDoor,
		Ladder,
		Rails,
		CobblestoneStairs,
		WallSign,
		Lever,
		StonePlate,
		IronDoorBlock,
		WoodPlate,
		RedstoneOre,
		GlowingRedstoneOre,
		RedstoneTorchOff,
		RedstoneTorchOn,
		StoneButton,
		Snow,
		Ice,
		SnowBlock,
		Cactus,
		Clay,
		SugarCaneBlock,
		JukeBox,
		Fence,
		Pumpkin,
		Netherrack,
		SoulSand,
		Glowstone,
		Portal,
		JackOLantern,
		CakeBlock,
		RedstoneRepeaterOff,
		RedstoneRepeaterOn,
		StainedGlass,
		TrapDoor,
		StoneMonsterEgg,
		StoneBrick,
		BrownMushroomBlock,
		RedMushroomBlock,
		IronBars,
		GlassPane,
		MelonBlock,
		PumpkinStem,
		MelonStem,
		Vine,
		OakFenceGate,
		BrickStairs,
		SmoothStairs,
		Mycel,
		WaterLily,
		NetherBrickBlock,
		NetherFence,
		NetherBrickStairs,
		NetherWartBlock,
		EnchantmentTable,
		BrewingStandBlock,
		CauldronBlock,
		EnderPortal,
		EnderPortalFrame,
		EnderStone,
		DragonEgg,
		RedstoneLampOff,
		RedstoneLampOn,
		WoodDoubleStep,
		WoodStep,
		CocoaBlock,
		SandstoneStairs,
		EmeraldOre,
		EnderChest,
		TripwireHook,
		TripwireString,
		EmeraldBlock,
		SpruceWoodStairs,
		BirchWoodStairs,
		JungleWoodStairs,
		CommandBlock,
		Beacon,
		CobblestoneWall,
		FlowerPotBlock,
		CarrotBlock,
		PotatoBlock,
		WoodButton,
		SkullBlock,
		Anvil,
		TrappedChest,
		GoldPlate,
		IronPlate,
		RedstoneComparatorOff,
		RedstoneComparatorOn,
		DaylightSensor,
		RedstoneBlock,
		QuartzOre,
		Hopper,
		QuartzBlock,
		QuartzStairs,
		ActivatorRail,
		Dropper,
		StainedClay,
		StainedGlassPane,
		Leaves2,
		Log2,
		AcaciaStairs,
		DarkOakStairs,
		SlimeBlock,
		Barrier,
		IronTrapdoor,
		Prismarine,
		SeaLantern,
		HayBlock,
		Carpet,
		HardClay,
		CoalBlock,
		PackedIce,
		DoublePlant,
		StandingBanner,
		WallBanner,
		DaylightSensorInverted,
		RedSandstone,
		RedSandstoneStairs,
		DoubleStoneSlab2,
		StoneSlab2,
		SpruceFenceGate,
		BirchFenceGate,
		JungleFenceGate,
		DarkOakFenceGate,
		AcaciaFenceGate,
		SpruceFence,
		BirchFence,
		JungleFence,
		DarkOakFence,
		AcaciaFence,
		SpruceDoorBlock,
		BirchDoorBlock,
		JungleDoorBlock,
		AcaciaDoorBlock,
		DarkOakDoorBlock,
		//
		IronShovel,
		IronPickaxe,
		IronAxe,
		FlintAndSteel,
		Apple,
		Bow,
		Arrow,
		Coal,
		Diamond,
		IronIngot,
		GoldIngot,
		IronSword,
		WoodSword,
		WoodShovel,
		WoodPickaxe,
		WoodAxe,
		StoneSword,
		StoneShovel,
		StonePickaxe,
		StoneAxe,
		DiamondSword,
		DiamondShovel,
		DiamondPickaxe,
		DiamondAxe,
		Stick,
		Bowl,
		MushroomSoup,
		GoldSword,
		GoldShovel,
		GoldPickaxe,
		GoldAxe,
		String,
		Feather,
		Gunpowder,
		WoodHoe,
		StoneHoe,
		IronHoe,
		DiamondHoe,
		GoldHoe,
		WheatSeeds,
		Wheat,
		Bread,
		LeatherHelmet,
		LeatherChestplate,
		LeatherLeggings,
		LeatherBoots,
		ChainmailHelmet,
		ChainmailChestplate,
		ChainmailLeggings,
		ChainmailBoots,
		IronHelmet,
		IronChestplate,
		IronLeggings,
		IronBoots,
		DiamondHelmet,
		DiamondChestplate,
		DiamondLeggings,
		DiamondBoots,
		GoldHelmet,
		GoldChestplate,
		GoldLeggings,
		GoldBoots,
		Flint,
		RawPorkchop,
		CookedPorkchop,
		Painting,
		GoldenApple,
		Sign,
		WoodDoor,
		Bucket,
		WaterBucket,
		LavaBucket,
		Minecart,
		Saddle,
		IronDoor,
		Redstone,
		Snowball,
		Boat,
		Leather,
		MilkBucket,
		Brick,
		ClayBall,
		SugarCane,
		Paper,
		Book,
		SlimeBall,
		StorageMinecart,
		PoweredMinecart,
		Egg,
		Compass,
		FishingRod,
		Clock,
		GlowstoneDust,
		RawFish,
		CookedFish,
		InkSack,
		Bone,
		Sugar,
		Cake,
		Bed,
		RedstoneRepeater,
		Cookie,
		FilledMap,
		Shears,
		Melon,
		PumpkinSeeds,
		MelonSeeds,
		RawBeef,
		CookedBeef,
		RawChicken,
		CookedChicken,
		RottenFlesh,
		EnderPearl,
		BlazeRod,
		GhastTear,
		GoldNugget,
		NetherWart,
		Potion,
		GlassBottle,
		SpiderEye,
		FermentedSpiderEye,
		BlazePowder,
		MagmaCream,
		BrewingStand,
		Cauldron,
		EyeOfEnder,
		GlisteringMelon,
		MonsterEgg,
		ExperienceBottle,
		Fireball,
		BookAndQuill,
		WrittenBook,
		Emerald,
		ItemFrame,
		FlowerPot,
		Carrot,
		Potato,
		BakedPotato,
		PoisonousPotato,
		EmptyMap,
		GoldenCarrot,
		Skull,
		CarrotOnAStick,
		NetherStar,
		PumpkinPie,
		Firework,
		FireworkCharge,
		EnchantedBook,
		RedstoneComparator,
		NetherBrick,
		Quartz,
		ExplosiveMinecart,
		HopperMinecart,
		PrismarineShard,
		PrismarineCrystals,
		RawRabbit,
		CookedRabbit,
		RabbitStew,
		RabbitFoot,
		RabbitHide,
		ArmorStand,
		IronHorseArmor,
		GoldHorseArmor,
		DiamondHorseArmor,
		Leash,
		NameTag,
		CommandMinecart,
		Mutton,
		CookedMutton,
		Banner,
		SpruceDoor,
		BirchDoor,
		JungleDoor,
		AcaciaDoor,
		DarkOakDoor,
		Record13,
		RecordCat,
		RecordBlocks,
		RecordChirp,
		RecordFar,
		RecordMall,
		RecordMellohi,
		RecordStall,
		RecordStrad,
		RecordWard,
		Record11,
		RecordWait,
	}
)

func newMaterial(id uint16) Material {
	return Material{id: id, maxStack: 64}
}

func newMaterialWithStack(id uint16, maxStack uint8) Material {
	return Material{id: id, maxStack: maxStack}
}

func newMaterialWithStackDurability(id uint16, maxStack uint8, durability uint16) Material {
	return Material{id: id, maxStack: maxStack, maxDurability: durability}
}

func FindMaterialByID(id uint16) Material {
	for _, m := range materialArray {
		if m.id == id {
			return m
		}
	}
	return Air
}
