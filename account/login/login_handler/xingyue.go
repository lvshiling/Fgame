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
	"strings"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypeXingYue, login.LoginHandlerFunc(handleXingYue))
}

func handleXingYue(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	xingYue := csAccountLogin.GetXingYue()
	if xingYue == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆星月手游,数据是空的")
		return
	}
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆星月手游,sdk配置为空")
		return
	}
	xingYueConfig, ok := sdkConfig.(*sdksdk.XingYueConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆星月手游,sdk配置强制转换失败")
		return
	}

	cpId := xingYueConfig.GetCpId(devicePlatformType)
	gameId := xingYueConfig.GetGameId(devicePlatformType)
	paramKey := xingYueConfig.GetParamKey(devicePlatformType)
	userId = xingYue.GetSid()
	if len(userId) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:登陆星月手游,userId为空")
		return
	}

	//登录认证
	flag, err = xingYueLogin(cpId, gameId, userId, paramKey)
	if err != nil {
		log.WithFields(
			log.Fields{
				"cpId":     cpId,
				"gameId":   gameId,
				"userId":   userId,
				"paramKey": paramKey,
				"error":    err,
			}).Warn("login:登陆星月手游,失败")
		return
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"cpId":     cpId,
				"gameId":   gameId,
				"userId":   userId,
				"paramKey": paramKey,
			}).Info("login:登陆星月手游,失败")
		return
	}
	returnPlatform = int32(platform)
	flag = true
	log.WithFields(
		log.Fields{
			"cpId":     cpId,
			"gameId":   gameId,
			"userId":   userId,
			"paramKey": paramKey,
		}).Info("login:登陆星月手游,登陆成功")
	return
}

const (
	xingYueCheckUrl = "https://xingws.xingyue189.com/GetPlayerInfo.ashx"
)

type XingYueCheckData struct {
	UserId        string `json:"userId`
	UserName      string `json:"userName"`
	Sid           string `json:"sid"`
	CpId          string `json:"cpId"`
	GameId        string `json:"gameId"`
	LastVisitTime string `json:"lastVisitTime"`
	Ip            string `json:"ip"`
	FromId        string `json:"fromId"`
	CreatedTime   string `json:"createdt"`
}

type XingYueLoginCheckResult struct {
	Code int32             `json:"code"` //返回状态码
	Msg  string            `json:"msg"`  //请求相应详细
	Data *XingYueCheckData `json:"data"` //返回信息
}

func xingYueLogin(cpId, gameId, userId, paramKey string) (flag bool, err error) {
	sign := getXingYueSign(cpId, gameId, userId, paramKey)

	params := make(map[string][]string)
	params["cpId"] = []string{cpId}
	params["PlayerId"] = []string{userId}
	params["gameId"] = []string{gameId}
	params["paramSign"] = []string{sign}
	data := url.Values(params)
	resp, err := http.PostForm(xingYueCheckUrl, data)
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
				"cpId":     cpId,
				"gameId":   gameId,
				"userId":   userId,
				"paramKey": paramKey,
				"err":      err,
			}).Error("login:登陆星月手游,错误")
		return
	}

	result := &XingYueLoginCheckResult{}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		return
	}

	statusInt := result.Code
	if statusInt != 1 {
		log.WithFields(
			log.Fields{
				"cpId":     cpId,
				"gameId":   gameId,
				"userId":   userId,
				"paramKey": paramKey,
				"status":   result.Code,
				"msg":      result.Msg,
				"data":     result.Data,
			}).Warn("login:登陆星月手游,登录验证失败")
		return
	}

	flag = true
	return
}

func getXingYueSign(cpId, gameId, userId, paramKey string) (sign string) {
	allStr := ""
	allStr += cpId
	allStr += userId
	allStr += gameId
	allStr += paramKey
	hw := md5.Sum([]byte(allStr))
	return strings.ToUpper(fmt.Sprintf("%x", hw))
}
