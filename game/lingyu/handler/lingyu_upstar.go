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
	lingyulogic "fgame/fgame/game/lingyu/logic"
	"fgame/fgame/game/lingyu/pbutil"
	playerlingyu "fgame/fgame/game/lingyu/player"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGYU_UPSTAR_TYPE), dispatch.HandlerFunc(handleLingYuUpstar))
}

//处理领域皮肤升星信息
func handleLingYuUpstar(s session.Session, msg interface{}) (err error) {
	log.Debug("lingyu:处理领域皮肤升星信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingYuUpstar := msg.(*uipb.CSLingYuUpstar)
	lingYuId := csLingYuUpstar.GetLingYuId()

	err = lingYuUpstar(tpl, lingYuId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"lingYuId": lingYuId,
				"error":    err,
			}).Error("lingyu:处理领域皮肤升星信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lingyu:处理领域皮肤升星完成")
	return nil
}

//领域皮肤升星的逻辑
func lingYuUpstar(pl player.Player, lingYuId int32) (err error) {
	lingYuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyu(int(lingYuId))
	if lingYuTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"lingYuId": lingYuId,
		}).Warn("lingyu:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if lingYuTemplate.FieldUpstarBeginId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"lingYuId": lingYuId,
		}).Warn("lingyu:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	lingyuManager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingYuOtherInfo, flag := lingyuManager.IfLingYuSkinExist(lingYuId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"lingYuId": lingYuId,
		}).Warn("lingyu:未激活的领域皮肤,无法升星")
		playerlogic.SendSystemMessage(pl, lang.LingyuSkinUpstarNoActive)
		return
	}

	_, flag = lingyuManager.IfCanUpStar(lingYuId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"lingYuId": lingYuId,
		}).Warn("lingyu:领域皮肤已满星")
		playerlogic.SendSystemMessage(pl, lang.LingyuSkinReacheFullStar)
		return
	}

	//升星需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	level := lingYuOtherInfo.Level
	nextLevel := level + 1
	to := lingyutemplate.GetLingyuTemplateService().GetLingyu(int(lingYuId))
	if to == nil {
		return
	}
	lingYuUpstarTemplate := to.GetLingYuUpstarByLevel(nextLevel)
	if lingYuUpstarTemplate == nil {
		return
	}

	needItems := lingYuUpstarTemplate.GetNeedItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"lingYuId": lingYuId,
			}).Warn("lingyu:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗物品
	if len(needItems) != 0 {
		reasonText := commonlog.InventoryLogReasonLingYuUpstar.String()
		flag := inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonLingYuUpstar, reasonText)
		if !flag {
			panic(fmt.Errorf("lingyu: lingyuUpstar use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//领域皮肤升星判断
	pro, _, sucess := lingyulogic.LingYuSkinUpstar(lingYuOtherInfo.UpNum, lingYuOtherInfo.UpPro, lingYuUpstarTemplate)
	flag = lingyuManager.Upstar(lingYuId, pro, sucess)
	if !flag {
		panic(fmt.Errorf("lingyu: mountUpstar should be ok"))
	}
	if sucess {
		lingyulogic.LingyuPropertyChanged(pl)
	}
	scMountUpstar := pbutil.BuildSCLingYuUpstar(lingYuId, lingYuOtherInfo.Level, lingYuOtherInfo.UpPro)
	pl.SendMsg(scMountUpstar)
	return
}
