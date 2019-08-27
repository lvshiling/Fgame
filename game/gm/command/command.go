package command

import (
	"fgame/fgame/common/lang"
	commandtypes "fgame/fgame/game/gm/command/types"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type Command struct {
	Type string
	Args []string
}

func ParseCommand(cmd string) (c *Command, err error) {
	cmd = strings.TrimSpace(cmd)
	arrs := strings.Split(cmd, " ")
	if len(arrs) <= 0 {
		err = fmt.Errorf("command [%s] no valid", cmd)
		return
	}

	c = &Command{
		Type: arrs[0],
		Args: arrs[1:],
	}
	return
}

type CommandHandler interface {
	Handle(pl scene.Player, c *Command) (err error)
}

type CommandHandlerFunc func(pl scene.Player, c *Command) (err error)

func (chf CommandHandlerFunc) Handle(pl scene.Player, c *Command) (err error) {
	return chf(pl, c)
}

var (
	handlerMap = make(map[string]CommandHandler)
)

func Register(cmd commandtypes.CommandType, ch CommandHandler) {
	lowercaseCmd := strings.ToLower(cmd.String())
	_, exist := handlerMap[lowercaseCmd]
	if exist {
		panic(fmt.Errorf("repeat register %s command", cmd))
	}
	handlerMap[lowercaseCmd] = ch
}

//执行命令
func RunCommand(pl scene.Player, cmd *Command) (err error) {

	lowcaseType := strings.ToLower(cmd.Type)
	h, ok := handlerMap[lowcaseType]

	if !ok {
		log.WithFields(
			log.Fields{
				"id":  pl.GetId(),
				"cmd": cmd,
			}).Warn("gm: 命令处理器不存在")
		err = nil
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)

		return
	}

	err = h.Handle(pl, cmd)
	if err != nil {
		return
	}
	return
}

var (
	crossMap = make(map[string]bool)
)

func RegisterCross(cmd commandtypes.CommandType) {
	lowercaseCmd := strings.ToLower(cmd.String())
	_, exist := crossMap[lowercaseCmd]
	if exist {
		panic(fmt.Errorf("repeat register %s command", cmd))
	}
	crossMap[lowercaseCmd] = true
}

func IsCross(cmd string) bool {
	lowercaseCmd := strings.ToLower(cmd)
	return crossMap[lowercaseCmd]
}
