package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	quiztypes "fgame/fgame/game/quiz/types"
	"fmt"
)

//仙尊问答配置
type QuizTemplate struct {
	*QuizTemplateVO
	rewData         *propertytypes.RewData
	rewErrorData    *propertytypes.RewData
	rewItemMap      map[int32]int32
	errorRewItemMap map[int32]int32
	rightAnswer     quiztypes.QuizAnswerType
	answerStrMap    map[quiztypes.QuizAnswerType]string
}

func (et *QuizTemplate) TemplateId() int {
	return et.Id
}

func (et *QuizTemplate) GetRightAnswer() quiztypes.QuizAnswerType {
	return et.rightAnswer
}

func (et *QuizTemplate) GetRewData() *propertytypes.RewData {
	return et.rewData
}

func (et *QuizTemplate) GetRewErrorData() *propertytypes.RewData {
	return et.rewErrorData
}

func (et *QuizTemplate) GetRewItemMap() map[int32]int32 {
	return et.rewItemMap
}

func (et *QuizTemplate) GetErrorRewItemMap() map[int32]int32 {
	return et.errorRewItemMap
}

func (et *QuizTemplate) GetAnswerStrByType(typ quiztypes.QuizAnswerType) string {
	return et.answerStrMap[typ]
}

func (et *QuizTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	//答案
	et.rightAnswer = quiztypes.QuizAnswerType(et.RightAnswer)

	//答案内容
	et.answerStrMap = make(map[quiztypes.QuizAnswerType]string)
	et.answerStrMap[quiztypes.QuizAnswerTypeA] = et.AnswerA
	et.answerStrMap[quiztypes.QuizAnswerTypeB] = et.AnswerB
	et.answerStrMap[quiztypes.QuizAnswerTypeC] = et.AnswerC
	et.answerStrMap[quiztypes.QuizAnswerTypeD] = et.AnswerD

	//答对奖励资源
	if et.RewExp > 0 || et.RewExpPoint > 0 || et.RewardSilver > 0 || et.RewardBindGold > 0 || et.RewardGold > 0 {
		et.rewData = propertytypes.CreateRewData(et.RewExp, et.RewExpPoint, et.RewardSilver, et.RewardBindGold, et.RewardGold)
	}

	//答错奖励资源
	if et.ErrorRewExp > 0 || et.ErrorRewExpPoint > 0 || et.ErrorRewardSilver > 0 || et.ErrorRewardBindGold > 0 || et.ErrorRewardGold > 0 {
		et.rewErrorData = propertytypes.CreateRewData(et.ErrorRewExp, et.ErrorRewExpPoint, et.ErrorRewardSilver, et.ErrorRewardBindGold, et.ErrorRewardGold)
	}

	//答对奖励物品
	et.rewItemMap = make(map[int32]int32)
	rewItemIdList, err := utils.SplitAsIntArray(et.RewItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", et.RewItemId)
		return template.NewTemplateFieldError("RewItemId", err)
	}
	rewItemCountList, err := utils.SplitAsIntArray(et.RewItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", et.RewItemCount)
		return template.NewTemplateFieldError("RewItemCount", err)
	}
	if len(rewItemIdList) > 0 {
		if len(rewItemIdList) == len(rewItemCountList) {
			//组合数据
			for index, itemId := range rewItemIdList {
				_, ok := et.rewItemMap[itemId]
				if ok {
					et.rewItemMap[itemId] += rewItemCountList[index]
				} else {
					et.rewItemMap[itemId] = rewItemCountList[index]
				}
			}
		} else {
			err = fmt.Errorf("[%s] [%s] len invalid", et.RewItemId, et.RewItemCount)
			return template.NewTemplateFieldError("RewItemId,RewItemCount", err)
		}
	}

	//答错奖励物品
	et.errorRewItemMap = make(map[int32]int32)
	errorRewItemIdList, err := utils.SplitAsIntArray(et.ErrorRewItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", et.ErrorRewItemId)
		return template.NewTemplateFieldError("ErrorRewItemId", err)
	}
	errorRewItemCountList, err := utils.SplitAsIntArray(et.ErrorRewItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", et.ErrorRewItemCount)
		return template.NewTemplateFieldError("ErrorRewItemCount", err)
	}
	if len(errorRewItemIdList) > 0 {
		if len(errorRewItemIdList) == len(errorRewItemCountList) {
			//组合数据
			for index, itemId := range errorRewItemIdList {
				_, ok := et.errorRewItemMap[itemId]
				if ok {
					et.errorRewItemMap[itemId] += errorRewItemCountList[index]
				} else {
					et.errorRewItemMap[itemId] = errorRewItemCountList[index]
				}
			}
		} else {
			err = fmt.Errorf("[%s] [%s] len invalid", et.ErrorRewItemId, et.ErrorRewItemCount)
			return template.NewTemplateFieldError("ErrorRewItemId,ErrorRewItemCount", err)
		}
	}

	return nil
}
func (et *QuizTemplate) PatchAfterCheck() {

}
func (et *QuizTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	//答案
	if !et.rightAnswer.Valid() {
		err = fmt.Errorf("[%d] invalid", et.RightAnswer)
		return template.NewTemplateFieldError("RightAnswer", err)
	}

	//出题概率（权重）
	err = validator.RangeValidate(float64(et.QuanZhong), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		return template.NewTemplateFieldError("QuanZhong", err)
	}

	//答对奖励rewData
	if !et.rewData.Valid() {
		return template.NewTemplateFieldError("RewExp, RewExpPoint, RewardSilver, RewardBindGold, RewardGold", err)
	}

	//答对奖励物品的概率
	err = validator.RangeValidate(float64(et.RewardItemRate), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		return template.NewTemplateFieldError("RewardItemRate", err)
	}

	//答错奖励rewErrorData
	if !et.rewErrorData.Valid() {
		return template.NewTemplateFieldError("ErrorRewExp, ErrorRewExpPoint, ErrorRewardSilver, ErrorRewardBindGold, ErrorRewardGold", err)
	}

	//答对奖励物品的概率
	err = validator.RangeValidate(float64(et.ErrorRewardItemRate), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		return template.NewTemplateFieldError("ErrorRewardItemRate", err)
	}

	// 答对奖励物品id
	for itemId, num := range et.rewItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("RewItemId", fmt.Errorf("[%d] invalid", itemId))
		}
		if err = validator.MinValidate(float64(num), float64(1), true); err != nil {
			err = template.NewTemplateFieldError("RewItemCount", err)
			return
		}
	}

	// 答错奖励物品id
	for itemId, num := range et.errorRewItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("ErrorRewItemId", fmt.Errorf("[%d] invalid", itemId))
		}
		if err = validator.MinValidate(float64(num), float64(1), true); err != nil {
			err = template.NewTemplateFieldError("ErrorRewItemCount", err)
			return
		}
	}

	return nil
}

func (edt *QuizTemplate) FileName() string {
	return "tb_quiz.json"
}

func init() {
	template.Register((*QuizTemplate)(nil))
}
