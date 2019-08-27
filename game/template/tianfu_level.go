package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/player/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
)

//天赋等级配置
type TianFuLevelTemplate struct {
	*TianFuLevelTemplateVO
	nextTianFuLevelTemplate *TianFuLevelTemplate
	needItemMap             map[types.RoleType]map[int32]int32 //需要物品
	areaType                skilltypes.SkillAreaType           //技能范围
	skillArea               skilltypes.SkillArea

	specialEffectType skilltypes.SkillSpecialEffectType
	//距离
	specialDistance float64
	//时间
	specialTime float64
	//表现时间
	specialAnimationTime float64
	//加buff
	buffRateMap map[int32]int32
	//buff动态模板
	buffDongTaiMap map[int32]int32
}

func (t *TianFuLevelTemplate) TemplateId() int {
	return t.Id
}

func (t *TianFuLevelTemplate) GetAreaType() skilltypes.SkillAreaType {
	return t.areaType
}

func (t *TianFuLevelTemplate) GetSkillArea() skilltypes.SkillArea {
	return t.skillArea
}

func (t *TianFuLevelTemplate) GetSpecialEffectType() skilltypes.SkillSpecialEffectType {
	return t.specialEffectType
}

func (t *TianFuLevelTemplate) GetSpecialTime() float64 {
	return t.specialTime
}
func (t *TianFuLevelTemplate) GetSpecialAnimationTime() float64 {
	return t.specialAnimationTime
}

func (t *TianFuLevelTemplate) GetSpecialDistance() float64 {
	return t.specialDistance
}

func (t *TianFuLevelTemplate) GetBuffRateMap() map[int32]int32 {
	return t.buffRateMap
}

func (t *TianFuLevelTemplate) GetBuffDongTaiMap() map[int32]int32 {
	return t.buffDongTaiMap
}

func (t *TianFuLevelTemplate) GetNeedItemMap(roleType types.RoleType) map[int32]int32 {
	itemMap, ok := t.needItemMap[roleType]
	if !ok {
		return nil
	}
	return itemMap
}

func (t *TianFuLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//作用半径
	if err = validator.MinValidate(float64(t.AreaRadius), float64(0), true); err != nil {
		return template.NewTemplateFieldError("AreaRadius", err)
	}

	t.areaType = skilltypes.SkillAreaType(t.AreaType)
	if !t.areaType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.AreaType)
		return template.NewTemplateFieldError("AreaType", err)
	}

	//组成技能范围
	switch t.areaType {
	case skilltypes.SkillAreaTypeLine:
		t.skillArea = skilltypes.NewSkillAreaRectangle(float64(t.AreaRange)/common.MILL_METER, float64(t.AreaRadius)/common.MILL_METER)
	case skilltypes.SkillAreaTypeFan:
		t.skillArea = skilltypes.NewSkillAreaFan(float64(t.AreaRange), float64(t.AreaRadius)/common.MILL_METER)
	case skilltypes.SkillAreaTypeRound:
		t.skillArea = skilltypes.NewSkillAreaRound(float64(t.AreaRadius) / common.MILL_METER)
	}

	//作用范围0-360
	if t.areaType == skilltypes.SkillAreaTypeFan || t.areaType == skilltypes.SkillAreaTypeRound {
		if err = validator.RangeValidate(float64(t.AreaRange), float64(common.MIN_ANGLE), true, float64(common.MAX_ANGLE), true); err != nil {
			return template.NewTemplateFieldError("AreaRange", err)
		}
	} else if t.areaType == skilltypes.SkillAreaTypeLine {
		if err = validator.MinValidate(float64(t.AreaRange), float64(0), false); err != nil {
			return template.NewTemplateFieldError("AreaRange", err)
		}
	}

	//特殊效果
	t.specialEffectType = skilltypes.SkillSpecialEffectType(t.SpecialEffect)
	if !t.specialEffectType.Valid() {
		return fmt.Errorf("specialEffect [%d] invalid", t.SpecialEffect)
	}

	t.specialDistance = float64(t.SpecialEffectValue) / common.MILL_METER
	t.specialTime = float64(t.SpecialEffectValue2) / float64(common.SECOND)
	t.specialAnimationTime = float64(t.SpecialEffectValue3) / float64(common.SECOND)

	//状态id
	if t.AddStatus != "" {
		t.buffRateMap = make(map[int32]int32)
		t.buffDongTaiMap = make(map[int32]int32)
		addStatusArr, err := utils.SplitAsIntArray(t.AddStatus)
		if err != nil {
			return err
		}

		addStatusRateArr, err := utils.SplitAsIntArray(t.AddStatusRate)
		if err != nil {
			return err
		}
		if len(addStatusRateArr) != 0 {
			if len(addStatusArr) != len(addStatusRateArr) {
				err = fmt.Errorf("[%s] invalid", t.AddStatus)
				return template.NewTemplateFieldError("AddStatus", err)
			}
		}

		dongTaiBuffIdArr, err := utils.SplitAsIntArray(t.BuffDongTaiId)
		if err != nil {
			return template.NewTemplateFieldError("BuffDongTaiId", err)
		}
		if len(addStatusArr) != len(dongTaiBuffIdArr) {
			err = fmt.Errorf("AddStatus[%s],BuffDongTaiId[%s]长度不一致", t.AddStatus, t.BuffDongTaiId)
			return template.NewTemplateFieldError("BuffDongTaiId", err)
		}

		for index, addStatus := range addStatusArr {
			if len(addStatusRateArr) != 0 {
				t.buffRateMap[addStatus] = addStatusRateArr[index]
			}
			t.buffDongTaiMap[addStatus] = dongTaiBuffIdArr[index]
		}
	}

	t.needItemMap = make(map[types.RoleType]map[int32]int32)
	//开天
	if t.UseItemId != 0 {
		tempItemTemplate := template.GetTemplateService().Get(int(t.UseItemId), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItemId)
			return template.NewTemplateFieldError("UseItemId", err)
		}

		err = validator.MinValidate(float64(t.UseItemCount), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("UseItemCount", err)
			return
		}
		needItemMap := make(map[int32]int32)
		needItemMap[t.UseItemId] = t.UseItemCount
		t.needItemMap[types.RoleTypeKaiTian] = needItemMap
	}

	//奕剑
	if t.UseItemCountYiJian != 0 {
		tempItemTemplate := template.GetTemplateService().Get(int(t.UseItemIdYiJian), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItemIdYiJian)
			return template.NewTemplateFieldError("UseItemIdYiJian", err)
		}

		err = validator.MinValidate(float64(t.UseItemCountYiJian), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("UseItemCountYiJian", err)
			return
		}
		needItemMap := make(map[int32]int32)
		needItemMap[t.UseItemIdYiJian] = t.UseItemCountYiJian
		t.needItemMap[types.RoleTypeYiJian] = needItemMap
	}

	//破月
	if t.UseItemCountPoYue != 0 {
		tempItemTemplate := template.GetTemplateService().Get(int(t.UseItemIdPoYue), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItemIdPoYue)
			return template.NewTemplateFieldError("UseItemIdPoYue", err)
		}

		err = validator.MinValidate(float64(t.UseItemCountPoYue), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("UseItemCountPoYue", err)
			return
		}
		needItemMap := make(map[int32]int32)
		needItemMap[t.UseItemIdPoYue] = t.UseItemCountPoYue
		t.needItemMap[types.RoleTypePoYue] = needItemMap
	}

	return nil
}

