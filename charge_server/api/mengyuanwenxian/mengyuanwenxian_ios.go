package mengyuanwenxian

import (
	"crypto/md5"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/pkg/timeutils"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type MengYuanWenXianRequest struct {
	AppId        string `form:"app_id" json:"app_id"`
	CpOrderId    string `form:"cp_order_id" json:"cp_order_id"`
	MemId        string `form:"mem_id" json:"mem_id"`
	OrderId      string `form:"order_id" json:"order_id"`
	OrderStatus  string `form:"order_status" json:"order_status"`
	PayTime      string `form:"pay_time" json:"pay_time"`
	ProductId    string `form:"product_id" json:"product_id"`
	ProductName  string `form:"product_name" json:"product_name"`
	ProductPrice string `form:"product_price" json:"product_price"`
	Sign         string `form:"sign" json:"sign"`
	Ext          string `form:"ext" json:"ext"`
}

func handleMengYuanWenXianIOS(rw http.ResponseWriter, req *http.Request) {
	reqForm := &MengYuanWenXianRequest{}
	err := httputils.Bind(req, reqForm)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:飞扬梦缘问仙ios充值请求，参数解析错误")
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
			}).Warn("charge:飞扬梦缘问仙ios充值请求,订单未充值成功")
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
			}).Warn("charge:飞扬梦缘问仙ios充值请求，商品价格解析错误")
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
		}).Info("charge:飞扬梦缘问仙ios充值请求")

	sdkType := logintypes.SDKTypeMengYuanWenXian
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:飞扬梦缘问仙ios充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	feiYnagConfig, ok := sdkConfig.(*sdksdk.FeiYangMengYuanWenXianConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:飞扬梦缘问仙ios充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeIOS
	appKey := feiYnagConfig.GetAppKey(devicePlatformType)

	//TODO 验证签名
	getSign := GetMengYuanWenXianSign(reqForm, appKey)
	if reqForm.Sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":           req.RemoteAddr,
				"appId":        reqForm.AppId,
				"appKey":       appKey,
				"cpOrderId":    reqForm.CpOrderId,
				"memId":        reqForm.MemId,
				"orderId":      reqForm.OrderId,
				"orderStatus":  reqForm.OrderStatus,
				"payTime":      reqForm.PayTime,
				"productId":    reqForm.ProductId,
				"productName":  reqForm.ProductName,
				"productPrice": reqForm.ProductPrice,
				"sign":         reqForm.Sign,
			}).Warn("charge:飞扬梦缘问仙ios充值请求,签名错误")
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
			}).Error("charge:飞扬梦缘问仙ios请求,错误")
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
			}).Warn("charge:飞扬梦缘问仙ios请求,订单不存在")
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
		}).Info("charge:飞扬梦缘问仙ios充值请求")
}

func GetMengYuanWenXianSign(reqForm *MengYuanWenXianRequest, appKey string) (sign string) {

	signMap := make(map[string]string)
	signMap["app_id"] = reqForm.AppId
	signMap["cp_order_id"] = reqForm.CpOrderId
	signMap["mem_id"] = reqForm.MemId
	signMap["order_id"] = reqForm.OrderId
	signMap["order_status"] = reqForm.OrderStatus
	signMap["pay_time"] = reqForm.PayTime
	signMap["product_id"] = reqForm.ProductId
	signMap["product_name"] = url.QueryEscape(reqForm.ProductName)
	signMap["product_price"] = reqForm.ProductPrice

	signMap["app_key"] = appKey
	keyList := []string{"app_id", "cp_order_id", "mem_id", "order_id", "order_status", "pay_time", "product_id", "product_name", "product_price", "app_key"}
	allStr := ""
	for _, key := range keyList {
		keyValue := fmt.Sprintf("%s=%s&", key, signMap[key])
		allStr += keyValue
	}
	if len(allStr) > 0 {
		allStr = allStr[:len(allStr)-1]
	}
	hw := md5.Sum([]byte(allStr))
	return fmt.Sprintf("%x", hw)
}
