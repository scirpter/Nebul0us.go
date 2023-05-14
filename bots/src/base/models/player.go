package models

import (
	"math"
	"neb/src/base/enums"
	"neb/src/common"
	"sync"
)

type Player struct {
	NetID     common.NET_ID
	Name      *string
	Cr2Token2 uint32 // as of now only set in JoinResult
	AccountID int32
	ClanName  *string
	// -2 = bot, -1 = signed out player
	Level    int16
	skin     enums.SKIN
	hat      common.Undefined
	halo     common.Undefined
	Blobs    map[common.NET_ID]*Blob
	BlobMutx sync.RWMutex
	particle common.Undefined
}

func NewPlayer(netID common.NET_ID, name *string, accountID int32) *Player {
	return &Player{
		NetID:     netID,
		Name:      name,
		Cr2Token2: 0,
		AccountID: accountID,
		ClanName:  nil,
		Level:     -1,
		skin:      enums.SKIN_misc_none,
		hat:       common.ValuedUndefined,
		halo:      common.ValuedUndefined,
		Blobs:     make(map[common.NET_ID]*Blob),
		particle:  common.ValuedUndefined,
	}
}

func (p *Player) GetMass() float64 {
	var mass float64 = 0

	p.BlobMutx.RLock()
	for _, blob := range p.Blobs {
		if math.IsInf(blob.Mass, -1) {
			continue
		}
		mass += blob.Mass
	}
	p.BlobMutx.RUnlock()

	return mass
}

func (p *Player) GetRadius() float64 {
	return math.Sqrt(p.GetMass()) * 0.5
}

func (p *Player) GetAvgPos() (float32, float32) {
	var x, y float32 = 0, 0
	var ct uint8 = 0

	p.BlobMutx.RLock()
	for _, blob := range p.Blobs {
		if math.IsInf(blob.Mass, -1) {
			continue
		}
		x += blob.X
		y += blob.Y
		ct++
	}
	p.BlobMutx.RUnlock()
	if ct == 0 { // do not divide through 0
		return 0, 0
	}
	len := float32(ct)

	return x / len, y / len
}
