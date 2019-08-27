package qia

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

type QiARequest struct {
	OrderId  string `form:"orderid" json:"orderid"`
	UserName string `form:"username" json:"username"`
	GameId   int32  `form:"gameid" json:"gameid"`
	RoleId   string `form:"roleid" json:"roleid"`
	ServerId int32  `form:"serverid" json:"serverid"`
	PayType  string `form:"paytype" json:"paytype"`
	Amount   int32  `form:"amount" json:"amount"`
	PayTime  int64  `form:"paytime" json:"paytime"`
	Attache  string `form:"attach" json:"attach"`
	Sign     string `form:"sign" json:"sign"`
}

func handleQiAIOS(rw http.ResponseWriter, req *http.Request) {
	reqForm := &QiARequest{}
	err := httputils.Bind(req, reqForm)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:7Aios充值请求，参数解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.WithFields(
		log.Fields{
			"ip":       req.RemoteAddr,
			"orderId":  reqForm.OrderId,
			"userName": reqForm.UserName,
			"gameId":   reqForm.GameId,
			"roleId":   reqForm.RoleId,
			"serverId": reqForm.ServerId,
			"payType":  reqForm.PayType,
			"amount":   reqForm.Amount,
			"payTime":  reqForm.PayTime,
			"attache":  reqForm.Attache,
			"sign":     reqForm.Sign,
		}).Info("charge:7Aios充值请求")

	sdkType := logintypes.SDKTypeQiA
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:7Aios充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	qiAConfig, ok := sdkConfig.(*sdksdk.QiAConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:7Aios充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeIOS
	appKey := qiAConfig.GetAppKey(devicePlatformType)

	//TODO 验证签名
	getSign := GetQiASign(reqForm, appKey)
	if reqForm.Sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserName,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"payType":  reqForm.PayType,
				"amount":   reqForm.Amount,
				"payTime":  reqForm.PayTime,
				"attache":  reqForm.Attache,
				"sign":     reqForm.Sign,
				"appKey":   appKey,
			}).Warn("charge:7Aios充值请求,签名错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	now := timeutils.TimeToMillisecond(time.Now())
	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(reqForm.Attache, reqForm.OrderId, logintypes.SDKTypeQiA, reqForm.Amount, reqForm.RoleId, now)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserName,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"payType":  reqForm.PayType,
				"amount":   reqForm.Amount,
				"payTime":  reqForm.PayTime,
				"attache":  reqForm.Attache,
				"sign":     reqForm.Sign,
				"error":    err,
			}).Error("charge:7Aios请求,错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserName,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"payType":  reqForm.PayType,
				"amount":   reqForm.Amount,
				"payTime":  reqForm.PayTime,
				"attache":  reqForm.Attache,
				"sign":     reqForm.Sign,
			}).Warn("charge:7Aios请求,订单不存在")
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
			"ip":       req.RemoteAddr,
			"orderId":  reqForm.OrderId,
			"userName": reqForm.UserName,
			"gameId":   reqForm.GameId,
			"roleId":   reqForm.RoleId,
			"serverId": reqForm.ServerId,
			"payType":  reqForm.PayType,
			"amount":   reqForm.Amount,
			"payTime":  reqForm.PayTime,
			"attache":  reqForm.Attache,
			"sign":     reqForm.Sign,
		}).Info("charge:7Aios充值请求")
}

func GetQiASign(reqForm *QiARequest, appKey string) (sign string) {
	gameIdStr := fmt.Sprintf("%d", reqForm.GameId)
	serverIdStr := fmt.Sprintf("%d", reqForm.ServerId)
	amountStr := fmt.Sprintf("%d", reqForm.Amount)
	payTimeStr := fmt.Sprintf("%d", reqForm.PayTime)

	signMap := make(map[string]string)
	signMap["orderid"] = reqForm.OrderId
	signMap["username"] = reqForm.UserName
	signMap["gameid"] = gameIdStr
	signMap["roleid"] = reqForm.RoleId
	signMap["serverid"] = serverIdStr
	signMap["paytype"] = reqForm.PayType
	signMap["amount"] = amountStr
	signMap["paytime"] = payTimeStr
	signMap["attach"] = reqForm.Attache
	signMap["appkey"] = appKey

	keyList := []string{"orderid", "username", "gameid", "roleid", "serverid", "paytype", "amount", "paytime", "attach", "appkey"}
	allStr := ""
	for _, key := range keyList {
		keyValue := fmt.Sprintf("%s=%s&", key, signMap[key])
		allStr += keyValue
	}
	if len(allStr) > 0 {
		allStr = allStr[:len(allStr)-1]
	}
	fmt.Println(allStr)
	hw := md5.Sum([]byte(allStr))
	return fmt.Sprintf("%x", hw)
}
