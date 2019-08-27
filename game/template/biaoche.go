package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	constanttypes "fgame/fgame/game/constant/types"
	transportationtypes "fgame/fgame/game/transportation/types"
	"fmt"
)

//镖车配置
type BiaocheTemplate struct {
	*BiaocheTemplateVO
	transportType    transportationtypes.TransportationType
	finishRewItemMap map[int32]int32
	robRewItemMap    map[int32]int32
	lastRewItemMap   map[int32]int32
	biaocheTempObj   *BiologyTemplate
}

func (t *BiaocheTemplate) TemplateId() int {
	return t.Id
}

func (t *BiaocheTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//押镖成功获得物品数量
	t.finishRewItemMap = make(map[int32]int32)
	intBiaocheAwardItemIdArr, err := utils.SplitAsIntArray(t.BiaocheAwardItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.BiaocheAwardItemId)
		return template.NewTemplateFieldError("BiaocheAwardItemId", err)
	}
	intBiaocheAwardItemCountArr, err := utils.SplitAsIntArray(t.BiaocheAwardItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.BiaocheAwardItemCount)
		return template.NewTemplateFieldError("BiaocheAwardItemCount", err)
	}
	if len(intBiaocheAwardItemIdArr) != len(intBiaocheAwardItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.BiaocheAwardItemId, t.BiaocheAwardItemCount)
		return template.NewTemplateFieldError("BiaocheAwardItemId or BiaocheAwardItemCount", err)
	}
	if len(intBiaocheAwardItemIdArr) > 0 {
		//组合数据
		for index, itemId := range intBiaocheAwardItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.BiaocheAwardItemId)
				return template.NewTemplateFieldError("BiaocheAwardItemId", err)
			}

			err = validator.MinValidate(float64(intBiaocheAwardItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("BiaocheAwardItemCount", err)
			}

			t.finishRewItemMap[itemId] = intBiaocheAwardItemCountArr[index]
		}
	}

	//劫镖成功获得物品数量
	t.robRewItemMap = make(map[int32]int32)
	intJiebiaoAwardItemIdArr, err := utils.SplitAsIntArray(t.JiebiaoAwardItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.JiebiaoAwardItemId)
		return template.NewTemplateFieldError("JiebiaoAwardItemId", err)
	}
	intJiebiaoAwardItemCountArr, err := utils.SplitAsIntArray(t.JiebiaoAwardItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.JiebiaoAwardItemCount)
		return template.NewTemplateFieldError("JiebiaoAwardItemCount", err)
	}
	if len(intJiebiaoAwardItemIdArr) != len(intJiebiaoAwardItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.JiebiaoAwardItemId, t.JiebiaoAwardItemCount)
		return template.NewTemplateFieldError("JiebiaoAwardItemId or JiebiaoAwardItemCount", err)
	}
	if len(intJiebiaoAwardItemIdArr) > 0 {
		//组合数据
		for index, itemId := range intJiebiaoAwardItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.JiebiaoAwardItemId)
				return template.NewTemplateFieldError("JiebiaoAwardItemId", err)
			}

			err = validator.MinValidate(float64(intJiebiaoAwardItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("JiebiaoAwardItemCount", err)
			}

			t.robRewItemMap[itemId] = intJiebiaoAwardItemCountArr[index]
		}
	}

	//被劫镖剩余获得物品数量
	t.lastRewItemMap = make(map[int32]int32)
	intLastAwardItemIdArr, err := utils.SplitAsIntArray(t.BiaocheLoseItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.BiaocheLoseItemId)
		return template.NewTemplateFieldError("BiaocheLoseItemId", err)
	}
	intLastAwardItemCountArr, err := utils.SplitAsIntArray(t.BiaocheLoseItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.BiaocheLoseItemCount)
		return template.NewTemplateFieldError("BiaocheLoseItemCount", err)
	}
	if len(intLastAwardItemIdArr) != len(intLastAwardItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.BiaocheLoseItemId, t.BiaocheLoseItemCount)
		return template.NewTemplateFieldError("BiaocheLoseItemId or BiaocheLoseItemCount", err)
	}
	if len(intLastAwardItemIdArr) > 0 {
		//组合数据
		for index, itemId := range intLastAwardItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.BiaocheLoseItemId)
				return template.NewTemplateFieldError("BiaocheLoseItemId", err)
			}

			err = validator.MinValidate(float64(intLastAwardItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("BiaocheLoseItemCount", err)
			}

			t.lastRewItemMap[itemId] = intLastAwardItemCountArr[index]
		}
	}

	//镖车id
	tempBilogyTempObj := template.GetTemplateService().Get(int(t.BiaocheId), (*BiologyTemplate)(nil))
	if tempBilogyTempObj == nil {
		err = fmt.Errorf("[%d] invalid", t.BiaocheId)
		return template.NewTemplateFieldError("BiaocheId", err)
	}
	t.biaocheTempObj = tempBilogyTempObj.(*BiologyTemplate)

	//起点NPC
	NPCTempObj := template.GetTemplateService().Get(int(t.BiologyId), (*BiologyTemplate)(nil))
	if NPCTempObj == nil {
		err = fmt.Errorf("[%d] invalid", t.BiologyId)
		return template.NewTemplateFieldError("BiologyId", err)
	}
	return nil
}

