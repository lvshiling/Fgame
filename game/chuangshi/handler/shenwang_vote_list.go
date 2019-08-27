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
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_SHEN_WANG_TOU_PIAO_LIST_TYPE), dispatch.HandlerFunc(handleChuangShiShenWangVoteList))

// }

// //处理玩家创世之战神王投票列表
// func handleChuangShiShenWangVoteList(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理玩家创世之战神王投票列表")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	err = chuangShiShenWangVoteList(tpl)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理玩家创世之战神王投票列表,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理玩家创世之战神王投票列表,成功")
// 	return nil
// }

// func chuangShiShenWangVoteList(pl player.Player) (err error) {
// 	voteList := chuangshi.GetChuangShiService().GetShenWangVoteList(pl.GetId())
// 	scMsg := pbutil.BuildSCChuangShiShengWangTouPiaoList(voteList)
// 	pl.SendMsg(scMsg)
// 	return
// }
