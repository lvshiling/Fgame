package authutils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var (
	juHeUrl = "http://op.juhe.cn/idcard/query"
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

	query := "key=" + key + "&" + "idcard=" + idCard + "&" + "realname=" + name
	eQuery := url.PathEscape(query)
	getUrl := juHeUrl + "?" + eQuery
	//TODO 稍后尝试
	resp, err := http.Get(getUrl)
	if err != nil {
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
		return false, fmt.Errorf("rest error code %d", r.ErrorCode)
	}
	if r.Result == nil {
		return false, nil
	}
	if r.Result.Res == 0 {
		return false, nil
	}
	return true, nil
}
