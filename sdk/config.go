package sdk

import (
	"fgame/fgame/account/login/types"
	"fmt"
	"reflect"
)

type SDKConfig interface {
	Platform() types.SDKType
	FileName() string
}

var (
	sdkConfigMap map[string]reflect.Type
)

func init() {
	sdkConfigMap = make(map[string]reflect.Type)
}

func RegisterSDK(to SDKConfig) {
	_, exist := sdkConfigMap[to.FileName()]
	if exist {
		panic(fmt.Sprintf("repeate register %s sdk object", to.FileName()))
	}
	sdkConfigMap[to.FileName()] = reflect.TypeOf(to).Elem()
}
