package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_PLAYER_ACTIVITY_RANK_DATA_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerActivityRankDataChanged))
}

//处理活动排行榜变化
func handlePlayerActivityRankDataChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("activity:处理跨服排行榜数据变化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isMsg := msg.(*crosspb.ISPlayerActivityRankDataChanged)
	activityType := isMsg.GetActivityType()
	rankType := isMsg.GetRankType()
	val := isMsg.GetVal()

	err = playerActivityRankDataChanged(tpl, activitytypes.ActivityType(activityType), rankType, val)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("activity:处理跨服排行榜数据变化,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("activity:处理跨服排行榜数据变化,完成")
	return nil
}

//跨服采集完成
func playerActivityRankDataChanged(pl player.Player, activityType activitytypes.ActivityType, rankType int32, val int64) (err error) {
	fa := scene.GetActivityRankTypeFactory(activityType)
	if fa == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"activityType": activityType,
			}).Warnln("activity:处理跨服排行榜数据变化,排行榜类型工厂未注册")
		return
	}

	activityRankType := fa.CreateActivityRankType(rankType)
	pl.UpdateActivityRankValue(activityType, activityRankType, val)
	return
}
