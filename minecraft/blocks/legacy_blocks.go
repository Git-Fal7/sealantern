package blocks

type LegacyBlock struct {
	Id   int
	Data int
	Name string
}

var legacy = []LegacyBlock{
	{
		0, 0, "minecraft:air",
	},
	{
		1, 0, "minecraft:stone",
	},
	{
		1, 1, "minecraft:granite",
	},
	{
		1, 2, "minecraft:polished_granite",
	},
	{
		1, 3, "minecraft:diorite",
	},
	{
		1, 4, "minecraft:polished_diorite",
	},
	{
		1, 5, "minecraft:andesite",
	},
	{
		1, 6, "minecraft:polished_andesite",
	},
	{
		2, 0, "minecraft:grass_block[snowy=false]",
	},
	{
		3, 0, "minecraft:dirt",
	},
	{
		3, 1, "minecraft:coarse_dirt",
	},
	{
		3, 2, "minecraft:podzol[snowy=false]",
	},
	{
		4, 0, "minecraft:cobblestone",
	},
	{
		5, 0, "minecraft:oak_planks",
	},
	{
		5, 1, "minecraft:spruce_planks",
	},
	{
		5, 2, "minecraft:birch_planks",
	},
	{
		5, 3, "minecraft:jungle_planks",
	},
	{
		5, 4, "minecraft:acacia_planks",
	},
	{
		5, 5, "minecraft:dark_oak_planks",
	},
	{
		6, 0, "minecraft:oak_sapling[stage=0]",
	},
	{
		6, 1, "minecraft:spruce_sapling[stage=0]",
	},
	{
		6, 2, "minecraft:birch_sapling[stage=0]",
	},
	{
		6, 3, "minecraft:jungle_sapling[stage=0]",
	},
	{
		6, 4, "minecraft:acacia_sapling[stage=0]",
	},
	{
		6, 5, "minecraft:dark_oak_sapling[stage=0]",
	},
	{
		6, 8, "minecraft:oak_sapling[stage=1]",
	},
	{
		6, 9, "minecraft:spruce_sapling[stage=1]",
	},
	{
		6, 10, "minecraft:birch_sapling[stage=1]",
	},
	{
		6, 11, "minecraft:jungle_sapling[stage=1]",
	},
	{
		6, 12, "minecraft:acacia_sapling[stage=1]",
	},
	{
		6, 13, "minecraft:dark_oak_sapling[stage=1]",
	},
	{
		7, 0, "minecraft:bedrock",
	},
	{
		8, 0, "minecraft:water[level=0]",
	},
	{
		8, 1, "minecraft:water[level=1]",
	},
	{
		8, 2, "minecraft:water[level=2]",
	},
	{
		8, 3, "minecraft:water[level=3]",
	},
	{
		8, 4, "minecraft:water[level=4]",
	},
	{
		8, 5, "minecraft:water[level=5]",
	},
	{
		8, 6, "minecraft:water[level=6]",
	},
	{
		8, 7, "minecraft:water[level=7]",
	},
	{
		8, 8, "minecraft:water[level=8]",
	},
	{
		8, 9, "minecraft:water[level=9]",
	},
	{
		8, 10, "minecraft:water[level=10]",
	},
	{
		8, 11, "minecraft:water[level=11]",
	},
	{
		8, 12, "minecraft:water[level=12]",
	},
	{
		8, 13, "minecraft:water[level=13]",
	},
	{
		8, 14, "minecraft:water[level=14]",
	},
	{
		8, 15, "minecraft:water[level=15]",
	},
	{
		9, 0, "minecraft:water[level=0]",
	},
	{
		9, 1, "minecraft:water[level=1]",
	},
	{
		9, 2, "minecraft:water[level=2]",
	},
	{
		9, 3, "minecraft:water[level=3]",
	},
	{
		9, 4, "minecraft:water[level=4]",
	},
	{
		9, 5, "minecraft:water[level=5]",
	},
	{
		9, 6, "minecraft:water[level=6]",
	},
	{
		9, 7, "minecraft:water[level=7]",
	},
	{
		9, 8, "minecraft:water[level=8]",
	},
	{
		9, 9, "minecraft:water[level=9]",
	},
	{
		9, 10, "minecraft:water[level=10]",
	},
	{
		9, 11, "minecraft:water[level=11]",
	},
	{
		9, 12, "minecraft:water[level=12]",
	},
	{
		9, 13, "minecraft:water[level=13]",
	},
	{
		9, 14, "minecraft:water[level=14]",
	},
	{
		9, 15, "minecraft:water[level=15]",
	},
	{
		10, 0, "minecraft:lava[level=0]",
	},
	{
		10, 1, "minecraft:lava[level=1]",
	},
	{
		10, 2, "minecraft:lava[level=2]",
	},
	{
		10, 3, "minecraft:lava[level=3]",
	},
	{
		10, 4, "minecraft:lava[level=4]",
	},
	{
		10, 5, "minecraft:lava[level=5]",
	},
	{
		10, 6, "minecraft:lava[level=6]",
	},
	{
		10, 7, "minecraft:lava[level=7]",
	},
	{
		10, 8, "minecraft:lava[level=8]",
	},
	{
		10, 9, "minecraft:lava[level=9]",
	},
	{
		10, 10, "minecraft:lava[level=10]",
	},
	{
		10, 11, "minecraft:lava[level=11]",
	},
	{
		10, 12, "minecraft:lava[level=12]",
	},
	{
		10, 13, "minecraft:lava[level=13]",
	},
	{
		10, 14, "minecraft:lava[level=14]",
	},
	{
		10, 15, "minecraft:lava[level=15]",
	},
	{
		11, 0, "minecraft:lava[level=0]",
	},
	{
		11, 1, "minecraft:lava[level=1]",
	},
	{
		11, 2, "minecraft:lava[level=2]",
	},
	{
		11, 3, "minecraft:lava[level=3]",
	},
	{
		11, 4, "minecraft:lava[level=4]",
	},
	{
		11, 5, "minecraft:lava[level=5]",
	},
	{
		11, 6, "minecraft:lava[level=6]",
	},
	{
		11, 7, "minecraft:lava[level=7]",
	},
	{
		11, 8, "minecraft:lava[level=8]",
	},
	{
		11, 9, "minecraft:lava[level=9]",
	},
	{
		11, 10, "minecraft:lava[level=10]",
	},
	{
		11, 11, "minecraft:lava[level=11]",
	},
	{
		11, 12, "minecraft:lava[level=12]",
	},
	{
		11, 13, "minecraft:lava[level=13]",
	},
	{
		11, 14, "minecraft:lava[level=14]",
	},
	{
		11, 15, "minecraft:lava[level=15]",
	},
	{
		12, 0, "minecraft:sand",
	},
	{
		12, 1, "minecraft:red_sand",
	},
	{
		13, 0, "minecraft:gravel",
	},
	{
		14, 0, "minecraft:gold_ore",
	},
	{
		15, 0, "minecraft:iron_ore",
	},
	{
		16, 0, "minecraft:coal_ore",
	},
	{
		17, 0, "minecraft:oak_log[axis=y]",
	},
	{
		17, 1, "minecraft:spruce_log[axis=y]",
	},
	{
		17, 2, "minecraft:birch_log[axis=y]",
	},
	{
		17, 3, "minecraft:jungle_log[axis=y]",
	},
	{
		17, 4, "minecraft:oak_log[axis=x]",
	},
	{
		17, 5, "minecraft:spruce_log[axis=x]",
	},
	{
		17, 6, "minecraft:birch_log[axis=x]",
	},
	{
		17, 7, "minecraft:jungle_log[axis=x]",
	},
	{
		17, 8, "minecraft:oak_log[axis=z]",
	},
	{
		17, 9, "minecraft:spruce_log[axis=z]",
	},
	{
		17, 10, "minecraft:birch_log[axis=z]",
	},
	{
		17, 11, "minecraft:jungle_log[axis=z]",
	},
	{
		17, 12, "minecraft:oak_wood",
	},
	{
		17, 13, "minecraft:spruce_wood",
	},
	{
		17, 14, "minecraft:birch_wood",
	},
	{
		17, 15, "minecraft:jungle_wood",
	},
	{
		18, 0, "minecraft:oak_leaves[persistent=false,distance=1]",
	},
	{
		18, 1, "minecraft:spruce_leaves[persistent=false,distance=1]",
	},
	{
		18, 2, "minecraft:birch_leaves[persistent=false,distance=1]",
	},
	{
		18, 3, "minecraft:jungle_leaves[persistent=false,distance=1]",
	},
	{
		18, 4, "minecraft:oak_leaves[persistent=true,distance=1]",
	},
	{
		18, 5, "minecraft:spruce_leaves[persistent=true,distance=1]",
	},
	{
		18, 6, "minecraft:birch_leaves[persistent=true,distance=1]",
	},
	{
		18, 7, "minecraft:jungle_leaves[persistent=true,distance=1]",
	},
	{
		18, 8, "minecraft:oak_leaves[persistent=false,distance=1]",
	},
	{
		18, 9, "minecraft:spruce_leaves[persistent=false,distance=1]",
	},
	{
		18, 10, "minecraft:birch_leaves[persistent=false,distance=1]",
	},
	{
		18, 11, "minecraft:jungle_leaves[persistent=false,distance=1]",
	},
	{
		18, 12, "minecraft:oak_leaves[persistent=true,distance=1]",
	},
	{
		18, 13, "minecraft:spruce_leaves[persistent=true,distance=1]",
	},
	{
		18, 14, "minecraft:birch_leaves[persistent=true,distance=1]",
	},
	{
		18, 15, "minecraft:jungle_leaves[persistent=true,distance=1]",
	},
	{
		19, 0, "minecraft:sponge",
	},
	{
		19, 1, "minecraft:wet_sponge",
	},
	{
		20, 0, "minecraft:glass",
	},
	{
		21, 0, "minecraft:lapis_ore",
	},
	{
		22, 0, "minecraft:lapis_block",
	},
	{
		23, 0, "minecraft:dispenser[triggered=false,facing=down]",
	},
	{
		23, 1, "minecraft:dispenser[triggered=false,facing=up]",
	},
	{
		23, 2, "minecraft:dispenser[triggered=false,facing=north]",
	},
	{
		23, 3, "minecraft:dispenser[triggered=false,facing=south]",
	},
	{
		23, 4, "minecraft:dispenser[triggered=false,facing=west]",
	},
	{
		23, 5, "minecraft:dispenser[triggered=false,facing=east]",
	},
	{
		23, 8, "minecraft:dispenser[triggered=true,facing=down]",
	},
	{
		23, 9, "minecraft:dispenser[triggered=true,facing=up]",
	},
	{
		23, 10, "minecraft:dispenser[triggered=true,facing=north]",
	},
	{
		23, 11, "minecraft:dispenser[triggered=true,facing=south]",
	},
	{
		23, 12, "minecraft:dispenser[triggered=true,facing=west]",
	},
	{
		23, 13, "minecraft:dispenser[triggered=true,facing=east]",
	},
	{
		24, 0, "minecraft:sandstone",
	},
	{
		24, 1, "minecraft:chiseled_sandstone",
	},
	{
		24, 2, "minecraft:cut_sandstone",
	},
	{
		25, 0, "minecraft:note_block",
	},
	{
		26, 0, "minecraft:red_bed[part=foot,facing=south,occupied=false]",
	},
	{
		26, 1, "minecraft:red_bed[part=foot,facing=west,occupied=false]",
	},
	{
		26, 2, "minecraft:red_bed[part=foot,facing=north,occupied=false]",
	},
	{
		26, 3, "minecraft:red_bed[part=foot,facing=east,occupied=false]",
	},
	{
		26, 4, "minecraft:red_bed[part=foot,facing=south,occupied=true]",
	},
	{
		26, 5, "minecraft:red_bed[part=foot,facing=west,occupied=true]",
	},
	{
		26, 6, "minecraft:red_bed[part=foot,facing=north,occupied=true]",
	},
	{
		26, 7, "minecraft:red_bed[part=foot,facing=east,occupied=true]",
	},
	{
		26, 8, "minecraft:red_bed[part=head,facing=south,occupied=false]",
	},
	{
		26, 9, "minecraft:red_bed[part=head,facing=west,occupied=false]",
	},
	{
		26, 10, "minecraft:red_bed[part=head,facing=north,occupied=false]",
	},
	{
		26, 11, "minecraft:red_bed[part=head,facing=east,occupied=false]",
	},
	{
		26, 12, "minecraft:red_bed[part=head,facing=south,occupied=true]",
	},
	{
		26, 13, "minecraft:red_bed[part=head,facing=west,occupied=true]",
	},
	{
		26, 14, "minecraft:red_bed[part=head,facing=north,occupied=true]",
	},
	{
		26, 15, "minecraft:red_bed[part=head,facing=east,occupied=true]",
	},
	{
		27, 0, "minecraft:powered_rail[shape=north_south,powered=false]",
	},
	{
		27, 1, "minecraft:powered_rail[shape=east_west,powered=false]",
	},
	{
		27, 2, "minecraft:powered_rail[shape=ascending_east,powered=false]",
	},
	{
		27, 3, "minecraft:powered_rail[shape=ascending_west,powered=false]",
	},
	{
		27, 4, "minecraft:powered_rail[shape=ascending_north,powered=false]",
	},
	{
		27, 5, "minecraft:powered_rail[shape=ascending_south,powered=false]",
	},
	{
		27, 8, "minecraft:powered_rail[shape=north_south,powered=true]",
	},
	{
		27, 9, "minecraft:powered_rail[shape=east_west,powered=true]",
	},
	{
		27, 10, "minecraft:powered_rail[shape=ascending_east,powered=true]",
	},
	{
		27, 11, "minecraft:powered_rail[shape=ascending_west,powered=true]",
	},
	{
		27, 12, "minecraft:powered_rail[shape=ascending_north,powered=true]",
	},
	{
		27, 13, "minecraft:powered_rail[shape=ascending_south,powered=true]",
	},
	{
		28, 0, "minecraft:detector_rail[shape=north_south,powered=false]",
	},
	{
		28, 1, "minecraft:detector_rail[shape=east_west,powered=false]",
	},
	{
		28, 2, "minecraft:detector_rail[shape=ascending_east,powered=false]",
	},
	{
		28, 3, "minecraft:detector_rail[shape=ascending_west,powered=false]",
	},
	{
		28, 4, "minecraft:detector_rail[shape=ascending_north,powered=false]",
	},
	{
		28, 5, "minecraft:detector_rail[shape=ascending_south,powered=false]",
	},
	{
		28, 8, "minecraft:detector_rail[shape=north_south,powered=true]",
	},
	{
		28, 9, "minecraft:detector_rail[shape=east_west,powered=true]",
	},
	{
		28, 10, "minecraft:detector_rail[shape=ascending_east,powered=true]",
	},
	{
		28, 11, "minecraft:detector_rail[shape=ascending_west,powered=true]",
	},
	{
		28, 12, "minecraft:detector_rail[shape=ascending_north,powered=true]",
	},
	{
		28, 13, "minecraft:detector_rail[shape=ascending_south,powered=true]",
	},
	{
		29, 0, "minecraft:sticky_piston[facing=down,extended=false]",
	},
	{
		29, 1, "minecraft:sticky_piston[facing=up,extended=false]",
	},
	{
		29, 2, "minecraft:sticky_piston[facing=north,extended=false]",
	},
	{
		29, 3, "minecraft:sticky_piston[facing=south,extended=false]",
	},
	{
		29, 4, "minecraft:sticky_piston[facing=west,extended=false]",
	},
	{
		29, 5, "minecraft:sticky_piston[facing=east,extended=false]",
	},
	{
		29, 8, "minecraft:sticky_piston[facing=down,extended=true]",
	},
	{
		29, 9, "minecraft:sticky_piston[facing=up,extended=true]",
	},
	{
		29, 10, "minecraft:sticky_piston[facing=north,extended=true]",
	},
	{
		29, 11, "minecraft:sticky_piston[facing=south,extended=true]",
	},
	{
		29, 12, "minecraft:sticky_piston[facing=west,extended=true]",
	},
	{
		29, 13, "minecraft:sticky_piston[facing=east,extended=true]",
	},
	{
		30, 0, "minecraft:cobweb",
	},
	{
		31, 0, "minecraft:dead_bush",
	},
	{
		31, 1, "minecraft:grass",
	},
	{
		31, 2, "minecraft:fern",
	},
	{
		32, 0, "minecraft:dead_bush",
	},
	{
		33, 0, "minecraft:piston[facing=down,extended=false]",
	},
	{
		33, 1, "minecraft:piston[facing=up,extended=false]",
	},
	{
		33, 2, "minecraft:piston[facing=north,extended=false]",
	},
	{
		33, 3, "minecraft:piston[facing=south,extended=false]",
	},
	{
		33, 4, "minecraft:piston[facing=west,extended=false]",
	},
	{
		33, 5, "minecraft:piston[facing=east,extended=false]",
	},
	{
		33, 8, "minecraft:piston[facing=down,extended=true]",
	},
	{
		33, 9, "minecraft:piston[facing=up,extended=true]",
	},
	{
		33, 10, "minecraft:piston[facing=north,extended=true]",
	},
	{
		33, 11, "minecraft:piston[facing=south,extended=true]",
	},
	{
		33, 12, "minecraft:piston[facing=west,extended=true]",
	},
	{
		33, 13, "minecraft:piston[facing=east,extended=true]",
	},
	{
		34, 0, "minecraft:piston_head[short=false,facing=down,type=normal]",
	},
	{
		34, 1, "minecraft:piston_head[short=false,facing=up,type=normal]",
	},
	{
		34, 2, "minecraft:piston_head[short=false,facing=north,type=normal]",
	},
	{
		34, 3, "minecraft:piston_head[short=false,facing=south,type=normal]",
	},
	{
		34, 4, "minecraft:piston_head[short=false,facing=west,type=normal]",
	},
	{
		34, 5, "minecraft:piston_head[short=false,facing=east,type=normal]",
	},
	{
		34, 8, "minecraft:piston_head[short=false,facing=down,type=sticky]",
	},
	{
		34, 9, "minecraft:piston_head[short=false,facing=up,type=sticky]",
	},
	{
		34, 10, "minecraft:piston_head[short=false,facing=north,type=sticky]",
	},
	{
		34, 11, "minecraft:piston_head[short=false,facing=south,type=sticky]",
	},
	{
		34, 12, "minecraft:piston_head[short=false,facing=west,type=sticky]",
	},
	{
		34, 13, "minecraft:piston_head[short=false,facing=east,type=sticky]",
	},
	{
		35, 0, "minecraft:white_wool",
	},
	{
		35, 1, "minecraft:orange_wool",
	},
	{
		35, 2, "minecraft:magenta_wool",
	},
	{
		35, 3, "minecraft:light_blue_wool",
	},
	{
		35, 4, "minecraft:yellow_wool",
	},
	{
		35, 5, "minecraft:lime_wool",
	},
	{
		35, 6, "minecraft:pink_wool",
	},
	{
		35, 7, "minecraft:gray_wool",
	},
	{
		35, 8, "minecraft:light_gray_wool",
	},
	{
		35, 9, "minecraft:cyan_wool",
	},
	{
		35, 10, "minecraft:purple_wool",
	},
	{
		35, 11, "minecraft:blue_wool",
	},
	{
		35, 12, "minecraft:brown_wool",
	},
	{
		35, 13, "minecraft:green_wool",
	},
	{
		35, 14, "minecraft:red_wool",
	},
	{
		35, 15, "minecraft:black_wool",
	},
	{
		36, 0, "minecraft:moving_piston[facing=down,type=normal]",
	},
	{
		36, 1, "minecraft:moving_piston[facing=up,type=normal]",
	},
	{
		36, 2, "minecraft:moving_piston[facing=north,type=normal]",
	},
	{
		36, 3, "minecraft:moving_piston[facing=south,type=normal]",
	},
	{
		36, 4, "minecraft:moving_piston[facing=west,type=normal]",
	},
	{
		36, 5, "minecraft:moving_piston[facing=east,type=normal]",
	},
	{
		36, 8, "minecraft:moving_piston[facing=down,type=sticky]",
	},
	{
		36, 9, "minecraft:moving_piston[facing=up,type=sticky]",
	},
	{
		36, 10, "minecraft:moving_piston[facing=north,type=sticky]",
	},
	{
		36, 11, "minecraft:moving_piston[facing=south,type=sticky]",
	},
	{
		36, 12, "minecraft:moving_piston[facing=west,type=sticky]",
	},
	{
		36, 13, "minecraft:moving_piston[facing=east,type=sticky]",
	},
	{
		37, 0, "minecraft:dandelion",
	},
	{
		38, 0, "minecraft:poppy",
	},
	{
		38, 1, "minecraft:blue_orchid",
	},
	{
		38, 2, "minecraft:allium",
	},
	{
		38, 3, "minecraft:azure_bluet",
	},
	{
		38, 4, "minecraft:red_tulip",
	},
	{
		38, 5, "minecraft:orange_tulip",
	},
	{
		38, 6, "minecraft:white_tulip",
	},
	{
		38, 7, "minecraft:pink_tulip",
	},
	{
		38, 8, "minecraft:oxeye_daisy",
	},
	{
		39, 0, "minecraft:brown_mushroom",
	},
	{
		40, 0, "minecraft:red_mushroom",
	},
	{
		41, 0, "minecraft:gold_block",
	},
	{
		42, 0, "minecraft:iron_block",
	},
	{
		43, 0, "minecraft:stone_slab[type=double]",
	},
	{
		43, 1, "minecraft:sandstone_slab[type=double]",
	},
	{
		43, 2, "minecraft:petrified_oak_slab[type=double]",
	},
	{
		43, 3, "minecraft:cobblestone_slab[type=double]",
	},
	{
		43, 4, "minecraft:brick_slab[type=double]",
	},
	{
		43, 5, "minecraft:stone_brick_slab[type=double]",
	},
	{
		43, 6, "minecraft:nether_brick_slab[type=double]",
	},
	{
		43, 7, "minecraft:quartz_slab[type=double]",
	},
	{
		43, 8, "minecraft:smooth_stone",
	},
	{
		43, 9, "minecraft:smooth_sandstone",
	},
	{
		43, 10, "minecraft:petrified_oak_slab[type=double]",
	},
	{
		43, 11, "minecraft:cobblestone_slab[type=double]",
	},
	{
		43, 12, "minecraft:brick_slab[type=double]",
	},
	{
		43, 13, "minecraft:stone_brick_slab[type=double]",
	},
	{
		43, 14, "minecraft:nether_brick_slab[type=double]",
	},
	{
		43, 15, "minecraft:smooth_quartz",
	},
	{
		44, 0, "minecraft:stone_slab[type=bottom]",
	},
	{
		44, 1, "minecraft:sandstone_slab[type=bottom]",
	},
	{
		44, 2, "minecraft:petrified_oak_slab[type=bottom]",
	},
	{
		44, 3, "minecraft:cobblestone_slab[type=bottom]",
	},
	{
		44, 4, "minecraft:brick_slab[type=bottom]",
	},
	{
		44, 5, "minecraft:stone_brick_slab[type=bottom]",
	},
	{
		44, 6, "minecraft:nether_brick_slab[type=bottom]",
	},
	{
		44, 7, "minecraft:quartz_slab[type=bottom]",
	},
	{
		44, 8, "minecraft:stone_slab[type=top]",
	},
	{
		44, 9, "minecraft:sandstone_slab[type=top]",
	},
	{
		44, 10, "minecraft:petrified_oak_slab[type=top]",
	},
	{
		44, 11, "minecraft:cobblestone_slab[type=top]",
	},
	{
		44, 12, "minecraft:brick_slab[type=top]",
	},
	{
		44, 13, "minecraft:stone_brick_slab[type=top]",
	},
	{
		44, 14, "minecraft:nether_brick_slab[type=top]",
	},
	{
		44, 15, "minecraft:quartz_slab[type=top]",
	},
	{
		45, 0, "minecraft:bricks",
	},
	{
		46, 0, "minecraft:tnt[unstable=false]",
	},
	{
		46, 1, "minecraft:tnt[unstable=true]",
	},
	{
		47, 0, "minecraft:bookshelf",
	},
	{
		48, 0, "minecraft:mossy_cobblestone",
	},
	{
		49, 0, "minecraft:obsidian",
	},
	{
		50, 1, "minecraft:wall_torch[facing=east]",
	},
	{
		50, 2, "minecraft:wall_torch[facing=west]",
	},
	{
		50, 3, "minecraft:wall_torch[facing=south]",
	},
	{
		50, 4, "minecraft:wall_torch[facing=north]",
	},
	{
		50, 5, "minecraft:torch",
	},
	{
		51, 0, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=0]",
	},
	{
		51, 1, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=1]",
	},
	{
		51, 2, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=2]",
	},
	{
		51, 3, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=3]",
	},
	{
		51, 4, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=4]",
	},
	{
		51, 5, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=5]",
	},
	{
		51, 6, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=6]",
	},
	{
		51, 7, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=7]",
	},
	{
		51, 8, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=8]",
	},
	{
		51, 9, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=9]",
	},
	{
		51, 10, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=10]",
	},
	{
		51, 11, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=11]",
	},
	{
		51, 12, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=12]",
	},
	{
		51, 13, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=13]",
	},
	{
		51, 14, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=14]",
	},
	{
		51, 15, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=15]",
	},
	{
		52, 0, "minecraft:spawner",
	},
	{
		53, 0, "minecraft:oak_stairs[half=bottom,shape=outer_right,facing=east]",
	},
	{
		53, 1, "minecraft:oak_stairs[half=bottom,shape=outer_right,facing=west]",
	},
	{
		53, 2, "minecraft:oak_stairs[half=bottom,shape=outer_right,facing=south]",
	},
	{
		53, 3, "minecraft:oak_stairs[half=bottom,shape=outer_right,facing=north]",
	},
	{
		53, 4, "minecraft:oak_stairs[half=top,shape=outer_right,facing=east]",
	},
	{
		53, 5, "minecraft:oak_stairs[half=top,shape=outer_right,facing=west]",
	},
	{
		53, 6, "minecraft:oak_stairs[half=top,shape=outer_right,facing=south]",
	},
	{
		53, 7, "minecraft:oak_stairs[half=top,shape=outer_right,facing=north]",
	},
	{
		54, 2, "minecraft:chest[facing=north,type=single]",
	},
	{
		54, 3, "minecraft:chest[facing=south,type=single]",
	},
	{
		54, 4, "minecraft:chest[facing=west,type=single]",
	},
	{
		54, 5, "minecraft:chest[facing=east,type=single]",
	},
	{
		55, 0, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=0]",
	},
	{
		55, 1, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=1]",
	},
	{
		55, 2, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=2]",
	},
	{
		55, 3, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=3]",
	},
	{
		55, 4, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=4]",
	},
	{
		55, 5, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=5]",
	},
	{
		55, 6, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=6]",
	},
	{
		55, 7, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=7]",
	},
	{
		55, 8, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=8]",
	},
	{
		55, 9, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=9]",
	},
	{
		55, 10, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=10]",
	},
	{
		55, 11, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=11]",
	},
	{
		55, 12, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=12]",
	},
	{
		55, 13, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=13]",
	},
	{
		55, 14, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=14]",
	},
	{
		55, 15, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=15]",
	},
	{
		56, 0, "minecraft:diamond_ore",
	},
	{
		57, 0, "minecraft:diamond_block",
	},
	{
		58, 0, "minecraft:crafting_table",
	},
	{
		59, 0, "minecraft:wheat[age=0]",
	},
	{
		59, 1, "minecraft:wheat[age=1]",
	},
	{
		59, 2, "minecraft:wheat[age=2]",
	},
	{
		59, 3, "minecraft:wheat[age=3]",
	},
	{
		59, 4, "minecraft:wheat[age=4]",
	},
	{
		59, 5, "minecraft:wheat[age=5]",
	},
	{
		59, 6, "minecraft:wheat[age=6]",
	},
	{
		59, 7, "minecraft:wheat[age=7]",
	},
	{
		60, 0, "minecraft:farmland[moisture=0]",
	},
	{
		60, 1, "minecraft:farmland[moisture=1]",
	},
	{
		60, 2, "minecraft:farmland[moisture=2]",
	},
	{
		60, 3, "minecraft:farmland[moisture=3]",
	},
	{
		60, 4, "minecraft:farmland[moisture=4]",
	},
	{
		60, 5, "minecraft:farmland[moisture=5]",
	},
	{
		60, 6, "minecraft:farmland[moisture=6]",
	},
	{
		60, 7, "minecraft:farmland[moisture=7]",
	},
	{
		61, 2, "minecraft:furnace[facing=north,lit=false]",
	},
	{
		61, 3, "minecraft:furnace[facing=south,lit=false]",
	},
	{
		61, 4, "minecraft:furnace[facing=west,lit=false]",
	},
	{
		61, 5, "minecraft:furnace[facing=east,lit=false]",
	},
	{
		62, 2, "minecraft:furnace[facing=north,lit=true]",
	},
	{
		62, 3, "minecraft:furnace[facing=south,lit=true]",
	},
	{
		62, 4, "minecraft:furnace[facing=west,lit=true]",
	},
	{
		62, 5, "minecraft:furnace[facing=east,lit=true]",
	},
	{
		63, 0, "minecraft:sign[rotation=0]",
	},
	{
		63, 1, "minecraft:sign[rotation=1]",
	},
	{
		63, 2, "minecraft:sign[rotation=2]",
	},
	{
		63, 3, "minecraft:sign[rotation=3]",
	},
	{
		63, 4, "minecraft:sign[rotation=4]",
	},
	{
		63, 5, "minecraft:sign[rotation=5]",
	},
	{
		63, 6, "minecraft:sign[rotation=6]",
	},
	{
		63, 7, "minecraft:sign[rotation=7]",
	},
	{
		63, 8, "minecraft:sign[rotation=8]",
	},
	{
		63, 9, "minecraft:sign[rotation=9]",
	},
	{
		63, 10, "minecraft:sign[rotation=10]",
	},
	{
		63, 11, "minecraft:sign[rotation=11]",
	},
	{
		63, 12, "minecraft:sign[rotation=12]",
	},
	{
		63, 13, "minecraft:sign[rotation=13]",
	},
	{
		63, 14, "minecraft:sign[rotation=14]",
	},
	{
		63, 15, "minecraft:sign[rotation=15]",
	},
	{
		64, 0, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
	},
	{
		64, 1, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
	},
	{
		64, 2, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
	},
	{
		64, 3, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
	},
	{
		64, 4, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
	},
	{
		64, 5, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
	},
	{
		64, 6, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
	},
	{
		64, 7, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
	},
	{
		64, 8, "minecraft:oak_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
	},
	{
		64, 9, "minecraft:oak_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
	},
	{
		64, 10, "minecraft:oak_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
	},
	{
		64, 11, "minecraft:oak_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
	},
	{
		65, 2, "minecraft:ladder[facing=north]",
	},
	{
		65, 3, "minecraft:ladder[facing=south]",
	},
	{
		65, 4, "minecraft:ladder[facing=west]",
	},
	{
		65, 5, "minecraft:ladder[facing=east]",
	},
	{
		66, 0, "minecraft:rail[shape=north_south]",
	},
	{
		66, 1, "minecraft:rail[shape=east_west]",
	},
	{
		66, 2, "minecraft:rail[shape=ascending_east]",
	},
	{
		66, 3, "minecraft:rail[shape=ascending_west]",
	},
	{
		66, 4, "minecraft:rail[shape=ascending_north]",
	},
	{
		66, 5, "minecraft:rail[shape=ascending_south]",
	},
	{
		66, 6, "minecraft:rail[shape=south_east]",
	},
	{
		66, 7, "minecraft:rail[shape=south_west]",
	},
	{
		66, 8, "minecraft:rail[shape=north_west]",
	},
	{
		66, 9, "minecraft:rail[shape=north_east]",
	},
	{
		67, 0, "minecraft:cobblestone_stairs[half=bottom,shape=straight,facing=east]",
	},
	{
		67, 1, "minecraft:cobblestone_stairs[half=bottom,shape=straight,facing=west]",
	},
	{
		67, 2, "minecraft:cobblestone_stairs[half=bottom,shape=straight,facing=south]",
	},
	{
		67, 3, "minecraft:cobblestone_stairs[half=bottom,shape=straight,facing=north]",
	},
	{
		67, 4, "minecraft:cobblestone_stairs[half=top,shape=straight,facing=east]",
	},
	{
		67, 5, "minecraft:cobblestone_stairs[half=top,shape=straight,facing=west]",
	},
	{
		67, 6, "minecraft:cobblestone_stairs[half=top,shape=straight,facing=south]",
	},
	{
		67, 7, "minecraft:cobblestone_stairs[half=top,shape=straight,facing=north]",
	},
	{
		68, 2, "minecraft:wall_sign[facing=north]",
	},
	{
		68, 3, "minecraft:wall_sign[facing=south]",
	},
	{
		68, 4, "minecraft:wall_sign[facing=west]",
	},
	{
		68, 5, "minecraft:wall_sign[facing=east]",
	},
	{
		69, 0, "minecraft:lever[powered=false,facing=north,face=ceiling]",
	},
	{
		69, 1, "minecraft:lever[powered=false,facing=east,face=wall]",
	},
	{
		69, 2, "minecraft:lever[powered=false,facing=west,face=wall]",
	},
	{
		69, 3, "minecraft:lever[powered=false,facing=south,face=wall]",
	},
	{
		69, 4, "minecraft:lever[powered=false,facing=north,face=wall]",
	},
	{
		69, 5, "minecraft:lever[powered=false,facing=east,face=floor]",
	},
	{
		69, 6, "minecraft:lever[powered=false,facing=north,face=floor]",
	},
	{
		69, 7, "minecraft:lever[powered=false,facing=east,face=ceiling]",
	},
	{
		69, 8, "minecraft:lever[powered=true,facing=north,face=ceiling]",
	},
	{
		69, 9, "minecraft:lever[powered=true,facing=east,face=wall]",
	},
	{
		69, 10, "minecraft:lever[powered=true,facing=west,face=wall]",
	},
	{
		69, 11, "minecraft:lever[powered=true,facing=south,face=wall]",
	},
	{
		69, 12, "minecraft:lever[powered=true,facing=north,face=wall]",
	},
	{
		69, 13, "minecraft:lever[powered=true,facing=east,face=floor]",
	},
	{
		69, 14, "minecraft:lever[powered=true,facing=north,face=floor]",
	},
	{
		69, 15, "minecraft:lever[powered=true,facing=east,face=ceiling]",
	},
	{
		70, 0, "minecraft:stone_pressure_plate[powered=false]",
	},
	{
		70, 1, "minecraft:stone_pressure_plate[powered=true]",
	},
	{
		71, 0, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
	},
	{
		71, 1, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
	},
	{
		71, 2, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
	},
	{
		71, 3, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
	},
	{
		71, 4, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
	},
	{
		71, 5, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
	},
	{
		71, 6, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
	},
	{
		71, 7, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
	},
	{
		71, 8, "minecraft:iron_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
	},
	{
		71, 9, "minecraft:iron_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
	},
	{
		71, 10, "minecraft:iron_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
	},
	{
		71, 11, "minecraft:iron_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
	},
	{
		72, 0, "minecraft:oak_pressure_plate[powered=false]",
	},
	{
		72, 1, "minecraft:oak_pressure_plate[powered=true]",
	},
	{
		73, 0, "minecraft:redstone_ore[lit=false]",
	},
	{
		74, 0, "minecraft:redstone_ore[lit=true]",
	},
	{
		75, 1, "minecraft:redstone_wall_torch[facing=east,lit=false]",
	},
	{
		75, 2, "minecraft:redstone_wall_torch[facing=west,lit=false]",
	},
	{
		75, 3, "minecraft:redstone_wall_torch[facing=south,lit=false]",
	},
	{
		75, 4, "minecraft:redstone_wall_torch[facing=north,lit=false]",
	},
	{
		75, 5, "minecraft:redstone_torch[lit=false]",
	},
	{
		76, 1, "minecraft:redstone_wall_torch[facing=east,lit=true]",
	},
	{
		76, 2, "minecraft:redstone_wall_torch[facing=west,lit=true]",
	},
	{
		76, 3, "minecraft:redstone_wall_torch[facing=south,lit=true]",
	},
	{
		76, 4, "minecraft:redstone_wall_torch[facing=north,lit=true]",
	},
	{
		76, 5, "minecraft:redstone_torch[lit=true]",
	},
	{
		77, 0, "minecraft:stone_button[powered=false,facing=east,face=ceiling]",
	},
	{
		77, 1, "minecraft:stone_button[powered=false,facing=east,face=wall]",
	},
	{
		77, 2, "minecraft:stone_button[powered=false,facing=west,face=wall]",
	},
	{
		77, 3, "minecraft:stone_button[powered=false,facing=south,face=wall]",
	},
	{
		77, 4, "minecraft:stone_button[powered=false,facing=north,face=wall]",
	},
	{
		77, 5, "minecraft:stone_button[powered=false,facing=east,face=floor]",
	},
	{
		77, 8, "minecraft:stone_button[powered=true,facing=south,face=ceiling]",
	},
	{
		77, 9, "minecraft:stone_button[powered=true,facing=east,face=wall]",
	},
	{
		77, 10, "minecraft:stone_button[powered=true,facing=west,face=wall]",
	},
	{
		77, 11, "minecraft:stone_button[powered=true,facing=south,face=wall]",
	},
	{
		77, 12, "minecraft:stone_button[powered=true,facing=north,face=wall]",
	},
	{
		77, 13, "minecraft:stone_button[powered=true,facing=south,face=floor]",
	},
	{
		78, 0, "minecraft:snow[layers=1]",
	},
	{
		78, 1, "minecraft:snow[layers=2]",
	},
	{
		78, 2, "minecraft:snow[layers=3]",
	},
	{
		78, 3, "minecraft:snow[layers=4]",
	},
	{
		78, 4, "minecraft:snow[layers=5]",
	},
	{
		78, 5, "minecraft:snow[layers=6]",
	},
	{
		78, 6, "minecraft:snow[layers=7]",
	},
	{
		78, 7, "minecraft:snow[layers=8]",
	},
	{
		79, 0, "minecraft:ice",
	},
	{
		80, 0, "minecraft:snow_block",
	},
	{
		81, 0, "minecraft:cactus[age=0]",
	},
	{
		81, 1, "minecraft:cactus[age=1]",
	},
	{
		81, 2, "minecraft:cactus[age=2]",
	},
	{
		81, 3, "minecraft:cactus[age=3]",
	},
	{
		81, 4, "minecraft:cactus[age=4]",
	},
	{
		81, 5, "minecraft:cactus[age=5]",
	},
	{
		81, 6, "minecraft:cactus[age=6]",
	},
	{
		81, 7, "minecraft:cactus[age=7]",
	},
	{
		81, 8, "minecraft:cactus[age=8]",
	},
	{
		81, 9, "minecraft:cactus[age=9]",
	},
	{
		81, 10, "minecraft:cactus[age=10]",
	},
	{
		81, 11, "minecraft:cactus[age=11]",
	},
	{
		81, 12, "minecraft:cactus[age=12]",
	},
	{
		81, 13, "minecraft:cactus[age=13]",
	},
	{
		81, 14, "minecraft:cactus[age=14]",
	},
	{
		81, 15, "minecraft:cactus[age=15]",
	},
	{
		82, 0, "minecraft:clay",
	},
	{
		83, 0, "minecraft:sugar_cane[age=0]",
	},
	{
		83, 1, "minecraft:sugar_cane[age=1]",
	},
	{
		83, 2, "minecraft:sugar_cane[age=2]",
	},
	{
		83, 3, "minecraft:sugar_cane[age=3]",
	},
	{
		83, 4, "minecraft:sugar_cane[age=4]",
	},
	{
		83, 5, "minecraft:sugar_cane[age=5]",
	},
	{
		83, 6, "minecraft:sugar_cane[age=6]",
	},
	{
		83, 7, "minecraft:sugar_cane[age=7]",
	},
	{
		83, 8, "minecraft:sugar_cane[age=8]",
	},
	{
		83, 9, "minecraft:sugar_cane[age=9]",
	},
	{
		83, 10, "minecraft:sugar_cane[age=10]",
	},
	{
		83, 11, "minecraft:sugar_cane[age=11]",
	},
	{
		83, 12, "minecraft:sugar_cane[age=12]",
	},
	{
		83, 13, "minecraft:sugar_cane[age=13]",
	},
	{
		83, 14, "minecraft:sugar_cane[age=14]",
	},
	{
		83, 15, "minecraft:sugar_cane[age=15]",
	},
	{
		84, 0, "minecraft:jukebox[has_record=false]",
	},
	{
		84, 1, "minecraft:jukebox[has_record=true]",
	},
	{
		85, 0, "minecraft:oak_fence[east=false,south=false,north=false,west=false]",
	},
	{
		86, 0, "minecraft:carved_pumpkin[facing=south]",
	},
	{
		86, 1, "minecraft:carved_pumpkin[facing=west]",
	},
	{
		86, 2, "minecraft:carved_pumpkin[facing=north]",
	},
	{
		86, 3, "minecraft:carved_pumpkin[facing=east]",
	},
	{
		87, 0, "minecraft:netherrack",
	},
	{
		88, 0, "minecraft:soul_sand",
	},
	{
		89, 0, "minecraft:glowstone",
	},
	{
		90, 1, "minecraft:nether_portal[axis=x]",
	},
	{
		90, 2, "minecraft:nether_portal[axis=z]",
	},
	{
		91, 0, "minecraft:jack_o_lantern[facing=south]",
	},
	{
		91, 1, "minecraft:jack_o_lantern[facing=west]",
	},
	{
		91, 2, "minecraft:jack_o_lantern[facing=north]",
	},
	{
		91, 3, "minecraft:jack_o_lantern[facing=east]",
	},
	{
		92, 0, "minecraft:cake[bites=0]",
	},
	{
		92, 1, "minecraft:cake[bites=1]",
	},
	{
		92, 2, "minecraft:cake[bites=2]",
	},
	{
		92, 3, "minecraft:cake[bites=3]",
	},
	{
		92, 4, "minecraft:cake[bites=4]",
	},
	{
		92, 5, "minecraft:cake[bites=5]",
	},
	{
		92, 6, "minecraft:cake[bites=6]",
	},
	{
		93, 0, "minecraft:repeater[delay=1,facing=south,locked=false,powered=false]",
	},
	{
		93, 1, "minecraft:repeater[delay=1,facing=west,locked=false,powered=false]",
	},
	{
		93, 2, "minecraft:repeater[delay=1,facing=north,locked=false,powered=false]",
	},
	{
		93, 3, "minecraft:repeater[delay=1,facing=east,locked=false,powered=false]",
	},
	{
		93, 4, "minecraft:repeater[delay=2,facing=south,locked=false,powered=false]",
	},
	{
		93, 5, "minecraft:repeater[delay=2,facing=west,locked=false,powered=false]",
	},
	{
		93, 6, "minecraft:repeater[delay=2,facing=north,locked=false,powered=false]",
	},
	{
		93, 7, "minecraft:repeater[delay=2,facing=east,locked=false,powered=false]",
	},
	{
		93, 8, "minecraft:repeater[delay=3,facing=south,locked=false,powered=false]",
	},
	{
		93, 9, "minecraft:repeater[delay=3,facing=west,locked=false,powered=false]",
	},
	{
		93, 10, "minecraft:repeater[delay=3,facing=north,locked=false,powered=false]",
	},
	{
		93, 11, "minecraft:repeater[delay=3,facing=east,locked=false,powered=false]",
	},
	{
		93, 12, "minecraft:repeater[delay=4,facing=south,locked=false,powered=false]",
	},
	{
		93, 13, "minecraft:repeater[delay=4,facing=west,locked=false,powered=false]",
	},
	{
		93, 14, "minecraft:repeater[delay=4,facing=north,locked=false,powered=false]",
	},
	{
		93, 15, "minecraft:repeater[delay=4,facing=east,locked=false,powered=false]",
	},
	{
		94, 0, "minecraft:repeater[delay=1,facing=south,locked=false,powered=true]",
	},
	{
		94, 1, "minecraft:repeater[delay=1,facing=west,locked=false,powered=true]",
	},
	{
		94, 2, "minecraft:repeater[delay=1,facing=north,locked=false,powered=true]",
	},
	{
		94, 3, "minecraft:repeater[delay=1,facing=east,locked=false,powered=true]",
	},
	{
		94, 4, "minecraft:repeater[delay=2,facing=south,locked=false,powered=true]",
	},
	{
		94, 5, "minecraft:repeater[delay=2,facing=west,locked=false,powered=true]",
	},
	{
		94, 6, "minecraft:repeater[delay=2,facing=north,locked=false,powered=true]",
	},
	{
		94, 7, "minecraft:repeater[delay=2,facing=east,locked=false,powered=true]",
	},
	{
		94, 8, "minecraft:repeater[delay=3,facing=south,locked=false,powered=true]",
	},
	{
		94, 9, "minecraft:repeater[delay=3,facing=west,locked=false,powered=true]",
	},
	{
		94, 10, "minecraft:repeater[delay=3,facing=north,locked=false,powered=true]",
	},
	{
		94, 11, "minecraft:repeater[delay=3,facing=east,locked=false,powered=true]",
	},
	{
		94, 12, "minecraft:repeater[delay=4,facing=south,locked=false,powered=true]",
	},
	{
		94, 13, "minecraft:repeater[delay=4,facing=west,locked=false,powered=true]",
	},
	{
		94, 14, "minecraft:repeater[delay=4,facing=north,locked=false,powered=true]",
	},
	{
		94, 15, "minecraft:repeater[delay=4,facing=east,locked=false,powered=true]",
	},
	{
		95, 0, "minecraft:white_stained_glass",
	},
	{
		95, 1, "minecraft:orange_stained_glass",
	},
	{
		95, 2, "minecraft:magenta_stained_glass",
	},
	{
		95, 3, "minecraft:light_blue_stained_glass",
	},
	{
		95, 4, "minecraft:yellow_stained_glass",
	},
	{
		95, 5, "minecraft:lime_stained_glass",
	},
	{
		95, 6, "minecraft:pink_stained_glass",
	},
	{
		95, 7, "minecraft:gray_stained_glass",
	},
	{
		95, 8, "minecraft:light_gray_stained_glass",
	},
	{
		95, 9, "minecraft:cyan_stained_glass",
	},
	{
		95, 10, "minecraft:purple_stained_glass",
	},
	{
		95, 11, "minecraft:blue_stained_glass",
	},
	{
		95, 12, "minecraft:brown_stained_glass",
	},
	{
		95, 13, "minecraft:green_stained_glass",
	},
	{
		95, 14, "minecraft:red_stained_glass",
	},
	{
		95, 15, "minecraft:black_stained_glass",
	},
	{
		96, 0, "minecraft:oak_trapdoor[half=bottom,facing=north,open=false,powered=false]",
	},
	{
		96, 1, "minecraft:oak_trapdoor[half=bottom,facing=south,open=false,powered=false]",
	},
	{
		96, 2, "minecraft:oak_trapdoor[half=bottom,facing=west,open=false,powered=false]",
	},
	{
		96, 3, "minecraft:oak_trapdoor[half=bottom,facing=east,open=false,powered=false]",
	},
	{
		96, 4, "minecraft:oak_trapdoor[half=bottom,facing=north,open=true,powered=true]",
	},
	{
		96, 5, "minecraft:oak_trapdoor[half=bottom,facing=south,open=true,powered=true]",
	},
	{
		96, 6, "minecraft:oak_trapdoor[half=bottom,facing=west,open=true,powered=true]",
	},
	{
		96, 7, "minecraft:oak_trapdoor[half=bottom,facing=east,open=true,powered=true]",
	},
	{
		96, 8, "minecraft:oak_trapdoor[half=top,facing=north,open=false,powered=false]",
	},
	{
		96, 9, "minecraft:oak_trapdoor[half=top,facing=south,open=false,powered=false]",
	},
	{
		96, 10, "minecraft:oak_trapdoor[half=top,facing=west,open=false,powered=false]",
	},
	{
		96, 11, "minecraft:oak_trapdoor[half=top,facing=east,open=false,powered=false]",
	},
	{
		96, 12, "minecraft:oak_trapdoor[half=top,facing=north,open=true,powered=true]",
	},
	{
		96, 13, "minecraft:oak_trapdoor[half=top,facing=south,open=true,powered=true]",
	},
	{
		96, 14, "minecraft:oak_trapdoor[half=top,facing=west,open=true,powered=true]",
	},
	{
		96, 15, "minecraft:oak_trapdoor[half=top,facing=east,open=true,powered=true]",
	},
	{
		97, 0, "minecraft:infested_stone",
	},
	{
		97, 1, "minecraft:infested_cobblestone",
	},
	{
		97, 2, "minecraft:infested_stone_bricks",
	},
	{
		97, 3, "minecraft:infested_mossy_stone_bricks",
	},
	{
		97, 4, "minecraft:infested_cracked_stone_bricks",
	},
	{
		97, 5, "minecraft:infested_chiseled_stone_bricks",
	},
	{
		98, 0, "minecraft:stone_bricks",
	},
	{
		98, 1, "minecraft:mossy_stone_bricks",
	},
	{
		98, 2, "minecraft:cracked_stone_bricks",
	},
	{
		98, 3, "minecraft:chiseled_stone_bricks",
	},
	{
		99, 0, "minecraft:brown_mushroom_block[north=false,east=false,south=false,west=false,up=false,down=false]",
	},
	{
		99, 1, "minecraft:brown_mushroom_block[north=true,east=false,south=false,west=true,up=true,down=false]",
	},
	{
		99, 2, "minecraft:brown_mushroom_block[north=true,east=false,south=false,west=false,up=true,down=false]",
	},
	{
		99, 3, "minecraft:brown_mushroom_block[north=true,east=true,south=false,west=false,up=true,down=false]",
	},
	{
		99, 4, "minecraft:brown_mushroom_block[north=false,east=false,south=false,west=true,up=true,down=false]",
	},
	{
		99, 5, "minecraft:brown_mushroom_block[north=false,east=false,south=false,west=false,up=true,down=false]",
	},
	{
		99, 6, "minecraft:brown_mushroom_block[north=false,east=true,south=false,west=false,up=true,down=false]",
	},
	{
		99, 7, "minecraft:brown_mushroom_block[north=false,east=false,south=true,west=true,up=true,down=false]",
	},
	{
		99, 8, "minecraft:brown_mushroom_block[north=false,east=false,south=true,west=false,up=true,down=false]",
	},
	{
		99, 9, "minecraft:brown_mushroom_block[north=false,east=true,south=true,west=false,up=true,down=false]",
	},
	{
		99, 10, "minecraft:mushroom_stem[north=true,east=true,south=true,west=true,up=false,down=false]",
	},
	{
		99, 14, "minecraft:brown_mushroom_block[north=true,east=true,south=true,west=true,up=true,down=true]",
	},
	{
		99, 15, "minecraft:mushroom_stem[north=true,east=true,south=true,west=true,up=true,down=true]",
	},
	{
		100, 0, "minecraft:red_mushroom_block[north=false,east=false,south=false,west=false,up=false,down=false]",
	},
	{
		100, 1, "minecraft:red_mushroom_block[north=true,east=false,south=false,west=true,up=true,down=false]",
	},
	{
		100, 2, "minecraft:red_mushroom_block[north=true,east=false,south=false,west=false,up=true,down=false]",
	},
	{
		100, 3, "minecraft:red_mushroom_block[north=true,east=true,south=false,west=false,up=true,down=false]",
	},
	{
		100, 4, "minecraft:red_mushroom_block[north=false,east=false,south=false,west=true,up=true,down=false]",
	},
	{
		100, 5, "minecraft:red_mushroom_block[north=false,east=false,south=false,west=false,up=true,down=false]",
	},
	{
		100, 6, "minecraft:red_mushroom_block[north=false,east=true,south=false,west=false,up=true,down=false]",
	},
	{
		100, 7, "minecraft:red_mushroom_block[north=false,east=false,south=true,west=true,up=true,down=false]",
	},
	{
		100, 8, "minecraft:red_mushroom_block[north=false,east=false,south=true,west=false,up=true,down=false]",
	},
	{
		100, 9, "minecraft:red_mushroom_block[north=false,east=true,south=true,west=false,up=true,down=false]",
	},
	{
		100, 10, "minecraft:mushroom_stem[north=true,east=true,south=true,west=true,up=false,down=false]",
	},
	{
		100, 14, "minecraft:red_mushroom_block[north=true,east=true,south=true,west=true,up=true,down=true]",
	},
	{
		100, 15, "minecraft:mushroom_stem[north=true,east=true,south=true,west=true,up=true,down=true]",
	},
	{
		101, 0, "minecraft:iron_bars[east=false,south=false,north=false,west=false]",
	},
	{
		102, 0, "minecraft:glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		103, 0, "minecraft:melon",
	},
	{
		104, 0, "minecraft:pumpkin_stem[age=0]",
	},
	{
		104, 1, "minecraft:pumpkin_stem[age=1]",
	},
	{
		104, 2, "minecraft:pumpkin_stem[age=2]",
	},
	{
		104, 3, "minecraft:pumpkin_stem[age=3]",
	},
	{
		104, 4, "minecraft:pumpkin_stem[age=4]",
	},
	{
		104, 5, "minecraft:pumpkin_stem[age=5]",
	},
	{
		104, 6, "minecraft:pumpkin_stem[age=6]",
	},
	{
		104, 7, "minecraft:pumpkin_stem[age=7]",
	},
	{
		105, 0, "minecraft:melon_stem[age=0]",
	},
	{
		105, 1, "minecraft:melon_stem[age=1]",
	},
	{
		105, 2, "minecraft:melon_stem[age=2]",
	},
	{
		105, 3, "minecraft:melon_stem[age=3]",
	},
	{
		105, 4, "minecraft:melon_stem[age=4]",
	},
	{
		105, 5, "minecraft:melon_stem[age=5]",
	},
	{
		105, 6, "minecraft:melon_stem[age=6]",
	},
	{
		105, 7, "minecraft:melon_stem[age=7]",
	},
	{
		106, 0, "minecraft:vine[east=false,south=false,north=false,west=false,up=false]",
	},
	{
		106, 1, "minecraft:vine[east=false,south=true,north=false,west=false,up=false]",
	},
	{
		106, 2, "minecraft:vine[east=false,south=false,north=false,west=true,up=false]",
	},
	{
		106, 3, "minecraft:vine[east=false,south=true,north=false,west=true,up=false]",
	},
	{
		106, 4, "minecraft:vine[east=false,south=false,north=true,west=false,up=false]",
	},
	{
		106, 5, "minecraft:vine[east=false,south=true,north=true,west=false,up=false]",
	},
	{
		106, 6, "minecraft:vine[east=false,south=false,north=true,west=true,up=false]",
	},
	{
		106, 7, "minecraft:vine[east=false,south=true,north=true,west=true,up=false]",
	},
	{
		106, 8, "minecraft:vine[east=true,south=false,north=false,west=false,up=false]",
	},
	{
		106, 9, "minecraft:vine[east=true,south=true,north=false,west=false,up=false]",
	},
	{
		106, 10, "minecraft:vine[east=true,south=false,north=false,west=true,up=false]",
	},
	{
		106, 11, "minecraft:vine[east=true,south=true,north=false,west=true,up=false]",
	},
	{
		106, 12, "minecraft:vine[east=true,south=false,north=true,west=false,up=false]",
	},
	{
		106, 13, "minecraft:vine[east=true,south=true,north=true,west=false,up=false]",
	},
	{
		106, 14, "minecraft:vine[east=true,south=false,north=true,west=true,up=false]",
	},
	{
		106, 15, "minecraft:vine[east=true,south=true,north=true,west=true,up=false]",
	},
	{
		107, 0, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=south,open=false]",
	},
	{
		107, 1, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=west,open=false]",
	},
	{
		107, 2, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=north,open=false]",
	},
	{
		107, 3, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=east,open=false]",
	},
	{
		107, 4, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=south,open=true]",
	},
	{
		107, 5, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=west,open=true]",
	},
	{
		107, 6, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=north,open=true]",
	},
	{
		107, 7, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=east,open=true]",
	},
	{
		107, 8, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=south,open=false]",
	},
	{
		107, 9, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=west,open=false]",
	},
	{
		107, 10, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=north,open=false]",
	},
	{
		107, 11, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=east,open=false]",
	},
	{
		107, 12, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=south,open=true]",
	},
	{
		107, 13, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=west,open=true]",
	},
	{
		107, 14, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=north,open=true]",
	},
	{
		107, 15, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=east,open=true]",
	},
	{
		108, 0, "minecraft:brick_stairs[half=bottom,shape=straight,facing=east]",
	},
	{
		108, 1, "minecraft:brick_stairs[half=bottom,shape=straight,facing=west]",
	},
	{
		108, 2, "minecraft:brick_stairs[half=bottom,shape=straight,facing=south]",
	},
	{
		108, 3, "minecraft:brick_stairs[half=bottom,shape=straight,facing=north]",
	},
	{
		108, 4, "minecraft:brick_stairs[half=top,shape=straight,facing=east]",
	},
	{
		108, 5, "minecraft:brick_stairs[half=top,shape=straight,facing=west]",
	},
	{
		108, 6, "minecraft:brick_stairs[half=top,shape=straight,facing=south]",
	},
	{
		108, 7, "minecraft:brick_stairs[half=top,shape=straight,facing=north]",
	},
	{
		109, 0, "minecraft:stone_brick_stairs[half=bottom,shape=straight,facing=east]",
	},
	{
		109, 1, "minecraft:stone_brick_stairs[half=bottom,shape=straight,facing=west]",
	},
	{
		109, 2, "minecraft:stone_brick_stairs[half=bottom,shape=straight,facing=south]",
	},
	{
		109, 3, "minecraft:stone_brick_stairs[half=bottom,shape=straight,facing=north]",
	},
	{
		109, 4, "minecraft:stone_brick_stairs[half=top,shape=straight,facing=east]",
	},
	{
		109, 5, "minecraft:stone_brick_stairs[half=top,shape=straight,facing=west]",
	},
	{
		109, 6, "minecraft:stone_brick_stairs[half=top,shape=straight,facing=south]",
	},
	{
		109, 7, "minecraft:stone_brick_stairs[half=top,shape=straight,facing=north]",
	},
	{
		110, 0, "minecraft:mycelium[snowy=false]",
	},
	{
		111, 0, "minecraft:lily_pad",
	},
	{
		112, 0, "minecraft:nether_bricks",
	},
	{
		113, 0, "minecraft:nether_brick_fence[east=false,south=false,north=false,west=false]",
	},
	{
		114, 0, "minecraft:nether_brick_stairs[half=bottom,shape=straight,facing=east]",
	},
	{
		114, 1, "minecraft:nether_brick_stairs[half=bottom,shape=straight,facing=west]",
	},
	{
		114, 2, "minecraft:nether_brick_stairs[half=bottom,shape=straight,facing=south]",
	},
	{
		114, 3, "minecraft:nether_brick_stairs[half=bottom,shape=straight,facing=north]",
	},
	{
		114, 4, "minecraft:nether_brick_stairs[half=top,shape=straight,facing=east]",
	},
	{
		114, 5, "minecraft:nether_brick_stairs[half=top,shape=straight,facing=west]",
	},
	{
		114, 6, "minecraft:nether_brick_stairs[half=top,shape=straight,facing=south]",
	},
	{
		114, 7, "minecraft:nether_brick_stairs[half=top,shape=straight,facing=north]",
	},
	{
		115, 0, "minecraft:nether_wart[age=0]",
	},
	{
		115, 1, "minecraft:nether_wart[age=1]",
	},
	{
		115, 2, "minecraft:nether_wart[age=2]",
	},
	{
		115, 3, "minecraft:nether_wart[age=3]",
	},
	{
		116, 0, "minecraft:enchanting_table",
	},
	{
		117, 0, "minecraft:brewing_stand[has_bottle_0=false,has_bottle_1=false,has_bottle_2=false]",
	},
	{
		117, 1, "minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=false,has_bottle_2=false]",
	},
	{
		117, 2, "minecraft:brewing_stand[has_bottle_0=false,has_bottle_1=true,has_bottle_2=false]",
	},
	{
		117, 3, "minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=true,has_bottle_2=false]",
	},
	{
		117, 4, "minecraft:brewing_stand[has_bottle_0=false,has_bottle_1=false,has_bottle_2=true]",
	},
	{
		117, 5, "minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=false,has_bottle_2=true]",
	},
	{
		117, 6, "minecraft:brewing_stand[has_bottle_0=false,has_bottle_1=true,has_bottle_2=true]",
	},
	{
		117, 7, "minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=true,has_bottle_2=true]",
	},
	{
		118, 0, "minecraft:cauldron[level=0]",
	},
	{
		118, 1, "minecraft:cauldron[level=1]",
	},
	{
		118, 2, "minecraft:cauldron[level=2]",
	},
	{
		118, 3, "minecraft:cauldron[level=3]",
	},
	{
		119, 0, "minecraft:end_portal",
	},
	{
		120, 0, "minecraft:end_portal_frame[eye=false,facing=south]",
	},
	{
		120, 1, "minecraft:end_portal_frame[eye=false,facing=west]",
	},
	{
		120, 2, "minecraft:end_portal_frame[eye=false,facing=north]",
	},
	{
		120, 3, "minecraft:end_portal_frame[eye=false,facing=east]",
	},
	{
		120, 4, "minecraft:end_portal_frame[eye=true,facing=south]",
	},
	{
		120, 5, "minecraft:end_portal_frame[eye=true,facing=west]",
	},
	{
		120, 6, "minecraft:end_portal_frame[eye=true,facing=north]",
	},
	{
		120, 7, "minecraft:end_portal_frame[eye=true,facing=east]",
	},
	{
		121, 0, "minecraft:end_stone",
	},
	{
		122, 0, "minecraft:dragon_egg",
	},
	{
		123, 0, "minecraft:redstone_lamp[lit=false]",
	},
	{
		124, 0, "minecraft:redstone_lamp[lit=true]",
	},
	{
		125, 0, "minecraft:oak_slab[type=double]",
	},
	{
		125, 1, "minecraft:spruce_slab[type=double]",
	},
	{
		125, 2, "minecraft:birch_slab[type=double]",
	},
	{
		125, 3, "minecraft:jungle_slab[type=double]",
	},
	{
		125, 4, "minecraft:acacia_slab[type=double]",
	},
	{
		125, 5, "minecraft:dark_oak_slab[type=double]",
	},
	{
		126, 0, "minecraft:oak_slab[type=bottom]",
	},
	{
		126, 1, "minecraft:spruce_slab[type=bottom]",
	},
	{
		126, 2, "minecraft:birch_slab[type=bottom]",
	},
	{
		126, 3, "minecraft:jungle_slab[type=bottom]",
	},
	{
		126, 4, "minecraft:acacia_slab[type=bottom]",
	},
	{
		126, 5, "minecraft:dark_oak_slab[type=bottom]",
	},
	{
		126, 8, "minecraft:oak_slab[type=top]",
	},
	{
		126, 9, "minecraft:spruce_slab[type=top]",
	},
	{
		126, 10, "minecraft:birch_slab[type=top]",
	},
	{
		126, 11, "minecraft:jungle_slab[type=top]",
	},
	{
		126, 12, "minecraft:acacia_slab[type=top]",
	},
	{
		126, 13, "minecraft:dark_oak_slab[type=top]",
	},
	{
		127, 0, "minecraft:cocoa[facing=south,age=0]",
	},
	{
		127, 1, "minecraft:cocoa[facing=west,age=0]",
	},
	{
		127, 2, "minecraft:cocoa[facing=north,age=0]",
	},
	{
		127, 3, "minecraft:cocoa[facing=east,age=0]",
	},
	{
		127, 4, "minecraft:cocoa[facing=south,age=1]",
	},
	{
		127, 5, "minecraft:cocoa[facing=west,age=1]",
	},
	{
		127, 6, "minecraft:cocoa[facing=north,age=1]",
	},
	{
		127, 7, "minecraft:cocoa[facing=east,age=1]",
	},
	{
		127, 8, "minecraft:cocoa[facing=south,age=2]",
	},
	{
		127, 9, "minecraft:cocoa[facing=west,age=2]",
	},
	{
		127, 10, "minecraft:cocoa[facing=north,age=2]",
	},
	{
		127, 11, "minecraft:cocoa[facing=east,age=2]",
	},
	{
		128, 0, "minecraft:sandstone_stairs[half=bottom,shape=straight,facing=east]",
	},
	{
		128, 1, "minecraft:sandstone_stairs[half=bottom,shape=straight,facing=west]",
	},
	{
		128, 2, "minecraft:sandstone_stairs[half=bottom,shape=straight,facing=south]",
	},
	{
		128, 3, "minecraft:sandstone_stairs[half=bottom,shape=straight,facing=north]",
	},
	{
		128, 4, "minecraft:sandstone_stairs[half=top,shape=straight,facing=east]",
	},
	{
		128, 5, "minecraft:sandstone_stairs[half=top,shape=straight,facing=west]",
	},
	{
		128, 6, "minecraft:sandstone_stairs[half=top,shape=straight,facing=south]",
	},
	{
		128, 7, "minecraft:sandstone_stairs[half=top,shape=straight,facing=north]",
	},
	{
		129, 0, "minecraft:emerald_ore",
	},
	{
		130, 2, "minecraft:ender_chest[facing=north]",
	},
	{
		130, 3, "minecraft:ender_chest[facing=south]",
	},
	{
		130, 4, "minecraft:ender_chest[facing=west]",
	},
	{
		130, 5, "minecraft:ender_chest[facing=east]",
	},
	{
		131, 0, "minecraft:tripwire_hook[powered=false,attached=false,facing=south]",
	},
	{
		131, 1, "minecraft:tripwire_hook[powered=false,attached=false,facing=west]",
	},
	{
		131, 2, "minecraft:tripwire_hook[powered=false,attached=false,facing=north]",
	},
	{
		131, 3, "minecraft:tripwire_hook[powered=false,attached=false,facing=east]",
	},
	{
		131, 4, "minecraft:tripwire_hook[powered=false,attached=true,facing=south]",
	},
	{
		131, 5, "minecraft:tripwire_hook[powered=false,attached=true,facing=west]",
	},
	{
		131, 6, "minecraft:tripwire_hook[powered=false,attached=true,facing=north]",
	},
	{
		131, 7, "minecraft:tripwire_hook[powered=false,attached=true,facing=east]",
	},
	{
		131, 8, "minecraft:tripwire_hook[powered=true,attached=false,facing=south]",
	},
	{
		131, 9, "minecraft:tripwire_hook[powered=true,attached=false,facing=west]",
	},
	{
		131, 10, "minecraft:tripwire_hook[powered=true,attached=false,facing=north]",
	},
	{
		131, 11, "minecraft:tripwire_hook[powered=true,attached=false,facing=east]",
	},
	{
		131, 12, "minecraft:tripwire_hook[powered=true,attached=true,facing=south]",
	},
	{
		131, 13, "minecraft:tripwire_hook[powered=true,attached=true,facing=west]",
	},
	{
		131, 14, "minecraft:tripwire_hook[powered=true,attached=true,facing=north]",
	},
	{
		131, 15, "minecraft:tripwire_hook[powered=true,attached=true,facing=east]",
	},
	{
		132, 0, "minecraft:tripwire[disarmed=false,east=false,powered=false,south=false,north=false,west=false,attached=false]",
	},
	{
		132, 1, "minecraft:tripwire[disarmed=false,east=false,powered=true,south=false,north=false,west=false,attached=false]",
	},
	{
		132, 4, "minecraft:tripwire[disarmed=false,east=false,powered=false,south=false,north=false,west=false,attached=true]",
	},
	{
		132, 5, "minecraft:tripwire[disarmed=false,east=false,powered=true,south=false,north=false,west=false,attached=true]",
	},
	{
		132, 8, "minecraft:tripwire[disarmed=true,east=false,powered=false,south=false,north=false,west=false,attached=false]",
	},
	{
		132, 9, "minecraft:tripwire[disarmed=true,east=false,powered=true,south=false,north=false,west=false,attached=false]",
	},
	{
		132, 12, "minecraft:tripwire[disarmed=true,east=false,powered=false,south=false,north=false,west=false,attached=true]",
	},
	{
		132, 13, "minecraft:tripwire[disarmed=true,east=false,powered=true,south=false,north=false,west=false,attached=true]",
	},
	{
		133, 0, "minecraft:emerald_block",
	},
	{
		134, 0, "minecraft:spruce_stairs[half=bottom,shape=straight,facing=east]",
	},
	{
		134, 1, "minecraft:spruce_stairs[half=bottom,shape=straight,facing=west]",
	},
	{
		134, 2, "minecraft:spruce_stairs[half=bottom,shape=straight,facing=south]",
	},
	{
		134, 3, "minecraft:spruce_stairs[half=bottom,shape=straight,facing=north]",
	},
	{
		134, 4, "minecraft:spruce_stairs[half=top,shape=straight,facing=east]",
	},
	{
		134, 5, "minecraft:spruce_stairs[half=top,shape=straight,facing=west]",
	},
	{
		134, 6, "minecraft:spruce_stairs[half=top,shape=straight,facing=south]",
	},
	{
		134, 7, "minecraft:spruce_stairs[half=top,shape=straight,facing=north]",
	},
	{
		135, 0, "minecraft:birch_stairs[half=bottom,shape=straight,facing=east]",
	},
	{
		135, 1, "minecraft:birch_stairs[half=bottom,shape=straight,facing=west]",
	},
	{
		135, 2, "minecraft:birch_stairs[half=bottom,shape=straight,facing=south]",
	},
	{
		135, 3, "minecraft:birch_stairs[half=bottom,shape=straight,facing=north]",
	},
	{
		135, 4, "minecraft:birch_stairs[half=top,shape=straight,facing=east]",
	},
	{
		135, 5, "minecraft:birch_stairs[half=top,shape=straight,facing=west]",
	},
	{
		135, 6, "minecraft:birch_stairs[half=top,shape=straight,facing=south]",
	},
	{
		135, 7, "minecraft:birch_stairs[half=top,shape=straight,facing=north]",
	},
	{
		136, 0, "minecraft:jungle_stairs[half=bottom,shape=straight,facing=east]",
	},
	{
		136, 1, "minecraft:jungle_stairs[half=bottom,shape=straight,facing=west]",
	},
	{
		136, 2, "minecraft:jungle_stairs[half=bottom,shape=straight,facing=south]",
	},
	{
		136, 3, "minecraft:jungle_stairs[half=bottom,shape=straight,facing=north]",
	},
	{
		136, 4, "minecraft:jungle_stairs[half=top,shape=straight,facing=east]",
	},
	{
		136, 5, "minecraft:jungle_stairs[half=top,shape=straight,facing=west]",
	},
	{
		136, 6, "minecraft:jungle_stairs[half=top,shape=straight,facing=south]",
	},
	{
		136, 7, "minecraft:jungle_stairs[half=top,shape=straight,facing=north]",
	},
	{
		137, 0, "minecraft:command_block[conditional=false,facing=down]",
	},
	{
		137, 1, "minecraft:command_block[conditional=false,facing=up]",
	},
	{
		137, 2, "minecraft:command_block[conditional=false,facing=north]",
	},
	{
		137, 3, "minecraft:command_block[conditional=false,facing=south]",
	},
	{
		137, 4, "minecraft:command_block[conditional=false,facing=west]",
	},
	{
		137, 5, "minecraft:command_block[conditional=false,facing=east]",
	},
	{
		137, 8, "minecraft:command_block[conditional=true,facing=down]",
	},
	{
		137, 9, "minecraft:command_block[conditional=true,facing=up]",
	},
	{
		137, 10, "minecraft:command_block[conditional=true,facing=north]",
	},
	{
		137, 11, "minecraft:command_block[conditional=true,facing=south]",
	},
	{
		137, 12, "minecraft:command_block[conditional=true,facing=west]",
	},
	{
		137, 13, "minecraft:command_block[conditional=true,facing=east]",
	},
	{
		138, 0, "minecraft:beacon",
	},
	{
		139, 0, "minecraft:cobblestone_wall[east=false,south=false,north=false,west=false,up=false]",
	},
	{
		139, 1, "minecraft:mossy_cobblestone_wall[east=false,south=false,north=false,west=false,up=false]",
	},
	{
		140, 0, "minecraft:flower_pot",
	},
	{
		140, 1, "minecraft:potted_poppy",
	},
	{
		140, 2, "minecraft:potted_dandelion",
	},
	{
		140, 3, "minecraft:potted_oak_sapling",
	},
	{
		140, 4, "minecraft:potted_spruce_sapling",
	},
	{
		140, 5, "minecraft:potted_birch_sapling",
	},
	{
		140, 6, "minecraft:potted_jungle_sapling",
	},
	{
		140, 7, "minecraft:potted_red_mushroom",
	},
	{
		140, 8, "minecraft:potted_brown_mushroom",
	},
	{
		140, 9, "minecraft:potted_cactus",
	},
	{
		140, 10, "minecraft:potted_dead_bush",
	},
	{
		140, 11, "minecraft:potted_fern",
	},
	{
		140, 12, "minecraft:potted_acacia_sapling",
	},
	{
		140, 13, "minecraft:potted_dark_oak_sapling",
	},
	{
		140, 14, "minecraft:potted_blue_orchid",
	},
	{
		140, 15, "minecraft:potted_allium",
	},
	{
		141, 0, "minecraft:carrots[age=0]",
	},
	{
		141, 1, "minecraft:carrots[age=1]",
	},
	{
		141, 2, "minecraft:carrots[age=2]",
	},
	{
		141, 3, "minecraft:carrots[age=3]",
	},
	{
		141, 4, "minecraft:carrots[age=4]",
	},
	{
		141, 5, "minecraft:carrots[age=5]",
	},
	{
		141, 6, "minecraft:carrots[age=6]",
	},
	{
		141, 7, "minecraft:carrots[age=7]",
	},
	{
		142, 0, "minecraft:potatoes[age=0]",
	},
	{
		142, 1, "minecraft:potatoes[age=1]",
	},
	{
		142, 2, "minecraft:potatoes[age=2]",
	},
	{
		142, 3, "minecraft:potatoes[age=3]",
	},
	{
		142, 4, "minecraft:potatoes[age=4]",
	},
	{
		142, 5, "minecraft:potatoes[age=5]",
	},
	{
		142, 6, "minecraft:potatoes[age=6]",
	},
	{
		142, 7, "minecraft:potatoes[age=7]",
	},
	{
		143, 0, "minecraft:oak_button[powered=false,facing=east,face=ceiling]",
	},
	{
		143, 1, "minecraft:oak_button[powered=false,facing=east,face=wall]",
	},
	{
		143, 2, "minecraft:oak_button[powered=false,facing=west,face=wall]",
	},
	{
		143, 3, "minecraft:oak_button[powered=false,facing=south,face=wall]",
	},
	{
		143, 4, "minecraft:oak_button[powered=false,facing=north,face=wall]",
	},
	{
		143, 5, "minecraft:oak_button[powered=false,facing=east,face=floor]",
	},
	{
		143, 8, "minecraft:oak_button[powered=true,facing=south,face=ceiling]",
	},
	{
		143, 9, "minecraft:oak_button[powered=true,facing=east,face=wall]",
	},
	{
		143, 10, "minecraft:oak_button[powered=true,facing=west,face=wall]",
	},
	{
		143, 11, "minecraft:oak_button[powered=true,facing=south,face=wall]",
	},
	{
		143, 12, "minecraft:oak_button[powered=true,facing=north,face=wall]",
	},
	{
		143, 13, "minecraft:oak_button[powered=true,facing=south,face=floor]",
	},
	{
		144, 0, "minecraft:skeleton_skull[rotation=0]",
	},
	{
		144, 1, "minecraft:skeleton_skull[rotation=4]",
	},
	{
		144, 2, "minecraft:skeleton_wall_skull[facing=north]",
	},
	{
		144, 3, "minecraft:skeleton_wall_skull[facing=south]",
	},
	{
		144, 4, "minecraft:skeleton_wall_skull[facing=west]",
	},
	{
		144, 5, "minecraft:skeleton_wall_skull[facing=east]",
	},
	{
		144, 8, "minecraft:skeleton_skull[rotation=8]",
	},
	{
		144, 9, "minecraft:skeleton_skull[rotation=12]",
	},
	{
		144, 10, "minecraft:skeleton_wall_skull[facing=north]",
	},
	{
		144, 11, "minecraft:skeleton_wall_skull[facing=south]",
	},
	{
		144, 12, "minecraft:skeleton_wall_skull[facing=west]",
	},
	{
		144, 13, "minecraft:skeleton_wall_skull[facing=east]",
	},
	{
		145, 0, "minecraft:anvil[facing=south]",
	},
	{
		145, 1, "minecraft:anvil[facing=west]",
	},
	{
		145, 2, "minecraft:anvil[facing=north]",
	},
	{
		145, 3, "minecraft:anvil[facing=east]",
	},
	{
		145, 4, "minecraft:chipped_anvil[facing=south]",
	},
	{
		145, 5, "minecraft:chipped_anvil[facing=west]",
	},
	{
		145, 6, "minecraft:chipped_anvil[facing=north]",
	},
	{
		145, 7, "minecraft:chipped_anvil[facing=east]",
	},
	{
		145, 8, "minecraft:damaged_anvil[facing=south]",
	},
	{
		145, 9, "minecraft:damaged_anvil[facing=west]",
	},
	{
		145, 10, "minecraft:damaged_anvil[facing=north]",
	},
	{
		145, 11, "minecraft:damaged_anvil[facing=east]",
	},
	{
		146, 2, "minecraft:trapped_chest[facing=north,type=single]",
	},
	{
		146, 3, "minecraft:trapped_chest[facing=south,type=single]",
	},
	{
		146, 4, "minecraft:trapped_chest[facing=west,type=single]",
	},
	{
		146, 5, "minecraft:trapped_chest[facing=east,type=single]",
	},
	{
		147, 0, "minecraft:light_weighted_pressure_plate[power=0]",
	},
	{
		147, 1, "minecraft:light_weighted_pressure_plate[power=1]",
	},
	{
		147, 2, "minecraft:light_weighted_pressure_plate[power=2]",
	},
	{
		147, 3, "minecraft:light_weighted_pressure_plate[power=3]",
	},
	{
		147, 4, "minecraft:light_weighted_pressure_plate[power=4]",
	},
	{
		147, 5, "minecraft:light_weighted_pressure_plate[power=5]",
	},
	{
		147, 6, "minecraft:light_weighted_pressure_plate[power=6]",
	},
	{
		147, 7, "minecraft:light_weighted_pressure_plate[power=7]",
	},
	{
		147, 8, "minecraft:light_weighted_pressure_plate[power=8]",
	},
	{
		147, 9, "minecraft:light_weighted_pressure_plate[power=9]",
	},
	{
		147, 10, "minecraft:light_weighted_pressure_plate[power=10]",
	},
	{
		147, 11, "minecraft:light_weighted_pressure_plate[power=11]",
	},
	{
		147, 12, "minecraft:light_weighted_pressure_plate[power=12]",
	},
	{
		147, 13, "minecraft:light_weighted_pressure_plate[power=13]",
	},
	{
		147, 14, "minecraft:light_weighted_pressure_plate[power=14]",
	},
	{
		147, 15, "minecraft:light_weighted_pressure_plate[power=15]",
	},
	{
		148, 0, "minecraft:heavy_weighted_pressure_plate[power=0]",
	},
	{
		148, 1, "minecraft:heavy_weighted_pressure_plate[power=1]",
	},
	{
		148, 2, "minecraft:heavy_weighted_pressure_plate[power=2]",
	},
	{
		148, 3, "minecraft:heavy_weighted_pressure_plate[power=3]",
	},
	{
		148, 4, "minecraft:heavy_weighted_pressure_plate[power=4]",
	},
	{
		148, 5, "minecraft:heavy_weighted_pressure_plate[power=5]",
	},
	{
		148, 6, "minecraft:heavy_weighted_pressure_plate[power=6]",
	},
	{
		148, 7, "minecraft:heavy_weighted_pressure_plate[power=7]",
	},
	{
		148, 8, "minecraft:heavy_weighted_pressure_plate[power=8]",
	},
	{
		148, 9, "minecraft:heavy_weighted_pressure_plate[power=9]",
	},
	{
		148, 10, "minecraft:heavy_weighted_pressure_plate[power=10]",
	},
	{
		148, 11, "minecraft:heavy_weighted_pressure_plate[power=11]",
	},
	{
		148, 12, "minecraft:heavy_weighted_pressure_plate[power=12]",
	},
	{
		148, 13, "minecraft:heavy_weighted_pressure_plate[power=13]",
	},
	{
		148, 14, "minecraft:heavy_weighted_pressure_plate[power=14]",
	},
	{
		148, 15, "minecraft:heavy_weighted_pressure_plate[power=15]",
	},
	{
		149, 0, "minecraft:comparator[mode=compare,powered=false,facing=south]",
	},
	{
		149, 1, "minecraft:comparator[mode=compare,powered=false,facing=west]",
	},
	{
		149, 2, "minecraft:comparator[mode=compare,powered=false,facing=north]",
	},
	{
		149, 3, "minecraft:comparator[mode=compare,powered=false,facing=east]",
	},
	{
		149, 4, "minecraft:comparator[mode=subtract,powered=false,facing=south]",
	},
	{
		149, 5, "minecraft:comparator[mode=subtract,powered=false,facing=west]",
	},
	{
		149, 6, "minecraft:comparator[mode=subtract,powered=false,facing=north]",
	},
	{
		149, 7, "minecraft:comparator[mode=subtract,powered=false,facing=east]",
	},
	{
		149, 8, "minecraft:comparator[mode=compare,powered=false,facing=south]",
	},
	{
		149, 9, "minecraft:comparator[mode=compare,powered=false,facing=west]",
	},
	{
		149, 10, "minecraft:comparator[mode=compare,powered=false,facing=north]",
	},
	{
		149, 11, "minecraft:comparator[mode=compare,powered=false,facing=east]",
	},
	{
		149, 12, "minecraft:comparator[mode=subtract,powered=false,facing=south]",
	},
	{
		149, 13, "minecraft:comparator[mode=subtract,powered=false,facing=west]",
	},
	{
		149, 14, "minecraft:comparator[mode=subtract,powered=false,facing=north]",
	},
	{
		149, 15, "minecraft:comparator[mode=subtract,powered=false,facing=east]",
	},
	{
		150, 0, "minecraft:comparator[mode=compare,powered=true,facing=south]",
	},
	{
		150, 1, "minecraft:comparator[mode=compare,powered=true,facing=west]",
	},
	{
		150, 2, "minecraft:comparator[mode=compare,powered=true,facing=north]",
	},
	{
		150, 3, "minecraft:comparator[mode=compare,powered=true,facing=east]",
	},
	{
		150, 4, "minecraft:comparator[mode=subtract,powered=true,facing=south]",
	},
	{
		150, 5, "minecraft:comparator[mode=subtract,powered=true,facing=west]",
	},
	{
		150, 6, "minecraft:comparator[mode=subtract,powered=true,facing=north]",
	},
	{
		150, 7, "minecraft:comparator[mode=subtract,powered=true,facing=east]",
	},
	{
		150, 8, "minecraft:comparator[mode=compare,powered=true,facing=south]",
	},
	{
		150, 9, "minecraft:comparator[mode=compare,powered=true,facing=west]",
	},
	{
		150, 10, "minecraft:comparator[mode=compare,powered=true,facing=north]",
	},
	{
		150, 11, "minecraft:comparator[mode=compare,powered=true,facing=east]",
	},
	{
		150, 12, "minecraft:comparator[mode=subtract,powered=true,facing=south]",
	},
	{
		150, 13, "minecraft:comparator[mode=subtract,powered=true,facing=west]",
	},
	{
		150, 14, "minecraft:comparator[mode=subtract,powered=true,facing=north]",
	},
	{
		150, 15, "minecraft:comparator[mode=subtract,powered=true,facing=east]",
	},
	{
		151, 0, "minecraft:daylight_detector[inverted=false,power=0]",
	},
	{
		151, 1, "minecraft:daylight_detector[inverted=false,power=1]",
	},
	{
		151, 2, "minecraft:daylight_detector[inverted=false,power=2]",
	},
	{
		151, 3, "minecraft:daylight_detector[inverted=false,power=3]",
	},
	{
		151, 4, "minecraft:daylight_detector[inverted=false,power=4]",
	},
	{
		151, 5, "minecraft:daylight_detector[inverted=false,power=5]",
	},
	{
		151, 6, "minecraft:daylight_detector[inverted=false,power=6]",
	},
	{
		151, 7, "minecraft:daylight_detector[inverted=false,power=7]",
	},
	{
		151, 8, "minecraft:daylight_detector[inverted=false,power=8]",
	},
	{
		151, 9, "minecraft:daylight_detector[inverted=false,power=9]",
	},
	{
		151, 10, "minecraft:daylight_detector[inverted=false,power=10]",
	},
	{
		151, 11, "minecraft:daylight_detector[inverted=false,power=11]",
	},
	{
		151, 12, "minecraft:daylight_detector[inverted=false,power=12]",
	},
	{
		151, 13, "minecraft:daylight_detector[inverted=false,power=13]",
	},
	{
		151, 14, "minecraft:daylight_detector[inverted=false,power=14]",
	},
	{
		151, 15, "minecraft:daylight_detector[inverted=false,power=15]",
	},
	{
		152, 0, "minecraft:redstone_block",
	},
	{
		153, 0, "minecraft:nether_quartz_ore",
	},
	{
		154, 0, "minecraft:hopper[facing=down,enabled=true]",
	},
	{
		154, 2, "minecraft:hopper[facing=north,enabled=true]",
	},
	{
		154, 3, "minecraft:hopper[facing=south,enabled=true]",
	},
	{
		154, 4, "minecraft:hopper[facing=west,enabled=true]",
	},
	{
		154, 5, "minecraft:hopper[facing=east,enabled=true]",
	},
	{
		154, 8, "minecraft:hopper[facing=down,enabled=false]",
	},
	{
		154, 10, "minecraft:hopper[facing=north,enabled=false]",
	},
	{
		154, 11, "minecraft:hopper[facing=south,enabled=false]",
	},
	{
		154, 12, "minecraft:hopper[facing=west,enabled=false]",
	},
	{
		154, 13, "minecraft:hopper[facing=east,enabled=false]",
	},
	{
		155, 0, "minecraft:quartz_block",
	},
	{
		155, 1, "minecraft:chiseled_quartz_block",
	},
	{
		155, 2, "minecraft:quartz_pillar[axis=y]",
	},
	{
		155, 3, "minecraft:quartz_pillar[axis=x]",
	},
	{
		155, 4, "minecraft:quartz_pillar[axis=z]",
	},
	{
		155, 6, "minecraft:quartz_pillar[axis=x]",
	},
	{
		155, 10, "minecraft:quartz_pillar[axis=z]",
	},
	{
		156, 0, "minecraft:quartz_stairs[half=bottom,shape=straight,facing=east]",
	},
	{
		156, 1, "minecraft:quartz_stairs[half=bottom,shape=straight,facing=west]",
	},
	{
		156, 2, "minecraft:quartz_stairs[half=bottom,shape=straight,facing=south]",
	},
	{
		156, 3, "minecraft:quartz_stairs[half=bottom,shape=straight,facing=north]",
	},
	{
		156, 4, "minecraft:quartz_stairs[half=top,shape=straight,facing=east]",
	},
	{
		156, 5, "minecraft:quartz_stairs[half=top,shape=straight,facing=west]",
	},
	{
		156, 6, "minecraft:quartz_stairs[half=top,shape=straight,facing=south]",
	},
	{
		156, 7, "minecraft:quartz_stairs[half=top,shape=straight,facing=north]",
	},
	{
		157, 0, "minecraft:activator_rail[shape=north_south,powered=false]",
	},
	{
		157, 1, "minecraft:activator_rail[shape=east_west,powered=false]",
	},
	{
		157, 2, "minecraft:activator_rail[shape=ascending_east,powered=false]",
	},
	{
		157, 3, "minecraft:activator_rail[shape=ascending_west,powered=false]",
	},
	{
		157, 4, "minecraft:activator_rail[shape=ascending_north,powered=false]",
	},
	{
		157, 5, "minecraft:activator_rail[shape=ascending_south,powered=false]",
	},
	{
		157, 8, "minecraft:activator_rail[shape=north_south,powered=true]",
	},
	{
		157, 9, "minecraft:activator_rail[shape=east_west,powered=true]",
	},
	{
		157, 10, "minecraft:activator_rail[shape=ascending_east,powered=true]",
	},
	{
		157, 11, "minecraft:activator_rail[shape=ascending_west,powered=true]",
	},
	{
		157, 12, "minecraft:activator_rail[shape=ascending_north,powered=true]",
	},
	{
		157, 13, "minecraft:activator_rail[shape=ascending_south,powered=true]",
	},
	{
		158, 0, "minecraft:dropper[triggered=false,facing=down]",
	},
	{
		158, 1, "minecraft:dropper[triggered=false,facing=up]",
	},
	{
		158, 2, "minecraft:dropper[triggered=false,facing=north]",
	},
	{
		158, 3, "minecraft:dropper[triggered=false,facing=south]",
	},
	{
		158, 4, "minecraft:dropper[triggered=false,facing=west]",
	},
	{
		158, 5, "minecraft:dropper[triggered=false,facing=east]",
	},
	{
		158, 8, "minecraft:dropper[triggered=true,facing=down]",
	},
	{
		158, 9, "minecraft:dropper[triggered=true,facing=up]",
	},
	{
		158, 10, "minecraft:dropper[triggered=true,facing=north]",
	},
	{
		158, 11, "minecraft:dropper[triggered=true,facing=south]",
	},
	{
		158, 12, "minecraft:dropper[triggered=true,facing=west]",
	},
	{
		158, 13, "minecraft:dropper[triggered=true,facing=east]",
	},
	{
		159, 0, "minecraft:white_terracotta",
	},
	{
		159, 1, "minecraft:orange_terracotta",
	},
	{
		159, 2, "minecraft:magenta_terracotta",
	},
	{
		159, 3, "minecraft:light_blue_terracotta",
	},
	{
		159, 4, "minecraft:yellow_terracotta",
	},
	{
		159, 5, "minecraft:lime_terracotta",
	},
	{
		159, 6, "minecraft:pink_terracotta",
	},
	{
		159, 7, "minecraft:gray_terracotta",
	},
	{
		159, 8, "minecraft:light_gray_terracotta",
	},
	{
		159, 9, "minecraft:cyan_terracotta",
	},
	{
		159, 10, "minecraft:purple_terracotta",
	},
	{
		159, 11, "minecraft:blue_terracotta",
	},
	{
		159, 12, "minecraft:brown_terracotta",
	},
	{
		159, 13, "minecraft:green_terracotta",
	},
	{
		159, 14, "minecraft:red_terracotta",
	},
	{
		159, 15, "minecraft:black_terracotta",
	},
	{
		160, 0, "minecraft:white_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 1, "minecraft:orange_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 2, "minecraft:magenta_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 3, "minecraft:light_blue_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 4, "minecraft:yellow_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 5, "minecraft:lime_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 6, "minecraft:pink_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 7, "minecraft:gray_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 8, "minecraft:light_gray_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 9, "minecraft:cyan_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 10, "minecraft:purple_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 11, "minecraft:blue_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 12, "minecraft:brown_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 13, "minecraft:green_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 14, "minecraft:red_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		160, 15, "minecraft:black_stained_glass_pane[east=false,south=false,north=false,west=false]",
	},
	{
		161, 0, "minecraft:acacia_leaves[persistent=false,distance=1]",
	},
	{
		161, 1, "minecraft:dark_oak_leaves[persistent=false,distance=1]",
	},
	{
		161, 4, "minecraft:acacia_leaves[persistent=true,distance=1]",
	},
	{
		161, 5, "minecraft:dark_oak_leaves[persistent=true,distance=1]",
	},
	{
		161, 8, "minecraft:acacia_leaves[persistent=false,distance=1]",
	},
	{
		161, 9, "minecraft:dark_oak_leaves[persistent=false,distance=1]",
	},
	{
		161, 12, "minecraft:acacia_leaves[persistent=true,distance=1]",
	},
	{
		161, 13, "minecraft:dark_oak_leaves[persistent=true,distance=1]",
	},
	{
		162, 0, "minecraft:acacia_log[axis=y]",
	},
	{
		162, 1, "minecraft:dark_oak_log[axis=y]",
	},
	{
		162, 4, "minecraft:acacia_log[axis=x]",
	},
	{
		162, 5, "minecraft:dark_oak_log[axis=x]",
	},
	{
		162, 8, "minecraft:acacia_log[axis=z]",
	},
	{
		162, 9, "minecraft:dark_oak_log[axis=z]",
	},
	{
		162, 12, "minecraft:acacia_wood",
	},
	{
		162, 13, "minecraft:dark_oak_wood",
	},
	{
		163, 0, "minecraft:acacia_stairs[half=bottom,shape=straight,facing=east]",
	},
	{
		163, 1, "minecraft:acacia_stairs[half=bottom,shape=straight,facing=west]",
	},
	{
		163, 2, "minecraft:acacia_stairs[half=bottom,shape=straight,facing=south]",
	},
	{
		163, 3, "minecraft:acacia_stairs[half=bottom,shape=straight,facing=north]",
	},
	{
		163, 4, "minecraft:acacia_stairs[half=top,shape=straight,facing=east]",
	},
	{
		163, 5, "minecraft:acacia_stairs[half=top,shape=straight,facing=west]",
	},
	{
		163, 6, "minecraft:acacia_stairs[half=top,shape=straight,facing=south]",
	},
	{
		163, 7, "minecraft:acacia_stairs[half=top,shape=straight,facing=north]",
	},
	{
		164, 0, "minecraft:dark_oak_stairs[half=bottom,shape=straight,facing=east]",
	},
	{
		164, 1, "minecraft:dark_oak_stairs[half=bottom,shape=straight,facing=west]",
	},
	{
		164, 2, "minecraft:dark_oak_stairs[half=bottom,shape=straight,facing=south]",
	},
	{
		164, 3, "minecraft:dark_oak_stairs[half=bottom,shape=straight,facing=north]",
	},
	{
		164, 4, "minecraft:dark_oak_stairs[half=top,shape=straight,facing=east]",
	},
	{
		164, 5, "minecraft:dark_oak_stairs[half=top,shape=straight,facing=west]",
	},
	{
		164, 6, "minecraft:dark_oak_stairs[half=top,shape=straight,facing=south]",
	},
	{
		164, 7, "minecraft:dark_oak_stairs[half=top,shape=straight,facing=north]",
	},
	{
		165, 0, "minecraft:slime_block",
	},
	{
		166, 0, "minecraft:barrier",
	},
	{
		167, 0, "minecraft:iron_trapdoor[half=bottom,facing=north,open=false]",
	},
	{
		167, 1, "minecraft:iron_trapdoor[half=bottom,facing=south,open=false]",
	},
	{
		167, 2, "minecraft:iron_trapdoor[half=bottom,facing=west,open=false]",
	},
	{
		167, 3, "minecraft:iron_trapdoor[half=bottom,facing=east,open=false]",
	},
	{
		167, 4, "minecraft:iron_trapdoor[half=bottom,facing=north,open=true]",
	},
	{
		167, 5, "minecraft:iron_trapdoor[half=bottom,facing=south,open=true]",
	},
	{
		167, 6, "minecraft:iron_trapdoor[half=bottom,facing=west,open=true]",
	},
	{
		167, 7, "minecraft:iron_trapdoor[half=bottom,facing=east,open=true]",
	},
	{
		167, 8, "minecraft:iron_trapdoor[half=top,facing=north,open=false]",
	},
	{
		167, 9, "minecraft:iron_trapdoor[half=top,facing=south,open=false]",
	},
	{
		167, 10, "minecraft:iron_trapdoor[half=top,facing=west,open=false]",
	},
	{
		167, 11, "minecraft:iron_trapdoor[half=top,facing=east,open=false]",
	},
	{
		167, 12, "minecraft:iron_trapdoor[half=top,facing=north,open=true]",
	},
	{
		167, 13, "minecraft:iron_trapdoor[half=top,facing=south,open=true]",
	},
	{
		167, 14, "minecraft:iron_trapdoor[half=top,facing=west,open=true]",
	},
	{
		167, 15, "minecraft:iron_trapdoor[half=top,facing=east,open=true]",
	},
	{
		168, 0, "minecraft:prismarine",
	},
	{
		168, 1, "minecraft:prismarine_bricks",
	},
	{
		168, 2, "minecraft:dark_prismarine",
	},
	{
		169, 0, "minecraft:sea_lantern",
	},
	{
		170, 0, "minecraft:hay_block[axis=y]",
	},
	{
		170, 4, "minecraft:hay_block[axis=x]",
	},
	{
		170, 8, "minecraft:hay_block[axis=z]",
	},
	{
		171, 0, "minecraft:white_carpet",
	},
	{
		171, 1, "minecraft:orange_carpet",
	},
	{
		171, 2, "minecraft:magenta_carpet",
	},
	{
		171, 3, "minecraft:light_blue_carpet",
	},
	{
		171, 4, "minecraft:yellow_carpet",
	},
	{
		171, 5, "minecraft:lime_carpet",
	},
	{
		171, 6, "minecraft:pink_carpet",
	},
	{
		171, 7, "minecraft:gray_carpet",
	},
	{
		171, 8, "minecraft:light_gray_carpet",
	},
	{
		171, 9, "minecraft:cyan_carpet",
	},
	{
		171, 10, "minecraft:purple_carpet",
	},
	{
		171, 11, "minecraft:blue_carpet",
	},
	{
		171, 12, "minecraft:brown_carpet",
	},
	{
		171, 13, "minecraft:green_carpet",
	},
	{
		171, 14, "minecraft:red_carpet",
	},
	{
		171, 15, "minecraft:black_carpet",
	},
	{
		172, 0, "minecraft:terracotta",
	},
	{
		173, 0, "minecraft:coal_block",
	},
	{
		174, 0, "minecraft:packed_ice",
	},
	{
		175, 0, "minecraft:sunflower[half=lower]",
	},
	{
		175, 1, "minecraft:lilac[half=lower]",
	},
	{
		175, 2, "minecraft:tall_grass[half=lower]",
	},
	{
		175, 3, "minecraft:large_fern[half=lower]",
	},
	{
		175, 4, "minecraft:rose_bush[half=lower]",
	},
	{
		175, 5, "minecraft:peony[half=lower]",
	},
	{
		175, 8, "minecraft:sunflower[half=upper]",
	},
	{
		175, 9, "minecraft:lilac[half=upper]",
	},
	{
		175, 10, "minecraft:tall_grass[half=upper]",
	},
	{
		175, 11, "minecraft:large_fern[half=upper]",
	},
	{
		175, 12, "minecraft:rose_bush[half=upper]",
	},
	{
		175, 13, "minecraft:peony[half=upper]",
	},
	{
		176, 0, "minecraft:white_banner[rotation=0]",
	},
	{
		176, 1, "minecraft:white_banner[rotation=1]",
	},
	{
		176, 2, "minecraft:white_banner[rotation=2]",
	},
	{
		176, 3, "minecraft:white_banner[rotation=3]",
	},
	{
		176, 4, "minecraft:white_banner[rotation=4]",
	},
	{
		176, 5, "minecraft:white_banner[rotation=5]",
	},
	{
		176, 6, "minecraft:white_banner[rotation=6]",
	},
	{
		176, 7, "minecraft:white_banner[rotation=7]",
	},
	{
		176, 8, "minecraft:white_banner[rotation=8]",
	},
	{
		176, 9, "minecraft:white_banner[rotation=9]",
	},
	{
		176, 10, "minecraft:white_banner[rotation=10]",
	},
	{
		176, 11, "minecraft:white_banner[rotation=11]",
	},
	{
		176, 12, "minecraft:white_banner[rotation=12]",
	},
	{
		176, 13, "minecraft:white_banner[rotation=13]",
	},
	{
		176, 14, "minecraft:white_banner[rotation=14]",
	},
	{
		176, 15, "minecraft:white_banner[rotation=15]",
	},
	{
		177, 2, "minecraft:white_wall_banner[facing=north]",
	},
	{
		177, 3, "minecraft:white_wall_banner[facing=south]",
	},
	{
		177, 4, "minecraft:white_wall_banner[facing=west]",
	},
	{
		177, 5, "minecraft:white_wall_banner[facing=east]",
	},
	{
		178, 0, "minecraft:daylight_detector[inverted=true,power=0]",
	},
	{
		178, 1, "minecraft:daylight_detector[inverted=true,power=1]",
	},
	{
		178, 2, "minecraft:daylight_detector[inverted=true,power=2]",
	},
	{
		178, 3, "minecraft:daylight_detector[inverted=true,power=3]",
	},
	{
		178, 4, "minecraft:daylight_detector[inverted=true,power=4]",
	},
	{
		178, 5, "minecraft:daylight_detector[inverted=true,power=5]",
	},
	{
		178, 6, "minecraft:daylight_detector[inverted=true,power=6]",
	},
	{
		178, 7, "minecraft:daylight_detector[inverted=true,power=7]",
	},
	{
		178, 8, "minecraft:daylight_detector[inverted=true,power=8]",
	},
	{
		178, 9, "minecraft:daylight_detector[inverted=true,power=9]",
	},
	{
		178, 10, "minecraft:daylight_detector[inverted=true,power=10]",
	},
	{
		178, 11, "minecraft:daylight_detector[inverted=true,power=11]",
	},
	{
		178, 12, "minecraft:daylight_detector[inverted=true,power=12]",
	},
	{
		178, 13, "minecraft:daylight_detector[inverted=true,power=13]",
	},
	{
		178, 14, "minecraft:daylight_detector[inverted=true,power=14]",
	},
	{
		178, 15, "minecraft:daylight_detector[inverted=true,power=15]",
	},
	{
		179, 0, "minecraft:red_sandstone",
	},
	{
		179, 1, "minecraft:chiseled_red_sandstone",
	},
	{
		179, 2, "minecraft:cut_red_sandstone",
	},
	{
		180, 0, "minecraft:red_sandstone_stairs[half=bottom,shape=straight,facing=east]",
	},
	{
		180, 1, "minecraft:red_sandstone_stairs[half=bottom,shape=straight,facing=west]",
	},
	{
		180, 2, "minecraft:red_sandstone_stairs[half=bottom,shape=straight,facing=south]",
	},
	{
		180, 3, "minecraft:red_sandstone_stairs[half=bottom,shape=straight,facing=north]",
	},
	{
		180, 4, "minecraft:red_sandstone_stairs[half=top,shape=straight,facing=east]",
	},
	{
		180, 5, "minecraft:red_sandstone_stairs[half=top,shape=straight,facing=west]",
	},
	{
		180, 6, "minecraft:red_sandstone_stairs[half=top,shape=straight,facing=south]",
	},
	{
		180, 7, "minecraft:red_sandstone_stairs[half=top,shape=straight,facing=north]",
	},
	{
		181, 0, "minecraft:red_sandstone_slab[type=double]",
	},
	{
		181, 8, "minecraft:smooth_red_sandstone",
	},
	{
		182, 0, "minecraft:red_sandstone_slab[type=bottom]",
	},
	{
		182, 8, "minecraft:red_sandstone_slab[type=top]",
	},
	{
		183, 0, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=south,open=false]",
	},
	{
		183, 1, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=west,open=false]",
	},
	{
		183, 2, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=north,open=false]",
	},
	{
		183, 3, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=east,open=false]",
	},
	{
		183, 4, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=south,open=true]",
	},
	{
		183, 5, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=west,open=true]",
	},
	{
		183, 6, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=north,open=true]",
	},
	{
		183, 7, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=east,open=true]",
	},
	{
		183, 8, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=south,open=false]",
	},
	{
		183, 9, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=west,open=false]",
	},
	{
		183, 10, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=north,open=false]",
	},
	{
		183, 11, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=east,open=false]",
	},
	{
		183, 12, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=south,open=true]",
	},
	{
		183, 13, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=west,open=true]",
	},
	{
		183, 14, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=north,open=true]",
	},
	{
		183, 15, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=east,open=true]",
	},
	{
		184, 0, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=south,open=false]",
	},
	{
		184, 1, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=west,open=false]",
	},
	{
		184, 2, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=north,open=false]",
	},
	{
		184, 3, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=east,open=false]",
	},
	{
		184, 4, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=south,open=true]",
	},
	{
		184, 5, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=west,open=true]",
	},
	{
		184, 6, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=north,open=true]",
	},
	{
		184, 7, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=east,open=true]",
	},
	{
		184, 8, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=south,open=false]",
	},
	{
		184, 9, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=west,open=false]",
	},
	{
		184, 10, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=north,open=false]",
	},
	{
		184, 11, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=east,open=false]",
	},
	{
		184, 12, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=south,open=true]",
	},
	{
		184, 13, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=west,open=true]",
	},
	{
		184, 14, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=north,open=true]",
	},
	{
		184, 15, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=east,open=true]",
	},
	{
		185, 0, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=south,open=false]",
	},
	{
		185, 1, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=west,open=false]",
	},
	{
		185, 2, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=north,open=false]",
	},
	{
		185, 3, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=east,open=false]",
	},
	{
		185, 4, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=south,open=true]",
	},
	{
		185, 5, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=west,open=true]",
	},
	{
		185, 6, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=north,open=true]",
	},
	{
		185, 7, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=east,open=true]",
	},
	{
		185, 8, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=south,open=false]",
	},
	{
		185, 9, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=west,open=false]",
	},
	{
		185, 10, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=north,open=false]",
	},
	{
		185, 11, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=east,open=false]",
	},
	{
		185, 12, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=south,open=true]",
	},
	{
		185, 13, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=west,open=true]",
	},
	{
		185, 14, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=north,open=true]",
	},
	{
		185, 15, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=east,open=true]",
	},
	{
		186, 0, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=south,open=false]",
	},
	{
		186, 1, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=west,open=false]",
	},
	{
		186, 2, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=north,open=false]",
	},
	{
		186, 3, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=east,open=false]",
	},
	{
		186, 4, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=south,open=true]",
	},
	{
		186, 5, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=west,open=true]",
	},
	{
		186, 6, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=north,open=true]",
	},
	{
		186, 7, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=east,open=true]",
	},
	{
		186, 8, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=south,open=false]",
	},
	{
		186, 9, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=west,open=false]",
	},
	{
		186, 10, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=north,open=false]",
	},
	{
		186, 11, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=east,open=false]",
	},
	{
		186, 12, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=south,open=true]",
	},
	{
		186, 13, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=west,open=true]",
	},
	{
		186, 14, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=north,open=true]",
	},
	{
		186, 15, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=east,open=true]",
	},
	{
		187, 0, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=south,open=false]",
	},
	{
		187, 1, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=west,open=false]",
	},
	{
		187, 2, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=north,open=false]",
	},
	{
		187, 3, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=east,open=false]",
	},
	{
		187, 4, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=south,open=true]",
	},
	{
		187, 5, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=west,open=true]",
	},
	{
		187, 6, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=north,open=true]",
	},
	{
		187, 7, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=east,open=true]",
	},
	{
		187, 8, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=south,open=false]",
	},
	{
		187, 9, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=west,open=false]",
	},
	{
		187, 10, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=north,open=false]",
	},
	{
		187, 11, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=east,open=false]",
	},
	{
		187, 12, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=south,open=true]",
	},
	{
		187, 13, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=west,open=true]",
	},
	{
		187, 14, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=north,open=true]",
	},
	{
		187, 15, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=east,open=true]",
	},
	{
		188, 0, "minecraft:spruce_fence[east=false,south=false,north=false,west=false]",
	},
	{
		189, 0, "minecraft:birch_fence[east=false,south=false,north=false,west=false]",
	},
	{
		190, 0, "minecraft:jungle_fence[east=false,south=false,north=false,west=false]",
	},
	{
		191, 0, "minecraft:dark_oak_fence[east=false,south=false,north=false,west=false]",
	},
	{
		192, 0, "minecraft:acacia_fence[east=false,south=false,north=false,west=false]",
	},
	{
		193, 0, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
	},
	{
		193, 1, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
	},
	{
		193, 2, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
	},
	{
		193, 3, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
	},
	{
		193, 4, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
	},
	{
		193, 5, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
	},
	{
		193, 6, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
	},
	{
		193, 7, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
	},
	{
		193, 8, "minecraft:spruce_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
	},
	{
		193, 9, "minecraft:spruce_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
	},
	{
		193, 10, "minecraft:spruce_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
	},
	{
		193, 11, "minecraft:spruce_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
	},
	{
		194, 0, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
	},
	{
		194, 1, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
	},
	{
		194, 2, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
	},
	{
		194, 3, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
	},
	{
		194, 4, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
	},
	{
		194, 5, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
	},
	{
		194, 6, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
	},
	{
		194, 7, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
	},
	{
		194, 8, "minecraft:birch_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
	},
	{
		194, 9, "minecraft:birch_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
	},
	{
		194, 10, "minecraft:birch_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
	},
	{
		194, 11, "minecraft:birch_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
	},
	{
		195, 0, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
	},
	{
		195, 1, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
	},
	{
		195, 2, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
	},
	{
		195, 3, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
	},
	{
		195, 4, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
	},
	{
		195, 5, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
	},
	{
		195, 6, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
	},
	{
		195, 7, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
	},
	{
		195, 8, "minecraft:jungle_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
	},
	{
		195, 9, "minecraft:jungle_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
	},
	{
		195, 10, "minecraft:jungle_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
	},
	{
		195, 11, "minecraft:jungle_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
	},
	{
		196, 0, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
	},
	{
		196, 1, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
	},
	{
		196, 2, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
	},
	{
		196, 3, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
	},
	{
		196, 4, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
	},
	{
		196, 5, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
	},
	{
		196, 6, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
	},
	{
		196, 7, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
	},
	{
		196, 8, "minecraft:acacia_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
	},
	{
		196, 9, "minecraft:acacia_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
	},
	{
		196, 10, "minecraft:acacia_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
	},
	{
		196, 11, "minecraft:acacia_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
	},
	{
		197, 0, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
	},
	{
		197, 1, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
	},
	{
		197, 2, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
	},
	{
		197, 3, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
	},
	{
		197, 4, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
	},
	{
		197, 5, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
	},
	{
		197, 6, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
	},
	{
		197, 7, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
	},
	{
		197, 8, "minecraft:dark_oak_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
	},
	{
		197, 9, "minecraft:dark_oak_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
	},
	{
		197, 10, "minecraft:dark_oak_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
	},
	{
		197, 11, "minecraft:dark_oak_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
	},
}

