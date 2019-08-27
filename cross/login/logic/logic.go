package logic

import (
	"context"
	"fgame/fgame/cross/player/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

//登陆
func Login(p *player.Player) bool {
	//登陆服务器
	flag := player.GetOnlinePlayerManager().PlayerEnterServer(p)
	if !flag {
		log.WithFields(
			log.Fields{
				"id": p.GetId(),
			}).Warn("player:进入游戏服务器,失败")
		//TODO 断开连接
		p.Close(nil)
		return false
	}
	return true
}

//登出前的操作
func Logout(p *player.Player) {

	flag := p.LogoutSave()
	if !flag {
		panic("login:玩家登出应该成功")
	}

	//退出服务器
	player.GetOnlinePlayerManager().PlayerLeaveServer(p)
	return
}

func onLogout(ctx context.Context, result interface{}, err error) (rerr error) {
	pl := scene.PlayerInContext(ctx)
	p := pl.(*player.Player)
	logout(p)
	return nil
}

func logout(p *player.Player) {
	//退出服务器
	p.Logout()
	player.GetOnlinePlayerManager().PlayerLeaveServer(p)
}

//异步回调
func onPlayerSceneExit(ctx context.Context, result interface{}, err error) (rerr error) {
	pl := scene.PlayerInContext(ctx)
	p := pl.(*player.Player)
	//退出场景
	scenelogic.PlayerExitScene(p, false)
	logout(p)
	return
}
