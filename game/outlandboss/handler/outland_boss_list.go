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
	"fgame/fgame/game/outlandboss/pbutil"
	playeroutlandboss "fgame/fgame/game/outlandboss/player"
	"fgame/fgame/game/outlandboss/outlandboss"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OUTLAND_BOSS_LIST_TYPE), dispatch.HandlerFunc(handlerOutlandBossList))
}

//外域BOSS列表请求
func handlerOutlandBossList(s session.Session, msg interface{}) (err error) {
	log.Debug("outlandboss:处理外域BOSS列表请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = outlandBossList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("outlandboss:处理外域BOSS列表请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("outlandboss:处理外域BOSS列表请求完成")

	return
}

func outlandBossList(pl player.Player) (err error) {
	bossList := outlandboss.GetOutlandBossService().GetOutlandBossList()
	outlandbossManager := pl.GetPlayerDataManager(playertypes.PlayerOutlandBossDataManagerType).(*playeroutlandboss.PlayerOutlandBossDataManager)
	outlandbossManager.RefreshZhuoQi()

	curZhuoQi := outlandbossManager.GetCurZhuoQiNum()
	scMsg := pbutil.BuildSCOutlandBossList(bossList, curZhuoQi)
	pl.SendMsg(scMsg)
	return
}
