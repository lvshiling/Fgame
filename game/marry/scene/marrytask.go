package scene

import (
	"fgame/fgame/game/global"
	"time"
)

const (
	marrySceneTaskTime = time.Second
)

type marrySceneTask struct {
	sd MarrySceneData
}

//婚礼清场
func (mst *marrySceneTask) Run() {
	now := global.GetGame().GetTimeService().Now()
	clearTime := mst.sd.GetClearTime()
	if clearTime == 0 {
		return
	}
	if now >= clearTime {
		sceneAllPlayer := mst.sd.GetScene().GetAllPlayers()
		for _, spl := range sceneAllPlayer {
			spl.BackLastScene()
		}
		mst.sd.ResetClearTime()
	}
}

func (mst *marrySceneTask) ElapseTime() time.Duration {
	return marrySceneTaskTime
}

func CreateMarrySceneTask(sd MarrySceneData) *marrySceneTask {
	marryScenetask := &marrySceneTask{
		sd: sd,
	}
	return marryScenetask
}
