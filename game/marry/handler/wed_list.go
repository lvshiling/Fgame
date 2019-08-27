package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/marry/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_WED_LIST_TYPE), dispatch.HandlerFunc(handleMarryWedList))
}

//处理婚期举办时间列表信息
func handleMarryWedList(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理婚期举办时间列表消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = marryWedList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("marry:处理婚期举办时间列表消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理婚期举办时间列表消息完成")
	return nil
}

//处理婚期举办时间列表信息逻辑
func marryWedList(pl player.Player) (err error) {
	weddingList := marry.GetMarryService().GetMarryWeddingList()
	scMarryWedList := pbuitl.BuildSCMarryWedList(weddingList)
	pl.SendMsg(scMarryWedList)
	return
}
