package xiaoyao

import (
	"encoding/json"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

func handleXiaoYaoIOS(rw http.ResponseWriter, req *http.Request) {
	log.WithFields(
		log.Fields{
			"ip": req.RemoteAddr,
		}).Info("charge: 逍遥ios充值请求, 回调")
	errData := &XiaoYaoRespData{
		Code: "2",
	}

	form := &XiaoYaoReqData{}
	err := req.ParseForm()
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge: 逍遥充值请求，参数解析错误")
		data, _ := json.Marshal(errData)
		rw.Write(data)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = httputils.BindForm(form, req.Form, nil)

	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge: 逍遥充值请求，参数解析错误")
		data, _ := json.Marshal(errData)
		rw.Write(data)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	sid := form.Sid
	uid := form.Uid
	oid := form.Oid
	money := form.Money
	gold := form.Gold
	time := form.Time
	gameSN := form.GameSN
	gameAttach := form.GameAttach

	sdkType := logintypes.SDKTypeXiaoYao
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge: 逍遥ios充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(errData)
		rw.Write(data)
		return
	}
	xiaoYaoConfig, ok := sdkConfig.(*sdksdk.XiaoYaoConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge: 逍遥ios充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(errData)
		rw.Write(data)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeIOS
	apiKey := xiaoYaoConfig.GetApiKey(devicePlatformType)
	secretKey := xiaoYaoConfig.GetSecretKey(devicePlatformType)
	if apiKey != form.ApiKey {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge: 逍遥ios充值请求,传入apiKey与本地不同")
		rw.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(errData)
		rw.Write(data)
		return
	}

	sign := getSign(form, secretKey)
	if sign != form.Sign {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"sid":        sid,
				"uid":        uid,
				"oid":        oid,
				"money":      money,
				"gold":       gold,
				"time":       time,
				"gameSN":     gameSN,
				"gameAttach": gameAttach,
				"api_key":    apiKey,
				"sign":       sign,
				"err":        err,
			}).Warn("charge: 逍遥ios充值请求,签名错误")
		rw.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(errData)
		rw.Write(data)
		return
	}
	uidStr := fmt.Sprintf("%d", uid)
	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(gameSN, oid, sdkType, money, uidStr, time)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"sid":        sid,
				"uid":        uid,
				"oid":        oid,
				"money":      money,
				"gold":       gold,
				"time":       time,
				"gameSN":     gameSN,
				"gameAttach": gameAttach,
				"api_key":    apiKey,
				"sign":       sign,
				"err":        err,
			}).Warn("charge: 逍遥ios充值请求,订单处理错误")
		rw.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(errData)
		rw.Write(data)
		return
	}
	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":         req.RemoteAddr,
				"sid":        sid,
				"uid":        uid,
				"oid":        oid,
				"money":      money,
				"gold":       gold,
				"time":       time,
				"gameSN":     gameSN,
				"gameAttach": gameAttach,
				"api_key":    apiKey,
				"sign":       sign,
				"err":        err,
			}).Warn("charge: 逍遥ios充值请求,订单不存在")
		rw.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(errData)
		rw.Write(data)
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

	successData := &XiaoYaoRespData{
		Code: "1",
	}
	data, _ := json.Marshal(successData)
	rw.Write(data)
	rw.WriteHeader(http.StatusOK)

	log.WithFields(
		log.Fields{
			"ip":         req.RemoteAddr,
			"sid":        sid,
			"uid":        uid,
			"oid":        oid,
			"money":      money,
			"gold":       gold,
			"time":       time,
			"gameSN":     gameSN,
			"gameAttach": gameAttach,
			"api_key":    apiKey,
			"sign":       sign,
		}).Info("charge: 逍遥ios充值请求, 成功")
}
