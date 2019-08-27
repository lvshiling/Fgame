package zuowan

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
	"github.com/xozrc/pkg/httputils"
)

func handleZuoWanIOS(rw http.ResponseWriter, req *http.Request) {

	form := &ZuoWanForm{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:佐玩ios充值请求，参数解析错误")
		result := "参数解析错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	ordNo := form.OrdNo
	ordGoods := form.OrdGoods
	ordGoodsDescription := form.OrdGoodsDescription
	tradeStatus := form.TradeStatus
	amount := form.Amount
	paymentType := form.PaymentType
	ordTime := form.OrdTime
	gameId := form.GameId
	payTradeId := form.PayTradeId
	gameCallbackDeliveryValue := form.GameCallbackDeliveryValue
	sid := form.Sid
	sign := form.Sign
	sdkType := logintypes.SDKTypeZuoWan
	log.WithFields(
		log.Fields{
			"ip":                        req.RemoteAddr,
			"ordNo":                     ordNo,
			"ordGoods":                  ordGoods,
			"ordGoodsDescription":       ordGoodsDescription,
			"tradeStatus":               tradeStatus,
			"amount":                    amount,
			"paymentType":               paymentType,
			"ordTime":                   ordTime,
			"gameId":                    gameId,
			"payTradeId":                payTradeId,
			"gameCallbackDeliveryValue": gameCallbackDeliveryValue,
			"sid": sid,
		}).Info("charge:佐玩ios充值请求")

	//交易失败
	if tradeStatus != 1 {
		result := "订单交易失败"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	money := int32(form.Amount)
	receiveTime := timeutils.TimeToMillisecond(time.Now())

	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:佐玩ios充值请求,sdk配置为空")
		result := "服务器错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}
	zuoWanConfig, ok := sdkConfig.(*sdksdk.ZuoWanConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:佐玩ios充值请求,sdk配置强制转换失败")
		result := "服务器错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	devicePlatformType := logintypes.DevicePlatformTypeIOS
	payKey := zuoWanConfig.GetPrivatekey(devicePlatformType)

	dataMap := convertZuoWanDataToMap(form)
	getSign := GetZuoWanSign(dataMap, payKey)
	if getSign != sign {
		log.WithFields(
			log.Fields{
				"ip":                        req.RemoteAddr,
				"ordNo":                     ordNo,
				"ordGoods":                  ordGoods,
				"ordGoodsDescription":       ordGoodsDescription,
				"tradeStatus":               tradeStatus,
				"amount":                    amount,
				"paymentType":               paymentType,
				"ordTime":                   ordTime,
				"gameId":                    gameId,
				"payTradeId":                payTradeId,
				"gameCallbackDeliveryValue": gameCallbackDeliveryValue,
				"sid": sid,
				"err": err,
			}).Warn("charge:佐玩ios充值请求,签名错误")
		result := "签名错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	//缺少游戏服的订单id
	obj, repeat, err := chargeService.OrderPay(gameCallbackDeliveryValue, ordNo, sdkType, money, sid, receiveTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":                        req.RemoteAddr,
				"ordNo":                     ordNo,
				"ordGoods":                  ordGoods,
				"ordGoodsDescription":       ordGoodsDescription,
				"tradeStatus":               tradeStatus,
				"amount":                    amount,
				"paymentType":               paymentType,
				"ordTime":                   ordTime,
				"gameId":                    gameId,
				"payTradeId":                payTradeId,
				"gameCallbackDeliveryValue": gameCallbackDeliveryValue,
				"sid":   sid,
				"error": err,
			}).Error("charge:佐玩ios请求,订单处理错误")
		result := "订单处理错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":                        req.RemoteAddr,
				"ordNo":                     ordNo,
				"ordGoods":                  ordGoods,
				"ordGoodsDescription":       ordGoodsDescription,
				"tradeStatus":               tradeStatus,
				"amount":                    amount,
				"paymentType":               paymentType,
				"ordTime":                   ordTime,
				"gameId":                    gameId,
				"payTradeId":                payTradeId,
				"gameCallbackDeliveryValue": gameCallbackDeliveryValue,
				"sid":   sid,
				"error": err,
			}).Warn("charge:佐玩ios请求,订单不存在")
		result := "订单不存在"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
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
			"ip":                        req.RemoteAddr,
			"ordNo":                     ordNo,
			"ordGoods":                  ordGoods,
			"ordGoodsDescription":       ordGoodsDescription,
			"tradeStatus":               tradeStatus,
			"amount":                    amount,
			"paymentType":               paymentType,
			"ordTime":                   ordTime,
			"gameId":                    gameId,
			"payTradeId":                payTradeId,
			"gameCallbackDeliveryValue": gameCallbackDeliveryValue,
			"sid": sid,
		}).Info("charge:佐玩ios充值请求,成功")
}
