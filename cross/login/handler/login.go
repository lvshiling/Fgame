package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/arena/arena"
	loginlogic "fgame/fgame/cross/login/logic"
	"fgame/fgame/cross/login/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/cross/teamcopy/teamcopy"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_LOGIN_TYPE), dispatch.HandlerFunc(handleCrossLogin))
}

//处理登录
func handleCrossLogin(s session.Session, msg interface{}) error {
	log.Info("login:处理跨服登陆消息")

	gcs := gamesession.SessionInContext(s.Context())
	//玩家重复登录
	if gcs.State() != gamesession.SessionStateInit {
		//TODO 发送错误信息
		log.WithFields(
			log.Fields{
				"playerId": gcs.Player().GetId(),
			}).Warn("login:玩家重复登陆")
		return nil
	}
	siLogin := msg.(*crosspb.SILogin)
	playerId := siLogin.GetPlayerId()
	err := crossLogin(gcs, playerId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": gcs.Player().GetId(),
			}).Error("login:玩家跨服登陆,失败")
		return err
	}

	log.Info("login:处理跨服登陆消息完成")
	return nil
}

//登陆
func crossLogin(gs gamesession.Session, playerId int64) (err error) {
	//玩家认证
	p := player.NewPlayer(gs, playerId)
	gs.Auth(p)
	flag := loginlogic.Login(p)
	if !flag {
		return
	}
	//判断是否在比赛
	arenaTeam := arena.GetArenaService().GetArenaTeamByPlayerId(playerId)
	match := false
	if arenaTeam != nil {
		match = true
	}
	//组队副本
	teamCopyData := teamcopy.GetTeamCopyService().GetTeamCopyDataByPlayerId(playerId)
	if teamCopyData != nil {
		match = true
	}
	isLogin := pbutil.BuildISLogin(playerId, match)
	//创建新玩家
	p.SendMsg(isLogin)
	return
}
