package logic

import (
	"crypto/md5"
	"encoding/json"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/chatproxy/sdk"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	log "github.com/Sirupsen/logrus"
)

func YouQuProxy(sdkType logintypes.SDKType, deviceType logintypes.DevicePlatformType, serverId int32, sdkUserId string, userId int64, playerName string, targetPlayerId int64, targetPlayerName string, chatType string, chatTime int64, body string) {
	sdkConfig := sdk.GetSdkService().GetSdkConfig(sdkType)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"sdkType": sdkType.String(),
			}).Error("配置不存在")
		return
	}
	jianKongUrl := sdkConfig.GetUrl(deviceType)
	gameId := sdkConfig.GetGame(deviceType)
	key := sdkConfig.GetKey(deviceType)
	genre := int32(1)
	queryMap := make(map[string]string)
	queryMap["genre"] = fmt.Sprintf("%d", genre)
	queryMap["game"] = fmt.Sprintf("%d", gameId)
	queryMap["uid"] = sdkUserId
	queryMap["mid"] = fmt.Sprintf("%d", targetPlayerId)
	now := timeutils.TimeToMillisecond(time.Now()) / 1000
	queryMap["type"] = fmt.Sprintf("%s", chatType)
	queryMap["time"] = fmt.Sprintf("%d", now)
	queryMap["chat"] = fmt.Sprintf("%d", chatTime/1000)
	queryMap["uame"] = playerName
	queryMap["mame"] = targetPlayerName
	queryMap["server"] = fmt.Sprintf("%d", serverId)
	queryMap["body"] = body
	queryMap["sign"] = youQuSign(gameId, sdkUserId, now, key)
	query := ""
	for key, val := range queryMap {
		query += (key + "=" + val + "&")
	}
	query = query[:len(query)-1]
	eQuery := url.PathEscape(query)

	getUrl := jianKongUrl + "?" + eQuery

	//TODO 稍后尝试
	resp, err := http.Get(getUrl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"url":     jianKongUrl,
				"getUrl":  getUrl,
				"sdkType": sdkType.String(),
				"err":     err,
			}).Warn("请求错误")
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.WithFields(
			log.Fields{
				"url":     jianKongUrl,
				"getUrl":  getUrl,
				"sdkType": sdkType.String(),
				"status":  resp.StatusCode,
			}).Warn("状态码错误")
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(
			log.Fields{
				"url":     jianKongUrl,
				"getUrl":  getUrl,
				"sdkType": sdkType.String(),
				"err":     err,
			}).Warn("读取回复错误")
		return
	}
	respMap := make(map[string]interface{})
	err = json.Unmarshal(respBody, &respMap)
	if err != nil {
		log.WithFields(
			log.Fields{
				"respBody": string(respBody),
				"url":      jianKongUrl,
				"getUrl":   getUrl,
				"sdkType":  sdkType.String(),
				"err":      err,
			}).Warn("解析错误")
		return
	}
	respCode, ok := respMap["code"]
	if !ok {
		log.WithFields(
			log.Fields{
				"respBody": string(respBody),
				"url":      jianKongUrl,
				"getUrl":   getUrl,
				"sdkType":  sdkType.String(),
			}).Warn("解析错误")
		return
	}
	code, ok := respCode.(float64)
	if !ok {
		log.WithFields(
			log.Fields{

				"respBody": string(respBody),
				"url":      jianKongUrl,
				"getUrl":   getUrl,
				"sdkType":  sdkType.String(),
			}).Warn("解析错误")
		return
	}
	// codeInt, err := code.Int64()
	// if err != nil {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"respBody": string(respBody),
	// 			"url":      jianKongUrl,
	// 			"getUrl":   getUrl,
	// 			"sdkType":  sdkType.String(),
	// 			"err":      err,
	// 		}).Warn("解析错误")
	// 	return
	// }
	if int(code) != 1 {
		log.WithFields(
			log.Fields{
				"respBody": string(respBody),
				"url":      jianKongUrl,
				"getUrl":   getUrl,
				"sdkType":  sdkType.String(),
				"codeInt":  code,
			}).Warn("状态错误")
		return
	}
	log.WithFields(
		log.Fields{
			"url":     jianKongUrl,
			"getUrl":  getUrl,
			"sdkType": sdkType.String(),
		}).Info("发送成功")
}

func youQuSign(game int32, uid string, time int64, key string) string {
	allStr := fmt.Sprintf("%d%s%d%s", game, uid, time, key)
	hw := md5.Sum([]byte(allStr))
	return fmt.Sprintf("%x", hw)
}
