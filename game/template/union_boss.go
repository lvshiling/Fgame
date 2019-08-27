package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/core/utils"
	constanttypes "fgame/fgame/game/constant/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//仙盟boss配置
type UnionBossTemplate struct {
	*UnionBossTemplateVO
	biologyTemplate  *BiologyTemplate   //生物模板
	mzRewItemMap     map[int32]int32    //盟主奖励物品
	cyRewItemMap     map[int32]int32    //仙盟成员奖励物品
	mapTemplate      *MapTemplate       //地图模板
	nextBossTemplate *UnionBossTemplate //下一个模板
	pos              coretypes.Position //位置
}

func (t *UnionBossTemplate) TemplateId() int {
	return t.Id
}

func (t *UnionBossTemplate) GetPos() coretypes.Position {
	return t.pos
}

func (t *UnionBossTemplate) GetBiologyTemplate() *BiologyTemplate {
	return t.biologyTemplate
}

func (t *UnionBossTemplate) GetMapTemplate() *MapTemplate {
	return t.mapTemplate
}

func (t *UnionBossTemplate) GetMzRewItemMap() map[int32]int32 {
	return t.mzRewItemMap
}

func (t *UnionBossTemplate) GetCyRewItemMap() map[int32]int32 {
	return t.cyRewItemMap
}

func (t *UnionBossTemplate) GetNextTemplate() *UnionBossTemplate {
	return t.nextBossTemplate
}

