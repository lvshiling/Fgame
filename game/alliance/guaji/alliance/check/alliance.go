package check

import (
	"fgame/fgame/common/lang"
	alliancelogic "fgame/fgame/game/alliance/logic"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
)

func init() {
	guaji.RegisterGuaJiCheckHandler(guajitypes.GuaJiCheckTypeAlliance, guaji.GuaJiCheckHandlerFunc(allianceGuaJiCheck))
}

func allianceGuaJiCheck(pl player.Player) {

	flag := alliancelogic.CheckPlayerIfCanBatchJoinAlliance(pl)
	if !flag {
		return
	}
	alliancelogic.HandleAllianceJoinBatch(pl)
	playerlogic.SendSystemMessage(pl, lang.GuaJiAllianceBatchJoin)
}
