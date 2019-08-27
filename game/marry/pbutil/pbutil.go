package pbuitl

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	marrymarry "fgame/fgame/game/marry/marry"
	playermarry "fgame/fgame/game/marry/player"
	marryscene "fgame/fgame/game/marry/scene"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playercommon "fgame/fgame/game/player/common"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

const (
	resultSucess  = int32(1)
	resultFailure = int32(2)
)

const (
	//婚礼巡游
	wedStart = int32(1)
	//婚宴
	wedBanquet = int32(2)
	//婚礼结束
	wedEnd = int32(3)
)

//TODO ylz:多线程引用
func BuildSCMarryGet(pl player.Player, marryInfo *playermarry.PlayerMarryObject, period int32, isProposal bool, playerSuitMap map[int32]map[int32]int32, spouseSuitMap map[int32]map[int32]int32) *uipb.SCMarryInfo {
	marryGet := &uipb.SCMarryInfo{}
	status := int32(marryInfo.Status)
	marryGet.Status = &status
	marryGet.Period = &period
	marryGet.IsProposal = &isProposal
	marryGet.MemList = append(marryGet.MemList, buildMember(pl, marryInfo, playerSuitMap))
	if marryInfo.Status != marrytypes.MarryStatusTypeUnmarried &&
		marryInfo.Status != marrytypes.MarryStatusTypeDivorce {
		spousePl := player.GetOnlinePlayerManager().GetPlayerById(marryInfo.SpouseId)
		if spousePl != nil {
			manager := spousePl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
			spouseInfo := manager.GetMarryInfo()
			marryGet.MemList = append(marryGet.MemList, buildMember(spousePl, spouseInfo, spouseSuitMap))
		} else {
			playerInfo, err := player.GetPlayerService().GetPlayerInfo(marryInfo.SpouseId)
			if err == nil {
				marryGet.MemList = append(marryGet.MemList, buildMemberByPlayerInfo(playerInfo, spouseSuitMap))
			}
		}
	}
	return marryGet
}

func BuildSCMarryPushWedCard(wedCard *marrymarry.MarryWedCardObject) *uipb.SCMarryPushWedCard {
	marryPushWedCard := &uipb.SCMarryPushWedCard{}
	marryPushWedCard.WedCard = buildWedCard(wedCard)
	return marryPushWedCard
}

func BuildSCMarryAfterLogin(wedCardList []*marrymarry.MarryWedCardObject) *uipb.SCMarryAfterLogin {
	marryAfterLogin := &uipb.SCMarryAfterLogin{}
	for _, wedCard := range wedCardList {
		marryAfterLogin.WedCardList = append(marryAfterLogin.WedCardList, buildWedCard(wedCard))
	}
	return marryAfterLogin
}

func BuildSCMarryProposal(result int32) *uipb.SCMarryProposal {
	marryProposal := &uipb.SCMarryProposal{}
	marryProposal.Result = &result
	return marryProposal
}

func BuildSCMarryPushProposal(pl player.Player, ring int32) *uipb.SCMarryPushProposal {
	marryPushProposal := &uipb.SCMarryPushProposal{}

	playerId := pl.GetId()
	name := pl.GetName()
	role := int32(pl.GetRole())
	sex := int32(pl.GetSex())
	fashionId := pl.GetFashionId()

	marryPushProposal.PlayerId = &playerId
	marryPushProposal.Name = &name
	marryPushProposal.Role = &role
	marryPushProposal.Sex = &sex
	marryPushProposal.FashionId = &fashionId
	marryPushProposal.Ring = &ring
	return marryPushProposal
}

func BuildSCMarryProposalResult(agree bool, name string) *uipb.SCMarryProposalResult {
	marryProposalResult := &uipb.SCMarryProposalResult{}
	result := resultFailure
	if agree {
		result = resultSucess
	}
	marryProposalResult.Result = &result
	marryProposalResult.Name = &name
	return marryProposalResult
}

func BuildSCMarryDivorce(typ int32, online bool) *uipb.SCMarryDivorce {
	marryDivorce := &uipb.SCMarryDivorce{}
	result := resultFailure
	if online {
		result = resultSucess
	}
	marryDivorce.Typ = &typ
	marryDivorce.Result = &result
	return marryDivorce
}

