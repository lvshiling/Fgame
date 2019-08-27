package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	"fgame/fgame/game/bagua/pbutil"
	playerbagua "fgame/fgame/game/bagua/player"
	baguatemplate "fgame/fgame/game/bagua/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerBaGuaAfterLogin(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := p.GetPlayerDataManager(playertypes.PlayerBaGuaDataManagerType).(*playerbagua.PlayerBaGuaDataManager)
	level := manager.GetLevel()

	scBaGuaLevel := pbutil.BuildSCBaGuaLevel(level)
	p.SendMsg(scBaGuaLevel)

	// 八卦秘境补偿
	if manager.IsBuChang() {
		emailBaGuaMiJingBuChang(p, level)
		manager.SetIsBuChang()
	}
	return
}

func emailBaGuaMiJingBuChang(pl player.Player, level int32) {
	emailItemMap := make(map[int32]int32)
	for level > 0 {
		buchangTemplate := baguatemplate.GetBaGuaTemplateService().GetBaGuaMiJingBuChangTemplateByLevel(level)
		if buchangTemplate == nil {
			level--
			continue
		}
		itemMap := buchangTemplate.GetItemMap()
		for id, count := range itemMap {
			num, ok := emailItemMap[id]
			if !ok {
				emailItemMap[id] = count
				continue
			}
			emailItemMap[id] = (num + count)
		}

		level--
	}
	if len(emailItemMap) != 0 {
		title := lang.GetLangService().ReadLang(lang.BaGuaEmailTitleBuChang)
		content := lang.GetLangService().ReadLang(lang.BaGuaEmailContentBuChang)
		emaillogic.AddEmail(pl, title, content, emailItemMap)
	}
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerBaGuaAfterLogin))
}
