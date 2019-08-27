package piaomiao

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
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type PiaoMiaoRequest struct {
	OutTradeNo string  `form:"out_trade_no" json:"out_trade_no"`
	Price      float32 `form:"price" json:"price"`
	PayStatus  int32   `form:"pay_status" json:"pay_status"`
	Extend     string  `form:"extend" json:"extend"`
	Signtype   string  `form:"signType:" json:"signType:"`
	Sign       string  `form:"sign" json:"sign"`
}

func handlePiaoMiaoIOS(rw http.ResponseWriter, req *http.Request) {
	reqForm := &PiaoMiaoRequest{}
	err := httputils.Bind(req, reqForm)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:缥缈ios充值请求，参数解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.WithFields(
		log.Fields{
			"ip":         req.RemoteAddr,
			"OutTradeNo": reqForm.OutTradeNo,
			"Price":      reqForm.Price,
			"PayStatus":  reqForm.PayStatus,
			"Extend":     reqForm.Extend,
			"Signtype":   reqForm.Signtype,
			"sign":       reqForm.Sign,
		}).Info("charge:缥缈ios充值请求")

	sdkType := logintypes.SDKTypePiaoMiao
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:缥缈ios充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	piaoMiaoConfig, ok := sdkConfig.(*sdksdk.PiaoMiaoConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:缥缈ios充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeIOS
	gameKey := piaoMiaoConfig.GetGameKey(devicePlatformType)

	if reqForm.PayStatus != 1 {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"OutTradeNo": reqForm.OutTradeNo,
				"Price":      reqForm.Price,
				"PayStatus":  reqForm.PayStatus,
				"Extend":     reqForm.Extend,
				"Signtype":   reqForm.Signtype,
				"sign":       reqForm.Sign,
				"gameKey":    gameKey,
			}).Warn("charge:缥缈ios充值请求,失败不处理")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO 验证签名
	getSign := GetPiaoMiaoSign(reqForm, gameKey)
	if reqForm.Sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"OutTradeNo": reqForm.OutTradeNo,
				"Price":      reqForm.Price,
				"PayStatus":  reqForm.PayStatus,
				"Extend":     reqForm.Extend,
				"Signtype":   reqForm.Signtype,
				"sign":       reqForm.Sign,
				"gameKey":    gameKey,
			}).Warn("charge:缥缈ios充值请求,签名错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	now := timeutils.TimeToMillisecond(time.Now())
	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(reqForm.Extend, reqForm.OutTradeNo, logintypes.SDKTypePiaoMiao, int32(reqForm.Price), "", now)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"OutTradeNo": reqForm.OutTradeNo,
				"Price":      reqForm.Price,
				"PayStatus":  reqForm.PayStatus,
				"Extend":     reqForm.Extend,
				"Signtype":   reqForm.Signtype,
				"sign":       reqForm.Sign,
				"error":      err,
			}).Error("charge:缥缈ios请求,错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"OutTradeNo": reqForm.OutTradeNo,
				"Price":      reqForm.Price,
				"PayStatus":  reqForm.PayStatus,
				"Extend":     reqForm.Extend,
				"Signtype":   reqForm.Signtype,
				"sign":       reqForm.Sign,
			}).Warn("charge:缥缈ios请求,订单不存在")
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

	result := "success"
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(result))
	log.WithFields(
		log.Fields{
			"ip":         req.RemoteAddr,
			"OutTradeNo": reqForm.OutTradeNo,
			"Price":      reqForm.Price,
			"PayStatus":  reqForm.PayStatus,
			"Extend":     reqForm.Extend,
			"Signtype":   reqForm.Signtype,
			"sign":       reqForm.Sign,
		}).Info("charge:缥缈ios充值请求")
}

func GetPiaoMiaoSign(reqForm *PiaoMiaoRequest, gameKey string) (sign string) {
	price := fmt.Sprintf("%.2f", reqForm.Price)
	payStatus := fmt.Sprintf("%d", reqForm.PayStatus)

	signMap := make(map[string]string)
	signMap["outTradeNo"] = reqForm.OutTradeNo
	signMap["price"] = price
	signMap["payStatus"] = payStatus
	signMap["extend"] = reqForm.Extend
	signMap["gameKey"] = gameKey

	keyList := []string{"outTradeNo", "price", "payStatus", "extend", "gameKey"}
	allStr := ""
	for _, key := range keyList {
		keyValue := fmt.Sprintf("%s", signMap[key])
		allStr += keyValue
	}
	fmt.Println(allStr)
	hw := md5.Sum([]byte(allStr))
	return fmt.Sprintf("%x", hw)
}