func (t *TianFuLevelTemplate) PatchAfterCheck() {

}

func (t *TianFuLevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.NextId != 0 {
		tempTianFuTemplate := template.GetTemplateService().Get(int(t.NextId), (*TianFuLevelTemplate)(nil))
		if tempTianFuTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		toNextTemplate := tempTianFuTemplate.(*TianFuLevelTemplate)
		if toNextTemplate.Level-t.Level != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		t.nextTianFuLevelTemplate = toNextTemplate
	}

	//特殊效果概率
	if err = validator.MinValidate(float64(t.SpecialEffectRate), float64(0), true); err != nil {
		return template.NewTemplateFieldError("SpecialEffectRate", err)
	}

	//特殊效果距离
	if err = validator.MinValidate(float64(t.SpecialEffectValue), float64(0), true); err != nil {
		return template.NewTemplateFieldError("SpecialEffectValue", err)
	}
	//特殊效果时间
	if err = validator.MinValidate(float64(t.SpecialEffectValue2), float64(0), true); err != nil {
		return template.NewTemplateFieldError("SpecialEffectValue2", err)
	}
	//验证
	for buffId, buffDongTaiId := range t.buffDongTaiMap {
		tempBuffTemplate := template.GetTemplateService().Get(int(buffId), (*BuffTemplate)(nil))
		if tempBuffTemplate == nil {
			err = fmt.Errorf("[%s] 无效", t.AddStatus)
			return template.NewTemplateFieldError("AddStatus", err)
		}
		tempBuffDongTaiTemplate := template.GetTemplateService().Get(int(buffDongTaiId), (*BuffDongTaiTemplate)(nil))
		if tempBuffDongTaiTemplate == nil {
			err = fmt.Errorf("[%s] 无效", t.BuffDongTaiId)
			return template.NewTemplateFieldError("BuffDongTaiId", err)
		}
	}
	return nil
}

func (t *TianFuLevelTemplate) FileName() string {
	return "tb_tianfu_level.json"
}

func init() {
	template.Register((*TianFuLevelTemplate)(nil))
}
