package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	tulonglogic "fgame/fgame/game/tulong/logic"
	"fgame/fgame/game/tulong/pbutil"
	tulongtemplate "fgame/fgame/game/tulong/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_TULONG_KILL_BOSS_TYPE), dispatch.HandlerFunc(handleTuLongKillBoss))
}

//处理跨服屠龙击杀boss
func handleTuLongKillBoss(s session.Session, msg interface{}) (err error) {
	log.Debug("tulong:处理跨服屠龙击杀boss")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isTuLongKillBoss := msg.(*crosspb.ISTuLongKillBoss)
	bossId := isTuLongKillBoss.GetBossId()
	err = tuLongKillBoss(tpl, bossId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bossId":   bossId,
			}).Error("tulong:处理跨服屠龙击杀boss,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"bossId":   bossId,
		}).Debug("tulong:处理跨服屠龙击杀boss,完成")
	return nil

}

//跨服屠龙击杀boss
func tuLongKillBoss(pl player.Player, bossId int32) (err error) {
	tuLongTemplate, flag := tulongtemplate.GetTuLongTemplateService().GetTuLongTemplateByBiologyId(bossId)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bossId":   bossId,
			}).Warn("tulong:处理跨服屠龙击杀boss,模板不存在")
		return
	}

	allianceId := pl.GetAllianceId()
	al := alliance.GetAllianceService().GetAlliance(allianceId)
	if al == nil {
		panic(fmt.Errorf("tulong: 仙盟不应该为空"))
	}
	tulonglogic.GiveTuLongKillBossReward(al, tuLongTemplate)
	siTuLongKillBoss := pbutil.BuildSITuLongKillBoss()
	pl.SendCrossMsg(siTuLongKillBoss)
	return
}
