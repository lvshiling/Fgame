package rest

import (
	"encoding/json"
)

//rest接口
type RestResult struct {
	ErrorCode int             `json:"errorCode"`
	Result    json.RawMessage `json:"result"`
}
