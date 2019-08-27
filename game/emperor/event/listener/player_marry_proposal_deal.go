package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/emperor/emperor"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/player"
)

//被求婚者决策
func playerMarryProposalDeal(target event.EventTarget, data event.EventData) (err error) {
	dpl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*marryeventtypes.MarryProposalDealEventData)
	if !ok {
		return
	}

	agree := eventData.GetAgree()
	if !agree {
		return
	}
	proposalId := eventData.GetProposalId()
	ppl := player.GetOnlinePlayerManager().GetPlayerById(proposalId)
	if ppl == nil {
		return
	}
	name := dpl.GetName()
	dealId := dpl.GetId()
	peerName := ppl.GetName()
	emperor.GetEmperorService().SetEmperorSpouseName(proposalId, name, dealId, peerName)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryProposalDeal, event.EventListenerFunc(playerMarryProposalDeal))
}
