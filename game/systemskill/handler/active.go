package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/systemskill/pbutil"
	playersysskill "fgame/fgame/game/systemskill/player"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltemplate "fgame/fgame/game/systemskill/template"
	systemskilltypes "fgame/fgame/game/systemskill/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SYSTEM_SKILL_ACTIVE_TYPE), dispatch.HandlerFunc(handleSystemActive))
}

//处理系统技能激活信息
func handleSystemActive(s session.Session, msg interface{}) (err error) {
	log.Debug("systemskill:处理系统技能激活信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSystemSkillActive := msg.(*uipb.CSSystemSkillActive)
	typ := csSystemSkillActive.GetTag()
	subType := csSystemSkillActive.GetSubType()

	err = systemActive(tpl, systemskilltypes.SystemSkillType(typ), systemskilltypes.SystemSkillSubType(subType))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
				"error":    err,
			}).Error("systemskill:处理系统技能激活信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"subType":  subType,
			"typ":      typ,
		}).Debug("systemskill:处理系统技能激活信息完成")
	return nil
}

//处理系统技能激活信息逻辑
func systemActive(pl player.Player, typ systemskilltypes.SystemSkillType, subType systemskilltypes.SystemSkillSubType) (err error) {
	if !typ.Valid() || !subType.Valid() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
			"subType":  subType,
		}).Warn("systemskill:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerSystemSkillDataManagerType).(*playersysskill.PlayerSystemSkillDataManager)
	flag := manager.IfSystemSkillExist(typ, subType)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
			"subType":  subType,
		}).Warn("systemskill:该系统技能已激活,无需激活")
		playerlogic.SendSystemMessage(pl, lang.SystemSkillRepeatActive)
		return
	}

	skTemplate := sysskilltemplate.GetSystemSkillTemplateService().GetSystemSkillTemplateByTypeAndLevel(typ, subType, 1)
	if skTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
			"subType":  subType,
		}).Warn("systemskill:模板不存在")
		return
	}

	//获取对应系统阶数
	curNumber := systemskill.GetSystemAdvancedNum(pl, typ)
	if curNumber > 0 && curNumber < skTemplate.GetNumber() {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"typ":        typ,
			"subType":    subType,
			"curNumber":  curNumber,
			"needNumber": skTemplate.GetNumber(),
		}).Warn("systemskill:系统阶数不足,无法激活")
		playerlogic.SendSystemMessage(pl, lang.SystemSkillActiveNoNumber)
		return
	}

	// 装备数量
	if additionType, ok := typ.ConvertToAdditionSysType(); ok {
		additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
		equipNum := additionsysManager.GetAdditionSysEquipNum(additionType, skTemplate.GetNeedEquipQuality())
		if equipNum < skTemplate.GetNeedEquipCount() {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"typ":          typ,
					"subType":      subType,
					"equipNum":     equipNum,
					"needEquipNum": skTemplate.GetNumber(),
				}).Warn("systemskill:系统技能激活，装备数量不足")
			playerlogic.SendSystemMessage(pl, lang.SystemSkillActiveNotEnoughEquipNum)
			return
		}
	}

	needItemMap := skTemplate.GetNeedItemMap()
	needYinLiang := int64(skTemplate.GetCostSilver())
	needGold := int64(skTemplate.GetCostGold())

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//银两判断
	if needYinLiang != 0 {
		flag := propertyManager.HasEnoughSilver(needYinLiang)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
			}).Warn("systemskill:银两不足，无法激活")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//判断元宝
	if needGold != 0 {
		flag := propertyManager.HasEnoughGold(needGold, false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
			}).Warn("systemskill:元宝不足，无法激活")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//物品判断
	if len(needItemMap) != 0 {
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItems(needItemMap)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
			}).Warn("systemskill:道具不足，无法激活")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		inventoryReason := commonlog.InventoryLogReasonSystemSkillActive
		reasonText := fmt.Sprintf(inventoryReason.String(), skTemplate.TemplateId())
		flag = inventoryManager.BatchRemove(needItemMap, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("systemskill: systemSkillActive use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//消耗银两
	if needYinLiang != 0 || needGold != 0 {
		goldReason := commonlog.GoldLogReasonSystemSkillActive
		goldReasonText := fmt.Sprintf(goldReason.String(), skTemplate.TemplateId())
		silverReason := commonlog.SilverLogReasonSystemSkillActive
		silverReasonText := fmt.Sprintf(silverReason.String(), skTemplate.TemplateId())
		flag := propertyManager.Cost(0, needGold, goldReason, goldReasonText, needYinLiang, silverReason, silverReasonText)
		if !flag {
			panic(fmt.Errorf("systemskill: systemSkillActive Cost  should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	obj, flag := manager.SystemSkillActive(typ, subType)
	if !flag {
		panic(fmt.Errorf("systemskill: systemSkillActive should be ok"))
	}

	scSystemSkillActive := pbutil.BuildSCSystemSkillActive(obj)
	pl.SendMsg(scSystemSkillActive)
	return
}
