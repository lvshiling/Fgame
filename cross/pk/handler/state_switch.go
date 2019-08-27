package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/processor"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/pk/pbutil"
	pktypes "fgame/fgame/game/pk/types"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_PK_STATE_SWITCH_TYPE), dispatch.HandlerFunc(handlePkStateSwitch))
}

//处理pk状态改变
func handlePkStateSwitch(s session.Session, msg interface{}) (err error) {
	log.Debug("pk:处理pk状态改变消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(scene.Player)
	csPkStateSwitch := msg.(*uipb.CSPkStateSwitch)
	pkStateInt := csPkStateSwitch.GetPkState()

	pkState := pktypes.PkState(pkStateInt)
	if !pkState.Valid() {
		playerlogic.SendSystemMessage(pl, lang.PkStateInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pkState":  pkStateInt,
			}).Warn("pk:处理pk状态改变消息,状态无效")
		return
	}

	err = pkStateSwitch(pl, pkState)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pkState":  pkStateInt,
				"error":    err,
			}).Error("pk:处理pk状态改变消息,失败")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"pkState":  pkStateInt,
		}).Debug("pk:处理pk状态改变消息,成功")
	return
}

//切换pk模式
func pkStateSwitch(pl scene.Player, state pktypes.PkState) (err error) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("pk:玩家不在场景中")
		return
	}
	protectedLevel := constant.GetConstantService().GetConstant(constanttypes.ConstantTypePKProtectLevel)
	if pl.GetLevel() < protectedLevel {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("pk:玩家pk保护中")
		playerlogic.SendSystemMessage(pl, lang.PkStateProtect, fmt.Sprintf("%d", protectedLevel))
		return
	}

	if s.MapTemplate().LimitPkMode&state.Mask() == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"state":    state.String(),
			}).Warn("pk:场景禁止pk模式")
		playerlogic.SendSystemMessage(pl, lang.PkStateForbid)
		return
	}

	pl.SwitchPkState(state, pktypes.PkCommonCampDefault)
	scPkStateSwitch := pbutil.BuildSCPkStateSwitch(state)
	pl.SendMsg(scPkStateSwitch)
	return
}
