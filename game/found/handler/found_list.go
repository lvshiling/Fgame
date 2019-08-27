package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/found/pbutil"
	playerfound "fgame/fgame/game/found/player"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FOUND_RESOUCE_LIST_TYPE), dispatch.HandlerFunc(handlerFoundResList))
}

//资源找回列表
func handlerFoundResList(s session.Session, msg interface{}) (err error) {
	log.Debug("found:处理资源找回列表获取请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	err = foundResList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("found:处理资源找回列表获取请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("found：处理资源找回列表获取请求完成")

	return
}

func foundResList(pl player.Player) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeResFind) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("found：处理资源找回列表,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	//获取信息List
	foundManager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)
	resList := foundManager.GetPreDayFoundList()

	scFoundResouceList := pbutil.BuildSCFoundResouceList(resList)
	pl.SendMsg(scFoundResouceList)
	return
}
