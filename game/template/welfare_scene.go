package template

import (
	"fgame/fgame/core/template"
	welfarescenetypes "fgame/fgame/game/welfarescene/types"
	"fmt"
)

func init() {
	template.Register((*WelfareSceneTemplate)(nil))
}

type WelfareSceneTemplate struct {
	*WelfareSceneTemplateVO
	qiYuTemp *WelfareSceneQiYuTemplate //奇遇模板
	wsType   welfarescenetypes.WelfareSceneType
}

func (t *WelfareSceneTemplate) TemplateId() int {
	return t.Id
}

func (t *WelfareSceneTemplate) GetQiYuTemp() *WelfareSceneQiYuTemplate {
	return t.qiYuTemp
}

func (t *WelfareSceneTemplate) GetWelfareSceneType() welfarescenetypes.WelfareSceneType {
	return t.wsType
}

func (t *WelfareSceneTemplate) FileName() string {
	return "tb_yunying_scene.json"
}

//组合成需要的数据
func (t *WelfareSceneTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	t.wsType = welfarescenetypes.WelfareSceneType(t.Type)
	if !t.wsType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	switch t.wsType {
	case welfarescenetypes.WelfareSceneTypeQiYuDao:
		to := template.GetTemplateService().Get(int(t.SubTemplateId), (*WelfareSceneQiYuTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.SubTemplateId)
			return template.NewTemplateFieldError("SubTemplateId", err)
		}
		t.qiYuTemp = to.(*WelfareSceneQiYuTemplate)
	}

	return nil
}

//检查有效性
func (t *WelfareSceneTemplate) Check() (err error) {
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

	return nil
}

//检验后组合
func (t *WelfareSceneTemplate) PatchAfterCheck() {
}
