package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	"fgame/fgame/core/session"
// 	"fgame/fgame/game/chuangshi/chuangshi"
// 	chuangshidata "fgame/fgame/game/chuangshi/data"
// 	"fgame/fgame/game/chuangshi/pbutil"
// 	playerchuangshi "fgame/fgame/game/chuangshi/player"
// 	"fgame/fgame/game/common/common"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	playertypes "fgame/fgame/game/player/types"
// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_CAMP_PAY_SCHEDULE_TYPE), dispatch.HandlerFunc(handleChuangShiCampPaySchedule))
// }

// //处理创世阵营工资分配
// func handleChuangShiCampPaySchedule(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理创世阵营工资分配")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	csMsg := msg.(*uipb.CSChuangShiCampPaySchedule)
// 	paramList := chuangshidata.ConvertToCampPayScheduleParamList(csMsg.GetScheduleList())

// 	err = chuangShiCampPaySchedule(tpl, paramList)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理创世阵营工资分配,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理创世阵营工资分配,成功")
// 	return nil
// }

// func chuangShiCampPaySchedule(pl player.Player, paramList []*chuangshidata.CamPayScheduleParam) (err error) {
// 	chuangshiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	chuangshiInfo := chuangshiManager.GetPlayerChuangShiInfo()

// 	// 官职判断
// 	if !chuangshiInfo.IfShenWang() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世阵营工资分配,不是神王")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotShenWang)
// 		return
// 	}

// 	// 城主存在判断
// 	totalRatio := int32(0)
// 	for _, param := range paramList {
// 		totalRatio += param.Ratio
// 	}
// 	if totalRatio != common.MAX_RATE {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":   pl.GetId(),
// 				"totalRatio": totalRatio,
// 			}).Warnln("chuangshi:处理创世阵营工资分配,超过最大分配比例")
// 		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	err = chuangshi.GetChuangShiService().CampPaySchedule(pl.GetId(), paramList)
// 	if err != nil {
// 		return
// 	}

// 	scMsg := pbutil.BuildSCChuangShiCampPaySchedule()
// 	pl.SendMsg(scMsg)
// 	return
// }
