package activity_handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	gametemplate "fgame/fgame/game/template"
	yuxiscene "fgame/fgame/game/yuxi/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeYuXi, activity.ActivityAttendHandlerFunc(playerEnterYuXiScene))
}

func playerEnterYuXiScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {

	// 获取场景
	s, isEnd := yuxiscene.GetYuXiScene(activityTemplate)
	if isEnd {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("yuxi:玉玺之战已经结束")
		playerlogic.SendSystemMessage(pl, lang.ActivityHadFinish)
		return
	}
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("yuxi:玉玺之战场景不存在")
		playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
		return
	}

	bornPos := s.MapTemplate().GetBornPos()
	if !scenelogic.PlayerEnterScene(pl, s, bornPos) {
		return
	}
	flag = true
	return
}
