package template

import (
	"fgame/fgame/core/template"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fmt"
)

type ChuangShiChengFangTemplate struct {
	*ChuangShiChengFangTemplateVO
	jianSheType    chuangshitypes.ChuangShiCityJianSheType
	startTemp      *ChuangShiChengFangJianSheTemplate //关联升级起始
	jianSheTempMap map[int32]*ChuangShiChengFangJianSheTemplate
}

func (t *ChuangShiChengFangTemplate) GetJianSheType() chuangshitypes.ChuangShiCityJianSheType {
	return t.jianSheType
}

func (t *ChuangShiChengFangTemplate) GetJianSheLevelTemp(level int32) *ChuangShiChengFangJianSheTemplate {
	return t.jianSheTempMap[level]
}

func (t *ChuangShiChengFangTemplate) TemplateId() int {
	return t.Id
}

func (t *ChuangShiChengFangTemplate) FileName() string {
	return "tb_chuangshi_chengfang.json"
}

func (t *ChuangShiChengFangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 类型
	t.jianSheType = chuangshitypes.ChuangShiCityJianSheType(t.Type)
	if !t.jianSheType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//起始id
	to := template.GetTemplateService().Get(int(t.LevelBeginId), (*ChuangShiChengFangJianSheTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.LevelBeginId)
		return template.NewTemplateFieldError("LevelBeginId", err)
	}
	t.startTemp = to.(*ChuangShiChengFangJianSheTemplate)
	return
}

func (t *ChuangShiChengFangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	to := template.GetTemplateService().Get(int(t.LevelItemId), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.LevelItemId)
		err = template.NewTemplateFieldError("LevelItemId", err)
		return
	}

	return
}

func (t *ChuangShiChengFangTemplate) PatchAfterCheck() {
	t.jianSheTempMap = make(map[int32]*ChuangShiChengFangJianSheTemplate)
	for initTemp := t.startTemp; initTemp != nil; initTemp = initTemp.GetNextTemp() {
		t.jianSheTempMap[initTemp.Level] = initTemp
	}

}

func init() {
	template.Register((*ChuangShiChengFangTemplate)(nil))
}
