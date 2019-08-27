package logic

import (
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//退出战斗
func ExitBattle(pl scene.Player) {
	objectBattle := pbutil.BuildObjectBattle(pl, false)
	pl.SendMsg(objectBattle)
}

//进入战斗
func EnterBattle(pl scene.Player) {
	objectBattle := pbutil.BuildObjectBattle(pl, true)
	pl.SendMsg(objectBattle)
}

func FindHatestEnemy(bo scene.BattleObject) (e *scene.Enemy) {
	s := bo.GetScene()
	if s == nil {
		return
	}
	isPlayer := bo.GetSceneObjectType() == scenetypes.BiologyTypePlayer
	maxHate := 0
	for _, te := range bo.GetEnemies() {
		if te.GetHate() > maxHate {
			if isPlayer {
				switch te.BattleObject.(type) {
				case scene.Player:
					{
						//检查是否在安全区内
						if s.MapTemplate().IsSafe(te.BattleObject.GetPosition()) {
							continue
						}
						break
					}
				}
			}
			e = te
			maxHate = te.GetHate()
		}
	}
	return
}
