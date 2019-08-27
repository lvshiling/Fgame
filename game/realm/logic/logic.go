package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/friend/friend"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemskilllogic "fgame/fgame/game/itemskill/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	realmeventtypes "fgame/fgame/game/realm/event/types"
	"fgame/fgame/game/realm/pbutil"
	playerrealm "fgame/fgame/game/realm/player"
	"fgame/fgame/game/realm/realm"
	realmtemplate "fgame/fgame/game/realm/template"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func CheckIfCanEnterTianJieTa(pl player.Player) (flag bool) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	s := pl.GetScene()
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeTianJieTa {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)

	if manager.IfFullLevel() {
		return
	}
	return true
}

//获取前往击杀界面信息的逻辑
func HandleTianJieTa(pl player.Player) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	s := pl.GetScene()
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeTianJieTa {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	flag := manager.IfFullLevel()
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("realm:境界已达最高层")
		playerlogic.SendSystemMessage(pl, lang.RealmReachLimit)
		return
	}

	curlevel := manager.GetTianJieTaLevel()
	nextLevel := curlevel + 1
	//调用场景接口
	_, flag = PlayerEnterTianJieTa(pl, 0, nextLevel)
	if !flag {
		panic(fmt.Errorf("realm: realmToKill  should be ok"))
	}
	scSoulRuinsChallenge := pbutil.BuildSCRealmToKill(nextLevel)
	pl.SendMsg(scSoulRuinsChallenge)
	return
}

func PlayerEnterTianJieTa(pl player.Player, spouseId int64, level int32) (s scene.Scene, flag bool) {
	tianJieTaTemplate := realmtemplate.GetRealmTemplateService().GetTianJieTaTemplateByLevel(level)
	if tianJieTaTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"level":    level,
			}).Warn("realm:处理跳转天劫塔,天截塔不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	sh := CreateTienJieTaSceneData(pl.GetId(), spouseId, tianJieTaTemplate)
	s = scene.CreateFuBenScene(tianJieTaTemplate.MapId, sh)
	if s == nil {
		panic(fmt.Errorf("realm:创建副本应该成功"))
	}
	scenelogic.PlayerEnterSingleFuBenScene(pl, s)
	flag = true
	return
}

//TODO 奖励
func onTianJieTaFinish(pl player.Player, spouseId int64, tianJieTaTemplate *gametemplate.TianJieTaTemplate, successful bool) (err error) {
	if tianJieTaTemplate == nil {
		return
	}
	err = ToKillFinished(pl, spouseId, tianJieTaTemplate, successful)
	return
}

//下发场景信息
func onPushSceneInfo(pl player.Player, starTime int64, ownerId int64, spouseId int64, level int32) (err error) {
	//推送给前端
	scRealmScene := pbutil.BuildSCRealmScene(starTime, level, ownerId, spouseId)
	pl.SendMsg(scRealmScene)
	return
}

//前往击杀结束
func ToKillFinished(pl player.Player, spouseId int64, tianJieTaTemplate *gametemplate.TianJieTaTemplate, successful bool) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	beforelevel := manager.GetTianJieTaLevel()
	level := tianJieTaTemplate.Level
	diffLevel := level - beforelevel
	if diffLevel != 1 {
		return nil
	}

	if successful {
		now := global.GetGame().GetTimeService().Now()
		flag := manager.UpgradeTianJieLevel(now)
		if !flag {
			panic(fmt.Errorf("realm: ToKillFinished UpgradeTianJieLevel() should be ok"))
		}

		//刷新天劫塔排名
		realm.GetRealmRankService().RefreshTianJieTaRank(pl.GetId(), pl.GetName(), level, now)

		//通关奖励物品
		rewData := tianJieTaTemplate.GetRewData()
		rewItemMap := tianJieTaTemplate.GetRewItemMap()
		realmPassRewItem(pl, rewData, rewItemMap, level)
		//增加亲密度
		addPoint := tianJieTaTemplate.GetRewQinMiDu()
		addPairPoint(pl, spouseId, addPoint)
	}
	pushRealmResult(pl, spouseId, successful, level)
	//发送事件
	gameevent.Emit(realmeventtypes.EventTypeRealmResult, pl, successful)
	return
}

