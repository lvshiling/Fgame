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
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/xuechi/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XUECHI_BLOODLINE_TYPE), dispatch.HandlerFunc(handleXueChiBloodLine))
}

//处理设置血池线
func handleXueChiBloodLine(s session.Session, msg interface{}) (err error) {
	log.Debug("xuechi:处理设置血池线")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)

	csXueChiBloodLine := msg.(*uipb.CSXueChiBloodLine)
	bloodLine := csXueChiBloodLine.GetBloodLine()

	err = xueChiBloodLine(tpl, bloodLine)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"bloodLine": bloodLine,
				"error":     err,
			}).Error("xuechi:处理设置血池线,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"bloodLine": bloodLine,
		}).Debug("xuechi:处理设置血池线完成")
	return nil
}

//处理设置血池线逻辑
func xueChiBloodLine(pl scene.Player, bloodLine int32) (err error) {
	minLineLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeRecoverHpLowerLimit)
	maxLineLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeRecoverHpUpperLimit)
	if bloodLine < minLineLimit || bloodLine > maxLineLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("xuechi:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	pl.SetBloodLine(bloodLine)
	scXueChiBloodLine := pbutil.BuildSCXueChiBloodLine(bloodLine)
	pl.SendMsg(scXueChiBloodLine)
	return
}
