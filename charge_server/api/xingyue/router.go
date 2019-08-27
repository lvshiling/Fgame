package xingyue

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/pkg/timeutils"
	"time"

	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/xozrc/pkg/httputils"
)

const (
	xingYuePath = "/xingyue"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(xingYuePath).Subrouter()
	sr.Path("/android").Handler(http.HandlerFunc(handleXingYueAndroid))
	sr.Path("/ios").Handler(http.HandlerFunc(handleXingYueIOS))
}

type XingYueForm struct {
	AsyxOrderId  string `form:"asyx_order_id" json:"asyx_order_id"`
	Subject      string `form:"subject" json:"subject"`
	SubjectDesc  string `form:"subject_desc" json:"subject_desc"`
	TradeStatus  string `form:"trade_status" json:"trade_status"`
	Amount       string `form:"amount" json:"amount"`
	Channel      string `form:"channel" json:"channel"`
	OrderCreatdt string `form:"order_creatdt" json:"order_creatdt"`
	OrderPaydt   string `form:"order_paydt" json:"order_paydt"`
	AsyxGameId   string `form:"asyx_game_id" json:"asyx_game_id"`
	PayOrderId   string `form:"pay_order_id" json:"pay_order_id"`
	Sign         string `form:"sign" json:"sign"`
	GameZero     string `form:"GameZero" json:"GameZero"`
	Memo         string `form:"Memo" json:"Memo"`
	Uid          string `form:"uid" json:"uid"`
}

func handleXingYueAndroid(rw http.ResponseWriter, req *http.Request) {

	form := &XingYueForm{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:星月安卓充值请求，参数解析错误")
		rw.WriteHeader(http.StatusBadRequest)
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
	sdkType := logintypes.SDKTypeXingYue
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
		}).Info("charge:星月安卓充值请求")

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
			}).Warn("charge:星月安卓充值请求，解析错误")
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
			}).Warn("charge:星月安卓充值请求,sdk配置为空")
		result := "服务器错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}
	xingYueConfig, ok := sdkConfig.(*sdksdk.XingYueConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:星月安卓充值请求,sdk配置强制转换失败")
		result := "服务器错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	devicePlatformType := logintypes.DevicePlatformTypeAndroid
	payKey := xingYueConfig.GetPayKey(devicePlatformType)

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
			}).Warn("charge:星月安卓充值请求,签名错误")
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
			}).Error("charge:星月安卓请求,订单处理错误")
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
			}).Warn("charge:星月安卓请求,订单不存在")
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
		}).Info("charge:星月安卓充值请求,成功")
}

func convertXingYueDataToMap(data *XingYueForm) (dataMap map[string]string) {
	dataMap = make(map[string]string)
	dataMap["asyx_order_id"] = data.AsyxOrderId
	dataMap["subject"] = data.Subject
	dataMap["subject_desc"] = data.SubjectDesc
	dataMap["trade_status"] = data.TradeStatus
	dataMap["amount"] = data.Amount
	dataMap["channel"] = data.Channel
	dataMap["order_creatdt"] = data.OrderCreatdt
	dataMap["order_paydt"] = data.OrderPaydt
	dataMap["asyx_game_id"] = data.AsyxGameId
	dataMap["pay_order_id"] = data.PayOrderId
	return
}

func GetSign(reqForm *XingYueForm, payKey string) (sign string) {
	allStr := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s",
		reqForm.AsyxOrderId,
		reqForm.Subject,
		reqForm.SubjectDesc,
		reqForm.TradeStatus,
		reqForm.Amount,
		reqForm.Channel,
		reqForm.OrderCreatdt,
		reqForm.OrderPaydt,
		reqForm.AsyxGameId,
		reqForm.PayOrderId)
	fmt.Println(allStr)
	hmac := hmac.New(md5.New, []byte(payKey))
	hmac.Write([]byte(allStr))
	sign = hex.EncodeToString(hmac.Sum([]byte("")))
	return
}
