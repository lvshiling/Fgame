package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/tianshu/pbutil"
	playertianshu "fgame/fgame/game/tianshu/player"
	tianshutemplate "fgame/fgame/game/tianshu/template"
	tianshutypes "fgame/fgame/game/tianshu/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TIANSHU_UPLEVEL_TYPE), dispatch.HandlerFunc(handleTianShuUplevel))
}

//处理天书升级
func handleTianShuUplevel(s session.Session, msg interface{}) (err error) {
	log.Debug("tianshu:处理升级天书消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSTianShuUplevel)
	typ := csMsg.GetType()

	tianshuType := tianshutypes.TianShuType(typ)
	if !tianshuType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Warn("tianshu:处理升级天书消息,错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = tianshuUplevel(tpl, tianshuType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("tianshu:处理升级天书消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("tianshu:处理升级天书消息完成")
	return nil
}

//升级天书信息
func tianshuUplevel(pl player.Player, typ tianshutypes.TianShuType) (err error) {
	tianshuManager := pl.GetPlayerDataManager(playertypes.PlayerTianShuDataManagerType).(*playertianshu.PlayerTianShuDataManager)
	if !tianshuManager.IsActivateTianShu(typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("tianshu:处理领取天书奖励，天书未激活")
		playerlogic.SendSystemMessage(pl, lang.TianShuNotActivate)
		return
	}

	tianshuLevel := tianshuManager.GetTianShuLevel(typ)
	tianshuTemp := tianshutemplate.GetTianShuTemplateService().GetTianShuTemplate(typ, tianshuLevel)
	if tianshuTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"level":    initLevel,
			}).Warn("tianshu:处理升级天书，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	nextTemplate := tianshuTemp.GetNextTemplate()
	if nextTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"level":    initLevel,
			}).Warn("tianshu:处理升级天书，天书满级")
		playerlogic.SendSystemMessage(pl, lang.TianShuHadFullLevel)
		return
	}

	needItemMap := nextTemplate.GetNeedItemMap()
	if len(needItemMap) > 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		if !inventoryManager.HasEnoughItems(needItemMap) {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"needItemMap": needItemMap,
				}).Warn("tianshu:处理升级天书，物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		itemUseReason := commonlog.InventoryLogReasonTianShuUplevel
		reasonText := fmt.Sprintf(itemUseReason.String(), typ.String())
		flag := inventoryManager.BatchRemove(needItemMap, itemUseReason, reasonText)
		if !flag {
			panic("tianshu:批量移除物品应该成功")
		}

		inventorylogic.SnapInventoryChanged(pl)
	}

	isSuccess := mathutils.RandomHit(common.MAX_RATE, int(nextTemplate.SuccessRate))
	if isSuccess {
		tianshuManager.UplevelTianShu(typ)
	}

	scMsg := pbutil.BuildSCTianShuUplevel(isSuccess, typ, nextTemplate.Level)
	pl.SendMsg(scMsg)
	return
}
