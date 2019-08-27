package login_handler

import (
	"encoding/json"
	"fgame/fgame/account/login/login"
	"fgame/fgame/account/login/types"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypeYiYun, login.LoginHandlerFunc(handleYiYun))
}

func handleYiYun(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	yiYun := csAccountLogin.GetYiYun()
	if yiYun == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆亦云间手游,数据是空的")
		return
	}

	token := yiYun.GetToken()
	if len(token) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:登陆亦云间手游,token")
		return
	}
	ugid := yiYun.GetUgid()
	if len(ugid) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:登陆亦云间手游,ugid")
		return
	}

	// 登录认证
	flag, err = yiYunLogin(token, ugid)
	if err != nil {
		return
	}
	if !flag {
		return
	}
	userId = ugid
	platform := types.SDKType(csAccountLogin.GetPlatform())
	returnPlatform = int32(platform)
	log.WithFields(
		log.Fields{
			"token": token,
			"ugid":  ugid,
		}).Info("login:登陆亦云间手游,登陆成功")
	return
}

const (
	yiYunPath = "http://napi.77l.com/sdkv2/basic/checkugid"
)

type yiYunLoginCheckResult struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"` //返回信息
}

func yiYunLogin(token, ugid string) (flag bool, err error) {
	params := make(map[string][]string)
	params["token"] = []string{token}
	params["ugid"] = []string{ugid}

	data := url.Values(params)
	// TODO 请求超时
	resp, err := http.PostForm(yiYunPath, data)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆亦云间手游,登录验证请求失败")
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
			}).Warnf("login:登陆亦云间手游,回包数据读取错误")
		return false, nil
	}

	result := &yiYunLoginCheckResult{}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆亦云间手游,回包数据解析错误")
		return false, nil
	}

	if result.Code != 0 {
		log.WithFields(
			log.Fields{
				"code": result.Code,
				"msg":  result.Msg,
				"data": result.Data,
			}).Warnf("login:登陆亦云间手游,登录验证失败")
		return
	}

	flag = true
	return
}
