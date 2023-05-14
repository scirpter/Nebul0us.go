package models

import "neb/src/common"

type Ejection struct {
	netID           common.NET_ID
	FromPlayerNetID common.NET_ID
	X, Y, Mass      float32
	Angle           float32
}

func NewEjection(netID common.NET_ID, x float32, y float32, mass float32) *Ejection {
	return &Ejection{
		netID:           netID,
		FromPlayerNetID: common.NET_ID(0),
		X:               x,
		Y:               y,
		Mass:            mass,
		Angle:           0.0,
	}
}
