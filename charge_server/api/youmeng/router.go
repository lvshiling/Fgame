package youmeng

import (
	"crypto/md5"
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
	youmengPath = "/youmeng"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(youmengPath).Subrouter()
	sr.Path("/ios").Handler(http.HandlerFunc(handleYouMengIOS))
	sr.Path("/android").Handler(http.HandlerFunc(handleYouMengAndroid))
}

func handleYouMengAndroid(rw http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	userIdStr := query.Get("uid")
	moneyStr := query.Get("amount")
	appIdStr := query.Get("appid")
	// 对方订单号
	order := query.Get("orderid")
	// 我方订单号
	cpOrder := query.Get("cp_orderid")
	timeStr := query.Get("time")
	sign := query.Get("sign")
	log.WithFields(
		log.Fields{
			"ip":         req.RemoteAddr,
			"userId":     userIdStr,
			"appid":      appIdStr,
			"money":      moneyStr,
			"cp_orderid": cpOrder,
			"orderid":    order,
			"timeStr":    timeStr,
			"sign":       sign,
		}).Info("charge:游梦江湖安卓充值请求")

	appId, err := strconv.ParseInt(appIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"userId":     userIdStr,
				"appid":      appId,
				"money":      moneyStr,
				"cp_orderid": cpOrder,
				"orderid":    order,
				"timeStr":    timeStr,
				"sign":       sign,
				"error":      err,
			}).Warn("charge:游梦江湖安卓充值请求，解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"userId":     userIdStr,
				"appid":      appId,
				"money":      moneyStr,
				"cp_orderid": cpOrder,
				"orderid":    order,
				"timeStr":    timeStr,
				"sign":       sign,
				"error":      err,
			}).Warn("charge:游梦江湖安卓充值请求，解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	moneyFloat, err := strconv.ParseFloat(moneyStr, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"userId":     userIdStr,
				"appid":      appId,
				"money":      moneyStr,
				"cp_orderid": cpOrder,
				"orderid":    order,
				"timeStr":    timeStr,
				"sign":       sign,
				"error":      err,
			}).Warn("charge:游梦江湖安卓充值请求，解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	money := int32(moneyFloat)
	receiveTime, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"userId":     userIdStr,
				"appid":      appId,
				"money":      moneyStr,
				"cp_orderid": cpOrder,
				"orderid":    order,
				"timeStr":    timeStr,
				"sign":       sign,
				"error":      err,
			}).Warn("charge:游梦江湖安卓充值请求，解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	sdkType := logintypes.SDKTypeYouMeng
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:游梦江湖安卓充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	youMengConfig, ok := sdkConfig.(*sdksdk.YouMengConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:游梦江湖安卓充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeAndroid
	chargeKey := youMengConfig.GetChargeKey(devicePlatformType)

	//TODO 验证签名
	getSign := youMengSign(cpOrder, order, appId, userId, money, receiveTime, chargeKey)
	if sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"userId":     userIdStr,
				"appid":      appId,
				"money":      moneyStr,
				"cp_orderid": cpOrder,
				"orderid":    order,
				"timeStr":    timeStr,
				"sign":       sign,
				"getSign":    getSign,
			}).Warn("charge:游梦江湖安卓充值请求,签名错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(cpOrder, order, logintypes.SDKTypeYouMeng, money, userIdStr, receiveTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"userId":     userIdStr,
				"appid":      appId,
				"money":      moneyStr,
				"cp_orderid": cpOrder,
				"orderid":    order,
				"timeStr":    timeStr,
				"sign":       sign,
				"error":      err,
			}).Error("charge:游梦江湖安卓请求,错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"userId":     userIdStr,
				"appid":      appId,
				"money":      moneyStr,
				"cp_orderid": cpOrder,
				"orderid":    order,
				"timeStr":    timeStr,
				"sign":       sign,
				"error":      err,
			}).Warn("charge:游梦江湖安卓请求,订单不存在")
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
			"ip":         req.RemoteAddr,
			"userId":     userIdStr,
			"appid":      appId,
			"money":      moneyStr,
			"cp_orderid": cpOrder,
			"orderid":    order,
			"timeStr":    timeStr,
			"sign":       sign,
		}).Info("charge:游梦江湖安卓充值请求")
}

func youMengSign(cpOrder string, order string, appId int64, uid int64, money int32, timeInt int64, key string) (sign string) {
	allStr := fmt.Sprintf("%s%s%d%d%d%d%s", cpOrder, order, appId, uid, money, timeInt, key)
	hw := md5.Sum([]byte(allStr))
	return fmt.Sprintf("%x", hw)
}
