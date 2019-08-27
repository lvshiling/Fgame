package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_INFO_TYPE), dispatch.HandlerFunc(handleBabyInfo))
}

//处理宝宝信息
func handleBabyInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理宝宝信息消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = handlerBabyInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("baby:处理宝宝信息消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("baby:处理宝宝信息消息完成")
	return nil

}

// 信息
func handlerBabyInfo(pl player.Player) (err error) {
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	babyList := babyManager.GetBabyInfoList()
	pregnantInfo := babyManager.GetPregnantInfo()
	allToySlotMap := babyManager.GetAllToySlotMap()
	scMsg := pbutil.BuildSCBabyInfo(pregnantInfo, babyList, allToySlotMap)
	pl.SendMsg(scMsg)
	return
}
