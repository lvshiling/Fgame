package yiyun

import (
	"crypto/md5"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/xozrc/pkg/httputils"
)

const (
	yiyunPath = "/yiyun"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(yiyunPath).Subrouter()
	sr.Path("/ios").Handler(http.HandlerFunc(handleYiYunIOS))
	sr.Path("/android").Handler(http.HandlerFunc(handleYiYunAndroid))
}

type YiYunReqData struct {
	OrderId     string `json:"orderId" form:"orderId"`
	GameOrderId string `json:"gameOrderId" form:"gameOrderId"`
	ProductName string `json:"productName" form:"productName"`
	Money       int32  `json:"money" form:"money"`
	Ext         string `json:"ext" form:"ext"`
	Ugid        int64  `json:"ugid" form:"ugid"`
	Sign        string `json:"sign" form:"sign"`
}

func handleYiYunAndroid(rw http.ResponseWriter, req *http.Request) {

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
		}).Info("charge:亦云间安卓充值请求")

	sdkType := logintypes.SDKTypeYiYun
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:亦云间安卓充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	yiYunConfig, ok := sdkConfig.(*sdksdk.YiYunConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:亦云间安卓充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeAndroid
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
			}).Warn("charge:亦云间安卓充值请求,签名错误")
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
			}).Error("charge:亦云间安卓请求,错误")
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
			}).Warn("charge:亦云间安卓请求,订单不存在")
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
		}).Info("charge:亦云间安卓充值请求")
}

func yiYunSign(orderId, gameOrderId, productName, money, ext, ugid, key string) string {
	allStr := fmt.Sprintf("ext=%s&gameOrderId=%s&money=%s&orderId=%s&productName=%s&ugid=%s#%s", ext, gameOrderId, money, orderId, productName, ugid, key)
	hw := md5.Sum([]byte(allStr))
	return strings.ToUpper(fmt.Sprintf("%x", hw))
}
