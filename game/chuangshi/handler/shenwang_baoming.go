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
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_SHEN_WANG_BAO_MING_TYPE), dispatch.HandlerFunc(handleChuangShiShenWangBaoMing))
// }

// //处理玩家创世之战神王报名
// func handleChuangShiShenWangBaoMing(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理玩家创世之战神王报名")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	err = chuangShiShenWangBaoMing(tpl)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理玩家创世之战神王报名,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理玩家创世之战神王报名,成功")
// 	return nil
// }

// func chuangShiShenWangBaoMing(pl player.Player) (err error) {
// 	signObj := chuangshi.GetChuangShiService().GetShenWangSignUp(pl.GetId())
// 	if signObj != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理玩家创世之战神王报名,正在报名")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiJianSheSigning)
// 		return
// 	}

// 	chuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	playerSignObj := chuangShiManager.GetPlayerChuangShiSignInfo()
// 	if playerSignObj.IfShenWangSignUp() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理玩家创世之战神王报名,已经报名过")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiHadShenWangSignUp)
// 		return
// 	}

// 	flag := chuangShiManager.ShenWangSignUp()
// 	if !flag {
// 		panic(fmt.Errorf("神王报名应该成功"))
// 	}

// 	chuangshi.GetChuangShiService().ShenWangSignUp(pl.GetId())
// 	return
// }
