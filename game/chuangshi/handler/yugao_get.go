package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/chuangshi/chuangshi"
	"fgame/fgame/game/chuangshi/pbutil"
	playerchuangshi "fgame/fgame/game/chuangshi/player"
	"fgame/fgame/game/player"

	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_YUGAO_TYPE), dispatch.HandlerFunc(handleChuangShiYuGaoGet))

}

//处理玩家创世之战预告消息
func handleChuangShiYuGaoGet(s session.Session, msg interface{}) (err error) {
	log.Debug("chuangshi:处理玩家创世之战预告消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = chuangShiYuGaoGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("chuangshi:处理玩家创世之战预告消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("chuangshi:处理玩家创世之战预告消息,成功")
	return nil

}

func chuangShiYuGaoGet(pl player.Player) (err error) {
	playerChuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
	isJoin := playerChuangShiManager.IsJoin()
	num := chuangshi.GetChuangShiService().GetBaoMingChuangShiPlayerNum()

	scMsg := pbutil.BuildSCChuangShiYuGaoInfo(isJoin, num)
	pl.SendMsg(scMsg)
	return
}
