package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	pbutil "fgame/fgame/game/xiantao/pbutil"
	playerxiantao "fgame/fgame/game/xiantao/player"
	xiantaotemplate "fgame/fgame/game/xiantao/template"
	xiantaotypes "fgame/fgame/game/xiantao/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANTAO_PEACH_COMMIT_TYPE), dispatch.HandlerFunc(handleXianTaoCommit))
}

//处理仙桃提交
func handleXianTaoCommit(s session.Session, msg interface{}) (err error) {
	log.Debug("xiantao:处理仙桃提交消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = xianTaoCommit(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("xiantao:处理仙桃提交消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("xiantao:处理仙桃提交消息完成")
	return nil
}

//处理仙桃提交信息逻辑
func xianTaoCommit(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xiantao:处理玩家采集仙桃,场景为空")
		return
	}

	mapType := s.MapTemplate().GetMapType()
	if mapType != scenetypes.SceneTypeXianTaoDaHui {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xiantao:处理玩家采集仙桃,不是仙桃大会场景")
		return
	}

	xiantaotemplateService := xiantaotemplate.GetXianTaoTemplateService()
	xiantaoConst := xiantaotemplateService.GetXianTaoConstTemplate()

	taoXianNPC := scenetemplate.GetSceneTemplateService().GetQuestNPC(int32(xiantaoConst.GetTaoXianTemp().TemplateId()))
	distanceOk := false
	if taoXianNPC != nil {
		distance := coreutils.Distance(taoXianNPC.GetPos(), pl.GetPos())
		commitDistance := float64(xiantaoConst.TiJiaoService) / float64(1000)
		if distance <= float64(commitDistance) {
			distanceOk = true
		}
	}

	if !distanceOk {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("xiantao:处理玩家采集仙桃,不在提交范围内")
		playerlogic.SendSystemMessage(pl, lang.XianTaoCommitNoDistance)
		return
	}

	xianTaoManager := pl.GetPlayerDataManager(types.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
	xianTaoObj := xianTaoManager.GetXianTaoObject()
	if xianTaoObj.HighPeachCount <= 0 && xianTaoObj.JuniorPeachCount <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xiantao:处理玩家采集仙桃,仙桃数量不足")
		playerlogic.SendSystemMessage(pl, lang.XianTaoPeachNumNoEnough)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	rewSilver := int32(0)
	rewBindGold := int32(0)
	rewGold := int32(0)
	rewExp := int32(0)
	rewExpPoint := int32(0)
	totalItemMap := make(map[int32]int32)

	qXianTaoTemp := xiantaotemplateService.GetXianTaoTempByArg(xiantaotypes.XianTaoTypeQianNian, xianTaoObj.HighPeachCount)
	if qXianTaoTemp != nil {
		rewSilver += qXianTaoTemp.RewSilver
		rewBindGold += qXianTaoTemp.RewBindGold
		rewGold += qXianTaoTemp.RewGold
		rewExp += qXianTaoTemp.RewExp
		rewExpPoint += qXianTaoTemp.RewExpPoint
		for itemId, num := range qXianTaoTemp.GetRewItemMap() {
			_, ok := totalItemMap[itemId]
			if ok {
				totalItemMap[itemId] += num
			} else {
				totalItemMap[itemId] = num
			}
		}
	}
	bXianTaoTemp := xiantaotemplateService.GetXianTaoTempByArg(xiantaotypes.XianTaoTypeBaiNian, xianTaoObj.JuniorPeachCount)
	if bXianTaoTemp != nil {
		rewSilver += bXianTaoTemp.RewSilver
		rewBindGold += bXianTaoTemp.RewBindGold
		rewGold += bXianTaoTemp.RewGold
		rewExp += bXianTaoTemp.RewExp
		rewExpPoint += bXianTaoTemp.RewExpPoint
		for itemId, num := range bXianTaoTemp.GetRewItemMap() {
			_, ok := totalItemMap[itemId]
			if ok {
				totalItemMap[itemId] += num
			} else {
				totalItemMap[itemId] = num
			}
		}
	}

	//背包空间
	if !inventoryManager.HasEnoughSlots(totalItemMap) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xiantao:处理玩家采集仙桃，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//清空仙桃数量
	xianTaoManager.ClearAllPeachCount()

	reasonGold := commonlog.GoldLogReasonXianTaoCommit
	reasonSilver := commonlog.SilverLogReasonXianTaoCommit
	reasonLevel := commonlog.LevelLogReasonXianTaoCommit
	reasonGoldText := reasonGold.String()
	reasonSilverText := reasonSilver.String()
	reasonLevelText := reasonLevel.String()

	totalRewData := propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)
	flag := propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSilverText, reasonLevel, reasonLevelText)
	if !flag {
		panic("xiantao:xiantao rewards add RewData should be ok")
	}

	//增加物品
	itemGetReason := commonlog.InventoryLogReasonXianTaoCommit
	itemGetReasonText := itemGetReason.String()
	flag = inventoryManager.BatchAdd(totalItemMap, itemGetReason, itemGetReasonText)
	if !flag {
		panic("xiantao:xiantao rewards add item should be ok")
	}
	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCXiantaoPeachCommit(totalRewData, totalItemMap)
	pl.SendMsg(scMsg)
	return
}
