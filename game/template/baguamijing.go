package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	propertytypes "fgame/fgame/game/property/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//八卦秘境配置
type BaGuaMiJingTemplate struct {
	*BaGuaMiJingTemplateVO
	rewItemMap  map[int32]int32        //奖励物品
	rewData     *propertytypes.RewData //奖励属性
	mapTemplate *MapTemplate           //副本地图
}

func (t *BaGuaMiJingTemplate) TemplateId() int {
	return t.Id
}

func (t *BaGuaMiJingTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *BaGuaMiJingTemplate) GetRewData() *propertytypes.RewData {
	return t.rewData
}

func (t *BaGuaMiJingTemplate) GetMapTemplate() *MapTemplate {
	return t.mapTemplate
}

func (t *BaGuaMiJingTemplate) GetRewQinMiDu() int32 {
	return t.RewQinMiDu
}

func (t *BaGuaMiJingTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 map_id
	tempMapTemplate := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if tempMapTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	t.mapTemplate = tempMapTemplate.(*MapTemplate)

	itemArr, err := coreutils.SplitAsIntArray(t.RewItem)
	if err != nil {
		return template.NewTemplateFieldError("RewItem", err)
	}
	numArr, err := coreutils.SplitAsIntArray(t.RewCount)
	if err != nil {
		return template.NewTemplateFieldError("RewCount", err)
	}
	if len(itemArr) != len(numArr) {
		err = fmt.Errorf("RewItem[%s]RewCount[%s]长度不相等", t.RewItem, t.RewCount)
		return template.NewTemplateFieldError("RewItemAmount", err)
	}

	t.rewItemMap = make(map[int32]int32)
	for i := 0; i < len(itemArr); i++ {
		t.rewItemMap[itemArr[i]] = numArr[i]
	}

	//验证 rew_exp
	err = validator.MinValidate(float64(t.RewExp), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewExp", err)
		return
	}

	//验证 rew_exp_point
	err = validator.MinValidate(float64(t.RewExpPoint), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewExpPoint", err)
		return
	}

	//验证 rew_silver
	err = validator.MinValidate(float64(t.RewSilver), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewSilver", err)
		return
	}

	//验证 rew_gold
	err = validator.MinValidate(float64(t.RewGold), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewGold", err)
		return
	}

	//验证 rew_bind_gold
	err = validator.MinValidate(float64(t.RewBindGold), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewBindGold", err)
		return
	}

	if t.RewExp > 0 || t.RewExpPoint > 0 || t.RewSilver > 0 || t.RewGold > 0 || t.RewBindGold > 0 {
		t.rewData = propertytypes.CreateRewData(t.RewExp, t.RewExpPoint, t.RewSilver, t.RewGold, t.RewBindGold)
	}

	return nil
}

func (t *BaGuaMiJingTemplate) PatchAfterCheck() {

}

func (t *BaGuaMiJingTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	for itemId, itemNum := range t.rewItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", itemId)
			return template.NewTemplateFieldError("RewItem", err)
		}

		//验证 rew_count
		err = validator.MinValidate(float64(itemNum), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("RewCount", err)
			return
		}
	}

	//验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*BaGuaMiJingTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		bTemplate := to.(*BaGuaMiJingTemplate)
		//验证level
		diffLevel := bTemplate.Level - t.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", bTemplate.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证 level
	err = validator.MinValidate(float64(t.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//map_id 类型
	mapType := t.mapTemplate.GetMapType()
	if mapType != scenetypes.SceneTypeBaGuaMiJing {
		err = fmt.Errorf("[%d]  invalid", t.MapId)
		err = template.NewTemplateFieldError("MapId", err)
		return
	}

	//验证 rew_qinmidu
	err = validator.MinValidate(float64(t.RewQinMiDu), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewQinMiDu)
		return template.NewTemplateFieldError("RewQinMiDu", err)
	}

	return nil
}

func (t *BaGuaMiJingTemplate) FileName() string {
	return "tb_baguamijing.json"
}

func init() {
	template.Register((*BaGuaMiJingTemplate)(nil))
}
