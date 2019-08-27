package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/lingyu/pbutil"
	playerlingyu "fgame/fgame/game/lingyu/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGYU_GET_TYPE), dispatch.HandlerFunc(handleLingyuGet))

}

//处理领域信息
func handleLingyuGet(s session.Session, msg interface{}) (err error) {
	log.Debug("Lingyu:处理获取领域消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = lingyuGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("Lingyu:处理获取领域消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("Lingyu:处理获取领域消息完成")
	return nil

}

//获取领域界面信息逻辑
func lingyuGet(pl player.Player) (err error) {
	lingyuManager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingyuInfo := lingyuManager.GetLingyuInfo()
	lingYuOtherMap := lingyuManager.GetLingyuOtherMap()
	scLingyuGet := pbutil.BuildSCLingyuGet(lingyuInfo, lingYuOtherMap)
	pl.SendMsg(scLingyuGet)
	return
}
