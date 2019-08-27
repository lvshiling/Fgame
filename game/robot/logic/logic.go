package logic

import (
	"fgame/fgame/game/robot/robot"
	"fgame/fgame/game/scene/scene"
)

//移除机器人
func RemoveRobot(p scene.RobotPlayer) {
	robot.GetRobotService().RemoveRobot(p.GetId())
	s := p.GetScene()
	if s != nil {
		s.RemoveSceneObject(p, true)
	}
}
