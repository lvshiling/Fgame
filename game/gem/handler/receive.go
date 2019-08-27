package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/gem/pbutil"
	playergem "fgame/fgame/game/gem/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GEM_MINE_RECEIVE_TYPE), dispatch.HandlerFunc(handleGemMineReceive))
}

//处理领取收益信息
func handleGemMineReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("gem:处理领取收益信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = gemMineReceive(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("gem:处理领取收益信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("gem:处理领取收益信息完成")
	return nil
}

//处理领取收益信息逻辑
func gemMineReceive(pl player.Player) (err error) {
	gemManager := pl.GetPlayerDataManager(types.PlayerGemDataManagerType).(*playergem.PlayerGemDataManager)
	//获取当前库存
	curStorage := gemManager.GetMineStorage()
	if curStorage <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("gem:当前无原石可供领取,稍后再来")
		playerlogic.SendSystemMessage(pl, lang.GemNotReceiveStone)
		return
	}
	//领取收益
	mine := gemManager.Receive()
	scGemMineGet := pbutil.BuildSCGemMineGet(mine)
	//飘字使用统一发get
	pl.SendMsg(scGemMineGet)

	// stone := gemManager.GetMineStone()
	// lastTime := gemManager.GetMineLastTime()
	// scGemMineReceive := pbutil.BuildSCGemMineReceive(stone, lastTime)
	// pl.SendMsg(scGemMineReceive)
	return
}
