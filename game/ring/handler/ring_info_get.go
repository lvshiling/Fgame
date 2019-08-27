package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/ring/pbutil"
	playerring "fgame/fgame/game/ring/player"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RING_INFO_GET_TYPE), dispatch.HandlerFunc(handleRingInfoGet))
}

func handleRingInfoGet(s session.Session, msg interface{}) (err error) {
	log.Debug("ring: 开始处理特戒信息请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = ringInfoGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("ring: 处理特戒信息请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("ring: 处理特戒信息请求消息,成功")

	return
}

func ringInfoGet(pl player.Player) (err error) {
	ringManager := pl.GetPlayerDataManager(playertypes.PlayerRingDataManagerType).(*playerring.PlayerRingDataManager)
	ringObjMap := ringManager.GetPlayerRingObjectMap()
	scRingInfoGet := pbutil.BuildSCRingInfoGet(ringObjMap)
	pl.SendMsg(scRingInfoGet)
	return
}
