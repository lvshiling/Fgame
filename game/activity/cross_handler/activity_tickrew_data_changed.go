package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_PLAYER_ACTIVITY_TICKREW_DATA_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerActivityTickRewDataChanged))
}

//处理活动定时奖励变化
func handlePlayerActivityTickRewDataChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("activity:处理跨服活动定时奖励数据变化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isMsg := msg.(*crosspb.ISPlayerActivityTickRewDataChanged)

	resMap := make(map[int32]int32)
	for _, resData := range isMsg.GetPropertyList() {
		resMap[resData.GetKey()] += int32(resData.GetValue())
	}

	specialResMap := make(map[int32]int32)
	for _, resData := range isMsg.GetSpecialPropertyList() {
		specialResMap[resData.GetKey()] += int32(resData.GetValue())
	}

	err = playerActivityTickRewDataChanged(tpl, resMap, specialResMap)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("activity:处理跨服活动定时奖励数据变化,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("activity:处理跨服活动定时奖励数据变化,完成")
	return nil
}

//跨服完成
func playerActivityTickRewDataChanged(pl player.Player, resMap, specialResMap map[int32]int32) (err error) {
	pl.AddActivityTickRew(resMap, specialResMap)
	return
}
