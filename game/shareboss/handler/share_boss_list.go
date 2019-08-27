package handler

//废弃:zrc
// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	"fgame/fgame/core/session"
// 	funcopentypes "fgame/fgame/game/funcopen/types"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"
// 	"fgame/fgame/game/shareboss/pbutil"
// 	"fgame/fgame/game/shareboss/shareboss"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_SHARE_BOSS_LIST_TYPE), dispatch.HandlerFunc(handlerShareBossList))
// }

// //跨服世界BOSS列表请求
// func handlerShareBossList(s session.Session, msg interface{}) (err error) {
// 	log.Debug("shareBoss:处理跨服世界BOSS列表请求")

// 	pl := gamesession.SessionInContext(s.Context()).Player()
// 	tpl := pl.(player.Player)

// 	err = shareBossList(tpl)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": tpl.GetId(),
// 				"err":      err,
// 			}).Error("shareBoss:处理跨服世界BOSS列表请求，错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": tpl.GetId(),
// 		}).Debug("shareBoss:处理跨服世界BOSS列表请求完成")

// 	return
// }

// func shareBossList(pl player.Player) (err error) {
// 	//功能开启判断
// 	flag := pl.IsFuncOpen(funcopentypes.FuncOpenTypeCrossWorldBoss)
// 	if !flag {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warn("shareBoss:跨服跨服世界BOSS列表错误，功能未开启")

// 		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
// 		return
// 	}

// 	bossList := shareboss.GetShareBossService().GetShareBossList()
// 	scShareBossList := pbutil.BuildSCShareBossList(bossList)
// 	pl.SendMsg(scShareBossList)

// 	return
// }
