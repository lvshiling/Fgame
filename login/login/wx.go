package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type WxLoginResult struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionId      string `json:"unionId"`
}

//TODO 重试
func WxLogin(appId string, secret string, code string) (result *WxLoginResult, err error) {

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", appId, secret, code)

	resp, err := http.Get(url)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("微信不可访问")
		return
	}

	defer resp.Body.Close()
	result = &WxLoginResult{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return
	}
	return
}

type WxRefreshResult struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
}

func WxRefreshToken(appId string, refreshToken string) (result *WxRefreshResult, err error) {

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s", appId, refreshToken)
	resp, err := http.Get(url)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("微信不可访问")
		return
	}

	defer resp.Body.Close()
	result = &WxRefreshResult{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return
	}
	return
}

type WxUserInfoResult struct {
	ErrCode    int      `json:"errcode"`
	ErrMsg     string   `json:"errmsg"`
	NickName   string   `json:"nickname"`
	OpenId     string   `json:"openid"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headImgUrl"`
	Privilege  []string `json:"privilege"`
	UnionId    string   `json:"unionId"`
}

func WxGetUserInfo(openId string, accessToken string) (result *WxUserInfoResult, err error) {

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", accessToken, openId)
	resp, err := http.Get(url)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("微信不可访问")
		return
	}

	defer resp.Body.Close()
	result = &WxUserInfoResult{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return
	}
	return
}

type WxAuthResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func WxAuth(openId string, accessToken string) (result *WxAuthResult, err error) {

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s", accessToken, openId)
	resp, err := http.Get(url)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("微信不可访问")
		return
	}

	defer resp.Body.Close()
	result = &WxAuthResult{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return
	}
	return
}
