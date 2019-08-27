package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"math"
	"math/rand"
)

//红包配置
type HongBaoTemplate struct {
	*HongBaoTemplateVO
	hongBaoType     itemtypes.ItemHongBaoSubType //红包类型
	useItemTemplate *ItemTemplate
	baoDiVal        int32
	bestMaxVal      int32
}

func (et *HongBaoTemplate) TemplateId() int {
	return et.Id
}

func (et *HongBaoTemplate) GetHongBaoType() itemtypes.ItemHongBaoSubType {
	return et.hongBaoType
}

func (et *HongBaoTemplate) CheckNeedCondition(vipLev, zhuanShu int32) bool {
	if vipLev >= et.NeedVipLevel || zhuanShu >= et.NeedZhuanShu {
		return true
	}
	return false
}

func (et *HongBaoTemplate) GetUseItemTemplate() *ItemTemplate {
	return et.useItemTemplate
}

func (et *HongBaoTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	//验证 need_item
	to := template.GetTemplateService().Get(int(et.ItemId), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", et.ItemId)
		err = template.NewTemplateFieldError("ItemId", err)
		return
	}
	et.useItemTemplate = to.(*ItemTemplate)

	return nil
}
func (et *HongBaoTemplate) PatchAfterCheck() {

}
func (et *HongBaoTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	//验证 hongBaoType
	et.hongBaoType = itemtypes.ItemHongBaoSubType(et.Typ)
	if !et.hongBaoType.Valid() {
		err = fmt.Errorf("[%d] invalid", et.Typ)
		err = template.NewTemplateFieldError("Typ", err)
		return
	}

	//验证 领取红包需要的VIP最低等级
	err = validator.MinValidate(float64(et.NeedVipLevel), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("NeedVipLevel", err)
		return
	}

	//验证 领取红包需要的转数
	err = validator.MinValidate(float64(et.NeedZhuanShu), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("NeedZhuanShu", err)
		return
	}

	//验证 奖励总量
	err = validator.MinValidate(float64(et.GeneralRew), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("GeneralRew", err)
		return
	}

	//验证 领取人数
	err = validator.RangeValidate(float64(et.CountMin), float64(0), false, float64(et.CountMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.CountMin)
		err = template.NewTemplateFieldError("CountMin", err)
		return
	}
	err = validator.MinValidate(float64(et.CountMax), float64(et.CountMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.CountMax)
		err = template.NewTemplateFieldError("CountMax", err)
		return
	}

	//验证 最佳手气最大值
	err = validator.MinValidate(float64(et.GoodProportionMax), float64(0), false)
	if err != nil {
		err = template.NewTemplateFieldError("GoodProportionMax", err)
		return
	}

	//验证 最佳手气最小值
	err = validator.MinValidate(float64(et.GoodProportionMin), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("GoodProportionMin", err)
		return
	}

	//验证 普通手气保底值
	err = validator.MinValidate(float64(et.ProportionMin), float64(0), false)
	if err != nil {
		err = template.NewTemplateFieldError("ProportionMin", err)
		return
	}

	//验证 珍稀红包 掉落id
	if et.hongBaoType == itemtypes.ItemHongBaoSubTypeZhenXi {
		tmpObj := template.GetTemplateService().Get(int(et.GoodProportionMax), (*DropTemplate)(nil))
		if tmpObj == nil {
			return template.NewTemplateFieldError("GoodProportionMax", fmt.Errorf("[%s] invalid", et.GoodProportionMax))
		}

		tmpObj = template.GetTemplateService().Get(int(et.ProportionMin), (*DropTemplate)(nil))
		if tmpObj == nil {
			return template.NewTemplateFieldError("ProportionMin", fmt.Errorf("[%s] invalid", et.ProportionMin))
		}
	} else {
		//验证 最佳手气最小值
		err = validator.RangeValidate(float64(et.GoodProportionMin), float64(0), true, float64(et.GoodProportionMax), true)
		if err != nil {
			err = template.NewTemplateFieldError("GoodProportionMin", err)
			return
		}

		//平均值应该小于最小手气
		if float64(common.MAX_RATE)/float64(et.CountMin) >= float64(et.GoodProportionMin) {
			return template.NewTemplateFieldError("GoodProportionMin", fmt.Errorf("[%d] 不大于平均值", et.ProportionMin))
		}

		//验证随机到最大有保底值
		et.bestMaxVal = int32(math.Ceil(float64(et.GoodProportionMax) / float64(common.MAX_RATE) * float64(et.GeneralRew)))
		et.baoDiVal = int32(math.Ceil(float64(et.ProportionMin) / float64(common.MAX_RATE) * float64(et.GeneralRew)))

		minVal := et.bestMaxVal + et.baoDiVal*(et.CountMax-1)
		if minVal > et.GeneralRew {
			return template.NewTemplateFieldError("ProportionMin", fmt.Errorf("[%d] 保底值太高或[%d]太高", et.ProportionMin, et.GoodProportionMax))
		}
	}

	return nil
}

//TODO 修改为通用算法
//获取奖励
func (edt *HongBaoTemplate) GetAwardNumList(countMax int32) []int32 {
	if edt.hongBaoType == itemtypes.ItemHongBaoSubTypeZhenXi {
		return nil
	}
	totalVal := int32(0)
	totalArr := make([]int32, 0, countMax)
	min := int(edt.GoodProportionMin)
	max := int(edt.GoodProportionMax)
	goodProportion := int32(mathutils.RandomRange(min, max))
	//最大值
	bestVal := int32(math.Ceil(float64(goodProportion) / float64(common.MAX_RATE) * float64(edt.GeneralRew)))
	totalArr = append(totalArr, bestVal)
	totalVal += bestVal
	leftVal := edt.GeneralRew - bestVal
	leftCount := countMax - 1

	baoDiVal := edt.baoDiVal

	for leftCount > 0 {
		maxVal := bestVal
		maxAssignVal := leftVal - leftCount*baoDiVal
		if maxAssignVal < maxVal {
			maxVal = maxAssignVal
		}
		//最大
		minVal := int32(math.Floor(float64(leftVal) / float64(leftCount)))

		elapse := maxVal - minVal
		randomVal := minVal
		if elapse > 0 {
			//随机最大
			randomVal = int32(rand.Intn(int(elapse))) + minVal
		}
		leftVal -= randomVal
		leftCount -= 1
		totalVal += randomVal
		totalArr = append(totalArr, randomVal)
	}

	if totalVal != edt.GeneralRew {
		panic(fmt.Errorf("红包得到的总值[%d],[%d]不相等", totalVal, edt.GeneralRew))
	}
	intArr := rand.Perm(len(totalArr))
	for k, v := range intArr {
		totalArr[k], totalArr[v] = totalArr[v], totalArr[k]
	}
	return totalArr
}

func (edt *HongBaoTemplate) FileName() string {
	return "tb_hongbao.json"
}

func init() {
	template.Register((*HongBaoTemplate)(nil))
}
