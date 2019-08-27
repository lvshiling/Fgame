package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancetypes "fgame/fgame/game/alliance/types"
	chatlogic "fgame/fgame/game/chat/logic"
	chatpbutil "fgame/fgame/game/chat/pbutil"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

//成员加入
func memberJoin(target event.EventTarget, data event.EventData) (err error) {
	al := target.(*alliance.Alliance)
	memObj := data.(*alliance.AllianceMemberObject)
	memId := memObj.GetMemberId()

	//成员变更推送
	scAllianceMemberChanged := pbutil.BuildSCAllianceMemberChanged(al.GetMemberList())
	for _, mem := range al.GetMemberList() {
		p := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if p == nil {
			continue
		}

		p.SendMsg(scAllianceMemberChanged)
	}

	//广播帮派
	if al.GetAllianceMengZhuId() != memId {
		joinName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, memObj.GetName())
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceJoinNotice), joinName)
		chatlogic.SystemBroadcastAlliance(al, chattypes.MsgTypeText, []byte(content))
	}

	pl := player.GetOnlinePlayerManager().GetPlayerById(memId)
	if pl == nil {
		return
	}
	//发送个人信息
	if pl.GetId() != al.GetAllianceMengZhuId() {
		personalContent := lang.GetLangService().ReadLang(lang.AllianceJoinPersonalNotice)
		chatRecv := chatpbutil.BuildSCChatRecvWithCliArgs(al.GetAllianceMengZhuId(), al.GetMengzhuName(), chattypes.ChannelTypePerson, memId, chattypes.MsgTypeText, []byte(personalContent), "")
		pl.SendMsg(chatRecv)
	}

	ctx := scene.WithPlayer(context.Background(), pl)
	memberJoinMsg := message.NewScheduleMessage(onMemberJoin, ctx, memObj, nil)
	pl.Post(memberJoinMsg)

	return
}

//成员加入回调
func onMemberJoin(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	mem := result.(*alliance.AllianceMemberObject)
	al := mem.GetAlliance()

	allianceId := al.GetAllianceId()
	mengzhuId := al.GetAllianceMengZhuId()
	allianceName := al.GetAllianceName()
	allianceLevel := al.GetAllianceLevel()
	pos := mem.GetPosition()

	allianceManager := tpl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceManager.JoinAlliance(allianceId, mengzhuId, allianceName, allianceLevel, pos)

	tpl.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeLingyuAura.Mask())
	return nil
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceMemberJoin, event.EventListenerFunc(memberJoin))
}