func BuildSCMarryDivorceDealPushPeer(name string) *uipb.SCMarryDivorceDealPushPeer {
	marryDivorceDealPushPeer := &uipb.SCMarryDivorceDealPushPeer{}
	marryDivorceDealPushPeer.Name = &name
	return marryDivorceDealPushPeer
}

func BuildSCMarryInfoStatusChange(status int32) *uipb.SCMarryInfoChange {
	marryInfoChange := &uipb.SCMarryInfoChange{}
	marryInfoChange.Status = &status
	return marryInfoChange
}

func BuildSCMarryInfoChangeWedding(status int32, period int32) *uipb.SCMarryInfoChange {
	marryInfoChange := &uipb.SCMarryInfoChange{}
	marryInfoChange.Status = &status
	marryInfoChange.Period = &period
	return marryInfoChange
}

func BuildSCMarryPushDivorce(name string) *uipb.SCMarryPushDivorce {
	marryPushDivorce := &uipb.SCMarryPushDivorce{}
	marryPushDivorce.Name = &name
	return marryPushDivorce
}

func BuildSCMarryRingReplace(ring int32) *uipb.SCMarryRingReplace {
	marryRingReplace := &uipb.SCMarryRingReplace{}
	marryRingReplace.Ring = &ring
	return marryRingReplace
}

func BuildSCMarryRingFeed(level int32, progress int32) *uipb.SCMarryRingFeed {
	marryRingFeed := &uipb.SCMarryRingFeed{}
	marryRingFeed.RLevel = &level
	marryRingFeed.Progress = &progress
	return marryRingFeed
}

func BuildSCMarryTreeFeed(level int32, progress int32) *uipb.SCMarryTreeFeed {
	marryTreeFeed := &uipb.SCMarryTreeFeed{}
	marryTreeFeed.TLevel = &level
	marryTreeFeed.Progress = &progress
	return marryTreeFeed
}

func BuildSCMarryWedGrade(result int32, marryGrade *marrytypes.MarryGrade, period int32) *uipb.SCMarryWedGrade {
	marryWedGrade := &uipb.SCMarryWedGrade{}
	marryWedGrade.Result = &result

	marryWedGrade.Grade = buildGrade(marryGrade)
	marryWedGrade.Period = &period
	return marryWedGrade
}

func BuildSCMarryWedGift(buffId int32, period int32, grade int32, autoFlag bool) *uipb.SCMarryWedGift {
	marryWedGift := &uipb.SCMarryWedGift{}
	marryWedGift.BuffId = &buffId
	marryWedGift.Period = &period
	marryWedGift.Grade = &grade
	marryWedGift.AutoFlag = &autoFlag
	return marryWedGift
}

func BuildSCMarryForceChange(playerId int64, force int64) *uipb.SCMarrySpouseInfoChange {
	marrySpouseInfoChange := &uipb.SCMarrySpouseInfoChange{}
	marrySpouseInfoChange.PlayerId = &playerId
	marrySpouseInfoChange.Force = &force
	return marrySpouseInfoChange
}

func BuildSCMarryLevelChange(playerId int64, level int32) *uipb.SCMarrySpouseInfoChange {
	marrySpouseInfoChange := &uipb.SCMarrySpouseInfoChange{}
	marrySpouseInfoChange.PlayerId = &playerId
	marrySpouseInfoChange.Level = &level
	return marrySpouseInfoChange
}

func BuildSCMarryTLevelChange(playerId int64, tLevel int32) *uipb.SCMarrySpouseInfoChange {
	marrySpouseInfoChange := &uipb.SCMarrySpouseInfoChange{}
	marrySpouseInfoChange.PlayerId = &playerId
	marrySpouseInfoChange.TLevel = &tLevel
	return marrySpouseInfoChange
}

func BuildSCMarryRLevelChange(playerId int64, rLevel int32) *uipb.SCMarrySpouseInfoChange {
	marrySpouseInfoChange := &uipb.SCMarrySpouseInfoChange{}
	marrySpouseInfoChange.PlayerId = &playerId
	marrySpouseInfoChange.RLevel = &rLevel
	return marrySpouseInfoChange
}

