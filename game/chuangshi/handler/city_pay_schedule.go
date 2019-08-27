package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	"fgame/fgame/core/session"
// 	"fgame/fgame/game/alliance/alliance"
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
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_CITY_PAY_SCHEDULE_TYPE), dispatch.HandlerFunc(handleChuangShiCityPaySchedule))
// }

// //处理创世城主工资分配
// func handleChuangShiCityPaySchedule(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理创世城主工资分配")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	csMsg := msg.(*uipb.CSChuangShiCityPaySchedule)
// 	paramList := chuangshidata.ConvertToCityPayScheduleParamList(csMsg.GetScheduleList())

// 	err = chuangShiCityPaySchedule(tpl, paramList)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理创世城主工资分配,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理创世城主工资分配,成功")
// 	return nil
// }

// func chuangShiCityPaySchedule(pl player.Player, paramList []*chuangshidata.CityPayScheduleParam) (err error) {
// 	chuangshiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	chuangshiInfo := chuangshiManager.GetPlayerChuangShiInfo()

// 	// 官职判断
// 	if !chuangshiInfo.IfChengZhu() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世城主工资分配,不是城主")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotShenWang)
// 		return
// 	}

// 	al := alliance.GetAllianceService().GetAlliance(pl.GetAllianceId())
// 	totalRatio := int32(0)
// 	for _, param := range paramList {
// 		posCount := al.GetManagerNum(param.AlPos)
// 		totalRatio += param.Ratio * posCount
// 	}
// 	if totalRatio > common.MAX_RATE {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":   pl.GetId(),
// 				"totalRatio": totalRatio,
// 			}).Warnln("chuangshi:处理创世阵营工资分配,超过最大分配比例")
// 		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	err = chuangshi.GetChuangShiService().CityPaySchedule(pl.GetId(), paramList)
// 	if err != nil {
// 		return
// 	}

// 	scMsg := pbutil.BuildSCChuangShiCityPaySchedule()
// 	pl.SendMsg(scMsg)
// 	return
// }
