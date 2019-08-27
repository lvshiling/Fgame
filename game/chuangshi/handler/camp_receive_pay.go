package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	"fgame/fgame/core/session"
// 	"fgame/fgame/game/chuangshi/chuangshi"
// 	"fgame/fgame/game/chuangshi/pbutil"
// 	playerchuangshi "fgame/fgame/game/chuangshi/player"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	playertypes "fgame/fgame/game/player/types"
// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_CAMP_PAY_RECEIVE_TYPE), dispatch.HandlerFunc(handleChuangShiCampPayReceive))
// }

// //处理创世阵营工资领取
// func handleChuangShiCampPayReceive(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理创世阵营工资领取")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	err = chuangShiCampPayReceive(tpl)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理创世阵营工资领取,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理创世阵营工资领取,成功")
// 	return nil
// }

// func chuangShiCampPayReceive(pl player.Player) (err error) {
// 	chuangshiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	chuangshiInfo := chuangshiManager.GetPlayerChuangShiInfo()

// 	// 官职判断
// 	if !chuangshiInfo.IfShenWang() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世阵营工资领取,不是神王")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotShenWang)
// 		return
// 	}

// 	err = chuangshi.GetChuangShiService().CampPayReceive(pl.GetId())
// 	if err != nil {
// 		return
// 	}

// 	scMsg := pbutil.BuildSCChuangShiCampPayReceive()
// 	pl.SendMsg(scMsg)
// 	return
// }
