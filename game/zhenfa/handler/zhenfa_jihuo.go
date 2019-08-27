package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	zhenfalogic "fgame/fgame/game/zhenfa/logic"
	"fgame/fgame/game/zhenfa/pbutil"
	playerzhenfa "fgame/fgame/game/zhenfa/player"
	zhenfatemplate "fgame/fgame/game/zhenfa/template"
	zhenfatypes "fgame/fgame/game/zhenfa/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ZHENFA_ACTIVATE_TYPE), dispatch.HandlerFunc(handleZhenFaActivate))
}

//处理阵法激活信息
func handleZhenFaActivate(s session.Session, msg interface{}) (err error) {
	log.Debug("zhenfa:处理阵法激活信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csZhenFaActivate := msg.(*uipb.CSZhenFaActivate)
	zhenFaType := csZhenFaActivate.GetZhenFaType()
	err = zhenFaActivate(tpl, zhenfatypes.ZhenFaType(zhenFaType))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"zhenFaType": zhenFaType,
				"error":      err,
			}).Error("zhenfa:处理阵法激活信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("zhenfa:处理阵法激活信息完成")
	return nil
}

//处理阵法激活信息逻辑
func zhenFaActivate(pl player.Player, zhenFaType zhenfatypes.ZhenFaType) (err error) {
	if !zhenFaType.Vaild() {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"zhenFaType": zhenFaType,
		}).Warn("zhenfa:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeZhenFa) {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"zhenFaType": zhenFaType,
		}).Warn("zhenfa:功能未开启")
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerZhenFaDataManagerType).(*playerzhenfa.PlayerZhenFaDataManager)
	obj := manager.GetZhenFaByType(zhenFaType)
	if obj != nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"zhenFaType": zhenFaType,
		}).Warn("zhenfa:该阵法已经激活过")
		playerlogic.SendSystemMessage(pl, lang.ZhenFaActivated)
		return
	}

	zhenFaJiHuoTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaJiHuoTemplate(zhenFaType)
	if zhenFaJiHuoTemplate == nil {
		return
	}
	if zhenFaJiHuoTemplate.NeedZhenFaType != 0 {
		needZhenFaType := zhenfatypes.ZhenFaType(zhenFaJiHuoTemplate.NeedZhenFaType)
		needObj := manager.GetZhenFaByType(needZhenFaType)
		if needObj == nil || needObj.GetLevel() < zhenFaJiHuoTemplate.NeedZhenFaLevel {
			log.WithFields(log.Fields{
				"playerId":   pl.GetId(),
				"zhenFaType": zhenFaType,
			}).Warn("zhenfa:前置条件不满足")
			needLevelStr := fmt.Sprintf("%d", zhenFaJiHuoTemplate.NeedZhenFaLevel)
			playerlogic.SendSystemMessage(pl, lang.ZhenFaActivatePreCond, zhenFaJiHuoTemplate.Name, needLevelStr)
			return
		}
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItemMap := zhenFaJiHuoTemplate.GetNeedItemMap()
	if len(needItemMap) != 0 {
		if !inventoryManager.HasEnoughItems(needItemMap) {
			log.WithFields(log.Fields{
				"playerId":   pl.GetId(),
				"zhenFaType": zhenFaType,
			}).Warn("zhenfa:物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		itemUsereason := commonlog.InventoryLogReasonZhenFaActivate
		if flag := inventoryManager.BatchRemove(needItemMap, itemUsereason, itemUsereason.String()); !flag {
			panic(fmt.Errorf("zhenfa: zhenFaActivate use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	obj, flag := manager.Activate(zhenFaType)
	if !flag {
		panic(fmt.Errorf("zhenfa: zhenFaActivate Activate should be ok"))
	}

	// 属性变化
	zhenfalogic.ZhenFaPropertyChanged(pl)

	scZhenFaActivate := pbutil.BuildSCZhenFaActivate(obj)
	pl.SendMsg(scZhenFaActivate)
	return
}
