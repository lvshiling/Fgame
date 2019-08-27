package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	babytypes "fgame/fgame/game/baby/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

//宝宝怀孕配置
type BabyPregnantTemplate struct {
	*BabyPregnantTemplateVO
	failReturnItemMap map[int32]int32
	tonicRange        randomGroup            //补品区间
	beiShuRange       randomGroup            //生命区间
	talentList        []babytypes.TalentInfo // 天赋池
	weights           []int64                //天赋权重
	startTalentTemp   *BabyTalentTemplate    //起始天赋配置
}

func (t *BabyPregnantTemplate) TemplateId() int {
	return t.Id
}

func (t *BabyPregnantTemplate) GetTonicProMax() int32 {
	return t.tonicRange.max
}

func (t *BabyPregnantTemplate) GetInitRandonmTalentList() (talentList []*babytypes.TalentInfo) {
	for num := int32(1); num <= t.TalentCount; num += 1 {
		// 随机
		index := mathutils.RandomWeights(t.weights)
		if index == -1 {
			return
		}

		talent := babytypes.NewTalentInfo(t.talentList[index].SkillId, t.talentList[index].Status, t.talentList[index].Type)
		talentList = append(talentList, talent)
	}

	return
}

func (t *BabyPregnantTemplate) GetRandonmTalent() *babytypes.TalentInfo {
	// 随机
	index := mathutils.RandomWeights(t.weights)
	if index == -1 {
		return nil
	}

	return babytypes.NewTalentInfo(t.talentList[index].SkillId, t.talentList[index].Status, t.talentList[index].Type)
}

func (t *BabyPregnantTemplate) GetRandonmBeiShu() int32 {
	min := int(t.beiShuRange.min)
	max := int(t.beiShuRange.max + 1)
	randomNum := mathutils.RandomRange(min, max)
	return int32(randomNum)
}

func (t *BabyPregnantTemplate) GetBabyName(sex playertypes.SexType) string {
	if sex == playertypes.SexTypeMan {
		return t.NameNan
	} else {
		return t.NameNv
	}
}

func (t *BabyPregnantTemplate) PatchAfterCheck() {
	// 加载池所有
	for initTemp := t.startTalentTemp; initTemp != nil; initTemp = initTemp.GetNextTemplate() {
		t.weights = append(t.weights, int64(initTemp.Rate))

		talentInfo := babytypes.TalentInfo{
			SkillId: initTemp.SkillId,
			Status:  babytypes.SkillStatusTypeUnLock,
			Type:    initTemp.GetSkillType(),
		}
		t.talentList = append(t.talentList, talentInfo)
	}

}

func (t *BabyPregnantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 补品值区间
	tonicArr, err := utils.SplitAsIntArray(t.BupinQujian)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.BupinQujian)
		return template.NewTemplateFieldError("BupinQujian", err)
	}
	if len(tonicArr) != 2 {
		err = fmt.Errorf("[%s] invalid", t.BupinQujian)
		return template.NewTemplateFieldError("BupinQujian", err)
	}

	t.tonicRange = randomGroup{
		min: tonicArr[0],
		max: tonicArr[1],
	}

	// 倍数区间
	beiShuArr, err := utils.SplitAsIntArray(t.DanBeiQujian)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.DanBeiQujian)
		return template.NewTemplateFieldError("DanBeiQujian", err)
	}
	if len(beiShuArr) != 2 {
		err = fmt.Errorf("[%s] invalid", t.DanBeiQujian)
		return template.NewTemplateFieldError("DanBeiQujian", err)
	}

	t.beiShuRange = randomGroup{
		min: beiShuArr[0],
		max: beiShuArr[1],
	}

	return nil
}

func (t *BabyPregnantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.tonicRange.min > t.tonicRange.max {
		err = fmt.Errorf("[%s] invalid", t.BupinQujian)
		return template.NewTemplateFieldError("BupinQujian", err)
	}
	if t.beiShuRange.min > t.beiShuRange.max {
		err = fmt.Errorf("[%s] invalid", t.DanBeiQujian)
		return template.NewTemplateFieldError("DanBeiQujian", err)
	}

	//天赋起始配置
	to := template.GetTemplateService().Get(int(t.TalentBeginId), (*BabyTalentTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%s] invalid", t.TalentBeginId)
		return template.NewTemplateFieldError("TalentBeginId", err)
	}
	t.startTalentTemp = to.(*BabyTalentTemplate)

	//天赋数量
	err = validator.MinValidate(float64(t.TalentCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TalentCount)
		return template.NewTemplateFieldError("TalentCount", err)
	}

	//成长
	err = validator.MinValidate(float64(t.GrowthNum), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GrowthNum)
		return template.NewTemplateFieldError("GrowthNum", err)
	}

	//读书最高等级
	err = validator.MinValidate(float64(t.LevelMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevelMax)
		return template.NewTemplateFieldError("LevelMax", err)
	}

	//分享万分比
	err = validator.MinValidate(float64(t.AttrShareRate), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AttrShareRate)
		return template.NewTemplateFieldError("AttrShareRate", err)
	}

	return nil
}

func (t *BabyPregnantTemplate) FileName() string {
	return "tb_baobao_huaiyun.json"
}

func init() {
	template.Register((*BabyPregnantTemplate)(nil))
}
