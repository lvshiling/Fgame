package module

import (
	"fmt"
)

//模块初始化
type Module interface {
	InitTemplate() error
	Init() error
	Start()
	String() string
	Stop()
}

var (
	baseModuleMap     = make(map[string]Module)
	moduleMap         = make(map[string]Module)
	activityModuleMap = make(map[string]Module)

	moduleMapOfMap = make(map[string]map[string]Module)
)

const (
	base     = "base"
	game     = "game"
	activity = "activity"
)

func register(part string, mi Module) {
	moduleMap, ok := moduleMapOfMap[part]
	if !ok {
		moduleMap = make(map[string]Module)
		moduleMapOfMap[part] = moduleMap
	}
	_, exist := moduleMap[mi.String()]
	if exist {
		panic(fmt.Errorf("module:重复注册模块%s", mi.String()))
	}

	moduleMap[mi.String()] = mi
}

func RegisterBase(mi Module) {
	register(base, mi)
}

func Register(mi Module) {
	register(game, mi)
}

func RegisterActivityModule(mi Module) {
	register(activity, mi)
}

func GetBaseModules() map[string]Module {
	return moduleMapOfMap[base]
}

func GetModules() map[string]Module {
	return moduleMapOfMap[game]
}

func GetActivityModules() map[string]Module {
	return moduleMapOfMap[activity]
}

func RegisterModule(part string, m Module) {
	register(part, m)
}
