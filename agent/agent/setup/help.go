package setup

import gameserver "fgame/fgame/game/server"

func SetupGameServer(baseOptions *gameserver.GameServerOptions, dir string, serverId int32) (err error) {
	s := NewConfigService(baseOptions)
	err = s.SetupGameServer(dir, serverId)
	if err != nil {
		return
	}
	return
}
