package login_handler

import (
	"fgame/fgame/cross/arena/arena"
	arenalogic "fgame/fgame/cross/arena/logic"
	"fgame/fgame/cross/login/login"
	"fgame/fgame/cross/player/player"
	crosstypes "fgame/fgame/game/cross/types"
	scenelogic "fgame/fgame/game/scene/logic"
)

func init() {
	login.RegisterLoginHandler(crosstypes.CrossTypeArena, login.LogincHandlerFunc(arenaLogin))
}

func arenaLogin(pl *player.Player, ct crosstypes.CrossType, args ...string) bool {
	s := arena.GetArenaService().GetArenaSceneByPlayerId(pl.GetId())
	if s != nil {
		if s.MapTemplate().IsArena() {
			arenalogic.PlayerEnterArenaScene(s, pl)
		} else {
			//四圣兽场景
			pos := s.MapTemplate().GetBornPos()
			scenelogic.PlayerEnterScene(pl, s, pos)
		}
		return true
	}

	return false
}
