package authutils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

var (
	juHeUrl = "http://op.juhe.cn/idcard/query"
)
var (
	ErrorIdCardInvalid = fmt.Errorf("id card invalid")
)

func JuHeIdCardQuery(key string, name string, idCard string) (match bool, err error) {
	type JuHeIdCardResult struct {
		RealName string `json:"realname"`
		IdCard   string `json:"idCard"`
		Res      int32  `json:"res"`
	}
	type JuHeQueryRes struct {
		ErrorCode int               `json:"error_code"`
		Reason    string            `json:"reason"`
		Result    *JuHeIdCardResult `json:"result"`
	}

	query := "key=" + url.QueryEscape(key) + "&" + "idcard=" + url.QueryEscape(idCard) + "&" + "realname=" + url.QueryEscape(name)
	// eQuery := url.QueryEscape(query)
	getUrl := juHeUrl + "?" + query
	fmt.Println(getUrl)
	//TODO 稍后尝试
	// resp, err := http.Get(getUrl)
	httpClient := &http.Client{}
	httpClient.Timeout = remoteRequestTimeout
	resp, err := httpClient.Get(getUrl)
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
	r := &JuHeQueryRes{}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return
	}
	if r.ErrorCode != 0 {
		if r.ErrorCode == 210301 {
			return false, ErrorIdCardInvalid
		}
		return false, fmt.Errorf("rest error code %d", r.ErrorCode)
	}
	if r.Result == nil {
		return false, nil
	}
	if r.Result.Res == 1 {
		return true, nil
	}
	return false, nil
}

func IsIdCardValid(idCard string) (flag bool, err error) {
	r, err := regexp.Compile(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`)
	if err != nil {
		return
	}
	flag = r.MatchString(idCard)
	return
}

func IsPhoneValid(idCard string) (flag bool, err error) {
	r, err := regexp.Compile(`^[1][0-9]{10}$`)
	if err != nil {
		return
	}
	flag = r.MatchString(idCard)
	return
}
