package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	soultypes "fgame/fgame/game/soul/types"
	"fmt"
)

//帝魂配置
type SoulTemplate struct {
	*SoulTemplateVO
	kindType           soultypes.SoulKindType //帝魂种类
	soulType           soultypes.SoulType     //帝魂类型
	needItemMap        map[int32]int32        //激活物品
	preSoulTemplate    *SoulTemplate          //激活前置帝魂条件
	battleAttrTemplate *AttrTemplate          //阶别属性
}

func (st *SoulTemplate) TemplateId() int {
	return st.Id
}

func (st *SoulTemplate) GetNeedItemMap() map[int32]int32 {
	return st.needItemMap
}

func (st *SoulTemplate) GetPreSoulCond() *SoulTemplate {
	return st.preSoulTemplate
}

func (st *SoulTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return st.battleAttrTemplate
}

func (st *SoulTemplate) GetSoulType() soultypes.SoulType {
	return st.soulType
}

func (st *SoulTemplate) GetKindType() soultypes.SoulKindType {
	return st.kindType
}

func (st *SoulTemplate) PatchAfterCheck() {

}

func (st *SoulTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(st.FileName(), st.TemplateId(), err)
			return
		}
	}()
	//type
	st.kindType = soultypes.SoulKindType(st.Type)
	if !st.kindType.Valid() {
		err = fmt.Errorf("[%d] invalid", st.Type)
		err = template.NewTemplateFieldError("Type", err)
		return err
	}

	//soul_typ
	st.soulType = soultypes.SoulType(st.SoulType)
	if !st.soulType.Valid() {
		err = fmt.Errorf("[%d] invalid", st.SoulType)
		err = template.NewTemplateFieldError("SoulType", err)
		return err
	}

	//need_item_id
	if st.NeedItemId != "" {
		itemArr, err := utils.SplitAsIntArray(st.NeedItemId)
		if err != nil {
			return err
		}
		numArr, err := utils.SplitAsIntArray(st.NeedItemCount)
		if err != nil {
			return err
		}
		if len(itemArr) <= 0 || len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", st.NeedItemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return err
		}
		st.needItemMap = make(map[int32]int32)
		for i := 0; i < len(itemArr); i++ {
			st.needItemMap[itemArr[i]] = numArr[i]
		}

		//验证 level
		if st.Level != 1 {
			err = fmt.Errorf("[%d] invalid", st.Level)
			err = template.NewTemplateFieldError("Level", err)
			return err
		}
	}

	for itemId, num := range st.needItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", st.NeedItemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return err
		}
		//need_item_count
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", st.NeedItemCount)
			return template.NewTemplateFieldError("NeedItemCount", err)
		}
	}

	//uplevel_exp
	if st.UplevelExp != 0 {
		err = validator.MinValidate(float64(st.UplevelExp), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", st.UplevelExp)
			return template.NewTemplateFieldError("UplevelExp", err)
		}
		if st.DevourId == 0 {
			err = fmt.Errorf("[%d] invalid", st.DevourId)
			return template.NewTemplateFieldError("DevourId", err)
		}
	}

	//devour_id
	if st.DevourId != 0 {
		to := template.GetTemplateService().Get(int(st.DevourId), (*SoulDevourTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", st.DevourId)
			err = template.NewTemplateFieldError("DevourId", err)
			return err
		}
	}

	//验证 need_soul
	if st.NeedSoul != 0 {
		to := template.GetTemplateService().Get(int(st.NeedSoul), (*SoulTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", st.NeedSoul)
			return template.NewTemplateFieldError("NeedSoul", err)
		}
		st.preSoulTemplate = to.(*SoulTemplate)
	}

	//阶别attr属性
	to := template.GetTemplateService().Get(int(st.AttrId), (*AttrTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", st.AttrId)
		return template.NewTemplateFieldError("AttrId", err)
	}
	attrTemplate, _ := to.(*AttrTemplate)
	st.battleAttrTemplate = attrTemplate

	return nil
}

func (st *SoulTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(st.FileName(), st.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if st.NextId != 0 {
		diff := st.NextId - int32(st.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", st.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return err
		}
		to := template.GetTemplateService().Get(int(st.NextId), (*SoulTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", st.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		sto := to.(*SoulTemplate)
		//验证 level
		diffLeve := sto.Level - st.Level
		if diffLeve != 1 {
			err = fmt.Errorf("[%d] invalid", st.Level)
			return template.NewTemplateFieldError("Level", err)
		}

		//验证 type
		if sto.Type != st.Type {
			err = fmt.Errorf("[%d] invalid", sto.Type)
			return template.NewTemplateFieldError("Type", err)
		}

		//验证 soul_type
		if sto.soulType != st.soulType {
			err = fmt.Errorf("[%d] invalid", sto.soulType)
			return template.NewTemplateFieldError("soulType", err)
		}

		//验证 devour_id
		if sto.DevourId != st.DevourId {
			err = fmt.Errorf("[%d] invalid", sto.DevourId)
			return template.NewTemplateFieldError("DevourId", err)
		}
	}

	return nil
}

func (st *SoulTemplate) FileName() string {
	return "tb_soul.json"
}

func init() {
	template.Register((*SoulTemplate)(nil))
}
