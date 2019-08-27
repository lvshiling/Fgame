package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/feisheng/pbutil"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	feishengtemplate "fgame/fgame/game/feisheng/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FEI_SHENG_EAT_DAN_TYPE), dispatch.HandlerFunc(handleFeiShengEatDan))
}

//处理飞升食丹
func handleFeiShengEatDan(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理飞升食丹消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSFeiShengEatDan)
	eatNum := csMsg.GetEatNum()

	err = feiShengEatDan(tpl, eatNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("feisheng:处理飞升食丹消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("feisheng:处理飞升食丹消息完成")
	return nil
}

//飞升食丹界面逻辑
func feiShengEatDan(pl player.Player, eatNum int32) (err error) {

	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	feiShengInfo := feiManager.GetFeiShengInfo()
	feiTemplate := feishengtemplate.GetFeiShengTemplateService().GetFeiShengTemplate(feiShengInfo.GetFeiLevel())
	if feiTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("feisheng:飞升食丹失败，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//物品
	needItemId := feiTemplate.ItemId
	needItemCount := feiTemplate.ItemCount * eatNum
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughItem(needItemId, needItemCount) {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
			}).Warn("feisheng:飞升食丹失败，道具不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//概率
	if feiManager.IsFullRate() {
		log.WithFields(
			log.Fields{
				"playerid":   pl.GetId(),
				"curLevel":   feiShengInfo.GetFeiLevel(),
				"curAddRate": feiShengInfo.GetAddRate(),
			}).Warn("feisheng:飞升食丹失败，成功率已满")
		playerlogic.SendSystemMessage(pl, lang.FeiShengFullRate)
		return
	}

	//消耗物品
	if needItemCount > 0 {
		useItemReason := commonlog.InventoryLogReasonFeiShenEatDan
		useItemReasonText := fmt.Sprintf(useItemReason.String(), feiShengInfo.GetFeiLevel())
		flag := inventoryManager.UseItem(needItemId, needItemCount, useItemReason, useItemReasonText)
		if !flag {
			panic(fmt.Errorf("feisheng: feisheng eat dan use item should be ok"))
		}
	}

	feiManager.EatRateDan(eatNum)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCFeiShengEatDan(eatNum, feiShengInfo.GetAddRate())
	pl.SendMsg(scMsg)
	return
}
