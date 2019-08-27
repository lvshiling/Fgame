package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/arena/arena"
	"fgame/fgame/cross/arena/pbutil"
	arenascene "fgame/fgame/cross/arena/scene"
	"fgame/fgame/cross/player/player"

	"fgame/fgame/cross/processor"
	arenatypes "fgame/fgame/game/arena/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_TEAM_INFO_LIST_TYPE), dispatch.HandlerFunc(handleArenaFourGodTeamInfoList))
}

//处理获取四圣兽队伍信息
func handleArenaFourGodTeamInfoList(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理获取四圣兽队伍信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)
	csArenaFourGodTeamInfoList := msg.(*uipb.CSArenaFourGodTeamInfoList)
	fourGodType := arenatypes.FourGodType(csArenaFourGodTeamInfoList.GetFourGodType())
	err = arenaFourGodTeamInfoList(tpl, fourGodType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理获取四圣兽信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理选择四圣兽")
	return nil

}

//四圣兽队伍信息
func arenaFourGodTeamInfoList(pl *player.Player, fourGodType arenatypes.FourGodType) (err error) {
	//判断是否有组队
	s := arena.GetArenaService().GetFourGodScene(fourGodType)
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"fourGodType": fourGodType.String(),
			}).Warn("arena:处理获取四圣兽场景,已经关闭")
		return
	}
	sd := s.SceneDelegate().(arenascene.FourGodSceneData)
	scArenaFourGodTeamInfoList := pbutil.BuildSCArenaFourGodTeamInfoList(sd)
	pl.SendMsg(scArenaFourGodTeamInfoList)
	return
}
