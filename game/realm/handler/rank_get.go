package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/realm/pbutil"
	"fgame/fgame/game/realm/realm"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_REALM_RANK_GET_TYPE), dispatch.HandlerFunc(handleRankRealmGet))
}

//处理天劫塔排行榜信息
func handleRankRealmGet(s session.Session, msg interface{}) (err error) {
	log.Debug("realm:处理获取天劫塔排行榜消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = rankRealmGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("realm:处理获取天劫塔排行榜消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("realm:处理获取天劫塔排行榜消息完成")
	return nil

}

//获取天劫塔排行榜界面信息的逻辑
func rankRealmGet(pl player.Player) (err error) {
	dataList, pos := realm.GetRealmRankService().GetTianJieTaTopThreeAndMyPos(pl.GetId())
	scRealmRankGet := pbutil.BuildSCRealmRankGet(dataList, pos)
	pl.SendMsg(scRealmRankGet)
	return
}
