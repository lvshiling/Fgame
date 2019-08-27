package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/marry/marry"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/scene/scene"
	"fmt"

	marrylogic "fgame/fgame/game/marry/logic"

	log "github.com/Sirupsen/logrus"
)

type proporalDealData struct {
	SpouseId   int64
	SpouseName string
	RingType   marrytypes.MarryRingType
}

func newProporalDealData(spouseId int64, spouseName string, ringType marrytypes.MarryRingType) *proporalDealData {
	d := &proporalDealData{
		SpouseId:   spouseId,
		SpouseName: spouseName,
		RingType:   ringType,
	}
	return d
}

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
	proposalId := eventData.GetProposalId()
	ppl := player.GetOnlinePlayerManager().GetPlayerById(proposalId)
	if dpl == nil || ppl == nil {
		return
	}

	agree := eventData.GetAgree()
	manager := dpl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	//获取婚戒类型
	ringType := manager.GetProposalRingType(proposalId)

	if agree {

		proporalDealData := newProporalDealData(dpl.GetId(), dpl.GetName(), ringType)
		pplCtx := scene.WithPlayer(context.Background(), ppl)
		playerProposalSucess := message.NewScheduleMessage(onProposalSucess, pplCtx, proporalDealData, nil)
		ppl.Post(playerProposalSucess)
		manager.ProposalMarry(ppl.GetId(), ppl.GetName(), ringType, false)

		marryInfo := manager.GetMarryInfo()

		playerSuitMap := manager.GetAllDingQingMap()

		scdMarryGet := pbuitl.BuildSCMarryGet(dpl, marryInfo, -1, true, playerSuitMap, nil)
		dpl.SendMsg(scdMarryGet)
		// sendMarryCountChange(dpl, marryInfo.MarryCount)
		//定情信物求婚成功后同步数据
		marrylogic.SyncMarryPlayerDingQing(dpl)

		pplName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(ppl.GetName()))
		dplName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(dpl.GetName()))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryProposalSucess), pplName, dplName)
		//跑马灯
		noticelogic.NoticeNumBroadcast([]byte(content), 0, int32(1))
		//系统频道
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	} else { //拒绝

	}
	scMarryProposalResult := pbuitl.BuildSCMarryProposalResult(agree, dpl.GetName())
	ppl.SendMsg(scMarryProposalResult)

	return
}

func onProposalSucess(ctx context.Context, result interface{}, err error) error {

	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	dealData := result.(*proporalDealData)
	ringType := dealData.RingType
	spouseId := dealData.SpouseId
	spouseName := dealData.SpouseName

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)

	manager.ProposalMarry(spouseId, spouseName, ringType, true)

	playerSuitMap := manager.GetAllDingQingMap()

	//推送婚烟信息
	marryInfo := manager.GetMarryInfo()

	scpMarryGet := pbuitl.BuildSCMarryGet(pl, marryInfo, -1, true, playerSuitMap, nil)
	pl.SendMsg(scpMarryGet)
	// sendMarryCountChange(pl, marryInfo.MarryCount)

	marrylogic.SyncMarryPlayerDingQing(pl)

	curRingLevel := marryInfo.RingLevel
	marry.GetMarryService().MarryRingLevel(pl.GetId(), curRingLevel)
	//将求婚的己算到消费记录
	ringItem := marrytemplate.GetMarryTemplateService().GetMarryBanquetTeamplate(marrytypes.MarryBanquetTypeRing, ringType.BanquetSubTypeRing())
	if ringItem == nil {
		log.WithFields(log.Fields{
			"playerId":             pl.GetId(),
			"ringType":             ringType,
			"MarryBanquetTypeRing": marrytypes.MarryBanquetTypeRing,
			"merrySubRing":         ringType.BanquetSubTypeRing(),
		}).Warn("marry:结婚同意,结婚戒指类型错误")
		return nil
	}
	costGold := ringItem.UseGold
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateCycleCostInfo(int64(costGold))
	gameevent.Emit(propertyeventtypes.EventTypePlayerGoldCost, pl, int64(costGold))
	gameevent.Emit(propertyeventtypes.EventTypePlayerGoldCostIncludeBind, pl, int64(costGold))
	return nil

}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryProposalDeal, event.EventListenerFunc(playerMarryProposalDeal))
}
