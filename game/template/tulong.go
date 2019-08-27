package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	consttypes "fgame/fgame/game/constant/types"
	propertytypes "fgame/fgame/game/property/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/tulong/types"
	"fmt"
)

//屠龙配置
type TuLongTemplate struct {
	*TuLongTemplateVO
	tulongType      types.TuLongBossType   //类型
	biologyTemplate *BiologyTemplate       //生物模板
	rewData         *propertytypes.RewData //奖励属性
	rewItemMap      map[int32]int32        //奖励物品
	offRewItemMap   map[int32]int32        //离线奖励物品
}

func (tl *TuLongTemplate) TemplateId() int {
	return tl.Id
}

func (tl *TuLongTemplate) GetRewItemMap() map[int32]int32 {
	return tl.rewItemMap
}

func (tl *TuLongTemplate) GetOffRewItemMap() map[int32]int32 {
	return tl.offRewItemMap
}

func (tl *TuLongTemplate) GetTuLongType() types.TuLongBossType {
	return tl.tulongType
}

func (tl *TuLongTemplate) GetRewData() *propertytypes.RewData {
	return tl.rewData
}

func (tl *TuLongTemplate) GetBiologyTemplate() *BiologyTemplate {
	return tl.biologyTemplate
}

func (tl *TuLongTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tl.FileName(), tl.TemplateId(), err)
			return
		}
	}()

	//验证类型
	tl.tulongType = types.TuLongBossType(tl.Type)
	if !tl.tulongType.Valid() {
		err = fmt.Errorf("[%d] invalid", tl.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	tl.rewItemMap = make(map[int32]int32)
	if tl.RewItemId != "" {
		rewItemIdArr, err := utils.SplitAsIntArray(tl.RewItemId)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", tl.RewItemId)
			return template.NewTemplateFieldError("RewItemId", err)
		}
		rewItemCountArr, err := utils.SplitAsIntArray(tl.RewCount)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", tl.RewCount)
			return template.NewTemplateFieldError("RewCount", err)
		}
		if len(rewItemIdArr) == 0 || len(rewItemCountArr) != len(rewItemIdArr) {
			err = fmt.Errorf("[%s] invalid", tl.RewItemId)
			return template.NewTemplateFieldError("RewItemId", err)
		}
		for i := 0; i < len(rewItemIdArr); i++ {
			tl.rewItemMap[rewItemIdArr[i]] = rewItemCountArr[i]
		}
	}

	//验证 rew_silver
	err = validator.MinValidate(float64(tl.RewSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tl.RewSilver)
		return template.NewTemplateFieldError("RewSilver", err)
	}

	//验证 rew_exp
	err = validator.MinValidate(float64(tl.RewExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tl.RewExp)
		return template.NewTemplateFieldError("RewExp", err)
	}

	//验证 rew_gold
	err = validator.MinValidate(float64(tl.RewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tl.RewGold)
		return template.NewTemplateFieldError("RewGold", err)
	}

	//验证 rew_bindgold
	err = validator.MinValidate(float64(tl.RewBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tl.RewBindGold)
		return template.NewTemplateFieldError("RewBindGold", err)
	}

	if tl.RewSilver > 0 || tl.RewGold > 0 || tl.RewExp > 0 || tl.RewBindGold > 0 {
		tl.rewData = propertytypes.CreateRewData(tl.RewExp, 0, tl.RewSilver, tl.RewGold, tl.RewBindGold)
	}

	tl.offRewItemMap = make(map[int32]int32)
	for itemId, num := range tl.rewItemMap {
		tl.offRewItemMap[itemId] = num
	}
	if tl.RewSilver > 0 {
		num := tl.offRewItemMap[consttypes.SilverItem]
		num += tl.RewSilver
		tl.offRewItemMap[consttypes.SilverItem] = num
	}
	if tl.RewGold > 0 {
		num := tl.offRewItemMap[consttypes.GoldItem]
		num += tl.RewGold
		tl.offRewItemMap[consttypes.GoldItem] = num
	}
	if tl.RewExp > 0 {
		num := tl.offRewItemMap[consttypes.ExpItem]
		num += tl.RewExp
		tl.offRewItemMap[consttypes.ExpItem] = num
	}
	if tl.RewBindGold > 0 {
		num := tl.offRewItemMap[consttypes.BindGoldItem]
		num += tl.RewBindGold
		tl.offRewItemMap[consttypes.BindGoldItem] = num
	}

	return nil
}

func (tl *TuLongTemplate) PatchAfterCheck() {

}

func (tl *TuLongTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tl.FileName(), tl.TemplateId(), err)
			return
		}
	}()

	for itemId, num := range tl.rewItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", tl.RewItemId)
			err = template.NewTemplateFieldError("RewItemId", err)
			return
		}
		if num <= 0 {
			err = fmt.Errorf("[%s] invalid", tl.RewCount)
			err = template.NewTemplateFieldError("RewCount", err)
			return
		}
	}

	//验证 RefreshTime
	err = validator.MinValidate(float64(tl.BiaoShi), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tl.BiaoShi)
		return template.NewTemplateFieldError("BiaoShi", err)
	}

	to := template.GetTemplateService().Get(int(tl.BiologyId), (*BiologyTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", tl.BiologyId)
		return template.NewTemplateFieldError("BiologyId", err)
	}
	tl.biologyTemplate = to.(*BiologyTemplate)
	switch tl.tulongType {
	case types.TuLongBossTypeBig:
		{
			if tl.biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeCrossBigBoss {
				err = fmt.Errorf("[%d] invalid", tl.BiologyId)
				return template.NewTemplateFieldError("BiologyId", err)
			}
		}
	case types.TuLongBossTypeSmall:
		{
			if tl.biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeCrossSmallBoss {
				err = fmt.Errorf("[%d] invalid", tl.BiologyId)
				return template.NewTemplateFieldError("BiologyId", err)
			}
		}
	}

	return nil
}

func (tl *TuLongTemplate) FileName() string {
	return "tb_tulong.json"
}

func init() {
	template.Register((*TuLongTemplate)(nil))
}
