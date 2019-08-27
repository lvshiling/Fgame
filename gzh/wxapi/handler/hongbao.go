package handler

import (
	"fgame/fgame/gzh/wxapi"
	"fgame/fgame/gzh/wxpay"
	"fmt"

	log "github.com/Sirupsen/logrus"
	wxcontext "github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/message"
)

func HandleHongBao(v message.MixMessage, ctx *wxcontext.Context) *message.Reply {

	//加前缀

	remoteService := wxapi.RemoteServiceInContext(ctx.Request.Context())
	returnCode, returnMsg, orderId, money, clientIp, err := remoteService.GetOrder(v.FromUserName, v.Content)
	if err != nil {
		log.WithFields(
			log.Fields{
				"code": v.Content,
				"wxId": v.FromUserName,
				"err":  err,
			}).Error("兑换异常")
		trans := message.NewText("兑换异常")
		reply := &message.Reply{
			MsgType: message.MsgTypeText,
			MsgData: trans,
		}
		return reply
	}
	if returnCode != 0 {
		log.WithFields(
			log.Fields{
				"code":      v.Content,
				"wxId":      v.FromUserName,
				"returnMsg": returnMsg,
			}).Warn("兑换错误")
		trans := message.NewText(returnMsg)
		reply := &message.Reply{
			MsgType: message.MsgTypeText,
			MsgData: trans,
		}
		return reply
	}
	openId := v.FromUserName
	desc := "红包"
	amount := fmt.Sprintf("%d", money)
	_, errMsg, err := wxpay.GetWxPayServiceInstance().WithdrawMoney(openId, amount, orderId, desc, clientIp)
	if err != nil {
		log.WithFields(
			log.Fields{
				"code": v.Content,
				"wxId": v.FromUserName,
				"err":  err,
			}).Error("兑换异常")
		trans := message.NewText("兑换异常")
		reply := &message.Reply{
			MsgType: message.MsgTypeText,
			MsgData: trans,
		}
		return reply
	}
	if len(errMsg) != 0 {
		log.WithFields(
			log.Fields{
				"code":   v.Content,
				"wxId":   v.FromUserName,
				"errMsg": errMsg,
			}).Warn("兑换失败")
		trans := message.NewText(errMsg)
		reply := &message.Reply{
			MsgType: message.MsgTypeText,
			MsgData: trans,
		}
		return reply
	}

	trans := message.NewText("兑换成功")
	reply := &message.Reply{
		MsgType: message.MsgTypeText,
		MsgData: trans,
	}
	return reply

	// trans := message.NewTransferCustomer("")
	// reply := &message.Reply{
	// 	MsgType: message.MsgTypeTransfer,
	// 	MsgData: trans,
	// }
	// return reply
}
