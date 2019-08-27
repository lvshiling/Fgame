package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	inventorypbutil "fgame/fgame/game/inventory/pbutil"
	playertrade "fgame/fgame/game/trade/player"
	"fgame/fgame/game/trade/trade"
	tradetypes "fgame/fgame/game/trade/types"
)

func BuildSCTradeInfoList(globalTradeItemObjList []*trade.GlobalTradeItemObject, currentPage int32, totalPage int32) *uipb.SCTradeInfoList {
	scTradeInfoList := &uipb.SCTradeInfoList{}
	scTradeInfoList.TradeInfoList = BuildGlobalTradeInfoList(globalTradeItemObjList)
	scTradeInfoList.TotalPage = &totalPage
	scTradeInfoList.CurrentPage = &currentPage
	return scTradeInfoList
}

func BuildGlobalTradeInfoList(globalTradeItemObjList []*trade.GlobalTradeItemObject) []*uipb.GlobalTradeInfo {
	infoList := make([]*uipb.GlobalTradeInfo, 0, len(globalTradeItemObjList))
	for _, globalTradeItemObj := range globalTradeItemObjList {
		infoList = append(infoList, BuildGlobalTradeInfo(globalTradeItemObj))
	}
	return infoList
}

func BuildGlobalTradeInfo(globalTradeItemObj *trade.GlobalTradeItemObject) *uipb.GlobalTradeInfo {
	info := &uipb.GlobalTradeInfo{}
	tradeId := globalTradeItemObj.GetId()
	serverId := globalTradeItemObj.GetServerId()
	playerId := globalTradeItemObj.GetPlayerId()
	playerName := globalTradeItemObj.GetPlayerName()
	itemId := globalTradeItemObj.GetItemId()
	itemNum := globalTradeItemObj.GetItemNum()
	propertyData := globalTradeItemObj.GetPropertyData()
	gold := globalTradeItemObj.GetGold()
	createTime := globalTradeItemObj.GetCreateTime()
	info.TradeId = &tradeId
	info.ServerId = &serverId
	info.PlayerId = &playerId
	info.PlayerName = &playerName
	info.ItemId = &itemId
	info.ItemNum = &itemNum
	info.Gold = &gold
	info.CreateTime = &createTime
	info.PropertyData = inventorypbutil.BuildItemPropertyData(propertyData)
	lev := globalTradeItemObj.GetLevel()
	info.Level = &lev
	return info
}

func BuildSCSelfTradeInfoList(tradeItemList []*trade.TradeItemObject) *uipb.SCSelfTradeInfoList {
	scSelfTradeInfoList := &uipb.SCSelfTradeInfoList{}
	scSelfTradeInfoList.TradeInfoList = BuildTradeInfoList(tradeItemList)
	return scSelfTradeInfoList
}

func BuildTradeInfoList(tradeItemObjList []*trade.TradeItemObject) []*uipb.TradeInfo {
	infoList := make([]*uipb.TradeInfo, 0, len(tradeItemObjList))
	for _, tradeItemObj := range tradeItemObjList {
		infoList = append(infoList, BuildTradeInfo(tradeItemObj))
	}
	return infoList
}

func BuildTradeInfo(tradeItemObj *trade.TradeItemObject) *uipb.TradeInfo {
	info := &uipb.TradeInfo{}
	tradeId := tradeItemObj.GetId()
	itemId := tradeItemObj.GetItemId()
	itemNum := tradeItemObj.GetNum()
	propertyData := tradeItemObj.GetPropertyData()
	gold := tradeItemObj.GetGold()
	createTime := tradeItemObj.GetCreateTime()
	info.TradeId = &tradeId
	info.ItemId = &itemId
	info.ItemNum = &itemNum
	info.PropertyData = inventorypbutil.BuildItemPropertyData(propertyData)
	info.Gold = &gold
	info.CreateTime = &createTime
	info.PropertyData = inventorypbutil.BuildItemPropertyData(propertyData)
	lev := tradeItemObj.GetLevel()
	info.Level = &lev
	status := int32(tradeItemObj.GetStatus())
	info.Status = &status
	return info
}

func BuildSCTradeUploadItem(tradeItemObj *trade.TradeItemObject) *uipb.SCTradeUploadItem {
	scTradeUploadItem := &uipb.SCTradeUploadItem{}
	tradeInfo := BuildTradeInfo(tradeItemObj)
	scTradeUploadItem.TradeInfo = tradeInfo
	return scTradeUploadItem
}

