package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sort"
	"sync"
)

//打宝塔配置服务
type TowerTemplaterService interface {
	// 打宝塔配置
	GetTowerTemplate(floor int32) *gametemplate.TowerTemplate
	// 打宝塔地图id
	GetTowerMapIdMap() map[int32]*gametemplate.TowerTemplate
	//获取推荐打宝塔
	GetRecommentTower(lev int32) *gametemplate.TowerTemplate
}

type towerTemplaterService struct {
	floorMap          map[int32]*gametemplate.TowerTemplate
	mapIdMap          map[int32]*gametemplate.TowerTemplate
	towerTemplateList []*gametemplate.TowerTemplate
}

type towerTemplateList []*gametemplate.TowerTemplate

func (l towerTemplateList) Len() int {
	return len(l)
}

func (l towerTemplateList) Less(i, j int) bool {
	return l[i].LevelMin < l[j].LevelMax
}

func (l towerTemplateList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

//初始化
func (ts *towerTemplaterService) init() error {
	ts.floorMap = make(map[int32]*gametemplate.TowerTemplate)
	ts.mapIdMap = make(map[int32]*gametemplate.TowerTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.TowerTemplate)(nil))
	for _, temp := range templateMap {
		towerTemplate := temp.(*gametemplate.TowerTemplate)
		ts.floorMap[int32(towerTemplate.Id)] = towerTemplate

		// 地图id
		_, ok := ts.mapIdMap[towerTemplate.MapId]
		if !ok {
			ts.mapIdMap[towerTemplate.MapId] = towerTemplate
		}
		ts.towerTemplateList = append(ts.towerTemplateList, towerTemplate)
	}

	sort.Sort(towerTemplateList(ts.towerTemplateList))
	return nil
}

func (ts *towerTemplaterService) GetTowerTemplate(floor int32) *gametemplate.TowerTemplate {
	temp, ok := ts.floorMap[floor]
	if !ok {
		return nil
	}

	return temp
}

func (s *towerTemplaterService) GetTowerMapIdMap() map[int32]*gametemplate.TowerTemplate {
	return s.mapIdMap
}

func (s *towerTemplaterService) GetRecommentTower(lev int32) *gametemplate.TowerTemplate {
	for _, towerTemplate := range s.towerTemplateList {
		if lev < towerTemplate.LevelMin {
			return nil
		}
		if towerTemplate.LevelMin <= lev && lev <= towerTemplate.LevelMax {
			return towerTemplate
		}
	}
	return nil
}

var (
	once sync.Once
	cs   *towerTemplaterService
)

func Init() (err error) {
	once.Do(func() {
		cs = &towerTemplaterService{}
		err = cs.init()
	})
	return err
}

func GetTowerTemplateService() TowerTemplaterService {
	return cs
}
