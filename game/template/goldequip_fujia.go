package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

//金装附件属性随机配置
type GoldEquipFuJiaTemplate struct {
	*GoldEquipFuJiaTemplateVO
	startAttrMap      map[int32]int32                         //起始属性池map
	attrPoolMap       map[int32][]*GoldEquipFuJiaAttrTemplate //属性池map
	attrNumWeightList []int64                                 //随机属性数量权重列表
}

func (t *GoldEquipFuJiaTemplate) TemplateId() int {
	return t.Id
}

func (t *GoldEquipFuJiaTemplate) RandomAttr() (fujiaList []int32) {
	//随机附件数
	index := mathutils.RandomWeights(t.attrNumWeightList)
	if index < 0 {
		return
	}
	attrNum := int32(index)

	//随机属性
	for attrIndex := int32(1); attrIndex <= attrNum; attrIndex++ {
		attrList := t.attrPoolMap[attrIndex]
		var attrWeightList []int64
		for _, attrTemp := range attrList {
			attrWeightList = append(attrWeightList, int64(attrTemp.Rate))
		}

		index := mathutils.RandomWeights(attrWeightList)
		if index < 0 {
			continue
		}

		fujiaList = append(fujiaList, int32(attrList[index].Id))
	}

	return
}

func (t *GoldEquipFuJiaTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//随机属性数量权重列表
	t.attrNumWeightList = append(t.attrNumWeightList, int64(t.Rate0))
	t.attrNumWeightList = append(t.attrNumWeightList, int64(t.Rate1))
	t.attrNumWeightList = append(t.attrNumWeightList, int64(t.Rate2))
	t.attrNumWeightList = append(t.attrNumWeightList, int64(t.Rate3))
	t.attrNumWeightList = append(t.attrNumWeightList, int64(t.Rate4))
	t.attrNumWeightList = append(t.attrNumWeightList, int64(t.Rate5))
	t.attrNumWeightList = append(t.attrNumWeightList, int64(t.Rate6))

	// 属性池id
	t.startAttrMap = make(map[int32]int32)
	t.startAttrMap[1] = t.Pool1
	t.startAttrMap[2] = t.Pool2
	t.startAttrMap[3] = t.Pool3
	t.startAttrMap[4] = t.Pool4
	t.startAttrMap[5] = t.Pool5
	t.startAttrMap[6] = t.Pool6
	for index, poolId := range t.startAttrMap {
		goldequipAttrTempObj := template.GetTemplateService().Get(int(poolId), (*GoldEquipFuJiaAttrTemplate)(nil))
		if goldequipAttrTempObj == nil {
			err = fmt.Errorf("Pool%d [%d] invalid", index, poolId)
			return template.NewTemplateFieldError("Pool", err)
		}
	}
	return nil
}
func (t *GoldEquipFuJiaTemplate) PatchAfterCheck() {
	//动态属性池模板
	t.attrPoolMap = make(map[int32][]*GoldEquipFuJiaAttrTemplate)
	//赋值 attrPoolMap
	for index, poolId := range t.startAttrMap {
		tempObj := template.GetTemplateService().Get(int(poolId), (*GoldEquipFuJiaAttrTemplate)(nil))
		startTemplate := tempObj.(*GoldEquipFuJiaAttrTemplate)
		for startTemplate != nil {
			t.attrPoolMap[index] = append(t.attrPoolMap[index], startTemplate)
			startTemplate = startTemplate.nextTemp
		}
	}

}
func (t *GoldEquipFuJiaTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (edt *GoldEquipFuJiaTemplate) FileName() string {
	return "tb_goldequip_fujia.json"
}

func init() {
	template.Register((*GoldEquipFuJiaTemplate)(nil))
}
