package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/house/house"
	"fgame/fgame/game/house/pbutil"
	playerhouse "fgame/fgame/game/house/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_HOUSE_LIST_GET_TYPE), dispatch.HandlerFunc(handleHouseGet))
}

//处理房子信息
func handleHouseGet(s session.Session, msg interface{}) (err error) {
	log.Debug("house:处理获取房子消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = houseGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("house:处理获取房子消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("house:处理获取房子消息完成")
	return nil
}

//获取房子信息
func houseGet(pl player.Player) (err error) {
	houseManager := pl.GetPlayerDataManager(playertypes.PlayerHouseDataManagerType).(*playerhouse.PlayerHouseDataManager)
	houseMap := houseManager.GetHouseMap()
	logList := house.GetHouseService().GetLogByTime(0)
	scMsg := pbutil.BuildSCHouseListGet(houseMap, logList)
	pl.SendMsg(scMsg)
	return
}
