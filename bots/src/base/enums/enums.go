package enums

type GAME_EVENT byte

const (
	GAME_EVENT_UNKNOWN GAME_EVENT = iota
	GAME_EVENT_EAT_DOTS
	GAME_EVENT_EAT_BLOB
	GAME_EVENT_EAT_SMBH
	GAME_EVENT_BLOB_EXPLODE
	GAME_EVENT_BLOB_LOST
	GAME_EVENT_EJECT
	GAME_EVENT_SPLIT
	GAME_EVENT_RECOMBINE
	GAME_EVENT_TIMER_WARNING
	GAME_EVENT_CTF_SCORE
	GAME_EVENT_CTF_FLAG_RETURNED
	GAME_EVENT_CTF_FLAG_STOLEN
	GAME_EVENT_CTF_FLAG_DROPPED
	GAME_EVENT_ACHIEVEMENT_EARNED
	GAME_EVENT_XP_GAINED
	GAME_EVENT_UNUSED_2
	GAME_EVENT_XP_SET
	GAME_EVENT_DQ_SET
	GAME_EVENT_DQ_COMPLETED
	GAME_EVENT_DQ_PROGRESS
	GAME_EVENT_EAT_SERVER_BLOB
	GAME_EVENT_EAT_SPECIAL_OBJECTS
	GAME_EVENT_SO_SET
	GAME_EVENT_LEVEL_UP
	GAME_EVENT_ARENA_RANK_ACHIEVED
	GAME_EVENT_DOM_CP_LOST
	GAME_EVENT_DOM_CP_GAINED
	GAME_EVENT_UNUSED_1
	GAME_EVENT_CTF_GAINED
	GAME_EVENT_GAME_OVER
	GAME_EVENT_BLOB_STATUS
	GAME_EVENT_TELEPORT
	GAME_EVENT_SHOOT
	GAME_EVENT_CLAN_WAR_WON
	GAME_EVENT_PLASMA_REWARD
	GAME_EVENT_EMOTE
	GAME_EVENT_END_MISSION
	GAME_EVENT_XP_GAINED_2
	GAME_EVENT_EAT_CAKE
	GAME_EVENT_COIN_COUNT
	GAME_EVENT_CLEAR_EFFECTS
	GAME_EVENT_SPEED
	GAME_EVENT_TRICK
	GAME_EVENT_DESTROY_ASTEROID
	GAME_EVENT_ACCOLADE
	GAME_EVENT_INVIS
	GAME_EVENT_KILLED_BY
	GAME_EVENT_RADIATION_CLOUD
	GAME_EVENT_CHARGE
	GAME_EVENT_LP_COUNT
	GAME_EVENT_BR_BOUNDS
	GAME_EVENT_MINIMAP
	GAME_EVENT_RLGL_DEATH
	GAME_EVENT_RLGL_STATE
)

type CLAN_RANK byte

const (
	CLAN_RANK_INVALID CLAN_RANK = iota
	CLAN_RANK_MEMBER
	CLAN_RANK_ADMIN
	CLAN_RANK_LEADER
	CLAN_RANK_ELDER
	CLAN_RANK_DIAMOND
	CLAN_RANK_INITIATE
)

type CLICK_TYPE byte

const (
	CT_NORMAL CLICK_TYPE = iota
	CT_AUTO
	CT_ULTRA
)

type COLOR_CYCLE byte

const (
	COLOR_CYCLE_NONE COLOR_CYCLE = iota
	COLOR_CYCLE_COLOR_CYCLE_SLOW
	COLOR_CYCLE_COLOR_CYCLE_FAST
	COLOR_CYCLE_RAINBOW_HORIZONTAL_SLOW
	COLOR_CYCLE_RAINBOW_HORIZONTAL_FAST
	COLOR_CYCLE_RAINBOW_VERTICAL_SLOW
	COLOR_CYCLE_RAINBOW_VERTICAL_FAST
)

type CONNECT_RESULT byte

const (
	CONNECT_RESULT_SUCCESS CONNECT_RESULT = iota
	CONNECT_RESULT_GAME_NOT_FOUND
	CONNECT_RESULT_UNKNOWN
	CONNECT_RESULT_ACCOUNT_ALREADY_SIGNED_IN
)

type CONNECTION_STATE byte

const (
	CONNECTION_STATE_DISCONNECTED CONNECTION_STATE = iota
	CONNECTION_STATE_CONNECTING
	CONNECTION_STATE_CONNECTED
)

type DIFFICULTY byte

const (
	DIFFICULTY_EASY DIFFICULTY = iota
	DIFFICULTY_MEDIUM
	DIFFICULTY_HARD
	DIFFICULTY_IMPOSSIBLE
)

type GAME_MODE byte

const (
	GAME_MODE_FFA GAME_MODE = iota
	GAME_MODE_FFA_TIME
	GAME_MODE_TEAMS
	GAME_MODE_TEAMS_TIME
	GAME_MODE_CTF
	GAME_MODE_SURVIVAL
	GAME_MODE_SOCCER
	GAME_MODE_FFA_CLASSIC
	GAME_MODE_DOMINATION
	GAME_MODE_FFA_ULTRA
	GAME_MODE_ZOMBIE_APOCALYPSE // NOTE: renamed from "ZA"
	GAME_MODE_PAINT
	GAME_MODE_TEAM_DEATHMATCH
	GAME_MODE_X
	GAME_MODE_X2
	GAME_MODE_X3
	GAME_MODE_X4
	GAME_MODE_X5
	GAME_MODE_SPLIT_16X
	GAME_MODE_X6
	GAME_MODE_X7
	GAME_MODE_CAMPAIGN
	GAME_MODE_ROYALEDUO
	GAME_MODE_X8
	GAME_MODE_TRICK_MODE  // NOTE: renamed from "X9"
	GAME_MODE_PLASMA_HUNT // NOTE: renamed from "X10"
	GAME_MODE_X11
	GAME_MODE_X12
	GAME_MODE_X13
	GAME_MODE_X14
	GAME_MODE_X15
	GAME_MODE_X16
	GAME_MODE_X17
	GAME_MODE_X18
	GAME_MODE_X19
	GAME_MODE_CRAZY_SPLIT
	GAME_MODE_BATTLE_ROYALE
	GAME_MODE_X20
	GAME_MODE_X21
	GAME_MODE_MEGA_SPLIT
)

func AllGameModes() []string {
	return []string{
		"FFA",
		"FFA Time",
		"Teams",
		"Teams Time",
		"CTF",
		"Survival",
		"Soccer",
		"FFA Classic",
		"Domination",
		"FFA Ultra",
		"Zombie Apocalypse",
		"Paint",
		"Team Deathmatch",
		"X",
		"X2",
		"X3",
		"X4",
		"X5",
		"Split 16x",
		"X6",
		"X7",
		"Campaign",
		"Royale Duo",
		"X8",
		"Trick Mode",
		"Plasma Hunt",
		"X11",
		"X12",
		"X13",
		"X14",
		"X15",
		"X16",
		"X17",
		"X18",
		"X19",
		"Crazy Split",
		"Battle Royale",
		"X20",
		"X21",
		"Mega Split",
	}
}

func FromNormGameModeToEnum(norm *string) GAME_MODE {
	switch *norm {
	case "FFA":
		return GAME_MODE_FFA
	case "FFA Time":
		return GAME_MODE_FFA_TIME
	case "Teams":
		return GAME_MODE_TEAMS
	case "Teams Time":
		return GAME_MODE_TEAMS_TIME
	case "CTF":
		return GAME_MODE_CTF
	case "Survival":
		return GAME_MODE_SURVIVAL
	case "Soccer":
		return GAME_MODE_SOCCER
	case "FFA Classic":
		return GAME_MODE_FFA_CLASSIC
	case "Domination":
		return GAME_MODE_DOMINATION
	case "FFA Ultra":
		return GAME_MODE_FFA_ULTRA
	case "Zombie Apocalypse":
		return GAME_MODE_ZOMBIE_APOCALYPSE
	case "Paint":
		return GAME_MODE_PAINT
	case "Team Deathmatch":
		return GAME_MODE_TEAM_DEATHMATCH
	case "X":
		return GAME_MODE_X
	case "X2":
		return GAME_MODE_X2
	case "X3":
		return GAME_MODE_X3
	case "X4":
		return GAME_MODE_X4
	case "X5":
		return GAME_MODE_X5
	case "Split 16x":
		return GAME_MODE_SPLIT_16X
	case "X6":
		return GAME_MODE_X6
	case "X7":
		return GAME_MODE_X7
	case "Campaign":
		return GAME_MODE_CAMPAIGN
	case "Royale Duo":
		return GAME_MODE_ROYALEDUO
	case "X8":
		return GAME_MODE_X8
	case "Trick Mode":
		return GAME_MODE_TRICK_MODE
	case "Plasma Hunt":
		return GAME_MODE_PLASMA_HUNT
	case "X11":
		return GAME_MODE_X11
	case "X12":
		return GAME_MODE_X12
	case "X13":
		return GAME_MODE_X13
	case "X14":
		return GAME_MODE_X14
	case "X15":
		return GAME_MODE_X15
	case "X16":
		return GAME_MODE_X16
	case "X17":
		return GAME_MODE_X17
	case "X18":
		return GAME_MODE_X18
	case "X19":
		return GAME_MODE_X19
	case "Crazy Split":
		return GAME_MODE_CRAZY_SPLIT
	case "Battle Royale":
		return GAME_MODE_BATTLE_ROYALE
	case "X20":
		return GAME_MODE_X20
	case "X21":
		return GAME_MODE_X21
	case "Mega Split":
		return GAME_MODE_MEGA_SPLIT
	default:
		return GAME_MODE_FFA
	}
}

type FONT byte

