package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/core/session"
// 	"fgame/fgame/game/chuangshi/chuangshi"
// 	"fgame/fgame/game/chuangshi/pbutil"
// 	"fgame/fgame/game/player"

// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_SHEN_WANG_BAO_MING_LIST_TYPE), dispatch.HandlerFunc(handleChuangShiShenWangBaoMingList))

// }

// //处理玩家创世之战神王报名列表
// func handleChuangShiShenWangBaoMingList(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理玩家创世之战神王报名列表")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	err = chuangShiShenWangBaoMingList(tpl)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理玩家创世之战神王报名列表,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理玩家创世之战神王报名列表,成功")
// 	return nil
// }

// func chuangShiShenWangBaoMingList(pl player.Player) (err error) {
// 	signList := chuangshi.GetChuangShiService().GetShenWangSignUpList(pl.GetId())
// 	scMsg := pbutil.BuildSCChuangShiShengWangBaoMingList(signList)
// 	pl.SendMsg(scMsg)
// 	return
// }
