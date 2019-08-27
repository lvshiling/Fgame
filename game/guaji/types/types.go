package types

import (
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//挂机提升
type GuaJiCheckType int32

const (
	GuaJiCheckTypeBag GuaJiCheckType = iota
	// GuaJiCheckTypeMount
	GuaJiCheckTypeSoul
	GuaJiCheckTypeSkill
	GuaJiCheckTypeBiaoChe
	GuaJiCheckTypeXueChi
	GuaJiCheckTypeAlliance
	GuaJiCheckTypeFuncOpen
	GuaJiCheckTypeEmail
)

var (
	guaJiCheckTypeMap = map[GuaJiCheckType]string{
		GuaJiCheckTypeBag: "背包",
		// GuaJiCheckTypeMount:    "坐骑",
		GuaJiCheckTypeSoul:     "帝魂",
		GuaJiCheckTypeSkill:    "技能",
		GuaJiCheckTypeBiaoChe:  "镖车",
		GuaJiCheckTypeXueChi:   "血池",
		GuaJiCheckTypeAlliance: "仙盟",
		GuaJiCheckTypeFuncOpen: "功能开启",
		GuaJiCheckTypeEmail:    "邮件",
	}
)

func (t GuaJiCheckType) String() string {
	return guaJiCheckTypeMap[t]
}

//挂机类型
type GuaJiType int32

const (
	GuaJiTypeMainQuest GuaJiType = iota
	GuaJiTypeDailyQuest
	GuaJiTypeMaterial
	GuaJiTypeXianFuSilver
	GuaJiTypeXianFuExp
	GuaJiTypeShuangXiu
	GuaJiTypeBiaoChe
	GuaJiTypeActivity
	GuaJiTypeWorldboss
	GuaJiTypeTianJieTa
	GuaJiTypeBaGua
	GuaJiTypeTeamFuBen
	GuaJiTypeDaBaoTa
	GuaJiTypeOutlandBoss
	GuaJiTypeUnrealBoss
	GuaJiTypeCangJingGe

	GuaJiTypeSoulRuins = 100
)

var (
	guaJiTypeMap = map[GuaJiType]string{
		GuaJiTypeMainQuest:    "主线任务",
		GuaJiTypeDailyQuest:   "日常任务",
		GuaJiTypeMaterial:     "材料副本",
		GuaJiTypeXianFuSilver: "银两副本",
		GuaJiTypeXianFuExp:    "经验副本",
		GuaJiTypeShuangXiu:    "双休副本",
		GuaJiTypeBiaoChe:      "镖车",
		GuaJiTypeActivity:     "活动",
		GuaJiTypeWorldboss:    "世界boss",
		GuaJiTypeTianJieTa:    "天劫塔",
		GuaJiTypeBaGua:        "八卦秘境",
		GuaJiTypeTeamFuBen:    "组队副本",
		GuaJiTypeDaBaoTa:      "打宝塔",
		GuaJiTypeOutlandBoss:  "外域boss",
		GuaJiTypeUnrealBoss:   "幻境boss",
		GuaJiTypeSoulRuins:    "帝魂",
		GuaJiTypeCangJingGe:   "藏经阁",
	}
)

func (t GuaJiType) String() string {
	return guaJiTypeMap[t]
}
func (t GuaJiType) Valid() bool {
	switch t {
	case GuaJiTypeMainQuest,
		GuaJiTypeDailyQuest,
		GuaJiTypeMaterial,
		GuaJiTypeXianFuSilver,
		GuaJiTypeXianFuExp,
		GuaJiTypeShuangXiu,
		GuaJiTypeBiaoChe,
		GuaJiTypeActivity,
		GuaJiTypeWorldboss,
		GuaJiTypeTianJieTa,
		GuaJiTypeBaGua,
		GuaJiTypeTeamFuBen,
		GuaJiTypeDaBaoTa,
		GuaJiTypeOutlandBoss,
		GuaJiTypeUnrealBoss,
		GuaJiTypeSoulRuins,
		GuaJiTypeCangJingGe:
		return true
	}
	return false
}

var (
	guaJiTypeAIMap = map[GuaJiType]scenetypes.GuaJiType{
		GuaJiTypeMainQuest:    scenetypes.GuaJiTypeQuest,
		GuaJiTypeDailyQuest:   scenetypes.GuaJiTypeDailyQuest,
		GuaJiTypeMaterial:     scenetypes.GuaJiTypeMaterial,
		GuaJiTypeXianFuSilver: scenetypes.GuaJiTypeXianFuSilver,
		GuaJiTypeXianFuExp:    scenetypes.GuaJiTypeXianFuExp,
		GuaJiTypeShuangXiu:    scenetypes.GuaJiTypeXianShuangXiu,
		GuaJiTypeBiaoChe:      scenetypes.GuaJiTypeBiaoChe,
		GuaJiTypeActivity:     scenetypes.GuaJiTypeActivity,
		GuaJiTypeWorldboss:    scenetypes.GuaJiTypeWorldBoss,
		GuaJiTypeTianJieTa:    scenetypes.GuaJiTypeTianJieTa,
		GuaJiTypeBaGua:        scenetypes.GuaJiTypeBaGua,
		GuaJiTypeDaBaoTa:      scenetypes.GuaJiTypeTower,
		GuaJiTypeOutlandBoss:  scenetypes.GuaJiTypeOutlandBoss,
		GuaJiTypeUnrealBoss:   scenetypes.GuaJiTypeUnrealBoss,
		GuaJiTypeCangJingGe:   scenetypes.GuaJiTypeCangJingGe,
	}
)

func (t GuaJiType) GuaJiAIType() (guaJiAIType scenetypes.GuaJiType, flag bool) {
	guaJiAIType, flag = guaJiTypeAIMap[t]
	return
}

type GuaJiOptionType interface {
	GetType() int32
	Valid() bool
	String() string
}

type GuaJiOptionTypeFactory interface {
	CreateGuaJiOptionType(typ int32) GuaJiOptionType
}

type GuaJiOptionTypeFactoryFunc func(typ int32) GuaJiOptionType

func (f GuaJiOptionTypeFactoryFunc) CreateGuaJiOptionType(typ int32) GuaJiOptionType {
	return f(typ)
}

var (
	guaJiOptionFactoryMap = map[GuaJiType]GuaJiOptionTypeFactory{}
)

func RegisterGuaJiOptionTypeFactory(typ GuaJiType, f GuaJiOptionTypeFactory) {
	_, ok := guaJiOptionFactoryMap[typ]
	if ok {
		panic(fmt.Errorf("重复注册[%s]挂机选项", typ.String()))
	}
	guaJiOptionFactoryMap[typ] = f
}

func GetGuaJiOptionTypeFactory(typ GuaJiType) GuaJiOptionTypeFactory {
	f, ok := guaJiOptionFactoryMap[typ]
	if !ok {
		return nil
	}
	return f
}

//挂机数据
type GuaJiData struct {
	typ       GuaJiType
	optionMap map[GuaJiOptionType]int32
}

func (d *GuaJiData) GetType() GuaJiType {
	return d.typ
}

func (d *GuaJiData) GetOptions() map[GuaJiOptionType]int32 {
	return d.optionMap
}
func (d *GuaJiData) GetOptionValue(option GuaJiOptionType) int32 {
	return d.optionMap[option]
}

func (d *GuaJiData) String() string {
	returnStr := fmt.Sprintf("挂机类型:%s", d.typ.String())
	for opType, opValue := range d.optionMap {
		gauJiOptionStr := fmt.Sprintf("挂机选项类型:%s,值:%d,", opType.String(), opValue)
		returnStr += gauJiOptionStr
	}
	return returnStr
}

func CreateGuaJiData(typ GuaJiType, optionMap map[GuaJiOptionType]int32) *GuaJiData {
	d := &GuaJiData{}
	d.typ = typ
	d.optionMap = optionMap
	return d
}
