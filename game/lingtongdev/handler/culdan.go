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

	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_CULDAN_TYPE), dispatch.HandlerFunc(handleLingTongDevCulDan))

}

//处理灵童养成类食培养丹信息
func handleLingTongDevCulDan(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtongdev:处理灵童养成类食培养丹信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongDevCulDan := msg.(*uipb.CSLingTongDevCulDan)
	classType := csLingTongDevCulDan.GetClassType()
	num := csLingTongDevCulDan.GetNum()

	err = lingTongDevCulDan(tpl, types.LingTongDevSysType(classType), num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"num":       num,
				"classType": classType,
				"error":     err,
			}).Error("lingtongdev:处理灵童养成类食培养丹信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"num":       num,
			"classType": classType,
		}).Debug("lingtongdev:处理灵童养成类食培养丹信息完成")
	return nil

}

//灵童养成类食培养丹逻辑
func lingTongDevCulDan(pl player.Player, classType types.LingTongDevSysType, num int32) (err error) {
	if !classType.Vaild() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:参数错误")
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
	culLevel := lingTongInfo.GetCulLevel()
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, advancedId)
	if lingTongDevTemplate == nil {
		return
	}
	lingTongDevPeiYangTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevPeiYangTemplate(classType, culLevel+1)
	if lingTongDevPeiYangTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:培养丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevEatCulDanReachedFull, classType.String())
		return
	}

	if culLevel >= lingTongDevTemplate.GetCulDanLimit() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:培养丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevEatCulDanReachedLimit, classType.String())
		return
	}

	reachCaoLiaoTemplate, flag := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevPeiYangUpgrade(classType, culLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachCaoLiaoTemplate.GetLevel() > lingTongDevTemplate.GetCulDanLimit() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"num":       num,
		}).Warn("lingtongdev:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := lingTongDevPeiYangTemplate.GetItemId()
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"num":       num,
			}).Warn("lingtongdev:当前培养丹数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonLingTongDevEatClu.String(), classType.String())
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonLingTongDevEatClu, reasonText)
		if !flag {
			panic(fmt.Errorf("lingtongdev:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	manager.EatCulDan(classType, reachCaoLiaoTemplate.GetLevel())
	lingtongdevlogic.LingTongDevPropertyChanged(pl, classType)

	scLingTongDevCulDan := pbutil.BuildSCLingTongDevCulDan(int32(classType), lingTongInfo.GetCulLevel(), lingTongInfo.GetCulPro())
	pl.SendMsg(scLingTongDevCulDan)
	return
}
