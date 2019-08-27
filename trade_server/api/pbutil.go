package api

import (
	"fgame/fgame/trade_server/pb"
	"fgame/fgame/trade_server/trade"
)

func BuildTradeItemList(tradeItemList []*trade.TradeObject) []*pb.GlobalTradeItem {
	tradeObjList := make([]*pb.GlobalTradeItem, 0, len(tradeItemList))
	for _, tradeItem := range tradeItemList {
		tempTradeItem := BuildTradeItem(tradeItem)
		tradeObjList = append(tradeObjList, tempTradeItem)
	}
	return tradeObjList
}

func BuildTradeItem(tradeItem *trade.TradeObject) *pb.GlobalTradeItem {
	globalTradeItem := &pb.GlobalTradeItem{}
	globalTradeItem.GlobalTradeId = tradeItem.GetId()
	globalTradeItem.ServerId = tradeItem.GetServerId()
	globalTradeItem.TradeId = tradeItem.GetTradeId()
	globalTradeItem.PlayerId = tradeItem.GetPlayerId()
	globalTradeItem.PlayerName = tradeItem.GetPlayerName()
	globalTradeItem.ItemId = tradeItem.GetItemId()
	globalTradeItem.ItemNum = tradeItem.GetItemNum()
	globalTradeItem.Gold = tradeItem.GetGold()
	globalTradeItem.PropertyData = tradeItem.GetPropertyData()
	globalTradeItem.BuyPlayerPlatform = tradeItem.GetBuyPlayerPlatform()
	globalTradeItem.BuyPlayerServerId = tradeItem.GetBuyPlayerServerId()
	globalTradeItem.BuyPlayerId = tradeItem.GetBuyPlayerId()
	globalTradeItem.BuyPlayerName = tradeItem.GetBuyPlayerName()
	globalTradeItem.CreateTime = tradeItem.GetCreateTime()
	globalTradeItem.DeleteTime = tradeItem.GetDeleteTime()
	globalTradeItem.UpdateTime = tradeItem.GetUpdateTime()
	globalTradeItem.Level = tradeItem.GetLevel()
	return globalTradeItem
}
