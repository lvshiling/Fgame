package check_attack

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterCheckNPCAttackHandler(scenetypes.BiologyScriptTypeOutlandBoss, scene.NPCCheckAttackHandlerFunc(bossCheckAttack))
}

func bossCheckAttack(attackObject scene.BattleObject, defenceNpc scene.NPC) bool {
	attackPl, ok := attackObject.(player.Player)
	if !ok {
		return true
	}

	//判断是浊气上限了
	//needZhuoQi := defenceNpc.GetBiologyTemplate().ZhuoQi
	//manager := pl.GetPlayerDataManager(playertypes.PlayerOutlandBossDataManagerType).(*playeroutlandboss.PlayerOutlandBossDataManager)
	//manager.RefreshZhuoQi()
	if attackPl.IsZhuoQiLimit() {
		if !attackPl.IsZhuoQiNoticeCd() {
			playerlogic.SendSystemMessage(attackPl, lang.OutlandBossZhuoQiNumNotEnough)
		}
		return false
	}

	return true
}
