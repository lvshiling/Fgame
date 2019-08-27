package feiyang

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
	"github.com/gorilla/mux"
	"github.com/xozrc/pkg/httputils"
)

const (
	feiyangPath = "/feiyang"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(feiyangPath).Subrouter()
	sr.Path("/android").Handler(http.HandlerFunc(handleFeiYangAndroid))
	sr.Path("/ios").Handler(http.HandlerFunc(handleFeiYangIOS))
}

func handleFeiYangAndroid(rw http.ResponseWriter, req *http.Request) {

	reqForm := &FeiYangRequest{}
	err := httputils.Bind(req, reqForm)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:飞扬ios充值请求，参数解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if reqForm.OrderStatus != "2" {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"appId":        reqForm.AppId,
				"cpOrderId":    reqForm.CpOrderId,
				"memId":        reqForm.MemId,
				"orderId":      reqForm.OrderId,
				"orderStatus":  reqForm.OrderStatus,
				"payTime":      reqForm.PayTime,
				"productId":    reqForm.ProductId,
				"productName":  reqForm.ProductName,
				"productPrice": reqForm.ProductPrice,
				"sign":         reqForm.Sign,
			}).Warn("charge:飞扬ios充值请求,订单未充值成功")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	money, err := strconv.ParseFloat(reqForm.ProductPrice, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"appId":        reqForm.AppId,
				"cpOrderId":    reqForm.CpOrderId,
				"memId":        reqForm.MemId,
				"orderId":      reqForm.OrderId,
				"orderStatus":  reqForm.OrderStatus,
				"payTime":      reqForm.PayTime,
				"productId":    reqForm.ProductId,
				"productName":  reqForm.ProductName,
				"productPrice": reqForm.ProductPrice,
				"sign":         reqForm.Sign,
				"err":          err,
			}).Warn("charge:飞扬ios充值请求，商品价格解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.WithFields(
		log.Fields{
			"ip":           req.RemoteAddr,
			"appId":        reqForm.AppId,
			"cpOrderId":    reqForm.CpOrderId,
			"memId":        reqForm.MemId,
			"orderId":      reqForm.OrderId,
			"orderStatus":  reqForm.OrderStatus,
			"payTime":      reqForm.PayTime,
			"productId":    reqForm.ProductId,
			"productName":  reqForm.ProductName,
			"productPrice": reqForm.ProductPrice,
			"sign":         reqForm.Sign,
		}).Info("charge:飞扬ios充值请求")

	sdkType := logintypes.SDKTypeFeiYang
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:飞扬安卓充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	feiYangConfig, ok := sdkConfig.(*sdksdk.FeiYangConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:飞扬安卓充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeAndroid
	appKey := feiYangConfig.GetAppKey(devicePlatformType)

	//TODO 验证签名
	getSign := GetFeiYangSign(reqForm, appKey)
	if reqForm.Sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"appId":        reqForm.AppId,
				"cpOrderId":    reqForm.CpOrderId,
				"memId":        reqForm.MemId,
				"orderId":      reqForm.OrderId,
				"orderStatus":  reqForm.OrderStatus,
				"payTime":      reqForm.PayTime,
				"productId":    reqForm.ProductId,
				"productName":  reqForm.ProductName,
				"productPrice": reqForm.ProductPrice,
				"sign":         reqForm.Sign,
			}).Warn("charge:飞扬安卓充值请求,签名错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	now := timeutils.TimeToMillisecond(time.Now())
	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(reqForm.CpOrderId, reqForm.OrderId, logintypes.SDKTypeFeiYang, int32(money), reqForm.MemId, now)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"appId":        reqForm.AppId,
				"cpOrderId":    reqForm.CpOrderId,
				"memId":        reqForm.MemId,
				"orderId":      reqForm.OrderId,
				"orderStatus":  reqForm.OrderStatus,
				"payTime":      reqForm.PayTime,
				"productId":    reqForm.ProductId,
				"productName":  reqForm.ProductName,
				"productPrice": reqForm.ProductPrice,
				"sign":         reqForm.Sign,
				"error":        err,
			}).Error("charge:飞扬安卓请求,错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"appId":        reqForm.AppId,
				"cpOrderId":    reqForm.CpOrderId,
				"memId":        reqForm.MemId,
				"orderId":      reqForm.OrderId,
				"orderStatus":  reqForm.OrderStatus,
				"payTime":      reqForm.PayTime,
				"productId":    reqForm.ProductId,
				"productName":  reqForm.ProductName,
				"productPrice": reqForm.ProductPrice,
				"sign":         reqForm.Sign,
			}).Warn("charge:飞扬安卓请求,订单不存在")
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

	result := "SUCCESS"
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(result))
	log.WithFields(
		log.Fields{
			"ip":           req.RemoteAddr,
			"appId":        reqForm.AppId,
			"cpOrderId":    reqForm.CpOrderId,
			"memId":        reqForm.MemId,
			"orderId":      reqForm.OrderId,
			"orderStatus":  reqForm.OrderStatus,
			"payTime":      reqForm.PayTime,
			"productId":    reqForm.ProductId,
			"productName":  reqForm.ProductName,
			"productPrice": reqForm.ProductPrice,
			"sign":         reqForm.Sign,
		}).Info("charge:飞扬安卓充值请求")
}