const (
	FONT_DEFAULT FONT = iota
	FONT_XXRAYTID
	FONT_XTRUSION
	FONT_XXON
	FONT_XEFUS
	FONT_XENOPHOBIA
	FONT_XENOWORT
	FONT_CENOBYTE
	FONT_NM_HERO
	FONT_SKIN_xmas
	FONT_XLINES
	FONT_XEROX_MALFUNCTION
	FONT_KAUSHAN_SCRIPT
	FONT_BALL
	FONT_LARSON
	FONT_SUPERHET
	FONT_GETTHEME
	FONT_DEPHUN2
	FONT_CHRISTMAS
	FONT_FIRE
	FONT_BEYNO
	FONT_KINGTHINGS
)

type SPELL_TYPE byte

const (
	SPELL_TYPE_FREEZE SPELL_TYPE = iota
	SPELL_TYPE_POISON
	SPELL_TYPE_BOMB
	SPELL_TYPE_SHOCK
	SPELL_TYPE_SPEED
	SPELL_TYPE_SHIELD
	SPELL_TYPE_RECOMBINE
	SPELL_TYPE_HEAL
	SPELL_TYPE_ATTRACTOR
	SPELL_TYPE_HOOK
	SPELL_TYPE_BLIND
	SPELL_TYPE_RNG
	SPELL_TYPE_TP
	SPELL_TYPE_SWAP
	SPELL_TYPE_PUSH
	SPELL_TYPE_PHASE
	SPELL_TYPE_UNUSED_1
	SPELL_TYPE_UNUSED_2
	SPELL_TYPE_MAGNET
	SPELL_TYPE_SHROOM
	SPELL_TYPE_CLONE
	SPELL_TYPE_RADIATION
	SPELL_TYPE_MINIMAP
	SPELL_TYPE_UNKNOWN
)

type HOLE_TYPE byte

const (
	HOLE_TYPE_NORMAL HOLE_TYPE = iota
	HOLE_TYPE_SUPERMASSIVE
	HOLE_TYPE_TELEPORT
	HOLE_TYPE_NEBU
)

func GoodHoles() []HOLE_TYPE {
	return []HOLE_TYPE{HOLE_TYPE_NEBU, HOLE_TYPE_SUPERMASSIVE}
}

type SPELL_STATUS byte

const (
	SPELL_STATUS_DEAD SPELL_STATUS = iota
	SPELL_STATUS_INACTIVE
	SPELL_STATUS_ACTIVE
)

type PACKET_TYPE byte

const (
	PACKET_TYPE_INVALID PACKET_TYPE = iota
	PACKET_TYPE_CONNECT_RESULT_2
	PACKET_TYPE_CONTROL
	PACKET_TYPE_KEEP_ALIVE
	PACKET_TYPE_INVALIDATE_CLIENT
	PACKET_TYPE_START_GAME_INTERNAL
	PACKET_TYPE_CONNECT_REQUEST
	PACKET_TYPE_DISCONNECT
	PACKET_TYPE_GAME_CHAT_MESSAGE
	PACKET_TYPE_CLAN_CHAT_MESSAGE
	PACKET_TYPE_JOIN_REQUEST
	PACKET_TYPE_JOIN_RESULT
	PACKET_TYPE_TTL_REFRESH_RESPONSE_INTERNAL
	PACKET_TYPE_SHUTDOWN_NODE_INTERNAL
	PACKET_TYPE_SET_GS_ADDR
	PACKET_TYPE_CLIENT_PREFERENCES
	PACKET_TYPE_SPECTATE_CHANGE
	PACKET_TYPE_CLAN_WAR_LIST_REQUEST
	PACKET_TYPE_CLAN_WAR_LIST_RESULT
	PACKET_TYPE_CLAN_WAR_NOTIFICATION
	PACKET_TYPE_TOP_SCORES
	PACKET_TYPE_SERVER_SHUTDOWN_WARNING
	PACKET_TYPE_GAME_UPDATE
	PACKET_TYPE_GROUP_LOBBY_LIST_REQUEST
	PACKET_TYPE_GROUP_LOBBY_LIST_RESULT
	PACKET_TYPE_PUBLIC_CHAT_MESSAGE
	PACKET_TYPE_ADMIN_INTERNAL
	PACKET_TYPE_GROUP_LOBBY_CREATE_REQUEST
	PACKET_TYPE_GROUP_LOBBY_CREATE_RESULT
	PACKET_TYPE_GROUP_LOBBY_JOIN_REQUEST
	PACKET_TYPE_GROUP_LOBBY_JOIN_RESULT
	PACKET_TYPE_GROUP_LOBBY_UPDATE
	PACKET_TYPE_GROUP_LOBBY_LEAVE
	PACKET_TYPE_ARENA_LIST_REQUEST
	PACKET_TYPE_CLIENT_PREFERENCES_INTERNAL
	PACKET_TYPE_GAME_CRASH_INTERNAL
	PACKET_TYPE_PRIVATE_CHAT_MESSAGE
	PACKET_TYPE_ARENA_LEAVE_QUEUE_REQUEST
	PACKET_TYPE_REMOVE_GAME_INTERNAL
	PACKET_TYPE_GROUP_LOBBY_WARN
	PACKET_TYPE_ENTER_GAME_REQUEST
	PACKET_TYPE_ENTER_GAME_RESULT
	PACKET_TYPE_PLAYER_SESSION_STATS_UPDATE_INTERNAL
	PACKET_TYPE_PLAYER_WS_ACCOUNT_UPDATE_INTERNAL
	PACKET_TYPE_ACCOUNT_STATUS_REQUEST
	PACKET_TYPE_ACCOUNT_STATUS_RESULT
	PACKET_TYPE_FRIEND_CHAT_MESSAGE
	PACKET_TYPE_CLIENT_STATUS_CHANGE_REQUEST
	PACKET_TYPE_CLIENT_STATUS_CHANGE_RESULT
	PACKET_TYPE_CLAN_WAR_CONTROL
	PACKET_TYPE_CLAN_WAR_UPDATE
	PACKET_TYPE_ARENA_LIST_RESULT
	PACKET_TYPE_ADMIN_INTERNAL2
	PACKET_TYPE_NODE_RESET_REQUEST_INTERNAL
	PACKET_TYPE_CLAN_WAR_RESULT_INTERNAL
	PACKET_TYPE_CLAN_WAR_FORFEIT_INTERNAL
	PACKET_TYPE_SPECTATE_GAME_REQUEST
	PACKET_TYPE_GET_PLAYER_STATS_INTERNAL
	PACKET_TYPE_ARENA_QUEUE_REQUEST
	PACKET_TYPE_ARENA_STATUS
	PACKET_TYPE_ADMIN_INTERNAL3
	PACKET_TYPE_ARENA_RESULT_INTERNAL
	PACKET_TYPE_ADMIN_INTERNAL4
	PACKET_TYPE_TEAM_ARENA_RESULT_INTERNAL
	PACKET_TYPE_TEAM_ARENA_STATUS_RESULT
	PACKET_TYPE_TEAM_ARENA_STATUS_REQUEST
	PACKET_TYPE_TEAM_ARENA_LIST_REQUEST
	PACKET_TYPE_TEAM_ARENA_LIST_RESULT
	PACKET_TYPE_TEAM_ARENA_QUEUE_REQUEST
	PACKET_TYPE_TEAM_ARENA_LEAVE_QUEUE_REQEUST
	PACKET_TYPE_TEAM_ARENA_UPDATE
	PACKET_TYPE_CLAN_HOUSE_UPDATE_INTERNAL
	PACKET_TYPE_ADMIN_INTERNAL5
	PACKET_TYPE_CLAN_HOUSE_UPDATE_INTERNAL2
	PACKET_TYPE_NODE_CONNECT_REQUEST_INTERNAL
	PACKET_TYPE_GAME_DATA
	PACKET_TYPE_CHALLENGE
	PACKET_TYPE_CHALLENGE_RESULT
	PACKET_TYPE_FWD_TO_CLIENT_INTERNAL
	PACKET_TYPE_TTL_REFRESH_REQUEST_INTERNAL
	PACKET_TYPE_CONNECT_REQUEST_2
	PACKET_TYPE_CONNECT_RESULT
	PACKET_TYPE_ADMIN_INTERNAL6
	PACKET_TYPE_CLAN_HOUSE_UPDATE_INTERNAL3
	PACKET_TYPE_TOURNEY_LIST_REQUEST
	PACKET_TYPE_TOURNEY_LIST_RESULT
	PACKET_TYPE_TOURNEY_ACTION
	PACKET_TYPE_TOURNEY_MATCH_RESULT_INTERNAL
	PACKET_TYPE_TOURNEY_START_INTERNAL
	PACKET_TYPE_TOURNEY_STATUS_UPDATE
	PACKET_TYPE_ADMIN_INTERNAL7
	PACKET_TYPE_MUTE_INTERNAL
	PACKET_TYPE_JOINED_GAME_INTERNAL
	PACKET_TYPE_CLAN_HOUSE_UPDATE_INTERNAL4
	PACKET_TYPE_CLAN_HOUSE_CONFIG
	PACKET_TYPE_INVITE
	PACKET_TYPE_DESIRED_DUO_PARTNER
	PACKET_TYPE_EMOTE_REQUEST
	PACKET_TYPE_UDP_KEEPALIVE
	PACKET_TYPE_GROUP_CHAT_CREATE_REQUEST
	PACKET_TYPE_GROUP_CHAT_JOIN_REQUEST
	PACKET_TYPE_GROUP_CHAT_LEAVE_REQUEST
	PACKET_TYPE_GROUP_CHAT_RESULT
	PACKET_TYPE_GROUP_CHAT_STATUS
	PACKET_TYPE_GROUP_CHAT_MESSAGE
	PACKET_TYPE_SESSION_STATS
	PACKET_TYPE_ACCOLADE
	PACKET_TYPE_VOICE_CONTROL
	PACKET_TYPE_VOICE_DATA
	PACKET_TYPE_MINIMAP_UPDATE
	PACKET_TYPE_GAME_STOP_INTERNAL
	PACKET_TYPE_BATTLE_ROYALE_ACTION
	PACKET_TYPE_BATTLE_ROYALE_LIST_REQUEST
	PACKET_TYPE_BATTLE_ROYALE_LIST_RESULT
	PACKET_TYPE_BATTLE_ROYALE_STATUS_UPDATE
	PACKET_TYPE_BATTLE_ROYALE_RESULT_INTERNAL
	PACKET_TYPE_ADMIN_INTERNAL8
	PACKET_TYPE_PING_MESSAGE
	PACKET_TYPE_CONNECT_REQUEST_3
	PACKET_TYPE_ARENA_CD_INTERNAL
)

