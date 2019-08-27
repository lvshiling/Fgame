package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BUFF_LIST_TYPE), dispatch.HandlerFunc(handleBuffList))
}

//处理buff列表包
func handleBuffList(s session.Session, msg interface{}) error {
	log.Debugln("scene:处理对象buff列表")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)

	getBuffList(tpl)

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("scene:处理对象buff列表,完成")
	return nil
}

func getBuffList(pl scene.Player) {
	scBuffList := pbutil.BuildSCBuffList(pl.GetBuffs())
	pl.SendMsg(scBuffList)
}
