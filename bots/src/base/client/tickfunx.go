package client

import (
	"neb/src/base/enums"
	"neb/src/base/models"
	"neb/src/utils/cnsl"
)

func RUN_FUNCTIONS(client *Client, player *models.Player) {
	if !client.Statery.IsFarmingPlasma && client.Statery.PlasmaFarmingTargetPlayerName == nil && client.Statery.PlasmaFarmingTargetPlayerAccountID == -1 && player.GetMass() < float64(client.Statery.MassThreshold) {
		client.TargetHoles(player, enums.GoodHoles())
		client.Split()
	} else if client.Statery.PlasmaFarmingTargetPlayerName != nil || client.Statery.PlasmaFarmingTargetPlayerAccountID != -1 {
		if client.net.World.TimeLeft < 30 {
			var target *models.Player

			client.net.World.PlayerMutx.RLock()
			for _, value := range client.net.World.Players {
				if client.Statery.PlasmaFarmingTargetPlayerName != nil && *value.Name == *client.Statery.PlasmaFarmingTargetPlayerName {
					target = value
					break
				} else if client.Statery.PlasmaFarmingTargetPlayerAccountID != -1 && value.AccountID == client.Statery.PlasmaFarmingTargetPlayerAccountID {
					target = value
					break
				}
			}
			client.net.World.PlayerMutx.RUnlock()

			if target == nil {
				cnsl.Error("target player not loaded. are you sure the target player has not left and all bots are in the same lobby?")
				client.Move(client.controlData.Angle, 0)
				client.controlData.SplitCt = 0
				client.Statery.PlasmaFarmingTargetPlayerName = nil
				client.Statery.PlasmaFarmingTargetPlayerAccountID = -1
				return
			}
			client.TargetPlayer(player, target)
			client.Split()
			return
		}
		client.TargetPlasma(player)
		client.Split()
	} else if client.Statery.IsFarmingPlasma {
		client.TargetPlasma(player)
		client.Split()
	} else if client.Statery.EmotionalSupportPlayerName != nil || client.Statery.EmotionalSupportPlayerAccountID != -1 {
		var target *models.Player

		client.net.World.PlayerMutx.RLock()
		for _, value := range client.net.World.Players {
			if client.Statery.EmotionalSupportPlayerName != nil && *value.Name == *client.Statery.EmotionalSupportPlayerName {
				target = value
				break
			} else if client.Statery.EmotionalSupportPlayerAccountID != -1 && value.AccountID == client.Statery.EmotionalSupportPlayerAccountID {
				target = value
				break
			}
		}
		client.net.World.PlayerMutx.RUnlock()

		if target == nil {
			cnsl.Error("target player not loaded. are you sure the target player has not left and all bots are in the same lobby?")
			client.Move(client.controlData.Angle, 0)
			client.controlData.SplitCt = 0
			client.Statery.EmotionalSupportPlayerName = nil
			client.Statery.EmotionalSupportPlayerAccountID = -1
			return
		}
		// aka dead
		if target.GetMass() < 3 {
			client.Move(client.controlData.Angle, 0)
			client.controlData.SplitCt = 0
			return
		}
		circleDist := client.TargetPlayer(player, target)
		_ = circleDist
		client.Split()
	} else if client.Statery.IsFarmingHoles {
		client.TargetHoles(player, enums.GoodHoles())
		client.Split()
	} else {
		client.Move(client.controlData.Angle, 0)
		client.controlData.SplitCt = 0
	}
}
