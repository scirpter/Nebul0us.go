package models

import (
	"neb/src/base/enums"
	"neb/src/common"
)

type Item struct {
	netID      common.NET_ID
	Typeof     enums.ITEM_TYPE
	X, Y, Dist float32
}

func NewItem(netID common.NET_ID, typeof enums.ITEM_TYPE, x float32, y float32) *Item {
	return &Item{
		netID:  netID,
		Typeof: typeof,
		X:      x,
		Y:      y,
	}
}
