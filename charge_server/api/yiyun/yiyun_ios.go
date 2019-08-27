package yiyun

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

func handleYiYunIOS(rw http.ResponseWriter, req *http.Request) {

	form := &YiYunReqData{}
	err := req.ParseForm()
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge: 亦云充值请求，参数解析错误")

		rw.Write([]byte("解析失败"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = httputils.BindForm(form, req.Form, nil)

	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge: 亦云充值请求，参数解析错误")
		rw.Write([]byte("解析失败"))

		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	orderId := form.OrderId
	gameOrderId := form.GameOrderId
	productName := form.ProductName
	moneyInt := form.Money
	ext := form.Ext
	ugidInt := form.Ugid
	sign := form.Sign
	ugid := fmt.Sprintf("%d", ugidInt)
	money := fmt.Sprintf("%d", moneyInt)
	log.WithFields(
		log.Fields{
			"ip":          req.RemoteAddr,
			"orderId":     orderId,
			"gameOrderId": gameOrderId,
			"productName": productName,
			"money":       money,
			"ext":         ext,
			"ugid":        ugid,
			"sign":        sign,
		}).Info("charge:亦云间ios充值请求")

	sdkType := logintypes.SDKTypeYiYun
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:亦云间ios充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	yiYunConfig, ok := sdkConfig.(*sdksdk.YiYunConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:亦云间ios充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeIOS
	chargeKey := yiYunConfig.GetKey(devicePlatformType)

	//TODO 验证签名
	getSign := yiYunSign(orderId, gameOrderId, productName, money, ext, ugid, chargeKey)
	if sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"orderId":     orderId,
				"gameOrderId": gameOrderId,
				"productName": productName,
				"money":       money,
				"ext":         ext,
				"ugid":        ugid,
				"sign":        sign,
				"getSign":     getSign,
			}).Warn("charge:亦云间ios充值请求,签名错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	receiveTime := int64(0)
	// receiveTime 未知
	obj, repeat, err := chargeService.OrderPay(gameOrderId, orderId, logintypes.SDKTypeYiYun, moneyInt/100, ugid, receiveTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"orderId":     orderId,
				"gameOrderId": gameOrderId,
				"productName": productName,
				"money":       money,
				"ext":         ext,
				"ugid":        ugid,
				"sign":        sign,
				"error":       err,
			}).Error("charge:亦云间IOS请求,错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"orderId":     orderId,
				"gameOrderId": gameOrderId,
				"productName": productName,
				"money":       money,
				"ext":         ext,
				"ugid":        ugid,
				"sign":        sign,
				"error":       err,
			}).Warn("charge:亦云间IOS请求,订单不存在")
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
	result := "ok"
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(result))
	log.WithFields(
		log.Fields{
			"orderId":     orderId,
			"gameOrderId": gameOrderId,
			"productName": productName,
			"money":       money,
			"ext":         ext,
			"ugid":        ugid,
			"sign":        sign,
		}).Info("charge:亦云间IOS充值请求")
}
