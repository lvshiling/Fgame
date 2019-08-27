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
	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONG_UPGRADE_TYPE), dispatch.HandlerFunc(handleLingTongUpgrade))
}

//处理灵童升级信息
func handleLingTongUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtong:处理灵童升级信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongUpgrade := msg.(*uipb.CSLingTongUpgrade)
	lingTongId := csLingTongUpgrade.GetLingTongId()

	err = lingTongUpgrade(tpl, lingTongId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"lingTongId": lingTongId,
				"error":      err,
			}).Error("lingtong:处理灵童升级信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Debug("lingtong:处理灵童升级完成")
	return nil
}

//灵童升级的逻辑
func lingTongUpgrade(pl player.Player, lingTongId int32) (err error) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Warn("lingtong:模板为空")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongInfo, flag := manager.GetLingTongInfo(lingTongId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Warn("lingtong:未激活该灵童")
		playerlogic.SendSystemMessage(pl, lang.LingTongNoActive)
		return
	}

	if lingTongTemplate.LingTongShengJiId == 0 {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Warn("lingtong:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	curLevel := lingTongInfo.GetLevel()
	nextLevel := curLevel + 1
	nextLingTongShengJiTemplate := lingTongTemplate.GetLingTongShengJiByLevel(nextLevel)
	if nextLingTongShengJiTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Warn("lingtong:已达满级")
		playerlogic.SendSystemMessage(pl, lang.LingTongShengJiReachFull)
		return
	}

	//升星需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItems := nextLingTongShengJiTemplate.GetItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":   pl.GetId(),
				"lingTongId": lingTongId,
			}).Warn("lingtong:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗物品
	if len(needItems) != 0 {
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonLingTongUpgrade.String(), lingTongId, curLevel)
		flag := inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonLingTongUpgrade, reasonText)
		if !flag {
			panic(fmt.Errorf("lingtong: lingTongUpgrade BatchRemove item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//灵童升级判断
	pro, _, sucess := lingtonglogic.LingTongShengJi(lingTongInfo.GetNum(), lingTongInfo.GetPro(), nextLingTongShengJiTemplate)
	flag = manager.ShengJi(lingTongId, pro, sucess)
	if !flag {
		panic(fmt.Errorf("lingtong: lingTongUpgrade should be ok"))
	}
	if sucess {
		manager.AddLevel()
		lingtonglogic.LingTongPropertyChanged(pl)
	}
	scLingTongUpgrade := pbutil.BuildSCLingTongUpgrade(lingTongId, lingTongInfo.GetLevel(), lingTongInfo.GetPro(), sucess)
	pl.SendMsg(scLingTongUpgrade)
	return
}
