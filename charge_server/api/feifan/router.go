package feifan

import (
	"crypto/md5"
	"encoding/json"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/pkg/timeutils"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

const (
	feiFanPath = "/feifan"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(feiFanPath).Subrouter()
	sr.Path("/android").Handler(http.HandlerFunc(handleFeiFanAndroid))
	sr.Path("/ios").Handler(http.HandlerFunc(handleFeiFanIOS))
}

type FeiFanForm struct {
	State     int32       `form:"state" json:"state"`
	SdkUserId string      `form:"sdkUserID" json:"sdkUserID"`
	Data      *FeiFanData `form:"data" json:"data"`
}

type FeiFanData struct {
	ProductId string `form:"productID" json:"productID"`
	OrderId   int64  `form:"orderID" json:"orderID"`
	UserId    int64  `form:"userID" json:"userID"`
	ChannelId int32  `form:"channelID" json:"channelID"`
	GameId    int32  `form:"gameId" json:gameId`
	ServerId  string `form:"serverID" json:serverID`
	Money     int32  `form:"money" json:"money"`
	Currency  string `form:"currency" json:"currency"`
	Extension string `form:"extension" json:"extension"`
	SignType  string `form:"signType" json:signType`
	Sign      string `form:"sign" json:sign`
}

func handleFeiFanAndroid(rw http.ResponseWriter, req *http.Request) {

	form := &FeiFanForm{}
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(form)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":  req.RemoteAddr,
				"err": err,
			}).Warn("charge:非凡安卓充值请求,参数错误")
		result := "参数解析错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	state := form.State
	sdkUserIdStr := form.SdkUserId
	data := form.Data
	sdkType := logintypes.SDKTypeFeiFan
	log.WithFields(
		log.Fields{
			"ip":        req.RemoteAddr,
			"state":     state,
			"sdkUserId": sdkUserIdStr,
			"data":      data,
		}).Info("charge:非凡安卓充值请求")

	serverIdInt, err := strconv.ParseInt(data.ServerId, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":        req.RemoteAddr,
				"state":     state,
				"sdkUserId": sdkUserIdStr,
				"data":      data,
				"error":     err,
			}).Warn("charge:非凡安卓充值请求，解析错误")
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
			}).Warn("charge:非凡安卓充值请求，充值失败")
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
			}).Warn("charge:非凡安卓充值请求,sdk配置为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	feiFanConfig, ok := sdkConfig.(*sdksdk.FeiFanConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Warn("charge:非凡安卓充值请求,sdk配置强制转换失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	devicePlatformType := logintypes.DevicePlatformTypeAndroid
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
			}).Warn("charge:非凡安卓充值请求,签名类型错误")

		result := "签名类型错误"
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
		return
	}

	//TODO 验证签名
	sign := data.Sign
	dataMap := convertFeiFanDataToMap(data)
	originalData := GetFeiFanOriginalData(dataMap, secretKey)
	flag, err := checkSign(signType, publicKey, sign, originalData)
	if err != nil || !flag {
		log.WithFields(
			log.Fields{
				"ip":        req.RemoteAddr,
				"state":     state,
				"sdkUserId": sdkUserIdStr,
				"data":      data,
				"err":       err,
			}).Warn("charge:非凡安卓充值请求,签名错误")
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
			}).Error("charge:非凡安卓请求,订单处理错误")
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
			}).Warn("charge:非凡安卓请求,订单不存在")
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
		}).Info("charge:非凡安卓充值请求,成功")
}

func convertFeiFanDataToMap(data *FeiFanData) (dataMap map[string]string) {
	dataMap = make(map[string]string)
	dataMap["productID"] = data.ProductId
	dataMap["orderID"] = fmt.Sprintf("%d", data.OrderId)
	dataMap["userID"] = fmt.Sprintf("%d", data.UserId)
	dataMap["channelID"] = fmt.Sprintf("%d", data.ChannelId)
	dataMap["gameID"] = fmt.Sprintf("%d", data.GameId)
	dataMap["serverID"] = data.ServerId
	dataMap["money"] = fmt.Sprintf("%d", data.Money)
	dataMap["currency"] = data.Currency
	dataMap["extension"] = data.Extension
	dataMap["signType"] = data.SignType
	dataMap["sign"] = data.Sign
	return
}

func GetFeiFanOriginalData(dataStrMap map[string]string, secretKey string) (sign string) {
	dataKeyList := make([]string, 0, 8)
	for key, _ := range dataStrMap {
		if key == "signType" || key == "sign" {
			continue
		}
		dataKeyList = append(dataKeyList, key)
	}
	sort.Sort(sort.StringSlice(dataKeyList))

	allStr := ""
	for _, key := range dataKeyList {
		keyValue := fmt.Sprintf("%s=%s&", key, dataStrMap[key])
		allStr += keyValue
	}
	allStr += secretKey
	return allStr
}

func checkSign(signType, publicKey, sign, originalData string) (flag bool, err error) {
	switch signType {
	case "rsa":
		{
			err = coreutils.CheckRsaSign(publicKey, sign, originalData)
			flag = true
			return
		}
	case "md5":
		{
			hw := md5.Sum([]byte(originalData))
			curSign := fmt.Sprintf("%x", hw)
			if sign == curSign {
				flag = true
			}
			return
		}
	}
	return
}
