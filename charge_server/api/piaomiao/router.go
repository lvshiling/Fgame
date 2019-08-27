package piaomiao

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
	piaomiaoPath = "/piaomiao"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(piaomiaoPath).Subrouter()
	sr.Path("/android").Handler(http.HandlerFunc(handlePiaoMiaoAndroid))
	sr.Path("/ios").Handler(http.HandlerFunc(handlePiaoMiaoIOS))
}

func handlePiaoMiaoAndroid(rw http.ResponseWriter, req *http.Request) {
	reqForm := &PiaoMiaoRequest{}
	err := httputils.Bind(req, reqForm)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:缥缈安卓充值请求，参数解析错误")
		result := "error"
		rw.Write([]byte(result))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.WithFields(
		log.Fields{
			"ip":         req.RemoteAddr,
			"OutTradeNo": reqForm.OutTradeNo,
			"Price":      reqForm.Price,
			"PayStatus":  reqForm.PayStatus,
			"Extend":     reqForm.Extend,
			"Signtype":   reqForm.Signtype,
			"sign":       reqForm.Sign,
		}).Info("charge:缥缈安卓充值请求")

	sdkType := logintypes.SDKTypePiaoMiao
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:缥缈安卓充值请求,sdk配置为空")
		result := "error"
		rw.Write([]byte(result))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	piaoMiaoConfig, ok := sdkConfig.(*sdksdk.PiaoMiaoConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:缥缈安卓充值请求,sdk配置强制转换失败")
		result := "error"
		rw.Write([]byte(result))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeAndroid
	gameKey := piaoMiaoConfig.GetGameKey(devicePlatformType)

	//TODO 验证签名
	getSign := GetPiaoMiaoSign(reqForm, gameKey)
	if reqForm.Sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"OutTradeNo": reqForm.OutTradeNo,
				"Price":      reqForm.Price,
				"PayStatus":  reqForm.PayStatus,
				"Extend":     reqForm.Extend,
				"Signtype":   reqForm.Signtype,
				"sign":       reqForm.Sign,
				"gameKey":    gameKey,
			}).Warn("charge:缥缈安卓充值请求,签名错误")
		result := "errorSign"
		rw.Write([]byte(result))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	now := timeutils.TimeToMillisecond(time.Now())
	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(reqForm.Extend, reqForm.OutTradeNo, logintypes.SDKTypePiaoMiao, int32(reqForm.Price), "", now)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"OutTradeNo": reqForm.OutTradeNo,
				"Price":      reqForm.Price,
				"PayStatus":  reqForm.PayStatus,
				"Extend":     reqForm.Extend,
				"Signtype":   reqForm.Signtype,
				"sign":       reqForm.Sign,
				"error":      err,
			}).Error("charge:缥缈安卓请求,错误")
		result := "error"
		rw.Write([]byte(result))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"OutTradeNo": reqForm.OutTradeNo,
				"Price":      reqForm.Price,
				"PayStatus":  reqForm.PayStatus,
				"Extend":     reqForm.Extend,
				"Signtype":   reqForm.Signtype,
				"sign":       reqForm.Sign,
			}).Warn("charge:缥缈安卓请求,订单不存在")
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
			"ip":         req.RemoteAddr,
			"OutTradeNo": reqForm.OutTradeNo,
			"Price":      reqForm.Price,
			"PayStatus":  reqForm.PayStatus,
			"Extend":     reqForm.Extend,
			"Signtype":   reqForm.Signtype,
			"sign":       reqForm.Sign,
		}).Info("charge:缥缈安卓充值请求")
}
