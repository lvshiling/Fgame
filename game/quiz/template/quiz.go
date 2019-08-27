package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sync"
)

//仙尊问答接口处理
type QuizTemplateService interface {
	//随机获取仙尊问答配置
	GetQuizTemplateRandom(excludeIdList []int32) *gametemplate.QuizTemplate
	//获取仙尊问答配置
	GetQuizByTemplateId(quizId int32) *gametemplate.QuizTemplate
	//获取仙尊问答常量配置
	GetQuizConstantTemplate() *gametemplate.QuizConstantTemplate
	//系统假答题生成间隔
	GetRandomAnswerChatTime() int64
}

type quizTemplateService struct {
	//仙尊问答map
	quizMap map[int32]*gametemplate.QuizTemplate
	//仙尊问答常量map
	quizConstant *gametemplate.QuizConstantTemplate
}

//初始化
func (cs *quizTemplateService) init() error {
	cs.quizMap = make(map[int32]*gametemplate.QuizTemplate)
	//仙尊问答
	templateMap := template.GetTemplateService().GetAll((*gametemplate.QuizTemplate)(nil))
	for _, templateObject := range templateMap {
		quizTemplate, _ := templateObject.(*gametemplate.QuizTemplate)
		cs.quizMap[int32(quizTemplate.TemplateId())] = quizTemplate
	}

	//仙尊问答常量表
	tempQuizConstantTemplate := template.GetTemplateService().Get(1, (*gametemplate.QuizConstantTemplate)(nil))
	if tempQuizConstantTemplate == nil {
		return fmt.Errorf("仙尊问答常量表不存在")
	}
	cs.quizConstant = tempQuizConstantTemplate.(*gametemplate.QuizConstantTemplate)

	return nil
}

//随机获取仙尊问答配置
func (cs *quizTemplateService) GetQuizTemplateRandom(excludeIdList []int32) *gametemplate.QuizTemplate {
	weights := make([]int64, 0, len(cs.quizMap)-1)
	tempQuizList := make([]*gametemplate.QuizTemplate, 0, len(cs.quizMap)-1)
	for _, ch := range cs.quizMap {
		isExclude := false
		for _, excludeId := range excludeIdList {
			if excludeId == int32(ch.TemplateId()) {
				isExclude = true
				break
			}
		}
		if isExclude {
			continue
		}
		weights = append(weights, int64(ch.QuanZhong))
		tempQuizList = append(tempQuizList, ch)
	}
	index := mathutils.RandomWeights(weights)
	if index == -1 {
		return nil
	}
	ch := tempQuizList[index]
	return ch
}

//获取仙尊问答配置
func (cs *quizTemplateService) GetQuizByTemplateId(quizId int32) *gametemplate.QuizTemplate {
	quiz, ok := cs.quizMap[quizId]
	if !ok {
		return nil
	}

	return quiz
}

//获取仙尊问答配置
func (cs *quizTemplateService) GetQuizConstantTemplate() *gametemplate.QuizConstantTemplate {
	return cs.quizConstant
}

//系统假答题生成间隔
func (cs *quizTemplateService) GetRandomAnswerChatTime() int64 {
	min := int(cs.GetQuizConstantTemplate().MsgTimeMin)
	max := int(cs.GetQuizConstantTemplate().MsgTimeMax)
	randTime := int64(mathutils.RandomRange(min, max))
	return randTime
}

var (
	once sync.Once
	cs   *quizTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &quizTemplateService{}
		err = cs.init()
	})
	return err
}

func GetQuizTemplateService() QuizTemplateService {
	return cs
}
