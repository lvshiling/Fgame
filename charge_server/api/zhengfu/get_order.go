package zhengfu

import (
	"encoding/json"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
)

const (
	getOrderUrl = "https://issue.23youxi.com/pay/getOrderID"
)

type sdkOrderData struct {
	SdkOrderId string `json:"orderID"`
	Extension  string `json:"extension"`
}

type sdkOrderResult struct {
	State int32         `json:"state"` //返回状态码
	Data  *sdkOrderData `json:"data"`  //返回信息
}

func init() {
	charge.RegisterSDKOrderHandler(logintypes.SDKTypeZhengFu, charge.GetSDKOrderHandlerFunc(GetSDKOrder))
}

//获取第三方下单
func GetSDKOrder(devicePlatformType logintypes.DevicePlatformType, platformUserId string, productId int32, money int32, roleId int64, roleName string, serverId int32, orderId string) (notifyUrl string, sdkOrderId string, extension string, flag bool) {

	sdkConfig := sdk.GetSdkService().GetSdkConfig(logintypes.SDKTypeZhengFu)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"platformUserId": platformUserId,
				"prodoctId":      productId,
				"money":          money,
				"roleId":         roleId,
				"roleName":       roleName,
				"serverId":       serverId,
				"orderId":        orderId,
			}).Warn("charge:征服下单,sdk配置为空")
		return
	}
	zhengFuConfig, ok := sdkConfig.(*sdksdk.ZhengFuConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"platformUserId": platformUserId,
				"prodoctId":      productId,
				"money":          money,
				"roleId":         roleId,
				"roleName":       roleName,
				"serverId":       serverId,
				"orderId":        orderId,
			}).Warn("charge:征服下单,sdk配置强制转换失败")
		return
	}

	notifyUrl = zhengFuConfig.GetNotifyUrl(devicePlatformType)
	appSecret := zhengFuConfig.GetSecretKey(devicePlatformType)
	privateKey := zhengFuConfig.GetPrivateKey(devicePlatformType)

	tSdkOrerId, textension, err := GetSdkOrderId(appSecret, privateKey, notifyUrl, orderId, platformUserId, productId, money, roleId, roleName, serverId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"platformUserId": platformUserId,
				"prodoctId":      productId,
				"money":          money,
				"roleId":         roleId,
				"roleName":       roleName,
				"serverId":       serverId,
				"orderId":        orderId,
				"error":          err,
			}).Warn("charge:征服安卓下单,下单错误")
		return
	}
	if tSdkOrerId == "" {
		return
	}
	sdkOrderId = tSdkOrerId
	extension = textension
	flag = true
	return
}

func GetSdkOrderId(appSecret string, privateKey string, notifyUrl string, orderId string, platformUserId string, produtId int32, money int32, roleId int64, roleName string, serverId int32) (sdkOrderId string, extension string, err error) {

	userIdStr := platformUserId
	productIdStr := fmt.Sprintf("%d", produtId)
	moneyStr := fmt.Sprintf("%d", money*100)
	roleIdStr := fmt.Sprintf("%d", roleId)
	serverIdStr := fmt.Sprintf("%d", serverId)

	serverName := serverIdStr
	signType := "rsa"

	signMap := make(map[string]string)
	signMap["userID"] = userIdStr
	signMap["productID"] = productIdStr
	signMap["productName"] = ""
	signMap["productDesc"] = ""
	signMap["money"] = moneyStr
	signMap["roleID"] = roleIdStr
	signMap["roleName"] = roleName
	signMap["serverID"] = serverIdStr
	signMap["serverName"] = serverName
	signMap["extension"] = orderId
	signMap["notifyUrl"] = notifyUrl
	keyList := []string{"userID", "productID", "productName", "productDesc", "money", "roleID", "roleName", "serverID", "serverName", "extension", "notifyUrl"}
	allStr := ""
	for _, key := range keyList {
		keyValue := fmt.Sprintf("%s=%s&", key, signMap[key])
		allStr += keyValue
	}
	if len(allStr) > 0 {
		allStr = allStr[:len(allStr)-1]
	}
	allStr += appSecret
	allStr = url.QueryEscape(allStr)
	sign, err := coreutils.RsaSign(allStr, privateKey)
	if err != nil {
		return
	}

	params := make(map[string][]string)
	params["userID"] = []string{userIdStr}
	params["productID"] = []string{productIdStr}
	params["productName"] = []string{""}
	params["productDesc"] = []string{""}
	params["money"] = []string{moneyStr}
	params["roleID"] = []string{roleIdStr}
	params["roleName"] = []string{roleName}
	params["serverID"] = []string{serverIdStr}
	params["serverName"] = []string{serverName}
	params["extension"] = []string{orderId}
	params["notifyUrl"] = []string{notifyUrl}
	params["signType"] = []string{signType}
	params["sign"] = []string{sign}

	data := url.Values(params)
	resp, err := http.PostForm(getOrderUrl, data)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status code %d", resp.StatusCode)
		return
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	sdkResult := &sdkOrderResult{}
	err = json.Unmarshal(respBody, sdkResult)
	if err != nil {
		return
	}
	statusInt := sdkResult.State
	if statusInt != 1 {
		log.WithFields(
			log.Fields{
				"status": sdkResult.State,
				"data":   sdkResult.Data,
			}).Warn("charge:获取订单号征服手游,获取订单号验证失败")
		return
	}

	if sdkResult.Data == nil {
		log.WithFields(
			log.Fields{
				"status": sdkResult.State,
				"data":   sdkResult.Data,
			}).Warnf("charge:获取订单号征服手游,数据是空")
		return
	}

	sdkOrderId = sdkResult.Data.SdkOrderId
	extension = sdkResult.Data.Extension
	return
}