func BuildSCMarryRingChange(playerId int64, ring int32) *uipb.SCMarrySpouseInfoChange {
	marrySpouseInfoChange := &uipb.SCMarrySpouseInfoChange{}
	marrySpouseInfoChange.PlayerId = &playerId
	marrySpouseInfoChange.Ring = &ring
	return marrySpouseInfoChange
}

func BuildSCMarryFashionChange(playerId int64, fashionId int32) *uipb.SCMarrySpouseInfoChange {
	marrySpouseInfoChange := &uipb.SCMarrySpouseInfoChange{}
	marrySpouseInfoChange.PlayerId = &playerId
	marrySpouseInfoChange.FashionId = &fashionId
	return marrySpouseInfoChange
}

func BuildSCMarryNameChange(playerId int64, name string) *uipb.SCMarrySpouseInfoChange {
	marrySpouseInfoChange := &uipb.SCMarrySpouseInfoChange{}
	marrySpouseInfoChange.PlayerId = &playerId
	marrySpouseInfoChange.SpouseName = &name
	return marrySpouseInfoChange
}

func BuildSCMarryDevelopLevelChange(playerId int64, developLevel int32) *uipb.SCMarrySpouseInfoChange {
	marrySpouseInfoChange := &uipb.SCMarrySpouseInfoChange{}
	marrySpouseInfoChange.PlayerId = &playerId
	marrySpouseInfoChange.DLevel = &developLevel
	return marrySpouseInfoChange
}

func BuildSCMarryDingQingChange(playerId int64, suitMap map[int32]map[int32]int32) *uipb.SCMarrySpouseInfoChange {
	marrySpouseInfoChange := &uipb.SCMarrySpouseInfoChange{}
	marrySpouseInfoChange.PlayerId = &playerId
	marrySpouseInfoChange.DingQingList = buildMarryDingQingSuitInfoList(suitMap)
	return marrySpouseInfoChange
}

func BuildSCMarryWingChange(playerId int64, wingId int32) *uipb.SCMarrySpouseInfoChange {
	marrySpouseInfoChange := &uipb.SCMarrySpouseInfoChange{}
	marrySpouseInfoChange.PlayerId = &playerId
	marrySpouseInfoChange.WingId = &wingId
	return marrySpouseInfoChange
}

func BuildSCMarryWeaponChange(playerId int64, weaponId int32) *uipb.SCMarrySpouseInfoChange {
	marrySpouseInfoChange := &uipb.SCMarrySpouseInfoChange{}
	marrySpouseInfoChange.PlayerId = &playerId
	marrySpouseInfoChange.WeaponId = &weaponId
	return marrySpouseInfoChange
}

func BuildSCMarryMarryCountChange(playerId int64, marryCount int32) *uipb.SCMarrySpouseInfoChange {
	marrySpouseInfoChange := &uipb.SCMarrySpouseInfoChange{}
	marrySpouseInfoChange.PlayerId = &playerId
	marrySpouseInfoChange.MarryCount = &marryCount
	return marrySpouseInfoChange
}

func BuildSCMarryViewWedCard(wedCardId int64) *uipb.SCMarryViewWedCard {
	marryViewWedCard := &uipb.SCMarryViewWedCard{}
	marryViewWedCard.CardId = &wedCardId
	return marryViewWedCard
}

func BuildSCMarryWedList(marryWedObjList []*marrymarry.MarryWedObject) *uipb.SCMarryWedList {
	marryWedList := &uipb.SCMarryWedList{}
	for _, marryWedObj := range marryWedObjList {
		if marryWedObj.Period == -1 {
			continue
		}
		marryWedList.WedList = append(marryWedList.WedList, buildWedding(marryWedObj))
	}
	return marryWedList
}

func BuildSCMarryCancle(period int32) *uipb.SCMarryCancel {
	marryCancel := &uipb.SCMarryCancel{}
	marryCancel.Period = &period
	return marryCancel
}