type JOIN_RESULT byte

const (
	JOIN_RESULT_SUCCESS = iota
	JOIN_RESULT_NAME_TAKEN
	JOIN_RESULT_NAME_INVALID
	JOIN_RESULT_FULL
	JOIN_RESULT_GAME_NOT_FOUND
	JOIN_RESULT_FRIEND_NOT_FOUND
	JOIN_RESULT_UNKNOWN_ERROR
	JOIN_RESULT_DIED_THIS_ROUND
	JOIN_RESULT_CLAN_WAR_NOT_FOUND
	JOIN_RESULT_CLAN_NOT_FOUND
	JOIN_RESULT_ACCOUNT_NOT_FOUND
	JOIN_RESULT_LACK_PERMISSION
	JOIN_RESULT_REQUEST_TIMED_OUT
	JOIN_RESULT_YOU_ARE_SPECTATING
	JOIN_RESULT_PLEASE_WAIT
	JOIN_RESULT_IS_ARENA
	JOIN_RESULT_ACCOUNT_IN_USE
	JOIN_RESULT_UPDATE_AVAILABLE
	JOIN_RESULT_INVALID_TOKEN
	JOIN_RESULT_BANNED
	JOIN_RESULT_NOT_SIGNED_IN
	JOIN_RESULT_TOURNAMENTS_DISABLED
	JOIN_RESULT_MUTED
	JOIN_RESULT_FRIEND_ALREADY_TEAMED
	JOIN_RESULT_GROUP_NOT_FOUND
	JOIN_RESULT_COMP_BANNED
	JOIN_RESULT_QUEUE_POSITION_UPDATE
	JOIN_RESULT_CHAT_BANNED
	JOIN_RESULT_KICKED
	JOIN_RESULT_INCOMPATIBLE_VERSION
)

func (j JOIN_RESULT) String() string {
	switch j {
	case JOIN_RESULT_SUCCESS:
		return "SUCCESS"
	case JOIN_RESULT_NAME_TAKEN:
		return "NAME_TAKEN"
	case JOIN_RESULT_NAME_INVALID:
		return "NAME_INVALID"
	case JOIN_RESULT_FULL:
		return "FULL"
	case JOIN_RESULT_GAME_NOT_FOUND:
		return "GAME_NOT_FOUND"
	case JOIN_RESULT_FRIEND_NOT_FOUND:
		return "FRIEND_NOT_FOUND"
	case JOIN_RESULT_UNKNOWN_ERROR:
		return "UNKNOWN_ERROR"
	case JOIN_RESULT_DIED_THIS_ROUND:
		return "DIED_THIS_ROUND"
	case JOIN_RESULT_CLAN_WAR_NOT_FOUND:
		return "CLAN_WAR_NOT_FOUND"
	case JOIN_RESULT_CLAN_NOT_FOUND:
		return "CLAN_NOT_FOUND"
	case JOIN_RESULT_ACCOUNT_NOT_FOUND:
		return "ACCOUNT_NOT_FOUND"
	case JOIN_RESULT_LACK_PERMISSION:
		return "LACK_PERMISSION"
	case JOIN_RESULT_REQUEST_TIMED_OUT:
		return "REQUEST_TIMED_OUT"
	case JOIN_RESULT_YOU_ARE_SPECTATING:
		return "YOU_ARE_SPECTATING"
	case JOIN_RESULT_PLEASE_WAIT:
		return "PLEASE_WAIT"
	case JOIN_RESULT_IS_ARENA:
		return "IS_ARENA"
	case JOIN_RESULT_ACCOUNT_IN_USE:
		return "ACCOUNT_IN_USE"
	case JOIN_RESULT_UPDATE_AVAILABLE:
		return "UPDATE_AVAILABLE"
	case JOIN_RESULT_INVALID_TOKEN:
		return "INVALID_TOKEN"
	case JOIN_RESULT_BANNED:
		return "BANNED"
	case JOIN_RESULT_NOT_SIGNED_IN:
		return "NOT_SIGNED_IN"
	case JOIN_RESULT_TOURNAMENTS_DISABLED:
		return "TOURNAMENTS_DISABLED"
	case JOIN_RESULT_MUTED:
		return "MUTED"
	case JOIN_RESULT_FRIEND_ALREADY_TEAMED:
		return "FRIEND_ALREADY_TEAMED"
	case JOIN_RESULT_GROUP_NOT_FOUND:
		return "GROUP_NOT_FOUND"
	case JOIN_RESULT_COMP_BANNED:
		return "COMP_BANNED"
	case JOIN_RESULT_QUEUE_POSITION_UPDATE:
		return "QUEUE_POSITION_UPDATE"
	case JOIN_RESULT_CHAT_BANNED:
		return "CHAT_BANNED"
	case JOIN_RESULT_KICKED:
		return "KICKED"
	case JOIN_RESULT_INCOMPATIBLE_VERSION:
		return "INCOMPATIBLE_VERSION"
	default:
		return "UNKNOWN"
	}
}

type PROFILE_VISIBILITY byte

const (
	PROFILE_VISIBILITY_ONLINE PROFILE_VISIBILITY = iota
	PROFILE_VISIBILITY_APPEAR_OFFLINE
	PROFILE_VISIBILITY_HIDDEN
	PROFILE_VISIBILITY_DND
)

type SERVER string

const (
	SERVER_US_EAST       SERVER = "45.56.113.95"
	SERVER_US_WEST       SERVER = "45.79.69.110"
	SERVER_EUROPE        SERVER = "172.105.248.252"
	SERVER_SOUTH_KOREA   SERVER = "158.247.231.199"
	SERVER_ASIA          SERVER = "139.162.49.99"
	SERVER_SOUTH_AMERICA SERVER = "216.238.98.140"
	SERVER_AUSTRALIA     SERVER = "45.79.238.85"
	SERVER_JAPAN         SERVER = "139.162.86.191"
	SERVER_MIDDLE_EAST   SERVER = "15.185.65.160"
	SERVER_SOUTH_AFRICA  SERVER = "139.84.232.5"
	SERVER_INDIA         SERVER = "194.195.115.5"
)

func AllServerIPs() []SERVER {
	return []SERVER{
		SERVER_US_EAST,
		SERVER_US_WEST,
		SERVER_EUROPE,
		SERVER_SOUTH_KOREA,
		SERVER_ASIA,
		SERVER_SOUTH_AMERICA,
		SERVER_AUSTRALIA,
		SERVER_JAPAN,
		SERVER_MIDDLE_EAST,
		SERVER_SOUTH_AFRICA,
		SERVER_INDIA,
	}
}

func AllServerNames() []string {
	return []string{
		"US East",
		"US West",
		"Europe",
		"South Korea",
		"Asia",
		"South America",
		"Australia",
		"Japan",
		"Middle East",
		"South Africa",
		"India",
	}
}

func ServerNameFromIP(ip SERVER) string {
	switch ip {
	case SERVER_US_EAST:
		return "US East"
	case SERVER_US_WEST:
		return "US West"
	case SERVER_EUROPE:
		return "Europe"
	case SERVER_SOUTH_KOREA:
		return "South Korea"
	case SERVER_ASIA:
		return "Asia"
	case SERVER_SOUTH_AMERICA:
		return "South America"
	case SERVER_AUSTRALIA:
		return "Australia"
	case SERVER_JAPAN:
		return "Japan"
	case SERVER_MIDDLE_EAST:
		return "Middle East"
	case SERVER_SOUTH_AFRICA:
		return "South Africa"
	case SERVER_INDIA:
		return "India"
	default:
		return "India"
	}
}

func ServerIPFromName(name *string) SERVER {
	switch *name {
	case "US East":
		return SERVER_US_EAST
	case "US West":
		return SERVER_US_WEST
	case "Europe":
		return SERVER_EUROPE
	case "South Korea":
		return SERVER_SOUTH_KOREA
	case "Asia":
		return SERVER_ASIA
	case "South America":
		return SERVER_SOUTH_AMERICA
	case "Australia":
		return SERVER_AUSTRALIA
	case "Japan":
		return SERVER_JAPAN
	case "Middle East":
		return SERVER_MIDDLE_EAST
	case "South Africa":
		return SERVER_SOUTH_AFRICA
	case "India":
		return SERVER_INDIA
	default:
		return SERVER_INDIA
	}
}

type ITEM_TYPE byte

const (
	ITEM_TYPE_PUNPKIN ITEM_TYPE = iota
	ITEM_TYPE_SNOWFLAKE
	ITEM_TYPE_HEART
	ITEM_TYPE_LEAF
	ITEM_TYPE_BIGDOT
	ITEM_TYPE_COIN
	ITEM_TYPE_PRESENT
	ITEM_TYPE_BEAD
	ITEM_TYPE_EGG
	ITEM_TYPE_RAINDROP
	ITEM_TYPE_NEBULA
	ITEM_TYPE_CANDY
	ITEM_TYPE_SUN
	ITEM_TYPE_MOON
	ITEM_TYPE_NOTE
	ITEM_TYPE_CAKE_PLASMA
	ITEM_TYPE_CAKE_XP
)

