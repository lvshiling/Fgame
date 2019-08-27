package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	tulongequiplogic "fgame/fgame/game/tulongequip/logic"
	"fgame/fgame/game/tulongequip/pbutil"
	playertulongequip "fgame/fgame/game/tulongequip/player"
	tulongequiptemplate "fgame/fgame/game/tulongequip/template"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TULONG_EQUIP_SKILL_UPGRADE_TYPE), dispatch.HandlerFunc(handleTuLongEquipSkillUpgrade))
}

//处理屠龙装备技能升级
func handleTuLongEquipSkillUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("tulongequipSkill:处理屠龙装备技能升级")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSTuLongEquipSkillUpgrade)
	suitInt := csMsg.GetSuitType()

	suitType := tulongequiptypes.TuLongSuitType(suitInt)
	if !suitType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitType": suitType,
			}).Warn("tulongequipSkill:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = tulongequipSkillUpgrade(tpl, suitType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitType": suitType,
				"error":    err,
			}).Error("tulongequipSkill:处理屠龙装备技能升级,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"suitType": suitType,
		}).Debug("tulongequipSkill:处理屠龙装备技能升级,完成")
	return nil

}

//屠龙装备技能升级
func tulongequipSkillUpgrade(pl player.Player, suitType tulongequiptypes.TuLongSuitType) (err error) {
	tulongequipManager := pl.GetPlayerDataManager(playertypes.PlayerTuLongEquipDataManagerType).(*playertulongequip.PlayerTuLongEquipDataManager)

	curLevel := tulongequipManager.GetSuitSkillLevel(suitType)
	nextLevel := curLevel + 1
	nextSkillTemp := tulongequiptemplate.GetTuLongEquipTemplateService().GetTuLongEquipTemplateSkill(suitType, nextLevel)
	if nextSkillTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"suitType":  suitType,
				"nextLevel": nextLevel,
			}).Warn("tulongequipSkill:屠龙装备技能升级，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	// 激活条件
	equipNum := tulongequipManager.GetTuLongEquipNumByJieShu(suitType, nextSkillTemp.NeedJieShu)
	totalLevel := tulongequipManager.GetTuLongEquipTotalLevel(suitType)
	if equipNum < nextSkillTemp.NeedEquipNum || totalLevel < nextSkillTemp.NeedStrengthenLevel {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"nextLevel":  nextLevel,
				"equipNum":   equipNum,
				"totalLevel": totalLevel,
			}).Warn("tulongequipSkill:屠龙装备技能不满足条件，无法升级")
		playerlogic.SendSystemMessage(pl, lang.TuLongEquipSkillFailed)
		return
	}

	//物品
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	useItemMap := nextSkillTemp.GetUseItemMap()
	if len(useItemMap) > 0 {
		if !inventoryManager.HasEnoughItems(useItemMap) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"useItemMap": useItemMap,
				}).Warn("tulongequip:强化升级失败,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		//消耗材料
		reason := commonlog.InventoryLogReasonTuLongEquipStrengthenUse
		flag := inventoryManager.BatchRemove(useItemMap, reason, reason.String())
		if !flag {
			panic(fmt.Errorf("tulongequip:屠龙装备升级技能移除材料应该成功"))
		}
	}

	//计算成功
	success := mathutils.RandomHit(common.MAX_RATE, int(nextSkillTemp.UplevelRate))
	if success {
		tulongequipManager.UpgradeSuitSkill(suitType)
	}

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)
	tulongequiplogic.TuLongEquipPropertyChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCTuLongEquipSkillUpgrade(int32(suitType), nextLevel, success)
	pl.SendMsg(scMsg)
	return
}
