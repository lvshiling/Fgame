package logic

import (
	coretypes "fgame/fgame/core/types"
	chuangshiscene "fgame/fgame/cross/chuangshi/scene"
	chuangshitemplate "fgame/fgame/game/chuangshi/template"
	"fgame/fgame/game/scene/scene"
)

const (
	numOfDoor = 1
)

func ChuangShiWarCheckMove(p scene.Player, pos coretypes.Position) (flag bool, fixPos coretypes.Position) {
	s := p.GetScene()
	if s == nil {
		flag = true
		return
	}
	sd, ok := s.SceneDelegate().(chuangshiscene.FuShuSceneData)
	if !ok {
		flag = true
		return
	}

	defenCampType := sd.GetInitDefendCampType()
	warTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarTemp(defenCampType)
	destArea := warTemp.GetArea(pos)
	//TODO 优化3个门
	//城门全破了
	isDefend := sd.GetCurrentDefendCampType() == p.GetCamp()
	currentDoor := sd.GetCurrentDoor()
	if currentDoor >= numOfDoor {
		flag = true
		return
	}
	switch currentDoor {
	case 0:
		if isDefend && (destArea <= 1) {
			flag = true
			return
		}
		if destArea <= 0 {
			flag = true
			return
		}
		flag = false
		sourceArea := warTemp.GetArea(p.GetPosition())
		if sourceArea != 0 {
			fixPos = s.MapTemplate().GetBornPos()
		} else {
			tempFixPos, tflag := warTemp.GetFixPos(currentDoor)
			if tflag {
				if s.MapTemplate().GetMap().IsMask(tempFixPos.X, tempFixPos.Z) {
					fixPos = tempFixPos
					return
				}
			}
			fixPos = tempFixPos
			fixPos.Y = s.MapTemplate().GetMap().GetHeight(fixPos.X, fixPos.Z)
			return
		}
		return
	default:
		if destArea <= currentDoor {
			flag = true
			return
		}
		sourceArea := warTemp.GetArea(p.GetPosition())
		if sourceArea >= destArea {
			fixPos = s.MapTemplate().GetBornPos()
		} else {
			tempFixPos, tflag := warTemp.GetFixPos(currentDoor)
			if tflag {
				if s.MapTemplate().GetMap().IsMask(tempFixPos.X, tempFixPos.Z) {
					fixPos = tempFixPos
					return
				}
			}
			fixPos = tempFixPos
			fixPos.Y = s.MapTemplate().GetMap().GetHeight(fixPos.X, fixPos.Z)
			return
		}
		return
	}
	flag = true
	return
}
