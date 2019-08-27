package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playerfuncopen "fgame/fgame/game/funcopen/player"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

type proposalData struct {
	ProposalId int64
	RingType   marrytypes.MarryRingType
}

func newProposalData(proposalId int64, ringType marrytypes.MarryRingType) *proposalData {
	d := &proposalData{
		ProposalId: proposalId,
		RingType:   ringType,
	}
	return d
}

//玩家求婚
func playerMarryProposal(target event.EventTarget, data event.EventData) (err error) {
	eventData, ok := data.(*marryeventtypes.MarryProposalEventData)
	if !ok {
		return
	}
	playerId := eventData.GetPlayerId()
	spouseId := eventData.GetSpouseId()
	ringType := eventData.GetRingType()

	pl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	if pl == nil {
		return
	}

	proposalData := newProposalData(playerId, ringType)
	ctx := scene.WithPlayer(context.Background(), pl)
	playerProposalMsg := message.NewScheduleMessage(onProposal, ctx, proposalData, nil)
	pl.Post(playerProposalMsg)

	return
}

func onProposal(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	data := result.(*proposalData)
	playerId := data.ProposalId
	ringType := data.RingType

	funcopenManager := pl.GetPlayerDataManager(types.PlayerFuncOpenDataManagerType).(*playerfuncopen.PlayerFuncOpenDataManager)
	ppl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if ppl == nil {
		return nil
	}
	//对方结婚功能未开启
	if !funcopenManager.IsOpen(funcopentypes.FuncOpenTypeMarry) {
		pplCtx := scene.WithPlayer(context.Background(), ppl)
		//返还婚戒
		playerRingReplaceMsg := message.NewScheduleMessage(onPlayerRingGiveBack, pplCtx, ringType, nil)
		ppl.Post(playerRingReplaceMsg)
		return nil
	}

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.Proposaled(playerId, ringType)

	itemTemplate := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWedRing, ringType.ItemWedRingSubType())
	ringItem := int32(itemTemplate.TemplateId())
	//推送婚戒信息
	scMarryPushProposal := pbuitl.BuildSCMarryPushProposal(ppl, ringItem)
	pl.SendMsg(scMarryPushProposal)
	return nil
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryProposal, event.EventListenerFunc(playerMarryProposal))
}
