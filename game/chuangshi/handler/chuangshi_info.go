package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/core/session"
// 	chuangshilogic "fgame/fgame/game/chuangshi/logic"
// 	"fgame/fgame/game/player"

// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_INFO_TYPE), dispatch.HandlerFunc(handleChuangShiInfo))

// }

// //处理玩家创世之战信息
// func handleChuangShiInfo(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理玩家创世之战信息")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)
// 	err = chuangShiInfo(tpl)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理玩家创世之战信息,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理玩家创世之战信息,成功")
// 	return nil
// }

// func chuangShiInfo(pl player.Player) (err error) {
// 	chuangshilogic.SendPlayerChuangShiInfo(pl)
// 	return
// }
