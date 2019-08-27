package check

import (
	emaillogic "fgame/fgame/game/email/logic"
	playeremail "fgame/fgame/game/email/player"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	guaji.RegisterGuaJiCheckHandler(guajitypes.GuaJiCheckTypeEmail, guaji.GuaJiCheckHandlerFunc(emailGuaJiCheck))
}

func emailGuaJiCheck(pl player.Player) {
	flag := emaillogic.CheckPlayerIfCanGetEmailAttachementBatch(pl)
	if flag {
		emaillogic.HandleGetEmailAttachementBatch(pl)
		return
	}
	emailManager := pl.GetPlayerDataManager(playertypes.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)
	for _, m := range emailManager.GetEmails() {
		if !emaillogic.CheckPlayerIfCanGetEmailAttachement(pl, m.GetEmailId()) {
			continue
		}
		emaillogic.HandleGetEmailAttachement(pl, m.GetEmailId())
	}
}
