package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	marrytypes "fgame/fgame/game/marry/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//结婚常量配置
type MarryTemplate struct {
	*MarryTemplateVO
	marryMapTemplate *MapTemplate
	carMapTemplate   *MapTemplate
	jiNianSjItemMap  map[int32]int32 //结婚纪念赠送时装物品
	qinMiDuMap       map[marrytypes.MarryHoutaiType]int32
}

func (mt *MarryTemplate) TemplateId() int {
	return mt.Id
}

func (mt *MarryTemplate) GetMarryMapTemplate() *MapTemplate {
	return mt.marryMapTemplate
}

func (mt *MarryTemplate) GetCarMapTemplate() *MapTemplate {
	return mt.carMapTemplate
}

func (mt *MarryTemplate) GetQinMiDuByVersion(version marrytypes.MarryHoutaiType) (qinMiDu int32, flag bool) {
	qinMiDu, ok := mt.qinMiDuMap[version]
	if !ok {
		return
	}
	flag = true
	return
}

func (mt *MarryTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//验证 marry_map_id
	to := template.GetTemplateService().Get(int(mt.MarryMapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", mt.MarryMapId)
		err = template.NewTemplateFieldError("MarryMapId", err)
		return
	}
	mt.marryMapTemplate = to.(*MapTemplate)

	//验证 car_map_id
	to = template.GetTemplateService().Get(int(mt.CarMapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", mt.CarMapId)
		err = template.NewTemplateFieldError("CarMapId", err)
		return
	}
	mt.carMapTemplate = to.(*MapTemplate)

	//纪念时装物品填充校验
	mt.jiNianSjItemMap = make(map[int32]int32)
	spilitJiNianSjItemArray, err := utils.SplitAsIntArray(mt.MarryItem)
	if err != nil {
		err = template.NewTemplateFieldError("MarryItem", err)
		return
	}
	spilitJiNianSjItemCountArray, err := utils.SplitAsIntArray(mt.MarryItemCount)
	if err != nil {
		err = template.NewTemplateFieldError("MarryItemCount", err)
		return
	}
	if len(spilitJiNianSjItemArray) != len(spilitJiNianSjItemCountArray) {
		err = fmt.Errorf("MarryItem [%s],MarryItemCount [%s] 数量不一样", mt.MarryItem, mt.MarryItemCount)
		err = template.NewTemplateFieldError("MarryItem,MarryItemCount", err)
		return
	}
	for index, value := range spilitJiNianSjItemArray {
		item := template.GetTemplateService().Get(int(value), (*ItemTemplate)(nil))
		if item == nil {
			err = fmt.Errorf("MarryItem [%d]无效", value)
			err = template.NewTemplateFieldError("MarryItem", err)
			return err
		}
		mt.jiNianSjItemMap[value] = spilitJiNianSjItemCountArray[index]
	}

	// 亲密度校验
	mt.qinMiDuMap = make(map[marrytypes.MarryHoutaiType]int32)
	spilitQinmiduHoutaiType, err := utils.SplitAsIntArray(mt.QinmiduHoutaiType)
	if err != nil {
		err = template.NewTemplateFieldError("QinmiduHoutaiType", err)
		return
	}
	spilitMarryQinmidu, err := utils.SplitAsIntArray(mt.MarryQinmidu)
	if err != nil {
		err = template.NewTemplateFieldError("MarryQinmidu", err)
		return
	}
	if len(spilitQinmiduHoutaiType) != len(spilitMarryQinmidu) {
		err = fmt.Errorf("QinmiduHoutaiType [%s],MarryQinmidu [%s] 长度不匹配", mt.QinmiduHoutaiType, mt.MarryQinmidu)
		err = template.NewTemplateFieldError("QinmiduHoutaiType,MarryQinmidu", err)
		return
	}
	for index, value := range spilitQinmiduHoutaiType {
		houTaiType := marrytypes.MarryHoutaiType(value)
		if !houTaiType.Valid() {
			err = template.NewTemplateFieldError("QinmiduHoutaiType", err)
			return
		}
		mt.qinMiDuMap[houTaiType] = spilitMarryQinmidu[index]
	}
	return nil
}

func (mt *MarryTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//验证 divorce_qinmidu
	err = validator.RangeValidate(float64(mt.DivorceQinmidu), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.DivorceQinmidu)
		err = template.NewTemplateFieldError("DivorceQinmidu", err)
		return
	}

	//验证 ring_level_gap
	err = validator.MinValidate(float64(mt.RingLevelGap), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.RingLevelGap)
		err = template.NewTemplateFieldError("RingLevelGap", err)
		return
	}

	//验证 marry_first_time
	if mt.MarryFirstTime == "" {
		err = fmt.Errorf("[%s] invalid", mt.MarryFirstTime)
		err = template.NewTemplateFieldError("MarryFirstTime", err)
		return
	}

	//验证 marry_amount
	err = validator.MinValidate(float64(mt.MarryAmount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.MarryAmount)
		err = template.NewTemplateFieldError("MarryAmount", err)
		return
	}

	to := template.GetTemplateService().Get(int(mt.CarId), (*BiologyTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", mt.CarId)
		err = template.NewTemplateFieldError("CarId", err)
		return
	}
	biologyTemplate := to.(*BiologyTemplate)
	if biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeWeddingCar {
		err = fmt.Errorf("[%d] invalid", mt.CarId)
		err = template.NewTemplateFieldError("CarId", err)
		return
	}

	if mt.marryMapTemplate.GetMapType() != scenetypes.SceneTypeMarry {
		err = fmt.Errorf("[%d] invalid", mt.MarryMapId)
		err = template.NewTemplateFieldError("MarryMapId", err)
		return
	}

	err = validator.MinValidate(float64(mt.MarryTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.MarryTime)
		err = template.NewTemplateFieldError("MarryTime", err)
		return
	}

	err = validator.MinValidate(float64(mt.QingChangTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.QingChangTime)
		err = template.NewTemplateFieldError("QingChangTime", err)
		return
	}

	return nil
}
func (mt *MarryTemplate) PatchAfterCheck() {

}
func (mt *MarryTemplate) FileName() string {
	return "tb_marry.json"
}

func (mt *MarryTemplate) GetJiNianSjItemMap() map[int32]int32 {
	return mt.jiNianSjItemMap
}

func init() {
	template.Register((*MarryTemplate)(nil))
}
