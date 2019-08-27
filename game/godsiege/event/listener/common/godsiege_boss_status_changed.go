package common

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	godsiegeeventtypes "fgame/fgame/game/godsiege/event/types"
	"fgame/fgame/game/godsiege/pbutil"
	godsiegescene "fgame/fgame/game/godsiege/scene"
	godsiegetemplate "fgame/fgame/game/godsiege/template"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fmt"

	"github.com/golang/protobuf/proto"
)

//神兽攻城Boss状态刷新
func godSiegeBossStatusChanged(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(godsiegescene.GodSiegeSceneData)
	mapId := data.(int32)

	boss := sd.GetBoss()
	bossStatus := boss.GetBossStatus()
	godType := sd.GetGodType()
	var godSiegeBossMsg proto.Message
	switch bossStatus {
	case godsiegetypes.GodSiegeBossStatusTypeDead:
		{
			godSiegeBossMsg = pbutil.BuildSCGodSiegeBossDead(int32(godType), int32(bossStatus))
			break
		}
	case godsiegetypes.GodSiegeBossStatusTypeLive:
		{
			pos := boss.GetNpc().GetPosition()
			godSiegeBossMsg = pbutil.BuildSCGodSiegeBossRefresh(int32(godType), int32(bossStatus), pos)

			//公告
			constTemplate := godsiegetemplate.GetGodSiegeTemplateService().GetConstantTemplate()
			bossTemplate := constTemplate.GetBiologyTemplate(mapId)
			systemName := coreutils.FormatColor(chattypes.ColorTypeModuleName, constTemplate.GetSystemName(mapId))
			bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, bossTemplate.Name)
			bossBornContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.GodSiegeBossRefreshNotice), coreutils.FormatNoticeStr(bossName), coreutils.FormatNoticeStr(systemName))
			noticelogic.NoticeNumBroadcastScene(sd.GetScene(), []byte(bossBornContent), 0, int32(1))
			chatlogic.BroadcastScene(sd.GetScene(), chattypes.MsgTypeText, []byte(bossBornContent))
			break
		}
	}

	sd.GetScene().BroadcastMsg(godSiegeBossMsg)
	return
}

func init() {
	gameevent.AddEventListener(godsiegeeventtypes.EventTypeGodSiegeBossStatusRefresh, event.EventListenerFunc(godSiegeBossStatusChanged))
}
