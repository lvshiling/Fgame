package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	worldbosstypes "fgame/fgame/game/worldboss/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_PLAYER_BOSS_RELIVE_SYNC_TYPE), dispatch.HandlerFunc(handlePlayerBossReliveSync))
}

//处理采集完成
func handlePlayerBossReliveSync(s session.Session, msg interface{}) (err error) {
	log.Debug("worldboss:处理跨服boss复活数据完成")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isPlayerBossReliveSync := msg.(*crosspb.ISPlayerBossReliveSync)
	reliveData := isPlayerBossReliveSync.GetPlayerBossReliveData()
	bossType := reliveData.GetBossType()
	reliveTime := reliveData.GetReliveTime()
	err = reliveDataSync(tpl, worldbosstypes.BossType(bossType), reliveTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"bossType":   bossType,
				"reliveTime": reliveTime,
			}).Error("worldboss:处理跨服boss复活数据完成,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"bossType":   bossType,
			"reliveTime": reliveTime,
		}).Debug("worldboss:处理跨服boss复活数据完成,完成")
	return nil

}

//跨服采集完成
func reliveDataSync(pl player.Player, bossType worldbosstypes.BossType, reliveTime int32) (err error) {
	log.Debug("worldboss:处理跨服boss复活数据完成")
	pl.PlayerBossReliveSync(bossType, reliveTime)
	return
}
