package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	"fgame/fgame/core/session"
// 	"fgame/fgame/game/chuangshi/chuangshi"
// 	"fgame/fgame/game/chuangshi/pbutil"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_GONGCHENG_TARGET_TYPE), dispatch.HandlerFunc(handlerChooesTarget))
// }

// //选择攻城目标请求
// func handlerChooesTarget(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理选择攻城目标请求")

// 	pl := gamesession.SessionInContext(s.Context()).Player()
// 	tpl := pl.(player.Player)
// 	csMsg := msg.(*uipb.CSChuangShiGongChengTarget)
// 	cityId := csMsg.GetCityId()

// 	err = chooseGongChengTarget(tpl, cityId)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": tpl.GetId(),
// 				"err":      err,
// 			}).Error("chuangshi:处理选择攻城目标请求，错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": tpl.GetId(),
// 		}).Debug("chuangshi:处理选择攻城目标请求完成")

// 	return
// }

// func chooseGongChengTarget(pl player.Player, cityId int64) (err error) {

// 	mem := chuangshi.GetChuangShiService().GetMember(pl.GetId())
// 	if mem == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世设置攻城目标,不是阵营成员")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotCamp)
// 		return
// 	}

// 	// 不是神王
// 	if !mem.IfShenWang() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世设置攻城目标,不是神王")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotShenWang)
// 		return
// 	}

// 	// 城池不存在
// 	targetCityData := chuangshi.GetChuangShiService().GetCity(cityId)
// 	if targetCityData == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"cityId":   cityId,
// 			}).Warnln("chuangshi:处理创世设置攻城目标,城池不存在")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiCityNotExist)
// 		return
// 	}

// 	// 不是附属城池
// 	if !targetCityData.IfFuShu() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"CityType": targetCityData.CityType,
// 			}).Warnln("chuangshi:处理创世设置攻城目标,城池不是附属城")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiCityNotFuShu)
// 		return
// 	}

// 	flag := chuangshi.GetChuangShiService().GongChengTargetFuShu(pl.GetId(), cityId)
// 	if !flag {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"cityId":   cityId,
// 			}).Warnln("chuangshi:处理选择攻城目标失败")
// 		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	scMsg := pbutil.BuildSCChuangShiGongChengTarget()
// 	pl.SendMsg(scMsg)
// 	return
// }
