package template

import (
	"fgame/fgame/core/template"
	suipiantypes "fgame/fgame/game/yinglingpu/types"
	"fmt"
)

type YinglingpuTemplate struct {
	*YinglingpuTemplateVO
	suiPianMap map[int32]*YinglingPuSuiPianTemplate //英灵谱碎片列表
	// levelMap             map[int32]*YinglingpuLevelTemplate   //英灵普升级信息
	// startLevelUp         *YinglingpuLevelTemplate             //初始升级图鉴信息
	startSuiPianTemplate *YinglingPuSuiPianTemplate
	levelSuan            *YinglingpuLevelSuanTemplate
}

func (t *YinglingpuTemplate) GetSuiPianMap() map[int32]*YinglingPuSuiPianTemplate {
	return t.suiPianMap
}

// func (t *YinglingpuTemplate) GetLevelMap() map[int32]*YinglingpuLevelTemplate {
// 	return t.levelMap
// }

func (t *YinglingpuTemplate) GetLevelSuan() *YinglingpuLevelSuanTemplate {
	return t.levelSuan
}

func (t *YinglingpuTemplate) TemplateId() int {
	return t.Id
}

func (t *YinglingpuTemplate) FileName() string {
	return "tb_yinglingpu.json"
}

//组合成需要的数据
func (t *YinglingpuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	tempLevelSuanTemplate := template.GetTemplateService().Get(int(t.LevelBeginId), (*YinglingpuLevelSuanTemplate)(nil))
	if tempLevelSuanTemplate == nil {
		err = fmt.Errorf("[%d]无效", t.LevelBeginId)
		err = template.NewTemplateFieldError("LevelBeginId", err)
		return err
	}
	t.levelSuan = tempLevelSuanTemplate.(*YinglingpuLevelSuanTemplate)

	// if t.LevelBeginId != 0 {
	// 	tempLevelTemplate := template.GetTemplateService().Get(int(t.LevelBeginId), (*YinglingpuLevelTemplate)(nil))
	// 	if tempLevelTemplate == nil {
	// 		err = fmt.Errorf("[%d]无效", t.LevelBeginId)
	// 		err = template.NewTemplateFieldError("LevelBeginId", err)
	// 		return err
	// 	}
	// 	t.startLevelUp = tempLevelTemplate.(*YinglingpuLevelTemplate)
	// }
	if t.SuiPianBeginId != 0 {
		tempSuiPian := template.GetTemplateService().Get(int(t.SuiPianBeginId), (*YinglingPuSuiPianTemplate)(nil))
		if tempSuiPian == nil {
			err = fmt.Errorf("[%d]无效", t.SuiPianBeginId)
			err = template.NewTemplateFieldError("SuiPianBeginId", err)
			return err
		}
		t.startSuiPianTemplate = tempSuiPian.(*YinglingPuSuiPianTemplate)
	}
	return nil
}

//检查有效性
func (t *YinglingpuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	tujianType := suipiantypes.YingLingPuTuJianType(t.Type)
	if !tujianType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	return nil
}

//检验后组合
func (t *YinglingpuTemplate) PatchAfterCheck() {
	// t.levelMap = make(map[int32]*YinglingpuLevelTemplate)

	// for levelTemplate := t.startLevelUp; levelTemplate != nil; levelTemplate = levelTemplate.GetNextLevelTemplate() {
	// 	t.levelMap[levelTemplate.Level] = levelTemplate
	// }

	t.suiPianMap = make(map[int32]*YinglingPuSuiPianTemplate)
	for suiPianTemplate := t.startSuiPianTemplate; suiPianTemplate != nil; suiPianTemplate = suiPianTemplate.nextSuiPianTemplate {
		t.suiPianMap[suiPianTemplate.SuipianId] = suiPianTemplate
	}
}

func (t *YinglingpuTemplate) GetConsumeMap(level int32) map[int32]int32 {

	consumeMap := make(map[int32]int32)
	num := t.levelSuan.GetConsumeNum(level)
	if num <= 0 {
		return consumeMap
	}
	for _, suiPian := range t.suiPianMap {
		consumeMap[suiPian.UseItemId] = num
	}
	return consumeMap
}

func init() {
	template.Register((*YinglingpuTemplate)(nil))
}