func BuildSCTradeWithDrawItem(tradeId int64) *uipb.SCTradeWithDrawItem {
	scTradeWithDrawItem := &uipb.SCTradeWithDrawItem{}
	scTradeWithDrawItem.TradeId = &tradeId
	return scTradeWithDrawItem
}

func BuildSCTradeItem(tradeId int64) *uipb.SCTradeItem {
	scTradeItem := &uipb.SCTradeItem{}
	scTradeItem.TradeId = &tradeId
	return scTradeItem
}

func BuildSCTradeLogList(tradeLogList []*playertrade.PlayerTradeLogObject) *uipb.SCTradeLogList {
	scTradeLogList := &uipb.SCTradeLogList{}
	for _, tradeLog := range tradeLogList {
		tradeLogMsg := BuildTradeLog(tradeLog)
		if tradeLogMsg == nil {
			continue
		}
		scTradeLogList.TradeInfoList = append(scTradeLogList.TradeInfoList, tradeLogMsg)
	}
	return scTradeLogList
}

func BuildTradeLog(tradeLog *playertrade.PlayerTradeLogObject) *uipb.TradeLog {
	tradeLogMsg := &uipb.TradeLog{}
	logType := int32(tradeLog.GetLogType())
	switch tradeLog.GetLogType() {
	case tradetypes.TradeLogTypeBuy:
		tradeLogMsg.BuyLog = BuildTradeBuyLog(tradeLog)
		break
	case tradetypes.TradeLogTypeSell:
		tradeLogMsg.SellLog = BuildTradeSellLog(tradeLog)
		break
	default:
		return nil
	}
	tradeLogMsg.LogType = &logType
	return tradeLogMsg
}

func BuildTradeSellLog(tradeLog *playertrade.PlayerTradeLogObject) *uipb.TradeSellLog {
	tradeSellLog := &uipb.TradeSellLog{}
	itemId := tradeLog.GetItemId()
	tradeSellLog.ItemId = &itemId
	itemNum := tradeLog.GetItemNum()
	tradeSellLog.ItemNum = &itemNum
	tradeSellLog.PropertyData = inventorypbutil.BuildItemPropertyData(tradeLog.GetPropertyData())
	tradeTime := tradeLog.GetCreateTime()
	tradeSellLog.TradeTime = &tradeTime
	buyServerId := tradeLog.GetBuyServerId()
	tradeSellLog.BuyServerId = &buyServerId
	buyPlayerId := tradeLog.GetBuyPlayerId()
	tradeSellLog.BuyPlayerId = &buyPlayerId
	buyPlayerName := tradeLog.GetBuyPlayerName()
	tradeSellLog.BuyPlayerName = &buyPlayerName
	gold := tradeLog.GetGold()
	tradeSellLog.BuyGold = &gold
	getGold := tradeLog.GetGetGold()
	tradeSellLog.GetGold = &getGold
	fee := tradeLog.GetFee()
	tradeSellLog.Fee = &fee
	level := tradeLog.GetLevel()
	tradeSellLog.Level = &level
	feeRate := tradeLog.GetFeeRate()
	tradeSellLog.FeeRate = &feeRate
	return tradeSellLog
}

func BuildTradeBuyLog(tradeLog *playertrade.PlayerTradeLogObject) *uipb.TradeBuyLog {
	tradeBuyLog := &uipb.TradeBuyLog{}
	itemId := tradeLog.GetItemId()
	tradeBuyLog.ItemId = &itemId
	itemNum := tradeLog.GetItemNum()
	tradeBuyLog.ItemNum = &itemNum
	tradeBuyLog.PropertyData = inventorypbutil.BuildItemPropertyData(tradeLog.GetPropertyData())
	tradeTime := tradeLog.GetCreateTime()
	tradeBuyLog.TradeTime = &tradeTime
	sellServerId := tradeLog.GetSellServerId()
	tradeBuyLog.SellServerId = &sellServerId
	sellPlayerId := tradeLog.GetSellPlayerId()
	tradeBuyLog.SellPlayerId = &sellPlayerId
	sellPlayerName := tradeLog.GetSellPlayerName()
	tradeBuyLog.SellPlayerName = &sellPlayerName
	gold := tradeLog.GetGold()
	tradeBuyLog.BuyGold = &gold
	level := tradeLog.GetLevel()
	tradeBuyLog.Level = &level
	return tradeBuyLog
}
