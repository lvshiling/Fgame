package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/game/chess/chess"
	playerchess "fgame/fgame/game/chess/player"
	chesstemplate "fgame/fgame/game/chess/template"
	chesstypes "fgame/fgame/game/chess/types"
	droptemplate "fgame/fgame/game/drop/template"
)

func BuildSCChessInfoGet(objMap map[chesstypes.ChessType]*playerchess.PlayerChessObject, logList []*chess.ChessLogObject) *uipb.SCChessInfoGet {
	scChessInfoGet := &uipb.SCChessInfoGet{}
	for _, obj := range objMap {
		scChessInfoGet.ChessInfo = append(scChessInfoGet.ChessInfo, buildChessBreifInfo(obj))

	}
	for _, log := range logList {
		scChessInfoGet.LogList = append(scChessInfoGet.LogList, buildChessLog(log))
	}
	return scChessInfoGet
}

func BuildSCChessLogIncr(logList []*chess.ChessLogObject) *uipb.SCChessLogIncr {
	scChessLogIncr := &uipb.SCChessLogIncr{}
	for _, log := range logList {
		scChessLogIncr.LogList = append(scChessLogIncr.LogList, buildChessLog(log))
	}
	return scChessLogIncr
}

func buildChessBreifInfo(obj *playerchess.PlayerChessObject) *uipb.ChessBreifInfo {
	chessBreifInfo := &uipb.ChessBreifInfo{}
	chessTemplate := chesstemplate.GetChessTemplateService().GetChessByTypAndChessId(obj.GetChessType(), obj.GetChessId())
	chessDropId := chessTemplate.DropId
	chessBreifInfo.ChessDropId = &chessDropId
	typ := int32(obj.GetChessType())
	chessBreifInfo.Typ = &typ
	useTimes := obj.GetAttendTimes()
	chessBreifInfo.UseTimes = &useTimes
	totalUseTimes := int32(obj.GetToatalAttendTimes())
	chessBreifInfo.TotalUseTimes = &totalUseTimes

	return chessBreifInfo
}

func buildChessLog(log *chess.ChessLogObject) *uipb.ChessLog {
	chessLog := &uipb.ChessLog{}
	playerName := log.GetPlayerName()
	itemId := log.GetItemId()
	itemNum := log.GetItemNum()
	time := log.GetUpdateTime()
	chessLog.PlayerName = &playerName
	chessLog.ItemId = &itemId
	chessLog.ItemNum = &itemNum
	chessLog.CreateTime = &time

	return chessLog
}

func BuildSCChessAttend(rewItemList []*droptemplate.DropItemData, typ chesstypes.ChessType, logList []*chess.ChessLogObject, autoFlag bool) *uipb.SCChessAttend {
	scChessAttend := &uipb.SCChessAttend{}
	for i := int(0); i < len(rewItemList); i++ {
		itemId := rewItemList[i].GetItemId()
		num := rewItemList[i].GetNum()
		level := rewItemList[i].GetLevel()

		scChessAttend.DropInfo = append(scChessAttend.DropInfo, buildDropInfo(itemId, num, level))
	}

	chessType := int32(typ)
	scChessAttend.Typ = &chessType
	scChessAttend.AutoFlag = &autoFlag
	for _, log := range logList {
		scChessAttend.LogList = append(scChessAttend.LogList, buildChessLog(log))
	}
	return scChessAttend
}

func BuildSCChessAttendBatch(rewItemList []*droptemplate.DropItemData, typ chesstypes.ChessType, logList []*chess.ChessLogObject, autoFlag bool) *uipb.SCChessAttendBatch {
	scChessAttendBatch := &uipb.SCChessAttendBatch{}

	for i := int(0); i < len(rewItemList); i++ {
		itemId := rewItemList[i].GetItemId()
		num := rewItemList[i].GetNum()
		level := rewItemList[i].GetLevel()

		scChessAttendBatch.DropList = append(scChessAttendBatch.DropList, buildDropInfo(itemId, num, level))
	}

	chessType := int32(typ)
	scChessAttendBatch.Typ = &chessType
	scChessAttendBatch.AutoFlag = &autoFlag
	for _, log := range logList {
		scChessAttendBatch.LogList = append(scChessAttendBatch.LogList, buildChessLog(log))
	}
	return scChessAttendBatch
}

func buildDropInfo(itemId, num, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level

	return dropInfo
}

func BuildSCChessChanged(dropId int32) *uipb.SCChessChanged {
	scChessChanged := &uipb.SCChessChanged{}
	scChessChanged.ChessDropId = &dropId

	return scChessChanged
}
