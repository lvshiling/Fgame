package login_handler

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fgame/fgame/account/login/login"
	"fgame/fgame/account/login/types"
	uipb "fgame/fgame/common/codec/pb/ui"

	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypeZuoWan, login.LoginHandlerFunc(handleZuoWan))
}

func handleZuoWan(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	zuoWan := csAccountLogin.GetZuoWan()
	if zuoWan == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆佐玩手游,数据是空的")
		return
	}
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆佐玩手游,sdk配置为空")
		return
	}
	zuoWanConfig, ok := sdkConfig.(*sdksdk.ZuoWanConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆佐玩手游,sdk配置强制转换失败")
		return
	}

	cpId := zuoWanConfig.GetCpId(devicePlatformType)
	gameId := zuoWanConfig.GetGameInstId(devicePlatformType)
	paramKey := zuoWanConfig.GetPrivatekey(devicePlatformType)
	userId = zuoWan.GetSid()
	if len(userId) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:登陆佐玩手游,userId为空")
		return
	}
	token := zuoWan.GetToken()
	if len(token) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:登陆佐玩手游,token为空")
		return
	}

	//登录认证
	flag, err = zuoWanLogin(cpId, userId, gameId, token, paramKey)
	if err != nil {
		log.WithFields(
			log.Fields{
				"cpId":     cpId,
				"gameId":   gameId,
				"userId":   userId,
				"token":    token,
				"paramKey": paramKey,
				"error":    err,
			}).Warn("login:登陆佐玩手游,失败")
		return
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"cpId":     cpId,
				"gameId":   gameId,
				"userId":   userId,
				"token":    token,
				"paramKey": paramKey,
			}).Warn("login:登陆佐玩手游,失败")
		return
	}
	returnPlatform = int32(platform)
	flag = true
	log.WithFields(
		log.Fields{
			"cpId":     cpId,
			"gameId":   gameId,
			"userId":   userId,
			"token":    token,
			"paramKey": paramKey,
		}).Info("login:登陆佐玩手游,登陆成功")
	return
}

const (
	zuoWanCheckUrl = "https://gateway.91zuowan.com/gateway/call/playervalidatepost"
)

type ZuoWanCheckData struct {
	loginName     string `json:"loginName"`
	GameId        string `json:"gameId"`
	Sid           string `json:"sid"`
	RegTime       string `json:"regTime"`
	LastLoginTime string `json:"lastLoginTime"`
	LastLoginIP   string `json:"lastLoginIP"`
}

type ZuoWanLoginCheckResult struct {
	Code int32            `json:"code"` //返回状态码
	Msg  string           `json:"msg"`  //请求相应详细
	Data *ZuoWanCheckData `json:"data"` //返回信息
}

func zuoWanLogin(cpId, sid, gameId, loginToken, paramKey string) (flag bool, err error) {
	sign := GetZuoWanSign(cpId, gameId, sid, loginToken, paramKey)

	params := make(map[string]string)
	params["cpid"] = cpId
	params["sid"] = sid
	params["gameid"] = gameId
	params["logintoken"] = loginToken
	params["paramsign"] = sign
	bodyBytes, err := json.Marshal(params)
	if err != nil {
		return
	}
	request, err := http.NewRequest("POST", zuoWanCheckUrl, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status code %d", resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status code %d", resp.StatusCode)
		return
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(
			log.Fields{
				"cpId":       cpId,
				"sid":        sid,
				"gameId":     gameId,
				"loginToken": loginToken,
				"err":        err,
			}).Error("login:登陆佐玩手游,错误")
		return
	}

	result := &ZuoWanLoginCheckResult{}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		return
	}

	statusInt := result.Code
	if statusInt != 0 {
		log.WithFields(
			log.Fields{
				"cpId":       cpId,
				"sid":        sid,
				"gameId":     gameId,
				"loginToken": loginToken,
				"status":     result.Code,
				"msg":        result.Msg,
				"data":       result.Data,
				"sign":       sign,
			}).Warn("login:登陆佐玩手游,登录验证失败")
		return
	}

	flag = true
	return
}

func GetZuoWanSign(cpId, gameId, userId, token, paramKey string) (sign string) {

	signMap := make(map[string]string)
	signMap["cpid"] = cpId
	signMap["gameid"] = gameId
	signMap["logintoken"] = token
	signMap["sid"] = userId
	keyList := []string{"cpid", "gameid", "logintoken", "sid"}
	allStr := ""
	for _, key := range keyList {
		keyValue := fmt.Sprintf("%s=%s&", key, signMap[key])
		allStr += keyValue
	}
	if len(allStr) > 0 {
		allStr = allStr[:len(allStr)-1]
	}
	allStr += paramKey
	fmt.Println(allStr)
	hw := md5.Sum([]byte(allStr))
	return strings.ToUpper(fmt.Sprintf("%x", hw))
}
