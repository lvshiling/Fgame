package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

func init() {
	template.Register((*ZhenXiBossUseTemplate)(nil))
}

type ZhenXiBossUseTemplate struct {
	*ZhenXiBossUseTemplateVO
	useItemMap map[int32]int32
}

func (t *ZhenXiBossUseTemplate) TemplateId() int {
	return t.Id
}

func (t *ZhenXiBossUseTemplate) FileName() string {
	return "tb_zhenxi_boss_use.json"
}

func (t *ZhenXiBossUseTemplate) GetUseItemMap() map[int32]int32 {
	return t.useItemMap
}

//组合成需要的数据
func (t *ZhenXiBossUseTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.useItemMap = make(map[int32]int32)
	t.useItemMap[t.UseItem] = t.UseCount
	return nil
}

//检查有效性
func (t *ZhenXiBossUseTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//地图
	mto := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if mto == nil {
		err = fmt.Errorf("mapId [%d] no exist", t.MapId)
		return err
	}
	_, ok := mto.(*MapTemplate)
	if !ok {
		err = fmt.Errorf("mapId [%d] no exist", t.MapId)
		return
	}

	for itemId, itemNum := range t.useItemMap {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%d] 无效", t.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}
		err = validator.MinValidate(float64(itemNum), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("UseCount", err)
		}
	}

	return nil
}

//检验后组合
func (t *ZhenXiBossUseTemplate) PatchAfterCheck() {
}
