package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	housetypes "fgame/fgame/game/house/types"
	"fmt"
)

//房子常量配置
type HouseConstantTemplate struct {
	*HouseConstantTemplateVO
	logTimeRange  *randomGroup
	initHouseType housetypes.HouseType
}

func (t *HouseConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *HouseConstantTemplate) GetDummyLogTime() (int, int) {
	return int(t.logTimeRange.min), int(t.logTimeRange.max)
}

func (t *HouseConstantTemplate) GetInitHouseType() housetypes.HouseType {
	return t.initHouseType
}

func (t *HouseConstantTemplate) PatchAfterCheck() {}

func (t *HouseConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	dummyLogTime, err := utils.SplitAsIntArray(t.JiaRiZhiTime)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.JiaRiZhiTime)
		return template.NewTemplateFieldError("JiaRiZhiTime", err)
	}
	if len(dummyLogTime) != 2 {
		err = fmt.Errorf("[%s] invalid", t.JiaRiZhiTime)
		return template.NewTemplateFieldError("JiaRiZhiTime", err)
	}
	t.logTimeRange = &randomGroup{
		min: dummyLogTime[0],
		max: dummyLogTime[1],
	}

	return nil
}

func (t *HouseConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	if t.logTimeRange.min > t.logTimeRange.max {
		err = fmt.Errorf("[%s] invalid", t.JiaRiZhiTime)
		return template.NewTemplateFieldError("JiaRiZhiTime", err)
	}

	//损坏CD
	err = validator.MinValidate(float64(t.BrokenCd), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BrokenCd)
		return template.NewTemplateFieldError("BrokenCd", err)
	}

	//装修次数
	err = validator.MinValidate(float64(t.UplevLimitCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UplevLimitCount)
		return template.NewTemplateFieldError("UplevLimitCount", err)
	}

	//初始类型
	t.initHouseType = housetypes.HouseType(t.FirstFangZiType)
	if !t.initHouseType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.FirstFangZiType)
		return template.NewTemplateFieldError("FirstFangZiType", err)
	}

	//最大房子数量
	err = validator.MinValidate(float64(t.FangZiCountMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FangZiCountMax)
		return template.NewTemplateFieldError("FangZiCountMax", err)
	}

	return nil
}

func (t *HouseConstantTemplate) FileName() string {
	return "tb_fangzi_constant.json"
}

func init() {
	template.Register((*HouseConstantTemplate)(nil))
}
