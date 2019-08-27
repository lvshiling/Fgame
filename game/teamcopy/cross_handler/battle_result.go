package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	teamtypes "fgame/fgame/game/team/types"
	teamcopylogic "fgame/fgame/game/teamcopy/logic"
	"fgame/fgame/game/teamcopy/pbutil"
	playerteamcopy "fgame/fgame/game/teamcopy/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_TEAMCOPY_BATTLE_RESULT_TYPE), dispatch.HandlerFunc(handleTeamCopyBattleResult))
}

//处理组队副本战斗结果
func handleTeamCopyBattleResult(s session.Session, msg interface{}) (err error) {
	log.Debug("teamcopy:处理组队副本战斗结果")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isTeamCopyBattleResult := msg.(*crosspb.ISTeamCopyBattleResult)
	purpose := isTeamCopyBattleResult.GetPurpose()
	sucess := isTeamCopyBattleResult.GetSucess()

	err = teamCopyBattleResult(tpl, purpose, sucess)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"purpose":  purpose,
				"sucess":   sucess,
			}).Error("teamcopy:处理组队副本战斗结果,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"sucess":   sucess,
		}).Debug("teamcopy:处理组队副本战斗结果,完成")
	return nil

}

//战斗结果
func teamCopyBattleResult(pl player.Player, purpose int32, sucess bool) (err error) {
	teamPurpose := teamtypes.TeamPurposeType(purpose)
	manager := pl.GetPlayerDataManager(types.PlayerTeamCopyDataManagerType).(*playerteamcopy.PlayerTeamCopyDataManager)
	obj, isRew := manager.FinishPurpose(teamPurpose, sucess)
	if isRew {
		err = teamcopylogic.GiveRewardTeamCopyPurpose(pl, teamPurpose)
		if err != nil {
			return
		}
	}
	scTeamCopyResult := pbutil.BuildSCTeamCopyResult(obj, sucess, isRew)
	pl.SendMsg(scTeamCopyResult)
	siTeamCopyBattleResult := pbutil.BuildSITeamCopyBattleResult()
	pl.SendCrossMsg(siTeamCopyBattleResult)
	return
}
