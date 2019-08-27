package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	pbutil "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_JINIAN_GET), dispatch.HandlerFunc(handleMarryWedJiNianList))
}

//处理婚宴纪念的查询
func handleMarryWedJiNianList(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理婚宴纪念的查询")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = jinianList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("marry:处理婚宴纪念的查询,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理婚宴纪念的查询")
	return nil
}

func jinianList(pl player.Player) error {
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	mapInfo := manager.GetJiNianMap()
	msg := pbutil.BuildScMarryJiNianList(mapInfo)
	pl.SendMsg(msg)
	return nil
}
