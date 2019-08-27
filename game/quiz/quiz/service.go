package quiz

import (
	"fgame/fgame/core/runner"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/quiz/dao"
	quizeventtypes "fgame/fgame/game/quiz/event/types"
	quiztemplate "fgame/fgame/game/quiz/template"
	quiztypes "fgame/fgame/game/quiz/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
)

const (
	quizAnswerDisorderRate int64 = 2500 //打乱答案的均分权重
)

//仙尊问答接口处理
type QuizService interface {
	runner.Task
	//获取仙尊问答
	GetQuizObj() *QuizObject
	//检测仙尊问答答题时间
	CheckQuizAnswerTime() bool

	//添加答题玩家
	AddAnswerPlayer(playerId int64)
	//获取答题玩家
	IsAnswerPlayerById(playerId int64) bool

	//Gm系统出题
	GmAssignQuiz(quizId int32) error
}

type quizService struct {
	rwm sync.RWMutex
	//仙尊问答对象
	quizObj *QuizObject
	//做过当前题目的玩家
	didPlayerList []int64
	//出过的题目
	didQuizList []int32
	//添加假答题信息最新时间
	lastAddDummyAnswerChatTime int64
}

//获取仙尊问答
func (s *quizService) GetQuizObj() *QuizObject {
	return s.quizObj
}

//TODO cjb:可能答的是上一题 但是记录到下一题了
//添加答题玩家
func (s *quizService) AddAnswerPlayer(playerId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.didPlayerList = append(s.didPlayerList, playerId)
	return
}

//TODO 命名规范点 Get就不要返回bool
//获取答题玩家
func (s *quizService) IsAnswerPlayerById(playerId int64) bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	for _, val := range s.didPlayerList {
		if val == playerId {
			return true
		}
	}
	return false
}

//清理答题玩家
func (s *quizService) clearAnswerPlayers() {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	//TODO 直接赋值空
	s.didPlayerList = nil
}

//清理题目
func (s *quizService) clearDidQuizs() {
	s.didQuizList = nil
}

//刷新仙尊问答
func (s *quizService) refreshQuizs() error {
	quizConstantTemplate := quiztemplate.GetQuizTemplateService().GetQuizConstantTemplate()
	now := global.GetGame().GetTimeService().Now()
	begin, err := quizConstantTemplate.GetBeginTime(now)
	if err != nil {
		err = fmt.Errorf("quiz:仙尊问答，获取活动开始时间错误")
		return err
	}

	if s.quizObj.lastQuizTime < begin {
		s.clearDidQuizs()
	}
	s.clearAnswerPlayers()
	return nil
}

//初始化
func (s *quizService) init() (err error) {
	serverId := global.GetGame().GetServerIndex()
	quizEntity, err := dao.GetQuizDao().GetQuizEntity(serverId)
	if err != nil {
		return err
	}
	if quizEntity == nil {
		s.initQuizObject()
	} else {
		s.quizObj = NewQuizObject()
		err = s.quizObj.FromEntity(quizEntity)
		if err != nil {
			return err
		}
	}
	now := global.GetGame().GetTimeService().Now()
	if s.GetQuizObj().lastQuizTime > now {
		err := s.assignQuiz()
		if err != nil {
			return err
		}
	}
	return
}

//第一次初始化
func (es *quizService) initQuizObject() {
	peo := NewQuizObject()
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	peo.id = id
	peo.serverId = global.GetGame().GetServerIndex()
	peo.createTime = now
	es.quizObj = peo
	peo.SetModified()
}

//TODO 修改为n秒加一次
//心跳
func (s *quizService) Heartbeat() {
	if s.CheckQuizAnswerTime() {
		//答题时间就不要刷题了
		s.addQuizAnswerChat()
		return
	}

	if !s.checkQuizTime() {
		return
	}

	err := s.assignQuiz()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("quiz:系统题目,错误")
		return
	}
}

//系统出题
func (s *quizService) assignQuiz() (err error) {
	ok := s.refreshQuizs()
	if ok != nil {
		return
	}
	quizTemplate := quiztemplate.GetQuizTemplateService().GetQuizTemplateRandom(s.didQuizList)
	if quizTemplate == nil {
		return
	}

	var answerOrderList []quiztypes.QuizAnswerType
	answerOrderList = s.disorderAnswer(quizTemplate.GetRightAnswer())
	s.didQuizList = append(s.didQuizList, int32(quizTemplate.TemplateId()))

	now := global.GetGame().GetTimeService().Now()
	s.quizObj.lastQuizTime = now
	s.quizObj.lastQuizId = int32(quizTemplate.TemplateId())
	s.quizObj.answerList = answerOrderList
	s.quizObj.updateTime = now

	s.quizObj.SetModified()
	//发布题目事件
	data := quizeventtypes.CreateQuizAssignEventData(s.quizObj.lastQuizId, s.quizObj.answerList)
	gameevent.Emit(quizeventtypes.EventTypeAssignQuiz, nil, data)
	return
}

