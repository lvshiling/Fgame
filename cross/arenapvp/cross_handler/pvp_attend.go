package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/arenapvp/arenapvp"
	"fgame/fgame/cross/arenapvp/pbutil"
	arenapvpscene "fgame/fgame/cross/arenapvp/scene"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	gamearenapvptypes "fgame/fgame/game/arenapvp/types"
	"fgame/fgame/game/global"
	playerlogic "fgame/fgame/game/player/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_ARENAPVP_ATTEND_TYPE), dispatch.HandlerFunc(handleArenapvpAttend))
}

//处理参与pvp海选
func handleArenapvpAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("arenapvp:处理参与pvp海选")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = arenapvpAttend(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arenapvp:处理参与pvp海选,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arenapvp:处理参与pvp海选,完成")
	return nil

}

//参与pvp海选
func arenapvpAttend(pl *player.Player) (err error) {

	isLineup := false
	sceneId := int64(0)
	battleS := arenapvp.GetArenapvpService().GetArenapvpBattleScene(pl.GetId())
	if battleS == nil {

		now := global.GetGame().GetTimeService().Now()
		activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeArenapvp)
		activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, 0, 0)
		if err != nil {
			return err
		}
		if activityTimeTemplate == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"now":      now,
					"type":     activitytypes.ActivityTypeArenapvp,
				}).Warn("pvp:处理参加pvp海选,活动时间模板不存在")
			playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
			pl.Close(nil)
			return nil
		}

		endTime, _ := activityTimeTemplate.GetEndTime(now)
		pvpTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpTemplate(gamearenapvptypes.ArenapvpTypeElection)
		sceneEndTime := pvpTemp.GetEndTime(now)
		if now > sceneEndTime {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"Mapid":        activityTemplate.Mapid,
					"sceneEndTime": sceneEndTime,
				}).Warn("pvp:pvp海选已经结束")
			playerlogic.SendSystemMessage(pl, lang.ArenapvpElectionEnd)
			pl.Close(nil)
			return nil
		}

		s := arenapvp.GetArenapvpService().GetArenapvpSceneElection()
		if s == nil {
			s = arenapvp.GetArenapvpService().CreateArenapvpSceneElection(endTime)
			if s == nil {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
					}).Warn("pvp:处理参加pvp海选,活动未开始")
				playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
				pl.Close(nil)
				return nil
			}
		}

		sd, ok := s.SceneDelegate().(arenapvpscene.ArenapvpSceneData)
		if !ok {
			return nil
		}
		isLineup = sd.IfLineup()
		sceneId = s.Id()
	}

	isMsg := pbutil.BuildISArenapvpAttend(isLineup, sceneId)
	pl.SendMsg(isMsg)
	return
}
