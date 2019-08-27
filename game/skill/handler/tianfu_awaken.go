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
	processor.Register(codec.MessageType(uipb.MessageType_CS_SKILL_TIANFU_AWAKEN_TYPE), dispatch.HandlerFunc(handleSkillTianFuAwaken))
}

//处理职业技能天赋信息
func handleSkillTianFuAwaken(s session.Session, msg interface{}) (err error) {
	log.Debug("skill:处理职业技能天赋信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSkillTianFuAwaken := msg.(*uipb.CSSkillTianFuAwaken)
	skillId := csSkillTianFuAwaken.GetSkillId()
	tianFuId := csSkillTianFuAwaken.GetTianFuId()

	err = skillTianFuAwaken(tpl, skillId, tianFuId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"skillId":  skillId,
				"tianFuId": tianFuId,
				"error":    err,
			}).Error("skill:处理职业技能天赋信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
			"tianFuId": tianFuId,
		}).Debug("skill:处理职业技能天赋信息完成")
	return nil
}

//处理技能升级的逻辑
func skillTianFuAwaken(pl player.Player, skillId int32, tianFuId int32) (err error) {
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
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
			"tianFuId": tianFuId,
		}).Warn("skill:当前天赋已经觉醒过")
		playerlogic.SendSystemMessage(pl, lang.SkillTianFuHasedAwaken)
		return
	}

	parentTianFuTemplate := skilltemplate.GetSkillTemplateService().GetSkillParentTianFuTemplate(skillId, tianFuId)
	if parentTianFuTemplate != nil {
		parentTianFuId := int32(parentTianFuTemplate.TemplateId())
		if !skillManager.HasedTianFuAwaken(skillId, parentTianFuId) {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"skillId":  skillId,
				"tianFuId": tianFuId,
			}).Warn("skill:您还未觉醒%s,无法觉醒当前天赋", parentTianFuTemplate.Name)
			playerlogic.SendSystemMessage(pl, lang.SkillTianFuAwakenNoAwakenParent, parentTianFuTemplate.Name)
			return
		}
	}

	tianFuLevelTemplate := tianFuTemplate.GetTianFuLevelByLevel(1)
	if tianFuLevelTemplate == nil {
		return
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItemMap := tianFuLevelTemplate.GetNeedItemMap(pl.GetRole())
	if len(needItemMap) != 0 {
		flag = inventoryManager.HasEnoughItems(needItemMap)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"skillId":  skillId,
				"tianFuId": tianFuId,
			}).Warn("skill:道具不足,无法觉醒")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonTianFuAwaken.String(), skillId, tianFuId)
		flag := inventoryManager.BatchRemove(needItemMap, commonlog.InventoryLogReasonTianFuAwaken, reasonText)
		if !flag {
			panic(fmt.Errorf("skill: skillTianFuAwaken use item should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	curLevel := int32(0)
	sucess := false
	if mathutils.RandomHit(common.MAX_RATE, int(tianFuLevelTemplate.UpdateWfb)) {
		//天赋觉醒
		flag = skillManager.TianFuAwaken(skillId, tianFuId)
		if !flag {
			panic(fmt.Errorf("skill: skillTianFuAwaken TianFuAwaken should be ok"))
		}
		sucess = true
		curLevel = 1
	}

	scSkillTianFuAwaken := pbutil.BuildSCSkillTianFuAwaken(skillId, tianFuId, curLevel, sucess)
	pl.SendMsg(scSkillTianFuAwaken)
	return
}
