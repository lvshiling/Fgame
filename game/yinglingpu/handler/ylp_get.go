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
	ylppbutil "fgame/fgame/game/yinglingpu/pbutil"
	ylpplayer "fgame/fgame/game/yinglingpu/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_YLPU_QUERY_TYPE), dispatch.HandlerFunc(handleYlpGet))
}

//英灵普获取
func handleYlpGet(s session.Session, msg interface{}) (err error) {
	log.Debug("yinglingpu:获取英灵普信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = getYlp(tpl)
	if err != nil {
		log.WithFields(log.Fields{
			"playerId": tpl.GetId(),
			"error":    err,
		}).Error("获取英灵普信息失败")
		//需要返回
		return
	}
	log.Debug("yinglingpu:获取信息完成")
	return
}

func getYlp(pl player.Player) (err error) {
	ylpManager := pl.GetPlayerDataManager(playertypes.PlayerYingLingPuManagerType).(*ylpplayer.PlayerYingLingPuManager)
	ylpList := ylpManager.GetAllYingLingPu()
	ylpSpList := ylpManager.GetAllYingLingPuSuiPian()
	sendInfo := ylppbutil.BuildYlpInfo(ylpList, ylpSpList)
	pl.SendMsg(sendInfo)
	return
}
