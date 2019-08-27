package sdk

import (
	"encoding/json"
	"fgame/fgame/account/login/types"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type SdkService interface {
	GetSdkConfig(platform types.SDKType) SDKConfig
}

type sdkService struct {
	sdkDir string
	sdkMap map[types.SDKType]SDKConfig
}

func (s *sdkService) init(dir string) (err error) {
	s.sdkMap = make(map[types.SDKType]SDKConfig)
	s.sdkDir = dir
	err = s.initSDK()
	if err != nil {
		return
	}
	return nil
}

func (s *sdkService) initSDK() (err error) {
	absolultDir, err := filepath.Abs(s.sdkDir)
	if err != nil {
		return
	}
	err = filepath.Walk(absolultDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		log.WithFields(log.Fields{
			"name": info.Name(),
		}).Info("sdk:读取sdk配置")
		typ, ok := sdkConfigMap[info.Name()]
		if !ok {
			log.WithFields(log.Fields{
				"name": info.Name(),
			}).Warn("sdk:结构不存在")
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		x := reflect.New(typ)
		val := x.Interface()
		err = json.Unmarshal(content, val)
		if err != nil {
			return err
		}
		sdkConfig, ok := val.(SDKConfig)
		if !ok {
			log.WithFields(log.Fields{
				"name": info.Name(),
			}).Warn("sdk:不是sdk配置结构")
			return nil
		}
		_, exist := s.sdkMap[sdkConfig.Platform()]
		if exist {
			err = fmt.Errorf("sdk:重复配置[%s],平台[%s]", sdkConfig.FileName(), sdkConfig.Platform().String())
			return err
		}
		s.sdkMap[sdkConfig.Platform()] = sdkConfig
		log.WithFields(log.Fields{
			"name":     info.Name(),
			"platform": sdkConfig.Platform().String(),
		}).Info("sdk:初始化sdk配置,成功")
		return nil
	})
	return
}

func (s *sdkService) GetSdkConfig(platform types.SDKType) SDKConfig {
	sdkConfig, ok := s.sdkMap[platform]
	if !ok {
		return nil
	}
	return sdkConfig
}

func newSdkService() *sdkService {
	s := &sdkService{}
	return s
}

var (
	once sync.Once
	s    *sdkService
)

func Init(sdkDir string) (err error) {
	once.Do(func() {
		s = newSdkService()
		err = s.init(sdkDir)
		return
	})
	return
}

func GetSdkService() SdkService {
	return s
}
