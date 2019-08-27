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
	ringtypes "fgame/fgame/game/ring/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RING_BAOKU_INFO_TYPE), dispatch.HandlerFunc(handleBaoKuInfoGet))
}

func handleBaoKuInfoGet(s session.Session, msg interface{}) (err error) {
	log.Debug("ring: 开始处理特戒宝库信息请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = baoKuInfoGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("ring: 处理特戒宝库信息请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("ring: 处理特戒宝库信息请求消息,成功")

	return
}

func baoKuInfoGet(pl player.Player) (err error) {
	ringManager := pl.GetPlayerDataManager(playertypes.PlayerRingDataManagerType).(*playerring.PlayerRingDataManager)
	ringObj := ringManager.GetPlayerBaoKuObject(ringtypes.BaoKuTypeRing)
	scRingBaoKuInfo := pbutil.BuildSCRingBaoKuInfo(ringObj)
	pl.SendMsg(scRingBaoKuInfo)
	return
}
