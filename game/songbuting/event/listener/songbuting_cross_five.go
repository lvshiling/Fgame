package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	songbutingeventtypes "fgame/fgame/game/songbuting/event/types"
	"fgame/fgame/game/songbuting/pbutil"
	playersongbuting "fgame/fgame/game/songbuting/player"
	songbutingtemplate "fgame/fgame/game/songbuting/template"
	"fmt"
)

//送不停跨5点
func songBuTingReward(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerSongBuTingDataManagerType).(*playersongbuting.PlayerSongBuTingDataManager)
	songBuTingObj := data.(*playersongbuting.PlayerSongBuTingObject)
	if songBuTingObj.GetIsReceive() &&
		songBuTingObj.GetTimes() == 0 {
		songBuTingTemplate := songbutingtemplate.GetSongBuTingTemplateService().GetSongBuTingTemplate()
		offItemMap := songBuTingTemplate.GetRewOffItemMap()
		emailTitle := fmt.Sprintf(lang.GetLangService().ReadLang(lang.SongBuTingTitle), songBuTingTemplate.NeedGold)
		emailContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.SongBuTingContent), songBuTingTemplate.NeedGold)
		emaillogic.AddEmail(pl, emailTitle, emailContent, offItemMap)
	}
	obj := manager.CrossFiveReset()
	scSongBuTingChanged := pbutil.BuildSCSongBuTingChanged(obj)
	pl.SendMsg(scSongBuTingChanged)
	return nil
}

func init() {
	gameevent.AddEventListener(songbutingeventtypes.EventTypeSongBuTingCrossFive, event.EventListenerFunc(songBuTingReward))
}
