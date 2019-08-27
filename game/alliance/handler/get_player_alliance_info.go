package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_PLAYER_INFO_TYPE), dispatch.HandlerFunc(handleAlliancePlayerInfo))
}

//请求仙盟个人信息
func handleAlliancePlayerInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:请求仙盟个人信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = alliancePlayerInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:请求仙盟个人信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:请求仙盟个人信息,完成")
	return nil

}

func alliancePlayerInfo(pl player.Player) (err error) {
	allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceManager.RefreshAlliancePlayerInfo()
	//发送仙盟个人信息

	scAlliancePlayerInfo := pbutil.BuildSCAlliancePlayerInfo(allianceManager.GetPlayerAllianceObject(), allianceManager.GetPlayerAllianceSkillMap())
	pl.SendMsg(scAlliancePlayerInfo)
	return
}
