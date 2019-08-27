package login

import (
	"fgame/fgame/account/login/types"
	"fmt"
)

type LoginHandler interface {
	Login(devicePlatformType types.DevicePlatformType, data interface{}) (flag bool, returnPlatform int32, userId string, err error)
}

type LoginHandlerFunc func(devicePlatformType types.DevicePlatformType, data interface{}) (flag bool, returnPlatform int32, userId string, err error)

func (f LoginHandlerFunc) Login(devicePlatformType types.DevicePlatformType, data interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	return f(devicePlatformType, data)
}

var (
	loginMap = make(map[types.SDKType]LoginHandler)
)

func RegisterLogin(sdkType types.SDKType, h LoginHandler) {
	_, ok := loginMap[sdkType]
	if ok {
		panic(fmt.Errorf("login:重复注册[%s]平台", sdkType.String()))
	}
	loginMap[sdkType] = h
}

func GetLoginHandler(sdkType types.SDKType) LoginHandler {
	h, ok := loginMap[sdkType]
	if !ok {
		return nil
	}
	return h
}
