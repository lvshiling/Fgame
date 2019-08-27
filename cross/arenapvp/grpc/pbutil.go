package grpc

import (
	"fgame/fgame/cross/arenapvp/arenapvp"
	arenapvppb "fgame/fgame/cross/arenapvp/pb"
	arenapvpscene "fgame/fgame/cross/arenapvp/scene"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	"fgame/fgame/game/scene/scene"
)

func BuildArenapvpData(playerList []*arenapvpdata.PvpPlayerInfo, baZhuList []*arenapvp.ArenapvpBaZhuObject, guessDataList []*arenapvpdata.GuessData, sceneMap map[int64]scene.Scene) (data *arenapvppb.ArenapvpData) {
	data = &arenapvppb.ArenapvpData{}
	data.PlayerList = buildArenapvpBattlePlayerList(playerList)
	data.BaZhuList = buildBaZhuDataList(baZhuList)
	data.ElectionDataList = buildElectionDataList(sceneMap)
	data.GuessDataList = buildGuessDataList(guessDataList)
	return data
}

func buildArenapvpBattlePlayerList(playerList []*arenapvpdata.PvpPlayerInfo) (infoList []*arenapvppb.ArenapvpBattlePlayer) {
	for _, playerInfo := range playerList {
		infoList = append(infoList, buildArenapvpBattlePlayer(playerInfo))
	}
	return infoList
}

func buildArenapvpBattlePlayer(playerInfo *arenapvpdata.PvpPlayerInfo) *arenapvppb.ArenapvpBattlePlayer {
	info := &arenapvppb.ArenapvpBattlePlayer{}
	info.ServerId = playerInfo.ServerId
	info.PlayerId = playerInfo.PlayerId
	info.PlayerName = playerInfo.PlayerName
	info.Platform = playerInfo.Platform
	info.Role = playerInfo.Role
	info.Sex = playerInfo.Sex
	info.WingId = playerInfo.WingId
	info.WeaponId = playerInfo.WeaponId
	info.FashionId = playerInfo.FashionId
	info.WeaponState = playerInfo.WeaponState
	info.XianTiId = playerInfo.XianTiId
	info.LingYuId = playerInfo.LingYuId
	info.FaBaoId = playerInfo.FaBaoId

	info.ResultList = buildBattleResultList(playerInfo.BattleDataList)

	return info
}

func buildBaZhuDataList(baZhuList []*arenapvp.ArenapvpBaZhuObject) (infoList []*arenapvppb.BaZhuData) {
	for _, baZhu := range baZhuList {

		info := &arenapvppb.BaZhuData{}
		info.RaceNumber = baZhu.RaceNumber
		info.Platform = baZhu.PlayerPlatform
		info.ServerId = baZhu.PlayerServerId
		info.PlayerId = baZhu.PlayerId
		info.PlayerName = baZhu.PlayerName
		info.Role = baZhu.Role
		info.Sex = baZhu.Sex
		info.WingId = baZhu.WingId
		info.WeaponId = baZhu.WeaponId
		info.FashionId = baZhu.FashionId

		infoList = append(infoList, info)
	}
	return infoList
}

func buildBattleResultList(resultList []*arenapvpdata.BattleResultData) (infoList []*arenapvppb.BattleResult) {
	for _, data := range resultList {

		info := &arenapvppb.BattleResult{}
		info.WinnerId = data.WinnerId
		info.BattleId1 = data.BattleId1
		info.BattleId2 = data.BattleId2
		info.PvpType = int32(data.PvpType)
		info.Index = data.Index

		infoList = append(infoList, info)

	}
	return infoList
}

func buildElectionDataList(sceneMap map[int64]scene.Scene) (infoList []*arenapvppb.ElectionData) {

	for _, s := range sceneMap {
		sd, ok := s.SceneDelegate().(arenapvpscene.ArenapvpSceneData)
		if !ok {
			continue
		}
		index := sd.GetElectionIndex()
		num := int32(len(s.GetAllPlayers()))
		lastLuckTime, lastLuckPlayerList := sd.GetLastLuckRewInfo()
		LuckyNameText := ""
		for _, spl := range lastLuckPlayerList {
			if len(LuckyNameText) == 0 {
				LuckyNameText += spl.GetName()
			} else {
				LuckyNameText += ", " + spl.GetName()
			}
		}

		info := &arenapvppb.ElectionData{}
		info.ElectionIndex = index
		info.PlNumber = num
		info.LastLuckyTime = lastLuckTime
		info.LuckyNameText = LuckyNameText

		infoList = append(infoList, info)
	}
	return infoList
}

func buildGuessDataList(guessDataList []*arenapvpdata.GuessData) (infoList []*arenapvppb.GuessData) {
	for _, data := range guessDataList {
		infoList = append(infoList, buildGuessData(data))
	}
	return infoList
}

func buildGuessData(guessData *arenapvpdata.GuessData) *arenapvppb.GuessData {
	info := &arenapvppb.GuessData{}
	if guessData != nil {
		info.PvpType = int32(guessData.PvpType)
		info.RaceNumber = guessData.RaceNumber

		for _, battlePl := range guessData.PlayerList {
			info.PlayerList = append(info.PlayerList, buildArenapvpBattlePlayer(battlePl))
		}
	}

	return info
}
