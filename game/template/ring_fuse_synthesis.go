package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

// 特戒融合合成配置
type RingFuseSynthesisTemplate struct {
	*RingFuseSynthesisTemplateVO
	needItemMap    map[int32]int32 // 需要的物品id
	receiveItemMap map[int32]int32 // 合成的物品id
}

func (t *RingFuseSynthesisTemplate) TemplateId() int {
	return t.Id
}

func (t *RingFuseSynthesisTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.needItemMap = make(map[int32]int32)
	//验证物品id
	if t.NeedItemId2 != 0 {
		to := template.GetTemplateService().Get(int(t.NeedItemId2), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NeedItemId2)
			return template.NewTemplateFieldError("NeedItemId2", err)
		}

		err = validator.MinValidate(float64(t.NeedItemCount2), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.NeedItemCount2)
			return template.NewTemplateFieldError("NeedItemCount2", err)
		}
		t.needItemMap[t.NeedItemId2] = t.NeedItemCount2
	}

	//验证物品id
	if t.ItemId != 0 {
		to := template.GetTemplateService().Get(int(t.ItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.ItemId)
			return template.NewTemplateFieldError("ItemId", err)
		}

		err = validator.MinValidate(float64(t.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.ItemCount)
			return template.NewTemplateFieldError("ItemCount", err)
		}
		t.needItemMap[t.ItemId] = t.ItemCount
	}

	return
}

func (t *RingFuseSynthesisTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 need_silver
	err = validator.MinValidate(float64(t.NeedSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedSilver)
		return template.NewTemplateFieldError("NeedSilver", err)
	}

	// 验证 need_gold
	err = validator.MinValidate(float64(t.NeedGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedGold)
		return template.NewTemplateFieldError("NeedGold", err)
	}

	// 验证 need_bind_gold
	err = validator.MinValidate(float64(t.NeedBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedBindGold)
		return template.NewTemplateFieldError("NeedBindGold", err)
	}

	return nil
}

func (t *RingFuseSynthesisTemplate) PatchAfterCheck() {
}

func (t *RingFuseSynthesisTemplate) FileName() string {
	return "tb_tejie_ronghe_synthesis.json"
}

func init() {
	template.Register((*RingFuseSynthesisTemplate)(nil))
}
