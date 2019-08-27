package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	pbutil "fgame/fgame/cross/player/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_SYSTEM_BATTLE_PROPERTY_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerSystemBattlePropertyChanged))
}

//处理跨服系统属性推送
func handlePlayerSystemBattlePropertyChanged(s session.Session, msg interface{}) error {
	log.Debug("login:处理跨服系统属性推送消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(*player.Player)
	siPlayerSystemBattlePropertyChanged := msg.(*crosspb.SIPlayerSystemBattlePropertyChanged)
	battlePropertyData := siPlayerSystemBattlePropertyChanged.GetBattlePropertyData()
	battleProperties := pbutil.ConvertFromBattleProperty(battlePropertyData)
	power := siPlayerSystemBattlePropertyChanged.GetPower()
	err := playerSystemBattlePropertyChanged(pl, battleProperties, power)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": gcs.Player().GetId(),
			}).Error("login:玩家跨服登陆,失败")
		return err
	}

	log.Debug("login:处理跨服登陆消息完成")
	return nil
}

//玩家系统属性变化
func playerSystemBattlePropertyChanged(pl *player.Player, battleProperties map[int32]int64, power int64) (err error) {
	pl.UpdateForce(power)
	pl.UpdateSystemBattleProperty(battleProperties)
	pl.Calculate()
	return
}
