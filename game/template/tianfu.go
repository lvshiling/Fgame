package template

import (
	"fgame/fgame/core/template"
	"fmt"
)

//天赋配置
type TianFuTemplate struct {
	*TianFuTemplateVO
	tianFuLevelTemplateMap map[int32]*TianFuLevelTemplate //天赋升级map
	tianFuLevelTemplate    *TianFuLevelTemplate           //天赋升级
	nextTianFuTemplate     *TianFuTemplate                //下一个天赋
}

func (t *TianFuTemplate) TemplateId() int {
	return t.Id
}

func (t *TianFuTemplate) GetTianFuLevelByLevel(level int32) *TianFuLevelTemplate {
	if v, ok := t.tianFuLevelTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (t *TianFuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 level_begin
	if t.LevelBegin != 0 {
		to := template.GetTemplateService().Get(int(t.LevelBegin), (*TianFuLevelTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.LevelBegin)
			return template.NewTemplateFieldError("LevelBegin", err)
		}

		tianFuLevelTemplate, ok := to.(*TianFuLevelTemplate)
		if !ok {
			return fmt.Errorf("LevelBegin [%d] invalid", t.LevelBegin)
		}
		if tianFuLevelTemplate.Level != 1 {
			return fmt.Errorf("LevelBegin Level [%d] invalid", tianFuLevelTemplate.Level)
		}
		t.tianFuLevelTemplate = tianFuLevelTemplate
	}

	return nil
}

func (t *TianFuTemplate) PatchAfterCheck() {
	if t.tianFuLevelTemplate != nil {
		t.tianFuLevelTemplateMap = make(map[int32]*TianFuLevelTemplate)
		//赋值tianFuLevelTemplateMap
		for tempTemplate := t.tianFuLevelTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextTianFuLevelTemplate {
			level := tempTemplate.Level
			t.tianFuLevelTemplateMap[level] = tempTemplate
		}
	}
}

func (t *TianFuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.NextId != 0 {
		tempTianFuTemplate := template.GetTemplateService().Get(int(t.NextId), (*TianFuTemplate)(nil))
		if tempTianFuTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTianFuTemplate = tempTianFuTemplate.(*TianFuTemplate)
	}

	if t.ParentId != 0 {
		tempTianFuTemplate := template.GetTemplateService().Get(int(t.ParentId), (*TianFuTemplate)(nil))
		if tempTianFuTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.ParentId)
			return template.NewTemplateFieldError("ParentId", err)
		}
	}

	return nil
}

func (t *TianFuTemplate) FileName() string {
	return "tb_tianfu.json"
}

func init() {
	template.Register((*TianFuTemplate)(nil))
}
