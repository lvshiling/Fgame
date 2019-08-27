package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arenapvp/arenapvp"
	"fgame/fgame/game/arenapvp/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENAPVP_BAZHU_INFO_TYPE), dispatch.HandlerFunc(handleArenapvpBaZhu))
}

//处理历届霸主
func handleArenapvpBaZhu(s session.Session, msg interface{}) (err error) {
	log.Debug("arenapvp:处理历届霸主")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSArenapvpBaZhuInfo)
	page := csMsg.GetPage()
	pageNum := csMsg.GetPageNum()


	err = arenapvpBaZhu(tpl, page, pageNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arenapvp:处理历届霸主,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arenapvp:处理历届霸主,完成")
	return nil
}

func arenapvpBaZhu(pl player.Player, page, pageNum int32) (err error) {
	baZhuInfoList, totalPage := arenapvp.GetArenapvpService().GetAreanapvpBaZhuList(page, pageNum)
	scMsg := pbutil.BuildSCArenapvpBaZhuInfo(page, pageNum, totalPage, baZhuInfoList)
	pl.SendMsg(scMsg)

	return
}
