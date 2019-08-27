package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	"fmt"
	"sort"
)

//装备宝库配置
type EquipBaoKuTemplate struct {
	*EquipBaoKuTemplateVO
	baoKuType      equipbaokutypes.BaoKuType
	dropByTimesMap map[int32]int32 //按次数必定掉落map
	timesList      []int           //循环掉落
	nextTemplate   *EquipBaoKuTemplate
}

func (t *EquipBaoKuTemplate) TemplateId() int {
	return t.Id
}

func (t *EquipBaoKuTemplate) GetBaoKuType() equipbaokutypes.BaoKuType {
	return t.baoKuType
}

func (t *EquipBaoKuTemplate) GetNextTemplate() *EquipBaoKuTemplate {
	return t.nextTemplate
}

func (t *EquipBaoKuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*EquipBaoKuTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemplate = to.(*EquipBaoKuTemplate)
	}
	//按次数必定掉落map
	t.dropByTimesMap = make(map[int32]int32)
	countArr, err := coreutils.SplitAsIntArray(t.MustAmount1)
	if err != nil {
		return template.NewTemplateFieldError("MustAmount1", err)
	}
	dropIdArr, err := coreutils.SplitAsIntArray(t.MustGet1)
	if err != nil {
		return template.NewTemplateFieldError("MustGet1", err)
	}
	if len(countArr) != len(dropIdArr) {
		return template.NewTemplateFieldError("MustAmount1,MustGet1 len not equality", err)
	}
	for i, count := range countArr {
		if _, ok := t.dropByTimesMap[count]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount1)
			return template.NewTemplateFieldError("MustAmount1", err)
		}
		t.dropByTimesMap[count] = dropIdArr[i]
		t.timesList = append(t.timesList, int(count))
	}

	return
}

func (t *EquipBaoKuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证类型
	typ := equipbaokutypes.BaoKuType(t.Type)
	if !typ.Valid() {
		return template.NewTemplateFieldError("Type", err)
	}
	t.baoKuType = typ

	//最小转数
	err = validator.MinValidate(float64(t.ZhuanshuMin), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("ZhuanshuMin", err)
	}

	//最小等级
	err = validator.MinValidate(float64(t.LevelMin), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("LevelMin", err)
	}

	// 银两消耗
	err = validator.MinValidate(float64(t.SilverUse), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("SilverUse", err)
	}
	// 元宝消耗
	err = validator.MinValidate(float64(t.GoldUse), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("GoldUse", err)
	}
	// 绑元消耗
	err = validator.MinValidate(float64(t.BindGoldUse), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("BindGoldUse", err)
	}

	// 探寻一次获得的积分
	err = validator.MinValidate(float64(t.GiftJiFen), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("GiftJiFen", err)
	}

	// 探寻一次获得的幸运值
	err = validator.MinValidate(float64(t.GiftXingYunZhi), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("GiftXingYunZhi", err)
	}

	// 幸运宝箱需要的幸运值
	err = validator.MinValidate(float64(t.NeedXingYunZhi), float64(0), false)
	if err != nil {
		return template.NewTemplateFieldError("NeedXingYunZhi", err)
	}

	tmpObj := template.GetTemplateService().Get(int(t.ScriptXingYun), (*DropTemplate)(nil))
	if tmpObj == nil {
		return template.NewTemplateFieldError("ScriptXingYun", fmt.Errorf("[%s] invalid", t.ScriptXingYun))
	}

	return nil
}

func (t *EquipBaoKuTemplate) PatchAfterCheck() {
}

func (t *EquipBaoKuTemplate) FileName() string {
	return "tb_equip_baoku.json"
}

func (t *EquipBaoKuTemplate) GetRewDropMap() map[int32]int32 {
	return t.dropByTimesMap
}

func (t *EquipBaoKuTemplate) GetDropTimesDescList() []int {
	newList := make([]int, len(t.timesList))
	copy(newList, t.timesList)
	sort.Sort(sort.Reverse(sort.IntSlice(newList)))
	return newList
}

func init() {
	template.Register((*EquipBaoKuTemplate)(nil))
}
