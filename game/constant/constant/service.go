package constant

import (
	"fgame/fgame/core/template"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"sync"
)

//快捷缓存
//常量配置的整合
type ConstantService interface {
	GetPlayerCreateTemplate(role types.RoleType, sex types.SexType) *gametemplate.PlayerCreateTemplate
	GetForceTemplate() *gametemplate.ForceMouldTemplate
	GetConstant(ct constanttypes.ConstantType) int32
	//传送地图消耗
	GetChangeSceneItems() map[int32]int32
	GetActivityResetTime() int64
	// VIP礼包重置时间
	GetVipLiBaoResetTime() int64
}

//快捷缓存
//常量配置的整合
type constantService struct {
	//角色初始化表
	playerCreateTemplateMapOfMap map[types.RoleType]map[types.SexType]*gametemplate.PlayerCreateTemplate
	//常量表
	constantTemplateMap map[constanttypes.ConstantType]*gametemplate.ConstantTemplate
	//战力表
	forceTemplate *gametemplate.ForceMouldTemplate
	//传送地图消耗
	changeSceneItemMap map[int32]int32
	//获取活动重置时间
	activityResetTime int64
	//获取vip礼包重置时间
	vipLiBaoResetTime int64
}

func (cs *constantService) init() (err error) {
	err = cs.initPlayerCreate()
	if err != nil {
		return
	}
	err = cs.initConstant()
	if err != nil {
		return
	}
	err = cs.initForce()
	if err != nil {
		return
	}
	return nil
}

//初始化角色创建
func (cs *constantService) initPlayerCreate() (err error) {
	cs.playerCreateTemplateMapOfMap = make(map[types.RoleType]map[types.SexType]*gametemplate.PlayerCreateTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.PlayerCreateTemplate)(nil))
	for _, templateObject := range templateMap {
		playerCreateTemplate, _ := templateObject.(*gametemplate.PlayerCreateTemplate)
		playerCreateTemplateMap, exist := cs.playerCreateTemplateMapOfMap[playerCreateTemplate.GetRole()]
		if !exist {
			playerCreateTemplateMap = make(map[types.SexType]*gametemplate.PlayerCreateTemplate)
			cs.playerCreateTemplateMapOfMap[playerCreateTemplate.GetRole()] = playerCreateTemplateMap
		}
		playerCreateTemplateMap[playerCreateTemplate.GetSex()] = playerCreateTemplate
	}
	//TODO 检查各种角色和性别配置
	return nil
}

//初始化常量配置
func (cs *constantService) initConstant() (err error) {
	cs.constantTemplateMap = make(map[constanttypes.ConstantType]*gametemplate.ConstantTemplate)

	templateConstantMap := template.GetTemplateService().GetAll((*gametemplate.ConstantTemplate)(nil))
	for _, templateObject := range templateConstantMap {
		tempConstant, _ := templateObject.(*gametemplate.ConstantTemplate)
		cs.constantTemplateMap[tempConstant.GetConstantType()] = tempConstant
	}

	for ct := constanttypes.ConstantTypeMin; ct <= constanttypes.ConstantTypeMax; ct++ {
		_, exist := cs.constantTemplateMap[ct]
		if !exist {
			return fmt.Errorf("constant:%d no exist", ct)
		}
	}

	cs.changeSceneItemMap = make(map[int32]int32)
	itemId := int32(cs.constantTemplateMap[constanttypes.ConstantTypeChangeSceneCostItem].Value)
	itemNum := int32(cs.constantTemplateMap[constanttypes.ConstantTypeChangeSceneCostItemNum].Value)
	if itemNum > 0 {
		cs.changeSceneItemMap[itemId] = itemNum
	}
	err = cs.initActivityResetTime()
	if err != nil {
		return
	}
	err = cs.initVipLiBaoResetTime()
	if err != nil {
		return
	}
	return nil
}

func (cs *constantService) initActivityResetTime() (err error) {
	constantTemplate, ok := cs.constantTemplateMap[constanttypes.ConstantTypeKaiFuHuoDongChongZhi]
	if !ok {
		return fmt.Errorf("开服活动重置时间没配")
	}
	timeStr := fmt.Sprintf("%d", constantTemplate.Value)
	resetTime, err := timeutils.ParseDayOfYYYYMMDDHHMMSS(timeStr)
	if err != nil {
		return err
	}
	cs.activityResetTime = resetTime
	return nil
}

func (cs *constantService) initVipLiBaoResetTime() (err error) {
	constantTemplate, ok := cs.constantTemplateMap[constanttypes.ConstantTypeVipLiBaoResetTime]
	if !ok {
		return fmt.Errorf("开服活动重置时间没配")
	}
	timeStr := fmt.Sprintf("%d", constantTemplate.Value)
	resetTime, err := timeutils.ParseDayOfYYYYMMDDHHMMSS(timeStr)
	if err != nil {
		return err
	}
	cs.vipLiBaoResetTime = resetTime
	return nil
}

//初始化战力表
func (cs *constantService) initForce() (err error) {
	tempForceTemplate := template.GetTemplateService().GetAll((*gametemplate.ForceMouldTemplate)(nil))[1]
	if tempForceTemplate == nil {
		return fmt.Errorf("constant:战力表不存在")
	}
	cs.forceTemplate = tempForceTemplate.(*gametemplate.ForceMouldTemplate)
	return nil
}

//获取角色创建模板
func (cs *constantService) GetPlayerCreateTemplate(role types.RoleType, sex types.SexType) *gametemplate.PlayerCreateTemplate {
	return cs.playerCreateTemplateMapOfMap[role][sex]
}

func (cs *constantService) GetConstant(ct constanttypes.ConstantType) (val int32) {
	tem, exist := cs.constantTemplateMap[ct]
	if !exist {
		return 0
	}
	return int32(tem.Value)
}

func (cs *constantService) GetForceTemplate() *gametemplate.ForceMouldTemplate {
	return cs.forceTemplate
}

func (cs *constantService) GetChangeSceneItems() map[int32]int32 {
	return cs.changeSceneItemMap
}

func (cs *constantService) GetActivityResetTime() int64 {
	return cs.activityResetTime
}

func (cs *constantService) GetVipLiBaoResetTime() int64 {
	return cs.vipLiBaoResetTime
}

var (
	once sync.Once
	cs   *constantService
)

func Init() (err error) {
	once.Do(func() {
		cs = &constantService{}
		err = cs.init()
	})
	return err
}

func GetConstantService() ConstantService {
	return cs
}