/*
TODO: create string representation of item types
example:
var (
    capabilitiesMap = map[string]Capability{
        "read":   Read,
        "create": Create,
        "update": Update,
        "delete": Delete,
        "list":   List,
    }
)
*/

type SKIN byte

const (
	SKIN_misc_none = iota
	SKIN_misc_8ball
	SKIN_misc_circuit
	SKIN_misc_glossyball
	SKIN_misc_lu
	SKIN_misc_matrix
	SKIN_misc_paint
	SKIN_misc_soccer
	SKIN_misc_warning
	SKIN_misc_yinyang
	SKIN_misc_doge
	SKIN_misc_waffle
	SKIN_misc_clock
	SKIN_misc_no_smoking
	SKIN_misc_pig
	SKIN_misc_turtle
	SKIN_misc_hell_doge
	SKIN_misc_polkadots
	SKIN_misc_polkadots2
	SKIN_misc_wheel
	SKIN_misc_compass
	SKIN_misc_sanik
	SKIN_misc_radiation
	SKIN_misc_radar
	SKIN_misc_creeper
	SKIN_misc_biohazard
	SKIN_misc_lambda
	SKIN_misc_tesla
	SKIN_misc_cheshire
	SKIN_misc_jack
	SKIN_misc_colors
	SKIN_misc_baseball
	SKIN_misc_basketball
	SKIN_misc_beachball
	SKIN_misc_doom
	SKIN_misc_euro
	SKIN_misc_eye
	SKIN_misc_grumpy
	SKIN_misc_lightning
	SKIN_misc_pball
	SKIN_misc_penny
	SKIN_misc_peppermint
	SKIN_misc_pineapple
	SKIN_misc_poker
	SKIN_misc_record
	SKIN_misc_wheatley
	SKIN_misc_clouds
	SKIN_misc_crop
	SKIN_misc_sauron
	SKIN_misc_ship
	SKIN_misc_wheel_car
	SKIN_misc_zerg
	SKIN_misc_rose
	SKIN_misc_cd
	SKIN_misc_chute
	SKIN_misc_astro
	SKIN_misc_chess
	SKIN_misc_sign
	SKIN_misc_trollface
	SKIN_misc_megustaxcf
	SKIN_misc_y
	SKIN_misc_dragonball
	SKIN_misc_stained
	SKIN_misc_stained2
	SKIN_misc_bauble
	SKIN_misc_camo
	SKIN_misc_bulb
	SKIN_misc_spiderweb
	SKIN_misc_rain
	SKIN_misc_chomp_1
	SKIN_misc_snow
	SKIN_scifi_mercury
	SKIN_scifi_venus
	SKIN_scifi_earth
	SKIN_scifi_mars
	SKIN_scifi_saturn
	SKIN_scifi_jupiter
	SKIN_scifi_neptune
	SKIN_scifi_moon
	SKIN_scifi_planet1
	SKIN_scifi_planet2
	SKIN_scifi_planet3
	SKIN_scifi_pluto
	SKIN_scifi_sun
	SKIN_scifi_planet
	SKIN_scifi_deathstar
	SKIN_scifi_pastry_cat
	SKIN_scifi_galaxy
	SKIN_scifi_dust
	SKIN_country_australia
	SKIN_country_austria
	SKIN_country_belgium
	SKIN_country_brazil
	SKIN_country_bulgaria
	SKIN_country_canada
	SKIN_country_china
	SKIN_country_finland
	SKIN_country_france
	SKIN_country_germany
	SKIN_country_greece
	SKIN_country_india
	SKIN_country_italy
	SKIN_country_japan
	SKIN_country_mexico
	SKIN_country_netherlands
	SKIN_country_norway
	SKIN_country_poland
	SKIN_country_romania
	SKIN_country_russia
	SKIN_country_south_africa
	SKIN_country_southkorea
	SKIN_country_spain
	SKIN_country_sweden
	SKIN_country_turkey
	SKIN_country_uk
	SKIN_country_ukraine
	SKIN_country_usa
	SKIN_country_columbia
	SKIN_country_ecuador
	SKIN_country_ireland
	SKIN_country_puerto_rico
	SKIN_country_argentina
	SKIN_country_denmark
	SKIN_country_egypt
	SKIN_country_peru
	SKIN_country_portugal
	SKIN_country_georgia
	SKIN_country_morocco
	SKIN_country_croatia
	SKIN_country_israel
	SKIN_country_pakistan
	SKIN_country_bosnia
	SKIN_country_chile
	SKIN_country_dr
	SKIN_country_hungary
	SKIN_country_philipines
	SKIN_country_venezuela
	SKIN_country_trinidad
	SKIN_country_costarica
	SKIN_country_scotlan
	SKIN_country_belaruse
	SKIN_country_serbia
	SKIN_country_slovakia
	SKIN_country_tunisia
	SKIN_country_iran
	SKIN_country_thailand
	SKIN_country_switzerland
	SKIN_country_algeria
	SKIN_country_newzealand
	SKIN_country_lithuania
	SKIN_country_jamaica
	SKIN_country_guatemala
	SKIN_country_lebanon
	SKIN_country_albania
	SKIN_country_macedonia
	SKIN_country_latvia
	SKIN_country_azerbaijan
	SKIN_country_estonia
	SKIN_country_czech
	SKIN_country_wales
	SKIN_country_armenia
	SKIN_country_cuba
	SKIN_country_england
	SKIN_country_kazakhstan
	SKIN_country_iceland
	SKIN_country_indonesia
	SKIN_country_panama
	SKIN_country_cyprus
	SKIN_country_moldova
	SKIN_country_montenegro
	SKIN_country_honduras
	SKIN_country_elsalvador
	SKIN_country_uruguay
	SKIN_country_iraq
	SKIN_country_saudiaarabia
	SKIN_misc_awesome
	SKIN_misc_pug
	SKIN_misc_panda
	SKIN_misc_biblethump
	SKIN_misc_saw
	SKIN_country_bolivia
	SKIN_country_jordan
	SKIN_misc_penguin
	SKIN_misc_heart
	SKIN_misc_circuit_2
	SKIN_misc_cookie
	SKIN_misc_hourglass
	SKIN_misc_mona
	SKIN_misc_pointer
	SKIN_misc_shield
	SKIN_scifi_dust_blue
	SKIN_scifi_dust_red
	SKIN_country_bahrain
	SKIN_misc_pizza
	SKIN_misc_burger
	SKIN_misc_donut
	SKIN_misc_dragon
	SKIN_misc_seal
	SKIN_misc_woofer
	SKIN_misc_spiral
	SKIN_misc_spiral_2
	SKIN_misc_round
	SKIN_misc_devil_alien
	SKIN_misc_comet
	SKIN_misc_lightning_ball
	SKIN_misc_foxy
	SKIN_scifi_plasma
	SKIN_scifi_comet
	SKIN_misc_tennis
	SKIN_misc_atom
	SKIN_misc_cat
	SKIN_misc_dandelion
	SKIN_misc_lion
	SKIN_misc_potatocorn
	SKIN_misc_ourboros
	SKIN_misc_discoball
	SKIN_special_pumpkin
	SKIN_special_pumpkin_tree
	SKIN_special_cat
	SKIN_special_town
	SKIN_special_cemetary
	SKIN_country_paraguay
	SKIN_misc_butterfly
	SKIN_country_nicaragua
	SKIN_misc_bacteria
	SKIN_misc_dinosaur
	SKIN_misc_dinosaur_2
	SKIN_misc_fire_ice
	SKIN_misc_minion
	SKIN_achieve_cw_1
	SKIN_misc_spirit_wolf
	SKIN_country_qatar
	SKIN_country_slovenia
	SKIN_misc_bomb
	SKIN_scifi_universe
	SKIN_scifi_prism_dust
	SKIN_country_malaysia
	SKIN_country_syria
	SKIN_scifi_callisto
	SKIN_special_leaf
	SKIN_special_acorn
	SKIN_special_cornucopia
	SKIN_special_harvest
	SKIN_special_autumn_tree
	SKIN_misc_torus
	SKIN_misc_orange
	SKIN_misc_static
	SKIN_misc_buridog
	SKIN_scifi_vortex
	SKIN_scifi_line_sink
	SKIN_country_kuwait
	SKIN_misc_whirlpool
	SKIN_scifi_plasma_2
	SKIN_scifi_pink_hole
	SKIN_misc_turbine
	SKIN_scifi_happy_stars
	SKIN_misc_magma
	SKIN_misc_matrix_2
	SKIN_achieve_arena_10
	SKIN_achieve_arena_100
	SKIN_achieve_arena_1000
	SKIN_achieve_cw_10
	SKIN_achieve_cw_100
	SKIN_plasma_circles
	SKIN_plasma_f1topdown
	SKIN_plasma_face_1
	SKIN_plasma_face_2
	SKIN_plasma_face_3
	SKIN_plasma_face_4
	SKIN_plasma_face_5
	SKIN_plasma_face_6
	SKIN_plasma_face_7
	SKIN_plasma_face_8
	SKIN_plasma_face_9
	SKIN_plasma_happy_stars
	SKIN_plasma_peacock
	SKIN_plasma_plane
	SKIN_misc_SKIN_plasma_ball
	SKIN_plasma_reel
	SKIN_plasma_space
	SKIN_plasma_purpleplanetstars
	SKIN_plasma_pulsar
	SKIN_plasma_bulb_1
	SKIN_plasma_bulb_2
	SKIN_plasma_bulb_3
	SKIN_plasma_bulb_4
	SKIN_plasma_bulb_5
	SKIN_plasma_bulb_6
	SKIN_plasma_bulb_7
	SKIN_plasma_bulb_8
	SKIN_plasma_bulb_9
	SKIN_plasma_wormhole
	SKIN_plasma_coin
	SKIN_special_christmas_tree
	SKIN_special_christmas_wreath
	SKIN_special_santa
	SKIN_special_snow_globe
	SKIN_special_snow
	SKIN_special_gifts
	SKIN_plasma_clouds
	SKIN_plasma_zoom
	SKIN_plasma_clock
	SKIN_misc_beach
	SKIN_misc_hamster
	SKIN_plasma_glowstars
	SKIN_plasma_metal
	SKIN_plasma_neurons
	SKIN_plasma_shuttle
	SKIN_plasma_space_2
	SKIN_special_ice
	SKIN_special_iceberg
	SKIN_special_icicles
	SKIN_special_igloo
	SKIN_special_penguin
	SKIN_special_snowleaf
	SKIN_special_winter_tree
	SKIN_misc_fish
	SKIN_plasma_bouncing_ball
	SKIN_misc_fireball
	SKIN_cartoon_snail
	SKIN_cartoon_zebra
	SKIN_cartoon_hippo
	SKIN_cartoon_elephant
	SKIN_cartoon_cow
	SKIN_cartoon_goat
	SKIN_cartoon_bear
	SKIN_cartoon_mouse
	SKIN_cartoon_kangaroo
	SKIN_cartoon_dog
	SKIN_cartoon_owl
	SKIN_cartoon_cat
	SKIN_cartoon_beaver
	SKIN_cartoon_sheep_2
	SKIN_cartoon_penguin
	SKIN_cartoon_sheep
	SKIN_country_uae
	SKIN_level_earth
	SKIN_misc_checkers
	SKIN_country_yemen
	SKIN_level_phoenix_fire
	SKIN_level_phoenix_electric
	SKIN_level_tapestry
	SKIN_level_scales
	SKIN_special_masks
	SKIN_special_mask_1
	SKIN_special_mask_2
	SKIN_special_feathers
	SKIN_special_mask_3
	SKIN_special_easter_garden
	SKIN_special_egg_basket
	SKIN_special_easter_bunny
	SKIN_special_eggs
	SKIN_special_eggs_transparent
	SKIN_special_flowers
	SKIN_special_canopy
	SKIN_special_grass
	SKIN_special_ladybug
	SKIN_special_butterfly
	SKIN_level_asteroids
	SKIN_level_dotty
	SKIN_level_leopard
	SKIN_level_fire_wings
	SKIN_plasma_tesselation
	SKIN_plasma_earth
	SKIN_plasma_firework
	SKIN_plasma_brain
	SKIN_special_desert_planet
	SKIN_special_space_cat
	SKIN_special_satellite
	SKIN_special_starfield
	SKIN_special_galaxy
	SKIN_level_dragon
	SKIN_level_bh
	SKIN_plasma_dish
	SKIN_plasma_skull
	SKIN_plasma_fireball
	SKIN_plasma_phoenix
	SKIN_plasma_unicorn
	SKIN_country_nepal
	SKIN_level_bamboo
	SKIN_level_bluepulse
	SKIN_level_fish
	SKIN_level_green_spinnie
	SKIN_level_tracks
	SKIN_misc_arcadego
	SKIN_misc_chicken
	SKIN_plasma_blue_spinnie
	SKIN_plasma_fawkes
	SKIN_plasma_fist
	SKIN_plasma_red_spinnies
	SKIN_plasma_shark
	SKIN_special_cinamon
	SKIN_special_cupcake
	SKIN_special_gummy
	SKIN_special_lollipop
	SKIN_special_watermelon
	SKIN_level_telephone
	SKIN_level_missile
	SKIN_level_wall
	SKIN_plasma_ufo
	SKIN_plasma_eye
	SKIN_plasma_ra
	SKIN_special_summer_canopy
	SKIN_special_oasis
	SKIN_special_sunset
	SKIN_special_wave
	SKIN_special_sun
	SKIN_country_haiti
	SKIN_level_spin_cats
	SKIN_level_bubbles
	SKIN_level_puzzle
	SKIN_plasma_dot_illusion
	SKIN_plasma_binary
	SKIN_plasma_chip_ship
	SKIN_special_moon_stars
	SKIN_special_night_sky
	SKIN_special_eiffel
	SKIN_special_aurora
	SKIN_special_earth_night
	SKIN_level_bolts
	SKIN_level_mandala
	SKIN_level_wormhole
	SKIN_misc_hole
	SKIN_misc_oreo
	SKIN_misc_rubiks
	SKIN_plasma_invaders
	SKIN_plasma_letters_a
	SKIN_plasma_letters_b
	SKIN_plasma_letters_c
	SKIN_plasma_letters_d
	SKIN_plasma_letters_e
	SKIN_plasma_letters_f
	SKIN_plasma_letters_g
	SKIN_plasma_letters_h
	SKIN_plasma_letters_i
	SKIN_plasma_letters_j
	SKIN_plasma_letters_k
	SKIN_plasma_letters_l
	SKIN_plasma_letters_m
	SKIN_plasma_letters_n
	SKIN_plasma_letters_o
	SKIN_plasma_letters_p
	SKIN_plasma_letters_q
	SKIN_plasma_letters_r
	SKIN_plasma_letters_s
	SKIN_plasma_letters_t
	SKIN_plasma_letters_u
	SKIN_plasma_letters_v
	SKIN_plasma_letters_w
	SKIN_plasma_letters_x
	SKIN_plasma_letters_y
	SKIN_plasma_letters_z
	SKIN_plasma_pickaxe
	SKIN_plasma_plasma
	SKIN_special_drum
	SKIN_special_guitar
	SKIN_special_horn
	SKIN_special_ocarina
	SKIN_special_piano
	SKIN_country_afghanistan
	SKIN_country_srilanka
	SKIN_achieve_arena_10000
	SKIN_misc_purple_smile
	SKIN_misc_purple_planet
	SKIN_misc_space
	SKIN_misc_pulsar
	SKIN_misc_blackhole
	SKIN_level_lock
	SKIN_level_eye
	SKIN_level_pencils
	SKIN_level_bat
	SKIN_level_cursors
	SKIN_achieve_cw_1000
	SKIN_level_galactic
	SKIN_plasma_bee
	SKIN_misc_duck
	SKIN_tuber_rouvel
	SKIN_tuber_gaara
	SKIN_tuber_johnny
	SKIN_tuber_agnosia
	SKIN_tuber_benitez
	SKIN_tuber_badboy
	SKIN_tuber_piedra
	SKIN_tuber_bruxo
	SKIN_tuber_john
	SKIN_tuber_kamikaze
	SKIN_tuber_gago
	SKIN_standard_cottoncandy
	SKIN_standard_jurassic
	SKIN_fb_brasil
	SKIN_country_kyrgyzstan
	SKIN_level_colorwheel
	SKIN_level_sunface
	SKIN_level_bee
	SKIN_plasma_spinster
	SKIN_plasma_pizzacutter
	SKIN_plasma_fossil
	SKIN_tuber_maro
	SKIN_tuber_maroo
	SKIN_tuber_blarp
	SKIN_tuber_rowdy
	SKIN_tuber_ghost
	SKIN_tuber_pronebulous
	SKIN_tuber_gokhan
	SKIN_misc_amoeba
	SKIN_misc_neb_es
	SKIN_misc_christmas
	SKIN_misc_neb
	SKIN_misc_panda_2
	SKIN_misc_cookie_2
	SKIN_level_shuriken
	SKIN_level_bow
	SKIN_level_copter
	SKIN_level_maze
	SKIN_level_arrow
	SKIN_level_eye_raptor
	SKIN_level_eyeball
	SKIN_level_shuriken_2
	SKIN_level_hammer
	SKIN_level_boo1
	SKIN_misc_bauble_2
	SKIN_misc_neon
	SKIN_tuber_celso
	SKIN_tuber_ken
	SKIN_tuber_zaygod
	SKIN_tuber_kings
	SKIN_level_unicorn
	SKIN_level_tiger
	SKIN_tuber_trex
	SKIN_level_smoke
	SKIN_level_shell
	SKIN_level_eye_2
	SKIN_achieve_shark
	SKIN_achieve_dab
	SKIN_achieve_pinwheel
	SKIN_misc_hearts
	SKIN_tuber_cheetah
	SKIN_tuber_neiker
	SKIN_tuber_fena23
	SKIN_tuber_koalaoso
	SKIN_tuber_a3
	SKIN_fb_espanol
	SKIN_level_moon_2
	SKIN_level_moon_1
	SKIN_level_green_embroider
	SKIN_level_spin_ribbon
	SKIN_level_nebgod
	SKIN_tuber_juangui
	SKIN_fb_mundial
	SKIN_tuber_swag
	SKIN_tuber_victor
	SKIN_plasma_swirls
	SKIN_plasma_yinyang
	SKIN_plasma_tank
	SKIN_level_fractal
	SKIN_level_hand
	SKIN_level_firewheel
	SKIN_level_fireplane
	SKIN_level_portal
	SKIN_level_firepower
	SKIN_level_turtle
	SKIN_achievement_bubble
	SKIN_level_liquorice
	SKIN_level_prism_flower
	SKIN_level_illusion
	SKIN_misc_illusion
	SKIN_tuber_marioluigi
	SKIN_tuber_lom
	SKIN_tuber_blacky
	SKIN_tuber_rodrick
	SKIN_level_shark
	SKIN_plasma_alien
	SKIN_level_bluesun
	SKIN_achieve_dot_gulper
	SKIN_plasma_wings
	SKIN_misc_lion_face
	SKIN_misc_cute
	SKIN_level_SKIN_plasma_comet
	SKIN_level_plasma
	SKIN_level_illusion_2
	SKIN_level_dragons
	SKIN_level_colorwheel_2
	SKIN_level_adze
	SKIN_level_leopard_2
	SKIN_level_fireflower
	SKIN_plasma_horse
	SKIN_level_bird
	SKIN_country_bangladesh
	SKIN_plasma_fidget
	SKIN_level_hydra
	SKIN_level_skull
	SKIN_level_trident
	SKIN_level_SKIN_plasma_ball_2
	SKIN_level_crab
	SKIN_level_spider
	SKIN_plasma_artifact
	SKIN_level_illusion_3
	SKIN_level_fidget
	SKIN_achieve_fidget
	SKIN_plasma_fidget2
	SKIN_level_star
	SKIN_level_cardio
	SKIN_level_ghost
	SKIN_level_grid
	SKIN_level_bulb
	SKIN_plasma_squid_1
	SKIN_plasma_colorful
	SKIN_level_cubes
	SKIN_level_cells
	SKIN_level_hole_2
	SKIN_level_hole_3
	SKIN_level_SKIN_plasma_sky
	SKIN_tuber_naffyt
	SKIN_tuber_taco
	SKIN_level_plasmastorm
	SKIN_level_plasmastorm_2
	SKIN_level_chakras
	SKIN_level_probe
	SKIN_level_glare
	SKIN_level_swisscheese
	SKIN_level_nebula
	SKIN_plasma_target
	SKIN_plasma_trifire
	SKIN_level_arrows
	SKIN_misc_surreal
	SKIN_level_bubbles_2
	SKIN_level_turbine_2
	SKIN_level_pulsar
	SKIN_level_shockwave
	SKIN_level_phoenix
	SKIN_level_jellyfish
	SKIN_plasma_blades
	SKIN_achieve_nebulator
	SKIN_achieve_nebulator_2
	SKIN_tuber_mezo
	SKIN_level_parabola
	SKIN_level_orborous
	SKIN_level_leaf
	SKIN_level_fan
	SKIN_level_rings
	SKIN_tuber_chan
	SKIN_misc_bullethole
	SKIN_plasma_whaleshark
	SKIN_level_astronaut
	SKIN_level_paw
	SKIN_level_fallingstars
	SKIN_level_blades
	SKIN_level_emeralddove
	SKIN_level_pixelheart
	SKIN_level_rainbow
	SKIN_level_fireturbine
	SKIN_level_radioplasma
	SKIN_level_firebird
	SKIN_level_fistbump
	SKIN_level_eartheye
	SKIN_level_void
	SKIN_level_lightning
	SKIN_level_atom
	SKIN_level_firedragon
	SKIN_level_colorplane
	SKIN_level_colorfireballs
	SKIN_level_plasmablades
	SKIN_level_spindle
	SKIN_level_concentric_rings
	SKIN_level_sunmoondog
	SKIN_level_fireatom
	SKIN_level_firewater
	SKIN_level_hole_4
	SKIN_level_cylinders
	SKIN_special_pumpkindrop
	SKIN_level_magni
	SKIN_special_jack
	SKIN_level_illusionpattern
	SKIN_special_hand
	SKIN_level_stealth
	SKIN_level_lightexplosion
	SKIN_level_atom_2
	SKIN_tuber_noob
	SKIN_tuber_icy
	SKIN_special_leaf_wreath
	SKIN_special_leaves
	SKIN_country_singapore
	SKIN_level_squares
	SKIN_level_stars
	SKIN_level_SKIN_plasma_arrow
	SKIN_xmas1
	SKIN_xmas2
	SKIN_xmas3
	SKIN_xmas4
	SKIN_xmas5
	SKIN_xmas6
	SKIN_xmas7
	SKIN_xmas8
	SKIN_xmas9
	SKIN_xmas10
	SKIN_xmas11
	SKIN_xmas12
	SKIN_xmas13
	SKIN_xmas14
	SKIN_xmas15
	SKIN_xmas16
	SKIN_xmas17
	SKIN_xmas18
	SKIN_xmas19
	SKIN_xmas20
	SKIN_xmas21
	SKIN_xmas22
	SKIN_xmas23
	SKIN_xmas24
	SKIN_xmas25
	SKIN_xmas26
	SKIN_xmas27
	SKIN_scifi_uranus
	SKIN_tuber_duck
	SKIN_tuber_oriente
	SKIN_tuber_hard
	SKIN_level_clown_hole
	SKIN_level_wobble
	SKIN_level_SKIN_cartoon_hole
	SKIN_level_neon_spinner
	SKIN_level_laser_1
	SKIN_st_andre
	SKIN_st_taco
	SKIN_st_mep
	SKIN_tuber_amandinha
	SKIN_tuber_cuchi
	SKIN_st_mueka
	SKIN_level_spiral
	SKIN_level_dragonskin
	SKIN_level_rgbbulb
	SKIN_level_psyflower
	SKIN_level_turbineflower
	SKIN_st_exdreamz
	SKIN_st_sk
	SKIN_level_rgb
	SKIN_level_illusion_4
	SKIN_level_plasmabeams
	SKIN_level_rgbplasma
	SKIN_level_whitehall
	SKIN_st_emre
	SKIN_level_meander
	SKIN_level_plasmadna
	SKIN_level_psiturbine
	SKIN_level_panda
	SKIN_level_glowydonut
	SKIN_level_emojis
	SKIN_special_egg
	SKIN_level_illusion_5
	SKIN_level_flowerspiral
	SKIN_level_squarespiral
	SKIN_level_plasmatile
	SKIN_level_arrowinder
	SKIN_level_firetwirl
	SKIN_level_rgball
	SKIN_level_rgbturbine
	SKIN_tuber_qzl
	SKIN_level_fly
	SKIN_level_purplehole
	SKIN_level_fireplant
	SKIN_level_balloon
	SKIN_level_greenstar
	SKIN_country_cambodia
	SKIN_level_face
	SKIN_level_bubblemouse
	SKIN_level_oscope
	SKIN_level_hangloose
	SKIN_misc_ribbon
	SKIN_level_purple_spinny
	SKIN_level_beachball
	SKIN_level_marble
	SKIN_level_sunball
	SKIN_level_ball
	SKIN_level_blank
	SKIN_level_jadeflower
	SKIN_level_retroswirl
	SKIN_plasma_doggy
	SKIN_level_mandala_2
	SKIN_level_square_spiral
	SKIN_level_wormhole_2
	SKIN_level_plantbubble
	SKIN_level_rocketship
	SKIN_level_mobius
	SKIN_level_redblue_spiral
	SKIN_tuber_revoltz
	SKIN_level_gif1
	SKIN_level_gif2
	SKIN_level_gif3
	SKIN_misc_grapefruit
	SKIN_level_earthstar
	SKIN_level_glassflower
	SKIN_level_faces
	SKIN_level_goldlattice
	SKIN_level_gif4
	SKIN_level_psysun
	SKIN_misc_galaxy
	SKIN_level_moltenarmor
	SKIN_level_paintsplosion
	SKIN_st_victor
	SKIN_tuber_moises
	SKIN_level_purplebubbles
	SKIN_level_rgbvortex
	SKIN_level_eyeflower
	SKIN_tuber_pesadelo
	SKIN_level_energy
	SKIN_misc_spacerose
	SKIN_plasma_deadmau5
	SKIN_level_gif5
	SKIN_st_index
	SKIN_level_magneticfield
	SKIN_level_gif6
	SKIN_level_fractal_2
	SKIN_level_boomerang
	SKIN_level_kittycat
	SKIN_misc_greenstar
	SKIN_level_purpleshell
	SKIN_level_pulse
	SKIN_level_purplepulsar
	pack_ghost
	pack_eye
	pack_skeleton
	pack_jack
	pack_halloween
	SKIN_level_trippy
	SKIN_level_look
	SKIN_level_gif7
	SKIN_level_gif8
	SKIN_misc_star_spiral
	SKIN_level_happystars
	SKIN_level_bluedisc
	SKIN_level_plasmayang
	SKIN_level_abstractturbine
	SKIN_level_gif9
	SKIN_standard_galaxyglobe
	SKIN_level_purplelightning
	SKIN_level_nebrings
	SKIN_level_gif10
	SKIN_level_metalmaze
	pack_SKIN_xmas
	pack_snow
	pack_santa
	pack_snowflake
	pack_candytown
	SKIN_standard_illusion
	SKIN_level_turbine_stars
	SKIN_level_red_vortex
	SKIN_level_waterbowl
	SKIN_standard_SKIN_xmas15
	SKIN_standard_newyear16
	SKIN_standard_SKIN_xmas16
	SKIN_standard_newyear17
	SKIN_standard_SKIN_xmas17
	SKIN_standard_newyear18
	SKIN_standard_SKIN_xmas18
	SKIN_standard_newyear19
	SKIN_level_gif11
	SKIN_standard_redfalls
	SKIN_level_energyshield
	SKIN_level_prismzoom
	SKIN_level_lightningball
	pack_vday
	pack_heart
	pack_cupid
	pack_bouquet
	pack_bloom
	pack_bow_1
	SKIN_standard_peacocktail
	SKIN_level_goldenring
	SKIN_level_purplephoenix
	SKIN_level_goldenring2
	SKIN_level_gif12
	SKIN_level_blueturbine
	SKIN_level_stardust
	SKIN_level_skullflower
	SKIN_standard_strawberry
	SKIN_tuber_agnozia
	SKIN_achieve_vet_1
	SKIN_achieve_vet_2
	SKIN_achieve_vet_3
	SKIN_achieve_vet_4
	SKIN_achieve_vet_5
	SKIN_level_ufos
	SKIN_level_nebunaut
	SKIN_level_rosey
	SKIN_level_fireflower2
	SKIN_level_bluestar
	SKIN_level_firehorse
	SKIN_level_gif13
	SKIN_plasma_SKIN_plasma_ball
	SKIN_tuber_venom
	SKIN_level_purplespiral
	SKIN_level_psiflower
	SKIN_level_firebirdy
	SKIN_standard_1
	SKIN_standard_2
	SKIN_standard_3
	SKIN_standard_4
	SKIN_standard_5
	SKIN_standard_6
	SKIN_standard_7
	SKIN_standard_8
	SKIN_standard_9
	SKIN_standard_10
	SKIN_standard_11
	SKIN_standard_12
	SKIN_standard_13
	SKIN_standard_14
	SKIN_standard_15
	SKIN_standard_16
	SKIN_standard_17
	SKIN_standard_18
	SKIN_standard_19
	SKIN_standard_20
	SKIN_standard_21
	SKIN_standard_22
	SKIN_standard_23
	SKIN_standard_24
	SKIN_standard_25
	SKIN_standard_26
	SKIN_standard_27
	SKIN_standard_28
	SKIN_standard_29
	SKIN_standard_30
	SKIN_level_fire_wheel
	SKIN_level_ice_wheel
	SKIN_level_water_wheel
	SKIN_level_moskull
	SKIN_level_rgbplane
	SKIN_level_rgb_turbine
	SKIN_level_golden_glow
	SKIN_level_puzzle_pieces
	SKIN_level_popo
	SKIN_level_gif14
	SKIN_standard_31
	SKIN_standard_32
	SKIN_standard_33
	SKIN_standard_34
	SKIN_standard_35
	SKIN_standard_36
	SKIN_standard_37
	SKIN_standard_38
	SKIN_standard_39
	SKIN_standard_40
	SKIN_standard_41
	SKIN_standard_42
	SKIN_standard_43
	SKIN_standard_44
	SKIN_standard_45
	SKIN_standard_46
	SKIN_standard_47
	SKIN_standard_48
	SKIN_standard_49
	SKIN_standard_50
	SKIN_st_ferdios
	SKIN_level_peacock
	SKIN_level_hadron
	SKIN_level_wavy
	SKIN_level_bloodturbine
	SKIN_level_gif15
	SKIN_standard_51
	SKIN_level_wavy_rgb
	SKIN_level_flickerskull
	SKIN_level_snowturbine
	SKIN_level_lavaturbine
	SKIN_level_grassturbine
	SKIN_level_trippypurp
	SKIN_level_trippyteal
	SKIN_level_plasmaswirl
	SKIN_achieve_gif16
	SKIN_level_purplefeathers
	SKIN_level_blueflowers
	SKIN_level_anchorage
	SKIN_standard_spaceface
	SKIN_level_neuroticnebula
	SKIN_level_digitalwaterfall
	SKIN_level_wobblespin
	SKIN_level_stainedfractal
	SKIN_country_vietnam
	SKIN_level_goldhalo
	SKIN_level_prismaticdisc
	SKIN_level_madavian
	SKIN_tuber_sten
	SKIN_tuber_snay
	SKIN_tuber_klezxvers
	SKIN_level_skullbones
	SKIN_level_colorstar
	SKIN_level_symmetry
	SKIN_level_entropy
	SKIN_tuber_badgirl
	SKIN_tuber_futurdroidbr
	SKIN_tuber_melennie
	SKIN_tuber_keeiz
	SKIN_tuber_koji
	SKIN_level_gif16
	SKIN_country_oman
	SKIN_tuber_danzin
	SKIN_level_dab
	SKIN_standard_plasmaball
	SKIN_level_firegalaxy
	SKIN_level_electrobauble
	SKIN_level_leopardskin
	SKIN_level_orborous2
	SKIN_level_green_plasma
	SKIN_level_electric_plasma
	SKIN_tuber_absalom
	SKIN_standard_52
	SKIN_standard_53
	SKIN_level_binary_spiral
	SKIN_level_sun
	SKIN_level_lightball
	SKIN_level_bluespark
	SKIN_level_beetle
	SKIN_level_eyes
	SKIN_standard_spacemarble
	SKIN_level_dandylion
	SKIN_level_shuttlecraft
	SKIN_level_redspark
	SKIN_level_pastel
	SKIN_achieve_vet_6
	SKIN_level_wavytwist
	SKIN_level_mouse_pointer
	SKIN_level_mouse_hourglass
	SKIN_plasma_noob
	SKIN_level_chicken
	SKIN_level_pizza
	SKIN_level_pearl
	SKIN_country_suriname
	SKIN_level_colorballs
	SKIN_standard_purplegalaxy
	SKIN_level_apple
	SKIN_level_firefootball
	SKIN_standard_54
	SKIN_level_pinkswirl
	SKIN_level_nautilus
	SKIN_level_spacesymbols
	SKIN_country_libya
	SKIN_country_mauritania
	SKIN_country_somalia
	SKIN_country_palestine
	SKIN_country_comoros
	SKIN_country_djibouti
	SKIN_country_sudan
	SKIN_level_tribalbird
	SKIN_level_darkpurplespinner
	SKIN_level_darkorangetesselate
	SKIN_country_equatorial_guinea
	SKIN_standard_robl
	SKIN_tuber_moodi
	SKIN_tuber_viper
	SKIN_achieve_vet_7
	SKIN_tuber_mbb
	SKIN_st_maja
	SKIN_st_cedric
)

