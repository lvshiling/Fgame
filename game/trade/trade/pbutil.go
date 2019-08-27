package trade

import (
	trade_serverpb "fgame/fgame/trade_server/pb"
)

func ConvertFromGlobalTradeItemList(globalTradeItemList []*trade_serverpb.GlobalTradeItem) (objList []*GlobalTradeItemObject, err error) {
	for _, globalTradeItem := range globalTradeItemList {
		obj, err := ConvertFromGlobalTradeItem(globalTradeItem)
		if err != nil {
			return nil, err
		}
		if obj == nil {
			continue
		}
		objList = append(objList, obj)
	}
	return objList, nil
}

func ConvertFromGlobalTradeItem(globalTradeItem *trade_serverpb.GlobalTradeItem) (obj *GlobalTradeItemObject, err error) {
	obj, err = NewGlobalTradeItemObject(
		globalTradeItem.GetGlobalTradeId(),
		globalTradeItem.GetTradeId(),
		globalTradeItem.GetPlatform(),
		globalTradeItem.GetServerId(),
		globalTradeItem.GetPlayerId(),
		globalTradeItem.GetPlayerName(),
		globalTradeItem.GetItemId(),
		globalTradeItem.GetItemNum(),
		globalTradeItem.GetGold(),
		globalTradeItem.GetPropertyData(),
		globalTradeItem.GetLevel(),
		globalTradeItem.GetBuyPlayerPlatform(),
		globalTradeItem.GetBuyPlayerServerId(),
		globalTradeItem.GetBuyPlayerId(),
		globalTradeItem.GetBuyPlayerName(),
		globalTradeItem.GetUpdateTime(),
		globalTradeItem.GetCreateTime(),
		globalTradeItem.GetDeleteTime(),
	)
	return
}
