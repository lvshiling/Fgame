package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	alliancebossscene "fgame/fgame/game/alliance/boss_scene"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	alliancetemplate "fgame/fgame/game/alliance/template"
	chattypes "fgame/fgame/game/chat/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fmt"
)

//仙盟boss场景结束
func allianceBossSceneEnd(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(alliancebossscene.AllianceBossSceneData)
	s := sd.GetScene()
	if s == nil {
		return
	}
	sucess := data.(bool)
	allianceId := sd.GetAllianceId()

	scAllianceBossEnd := pbutil.BuildSCAllianceBossEnd(sucess)
	s.BroadcastMsg(scAllianceBossEnd)
	if sucess {
		level := sd.GetLevel()
		allianceBossReward(allianceId, level)
	}
	alliance.GetAllianceService().AllianceBossEnd(allianceId)

	return
}

func allianceBossReward(allianceId int64, level int32) {
	allianceBossTemplate := alliancetemplate.GetAllianceTemplateService().GetAllianceBossTemplate(level)
	if allianceBossTemplate == nil {
		return
	}
	biologyTemplate := allianceBossTemplate.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}
	biologyName := fmt.Sprintf("%d级•%s", level, biologyTemplate.Name)

	al := alliance.GetAllianceService().GetAlliance(allianceId)
	if al == nil {
		return
	}

	mengZhuId := al.GetAllianceMengZhuId()
	allianceMemberList := al.GetMemberList()

	bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(biologyName))
	emailTitle := lang.GetLangService().ReadLang(lang.AllianceBossTitle)
	emailContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceBossContent), bossName)

	mzRewItemMap := allianceBossTemplate.GetMzRewItemMap()
	cyRewItemMap := allianceBossTemplate.GetCyRewItemMap()
	//盟主
	if len(mzRewItemMap) != 0 {
		pl := player.GetOnlinePlayerManager().GetPlayerById(mengZhuId)
		if pl != nil {
			emaillogic.AddEmail(pl, emailTitle, emailContent, mzRewItemMap)
		} else {
			emaillogic.AddOfflineEmail(mengZhuId, emailTitle, emailContent, mzRewItemMap)
		}
	}
	//盟主
	for _, member := range allianceMemberList {
		playerId := member.GetMemberId()
		if playerId == mengZhuId {
			continue
		}
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl != nil {
			emaillogic.AddEmail(pl, emailTitle, emailContent, cyRewItemMap)
		} else {
			emaillogic.AddOfflineEmail(playerId, emailTitle, emailContent, cyRewItemMap)
		}
	}
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceBossSceneFinish, event.EventListenerFunc(allianceBossSceneEnd))
}
