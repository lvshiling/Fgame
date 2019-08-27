package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/core/utils"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/friend/friend"
	playerfriend "fgame/fgame/game/friend/player"
	"fgame/fgame/game/global"
	"fgame/fgame/game/marry/dao"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/marry"
	marrymarry "fgame/fgame/game/marry/marry"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marryscene "fgame/fgame/game/marry/scene"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	manager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	curRing := marryInfo.Ring

	//同步配偶名字
	if marryInfo.SpouseId != 0 {
		spouseName := marry.GetMarryService().GetSpouseName(pl.GetId())

		manager.SpouseNameChanged(spouseName)

		//同步配偶表白等级
		developLevel := marry.GetMarryService().GetSpouseDevelopLevel(pl.GetId())
		manager.UpdateCoupleMarryDevelopLevel(developLevel)
	}

	//婚戒返还
	marryRingObj := marry.GetMarryService().GetMarryProposalRing(pl.GetId())
	if marryRingObj != nil && marryRingObj.Status == marrytypes.MarryRingStatusTypeFail {
		marrylogic.PlayerMarryRingGiveBack(pl, marrytypes.MarryRingType(marryRingObj.Ring), marryRingObj.PeerName)
	}

	//下线时婚戒替换过
	ringType := marry.GetMarryService().GetMarryRing(pl.GetId())
	if curRing < ringType {
		manager.RingReplacedBySpouse(ringType)
	}

	now := global.GetGame().GetTimeService().Now()
	//下线后判断婚烟状态改变
	if err = marryChangeOffonline(pl, now); err != nil {
		return
	}

	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	sd := marry.GetMarryService().GetMarrySceneData()
	//判断是否好友或者同一仙盟
	isShowHunYan := false
	if pl.GetId() != sd.PlayerId && pl.GetId() != sd.SpouseId {
		if friendManager.IsFriend(sd.PlayerId) || friendManager.IsFriend(sd.SpouseId) {
			isShowHunYan = true
		} else { //统一仙盟
			marryPlayerInfo, _ := player.GetPlayerService().GetPlayerInfo(sd.PlayerId)
			if marryPlayerInfo != nil && marryPlayerInfo.AllianceId == pl.GetAllianceId() && pl.GetAllianceId() != 0 {
				isShowHunYan = true
			}
			spousePlayerInfo, _ := player.GetPlayerService().GetPlayerInfo(sd.SpouseId)
			if spousePlayerInfo != nil && spousePlayerInfo.AllianceId == pl.GetAllianceId() && pl.GetAllianceId() != 0 {
				isShowHunYan = true
			}
		}
	}

	//婚礼按钮推送
	if sd.Status != marryscene.MarrySceneStatusTypeInit {
		pushWedRecord := manager.GetPushWedRecord()
		scMarryWedPushStar := pbuitl.BuildSCMarryWedPushStatus(sd, true)
		scHiddenMarryWedPushStar := pbuitl.BuildSCMarryWedPushStatus(sd, false)
		switch sd.Status {
		case marryscene.MarrySceneStatusCruise:
			{
				if pushWedRecord.WedId != sd.Id {
					manager.PushWedRecordHunChe(sd.Id)
					pl.SendMsg(scMarryWedPushStar)
				} else {
					pl.SendMsg(scHiddenMarryWedPushStar)
				}
				break
			}
		case marryscene.MarrySceneStatusBanquet:
			{
				if pushWedRecord.BanquetTime == 0 || pushWedRecord.WedId != sd.Id {
					if isShowHunYan {
						manager.PushWedRecordBanquet(sd.Id)
						pl.SendMsg(scMarryWedPushStar)
					} else {
						pl.SendMsg(scHiddenMarryWedPushStar)
					}
				} else {
					pl.SendMsg(scHiddenMarryWedPushStar)
				}
				break
			}
		}
	}
	//同步玩家婚宴状态(婚宴礼服使用)
	wedStatus := marrytypes.MarryWedStatusSelfTypeNo
	if pl.GetId() == sd.PlayerId || pl.GetId() == sd.SpouseId {
		switch sd.Status {
		case marryscene.MarrySceneStatusCruise:
			{
				wedStatus = marrytypes.MarryWedStatusSelfTypeCruise
				break
			}
		case marryscene.MarrySceneStatusBanquet:
			{
				wedStatus = marrytypes.MarryWedStatusSelfTypeBanquet
				break
			}
		}
	}
	manager.SynchronousWedStatus(wedStatus)

	//获取喜帖
	viewCardList := manager.GetViewCardList()
	wedCardList := marry.GetMarryService().GetWeddingCardList()
	wedCardIdList := make([]int64, 0, 12)
	needViewWedCardList := make([]*marrymarry.MarryWedCardObject, 0, 12)
	for _, wedCard := range wedCardList {
		wedPlayerId := wedCard.PlayerId
		wedSpouseId := wedCard.SpouseId
		isFriend := friendManager.IsFriend(wedPlayerId)
		isSpouseFriend := friendManager.IsFriend(wedSpouseId)

		wedPlayerInfo, err := player.GetPlayerService().GetPlayerInfo(wedPlayerId)
		if err != nil {
			return err
		}
		wedSpousePlayerInfo, err := player.GetPlayerService().GetPlayerInfo(wedSpouseId)
		if err != nil {
			return err
		}

		wedAllianceId := wedPlayerInfo.AllianceId
		wedSpouseAllianceId := wedSpousePlayerInfo.AllianceId

		if isFriend || isSpouseFriend ||
			(pl.GetAllianceId() == wedAllianceId && wedAllianceId != 0) ||
			(pl.GetAllianceId() == wedSpouseAllianceId && wedSpouseAllianceId != 0) {
			flag := utils.ContainInt64(viewCardList, wedCard.Id)
			if flag {
				continue
			}
			if wedPlayerId == pl.GetId() || pl.GetId() == wedSpouseId {
				continue
			}
			wedCardIdList = append(wedCardIdList, wedCard.Id)
			needViewWedCardList = append(needViewWedCardList, wedCard)
		}
	}

	scMarryAfterLogin := pbuitl.BuildSCMarryAfterLogin(needViewWedCardList)
	pl.SendMsg(scMarryAfterLogin)

	period := marry.GetMarryService().GetWeddingPeriod(pl.GetId())

	//定情信物同步
	playerSuitMap := manager.GetAllDingQingMap()
	marrymarry.GetMarryService().SyncMarryDingQing(pl.GetId(), playerSuitMap)

	// playerSuitMap := marry.GetMarryService().GetMarryDingQing(pl.GetId())
	spoutId := marry.GetMarryService().GetSpouseId(pl.GetId())
	var spoutSuiteMap map[int32]map[int32]int32
	if spoutId > 0 {
		spoutSuiteMap = marry.GetMarryService().GetMarryDingQing(spoutId)
		manager.UpdateSpouseSuit(spoutSuiteMap) //加载的时候更新
	}

	scMarryGet := pbuitl.BuildSCMarryGet(pl, marryInfo, period, false, playerSuitMap, spoutSuiteMap)
	pl.SendMsg(scMarryGet)
	//下发结婚类型信息
	houtai := marrytemplate.GetMarryTemplateService().GetHouTaiType()
	scMarryBanquetSet := pbuitl.BuildSCMarryBanquetSetChangeMsg(houtai)
	pl.SendMsg(scMarryBanquetSet)
	return
}

