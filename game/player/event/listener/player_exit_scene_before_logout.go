package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	loginlogic "fgame/fgame/game/login/logic"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func playerExitSceneBeforeLogout(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家角色正在登出,退出场景前")
	s := p.GetScene()
	if s != nil {
		//退出场景
		ctx := scene.WithPlayer(context.Background(), p)
		//TODO: zrc:post在退出场景后就会造成玩家卡死服务器
		p.Post(message.NewScheduleMessage(onPlayerSceneExit, ctx, nil, nil))
	} else {
		ctx := scene.WithPlayer(context.Background(), p)
		global.GetGame().GetGlobalRunner().Post(message.NewScheduleMessage(onPlayerSceneExit, ctx, nil, nil))
	}
	return nil
}

func onPlayerSceneExit(ctx context.Context, result interface{}, err error) (rerr error) {

	pl := scene.PlayerInContext(ctx)
	p := pl.(player.Player)
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家角色正在登出,退出场景中")
	//防止退出场景奔溃
	defer func() {
		//保存数据 退出服务器
		loginlogic.Logout(p)
	}()
	//退出场景
	scenelogic.PlayerExitScene(p, false)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerExitSceneBeforeLogout, event.EventListenerFunc(playerExitSceneBeforeLogout))
}
