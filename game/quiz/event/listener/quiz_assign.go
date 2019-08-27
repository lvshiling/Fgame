package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	quizeventtypes "fgame/fgame/game/quiz/event/types"
	"fgame/fgame/game/quiz/pbutil"
)

//仙尊问答系统出题
func systemAssignQuiz(target event.EventTarget, data event.EventData) (err error) {
	assignEventData := data.(*quizeventtypes.QuizAssignEventData)
	quizId := assignEventData.GetQuizId()
	answerList := assignEventData.GetAnswerList()
	onlinePlayers := player.GetOnlinePlayerManager().GetAllPlayers()
	scMsg := pbutil.BuildSCQuizAssignInfo(quizId, answerList)
	for _, pl := range onlinePlayers {
		if pl.IsFuncOpen(funcopentypes.FuncOpenTypeQuiz) {
			pl.SendMsg(scMsg)
		}
	}
	return
}

func init() {
	gameevent.AddEventListener(quizeventtypes.EventTypeAssignQuiz, event.EventListenerFunc(systemAssignQuiz))
}
