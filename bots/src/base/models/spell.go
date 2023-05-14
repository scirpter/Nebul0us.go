package models

import (
	"neb/src/base/enums"
	"neb/src/common"
)

type Spell struct {
	netID  common.NET_ID
	Typeof enums.SPELL_TYPE
	Status enums.SPELL_STATUS
	X, Y   float32
}

func NewSpell(netID common.NET_ID, typeof enums.SPELL_TYPE, status enums.SPELL_STATUS, x float32, y float32) *Spell {
	return &Spell{
		netID:  netID,
		Typeof: typeof,
		Status: status,
		X:      x,
		Y:      y,
	}
}
