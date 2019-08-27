package login_handler

import (
	"fgame/fgame/cross/login/login"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/tulong/tulong"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/global"
	scenelogic "fgame/fgame/game/scene/logic"
	tulongtemplate "fgame/fgame/game/tulong/template"
	tulongtypes "fgame/fgame/game/tulong/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLoginHandler(crosstypes.CrossTypeTuLong, login.LogincHandlerFunc(tuLongLogin))
}

func tuLongLogin(pl *player.Player, ct crosstypes.CrossType, crossArgs ...string) (flag bool) {

	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("tulong:玩家不在仙盟")
		return
	}

	sc := tulong.GetTuLongService().GetTuLongScene()
	if sc == nil {
		now := global.GetGame().GetTimeService().Now()
		act := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeCoressTuLong)
		timeTemplate, err := act.GetActivityTimeTemplate(now, 0, 0)
		if err != nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"err":      err,
				}).Warn("tulong:获取活动模板,错误")
			return false
		}
		if timeTemplate == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("tulong:获取活动模板为空")
			return false
		}
		endTime, err := timeTemplate.GetEndTime(now)
		if err != nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"err":      err,
				}).Warn("tulong:获取活动结束事件,错误")
			return false
		}
		sc = tulong.GetTuLongService().CreateTuLongScene(act.Mapid, endTime)
		if sc == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("tulong:创建场景失败")
			return false
		}
	}
	bornBiaoShi, flag := tulong.GetTuLongService().GetPlayerBornBiaoShi(allianceId)
	if !flag {
		return false
	}

	tuLongPosTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongPosTemplate(tulongtypes.TuLongPosTypePlayer, bornBiaoShi)
	if tuLongPosTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("tulong:获取位置失败")
		return false
	}

	pos := tuLongPosTemplate.GetPos()
	if !scenelogic.PlayerEnterScene(pl, sc, pos) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("tulong:进入场景失败")
		return false
	}
	return true
}
