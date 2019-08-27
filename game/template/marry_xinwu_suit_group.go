package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type MarryXinWuSuitGroupTemplate struct {
	*MarryXinWuSuitGroupTemplateVO
	suitItemMap    map[int32]*MarryXinWuTemplate
	suitAddMap     map[int32]*MarryXinWuSuitTemplate
	itemIdPosMap   map[int32]int32
	itemNamePosMap map[int32]string
}

func (t *MarryXinWuSuitGroupTemplate) TemplateId() int {
	return t.Id
}

func (t *MarryXinWuSuitGroupTemplate) FileName() string {
	return "tb_marry_xinwu_suit_group.json"
}

func (t *MarryXinWuSuitGroupTemplate) Check() (err error) {
	err = validator.MinValidate(float64(t.KaifuTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.KaifuTime)
		err = template.NewTemplateFieldError("KaifuTime", err)
		return
	}
	err = validator.MinValidate(float64(t.MaxNum), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MaxNum)
		err = template.NewTemplateFieldError("MaxNum", err)
		return
	}
	item1 := template.GetTemplateService().Get(int(t.Pos1Id), (*ItemTemplate)(nil))
	if item1 == nil {
		err = fmt.Errorf("[%d]无效", t.Pos1Id)
		err = template.NewTemplateFieldError("Pos1Id", err)
		return err
	}
	item1 = template.GetTemplateService().Get(int(t.Pos2Id), (*ItemTemplate)(nil))
	if item1 == nil {
		err = fmt.Errorf("[%d]无效", t.Pos2Id)
		err = template.NewTemplateFieldError("Pos2Id", err)
		return err
	}
	item1 = template.GetTemplateService().Get(int(t.Pos3Id), (*ItemTemplate)(nil))
	if item1 == nil {
		err = fmt.Errorf("[%d]无效", t.Pos3Id)
		err = template.NewTemplateFieldError("Pos3Id", err)
		return err
	}
	item1 = template.GetTemplateService().Get(int(t.Pos4Id), (*ItemTemplate)(nil))
	if item1 == nil {
		err = fmt.Errorf("[%d]无效", t.Pos4Id)
		err = template.NewTemplateFieldError("Pos4Id", err)
		return err
	}
	item1 = template.GetTemplateService().Get(int(t.Pos5Id), (*ItemTemplate)(nil))
	if item1 == nil {
		err = fmt.Errorf("[%d]无效", t.Pos5Id)
		err = template.NewTemplateFieldError("Pos5Id", err)
		return err
	}
	item1 = template.GetTemplateService().Get(int(t.Pos6Id), (*ItemTemplate)(nil))
	if item1 == nil {
		err = fmt.Errorf("[%d]无效", t.Pos6Id)
		err = template.NewTemplateFieldError("Pos6Id", err)
		return err
	}
	item1 = template.GetTemplateService().Get(int(t.Pos7Id), (*ItemTemplate)(nil))
	if item1 == nil {
		err = fmt.Errorf("[%d]无效", t.Pos7Id)
		err = template.NewTemplateFieldError("Pos7Id", err)
		return err
	}
	item1 = template.GetTemplateService().Get(int(t.Pos8Id), (*ItemTemplate)(nil))
	if item1 == nil {
		err = fmt.Errorf("[%d]无效", t.Pos8Id)
		err = template.NewTemplateFieldError("Pos8Id", err)
		return err
	}
	return
}

