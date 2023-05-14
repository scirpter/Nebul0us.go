package packets

import (
	"math"
	event_handle "neb/src/base/client/events"
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/base/models"
	"neb/src/base/natives"
	"neb/src/base/sigs"
	"neb/src/common"
	"neb/src/net"
)

// ## ?
// Received to obtain raw
// game object data like hole
// positions, player splits, ...
type GameUpdate struct {
	enum            enums.PACKET_TYPE
	net             *specs.Net
	allNets         []*specs.Net
	data            *[]byte
	eventDispatcher *event_handle.EventDispatcher
}

func NewGameUpdatePacket(net *specs.Net, allNets []*specs.Net, data *[]byte, eventDispatcher *event_handle.EventDispatcher) *GameUpdate {
	return &GameUpdate{enum: enums.PACKET_TYPE_GAME_UPDATE, net: net, data: data, allNets: allNets, eventDispatcher: eventDispatcher}
}

func (p *GameUpdate) Parse() {
	dis := net.NewDataInputStream(p.data)

	dis.ReadByte2() // packet type
	dis.ReadByte2() // tick
	dis.ReadByte2() // tick
	love := natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, sigs.MAP_SIZE_SIG)
	dis.ReadByte2()
	natives.RNATIVE_LE_SHORT_INT(dis)
	p0 := goofyAhCalc(dis.ReadByte2())
	_ = p0
	q0 := goofyAhCalc(dis.ReadByte2())
	_ = q0
	playerCt := dis.ReadByte2()
	readByte := dis.ReadByte2()
	eventCt := readByte & 31
	io := (readByte & 224) >> 5
	ejectionCt := dis.ReadByte2()
	dotCt := goofyAhCalc(dis.ReadByte2())
	holeCt := dis.ReadByte2()
	ejectionMoveCt := dis.ReadByte2()
	readByte2 := dis.ReadByte2()
	itemCt := readByte2 & 31
	jo := (readByte2 & 224) >> 5
	readByte3 := dis.ReadByte2()
	spellCt := readByte3 & 31
	wallCt := (readByte3 & 224) >> 5
	m0 := goofyAhCalc(dis.ReadByte2())
	r0 := dis.ReadByte2()

	if io > 0 {
		netID := make([]common.NET_ID, io)
		x := make([]float32, io)
		y := make([]float32, io)

		for i := 0; i < int(io); i++ {
			netID[i] = common.NET_ID(dis.ReadByte2())
			x[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			y[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
		}
	}

	if ejectionCt > 0 {
		netID := make([]common.NET_ID, ejectionCt)
		fromPlayerNetID := make([]common.NET_ID, ejectionCt)
		x := make([]float32, ejectionCt)
		y := make([]float32, ejectionCt)
		mass := make([]float32, ejectionCt)
		angle := make([]float32, ejectionCt)

		for i := 0; i < int(ejectionCt); i++ {
			readByte4 := dis.ReadByte2()
			z8 := (readByte4 & 1) == 1
			netID[i] = common.NET_ID((readByte4 & 254) >> 1)
			fromPlayerNetID[i] = common.NET_ID(dis.ReadByte2())
			x[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			y[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			angle[i] = natives.RNATIVE_OBJ_ANGLE(dis, 0, 2*math.Pi)
			if !z8 {
				mass[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, 500000.0)
				ejection := models.NewEjection(netID[i], x[i], y[i], mass[i])
				ejection.Angle = angle[i]
				ejection.FromPlayerNetID = fromPlayerNetID[i]

				for _, net := range p.allNets {
					net.World.EjectionMutx.Lock()
					net.World.Ejections[netID[i]] = ejection
					net.World.EjectionMutx.Unlock()
				}
			} else {
				// ejection delete
				mass[i] = float32(math.Inf(-1))

				for _, net := range p.allNets {
					net.World.EjectionMutx.Lock()
					delete(net.World.Ejections, netID[i])
					net.World.EjectionMutx.Unlock()
				}
			}
		}
	}

	if ejectionMoveCt > 0 {
		netID := make([]common.NET_ID, ejectionMoveCt)
		x := make([]float32, ejectionMoveCt)
		y := make([]float32, ejectionMoveCt)

		for i := 0; i < int(ejectionMoveCt); i++ {
			netID[i] = common.NET_ID(dis.ReadByte2())
			x[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			y[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)

			for _, net := range p.allNets {
				net.World.EjectionMutx.Lock()
				for _, ejection := range net.World.Ejections {
					ejection.X = x[i]
					ejection.Y = y[i]
				}
				net.World.EjectionMutx.Unlock()
			}
		}
	}

	if dotCt > 0 {
		netID := make([]common.NET_ID, dotCt)
		x := make([]float32, dotCt)
		y := make([]float32, dotCt)

		for i := 0; i < int(dotCt); i++ {
			netID[i] = common.NET_ID(dis.ReadShort())
			x[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			y[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)

			for _, net := range p.allNets {
				net.World.DotMutx.Lock()
				net.World.Dots[netID[i]] = models.NewDot(netID[i], x[i], y[i])
				net.World.DotMutx.Unlock()
			}
		}
	}

	if itemCt > 0 {
		netID := make([]common.NET_ID, itemCt)
		typeof := make([]enums.ITEM_TYPE, itemCt)
		x := make([]float32, itemCt)
		y := make([]float32, itemCt)

		for i := 0; i < int(itemCt); i++ {
			netID[i] = common.NET_ID(dis.ReadByte2())
			typeof[i] = enums.ITEM_TYPE(dis.ReadByte2())
			x[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			y[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			item := models.NewItem(netID[i], typeof[i], x[i], y[i])

			for _, net := range p.allNets {
				net.World.ItemMutx.Lock()
				net.World.Items[netID[i]] = item
				net.World.ItemMutx.Unlock()
			}
		}
	}

	if holeCt > 0 {
		netID := make([]common.NET_ID, holeCt)
		typeof := make([]enums.HOLE_TYPE, holeCt)
		x := make([]float32, holeCt)
		y := make([]float32, holeCt)
		mass := make([]float32, holeCt)

		for i := 0; i < int(holeCt); i++ {
			readByte5 := dis.ReadByte2()
			netID[i] = common.NET_ID((readByte5 & 252) >> 2)
			typeof[i] = enums.HOLE_TYPE(readByte5 & 3)
			y[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			x[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			mass[i] = natives.RNATIVE_APPLY_SCALING_FACTOR(dis, 62.6)
			hole := models.NewHole(netID[i], typeof[i], x[i], y[i], mass[i])
			if mass[i] == 0 {
				// hole delete
				for _, net := range p.allNets {
					net.World.HoleMutx.Lock()
					delete(net.World.Holes, netID[i])
					net.World.HoleMutx.Unlock()
				}
			} else {
				for _, net := range p.allNets {
					net.World.HoleMutx.Lock()
					net.World.Holes[netID[i]] = hole
					net.World.HoleMutx.Unlock()
				}
			}
		}
	}

	if jo > 0 {
		f := make([]common.NET_ID, jo)
		g := make([]float32, jo)
		h := make([]float32, jo)
		i := make([]float32, jo)
		j := make([]float32, jo)

		for i21 := 0; i21 < int(jo); i21++ {
			f[i21] = common.NET_ID(dis.ReadByte2())
			g[i21] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			h[i21] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			i[i21] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			j[i21] = natives.RNATIVE_OBJ_ANGLE(dis, 0.0, 1.0)
		}
	}

	perBlobNetID := make([]common.NET_ID, 0)
	P := make([]uint16, 0)
	blobX := make([]float32, 0)
	blobY := make([]float32, 0)
	blobMass := make([]float64, 0)

	if playerCt > 0 {
		netID := make([]common.NET_ID, playerCt)
		M := make([]uint8, playerCt)
		blobCt := make([]uint8, playerCt)

		for i := 0; i < int(playerCt); i++ {
			readByte6 := dis.ReadByte2()
			netID[i] = common.NET_ID((readByte6 & 254) >> 1)
			z2 := false

			if (readByte6 & 1) != 0 {
				z2 = true
			}

			if z2 {
				M[i] = dis.ReadByte2()
			} else {
				M[i] = 0
			}

			blobCt[i] = goofyAhCalc(dis.ReadByte2())
			players := make([]*models.Player, 0)

			for _, net := range p.allNets {
				net.World.PlayerMutx.Lock()
				if player, ok := net.World.Players[netID[i]]; ok {
					player.Blobs = make(map[common.NET_ID]*models.Blob)
					players = append(players, player)
				} else {
					noname := "NULL"
					player := models.NewPlayer(netID[i], &noname, -1)
					net.World.Players[netID[i]] = player
					players = append(players, player)
				}
				net.World.PlayerMutx.Unlock()
			}

			for i23 := 0; i23 < int(blobCt[i]); i23++ {
				readByte7 := dis.ReadByte2()
				z3 := false

				if (readByte7 & 1) == 1 {
					z3 = true
				}

				z4 := false

				if ((readByte7 & 128) >> 7) == 1 {
					z4 = true
				}

				perBlobNetID = append(perBlobNetID, common.NET_ID((readByte7&126)>>1))

				if z3 {
					P = append(P, dis.ReadShort())
				} else {
					P = append(P, 0)
				}
				_ = P

				blobX = append(blobX, natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love))
				blobY = append(blobY, natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love))

				if !z4 {
					blobMass = append(blobMass, float64(natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, 500000.0)))
				} else {
					blobMass = append(blobMass, math.Inf(-1))
				}

				blob := models.NewBlob(
					perBlobNetID[len(perBlobNetID)-1],
					blobX[len(blobX)-1],
					blobY[len(blobY)-1],
					blobMass[len(blobMass)-1],
				)

				for _, player := range players {
					player.BlobMutx.Lock()
					player.Blobs[perBlobNetID[len(perBlobNetID)-1]] = blob
					player.BlobMutx.Unlock()
				}
			}
		}
	}

	if r0 > 0 {
		G := make([]common.NET_ID, r0)
		H := make([]bool, r0)
		I := make([]float32, r0)
		J := make([]float32, r0)
		K := make([]float32, r0)

		for i := 0; i < int(r0); i++ {
			readByte8 := dis.ReadByte2()
			G[i] = common.NET_ID(readByte8 & 0x7f) // 0x7f = Byte.MAX_VALUE
			z := false

			if ((readByte8 & 128) >> 7) != 0 {
				z = true
			}

			H[i] = z
			I[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			J[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			K[i] = natives.RNATIVE_OBJ_ANGLE(dis, 0.0, float32(math.Sqrt(2000)))
		}
	}

	if spellCt > 0 {
		netID := make([]common.NET_ID, spellCt)
		typeof := make([]enums.SPELL_TYPE, spellCt)
		status := make([]enums.SPELL_STATUS, spellCt)
		x := make([]float32, spellCt)
		y := make([]float32, spellCt)

		for i := 0; i < int(spellCt); i++ {
			netID[i] = common.NET_ID(dis.ReadByte2())
			readByte9 := dis.ReadByte2()
			status[i] = enums.SPELL_STATUS(readByte9 & 3)
			typeof[i] = enums.SPELL_TYPE((readByte9 & 252) >> 2)

			if typeof[i] >= (enums.SPELL_TYPE_UNKNOWN + 1) {
				typeof[i] = enums.SPELL_TYPE_UNKNOWN
			}

			x[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			y[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			spell := models.NewSpell(netID[i], typeof[i], status[i], x[i], y[i])

			for _, net := range p.allNets {
				net.World.SpellMutx.Lock()
				net.World.Spells[netID[i]] = spell
				net.World.SpellMutx.Unlock()
			}
		}
	}

	if wallCt > 0 {
		Y := make([]common.NET_ID, wallCt)
		Z := make([]float32, wallCt)
		a0 := make([]float32, wallCt)
		b0 := make([]float32, wallCt)
		c0 := make([]float32, wallCt)

		for i := 0; i < int(wallCt); i++ {
			Y[i] = common.NET_ID(dis.ReadByte2())
			Z[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			a0[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			b0[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			c0[i] = natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
		}
	}

	if m0 > 0 {
		d0 := make([]byte, m0)
		e0 := make([]byte, m0)
		f0 := make([]byte, m0)

		for i := 0; i < int(m0); i++ {
			d0[i] = goofyAhCalc(dis.ReadByte2())
			e0[i] = goofyAhCalc(dis.ReadByte2())
			f0[i] = dis.ReadByte2()
		}
	}

	if eventCt > 0 {
		for i28 := 0; i28 < int(eventCt); i28++ {
			event := enums.GAME_EVENT(dis.ReadByte2())

			if !(event >= enums.GAME_EVENT_UNKNOWN && event < (enums.GAME_EVENT_RLGL_STATE+1)) {
				event = enums.GAME_EVENT_UNKNOWN
			}

			switch event {
			case enums.GAME_EVENT_EAT_DOTS:
			case enums.GAME_EVENT_EAT_BLOB:
			case enums.GAME_EVENT_EAT_SMBH:
			case enums.GAME_EVENT_BLOB_EXPLODE:
				dis.ReadByte2()
				dis.ReadByte2()
			case enums.GAME_EVENT_BLOB_LOST:
			case enums.GAME_EVENT_EJECT:
				dis.ReadByte2()
				dis.ReadByte2()
			case enums.GAME_EVENT_SPLIT:
				playerNetID := common.NET_ID(dis.ReadByte2())
				p.net.World.PlayerMutx.RLock()
				player := p.net.World.Players[playerNetID]
				p.net.World.PlayerMutx.RUnlock()
				eventData := event_handle.EventData{
					"player": player,
				}
				p.eventDispatcher.Dispatch(event_handle.Event{Type: event_handle.OnPlayerSplit, Data: eventData})
			case enums.GAME_EVENT_RECOMBINE:
				dis.ReadByte2()
			case enums.GAME_EVENT_TIMER_WARNING:
			case enums.GAME_EVENT_CTF_SCORE:
			case enums.GAME_EVENT_CTF_FLAG_RETURNED:
			case enums.GAME_EVENT_CTF_FLAG_STOLEN:
			case enums.GAME_EVENT_ACHIEVEMENT_EARNED:
				dis.ReadShort()
			case enums.GAME_EVENT_XP_GAINED:
			case enums.GAME_EVENT_UNUSED_2:
			case enums.GAME_EVENT_XP_SET:
				dis.ReadLong()
				dis.ReadByte2()
				dis.ReadInt()
				dis.ReadByte2()
				natives.RNATIVE_LE_SHORT_INT(dis)
			case enums.GAME_EVENT_DQ_SET:
				dis.ReadByte2()
				dis.ReadBool()
			case enums.GAME_EVENT_DQ_COMPLETED:
				dis.ReadByte2()
			case enums.GAME_EVENT_DQ_PROGRESS:
				dis.ReadShort()
			case enums.GAME_EVENT_EAT_SERVER_BLOB:
			case enums.GAME_EVENT_EAT_SPECIAL_OBJECTS:
				dis.ReadByte2()
				dis.ReadByte2()
			case enums.GAME_EVENT_SO_SET:
				dis.ReadByte2()
				dis.ReadInt()
			case enums.GAME_EVENT_LEVEL_UP:
				dis.ReadShort()
			case enums.GAME_EVENT_ARENA_RANK_ACHIEVED:
				dis.ReadByte2()
				dis.ReadByte2()
			case enums.GAME_EVENT_DOM_CP_LOST:
			case enums.GAME_EVENT_DOM_CP_GAINED:
			case enums.GAME_EVENT_UNUSED_1:
			case enums.GAME_EVENT_CTF_GAINED:
			case enums.GAME_EVENT_GAME_OVER:
			case enums.GAME_EVENT_BLOB_STATUS:
				dis.ReadByte2()
				dis.ReadByte2()
				dis.ReadShort()
			case enums.GAME_EVENT_TELEPORT:
				dis.ReadByte2()
			case enums.GAME_EVENT_SHOOT:
				dis.ReadByte2()
				dis.ReadByte2()
				dis.ReadByte2()
			case enums.GAME_EVENT_CLAN_WAR_WON:
				dis.ReadShort()
			case enums.GAME_EVENT_PLASMA_REWARD:
				natives.RNATIVE_LE_SHORT_INT(dis)
				dis.ReadByte2()
			case enums.GAME_EVENT_EMOTE:
				dis.ReadByte2()
				dis.ReadByte2()
				emoteID := dis.ReadByte2()
				_ = emoteID
				dis.ReadInt()
			case enums.GAME_EVENT_END_MISSION:
				dis.ReadByte2()
				dis.ReadBool()
				dis.ReadByte2()
				natives.RNATIVE_LE_SHORT_INT(dis)
				dis.ReadShort()
			case enums.GAME_EVENT_XP_GAINED_2:
				i0 := natives.RNATIVE_APPLY_SCALING_FACTOR(dis, 1.0) + 1.0
				_ = i0
				natives.RNATIVE_LE_SHORT_INT(dis)
				natives.RNATIVE_LE_SHORT_INT(dis)
			case enums.GAME_EVENT_EAT_CAKE:
				natives.RNATIVE_LE_SHORT_INT(dis)
				natives.RNATIVE_LE_SHORT_INT(dis)
			case enums.GAME_EVENT_COIN_COUNT:
				dis.ReadByte2()
				dis.ReadShort()
			case enums.GAME_EVENT_CLEAR_EFFECTS:
			case enums.GAME_EVENT_SPEED:
				dis.ReadShort()
			case enums.GAME_EVENT_TRICK:
				dis.ReadByte2()
				dis.ReadShort()
				natives.RNATIVE_LE_SHORT_INT(dis)
			case enums.GAME_EVENT_DESTROY_ASTEROID:
			case enums.GAME_EVENT_ACCOLADE:
				dis.ReadByte2()
			case enums.GAME_EVENT_INVIS:
				dis.ReadShort()
			case enums.GAME_EVENT_KILLED_BY:
				dis.ReadByte2()
			case enums.GAME_EVENT_RADIATION_CLOUD:
				dis.ReadByte2()
				natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
				natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
				natives.RNATIVE_OBJ_ANGLE(dis, 0.0, 16.0)
			case enums.GAME_EVENT_CHARGE:
				dis.ReadByte2()
				dis.ReadByte2()
			case enums.GAME_EVENT_LP_COUNT:
				dis.ReadByte2()
			case enums.GAME_EVENT_BR_BOUNDS:
				natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
				natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
				natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
				natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
				natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
				natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
				natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
				natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, love)
			case enums.GAME_EVENT_MINIMAP:
			case enums.GAME_EVENT_RLGL_DEATH:
			case enums.GAME_EVENT_RLGL_STATE:
				dis.ReadByte2()
			}
		}
	}
}

func goofyAhCalc(rb uint8) uint8 {
	return uint8((int16(rb) + 128) % 256)
}
