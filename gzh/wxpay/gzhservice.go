package wxpay

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type GzhApiService struct {
}

type GzhUserListRespon struct {
	Total      int    `json:"total"`
	Count      int    `json:"count"`
	NextOpenID string `json:"next_openid"`
	Data       struct {
		OpenID []string `json:"openid"`
	} `json:"data"`
}

func (m *GzhApiService) GetUserList(p_accesstoken string, p_nextopenid string) (*GzhUserListRespon, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s", p_accesstoken, p_nextopenid)
	rst := &GzhUserListRespon{}
	err := httpget(url, rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func httpget(p_url string, p_rst interface{}) error {
	request, _ := http.NewRequest("GET", p_url, nil)
	client := http.DefaultClient
	if strings.Contains(p_url, "https") {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// fmt.Println("Visit " + apiPath + "Respon:" + string(respBody))

	err = json.Unmarshal(respBody, p_rst)
	if err != nil {
		return err
	}
	return nil
}
