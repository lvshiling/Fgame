package taifeng

import (
	"fgame/fgame/pkg/timeutils"
	"time"

	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/xozrc/pkg/httputils"
)

const (
	taifengPath = "/taifeng"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(taifengPath).Subrouter()
	sr.Path("/ios").Handler(http.HandlerFunc(handleTaiFengIOS))
	sr.Path("/android").Handler(http.HandlerFunc(handleTaiFengAndroid))
}

type TaiFengForm struct {
	AppId       int32  `form:"app_id" json:"app_id"`
	TfTradeNo   string `form:"tf_trade_no" json:"tf_trade_no"`
	ChlOrderNum string `form:"chl_order_num" json:"chl_order_num"`
	Extra       string `form:"extra" json:"extra"`
	CpTradeId   string `form:"cp_trade_id" json:"cp_trade_id"`
	RoleId      string `form:"role_id" json:"role_id"`
	ServerId    int32  `form:"server_id" json:"server_id"`
	MoneyType   int32  `form:"money_type" json:"money_type"`
	TotalFee    int32  `form:"total_fee" json:"total_fee"`
	PayType     int32  `form:"pay_type" json:"pay_type"`
	PayResult   int32  `form:"pay_result" json:"pay_result"`
	Sign        string `form:"sign" json:"sign"`
}

func handleTaiFengAndroid(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header)

	form := &TaiFengForm{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:泰逢ios充值请求，参数解析错误")
		result := "参数解析错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	appId := fmt.Sprintf("%d", form.AppId)
	tfTradeNo := form.TfTradeNo
	chlOrderNum := form.ChlOrderNum
	extra := form.Extra
	cpTradeId := form.CpTradeId
	roleId := form.RoleId
	serverId := fmt.Sprintf("%d", form.ServerId)
	moneyType := fmt.Sprintf("%d", form.MoneyType)
	totalFee := fmt.Sprintf("%d", form.TotalFee)
	payType := fmt.Sprintf("%d", form.PayType)
	payResult := fmt.Sprintf("%d", form.PayResult)
	sign := form.Sign
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
		}).Info("charge:泰逢安卓充值请求")

	money := form.TotalFee

	sdkType := logintypes.SDKTypeTaiFeng
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:泰逢安卓充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	taiFengConfig, ok := sdkConfig.(*sdksdk.TaiFengConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:泰逢安卓充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	devicePlatformType := logintypes.DevicePlatformTypeAndroid
	payKey := taiFengConfig.GetPayKey(devicePlatformType)

	//TODO 验证签名
	getSign := TaifengSign(appId, tfTradeNo, chlOrderNum, extra, cpTradeId, roleId, serverId, moneyType, totalFee, payType, payResult, payKey)
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
			}).Warn("charge:泰逢安卓充值请求,签名错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	now := timeutils.TimeToMillisecond(time.Now())
	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	moneyYuan := money / 100
	obj, repeat, err := chargeService.OrderPay(cpTradeId, tfTradeNo, logintypes.SDKTypeTaiFeng, moneyYuan, roleId, now)
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
			}).Error("charge:泰逢安卓请求,错误")
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
			}).Warn("charge:泰逢安卓请求,订单不存在")
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
	result := "SUCCESS"
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
		}).Info("charge:泰逢安卓充值请求")
}