//通关奖励物品
func realmPassRewItem(pl player.Player, rewData *propertytypes.RewData, rewItemMap map[int32]int32, level int32) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//奖励属性
	if rewData != nil {
		reasonGold := commonlog.GoldLogReasonRealmToKillRew
		reasonSilver := commonlog.SilverLogReasonRealmToKillRew
		reasonLevel := commonlog.LevelLogReasonRealmToKillRew
		reasonGoldText := fmt.Sprintf(reasonGold.String(), level)
		reasonSliverText := fmt.Sprintf(reasonSilver.String(), level)
		reasonlevelText := fmt.Sprintf(reasonLevel.String(), level)

		flag := propertyManager.AddRewData(rewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
		if !flag {
			panic(fmt.Errorf("realm: ToKillFinished AddRewData  should be ok"))
		}
	}
	//奖励物品
	if len(rewItemMap) != 0 {
		flag := inventoryManager.HasEnoughSlots(rewItemMap)
		if !flag {
			//写邮件
			emailTitle := fmt.Sprintf(lang.GetLangService().ReadLang(lang.RealmPassRewTitle), level)
			emailContent := lang.GetLangService().ReadLang(lang.RealmPassRewContent)
			emaillogic.AddEmail(pl, emailTitle, emailContent, rewItemMap)
		} else {
			reasonInventory := commonlog.InventoryLogReasonRealmToKillRew
			reasonInventoryText := fmt.Sprintf(reasonInventory.String(), level)
			flag = inventoryManager.BatchAdd(rewItemMap, reasonInventory, reasonInventoryText)
			if !flag {
				panic(fmt.Errorf("realm: ToKillFinished BatchAdd should be ok"))
			}
			//同步物品
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

	//同步属性
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeRealm.Mask())
	propertylogic.SnapChangedProperty(pl)
	return
}

//夫妻助战增加亲密度
func addPairPoint(pl player.Player, spouseId int64, addPoint int32) {
	if spouseId == 0 {
		return
	}
	friend.GetFriendService().AddPoint(pl, spouseId, addPoint)
}

//发送挑战结果
func pushRealmResult(pl player.Player, spouseId int64, sucessful bool, level int32) {
	if spouseId != 0 {
		scRealmPairResult := pbutil.BuildSCRealmPairResult(true, sucessful, level)
		pl.SendMsg(scRealmPairResult)
		spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
		if spl == nil {
			return
		}
		scRealmPairResult = pbutil.BuildSCRealmPairResult(false, sucessful, level)
		spl.SendMsg(scRealmPairResult)
	} else {
		scRealmToKillResult := pbutil.BuildSCRealmToKillResult(sucessful, level)
		pl.SendMsg(scRealmToKillResult)
	}
	return
}

//检测补偿
func CheckReissue(pl player.Player) {
	realmManager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	level := realmManager.GetTianJieTaLevel()
	if !realmManager.IsCheckReissue() {
		realmManager.SetCheckReissue()
		itemskilllogic.RestitutionHuTiDun(pl) //护体盾
		temp := realmtemplate.GetRealmTemplateService().GetTianJieTaBuChangTemplateByLevel(level)
		if temp == nil {
			return
		}

		attachmentInfo := temp.GetReturnItemMap()
		if len(attachmentInfo) == 0 {
			return
		}

		title := lang.GetLangService().ReadLang(lang.RealmReissueTitle)
		content := lang.GetLangService().ReadLang(lang.RealmReissueContent)
		emaillogic.AddEmail(pl, title, content, attachmentInfo)
	}
	return
}