func (t *BiaocheTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//镖车类型
	typ := transportationtypes.TransportationType(t.Type)
	if !typ.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}
	t.transportType = typ

	//押镖所需银两
	err = validator.MinValidate(float64(t.BiaocheSilver), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("BiaocheSilver", err)
	}
	//押镖所需元宝
	err = validator.MinValidate(float64(t.BiaocheGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("BiaocheGold", err)
	}
	//押镖成功奖励银两
	err = validator.MinValidate(float64(t.BiaocheAwardSilver), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("BiaocheAwardSilver", err)
	}
	//押镖成功奖励元宝
	err = validator.MinValidate(float64(t.BiaocheAwardGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("BiaocheAwardGold", err)
	}
	//劫镖获得银两
	err = validator.MinValidate(float64(t.JiebiaoAwardSilver), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("JiebiaoAwardSilver", err)
	}
	//劫镖获得元宝
	err = validator.MinValidate(float64(t.JiebiaoAwardGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("JiebiaoAwardGold", err)
	}
	//押镖失败获得银两
	err = validator.MinValidate(float64(t.BiaocheLoseSilver), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("BiaocheLoseSilver", err)
	}
	//押镖失败获得元宝
	err = validator.MinValidate(float64(t.BiaocheLoseGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("BiaocheLoseGold", err)
	}

	return nil
}

func (t *BiaocheTemplate) PatchAfterCheck() {
	if t.transportType == transportationtypes.TransportationTypeAlliance {
		if t.BiaocheAwardSilver != 0 {
			t.finishRewItemMap[constanttypes.SilverItem] = t.BiaocheAwardSilver
		}
		if t.BiaocheAwardGold != 0 {
			t.finishRewItemMap[constanttypes.GoldItem] = t.BiaocheAwardGold
		}
		if t.BiaocheLoseSilver != 0 {
			t.lastRewItemMap[constanttypes.SilverItem] = t.BiaocheLoseSilver
		}
		if t.BiaocheLoseGold != 0 {
			t.lastRewItemMap[constanttypes.GoldItem] = t.BiaocheLoseGold
		}
		if t.JiebiaoAwardSilver != 0 {
			t.robRewItemMap[constanttypes.SilverItem] = t.JiebiaoAwardSilver
		}
		if t.JiebiaoAwardGold != 0 {
			t.robRewItemMap[constanttypes.GoldItem] = t.JiebiaoAwardGold
		}
	}
}

func (t *BiaocheTemplate) FileName() string {
	return "tb_biaoche.json"
}

func (t *BiaocheTemplate) GetLastItemMap() map[int32]int32 {
	return t.lastRewItemMap
}

func (t *BiaocheTemplate) GetRobItemMap() map[int32]int32 {
	return t.robRewItemMap
}

func (t *BiaocheTemplate) GetFinishItemMap() map[int32]int32 {
	return t.finishRewItemMap
}

func (t *BiaocheTemplate) GetTransportType() transportationtypes.TransportationType {
	return t.transportType
}

func (t *BiaocheTemplate) GetBiaocheTemplate() *BiologyTemplate {
	return t.biaocheTempObj
}

func init() {
	template.Register((*BiaocheTemplate)(nil))
}
