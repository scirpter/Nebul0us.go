package models

import "neb/src/common"

type Dot struct {
	netID common.NET_ID
	X, Y  float32
}

func NewDot(netID common.NET_ID, x float32, y float32) *Dot {
	return &Dot{
		netID: netID,
		X:     x,
		Y:     y,
	}
}
