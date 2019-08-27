package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	"fgame/fgame/core/session"
// 	"fgame/fgame/game/chuangshi/chuangshi"
// 	playerchuangshi "fgame/fgame/game/chuangshi/player"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	playertypes "fgame/fgame/game/player/types"
// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"
// 	"fmt"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_SHEN_WANG_TOU_PIAO_TYPE), dispatch.HandlerFunc(handleChuangShiShenWangVote))
// }

// //处理玩家创世之战神王投票
// func handleChuangShiShenWangVote(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理玩家创世之战神王投票")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	csMsg := msg.(*uipb.CSChuangShiShengWangTouPiao)
// 	supportId := csMsg.GetSupportId()

// 	err = chuangShiShenWangVote(tpl, supportId)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理玩家创世之战神王投票,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理玩家创世之战神王投票,成功")
// 	return nil
// }

// func chuangShiShenWangVote(pl player.Player, supportId int64) (err error) {
// 	voteObj := chuangshi.GetChuangShiService().GetShenWangVote(pl.GetId())
// 	if voteObj != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理玩家创世之战神王投票,正在投票")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiJianSheVoting)
// 		return
// 	}

// 	chuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	playerVoteObj := chuangShiManager.GetPlayerChuangShiVoteInfo()
// 	if playerVoteObj.IfVote() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理玩家创世之战神王投票,已经投票过")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiHadShenWangSignUp)
// 		return
// 	}

// 	flag := chuangShiManager.AttendVote()
// 	if !flag {
// 		panic(fmt.Errorf("神王投票应该成功"))
// 	}

// 	chuangshi.GetChuangShiService().ShenWangVote(pl.GetId(), supportId)
// 	return
// }
