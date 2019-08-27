package zhangjian

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
	zhangjianPath = "/zhangjian"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(zhangjianPath).Subrouter()
	sr.Path("/ios").Handler(http.HandlerFunc(handleZhangJianIOS))
	sr.Path("/android").Handler(http.HandlerFunc(handleZhangJianAndroid))
}

func handleZhangJianAndroid(rw http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	userIdStr := query.Get("userid")
	serverStr := query.Get("server")
	moneyStr := query.Get("money")
	pay := query.Get("pay")
	order := query.Get("order")
	timeStr := query.Get("time")
	sign := query.Get("sign")
	log.WithFields(
		log.Fields{
			"ip":      req.RemoteAddr,
			"userId":  userIdStr,
			"server":  serverStr,
			"money":   moneyStr,
			"pay":     pay,
			"order":   order,
			"timeStr": timeStr,
			"sign":    sign,
		}).Info("charge:仗剑安卓充值请求")

	serverIdInt, err := strconv.ParseInt(serverStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":      req.RemoteAddr,
				"userId":  userIdStr,
				"server":  serverStr,
				"money":   moneyStr,
				"pay":     pay,
				"order":   order,
				"timeStr": timeStr,
				"sign":    sign,
				"error":   err,
			}).Warn("charge:仗剑安卓充值请求，解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	server := int32(serverIdInt)
	moneyFloat, err := strconv.ParseFloat(moneyStr, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":      req.RemoteAddr,
				"userId":  userIdStr,
				"server":  serverStr,
				"money":   moneyStr,
				"pay":     pay,
				"order":   order,
				"timeStr": timeStr,
				"sign":    sign,
				"error":   err,
			}).Warn("charge:仗剑安卓充值请求，解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	money := int32(moneyFloat)
	receiveTime, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":      req.RemoteAddr,
				"userId":  userIdStr,
				"server":  serverStr,
				"money":   moneyStr,
				"pay":     pay,
				"order":   order,
				"timeStr": timeStr,
				"sign":    sign,
				"error":   err,
			}).Warn("charge:仗剑安卓充值请求，解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	sdkType := logintypes.SDKTypeZhangJian
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:仗剑安卓充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	hengGeWanConfig, ok := sdkConfig.(*sdksdk.ZhangJianConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:仗剑安卓充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeAndroid
	chargeKey := hengGeWanConfig.GetChargeKey(devicePlatformType)

	//TODO 验证签名
	getSign := zhangJianSign(chargeKey, userIdStr, server, moneyStr, pay, order, receiveTime)
	if sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"orderId":     order,
				"userId":      userIdStr,
				"server":      server,
				"money":       money,
				"pay":         pay,
				"receiveTime": receiveTime,
				"sign":        sign,
				"getSign":     getSign,
			}).Warn("charge:仗剑安卓充值请求,签名错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(order, pay, logintypes.SDKTypeZhangJian, money, userIdStr, receiveTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"orderId":     order,
				"userId":      userIdStr,
				"server":      server,
				"money":       money,
				"pay":         pay,
				"receiveTime": receiveTime,
				"sign":        sign,
				"error":       err,
			}).Error("charge:仗剑安卓请求,错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"orderId":     order,
				"userId":      userIdStr,
				"server":      server,
				"money":       money,
				"pay":         pay,
				"receiveTime": receiveTime,
				"sign":        sign,
				"error":       err,
			}).Warn("charge:仗剑安卓请求,订单不存在")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	//放入回调队列中
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
			"orderId":     order,
			"userId":      userIdStr,
			"server":      server,
			"money":       money,
			"pay":         pay,
			"receiveTime": receiveTime,
			"sign":        sign,
		}).Info("charge:仗剑安卓充值请求")
}
