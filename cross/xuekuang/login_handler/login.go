package login_handler

import (
	"fgame/fgame/cross/login/login"
	"fgame/fgame/cross/player/player"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/global"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLoginHandler(crosstypes.CrossTypeXueKuang, login.LogincHandlerFunc(xueKuangLogin))
}

func xueKuangLogin(pl *player.Player, ct crosstypes.CrossType, crossArgs ...string) (flag bool) {
	now := global.GetGame().GetTimeService().Now()
	act := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeXueKuangCollect)
	timeTemplate, err := act.GetActivityTimeTemplate(now, 0, 0)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Warn("xuekuang:获取活动模板,错误")
		return false
	}
	if timeTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xuekuang:获取活动模板为空")
		return false
	}

	s := scene.GetSceneService().GetWorldSceneByMapId(act.Mapid)
	if s == nil {
		return
	}

	if !scenelogic.PlayerEnterScene(pl, s, s.MapTemplate().GetBornPos()) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xuekuang:进入场景失败")
		return false
	}
	return true

}
