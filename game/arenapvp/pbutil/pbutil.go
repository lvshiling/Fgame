package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
)

func BuildSCArenapvpRaceInfo(pvpType arenapvptypes.ArenapvpType, battlePlInfoList []*arenapvpdata.PvpPlayerInfo) *uipb.SCArenapvpRaceInfo {
	scMsg := &uipb.SCArenapvpRaceInfo{}
	typeInt := int32(pvpType)
	scMsg.Type = &typeInt

	maxType := pvpType
	if pvpType == arenapvptypes.ArenapvpTypeTop8 {
		maxType = arenapvptypes.ArenapvpTypeFinals
	}
	for minType := pvpType; minType <= maxType; minType++ {
		scMsg.PlayerList = append(scMsg.PlayerList, buildArenapvpBattlePlayer(minType, battlePlInfoList)...)
	}

	return scMsg
}

func buildArenapvpBattlePlayer(pvpType arenapvptypes.ArenapvpType, battlePlInfoList []*arenapvpdata.PvpPlayerInfo) (infoList []*uipb.ArenapvpBattlePlayer) {

	for _, battlePl := range battlePlInfoList {
		result := battlePl.GetBattleData(pvpType)
		if result == nil {
			continue
		}
		typeInt := int32(pvpType)

		info := &uipb.ArenapvpBattlePlayer{}
		info.Type = &typeInt
		info.Index = &result.Index
		info.WinnerId = &result.WinnerId
		info.BasicInfo = buildBattlePlayerBasicInfo(battlePl)
		infoList = append(infoList, info)

	}
	return infoList
}

func BuildSCArenapvpBaZhuInfo(page, pageNum, totalPage int32, baZhuInfoList []*arenapvpdata.BaZhuData) *uipb.SCArenapvpBaZhuInfo {
	scMsg := &uipb.SCArenapvpBaZhuInfo{}
	scMsg.Page = &page
	scMsg.TotalPage = &totalPage
	scMsg.PageNum = &pageNum
	scMsg.BaZhuList = buildArenapvpBaZhuList(baZhuInfoList)
	return scMsg
}

func buildArenapvpBaZhuList(baZhuInfoList []*arenapvpdata.BaZhuData) (infoList []*uipb.ArenapvpBaZhu) {
	for _, baZhu := range baZhuInfoList {
		infoList = append(infoList, buildArenapvpBaZhu(baZhu))
	}
	return infoList
}

func buildArenapvpBaZhu(baZhuInfo *arenapvpdata.BaZhuData) *uipb.ArenapvpBaZhu {
	info := &uipb.ArenapvpBaZhu{}
	info.BasicInfo = buildBattlePlayerBasicInfoWithBaZhu(baZhuInfo)
	info.RaceNumber = &baZhuInfo.RaceNumber
	return info
}

func buildBattlePlayerBasicInfoWithBaZhu(baZhuInfo *arenapvpdata.BaZhuData) *uipb.BattlePlayerBasicInfo {
	info := &uipb.BattlePlayerBasicInfo{}
	info.Platform = &baZhuInfo.Platform
	info.ServerId = &baZhuInfo.ServerId
	info.PlayerId = &baZhuInfo.PlayerId
	info.PlayerName = &baZhuInfo.PlayerName
	info.Role = &baZhuInfo.Role
	info.Sex = &baZhuInfo.Sex
	info.WingId = &baZhuInfo.WingId
	info.FashionId = &baZhuInfo.FashionId
	info.WeaponId = &baZhuInfo.WeaponId
	return info
}

func buildBattlePlayerBasicInfoList(playerInfoList []*arenapvpdata.PvpPlayerInfo) (infoList []*uipb.BattlePlayerBasicInfo) {
	for _, battlePl := range playerInfoList {
		infoList = append(infoList, buildBattlePlayerBasicInfo(battlePl))
	}
	return infoList
}

func buildBattlePlayerBasicInfo(battlePl *arenapvpdata.PvpPlayerInfo) *uipb.BattlePlayerBasicInfo {
	info := &uipb.BattlePlayerBasicInfo{}
	platform := battlePl.Platform
	info.Platform = &platform
	serverId := battlePl.ServerId
	info.ServerId = &serverId
	playerId := battlePl.PlayerId
	info.PlayerId = &playerId
	playerName := battlePl.PlayerName
	info.PlayerName = &playerName
	role := battlePl.Role
	info.Role = &role
	info.Sex = &battlePl.Sex
	weaponId := battlePl.WeaponId
	info.WeaponId = &weaponId
	wingid := battlePl.WingId
	info.WingId = &wingid
	fashionId := battlePl.FashionId
	info.FashionId = &fashionId
	weaponState := battlePl.WeaponState
	info.WeaponState = &weaponState
	xianTiId := battlePl.XianTiId
	info.XianTiId = &xianTiId
	lingYuId := battlePl.LingYuId
	info.LingYuId = &lingYuId
	faBaoId := battlePl.FaBaoId
	info.FaBaoId = &faBaoId

	return info
}

