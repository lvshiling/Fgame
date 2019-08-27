package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/fourgod/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sync"

	droptemplate "fgame/fgame/game/drop/template"
	gametemplate "fgame/fgame/game/template"
)

type FourGodTemplateService interface {
	//获取四神遗迹宝箱配置
	//GetFourGodBoxTemplate(keyNum int32) *gametemplate.FourGodBoxTemplate
	GetFourGodBoxTemplateByBiologyId(biologyId int32) *gametemplate.FourGodBoxTemplate
	//获取四神遗迹常量模板
	GetFourGodConstTemplate() *gametemplate.FourGodTemplate
	//特殊怪配置默认出生地点
	GetFourGodSpecialDefaultPos() (defaultPos coretypes.Position, flag bool)
	//获取boss出生点
	GetFourGodBossPos() (bornPos coretypes.Position, flag bool)
	//获取假人出生的随机钥匙
	GetRobotBornKey() (keyNum int32)
	//获取公告阈值
	GetBossThreshold(oldPercent int32, newPercent int32) (percent int32, flag bool)
}

type fourGodTemplateService struct {
	//四神遗迹常量模板
	fourGodConstTemplate *gametemplate.FourGodTemplate
	//特殊怪
	specialMap map[types.FourGodSpecialPosType]map[int32]*gametemplate.FourGodSpecialTemplate
	//四神遗迹宝箱钥匙
	//fourGodBoxMap map[int32]*gametemplate.FourGodBoxTemplate
	fourGodBoxBiologyMap map[int32]*gametemplate.FourGodBoxTemplate
}

//初始化
func (fgts *fourGodTemplateService) init() error {
	fgts.specialMap = make(map[types.FourGodSpecialPosType]map[int32]*gametemplate.FourGodSpecialTemplate)
	//fgts.fourGodBoxMap = make(map[int32]*gametemplate.FourGodBoxTemplate)
	fgts.fourGodBoxBiologyMap = make(map[int32]*gametemplate.FourGodBoxTemplate)
	//四神常量
	templateMap := template.GetTemplateService().GetAll((*gametemplate.FourGodTemplate)(nil))
	for _, templateObject := range templateMap {
		fgts.fourGodConstTemplate, _ = templateObject.(*gametemplate.FourGodTemplate)
		break
	}

	//四神宝箱
	boxTemplateMap := template.GetTemplateService().GetAll((*gametemplate.FourGodBoxTemplate)(nil))
	for _, templateObject := range boxTemplateMap {
		fourGodBoxTemplate, _ := templateObject.(*gametemplate.FourGodBoxTemplate)
		//fgts.fourGodBoxMap[int32(fourGodBoxTemplate.TemplateId())] = fourGodBoxTemplate
		fgts.fourGodBoxBiologyMap[fourGodBoxTemplate.BiologyId] = fourGodBoxTemplate

		//验证掉落
		dropIdList := fourGodBoxTemplate.GetDropIdList()
		isSureDropNum := int32(0)
		for _, dropId := range dropIdList {
			flag := droptemplate.GetDropTemplateService().CheckSureDrop(dropId)
			if flag {
				isSureDropNum++
			}
		}
		if isSureDropNum == 0 {
			return fmt.Errorf("fourgod: 宝箱配置掉落应该是必定掉落的")
		}
	}

	//四神特殊怪
	templateMap = template.GetTemplateService().GetAll((*gametemplate.FourGodSpecialTemplate)(nil))
	for _, templateObject := range templateMap {
		specialTemplate, _ := templateObject.(*gametemplate.FourGodSpecialTemplate)
		typ := specialTemplate.GetTyp()

		specialTypMap, exist := fgts.specialMap[typ]
		if !exist {
			specialTypMap = make(map[int32]*gametemplate.FourGodSpecialTemplate)
			fgts.specialMap[typ] = specialTypMap
		}
		specialTypMap[int32(specialTemplate.TemplateId())] = specialTemplate
	}

	_, exist := fgts.GetFourGodBossPos()
	if !exist {
		return fmt.Errorf("fourgod: boss的位置应该存在的")
	}

	_, exist = fgts.GetFourGodSpecialDefaultPos()
	if !exist {
		return fmt.Errorf("fourgod: 特殊怪的默认出生地点应该是存在的")
	}
	return nil
}

//获取四神遗迹宝箱配置
// func (fgts *fourGodTemplateService) GetFourGodBoxTemplate(keyNum int32) (boxTemplate *gametemplate.FourGodBoxTemplate) {
// 	if keyNum <= 0 {
// 		return
// 	}
// 	for _, fourGodBoxTemplate := range fgts.fourGodBoxMap {
// 		keyMin := fourGodBoxTemplate.KeyMin
// 		keyMax := fourGodBoxTemplate.KeyMax
// 		if keyMin <= keyNum && keyNum <= keyMax {
// 			boxTemplate = fourGodBoxTemplate
// 			break
// 		}
// 	}
// 	return
// }

func (fgts *fourGodTemplateService) GetFourGodBoxTemplateByBiologyId(biologyId int32) *gametemplate.FourGodBoxTemplate {
	fourGodBoxTemplate, ok := fgts.fourGodBoxBiologyMap[biologyId]
	if !ok {
		return nil
	}
	return fourGodBoxTemplate
}

//获取公告阈值
func (fgts *fourGodTemplateService) GetBossThreshold(oldPercent int32, newPercent int32) (percent int32, flag bool) {
	thresholdList := fgts.fourGodConstTemplate.GetGongGaoThresholdList()
	for _, threshold := range thresholdList {
		if !flag {
			percent = threshold
		}
		if threshold >= newPercent && threshold < oldPercent {
			flag = true
			if threshold <= percent {
				percent = threshold
			}
		}
	}
	return
}

//获取四神遗迹常量模板
func (fgts *fourGodTemplateService) GetFourGodConstTemplate() *gametemplate.FourGodTemplate {
	return fgts.fourGodConstTemplate
}

//获取假人出生的随机钥匙
func (fgts *fourGodTemplateService) GetRobotBornKey() (keyNum int32) {
	robotKeyMin := fgts.fourGodConstTemplate.RobotKeyMin
	robotKeyMax := fgts.fourGodConstTemplate.RobotKeyMax
	keyNum = int32(mathutils.RandomRange(int(robotKeyMin), int(robotKeyMax)))
	return
}

//特殊怪配置默认出生地点
func (fgts *fourGodTemplateService) GetFourGodSpecialDefaultPos() (defaultPos coretypes.Position, flag bool) {
	specialMap, exist := fgts.specialMap[types.FourGodSpecialPosTypeBorn]
	if !exist {
		return
	}
	for _, specialTemplate := range specialMap {
		defaultPos = specialTemplate.GetPos()
		break
	}
	flag = true
	return
}

//获取boss出生点
func (fgts *fourGodTemplateService) GetFourGodBossPos() (bornPos coretypes.Position, flag bool) {
	bossMap, exist := fgts.specialMap[types.FourGodBossBorn]
	if !exist {
		return
	}
	for _, specialTemplate := range bossMap {
		bornPos = specialTemplate.GetPos()
		break
	}
	flag = true
	return
}

var (
	once sync.Once
	cs   *fourGodTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &fourGodTemplateService{}
		err = cs.init()
	})
	return err
}

func GetFourGodTemplateService() FourGodTemplateService {
	return cs
}
