package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	handler "fgame/fgame/gzh/wxapi/handler"

	"fgame/fgame/gzh/wxapi"

	log "github.com/Sirupsen/logrus"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

func handlegzhapi(rw http.ResponseWriter, req *http.Request) {

	queryForm, _ := url.ParseQuery(req.URL.RawQuery)
	if len(queryForm["echostr"]) > 0 {
		echostr := []byte(queryForm["echostr"][0])
		fmt.Println(echostr)
		rw.WriteHeader(http.StatusOK)
		rw.Write(echostr)
		return
	}

	wcconfig := wxapi.WeChatServiceInContext(req.Context())

	wc := wechat.NewWechat(wcconfig.Wc)
	if wc == nil {
		log.Error("获取微信对象为空")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	server := wc.GetServer(req, rw)

	server.SetMessageHandler(func(v message.MixMessage) *message.Reply {
		btmsg, _ := json.Marshal(v)
		log.WithFields(log.Fields{
			"MsgType": v.MsgType,
			"Event":   v.Event,
			"allmsg":  string(btmsg),
		}).Info("处理消息中...")

		switch v.MsgType {
		case message.MsgTypeEvent: //推送事件
			switch v.Event {
			case message.EventSubscribe: //只接受订阅和取消订阅
				//do something
				sv := &handler.EventSubscribeHandler{}
				return sv.HandlerMsg(v, wc.Context)
			}
		case message.MsgTypeText: //文本信息,重新定义一个handler
			return handler.HandleHongBao(v, wc.Context)
		case message.MsgTypeImage: //图像信息
			tc := &handler.MsgTransferCustomerHandler{}
			return tc.HandlerMsg(v, wc.Context)
		case message.MsgTypeLink:
			tc := &handler.MsgTransferCustomerHandler{}
			return tc.HandlerMsg(v, wc.Context)
		case message.MsgTypeVideo:
			tc := &handler.MsgTransferCustomerHandler{}
			return tc.HandlerMsg(v, wc.Context)
		case message.MsgTypeVoice:
			tc := &handler.MsgTransferCustomerHandler{}
			return tc.HandlerMsg(v, wc.Context)
		case message.MsgTypeShortVideo:
			tc := &handler.MsgTransferCustomerHandler{}
			return tc.HandlerMsg(v, wc.Context)

		}
		return nil
	})

	serverErr := server.Serve()
	if serverErr != nil {
		log.WithFields(log.Fields{
			"失败": serverErr.Error(),
		}).Error("处理服务失败")
	}
	sendErr := server.Send()
	if sendErr != nil {
		log.WithFields(log.Fields{
			"Send失败": serverErr.Error(),
		}).Error("处理服务失败Send")
	}
}
