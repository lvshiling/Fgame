package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	materialtypes "fgame/fgame/game/material/types"
	"fmt"
)

func init() {
	template.Register((*MaterialTemplate)(nil))
}

type MaterialTemplate struct {
	*MaterialTemplateVO
	saodangItemMap       map[int32]int32
	saodangRewardItemMap map[int32]int32
	saodangRewardDropArr []int32
	mapTemplate          *MapTemplate
	materialType         materialtypes.MaterialType
}

func (t *MaterialTemplate) TemplateId() int {
	return t.Id
}
func (t *MaterialTemplate) FileName() string {
	return "tb_material.json"
}

//组合成需要的数据
func (t *MaterialTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//验证 map_id
	tempMapTemplate := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if tempMapTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}
	t.mapTemplate = tempMapTemplate.(*MapTemplate)

	//验证: 扫荡所需的物品ID,逗号隔开
	t.saodangItemMap = make(map[int32]int32)
	intSaodangNeedItemIdArr, err := utils.SplitAsIntArray(t.SaodangNeedItemId)
	if err != nil {
		return template.NewTemplateFieldError("SaodangNeedItemId", fmt.Errorf("[%s] invalid", t.SaodangNeedItemId))
	}
	intSaodangNeedItemCountArr, err := utils.SplitAsIntArray(t.SaodangNeedItemCount)
	if err != nil {
		return template.NewTemplateFieldError("SaodangNeedItemCount", fmt.Errorf("[%s] invalid", t.SaodangNeedItemCount))
	}
	if len(intSaodangNeedItemIdArr) != len(intSaodangNeedItemCountArr) {
		err = fmt.Errorf("[%s] invalid", t.SaodangNeedItemId)
		return template.NewTemplateFieldError("SaodangNeedItemId", err)
	}
	if len(intSaodangNeedItemIdArr) > 0 {
		//组合数据
		for index, itemId := range intSaodangNeedItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				return template.NewTemplateFieldError("SaodangNeedItemId", fmt.Errorf("[%s] invalid", t.SaodangNeedItemId))
			}

			err = validator.MinValidate(float64(intSaodangNeedItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("SaodangNeedItemCount", err)
			}

			t.saodangItemMap[itemId] = intSaodangNeedItemCountArr[index]
		}
	}

	//验证：扫荡奖励物品ID,逗号隔开
	//验证：扫荡奖励物品数量
	t.saodangRewardItemMap = make(map[int32]int32)
	intRawItemIdArr, err := utils.SplitAsIntArray(t.RawItemId)
	if err != nil {
		return template.NewTemplateFieldError("RawItemId", fmt.Errorf("[%s] invalid", t.RawItemId))
	}
	intRawItemCountArr, err := utils.SplitAsIntArray(t.RawItemCount)
	if err != nil {
		return template.NewTemplateFieldError("RawItemCount", fmt.Errorf("[%s] invalid", t.RawItemCount))
	}
	if len(intRawItemIdArr) != len(intRawItemCountArr) {
		err = fmt.Errorf("[%s] invalid", t.RawItemId)
		return template.NewTemplateFieldError("RawItemId", err)
	}
	if len(intRawItemIdArr) > 0 && len(intRawItemIdArr) == len(intRawItemCountArr) {
		for index, itemId := range intRawItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				return template.NewTemplateFieldError("RawItemId", fmt.Errorf("[%s] invalid", t.RawItemId))
			}

			err = validator.MinValidate(float64(intRawItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("RawItemCount", err)
			}

			//组合数据
			t.saodangRewardItemMap[itemId] = intRawItemCountArr[index]
		}
	}

	//验证：扫荡奖励掉落包ID，逗号隔开
	rawDropIdArr, err := utils.SplitAsIntArray(t.RawDropId)
	if err != nil {
		return template.NewTemplateFieldError("RawDropId", fmt.Errorf("[%s] invalid", t.RawDropId))
	}
	if len(rawDropIdArr) > 0 {
		for _, dropId := range rawDropIdArr {
			dropTmpObj := template.GetTemplateService().Get(int(dropId), (*DropTemplate)(nil))
			if dropTmpObj == nil {
				return template.NewTemplateFieldError("RawDropId", fmt.Errorf("[%s] invalid", t.RawDropId))
			}
			t.saodangRewardDropArr = append(t.saodangRewardDropArr, dropId)
		}
	}

	return nil
}

//检查有效性
func (t *MaterialTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//副本类型
	t.materialType = materialtypes.MaterialType(t.Type)
	if !t.materialType.Valid() {
		err = fmt.Errorf("[%d] invalid ", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//验证：扫荡消耗元宝
	err = validator.MinValidate(float64(t.SaodangNeedGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("SaodangNeedGold", err)
	}
	//验证：扫荡奖励经验值
	err = validator.MinValidate(float64(t.RawExp), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RawExp", err)
	}
	//验证：扫荡奖励经验点
	err = validator.MinValidate(float64(t.RawExpPoint), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RawExpPoint", err)
	}
	//验证：扫荡奖励银两
	err = validator.MinValidate(float64(t.RawSilver), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RawSilver", err)
	}
	//验证：扫荡奖励元宝
	err = validator.MinValidate(float64(t.RawGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RawGold", err)
	}
	//验证：扫荡奖励绑定元宝
	err = validator.MinValidate(float64(t.RawBindGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RawBindGold", err)
	}

	//验证：进入副本所需物品ID
	itemTmpObj := template.GetTemplateService().Get(int(t.NeedItemId), (*ItemTemplate)(nil))
	if itemTmpObj == nil {
		return template.NewTemplateFieldError("NeedItemId", fmt.Errorf("[%d] invalid", t.NeedItemId))
	}
	//验证：进入副本所需物品数量
	err = validator.MinValidate(float64(t.NeedItemCount), float64(1), true)
	if err != nil {
		return template.NewTemplateFieldError("NeedItemCount", err)
	}
	//验证：每日免费的次数
	err = validator.MinValidate(float64(t.Free), float64(1), true)
	if err != nil {
		return template.NewTemplateFieldError("Free", err)
	}
	//验证：扫荡波数限制
	err = validator.MinValidate(float64(t.GroupLimit), float64(0), false)
	if err != nil {
		return template.NewTemplateFieldError("GroupLimit", err)
	}

	return nil
}

//检验后组合
func (t *MaterialTemplate) PatchAfterCheck() {
}

func (t *MaterialTemplate) GetMapTemplate() *MapTemplate {
	return t.mapTemplate
}

func (t *MaterialTemplate) GetMaterialType() materialtypes.MaterialType {
	return t.materialType
}

//获取扫荡所需物品
func (t *MaterialTemplate) GetSaodangItemMap(saoDangNum int32) map[int32]int32 {
	if saoDangNum > 1 {
		newMap := make(map[int32]int32)
		for itemId, num := range t.saodangItemMap {
			newMap[itemId] = num * saoDangNum
		}

		return newMap
	}

	return t.saodangItemMap
}

//获取扫荡奖励物品
func (t *MaterialTemplate) GetSaodangRewardItemMap(saoDangNum int32) map[int32]int32 {
	if saoDangNum > 1 {
		newMap := make(map[int32]int32)
		for itemId, num := range t.saodangItemMap {
			newMap[itemId] = num * saoDangNum
		}

		return newMap
	}
	return t.saodangRewardItemMap
}

//获取扫荡掉落物品
func (t *MaterialTemplate) GetSaodangRewardDropArr() []int32 {
	return t.saodangRewardDropArr
}
