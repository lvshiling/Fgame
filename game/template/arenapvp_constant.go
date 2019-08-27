package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

type ArenapvpConstantTemplate struct {
	*ArenapvpConstantTemplateVO
	arenaMapTemplate *MapTemplate
	luckyItemMap     map[int32]int32
}

func (t *ArenapvpConstantTemplate) GetLuckyItemMap() map[int32]int32 {
	return t.luckyItemMap
}

func (t *ArenapvpConstantTemplate) IsJiaRen() bool {
	return t.IsJiaren != 0
}

func (t *ArenapvpConstantTemplate) GetRandomRobotRatio() float64 {
	robotRatio := float64(mathutils.RandomRange(int(t.AttrRatioMin), int(t.AttrRatioMax+1)))
	return robotRatio
}

func (t *ArenapvpConstantTemplate) GetLuckyPlayerNumber(times int32) (realNum int32, robotNum int32) {
	hitList := mathutils.RandomHits(int(common.MAX_RATE), int(t.XinyunRobotRate), times)
	for _, hit := range hitList {
		if hit {
			robotNum += 1
		} else {
			realNum += 1
		}
	}

	return
}

func (t *ArenapvpConstantTemplate) GetArenaMapTemplate() *MapTemplate {
	return t.arenaMapTemplate
}

func (t *ArenapvpConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return
}
func (t *ArenapvpConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	err = validator.RangeValidate(float64(t.JifenBekillPercent), 0, false, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.JifenBekillPercent)
		return template.NewTemplateFieldError("JifenBekillPercent", err)
	}

	err = validator.MinValidate(float64(t.RobotAddTime), 0, false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RobotAddTime)
		return template.NewTemplateFieldError("RobotAddTime", err)
	}

	err = validator.MinValidate(float64(t.IsJiaren), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.IsJiaren)
		return template.NewTemplateFieldError("IsJiaren", err)
	}

	err = validator.MinValidate(float64(t.XingYunFirstTime), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XingYunFirstTime)
		return template.NewTemplateFieldError("XingYunFirstTime", err)
	}

	err = validator.MinValidate(float64(t.XingYunTime), 0, false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XingYunTime)
		return template.NewTemplateFieldError("XingYunTime", err)
	}

	err = validator.MinValidate(float64(t.XingYunPlayerCount), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XingYunPlayerCount)
		return template.NewTemplateFieldError("XingYunPlayerCount", err)
	}

	err = validator.RangeValidate(float64(t.XinyunRobotRate), 0, true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XinyunRobotRate)
		return template.NewTemplateFieldError("XinyunRobotRate", err)
	}

	to := template.GetTemplateService().Get(int(t.XingYunItemId), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.XingYunItemId)
		err = template.NewTemplateFieldError("XingYunItemId", err)
		return
	}

	//验证 AttrRatioMin
	err = validator.MinValidate(float64(t.AttrRatioMin), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AttrRatioMin)
		err = template.NewTemplateFieldError("AttrRatioMin", err)
		return
	}

	//验证 max_stack
	err = validator.MinValidate(float64(t.AttrRatioMax), float64(1), true)
	if err != nil || t.AttrRatioMax < t.AttrRatioMin {
		err = fmt.Errorf("[%d] invalid", t.AttrRatioMax)
		err = template.NewTemplateFieldError("AttrRatioMax", err)
		return
	}
	return
}

func (t *ArenapvpConstantTemplate) PatchAfterCheck() {
	t.luckyItemMap = make(map[int32]int32)
	if t.XingYunItemId > 0 {
		t.luckyItemMap[t.XingYunItemId] = 1
	}
}

func (t *ArenapvpConstantTemplate) TemplateId() int {
	return t.Id
}

func (at *ArenapvpConstantTemplate) FileName() string {
	return "tb_biwudahui.json"
}

func init() {
	template.Register((*ArenapvpConstantTemplate)(nil))
}
