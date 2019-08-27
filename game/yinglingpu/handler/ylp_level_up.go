package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	ylplogic "fgame/fgame/game/yinglingpu/logic"
	ylppbutil "fgame/fgame/game/yinglingpu/pbutil"
	ylpplayer "fgame/fgame/game/yinglingpu/player"
	ylptmp "fgame/fgame/game/yinglingpu/template"
	ylptypes "fgame/fgame/game/yinglingpu/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_YLPU_SP_YLP_UP_TYPE), dispatch.HandlerFunc(handleYlpLevelup))
}

func handleYlpLevelup(s session.Session, msg interface{}) (err error) {
	log.Debug("yinglingpu:英灵普升级开始....")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csXiangQian := msg.(*uipb.CSYingLingPuUpLevel)
	tujianId := csXiangQian.GetTuJianId()
	tujianTypeId := csXiangQian.GetTuJianType()
	tujianType := ylptypes.YingLingPuTuJianType(tujianTypeId)

	err = levelUp(tpl, tujianId, tujianType)
	if err != nil {
		log.WithFields(log.Fields{
			"playerId": tpl.GetId(),
			"error":    err,
		}).Error("yinglingpu:英灵普升级失败")
		//需要返回
		return
	}
	log.Debug("yinglingpu:英灵普升级结束....")
	return
}

func levelUp(pl player.Player, tujianId int32, tujianType ylptypes.YingLingPuTuJianType) (err error) {
	if !tujianType.Valid() {
		//需要用warn
		log.WithFields(log.Fields{
			"tujianId":   tujianId,
			"tujianType": tujianType,
			"playerId":   pl.GetId(),
		}).Warn("yinglingpu:升级图鉴失败，图鉴类型错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	templateService := ylptmp.GetYingLingPuTemplateService()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	ylpManager := pl.GetPlayerDataManager(playertypes.PlayerYingLingPuManagerType).(*ylpplayer.PlayerYingLingPuManager)

	if !ylpManager.ExistsYlp(tujianId, tujianType) {
		log.WithFields(log.Fields{
			"tujianId":   tujianId,
			"tujianType": tujianType,
			"playerId":   pl.GetId(),
		}).Warn("yinglingpu:升级图鉴失败，已经存在")
		playerlogic.SendSystemMessage(pl, lang.YingLingPuNotExists)
		return
	}
	ylpTemplate := templateService.GetYingLingPuById(tujianId, tujianType)
	if ylpTemplate == nil {
		log.WithFields(log.Fields{
			"tujianId":   tujianId,
			"tujianType": tujianType,
			"playerId":   pl.GetId(),
		}).Warn("yinglingpu:升级图鉴失败，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.YingLingPuNotExistsTemplate)
		return
	}

	ylpInfo := ylpManager.GetYlpInfo(tujianId, tujianType)
	nextLevel := ylpInfo.Level + 1
	if nextLevel > ylpTemplate.GetLevelSuan().LevelMax {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"level":      nextLevel,
			"tujianId":   tujianId,
			"tujianType": tujianType,
		}).Warn("yinglingpu:升级图鉴失败，等级已经最高")
		playerlogic.SendSystemMessage(pl, lang.YingLingPuMaxLevel)
		return

	}
	// nextLevelTemplateInfo := templateService.GetYingLingPuLevel(tujianId, nextLevel, tujianType)
	// if nextLevelTemplateInfo == nil { //模板没有这个等级的
	// 	log.WithFields(log.Fields{
	// 		"playerId":   pl.GetId(),
	// 		"level":      nextLevel,
	// 		"tujianId":   tujianId,
	// 		"tujianType": tujianType,
	// 	}).Warn("yinglingpu:升级图鉴失败，等级已经最高")
	// 	playerlogic.SendSystemMessage(pl, lang.YingLingPuMaxLevel)
	// 	return
	// }
	useLevelTempCount := ylpTemplate.GetConsumeMap(nextLevel)

	if len(useLevelTempCount) > 0 {
		//判断背包里面是不是有足够的东西
		// useLevelTempCount := nextLevelTemplateInfo.GetUseItemMap()
		if !inventoryManager.HasEnoughItems(useLevelTempCount) { //不够
			log.WithFields(log.Fields{
				"playerId":   pl.GetId(),
				"level":      nextLevel,
				"tujianId":   tujianId,
				"tujianType": tujianType,
			}).Warn("yinglingpu:升级图鉴失败，升级物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonYingLingPuLevelUp.String(), tujianType.String(), tujianId, nextLevel)
		//扣掉物品
		flag := inventoryManager.BatchRemove(useLevelTempCount, commonlog.InventoryLogReasonYingLingPuLevelUp, reasonText)
		if !flag {
			panic(fmt.Errorf("yinglingpu:英灵谱升级，扣除物品失败"))
		}
	}

	//这下就都满足了
	flag := ylpManager.UpYingLingPu(tujianId, tujianType) //升级
	if !flag {
		panic(fmt.Errorf("yinglingpu:升级图鉴[%d],[%s]应该成功", tujianId, tujianType.String()))
	}

	inventorylogic.SnapInventoryChanged(pl)
	ylplogic.YingLingPuPropertyChanged(pl)
	responInfo := ylppbutil.BuildYlpLevel(tujianId, int32(tujianType), nextLevel)
	pl.SendMsg(responInfo)

	return
}
