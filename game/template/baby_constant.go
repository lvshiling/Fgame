package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"math"
)

//宝宝常量配置
type BabyConstantTemplate struct {
	*BabyConstantTemplateVO
	addTonicRange     randomGroup //补品增加值区间
	chaoshengGoldList []int32
	noticeTimeList    []int32
}

func (t *BabyConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *BabyConstantTemplate) GetBornNoticeTimeList() []int32 {
	return t.noticeTimeList
}

func (t *BabyConstantTemplate) GetRefreshTalentUseItemMap(refreshTimes int32) map[int32]int32 {
	if refreshTimes > t.XilianNum {
		refreshTimes = t.XilianNum
	}

	useItemMap := make(map[int32]int32)
	x := float64(refreshTimes + 1)
	y := float64(t.XiLianCoefficient) / float64(common.MAX_RATE)
	ratio := math.Pow(x, y)

	needItemNum := int32(math.Ceil(float64(t.XiLianItemCount)*ratio + float64(t.XiLianCoefficientFixed)))
	useItemMap[t.XiLianItemId] = needItemNum

	return useItemMap
}

func (t *BabyConstantTemplate) GetAddTonicPro() int32 {
	min := int(t.addTonicRange.min)
	max := int(t.addTonicRange.max + 1)
	randomNum := mathutils.RandomRange(min, max)
	return int32(randomNum)
}

func (t *BabyConstantTemplate) GetChaoShengGold(times int) int32 {
	maxIndex := len(t.chaoshengGoldList) - 1
	if times > maxIndex {
		return t.chaoshengGoldList[maxIndex]
	}

	return t.chaoshengGoldList[times]
}

func (t *BabyConstantTemplate) GetAccelerateNeedGold(now, pregnantTime int64) int64 {
	if now < pregnantTime {
		panic(fmt.Errorf("当前时间比怀孕时间早,now:%d,pregnant:%d", now, pregnantTime))
	}
	diff := t.BornTime - (now - pregnantTime)
	acclerateTime := int64(math.Ceil(float64(diff) / float64(common.MINUTE*10)))
	return int64(t.GoldZaoChan) * acclerateTime
}

func (t *BabyConstantTemplate) PatchAfterCheck() {}

func (t *BabyConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 补品值区间
	tonicArr, err := utils.SplitAsIntArray(t.BupinQujian)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.BupinQujian)
		return template.NewTemplateFieldError("BupinQujian", err)
	}
	if len(tonicArr) != 2 {
		err = fmt.Errorf("[%s] invalid", t.BupinQujian)
		return template.NewTemplateFieldError("BupinQujian", err)
	}

	t.addTonicRange = randomGroup{
		min: tonicArr[0],
		max: tonicArr[1],
	}

	// 超生消耗
	t.chaoshengGoldList, err = utils.SplitAsIntArray(t.GoldChaoSheng)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.GoldChaoSheng)
		return template.NewTemplateFieldError("GoldChaoSheng", err)
	}
	if len(t.chaoshengGoldList) == 0 {
		err = fmt.Errorf("[%s] invalid", t.GoldChaoSheng)
		return template.NewTemplateFieldError("GoldChaoSheng", err)
	}

	// 出生提示
	t.noticeTimeList, err = utils.SplitAsIntArray(t.TishiChushengTime)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.TishiChushengTime)
		return template.NewTemplateFieldError("TishiChushengTime", err)
	}

	return nil
}

func (t *BabyConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 字母河水
	to := template.GetTemplateService().Get(int(t.RiverItem), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.RiverItem)
		err = template.NewTemplateFieldError("RiverItem", err)
		return
	}

	// 补品
	to = template.GetTemplateService().Get(int(t.BupinItemId), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.BupinItemId)
		err = template.NewTemplateFieldError("BupinItemId", err)
		return
	}
	//数量
	err = validator.MinValidate(float64(t.BupinItemCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BupinItemCount)
		return template.NewTemplateFieldError("BupinItemCount", err)
	}
	//补品次数
	err = validator.MinValidate(float64(t.BupinMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BupinMax)
		return template.NewTemplateFieldError("BupinMax", err)
	}

	// 补品值随机区间
	if t.addTonicRange.min > t.addTonicRange.max {
		err = fmt.Errorf("[%s] invalid", t.BupinQujian)
		return template.NewTemplateFieldError("BupinQujian", err)
	}

	//出生时间
	err = validator.MinValidate(float64(t.BornTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BornTime)
		return template.NewTemplateFieldError("BornTime", err)
	}

	//加速
	err = validator.MinValidate(float64(t.GoldZaoChan), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GoldZaoChan)
		return template.NewTemplateFieldError("GoldZaoChan", err)
	}

	//宝宝基础数量
	err = validator.MinValidate(float64(t.BabyCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GoldZaoChan)
		return template.NewTemplateFieldError("GoldZaoChan", err)
	}

	//转世返还率
	err = validator.MinValidate(float64(t.ZsTianFuReturnRate), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZsTianFuReturnRate)
		return template.NewTemplateFieldError("ZsTianFuReturnRate", err)
	}

	// //超生消费
	// err = validator.MinValidate(float64(t.GoldChaoSheng), float64(0), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.GoldChaoSheng)
	// 	return template.NewTemplateFieldError("GoldChaoSheng", err)
	// }

	//改名卡
	to = template.GetTemplateService().Get(int(t.GaiMingKaId), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.GaiMingKaId)
		err = template.NewTemplateFieldError("GaiMingKaId", err)
		return
	}

	//天赋数量
	err = validator.MinValidate(float64(t.LimitSkiiNum), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LimitSkiiNum)
		return template.NewTemplateFieldError("LimitSkiiNum", err)
	}

	//s宝宝卡
	to = template.GetTemplateService().Get(int(t.BaoBaoCard), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.BaoBaoCard)
		err = template.NewTemplateFieldError("BaoBaoCard", err)
		return
	}
	//洗练消耗物品
	to = template.GetTemplateService().Get(int(t.XiLianItemId), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.XiLianItemId)
		err = template.NewTemplateFieldError("XiLianItemId", err)
		return
	}
	//洗练消耗数量
	err = validator.MinValidate(float64(t.XiLianItemCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XiLianItemCount)
		return template.NewTemplateFieldError("XiLianItemCount", err)
	}

	//洗练系数
	err = validator.MinValidate(float64(t.XiLianCoefficient), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XiLianCoefficient)
		return template.NewTemplateFieldError("XiLianCoefficient", err)
	}
	//洗练固定值
	err = validator.MinValidate(float64(t.XiLianCoefficientFixed), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XiLianCoefficientFixed)
		return template.NewTemplateFieldError("XiLianCoefficientFixed", err)
	}
	//成长值
	err = validator.MinValidate(float64(t.GrowthCoefficient), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GrowthCoefficient)
		return template.NewTemplateFieldError("GrowthCoefficient", err)
	}
	//洗练公式次数上限
	err = validator.MinValidate(float64(t.XilianNum), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XilianNum)
		return template.NewTemplateFieldError("XilianNum", err)
	}

	return nil
}

func (t *BabyConstantTemplate) FileName() string {
	return "tb_baobao_constant.json"
}

func init() {
	template.Register((*BabyConstantTemplate)(nil))
}
