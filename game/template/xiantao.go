package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	xiantaotypes "fgame/fgame/game/xiantao/types"
	"fmt"
)

//仙桃大会配置
type XianTaoTemplate struct {
	*XianTaoTemplateVO
	xianTaoRange randomGroup //提交仙桃数量区间
	rewItemMap   map[int32]int32
	xiantaoType  xiantaotypes.XianTaoType
	nextTemp     *XianTaoTemplate //下一条
}

func (mt *XianTaoTemplate) TemplateId() int {
	return mt.Id
}

func (mt *XianTaoTemplate) GetXianTaoType() xiantaotypes.XianTaoType {
	return mt.xiantaoType
}

func (mt *XianTaoTemplate) IsInXianTaoRange(count int32) bool {
	if count >= mt.xianTaoRange.min && count <= mt.xianTaoRange.max {
		return true
	}

	return false
}

func (mt *XianTaoTemplate) GetRewItemMap() map[int32]int32 {
	return mt.rewItemMap
}

func (mt *XianTaoTemplate) GetNextTemplate() *XianTaoTemplate {
	return mt.nextTemp
}

func (mt *XianTaoTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//下一条
	if mt.NextId != 0 {
		if mt.NextId-mt.Id != 1 {
			err = fmt.Errorf("[%d] invalid", mt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		to := template.GetTemplateService().Get(mt.NextId, (*XianTaoTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		mt.nextTemp = to.(*XianTaoTemplate)
	}

	// 提交数量区间
	xianTaoArr, err := utils.SplitAsIntArray(mt.XianTaoCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", mt.XianTaoCount)
		return template.NewTemplateFieldError("XianTaoCount", err)
	}
	if len(xianTaoArr) != 2 {
		err = fmt.Errorf("[%s] invalid", mt.XianTaoCount)
		return template.NewTemplateFieldError("XianTaoCount", err)
	}
	mt.xianTaoRange = randomGroup{
		min: xianTaoArr[0],
		max: xianTaoArr[1],
	}
	if mt.xianTaoRange.min > mt.xianTaoRange.max {
		err = fmt.Errorf("[%s] invalid", mt.XianTaoCount)
		return template.NewTemplateFieldError("XianTaoCount", err)
	}

	mt.rewItemMap = make(map[int32]int32)
	//验证 rew_item_id
	rewItemIdList, err := utils.SplitAsIntArray(mt.RewItemId)
	if err != nil {
		err = fmt.Errorf("[%s] split invalid", mt.RewItemId)
		return template.NewTemplateFieldError("RewItemId", err)
	}
	rewItemCountList, err := utils.SplitAsIntArray(mt.RewItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", mt.RewItemCount)
		return template.NewTemplateFieldError("RewItemCount", err)
	}
	if len(rewItemIdList) != len(rewItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid ", mt.RewItemId, mt.RewItemCount)
		err = template.NewTemplateFieldError("RewItemId or RewItemCount", err)
	}
	if len(rewItemIdList) > 0 {
		//组合数据
		for index, itemId := range rewItemIdList {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] nil invalid", mt.RewItemId)
				return template.NewTemplateFieldError("RewItemId", err)
			}

			err = validator.MinValidate(float64(rewItemCountList[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("RewItemCount", err)
			}

			_, ok := mt.rewItemMap[itemId]
			if !ok {
				mt.rewItemMap[itemId] = rewItemCountList[index]
			} else {
				mt.rewItemMap[itemId] += rewItemCountList[index]
			}
		}
	}

	return nil
}

func (mt *XianTaoTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if mt.nextTemp != nil {
		//区间连续校验
		if mt.nextTemp.xianTaoRange.min-mt.xianTaoRange.max != 1 {
			err = fmt.Errorf("[%d] invalid", mt.XianTaoCount)
			return template.NewTemplateFieldError("XianTaoCount", err)
		}
	}

	//类型
	mt.xiantaoType = xiantaotypes.XianTaoType(mt.Typ)
	if !mt.xiantaoType.Valid() {
		err = fmt.Errorf("[%d] invalid", mt.Typ)
		return template.NewTemplateFieldError("Typ", err)
	}

	//验证 rew_silver
	err = validator.MinValidate(float64(mt.RewSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.RewSilver)
		err = template.NewTemplateFieldError("RewSilver", err)
		return
	}

	//验证 rew_gold
	err = validator.MinValidate(float64(mt.RewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.RewGold)
		err = template.NewTemplateFieldError("RewGold", err)
		return
	}

	//验证 rew_bind_gold
	err = validator.MinValidate(float64(mt.RewBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.RewBindGold)
		err = template.NewTemplateFieldError("RewBindGold", err)
		return
	}

	//验证 rew_xp
	err = validator.MinValidate(float64(mt.RewExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.RewExp)
		err = template.NewTemplateFieldError("RewExp", err)
		return
	}

	//验证 rew_exp_point
	err = validator.MinValidate(float64(mt.RewExpPoint), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.RewExpPoint)
		err = template.NewTemplateFieldError("RewExpPoint", err)
		return
	}

	return nil
}

func (mt *XianTaoTemplate) PatchAfterCheck() {

}

func (mt *XianTaoTemplate) FileName() string {
	return "tb_xiantao.json"
}

func init() {
	template.Register((*XianTaoTemplate)(nil))
}