func (t *UnionBossTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.pos = coretypes.Position{
		X: t.PosX,
		Y: t.PosY,
		Z: t.PosZ,
	}

	t.mzRewItemMap = make(map[int32]int32)
	if t.MzAwardSilver != 0 {
		err = validator.MinValidate(float64(t.MzAwardSilver), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.MzAwardSilver)
			return template.NewTemplateFieldError("RewSilver", err)
		}
		t.mzRewItemMap[constanttypes.SilverItem] += t.MzAwardSilver
	}

	if t.MzAwardGold != 0 {
		err = validator.MinValidate(float64(t.MzAwardGold), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.MzAwardGold)
			return template.NewTemplateFieldError("MzAwardGold", err)
		}
		t.mzRewItemMap[constanttypes.GoldItem] += t.MzAwardGold
	}

	if t.MzAwardBindGold != 0 {
		err = validator.MinValidate(float64(t.MzAwardBindGold), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.MzAwardBindGold)
			return template.NewTemplateFieldError("MzAwardBindGold", err)
		}
		t.mzRewItemMap[constanttypes.BindGoldItem] += t.MzAwardBindGold
	}

	if t.MzAwardItemId != "" {
		rewItemIdArr, err := utils.SplitAsIntArray(t.MzAwardItemId)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.MzAwardItemId)
			return template.NewTemplateFieldError("MzAwardItemId", err)
		}
		rewItemCountArr, err := utils.SplitAsIntArray(t.MzAwardItemCount)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.MzAwardItemCount)
			return template.NewTemplateFieldError("MzAwardItemCount", err)
		}
		if len(rewItemIdArr) == 0 || len(rewItemCountArr) != len(rewItemIdArr) {
			err = fmt.Errorf("[%s] invalid", t.MzAwardItemId)
			return template.NewTemplateFieldError("MzAwardItemId", err)
		}
		for index, itemId := range rewItemIdArr {
			to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if to == nil {
				err = fmt.Errorf("[%s] invalid", t.MzAwardItemId)
				return template.NewTemplateFieldError("MzAwardItemId", err)
			}

			err = validator.MinValidate(float64(rewItemCountArr[index]), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%s] invalid", t.MzAwardItemCount)
				return template.NewTemplateFieldError("MzAwardItemCount", err)
			}
			t.mzRewItemMap[itemId] += rewItemCountArr[index]
		}
	}

	t.cyRewItemMap = make(map[int32]int32)
	if t.CyAwardSilver != 0 {
		err = validator.MinValidate(float64(t.CyAwardSilver), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.CyAwardSilver)
			return template.NewTemplateFieldError("CyAwardSilver", err)
		}
		t.mzRewItemMap[constanttypes.SilverItem] += t.CyAwardSilver
		t.cyRewItemMap[constanttypes.SilverItem] += t.CyAwardSilver
	}

	if t.CyAwardGold != 0 {
		err = validator.MinValidate(float64(t.CyAwardGold), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.CyAwardGold)
			return template.NewTemplateFieldError("CyAwardGold", err)
		}
		t.mzRewItemMap[constanttypes.GoldItem] += t.CyAwardGold
		t.cyRewItemMap[constanttypes.GoldItem] += t.CyAwardGold
	}

	if t.CyAwardBindGold != 0 {
		err = validator.MinValidate(float64(t.CyAwardBindGold), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.CyAwardBindGold)
			return template.NewTemplateFieldError("CyAwardBindGold", err)
		}
		t.mzRewItemMap[constanttypes.BindGoldItem] += t.CyAwardBindGold
		t.cyRewItemMap[constanttypes.BindGoldItem] += t.CyAwardBindGold
	}

	if t.CyAwardItemId != "" {
		rewItemIdArr, err := utils.SplitAsIntArray(t.CyAwardItemId)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.CyAwardItemId)
			return template.NewTemplateFieldError("CyAwardItemId", err)
		}
		rewItemCountArr, err := utils.SplitAsIntArray(t.CyAwardItemCount)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.CyAwardItemCount)
			return template.NewTemplateFieldError("CyAwardItemCount", err)
		}
		if len(rewItemIdArr) == 0 || len(rewItemCountArr) != len(rewItemIdArr) {
			err = fmt.Errorf("[%s] invalid", t.CyAwardItemId)
			return template.NewTemplateFieldError("CyAwardItemId", err)
		}
		for index, itemId := range rewItemIdArr {
			to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if to == nil {
				err = fmt.Errorf("[%s] invalid", t.CyAwardItemId)
				return template.NewTemplateFieldError("CyAwardItemId", err)
			}
			err = validator.MinValidate(float64(rewItemCountArr[index]), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%s] invalid", t.CyAwardItemId)
				return template.NewTemplateFieldError("CyAwardItemId", err)
			}
			t.mzRewItemMap[itemId] += rewItemCountArr[index]
			t.cyRewItemMap[itemId] += rewItemCountArr[index]
		}
	}

	//生物
	biologyTemplate := template.GetTemplateService().Get(int(t.BiologyId), (*BiologyTemplate)(nil))
	if biologyTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.BiologyId)
		return template.NewTemplateFieldError("BiologyId", err)
	}
	t.biologyTemplate = biologyTemplate.(*BiologyTemplate)

	//地图
	mapTemplate := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if mapTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}
	t.mapTemplate = mapTemplate.(*MapTemplate)

	//下一个next_id
	if t.NextId != 0 {
		nextTo := template.GetTemplateService().Get(int(t.NextId), (*UnionBossTemplate)(nil))
		if nextTo == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		t.nextBossTemplate = nextTo.(*UnionBossTemplate)
		if t.nextBossTemplate.Level-t.Level != 1 {
			err = fmt.Errorf("[%d] invalid", t.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}
	return nil
}

func (t *UnionBossTemplate) PatchAfterCheck() {

}

func (t *UnionBossTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	err = validator.MinValidate(float64(t.Experience), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Experience)
		return template.NewTemplateFieldError("Experience", err)
	}

	if t.biologyTemplate != nil {
		if t.biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeAllianceBoss {
			err = fmt.Errorf("[%d] invalid", t.BiologyId)
			return template.NewTemplateFieldError("BiologyId", err)
		}
	}

	if t.mapTemplate != nil {
		if t.mapTemplate.GetMapType() != scenetypes.SceneTypeAllianceBoss {
			err = fmt.Errorf("[%d] invalid", t.MapId)
			return template.NewTemplateFieldError("MapId", err)
		}

		mask := t.mapTemplate.GetMap().IsMask(t.pos.X, t.pos.Z)
		if !mask {
			err = fmt.Errorf("pos[%s] invalid", t.pos.String())
			err = template.NewTemplateFieldError("pos", err)
			return
		}
		y := t.mapTemplate.GetMap().GetHeight(t.pos.X, t.pos.Z)
		t.pos.Y = y
	}

	return nil
}

func (t *UnionBossTemplate) FileName() string {
	return "tb_union_boss.json"
}

func init() {
	template.Register((*UnionBossTemplate)(nil))
}
