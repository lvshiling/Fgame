package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	mounttypes "fgame/fgame/game/mount/types"

	"fmt"
	"strconv"
)

//坐骑配置
type MountTemplate struct {
	*MountTemplateVO
	typ                    mounttypes.MountType           //坐骑类型
	magicParamMap          map[int32]string               //幻化条件
	magicParamXUMap        map[int32]int32                //幻化条件1,2
	magicParamIMap         map[int32]int32                //幻化条件3消耗物品
	useItemTemplate        *ItemTemplate                  //进阶物品
	battleAttrTemplate     *AttrTemplate                  //阶别属性
	mountUpstarTemplateMap map[int32]*MountUpstarTemplate //坐骑皮肤升星map
	mountUpstarTemplate    *MountUpstarTemplate           //坐骑皮肤升星
}

func (mt *MountTemplate) TemplateId() int {
	return mt.Id
}

func (mt *MountTemplate) GetTyp() mounttypes.MountType {
	return mt.typ
}

func (mt *MountTemplate) GetUseItemTemplate() *ItemTemplate {
	return mt.useItemTemplate
}

func (mt *MountTemplate) GetMagicParamIMap() map[int32]int32 {
	return mt.magicParamIMap
}

func (mt *MountTemplate) GetMagicParamXUMap() map[int32]int32 {
	return mt.magicParamXUMap
}

func (mt *MountTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return mt.battleAttrTemplate
}

func (mt *MountTemplate) GetIsClear() bool {
	return mt.IsClear != 0
}

func (mt *MountTemplate) GetMountUpstarByLevel(level int32) *MountUpstarTemplate {
	if v, ok := mt.mountUpstarTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (mt *MountTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//坐骑类型
	mt.typ = mounttypes.MountType(mt.Type)
	if !mt.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", mt.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//阶别attr属性
	if mt.Attr != 0 {
		to := template.GetTemplateService().Get(int(mt.Attr), (*AttrTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.Attr)
			return template.NewTemplateFieldError("Attr", err)
		}
		attrTemplate, _ := to.(*AttrTemplate)
		mt.battleAttrTemplate = attrTemplate
	}

	//幻化条件
	mt.magicParamMap = make(map[int32]string)
	mt.magicParamXUMap = make(map[int32]int32)
	mt.magicParamIMap = make(map[int32]int32)

	mt.magicParamMap[mt.MagicConditionType1] = mt.MagicConditionParameter1
	mt.magicParamMap[mt.MagicConditionType2] = mt.MagicConditionParameter2
	mt.magicParamMap[mt.MagicConditionType3] = mt.MagicConditionParameter3

	for condType, condParam := range mt.magicParamMap {
		cType := mounttypes.MountUCondType(condType)
		if !cType.Valid() {
			err = fmt.Errorf("[%d] invalid", condType)
			return template.NewTemplateFieldError("magic_condition_type", err)
		}
		switch cType {
		case mounttypes.MountUCondTypeX,
			mounttypes.MountUCondTypeU:
			num, err := strconv.ParseInt(condParam, 10, 32)
			if err != nil {
				return err
			}
			mt.magicParamXUMap[condType] = int32(num)
			break
		case mounttypes.MountUCondTypeI:
			itemArr, err := utils.SplitAsIntArray(condParam)
			if err != nil {
				return err
			}
			if len(itemArr) != 2 {
				err = fmt.Errorf("[%s] invalid", condParam)
				return template.NewTemplateFieldError("magic_condition_parameter", err)
			}
			mt.magicParamIMap[itemArr[0]] = itemArr[1]
			break
		default:
			break
		}
	}
	//幻化条件1、2
	for condType, condParam := range mt.magicParamXUMap {
		if condType == int32(mounttypes.MountUCondTypeX) {
			err = validator.MinValidate(float64(condParam), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", condParam)
				err = template.NewTemplateFieldError("MagicConditionParameter", err)
				return
			}
		}
	}
	//幻化条件3
	for item, num := range mt.magicParamIMap {
		itemTemplate := template.GetTemplateService().Get(int(item), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("item [%d] invalid", item)
			err = template.NewTemplateFieldError("MagicConditionParameter", err)
			return
		}
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("num [%d] invalid", num)
			return template.NewTemplateFieldError("MagicConditionParameter", err)
		}
	}

	//验证 UseItem
	if mt.UseItem != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(mt.UseItem), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", mt.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		mt.useItemTemplate = useItemTemplateVO.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(mt.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mt.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	//验证 mount_upgrade_begin_id
	if mt.MountUpstarBeginId != 0 {
		to := template.GetTemplateService().Get(int(mt.MountUpstarBeginId), (*MountUpstarTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.MountUpstarBeginId)
			return template.NewTemplateFieldError("MountUpstarBeginId", err)
		}

		mountUpstarTemplate, ok := to.(*MountUpstarTemplate)
		if !ok {
			return fmt.Errorf("MountUpgradeBeginId [%d] invalid", mt.MountUpstarBeginId)
		}

		mt.mountUpstarTemplate = mountUpstarTemplate
	}

	return nil
}

func (mt *MountTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if mt.NextId != 0 {

		to := template.GetTemplateService().Get(int(mt.NextId), (*MountTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp := to.(*MountTemplate)

		diff := nextTemp.Number - int32(mt.Number)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", mt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	//验证Number
	err = validator.MinValidate(float64(mt.Number), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.Number)
		err = template.NewTemplateFieldError("Number", err)
		return
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(mt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证use_money
	err = validator.MinValidate(float64(mt.UseMoney), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.UseMoney)
		err = template.NewTemplateFieldError("UseMoney", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(mt.TimesMin), float64(0), true, float64(mt.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mt.TimesMax), float64(mt.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mt.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(mt.AddMin), float64(0), true, float64(mt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(mt.AddMax), float64(mt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 ZhufuMax
	err = validator.MinValidate(float64(mt.ZhufuMax), float64(mt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 UseYinliang
	err = validator.MinValidate(float64(mt.UseYinliang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.UseYinliang)
		err = template.NewTemplateFieldError("UseYinliang", err)
		return
	}

	//验证 ShidanLimit
	err = validator.MinValidate(float64(mt.ShidanLimit), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.ShidanLimit)
		err = template.NewTemplateFieldError("ShidanLimit", err)
		return
	}

	//验证 CulturingDanLimit
	err = validator.MinValidate(float64(mt.CulturingDanLimit), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.CulturingDanLimit)
		err = template.NewTemplateFieldError("CulturingDanLimit", err)
		return
	}

	err = validator.MinValidate(float64(mt.IsClear), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.IsClear)
		return template.NewTemplateFieldError("IsClear", err)
	}

	return nil
}
func (mt *MountTemplate) PatchAfterCheck() {
	if mt.mountUpstarTemplate != nil {
		mt.mountUpstarTemplateMap = make(map[int32]*MountUpstarTemplate)
		//赋值mountUpstarTemplateMap
		for tempTemplate := mt.mountUpstarTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextMountUpstarTemplate {
			level := tempTemplate.Level
			mt.mountUpstarTemplateMap[level] = tempTemplate
		}
	}

}
func (mt *MountTemplate) FileName() string {
	return "tb_mount.json"
}

func init() {
	template.Register((*MountTemplate)(nil))
}
