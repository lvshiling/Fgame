package types

import (
	quiztypes "fgame/fgame/game/quiz/types"
)

type QuizEventType string

const (
	//仙尊问答出题
	EventTypeAssignQuiz QuizEventType = "AssignQuiz"
	//仙尊问答答题频道
	EventTypeQuizAnswerChat QuizEventType = "QuizAnswerChat"
)

//仙尊问答出题类型
type QuizAssignEventData struct {
	quizId     int32
	answerList []quiztypes.QuizAnswerType
}

func CreateQuizAssignEventData(id int32, orderList []quiztypes.QuizAnswerType) *QuizAssignEventData {
	d := &QuizAssignEventData{
		quizId:     id,
		answerList: orderList,
	}
	return d
}

func (d *QuizAssignEventData) GetQuizId() int32 {
	return d.quizId
}

func (d *QuizAssignEventData) GetAnswerList() []quiztypes.QuizAnswerType {
	return d.answerList
}

//仙尊问答答题类型
type QuizAnswerChatEventData struct {
	sendId     int64
	answerType quiztypes.QuizAnswerType
}

func CreateQuizAnswerChatEventData(playerId int64, typ quiztypes.QuizAnswerType) *QuizAnswerChatEventData {
	d := &QuizAnswerChatEventData{
		sendId:     playerId,
		answerType: typ,
	}
	return d
}

func (d *QuizAnswerChatEventData) GetSendId() int64 {
	return d.sendId
}

func (d *QuizAnswerChatEventData) GetAnswerType() quiztypes.QuizAnswerType {
	return d.answerType
}
