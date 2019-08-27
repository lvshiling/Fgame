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
	processor.Register(codec.MessageType(uipb.MessageType_CS_SKILL_UPGRADE_ALL_TYPE), dispatch.HandlerFunc(handleRoleSkillUpgradeAll))
}

//处理升级全部职业技能信息
func handleRoleSkillUpgradeAll(s session.Session, msg interface{}) (err error) {
	log.Debug("skill:处理升级全部职业技能信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = skillUpgradeAll(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("skill:处理升级全部职业技能信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("skill:处理升级全部职业技能信息完成")
	return nil
}

//处理职业技能全部升级的逻辑
func skillUpgradeAll(pl player.Player) (err error) {
	skillManager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	skillIdMap := skillManager.CanUpgradeRoleSkills()
	if skillIdMap == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("skill:当前无可升级技能")
		playerlogic.SendSystemMessage(pl, lang.SKillNotHasUpgrade)
		return
	}

	//升级所有技能所需银量
	totalCostSilver := int64(0)
	for skillId, level := range skillIdMap {
		costSilver := skilllogic.UpgradeConsumeSilver(skillId, level+1)
		totalCostSilver += int64(costSilver)
	}

	//银量是否足够
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.HasEnoughSilver(totalCostSilver)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("skill:银两不足，无法升级")
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}

	//消耗银两
	reasonText := commonlog.SilverLogReasonSkillUpgradeAll.String()
	flag = propertyManager.CostSilver(totalCostSilver, commonlog.SilverLogReasonSkillUpgradeAll, reasonText)
	if !flag {
		panic(fmt.Errorf("skill: skillUpgrade cost sliver should be ok"))
	}
	propertylogic.SnapChangedProperty(pl)

	flag = skillManager.UpgradeSkillAll(skillIdMap)
	if !flag {
		panic(fmt.Errorf("skill: handleSkillUpgradeAll  should be ok"))
	}

	skilllogic.SkillPropertyChanged(pl)
	scSkillUpgradeAll := pbutil.BuildSCSkillUpgradeAll(skillIdMap)
	pl.SendMsg(scSkillUpgradeAll)
	return
}
