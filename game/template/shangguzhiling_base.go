package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	shangguzhilingtypes "fgame/fgame/game/shangguzhiling/types"
	"fmt"
)

type ShangguzhilingBaseTemplate struct {
	*ShangguzhilingBaseTemplateVO
	//类型
	lingshouType shangguzhilingtypes.LingshouType
	//所需灵兽模板
	needLingshouTemp *ShangguzhilingBaseTemplate
	//升级起始Id
	firstLevelTemp *ShangguzhilingLevelTemplate
	levelMap       map[int32]*ShangguzhilingLevelTemplate
	//进阶起始Id
	firstRankTemp *ShangguzhilingJinjieTemplate
	rankMap       map[int32]*ShangguzhilingJinjieTemplate
	//灵兽升级使用的物品列表
	lingshouUseItemIdList []int32
}

// 灵兽升级可使用的物品列表
func (t *ShangguzhilingBaseTemplate) GetLingShouUpLevelUseItemList() []int32 {
	return t.lingshouUseItemIdList
}

func (t *ShangguzhilingBaseTemplate) GetLingShouType() shangguzhilingtypes.LingshouType {
	return t.lingshouType
}

//解锁所需灵兽类型
func (t *ShangguzhilingBaseTemplate) GetNeedLingShouType() shangguzhilingtypes.LingshouType {
	if t.needLingshouTemp == nil {
		return t.lingshouType
	}
	return t.needLingshouTemp.GetLingShouType()
}

func (t *ShangguzhilingBaseTemplate) GetLevelTemp(level int32) *ShangguzhilingLevelTemplate {
	temp, ok := t.levelMap[level]
	if !ok {
		return nil
	}
	return temp
}

func (t *ShangguzhilingBaseTemplate) GetRankTemp(rank int32) *ShangguzhilingJinjieTemplate {
	temp, ok := t.rankMap[rank]
	if !ok {
		return nil
	}
	return temp
}

func (t *ShangguzhilingBaseTemplate) TemplateId() int {
	return t.Id
}

func (t *ShangguzhilingBaseTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShangguzhilingBaseTemplate) PatchAfterCheck() {
	t.levelMap = make(map[int32]*ShangguzhilingLevelTemplate)
	for temp := t.firstLevelTemp; temp != nil; temp = temp.GetNextLevelTemp() {
		t.levelMap[temp.Level] = temp
	}

	t.rankMap = make(map[int32]*ShangguzhilingJinjieTemplate)
	for temp := t.firstRankTemp; temp != nil; temp = temp.GetNextRankTemp() {
		t.rankMap[temp.Number] = temp
	}
}

func (t *ShangguzhilingBaseTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//类型
	lingshouType := shangguzhilingtypes.LingshouType(t.Type)
	if !lingshouType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}
	t.lingshouType = lingshouType

	//所需的上古之灵ID
	if t.NeedSgzlId != 0 {
		lingshouTempInterface := template.GetTemplateService().Get(int(t.NeedSgzlId), (*ShangguzhilingBaseTemplate)(nil))
		if lingshouTempInterface == nil {
			err = fmt.Errorf("ShangguzhilingBaseTemplate [%d] no exist", t.NeedSgzlId)
			return template.NewTemplateFieldError("NeedSgzlId", err)
		}
		lingshouTemp, ok := lingshouTempInterface.(*ShangguzhilingBaseTemplate)
		if !ok {
			err = fmt.Errorf("ShangguzhilingBaseTemplate assert [%d] no exist", t.NeedSgzlId)
			return template.NewTemplateFieldError("NeedSgzlId", err)
		}
		t.needLingshouTemp = lingshouTemp
	}

	//所需的上古之灵等级
	err = validator.MinValidate(float64(t.NeedSgzlLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedSgzlLevel)
		return template.NewTemplateFieldError("NeedSgzlLevel", err)
	}

	//上古之灵升级起始ID
	lingshouLevelTempInterface := template.GetTemplateService().Get(int(t.SgzlLevelBeginId), (*ShangguzhilingLevelTemplate)(nil))
	if lingshouLevelTempInterface == nil {
		err = fmt.Errorf("ShangguzhilingLevelTemplate [%d] no exist", t.SgzlLevelBeginId)
		return template.NewTemplateFieldError("SgzlLevelBeginId", err)
	}
	lingshouLevelTemp, ok := lingshouLevelTempInterface.(*ShangguzhilingLevelTemplate)
	if !ok {
		err = fmt.Errorf("ShangguzhilingLevelTemplate assert [%d] no exist", t.SgzlLevelBeginId)
		return template.NewTemplateFieldError("SgzlLevelBeginId", err)
	}
	t.firstLevelTemp = lingshouLevelTemp

	//上古之灵进阶起始ID
	lingshouRankTempInterface := template.GetTemplateService().Get(int(t.JinjieBeginId), (*ShangguzhilingJinjieTemplate)(nil))
	if lingshouRankTempInterface == nil {
		err = fmt.Errorf("ShangguzhilingJinjieTemplate [%d] no exist", t.JinjieBeginId)
		return template.NewTemplateFieldError("JinjieBeginId", err)
	}
	lingshouRankTemp, ok := lingshouRankTempInterface.(*ShangguzhilingJinjieTemplate)
	if !ok {
		err = fmt.Errorf("ShangguzhilingJinjieTemplate assert [%d] no exist", t.JinjieBeginId)
		return template.NewTemplateFieldError("JinjieBeginId", err)
	}
	t.firstRankTemp = lingshouRankTemp

	//hp
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//攻击
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//防御
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	//上古之灵升级使用的物品ID
	lingshouUseItemIdList, err := coreutils.SplitAsIntArray(t.SgzlLevelUseItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.SgzlLevelUseItemId)
		return template.NewTemplateFieldError("SgzlLevelUseItemId", err)
	}
	t.lingshouUseItemIdList = lingshouUseItemIdList

	return nil
}

func (t *ShangguzhilingBaseTemplate) FileName() string {
	return "tb_sgzl.json"
}

func init() {
	template.Register((*ShangguzhilingBaseTemplate)(nil))
}
