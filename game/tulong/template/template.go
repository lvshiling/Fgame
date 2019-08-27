package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/utils"
	gametemplate "fgame/fgame/game/template"
	tulongtypes "fgame/fgame/game/tulong/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sync"
)

//屠龙接口处理
type TuLongTemplateService interface {
	//获取屠龙小boss模板
	GetTuLongSmallBossTemplate(biaoshi int32) *gametemplate.TuLongTemplate
	//获取大boss模板
	GetTuLongBigBossTemplate() *gametemplate.TuLongTemplate
	//获取屠龙常量配置
	GetTuLongConstTemplate() *gametemplate.TuLongConstantTemplate

	//获取出生模板
	GetTuLongPosTemplate(posType tulongtypes.TuLongPosType, biaoShi int32) *gametemplate.TuLongPosTemplate
	//获取出生标识
	GetTuLongPosBiaoShi(posType tulongtypes.TuLongPosType, excludeBiaoShi []int32) (biaoShi int32, flag bool)

	//获取屠龙模板Map
	GetTuLongMap() map[tulongtypes.TuLongBossType]map[int32]*gametemplate.TuLongTemplate
	//获取屠龙模板大小
	GetTuLongLen() int32
	//获取屠龙模板通过生物id
	GetTuLongTemplateByBiologyId(biologyId int32) (*gametemplate.TuLongTemplate, bool)
}

type tuLongTemplateService struct {
	//屠龙模板
	tuLongTemplateMap map[tulongtypes.TuLongBossType]map[int32]*gametemplate.TuLongTemplate
	//屠龙模板
	tuLongBiologyTemplateMap map[int32]*gametemplate.TuLongTemplate
	//屠龙常量模板
	tuLongConstTemplate *gametemplate.TuLongConstantTemplate
	//屠龙出生模板
	tuLongPosTemplateMap map[tulongtypes.TuLongPosType]map[int32]*gametemplate.TuLongPosTemplate
}

//初始化
func (ts *tuLongTemplateService) init() error {
	ts.tuLongTemplateMap = make(map[tulongtypes.TuLongBossType]map[int32]*gametemplate.TuLongTemplate)
	ts.tuLongBiologyTemplateMap = make(map[int32]*gametemplate.TuLongTemplate)
	ts.tuLongPosTemplateMap = make(map[tulongtypes.TuLongPosType]map[int32]*gametemplate.TuLongPosTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.TuLongTemplate)(nil))
	for _, templateObject := range templateMap {
		tuLongTemplate, _ := templateObject.(*gametemplate.TuLongTemplate)

		tuLongBossType := tuLongTemplate.GetTuLongType()
		tuLongTypeMap, exist := ts.tuLongTemplateMap[tuLongBossType]
		if !exist {
			tuLongTypeMap = make(map[int32]*gametemplate.TuLongTemplate)
			ts.tuLongTemplateMap[tuLongBossType] = tuLongTypeMap
		}
		tuLongTypeMap[tuLongTemplate.BiaoShi] = tuLongTemplate

		//
		_, exist = ts.tuLongBiologyTemplateMap[tuLongTemplate.BiologyId]
		if exist {
			return fmt.Errorf("tulong:生物id应该是唯一的")
		}
		ts.tuLongBiologyTemplateMap[tuLongTemplate.BiologyId] = tuLongTemplate
	}

	constTemplateMap := template.GetTemplateService().GetAll((*gametemplate.TuLongConstantTemplate)(nil))
	if len(constTemplateMap) != 1 {
		return fmt.Errorf("tulong:屠龙常量模板有且只有一条")
	}
	ts.tuLongConstTemplate = constTemplateMap[1].(*gametemplate.TuLongConstantTemplate)

	posTemplateMap := template.GetTemplateService().GetAll((*gametemplate.TuLongPosTemplate)(nil))
	for _, templateObject := range posTemplateMap {
		tuLongPosTemplate, _ := templateObject.(*gametemplate.TuLongPosTemplate)

		tuLongPosType := tuLongPosTemplate.GetPosType()
		tuLongPosMap, exist := ts.tuLongPosTemplateMap[tuLongPosType]
		if !exist {
			tuLongPosMap = make(map[int32]*gametemplate.TuLongPosTemplate)
			ts.tuLongPosTemplateMap[tuLongPosType] = tuLongPosMap
		}
		tuLongPosMap[tuLongPosTemplate.BiaoShi] = tuLongPosTemplate
	}

	//校验
	bigBossMap, exist := ts.tuLongTemplateMap[tulongtypes.TuLongBossTypeBig]
	if !exist {
		return fmt.Errorf("tulong:大boss的配置应该是ok")
	}
	_, exist = bigBossMap[1]
	if !exist {
		return fmt.Errorf("tulong:大boss的配置应该是ok")
	}

	return nil
}

