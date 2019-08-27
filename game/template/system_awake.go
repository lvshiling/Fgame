package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fmt"
)

type SystemAwakeTemplate struct {
	*SystemAwakeTemplateVO
	sysType                 additionsystypes.AdditionSysType
	beginAwakeLevelTemplate *SystemAwakeLevelTemplate           // 觉醒起始等级模板
	awakeLevelTemplateMap   map[int32]*SystemAwakeLevelTemplate // 觉醒等级模板
}

func (t *SystemAwakeTemplate) TemplateId() int {
	return t.Id
}

func (t *SystemAwakeTemplate) GetSysType() additionsystypes.AdditionSysType {
	return t.sysType
}

func (t *SystemAwakeTemplate) GetAwakeLevelTemplate(level int32) *SystemAwakeLevelTemplate {
	temp, ok := t.awakeLevelTemplateMap[level]
	if !ok {
		return nil
	}
	return temp
}

func (t *SystemAwakeTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证SysType
	sysType := additionsystypes.AdditionSysType(t.SysType)
	if !sysType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.SysType)
		return template.NewTemplateFieldError("SysType", err)
	}

	t.sysType = sysType

	return
}

func (t *SystemAwakeTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*SystemAwakeTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		next := to.(*SystemAwakeTemplate)

		if next.SysType != t.SysType {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		diff := next.Number - int32(t.Number)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	//验证number
	err = validator.MinValidate(float64(t.Number), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Number)
		err = template.NewTemplateFieldError("Number", err)
		return
	}

	//验证 hp
	err = validator.MinValidate(float64(t.Hp), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}

	//验证 attack
	err = validator.MinValidate(float64(t.Attack), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}

	//验证 defence
	err = validator.MinValidate(float64(t.Defence), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	// 验证 level_begin_id
	if t.BeginId != 0 {
		awakeLevelTemp := template.GetTemplateService().Get(t.BeginId, (*SystemAwakeLevelTemplate)(nil))
		if awakeLevelTemp == nil {
			err = fmt.Errorf("[%d] invalid", t.BeginId)
			err = template.NewTemplateFieldError("BeginId", err)
			return
		}
		temp, ok := awakeLevelTemp.(*SystemAwakeLevelTemplate)
		if !ok {
			return fmt.Errorf("UpStarBeginId [%d] invalid", t.BeginId)
		}
		t.beginAwakeLevelTemplate = temp
	}

	return
}

func (t *SystemAwakeTemplate) PatchAfterCheck() {
	t.awakeLevelTemplateMap = make(map[int32]*SystemAwakeLevelTemplate)
	if t.BeginId != 0 {
		for temp := t.beginAwakeLevelTemplate; temp != nil; temp = temp.nextTemplate {
			t.awakeLevelTemplateMap[temp.Level] = temp
		}
	}
}

func (t *SystemAwakeTemplate) FileName() string {
	return "tb_system_juexing.json"
}

func init() {
	template.Register((*SystemAwakeTemplate)(nil))
}