func BuildSCMarryWedPushStatus(sd *marryscene.MarrySceneStatusData, isShow bool) *uipb.SCMarryWedPushStatus {
	marryWedPushStatus := &uipb.SCMarryWedPushStatus{}
	status := wedStart

	switch sd.Status {
	case marryscene.MarrySceneStatusTypeInit:
		{
			status = wedEnd
			break
		}
	case marryscene.MarrySceneStatusBanquet:
		{
			status = wedBanquet
			break
		}
	case marryscene.MarrySceneStatusCruise:
		{
			status = wedStart
		}
	}

	period := sd.Period
	playerId := sd.PlayerId
	name := sd.PlayerName
	role := sd.Role
	sex := sd.Sex

	spouseId := sd.SpouseId
	spouseName := sd.SpouseName
	sRole := sd.SpouseRole
	sSex := sd.SpouseSex
	marryWedPushStatus.IsShow = &isShow
	marryWedPushStatus.Status = &status
	marryWedPushStatus.Period = &period
	marryWedPushStatus.PlayerId = &playerId
	marryWedPushStatus.Name = &name
	marryWedPushStatus.Role = &role
	marryWedPushStatus.Sex = &sex
	marryWedPushStatus.Name = &name

	marryWedPushStatus.SpouseId = &spouseId
	marryWedPushStatus.SpouseName = &spouseName
	marryWedPushStatus.SRole = &sRole
	marryWedPushStatus.SSex = &sSex
	return marryWedPushStatus
}

func BuildSCMarryCancleToOther(marryObj *marrymarry.MarryObject, period int32, isOnline bool, sIsOnline bool, isShow bool) *uipb.SCMarryCancelToOther {
	scMarryCancelToOther := &uipb.SCMarryCancelToOther{}

	playerId := marryObj.PlayerId
	spouseId := marryObj.SpouseId
	playerName := marryObj.PlayerName
	spouseName := marryObj.SpouseName
	role := marryObj.Role
	spouseRole := marryObj.SpouseRole
	sex := marryObj.Sex
	spouseSex := marryObj.SpouseSex

	scMarryCancelToOther.IsShow = &isShow
	scMarryCancelToOther.Period = &period
	scMarryCancelToOther.PlayerId = &playerId
	scMarryCancelToOther.Name = &playerName
	scMarryCancelToOther.Role = &role
	scMarryCancelToOther.Sex = &sex
	scMarryCancelToOther.IsOnline = &isOnline
	scMarryCancelToOther.SpouseId = &spouseId
	scMarryCancelToOther.SpouseName = &spouseName
	scMarryCancelToOther.SRole = &spouseRole
	scMarryCancelToOther.SSex = &spouseSex
	scMarryCancelToOther.SIsOnline = &sIsOnline
	return scMarryCancelToOther
}

func BuildSCMarryBanquet(period int32, playerId int64, name string, spouseId int64, spouseName string, heroismList []*marryscene.MarryHeroism) *uipb.SCMarryBanquet {
	marryBanquet := &uipb.SCMarryBanquet{}
	marryBanquet.Period = &period
	marryPlayer := buildMarryWedInfo(playerId, name)
	marrySpouse := buildMarryWedInfo(spouseId, spouseName)
	marryBanquet.PlayerList = append(marryBanquet.PlayerList, marryPlayer, marrySpouse)

	for index, heroism := range heroismList {
		if index >= marrytypes.HeroisTopLen {
			break
		}
		marryBanquet.HeroismList = append(marryBanquet.HeroismList, buildHeroism(heroism))
	}
	return marryBanquet
}

func buildMarryWedInfo(playerId int64, playerName string) *uipb.MarryWedInfo {
	marryWedInfo := &uipb.MarryWedInfo{}
	marryWedInfo.PlayerId = &playerId
	marryWedInfo.Name = &playerName
	return marryWedInfo
}

func BuildSCMarryHeroismTopThree(heroismList []*marryscene.MarryHeroism) *uipb.SCMarryHeroismTopThree {
	marryHeroismTopThree := &uipb.SCMarryHeroismTopThree{}

	for index, heroism := range heroismList {
		if index >= marrytypes.HeroisTopLen {
			break
		}
		marryHeroismTopThree.HeroismList = append(marryHeroismTopThree.HeroismList, buildHeroism(heroism))
	}
	return marryHeroismTopThree
}

