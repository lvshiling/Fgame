package template

import (
	"fgame/fgame/core/template"
	babytypes "fgame/fgame/game/baby/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sync"
)

//宝宝配置服务
type BabyTemplaterService interface {
	// 宝宝常量模板
	GetBabyConstantTemplate() *gametemplate.BabyConstantTemplate
	// 宝宝怀孕模板
	GetBabyPregnantTemplateByQuality(quality int32) *gametemplate.BabyPregnantTemplate
	// 宝宝补品最大进度值
	GetBabyPregnantTonicMaxPro() int32
	// 宝宝天赋配置
	GetBabyTalentTemplate(skillId int32) *gametemplate.BabyTalentTemplate
	// 宝宝读书配置
	GetBabyLearnTemplate(level int32) *gametemplate.BabyLearnTemplate
	// 宝宝天赋解锁配置
	GetBabyUnlockTalentTemplate(times int32) *gametemplate.BabyUnlockTalentTemplate
	// 宝宝玩具升级配置
	GetBabyToyUplevelTemplate(suitType babytypes.ToySuitType, posType inventorytypes.BodyPositionType, level int32) *gametemplate.BabyToyUplevelTemplate
	// 宝宝洞房配置
	GetBabyDongFangTemplate(babyNum int32) *gametemplate.BabyDongFangTemplate
	// 宝宝属性配置
	GetBabyDanBeiTemplateMap() map[int32]*gametemplate.BabyBeiShuTemplate
	// 宝宝品质配置
	GetBabyQualityTemplate(tonicNum int32) *gametemplate.BabyQualityTemplate
	// 宝宝玩具套装模板
	GetBabyToyTemplateBySuitGroup(suitGroup int32) *gametemplate.BabyToySuitGroupTemplate
	// 获取宝宝类型模板
	GetBabyTypeTemplate() *gametemplate.BabyPregnantTemplate
}

const (
	maxRate int64 = 2000000000
)

type babyTemplaterService struct {
	babyConstantTemplate *gametemplate.BabyConstantTemplate
	//
	pregnantMap map[int32]*gametemplate.BabyPregnantTemplate
	//
	talentMap map[int32]*gametemplate.BabyTalentTemplate
	//
	learnMap map[int32]*gametemplate.BabyLearnTemplate
	//
	unlockTalentMap map[int32]*gametemplate.BabyUnlockTalentTemplate
	unlockMaxTimes  int32 //最大解锁次数
	//
	toyUplevelMap map[babytypes.ToySuitType]map[inventorytypes.BodyPositionType][]*gametemplate.BabyToyUplevelTemplate
	//
	toySuitGroupMap map[int32]*gametemplate.BabyToySuitGroupTemplate
	//
	dongfangMap map[int32]*gametemplate.BabyDongFangTemplate
	//
	danbeiMap map[int32]*gametemplate.BabyBeiShuTemplate
	//品质配置
	qualityTempMap       map[int32]*gametemplate.BabyQualityTemplate
	maxTonicNum          int32              //最大补品值
	qualityMap           map[int32]struct{} //所有品质
	babyTypeTemplateList []*gametemplate.BabyPregnantTemplate
}

//初始化
func (ts *babyTemplaterService) init() (err error) {

	// 宝宝常量配置
	err = ts.loadBaoBaoConstant()
	if err != nil {
		return
	}

	// 怀孕配置
	err = ts.loadBaoBaoPregnant()
	if err != nil {
		return
	}
	// 天赋配置
	err = ts.loadBaoBaoTalent()
	if err != nil {
		return
	}

	// 读书配置
	err = ts.loadBaoBaoLearn()
	if err != nil {
		return
	}

	// 天赋解锁配置
	err = ts.loadBaoBaoUnlockTalent()
	if err != nil {
		return
	}
	// 玩具升级配置
	err = ts.loadBaoBaoToyUplevel()
	if err != nil {
		return
	}
	// 洞房配置
	err = ts.loadBaoBaoDongFang()
	if err != nil {
		return
	}
	// 属性单倍配置
	err = ts.loadBaoBaoDanBei()
	if err != nil {
		return
	}
	// 宝宝品质配置
	err = ts.loadBaoBaoQuality()
	if err != nil {
		return
	}
	// 宝宝玩具套装配置
	err = ts.loadBaoBaoSuit()
	if err != nil {
		return
	}

	return nil
}

