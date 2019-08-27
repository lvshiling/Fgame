package jiumeng

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/pkg/timeutils"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/xozrc/pkg/httputils"
)

const (
	jiumengPath = "/jiumeng"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(jiumengPath).Subrouter()
	sr.Path("/android").Handler(http.HandlerFunc(handleJiuMengAndroid))
	sr.Path("/ios").Handler(http.HandlerFunc(handleJiuMengIOS))
}

func handleJiuMengAndroid(rw http.ResponseWriter, req *http.Request) {
	reqForm := &JiuMengRequest{}
	err := httputils.Bind(req, reqForm)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:九梦android充值请求，参数解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.WithFields(
		log.Fields{
			"ip":       req.RemoteAddr,
			"orderId":  reqForm.OrderId,
			"userName": reqForm.UserId,
			"gameId":   reqForm.GameId,
			"roleId":   reqForm.RoleId,
			"serverId": reqForm.ServerId,
			"amount":   reqForm.Money,
			"money":    reqForm.PayTime,
			"attach":   reqForm.Attach,
			"sign":     reqForm.Sign,
		}).Info("charge:九梦android充值请求")

	sdkType := logintypes.SDKTypeJiuMeng
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:九梦android充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	qiAConfig, ok := sdkConfig.(*sdksdk.JiuMengConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:九梦android充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeAndroid
	appKey := qiAConfig.GetAppKey(devicePlatformType)

	//TODO 验证签名
	getSign := GetJiuMengSign(reqForm, appKey)
	if reqForm.Sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserId,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"amount":   reqForm.Money,
				"money":    reqForm.PayTime,
				"attach":   reqForm.Attach,
				"sign":     reqForm.Sign,
				"appKey":   appKey,
			}).Warn("charge:九梦android充值请求,签名错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	money, err := strconv.ParseInt(reqForm.Money, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserId,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"amount":   reqForm.Money,
				"money":    reqForm.PayTime,
				"attach":   reqForm.Attach,
				"sign":     reqForm.Sign,
				"appKey":   appKey,
			}).Warn("charge:九梦android充值请求,解析错误")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := timeutils.TimeToMillisecond(time.Now())
	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(reqForm.Attach, reqForm.OrderId, logintypes.SDKTypeJiuMeng, int32(money)/100, fmt.Sprintf("%d", reqForm.UserId), now)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserId,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"amount":   reqForm.Money,
				"money":    reqForm.PayTime,
				"attach":   reqForm.Attach,
				"sign":     reqForm.Sign,
				"error":    err,
			}).Error("charge:九梦ios请求,错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserId,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"amount":   reqForm.Money,
				"money":    reqForm.PayTime,
				"attach":   reqForm.Attach,
				"sign":     reqForm.Sign,
			}).Warn("charge:九梦ios请求,订单不存在")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if !repeat {
		//放入回调队列中
		remoteService := remote.RemoteServiceInContext(ctx)
		flag := remoteService.Charge(obj)
		if !flag {
			panic(fmt.Errorf("charge:添加到回调队列应该成功"))
		}
	}

	result := "success"
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(result))
	log.WithFields(
		log.Fields{
			"ip":       req.RemoteAddr,
			"orderId":  reqForm.OrderId,
			"userName": reqForm.UserId,
			"gameId":   reqForm.GameId,
			"roleId":   reqForm.RoleId,
			"serverId": reqForm.ServerId,
			"amount":   reqForm.Money,
			"money":    reqForm.PayTime,
			"attach":   reqForm.Attach,
			"sign":     reqForm.Sign,
		}).Info("charge:九梦android充值请求")
}
