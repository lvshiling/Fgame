package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//元宝猪
type GroupTemplateFeedbackGoldPig struct {
	*welfaretemplate.GroupTemplateBase                                         //
	startGoldPigTemplate               *gametemplate.GoldPigTemplate           //起始模板
	goldPigTempMap                     map[int32]*gametemplate.GoldPigTemplate //模板集合
}

func (gt *GroupTemplateFeedbackGoldPig) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	for _, t := range gt.GetOpenTempMap() {
		//元宝猪起始配置
		tempObj := template.GetTemplateService().Get(int(t.Value1), (*gametemplate.GoldPigTemplate)(nil))
		if tempObj == nil {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}
		gt.startGoldPigTemplate = tempObj.(*gametemplate.GoldPigTemplate)
	}

	//加载元宝猪配置
	gt.goldPigTempMap = make(map[int32]*gametemplate.GoldPigTemplate)
	for startTemp := gt.startGoldPigTemplate; startTemp != nil; startTemp = startTemp.GetNextTemplate() {
		gt.goldPigTempMap[startTemp.NeedRecharge] = startTemp
	}

	return
}

//计算元宝猪档次条件
func (gt *GroupTemplateFeedbackGoldPig) CountGoldPigNeedCharge(curConditon, chargeGold int32) (newCondition int32) {
	temp := gt.GetFirstOpenTemp()
	if temp == nil {
		return
	}

	for _, temp := range gt.goldPigTempMap {
		if chargeGold < temp.NeedRecharge {
			continue
		}
		if newCondition >= temp.NeedRecharge {
			continue
		}

		newCondition = temp.NeedRecharge
	}

	return
}

//元宝猪返利比例
func (gt *GroupTemplateFeedbackGoldPig) GetGoldPigReturnRate(condition int32) int32 {
	rate := int32(0)
	goldPigTemp, ok := gt.goldPigTempMap[condition]
	if ok {
		rate = goldPigTemp.ReturnPercent
	}

	return rate
}

func CreateGroupTemplateFeedbackGoldPig(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateFeedbackGoldPig{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeGoldPig, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateFeedbackGoldPig))
}
