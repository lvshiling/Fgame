package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	"fgame/fgame/core/session"
// 	"fgame/fgame/cross/arenapvp/arenapvp"
// 	"fgame/fgame/cross/arenapvp/pbutil"
// 	"fgame/fgame/cross/player/player"
// 	"fgame/fgame/cross/processor"
// 	arenapvptypes "fgame/fgame/game/arenapvp/types"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	gamesession "fgame/fgame/game/session"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_SELECT_FOUR_GOD_TYPE), dispatch.HandlerFunc(handleArenapvpSelectFourGod))
// }

// //处理选择四圣兽
// func handleArenapvpSelectFourGod(s session.Session, msg interface{}) (err error) {
// 	log.Debug("arenapvp:处理选择四圣兽")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(*player.Player)
// 	csArenapvpSelectFourGod := msg.(*uipb.CSArenapvpSelectFourGod)
// 	fourGodTypeInt := csArenapvpSelectFourGod.GetFourGodType()
// 	fourGodType := arenapvptypes.FourGodType(fourGodTypeInt)
// 	if !fourGodType.Valid() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Error("arenapvp:处理选择四圣兽,参数错误")
// 		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	err = arenapvpSelectFourGod(tpl, fourGodType)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Error("arenapvp:处理选择四圣兽,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("arenapvp:处理选择四圣兽")
// 	return nil

// }

// //3v3匹配
// func arenapvpSelectFourGod(pl *player.Player, fourGodType arenapvptypes.FourGodType) (err error) {
// 	arenapvp.GetArenapvpService().PlayerEnterFourGod(pl.GetId(), fourGodType)
// 	scArenapvpSelectFourGod := pbutil.BuildSCArenapvpSelectFourGod(fourGodType)
// 	pl.SendMsg(scArenapvpSelectFourGod)
// 	return
// }
