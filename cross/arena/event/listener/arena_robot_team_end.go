package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	arenascene "fgame/fgame/cross/arena/scene"
	gameevent "fgame/fgame/game/event"
	robotlogic "fgame/fgame/game/robot/logic"
	"fgame/fgame/game/robot/robot"
	"fgame/fgame/game/scene/scene"
)

//竞技场机器人结束
func arenaRobotTeamEnd(target event.EventTarget, data event.EventData) (err error) {

	t := target.(*arenascene.ArenaTeam)
	for _, mem := range t.GetTeam().GetMemberList() {
		p := robot.GetRobotService().GetRobot(mem.GetPlayerId())
		if p != nil {
			ctx := scene.WithPlayer(context.Background(), p)
			p.Post(message.NewScheduleMessage(onRobotRemove, ctx, nil, nil))
		}
	}

	return
}

//机器人自动退出
func onRobotRemove(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	robotPl, ok := p.(scene.RobotPlayer)
	if !ok {
		return nil
	}
	robotlogic.RemoveRobot(robotPl)
	return nil
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaRobotTeamEnd, event.EventListenerFunc(arenaRobotTeamEnd))
}
