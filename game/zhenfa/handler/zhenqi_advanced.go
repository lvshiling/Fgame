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
	zhenfalogic "fgame/fgame/game/zhenfa/logic"
	"fgame/fgame/game/zhenfa/pbutil"
	playerzhenfa "fgame/fgame/game/zhenfa/player"
	zhenfatemplate "fgame/fgame/game/zhenfa/template"
	zhenfatypes "fgame/fgame/game/zhenfa/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ZHENQI_ADVANCED_TYPE), dispatch.HandlerFunc(handleZhenQiAdvanced))
}

//处理阵旗升阶信息
func handleZhenQiAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("zhenfa:处理阵旗升阶信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csZhenQiAdvanced := msg.(*uipb.CSZhenQiAdvanced)
	zhenFaType := csZhenQiAdvanced.GetZhenFaType()
	zhenQiPos := csZhenQiAdvanced.GetZhenQiPos()
	err = zhenQiAdvanced(tpl, zhenfatypes.ZhenFaType(zhenFaType), zhenfatypes.ZhenQiType(zhenQiPos))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"zhenFaType": zhenFaType,
				"zhenQiPos":  zhenQiPos,
				"error":      err,
			}).Error("zhenfa:处理阵旗升阶信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("zhenfa:处理阵旗升阶信息完成")
	return nil
}

//处理阵旗升阶信息逻辑
func zhenQiAdvanced(pl player.Player, zhenFaType zhenfatypes.ZhenFaType, zhenQiPos zhenfatypes.ZhenQiType) (err error) {
	if !zhenFaType.Vaild() || !zhenQiPos.Vaild() {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"zhenFaType": zhenFaType,
			"zhenQiPos":  zhenQiPos,
		}).Warn("zhenfa:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerZhenFaDataManagerType).(*playerzhenfa.PlayerZhenFaDataManager)
	zhenFaObj := manager.GetZhenFaByType(zhenFaType)
	if zhenFaObj == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"zhenFaType": zhenFaType,
			"zhenQiPos":  zhenQiPos,
		}).Warn("zhenfa:您还未激活相应的阵法,无法升阶该阵旗")
		playerlogic.SendSystemMessage(pl, lang.ZhenQiAdvancedNoActivate)
		return
	}

	obj := manager.GetZhenQi(zhenFaType, zhenQiPos)
	if obj == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"zhenFaType": zhenFaType,
			"zhenQiPos":  zhenQiPos,
		}).Warn("zhenfa:您还未激活相应的阵法,无法升阶该阵旗")
		playerlogic.SendSystemMessage(pl, lang.ZhenQiAdvancedNoActivate)
		return
	}

	zhenFaJiHuoTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaJiHuoTemplate(zhenFaType)
	if zhenFaJiHuoTemplate == nil {
		return
	}

	curNumber := obj.GetNumber()
	nextNumber := curNumber + 1
	zhenQiTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaZhenQiTemplate(zhenFaType, zhenQiPos, nextNumber)
	if zhenQiTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"zhenFaType": zhenFaType,
			"zhenQiPos":  zhenQiPos,
		}).Warn("zhenfa:您还未激活相应的阵法,无法升阶该阵旗")
		playerlogic.SendSystemMessage(pl, lang.ZhenQiAdvancedFullLevel)
		return
	}

	if zhenFaObj.GetLevel() < zhenQiTemplate.NeedLevel {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"zhenFaType": zhenFaType,
			"zhenQiPos":  zhenQiPos,
		}).Warn("zhenfa:阵法等级不足")

		needLevelStr := fmt.Sprintf("%d", zhenQiTemplate.NeedLevel)
		playerlogic.SendSystemMessage(pl, lang.ZhenFaZhenQiAdvancedNoEnoughLevel, zhenFaJiHuoTemplate.Name, needLevelStr)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItemMap := zhenQiTemplate.GetNeedItemMap()
	if len(needItemMap) != 0 {
		if !inventoryManager.HasEnoughItems(needItemMap) {
			log.WithFields(log.Fields{
				"playerId":   pl.GetId(),
				"zhenFaType": zhenFaType,
				"zhenQiPos":  zhenQiPos,
			}).Warn("zhenfa:物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		itemUsereason := commonlog.InventoryLogReasonZhenQiAdvanced
		if flag := inventoryManager.BatchRemove(needItemMap, itemUsereason, itemUsereason.String()); !flag {
			panic(fmt.Errorf("zhenfa: zhenQiAdvanced use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	pro, _, sucess := zhenfalogic.ZhenQiAdvanced(obj.GetNumberNum(), obj.GetNumberPro(), zhenQiTemplate)
	flag := manager.ZhenQiAdvanced(zhenFaType, zhenQiPos, sucess, pro)
	if !flag {
		panic(fmt.Errorf("zhenfa: ZhenQiAdvanced  should be ok"))
	}

	// 属性变化
	zhenfalogic.ZhenFaPropertyChanged(pl)
	scZhenQiAdvanced := pbutil.BuildSCZhenQiAdvanced(sucess, obj)
	pl.SendMsg(scZhenQiAdvanced)
	return
}
