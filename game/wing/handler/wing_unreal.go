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

	processor.Register(codec.MessageType(uipb.MessageType_CS_WING_UNREAL_TYPE), dispatch.HandlerFunc(handleWingUnreal))

}

//处理战翼幻化信息
func handleWingUnreal(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理战翼幻化信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csWingUnreal := msg.(*uipb.CSWingUnreal)
	wingId := csWingUnreal.GetWingId()
	err = wingUnreal(tpl, int(wingId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"wingId":   wingId,
				"error":    err,
			}).Error("wing:处理战翼幻化信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"wingId":   wingId,
		}).Debug("wing:处理战翼幻化信息完成")
	return nil

}

//战翼幻化的逻辑
func wingUnreal(pl player.Player, wingId int) (err error) {
	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingTemplate := wing.GetWingService().GetWing(int(wingId))
	//校验参数
	if wingTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"wingId":   wingId,
		}).Warn("Wing:幻化advancedId无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	advancedId := wingManager.GetWingAdvancedId()
	if advancedId <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"wingId":   wingId,
		}).Warn("wing:请先激活战翼系统")
		playerlogic.SendSystemMessage(pl, lang.WingUnrealActiveSystem)
		return
	}

	//是否已幻化
	flag := wingManager.IsUnrealed(wingId)
	if !flag {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//获取物品幻化条件3,消耗物品数
		paramItemsMap := wingTemplate.GetMagicParamIMap()
		if len(paramItemsMap) != 0 {
			flag := inventoryManager.HasEnoughItems(paramItemsMap)
			if !flag {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"wingId":   wingId,
				}).Warn("Wing:还有幻化条件未达成，无法解锁幻化")
				playerlogic.SendSystemMessage(pl, lang.WingUnrealCondNotReached)
				return
			}
		}

		flag = wingManager.IsCanUnreal(wingId)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"wingId":   wingId,
			}).Warn("Wing:还有幻化条件未达成，无法解锁幻化")
			playerlogic.SendSystemMessage(pl, lang.WingUnrealCondNotReached)
			return
		}

		//使用幻化条件3的物品
		inventoryReason := commonlog.InventoryLogReasonWingUnreal
		reasonText := fmt.Sprintf(inventoryReason.String(), wingId)
		if len(paramItemsMap) != 0 {
			flag := inventoryManager.BatchRemove(paramItemsMap, inventoryReason, reasonText)
			if !flag {
				panic(fmt.Errorf("wing:use item should be ok"))
			}
			//同步物品
			inventorylogic.SnapInventoryChanged(pl)
		}
		wingManager.AddUnrealInfo(wingId)
		winglogic.WingPropertyChanged(pl)
	}

	flag = wingManager.Unreal(wingId)
	if !flag {
		panic(fmt.Errorf("wing:幻化应该成功"))
	}
	scWingUnreal := pbutil.BuildSCWingUnreal(int32(wingId))
	pl.SendMsg(scWingUnreal)
	return
}