func (t *MarryXinWuSuitGroupTemplate) Patch() (err error) {
	t.suitItemMap = make(map[int32]*MarryXinWuTemplate)
	t.suitAddMap = make(map[int32]*MarryXinWuSuitTemplate)

	item1 := template.GetTemplateService().Get(int(t.SuitId1), (*MarryXinWuSuitTemplate)(nil))
	if item1 == nil {
		err = fmt.Errorf("[%d]无效", t.SuitId1)
		err = template.NewTemplateFieldError("SuitId1", err)
		return err
	}
	item1Template := item1.(*MarryXinWuSuitTemplate)
	t.suitAddMap[item1Template.Num] = item1Template

	item2 := template.GetTemplateService().Get(int(t.SuitId2), (*MarryXinWuSuitTemplate)(nil))
	if item2 == nil {
		err = fmt.Errorf("[%d]无效", t.SuitId2)
		err = template.NewTemplateFieldError("SuitId2", err)

		return err
	}
	item2Template := item2.(*MarryXinWuSuitTemplate)
	t.suitAddMap[item2Template.Num] = item2Template

	item3 := template.GetTemplateService().Get(int(t.SuitId3), (*MarryXinWuSuitTemplate)(nil))
	if item3 == nil {
		err = fmt.Errorf("[%d]无效", t.SuitId3)
		err = template.NewTemplateFieldError("SuitId3", err)

		return err
	}
	item3Template := item3.(*MarryXinWuSuitTemplate)
	t.suitAddMap[item3Template.Num] = item3Template

	item4 := template.GetTemplateService().Get(int(t.SuitId4), (*MarryXinWuSuitTemplate)(nil))
	if item4 == nil {
		err = fmt.Errorf("[%d]无效", t.SuitId4)
		err = template.NewTemplateFieldError("SuitId4", err)

		return err
	}
	item4Template := item4.(*MarryXinWuSuitTemplate)
	t.suitAddMap[item4Template.Num] = item4Template

	allXinWu := template.GetTemplateService().GetAll((*MarryXinWuTemplate)(nil))
	for _, value := range allXinWu {
		xinWuItem := value.(*MarryXinWuTemplate)
		if xinWuItem.SuitGroup != t.Id {
			continue
		}

		err = validator.MinValidate(float64(xinWuItem.Type), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", xinWuItem.Type)
			err = template.NewTemplateFieldError("xinWuItem.Type", err)
			return
		}
		err = validator.MaxValidate(float64(xinWuItem.Type), float64(t.MaxNum), false)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", xinWuItem.Type)
			err = template.NewTemplateFieldError("xinWuItem.Type", err)
			return
		}
		_, exists := t.suitItemMap[xinWuItem.Type]
		if exists {
			err = fmt.Errorf("[%d] invalid duplicate", xinWuItem.Type)
			err = template.NewTemplateFieldError("xinWuItem.Type", err)
			return
		}
		t.suitItemMap[xinWuItem.Type] = xinWuItem
	}

	for i := int32(0); i < t.MaxNum; i++ {
		_, exists := t.suitItemMap[i]
		if !exists {
			err = fmt.Errorf("[%d] invalid not enought", t.MaxNum)
			err = template.NewTemplateFieldError("MaxNum", err)
			return
		}
	}
	t.initBindGroup()

	return nil
}

func (t *MarryXinWuSuitGroupTemplate) PatchAfterCheck() {
	return
}

func (t *MarryXinWuSuitGroupTemplate) GetSuitItemMap() map[int32]*MarryXinWuTemplate {
	return t.suitItemMap
}

func (t *MarryXinWuSuitGroupTemplate) GetSuitAddMap() map[int32]*MarryXinWuSuitTemplate {
	return t.suitAddMap
}

func (t *MarryXinWuSuitGroupTemplate) initBindGroup() {
	t.itemIdPosMap = make(map[int32]int32)
	t.itemNamePosMap = make(map[int32]string)
	t.itemIdPosMap[0] = t.Pos1Id
	t.itemIdPosMap[1] = t.Pos2Id
	t.itemIdPosMap[2] = t.Pos3Id
	t.itemIdPosMap[3] = t.Pos4Id
	t.itemIdPosMap[4] = t.Pos5Id
	t.itemIdPosMap[5] = t.Pos6Id
	t.itemIdPosMap[6] = t.Pos7Id
	t.itemIdPosMap[7] = t.Pos8Id

	t.itemNamePosMap[0] = t.Pos1Name
	t.itemNamePosMap[1] = t.Pos2Name
	t.itemNamePosMap[2] = t.Pos3Name
	t.itemNamePosMap[3] = t.Pos4Name
	t.itemNamePosMap[4] = t.Pos5Name
	t.itemNamePosMap[5] = t.Pos6Name
	t.itemNamePosMap[6] = t.Pos7Name
	t.itemNamePosMap[7] = t.Pos8Name
}

func (t *MarryXinWuSuitGroupTemplate) GetXinWuItemId(pos int32) int32 {
	value, exists := t.itemIdPosMap[pos]
	if exists {
		return value
	}
	return int32(0)
}

func (t *MarryXinWuSuitGroupTemplate) GetXinWuItemCount() int {
	return len(t.itemIdPosMap)
}

func (t *MarryXinWuSuitGroupTemplate) GetXinWuItemName(pos int32) string {
	value, exists := t.itemNamePosMap[pos]
	if exists {
		return value
	}
	return ""
}

func init() {
	template.Register((*MarryXinWuSuitGroupTemplate)(nil))
}
