package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	xuedunlogic "fgame/fgame/game/xuedun/logic"
	"fgame/fgame/game/xuedun/pbutil"
	playerxuedun "fgame/fgame/game/xuedun/player"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeXueDunReset, command.CommandHandlerFunc(handleXueDunReset))
}

func handleXueDunReset(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)

	manager := pl.GetPlayerDataManager(types.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
	manager.GmSetXueDunReset()

	xueDunInfo := manager.GetXueDunInfo()
	//同步属性
	xuedunlogic.XueDunPropertyChanged(pl)
	scXueDunPeiYang := pbutil.BuildSCXueDunGet(xueDunInfo)
	pl.SendMsg(scXueDunPeiYang)
	return
}
