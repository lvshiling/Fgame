package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

type randomGroup struct {
	min int32 //最小区间
	max int32 //最大区间
}

//元宝拉霸概率配置
type GoldLaBaTemplate struct {
	*GoldLaBaTemplateVo
	weights     []int64
	randomGroup []randomGroup
}

func (t *GoldLaBaTemplate) TemplateId() int {
	return t.Id
}

func (t *GoldLaBaTemplate) RandomGold() int64 {
	// 随机奖励组
	index := mathutils.RandomWeights(t.weights)
	if index == -1 {
		return 0
	}

	// 随机元宝
	group := t.randomGroup[index]
	min := int(group.min)
	max := int(group.max)
	rewGold := int64(mathutils.RandomRange(min, max+1))
	return rewGold
}

func (t *GoldLaBaTemplate) RandomRuleGold() int64 {
	lastGroupIndex := int32(len(t.weights) - 1)
	group := t.randomGroup[lastGroupIndex]
	min := int(group.min)
	max := int(group.max)
	rewGold := int64(mathutils.RandomRange(min, max+1))
	return rewGold
}

func (t *GoldLaBaTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.weights = append(t.weights, int64(t.Percent1))
	t.weights = append(t.weights, int64(t.Percent2))

	group1 := randomGroup{
		min: t.ReturnMin1,
		max: t.ReturnMax1,
	}
	t.randomGroup = append(t.randomGroup, group1)

	group2 := randomGroup{
		min: t.ReturnMin2,
		max: t.ReturnMax2,
	}
	t.randomGroup = append(t.randomGroup, group2)

	if len(t.weights) != len(t.randomGroup) {
		err = fmt.Errorf("weights invalid")
		err = template.NewTemplateFieldError("随机池长度不一致", err)
		return
	}

	return nil
}

func (t *GoldLaBaTemplate) PatchAfterCheck() {
}

func (t *GoldLaBaTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//拉霸次数
	err = validator.MinValidate(float64(t.Times), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Times)
		return template.NewTemplateFieldError("Times", err)
	}

	//活动id
	err = validator.MinValidate(float64(t.GroupId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GroupId)
		return template.NewTemplateFieldError("GroupId", err)
	}

	// nextId
	nextTempObj := template.GetTemplateService().Get(int(t.NextId), (*GoldLaBaTemplate)(nil))
	if nextTempObj != nil {
		nextTemp := nextTempObj.(*GoldLaBaTemplate)
		diff := nextTemp.Times - t.Times
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.Times)
			return template.NewTemplateFieldError("Times", err)
		}
	}

	//充值
	err = validator.MinValidate(float64(t.InvestmentRecharge), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.InvestmentRecharge)
		return template.NewTemplateFieldError("InvestmentRecharge", err)
	}
	//所需元宝
	err = validator.MinValidate(float64(t.Investment), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Investment)
		return template.NewTemplateFieldError("Investment", err)
	}
	//返还下限
	err = validator.MinValidate(float64(t.ReturnMin1), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ReturnMin1)
		return template.NewTemplateFieldError("ReturnMin1", err)
	}
	//返还上限
	err = validator.MinValidate(float64(t.ReturnMax1), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ReturnMax1)
		return template.NewTemplateFieldError("ReturnMax1", err)
	}
	//概率
	err = validator.MinValidate(float64(t.Percent1), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Percent1)
		return template.NewTemplateFieldError("Percent1", err)
	}
	//返还下限
	err = validator.MinValidate(float64(t.ReturnMin2), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ReturnMin2)
		return template.NewTemplateFieldError("ReturnMin2", err)
	}
	//返还下限
	err = validator.MinValidate(float64(t.ReturnMax2), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ReturnMax2)
		return template.NewTemplateFieldError("ReturnMax2", err)
	}
	//概率
	err = validator.MinValidate(float64(t.Percent2), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Percent2)
		return template.NewTemplateFieldError("Percent2", err)
	}

	return nil
}

func (edt *GoldLaBaTemplate) FileName() string {
	return "tb_yuanbaolaba.json"
}

func init() {
	template.Register((*GoldLaBaTemplate)(nil))
}
