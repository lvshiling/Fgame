package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	additionsystypes "fgame/fgame/game/additionsys/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fmt"
)

//功能开启
type ModuleOpenedTemplate struct {
	*ModuleOpenedTemplateVO
	funcOpenType      funcopentypes.FuncOpenType
	funcOpenCheckType funcopentypes.FuncOpenCheckType
	parentFuncOpens   []funcopentypes.FuncOpenType
	needSysType       additionsystypes.AdditionSysType
	openedItems       map[int32]int32
	rewItems          map[int32]int32
	mailRewItems      map[int32]int32
}

func (t *ModuleOpenedTemplate) GetFuncOpenType() funcopentypes.FuncOpenType {
	return t.funcOpenType
}

func (t *ModuleOpenedTemplate) GetFuncOpenCheckType() funcopentypes.FuncOpenCheckType {
	return t.funcOpenCheckType
}

func (t *ModuleOpenedTemplate) GetParentFuncOpens() []funcopentypes.FuncOpenType {
	return t.parentFuncOpens
}

func (t *ModuleOpenedTemplate) GetNeedSysType() additionsystypes.AdditionSysType {
	return t.needSysType
}

func (t *ModuleOpenedTemplate) GetOpenedItems() map[int32]int32 {
	return t.openedItems
}

func (t *ModuleOpenedTemplate) GetRewItems() map[int32]int32 {
	return t.rewItems
}

func (t *ModuleOpenedTemplate) GetMailRewItems() map[int32]int32 {
	return t.mailRewItems
}

func (t *ModuleOpenedTemplate) TemplateId() int {
	return t.Id
}

func (t *ModuleOpenedTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.funcOpenType = funcopentypes.FuncOpenType(t.FuncId)
	t.funcOpenCheckType = funcopentypes.FuncOpenCheckType(t.JianHaoIsAuto)
	t.openedItems = make(map[int32]int32)
	if t.OpenedItemId != 0 {
		t.openedItems[t.OpenedItemId] = t.OpenedItemCount
	}

	parentIds, err := coreutils.SplitAsIntArray(t.ParentId)
	if err != nil {
		return template.NewTemplateFieldError("parentId", err)
	}
	for _, parentId := range parentIds {
		t.parentFuncOpens = append(t.parentFuncOpens, funcopentypes.FuncOpenType(parentId))
	}

	// 开启奖励
	t.rewItems = make(map[int32]int32)
	intRewItemIdArr, err := coreutils.SplitAsIntArray(t.RewItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RewItem)
		return template.NewTemplateFieldError("RewItem", err)
	}
	intRewItemCountArr, err := coreutils.SplitAsIntArray(t.RewItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RewItemCount)
		return template.NewTemplateFieldError("RewItemCount", err)
	}
	if len(intRewItemIdArr) != len(intRewItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.RewItem, t.RewItemCount)
		return template.NewTemplateFieldError("RewItem or RewItemCount", err)
	}
	if len(intRewItemIdArr) > 0 {
		//组合数据
		for index, itemId := range intRewItemIdArr {
			t.rewItems[itemId] = intRewItemCountArr[index]
		}
	}

	// 开启邮件奖励
	t.mailRewItems = make(map[int32]int32)
	intMailRewItemIdArr, err := coreutils.SplitAsIntArray(t.MailRewItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.MailRewItem)
		return template.NewTemplateFieldError("MailRewItem", err)
	}
	intMailRewItemCountArr, err := coreutils.SplitAsIntArray(t.MailRewItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.MailRewItemCount)
		return template.NewTemplateFieldError("MailRewItemCount", err)
	}
	if len(intMailRewItemIdArr) != len(intMailRewItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.MailRewItem, t.MailRewItemCount)
		return template.NewTemplateFieldError("MailRewItem or MailRewItemCount", err)
	}
	if len(intMailRewItemIdArr) > 0 {
		//组合数据
		for index, itemId := range intMailRewItemIdArr {
			t.mailRewItems[itemId] = intMailRewItemCountArr[index]
		}
	}

	return nil
}

func (t *ModuleOpenedTemplate) PatchAfterCheck() {

}

func (t *ModuleOpenedTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if err = validator.MinValidate(float64(t.OpenedLevel), float64(1), true); err != nil {
		return template.NewTemplateFieldError("openedLevel", err)
	}

	if err = validator.MinValidate(float64(t.OpenedZhuanShu), float64(0), true); err != nil {
		return template.NewTemplateFieldError("OpenedZhuanShu", err)
	}

	if !t.funcOpenType.Valid() {
		err = fmt.Errorf("[%d] invalid", int32(t.funcOpenType))
		return template.NewTemplateFieldError("func_id", err)
	}

	if !t.funcOpenCheckType.Valid() {
		err = fmt.Errorf("[%d] invalid", int32(t.funcOpenCheckType))
		return template.NewTemplateFieldError("JianHaoIsAuto", err)
	}

	for _, parentFuncOpen := range t.parentFuncOpens {
		if !parentFuncOpen.Valid() {
			err = fmt.Errorf("[%s] invalid", t.ParentId)
			return template.NewTemplateFieldError("parentId", err)
		}
	}

	for itemId, itemNum := range t.openedItems {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.OpenedItemId)
			return template.NewTemplateFieldError("openItemId", err)
		}
		//验证Time
		err = validator.MinValidate(float64(itemNum), float64(0), false)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", itemNum)
			return template.NewTemplateFieldError("openItemId", err)
		}
	}

	//开启奖励物品
	for itemId, num := range t.rewItems {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("RewItem", fmt.Errorf("[%s] invalid", t.RewItem))
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("RewItemCount", err)
		}
	}

	//开启邮件奖励
	for itemId, num := range t.mailRewItems {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("MailRewItem", fmt.Errorf("[%s] invalid", t.MailRewItem))
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("MailRewItemCount", err)
		}
	}

	if err = validator.MinValidate(float64(t.JianHaoDay), float64(0), true); err != nil {
		return template.NewTemplateFieldError("JianHaoDay", err)
	}

	//验证NeedSysType
	t.needSysType = additionsystypes.AdditionSysType(t.NeedSysType)
	if t.NeedSysType != 0 {
		if !t.needSysType.Valid() {
			err = fmt.Errorf("[%d] invalid", t.NeedSysType)
			return template.NewTemplateFieldError("NeedSysType", err)
		}
		err = validator.MinValidate(float64(t.NeedSysNum), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.NeedSysNum)
			return template.NewTemplateFieldError("NeedSysNum", err)
		}
	}

	return nil
}

func (t *ModuleOpenedTemplate) FileName() string {
	return "tb_module_opened.json"
}

func init() {
	template.Register((*ModuleOpenedTemplate)(nil))
}
