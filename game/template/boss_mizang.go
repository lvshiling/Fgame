package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	collecttypes "fgame/fgame/game/collect/types"
	"fmt"
)

type BossMiZang struct {
	itemMap  map[int32]int32
	dropList []int32
	rate     int32
}

func (z *BossMiZang) GetItemMap() map[int32]int32 {
	return z.itemMap
}

func (z *BossMiZang) GetDropList() []int32 {
	return z.dropList
}

func (z *BossMiZang) GetRate() int32 {
	return z.rate
}

func newBossMiZang(itemMap map[int32]int32, dropList []int32, rate int32) *BossMiZang {
	miZang := &BossMiZang{}
	miZang.itemMap = itemMap
	miZang.dropList = dropList
	miZang.rate = rate
	return miZang
}

//boss密藏
type BossMiZangTemplate struct {
	*BossMizangTemplateVO
	miZangMap       map[collecttypes.MiZangOpenType]*BossMiZang
	biologyTemplate *BiologyTemplate
}

func (t *BossMiZangTemplate) TemplateId() int {
	return t.Id
}

func (t *BossMiZangTemplate) GetMiZang(typ collecttypes.MiZangOpenType) *BossMiZang {
	miZang, ok := t.miZangMap[typ]
	if !ok {
		return nil
	}
	return miZang
}

func (t *BossMiZangTemplate) GetBiologyTemplate() *BiologyTemplate {
	return t.biologyTemplate
}

func (t *BossMiZangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.miZangMap = make(map[collecttypes.MiZangOpenType]*BossMiZang)
	silverItemMap := make(map[int32]int32)
	silverItemMap[t.SilverItemId] = t.SilverItemCount
	silverDropList, err := coreutils.SplitAsIntArray(t.SilverDrop)
	if err != nil {
		err = template.NewTemplateFieldError("SilverDrop", err)
		return
	}
	t.miZangMap[collecttypes.MiZangOpenTypeSilver] = newBossMiZang(silverItemMap, silverDropList, t.SilverRateAdd)
	goldItemMap := make(map[int32]int32)
	goldItemMap[t.GoldItemId] = t.GoldItemCount
	goldDropList, err := coreutils.SplitAsIntArray(t.GoldDrop)
	if err != nil {
		err = template.NewTemplateFieldError("GoldDrop", err)
		return
	}
	t.miZangMap[collecttypes.MiZangOpenTypeGold] = newBossMiZang(goldItemMap, goldDropList, t.GoldRateAdd)
	if t.CaijiBiologyId != 0 {
		tempBiologyTemplate := template.GetTemplateService().Get(int(t.CaijiBiologyId), (*BiologyTemplate)(nil))
		if tempBiologyTemplate != nil {
			t.biologyTemplate = tempBiologyTemplate.(*BiologyTemplate)
		}
	}
	return
}

func (t *BossMiZangTemplate) PatchAfterCheck() {

}
func (t *BossMiZangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	for _, bossMiZang := range t.miZangMap {
		for itemId, itemNum := range bossMiZang.GetItemMap() {
			tempItemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if tempItemTemplate == nil {
				err = fmt.Errorf("[%d] invalid", itemId)
				return template.NewTemplateFieldError("ItemId", err)
			}
			err = validator.MinValidate(float64(itemNum), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", itemNum)
				return template.NewTemplateFieldError("ItemNum", err)
			}
		}
		err = validator.MinValidate(float64(bossMiZang.GetRate()), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", bossMiZang.GetRate())
			return template.NewTemplateFieldError("Rate", err)
		}
	}
	if t.biologyTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.CaijiBiologyId)
		return template.NewTemplateFieldError("CaijiBiologyId", err)
	}
	return nil
}

func (t *BossMiZangTemplate) FileName() string {
	return "tb_boss_mizang.json"
}

func init() {
	template.Register((*BossMiZangTemplate)(nil))
}
