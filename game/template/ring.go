package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/property/types"
	propertytypes "fgame/fgame/game/property/types"
	ringtypes "fgame/fgame/game/ring/types"
	"fmt"
)

// 特戒配置
type RingTemplate struct {
	*RingTemplateVO
	ringType              ringtypes.RingType
	fuseSynthesisTemplate *RingFuseSynthesisTemplate                 // 融合合成模板
	beginAdvanceTemp      *RingAdvanceTemplate                       // 起始进阶模板
	beginStrengthenTemp   *RingStrengthenTemplate                    // 起始强化模板
	beginJingLingTemp     *RingJingLingTemplate                      // 起始净灵模板
	advanceMap            map[int32]*RingAdvanceTemplate             // 进阶
	strengthenMap         map[int32]*RingStrengthenTemplate          // 强化
	jingLingMap           map[int32]*RingJingLingTemplate            // 净灵
	battlePropertyMap     map[propertytypes.BattlePropertyType]int64 // 战斗属性
}

func (t *RingTemplate) TemplateId() int {
	return t.Id
}

func (t *RingTemplate) GetRingType() ringtypes.RingType {
	return t.ringType
}

func (t *RingTemplate) GetFuseSynthesisTemplate() *RingFuseSynthesisTemplate {
	return t.fuseSynthesisTemplate
}

func (t *RingTemplate) GetAdvanceTemplateMap() map[int32]*RingAdvanceTemplate {
	return t.advanceMap
}

func (t *RingTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *RingTemplate) GetAdvanceTemplate(advance int32) *RingAdvanceTemplate {
	temp, ok := t.advanceMap[advance]
	if !ok {
		return nil
	}
	return temp
}

func (t *RingTemplate) GetStrengthenTemplate(level int32) *RingStrengthenTemplate {
	temp, ok := t.strengthenMap[level]
	if !ok {
		return nil
	}
	return temp
}

func (t *RingTemplate) GetJingLingTemplate(level int32) *RingJingLingTemplate {
	temp, ok := t.jingLingMap[level]
	if !ok {
		return nil
	}
	return temp
}

func (t *RingTemplate) Patch() (err error) {
	return
}

func (t *RingTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.ringType = ringtypes.RingType(t.Type)
	if !t.ringType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	// 验证id
	if t.Id != 0 {
		temp := template.GetTemplateService().Get(t.Id, (*ItemTemplate)(nil))
		if temp == nil {
			err = fmt.Errorf("[%d] invalid", t.Id)
			err = template.NewTemplateFieldError("Id", err)
			return
		}
	}

	// 验证起始模板id
	if t.AdvanceBeginId != 0 {
		upStarTemplate := template.GetTemplateService().Get(int(t.AdvanceBeginId), (*RingAdvanceTemplate)(nil))
		if upStarTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.AdvanceBeginId)
			err = template.NewTemplateFieldError("AdvanceBeginId", err)
			return
		}
		temp, ok := upStarTemplate.(*RingAdvanceTemplate)
		if !ok {
			return fmt.Errorf("AdvanceBeginId [%d] invalid", t.AdvanceBeginId)
		}
		t.beginAdvanceTemp = temp
	}

	// 验证起始模板id
	if t.StrengthenBeginId != 0 {
		upStarTemplate := template.GetTemplateService().Get(int(t.StrengthenBeginId), (*RingStrengthenTemplate)(nil))
		if upStarTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.StrengthenBeginId)
			err = template.NewTemplateFieldError("StrengthenBeginId", err)
			return
		}
		temp, ok := upStarTemplate.(*RingStrengthenTemplate)
		if !ok {
			return fmt.Errorf("StrengthenBeginId [%d] invalid", t.StrengthenBeginId)
		}
		t.beginStrengthenTemp = temp
	}

	// 验证起始模板id
	if t.JingLingBeginId != 0 {
		upStarTemplate := template.GetTemplateService().Get(int(t.JingLingBeginId), (*RingJingLingTemplate)(nil))
		if upStarTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.JingLingBeginId)
			err = template.NewTemplateFieldError("JingLingBeginId", err)
			return
		}
		temp, ok := upStarTemplate.(*RingJingLingTemplate)
		if !ok {
			return fmt.Errorf("JingLingBeginId [%d] invalid", t.JingLingBeginId)
		}
		t.beginJingLingTemp = temp
	}

	// 验证等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 hp
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//验证 attack
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//验证 defence
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	//验证 skill_id
	if t.SkillId != 0 {
		to := template.GetTemplateService().Get(int(t.SkillId), (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.SkillId)
			return template.NewTemplateFieldError("SkillId", err)
		}
	}

	//验证 ronghe_synthesis_id
	if t.FuseSynthesisId != 0 {
		to := template.GetTemplateService().Get(int(t.FuseSynthesisId), (*RingFuseSynthesisTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.SkillId)
			return template.NewTemplateFieldError("SkillId", err)
		}
		temp := to.(*RingFuseSynthesisTemplate)
		t.fuseSynthesisTemplate = temp
	}

	return nil
}

func (t *RingTemplate) PatchAfterCheck() {
	if t.AdvanceBeginId != 0 {
		t.advanceMap = make(map[int32]*RingAdvanceTemplate)
		for tempTemplate := t.beginAdvanceTemp; tempTemplate != nil; tempTemplate = tempTemplate.nextTemp {
			advance := tempTemplate.Advance
			t.advanceMap[advance] = tempTemplate
		}
	}

	if t.StrengthenBeginId != 0 {
		t.strengthenMap = make(map[int32]*RingStrengthenTemplate)
		for tempTemplate := t.beginStrengthenTemp; tempTemplate != nil; tempTemplate = tempTemplate.nextTemp {
			level := tempTemplate.Level
			t.strengthenMap[level] = tempTemplate
		}
	}

	if t.JingLingBeginId != 0 {
		t.jingLingMap = make(map[int32]*RingJingLingTemplate)
		for tempTemplate := t.beginJingLingTemp; tempTemplate != nil; tempTemplate = tempTemplate.nextTemp {
			level := tempTemplate.Level
			t.jingLingMap[level] = tempTemplate
		}
	}

	t.battlePropertyMap = make(map[types.BattlePropertyType]int64)
	t.battlePropertyMap[types.BattlePropertyTypeMaxHP] = int64(t.Hp)
	t.battlePropertyMap[types.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[types.BattlePropertyTypeDefend] = int64(t.Defence)
}

func (t *RingTemplate) FileName() string {
	return "tb_tejie.json"
}

func init() {
	template.Register((*RingTemplate)(nil))
}
