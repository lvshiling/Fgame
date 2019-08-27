package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/lingyu/pbutil"
	playerlingyu "fgame/fgame/game/lingyu/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGYU_UNLOAD_TYPE), dispatch.HandlerFunc(handleLingyuUnload))
}

//处理领域卸下信息
func handleLingyuUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("lingyu:处理领域卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = lingyuUnload(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("lingyu:处理领域卸下信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lingyu:处理领域卸下信息完成")
	return nil

}

//领域卸下的逻辑
func lingyuUnload(pl player.Player) (err error) {
	lingyuManager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	obj := lingyuManager.GetLingyuInfo()
	if obj.LingyuId == 0 {
		playerlogic.SendSystemMessage(pl, lang.LingyuUnrealNoExist)
		return
	}
	lingyuManager.Unload()

	scLingyuUnload := pbutil.BuildSCLingyuUnload(0)
	pl.SendMsg(scLingyuUnload)
	return
}
