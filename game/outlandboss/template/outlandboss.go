package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"sync"
)

//外域BOSS配置服务
type OutlandBossTemplateService interface {
	// 外域BOSS配置
	GetOutlandBossTemplate(biologyId int32) *gametemplate.OutlandBossTemplate
	// 外域地图列表
	GetMapIdList() []int32
	//随机获取外域boss配置
	GetOutlandbossTemplateRandom(excludeId int32) *gametemplate.OutlandBossTemplate
}

type outlandbossTemplateService struct {
	bossMap   map[int32]*gametemplate.OutlandBossTemplate
	mapIdList []int32
}

//初始化
func (ts *outlandbossTemplateService) init() error {
	ts.bossMap = make(map[int32]*gametemplate.OutlandBossTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.OutlandBossTemplate)(nil))
	for _, temp := range templateMap {
		outlandbossTemplate, _ := temp.(*gametemplate.OutlandBossTemplate)
		ts.bossMap[outlandbossTemplate.BiologyId] = outlandbossTemplate

		if !ts.isExsitMap(outlandbossTemplate.MapId) {
			ts.mapIdList = append(ts.mapIdList, outlandbossTemplate.MapId)
		}
	}

	return nil
}

func (s *outlandbossTemplateService) GetMapIdList() []int32 {
	return s.mapIdList
}

func (ts *outlandbossTemplateService) GetOutlandBossTemplate(biologyId int32) *gametemplate.OutlandBossTemplate {
	temp, ok := ts.bossMap[biologyId]
	if !ok {
		return nil
	}

	return temp
}

func (s *outlandbossTemplateService) isExsitMap(mapId int32) bool {
	for _, value := range s.mapIdList {
		if value == mapId {
			return true
		}
	}

	return false
}

//随机获取外域boss配置
func (s *outlandbossTemplateService) GetOutlandbossTemplateRandom(excludeId int32) *gametemplate.OutlandBossTemplate {
	weights := make([]int64, 0, len(s.bossMap)-1)
	tempList := make([]*gametemplate.OutlandBossTemplate, 0, len(s.bossMap)-1)
	for _, ch := range s.bossMap {
		if ch.BiologyId == excludeId {
			continue
		}
		weights = append(weights, int64(ch.Rate))
		tempList = append(tempList, ch)
	}
	index := mathutils.RandomWeights(weights)
	if index == -1 {
		return nil
	}
	ch := tempList[index]
	return ch
}

var (
	once sync.Once
	cs   *outlandbossTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &outlandbossTemplateService{}
		err = cs.init()
	})
	return err
}

func GetOutlandBossTemplateService() OutlandBossTemplateService {
	return cs
}
