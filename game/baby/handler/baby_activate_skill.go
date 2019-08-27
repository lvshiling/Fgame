package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	babylogic "fgame/fgame/game/baby/logic"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_ACTIVATE_SKILL_TYPE), dispatch.HandlerFunc(handleBabyActivateSkill))
}

//处理宝宝激活技能
func handleBabyActivateSkill(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理宝宝激活技能消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSBabyActivateSkill)
	babyId := csMsg.GetBabyId()

	err = handlerActivateSkill(tpl, babyId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("baby:处理宝宝激活技能消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("baby:处理宝宝激活技能消息完成")
	return nil

}

// 激活技能
func handlerActivateSkill(pl player.Player, babyId int64) (err error) {
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	baby := babyManager.GetBabyInfo(babyId)
	if baby == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"babyId":   babyId,
			}).Warn("baby:处理宝宝激活技能, 宝宝不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//技能数量
	if babyManager.IsFullTalent(babyId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"babyId":   babyId,
			}).Warn("baby:处理宝宝激活技能, 没有可激活的技能")
		playerlogic.SendSystemMessage(pl, lang.BabyFailActivitySkill)
		return
	}

	activateTimes := baby.GetActivateTimes() + 1
	nextBabyUnlockTemp := babytemplate.GetBabyTemplateService().GetBabyUnlockTalentTemplate(activateTimes)
	needGold := int64(nextBabyUnlockTemp.UseGold)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(needGold, false) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
				"babyId":   babyId,
			}).Warn("baby:处理宝宝激活技能, 元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	goldCostReason := commonlog.GoldLogReasonBabyActivateSkillUse
	goldCostReasonText := fmt.Sprintf(goldCostReason.String(), babyId, activateTimes)
	flag := propertyManager.CostGold(needGold, false, goldCostReason, goldCostReasonText)
	if !flag {
		panic(fmt.Errorf("baby: 宝宝激活技能消耗元宝应该成功"))
	}

	flag = babyManager.ActivateSkill(babyId)
	if !flag {
		panic(fmt.Errorf("baby: 宝宝激活技能应该成功"))
	}

	propertylogic.SnapChangedProperty(pl)
	babylogic.BabyPropertyChanged(pl)

	scMsg := pbutil.BuildSCBabyActivateSkill(babyId, baby.GetSkillList())
	pl.SendMsg(scMsg)
	return
}
