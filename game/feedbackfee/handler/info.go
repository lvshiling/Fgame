package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/feedbackfee/pbutil"
	playerfeedbackfee "fgame/fgame/game/feedbackfee/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FEEDBACK_FEE_INFO_TYPE), dispatch.HandlerFunc(handlerFeedbackFeeGetInfo))
}

//处理获取逆付费信息
func handlerFeedbackFeeGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("feedbackfee:处理获取逆付费请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = getFeedbackFeeInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("feedbackfee:处理获取逆付费请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("feedbackfee:处理获取逆付费请求完成")

	return
}

//获取逆付费请求逻辑
func getFeedbackFeeInfo(pl player.Player) (err error) {
	feedbackfeeManager := pl.GetPlayerDataManager(playertypes.PlayerFeedbackFeeDataManagerType).(*playerfeedbackfee.PlayerFeedbackFeeManager)
	feeInfo := feedbackfeeManager.GetFeedbackFeeInfo()
	recordObj := feedbackfeeManager.GetCurrentRecord()
	xianJinflag := center.GetCenterService().IsXianJinOpen()
	//发送通知
	scFeedbackFeeInfo := pbutil.BuildSCFeedbackFeeInfo(xianJinflag, feeInfo, recordObj)
	pl.SendMsg(scFeedbackFeeInfo)
	return
}
