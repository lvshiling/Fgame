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
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	shenfalogic "fgame/fgame/game/shenfa/logic"
	"fgame/fgame/game/shenfa/pbutil"
	playershenfa "fgame/fgame/game/shenfa/player"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENFA_UPSTAR_TYPE), dispatch.HandlerFunc(handleShenFaUpstar))
}

//处理身法皮肤升星信息
func handleShenFaUpstar(s session.Session, msg interface{}) (err error) {
	log.Debug("shenfa:处理身法皮肤升星信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csShenFaUpstar := msg.(*uipb.CSShenFaUpstar)
	shenFaId := csShenFaUpstar.GetShenFaId()

	err = shenFaUpstar(tpl, shenFaId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shenFaId": shenFaId,
				"error":    err,
			}).Error("shenfa:处理身法皮肤升星信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenfa:处理身法皮肤升星完成")
	return nil
}

//身法皮肤升星的逻辑
func shenFaUpstar(pl player.Player, shenFaId int32) (err error) {
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(int(shenFaId))
	if shenfaTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"shenFaId": shenFaId,
		}).Warn("shenfa:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if shenfaTemplate.ShenfaUpstarBeginId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"shenFaId": shenFaId,
		}).Warn("shenfa:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	shenfaManager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	shenFaOtherInfo, flag := shenfaManager.IfShenFaSkinExist(shenFaId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"shenFaId": shenFaId,
		}).Warn("shenfa:未激活的身法皮肤,无法升星")
		playerlogic.SendSystemMessage(pl, lang.ShenfaSkinUpstarNoActive)
		return
	}

	_, flag = shenfaManager.IfCanUpStar(shenFaId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"shenFaId": shenFaId,
		}).Warn("shenfa:身法皮肤已满星")
		playerlogic.SendSystemMessage(pl, lang.ShenfaSkinReacheFullStar)
		return
	}

	//升星需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	level := shenFaOtherInfo.Level
	nextLevel := level + 1
	to := shenfatemplate.GetShenfaTemplateService().GetShenfa(int(shenFaId))
	if to == nil {
		return
	}
	shenFaUpstarTemplate := to.GetShenFaUpstarByLevel(nextLevel)
	if shenFaUpstarTemplate == nil {
		return
	}

	needItems := shenFaUpstarTemplate.GetNeedItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"shenFaId": shenFaId,
			}).Warn("shenfa:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗物品
	if len(needItems) != 0 {
		reasonText := commonlog.InventoryLogReasonShenFaUpstar.String()
		flag := inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonShenFaUpstar, reasonText)
		if !flag {
			panic(fmt.Errorf("shenfa: shenFaUpstar use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//身法皮肤升星判断
	pro, _, sucess := shenfalogic.ShenFaSkinUpstar(shenFaOtherInfo.UpNum, shenFaOtherInfo.UpPro, shenFaUpstarTemplate)
	flag = shenfaManager.Upstar(shenFaId, pro, sucess)
	if !flag {
		panic(fmt.Errorf("shenfa: shenfaUpstar should be ok"))
	}
	if sucess {
		shenfalogic.ShenfaPropertyChanged(pl)
	}
	scShenFaUpstar := pbutil.BuildSCShenFaUpstar(shenFaId, shenFaOtherInfo.Level, shenFaOtherInfo.UpPro)
	pl.SendMsg(scShenFaUpstar)
	return
}
