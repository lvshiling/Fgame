package check_attack

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterCheckNPCAttackHandler(scenetypes.BiologyScriptTypeXianMengNPC, scene.NPCCheckAttackHandlerFunc(xianMengNpcCheckAttack))
}

func xianMengNpcCheckAttack(attackObject scene.BattleObject, defenceNpc scene.NPC) bool {
	attackPl, ok := attackObject.(player.Player)
	if !ok {
		return attackObject.IsEnemy(defenceNpc)
	}

	if attackPl.IsEnemy(defenceNpc) {
		return true
	}

	//被召唤过而且是友方
	if defenceNpc.GetOwnerId() != 0 {
		return false
	}

	playerlogic.SendSystemMessage(attackPl, lang.AllianceSceneGuardInactive)
	return false
}
