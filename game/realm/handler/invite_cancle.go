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
	processor.Register(codec.MessageType(uipb.MessageType_CS_REALM_PAIR_CANCLE_TYPE), dispatch.HandlerFunc(handleRealmPairCancle))
}

//处理取消夫妻助战邀请信息
func handleRealmPairCancle(s session.Session, msg interface{}) (err error) {
	log.Debug("realm:处理取消夫妻助战邀请消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = realmPairCancle(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("realm:处理取消夫妻助战邀请消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("realm:处理取消夫妻助战邀请消息完成")
	return nil
}

//处理取消夫妻助战邀请信息逻辑
func realmPairCancle(pl player.Player) (err error) {
	codeResult := realm.GetRealmRankService().CanclePairInvite(pl)
	scRealmPairCancle := pbutil.BuildSCRealmPairCancle(int32(codeResult))
	pl.SendMsg(scRealmPairCancle)
	return
}
