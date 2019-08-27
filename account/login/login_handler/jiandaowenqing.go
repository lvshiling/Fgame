package login_handler

import (
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
	login.RegisterLogin(types.SDKTypeJianDao, login.LoginHandlerFunc(handleJianDao))
}

func handleJianDao(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	jianDao := csAccountLogin.GetJianDao()
	if jianDao == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆剑道问情手游,数据是空的")
		return
	}

	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆剑道问情手游,sdk配置为空")
		return
	}

	jianDaoConfig, ok := sdkConfig.(*sdksdk.JianDaoConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆剑道问情手游,sdk配置强制转换失败")
		return
	}

	appId := jianDaoConfig.GetAppId(devicePlatformType)
	userId = jianDao.GetUserid()
	userToken := jianDao.GetToken()
	userName := jianDao.GetUserName()
	if len(userToken) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:登陆剑道问情手游,userToken为空")
		return
	}

	// 登录认证
	flag, err = jianDaoLogin(appId, userId, userToken, userName)
	if err != nil {
		return
	}
	if !flag {
		return
	}
	returnPlatform = int32(platform)
	log.WithFields(
		log.Fields{
			"appId":     appId,
			"userId":    userId,
			"userName":  userName,
			"userToken": userToken,
		}).Info("login:登陆飞扬梦缘问仙手游,登陆成功")
	return
}

const (
	jianDaoPath = "http://i.5218yx.com/cpVerify.php"
)

type jianDaoLoginCheckResult struct {
	Id     string `json:"id"`
	Status int32  `json:"status"` //返回状态码
	Data   string `json:"msg"`    //返回信息
}

func jianDaoLogin(appId, userId, userToken, userName string) (flag bool, err error) {
	params := make(map[string][]string)
	params["id"] = []string{userId}
	params["appid"] = []string{appId}
	params["username"] = []string{userName}
	params["token"] = []string{userToken}

	data := url.Values(params)
	// TODO 请求超时
	resp, err := http.PostForm(jianDaoPath, data)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆剑道问情手游,登录验证请求失败")
		return false, nil
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
				"err": err,
			}).Warnf("login:登陆剑道问情手游,回包数据读取错误")
		return false, nil
	}

	result := &jianDaoLoginCheckResult{}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆剑道问情手游,回包数据解析错误")
		return false, nil
	}

	if result.Status != 1 {
		log.WithFields(
			log.Fields{
				"status": result.Status,
				"msg":    result.Data,
			}).Warnf("login:登陆剑道问情手游,登录验证失败")
		return
	}

	flag = true
	return
}
