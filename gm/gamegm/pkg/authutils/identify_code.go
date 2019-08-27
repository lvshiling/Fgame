package authutils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
)

var (
	smsUrl = "http://dysmsapi.aliyuncs.com/?Signature="
)

func randomCode() string {
	now := time.Now().UnixNano()
	rand.Seed(now)
	num := rand.Intn(1000000)
	return fmt.Sprintf("%06d", num)
}

var (
	ErrorPhoneInvalid    = fmt.Errorf("phone number invalid")
	ErrorSMSLimitControl = fmt.Errorf("sms limit control")
	ErrorTimeout         = fmt.Errorf("error timeout")
)

const (
	remoteRequestTimeout = 10 * time.Second
)

func GetIdentifyCode(accessKeyId string, accessSecret string, phoneNum string) (code string, err error) {
	code = randomCode()

	params := make(map[string]string)
	params["SignatureMethod"] = "HMAC-SHA1"
	params["SignatureNonce"] = uuid.NewV1().String()
	params["AccessKeyId"] = accessKeyId
	params["SignatureVersion"] = "1.0"
	gmt, err := time.LoadLocation("GMT")
	if err != nil {
		return
	}
	tz := time.Now().In(gmt).Format("2006-01-02T15:04:05Z")
	fmt.Println(tz)
	params["Timestamp"] = tz
	params["Format"] = "XML"
	params["Action"] = "SendSms"
	params["Version"] = "2017-05-25"
	params["RegionId"] = "cn-hangzhou"
	params["PhoneNumbers"] = phoneNum
	signName := "丁丁互娱"
	params["SignName"] = signName
	templateParam := fmt.Sprintf("{\"code\":\"%s\"}", code)
	params["TemplateParam"] = templateParam
	templateCode := "SMS_130200228"
	params["TemplateCode"] = templateCode

	q := Query(params)
	signString := sign(accessSecret, q)
	newUrl := smsUrl + signString + "&" + q

	//TODO 稍后尝试
	httpClient := &http.Client{}
	httpClient.Timeout = remoteRequestTimeout
	resp, err := httpClient.Get(newUrl)
	if err != nil {
		terr, ok := err.(*url.Error)
		if ok {
			if terr.Timeout() {
				err = ErrorTimeout
				return
			}
		}
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http response code %d", resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	type SMSResp struct {
		XMLName   xml.Name `xml:"SendSmsResponse"`
		Message   string   `xml:"Message"`
		RequestId string   `xml:"RequestId"`
		Code      string   `xml:"Code"`
		BizId     string   `xml:"BizId"`
	}
	smsResp := &SMSResp{}
	err = xml.Unmarshal(bs, smsResp)
	if err != nil {
		return "", err
	}
	if smsResp.Code == "OK" {
		return
	}

	log.WithFields(
		log.Fields{
			"code":      smsResp.Code,
			"bizId":     smsResp.BizId,
			"message":   smsResp.Message,
			"requestId": smsResp.RequestId,
		},
	).Warn("发送验证码,错误")
	if smsResp.Code == "isv.MOBILE_NUMBER_ILLEGAL" {
		err = ErrorPhoneInvalid
		return
	}
	if smsResp.Code == "isv.BUSINESS_LIMIT_CONTROL" {
		err = ErrorSMSLimitControl
		return
	}
	return "", fmt.Errorf("sms send code %s", smsResp.Code)
}

func sign(accessSecret string, query string) string {
	stringToSign := pop(query)
	accessSecret += "&"
	h := hmac.New(sha1.New, []byte(accessSecret))
	h.Write([]byte(stringToSign))
	signBytes := h.Sum(nil)
	base64Sign := base64.StdEncoding.EncodeToString(signBytes)
	finalSign := specialUrlEncode(base64Sign)
	return finalSign
}

func pop(query string) string {
	stringToSign := "GET&" + specialUrlEncode("/") + "&"
	stringToSign += specialUrlEncode(query)
	return stringToSign
}

func Query(params map[string]string) string {
	var buf bytes.Buffer
	keys := make([]string, 0, len(params))
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := params[k]
		prefix := specialUrlEncode(k) + "="
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(prefix)
		buf.WriteString(specialUrlEncode(v))

	}
	q := buf.String()
	return q
}

func specialUrlEncode(value string) string {
	t := url.QueryEscape(value)
	t = strings.Replace(t, "+", "%20", -1)
	t = strings.Replace(t, "*", "%2A", -1)
	t = strings.Replace(t, "%7E", "~", -1)
	return t
}
