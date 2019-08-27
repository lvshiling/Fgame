package arenapvp

import (
	arenapvppb "fgame/fgame/cross/arenapvp/pb"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
)

//跨服数据转换
func convertFromAreanapvpPlayerList(playerList []*arenapvppb.ArenapvpBattlePlayer) (dataList []*arenapvpdata.PvpPlayerInfo) {
	dataList = make([]*arenapvpdata.PvpPlayerInfo, 0, len(playerList))
	for _, playerInfo := range playerList {
		dataList = append(dataList, convertFromAreanapvpPlayer(playerInfo))
	}
	return dataList
}

//跨服数据转换
func convertFromAreanapvpPlayer(playerInfo *arenapvppb.ArenapvpBattlePlayer) *arenapvpdata.PvpPlayerInfo {
	data := arenapvpdata.NewPvpPlayerInfo()
	data.Platform = playerInfo.Platform
	data.ServerId = playerInfo.ServerId
	data.PlayerId = playerInfo.PlayerId
	data.PlayerName = playerInfo.PlayerName
	data.Role = playerInfo.Role
	data.Sex = playerInfo.Sex
	data.WingId = playerInfo.WingId
	data.WeaponId = playerInfo.WeaponId
	data.FashionId = playerInfo.FashionId
	data.WeaponState = playerInfo.WeaponState
	data.XianTiId = playerInfo.XianTiId
	data.LingYuId = playerInfo.LingYuId
	data.FaBaoId = playerInfo.FaBaoId

	data.BattleDataList = convertFromArenapvpBattleDataInfo(playerInfo.ResultList)
	return data
}

func convertFromArenapvpBattleDataInfo(resultList []*arenapvppb.BattleResult) (battleResultList []*arenapvpdata.BattleResultData) {
	for _, resultInfo := range resultList {
		result := &arenapvpdata.BattleResultData{}
		result.BattleId1 = resultInfo.BattleId1
		result.BattleId2 = resultInfo.BattleId2
		result.WinnerId = resultInfo.WinnerId
		result.PvpType = arenapvptypes.ArenapvpType(resultInfo.PvpType)
		result.Index = resultInfo.Index

		battleResultList = append(battleResultList, result)
	}

	return
}

//
func convertToGuessData(data *arenapvppb.GuessData) *arenapvpdata.GuessData {
	d := &arenapvpdata.GuessData{}
	d.PvpType = arenapvptypes.ArenapvpType(data.PvpType)
	d.RaceNumber = data.RaceNumber
	d.PlayerList = convertFromAreanapvpPlayerList(data.PlayerList)
	return d
}

func convertToGuessDataList(dataList []*arenapvppb.GuessData) (infoList []*arenapvpdata.GuessData) {
	for _, data := range dataList {
		infoList = append(infoList, convertToGuessData(data))
	}
	return infoList
}

//
func convertFromArenapvpElectionDataList(daList []*arenapvppb.ElectionData) (dataList []*arenapvpdata.ElectionData) {
	for _, d := range daList {
		dataList = append(dataList, convertFromArenapvpElectionData(d))
	}

	return dataList
}

func convertFromArenapvpElectionData(data *arenapvppb.ElectionData) *arenapvpdata.ElectionData {
	d := &arenapvpdata.ElectionData{}
	d.ElectionIndex = data.ElectionIndex
	d.PlNumber = data.PlNumber
	d.LastLuckyTime = data.LastLuckyTime
	d.LuckyNameText = data.LuckyNameText
	return d
}

//跨服数据转换
func convertToBaZhuList(baZhuList []*arenapvppb.BaZhuData) (dataList []*arenapvpdata.BaZhuData) {
	for _, data := range baZhuList {
		bazhu := &arenapvpdata.BaZhuData{}
		bazhu.Platform = data.Platform
		bazhu.PlayerId = data.PlayerId
		bazhu.PlayerName = data.PlayerName
		bazhu.ServerId = data.ServerId
		bazhu.Role = data.Role
		bazhu.Sex = data.Sex
		bazhu.RaceNumber = data.RaceNumber
		bazhu.WingId = data.WingId
		bazhu.WeaponId = data.WeaponId
		bazhu.FashionId = data.FashionId

		dataList = append(dataList, bazhu)
	}
	return dataList
}