func (ts *babyTemplaterService) loadBaoBaoConstant() (err error) {
	templateMap := template.GetTemplateService().GetAll((*gametemplate.BabyConstantTemplate)(nil))
	if len(templateMap) != 1 {
		return fmt.Errorf("baby:怀孕配置只有一条")
	}
	for _, temp := range templateMap {
		constantTemp, _ := temp.(*gametemplate.BabyConstantTemplate)
		ts.babyConstantTemplate = constantTemp
	}

	return
}

func (ts *babyTemplaterService) loadBaoBaoPregnant() (err error) {
	totalRate := int64(0)
	ts.pregnantMap = make(map[int32]*gametemplate.BabyPregnantTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.BabyPregnantTemplate)(nil))
	for _, temp := range templateMap {
		pregnantTemplate, _ := temp.(*gametemplate.BabyPregnantTemplate)
		_, ok := ts.pregnantMap[pregnantTemplate.Quality]
		if ok {
			return fmt.Errorf("宝宝怀孕配置，品质：%d重复", pregnantTemplate.Quality)
		}
		ts.pregnantMap[pregnantTemplate.Quality] = pregnantTemplate
		ts.babyTypeTemplateList = append(ts.babyTypeTemplateList, pregnantTemplate)
		totalRate += pregnantTemplate.Rate
	}

	if totalRate != maxRate {
		return fmt.Errorf("宝宝怀孕配置, 出生概率错误")
	}
	return
}

func (ts *babyTemplaterService) loadBaoBaoTalent() (err error) {
	ts.talentMap = make(map[int32]*gametemplate.BabyTalentTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.BabyTalentTemplate)(nil))
	for _, temp := range templateMap {
		template, _ := temp.(*gametemplate.BabyTalentTemplate)
		ts.talentMap[template.SkillId] = template
	}
	return
}

func (ts *babyTemplaterService) loadBaoBaoLearn() (err error) {
	ts.learnMap = make(map[int32]*gametemplate.BabyLearnTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.BabyLearnTemplate)(nil))
	for _, temp := range templateMap {
		template, _ := temp.(*gametemplate.BabyLearnTemplate)
		ts.learnMap[template.Level] = template
	}
	return
}

func (ts *babyTemplaterService) loadBaoBaoUnlockTalent() (err error) {
	ts.unlockTalentMap = make(map[int32]*gametemplate.BabyUnlockTalentTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.BabyUnlockTalentTemplate)(nil))
	for _, temp := range templateMap {
		template, _ := temp.(*gametemplate.BabyUnlockTalentTemplate)
		ts.unlockTalentMap[template.Times] = template

		if template.Times > ts.unlockMaxTimes {
			ts.unlockMaxTimes = template.Times
		}
	}
	return
}

func (ts *babyTemplaterService) loadBaoBaoToyUplevel() (err error) {
	ts.toyUplevelMap = make(map[babytypes.ToySuitType]map[inventorytypes.BodyPositionType][]*gametemplate.BabyToyUplevelTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.BabyToyUplevelTemplate)(nil))
	for _, temp := range templateMap {
		template, _ := temp.(*gametemplate.BabyToyUplevelTemplate)
		subMap, ok := ts.toyUplevelMap[template.GetSuitType()]
		if !ok {
			subMap = make(map[inventorytypes.BodyPositionType][]*gametemplate.BabyToyUplevelTemplate)
			ts.toyUplevelMap[template.GetSuitType()] = subMap
		}
		subMap[template.GetPosType()] = append(subMap[template.GetPosType()], template)
	}

	return
}

func (ts *babyTemplaterService) loadBaoBaoDongFang() (err error) {
	ts.dongfangMap = make(map[int32]*gametemplate.BabyDongFangTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.BabyDongFangTemplate)(nil))
	for _, temp := range templateMap {
		template, _ := temp.(*gametemplate.BabyDongFangTemplate)
		ts.dongfangMap[template.BabyCount] = template
	}

	return
}

func (ts *babyTemplaterService) loadBaoBaoDanBei() (err error) {
	ts.danbeiMap = make(map[int32]*gametemplate.BabyBeiShuTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.BabyBeiShuTemplate)(nil))
	for _, temp := range templateMap {
		template, _ := temp.(*gametemplate.BabyBeiShuTemplate)
		ts.danbeiMap[int32(template.Id)] = template
	}

	return
}

