package logic

import "fgame/fgame/game/scene/scene"

func CheckIfLingTongAndPlayerSameScene(lingTong scene.LingTong) bool {
	s := lingTong.GetScene()
	if s == nil {
		return false
	}
	owner := lingTong.GetOwner()
	if owner == nil {
		return false
	}
	if owner.GetScene() != s {
		return false
	}
	return true
}
