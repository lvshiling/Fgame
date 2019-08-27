package httputils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	if err != nil {
		return nil, err
	}
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

	rr := &RestResult{}
	err = json.Unmarshal(respBody, rr)
	if err != nil {
		return nil, err
	}
	if rr.ErrorCode != 0 {
		return rr.Result, fmt.Errorf("error_code %d", rr.ErrorCode)
	}
	return rr.Result, nil
}

type RawRestResult struct {
	ErrorMsg  string          `json:"error_msg"`
	ErrorCode int             `json:"error_code"`
	Result    json.RawMessage `json:"result"`
}

func PostJsonWithRawMessage(apiPath string, headers map[string]string, form interface{}) (result *RawRestResult, err error) {
	bodyBytes, err := json.Marshal(form)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", apiPath, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
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

	result = &RawRestResult{}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		return nil, err
	}
	return
	// if rr.ErrorCode != 0 {
	// 	return rr.Result, fmt.Errorf("error_code %d", rr.ErrorCode)
	// }
	// return rr.Result, nil
}
