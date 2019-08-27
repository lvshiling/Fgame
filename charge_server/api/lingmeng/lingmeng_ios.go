package lingmeng

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
	"github.com/xozrc/pkg/httputils"
)

func handleLingMengIOS(rw http.ResponseWriter, req *http.Request) {
	resultErr := "FAILURE"

	form := &LingMengForm{}
	err := req.ParseForm()
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:灵梦仙界ios充值请求，解析错误")

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(resultErr))
		return
	}

	err = httputils.BindForm(form, req.Form, nil)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:灵梦仙界ios充值请求，解析错误")

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(resultErr))
		return
	}
	orderId := form.OrderId
	memId := form.MemId
	appId := form.AppId
	money := form.Money
	orderStatus := form.OrderStatus
	payTime := form.Paytime
	attach := form.Attach
	sign := form.Sign
	log.WithFields(
		log.Fields{
			"ip":          req.RemoteAddr,
			"orderId":     orderId,
			"memId":       memId,
			"appId":       appId,
			"money":       money,
			"orderStatus": orderStatus,
			"payTime":     payTime,
			"attach":      attach,
			"sign":        sign,
		}).Info("charge:灵梦仙界ios充值请求")

	moneyFloat, err := strconv.ParseFloat(money, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"orderId":     orderId,
				"memId":       memId,
				"appId":       appId,
				"money":       money,
				"orderStatus": orderStatus,
				"payTime":     payTime,
				"attach":      attach,
				"sign":        sign,
				"error":       err,
			}).Warn("charge:灵梦仙界ios充值请求，解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(resultErr))
		return
	}
	moneyInt := int32(moneyFloat)

	receiveTime, err := strconv.ParseInt(payTime, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"orderId":     orderId,
				"memId":       memId,
				"appId":       appId,
				"money":       money,
				"orderStatus": orderStatus,
				"payTime":     payTime,
				"attach":      attach,
				"sign":        sign,
				"error":       err,
			}).Warn("charge:灵梦仙界ios充值请求，解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(resultErr))
		return
	}

	sdkType := logintypes.SDKTypeLingMeng
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:灵梦仙界ios充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(resultErr))
		return
	}
	lingMengConfig, ok := sdkConfig.(*sdksdk.LingMengConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:灵梦仙界ios充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(resultErr))
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeIOS
	appKey := lingMengConfig.GetAppKey(devicePlatformType)

	//TODO 验证签名
	getSign := lingMengSign(orderId, memId, appId, money, orderStatus, payTime, attach, appKey)
	if sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"orderId":     orderId,
				"memId":       memId,
				"appId":       appId,
				"money":       money,
				"orderStatus": orderStatus,
				"payTime":     payTime,
				"attach":      attach,
				"sign":        sign,
				"getSign":     getSign,
			}).Warn("charge:灵梦仙界ios充值请求,签名错误")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(resultErr))
		return
	}

	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(attach, orderId, logintypes.SDKTypeLingMeng, moneyInt, memId, receiveTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"orderId":     orderId,
				"memId":       memId,
				"appId":       appId,
				"money":       money,
				"orderStatus": orderStatus,
				"payTime":     payTime,
				"attach":      attach,
				"sign":        sign,
				"error":       err,
			}).Error("charge:灵梦仙界ios请求,错误")
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(resultErr))
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"orderId":     orderId,
				"memId":       memId,
				"appId":       appId,
				"money":       money,
				"orderStatus": orderStatus,
				"payTime":     payTime,
				"attach":      attach,
				"sign":        sign,
				"error":       err,
			}).Warn("charge:灵梦仙界ios请求,订单不存在")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(resultErr))
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
	result := "SUCCESS"
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(result))
	log.WithFields(
		log.Fields{
			"orderId":     orderId,
			"memId":       memId,
			"appId":       appId,
			"money":       money,
			"orderStatus": orderStatus,
			"payTime":     payTime,
			"attach":      attach,
			"sign":        sign,
		}).Info("charge:灵梦仙界ios充值请求")
}
