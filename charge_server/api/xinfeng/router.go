package xinfeng

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

const (
	xinfengPath = "/xinfeng"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(xinfengPath).Subrouter()
	sr.Path("/ios").Handler(http.HandlerFunc(handleXinFengIOS))
	sr.Path("/android").Handler(http.HandlerFunc(handleXinFengAndroid))
}

func handleXinFengAndroid(rw http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	cpOrderIdStr := query.Get("cp_orderid")
	orderIdStr := query.Get("orderid")
	appIdStr := query.Get("appid")
	uidStr := query.Get("uid")
	amountStr := query.Get("amount")
	timeStr := query.Get("time")
	sign := query.Get("sign")

	log.WithFields(
		log.Fields{
			"ip":           req.RemoteAddr,
			"cpOrderIdStr": cpOrderIdStr,
			"orderIdStr":   orderIdStr,
			"appIdStr":     appIdStr,
			"uidStr":       uidStr,
			"amountStr":    amountStr,
			"timeStr":      timeStr,
			"sign":         sign,
		}).Info("charge:新蜂安卓充值请求")

	amountFloat, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"cpOrderIdStr": cpOrderIdStr,
				"orderIdStr":   orderIdStr,
				"appIdStr":     appIdStr,
				"uidStr":       uidStr,
				"amountStr":    amountStr,
				"timeStr":      timeStr,
				"sign":         sign,
				"error":        err,
			}).Warn("charge:新蜂安卓充值请求，解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	amount := int32(amountFloat)

	receiveTime, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"cpOrderIdStr": cpOrderIdStr,
				"orderIdStr":   orderIdStr,
				"appIdStr":     appIdStr,
				"uid":          uidStr,
				"amountStr":    amountStr,
				"timeStr":      timeStr,
				"sign":         sign,
				"error":        err,
			}).Warn("charge:新蜂安卓充值请求，解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	sdkType := logintypes.SDKTypeXinFeng
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:新蜂安卓充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	xinFengConfig, ok := sdkConfig.(*sdksdk.XinFengConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:新蜂安卓充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeAndroid
	chargeKey := xinFengConfig.GetChargeKey(devicePlatformType)

	//TODO 验证签名
	getSign := xinFengSign(cpOrderIdStr, orderIdStr, appIdStr, uidStr, amount, receiveTime, chargeKey)
	if sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"cpOrderIdStr": cpOrderIdStr,
				"orderIdStr":   orderIdStr,
				"appIdStr":     appIdStr,
				"uid":          uidStr,
				"amountStr":    amountStr,
				"timeStr":      timeStr,
				"sign":         sign,
				"getSign":      getSign,
				"error":        err,
			}).Warn("charge:新蜂安卓充值请求,签名错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(cpOrderIdStr, orderIdStr, logintypes.SDKTypeXinFeng, amount, uidStr, receiveTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"cpOrderIdStr": cpOrderIdStr,
				"orderIdStr":   orderIdStr,
				"appIdStr":     appIdStr,
				"uid":          uidStr,
				"amountStr":    amountStr,
				"timeStr":      timeStr,
				"sign":         sign,
				"error":        err,
			}).Error("charge:新蜂ios请求,错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"cpOrderIdStr": cpOrderIdStr,
				"orderIdStr":   orderIdStr,
				"appIdStr":     appIdStr,
				"uid":          uidStr,
				"amountStr":    amountStr,
				"timeStr":      timeStr,
				"sign":         sign,
				"error":        err,
			}).Warn("charge:新蜂ios请求,订单不存在")
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
			"cpOrderIdStr": cpOrderIdStr,
			"orderIdStr":   orderIdStr,
			"appIdStr":     appIdStr,
			"uid":          uidStr,
			"amountStr":    amountStr,
			"timeStr":      timeStr,
			"sign":         sign,
		}).Info("charge:新蜂安卓充值请求")
}
