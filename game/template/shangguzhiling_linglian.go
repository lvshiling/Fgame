package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	shangguzhilingtypes "fgame/fgame/game/shangguzhiling/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

type ShangguzhilingLinglianTemplate struct {
	*ShangguzhilingLinglianTemplateVO
	//灵兽类型
	lingshouType shangguzhilingtypes.LingshouType
	//部位类型
	linglianPosType shangguzhilingtypes.LinglianPosType

	//属性池
	poolMap       map[int32]*ShangguzhilingLinglianPoolTemplate
	firstPoolTemp *ShangguzhilingLinglianPoolTemplate
	//初始时属性随机池
	beginPoolMap       map[int32]*ShangguzhilingLinglianPoolTemplate
	firstBeginPoolTemp *ShangguzhilingLinglianPoolTemplate

	//所有的属性随机池
	allPoolMap map[int32]*ShangguzhilingLinglianPoolTemplate
}

//灵兽类型
func (t *ShangguzhilingLinglianTemplate) GetLingShouType() shangguzhilingtypes.LingshouType {
	return t.lingshouType
}

//灵炼部位类型
func (t *ShangguzhilingLinglianTemplate) GetLingLianPosType() shangguzhilingtypes.LinglianPosType {
	return t.linglianPosType
}

//获取属性模板
func (t *ShangguzhilingLinglianTemplate) GetPoolTemp(biaoshi int32) *ShangguzhilingLinglianPoolTemplate {
	temp, ok := t.poolMap[biaoshi]
	if !ok {
		return nil
	}
	return temp
}

//根据id获取属性模板
func (t *ShangguzhilingLinglianTemplate) GetPoolTempById(id int32) *ShangguzhilingLinglianPoolTemplate {
	temp, ok := t.allPoolMap[id]
	if !ok {
		return nil
	}
	return temp
}

//随机属性
func (t *ShangguzhilingLinglianTemplate) GetRandomPoolTempMark() int32 {
	weights := []int64{}
	indexToBiaoshiMap := make(map[int]int32)
	for biaoshi, poolTemp := range t.poolMap {
		indexToBiaoshiMap[len(weights)] = biaoshi
		weights = append(weights, int64(poolTemp.Rate))
	}
	index := mathutils.RandomWeights(weights)
	if index < 0 {
		return 0
	}
	return int32(t.poolMap[indexToBiaoshiMap[index]].Id)
}

//获取初始属性模板
func (t *ShangguzhilingLinglianTemplate) GetBeginPoolTemp(biaoshi int32) *ShangguzhilingLinglianPoolTemplate {
	temp, ok := t.beginPoolMap[biaoshi]
	if !ok {
		return nil
	}
	return temp
}

//随机初始属性
func (t *ShangguzhilingLinglianTemplate) GetBeginRandomPoolTempMark() int32 {
	weights := []int64{}
	indexToBiaoshiMap := make(map[int]int32)
	for biaoshi, poolTemp := range t.beginPoolMap {
		indexToBiaoshiMap[len(weights)] = biaoshi
		weights = append(weights, int64(poolTemp.Rate))
	}
	index := mathutils.RandomWeights(weights)
	if index < 0 {
		return 0
	}
	return int32(t.beginPoolMap[indexToBiaoshiMap[index]].Id)
}

func (t *ShangguzhilingLinglianTemplate) TemplateId() int {
	return t.Id
}

func (t *ShangguzhilingLinglianTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShangguzhilingLinglianTemplate) PatchAfterCheck() {
	t.allPoolMap = make(map[int32]*ShangguzhilingLinglianPoolTemplate)

	t.poolMap = make(map[int32]*ShangguzhilingLinglianPoolTemplate)
	for temp := t.firstPoolTemp; temp != nil; temp = temp.GetNextPoolTemplate() {
		t.poolMap[temp.Biaoshi] = temp
		t.allPoolMap[int32(temp.Id)] = temp
	}

	t.beginPoolMap = make(map[int32]*ShangguzhilingLinglianPoolTemplate)
	for temp := t.firstBeginPoolTemp; temp != nil; temp = temp.GetNextPoolTemplate() {
		t.beginPoolMap[temp.Biaoshi] = temp
		t.allPoolMap[int32(temp.Id)] = temp
	}
}

func (t *ShangguzhilingLinglianTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//灵兽类型
	lingshouType := shangguzhilingtypes.LingshouType(t.Type)
	if !lingshouType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}
	t.lingshouType = lingshouType

	//灵炼部位类型
	linglianPosType := shangguzhilingtypes.LinglianPosType(t.SubType)
	if !linglianPosType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.SubType)
		return template.NewTemplateFieldError("SubType", err)
	}
	t.linglianPosType = linglianPosType

	//所需的上古之灵等级
	err = validator.MinValidate(float64(t.NeedSgzlLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedSgzlLevel)
		return template.NewTemplateFieldError("NeedSgzlLevel", err)
	}

	//初始时的属性池
	poolTempInterface := template.GetTemplateService().Get(int(t.ChushiAttrPoolBeginId), (*ShangguzhilingLinglianPoolTemplate)(nil))
	if poolTempInterface == nil {
		err = fmt.Errorf("ShangguzhilingLinglianPoolTemplate [%d] no exist", t.ChushiAttrPoolBeginId)
		return template.NewTemplateFieldError("ChushiAttrPoolBeginId", err)
	}
	poolTemp, ok := poolTempInterface.(*ShangguzhilingLinglianPoolTemplate)
	if !ok {
		err = fmt.Errorf("ShangguzhilingLinglianPoolTemplate assert [%d] no exist", t.ChushiAttrPoolBeginId)
		return template.NewTemplateFieldError("ChushiAttrPoolBeginId", err)
	}
	t.firstBeginPoolTemp = poolTemp

	//关联属性池
	poolTempInterface = template.GetTemplateService().Get(int(t.AttrPoolBeginId), (*ShangguzhilingLinglianPoolTemplate)(nil))
	if poolTempInterface == nil {
		err = fmt.Errorf("ShangguzhilingLinglianPoolTemplate [%d] no exist", t.AttrPoolBeginId)
		return template.NewTemplateFieldError("AttrPoolBeginId", err)
	}
	poolTemp, ok = poolTempInterface.(*ShangguzhilingLinglianPoolTemplate)
	if !ok {
		err = fmt.Errorf("ShangguzhilingLinglianPoolTemplate assert [%d] no exist", t.AttrPoolBeginId)
		return template.NewTemplateFieldError("AttrPoolBeginId", err)
	}
	t.firstPoolTemp = poolTemp

	return nil
}

func (t *ShangguzhilingLinglianTemplate) FileName() string {
	return "tb_sgzl_linglian_pos.json"
}

func init() {
	template.Register((*ShangguzhilingLinglianTemplate)(nil))
}