var legacyFromInt, legacyFromString = generateLegacyBlockMapping()

func generateLegacyBlockMapping() (map[int]*LegacyBlock, map[string]*LegacyBlock) {
	legacyFromInt := make(map[int]*LegacyBlock)
	legacyFromString := make(map[string]*LegacyBlock)
	for i, v := range legacy {
		bigId := v.GetBlockState()
		legacyFromInt[bigId] = &legacy[i]
		legacyFromString[v.Name] = &legacy[i]
	}

	return legacyFromInt, legacyFromString
}

func GetLegacyFromState(legacy int) *LegacyBlock {
	if val, ok := legacyFromInt[legacy]; ok {
		return val
	} else {
		return legacyFromString["minecraft:air"]
	}
}

func GetLegacyFromName(name string) *LegacyBlock {
	if val, ok := legacyFromString[name]; ok {
		return val
	} else {
		return legacyFromString["minecraft:air"]
	}
}

func GetLegacyMapping() map[int]*LegacyBlock {
	return legacyFromInt
}

func (block *LegacyBlock) GetBlockState() int {
	return block.Id<<4 | (block.Data & 0xF)
}

func GetLegacyBlockState(id int, data int) int {
	return id<<4 | (data & 0xF)
}

func GetTypeDataFromLegacy(legacy int) (int, int) {
	return legacy >> 4, legacy & 0xF
}
