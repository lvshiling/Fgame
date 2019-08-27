package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	crosslogic "fgame/fgame/game/cross/logic"
	crossloginflow "fgame/fgame/game/cross/loginflow"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_LOGIN_TYPE), dispatch.HandlerFunc(handleLogin))
}

//处理跨服登录
func handleLogin(s session.Session, msg interface{}) error {
	log.Info("cross:处理跨服登录")
	gameS := gamesession.SessionInContext(s.Context())
	isLogin := msg.(*crosspb.ISLogin)
	match := isLogin.GetMatch()
	pl := gameS.Player().(player.Player)
	err := crossLogin(pl, match)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("cross:玩家跨服登陆,失败")
		return err
	}

	log.Info("cross:处理跨服登陆消息完成")
	return nil
}

//登陆
func crossLogin(pl player.Player, match bool) (err error) {

	//已经在比赛中
	if match {
		//跨服传送数据
		crosslogic.CrossPlayerDataLogin(pl)
		return
	}

	//判断跨服模式
	crossType := pl.GetCrossType()
	crossArgs := pl.GetCrossArgs()
	if len(crossArgs) >= 1 {
		crossArgs = crossArgs[1:]
	}
	switch crossType {
	case crosstypes.CrossTypeArena, //3v3
		crosstypes.CrossTypeLianYu,          //无间炼狱
		crosstypes.CrossTypeGodSiegeQiLin,   //神兽攻城
		crosstypes.CrossTypeGodSiegeHuoFeng, //
		crosstypes.CrossTypeGodSiegeDuLong,  //
		crosstypes.CrossTypeTeamCopy,        //组队副本
		crosstypes.CrossTypeDenseWat,        //金银密窟
		crosstypes.CrossTypeShenMoWar,       //神魔战场
		crosstypes.CrossTypeArenapvp,        //比武大会
		crosstypes.CrossTypeChuangShi,       //创世之战
		crosstypes.CrossTypeTuLong:          //跨服屠龙
		{
			return crossloginflow.CrossLoginFlowHandle(pl, crossType, crossArgs...)
		}
	}

	//非匹配模式 进入活动场景
	crosslogic.CrossPlayerDataLogin(pl)
	return

}
