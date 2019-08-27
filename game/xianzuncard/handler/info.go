package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/xianzuncard/pbutil"
	playerxianzuncard "fgame/fgame/game/xianzuncard/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIAN_ZUN_CARD_INFO_TYPE), dispatch.HandlerFunc(handleXianZunCardInfo))
}

func handleXianZunCardInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("xianzuncard: 开始处理仙尊特权卡请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = getXianZunCardInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("xianzuncard: 处理仙尊特权卡请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("xianzuncard: 处理仙尊特权卡请求消息,成功")

	return
}

func getXianZunCardInfo(pl player.Player) (err error) {
	xianZunManager := pl.GetPlayerDataManager(playertypes.PlayerXianZunCardManagerType).(*playerxianzuncard.PlayerXianZunCardDataManager)
	xianZunObjMap := xianZunManager.GetXianZunCardObjectMap()
	scXianZunCardInfo := pbutil.BuildSCXianZunCardInfo(xianZunObjMap)
	pl.SendMsg(scXianZunCardInfo)
	return
}
