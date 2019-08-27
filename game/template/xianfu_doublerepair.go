package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	coreutils "fgame/fgame/core/utils"
	majortypes "fgame/fgame/game/major/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//双修配置
type XianFuDoubleRepairTemplate struct {
	*XianFuDoubleRepairTemplateVO
	rewItemMap           map[int32]int32
	majorMapTemplate     *MapTemplate
	saodangItemMap       map[int32]int32
	saodangRewardItemMap map[int32]int32
	saodangRewardDropArr []int32
}

func (t *XianFuDoubleRepairTemplate) TemplateId() int {
	return t.Id
}

func (t *XianFuDoubleRepairTemplate) GetMapId() int32 {
	return t.MapId
}

func (t *XianFuDoubleRepairTemplate) GetBossId() int32 {
	return t.BossId
}

func (t *XianFuDoubleRepairTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *XianFuDoubleRepairTemplate) GetMajorType() majortypes.MajorType {
	return majortypes.MajorTypeShuangXiu
}

func (t *XianFuDoubleRepairTemplate) GetSaodangNeedGold() int32 {
	return t.SaodangNeedGold
}

func (t *XianFuDoubleRepairTemplate) GetRawSilver() int64 {
	return t.RawSilver
}

func (t *XianFuDoubleRepairTemplate) GetRawBindGold() int32 {
	return t.RawBindGold
}

func (t *XianFuDoubleRepairTemplate) GetRawGold() int32 {
	return t.RawGold
}

func (t *XianFuDoubleRepairTemplate) GetRawExp() int64 {
	return t.RawExp
}

func (t *XianFuDoubleRepairTemplate) GetRawExpPoint() int64 {
	return t.RawExpPoint
}

//获取扫荡所需物品
func (t *XianFuDoubleRepairTemplate) GetSaodangItemMap(saoDangNum int32) map[int32]int32 {
	newMap := make(map[int32]int32)
	if saoDangNum < 1 {
		return nil
	}
	for itemId, num := range t.saodangItemMap {
		newMap[itemId] = num * saoDangNum
	}
	return newMap
}

//获取扫荡奖励物品
func (t *XianFuDoubleRepairTemplate) GetSaodangRewardItemMap(saoDangNum int32) map[int32]int32 {
	newMap := make(map[int32]int32)
	if saoDangNum < 1 {
		return nil
	}
	for itemId, num := range t.saodangRewardItemMap {
		newMap[itemId] = num * saoDangNum
	}
	return newMap
}

//获取扫荡掉落物品
func (t *XianFuDoubleRepairTemplate) GetSaodangRewardDropArr() []int32 {
	return t.saodangRewardDropArr
}

func (t *XianFuDoubleRepairTemplate) PatchAfterCheck() {
}

func (t *XianFuDoubleRepairTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.rewItemMap = make(map[int32]int32)
	if t.GetItemId != "" {
		itemArr, err := coreutils.SplitAsIntArray(t.GetItemId)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.GetItemId)
			return template.NewTemplateFieldError("GetItemId", err)
		}
		itemCountArr, err := coreutils.SplitAsIntArray(t.GetItemCount)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.GetItemCount)
			return template.NewTemplateFieldError("GetItemCount", err)
		}
		if len(itemArr) != len(itemCountArr) {
			err = fmt.Errorf("[%s] [%s] invalid", t.GetItemId, t.GetItemCount)
			return template.NewTemplateFieldError("GetItemId or GetItemCount", err)
		}
		for i := 0; i < len(itemArr); i++ {
			itemId := itemArr[i]
			itemCount := itemCountArr[i]
			tempItemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if tempItemTemplate == nil {
				err = fmt.Errorf("[%s] invalid", t.GetItemId)
				return template.NewTemplateFieldError("GetItemId", err)
			}
			//验证数量至少1
			err = validator.MinValidate(float64(itemCount), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%s] invalid", t.GetItemCount)
				return template.NewTemplateFieldError("GetItemCount", err)
			}
			t.rewItemMap[itemId] = itemCount
		}
	}

	//验证 map_id
	to := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		err = template.NewTemplateFieldError("MapId", err)
		return
	}
	t.majorMapTemplate = to.(*MapTemplate)

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

func (t *XianFuDoubleRepairTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	to := template.GetTemplateService().Get(int(t.BossId), (*BiologyTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.BossId)
		err = template.NewTemplateFieldError("BossId", err)
		return
	}
	biologyTemplate := to.(*BiologyTemplate)
	if biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeMonster {
		err = fmt.Errorf("[%d] invalid", t.BossId)
		err = template.NewTemplateFieldError("BossId", err)
		return
	}

	if t.majorMapTemplate.GetMapType() != scenetypes.SceneTypeMajor {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		err = template.NewTemplateFieldError("MapId", err)
		return
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

	return nil
}

func (t *XianFuDoubleRepairTemplate) FileName() string {
	return "tb_xianfu_doublerepair.json"
}

func init() {
	template.Register((*XianFuDoubleRepairTemplate)(nil))
}
