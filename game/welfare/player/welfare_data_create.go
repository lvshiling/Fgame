package player

import (
	"encoding/json"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
	"reflect"
)

var (
	activityDataMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]reflect.Type)
)

func RegisterOpenActivityData(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, od welfaretypes.OpenActivityData) {
	subDataMap, ok := activityDataMap[typ]
	if !ok {
		subDataMap = make(map[welfaretypes.OpenActivitySubType]reflect.Type)
		activityDataMap[typ] = subDataMap
	}
	_, ok = subDataMap[subType]
	if ok {
		panic(fmt.Errorf("welfare:repeat register open activity; type:%d,subType:%d", typ, subType))
	}

	subDataMap[subType] = reflect.TypeOf(od).Elem()
}

func CreateOpenActivityData(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, content string) (data welfaretypes.OpenActivityData, err error) {
	subDataMap, ok := activityDataMap[typ]
	if !ok {
		return struct{}{}, nil
	}
	dataType, ok := subDataMap[subType]
	if !ok {
		return struct{}{}, nil
	}

	x := reflect.New(dataType)
	err = json.Unmarshal([]byte(content), x.Interface())
	if err != nil {
		return
	}
	data = x.Interface().(welfaretypes.OpenActivityData)
	return
}

func CreateEmptyOpenActivityData(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) (data welfaretypes.OpenActivityData) {
	subDataMap, ok := activityDataMap[typ]
	if !ok {
		return struct{}{}
	}
	dataType, ok := subDataMap[subType]
	if !ok {
		return struct{}{}
	}

	x := reflect.New(dataType)
	data = x.Interface().(welfaretypes.OpenActivityData)
	return
}
