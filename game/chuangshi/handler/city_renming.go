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
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_CITY_REN_MING), dispatch.HandlerFunc(handleChuangShiCityRenMing))
// }

// //处理创世城主任命
// func handleChuangShiCityRenMing(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理创世城主任命")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	csMsg := msg.(*uipb.CSChuangShiCityRenMing)
// 	cityId := csMsg.GetCityId()
// 	// campTypeInt := csMsg.GetOrignalCampType()
// 	// cityTypeInt := csMsg.GetCityType()
// 	// index := csMsg.GetIndex()
// 	beCommitId := csMsg.GetBeCommitId()

// 	// campType := chuangshitypes.ChuangShiCampType(campTypeInt)
// 	// if !campType.Valid() {
// 	// 	log.WithFields(
// 	// 		log.Fields{
// 	// 			"playerId": pl.GetId(),
// 	// 			"campType": campType,
// 	// 		}).Warnln("chuangshi:处理创世城主任命,参数错误")
// 	// 	playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
// 	// 	return
// 	// }

// 	// cityType := chuangshitypes.ChuangShiCityType(cityTypeInt)
// 	// if !cityType.Valid() {
// 	// 	log.WithFields(
// 	// 		log.Fields{
// 	// 			"playerId": pl.GetId(),
// 	// 			"cityType": cityType,
// 	// 		}).Warnln("chuangshi:处理创世城主任命,参数错误")
// 	// 	playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
// 	// 	return
// 	// }

// 	err = chuangShiCityRenMing(tpl, cityId, beCommitId)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理创世城主任命,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理创世城主任命,成功")
// 	return nil
// }

// func chuangShiCityRenMing(pl player.Player, cityId, beCommitId int64) (err error) {
// 	mem := chuangshi.GetChuangShiService().GetMember(pl.GetId())
// 	if mem == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世设置攻城目标,不是阵营成员")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotCamp)
// 		return
// 	}

// 	// 官职判断
// 	if !mem.IfShenWang() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世城主任命,不是神王")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotShenWang)
// 		return
// 	}

// 	commitCity := chuangshi.GetChuangShiService().GetCity(cityId)
// 	if commitCity == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世城主任命,城池不存在")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiCityNotExist)
// 		return
// 	}
// 	if !commitCity.IfFuShu() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世城主任命,不是附属城")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiCityNotFuShu)
// 	}

// 	beCommitMem := chuangshi.GetChuangShiService().GetMember(beCommitId)
// 	if beCommitMem == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世设置攻城目标,不是阵营成员")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotCamp)
// 		return
// 	}

// 	if !beCommitMem.IfMengZhu() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世设置攻城目标,不是仙盟盟主")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotMengZhu)
// 		return
// 	}

// 	//已经是城主
// 	if beCommitMem.IfChengZhu() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世设置攻城目标,已经是城主")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiHasChengZhu)
// 		return
// 	}

// 	success, err := chuangshi.GetChuangShiService().CityRenMing(pl.GetId(), cityId, beCommitId)
// 	if err != nil {
// 		return
// 	}
// 	if !success {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"cityId":   cityId,
// 			}).Warnln("chuangshi:处理创世城主任命，任命失败")
// 		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	scMsg := pbutil.BuildSCChuangShiCityRenMing()
// 	pl.SendMsg(scMsg)
// 	return
// }
