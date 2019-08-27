package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	skilllogic "fgame/fgame/game/skill/logic"
	"fgame/fgame/game/skill/pbutil"
	playerskill "fgame/fgame/game/skill/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SKILL_UPGRADE_TYPE), dispatch.HandlerFunc(handleRoleSkillUpgrade))
}

//处理升级职业技能信息
func handleRoleSkillUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("skill:处理升级职业技能信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSkillUpgrade := msg.(*uipb.CSSkillUpgrade)
	skillId := csSkillUpgrade.GetSkillId()

	err = skillUpgrade(tpl, skillId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"skillId":  skillId,
				"error":    err,
			}).Error("skill:处理升级职业技能信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
		}).Debug("skill:处理升级职业技能信息完成")
	return nil
}

//处理技能升级的逻辑
func skillUpgrade(pl player.Player, skillId int32) (err error) {
	skillManager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	flag := skillManager.IsValid(skillId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
		}).Warn("skill:无效技能")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	flag = skillManager.IfSkillExist(skillId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
		}).Warn("skill:还没有该技能,请先获取")
		playerlogic.SendSystemMessage(pl, lang.SkillNotHas)
		return
	}
	skillInfo := skillManager.GetSkill(skillId)
	flag = skillManager.IfCanUpgradeSkill(skillId, skillInfo.Level)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
		}).Warn("skill:当前技能已达最高级，无法升级")
		playerlogic.SendSystemMessage(pl, lang.SkillReachLimit)
		return
	}

	//升级所需银量
	costSilver := skilllogic.UpgradeConsumeSilver(skillId, skillInfo.Level+1)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag = propertyManager.HasEnoughSilver(int64(costSilver))
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
		}).Warn("skill:银两不足，无法升级")
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}

	//消耗银两
	silverReason := commonlog.SilverLogReasonSkillUpgrade
	reasonText := fmt.Sprintf(silverReason.String(), skillId, skillInfo.Level)
	flag = propertyManager.CostSilver(int64(costSilver), silverReason, reasonText)
	if !flag {
		panic(fmt.Errorf("skill: skillUpgrade cost sliver should be ok"))
	}
	propertylogic.SnapChangedProperty(pl)

	flag = skillManager.UpgradeSkill(skillId, skillInfo.Level)
	if !flag {
		panic(fmt.Errorf("skill: skillUpgrade  should be ok"))
	}

	skilllogic.SkillPropertyChanged(pl)
	scSkillUpgrade := pbutil.BuildSCSkillUpgrade(skillId)
	pl.SendMsg(scSkillUpgrade)
	return
}
