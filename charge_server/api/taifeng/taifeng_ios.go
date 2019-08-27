package taifeng

import (
	"crypto/md5"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/game/global"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

type taiFengRequest struct {
	UserId int64  `form:"userid" json:"userid"`
	Server int32  `form:"server" json:"server"`
	Money  int32  `form:"money" json:"money"`
	Pay    int64  `form:"pay" json:"pay"`
	Order  string `form:"order" json:"order"`
	Time   int64  `form:"time" json:"time"`
	Sign   string `form:"sign" json:"sign"`
}

func handleTaiFengIOS(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	appId := query.Get("app_id")
	tfTradeNo := query.Get("tf_trade_no")
	chlOrderNum := query.Get("chl_order_num")
	extra := query.Get("extra")
	cpTradeId := query.Get("cp_trade_id")
	roleId := query.Get("role_id")
	serverId := query.Get("server_id")
	moneyType := query.Get("money_type")
	totalFee := query.Get("total_fee")
	payType := query.Get("pay_type")
	payResult := query.Get("pay_result")
	sign := query.Get("sign")
	log.WithFields(
		log.Fields{
			"ip":          req.RemoteAddr,
			"appId":       appId,
			"tfTradeNo":   tfTradeNo,
			"chlOrderNum": chlOrderNum,
			"extra":       extra,
			"cpTradeId":   cpTradeId,
			"roleId":      roleId,
			"serverId":    serverId,
			"moneyType":   moneyType,
			"totalFee":    totalFee,
			"payType":     payType,
			"payResult":   payResult,
			"sign":        sign,
		}).Info("charge:启灵ios充值请求")

	moneyFloat, err := strconv.ParseFloat(totalFee, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"appId":       appId,
				"tfTradeNo":   tfTradeNo,
				"chlOrderNum": chlOrderNum,
				"extra":       extra,
				"cpTradeId":   cpTradeId,
				"roleId":      roleId,
				"serverId":    serverId,
				"moneyType":   moneyType,
				"totalFee":    totalFee,
				"payType":     payType,
				"payResult":   payResult,
				"sign":        sign,
				"error":       err,
			}).Warn("charge:启灵ios充值请求，解析错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	money := int32(moneyFloat)

	sdkType := logintypes.SDKTypeTaiFeng
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:启灵ios充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	taiFengConfig, ok := sdkConfig.(*sdksdk.TaiFengConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:启灵ios充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeIOS
	appKey := taiFengConfig.GetAppKey(devicePlatformType)

	//TODO 验证签名
	getSign := TaifengSign(appId, tfTradeNo, chlOrderNum, extra, cpTradeId, roleId, serverId, moneyType, totalFee, payType, payResult, appKey)
	if sign != getSign {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"appId":       appId,
				"tfTradeNo":   tfTradeNo,
				"chlOrderNum": chlOrderNum,
				"extra":       extra,
				"cpTradeId":   cpTradeId,
				"roleId":      roleId,
				"serverId":    serverId,
				"moneyType":   moneyType,
				"totalFee":    totalFee,
				"payType":     payType,
				"payResult":   payResult,
				"sign":        sign,
				"getSign":     getSign,
				"error":       err,
			}).Warn("charge:启灵ios充值请求,签名错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	now := global.GetGame().GetTimeService().Now()
	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, repeat, err := chargeService.OrderPay(extra, tfTradeNo, logintypes.SDKTypeTaiFeng, money, roleId, now)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"appId":       appId,
				"tfTradeNo":   tfTradeNo,
				"chlOrderNum": chlOrderNum,
				"extra":       extra,
				"cpTradeId":   cpTradeId,
				"roleId":      roleId,
				"serverId":    serverId,
				"moneyType":   moneyType,
				"totalFee":    totalFee,
				"payType":     payType,
				"payResult":   payResult,
				"sign":        sign,
				"error":       err,
			}).Error("charge:启灵ios请求,错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if obj == nil {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"appId":       appId,
				"tfTradeNo":   tfTradeNo,
				"chlOrderNum": chlOrderNum,
				"extra":       extra,
				"cpTradeId":   cpTradeId,
				"roleId":      roleId,
				"serverId":    serverId,
				"moneyType":   moneyType,
				"totalFee":    totalFee,
				"payType":     payType,
				"payResult":   payResult,
				"sign":        sign,
				"error":       err,
			}).Warn("charge:启灵ios请求,订单不存在")
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
			"appId":       appId,
			"tfTradeNo":   tfTradeNo,
			"chlOrderNum": chlOrderNum,
			"extra":       extra,
			"cpTradeId":   cpTradeId,
			"roleId":      roleId,
			"serverId":    serverId,
			"moneyType":   moneyType,
			"totalFee":    totalFee,
			"payType":     payType,
			"payResult":   payResult,
			"sign":        sign,
		}).Info("charge:启灵ios充值请求")
}

func TaifengSign(appId, tfTradeNo, chlOrderNum, extra, cpTradeId, roleId, serverId, moneyType, totalFee, payType, payResult, payKey string) (sign string) {
	signMap := make(map[string]string)
	signMap["app_id"] = appId
	signMap["tf_trade_no"] = tfTradeNo
	signMap["extra"] = extra
	signMap["cp_trade_id"] = cpTradeId
	signMap["role_id"] = roleId
	signMap["server_id"] = serverId
	signMap["money_type"] = moneyType
	signMap["total_fee"] = totalFee
	signMap["pay_type"] = payType
	signMap["pay_result"] = payResult

	keyList := []string{"app_id", "tf_trade_no", "extra", "cp_trade_id", "role_id", "server_id", "money_type", "total_fee", "pay_type", "pay_result"}
	sort.Sort(sort.StringSlice(keyList))
	allStr := ""
	for _, key := range keyList {
		value := signMap[key]
		keyValue := fmt.Sprintf("%s=%s&", key, value)
		allStr += keyValue
	}
	payKey = fmt.Sprintf("%s", payKey)
	allStr += payKey
	fmt.Println(allStr)
	hw := md5.Sum([]byte(allStr))
	return fmt.Sprintf("%x", hw)
}
