package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	alliancetypes "fgame/fgame/game/alliance/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_ALLIANCE_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerAllianceChanged))
}

//玩家仙盟变化
func handlePlayerAllianceChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:玩家仙盟变化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siPlayerAllianceSync := msg.(*crosspb.SIPlayerAllianceSync)
	allianceId := siPlayerAllianceSync.GetAllianceData().GetAllianceId()
	allianceName := siPlayerAllianceSync.GetAllianceData().GetAllianceName()
	mengZhuId := siPlayerAllianceSync.GetAllianceData().GetMengZhuId()
	memPos := alliancetypes.AlliancePosition(siPlayerAllianceSync.GetAllianceData().GetMemPos())
	if !memPos.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("team:玩家队伍变化,职位错误")
	}
	err = playerAllianceChanged(tpl, allianceId, allianceName, mengZhuId, memPos)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("team:玩家队伍变化,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:玩家队伍变化,完成")
	return nil

}

//仙盟变化
func playerAllianceChanged(pl *player.Player, allianceId int64, allianceName string, mengZhuId int64, memPos alliancetypes.AlliancePosition) (err error) {
	pl.SyncAlliance(allianceId, allianceName, mengZhuId, memPos)
	return
}