func FromNormSkinToEnum(norm *string) SKIN {
	switch *norm {
	case "none":
		return SKIN_misc_none
	case "8ball":
		return SKIN_misc_8ball
	case "circuit":
		return SKIN_misc_circuit
	case "glossyball":
		return SKIN_misc_glossyball
	case "lu":
		return SKIN_misc_lu
	case "matrix":
		return SKIN_misc_matrix
	case "paint":
		return SKIN_misc_paint
	case "soccer":
		return SKIN_misc_soccer
	case "warning":
		return SKIN_misc_warning
	case "yinyang":
		return SKIN_misc_yinyang
	case "doge":
		return SKIN_misc_doge
	case "waffle":
		return SKIN_misc_waffle
	case "clock":
		return SKIN_misc_clock
	case "no_smoking":
		return SKIN_misc_no_smoking
	case "pig":
		return SKIN_misc_pig
	case "turtle":
		return SKIN_misc_turtle
	case "hell_doge":
		return SKIN_misc_hell_doge
	case "polkadots":
		return SKIN_misc_polkadots
	case "polkadots2":
		return SKIN_misc_polkadots2
	case "wheel":
		return SKIN_misc_wheel
	case "compass":
		return SKIN_misc_compass
	case "sanik":
		return SKIN_misc_sanik
	case "radiation":
		return SKIN_misc_radiation
	case "radar":
		return SKIN_misc_radar
	case "creeper":
		return SKIN_misc_creeper
	case "biohazard":
		return SKIN_misc_biohazard
	case "lambda":
		return SKIN_misc_lambda
	case "tesla":
		return SKIN_misc_tesla
	case "cheshire":
		return SKIN_misc_cheshire
	case "jack":
		return SKIN_misc_jack
	case "colors":
		return SKIN_misc_colors
	case "baseball":
		return SKIN_misc_baseball
	case "basketball":
		return SKIN_misc_basketball
	case "beachball":
		return SKIN_misc_beachball
	case "doom":
		return SKIN_misc_doom
	case "euro":
		return SKIN_misc_euro
	case "eye":
		return SKIN_misc_eye
	case "grumpy":
		return SKIN_misc_grumpy
	case "penny":
		return SKIN_misc_penny
	case "pball":
		return SKIN_misc_pball
	case "peppermint":
		return SKIN_misc_peppermint
	case "pineapple":
		return SKIN_misc_pineapple
	case "poker":
		return SKIN_misc_poker
	case "record":
		return SKIN_misc_record
	case "wheatley":
		return SKIN_misc_wheatley
	case "clouds":
		return SKIN_misc_clouds
	case "crop":
		return SKIN_misc_crop
	case "sauron":
		return SKIN_misc_sauron
	case "ship":
		return SKIN_misc_ship
	case "wheel_car":
		return SKIN_misc_wheel_car
	case "zerg":
		return SKIN_misc_zerg
	case "rose":
		return SKIN_misc_rose
	case "cd":
		return SKIN_misc_cd
	case "chute":
		return SKIN_misc_chute
	case "astro":
		return SKIN_misc_astro
	case "chess":
		return SKIN_misc_chess
	case "sign":
		return SKIN_misc_sign
	case "trollface":
		return SKIN_misc_trollface
	case "megustaxcf":
		return SKIN_misc_megustaxcf
	case "y":
		return SKIN_misc_y
	case "dragonball":
		return SKIN_misc_dragonball
	case "stained":
		return SKIN_misc_stained
	case "stained2":
		return SKIN_misc_stained2
	case "bauble":
		return SKIN_misc_bauble
	case "camo":
		return SKIN_misc_camo
	case "bulb":
		return SKIN_misc_bulb
	case "spiderweb":
		return SKIN_misc_spiderweb
	case "rain":
		return SKIN_misc_rain
	case "chomp1":
		return SKIN_misc_chomp_1
	case "snow":
		return SKIN_misc_snow
	case "mercury":
		return SKIN_scifi_mercury
	case "venus":
		return SKIN_scifi_venus
	case "earth":
		return SKIN_scifi_earth
	case "mars":
		return SKIN_scifi_mars
	case "saturn":
		return SKIN_scifi_saturn
	case "jupiter":
		return SKIN_scifi_jupiter
	case "neptune":
		return SKIN_scifi_neptune
	case "moon":
		return SKIN_scifi_moon
	case "planet1":
		return SKIN_scifi_planet1
	case "planet2":
		return SKIN_scifi_planet2
	case "planet3":
		return SKIN_scifi_planet3
	case "pluto":
		return SKIN_scifi_pluto
	case "sun":
		return SKIN_scifi_sun
	case "planet":
		return SKIN_scifi_planet
	case "deathstar":
		return SKIN_scifi_deathstar
	case "pastry_cat":
		return SKIN_scifi_pastry_cat
	case "galaxy":
		return SKIN_scifi_galaxy
	case "dust":
		return SKIN_scifi_dust
	case "australia":
		return SKIN_country_australia
	case "austria":
		return SKIN_country_austria
	case "belgium":
		return SKIN_country_belgium
	case "brazil":
		return SKIN_country_brazil
	case "bulgaria":
		return SKIN_country_bulgaria
	case "canada":
		return SKIN_country_canada
	case "china":
		return SKIN_country_china
	case "finland":
		return SKIN_country_finland
	case "france":
		return SKIN_country_france
	case "germany":
		return SKIN_country_germany
	case "greece":
		return SKIN_country_greece
	case "india":
		return SKIN_country_india
	case "italy":
		return SKIN_country_italy
	case "japan":
		return SKIN_country_japan
	case "mexico":
		return SKIN_country_mexico
	case "netherlands":
		return SKIN_country_netherlands
	case "norway":
		return SKIN_country_norway
	case "poland":
		return SKIN_country_poland
	case "romania":
		return SKIN_country_romania
	case "russia":
		return SKIN_country_russia
	case "southafrica":
		return SKIN_country_south_africa
	case "southkorea":
		return SKIN_country_southkorea
	case "spain":
		return SKIN_country_spain
	case "sweden":
		return SKIN_country_sweden
	case "turkey":
		return SKIN_country_turkey
	case "uk":
		return SKIN_country_uk
	case "ukraine":
		return SKIN_country_ukraine
	case "usa":
		return SKIN_country_usa
	case "columbia":
		return SKIN_country_columbia
	case "ecuador":
		return SKIN_country_ecuador
	case "ireland":
		return SKIN_country_ireland
	case "puerto_rico":
		return SKIN_country_puerto_rico
	case "argentina":
		return SKIN_country_argentina
	case "denmark":
		return SKIN_country_denmark
	case "egypt":
		return SKIN_country_egypt
	case "peru":
		return SKIN_country_peru
	case "portugal":
		return SKIN_country_portugal
	case "georgia":
		return SKIN_country_georgia
	case "croatia":
		return SKIN_country_croatia
	case "morocco":
		return SKIN_country_morocco
	case "israel":
		return SKIN_country_israel
	case "pakistan":
		return SKIN_country_pakistan
	case "bosnia":
		return SKIN_country_bosnia
	case "chile":
		return SKIN_country_chile
	case "hungary":
		return SKIN_country_hungary
	case "dr":
		return SKIN_country_dr
	case "philipines":
		return SKIN_country_philipines
	case "venezuela":
		return SKIN_country_venezuela
	case "trinidad":
		return SKIN_country_trinidad
	case "costarica":
		return SKIN_country_costarica
	case "scotlan":
		return SKIN_country_scotlan
	case "belaruse":
		return SKIN_country_belaruse
	case "serbia":
		return SKIN_country_serbia
	case "slovakia":
		return SKIN_country_slovakia
	case "tunisia":
		return SKIN_country_tunisia
	case "iran":
		return SKIN_country_iran
	case "thailand":
		return SKIN_country_thailand
	case "switzerland":
		return SKIN_country_switzerland
	case "algeria":
		return SKIN_country_algeria
	case "newzealand":
		return SKIN_country_newzealand
	case "lithuania":
		return SKIN_country_lithuania
	case "jamaica":
		return SKIN_country_jamaica
	case "guatemala":
		return SKIN_country_guatemala
	case "lebanon":
		return SKIN_country_lebanon
	case "albania":
		return SKIN_country_albania
	case "macedonia":
		return SKIN_country_macedonia
	case "latvia":
		return SKIN_country_latvia
	case "azerbaijan":
		return SKIN_country_azerbaijan
	case "estonia":
		return SKIN_country_estonia
	case "czech":
		return SKIN_country_czech
	case "wales":
		return SKIN_country_wales
	case "armenia":
		return SKIN_country_armenia
	case "country_cuba":
		return SKIN_country_cuba
	case "england":
		return SKIN_country_england
	case "kazakhstan":
		return SKIN_country_kazakhstan
	case "iceland":
		return SKIN_country_iceland
	case "indonesia":
		return SKIN_country_indonesia
	case "panama":
		return SKIN_country_panama
	case "cyprus":
		return SKIN_country_cyprus
	case "moldova":
		return SKIN_country_moldova
	case "montenegro":
		return SKIN_country_montenegro
	case "honduras":
		return SKIN_country_honduras
	case "elsalvador":
		return SKIN_country_elsalvador
	case "uruguay":
		return SKIN_country_uruguay
	case "iraq":
		return SKIN_country_iraq
	case "saudiaarabia":
		return SKIN_country_saudiaarabia
	case "awesome":
		return SKIN_misc_awesome
	case "pug":
		return SKIN_misc_pug
	case "panda":
		return SKIN_misc_panda
	case "saw":
		return SKIN_misc_saw
	case "jordan":
		return SKIN_country_jordan
	case "penguin":
		return SKIN_misc_penguin
	case "heart":
		return SKIN_misc_heart
	case "circuit_2":
		return SKIN_misc_circuit_2
	case "cookie":
		return SKIN_misc_cookie
	case "hourglass":
		return SKIN_misc_hourglass
	case "mona":
		return SKIN_misc_mona
	case "pointer":
		return SKIN_misc_pointer
	case "shield":
		return SKIN_misc_shield
	case "dust_blue":
		return SKIN_scifi_dust_blue
	case "dust_red":
		return SKIN_scifi_dust_red
	default:
		return SKIN_misc_none
	}
}
func AllSkins() []string {
	return []string{
		"none",
		"8ball",
		"circuit",
		"glossyball",
		"lu",
		"matrix",
		"paint",
		"soccer",
		"warning",
		"yinyang",
		"doge",
		"waffle",
		"clock",
		"no_smoking",
		"pig",
		"turtle",
		"hell_doge",
		"polkadots",
		"polkadots2",
		"wheel",
		"compass",
		"sanik",
		"radiation",
		"radar",
		"creeper",
		"biohazard",
		"lambda",
		"tesla",
		"cheshire",
		"jack",
		"colors",
		"baseball",
		"basketball",
		"beachball",
		"doom",
		"euro",
		"eye",
		"grumpy",
		"lightning",
		"pball",
		"penny",
		"peppermint",
		"pineapple",
		"poker",
		"record",
		"wheatley",
		"clouds",
		"crop",
		"sauron",
		"ship",
		"wheel_car",
		"zerg",
		"rose",
		"cd",
		"chute",
		"astro",
		"chess",
		"sign",
		"trollface",
		"megustaxcf",
		"y",
		"dragonball",
		"stained",
		"stained2",
		"bauble",
		"camo",
		"bulb",
		"spiderweb",
		"rain",
		"chomp1",
		"snow",
		"mercury",
		"venus",
		"earth",
		"mars",
		"saturn",
		"jupiter",
		"neptune",
		"moon",
		"planet1",
		"planet2",
		"planet3",
		"pluto",
		"sun",
		"planet",
		"deathstar",
		"pastry_cat",
		"galaxy",
		"dust",
		"australia",
		"austria",
		"belgium",
		"brazil",
		"bulgaria",
		"canada",
		"china",
		"finland",
		"france",
		"germany",
		"greece",
		"india",
		"italy",
		"japan",
		"mexico",
		"netherlands",
		"norway",
		"poland",
		"romania",
		"russia",
		"southafrica",
		"southkorea",
		"spain",
		"southkorea",
		"spain",
		"sweden",
		"turkey",
		"uk",
		"ukraine",
		"usa",
		"columbia",
		"ecuador",
		"ireland",
		"puerto_rico",
		"argentina",
		"denmark",
		"egypt",
		"peru",
		"portugal",
		"georgia",
		"croatia",
		"morocco",
		"israel",
		"pakistan",
		"bosnia",
		"chile",
		"hungary",
		"dr",
		"philipines",
		"venezuela",
		"trinidad",
		"costarica",
		"scotlan",
		"belaruse",
		"serbia",
		"slovakia",
		"tunisia",
		"iran",
		"thailand",
		"switzerland",
		"algeria",
		"newzealand",
		"lithuania",
		"jamaica",
		"guatemala",
		"lebanon",
		"albania",
		"macedonia",
		"latvia",
		"azerbaijan",
		"estonia",
		"czech",
		"wales",
		"armenia",
		"country_cuba",
		"england",
		"kazakhstan",
		"iceland",
		"indonesia",
		"panama",
		"cyprus",
		"moldova",
		"montenegro",
		"honduras",
		"elsalvador",
		"uruguay",
		"iraq",
		"saudiaarabia",
		"awesome",
		"pug",
		"panda",
		"saw",
		"jordan",
		"penguin",
		"heart",
		"circuit_2",
		"cookie",
		"hourglass",
		"mona",
		"pointer",
		"shield",
		"dust_blue",
		"dust_red",
	}
}

type SPLIT_MULTIPLIER byte

const (
	SPLIT_MULTIPLIER_X8 SPLIT_MULTIPLIER = iota
	SPLIT_MULTIPLIER_X16
	SPLIT_MULTIPLIER_X32
	SPLIT_MULTIPLIER_X64
)

type WORLD_SIZE byte

const (
	WORLD_SIZE_TINY WORLD_SIZE = iota
	WORLD_SIZE_SMALL
	WORLD_SIZE_NORMAL
	WORLD_SIZE_LARGE
)
