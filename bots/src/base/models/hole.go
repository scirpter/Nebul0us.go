package models

import (
	"neb/src/base/enums"
	"neb/src/common"
)

type Hole struct {
	netID            common.NET_ID
	Typeof           enums.HOLE_TYPE
	X, Y, Mass, Dist float32
}

func NewHole(netID common.NET_ID, typeof enums.HOLE_TYPE, x float32, y float32, mass float32) *Hole {
	return &Hole{
		netID:  netID,
		X:      x,
		Y:      y,
		Mass:   mass,
		Typeof: typeof,
	}
}
