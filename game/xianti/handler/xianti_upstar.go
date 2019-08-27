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
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANTI_UPSTAR_TYPE), dispatch.HandlerFunc(handleXianTiUpstar))
}

//处理仙体皮肤升星信息
func handleXianTiUpstar(s session.Session, msg interface{}) (err error) {
	log.Debug("xianti:处理仙体皮肤升星信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csXianTiUpstar := msg.(*uipb.CSXiantiUpstar)
	xianTiId := csXianTiUpstar.GetXiantiId()

	err = xianTiUpstar(tpl, xianTiId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"xianTiId": xianTiId,
				"error":    err,
			}).Error("xianti:处理仙体皮肤升星信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("xianti:处理仙体皮肤升星完成")
	return nil
}

//仙体皮肤升星的逻辑
func xianTiUpstar(pl player.Player, xianTiId int32) (err error) {
	xianTiTemplate := xianti.GetXianTiService().GetXianTi(int(xianTiId))
	if xianTiTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"xianTiId": xianTiId,
		}).Warn("xianti:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if xianTiTemplate.XianTiUpstarBeginId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"xianTiId": xianTiId,
		}).Warn("xianti:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	xianTiManager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	xianTiOtherInfo, flag := xianTiManager.IfXianTiSkinExist(xianTiId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"xianTiId": xianTiId,
		}).Warn("xianti:未激活的仙体皮肤,无法升星")
		playerlogic.SendSystemMessage(pl, lang.XianTiSkinUpstarNoActive)
		return
	}

	_, flag = xianTiManager.IfCanUpStar(xianTiId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"xianTiId": xianTiId,
		}).Warn("xianti:仙体皮肤已满星")
		playerlogic.SendSystemMessage(pl, lang.XianTiSkinReacheFullStar)
		return
	}

	//升星需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	level := xianTiOtherInfo.Level
	nextLevel := level + 1
	to := xianti.GetXianTiService().GetXianTi(int(xianTiId))
	if to == nil {
		return
	}
	xianTiUpstarTemplate := to.GetXianTiUpstarByLevel(nextLevel)
	if xianTiUpstarTemplate == nil {
		return
	}

	needItems := xianTiUpstarTemplate.GetNeedItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"xianTiId": xianTiId,
			}).Warn("xianti:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗物品
	if len(needItems) != 0 {
		reasonText := commonlog.InventoryLogReasonXianTiUpstar.String()
		flag := inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonXianTiUpstar, reasonText)
		if !flag {
			panic(fmt.Errorf("xianti: xianTiUpstar use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//仙体皮肤升星判断
	pro, _, sucess := xiantilogic.XianTiSkinUpstar(xianTiOtherInfo.UpNum, xianTiOtherInfo.UpPro, xianTiUpstarTemplate)
	flag = xianTiManager.Upstar(xianTiId, pro, sucess)
	if !flag {
		panic(fmt.Errorf("xianti: xianTiUpstar should be ok"))
	}
	if sucess {
		xiantilogic.XianTiPropertyChanged(pl)
	}
	scXianTiUpstar := pbutil.BuildSCXianTiUpstar(xianTiId, xianTiOtherInfo.Level, xianTiOtherInfo.UpPro)
	pl.SendMsg(scXianTiUpstar)
	return
}
