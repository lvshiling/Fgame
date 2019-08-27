package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	crosslogic "fgame/fgame/game/cross/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_PLAYER_EXIT_CROSS_TYPE), dispatch.HandlerFunc(handleExitCross))
}

//处理跨服退出
func handleExitCross(s session.Session, msg interface{}) error {
	log.Debug("cross:处理跨服退出")
	gameS := gamesession.SessionInContext(s.Context())

	pl := gameS.Player().(player.Player)
	err := crossExit(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("cross:玩家跨服登陆,失败")
		return err
	}

	log.Debug("cross:处理跨服登陆消息完成")
	return nil
}

//登陆
func crossExit(pl player.Player) (err error) {
	//判断是否是组队
	//zrc:临时去掉离队
	// crossType := pl.GetCrossType()
	// switch crossType {
	// case crosstypes.CrossTypeArena,
	// 	crosstypes.CrossTypeTeamCopy:
	// 	team.GetTeamService().LeaveTeam(pl)
	// }
	crosslogic.PlayerExitCross(pl)
	return
}
