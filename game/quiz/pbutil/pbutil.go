package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	quiztypes "fgame/fgame/game/quiz/types"
)

func BuildSCQuizAnswer(result int32, answerIdx int32) *uipb.SCQuizAnswer {
	scMsg := &uipb.SCQuizAnswer{}
	scMsg.Result = &result
	scMsg.AnswerIdx = &answerIdx
	return scMsg
}

func BuildSCQuizAssignInfo(quizId int32, answerOrder []quiztypes.QuizAnswerType) *uipb.SCQuizAssignInfo {
	scMsg := &uipb.SCQuizAssignInfo{}
	scMsg.QuizId = &quizId
	for i := 0; i < len(answerOrder); i++ {
		scMsg.AnswerOrder = append(scMsg.AnswerOrder, int32(answerOrder[i]))
	}
	return scMsg
}
