package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//幻境BOSS配置服务
type UnrealBossTemplaterService interface {
	// 幻境BOSS配置
	GetUnrealBossTemplate(biologyId int32) *gametemplate.UnrealBossTemplate
	// 幻境地图列表
	GetMapIdList() []int32
}

type unrealbossTemplaterService struct {
	bossMap   map[int32]*gametemplate.UnrealBossTemplate
	mapIdList []int32
}

//初始化
func (ts *unrealbossTemplaterService) init() error {
	ts.bossMap = make(map[int32]*gametemplate.UnrealBossTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.UnrealBossTemplate)(nil))
	for _, temp := range templateMap {
		unrealbossTemplate, _ := temp.(*gametemplate.UnrealBossTemplate)
		ts.bossMap[unrealbossTemplate.BiologyId] = unrealbossTemplate

		if !ts.isExsitMap(unrealbossTemplate.MapId) {
			ts.mapIdList = append(ts.mapIdList, unrealbossTemplate.MapId)
		}

	}

	return nil
}

func (s *unrealbossTemplaterService) GetMapIdList() []int32 {
	return s.mapIdList
}

func (ts *unrealbossTemplaterService) GetUnrealBossTemplate(biologyId int32) *gametemplate.UnrealBossTemplate {
	temp, ok := ts.bossMap[biologyId]
	if !ok {
		return nil
	}

	return temp
}

func (s *unrealbossTemplaterService) isExsitMap(mapId int32) bool {
	for _, value := range s.mapIdList {
		if value == mapId {
			return true
		}
	}

	return false
}

var (
	once sync.Once
	cs   *unrealbossTemplaterService
)

func Init() (err error) {
	once.Do(func() {
		cs = &unrealbossTemplaterService{}
		err = cs.init()
	})
	return err
}

func GetUnrealBossTemplateService() UnrealBossTemplaterService {
	return cs
}