func (ts *babyTemplaterService) loadBaoBaoQuality() (err error) {
	ts.qualityTempMap = make(map[int32]*gametemplate.BabyQualityTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.BabyQualityTemplate)(nil))
	for _, temp := range templateMap {
		template, _ := temp.(*gametemplate.BabyQualityTemplate)
		ts.qualityTempMap[int32(template.Id)] = template

		if ts.maxTonicNum < template.GetTonicProMax() {
			ts.maxTonicNum = template.GetTonicProMax()
		}

		for _, quality := range template.GetQualityList() {
			_, ok := ts.pregnantMap[quality]
			if !ok {
				return fmt.Errorf("宝宝怀孕配置错误，品质不存在:%d", quality)
			}
		}
	}
	return
}

func (ts *babyTemplaterService) loadBaoBaoSuit() (err error) {
	ts.toySuitGroupMap = make(map[int32]*gametemplate.BabyToySuitGroupTemplate)
	suitGroupTemplateMap := template.GetTemplateService().GetAll((*gametemplate.BabyToySuitGroupTemplate)(nil))
	for _, templateObject := range suitGroupTemplateMap {
		suitGroupTemplate, _ := templateObject.(*gametemplate.BabyToySuitGroupTemplate)
		ts.toySuitGroupMap[int32(suitGroupTemplate.TemplateId())] = suitGroupTemplate
	}
	return
}

func (ts *babyTemplaterService) GetBabyConstantTemplate() *gametemplate.BabyConstantTemplate {
	return ts.babyConstantTemplate
}

func (ts *babyTemplaterService) GetBabyQualityTemplate(tonicNum int32) *gametemplate.BabyQualityTemplate {
	for _, temp := range ts.qualityTempMap {
		if !temp.IsInTonicRange(tonicNum) {
			continue
		}
		return temp
	}

	return nil
}

func (ts *babyTemplaterService) GetBabyToyTemplateBySuitGroup(suitGroupId int32) *gametemplate.BabyToySuitGroupTemplate {
	temp, ok := ts.toySuitGroupMap[suitGroupId]
	if !ok {
		return nil
	}

	return temp
}

func (ts *babyTemplaterService) GetBabyPregnantTemplateByQuality(quality int32) *gametemplate.BabyPregnantTemplate {
	temp, ok := ts.pregnantMap[quality]
	if !ok {
		return nil
	}

	return temp
}

func (ts *babyTemplaterService) GetBabyTalentTemplate(skillId int32) *gametemplate.BabyTalentTemplate {
	temp, ok := ts.talentMap[skillId]
	if !ok {
		return nil
	}

	return temp
}

func (ts *babyTemplaterService) GetBabyPregnantTonicMaxPro() int32 {
	return ts.maxTonicNum
}

func (ts *babyTemplaterService) GetBabyLearnTemplate(level int32) *gametemplate.BabyLearnTemplate {
	temp, ok := ts.learnMap[level]
	if !ok {
		return nil
	}

	return temp
}

func (ts *babyTemplaterService) GetBabyUnlockTalentTemplate(times int32) *gametemplate.BabyUnlockTalentTemplate {
	temp, ok := ts.unlockTalentMap[times]
	if !ok {
		return ts.unlockTalentMap[ts.unlockMaxTimes]
	}

	return temp
}

func (ts *babyTemplaterService) GetBabyToyUplevelTemplate(suitType babytypes.ToySuitType, posType inventorytypes.BodyPositionType, level int32) *gametemplate.BabyToyUplevelTemplate {
	subMap, ok := ts.toyUplevelMap[suitType]
	if !ok {
		return nil
	}

	tempList, ok := subMap[posType]
	if !ok {
		return nil
	}

	for _, temp := range tempList {
		if temp.Level != level {
			continue
		}
		return temp
	}

	return nil
}

func (ts *babyTemplaterService) GetBabyDongFangTemplate(babyNum int32) *gametemplate.BabyDongFangTemplate {
	temp, ok := ts.dongfangMap[babyNum]
	if !ok {
		return nil
	}

	return temp
}

func (ts *babyTemplaterService) GetBabyDanBeiTemplateMap() map[int32]*gametemplate.BabyBeiShuTemplate {
	return ts.danbeiMap
}

func (ts *babyTemplaterService) GetBabyTypeTemplate() *gametemplate.BabyPregnantTemplate {
	var weight []int64
	for _, temp := range ts.babyTypeTemplateList {
		weight = append(weight, temp.Rate)
	}

	index := mathutils.RandomWeights(weight)

	return ts.babyTypeTemplateList[index]
}

var (
	once sync.Once
	cs   *babyTemplaterService
)

func Init() (err error) {
	once.Do(func() {
		cs = &babyTemplaterService{}
		err = cs.init()
	})
	return err
}

func GetBabyTemplateService() BabyTemplaterService {
	return cs
}
