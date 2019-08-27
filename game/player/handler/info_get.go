package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"

	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/pbutil"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_PLAYER_INFO_GET_TYPE), dispatch.HandlerFunc(handlePlayerInfoGet))
}

//处理玩家信息获取
func handlePlayerInfoGet(s session.Session, msg interface{}) (err error) {
	log.Debug("player:处理玩家信息获取")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(player.Player)
	csPlayerInfoGet := msg.(*uipb.CSPlayerInfoGet)
	getPlayerId := csPlayerInfoGet.GetPlayerId()
	err = playerInfoGet(pl, getPlayerId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"getPlayerId": getPlayerId,
				"error":       err,
			}).Error("player:处理玩家信息获取,创建失败")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":    pl.GetId(),
			"getPlayerId": getPlayerId,
		}).Debug("player:处理玩家信息获取,创建成功")
	return
}

func playerInfoGet(pl player.Player, getPlayerId int64) (err error) {
	info, err := player.GetPlayerService().GetPlayerInfo(getPlayerId)
	if err != nil {
		return
	}
	if info == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"getPlayerId": getPlayerId,
			}).Warn("player:处理玩家信息获取,玩家不存在")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoExist)
		return
	}
	scPlayerInfoGet := pbutil.BuildSCPlayerInfoGet(info)
	pl.SendMsg(scPlayerInfoGet)
	return
}
