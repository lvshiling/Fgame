package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	"fmt"
)

var (
	rankMsgCode = map[int32]lang.LangCode{
		1: lang.RankLoginFirst,
		2: lang.RankLoginSecond,
		3: lang.RankLoginThird,
	}
)

// 玩家加载后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	// 战力榜登录公告
	pos := rank.GetRankService().GetMyRankPos(ranktypes.RankClassTypeLocal, 0, ranktypes.RankTypeForce, pl.GetId())

	// // 名次条件
	// global.GetGame().GetPlatform()
	// template.GetTemplateService()
	// if pos > condition{
	// 	return
	// }

	msgCode, ok := rankMsgCode[pos]
	if !ok {
		return
	}

	plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	content := fmt.Sprintf(lang.GetLangService().ReadLang(msgCode), plName)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
