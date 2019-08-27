package template

import (
	"fgame/fgame/core/template"
)

//装备强化概率配置
type GoldEquipStrengthenRateTemplate struct {
	*GoldEquipStrengthenRateTemplateVO
	rateMap      map[int32]int32 //提供强化概率
	needItemList []int32
}

func (t *GoldEquipStrengthenRateTemplate) TemplateId() int {
	return t.Id
}

func (t *GoldEquipStrengthenRateTemplate) OfferRate(level int32) int32 {
	return t.rateMap[level]
}

func (t *GoldEquipStrengthenRateTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//提供的概率
	t.rateMap = make(map[int32]int32)
	if t.GiveRate1 != 0 {
		t.rateMap[0] = t.GiveRate1
	}
	if t.GiveRate2 != 0 {
		t.rateMap[1] = t.GiveRate2
	}
	if t.GiveRate3 != 0 {
		t.rateMap[2] = t.GiveRate3
	}
	if t.GiveRate4 != 0 {
		t.rateMap[3] = t.GiveRate4
	}
	if t.GiveRate5 != 0 {
		t.rateMap[4] = t.GiveRate5
	}
	if t.GiveRate6 != 0 {
		t.rateMap[5] = t.GiveRate6
	}
	if t.GiveRate7 != 0 {
		t.rateMap[6] = t.GiveRate7
	}
	if t.GiveRate8 != 0 {
		t.rateMap[7] = t.GiveRate8
	}
	if t.GiveRate9 != 0 {
		t.rateMap[8] = t.GiveRate9
	}
	if t.GiveRate10 != 0 {
		t.rateMap[9] = t.GiveRate10
	}

	return nil
}
func (t *GoldEquipStrengthenRateTemplate) PatchAfterCheck() {

}
func (t *GoldEquipStrengthenRateTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (edt *GoldEquipStrengthenRateTemplate) FileName() string {
	return "tb_goldequip_upstar_rate.json"
}

func init() {
	template.Register((*GoldEquipStrengthenRateTemplate)(nil))
}
