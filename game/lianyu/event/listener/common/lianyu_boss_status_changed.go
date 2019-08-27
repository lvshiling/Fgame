package common

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	lianyueventtypes "fgame/fgame/game/lianyu/event/types"
	"fgame/fgame/game/lianyu/pbutil"
	lianyuscene "fgame/fgame/game/lianyu/scene"
	lianyutemplate "fgame/fgame/game/lianyu/template"
	lianyutypes "fgame/fgame/game/lianyu/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fmt"

	"github.com/golang/protobuf/proto"
)

//无间炼狱Boss状态刷新
func lianYuBossStatusChanged(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(lianyuscene.LianYuSceneData)
	if !ok {
		return
	}

	s := sd.GetScene()
	boss := sd.GetBoss()
	bossStatus := boss.GetBossStatus()
	var lianYuBossMsg proto.Message
	switch bossStatus {
	case lianyutypes.LianYuBossStatusTypeDead:
		{
			lianYuBossMsg = pbutil.BuildSCLianYuBossDead(int32(bossStatus))
			break
		}
	case lianyutypes.LianYuBossStatusTypeLive:
		{
			pos := boss.GetNpc().GetPosition()
			lianYuBossMsg = pbutil.BuildSCLianYuBossRefresh(int32(bossStatus), pos)

			//公告
			constTemplate := lianyutemplate.GetLianYuTemplateService().GetConstantTemplate(sd.GetAcitvityType())
			if constTemplate == nil {
				return
			}
			bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(constTemplate.GetBiologyTemplate().Name))
			bossBornContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.LianYuBossRefreshNotice), bossName)
			noticelogic.NoticeNumBroadcastScene(s, []byte(bossBornContent), 0, int32(1))
			chatlogic.BroadcastScene(s, chattypes.MsgTypeText, []byte(bossBornContent))
			break
		}
	}

	s.BroadcastMsg(lianYuBossMsg)
	return
}

func init() {
	gameevent.AddEventListener(lianyueventtypes.EventTypeLianYuBossStatusRefresh, event.EventListenerFunc(lianYuBossStatusChanged))
}
