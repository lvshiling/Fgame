package handler

import (
	"fgame/fgame/core/session"
	"fmt"

	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/game/misc/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/timeutils"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EXIT_KASI_TYPE), dispatch.HandlerFunc(handleExitKaSi))

}

//脱离卡死
func handleExitKaSi(s session.Session, msg interface{}) error {
	log.Debug("misc:脱离卡死")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err := exitKaSi(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("misc:脱离卡死,错误")

		return err
	}
	log.Debug("misc:脱离卡死,完成")
	return nil
}

func exitKaSi(pl player.Player) (err error) {
	pls := pl.GetScene()
	if pls != nil && !pls.MapTemplate().IfChangeScenePvp() {
		//战斗状态
		if pl.IsPvpBattle() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("player:当前处于战斗状态")
			playerlogic.SendSystemMessage(pl, lang.PlayerInPVP)
			return
		}
	}

	if !pl.IfCanExitKaSi() {
		leftTimeMs := pl.ExitKaSiLeftTime()
		leftTime := timeutils.MillisecondToSecondCeil(leftTimeMs)
		leftTimeStr := fmt.Sprintf("%d", leftTime)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("misc:脱离卡死,cd")
		playerlogic.SendSystemMessage(pl, lang.MiscExitKaSiCd, leftTimeStr)
		return
	}
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("misc:脱离卡死,不在场景中")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoInScene)
		return
	}

	pl.ExitKaSi()
	scenelogic.FixPosition(pl, s.MapTemplate().GetBornPos())

	//发送卡死
	scExitKaSi := pbutil.BuildSCExitKaSi()
	err = pl.SendMsg(scExitKaSi)
	return
}
