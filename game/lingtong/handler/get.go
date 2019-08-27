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
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONG_GET_TYPE), dispatch.HandlerFunc(handleLingTongGet))

}

//处理灵童信息
func handleLingTongGet(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtong:处理获取灵童消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = lingTongGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("lingtong:处理获取灵童消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lingtong:处理获取灵童消息完成")
	return nil

}

//获取灵童界面信息逻辑
func lingTongGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongObj := manager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("lingtong:请先激活灵童系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongActiveSystem)
		return
	}
	lingTongId := lingTongObj.GetLingTongId()
	lingTongMap := manager.GetLingTongMap()
	lingTongFashion := manager.GetLingTongFashion()

	scLingTongGet := pbutil.BuildSCLingTongGet(lingTongId, lingTongMap, lingTongFashion)
	pl.SendMsg(scLingTongGet)
	return
}
