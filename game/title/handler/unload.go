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
	"fgame/fgame/game/title/pbutil"
	playertitle "fgame/fgame/game/title/player"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_TITLE_UNLOAD_TYPE), dispatch.HandlerFunc(handleTitleUnload))

}

//处理称号卸下信息
func handleTitleUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("title:处理称号卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = titleUnload(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("title:处理称号卸下信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("title:处理称号卸下信息完成")
	return nil

}

//称号卸下逻辑
func titleUnload(pl player.Player) (err error) {
	titleManager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	titleWear := titleManager.GetTitleId()
	if titleWear == 0 {
		playerlogic.SendSystemMessage(pl, lang.TitleNotHas)
		return
	}

	titleManager.TitleNoWear()
	scTitleUnload := pbutil.BuildSCTitleUnload(titleWear)
	pl.SendMsg(scTitleUnload)
	return
}
