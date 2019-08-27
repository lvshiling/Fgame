package httputils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type RestResult struct {
	ErrorMsg  string      `json:"error_msg"`
	ErrorCode int         `json:"error_code"`
	Result    interface{} `json:"result"`
}

func NewSuccessResult(result interface{}) *RestResult {
	return &RestResult{
		Result: result,
	}
}

func NewFailedResult(errorCode int) *RestResult {
	return &RestResult{
		ErrorCode: errorCode,
	}
}

func NewFailedResultWithMsg(errorCode int, errorMsg string) *RestResult {
	return &RestResult{
		ErrorCode: errorCode,
		ErrorMsg:  errorMsg,
	}
}

func PostJson(apiPath string, headers map[string]string, form interface{}) (result interface{}, err error) {

	bodyBytes, err := json.Marshal(form)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", apiPath, bytes.NewBuffer(bodyBytes))
	request.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	client := http.DefaultClient
	if strings.Contains(apiPath, "https") {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println("Visit " + apiPath + "Respon:" + string(respBody))

	rr := &RestResult{}
	err = json.Unmarshal(respBody, rr)
	if err != nil {
		return nil, err
	}
	if rr.ErrorCode != 0 {
		return rr.Result, fmt.Errorf("error_code %d, error_msg %s", rr.ErrorCode, rr.ErrorMsg)
	}
	return rr.Result, nil
}

func PostJsonReturnByte(apiPath string, headers map[string]string, form interface{}) (result []byte, err error) {

	bodyBytes, err := json.Marshal(form)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", apiPath, bytes.NewBuffer(bodyBytes))
	request.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
	// rr := &RestResult{}
	// err = json.Unmarshal(respBody, rr)
	// if err != nil {
	// 	return nil, err
	// }
	// if rr.ErrorCode != 0 {
	// 	return rr.Result, fmt.Errorf("error_code %d", rr.ErrorCode)
	// }
	// return rr.Result, nil
}
