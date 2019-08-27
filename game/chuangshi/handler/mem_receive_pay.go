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
// 	chuangshitemplate "fgame/fgame/game/chuangshi/template"
// 	"fgame/fgame/game/global"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	playertypes "fgame/fgame/game/player/types"
// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"
// 	"fmt"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_MY_PAY_RECEIVE_TYPE), dispatch.HandlerFunc(handleChuangShiMyPayReceive))
// }

// //处理创世阵营领取个人工资
// func handleChuangShiMyPayReceive(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理创世阵营领取个人工资")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	// csMsg := msg.(*uipb.CSChuangShiMyPayReceive)

// 	err = chuangShiMyPayReceive(tpl)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理创世阵营领取个人工资,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理创世阵营领取个人工资,成功")
// 	return nil
// }

// func chuangShiMyPayReceive(pl player.Player) (err error) {
// 	chuangshiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	chuangshiInfo := chuangshiManager.GetPlayerChuangShiInfo()

// 	now := global.GetGame().GetTimeService().Now()
// 	constantTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangshiConstantTemp()
// 	rewCount, _ := constantTemp.RewCout(chuangshiInfo.GetLastMyPayTime(), now)
// 	if rewCount <= 0 {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世阵营领取个人工资,没有可领取的奖励")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotMyRew)
// 		return
// 	}

// 	camp := chuangshi.GetChuangShiService().GetCamp(pl.GetId())
// 	if camp == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世阵营领取个人工资,没有阵营")
// 		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	flag := chuangshiManager.ReceiveMyPay(camp.CityList)
// 	if !flag {
// 		panic(fmt.Errorf("chuangshi:领取个人工资应该成功"))
// 	}

// 	scMsg := pbutil.BuildSCChuangShiMyPayReceive()
// 	pl.SendMsg(scMsg)
// 	return
// }
