package lingmeng

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
	"github.com/xozrc/pkg/httputils"
)

const (
	lingmengPath = "/lingmeng"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(lingmengPath).Subrouter()
	sr.Path("/ios").Handler(http.HandlerFunc(handleLingMengIOS))
	sr.Path("/android").Handler(http.HandlerFunc(handleLingMengAndroid))
}

type LingMengForm struct {
	OrderId     string `form:"order_id" json:"order_id"`
	MemId       string `form:"mem_id" json:"mem_id"`
	AppId       string `form:"app_id" json:"app_id"`
	Money       string `form:"money" json:"money"`
	OrderStatus string `form:"order_status" json:"order_status"`
	Paytime     string `form:"paytime" json:"paytime"`
	Attach      string `form:"attach" json:"attach"`
	Sign        string `form:"sign" json:"sign"`
}

func handleLingMengAndroid(rw http.ResponseWriter, req *http.Request) {
	resultErr := "FAILURE"

	// defer req.Body.Close()
	// content, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"ip":    req.RemoteAddr,
	// 			"error": err,
	// 		}).Warn("charge:灵梦仙界安卓充值请求，解析错误")
	// 	rw.WriteHeader(http.StatusBadRequest)
	// 	rw.Write([]byte(resultErr))
	// 	return
	// }
	// form := &LingMengForm{}
	// err = json.Unmarshal(content, form)
	// if err != nil {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"ip":      req.RemoteAddr,
	// 			"content": string(content),
	// 			"error":   err,
	// 		}).Warn("charge:灵梦仙界安卓充值请求，解析错误")
	// 	rw.WriteHeader(http.StatusOK)
	// 	rw.Write([]byte(resultErr))
	// 	return
	// }
	form := &LingMengForm{}
	err := req.ParseForm()
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:灵梦仙界安卓充值请求，解析错误")

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
			}).Warn("charge:灵梦仙界安卓充值请求，解析错误")

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
		}).Info("charge:灵梦仙界安卓充值请求")

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
			}).Warn("charge:灵梦仙界安卓充值请求，解析错误")
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
			}).Warn("charge:灵梦仙界安卓充值请求，解析错误")
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
			}).Warn("charge:灵梦仙界安卓充值请求,sdk配置为空")
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
			}).Warn("charge:灵梦仙界安卓充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(resultErr))
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeAndroid
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
			}).Warn("charge:灵梦仙界安卓充值请求,签名错误")
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
			}).Error("charge:灵梦仙界安卓请求,错误")
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
			}).Warn("charge:灵梦仙界安卓请求,订单不存在")
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
		}).Info("charge:灵梦仙界安卓充值请求")
}

func lingMengSign(orderId, memId, appId, money, orderStatus, payTime, attach, appKey string) (sign string) {
	allStr := fmt.Sprintf("order_id=%s&mem_id=%s&app_id=%s&money=%s&order_status=%s&paytime=%s&attach=%s&app_key=%s",
		orderId, memId, appId, money, orderStatus, payTime, attach, appKey)
	hw := md5.Sum([]byte(allStr))
	return fmt.Sprintf("%x", hw)
}
