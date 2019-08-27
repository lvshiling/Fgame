package grpc_pbutil

import (
	treasureboxpb "fgame/fgame/cross/treasurebox/pb"
	"fgame/fgame/cross/treasurebox/treasurebox"
)

func BuildTreasureBoxInfoList(logBoxList []*treasurebox.TreasureBoxLogObject) (boxLogList []*treasureboxpb.BoxLogInfo) {
	boxLogList = make([]*treasureboxpb.BoxLogInfo, 0, 8)
	for _, logBox := range logBoxList {
		boxLogList = append(boxLogList, buildTreasureBoxInfo(logBox))
	}
	return
}

func buildTreasureBoxInfo(logBox *treasurebox.TreasureBoxLogObject) *treasureboxpb.BoxLogInfo {
	boxLogInfo := &treasureboxpb.BoxLogInfo{}
	boxLogInfo.ServerId = logBox.ServerId
	boxLogInfo.PlayerName = logBox.PlayerName
	boxLogInfo.LastTime = logBox.LastTime
	for itemId, num := range logBox.ItemMap {
		boxLogInfo.ItemList = append(boxLogInfo.ItemList, buildItem(itemId, num))
	}
	return boxLogInfo
}

func buildItem(itemId int32, num int32) *treasureboxpb.ItemInfo {
	itemInfo := &treasureboxpb.ItemInfo{}
	itemInfo.ItemId = itemId
	itemInfo.Num = num
	return itemInfo
}
