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
	"fgame/fgame/game/shihunfan/pbutil"
	playershihunfan "fgame/fgame/game/shihunfan/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHIHUNFAN_GET_TYPE), dispatch.HandlerFunc(handleShiHunFanGet))
}

//处理噬魂幡系统信息
func handleShiHunFanGet(s session.Session, msg interface{}) (err error) {
	log.Debug("shihunfan:处理获取噬魂幡系统消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = shiHunFanGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("shihunfan:处理获取噬魂幡系统消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shihunfan:处理获取噬魂幡系统消息完成")
	return nil
}

//获取噬魂幡系统信息
func shiHunFanGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	shiHunFanObject := manager.GetShiHunFanInfo()
	scShiHunFanGet := pbutil.BuildSCShiHunFanGet(shiHunFanObject)
	pl.SendMsg(scShiHunFanGet)
	return
}
