package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/core/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/pbutil"
	gamesession "fgame/fgame/game/session"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_PLAYER_KILL_BIOLOGY_TYPE), dispatch.HandlerFunc(handlePlayerKillBiology))
}

func handlePlayerKillBiology(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:处理玩家击杀怪物")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isPlayerKillBiology := msg.(*crosspb.ISPlayerKillBiology)
	biologyId := isPlayerKillBiology.GetBiologyId()
	err = playerKillBoss(tpl, biologyId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Error("scene:处理玩家击杀怪物,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"biologyId": biologyId,
		}).Debug("scene:处理玩家击杀怪物,完成")
	return nil

}

func playerKillBoss(pl player.Player, biologyId int32) (err error) {
	to := template.GetTemplateService().Get(int(biologyId), (*gametemplate.BiologyTemplate)(nil))
	if to == nil {
		return
	}
	//跨服增加经验
	biologyTemplate := to.(*gametemplate.BiologyTemplate)
	expBase := biologyTemplate.ExpBase
	expPoint := biologyTemplate.ExpPoint
	propertylogic.AddExpKillMonster(pl, biologyId, int64(expBase), int64(expPoint))

	gameevent.Emit(sceneeventtypes.EventTypeCrossKillBiology, pl, biologyId)
	siTuLongKillBoss := pbutil.BuildSIPlayerKillBiology(biologyId)
	pl.SendMsg(siTuLongKillBoss)
	return
}
