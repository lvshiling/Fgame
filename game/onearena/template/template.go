package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/onearena/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//灵池争夺接口处理
type OneArenaTemplateService interface {
	//获取所有灵池
	GetOneArenaMap() map[types.OneArenaLevelType]map[int32]*gametemplate.OneArenaTemplate
	//获取灵池争夺配置
	GetOneArenaTemplateByLevel(level types.OneArenaLevelType, posId int32) *gametemplate.OneArenaTemplate
	//获取1v1鱼的存储上限
	GetOneArenaKunLimit() int32
}

type oneArenaTemplateService struct {
	oneArenaTemplateMap map[types.OneArenaLevelType]map[int32]*gametemplate.OneArenaTemplate
}

//初始化
func (os *oneArenaTemplateService) init() (err error) {
	os.oneArenaTemplateMap = make(map[types.OneArenaLevelType]map[int32]*gametemplate.OneArenaTemplate)

	//赋值oneArenaTemplateMap
	templateMap := template.GetTemplateService().GetAll((*gametemplate.OneArenaTemplate)(nil))
	for _, templateObject := range templateMap {
		oneArenaTemplate, _ := templateObject.(*gametemplate.OneArenaTemplate)

		typ := oneArenaTemplate.GetArenaType()

		oneArenaPosMap, ok := os.oneArenaTemplateMap[typ]
		if !ok {
			oneArenaPosMap = make(map[int32]*gametemplate.OneArenaTemplate)
			os.oneArenaTemplateMap[typ] = oneArenaPosMap
		}
		oneArenaPosMap[oneArenaTemplate.PosId] = oneArenaTemplate
	}

	return nil
}

func (os *oneArenaTemplateService) GetOneArenaMap() map[types.OneArenaLevelType]map[int32]*gametemplate.OneArenaTemplate {
	return os.oneArenaTemplateMap
}

func (os *oneArenaTemplateService) GetOneArenaTemplateByLevel(level types.OneArenaLevelType, posId int32) *gametemplate.OneArenaTemplate {
	posOneArenaMap, ok := os.oneArenaTemplateMap[level]
	if !ok {
		return nil
	}
	to, ok := posOneArenaMap[posId]
	if !ok {
		return nil
	}
	return to
}

func (os *oneArenaTemplateService) GetOneArenaKunLimit() int32 {
	return constant.GetConstantService().GetConstant(constanttypes.ConstantTypeOneArenaKunLimit)
}

var (
	once sync.Once
	cs   *oneArenaTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &oneArenaTemplateService{}
		err = cs.init()
	})
	return err
}

func GetOneArenaTemplateService() OneArenaTemplateService {
	return cs
}
