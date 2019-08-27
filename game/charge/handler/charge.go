package handler

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/charge/charge"
	chargelogic "fgame/fgame/game/charge/logic"
	chargetemplate "fgame/fgame/game/charge/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_CHARGE_TYPE), dispatch.HandlerFunc(handleCharge))
}

//处理充值
func handleCharge(s session.Session, msg interface{}) (err error) {
	log.Debug("charge:处理获取充值")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csCharge := msg.(*uipb.CSCharge)
	chargeId := csCharge.GetChargeId()

	err = chargeGold(tpl, chargeId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chargeId": chargeId,
				"error":    err,
			}).Error("charge:处理获取充值,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"chargeId": chargeId,
		}).Debug("charge:处理获取充值完成")
	return nil

}

//处理充值
func chargeGold(pl player.Player, chargeId int32) (err error) {
	chargeTemp := chargetemplate.GetChargeTemplateService().GetChargeTemplate(chargeId)
	if chargeTemp == nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"chargeId": chargeId,
			}).Warn("charge:处理充值,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if chargeTemp.GetType() != pl.GetSDKType() {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"chargeId": chargeId,
				"sdkType":  pl.GetSDKType(),
			}).Warn("charge:处理充值,sdk不对")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	//TODO 正式服去掉
	//pc平台直接下单充值
	if chargeTemp.GetType() == logintypes.SDKTypePC {
		if !pl.IsGm() {
			return
		}
		str := uuid.NewV4().String()
		orderId := strings.Replace(str, "-", "", -1)
		chargelogic.OnPlayerCharge(pl, orderId, chargeId)
		return
	}

	flag := charge.GetChargeService().GetOrder(pl, chargeId)
	if !flag {
		panic(fmt.Errorf("charge:下单应该成功"))
	}
	return
}
