package battle

import (
	"fgame/fgame/common/lang"
	relivelogic "fgame/fgame/cross/relive/logic"
	chuangshitemplate "fgame/fgame/game/chuangshi/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterReliveHandler(scenetypes.SceneTypeChuangShiZhiZhanFuShu, scene.ReliveHandlerFunc(Relive))
}

//城战原地复活
func Relive(pl scene.Player, autoBuy bool) {
	s := pl.GetScene()
	sd := s.SceneDelegate()
	if sd == nil {
		return
	}

	_, ok := pl.(player.Player)
	if !ok {
		return
	}

	reliveTime := pl.GetChuangShiReliveTimes()
	maxReliveTime := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarTemp().RebornSitu
	if reliveTime >= maxReliveTime {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"reliveTime":    reliveTime,
				"maxReliveTime": maxReliveTime,
			}).Warn("chuangshi:复活超过最大次数")
		playerlogic.SendSystemMessage(pl, lang.ReliveBaseYuanDiMaxTimes)
		return
	}

	relivelogic.Relive(pl)
}