func BuildSCMarryWorish() *uipb.SCMarryWorship {
	marryWorship := &uipb.SCMarryWorship{}
	return marryWorship
}

func BuildSCMarryWedGiftList(heroismList []*marryscene.MarryHeroism) *uipb.SCMarryWedGiftList {
	marryWedGiftList := &uipb.SCMarryWedGiftList{}

	for _, heroism := range heroismList {
		marryWedGiftList.HeroismList = append(marryWedGiftList.HeroismList, buildHeroism(heroism))
	}
	return marryWedGiftList
}

func BuildSCMarryWedEnd(marryData *marryscene.MarryData) *uipb.SCMarryWedEnd {
	scMarryWedEnd := &uipb.SCMarryWedEnd{}
	playerId := marryData.PlayerId
	playerName := marryData.PlayerName
	playerRole := marryData.PlayerRole
	playerSex := marryData.PlayerSex
	spouseId := marryData.SpouseId
	spouseName := marryData.SpouseName
	spouseRole := marryData.SpouseRole
	spouseSex := marryData.SpouseSex
	scMarryWedEnd.PlayerId = &playerId
	scMarryWedEnd.PlayerName = &playerName
	scMarryWedEnd.PlayerRole = &playerRole
	scMarryWedEnd.PlayerSex = &playerSex

	scMarryWedEnd.SpouseId = &spouseId
	scMarryWedEnd.SpouseName = &spouseName
	scMarryWedEnd.SpouseRole = &spouseRole
	scMarryWedEnd.SpouseSex = &spouseSex
	return scMarryWedEnd
}

func BuildSCMarryRecomment(friendIdList []int64, pList []scene.Player) *uipb.SCMarryRecommended {
	scMarryRecommended := &uipb.SCMarryRecommended{}
	for _, p := range pList {
		isFriend := false
		flag := utils.ContainInt64(friendIdList, p.GetId())
		if flag {
			isFriend = true
		}
		scMarryRecommended.RecommendList = append(scMarryRecommended.RecommendList, buildRecomment(isFriend, p))
	}
	return scMarryRecommended

}

func BuildSCMaryClickCar(clickTime int64, grade int32) *uipb.SCMarryClickCar {
	scMarryClickCar := &uipb.SCMarryClickCar{}
	scMarryClickCar.ClickTime = &clickTime
	scMarryClickCar.Grade = &grade
	return scMarryClickCar
}

func BuildSCMarryWedGradeToSpouse(period int32, playerName string, wedGrade *marrytypes.MarryGrade) *uipb.SCMarryWedGradeToSpouse {
	scMarryWedGradeToSpouse := &uipb.SCMarryWedGradeToSpouse{}
	scMarryWedGradeToSpouse.Period = &period
	scMarryWedGradeToSpouse.Name = &playerName
	scMarryWedGradeToSpouse.Grade = buildGrade(wedGrade)
	return scMarryWedGradeToSpouse
}

func BuildSCMarryWedGradeSpouseDeal(result bool) *uipb.SCMarryWedGradeSpouseDeal {
	scMarryWedGradeSpouseDeal := &uipb.SCMarryWedGradeSpouseDeal{}
	scMarryWedGradeSpouseDeal.Result = &result
	return scMarryWedGradeSpouseDeal
}

func BuildSCMarryWedGradeRefuseToPeer() *uipb.SCMarryWedGradeRefuseToPeer {
	scMarryWedGradeRefuseToPeer := &uipb.SCMarryWedGradeRefuseToPeer{}
	return scMarryWedGradeRefuseToPeer
}

func BuildSCMarryWedSucess(period int32) *uipb.SCMarryWedSucess {
	scMarryWedSucess := &uipb.SCMarryWedSucess{}
	scMarryWedSucess.Period = &period
	return scMarryWedSucess
}

// func BuildSCMarryDevelopSendGift(friendId int64, itemId, num int32, auto bool) *uipb.SCMarryDevelopSendGift {
// 	scMsg := &uipb.SCMarryDevelopSendGift{}
// 	scMsg.FriendId = &friendId
// 	scMsg.ItemId = &itemId
// 	scMsg.Num = &num
// 	scMsg.Auto = &auto
// 	return scMsg
// }

