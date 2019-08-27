package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/activity/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_PLAYER_ACTIVITY_PK_DATA_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerActivityPkDataChanged))
}

//处理采集完成
func handlePlayerActivityPkDataChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("activity:处理跨服pk数据变化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isPlayerActivityPkDataChanged := msg.(*crosspb.ISPlayerActivityPkDataChanged)
	killData := pbutil.ConvertToPlayerKillData(isPlayerActivityPkDataChanged.GetActivityPkData())
	err = playerActivityPkDataChanged(tpl, killData)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("activity:处理跨服pk数据变化,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("activity:处理跨服pk数据变化,完成")
	return nil

}

//跨服采集完成
func playerActivityPkDataChanged(pl player.Player, killData *scene.PlayerActvitiyKillData) (err error) {
	pl.SyncKillData(killData)
	siPlayerActivityPkDataChanged := pbutil.BuildSIPlayerActivityPkDataChanged(killData)
	pl.SendCrossMsg(siPlayerActivityPkDataChanged)
	return
}
