package zuowan

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/pkg/timeutils"
	"sort"
	"strings"
	"time"

	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/xozrc/pkg/httputils"
)

const (
	zuoWanPath = "/zuowan"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(zuoWanPath).Subrouter()
	sr.Path("/android").Handler(http.HandlerFunc(handleZuoWanAndroid))
	sr.Path("/ios").Handler(http.HandlerFunc(handleZuoWanIOS))
}

type ZuoWanForm struct {
	OrdNo                     string  `form:"ordno" json:"ordno"`
	OrdGoods                  string  `form:"ordgoods" json:"ordgoods"`
	OrdGoodsDescription       string  `form:"ordgoodsdescription" json:"ordgoodsdescription"`
	TradeStatus               int32   `form:"tradestatus" json:"tradestatus"`
	Amount                    float64 `form:"amount" json:"amount"`
	PaymentType               string  `form:"paymenttype" json:"paymenttype"`
	OrdTime                   string  `form:"ordtime" json:"ordtime"`
	PayTime                   string  `form:"paytime" json:"paytime"`
	GameId                    string  `form:"gameid" json:"gameid"`
	PayTradeId                string  `form:"paytradeid" json:"paytradeid"`
	GameCallbackDeliveryValue string  `form:"gamecallbackdeliveryvalue" json:"gamecallbackdeliveryvalue"`
	Sid                       string  `form:"sid" json:"sid"`
	Sign                      string  `form:"sign" json:"sign"`
}

func handleZuoWanAndroid(rw http.ResponseWriter, req *http.Request) {

	form := &ZuoWanForm{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:佐玩安卓充值请求，参数解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	ordNo := form.OrdNo
	ordGoods := form.OrdGoods
	ordGoodsDescription := form.OrdGoodsDescription
	tradeStatus := form.TradeStatus
	amount := form.Amount
	paymentType := form.PaymentType
	ordTime := form.OrdTime
	payTime := form.PayTime
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
			"payTime":                   payTime,
			"gameId":                    gameId,
			"payTradeId":                payTradeId,
			"gameCallbackDeliveryValue": gameCallbackDeliveryValue,
			"sid":  sid,
			"sign": sign,
		}).Info("charge:佐玩安卓充值请求")

	//交易失败
	if tradeStatus != 1 {
		result := "订单交易失败"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	money := int32(amount)
	receiveTime := timeutils.TimeToMillisecond(time.Now())

	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:佐玩安卓充值请求,sdk配置为空")
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
			}).Warn("charge:佐玩安卓充值请求,sdk配置强制转换失败")
		result := "服务器错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	devicePlatformType := logintypes.DevicePlatformTypeAndroid
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
				"payTime":                   payTime,
				"gameId":                    gameId,
				"payTradeId":                payTradeId,
				"gameCallbackDeliveryValue": gameCallbackDeliveryValue,
				"sid":     sid,
				"sign":    sign,
				"getSign": getSign,
				"err":     err,
			}).Warn("charge:佐玩安卓充值请求,签名错误")
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
			}).Error("charge:佐玩安卓请求,订单处理错误")
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
			}).Warn("charge:佐玩安卓请求,订单不存在")
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
		}).Info("charge:佐玩安卓充值请求,成功")
}

func convertZuoWanDataToMap(reqForm *ZuoWanForm) (dataMap map[string]string) {
	tradeStr := fmt.Sprintf("%d", reqForm.TradeStatus)
	amountStr := fmt.Sprintf("%.2f", reqForm.Amount)
	dataMap = make(map[string]string)
	dataMap["ordno"] = reqForm.OrdNo
	dataMap["ordgoods"] = reqForm.OrdGoods
	dataMap["ordgoodsdescription"] = reqForm.OrdGoodsDescription
	dataMap["tradestatus"] = tradeStr
	dataMap["amount"] = amountStr
	dataMap["paymenttype"] = reqForm.PaymentType
	dataMap["ordtime"] = reqForm.OrdTime
	dataMap["paytime"] = reqForm.PayTime
	dataMap["gameid"] = reqForm.GameId
	dataMap["paytradeid"] = reqForm.PayTradeId
	dataMap["gamecallbackdeliveryvalue"] = reqForm.GameCallbackDeliveryValue
	dataMap["sid"] = reqForm.Sid
	return
}

func GetZuoWanSign(dataStrMap map[string]string, payKey string) (sign string) {

	dataKeyList := make([]string, 0, 8)
	for key, _ := range dataStrMap {
		dataKeyList = append(dataKeyList, key)
	}
	sort.Sort(sort.StringSlice(dataKeyList))

	allStr := ""
	for _, key := range dataKeyList {
		keyValue := fmt.Sprintf("%s=%s&", key, dataStrMap[key])
		allStr += keyValue
	}
	if len(allStr) > 0 {
		allStr = allStr[:len(allStr)-1]
	}
	fmt.Println(allStr)
	hmac := hmac.New(md5.New, []byte(payKey))
	hmac.Write([]byte(allStr))
	sign = hex.EncodeToString(hmac.Sum([]byte("")))
	return strings.ToUpper(sign)
}
