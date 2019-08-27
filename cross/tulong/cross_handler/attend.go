package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/cross/tulong/pbutil"
	tulongscene "fgame/fgame/cross/tulong/scene"
	"fgame/fgame/cross/tulong/tulong"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/global"
	playerlogic "fgame/fgame/game/player/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_TULONG_ATTEND_TYPE), dispatch.HandlerFunc(handleTuLongAttend))
}

//处理参与屠龙
func handleTuLongAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("tuLong:处理参与屠龙")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = tuLongAttend(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("tuLong:处理参与屠龙,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("tuLong:处理参与屠龙,完成")
	return nil

}

//参与屠龙
func tuLongAttend(pl *player.Player) (err error) {
	isLineup := false
	sceneId := int64(0)
	sc := tulong.GetTuLongService().GetTuLongScene()
	if sc == nil {
		now := global.GetGame().GetTimeService().Now()
		activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeCoressTuLong)
		activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, 0, 0)
		if err != nil {
			return err
		}
		if activityTimeTemplate == nil {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"now":          now,
					"activityType": activitytypes.ActivityTypeCoressTuLong,
				}).Warn("tulong:处理参加屠龙,活动时间模板不存在")
			playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
			pl.Close(nil)
			return nil
		}

		endTime, _ := activityTimeTemplate.GetEndTime(now)
		sc = tulong.GetTuLongService().CreateTuLongScene(activityTemplate.Mapid, endTime)
		if sc == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("tulong:处理参加屠龙,创建场景失败")
			playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
			pl.Close(nil)
			return nil
		}

		sd, ok := sc.SceneDelegate().(tulongscene.TuLongSceneData)
		if !ok {
			pl.Close(nil)
			return nil
		}
		isLineup = sd.IfLineup()
		sceneId = sc.Id()
	}

	isMsg := pbutil.BuildISTuLongAttend(isLineup, sceneId)
	pl.SendMsg(isMsg)
	return
}
