package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	additionsyslogic "fgame/fgame/game/additionsys/logic"
	"fgame/fgame/game/additionsys/pbutil"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystemplate "fgame/fgame/game/additionsys/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_HUALING_EAT_TYPE), dispatch.HandlerFunc(handleAdditionSysHuaLingEat))
}

//处理附加系统化灵食用丹
func handleAdditionSysHuaLingEat(s session.Session, msg interface{}) (err error) {
	log.Debug("additionsys:处理附加系统化灵食用丹")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAdditionSysHualingEat := msg.(*uipb.CSAdditionSysHualingEat)
	num := csAdditionSysHualingEat.GetNum()
	sysTypeId := csAdditionSysHualingEat.GetSysType()
	sysType := additionsystypes.AdditionSysType(sysTypeId)

	//参数不对
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("additionsys:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	if !sysType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
			}).Warn("additionsys:系统化灵类型,错误")
		return
	}

	err = additionSysHuaLingEat(tpl, sysType, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
				"num":      num,
				"error":    err,
			}).Warn("additionsys:系统化灵类型,错误")
		return
	}

	log.Debug("additionsys:处理附加系统化灵,完成")
	return nil

}

//系统化灵食用丹逻辑
func additionSysHuaLingEat(pl player.Player, typ additionsystypes.AdditionSysType, num int32) (err error) {
	if !additionsyslogic.GetAdditionSysHuaLingFuncOpenByType(pl, typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("additionsyse:系统化灵食用丹,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	levelInfo := additionsysManager.GetAdditionSysLevelInfoByType(typ)
	culLevel := levelInfo.LingLevel

	//判断是否可以升
	nextHuaLingTemplate := levelInfo.GetNextHuaLingTemplate()
	if nextHuaLingTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"num":      num,
			}).Warn("additionsyse:系统化灵食用丹,已经满级")
		playerlogic.SendSystemMessage(pl, lang.AdditionSysLevelHighest)
		return
	}

	reachTemplate, flag := additionsystemplate.GetAdditionSysTemplateService().GetHuaLingReachByArg(typ, culLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ.String(),
			"num":      num,
		}).Warn("additionsyse:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	_, huaLingUseTemplate := additionsystemplate.GetAdditionSysTemplateService().GetHuaLingByArg(typ, nextHuaLingTemplate.Level)
	useItem := huaLingUseTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"num":      num,
			}).Warn("additionsyse:当前系统化灵食用丹数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		inventoryReason := commonlog.InventoryLogReasonAdditionSysHuaLing
		reasonText := fmt.Sprintf(inventoryReason.String(), typ.String(), levelInfo.LingLevel)
		flag = inventoryManager.UseItem(useItem, num, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("additionsyse:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	levelInfo.HuaLingUpgrade(reachTemplate.Level)
	//更新属性
	additionsyslogic.UpdataAdditionSysPropertyByType(pl, typ)

	scAdditionSysHuaLingEat := pbutil.BuildSCAdditionSysHuaLingEat(int32(levelInfo.SysType), levelInfo.LingLevel, levelInfo.LingPro)
	pl.SendMsg(scAdditionSysHuaLingEat)
	return
}
