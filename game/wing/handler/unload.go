package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/wing/pbutil"
	playerwing "fgame/fgame/game/wing/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WING_UNLOAD_TYPE), dispatch.HandlerFunc(handleWingUnload))
}

//处理战翼卸下信息
func handleWingUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理战翼卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = wingUnload(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("wing:处理战翼卸下信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("wing:处理战翼卸下信息完成")
	return nil

}

//战翼卸下的逻辑
func wingUnload(pl player.Player) (err error) {
	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	obj := wingManager.GetWingInfo()
	if obj.WingId == 0 {
		playerlogic.SendSystemMessage(pl, lang.WingUnrealNoExist)
		return
	}
	wingManager.Unload()
	scWingUnload := pbutil.BuildSCWingUnload(0)
	pl.SendMsg(scWingUnload)
	return
}
