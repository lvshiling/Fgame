package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/jieyi/pbutil"
	"fgame/fgame/game/player"

	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_LEAVE_WORD_INFO_TYPE), dispatch.HandlerFunc(handleJieYiLeaveWordInfo))
}

func handleJieYiLeaveWordInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理结义墙请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	p := gcs.Player()
	pl := p.(player.Player)

	err = jieYiLeaveWordInfo(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理结义墙请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("jieyi: 处理结义墙请求消息,成功")

	return
}

func jieYiLeaveWordInfo(pl player.Player) (err error) {
	objList := jieyi.GetJieYiService().GetJieYiLeaveWord()
	scMsg := pbutil.BuildSCJieYiLeaveWordInfo(objList)
	pl.SendMsg(scMsg)
	return
}
