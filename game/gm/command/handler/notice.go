package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	noticelogic "fgame/fgame/game/notice/logic"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeNotice, command.CommandHandlerFunc(handleNotice))
}

//跑马灯测试
func handleNotice(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理跑马灯")

	if len(c.Args) < 3 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	intervalTimeStr := c.Args[0]
	numStr := c.Args[1]
	content := []byte(c.Args[2])
	intervalTime, err := strconv.ParseInt(intervalTimeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"intervalTime": intervalTimeStr,
				"num":          numStr,
				"content":      content,
				"error":        err,
			}).Warn("gm:处理跑马灯,intervalTime不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"intervalTime": intervalTimeStr,
				"num":          numStr,
				"content":      content,
				"error":        err,
			}).Warn("gm:处理跑马灯,num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	noticelogic.NoticeNumBroadcast(content, intervalTime, int32(num))
	return
}