//获取屠龙小boss模板
func (ts *tuLongTemplateService) GetTuLongSmallBossTemplate(biaoshi int32) *gametemplate.TuLongTemplate {
	smallBossMap, exist := ts.tuLongTemplateMap[tulongtypes.TuLongBossTypeSmall]
	if !exist {
		return nil
	}
	tulongTemplate, exist := smallBossMap[biaoshi]
	if !exist {
		return nil
	}
	return tulongTemplate
}

//获取大boss模板
func (ts *tuLongTemplateService) GetTuLongBigBossTemplate() *gametemplate.TuLongTemplate {
	bigBossMap, exist := ts.tuLongTemplateMap[tulongtypes.TuLongBossTypeBig]
	if !exist {
		return nil
	}
	tulongTemplate, exist := bigBossMap[1]
	if !exist {
		return nil
	}
	return tulongTemplate
}

//获取屠龙常量配置
func (ts *tuLongTemplateService) GetTuLongConstTemplate() *gametemplate.TuLongConstantTemplate {
	return ts.tuLongConstTemplate
}

//获取屠龙出生模板
func (ts *tuLongTemplateService) GetTuLongPosTemplate(posType tulongtypes.TuLongPosType, biaoshi int32) *gametemplate.TuLongPosTemplate {
	posTypMap, exist := ts.tuLongPosTemplateMap[posType]
	if !exist {
		return nil
	}
	tulongPosTemplate, exist := posTypMap[biaoshi]
	if !exist {
		return nil
	}
	return tulongPosTemplate
}

func (ts *tuLongTemplateService) GetTuLongTemplateByBiologyId(biologyId int32) (*gametemplate.TuLongTemplate, bool) {
	tuLongTemplate, exist := ts.tuLongBiologyTemplateMap[biologyId]
	if !exist {
		return nil, false
	}
	return tuLongTemplate, true
}

//获取屠龙模板Map
func (ts *tuLongTemplateService) GetTuLongMap() map[tulongtypes.TuLongBossType]map[int32]*gametemplate.TuLongTemplate {
	return ts.tuLongTemplateMap
}

//获取屠龙模板大小
func (ts *tuLongTemplateService) GetTuLongLen() int32 {
	return int32(len(ts.tuLongBiologyTemplateMap))
}

//获取出生标识
func (ts *tuLongTemplateService) GetTuLongPosBiaoShi(posType tulongtypes.TuLongPosType, excludeBiaoShi []int32) (biaoShi int32, flag bool) {
	var biaoShiList []int32
	var weightList []int64

	for biaoShi, _ := range ts.tuLongPosTemplateMap[posType] {
		flag := utils.ContainInt32(excludeBiaoShi, biaoShi)
		if flag {
			continue
		}
		biaoShiList = append(biaoShiList, biaoShi)
		weightList = append(weightList, 1)
	}

	if len(weightList) == 0 {
		return
	}

	index := mathutils.RandomWeights(weightList)
	return biaoShiList[index], true
}

var (
	once sync.Once
	cs   *tuLongTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &tuLongTemplateService{}
		err = cs.init()
	})
	return err
}

func GetTuLongTemplateService() TuLongTemplateService {
	return cs
}
