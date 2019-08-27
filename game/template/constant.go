package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	constanttypes "fgame/fgame/game/constant/types"
	itemtypes "fgame/fgame/game/item/types"
	"fmt"
)

type ConstantTemplate struct {
	*ConstantTemplateVO
	constantType constanttypes.ConstantType
}

func (ct *ConstantTemplate) GetConstantType() constanttypes.ConstantType {
	return ct.constantType
}

func (ct *ConstantTemplate) TemplateId() int {
	return ct.Id
}

//组合数据
func (ct *ConstantTemplate) Patch() (err error) {
	//统一处理错误方式
	defer func() {
		if err != nil {
			err = template.NewTemplateError(ct.FileName(), ct.TemplateId(), err)
			return
		}
	}()
	ct.constantType = constanttypes.ConstantType(ct.Type)

	return nil
}

//检验数值
func (ct *ConstantTemplate) Check() (err error) {
	//统一处理错误方式
	defer func() {
		if err != nil {
			err = template.NewTemplateError(ct.FileName(), ct.TemplateId(), err)
			return
		}
	}()
	switch ct.constantType {
	case constanttypes.ConstantTypeBagTotalNum:
		err = validator.MinValidate(float64(ct.Value), float64(0), false)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", ct.Value)
			return template.NewTemplateFieldError(ct.constantType.String(), err)
		}
		break
	case constanttypes.ConstantTypeBagDefaultOpenNum:
		templateConstantTemplate := template.GetTemplateService().Get(int(constanttypes.ConstantTypeBagTotalNum), (*ConstantTemplate)(nil))
		if templateConstantTemplate == nil {
			err = fmt.Errorf("背包总数量,没设置")
			return template.NewTemplateFieldError(ct.constantType.String(), err)
		}
		bagTotalNumConstantTemplate := templateConstantTemplate.(*ConstantTemplate)
		bagTotalNum := bagTotalNumConstantTemplate.Value
		err = validator.RangeValidate(float64(ct.Value), float64(0), false, float64(bagTotalNum), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", ct.Value)
			return template.NewTemplateFieldError(ct.constantType.String(), err)
		}
		break
	case constanttypes.ConstantTypeGoldForSingleSlot:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), false)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeSlotNumForSingleOpen:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), false)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypePVP:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), false)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeCrit:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), false)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeBlock:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), false)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeHarmPercentMax:
		//TODO
	case constanttypes.ConstantTypeCuthurtPercentMax:
		//TODO
	case constanttypes.ConstantTypeInitalHit:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), false)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeExitBattle:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), false)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeGemBagTotalNum:
		//TODO
	case constanttypes.ConstantTypeBornQuestId:
		{
			tempQuestTemplate := template.GetTemplateService().Get(int(ct.Value), (*QuestTemplate)(nil))
			if tempQuestTemplate == nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeReliveItemId:
		{
			tempQuestTemplate := template.GetTemplateService().Get(int(ct.Value), (*ItemTemplate)(nil))
			if tempQuestTemplate == nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeReliveNoNeedItemsBeforeLevel:
	case constanttypes.ConstantTypeReliveTimesClearTime:
	case constanttypes.ConstantTypeFirstReliveItemNum:
	case constanttypes.ConstantTypeItemNumAddEveryRelive:
	case constanttypes.ConstantTypeInitMoveSpeed:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), false)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeSkillLimit:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), false)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}

	case constanttypes.ConstantTypeChangeSceneCostSilver:
	case constanttypes.ConstantTypeChangeSceneCostItem:
		{
			if ct.Value != 0 {
				tempChangeSceneItem := template.GetTemplateService().Get(int(ct.Value), (*ItemTemplate)(nil))
				if tempChangeSceneItem == nil {
					err = fmt.Errorf("[%d] invalid", ct.Value)
					return template.NewTemplateFieldError(ct.constantType.String(), err)
				}
			}
			break
		}
	case constanttypes.ConstantTypeChangeSceneCostItemNum:
	case constanttypes.ConstantTypeGiftId:
		{
			tempGiftTemplate := template.GetTemplateService().Get(int(ct.Value), (*ItemTemplate)(nil))
			if tempGiftTemplate == nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			giftTemplate := tempGiftTemplate.(*ItemTemplate)
			if giftTemplate.GetItemType() != itemtypes.ItemTypeXianHua {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}

	case constanttypes.ConstantTypeAllianceHuangGongOccupyTime:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeAllianceHuangGongOccupyFlagTime:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeAlliancePropertyPerHuFu:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeProtectBuff,
		constanttypes.ConstantTypeReliveProtectBuff:
		{
			buffTemplate := template.GetTemplateService().Get(int(ct.Value), (*BuffTemplate)(nil))
			if buffTemplate == nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeSecretCardCostGold:
		{
			err = validator.MinValidate(float64(ct.Value), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeResurrectionDanLimit:
		{
			err = validator.MinValidate(float64(ct.Value), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeNewGift:
		{
			tempGiftTemplate := template.GetTemplateService().Get(int(ct.Value), (*ItemTemplate)(nil))
			if tempGiftTemplate == nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypePKProtectBuff:
		{
			tempGiftTemplate := template.GetTemplateService().Get(int(ct.Value), (*BuffTemplate)(nil))
			if tempGiftTemplate == nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeHongBaoKeepTime:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeDropOwnerBuffId:
		{
			tempDropBuffTemplate := template.GetTemplateService().Get(int(ct.Value), (*BuffTemplate)(nil))
			if tempDropBuffTemplate == nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
			break
		}
	case constanttypes.ConstantTypeWushuangEssence:
		{
			itemTemplate := template.GetTemplateService().Get(int(ct.Value), (*ItemTemplate)(nil))
			itemTemp, _ := itemTemplate.(*ItemTemplate)
			if itemTemp == nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			} else if itemTemp.GetItemType() != itemtypes.ItemTypeWushuangWeaponEssence || itemTemp.GetItemSubType() != itemtypes.ItemWushuangWeaponEssenceSubTypeEssence {
				err = fmt.Errorf("[%d] invalid,Wrong WushuangWeaponEssence item", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
		}
	case constanttypes.ConstantTypeShengShouItemId:
		{
			itemTemplate := template.GetTemplateService().Get(int(ct.Value), (*ItemTemplate)(nil))
			itemTemp, _ := itemTemplate.(*ItemTemplate)
			if itemTemp == nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
		}
	case constanttypes.ConstantTypeShengShouItemNum:
		{
			err = validator.MinValidate(float64(ct.Value), float64(0), false)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", ct.Value)
				return template.NewTemplateFieldError(ct.constantType.String(), err)
			}
		}
	}

	return
}

func (ct *ConstantTemplate) PatchAfterCheck() {

}

func (ct *ConstantTemplate) FileName() string {
	return "tb_constant.json"
}

func init() {
	template.Register((*ConstantTemplate)(nil))
}
