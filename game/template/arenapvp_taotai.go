package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//3v3淘汰配置
type ArenapvpTaoTaiTemplate struct {
	*ArenapvpTaoTaiTemplateVO
	winItemMap   map[int32]int32
	loseItemMap  map[int32]int32
	arenapvpType arenapvptypes.ArenapvpType
	winRwd       *propertytypes.RewData
	loseRwd      *propertytypes.RewData
}

func (t *ArenapvpTaoTaiTemplate) GetArenapvpType() arenapvptypes.ArenapvpType {
	return t.arenapvpType
}

func (t *ArenapvpTaoTaiTemplate) TemplateId() int {
	return t.Id
}

func (t *ArenapvpTaoTaiTemplate) GetJiFen(win bool) int32 {
	if win {
		return t.WinJifen
	} else {
		return t.LoseJifen
	}
}

func (t *ArenapvpTaoTaiTemplate) GetRewData(win bool) *propertytypes.RewData {
	if win {
		return t.winRwd
	} else {
		return t.loseRwd
	}
}

func (t *ArenapvpTaoTaiTemplate) GetRewItemMap(win bool) map[int32]int32 {
	if win {
		return t.winItemMap
	} else {
		return t.loseItemMap
	}
}

func (t *ArenapvpTaoTaiTemplate) PatchAfterCheck() {
	t.winRwd = propertytypes.CreateRewData(0, 0, t.WinSilver, t.WinGold, t.WinBindGold)
	t.loseRwd = propertytypes.CreateRewData(0, 0, t.LoseSilver, t.LoseGold, t.LoseBindGold)
}

func (t *ArenapvpTaoTaiTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.winItemMap = make(map[int32]int32)
	if t.WinItemId != "" {
		itemArr, err := utils.SplitAsIntArray(t.WinItemId)
		if err != nil {
			return err
		}

		numArr, err := utils.SplitAsIntArray(t.WinItemCount)
		if err != nil {
			return err
		}

		if len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", t.WinItemCount)
			return template.NewTemplateFieldError("WinItemCount", err)
		}

		for index, itemId := range itemArr {
			t.winItemMap[itemId] = numArr[index]
		}
	}

	//
	t.loseItemMap = make(map[int32]int32)
	if t.LoseItemId != "" {
		itemArr, err := utils.SplitAsIntArray(t.LoseItemId)
		if err != nil {
			return err
		}

		numArr, err := utils.SplitAsIntArray(t.LoseItemCount)
		if err != nil {
			return err
		}

		if len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", t.LoseItemCount)
			return template.NewTemplateFieldError("LoseItemCount", err)
		}

		for index, itemId := range itemArr {
			t.loseItemMap[itemId] = numArr[index]
		}
	}

	return nil
}

func (t *ArenapvpTaoTaiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.arenapvpType = arenapvptypes.ArenapvpType(t.Type)
	if !t.arenapvpType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//验证 get_silver
	err = validator.MinValidate(float64(t.WinSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.WinSilver)
		err = template.NewTemplateFieldError("WinSilver", err)
		return
	}

	//验证 get_bind_gold
	err = validator.MinValidate(float64(t.WinBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.WinBindGold)
		err = template.NewTemplateFieldError("WinBindGold", err)
		return
	}

	//验证 get_gold
	err = validator.MinValidate(float64(t.WinGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.WinGold)
		err = template.NewTemplateFieldError("WinGold", err)
		return
	}

	for itemId, num := range t.winItemMap {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.WinItemId)
			err = template.NewTemplateFieldError("WinItemId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.WinItemCount)
			err = template.NewTemplateFieldError("WinItemCount", err)
			return
		}
	}

	//验证 get_silver
	err = validator.MinValidate(float64(t.LoseSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LoseSilver)
		err = template.NewTemplateFieldError("LoseSilver", err)
		return
	}

	//验证 get_bind_gold
	err = validator.MinValidate(float64(t.LoseBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LoseBindGold)
		err = template.NewTemplateFieldError("LoseBindGold", err)
		return
	}

	//验证 get_gold
	err = validator.MinValidate(float64(t.LoseGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LoseGold)
		err = template.NewTemplateFieldError("LoseGold", err)
		return
	}

	for itemId, num := range t.loseItemMap {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.LoseItemId)
			err = template.NewTemplateFieldError("LoseItemId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.LoseItemCount)
			err = template.NewTemplateFieldError("LoseItemCount", err)
			return
		}
	}

	return nil
}

func (t *ArenapvpTaoTaiTemplate) FileName() string {
	return "tb_biwudahui_taotai.json"
}

func init() {
	template.Register((*ArenapvpTaoTaiTemplate)(nil))
}
