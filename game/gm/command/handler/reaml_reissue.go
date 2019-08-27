package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerrealm "fgame/fgame/game/realm/player"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeSetRealmCheckReissue, command.CommandHandlerFunc(handleSetRealmCheckReissue))

}

func handleSetRealmCheckReissue(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	log.Debug("gm:处理天劫塔补偿重置")

	realmManager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	if realmManager.IsCheckReissue() {
		realmManager.GmSetCheckReissueOn()
	}

	return
}
