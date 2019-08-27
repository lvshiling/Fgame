package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancetemplate "fgame/fgame/game/alliance/template"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_SKILL_UPGRADE_TYPE), dispatch.HandlerFunc(handleAllianceSkillUpgrade))
}

//处理仙术升级
func handleAllianceSkillUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("allianceSkill:处理仙术升级")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csAllianceSkillUpgrade := msg.(*uipb.CSAllianceSkillUpgrade)
	typ := csAllianceSkillUpgrade.GetSkillType()

	skillType := alliancetypes.AllianceSkillType(typ)
	if !skillType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"skillType": skillType,
			}).Warn("allianceSkill:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = allianceSkillUpgrade(tpl, skillType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"skillType": skillType,
				"error":     err,
			}).Error("allianceSkill:处理仙术升级,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"skillType": skillType,
		}).Debug("allianceSkill:处理仙术升级,完成")
	return nil

}

//仙术升级
func allianceSkillUpgrade(pl player.Player, skillType alliancetypes.AllianceSkillType) (err error) {
	allianceManager := pl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	curLevel := allianceManager.GetAllianceSkillLevel(skillType)
	nextLevel := curLevel + 1

	nextSkillTemp := alliancetemplate.GetAllianceTemplateService().GetAllianceSkillTemplateByType(nextLevel, skillType)
	if nextSkillTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"nextLevel": nextLevel,
				"skillType": skillType,
			}).Warn("allianceSkill:仙术升级，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag := allianceManager.IsOpenAllianceSkill(nextLevel, skillType)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"nextLevel": nextLevel,
				"skillType": skillType,
			}).Warn("allianceSkill:仙术未开启，无法升级")
		playerlogic.SendSystemMessage(pl, lang.AllianceSkillNotOpen)
		return
	}

	flag = allianceManager.IsAllianceSkillFullLevel(skillType)
	if flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"nextLevel": nextLevel,
				"skillType": skillType,
			}).Warn("allianceSkill:仙术等级已达上限")
		playerlogic.SendSystemMessage(pl, lang.AllianceSkillMaxLevel)
		return
	}

	//贡献是否足够
	needContribution := nextSkillTemp.NeedContribution
	flag = allianceManager.IsEnoughGongXian(needContribution)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"skillType": skillType,
		}).Warn("allianceSkill:贡献不足，无法升级")
		playerlogic.SendSystemMessage(pl, lang.AllianceSkillNotEnoughGongXian)
		return
	}

	//扣除贡献
	flag = allianceManager.UseGongXian(needContribution)
	if !flag {
		panic("allianceSkill: cost gongxian should be ok")
	}

	flag = allianceManager.UpgradeAllianceSkill(skillType)
	if !flag {
		panic("allianceSkill:升级应该成功")
	}

	tempId := int32(nextSkillTemp.TemplateId())
	scAllianceSkillUpgrade := pbutil.BuildSCAllianceSkillUpgrade(tempId)
	pl.SendMsg(scAllianceSkillUpgrade)
	return
}
