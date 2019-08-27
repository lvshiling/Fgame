package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerjieyi "fgame/fgame/game/jieyi/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeJieYiClearLastInviteTime, command.CommandHandlerFunc(handleJieYiClearLastInviteTime))
}

func handleJieYiClearLastInviteTime(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:清除上一次邀请时间")

	pl := p.(player.Player)

	if len(c.Args) < 0 {
		playerlogic.SendSystemMessage(p, lang.GMFormatWrong)
		return
	}

	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	jieYiManager.GmClearLastInvite()

	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:清除上一次邀请时间,完成")
	return
}
