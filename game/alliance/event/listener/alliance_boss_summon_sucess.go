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
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fmt"
)

//仙盟boss召唤成功
func allianceBossSummonSucess(target event.EventTarget, data event.EventData) (err error) {
	al := target.(*alliance.Alliance)
	eventData := data.(*allianceeventtypes.AllianceBossSummonSucessEventData)
	pl := eventData.GetPlayer()
	s := eventData.GetScene()

	scMsg := pbutil.BuildSCAllianceBossSummon()
	pl.SendMsg(scMsg)
	bornPos := s.MapTemplate().GetBornPos()
	scenelogic.PlayerEnterScene(pl, s, bornPos)

	scAllianceBossSummonSucess := pbutil.BuildSCAllianceBossSummonSucess()
	for _, member := range al.GetMemberList() {
		playerId := member.GetMemberId()
		if playerId == pl.GetId() {
			continue
		}
		curPl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if curPl == nil {
			continue
		}
		curPl.SendMsg(scAllianceBossSummonSucess)
	}

	//仙盟频道
	if s == nil {
		return
	}
	sd := s.SceneDelegate().(alliancebossscene.AllianceBossSceneData)
	if sd == nil {
		return
	}
	level := sd.GetLevel()
	allianceBossTemplate := alliancetemplate.GetAllianceTemplateService().GetAllianceBossTemplate(level)
	if allianceBossTemplate == nil {
		return
	}
	biologyTemplate := allianceBossTemplate.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}
	mapTemplate := allianceBossTemplate.GetMapTemplate()
	if mapTemplate == nil {
		return
	}
	mapId := int64(mapTemplate.TemplateId())
	clientMapType := int64(mapTemplate.Type)
	biologyName := fmt.Sprintf("%d级•%s", level, biologyTemplate.Name)
	bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(biologyName))
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))

	args := []int64{int64(chattypes.ChatPlayerEnterScene), funcopentypes.FuncOpenTypeAllianceBoss, clientMapType, mapId}
	joinLink := coreutils.FormatLink(chattypes.ButtonTypeToGoNow, args)
	var allianceStr string
	if pl.GetId() == al.GetFuMengZhuId() {
		allianceStr = "盟主"
	} else {
		allianceStr = "副盟主"
	}
	bossRefreshContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceBossSummonSucessChat), allianceStr, playerName, bossName, joinLink)
	chatlogic.SystemBroadcastAlliance(al, chattypes.MsgTypeText, []byte(bossRefreshContent))
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EvnetTypeAllianceBossSummonSucess, event.EventListenerFunc(allianceBossSummonSucess))
}
