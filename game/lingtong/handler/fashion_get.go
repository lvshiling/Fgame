package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	playerlingtong "fgame/fgame/game/lingtong/player"

	"fgame/fgame/game/lingtong/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONG_FASHION_GET_TYPE), dispatch.HandlerFunc(handleLingTongFashionGet))

}

//处理灵童时装信息
func handleLingTongFashionGet(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtong:处理获取灵童时装消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = lingTongFashionGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("lingtong:处理获取灵童时装消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lingtong:处理获取灵童时装消息完成")
	return nil
}

//获取灵童时装界面信息逻辑
func lingTongFashionGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongObj := manager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("lingtong:请先激活灵童时装系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongActiveSystem)
		return
	}

	lingTongId := lingTongObj.GetLingTongId()
	fashionMap := manager.GetActivateFashionMap()
	scLingTongFashionGet := pbutil.BuildSCLingTongFashionGet(lingTongId, fashionMap)
	pl.SendMsg(scLingTongFashionGet)
	return
}