func BuildSCMarryDevelopUplevel(developLevel, developExp int32) *uipb.SCMarryDevelopUplevel {
	scMsg := &uipb.SCMarryDevelopUplevel{}
	scMsg.DevelopLevel = &developLevel
	scMsg.DevelopExp = &developExp
	return scMsg
}

func buildGrade(wedGrade *marrytypes.MarryGrade) *uipb.MarryGrade {
	marryGrade := &uipb.MarryGrade{}
	grade := int32(wedGrade.Grade)
	hunCheGrade := int32(wedGrade.HunCheGrade)
	sugarGrade := int32(wedGrade.SugarGrade)
	marryGrade.Grade = &grade
	marryGrade.HunCheGrade = &hunCheGrade
	marryGrade.SugarGrade = &sugarGrade
	return marryGrade
}

func buildRecomment(isFriend bool, p scene.Player) *uipb.MarryRecommend {
	marryRecommend := &uipb.MarryRecommend{}
	playerId := p.GetId()
	name := p.GetName()
	level := p.GetLevel()
	force := p.GetForce()
	role := int32(p.GetRole())
	sex := int32(p.GetSex())
	fashionId := p.GetFashionId()
	weaponId := p.GetWeaponId()

	marryRecommend.PlayerId = &playerId
	marryRecommend.Name = &name
	marryRecommend.Level = &level
	marryRecommend.Force = &force
	marryRecommend.Role = &role
	marryRecommend.Sex = &sex
	marryRecommend.FashionId = &fashionId
	marryRecommend.WeaponId = &weaponId
	marryRecommend.IsFriend = &isFriend
	return marryRecommend
}

func buildHeroism(heroismObj *marryscene.MarryHeroism) *uipb.MarryHeroism {
	marryHeroism := &uipb.MarryHeroism{}
	name := heroismObj.Name
	heroism := heroismObj.Heroism

	marryHeroism.Name = &name
	marryHeroism.Heroism = &heroism
	return marryHeroism
}

func buildWedding(marryWedObj *marrymarry.MarryWedObject) *uipb.MarryWed {
	marryWed := &uipb.MarryWed{}
	period := marryWedObj.Period
	name := marryWedObj.Name
	spouseName := marryWedObj.SpouseName
	marryWed.Period = &period
	marryWed.NameList = append(marryWed.NameList, name, spouseName)
	return marryWed
}

func buildMember(pl player.Player, marryInfo *playermarry.PlayerMarryObject, suitMap map[int32]map[int32]int32) *uipb.MarryMember {
	member := &uipb.MarryMember{}
	playerId := pl.GetId()
	name := pl.GetName()
	level := pl.GetLevel()
	force := pl.GetForce()
	role := int32(pl.GetRole())
	sex := int32(pl.GetSex())
	fashionId := pl.GetFashionId()
	ring := int32(0)
	itemTemplate := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWedRing, marryInfo.Ring.ItemWedRingSubType())
	if itemTemplate != nil {
		ring = int32(itemTemplate.TemplateId())
	}
	rLevel := marryInfo.RingLevel
	rProgress := marryInfo.RingExp
	tLevel := marryInfo.TreeLevel
	tProgress := marryInfo.TreeExp
	dLevel := marryInfo.GetDevelopLevel()
	dExp := marryInfo.GetDevelopExp()
	weaponId := pl.GetAllWeaponInfo().Wear
	wingId := pl.GetWingId()
	marryCount := marryInfo.MarryCount

	isProposal := false
	if marryInfo.IsProposal == 1 {
		isProposal = true
	}

	member.PlayerId = &playerId
	member.Name = &name
	member.Level = &level
	member.Force = &force
	member.Role = &role
	member.Sex = &sex
	member.FashionId = &fashionId
	member.Ring = &ring
	member.RLevel = &rLevel
	member.RProgress = &rProgress
	member.TLevel = &tLevel
	member.TProgress = &tProgress
	member.IsProposal = &isProposal
	member.DLevel = &dLevel
	member.DExp = &dExp
	member.DingQingList = buildMarryDingQingSuitInfoList(suitMap)
	member.WeaponId = &weaponId
	member.WingId = &wingId
	member.MarryCount = &marryCount

	return member
}

