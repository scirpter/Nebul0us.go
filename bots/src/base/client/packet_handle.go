package client

import (
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/base/packets"
)

func HandlePacket(client *Client, packet *[]byte) {
	id := enums.PACKET_TYPE((*packet)[0])
	var allNets []*specs.Net
	for _, appclient := range client.app.Clients {
		allNets = append(allNets, appclient.net)
	}

	switch id {
	case enums.PACKET_TYPE_CONNECT_RESULT_2:
		packets.NewConnectResult2Packet(client.net, packet, client.eventDispatcher).Parse()

	case enums.PACKET_TYPE_CLAN_CHAT_MESSAGE:
		packets.NewClanChatMessagePacket(packet, client.eventDispatcher).Parse()

	case enums.PACKET_TYPE_CLIENT_PREFERENCES:
		packets.NewClientPreferencesPacket(client.net, packet, client.eventDispatcher).Parse()

	case enums.PACKET_TYPE_GAME_UPDATE:
		if client.Statery.TaskLock {
			return // avoid any mutex issues
		}
		packets.NewGameUpdatePacket(client.net, allNets, packet, client.eventDispatcher).Parse()

	case enums.PACKET_TYPE_TOP_SCORES:
		packets.NewTopScoresPacket(client.net, packet).Parse()

	case enums.PACKET_TYPE_GAME_DATA:
		packets.NewGameDataPacket(client.net, allNets, client.eventDispatcher, packet).Parse()

	case enums.PACKET_TYPE_SESSION_STATS:
		packets.NewSessionStatsPacket(packet, client.eventDispatcher).Parse()

	case enums.PACKET_TYPE_JOIN_RESULT:
		packets.NewJoinResultPacket(packet, client.eventDispatcher, client.net, client.Statery).Parse()

	case enums.PACKET_TYPE_GAME_CHAT_MESSAGE:
		packets.NewGameChatMessagePacket(nil, client.blobData, packet, client.eventDispatcher, nil).Parse()

	case enums.PACKET_TYPE_ENTER_GAME_RESULT:
		packets.NewEnterGameResultPacket(packet, client.net, client.eventDispatcher).Parse()

	case enums.PACKET_TYPE_BATTLE_ROYALE_STATUS_UPDATE:
		packets.NewBattleRoyaleStatusUpdatePacket(packet, client.Statery, client.eventDispatcher).Parse()

	case enums.PACKET_TYPE_BATTLE_ROYALE_LIST_RESULT:
		packets.NewBattleRoyaleListResultPacket(packet, client.eventDispatcher).Parse()

	default:
		// loggin.ERROR("No handle for %d", id)
	}
}
