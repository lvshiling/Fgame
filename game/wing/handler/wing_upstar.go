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
	winglogic "fgame/fgame/game/wing/logic"
	"fgame/fgame/game/wing/pbutil"
	playerwing "fgame/fgame/game/wing/player"
	"fgame/fgame/game/wing/wing"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WING_UPSTAR_TYPE), dispatch.HandlerFunc(handleWingUpstar))
}

//处理战翼皮肤升星信息
func handleWingUpstar(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理战翼皮肤升星信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csWingUpstar := msg.(*uipb.CSWingUpstar)
	wingId := csWingUpstar.GetWingId()

	err = wingUpstar(tpl, wingId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"wingId":   wingId,
				"error":    err,
			}).Error("wing:处理战翼皮肤升星信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("wing:处理战翼皮肤升星完成")
	return nil
}

//战翼皮肤升星的逻辑
func wingUpstar(pl player.Player, wingId int32) (err error) {
	wingTemplate := wing.GetWingService().GetWing(int(wingId))
	if wingTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"wingId":   wingId,
		}).Warn("wing:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if wingTemplate.WingUpstarBeginId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"wingId":   wingId,
		}).Warn("wing:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingOtherInfo, flag := wingManager.IfWingSkinExist(wingId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"wingId":   wingId,
		}).Warn("wing:未激活的战翼皮肤,无法升星")
		playerlogic.SendSystemMessage(pl, lang.WingSkinUpstarNoActive)
		return
	}

	_, flag = wingManager.IfCanUpStar(wingId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"wingId":   wingId,
		}).Warn("wing:战翼皮肤已满星")
		playerlogic.SendSystemMessage(pl, lang.WingSkinReacheFullStar)
		return
	}

	//升星需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	level := wingOtherInfo.Level
	nextLevel := level + 1
	to := wing.GetWingService().GetWing(int(wingId))
	if to == nil {
		return
	}
	wingUpstarTemplate := to.GetWingUpstarByLevel(nextLevel)
	if wingUpstarTemplate == nil {
		return
	}

	needItems := wingUpstarTemplate.GetNeedItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"wingId":   wingId,
			}).Warn("wing:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗物品
	if len(needItems) != 0 {
		reasonText := commonlog.InventoryLogReasonWingUpstar.String()
		flag := inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonWingUpstar, reasonText)
		if !flag {
			panic(fmt.Errorf("wing: wingUpstar use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//皮肤升星判断
	pro, _, sucess := winglogic.WingSkinUpstar(wingOtherInfo.UpNum, wingOtherInfo.UpPro, wingUpstarTemplate)
	flag = wingManager.Upstar(wingId, pro, sucess)
	if !flag {
		panic(fmt.Errorf("wing: wingUpstar should be ok"))
	}
	if sucess {
		winglogic.WingPropertyChanged(pl)
	}
	scWingUpstar := pbutil.BuildSCWingUpstar(wingId, wingOtherInfo.Level, wingOtherInfo.UpPro)
	pl.SendMsg(scWingUpstar)
	return
}
