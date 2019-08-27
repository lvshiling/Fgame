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
	"fgame/fgame/game/shenfa/pbutil"
	playershenfa "fgame/fgame/game/shenfa/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENFA_UNLOAD_TYPE), dispatch.HandlerFunc(handleShenfaUnload))
}

//处理身法卸下信息
func handleShenfaUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("shenfa:处理身法卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = shenfaUnload(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("shenfa:处理身法卸下信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenfa:处理身法卸下信息完成")
	return nil

}

//身法卸下的逻辑
func shenfaUnload(pl player.Player) (err error) {
	shenfaManager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	obj := shenfaManager.GetShenfaInfo()
	if obj.ShenfaId == 0 {
		playerlogic.SendSystemMessage(pl, lang.ShenfaUnrealNoExist)
		return
	}

	shenfaManager.Unload()

	scShenfaUnload := pbutil.BuildSCShenfaUnload(0)
	pl.SendMsg(scShenfaUnload)
	return
}
