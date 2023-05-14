package models

import (
	"neb/src/base/enums"
	"neb/src/common"
	"sync"
)

type World struct {
	Name        *string
	TimeLeft    uint16
	GameMode    enums.GAME_MODE
	MaxPlayers  byte
	PlayerCt    byte
	SpectatorCt byte
	Tick        byte
	RawSize     float32
	WorldSize   enums.WORLD_SIZE
	Token       uint32

	Ejections    map[common.NET_ID]*Ejection
	EjectionMutx *sync.RWMutex

	Players    map[common.NET_ID]*Player
	PlayerMutx *sync.RWMutex

	Dots    map[common.NET_ID]*Dot
	DotMutx *sync.RWMutex

	Items    map[common.NET_ID]*Item
	ItemMutx *sync.RWMutex

	Spells    map[common.NET_ID]*Spell
	SpellMutx *sync.RWMutex

	Holes    map[common.NET_ID]*Hole
	HoleMutx *sync.RWMutex
}

func NewWorld() *World {
	return &World{
		Name:        nil,
		TimeLeft:    0,
		GameMode:    enums.GAME_MODE_FFA,
		MaxPlayers:  0,
		SpectatorCt: 0,
		Tick:        0,
		RawSize:     0,
		WorldSize:   enums.WORLD_SIZE_NORMAL,
		Token:       0,

		Ejections:    make(map[common.NET_ID]*Ejection),
		EjectionMutx: new(sync.RWMutex),
		Players:      make(map[common.NET_ID]*Player),
		PlayerMutx:   new(sync.RWMutex),
		Dots:         make(map[common.NET_ID]*Dot),
		DotMutx:      new(sync.RWMutex),
		Items:        make(map[common.NET_ID]*Item),
		ItemMutx:     new(sync.RWMutex),
		Spells:       make(map[common.NET_ID]*Spell),
		SpellMutx:    new(sync.RWMutex),
		Holes:        make(map[common.NET_ID]*Hole),
		HoleMutx:     new(sync.RWMutex),
	}
}

func (w *World) Reset() {
	w.Name = nil
	w.TimeLeft = 0
	w.GameMode = enums.GAME_MODE_FFA
	w.MaxPlayers = 0
	w.SpectatorCt = 0
	w.Tick = 0
	w.RawSize = 0
	w.WorldSize = enums.WORLD_SIZE_NORMAL
	w.Token = 0

	w.EjectionMutx.Lock()
	w.Ejections = make(map[common.NET_ID]*Ejection)
	w.EjectionMutx.Unlock()

	w.PlayerMutx.Lock()
	w.Players = make(map[common.NET_ID]*Player)
	w.PlayerMutx.Unlock()

	w.DotMutx.Lock()
	w.Dots = make(map[common.NET_ID]*Dot)
	w.DotMutx.Unlock()

	w.ItemMutx.Lock()
	w.Items = make(map[common.NET_ID]*Item)
	w.ItemMutx.Unlock()

	w.SpellMutx.Lock()
	w.Spells = make(map[common.NET_ID]*Spell)
	w.SpellMutx.Unlock()

	w.HoleMutx.Lock()
	w.Holes = make(map[common.NET_ID]*Hole)
	w.HoleMutx.Unlock()
}

func (w *World) Copy() *World {
	return &World{
		Name:        w.Name,
		TimeLeft:    w.TimeLeft,
		GameMode:    w.GameMode,
		MaxPlayers:  w.MaxPlayers,
		SpectatorCt: w.SpectatorCt,
		Tick:        w.Tick,
		RawSize:     w.RawSize,
		WorldSize:   w.WorldSize,
		Token:       w.Token,

		Ejections:    w.Ejections,
		EjectionMutx: w.EjectionMutx,
		Players:      w.Players,
		PlayerMutx:   w.PlayerMutx,
		Dots:         w.Dots,
		DotMutx:      w.DotMutx,
		Items:        w.Items,
		ItemMutx:     w.ItemMutx,
		Spells:       w.Spells,
		SpellMutx:    w.SpellMutx,
		Holes:        w.Holes,
		HoleMutx:     w.HoleMutx,
	}
}

type WorldProps struct {
	isHidden          bool
	minPlayers        byte
	maxPlayers        byte
	gameMode          enums.GAME_MODE
	worldSize         enums.WORLD_SIZE
	difficulty        enums.DIFFICULTY
	roomName          *string
	duration          uint16
	isMayhem          bool
	splitMultiplier   enums.SPLIT_MULTIPLIER
	allowUltraClick   bool
	allowMassBoost    bool
	allowRainbowHoles bool
	wallCt            byte
	allowAdmin        bool
	allowGuests       bool
	isArena           bool
}

func NewWorldProps() *WorldProps {
	return &WorldProps{
		isHidden:          false,
		minPlayers:        0,
		maxPlayers:        0,
		gameMode:          enums.GAME_MODE_FFA,
		worldSize:         enums.WORLD_SIZE_NORMAL,
		difficulty:        enums.DIFFICULTY_EASY,
		roomName:          nil,
		duration:          0,
		isMayhem:          false,
		splitMultiplier:   enums.SPLIT_MULTIPLIER_X8,
		allowUltraClick:   false,
		allowMassBoost:    false,
		allowRainbowHoles: false,
		wallCt:            0,
		allowAdmin:        false,
		allowGuests:       false,
		isArena:           false,
	}
}
