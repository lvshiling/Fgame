package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
	babytypes "fgame/fgame/game/baby/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_LOCK_SKILL_TYPE), dispatch.HandlerFunc(handleBabyLockSkill))
}

//处理宝宝锁定技能
func handleBabyLockSkill(s session.Session, msg interface{}) (err error) {
	log.Debug("baby:处理宝宝锁定技能消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSBabyLockSkill)
	babyId := csMsg.GetBabyId()
	skillIndex := csMsg.GetSkillIndex()
	operation := csMsg.GetOperation()

	operationType := babytypes.SkillStatusType(operation)
	if !operationType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"operationType": operationType,
			}).Warn("baby:处理宝宝锁定技能,类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = handlerLockSkill(tpl, babyId, skillIndex, operationType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"operationType": operationType,
				"babyId":        babyId,
				"skillIndex":    skillIndex,
				"error":         err,
			}).Error("baby:处理宝宝锁定技能消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":      pl.GetId(),
			"operationType": operationType,
			"babyId":        babyId,
			"skillIndex":    skillIndex,
		}).Debug("baby:处理宝宝锁定技能消息完成")
	return nil

}

// 锁定技能
func handlerLockSkill(pl player.Player, babyId int64, skillIndex int32, operationType babytypes.SkillStatusType) (err error) {
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	baby := babyManager.GetBabyInfo(babyId)
	if baby == nil {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"operationType": operationType,
				"babyId":        babyId,
				"skillIndex":    skillIndex,
			}).Warn("baby:处理宝宝锁定技能, 宝宝不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//
	talent := babyManager.GetTalentInfo(babyId, skillIndex)
	if talent == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"babyId":     babyId,
				"skillIndex": skillIndex,
			}).Warn("baby:处理宝宝锁定技能, 宝宝技能不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if talent.Status == operationType {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"babyId":     babyId,
				"skillIndex": skillIndex,
			}).Warn("baby:处理宝宝锁定技能, 重复操作")
		playerlogic.SendSystemMessage(pl, lang.BabyTalentLockFail)
		return
	}

	// 锁定消耗
	if operationType == babytypes.SkillStatusTypeLock {
		lockTimes := baby.GetLockTimes() + 1
		nextBabyUnlockTemp := babytemplate.GetBabyTemplateService().GetBabyUnlockTalentTemplate(lockTimes)
		needGold := int64(nextBabyUnlockTemp.SuodingGold)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		if !propertyManager.HasEnoughGold(needGold, false) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"needGold": needGold,
					"babyId":   babyId,
				}).Warn("baby:处理宝宝锁定技能, 元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
		goldCostReason := commonlog.GoldLogReasonBabyLockSkillUse
		goldCostReasonText := fmt.Sprintf(goldCostReason.String(), babyId, lockTimes)
		flag := propertyManager.CostGold(needGold, false, goldCostReason, goldCostReasonText)
		if !flag {
			panic(fmt.Errorf("baby: 宝宝锁定技能消耗元宝应该成功"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	flag := babyManager.LockBabySkill(babyId, skillIndex, operationType)
	if !flag {
		panic(fmt.Errorf("baby：宝宝锁定或解锁天赋技能应该成功,babyId:%d,skillIndex:%d,operat:%d", babyId, skillIndex, operationType))
	}

	scMsg := pbutil.BuildSCBabyLockSkill(babyId, skillIndex, int32(operationType))
	pl.SendMsg(scMsg)
	return
}
