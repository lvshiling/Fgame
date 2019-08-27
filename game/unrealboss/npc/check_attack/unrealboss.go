package check_attack

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterCheckNPCAttackHandler(scenetypes.BiologyScriptTypeUnrealBoss, scene.NPCCheckAttackHandlerFunc(bossCheckAttack))
}

func bossCheckAttack(attackObject scene.BattleObject, defenceNpc scene.NPC) bool {
	attackPl, ok := attackObject.(player.Player)
	if !ok {
		return true
	}

	//判断是否有疲劳值
	needPilao := defenceNpc.GetBiologyTemplate().NeedPilao
	if !attackPl.IsEnoughPilao(needPilao) && !attackPl.IsPilaoNoticeCd() {

		playerlogic.SendSystemMessage(attackPl, lang.UnrealBossPilaoNumNotEnough)
	}

	return true
}
