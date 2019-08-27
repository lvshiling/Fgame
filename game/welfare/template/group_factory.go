package template

import (
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

type GroupTemplateIFactory interface {
	CreateGroupTemplate(base *GroupTemplateBase) GroupTemplateI
}

type GroupTemplateIFactoryFunc func(base *GroupTemplateBase) GroupTemplateI

func (hf GroupTemplateIFactoryFunc) CreateGroupTemplate(base *GroupTemplateBase) GroupTemplateI {
	return hf(base)
}

var (
	groupFactoryMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]GroupTemplateIFactory)
)

func RegisterGroupTemplate(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, g GroupTemplateIFactory) {
	subDataMap, ok := groupFactoryMap[typ]
	if !ok {
		subDataMap = make(map[welfaretypes.OpenActivitySubType]GroupTemplateIFactory)
		groupFactoryMap[typ] = subDataMap
	}
	_, ok = subDataMap[subType]
	if ok {
		panic(fmt.Errorf("welfare:repeat register open activity; OpenActivitySubType:%d", subType))
	}

	subDataMap[subType] = g
}

func CreateGroupTemplateI(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, base *GroupTemplateBase) (data GroupTemplateI) {
	subFactoryMap, ok := groupFactoryMap[typ]
	if !ok {
		return base
	}
	h, ok := subFactoryMap[subType]
	if !ok {
		return base
	}

	return h.CreateGroupTemplate(base)
}
