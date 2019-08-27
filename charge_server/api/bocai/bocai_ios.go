package bocai

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
	"github.com/xozrc/pkg/httputils"
)

func handleBoCaiIOS(rw http.ResponseWriter, req *http.Request) {

	form := &BoCaiForm{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:菠菜ios充值请求，参数解析错误")
		result := "参数解析错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	asyxOrderId := form.AsyxOrderId
	subject := form.Subject
	subjectDesc := form.SubjectDesc
	tradeStatus := form.TradeStatus
	amount := form.Amount
	channel := form.Channel
	orderCreatdt := form.OrderCreatdt
	orderPaydt := form.OrderPaydt
	asyxGameId := form.AsyxGameId
	payOrderId := form.PayOrderId
	sign := form.Sign
	gameZero := form.GameZero
	memo := form.Memo
	uid := form.Uid
	sdkType := logintypes.SDKTypeBoCai
	log.WithFields(
		log.Fields{
			"ip":           req.RemoteAddr,
			"asyxOrderId":  asyxOrderId,
			"subject":      subject,
			"subjectDesc":  subjectDesc,
			"tradeStatus":  tradeStatus,
			"amount":       amount,
			"channel":      channel,
			"orderCreatdt": orderCreatdt,
			"orderPaydt":   orderPaydt,
			"asyxGameId":   asyxGameId,
			"payOrderId":   payOrderId,
			"sign":         sign,
			"gameZero":     gameZero,
			"memo":         memo,
			"uid":          uid,
		}).Info("charge:菠菜ios充值请求")

	//交易失败
	if tradeStatus != "1" {
		result := "订单交易失败"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"asyxOrderId":  asyxOrderId,
				"subject":      subject,
				"subjectDesc":  subjectDesc,
				"tradeStatus":  tradeStatus,
				"amount":       amount,
				"channel":      channel,
				"orderCreatdt": orderCreatdt,
				"orderPaydt":   orderPaydt,
				"asyxGameId":   asyxGameId,
				"payOrderId":   payOrderId,
				"sign":         sign,
				"gameZero":     gameZero,
				"memo":         memo,
				"uid":          uid,
				"error":        err,
			}).Warn("charge:菠菜ios充值请求，解析错误")
		result := "参数解析错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}
	money := int32(amountFloat)
	receiveTime := timeutils.TimeToMillisecond(time.Now())

	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:菠菜ios充值请求,sdk配置为空")
		result := "服务器错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}
	boCaiConfig, ok := sdkConfig.(*sdksdk.BoCaiConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:菠菜ios充值请求,sdk配置强制转换失败")
		result := "服务器错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	devicePlatformType := logintypes.DevicePlatformTypeIOS
	payKey := boCaiConfig.GetPayKey(devicePlatformType)

	getSign := GetSign(form, payKey)
	if getSign != sign {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"asyxOrderId":  asyxOrderId,
				"subject":      subject,
				"subjectDesc":  subjectDesc,
				"tradeStatus":  tradeStatus,
				"amount":       amount,
				"channel":      channel,
				"orderCreatdt": orderCreatdt,
				"orderPaydt":   orderPaydt,
				"asyxGameId":   asyxGameId,
				"payOrderId":   payOrderId,
				"sign":         sign,
				"gameZero":     gameZero,
				"memo":         memo,
				"uid":          uid,
				"err":          err,
			}).Warn("charge:菠菜ios充值请求,签名错误")
		result := "签名错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	//缺少游戏服的订单id
	obj, repeat, err := chargeService.OrderPay(memo, asyxOrderId, sdkType, money, uid, receiveTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"asyxOrderId":  asyxOrderId,
				"subject":      subject,
				"subjectDesc":  subjectDesc,
				"tradeStatus":  tradeStatus,
				"amount":       amount,
				"channel":      channel,
				"orderCreatdt": orderCreatdt,
				"orderPaydt":   orderPaydt,
				"asyxGameId":   asyxGameId,
				"payOrderId":   payOrderId,
				"sign":         sign,
				"gameZero":     gameZero,
				"memo":         memo,
				"uid":          uid,
				"error":        err,
			}).Error("charge:菠菜ios请求,订单处理错误")
		result := "订单处理错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"asyxOrderId":  asyxOrderId,
				"subject":      subject,
				"subjectDesc":  subjectDesc,
				"tradeStatus":  tradeStatus,
				"amount":       amount,
				"channel":      channel,
				"orderCreatdt": orderCreatdt,
				"orderPaydt":   orderPaydt,
				"asyxGameId":   asyxGameId,
				"payOrderId":   payOrderId,
				"sign":         sign,
				"gameZero":     gameZero,
				"memo":         memo,
				"uid":          uid,
				"error":        err,
			}).Warn("charge:菠菜ios请求,订单不存在")
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
			"ip":           req.RemoteAddr,
			"asyxOrderId":  asyxOrderId,
			"subject":      subject,
			"subjectDesc":  subjectDesc,
			"tradeStatus":  tradeStatus,
			"amount":       amount,
			"channel":      channel,
			"orderCreatdt": orderCreatdt,
			"orderPaydt":   orderPaydt,
			"asyxGameId":   asyxGameId,
			"payOrderId":   payOrderId,
			"sign":         sign,
			"gameZero":     gameZero,
			"memo":         memo,
			"uid":          uid,
		}).Info("charge:菠菜ios充值请求,成功")
}
