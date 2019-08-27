package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/lianyu/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lianyu/lianyu"
	playerlogic "fgame/fgame/game/player/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_LIANYU_ATTEND_TYPE), dispatch.HandlerFunc(handleLianYuAttend))
}

//处理参加无间炼狱
func handleLianYuAttend(s session.Session, msg interface{}) (err error) {
	log.Info("lianyu:处理参加无间炼狱")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = lianYuAttend(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("lianyu:处理参加无间炼狱,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("lianyu:处理参加无间炼狱,完成")
	return nil

}

//参加无间炼狱
func lianYuAttend(pl *player.Player) (err error) {
	//活动未开始
	s := lianyu.GetLianYuService().GetLianYuScene()
	if s == nil {
		now := global.GetGame().GetTimeService().Now()
		activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeLianYu)
		activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, 0, 0)
		if err != nil {
			return err
		}
		if activityTimeTemplate == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("lianyu:处理参加无间炼狱,活动时间模板不存在")
			playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
			pl.Close(nil)
			return nil
		}
		endTime, err := activityTimeTemplate.GetEndTime(now)
		if err != nil {
			return err
		}
		s = lianyu.GetLianYuService().CreateLianYuScene(activityTemplate.Mapid, endTime)
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("lianyu:处理参加无间炼狱,活动未开始")
			playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
			pl.Close(nil)
			return nil
		}
	}

	pos, flag := lianyu.GetLianYuService().GetHasLineUp(pl.GetId())
	if flag {
		isArenaMatchResult := pbutil.BuildISLianYuAttend(flag, pos)
		pl.SendMsg(isArenaMatchResult)
		return
	}

	pos, flag = lianyu.GetLianYuService().Attend(pl.GetId())
	isArenaMatchResult := pbutil.BuildISLianYuAttend(flag, pos)
	pl.SendMsg(isArenaMatchResult)
	return
}
