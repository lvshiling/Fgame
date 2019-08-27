package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	baguaeventtypes "fgame/fgame/game/bagua/event/types"
	"fgame/fgame/game/bagua/pbutil"
	playerbagua "fgame/fgame/game/bagua/player"
	baguatemplate "fgame/fgame/game/bagua/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/friend/friend"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//获取前往击杀界面信息的逻辑
func HandleBaGuaToKill(pl player.Player) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeBaGuaMiJing) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("bagua:八卦秘境功能未开启")
		return
	}
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	s := pl.GetScene()
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeBaGuaMiJing {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerBaGuaDataManagerType).(*playerbagua.PlayerBaGuaDataManager)
	flag := manager.IfFullLevel()
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("bagua:八卦秘境已达最高层")
		playerlogic.SendSystemMessage(pl, lang.BaGuaReachLimit)
		return
	}

	curlevel := manager.GetLevel()
	nextLevel := curlevel + 1
	//调用场景接口
	_, flag = PlayerEnterBaGua(pl, 0, nextLevel)
	if !flag {
		panic(fmt.Errorf("bagua: baGuaToKill  should be ok"))
	}
	scBaGuaToKill := pbutil.BuildSCBaGuaToKill(nextLevel)
	pl.SendMsg(scBaGuaToKill)
	return
}

func PlayerEnterBaGua(pl player.Player, spouseId int64, level int32) (s scene.Scene, flag bool) {
	baGuaTemplate := baguatemplate.GetBaGuaTemplateService().GetBaGuaTemplateByLevel(level)
	if baGuaTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"level":    level,
			}).Warn("bagua:处理跳转八卦秘境,八卦秘境不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	sh := CreateBaGuaSceneData(pl.GetId(), spouseId, baGuaTemplate)
	s = scene.CreateFuBenScene(baGuaTemplate.MapId, sh)
	if s == nil {
		panic(fmt.Errorf("bagua:创建副本应该成功"))
	}
	scenelogic.PlayerEnterSingleFuBenScene(pl, s)
	flag = true
	return
}

//TODO 奖励
func onBaGuaFinish(pl player.Player, spouseId int64, tianJieTaTemplate *gametemplate.BaGuaMiJingTemplate, successful bool) (err error) {
	if tianJieTaTemplate == nil {
		return
	}
	err = ToKillFinished(pl, spouseId, tianJieTaTemplate, successful)
	return
}

//下发场景信息
func onPushSceneInfo(pl player.Player, starTime int64, ownerId int64, spouseId int64, level int32) (err error) {
	//推送给前端
	scBaGuaScene := pbutil.BuildSCBaGuaScene(starTime, level, ownerId, spouseId)
	pl.SendMsg(scBaGuaScene)
	return
}

//前往击杀结束
func ToKillFinished(pl player.Player, spouseId int64, baGuaTaTemplate *gametemplate.BaGuaMiJingTemplate, successful bool) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerBaGuaDataManagerType).(*playerbagua.PlayerBaGuaDataManager)
	beforelevel := manager.GetLevel()
	level := baGuaTaTemplate.Level
	diffLevel := level - beforelevel
	if diffLevel != 1 {
		return nil
	}

	if successful {
		flag := manager.UpgradeLevel()
		if !flag {
			panic(fmt.Errorf("bagua: ToKillFinished UpgradeLevel() should be ok"))
		}

		//通关奖励物品
		rewData := baGuaTaTemplate.GetRewData()
		rewItemMap := baGuaTaTemplate.GetRewItemMap()
		baGuaPassRewItem(pl, rewData, rewItemMap, level)
		//增加亲密度
		addPoint := baGuaTaTemplate.GetRewQinMiDu()
		addPairPoint(pl, spouseId, addPoint)
	}
	pushBaGuaResult(pl, spouseId, successful, level)
	//发送事件
	gameevent.Emit(baguaeventtypes.EventTypeBaGuaResult, pl, successful)
	return
}

//通关奖励物品
func baGuaPassRewItem(pl player.Player, rewData *propertytypes.RewData, rewItemMap map[int32]int32, level int32) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//奖励属性
	if rewData != nil {
		reasonGold := commonlog.GoldLogReasonBaGuaToKillRew
		reasonSilver := commonlog.SilverLogReasonBaGuaToKillRew
		reasonLevel := commonlog.LevelLogReasonBaGuaToKillRew
		reasonGoldText := fmt.Sprintf(reasonGold.String(), level)
		reasonSliverText := fmt.Sprintf(reasonSilver.String(), level)
		reasonlevelText := fmt.Sprintf(reasonLevel.String(), level)

		flag := propertyManager.AddRewData(rewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
		if !flag {
			panic(fmt.Errorf("bagua: ToKillFinished AddRewData  should be ok"))
		}
	}
	//奖励物品
	if len(rewItemMap) != 0 {
		flag := inventoryManager.HasEnoughSlots(rewItemMap)
		if !flag {
			//写邮件
			emailTitle := fmt.Sprintf(lang.GetLangService().ReadLang(lang.BaGuaPassRewTitle), level)
			emailContent := lang.GetLangService().ReadLang(lang.BaGuaPassRewContent)
			emaillogic.AddEmail(pl, emailTitle, emailContent, rewItemMap)
		} else {
			reasonInventory := commonlog.InventoryLogReasonBaGuaToKillRew
			reasonInventoryText := fmt.Sprintf(reasonInventory.String(), level)
			flag = inventoryManager.BatchAdd(rewItemMap, reasonInventory, reasonInventoryText)
			if !flag {
				panic(fmt.Errorf("bagua: ToKillFinished BatchAdd should be ok"))
			}
			//同步物品
			inventorylogic.SnapInventoryChanged(pl)
		}
	}
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
func pushBaGuaResult(pl player.Player, spouseId int64, sucessful bool, level int32) {
	if spouseId != 0 {
		scBaGuaPairResult := pbutil.BuildSCBaGuaPairResult(true, sucessful, level)
		pl.SendMsg(scBaGuaPairResult)
		spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
		if spl == nil {
			return
		}
		scBaGuaPairResult = pbutil.BuildSCBaGuaPairResult(false, sucessful, level)
		spl.SendMsg(scBaGuaPairResult)
	} else {
		scBaGuaToKillResult := pbutil.BuildSCBaGuaToKillResult(sucessful, level)
		pl.SendMsg(scBaGuaToKillResult)
	}
	return
}
