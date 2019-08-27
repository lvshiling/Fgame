package judu

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/pkg/timeutils"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/xozrc/pkg/httputils"
)

const (
	juduPath = "/judu"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(juduPath).Subrouter()
	sr.Path("/android").Handler(http.HandlerFunc(handleJuDuAndroid))
	sr.Path("/ios").Handler(http.HandlerFunc(handleJuDuIOS))
}

func handleJuDuAndroid(rw http.ResponseWriter, req *http.Request) {
	reqForm := &JuDuRequest{}
	err := httputils.Bind(req, reqForm)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:巨都安卓充值请求，参数解析错误")
		result := "error"
		rw.Write([]byte(result))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.WithFields(
		log.Fields{
			"ip":       req.RemoteAddr,
			"orderId":  reqForm.OrderId,
			"userName": reqForm.UserName,
			"gameId":   reqForm.GameId,
			"roleId":   reqForm.RoleId,
			"serverId": reqForm.ServerId,
			"payType":  reqForm.PayType,
			"amount":   reqForm.Amount,
			"payTime":  reqForm.PayTime,
			"attache":  reqForm.Attache,
			"sign":     reqForm.Sign,
		}).Info("charge:巨都安卓充值请求")

	sdkType := logintypes.SDKTypeJuDu
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:巨都安卓充值请求,sdk配置为空")
		result := "error"
		rw.Write([]byte(result))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	juDuConfig, ok := sdkConfig.(*sdksdk.JuDuConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:巨都安卓充值请求,sdk配置强制转换失败")
		result := "error"
		rw.Write([]byte(result))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeAndroid
	appKey := juDuConfig.GetAppKey(devicePlatformType)

	//TODO 验证签名
	getSign := GetJuDuSign(reqForm, appKey)
	if reqForm.Sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserName,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"payType":  reqForm.PayType,
				"amount":   reqForm.Amount,
				"payTime":  reqForm.PayTime,
				"attache":  reqForm.Attache,
				"sign":     reqForm.Sign,
				"appKey":   appKey,
			}).Warn("charge:巨都安卓充值请求,签名错误")
		result := "errorSign"
		rw.Write([]byte(result))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	now := timeutils.TimeToMillisecond(time.Now())
	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(reqForm.Attache, reqForm.OrderId, logintypes.SDKTypeJuDu, reqForm.Amount, reqForm.RoleId, now)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserName,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"payType":  reqForm.PayType,
				"amount":   reqForm.Amount,
				"payTime":  reqForm.PayTime,
				"attache":  reqForm.Attache,
				"sign":     reqForm.Sign,
				"error":    err,
			}).Error("charge:巨都安卓请求,错误")
		result := "error"
		rw.Write([]byte(result))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserName,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"payType":  reqForm.PayType,
				"amount":   reqForm.Amount,
				"payTime":  reqForm.PayTime,
				"attache":  reqForm.Attache,
				"sign":     reqForm.Sign,
			}).Warn("charge:巨都安卓请求,订单不存在")
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
			"userName": reqForm.UserName,
			"gameId":   reqForm.GameId,
			"roleId":   reqForm.RoleId,
			"serverId": reqForm.ServerId,
			"payType":  reqForm.PayType,
			"amount":   reqForm.Amount,
			"payTime":  reqForm.PayTime,
			"attache":  reqForm.Attache,
			"sign":     reqForm.Sign,
		}).Info("charge:巨都安卓充值请求")
}