func BuildSCArenapvpGuessInfo(guessLogObj *playerarenapvp.PlayerArenapvpGuessLogObject, guessData *arenapvpdata.GuessData, guessLogList []*playerarenapvp.PlayerArenapvpGuessLogObject) *uipb.SCArenapvpGuessInfo {
	guessId := int64(0)

	scMsg := &uipb.SCArenapvpGuessInfo{}
	scMsg.GuessId = &guessId
	scMsg.LogList = buildArenapvpGuessLogList(guessLogList)
	if guessData != nil {
		scMsg.PlayerList = buildArenapvpBattlePlayer(guessData.PvpType, guessData.PlayerList)

		if guessLogObj != nil {
			if guessLogObj.GetRaceNum() == guessData.RaceNumber && guessLogObj.GetGuessType() == guessData.PvpType {
				guessId = guessLogObj.GetGuessId()
			}
		}
	}
	return scMsg
}

func buildArenapvpGuessLogList(guessLogList []*playerarenapvp.PlayerArenapvpGuessLogObject) (infoList []*uipb.ArenapvpGuessLog) {
	for _, log := range guessLogList {
		winnerId := log.GetWinnerId()
		if winnerId == 0 {
			continue
		}

		infoList = append(infoList, buildArenapvpGuessLog(log))
	}
	return infoList
}

func buildArenapvpGuessLog(log *playerarenapvp.PlayerArenapvpGuessLogObject) (info *uipb.ArenapvpGuessLog) {

	winnerId := log.GetWinnerId()

	guessType := int32(log.GetGuessType())
	guessId := log.GetGuessId()
	createTime := log.GetCreateTime()
	info = &uipb.ArenapvpGuessLog{}
	info.GuessId = &guessId
	info.GuessType = &guessType
	info.WinnerId = &winnerId
	info.CreateTime = &createTime
	return info
}

func BuildSCArenapvpGuessBeginNotice(guessData *arenapvpdata.GuessData) *uipb.SCArenapvpGuessBeginNotice {
	scMsg := &uipb.SCArenapvpGuessBeginNotice{}
	if guessData != nil {
		scMsg.PlayerList = buildArenapvpBattlePlayer(guessData.PvpType, guessData.PlayerList)
	}
	return scMsg
}

func BuildSCArenapvpGuess(guessId int64) *uipb.SCArenapvpGuess {
	scMsg := &uipb.SCArenapvpGuess{}
	scMsg.GuessId = &guessId
	return scMsg
}

func BuildSCArenapvpInfo(baZhuInfo *arenapvpdata.BaZhuData, obj *playerarenapvp.PlayerArenapvpObject) *uipb.SCArenapvpInfo {
	out := obj.GetOutStatus()
	jifen := obj.GetJiFen()
	notice := obj.GetGuessNotice()
	isTicket := obj.GetTicketFlag()

	scMsg := &uipb.SCArenapvpInfo{}
	scMsg.IsOut = &out
	scMsg.JiFen = &jifen
	scMsg.NoticeSetting = &notice
	scMsg.IsTicket = &isTicket
	if baZhuInfo != nil {
		scMsg.BaZhu = buildArenapvpBaZhu(baZhuInfo)
	}

	return scMsg
}

func BuildSCArenapvpElectionRaceInfo(electionInfoList []*arenapvpdata.ElectionData) *uipb.SCArenapvpElectionRaceInfo {
	scMsg := &uipb.SCArenapvpElectionRaceInfo{}
	for _, election := range electionInfoList {
		scMsg.ElectionList = append(scMsg.ElectionList, buildElectionRaceInfo(election))
	}
	return scMsg
}

func buildElectionRaceInfo(election *arenapvpdata.ElectionData) *uipb.ElectionRaceInfo {
	info := &uipb.ElectionRaceInfo{}
	info.PlayerCount = &election.PlNumber
	info.ElectionIndex = &election.ElectionIndex
	return info
}

func BuildSCArenapvpGuessNoticeSetting(notice int32) *uipb.SCArenapvpGuessNoticeSetting {
	scMsg := &uipb.SCArenapvpGuessNoticeSetting{}
	scMsg.Notice = &notice
	return scMsg
}

func BuildSCArenapvpJiFenChanged(jiFen int32) *uipb.SCArenapvpJiFenChanged {
	scMsg := &uipb.SCArenapvpJiFenChanged{}
	scMsg.JiFen = &jiFen
	return scMsg
}

func BuildSCArenapvpGuessInfoPush(log *playerarenapvp.PlayerArenapvpGuessLogObject) *uipb.SCArenapvpGuessInfoPush {
	scMsg := &uipb.SCArenapvpGuessInfoPush{}
	scMsg.Log = buildArenapvpGuessLog(log)
	return scMsg
}
