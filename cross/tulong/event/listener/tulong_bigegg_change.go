package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/cross/chat/logic"
	noticelogic "fgame/fgame/cross/notice/logic"
	tulongeventtypes "fgame/fgame/cross/tulong/event/types"
	"fgame/fgame/cross/tulong/pbutil"
	tulongscene "fgame/fgame/cross/tulong/scene"
	tulongtypes "fgame/fgame/cross/tulong/types"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	tulongtemplate "fgame/fgame/game/tulong/template"
	"fmt"
)

//大龙蛋状态改变
func tuLongBigEggStatusChange(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(tulongscene.TuLongSceneData)
	if !ok {
		return
	}
	bigEgg := sd.GetBigEgg()
	status := bigEgg.GetStatus()
	biaoShi := bigEgg.GetBornBiaoShi()
	scTuLongBossStatus := pbutil.BuildSCTuLongBossStatus(int32(status), biaoShi)
	sd.GetScene().BroadcastMsg(scTuLongBossStatus)
	//公告
	if status == tulongtypes.EggStatusTypeBoss {
		tuLongTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongBigBossTemplate()
		tuLongConstTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongConstTemplate()
		bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(tuLongTemplate.GetBiologyTemplate().Name))
		mapName := coreutils.FormatColor(chattypes.ColorTypeBossMap, coreutils.FormatNoticeStr(tuLongConstTemplate.GetMapTemplate().Name))
		bossBornContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.TuLongBossBorn), bossName, mapName)
		//TODO 跑马灯
		noticelogic.NoticeNumBroadcast([]byte(bossBornContent), 0, int32(1))
		//系统频道
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(bossBornContent))
	}
	return
}

func init() {
	gameevent.AddEventListener(tulongeventtypes.EventTypeTuLongBigEggStatusRefresh, event.EventListenerFunc(tuLongBigEggStatusChange))
}
