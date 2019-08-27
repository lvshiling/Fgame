package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

//宝宝品质配置
type BabyQualityTemplate struct {
	*BabyQualityTemplateVO
	tonicRange  randomGroup          //补品区间
	qualityList []int32              //品质
	weights     []int64              //权重
	nextTemp    *BabyQualityTemplate //下一级
}

func (t *BabyQualityTemplate) TemplateId() int {
	return t.Id
}

func (t *BabyQualityTemplate) PatchAfterCheck() {
}

func (t *BabyQualityTemplate) GetTonicProMax() int32 {
	return t.tonicRange.max
}

func (t *BabyQualityTemplate) GetQualityList() []int32 {
	return t.qualityList
}

func (t *BabyQualityTemplate) RandomQuality() int32 {
	// 随机
	index := mathutils.RandomWeights(t.weights)
	if index == -1 {
		return -1
	}

	// 随机品质
	return t.qualityList[index]
}

func (t *BabyQualityTemplate) IsInTonicRange(tonic int32) bool {
	if tonic >= t.tonicRange.min && tonic <= t.tonicRange.max {
		return true
	}

	return false
}

func (t *BabyQualityTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 补品值区间
	tonicArr, err := utils.SplitAsIntArray(t.QuJian)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.QuJian)
		return template.NewTemplateFieldError("QuJian", err)
	}
	if len(tonicArr) != 2 {
		err = fmt.Errorf("[%s] invalid", t.QuJian)
		return template.NewTemplateFieldError("QuJian", err)
	}

	t.tonicRange = randomGroup{
		min: tonicArr[0],
		max: tonicArr[1],
	}
	if t.tonicRange.min > t.tonicRange.max {
		err = fmt.Errorf("[%s] invalid", t.QuJian)
		return template.NewTemplateFieldError("QuJian", err)
	}

	//权重
	t.weights = append(t.weights, int64(t.Rate1))
	t.weights = append(t.weights, int64(t.Rate2))
	t.weights = append(t.weights, int64(t.Rate3))
	t.weights = append(t.weights, int64(t.Rate4))

	//品质
	t.qualityList = append(t.qualityList, t.Type1)
	t.qualityList = append(t.qualityList, t.Type2)
	t.qualityList = append(t.qualityList, t.Type3)
	t.qualityList = append(t.qualityList, t.Type4)

	//下一阶强化
	if t.NextId != 0 {
		if t.NextId-t.Id != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		to := template.GetTemplateService().Get(t.NextId, (*BabyQualityTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*BabyQualityTemplate)

	}

	return nil
}

func (t *BabyQualityTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	// 权重校验
	totalWeight := int64(0)
	for _, weight := range t.weights {
		err = validator.MinValidate(float64(weight), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.weights)
			return template.NewTemplateFieldError("Rate", err)
		}
		totalWeight += weight
	}
	if totalWeight == 0 {
		err = fmt.Errorf("Rate invalid")
		return template.NewTemplateFieldError("Rate", err)
	}

	// 品质
	for _, quality := range t.qualityList {
		err = validator.MinValidate(float64(quality), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.qualityList)
			return template.NewTemplateFieldError("Type", err)
		}
	}

	if t.nextTemp != nil {
		//区间连续校验
		if t.nextTemp.tonicRange.min-t.tonicRange.max != 1 {
			err = fmt.Errorf("[%d] invalid", t.QuJian)
			return template.NewTemplateFieldError("QuJian", err)
		}
	}

	return nil
}

func (t *BabyQualityTemplate) FileName() string {
	return "tb_baobao_quality.json"
}

func init() {
	template.Register((*BabyQualityTemplate)(nil))
}
