package listener

import (
	"fgame/fgame/core/event"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	quizeventtypes "fgame/fgame/game/quiz/event/types"
	quiz "fgame/fgame/game/quiz/quiz"
	quiztemplate "fgame/fgame/game/quiz/template"
)

//仙尊问答答题频道
func quizAnswerChat(target event.EventTarget, data event.EventData) (err error) {
	eventData := data.(*quizeventtypes.QuizAnswerChatEventData)
	sendId := eventData.GetSendId()
	answerType := eventData.GetAnswerType()
	if !answerType.Valid() {
		return
	}

	quizObj := quiz.GetQuizService().GetQuizObj()
	curTemp := quiztemplate.GetQuizTemplateService().GetQuizByTemplateId(quizObj.GetLastQuizId())
	if curTemp == nil {
		return
	}
	content := curTemp.GetAnswerStrByType(answerType)
	chatlogic.BroadcastQuiz(sendId, chattypes.MsgTypeText, []byte(content))
	return
}

func init() {
	gameevent.AddEventListener(quizeventtypes.EventTypeQuizAnswerChat, event.EventListenerFunc(quizAnswerChat))
}
