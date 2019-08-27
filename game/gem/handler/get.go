package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/gem/pbutil"
	playergem "fgame/fgame/game/gem/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GEM_MINE_GET_TYPE), dispatch.HandlerFunc(handleGemMineGet))
}

//处理获取挖矿信息
func handleGemMineGet(s session.Session, msg interface{}) (err error) {
	log.Debug("gem:处理获取挖矿消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = gemMineGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("gem:处理获取挖矿消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("gem:处理获取挖矿消息完成")
	return nil

}

//获取挖矿界面信息的逻辑
func gemMineGet(pl player.Player) (err error) {
	gemManager := pl.GetPlayerDataManager(types.PlayerGemDataManagerType).(*playergem.PlayerGemDataManager)
	mine := gemManager.GetMine()
	scGemMineGet := pbutil.BuildSCGemMineGet(mine)
	pl.SendMsg(scGemMineGet)
	return
}
