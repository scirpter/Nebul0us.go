package usernames

import (
	"math/rand"
	"time"
)

type UsernameManager struct {
	Available []string
}

func NewUsernameManager() *UsernameManager {
	return &UsernameManager{
		Available: usernames,
	}
}

func (u *UsernameManager) GetRandomUsername() *string {
	if len(u.Available) == 0 {
		return nil // this will never happen
	}
	// choose a random username
	randSrc := rand.NewSource(time.Now().UnixNano())
	RNG := rand.New(randSrc)
	index := RNG.Intn(len(u.Available))
	username := u.Available[index]

	return &username
}

var usernames = []string{
	"ZephyrX13",
	"Shadow_08",
	"KillerK_01",
	"Nova_Gamer",
	"DragonBlade9",
	"DarkKnightX",
	"PixelWizard",
	"IceQueen_22",
	"ArcaneRider",
	"PhoenixFire",
	"Thunderbolt7",
	"NeonNinja_",
	"LaserGamer_",
	"CyberSamurai",
	"MysticMage_",
	"VortexKnight",
	"CosmicViper",
	"ChronoGaming",
	"Ravenclaw_1",
	"XtremeGamer_",
	"BlazeDragon",
	"SpartanX_24",
	"IronWarrior_",
	"NightmareX_",
	"RetroGamer_",
	"JediMaster_7",
	"SamuraiWar_",
	"InfinityX_",
	"Thunderbird_",
	"BlizzardX_",
	"GalacticG_",
	"WarriorX_23",
	"ShadowWar_",
	"MysticalX_",
	"CyborgGamer",
	"AlphaGaming",
	"TheLegend_9",
	"StormKing_",
	"RagingStorm",
	"AtomicGamer",
	"OmegaGaming",
	"TitaniumX_",
	"LunarKnight_",
	"GamerKing_7",
	"ElectricX_",
	"NeoWarrior_",
	"ShadowG_11",
	"PhantomX_22",
	"DragonHeart_",
	"KnightRider_",
	"EagleEye_99",
	"ThunderWolf",
	"CosmicX_77",
	"SonicGamer_",
	"Warrior_XXI",
	"CyberDragon",
	"SamuraiX_24",
	"MysticX_007",
	"ChronoGamer_",
	"DragonFury_",
	"BlazeKnight_",
	"IronManX_99",
	"NightmareXx_",
	"RetroKnight_",
	"JediKnight_7",
	"SamuraiX_13",
	"InfinityG_",
	"ThunderX_007",
	"BlizzardG_22",
	"GalacticX_11",
	"WarriorXx_23",
	"ShadowXx_99",
	"MysticalG_",
	"CyborgX_22",
	"AlphaG_007",
	"TheLegend_13",
	"StormX_22",
	"RagingG_",
	"AtomicX_88",
	"OmegaX_007",
	"TitaniumG_",
	"LunarX_99",
	"GamerX_007",
	"ElectricG_22",
	"NeoX_99",
	"ShadowG_007",
	"PhantomX_99",
	"DragonXx_22",
	"KnightX_007",
	"EagleGamer_",
	"ThunderX_22",
	"CosmicG_11",
	"SonicX_99",
	"WarriorX_22",
	"CyberX_13",
	"SamuraiXx_",
	"MysticG_007",
	"ChronoX_11",
	"DragonKnight",
	"BlazeXx_22",
	"IronX_007",
	"NightmareG_",
	"RetroX_11",
	"JediG_007",
	"SamuraiG_",
	"InfinityKnight",
	"ThunderG_",
	"BlizzardXx_",
	"GalacticG_99",
	"WarriorXx_99",
	"ShadowX_007",
	"MysticalX_22",
	"CyborgKnight",
	"AlphaX_22",
	"TheLegendG_",
	"StormG_007",
	"RagingX_",
	"AtomicG_99",
	"OmegaG_22",
	"TitaniumXx_",
	"LunarG_007",
	"GamerG_99",
	"ElectricXx_",
	"NeoGamer_",
	"ShadowX_22",
	"PhantomG_11",
	"DragonRider_",
	"KnightG_007",
	"EagleXx_",
	"ThunderG_22",
	"CosmicKnight",
	"SonicG_22",
	"WarriorX_007",
	"CyberG_99",
	"SamuraiG_22",
	"MysticXx_",
	"ChronoKnight",
	"DragonX_007",
	"BlazeGamer_",
	"IronKnight_",
	"NightmareXx_",
	"RetroG_22",
	"JediX_99",
	"SamuraiG_007",
	"InfinityG_99",
	"ThunderXx_",
	"BlizzardG_007",
	"GalacticXx_",
	"WarriorG_99",
	"ShadowKnight",
	"MysticalG_22",
	"CyborgXx_",
	"AlphaGamer",
	"TheLegendX_",
	"StormXx_99",
	"RagingGamer",
	"AtomicXx_",
	"OmegaKnight",
	"TitaniumG_22",
	"LunarXx_99",
	"GamerX_22",
	"ElectricG_007",
	"NeoX_22",
	"ShadowG_99",
	"PhantomXx_",
	"DragonG_22",
	"KnightXx_",
	"EagleG_22",
	"ThunderX_99",
	"CosmicXx_",
	"SonicKnight",
	"WarriorG_007",
	"CyberXx_",
	"SamuraiXx_99",
	"MysticG_99",
	"ChronoG_22",
	"DragonXx_99",
	"BlazeX_22",
	"IronG_22",
	"NightmareKnight",
	"RetroXx_99",
	"JediKnight_",
	"SamuraiX_99",
	"InfinityXx_",
	"ThunderKnight",
	"BlizzardG_99",
	"GalacticX_007",
	"WarriorXx_22",
	"ShadowG_22",
	"MysticalXx_99",
	"CyborgG_22",
	"AlphaXx_",
	"TheLegendG_22",
	"StormX_007",
	"RagingKnight",
	"AtomicG_22",
	"OmegaXx_",
	"TitaniumGamer",
	"LunarKnight_",
	"GamerXx_",
	"ElectricKnight",
	"NeoG_22",
	"ShadowRider_",
	"DarkKnightX_",
	"PhantomGamer",
	"StormG_22",
	"CyberKnight_",
	"InfinityX_99",
	"BlizzardKnight",
	"TitaniumX_99",
	"NightmareG_99",
	"DragonRiderX",
	"SamuraiGamer",
	"MysticalKnight",
	"OmegaG_007",
	"IronGamer_",
	"WarriorXx_007",
	"ChronoXx_99",
	"GalacticG_22",
	"EagleKnight_",
	"ThunderG_99",
	"NeoG_22",
	"DragonKnightX",
	"ShadowXx_99",
	"MysticKnight_",
	"AtomicX_99",
	"LunarG_22",
	"BlazeG_22",
	"TheLegendKnight",
	"SamuraiXx_22",
	"RetroGamer_",
	"SonicKnightX",
	"WarriorG_22",
	"CosmicG_99",
	"PhantomX_22",
	"AlphaG_22",
	"ChronoGamer_",
	"StormKnightX",
	"ElectricX_99",
	"ShadowG_007",
	"GalacticXx_99",
	"OmegaKnightX",
	"TitaniumG_007",
	"IronKnightXx",
	"WarriorGamer_",
	"DragonXx_007",
	"MysticalX_007",
	"ThunderXx_99",
	"BlizzardG_22",
	"SamuraiKnight",
	"RagingXx_99",
	"NeoKnight_",
	"LunarX_22",
	"CyborgGamer_",
	"DragonGamer_",
	"InfinityG_22",
	"AtomicKnight_",
	"NightmareX_22",
	"TitaniumKnight",
	"EagleX_22",
	"GalacticGamer",
	"ChronoKnightX",
	"ShadowKnightX",
	"MysticXx_22",
	"WarriorX_22",
	"PhantomG_007",
	"TheLegendXx_",
	"StormGamer_",
	"AlphaX_99",
	"ElectricG_22",
	"DragonKnightG",
	"SamuraiX_22",
	"BlazeG_99",
	"ChronoG_007",
	"RetroKnight_",
	"SonicG_007",
	"CyborgX_22",
	"OmegaGamer_",
	"LunarKnightX",
	"ThunderKnightX",
	"CosmicX_99",
	"IronXx_99",
	"WarriorKnight_",
	"DragonRiderG",
	"ShadowG_22",
	"MysticalKnightX",
	"AtomicG_007",
	"TitaniumXx_22",
	"NightmareKnightX",
	"GalacticX_22",
	"ChronoGamerX",
	"StormXx_22",
	"PhantomKnight",
	"AlphaKnightX",
	"ElectricGamer_",
	"DragonX_22",
	"SamuraiG_99",
	"BlizzardX_22",
	"RetroXx_22",
	"SonicGamer_",
	"CyberXx_99",
	"OmegaX_22",
	"LunarGamer_",
	"DragonX_007",
	"MysticG_22",
	"ShadowKnightG",
	"AtomicKnightX",
	"BlazeGamer_",
	"CosmicKnightX",
	"ChronoG_22",
	"DragonRiderXx",
	"ElectricKnight",
	"GalacticKnight",
	"InfinityXx_22",
	"IronKnightG_",
	"LunarG_007",
	"NeoX_99",
	"NightmareGamer",
	"OmegaKnight_",
	"PhantomXx_22",
	"RetroG_99",
	"SamuraiKnightX",
	"SonicKnightG_",
	"StormX_007",
	"ThunderKnightG",
	"TitaniumG_22",
	"WarriorXx_22",
	"AlphaGamerX_",
	"BlizzardKnightX",
	"ChronoX_22",
	"CyborgG_22",
	"DragonG_22",
	"ElectricXx_99",
	"GalacticXx_22",
	"IronG_007",
	"LunarKnightG",
	"MysticalG_99",
	"NightmareKnight",
	"OmegaXx_22",
	"PhantomKnightX",
	"RagingGamer_",
	"SamuraiX_99",
	"ShadowXx_22",
	"SonicG_22",
	"StormKnightG_",
	"TitaniumX_22",
	"WarriorKnightX",
	"ChronoGamerG_",
	"DragonKnightXx",
	"ElectricG_99",
	"GalacticG_007",
	"InfinityG_99",
	"IronKnightX_",
	"LunarXx_99",
	"MysticalX_22",
	"NightmareXx_",
	"OmegaKnightG_",
	"PhantomG_22",
	"RetroKnightX_",
	"SamuraiG_007",
	"ShadowKnightXx",
	"SonicXx_22",
	"StormG_007",
	"ThunderX_22",
	"TitaniumKnightX",
	"WarriorG_99",
	"AlphaGamer_22",
	"BlazeX_22",
	"ChronoKnightG",
	"CyberKnightX",
	"DragonRiderG_",
	"ElectricGamerX",
	"GalacticKnightX",
	"InfinityGamer_",
	"IronX_22",
	"LunarGamerX_",
	"MysticKnightG_",
	"NightmareG_22",
	"OmegaX_99",
	"PhantomKnightG",
	"RagingX_22",
	"SamuraiXx_99",
	"ShadowGamer_",
	"SonicKnightXx",
	"StormX_22",
	"ThunderKnightXx",
	"TitaniumGamer",
	"WarriorX_99",
	"BlizzardXx_22",
	"ChronoGamerKnight",
	"CyborgKnight_",
	"DragonXx_22",
	"ElectricKnightX",
	"GalacticX_007",
	"InfinityKnightX",
	"IronGamerX_",
	"LunarKnightXx",
	"MysticalGamer",
	"NightmareX_99",
	"OmegaKnightXx",
	"PhantomG_99",
	"RetroXx_99",
}