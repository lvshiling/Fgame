package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/core/utils"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

//打宝塔配置
type TowerTemplate struct {
	*TowerTemplateVO
	nextTemp        *TowerTemplate
	bossTemp        *BiologyTemplate
	bossPos         coretypes.Position
	monsterIdList   []int32
	dummyItemIdList []int32
}

func (t *TowerTemplate) TemplateId() int {
	return t.Id
}

func (t *TowerTemplate) GetNextTemplate() *TowerTemplate {
	return t.nextTemp
}

func (t *TowerTemplate) GetBossBornPos() coretypes.Position {
	return t.bossPos
}

func (t *TowerTemplate) GetDummyItemIdList() []int32 {
	return t.dummyItemIdList
}

func (t *TowerTemplate) GetRandomBiologyDropItemId() (biologyId int32, itemId int32) {
	if len(t.monsterIdList) == 0 {
		return
	}
	indexMonster := mathutils.RandomRange(0, len(t.monsterIdList))
	biologyId = t.monsterIdList[indexMonster]

	if len(t.dummyItemIdList) == 0 {
		return
	}

	indexItem := mathutils.RandomRange(0, len(t.dummyItemIdList))
	itemId = t.dummyItemIdList[indexItem]
	return
}

func (t *TowerTemplate) PatchAfterCheck() {
}

func (t *TowerTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//小怪掉落itemId
	rewItemIdList, err := utils.SplitAsIntArray(t.XujiaItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.XujiaItem)
		return template.NewTemplateFieldError("XujiaItem", err)
	}
	for _, itemId := range rewItemIdList {
		itemTemp := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemp == nil {
			err = fmt.Errorf("[%d] invalid", t.XujiaItem)
			err = template.NewTemplateFieldError("XujiaItem", err)
			return
		}

		t.dummyItemIdList = append(t.dummyItemIdList, itemId)
	}

	//小怪id
	monsterIdList, err := utils.SplitAsIntArray(t.XujiaId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.XujiaId)
		return template.NewTemplateFieldError("XujiaId", err)
	}
	for _, monsterId := range monsterIdList {
		bilogyTemp := template.GetTemplateService().Get(int(monsterId), (*BiologyTemplate)(nil))
		if bilogyTemp == nil {
			err = fmt.Errorf("[%d] invalid", t.XujiaId)
			err = template.NewTemplateFieldError("XujiaId", err)
			return
		}

		t.monsterIdList = append(t.monsterIdList, monsterId)
	}

	return nil
}

func (t *TowerTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*TowerTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		t.nextTemp = to.(*TowerTemplate)
	}

	//验证 最低等级
	err = validator.MinValidate(float64(t.LevelMin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevelMin)
		err = template.NewTemplateFieldError("LevelMin", err)
		return
	}
	//验证 最高等级
	err = validator.MinValidate(float64(t.LevelMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevelMax)
		err = template.NewTemplateFieldError("LevelMax", err)
		return
	}
	//验证 地图id
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

	//怪物id
	if t.BossId > 0 {
		bilogyTemp := template.GetTemplateService().Get(int(t.BossId), (*BiologyTemplate)(nil))
		if bilogyTemp == nil {
			err = fmt.Errorf("[%d] invalid", t.BossId)
			return template.NewTemplateFieldError("BossId", err)
		}
		bossTemp := bilogyTemp.(*BiologyTemplate)
		if bossTemp == nil {
			err = fmt.Errorf("[%d] invalid", t.BossId)
			return template.NewTemplateFieldError("BossId", err)
		}
		t.bossTemp = bossTemp
	}

	//验证 礼包物品
	itemTemp := template.GetTemplateService().Get(int(t.ZhifeiItem), (*ItemTemplate)(nil))
	if itemTemp == nil {
		err = fmt.Errorf("[%d] invalid", t.ZhifeiItem)
		err = template.NewTemplateFieldError("ZhifeiItem", err)
		return
	}

	err = validator.MinValidate(float64(t.ZhifeiItemCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhifeiItemCount)
		err = template.NewTemplateFieldError("ZhifeiItemCount", err)
		return
	}

	// boss位置
	t.bossPos = coretypes.Position{
		X: t.BossPosX,
		Y: t.BossPosY,
		Z: t.BossPosZ,
	}

	return nil
}

func (t *TowerTemplate) FileName() string {
	return "tb_dabaota.json"
}

func init() {
	template.Register((*TowerTemplate)(nil))
}
