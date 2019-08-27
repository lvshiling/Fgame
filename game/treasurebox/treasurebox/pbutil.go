package treasurebox

import (
	treasureboxpb "fgame/fgame/cross/treasurebox/pb"
	droptemplate "fgame/fgame/game/drop/template"
)

func convertFromOpenInfo(serverId int32, playerName string, lastTime int64, itemMap []*droptemplate.DropItemData) *treasureboxpb.TreasureBoxOpenLogRequest {
	req := &treasureboxpb.TreasureBoxOpenLogRequest{}
	req.ServerId = serverId
	req.PlayerName = playerName
	req.LastTime = lastTime
	itemList := make([]*treasureboxpb.ItemInfo, 0, 8)

	for _, itemData := range itemMap {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()

		itemInfo := convertItemInfo(itemId, num)
		itemList = append(itemList, itemInfo)
	}

	req.ItemList = itemList
	return req

}

func convertItemInfo(itemId int32, num int32) *treasureboxpb.ItemInfo {
	itemInfo := &treasureboxpb.ItemInfo{}
	itemInfo.ItemId = itemId
	itemInfo.Num = num
	return itemInfo
}

func convertFromBoxLogInfo(boxLog *treasureboxpb.BoxLogInfo) *BoxLogInfo {
	boxLogInfo := &BoxLogInfo{}
	boxLogInfo.ServerId = boxLog.ServerId
	boxLogInfo.PlayerName = boxLog.PlayerName
	boxLogInfo.LastTime = boxLog.LastTime
	boxLogInfo.ItemMap = make(map[int32]int32)

	for _, itemInfo := range boxLog.GetItemList() {
		boxLogInfo.ItemMap[itemInfo.ItemId] = itemInfo.Num
	}
	return boxLogInfo
}

func convertFromBoxLogInfoList(boxLogInfoList []*treasureboxpb.BoxLogInfo) []*BoxLogInfo {
	logBoxList := make([]*BoxLogInfo, 0, len(boxLogInfoList))
	for _, boxLogInfo := range boxLogInfoList {
		logBoxList = append(logBoxList, convertFromBoxLogInfo(boxLogInfo))
	}
	return logBoxList
}
