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

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypeXiaoYao, login.LoginHandlerFunc(handleXiaoYao))
}

func handleXiaoYao(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)

	xiaoYao := csAccountLogin.GetXiaoYao()
	if xiaoYao == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆逍遥,数据是空的")
		return
	}

	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆逍遥,sdk配置为空")
		return
	}

	xiaoYaoConfig, ok := sdkConfig.(*sdksdk.XiaoYaoConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆逍遥,sdk配置强制转换失败")
		return
	}
	apiKey := xiaoYaoConfig.GetApiKey(devicePlatformType)
	secretKey := xiaoYaoConfig.GetSecretKey(devicePlatformType)
	scode := xiaoYao.GetScode()

	//验证签名正确性
	uid, err := XiaoYaoVerify(apiKey, secretKey, scode)
	if err != nil {
		return
	}
	userId = uid
	returnPlatform = int32(platform)
	flag = true
	log.WithFields(
		log.Fields{
			"apiKey": apiKey,
			"scode":  scode,
		}).Info("login:登陆逍遥,登陆成功")
	return
}

const (
	xiaoYaoUrl = "http://api.m.okwan.com/login/in"
)

type XiaoYaoData struct {
	Scode  string `json:"scode"`
	ApiKey string `json:"apiKey"`
	Sign   string `json:"sign"`
}

type XiaoYaoResult struct {
	Code   int32  `json:"code"`
	Mes    string `json:"mes"`
	Uid    int64  `json:"uid"`
	ApiKey string `json:"apiKey"`
	Sign   string `json:"sign"`
}

func XiaoYaoVerify(apiKey string, secretKey string, scode string) (uid string, err error) {
	ReqSign := getXiaoYaoRequestSign(scode, apiKey, secretKey)
	params := make(map[string][]string)
	params["scode"] = []string{scode}
	params["api_key"] = []string{apiKey}
	params["sign"] = []string{ReqSign}
	data := url.Values(params)

	resp, err := http.PostForm(xiaoYaoUrl, data)
	if err != nil {
		log.WithFields(
			log.Fields{
				"scode":  scode,
				"apiKey": apiKey,
				"sign":   ReqSign,
				"err":    err,
			}).Error("login:登陆逍遥手游,请求失败")
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
				"scode":  scode,
				"apiKey": apiKey,
				"err":    err,
			}).Error("login:登陆逍遥手游,错误")
		return
	}

	result := &XiaoYaoResult{}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		return
	}

	statusInt := result.Code
	if statusInt != 1 {
		log.WithFields(
			log.Fields{
				"code":   result.Code,
				"mes":    result.Mes,
				"userId": result.Uid,
				"apiKey": result.ApiKey,
				"sign":   result.Sign,
			}).Warn("login:登陆逍遥手游,登录验证失败")
		err = fmt.Errorf("result code %d", result.Code)
		return
	}

	log.WithFields(
		log.Fields{
			"body": string(respBody),
		}).Info("login:登陆逍遥手游,请求信息")

	// if result.ApiKey != apiKey {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"code":   result.Code,
	// 			"mes":    result.Mes,
	// 			"userId": result.Uid,
	// 			"apiKey": result.ApiKey,
	// 			"sign":   result.Sign,
	// 		}).Warn("login:登陆逍遥手游,传入apiKey与本地不同")
	// 	return
	// }
	uid = fmt.Sprintf("%d", result.Uid)
	respSign := getXiaoYaoResponseSign(result.Code, result.Mes, uid, apiKey, secretKey)
	if respSign != result.Sign {
		log.WithFields(
			log.Fields{
				"code":     result.Code,
				"mes":      result.Mes,
				"userId":   result.Uid,
				"apiKey":   result.ApiKey,
				"sign":     result.Sign,
				"respSign": respSign,
			}).Warn("login:登陆逍遥手游,登录签名验证失败")
		return
	}

	return
}

func getXiaoYaoRequestSign(scode string, apiKey string, secretKey string) (sign string) {
	signStr := fmt.Sprintf("scode=%s&api_key=%s", scode, apiKey)
	signByte := md5.Sum([]byte(signStr))
	signStr = fmt.Sprintf("%x%s", signByte, secretKey)
	signByte = md5.Sum([]byte(signStr))
	sign = fmt.Sprintf("%x", signByte)
	return
}

func getXiaoYaoResponseSign(code int32, mes string, uid string, apiKey string, secretKey string) (sign string) {
	signStr := fmt.Sprintf("code=%d&mes=%s&uid=%s&api_key=%s", code, mes, uid, apiKey)
	signByte := md5.Sum([]byte(signStr))
	signStr = fmt.Sprintf("%x%s", signByte, secretKey)
	signByte = md5.Sum([]byte(signStr))
	sign = fmt.Sprintf("%x", signByte)
	return
}
