package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/title/types"
	"fmt"
)

//称号配置
type TitleTemplate struct {
	*TitleTemplateVO
	titleType              types.TitleType                //类型
	titleSubType           types.TitleSubType             //子类型
	battleAttrTemplate     *AttrTemplate                  //阶别属性
	needItemMap            map[int32]int32                //激活需要物品
	upStarBeginTemplate    *TitleUpStarTemplate           //称号升星起始模板
	titleUpStarTemplateMap map[int32]*TitleUpStarTemplate //称号升星模板
}

func (tt *TitleTemplate) TemplateId() int {
	return tt.Id
}

func (tt *TitleTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return tt.battleAttrTemplate
}

func (tt *TitleTemplate) GetNeedItemMap() map[int32]int32 {
	return tt.needItemMap
}

func (tt *TitleTemplate) GetTitleType() types.TitleType {
	return tt.titleType
}

func (tt *TitleTemplate) GetTitleSubType() types.TitleSubType {
	return tt.titleSubType
}

func (tt *TitleTemplate) GetTitleUpStarTemplateByStarLev(starLev int32) *TitleUpStarTemplate {
	temp, exist := tt.titleUpStarTemplateMap[starLev]
	if !exist {
		return nil
	}
	return temp
}

func (tt *TitleTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	//验证类型
	tt.titleType = types.TitleType(tt.Type)
	if !tt.titleType.Valid() {
		err = fmt.Errorf("[%d] invalid", tt.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//验证 subType
	tt.titleSubType = types.CreateTitleSubType(tt.titleType, tt.SubType)
	if tt.titleSubType == nil || !tt.titleSubType.Valid() {
		err = fmt.Errorf("[%d] invalid", tt.SubType)
		return template.NewTemplateFieldError("subType", err)
	}

	if tt.Attr != 0 {
		//验证 Attr
		tempAttrTemplate := template.GetTemplateService().Get(int(tt.Attr), (*AttrTemplate)(nil))
		if tempAttrTemplate == nil {
			err = fmt.Errorf("[%d] invalid", tt.Attr)
			return template.NewTemplateFieldError("Attr", err)
		}
		attrTemplate, _ := tempAttrTemplate.(*AttrTemplate)
		tt.battleAttrTemplate = attrTemplate
	}
	//验证 NeedItemId
	tt.needItemMap = make(map[int32]int32)
	if tt.NeedItemId != 0 {
		needItemTemplateVO := template.GetTemplateService().Get(int(tt.NeedItemId), (*ItemTemplate)(nil))
		if needItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", tt.NeedItemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return
		}

		//验证 NeedItemCount
		err = validator.MinValidate(float64(tt.NeedItemCount), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("NeedItemCount", err)
			return
		}
		tt.needItemMap[tt.NeedItemId] = tt.NeedItemCount
	}

	return nil
}

func (tt *TitleTemplate) PatchAfterCheck() {
	if tt.UpStarBeginId != 0 {
		tt.titleUpStarTemplateMap = make(map[int32]*TitleUpStarTemplate)
		for tempTemplate := tt.upStarBeginTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextTitleUpStarTemplate {
			level := tempTemplate.Level
			tt.titleUpStarTemplateMap[level] = tempTemplate
		}
	}
}

func (tt *TitleTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	//验证 time
	if tt.titleType == types.TitleTypeActivity {
		if tt.Time == 0 {
			err = fmt.Errorf("[%d] invalid", tt.Time)
			err = template.NewTemplateFieldError("Time", err)
			return
		}
	}

	// 验证 UpStarBeginId
	if tt.UpStarBeginId != 0 {
		upStarTemplate := template.GetTemplateService().Get(int(tt.UpStarBeginId), (*TitleUpStarTemplate)(nil))
		if upStarTemplate == nil {
			err = fmt.Errorf("[%d] invalid", tt.UpStarBeginId)
			err = template.NewTemplateFieldError("UpStarBeginId", err)
			return
		}
		temp, ok := upStarTemplate.(*TitleUpStarTemplate)
		if !ok {
			return fmt.Errorf("UpStarBeginId [%d] invalid", tt.UpStarBeginId)
		}
		tt.upStarBeginTemplate = temp
	}

	return nil
}

func (tt *TitleTemplate) FileName() string {
	return "tb_title.json"
}

func init() {
	template.Register((*TitleTemplate)(nil))
}
