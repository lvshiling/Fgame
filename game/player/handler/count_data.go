package handler

import (
	"encoding/json"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"

	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/pbutil"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_PLAYER_COUNT_DATA_TYPE), dispatch.HandlerFunc(handlePlayerCountData))
}

//处理玩家统计数据
func handlePlayerCountData(s session.Session, msg interface{}) (err error) {
	log.Debug("player:处理玩家统计数据")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(player.Player)
	csMsg := msg.(*uipb.CSPlayerCountData)
	countList := csMsg.GetCountList()
	err = playerCountData(pl, countList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"countList": countList,
				"error":     err,
			}).Error("player:处理玩家统计数据, 失败")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"countList": countList,
		}).Debug("player:处理玩家统计数据, 成功")
	return
}

func playerCountData(pl player.Player, countList []*uipb.CountData) (err error) {
	statsData := make(map[string]int64)
	if len(countList) == 0 {
		return
	}
	for _, countData := range countList {
		statsData[countData.GetKey()] = countData.GetValue()
	}
	stasContent, err := json.Marshal(statsData)
	if err != nil {
		return
	}
	gameevent.Emit(playereventtypes.EventTypePlayerStats, pl, string(stasContent))
	scMsg := pbutil.BuildSCPlayerCountData()
	pl.SendMsg(scMsg)
	return
}
