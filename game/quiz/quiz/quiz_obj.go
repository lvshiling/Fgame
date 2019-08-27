package quiz

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	quizentity "fgame/fgame/game/quiz/entity"
	quiztypes "fgame/fgame/game/quiz/types"

	"github.com/pkg/errors"
)

//仙尊答题对象
type QuizObject struct {
	id           int64
	serverId     int32
	lastQuizId   int32
	lastQuizTime int64
	answerList   []quiztypes.QuizAnswerType
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewQuizObject() *QuizObject {
	o := &QuizObject{}
	return o
}

func convertNewQuizObjectToEntity(o *QuizObject) (*quizentity.QuizEntity, error) {
	answerBytes, err := json.Marshal(o.answerList)
	if err != nil {
		return nil, err
	}
	e := &quizentity.QuizEntity{
		Id:           o.id,
		ServerId:     o.serverId,
		LastQuizId:   o.lastQuizId,
		LastQuizTime: o.lastQuizTime,
		AnswerList:   string(answerBytes),
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *QuizObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *QuizObject) GetLastQuizId() int32 {
	return o.lastQuizId
}

func (o *QuizObject) GetLastQuizTime() int64 {
	return o.lastQuizTime
}

func (o *QuizObject) GetAnswerList() []quiztypes.QuizAnswerType {
	return o.answerList
}

func (o *QuizObject) GetDBId() int64 {
	return o.id
}

func (o *QuizObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewQuizObjectToEntity(o)
	return e, err
}

func (o *QuizObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*quizentity.QuizEntity)

	var answerOrderList []quiztypes.QuizAnswerType
	if err := json.Unmarshal([]byte(pse.AnswerList), &answerOrderList); err != nil {
		return err
	}

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.lastQuizId = pse.LastQuizId
	o.lastQuizTime = pse.LastQuizTime
	o.answerList = answerOrderList
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *QuizObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Quiz"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}
