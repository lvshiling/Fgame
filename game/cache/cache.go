package cache

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/cache/cache"
	"fgame/fgame/game/cache/dao"
	"fgame/fgame/game/global"
)

//注册管理器
import (
	_ "fgame/fgame/game/cache/event/listener"
	_ "fgame/fgame/game/cache/player"
)

//缓存模块
type cacheModule struct {
}

func (m *cacheModule) InitTemplate() (err error) {

	return
}

func (m *cacheModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = cache.Init()
	if err != nil {
		return
	}
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *cacheModule) Start() {

}

func (m *cacheModule) Stop() {

}

func (m *cacheModule) String() string {
	return "cache"
}

var (
	m = &cacheModule{}
)

func init() {
	module.Register(m)
}
