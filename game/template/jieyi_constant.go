package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

//结义常量配置
type JieYiConstantTemplate struct {
	*JieYiConstantTemplateVO
}

func (t *JieYiConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *JieYiConstantTemplate) Patch() (err error) {
	return
}

func (t *JieYiConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证改名卡
	to := template.GetTemplateService().Get(int(t.ChangeNameItemId), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.ChangeNameItemId)
		return template.NewTemplateFieldError("ChangeNameItemId", err)
	}

	// 验证求援CD
	err = validator.MinValidate(float64(t.QiuYuanCD), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.QiuYuanCD)
		return template.NewTemplateFieldError("QiuYuanCD", err)
	}

	// 验证发布结义CD
	err = validator.MinValidate(float64(t.FaBuCD), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FaBuCD)
		return template.NewTemplateFieldError("FaBuCD", err)
	}

	// 验证结义存在时间
	err = validator.MinValidate(float64(t.FaBuJieYiMaxTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FaBuJieYiMaxTime)
		return template.NewTemplateFieldError("FaBuJieYiMaxTime", err)
	}

	// 验证存储的最大声威值
	err = validator.MinValidate(float64(t.ShengWeiMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ShengWeiMax)
		return template.NewTemplateFieldError("ShengWeiMax", err)
	}

	// 验证邀请玩家结义CD
	err = validator.MinValidate(float64(t.YaoQingCD), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.YaoQingCD)
		return template.NewTemplateFieldError("YaoQingCD", err)
	}

	// 验证邀请有效时间
	err = validator.MinValidate(float64(t.YaoQingExistTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.YaoQingExistTime)
		return template.NewTemplateFieldError("YaoQingExistTime", err)
	}

	// 验证解除结义无法加入CD
	err = validator.MinValidate(float64(t.JieChuCD), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.JieChuCD)
		return template.NewTemplateFieldError("JieChuCD", err)
	}

	// 验证声威掉落概率
	err = validator.RangeValidate(float64(t.DropRate), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropRate)
		return template.NewTemplateFieldError("DropRate", err)
	}

	// 验证声威掉落最小比例
	err = validator.RangeValidate(float64(t.DropPercentMin), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropPercentMin)
		return template.NewTemplateFieldError("DropPercentMin", err)
	}

	// 验证声威掉落最大比例
	err = validator.RangeValidate(float64(t.DropPercentMax), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropPercentMax)
		return template.NewTemplateFieldError("DropPercentMax", err)
	}
	if t.DropPercentMin > t.DropPercentMax {
		err = fmt.Errorf("[%d] and [%d]invalid", t.DropPercentMin, t.DropPercentMax)
		return template.NewTemplateFieldError("DropPercentMin and DropPercentMax", err)
	}

	// 验证声威掉落冷却时间
	err = validator.MinValidate(float64(t.DropCD), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropCD)
		return template.NewTemplateFieldError("DropCD", err)
	}

	// 验证声威掉落保护时间
	err = validator.MinValidate(float64(t.DropProtectedTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropProtectedTime)
		return template.NewTemplateFieldError("DropProtectedTime", err)
	}

	// 验证声威掉落存活时间
	err = validator.MinValidate(float64(t.DropFailTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropFailTime)
		return template.NewTemplateFieldError("DropFailTime", err)
	}

	// 验证背包声威掉落最小堆数
	err = validator.RangeValidate(float64(t.DropMinStack), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropMinStack)
		return template.NewTemplateFieldError("DropMinStack", err)
	}

	// 验证背包声威掉落最大堆数
	err = validator.RangeValidate(float64(t.DropMaxStack), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropMaxStack)
		return template.NewTemplateFieldError("DropMaxStack", err)
	}
	if t.DropMinStack > t.DropMaxStack {
		err = fmt.Errorf("[%d] and [%d]invalid", t.DropMinStack, t.DropMaxStack)
		return template.NewTemplateFieldError("DropMinStack and DropMaxStack", err)
	}

	// 验证声威掉落被系统回收比例
	err = validator.RangeValidate(float64(t.DropSystemReturn), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropSystemReturn)
		return template.NewTemplateFieldError("DropSystemReturn", err)
	}

	// 验证结义人数
	err = validator.MinValidate(float64(t.MaxPeopleNum), 2, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MaxPeopleNum)
		return template.NewTemplateFieldError("MaxPeopleNum", err)
	}

	return nil
}

func (t *JieYiConstantTemplate) PatchAfterCheck() {
}

func (t *JieYiConstantTemplate) FileName() string {
	return "tb_jieyi_constant.json"
}

func init() {
	template.Register((*JieYiConstantTemplate)(nil))
}
