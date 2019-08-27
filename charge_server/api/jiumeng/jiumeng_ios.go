package jiumeng

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
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type JiuMengRequest struct {
	OrderId  string `form:"orderid" json:"orderid"`
	UserId   int64  `form:"userid" json:"userid"`
	GameId   string `form:"gameid" json:"gameid"`
	RoleId   string `form:"roleid" json:"roleid"`
	ServerId string `form:"serverid" json:"serverid"`

	Money   string `form:"money" json:"money"`
	PayTime string `form:"paytime" json:"paytime"`
	Attach  string `form:"attach" json:"attach"`
	Sign    string `form:"sign" json:"sign"`
}

func handleJiuMengIOS(rw http.ResponseWriter, req *http.Request) {
	reqForm := &JiuMengRequest{}
	err := httputils.Bind(req, reqForm)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:九梦ios充值请求，参数解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.WithFields(
		log.Fields{
			"ip":       req.RemoteAddr,
			"orderId":  reqForm.OrderId,
			"userName": reqForm.UserId,
			"gameId":   reqForm.GameId,
			"roleId":   reqForm.RoleId,
			"serverId": reqForm.ServerId,
			"amount":   reqForm.Money,
			"money":    reqForm.PayTime,
			"attach":   reqForm.Attach,
			"sign":     reqForm.Sign,
		}).Info("charge:九梦ios充值请求")

	sdkType := logintypes.SDKTypeJiuMeng
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:九梦ios充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	qiAConfig, ok := sdkConfig.(*sdksdk.JiuMengConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip": req.RemoteAddr,
			}).Warn("charge:九梦ios充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeIOS
	appKey := qiAConfig.GetAppKey(devicePlatformType)

	//TODO 验证签名
	getSign := GetJiuMengSign(reqForm, appKey)
	if reqForm.Sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserId,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"amount":   reqForm.Money,
				"money":    reqForm.PayTime,
				"attach":   reqForm.Attach,
				"sign":     reqForm.Sign,
				"appKey":   appKey,
			}).Warn("charge:九梦ios充值请求,签名错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	money, err := strconv.ParseInt(reqForm.Money, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserId,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"amount":   reqForm.Money,
				"money":    reqForm.PayTime,
				"attach":   reqForm.Attach,
				"sign":     reqForm.Sign,
				"appKey":   appKey,
			}).Warn("charge:九梦ios充值请求,解析错误")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := timeutils.TimeToMillisecond(time.Now())
	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(reqForm.Attach, reqForm.OrderId, logintypes.SDKTypeJiuMeng, int32(money)/100, fmt.Sprintf("%d", reqForm.UserId), now)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserId,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"amount":   reqForm.Money,
				"money":    reqForm.PayTime,
				"attach":   reqForm.Attach,
				"sign":     reqForm.Sign,
				"error":    err,
			}).Error("charge:九梦ios请求,错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":       req.RemoteAddr,
				"orderId":  reqForm.OrderId,
				"userName": reqForm.UserId,
				"gameId":   reqForm.GameId,
				"roleId":   reqForm.RoleId,
				"serverId": reqForm.ServerId,
				"amount":   reqForm.Money,
				"money":    reqForm.PayTime,
				"attach":   reqForm.Attach,
				"sign":     reqForm.Sign,
			}).Warn("charge:九梦ios请求,订单不存在")
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
			"userName": reqForm.UserId,
			"gameId":   reqForm.GameId,
			"roleId":   reqForm.RoleId,
			"serverId": reqForm.ServerId,
			"amount":   reqForm.Money,
			"money":    reqForm.PayTime,
			"attach":   reqForm.Attach,
			"sign":     reqForm.Sign,
		}).Info("charge:九梦ios充值请求")
}

func GetJiuMengSign(reqForm *JiuMengRequest, appKey string) (sign string) {
	//gameIdStr := fmt.Sprintf("%d", reqForm.GameId)
	//serverIdStr := fmt.Sprintf("%d", reqForm.ServerId)
	//amountStr := fmt.Sprintf("%d", reqForm.Money)
	//payTimeStr := fmt.Sprintf("%d", reqForm.PayTime)

	signMap := make(map[string]string)
	signMap["orderid"] = reqForm.OrderId
	signMap["userid"] = fmt.Sprintf("%d", reqForm.UserId)
	signMap["gameid"] = reqForm.GameId
	signMap["roleid"] = reqForm.RoleId
	signMap["serverid"] = reqForm.ServerId
	signMap["money"] = reqForm.Money
	signMap["paytime"] = reqForm.PayTime
	signMap["attach"] = reqForm.Attach
	signMap["appkey"] = appKey

	keyList := []string{"orderid", "userid", "gameid", "roleid", "serverid", "money", "paytime", "attach", "appkey"}
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
