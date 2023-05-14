package models

import "neb/src/common"

type Blob struct {
	netID common.NET_ID
	X, Y  float32
	Mass  float64
}

func NewBlob(netID common.NET_ID, x float32, y float32, mass float64) *Blob {
	return &Blob{
		netID: netID,
		X:     x,
		Y:     y,
		Mass:  mass,
	}
}
