package jiandaowenqing

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
	jianDaoPath = "/jiandao"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(jianDaoPath).Subrouter()
	sr.Path("/android").Handler(http.HandlerFunc(handleJianDaoAndroid))
	sr.Path("/ios").Handler(http.HandlerFunc(handleJianDaoIOS))
}

func handleJianDaoAndroid(rw http.ResponseWriter, req *http.Request) {
	errSign := "errSign"
	otherErr := "error"
	reqForm := &JianDaoRequest{}
	err := httputils.Bind(req, reqForm)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:剑道问情安卓充值请求，参数解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(otherErr))
		return
	}

	money, err := strconv.ParseFloat(reqForm.Amount, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"orderid":     reqForm.OrderId,
				"username":    reqForm.UserName,
				"appid":       reqForm.AppId,
				"roleid":      reqForm.RoleId,
				"serverid":    reqForm.ServerId,
				"amount":      reqForm.Amount,
				"paytime":     reqForm.PayTime,
				"attach":      reqForm.Attach,
				"productname": reqForm.ProductName,
				"sign":        reqForm.Sign,
			}).Warn("charge:剑道问情安卓充值请求，商品价格解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(otherErr))
		return
	}

	sdkType := logintypes.SDKTypeJianDao
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:剑道问情安卓充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(otherErr))
		return
	}
	jianDaoConfig, ok := sdkConfig.(*sdksdk.JianDaoConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:剑道问情安卓充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(otherErr))
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeAndroid
	appKey := jianDaoConfig.GetAppKey(devicePlatformType)

	getSign := GetJianDao(reqForm, appKey)

	if reqForm.Sign != getSign {
		log.WithFields(
			log.Fields{
				"orderid":     reqForm.OrderId,
				"username":    reqForm.UserName,
				"appid":       reqForm.AppId,
				"roleid":      reqForm.RoleId,
				"serverid":    reqForm.ServerId,
				"amount":      reqForm.Amount,
				"paytime":     reqForm.PayTime,
				"attach":      reqForm.Attach,
				"productname": reqForm.ProductName,
				"appkey":      appKey,
				"sign":        reqForm.Sign,
			}).Warn("charge:剑道问情安卓充值请求,签名错误")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(errSign))
		return
	}

	now := timeutils.TimeToMillisecond(time.Now())
	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(reqForm.Attach, reqForm.OrderId, logintypes.SDKTypeJianDao, int32(money), reqForm.AppId, now)
	if err != nil {
		log.WithFields(
			log.Fields{
				"orderid":     reqForm.OrderId,
				"username":    reqForm.UserName,
				"appid":       reqForm.AppId,
				"roleid":      reqForm.RoleId,
				"serverid":    reqForm.ServerId,
				"amount":      reqForm.Amount,
				"paytime":     reqForm.PayTime,
				"attach":      reqForm.Attach,
				"productname": reqForm.ProductName,
				"appkey":      appKey,
				"sign":        reqForm.Sign,
			}).Error("charge:剑道问情安卓请求,错误")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(otherErr))
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"orderid":     reqForm.OrderId,
				"username":    reqForm.UserName,
				"appid":       reqForm.AppId,
				"roleid":      reqForm.RoleId,
				"serverid":    reqForm.ServerId,
				"amount":      reqForm.Amount,
				"paytime":     reqForm.PayTime,
				"attach":      reqForm.Attach,
				"productname": reqForm.ProductName,
				"appkey":      appKey,
				"sign":        reqForm.Sign,
			}).Warn("charge:剑道问情安卓请求,订单不存在")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(otherErr))
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
			"orderid":     reqForm.OrderId,
			"username":    reqForm.UserName,
			"appid":       reqForm.AppId,
			"roleid":      reqForm.RoleId,
			"serverid":    reqForm.ServerId,
			"amount":      reqForm.Amount,
			"paytime":     reqForm.PayTime,
			"attach":      reqForm.Attach,
			"productname": reqForm.ProductName,
			"appkey":      appKey,
			"sign":        reqForm.Sign,
		}).Info("charge:剑道问情安卓充值请求")
}
