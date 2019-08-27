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
	lingtongdevlogic "fgame/fgame/game/lingtongdev/logic"
	"fgame/fgame/game/lingtongdev/pbutil"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	"fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_UNREALDAN_TYPE), dispatch.HandlerFunc(handleLingTongDevUnrealDan))

}

//处理灵童养成类食幻化丹信息
func handleLingTongDevUnrealDan(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtongdev:处理灵童养成类食幻化丹信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongDevUnrealDan := msg.(*uipb.CSLingTongDevUnrealDan)
	classType := csLingTongDevUnrealDan.GetClassType()
	num := csLingTongDevUnrealDan.GetNum()

	err = lingTongDevUnrealDan(tpl, types.LingTongDevSysType(classType), num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"num":       num,
				"error":     err,
			}).Error("lingtongdev:处理灵童养成类食幻化丹信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Debug("lingtongdev:处理灵童养成类食幻化丹信息完成")
	return nil

}

//灵童养成类食幻化丹的逻辑
func lingTongDevUnrealDan(pl player.Player, classType types.LingTongDevSysType, num int32) (err error) {
	if !classType.Vaild() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongInfo := manager.GetLingTongDevInfo(classType)
	if lingTongInfo == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:请先激活灵童养成类系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevActiveSystem, classType.String())
		return
	}
	advancedId := lingTongInfo.GetAdvancedId()
	unrealLevel := lingTongInfo.GetUnrealLevel()
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, advancedId)
	if lingTongDevTemplate == nil {
		return
	}

	lingTongDevHuanHuaTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevHuanHuaTemplate(classType, unrealLevel+1)
	if lingTongDevHuanHuaTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:幻化丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevEatUnDanReachedFull, classType.String())
		return
	}

	if unrealLevel >= lingTongDevTemplate.GetShiDanLimit() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:幻化丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevEatUnDanReachedLimit, classType.String())
		return
	}

	reachHuanHuaTemplate, flag := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevEatHuanHuanTemplate(classType, unrealLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachHuanHuaTemplate.GetLevel() > lingTongDevTemplate.GetShiDanLimit() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := lingTongDevHuanHuaTemplate.GetItemId()
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"num":       num,
			}).Warn("lingtongdev:当前幻化丹药数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonLingTongDevEatUn.String(), classType.String())
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonLingTongDevEatUn, reasonText)
		if !flag {
			panic(fmt.Errorf("lingtongdev:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	manager.EatUnrealDan(classType, reachHuanHuaTemplate.GetLevel())
	//同步属性
	lingtongdevlogic.LingTongDevPropertyChanged(pl, classType)

	scLingTongDevShiDan := pbutil.BuildSCLingTongDevUnrealDan(int32(classType), lingTongInfo.GetUnrealLevel(), lingTongInfo.GetUnrealPro())
	pl.SendMsg(scLingTongDevShiDan)
	return
}
