package handler

import (
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_CS_REENTER_SCENE_TYPE), dispatch.HandlerFunc(handleReenter))
}

//处理进入场景
func handleReenter(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:处理重新进入场景")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)
	err = playerReenterScene(tpl)
	if err != nil {
		return err
	}
	return nil
}

//玩家重新进入
func playerReenterScene(pl scene.Player) (err error) {
	scenelogic.PlayerReenterScene(pl)
	return nil
}
