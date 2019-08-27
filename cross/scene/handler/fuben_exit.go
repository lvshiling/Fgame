package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/arena/arena"
	"fgame/fgame/cross/arenapvp/arenapvp"
	arenapvppbutil "fgame/fgame/cross/arenapvp/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/cross/teamcopy/teamcopy"
	"fgame/fgame/game/scene/pbutil"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FUBEN_EXIT_TYPE), dispatch.HandlerFunc(handleFuBenExit))
}

//处理副本退出
func handleFuBenExit(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:处理副本退出场景")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)
	//TODO 判断
	err = fuBenExit(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("scene:处理副本退出场景,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("scene:处理副本退出场景,完成")

	return nil
}

//玩家退出副本
func fuBenExit(pl *player.Player) (err error) {
	//TODO 判断是否在副本
	s := pl.GetScene()
	if s == nil {
		return
	}

	//逃跑
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeArena || s.MapTemplate().GetMapType() == scenetypes.SceneTypeArenaShengShou {
		arena.GetArenaService().ArenaMemeberExit(pl)
	}

	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeCrossTeamCopy {
		teamcopy.GetTeamCopyService().TeamCopyMemeberExit(pl)
	}

	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeArenapvp || s.MapTemplate().GetMapType() == scenetypes.SceneTypeArenapvpHaiXuan {
		pvpPlayer, flag := arenapvp.GetArenapvpService().PvpPlayerExit(pl.GetId())
		if !flag {
			goto Exit
		}

		if s.MapTemplate().GetMapType() == scenetypes.SceneTypeArenapvp {
			scArenapvpPlayerStateChanged := arenapvppbutil.BuildSCArenapvpPlayerStateChanged(pvpPlayer)
			s.BroadcastMsg(scArenapvpPlayerStateChanged)
		}

		curBattleData := pvpPlayer.GetCurBatlleData()
		if curBattleData.WinnerId == 0 {
			isMsg := arenapvppbutil.BuildISArenapvpResultBattle(false, int32(curBattleData.PvpType))
			pl.SendMsg(isMsg)
		}
	}
Exit:
	scFuBenExit := pbutil.BuildSCFuBenExit()
	pl.SendMsg(scFuBenExit)

	pl.BackLastScene()
	return
}
