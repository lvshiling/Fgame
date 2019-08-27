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
	xiantilogic "fgame/fgame/game/xianti/logic"
	"fgame/fgame/game/xianti/pbutil"
	playerxianti "fgame/fgame/game/xianti/player"
	"fgame/fgame/game/xianti/xianti"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANTI_UNREAL_TYPE), dispatch.HandlerFunc(handleXianTiUnreal))
}

//处理仙体幻化信息
func handleXianTiUnreal(s session.Session, msg interface{}) (err error) {
	log.Debug("xianti:处理仙体幻化信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csXianTiUnreal := msg.(*uipb.CSXiantiUnreal)
	xianTiId := csXianTiUnreal.GetXiantiId()
	if xianTiId <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"xianTiId": xianTiId,
			}).Debug("xianti:处理仙体幻化信息,无效xianTiId")
		return nil
	}
	err = xianTiUnreal(tpl, int(xianTiId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"xianTiId": xianTiId,
				"error":    err,
			}).Error("xianti:处理仙体幻化信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"xianTiId": xianTiId,
		}).Debug("xianti:处理仙体幻化信息完成")
	return nil

}

//仙体幻化的逻辑
func xianTiUnreal(pl player.Player, xianTiId int) (err error) {
	xianTiManager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	xianTiTemplate := xianti.GetXianTiService().GetXianTi(int(xianTiId))
	//校验参数
	if xianTiTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"xianTiId": xianTiId,
		}).Warn("xianti:幻化xianTiId无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//是否已幻化
	flag := xianTiManager.IsUnrealed(xianTiId)
	if !flag {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//获取物品幻化条件3,消耗物品数
		paramItemsMap := xianTiTemplate.GetMagicParamIMap()
		if len(paramItemsMap) != 0 {
			flag := inventoryManager.HasEnoughItems(paramItemsMap)
			if !flag {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"xianTiId": xianTiId,
				}).Warn("xianti:还有幻化条件未达成，无法解锁幻化")
				playerlogic.SendSystemMessage(pl, lang.XianTiUnrealCondNotReached)
				return
			}

		}

		flag = xianTiManager.IsCanUnreal(xianTiId)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"xianTiId": xianTiId,
			}).Warn("xianti:还有条件未达成，无法解锁幻化")
			playerlogic.SendSystemMessage(pl, lang.XianTiUnrealCondNotReached)
			return
		}

		//使用幻化条件3的物品
		inventoryReason := commonlog.InventoryLogReasonXianTiUnreal
		reasonText := fmt.Sprintf(inventoryReason.String(), xianTiId)
		if len(paramItemsMap) != 0 {
			flag := inventoryManager.BatchRemove(paramItemsMap, inventoryReason, reasonText)
			if !flag {
				panic(fmt.Errorf("xianti:use item should be ok"))
			}
			//同步物品
			inventorylogic.SnapInventoryChanged(pl)
		}
		xianTiManager.AddUnrealInfo(xianTiId)
		xiantilogic.XianTiPropertyChanged(pl)
	}
	flag = xianTiManager.Unreal(xianTiId)
	if !flag {
		panic(fmt.Errorf("xianti:幻化应该成功"))
	}
	scXianTiUnreal := pbutil.BuildSCXianTiUnreal(int32(xianTiId))
	pl.SendMsg(scXianTiUnreal)
	return
}
