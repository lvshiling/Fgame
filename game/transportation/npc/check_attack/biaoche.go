package check_attack

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	biaochenpc "fgame/fgame/game/transportation/npc/biaoche"
	playertransportation "fgame/fgame/game/transportation/player"
	"fmt"
)

func init() {
	scene.RegisterCheckNPCAttackHandler(scenetypes.BiologyScriptTypeBiaoChe, scene.NPCCheckAttackHandlerFunc(biaoCheCheckAttack))
}

func biaoCheCheckAttack(attackObject scene.BattleObject, defenceNpc scene.NPC) bool {
	attackPl, ok := attackObject.(player.Player)
	if !ok {
		return true
	}

	if attackPl.GetPkState() == pktypes.PkStatePeach {
		playerlogic.SendSystemMessage(attackPl, lang.TransportationAttackedFailed)
		return false
	}
	biaoChe := defenceNpc.(*biaochenpc.BiaocheNPC)
	//判断是否打过三次以上
	//镖车管理器
	m := attackPl.GetPlayerDataManager(playertypes.PlayerTransportationType).(*playertransportation.PlayerTransportationDataManager)
	robTime := m.GetRobOfTimes()
	maxRobTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeRobTransportationTimes)

	//抢过镖车
	if m.IfRob(biaoChe.GetTransportationObject().GetTransportId()) {
		return true
	}

	//抢超过次数
	if robTime < maxRobTime {
		return true
	}
	playerlogic.SendSystemMessage(attackPl, lang.TransportationRobFull, fmt.Sprintf("%d", maxRobTime))
	return false
}
