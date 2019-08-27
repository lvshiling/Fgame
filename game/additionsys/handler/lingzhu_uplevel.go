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
	commonlogic "fgame/fgame/game/common/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_LINGTONG_LINGZHU_UPLEVEL_TYPE), dispatch.HandlerFunc(handleAdditionSysLingZhuUplevel))
}

func handleAdditionSysLingZhuUplevel(s session.Session, msg interface{}) (err error) {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSAdditionSysLingTongLingZhuUplevel)
	lingtongId := csMsg.GetLingtongId()
	sysType, _ := additionsystypes.ConvertLingTongIdToAdditionSysType(int(lingtongId))
	lingZhuType := additionsystypes.LingZhuType(csMsg.GetLingZhuType())
	//参数不对
	if !sysType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType,
			}).Warn("additionsys:灵珠升级请求类型,错误")
		return
	}
	if !lingZhuType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"lingZhuType": lingZhuType,
			}).Warn("additionsys:灵珠升级请求类型,错误")
		return
	}

	err = additionSysLingZhuUplevel(tpl, sysType, lingZhuType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"sysType":     sysType.String(),
				"lingZhuType": lingZhuType.String(),
				"error":       err,
			}).Warn("additionsys:灵珠升级请求类型,错误")
		return
	}

	return
}

func additionSysLingZhuUplevel(pl player.Player, typ additionsystypes.AdditionSysType, lingZhuType additionsystypes.LingZhuType) (err error) {
	if !additionsyslogic.GetAdditionSysLingZhuFuncOpenByType(pl, typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("inventory:升级失败,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	lingZhuObj := additionsysManager.GetAdditionSysLingZhu(typ, lingZhuType)
	curLevel := int32(0)
	if lingZhuObj != nil {
		curLevel = lingZhuObj.GetLevel()
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	// propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	lingzhuTemplate := additionsystemplate.GetAdditionSysTemplateService().GetLingZhuTemplate(lingZhuType)
	needItemId := lingzhuTemplate.UseItemId
	lingzhuNextLevelTemplate := lingzhuTemplate.GetLevelTemplate(curLevel + 1)
	if lingzhuNextLevelTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"typ":         typ,
				"lingZhuType": lingZhuType,
				"curLevel":    curLevel,
			}).Warn("additionsys:灵珠升级请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	needItemNum := lingzhuNextLevelTemplate.UseItemCount
	curNeedItemNum := inventoryManager.NumOfItems(needItemId)
	if curNeedItemNum < needItemNum {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"typ":         typ,
				"lingZhuType": lingZhuType,
				"curLevel":    curLevel,
				"needItemId":  needItemId,
				"needItemNum": needItemNum,
			}).Warn("additionsys:灵珠升级请求，物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//升级
	updateRate := lingzhuNextLevelTemplate.UpdateWfb
	blessMax := lingzhuNextLevelTemplate.ZhufuMax
	addMin := lingzhuNextLevelTemplate.AddMin
	addMax := lingzhuNextLevelTemplate.AddMax + 1
	randBless := int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes := int32(1)
	curTimesNum := int32(0)
	curBless := int32(0)
	if lingZhuObj != nil {
		curTimesNum = lingZhuObj.GetTimes()
		curBless = int32(lingZhuObj.GetBless())
	}
	curTimesNum += addTimes
	pro, sucess := commonlogic.AdvancedStatusAndProgress(curTimesNum, curBless, lingzhuNextLevelTemplate.TimesMin, lingzhuNextLevelTemplate.TimesMax, randBless, updateRate, blessMax)

	additionsysManager.LingZhuUpLevel(typ, lingZhuType, sucess, int64(pro))

	additionsyslogic.LingTongLingZhuActiveSkill(pl, typ)

	//同步物品（删掉吃掉的物品）
	useReason := commonlog.InventoryLogReasonAdditionSysLingZhuUplevel
	useReasonText := fmt.Sprintf(useReason.String(), typ.String(), lingZhuType.String())
	if needItemNum > 0 {
		flag := inventoryManager.UseItem(needItemId, needItemNum, useReason, useReasonText)
		if !flag {
			panic("inventory:移除物品应该是可以的")
		}
	}

	additionsyslogic.UpdataAdditionSysPropertyByType(pl, typ)
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	id, _ := typ.ConvertAdditionSysTypeToLingTongId()
	if lingZhuObj != nil {
		scMsg := pbutil.BuildSCAdditionSysLingZhuUplevel(int32(id), lingZhuObj)
		pl.SendMsg(scMsg)
	} else {
		lingZhuObj = additionsysManager.GetAdditionSysLingZhu(typ, lingZhuType)
		scMsg := pbutil.BuildSCAdditionSysLingZhuUplevel(int32(id), lingZhuObj)
		pl.SendMsg(scMsg)
	}

	return
}
