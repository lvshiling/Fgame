package logic

import "fgame/fgame/game/scene/scene"

//更新属性
func UpdateMountBattleProperty(p scene.Player) {
	p.Calculate()
	return
}
