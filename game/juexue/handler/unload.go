package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/juexue/pbutil"
	playerjx "fgame/fgame/game/juexue/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JUEXUE_UNLOAD_TYPE), dispatch.HandlerFunc(handleJueXueUnload))
}

//处理绝学卸下信息
func handleJueXueUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("juexue:处理绝学卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = juexueUnload(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("juexue:处理绝学卸下信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("juexue:处理绝学卸下信息完成")
	return nil

}

//绝学卸下的逻辑
func juexueUnload(pl player.Player) (err error) {
	jxManager := pl.GetPlayerDataManager(types.PlayerJueXueDataManagerType).(*playerjx.PlayerJueXueDataManager)
	useId := jxManager.GetJueXueUseId()
	if useId == 0 {
		playerlogic.SendSystemMessage(pl, lang.JueXueUseNoExist)
		return
	}

	flag := jxManager.Unload()
	if !flag {
		panic(fmt.Errorf("juexue: juexueUnload should be ok"))
	}
	scJueXueUnload := pbutil.BuildSCJueXueUnload(0)
	pl.SendMsg(scJueXueUnload)
	return
}
