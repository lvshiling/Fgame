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
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/skill/pbutil"
	playerskill "fgame/fgame/game/skill/player"
	skilltemplate "fgame/fgame/game/skill/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SKILL_TIANFU_UPGRADE_TYPE), dispatch.HandlerFunc(handleSkillTianFuUpgrade))
}

//处理职业技能天赋升级信息
func handleSkillTianFuUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("skill:处理职业技能天赋升级信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSkillTianFuUpgrade := msg.(*uipb.CSSkillTianFuUpgrade)
	skillId := csSkillTianFuUpgrade.GetSkillId()
	tianFuId := csSkillTianFuUpgrade.GetTianFuId()

	err = skillTianFuUpgrade(tpl, skillId, tianFuId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"skillId":  skillId,
				"tianFuId": tianFuId,
				"error":    err,
			}).Error("skill:处理职业技能天赋升级信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
			"tianFuId": tianFuId,
		}).Debug("skill:处理职业技能天赋升级信息完成")
	return nil
}

//处理技能升级的逻辑
func skillTianFuUpgrade(pl player.Player, skillId int32, tianFuId int32) (err error) {
	skillManager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	flag := skillManager.IsValid(skillId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
			"tianFuId": tianFuId,
		}).Warn("skill:无效技能")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	tianFuTemplate := skilltemplate.GetSkillTemplateService().GetSkillTianFuTemplate(skillId, tianFuId)
	if tianFuTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
			"tianFuId": tianFuId,
		}).Warn("skill:无效技能")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	flag = skillManager.IfSkillExist(skillId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
			"tianFuId": tianFuId,
		}).Warn("skill:还没有该技能,请先获取")
		playerlogic.SendSystemMessage(pl, lang.SkillNotHas)
		return
	}

	flag = skillManager.HasedTianFuAwaken(skillId, tianFuId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
			"tianFuId": tianFuId,
		}).Warn("skill:您的天赋还未觉醒,无法升级")
		playerlogic.SendSystemMessage(pl, lang.SKillTianFuUpgradeNoAwaken)
		return
	}

	curLevel, flag := skillManager.GetSkillTianFuLevel(skillId, tianFuId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
			"tianFuId": tianFuId,
		}).Warn("skill:您的天赋还未觉醒,无法升级")
		playerlogic.SendSystemMessage(pl, lang.SKillTianFuUpgradeNoAwaken)
		return
	}

	tianFuNextLevelTemplate := tianFuTemplate.GetTianFuLevelByLevel(curLevel + 1)
	if tianFuNextLevelTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
			"tianFuId": tianFuId,
		}).Warn("skill:您的天赋等级已达满级")
		playerlogic.SendSystemMessage(pl, lang.SkillTianFuUpgradeFull)
		return
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItemMap := tianFuNextLevelTemplate.GetNeedItemMap(pl.GetRole())
	if len(needItemMap) != 0 {
		flag = inventoryManager.HasEnoughItems(needItemMap)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"skillId":  skillId,
				"tianFuId": tianFuId,
			}).Warn("skill:道具不足,无法升级")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonTianFuUpgrade.String(), skillId, tianFuId, curLevel)
		flag := inventoryManager.BatchRemove(needItemMap, commonlog.InventoryLogReasonTianFuUpgrade, reasonText)
		if !flag {
			panic(fmt.Errorf("skill: skillTianFuUpgrade use item should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	nextLevel := curLevel
	sucess := false
	if mathutils.RandomHit(common.MAX_RATE, int(tianFuNextLevelTemplate.UpdateWfb)) {
		//天赋升级
		flag = skillManager.TianFuUpgrade(skillId, tianFuId)
		if !flag {
			panic(fmt.Errorf("skill: skillTianFuUpgrade TianFuUpgrade should be ok"))
		}
		sucess = true
		nextLevel += 1
	}

	scSkillTianFuUpgrade := pbutil.BuildSCSkillTianFuUpgrade(skillId, tianFuId, nextLevel, sucess)
	pl.SendMsg(scSkillTianFuUpgrade)
	return
}
