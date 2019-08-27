package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	"fgame/fgame/core/utils"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/mingge/pbutil"
	playermingge "fgame/fgame/game/mingge/player"
	minggetemplate "fgame/fgame/game/mingge/template"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := p.GetPlayerDataManager(playertypes.PlayerMingGeDataManagerType).(*playermingge.PlayerMingGeDataManager)
	level := p.GetLevel()
	zhuangShu := p.GetZhuanSheng()
	manager.CheckMingGongActivate(level, zhuangShu)
	//命盘
	mingGePanMap := manager.GetMingGePanMap()
	scMingGePanGet := pbutil.BuildSCMingGePanGet(mingGePanMap)
	p.SendMsg(scMingGePanGet)
	//命盘祭炼
	refinedMap := manager.GetMingGePanRefinedMap()
	scMingGeRefinedGet := pbutil.BuildSCMingGeRefinedGet(refinedMap)
	p.SendMsg(scMingGeRefinedGet)
	//命理
	mingLiMap := manager.GetMingLiMap()
	scMingGeMingLiGet := pbutil.BuildSCMingGeMingLiGet(mingLiMap)
	p.SendMsg(scMingGeMingLiGet)

	//补偿
	if !manager.IsBuchang() {
		manager.Buchang()
		returnItemMap := make(map[int32]int32)
		for _, num := range manager.GetBuchangList() {
			buchangTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeBuChang(num)
			if buchangTemplate == nil {
				continue
			}
			utils.MergeMap(returnItemMap, buchangTemplate.GetReturnItemMap())
		}
		if len(returnItemMap) > 0 {
			title := lang.GetLangService().ReadLang(lang.MingGeRefundMailTitle)
			content := lang.GetLangService().ReadLang(lang.MingGeRefundMailContent)
			emaillogic.AddEmail(p, title, content, returnItemMap)
		}
	}
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
