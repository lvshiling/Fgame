package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/email/pbutil"
	playeremail "fgame/fgame/game/email/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	//发送邮件列表
	p := target.(player.Player)
	emailManager := p.GetPlayerDataManager(playertypes.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)

	emailList := emailManager.GetEmails()
	scEmailsGet := pbutil.BuildSCEmailsGet(emailList)
	p.SendMsg(scEmailsGet)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
