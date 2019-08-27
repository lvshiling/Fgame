package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	constanttypes "fgame/fgame/game/constant/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//元宝送不停配置
type YuanBaoSongBuTingTemplate struct {
	*YuanBaoSongBuTingTemplateVO
	rewData       *propertytypes.RewData //奖励属性
	rewItemMap    map[int32]int32        //奖励物品
	rewOffItemMap map[int32]int32
}

func (t *YuanBaoSongBuTingTemplate) TemplateId() int {
	return t.Id
}

func (t *YuanBaoSongBuTingTemplate) GetRewData() *propertytypes.RewData {
	return t.rewData
}

func (t *YuanBaoSongBuTingTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *YuanBaoSongBuTingTemplate) GetRewOffItemMap() map[int32]int32 {
	return t.rewOffItemMap
}

func (t *YuanBaoSongBuTingTemplate) PatchAfterCheck() {
}

func (t *YuanBaoSongBuTingTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.rewOffItemMap = make(map[int32]int32)
	//验证 rew_silver
	err = validator.MinValidate(float64(t.RewSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewSilver)
		return template.NewTemplateFieldError("RewSilver", err)
	}
	if t.RewSilver != 0 {
		t.rewOffItemMap[constanttypes.SilverItem] += t.RewSilver
	}

	//验证 rew_gold
	err = validator.MinValidate(float64(t.RewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewGold)
		return template.NewTemplateFieldError("RewGold", err)
	}
	if t.RewGold != 0 {
		t.rewOffItemMap[constanttypes.GoldItem] += t.RewGold
	}

	//验证 rew_bind_gold
	err = validator.MinValidate(float64(t.RewBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewBindGold)
		return template.NewTemplateFieldError("RewBindGold", err)
	}
	if t.RewBindGold != 0 {
		t.rewOffItemMap[constanttypes.BindGoldItem] += t.RewBindGold
	}

	t.rewData = propertytypes.CreateRewData(0, 0, t.RewSilver, t.RewGold, t.RewBindGold)

	t.rewItemMap = make(map[int32]int32)
	if t.RewItem != "" {
		if t.RewItemCount == "" {
			err = fmt.Errorf("[%s] invalid", t.RewItem)
			return template.NewTemplateFieldError("RewItem", err)
		}

		itemArr, err := utils.SplitAsIntArray(t.RewItem)
		if err != nil {
			return err
		}
		numArr, err := utils.SplitAsIntArray(t.RewItemCount)
		if err != nil {
			return err
		}
		if len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", t.RewItemCount)
			return template.NewTemplateFieldError("RewItemCount", err)
		}

		for i := 0; i < len(itemArr); i++ {
			t.rewItemMap[itemArr[i]] = numArr[i]
			t.rewOffItemMap[itemArr[i]] += numArr[i]
		}
	}

	return nil
}

func (t *YuanBaoSongBuTingTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 need_gold
	err = validator.MinValidate(float64(t.NeedGold), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedGold)
		return template.NewTemplateFieldError("NeedGold", err)
	}

	return nil
}

func (t *YuanBaoSongBuTingTemplate) FileName() string {
	return "tb_yuanbaosongbuting.json"
}

func init() {
	template.Register((*YuanBaoSongBuTingTemplate)(nil))
}
