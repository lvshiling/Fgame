package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fmt"
)

//结义信物配置
type JieYiTokenTemplate struct {
	*JieYiTokenTemplateVO
	tokenType                jieyitypes.JieYiTokenType
	needItemMap              map[int32]int32
	beginJieYiTokenLevelTemp *JieYiTokenLevelTemplate           // 起始结义等级模板
	jieYiTokenLevelTempMap   map[int32]*JieYiTokenLevelTemplate // 结义等级模板
}

func (t *JieYiTokenTemplate) TemplateId() int {
	return t.Id
}

func (t *JieYiTokenTemplate) GetJieYiTokenLevelTemp(level int32) *JieYiTokenLevelTemplate {
	temp, ok := t.jieYiTokenLevelTempMap[level]
	if !ok {
		return nil
	}
	return temp
}

func (t *JieYiTokenTemplate) GetTokenType() jieyitypes.JieYiTokenType {
	return t.tokenType
}

func (t *JieYiTokenTemplate) GetNeedItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *JieYiTokenTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证信物类型
	t.tokenType = jieyitypes.JieYiTokenType(t.Type)
	if !t.tokenType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	t.needItemMap = make(map[int32]int32)
	//验证物品id
	if t.UseItemId != 0 {
		to := template.GetTemplateService().Get(int(t.UseItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItemId)
			return template.NewTemplateFieldError("UseItemId", err)
		}

		err = validator.MinValidate(float64(t.UseItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.UseItemCount)
			return template.NewTemplateFieldError("UseItemCount", err)
		}
		t.needItemMap[t.UseItemId] = t.UseItemCount

	}

	return
}

func (t *JieYiTokenTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证信物分享万分比
	err = validator.RangeValidate(float64(t.SharePercent), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.SharePercent)
		return template.NewTemplateFieldError("SharePercent", err)
	}

	//验证 hp
	// err = validator.MinValidate(float64(t.Hp), float64(1), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.Hp)
	// 	return template.NewTemplateFieldError("Hp", err)
	// }

	//验证 attack
	// err = validator.MinValidate(float64(t.Attack), float64(1), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.Attack)
	// 	return template.NewTemplateFieldError("Attack", err)
	// }

	//验证 defence
	// err = validator.MinValidate(float64(t.Defence), float64(1), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.Defence)
	// 	return template.NewTemplateFieldError("Defence", err)
	// }

	// 验证起始模板id
	if t.BeginId != 0 {
		upStarTemplate := template.GetTemplateService().Get(int(t.BeginId), (*JieYiTokenLevelTemplate)(nil))
		if upStarTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.BeginId)
			err = template.NewTemplateFieldError("BeginId", err)
			return
		}
		temp, ok := upStarTemplate.(*JieYiTokenLevelTemplate)
		if !ok {
			return fmt.Errorf("BeginId [%d] invalid", t.BeginId)
		}
		t.beginJieYiTokenLevelTemp = temp
	}

	return nil
}

func (t *JieYiTokenTemplate) PatchAfterCheck() {
	if t.BeginId != 0 {
		t.jieYiTokenLevelTempMap = make(map[int32]*JieYiTokenLevelTemplate)
		for tempTemplate := t.beginJieYiTokenLevelTemp; tempTemplate != nil; tempTemplate = tempTemplate.nextTemp {
			level := tempTemplate.Level
			t.jieYiTokenLevelTempMap[level] = tempTemplate
		}
	}
}

func (t *JieYiTokenTemplate) FileName() string {
	return "tb_jieyi_xinwu.json"
}

func init() {
	template.Register((*JieYiTokenTemplate)(nil))
}