//下线期间婚烟状态改变
func marryChangeOffonline(pl player.Player, now int64) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	status := marryInfo.Status
	spouseId := marryInfo.SpouseId
	curRingLevel := marryInfo.RingLevel
	marryObj := marry.GetMarryService().GetMarry(pl.GetId())
	//强制离婚和协议离婚
	if marryObj == nil {
		switch status {
		case marrytypes.MarryStatusTypeProposal,
			marrytypes.MarryStatusTypeEngagement,
			marrytypes.MarryStatusTypeMarried:
			{
				manager.Divorce()
				if err = divorceDealSubPoint(pl, spouseId, now); err != nil {
					return
				}
				break
			}
		}
	} else {
		marryStatus := marryObj.Status
		//求婚成功
		switch marryStatus {
		case marrytypes.MarryStatusTypeProposal,
			marrytypes.MarryStatusTypeEngagement:
			isProposal := false
			spouseId := marryObj.PlayerId
			spouseName := marryObj.PlayerName
			if marryObj.PlayerId == pl.GetId() {
				isProposal = true
				spouseId = marryObj.SpouseId
				spouseName = marryObj.SpouseName
			}

			if status == marrytypes.MarryStatusTypeUnmarried {
				manager.ProposalMarry(spouseId, spouseName, marryObj.Ring, isProposal)
			}
			break
		case marrytypes.MarryStatusTypeMarried:
			manager.Marry()
			break
		}

		//发生了订婚
		if marryStatus == marrytypes.MarryStatusTypeEngagement && status == marrytypes.MarryStatusTypeProposal {
			manager.DueToWedding()
		}
		//预定婚礼取消
		if marryStatus == marrytypes.MarryStatusTypeProposal && status == marrytypes.MarryStatusTypeEngagement {
			manager.CancleWedding()
		}
		//同步婚戒等级
		marry.GetMarryService().MarryRingLevel(pl.GetId(), curRingLevel)
	}
	return nil
}

//根据离婚类型扣除亲密度
func divorceDealSubPoint(pl player.Player, spouseId int64, now int64) (err error) {
	divorceConsentEntity, err := dao.GetMarryDao().GetConsentDivorce(pl.GetId())
	if err != nil {
		return err
	}

	//协议离婚成功
	if divorceConsentEntity != nil {
		percent := marrytemplate.GetMarryTemplateService().GetMarryDivorceLeftIntimacy()
		friend.GetFriendService().DivorceSubPoint(pl, spouseId, marrytypes.MarryDivorceTypeConsent, percent)
		divorceConsentObj := marry.NewMarryDivorceConsentObject()
		divorceConsentObj.FromEntity(divorceConsentEntity)
		divorceConsentObj.UpdateTime = now
		divorceConsentObj.DeleteTime = now
		divorceConsentObj.SetModified()
	} else { //下线后被强制离婚
		friend.GetFriendService().DivorceSubPoint(pl, spouseId, marrytypes.MarryDivorceTypeConsent, 0)
	}
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
