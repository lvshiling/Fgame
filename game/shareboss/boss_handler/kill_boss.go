package boss_handler

import (
	"fgame/fgame/common/lang"
	crosslogic "fgame/fgame/game/cross/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/shareboss/shareboss"
	sharebosstemplate "fgame/fgame/game/shareboss/template"
	"fgame/fgame/game/team/team"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	worldboss.RegistKillBossHandler(worldbosstypes.BossTypeShareBoss, worldboss.KillBossHandlerFunc(killShareBoss))

}

func killShareBoss(pl player.Player, typ worldbosstypes.BossType, biologyId int32) {
	//战斗状态
	if pl.IsPvpBattle() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bossId":   biologyId,
			}).Warn("scene:处理进入世界场景,正在pvp")
		playerlogic.SendSystemMessage(pl, lang.PlayerInPVP)
		return
	}

	bossTemp := sharebosstemplate.GetShareBossTemplateService().GetShareBossTemplateByBiologyId(typ, biologyId)
	if bossTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shareBoss:boss不存在")

		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	myTeam := team.GetTeamService().GetTeamByPlayerId(pl.GetId())
	if myTeam != nil && myTeam.IsMatch() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shareBoss:3v3匹配中")
		playerlogic.SendSystemMessage(pl, lang.TeamInMatch)
		return
	}
	crossType := typ.CrossType()
	//TODO 判断是否在跨服中
	arg := fmt.Sprintf("%d", biologyId)
	crosslogic.PlayerEnterCross(pl, crossType, arg)

	bossId := int32(bossTemp.GetBiologyTemplate().TemplateId())
	boss := shareboss.GetShareBossService().GetShareBoss(typ, bossId)
	scChallengeWorldBoss := pbutil.BuildSCChallengeWorldBoss(boss.GetPosition(), int32(typ))
	pl.SendMsg(scChallengeWorldBoss)
}
