package handler

//废弃:zrc
// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	"fgame/fgame/core/session"
// 	crosslogic "fgame/fgame/game/cross/logic"
// 	crosstypes "fgame/fgame/game/cross/types"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"
// 	"fgame/fgame/game/shareboss/pbutil"
// 	"fgame/fgame/game/shareboss/shareboss"
// 	sharebosstemplate "fgame/fgame/game/shareboss/template"
// 	"fgame/fgame/game/team/team"
// 	"fmt"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_SHARE_BOSS_CHALLENGE_TYPE), dispatch.HandlerFunc(handlerKillShareBoss))
// }

// //前往击杀跨服世界BOSS请求
// func handlerKillShareBoss(s session.Session, msg interface{}) (err error) {
// 	log.Debug("shareBoss:处理前往击杀跨服世界BOSS请求")

// 	pl := gamesession.SessionInContext(s.Context()).Player()
// 	tpl := pl.(player.Player)
// 	cs := msg.(*uipb.CSShareBossChallenge)
// 	bossBiologyId := cs.GetBiologyId()

// 	err = killShareBoss(tpl, bossBiologyId)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": tpl.GetId(),
// 				"err":      err,
// 			}).Error("shareBoss:处理前往击杀跨服世界BOSS请求，错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": tpl.GetId(),
// 		}).Debug("shareBoss:处理前往击杀跨服世界BOSS请求完成")

// 	return
// }

// func killShareBoss(pl player.Player, bossBiologyId int32) (err error) {
// 	//战斗状态
// 	if pl.IsPvpBattle() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"bossId":   bossBiologyId,
// 			}).Warn("scene:处理进入世界场景,正在pvp")
// 		playerlogic.SendSystemMessage(pl, lang.PlayerInPVP)
// 		return
// 	}

// 	bossTemp := sharebosstemplate.GetShareBossTemplateService().GetShareBossTemplateByBiologyId(bossBiologyId)
// 	if bossTemp == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warn("shareBoss:boss不存在")

// 		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	myTeam := team.GetTeamService().GetTeamByPlayerId(pl.GetId())
// 	if myTeam != nil && myTeam.IsMatch() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warn("shareBoss:3v3匹配中")
// 		playerlogic.SendSystemMessage(pl, lang.TeamInMatch)
// 		return
// 	}

// 	//TODO 判断是否在跨服中
// 	arg := fmt.Sprintf("%d", bossBiologyId)
// 	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeWorldboss, arg)

// 	bossId := int32(bossTemp.GetBiologyTemplate().TemplateId())
// 	boss := shareboss.GetShareBossService().GetShareBoss(bossId)
// 	scChallengeWorldBoss := pbutil.BuildSCShareBossChallenge(boss.GetPosition())
// 	pl.SendMsg(scChallengeWorldBoss)

// 	return
// }
