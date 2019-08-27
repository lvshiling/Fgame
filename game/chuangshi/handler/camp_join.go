package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"

	// chuangshilogic "fgame/fgame/game/chuangshi/logic"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_JOIN_CAMP_TYPE), dispatch.HandlerFunc(handlerJoinCamp))
}

//选择加入阵营请求
func handlerJoinCamp(s session.Session, msg interface{}) (err error) {
	log.Debug("chuangshi:处理选择加入阵营请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSChuangShiJoinCamp)
	campTypeInt := csMsg.GetCamp()

	campType := chuangshitypes.ChuangShiCampType(campTypeInt)
	if !campType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"campType": campType,
			}).Warnln("chuangshi:处理创世城池建设,阵营类型参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = joinCamp(tpl, campType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("chuangshi:处理选择加入阵营请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("chuangshi:处理选择加入阵营请求完成")

	return
}

func joinCamp(pl player.Player, campType chuangshitypes.ChuangShiCampType) (err error) {
	// success := chuangshilogic.PlayerJoinCamp(pl, campType)
	// if !success {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 		}).Warnln("chuangshi:处理创世加入阵营失败")
	// 	playerlogic.SendSystemMessage(pl, lang.ChuangShiJoinCampFailed)
	// 	return
	// }

	return
}
