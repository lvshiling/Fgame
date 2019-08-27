package sdk

import (
	"fgame/fgame/account/login/types"
	"fmt"
	"reflect"
)

type SDKConfig interface {
	Platform() types.SDKType
	GetGame(types.DevicePlatformType) int32
	GetKey(types.DevicePlatformType) string
	GetUrl(types.DevicePlatformType) string
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