//系统生成假答题消息
func (s *quizService) addQuizAnswerChat() {
	now := global.GetGame().GetTimeService().Now()
	lastTime := s.lastAddDummyAnswerChatTime
	diffTime := now - lastTime
	randTime := quiztemplate.GetQuizTemplateService().GetRandomAnswerChatTime()
	if diffTime < randTime {
		return
	}

	s.lastAddDummyAnswerChatTime = now
	//答题消息事件
	answerType := s.getRandomAnswerType()
	data := quizeventtypes.CreateQuizAnswerChatEventData(int64(0), answerType)
	gameevent.Emit(quizeventtypes.EventTypeQuizAnswerChat, nil, data)
	return
}

//是否答题时间内
func (s *quizService) CheckQuizAnswerTime() bool {
	now := global.GetGame().GetTimeService().Now()
	quizAnswerLimitTime := quiztemplate.GetQuizTemplateService().GetQuizConstantTemplate().DaTiTime
	if now > s.quizObj.GetLastQuizTime()+int64(quizAnswerLimitTime) {
		return false
	}
	return true
}

//检测是否出题时间
func (s *quizService) checkQuizTime() bool {
	quizConstantTemplate := quiztemplate.GetQuizTemplateService().GetQuizConstantTemplate()
	now := global.GetGame().GetTimeService().Now()
	begin, err := quizConstantTemplate.GetBeginTime(now)
	if err != nil {
		panic(errors.Wrap(err, "quiz:仙尊问答，获取活动开始时间错误"))
	}
	end, err := quizConstantTemplate.GetEndTime(now)
	if err != nil {
		panic(errors.Wrap(err, "quiz:仙尊问答，获取活动结束时间错误"))
	}

	if !(now >= begin && now <= end) {
		return false
	}

	lastQuizTime := s.quizObj.GetLastQuizTime()
	//定点刷题
	nearRefreshTime, err := quizConstantTemplate.GetNearRefreshTime(now)
	if err != nil {
		panic(errors.Wrap(err, "quiz:仙尊问答，获取活动定点时间错误"))
	}
	if nearRefreshTime > 0 && lastQuizTime < nearRefreshTime {
		return true
	}

	//cd刷题
	intervalTime := quizConstantTemplate.IntervalTime
	if lastQuizTime+int64(intervalTime) > now {
		return false
	}

	return true
}

//打乱顺序
func (s *quizService) disorderAnswer(rightAnswer quiztypes.QuizAnswerType) []quiztypes.QuizAnswerType {
	var weights []int64
	var answerList []quiztypes.QuizAnswerType
	for i := int32(0); i < int32(rightAnswer); i++ {
		weights = append(weights, quizAnswerDisorderRate)
		answerList = append(answerList, quiztypes.QuizAnswerType(i+1))
	}

	var answerOrderList []quiztypes.QuizAnswerType
	orderIdxList := mathutils.RandomListFromWeights(weights, int32(rightAnswer))
	for j := 0; j < len(orderIdxList); j++ {
		answerOrderList = append(answerOrderList, answerList[orderIdxList[j]])
	}
	return answerOrderList
}

//系统随机假答题选项
func (s *quizService) getRandomAnswerType() quiztypes.QuizAnswerType {
	curTemp := quiztemplate.GetQuizTemplateService().GetQuizByTemplateId(s.quizObj.GetLastQuizId())
	min := int32(quiztypes.QuizAnswerTypeA)
	max := int32(curTemp.GetRightAnswer())
	randIdx := int32(mathutils.RandomRange(int(min), int(max)))
	randType := quiztypes.QuizAnswerType(randIdx)
	return randType
}

//Gm系统出题
func (s *quizService) GmAssignQuiz(quizId int32) (err error) {
	ok := s.refreshQuizs()
	if ok != nil {
		return
	}
	quizTemplate := quiztemplate.GetQuizTemplateService().GetQuizByTemplateId(quizId)
	if quizTemplate == nil {
		quizTemplate = quiztemplate.GetQuizTemplateService().GetQuizTemplateRandom(s.didQuizList)
		if quizTemplate == nil {
			return
		}
	}

	var answerOrderList []quiztypes.QuizAnswerType
	answerOrderList = s.disorderAnswer(quizTemplate.GetRightAnswer())
	s.didQuizList = append(s.didQuizList, int32(quizTemplate.TemplateId()))

	now := global.GetGame().GetTimeService().Now()
	s.quizObj.lastQuizTime = now
	s.quizObj.lastQuizId = int32(quizTemplate.TemplateId())
	s.quizObj.answerList = answerOrderList
	s.quizObj.updateTime = now

	s.quizObj.SetModified()
	//发布题目事件
	data := quizeventtypes.CreateQuizAssignEventData(s.quizObj.lastQuizId, s.quizObj.answerList)
	gameevent.Emit(quizeventtypes.EventTypeAssignQuiz, nil, data)
	return
}

var (
	once sync.Once
	cs   *quizService
)

func Init() (err error) {
	once.Do(func() {
		cs = &quizService{}
		err = cs.init()
	})
	return err
}

func GetQuizService() QuizService {
	return cs
}
