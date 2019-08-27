package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/emperor/emperor"
	emperoreventtypes "fgame/fgame/game/emperor/event/types"
	emperorlogic "fgame/fgame/game/emperor/logic"
	"fgame/fgame/game/emperor/pbutil"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

//帝王改变
func emperorChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if pl != nil {
		emperorlogic.EmperorPropertyChanged(pl)
	}
	oldPlayerId := data.(int64)
	olpl := player.GetOnlinePlayerManager().GetPlayerById(oldPlayerId)

	needGold := emperor.GetEmperorService().GetLastEmperorCostGold()
	if needGold != 0 {
		itemMap := make(map[int32]int32)
		itemMap[constanttypes.BindGoldItem] = needGold
		curPlayerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStrUnderline(pl.GetName()))
		emailTitle := lang.GetLangService().ReadLang(lang.EmperorRobTitle)
		args := []int64{int64(chattypes.ChatPlayerName), pl.GetId()}
		infoLink := coreutils.FormatLink(curPlayerName, args)
		emailContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmperorRobGiveBlackContent), infoLink)
		if olpl == nil {
			emaillogic.AddOfflineEmail(oldPlayerId, emailTitle, emailContent, itemMap)
		} else {
			emaillogic.AddEmail(olpl, emailTitle, emailContent, itemMap)
		}
	}

	if olpl == nil {
		return
	}
	ctx := scene.WithPlayer(context.Background(), olpl)
	olpl.Post(message.NewScheduleMessage(onEmperorRob, ctx, nil, nil))

	//公告
	power := propertylogic.CulculateForce(emperorlogic.CountEmperorPower())
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	oldPlayerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(olpl.GetName()))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", power))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmperorRobNotice), playerName, oldPlayerName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func onEmperorRob(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)

	scEmperorRobbed := pbuitl.BuildSCEmperorRobbed()
	pl.SendMsg(scEmperorRobbed)
	//属性改变
	return emperorlogic.EmperorPropertyChanged(pl)
}

func init() {
	gameevent.AddEventListener(emperoreventtypes.EmperorEventTypeRobed, event.EventListenerFunc(emperorChanged))
}
