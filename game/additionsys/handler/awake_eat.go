package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/additionsys/additionsys"
	additionsyseventtypes "fgame/fgame/game/additionsys/event/types"
	additionsyslogic "fgame/fgame/game/additionsys/logic"
	"fgame/fgame/game/additionsys/pbutil"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystemplate "fgame/fgame/game/additionsys/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	commontypes "fgame/fgame/game/common/types"
	gamevent "fgame/fgame/game/event"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_AWAKE_EAT_TYPE), dispatch.HandlerFunc(handleAdditionSysAwakeEat))
}

func handleAdditionSysAwakeEat(s session.Session, msg interface{}) (err error) {
	log.Debug("additionsys: 处理附加系统觉醒丹使用")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAdditionSysAwakeEat := msg.(*uipb.CSAdditionSysAwakeEat)
	sysTypeId := csAdditionSysAwakeEat.GetSysType()
	sysType := additionsystypes.AdditionSysType(sysTypeId)

	//检查参数
	if !sysType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
			}).Warn("additionsys:系统觉醒类型,错误")
		return
	}

	err = additionSysAwakeEat(tpl, sysType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
				"error":    err,
			}).Warn("additionsys:系统觉醒类型,错误")
		return
	}

	log.Debug("additionsys:处理附加系统觉醒,完成")

	return
}

func additionSysAwakeEat(pl player.Player, typ additionsystypes.AdditionSysType) (err error) {
	if !additionsyslogic.GetAdditionSysAwakeFuncOpenByType(pl, typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("additionsys:系统觉醒食用丹,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	awakeInfo := additionsysManager.GetAdditionSysAwakeInfoByType(typ)
	isAwake := awakeInfo.IsAwake

	sysAdvanced := additionsys.GetSystemAdvancedNum(pl, typ)
	awakeTemp := additionsystemplate.GetAdditionSysTemplateService().GetAwakeByType(typ)
	if awakeTemp == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ.String(),
		}).Warn("additionsys:模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	minAdvanced := awakeTemp.GetMinAwakeAdvanced()
	//判断系统阶数是否足够
	if sysAdvanced < minAdvanced {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ.String(),
		}).Warn("additionsys:系统阶数不足，无法进行觉醒")
		playerlogic.SendSystemMessage(pl, lang.AdditionSysAdvancedNoEnough)
		return
	}

	level := awakeInfo.GetAwakeLevel() + 1
	awakeLevelTemp := additionsystemplate.GetAdditionSysTemplateService().GetAwakeLevelByArg(typ, sysAdvanced, level)
	if awakeLevelTemp == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ.String(),
		}).Warn("additionsys:觉醒等级已满")
		playerlogic.SendSystemMessage(pl, lang.AdditionSysAwakeLevelTop)
		return
	}

	useItemMap := make(map[int32]int32)
	useItemMap[awakeTemp.UseItem] = awakeLevelTemp.UseItemCount

	//判断背包内物品
	if len(useItemMap) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		flag := inventoryManager.HasEnoughItems(useItemMap)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("additionsys:当前系统觉醒食用丹数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		//扣除物品
		inventoryReason := commonlog.InventoryLogReasonAdditionSysAwake
		reasonText := fmt.Sprintf(inventoryReason.String(), typ.String())
		flag = inventoryManager.BatchRemove(useItemMap, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("additionsys: UseItem should be ok"))
		}

		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	//刷新数据
	awakeInfo.SysAwakeSuccess()

	//更新属性
	additionsyslogic.UpdataAdditionSysPropertyByType(pl, typ)

	//日志
	additionsysReason := commonlog.AdditionSysLogReasonAwake
	reasonText := fmt.Sprintf(additionsysReason.String(), typ.String(), commontypes.AdvancedTypeJinJieDan.String())
	data := additionsyseventtypes.CreatePlayerAdditionSysAwakeLogEventData(typ, isAwake, additionsysReason, reasonText)
	gamevent.Emit(additionsyseventtypes.EventTypeAdditionSysAwakeLog, pl, data)

	scAwake := pbutil.BuildSCAdditionSysAwakeEat(int32(typ), level)
	pl.SendMsg(scAwake)
	return
}