func buildMarryDingQingSuitInfoList(suitMap map[int32]map[int32]int32) []*uipb.MarryDingQingSuitInfo {
	rst := make([]*uipb.MarryDingQingSuitInfo, 0)
	for suitId, posMap := range suitMap {
		mySuitId := suitId
		for posId, _ := range posMap {
			myPostId := posId
			item := &uipb.MarryDingQingSuitInfo{}
			item.SuitId = &mySuitId
			item.PosId = &myPostId
			rst = append(rst, item)
		}
	}
	return rst
}

func buildMemberByPlayerInfo(playerInfo *playercommon.PlayerInfo, suitMap map[int32]map[int32]int32) *uipb.MarryMember {
	member := &uipb.MarryMember{}
	marryInfo := playerInfo.MarryInfo
	playerId := playerInfo.PlayerId
	name := playerInfo.Name
	level := playerInfo.Level
	force := playerInfo.Force
	role := int32(playerInfo.Role)
	sex := int32(playerInfo.Sex)
	fashionId := playerInfo.FashionId

	ring := int32(0)
	itemTemplate := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWedRing, marrytypes.MarryRingType(marryInfo.Ring).ItemWedRingSubType())
	if itemTemplate != nil {
		ring = int32(itemTemplate.TemplateId())
	}
	rLevel := marryInfo.RLevel
	rProgress := marryInfo.RProgress
	tLevel := marryInfo.TLevel
	tProgress := marryInfo.TProgress
	dLevel := marryInfo.DLevel
	dExp := marryInfo.DExp
	weaponId := playerInfo.AllWeaponInfo.Wear
	wingId := playerInfo.WingInfo.WingId
	marryCount := marryInfo.MarryCount

	isProposal := false
	if marryInfo.IsProposal == 1 {
		isProposal = true
	}

	member.PlayerId = &playerId
	member.Name = &name
	member.Level = &level
	member.Force = &force
	member.Role = &role
	member.Sex = &sex
	member.FashionId = &fashionId
	member.Ring = &ring
	member.RLevel = &rLevel
	member.RProgress = &rProgress
	member.TLevel = &tLevel
	member.TProgress = &tProgress
	member.IsProposal = &isProposal
	member.DLevel = &dLevel
	member.DExp = &dExp
	member.DingQingList = buildMarryDingQingSuitInfoList(suitMap)
	member.WeaponId = &weaponId
	member.WingId = &wingId
	member.MarryCount = &marryCount
	return member
}

func buildWedCard(wedCard *marrymarry.MarryWedCardObject) *uipb.MarryWedCard {
	marryWedCard := &uipb.MarryWedCard{}
	playerId := wedCard.PlayerId
	spouseId := wedCard.SpouseId
	role := int32(1)
	sex := int32(1)
	wRole := int32(1)
	wSex := int32(1)
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl != nil {
		role = int32(pl.GetRole())
		sex = int32(pl.GetSex())
	} else {
		playerInfo, _ := player.GetPlayerService().GetPlayerInfo(playerId)
		role = int32(playerInfo.Role)
		sex = int32(playerInfo.Sex)
	}
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	if spl != nil {
		wRole = int32(spl.GetRole())
		wSex = int32(spl.GetSex())
	} else {
		playerInfo, _ := player.GetPlayerService().GetPlayerInfo(spouseId)
		wRole = int32(playerInfo.Role)
		wSex = int32(playerInfo.Sex)
	}

	cardId := wedCard.Id
	mName := wedCard.PlayerName
	wName := wedCard.SpouseName
	hTime := wedCard.HoldTime
	marryWedCard.CardId = &cardId
	marryWedCard.PlayerId = &playerId
	marryWedCard.Role = &role
	marryWedCard.Sex = &sex
	marryWedCard.MName = &mName
	marryWedCard.SpouseId = &spouseId
	marryWedCard.WName = &wName
	marryWedCard.WRole = &wRole
	marryWedCard.WSex = &wSex
	marryWedCard.HTime = &hTime
	return marryWedCard
}

