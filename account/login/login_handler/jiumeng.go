package login_handler

import (
	"fgame/fgame/account/login/login"
	"fgame/fgame/account/login/types"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypeJiuMeng, login.LoginHandlerFunc(handleJiuMeng))
}

func handleJiuMeng(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)

	jiuMeng := csAccountLogin.GetJiuMeng()
	if jiuMeng == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆九梦天书,数据是空的")
		return
	}
	devicePlatform := types.DevicePlatformType(csAccountLogin.GetDevicePlatform())
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆九梦天书,sdk配置为空")
		return
	}
	jiuMengConfig, ok := sdkConfig.(*sdksdk.JiuMengConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆九梦天书,sdk配置强制转换失败")
		return
	}
	returnPlatform = int32(platform)

	timeStr := jiuMeng.GetLogintime()
	userId = jiuMeng.GetUserid()
	sign := jiuMeng.GetSign()
	appId := jiuMengConfig.GetAppId(devicePlatform)
	//验证签名正确性
	flag, err = loginVerify(appId, userId, timeStr, sign)
	if err != nil {
		log.WithFields(
			log.Fields{
				"userId":  userId,
				"timeStr": timeStr,
				"sign":    sign,
				"err":     err,
			}).Warn("login:登陆九梦天书,签名错误")
		return
	}

	flag = true
	log.WithFields(
		log.Fields{
			"timeStr": timeStr,
			"userId":  userId,
			"sign":    sign,
		}).Info("login:登陆九梦天书,登陆成功")
	return
}

const (
	jiuMengURL = "https://sdk.17wjz.com/Api/LoginSign"
)

type YouMengCheckData struct {
	AppId     string `json:"appid"`
	UserId    string `json:"userid"`
	LoginTime int64  `json:"logintime"`
	Sign      string `json:"sign"`
}

func loginVerify(appId string, userId string, loginTime string, sign string) (flag bool, err error) {
	url := fmt.Sprintf("%s?appid=%s&userid=%s&logintime=%s&sign=%s", jiuMengURL, appId, userId, loginTime, sign)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status code: %d", resp.StatusCode)
		return
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if string(respBody) == "success" {
		return true, nil
	}
	return false, fmt.Errorf(string(respBody))
}
