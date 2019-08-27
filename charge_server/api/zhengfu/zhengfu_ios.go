package zhengfu

import (
	"encoding/json"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/pkg/timeutils"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

func handleZhengFuIOS(rw http.ResponseWriter, req *http.Request) {

	form := &ZhengFuForm{}
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(form)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:征服ios充值请求,参数错误")
		result := "参数解析错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	state := form.State
	sdkUserIdStr := form.SdkUserId
	data := form.Data
	sdkType := logintypes.SDKTypeZhengFu
	log.WithFields(
		log.Fields{
			"ip":        req.RemoteAddr,
			"state":     state,
			"sdkUserId": sdkUserIdStr,
			"data":      data,
		}).Info("charge:征服ios充值请求")

	serverIdInt, err := strconv.ParseInt(data.ServerId, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":        req.RemoteAddr,
				"state":     state,
				"sdkUserId": sdkUserIdStr,
				"data":      data,
				"error":     err,
			}).Warn("charge:征服ios充值请求，解析错误")
		result := "参数解析错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}
	serverId := int32(serverIdInt)

	//充值失败
	if state != 1 {
		ctx := req.Context()
		chargeService := charge.ChargeServiceInContext(ctx)
		chargeService.OrderFail(data.Extension, sdkType)
		log.WithFields(
			log.Fields{
				"ip":        req.RemoteAddr,
				"state":     state,
				"sdkUserId": sdkUserIdStr,
				"data":      data,
				"error":     err,
			}).Warn("charge:征服ios充值请求，充值失败")
		result := "订单已经取消"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	moneyFen := data.Money
	//分转元
	money := int32(moneyFen / 100)
	receiveTime := timeutils.TimeToMillisecond(time.Now())

	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:征服ios充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	feiFanConfig, ok := sdkConfig.(*sdksdk.ZhengFuConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:征服ios充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	devicePlatformType := logintypes.DevicePlatformTypeIOS
	secretKey := feiFanConfig.GetSecretKey(devicePlatformType)
	publicKey := feiFanConfig.GetPublicKey(devicePlatformType)

	signType := strings.ToLower(data.SignType)
	if signType != "rsa" && signType != "md5" {
		log.WithFields(
			log.Fields{
				"ip":        req.RemoteAddr,
				"state":     state,
				"sdkUserId": sdkUserIdStr,
				"data":      data,
			}).Warn("charge:征服ios充值请求,签名类型错误")

		result := "签名类型错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	//TODO 验证签名
	sign := data.Sign
	dataMap := convertZhengFuDataToMap(data)
	originalData := GetZhengFuOriginalData(dataMap, secretKey)
	flag, err := checkSign(signType, publicKey, sign, originalData)
	if err != nil || !flag {
		log.WithFields(
			log.Fields{
				"ip":        req.RemoteAddr,
				"state":     state,
				"sdkUserId": sdkUserIdStr,
				"data":      data,
				"err":       err,
			}).Warn("charge:征服ios充值请求,签名错误")
		result := "签名错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(data.Extension, fmt.Sprintf("%d", data.OrderId), sdkType, money, fmt.Sprintf("%d", data.UserId), receiveTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":        req.RemoteAddr,
				"state":     state,
				"sdkUserId": sdkUserIdStr,
				"data":      data,
				"error":     err,
			}).Error("charge:征服ios请求,订单处理错误")
		result := "订单处理错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":        req.RemoteAddr,
				"state":     state,
				"sdkUserId": sdkUserIdStr,
				"data":      data,
				"error":     err,
			}).Warn("charge:征服ios请求,订单不存在")
		result := "订单不存在"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}
	if !repeat {
		//放入回调队列中
		remoteService := remote.RemoteServiceInContext(ctx)
		flag = remoteService.Charge(obj)
		if !flag {
			panic(fmt.Errorf("charge:添加到回调队列应该成功"))
		}
	}

	result := "SUCCESS"
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(result))
	log.WithFields(
		log.Fields{
			"orderId":     data.Extension,
			"sdkUserId":   sdkUserIdStr,
			"server":      serverId,
			"money":       money,
			"pay":         data.OrderId,
			"receiveTime": receiveTime,
			"sign":        sign,
		}).Info("charge:征服ios充值请求,成功")
}