func BuildSCMarryPreGift(itemMap map[int32]int32) *uipb.SCMarryPreGift {
	rst := &uipb.SCMarryPreGift{}
	rst.ItemId = make([]int32, 0)
	rst.ItemNum = make([]int32, 0)
	for key, value := range itemMap {
		rst.ItemId = append(rst.ItemId, key)
		rst.ItemNum = append(rst.ItemNum, value)
	}
	return rst
}

func BuildSCMarryPreGiftMsg(playerId int64, playerName string, giftType int32, exp int64) *uipb.SCMarryPreGiftMsg {
	rst := &uipb.SCMarryPreGiftMsg{}
	rst.PlayerId = &playerId
	rst.PlayerName = &playerName
	rst.GiftType = &giftType
	rst.Exp = &exp
	return rst
}

func BuildScMarryJiNianList(jiNianMap map[marrytypes.MarryBanquetSubTypeWed]*playermarry.PlayerMarryJiNianObject) *uipb.SCMarryJiNianMsg {
	rst := &uipb.SCMarryJiNianMsg{}
	rst.JiNianList = make([]*uipb.SCMarryJiNianInfo, 0)
	for _, value := range jiNianMap {
		item := &uipb.SCMarryJiNianInfo{}
		typeId := int32(value.JiNianType)
		typeCount := int32(value.JiNianCount)
		item.JiNianType = &typeId
		item.JiNianCount = &typeCount
		rst.JiNianList = append(rst.JiNianList, item)
	}
	return rst
}

func BuildSCMarryDingQingJiHuoMsg(spouseId int64, hasFlag bool, suitId int32, posId int32) *uipb.SCMarryDingQingJiHuoMsg {
	rst := &uipb.SCMarryDingQingJiHuoMsg{}
	result := int32(1)
	rst.Result = &result
	rst.SpouseHasFlag = &hasFlag
	rst.SpouseId = &spouseId
	rst.SuitId = &suitId
	rst.PosId = &posId
	return rst
}

func BuildSCMarryDingQingSuoYaoMsg(result int32) *uipb.SCMarryDingQingSuoYaoMsg {
	rst := &uipb.SCMarryDingQingSuoYaoMsg{}
	rst.Result = &result
	return rst
}

func BuildSCMarryDingQingSuoYaoRspMsg(spouseId int64, spouseName string, suitId int32, posId int32, content string) *uipb.SCMarryDingQingSuoYaoRspMsg {
	rst := &uipb.SCMarryDingQingSuoYaoRspMsg{}
	rst.SpouseId = &spouseId
	rst.SpouseName = &spouseName
	rst.SuitId = &suitId
	rst.PosId = &posId
	rst.Content = &content
	return rst
}

func BuildSCMarryDingQingSuoYaoDealMsg(flag int32) *uipb.SCMarryDingQingSuoYaoDealMsg {
	rst := &uipb.SCMarryDingQingSuoYaoDealMsg{}
	rst.Result = &flag
	return rst
}

func BuildSCMarryDingQingZengSongDealMsg(flag int32) *uipb.SCMarryDingQingZengSongMsg {
	rst := &uipb.SCMarryDingQingZengSongMsg{}
	rst.Result = &flag
	return rst
}

func BuildSCMarryBanquetSetChangeMsg(houtai marrytypes.MarryHoutaiType) *uipb.SCMarryBanquetSetChangeMsg {
	rst := &uipb.SCMarryBanquetSetChangeMsg{}
	houtaiInt := int32(houtai)
	rst.HoutaiId = &houtaiInt
	return rst
}

func BuildSCMarryDingQingYueMsg() *uipb.SCMarryDingQingYueMsg {
	rst := &uipb.SCMarryDingQingYueMsg{}
	return rst
}

func BuildSCMarryDingQingYueSpouseMsg(playerId int64, suitId int32, posId int32) *uipb.SCMarryDingQingYueSpouseMsg {
	rst := &uipb.SCMarryDingQingYueSpouseMsg{}
	rst.PlayerId = &playerId
	rst.SuitId = &suitId
	rst.PosId = &posId
	return rst
}
