package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_BATTLE_DATA_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerBattleDataChanged))
}

//战斗数据变化
func handlePlayerBattleDataChanged(s session.Session, msg interface{}) error {
	log.Debug("login:处理跨服战斗数据变化消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(*player.Player)
	siPlayerBattleDataChanged := msg.(*crosspb.SIPlayerBattleDataChanged)

	err := playerBattleData(pl, siPlayerBattleDataChanged.GetPlayerBattleData())
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": gcs.Player().GetId(),
			}).Error("login:玩家战斗数据变化,失败")
		return err
	}

	log.Debug("login:处理战斗数据变化,完成")
	return nil
}

//玩家显示变化
func playerBattleData(pl *player.Player, playerBattleData *crosspb.PlayerBattleData) (err error) {
	if playerBattleData.SoulAwakenNum != nil {
		pl.SyncSoulAwakenNum(playerBattleData.GetSoulAwakenNum())
	}
	if playerBattleData.Level != nil {
		pl.SyncLevel(playerBattleData.GetLevel())
	}

	if playerBattleData.ZhuanSheng != nil {
		pl.SyncZhuanSheng(playerBattleData.GetZhuanSheng())
	}

	return nil
}
