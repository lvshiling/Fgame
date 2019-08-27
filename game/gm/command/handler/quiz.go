package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	quiz "fgame/fgame/game/quiz/quiz"
	"fgame/fgame/game/scene/scene"
	"fmt"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeQuizAssign, command.CommandHandlerFunc(handleQuizAssign))
}

func handleQuizAssign(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	log.Debug("gm:处理仙尊问答出题")
	if len(c.Args) < 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	quizIdStr := c.Args[0]

	quizId, err := strconv.ParseInt(quizIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"quizId": quizId,
				"error":  err,
			}).Warn("gm:处理仙尊问答出题,类型quizId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	err = quizAssign(pl, int32(quizId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"quizId": quizId,
				"error":  err,
			}).Warn("gm:处理仙尊问答出题,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":     pl.GetId(),
			"quizId": quizId,
		}).Debug("gm:处理仙尊问答出题,完成")
	return
}

func quizAssign(pl player.Player, quizId int32) (err error) {
	quizService := quiz.GetQuizService()
	ok := quizService.GmAssignQuiz(quizId)
	if ok != nil {
		err = fmt.Errorf("gm:处理仙尊问答出题，出题错误")
	}
	return
}
