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
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/systemskill/pbutil"
	playersysskill "fgame/fgame/game/systemskill/player"
	"fgame/fgame/game/systemskill/systemskill"
	systemskilltemplate "fgame/fgame/game/systemskill/template"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SYSTEM_SKILL_UPGRADE_TYPE), dispatch.HandlerFunc(handleSystemSkillUpgrade))
}

//处理系统技能升级信息
func handleSystemSkillUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("systemskill:处理系统技能升级信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSystemSkillUpgrade := msg.(*uipb.CSSystemSkillUpgrade)
	typ := csSystemSkillUpgrade.GetTag()
	subType := csSystemSkillUpgrade.GetSubType()

	err = systemSkillUpgrade(tpl, sysskilltypes.SystemSkillType(typ), sysskilltypes.SystemSkillSubType(subType))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
				"error":    err,
			}).Error("systemskill:处理系统技能升级信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
			"subType":  subType,
		}).Debug("systemskill:处理系统技能升级完成")
	return nil
}

//系统技能升级的逻辑
func systemSkillUpgrade(pl player.Player, typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSystemSkillDataManagerType).(*playersysskill.PlayerSystemSkillDataManager)
	flag := manager.IfSystemSkillExist(typ, subType)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
			"subType":  subType,
		}).Warn("systemskill:未激活的系统技能,无法升级")
		playerlogic.SendSystemMessage(pl, lang.SystemSkillNotActiveNotUpgrade)
		return
	}

	flag = manager.IfCanUpgrade(typ, subType)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
			"subType":  subType,
		}).Warn("systemskill:系统技能已达最高级")
		playerlogic.SendSystemMessage(pl, lang.SystemSkillReacheFullUpgrade)
		return
	}

	curLevel := manager.GetSystemSkillLevelByTyp(typ, subType)
	nextLevel := curLevel + 1
	skTemplate := systemskilltemplate.GetSystemSkillTemplateService().GetSystemSkillTemplateByTypeAndLevel(typ, subType, nextLevel)
	if skTemplate == nil {
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
		}).Warn("systemskill:系统阶数不足")
		playerlogic.SendSystemMessage(pl, lang.SystemSkillActiveNoNumber)
		return
	}

	itemMap := skTemplate.GetNeedItemMap()
	needYinLiang := int64(skTemplate.GetCostSilver())
	needGold := int64(skTemplate.GetCostGold())
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//银两判断
	if needYinLiang != 0 {
		flag = propertyManager.HasEnoughSilver(needYinLiang)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
			}).Warn("systemskill:银两不足,无法升级")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//判断元宝
	if needGold != 0 {
		flag = propertyManager.HasEnoughGold(needGold, false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
			}).Warn("systemskill:元宝不足，无法升级")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//物品判断
	if len(itemMap) != 0 {
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItems(itemMap)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
			}).Warn("systemskill:道具不足，无法升级")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		inventoryReason := commonlog.InventoryLogReasonSystemSkillUpgrade
		reasonText := fmt.Sprintf(inventoryReason.String(), skTemplate.TemplateId())
		flag = inventoryManager.BatchRemove(itemMap, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("systemskill: systemSkillUpgarde use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//消耗银两
	if needYinLiang != 0 || needGold != 0 {
		goldReason := commonlog.GoldLogReasonSystemSkillUpgrade
		goldReasonText := fmt.Sprintf(goldReason.String(), skTemplate.TemplateId())
		silverReason := commonlog.SilverLogReasonSystemSkillUpgrade
		silverReasonText := fmt.Sprintf(silverReason.String(), skTemplate.TemplateId())
		flag = propertyManager.Cost(0, needGold, goldReason, goldReasonText, needYinLiang, silverReason, silverReasonText)
		if !flag {
			panic(fmt.Errorf("systemskill: systemSkillUpgrade Cost  should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	obj, flag := manager.Upgrade(typ, subType)
	if !flag {
		panic(fmt.Errorf("systemskill: Upgrade should be ok"))
	}

	scSystemSkillUpgrade := pbutil.BuildSCSystemSkillUpgrade(obj)
	pl.SendMsg(scSystemSkillUpgrade)
	return
}
