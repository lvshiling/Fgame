package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	skilltypes "fgame/fgame/game/skill/types"
	xinfatypes "fgame/fgame/game/xinfa/types"
	"fmt"
)

//心法配置
type XinFaTemplate struct {
	*XinFaTemplateVO
	xinfaTyp        xinfatypes.XinFaType //心法类型
	useItemTemplate *ItemTemplate        //是否消耗物品
}

func (xft *XinFaTemplate) TemplateId() int {
	return xft.Id
}

func (xft *XinFaTemplate) GetType() xinfatypes.XinFaType {
	return xft.xinfaTyp
}

func (xft *XinFaTemplate) GetUseItemTemplate() *ItemTemplate {
	return xft.useItemTemplate
}

func (xft *XinFaTemplate) PatchAfterCheck() {
}

func (xft *XinFaTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(xft.FileName(), xft.TemplateId(), err)
			return
		}
	}()

	//验证 type
	xft.xinfaTyp = xinfatypes.XinFaType(xft.Type)
	if !xft.xinfaTyp.Valid() {
		err = fmt.Errorf("[%d] invalid", xft.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//验证 need_item_id
	if xft.NeedItemId != 0 {
		to := template.GetTemplateService().Get(int(xft.NeedItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", xft.NeedItemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return
		}

		//验证 need_item_num
		err = validator.MinValidate(float64(xft.NeedItemNum), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", xft.NeedItemNum)
			return template.NewTemplateFieldError("NeedItemNum", err)
		}
		xft.useItemTemplate = to.(*ItemTemplate)
	}

	return nil
}

func (xft *XinFaTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(xft.FileName(), xft.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if xft.NextId != 0 {
		to := template.GetTemplateService().Get(int(xft.NextId), (*XinFaTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", xft.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}

		nextTo := to.(*XinFaTemplate)

		//验证type
		if xft.Type != nextTo.Type {
			err = fmt.Errorf("[%d] invalid", nextTo.Type)
			err = template.NewTemplateFieldError("Type", err)
			return
		}

		//验证level
		diffLevel := nextTo.Level - xft.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", nextTo.Level)
			err = template.NewTemplateFieldError("Level", err)
			return
		}
	}

	//验证 need_yinliang
	err = validator.MinValidate(float64(xft.NeedYinLiang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", xft.NeedYinLiang)
		return template.NewTemplateFieldError("NeedYinLiang", err)
	}

	//验证 skill_id
	to := template.GetTemplateService().Get(int(xft.SkillId), (*SkillTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", xft.SkillId)
		err = template.NewTemplateFieldError("SkillId", err)
		return
	}
	typ := to.(*SkillTemplate).GetSkillFirstType()
	if typ != skilltypes.SkillFirstTypeXinFa {
		err = fmt.Errorf("[%d] invalid", xft.SkillId)
		err = template.NewTemplateFieldError("SkillId", err)
		return
	}

	return nil
}

func (xft *XinFaTemplate) FileName() string {
	return "tb_xinfa.json"
}

func init() {
	template.Register((*XinFaTemplate)(nil))
}
