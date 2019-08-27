package login_handler

import (
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
	"net/url"
	"sort"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypeTaiFeng, login.LoginHandlerFunc(handleTaiFeng))
}

func handleTaiFeng(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	if devicePlatformType == types.DevicePlatformTypeIOS {
		flag = false
		return
	}

	csAccountLogin := msg.(*uipb.CSAccountLogin)
	taiFeng := csAccountLogin.GetTaiFeng()
	if taiFeng == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆泰逢,数据是空的")
		return
	}
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆泰逢,sdk配置为空")
		return
	}
	taiFengConfig, ok := sdkConfig.(*sdksdk.TaiFengConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆泰逢,sdk配置强制转换失败")
		return
	}
	appId := taiFengConfig.GetAppId(devicePlatformType)
	appKey := taiFengConfig.GetAppKey(devicePlatformType)
	timeStr := taiFeng.GetTime()
	token := taiFeng.GetToken()

	sign := TaiFengSign(appId, timeStr, token, appKey)
	userId, err = taiFengLogin(appId, timeStr, token, sign)
	if err != nil {
		log.WithFields(
			log.Fields{
				"appId":   appId,
				"appKey":  appKey,
				"timeStr": timeStr,
				"token":   token,
				"sign":    sign,
				"err":     err,
			}).Warn("login:登陆泰逢,登陆失败")
		return
	}

	returnPlatform = int32(platform)
	flag = true
	log.WithFields(
		log.Fields{
			"appId":   appId,
			"appKey":  appKey,
			"timeStr": timeStr,
			"token":   token,
			"userId":  userId,
		}).Info("login:登陆泰逢,登陆成功")
	return
}

const (
	taiFengCheckUrl = "http://sdk-login.cocowan.cn/AuthToken"
)

type TaiFengLoginCheckResult struct {
	Code int32           `json:"code"` //返回状态码
	Msg  json.RawMessage `json:"msg"`  //请求相应详细
}

type TaiFengkResultData struct {
	Uid string `json:"uid`
}

func taiFengLogin(appId, time, token, sign string) (userId string, err error) {

	params := make(map[string][]string)
	params["app_id"] = []string{appId}
	params["token"] = []string{token}
	params["time"] = []string{time}
	params["sign"] = []string{sign}

	data := url.Values(params)
	resp, err := http.PostForm(taiFengCheckUrl, data)
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
		log.WithFields(
			log.Fields{
				"appId": appId,
				"token": token,
				"time":  time,
				"sign":  sign,
				"err":   err,
			}).Error("login:登陆泰逢手游,错误")
		return
	}

	result := &TaiFengLoginCheckResult{}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"appId": appId,
				"token": token,
				"time":  time,
				"sign":  sign,
				"body":  string(respBody),
			}).Warn("login:登陆泰逢手游,登录验证失败")
		return
	}

	statusInt := result.Code
	if statusInt != 0 {
		log.WithFields(
			log.Fields{
				"appId":  appId,
				"token":  token,
				"time":   time,
				"sign":   sign,
				"status": result.Code,
				"msg":    string(result.Msg),
			}).Warn("login:登陆泰逢手游,登录验证失败")
		err = fmt.Errorf("status code %d", statusInt)
		return
	}

	resultData := &TaiFengkResultData{}
	err = json.Unmarshal([]byte(result.Msg), resultData)
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{
			"appId": appId,
			"uId":   resultData.Uid,
			"token": token,
			"sign":  sign,
			"body":  string(respBody),
		}).Info("login:登陆泰逢,登陆成功")
	return resultData.Uid, nil
}

func TaiFengSign(appId string, timeStr string, token string, appKey string) (sign string) {
	signMap := make(map[string]string)
	signMap["app_id"] = appId
	signMap["time"] = timeStr
	signMap["token"] = url.QueryEscape(token)
	keyList := []string{"app_id", "time", "token"}
	sort.Sort(sort.StringSlice(keyList))
	allStr := ""
	for _, key := range keyList {
		keyValue := fmt.Sprintf("%s=%s&", key, signMap[key])
		allStr += keyValue
	}

	appKey = fmt.Sprintf("%s", appKey)
	allStr += appKey
	fmt.Println(allStr)
	hw := md5.Sum([]byte(allStr))
	return fmt.Sprintf("%x", hw)
}
