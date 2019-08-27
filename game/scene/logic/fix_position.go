package logic

import (
	coretypes "fgame/fgame/core/types"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func FixPosition(bo scene.BattleObject, pos coretypes.Position) {
	s := bo.GetScene()
	if s == nil {
		return
	}

	spl, ok := bo.(scene.Player)
	if ok {
		if !scene.CheckFixPosScene(spl, s) {
			return
		}
	}

	log.WithFields(log.Fields{
		"id":              bo.GetId(),
		"sceneObjectType": bo.GetSceneObjectType(),
		"pos":             pos,
	}).Debugln("scene: 瞬移")

	//保存数据
	bo.Move(pos, bo.GetAngle())
	//保存aoi数据
	s.Move(bo, pos)
	msg := pbutil.BuildSCObjectFixPosition(bo, pos)
	BroadcastNeighborIncludeSelf(bo, msg)
	gameevent.Emit(sceneeventtypes.EventTypeBattleObjectMove, bo, nil)

	if spl != nil {
		if spl.GetLingTong() != nil && !spl.IsLingTongHidden() {
			if CheckIfLingTongAndPlayerSameScene(spl.GetLingTong()) {
				FixPosition(spl.GetLingTong(), pos)
			}
		}
	}
}
