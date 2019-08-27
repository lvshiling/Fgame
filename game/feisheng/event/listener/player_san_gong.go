package listener

import (
	"context"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	feishengeventtypes "fgame/fgame/game/feisheng/event/types"
	"fgame/fgame/game/feisheng/pbutil"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	feishengtemplate "fgame/fgame/game/feisheng/template"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

type sanGongData struct {
	plName  string
	giveExp int64
}

// 玩家散功
func playerSanGong(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	costExp, ok := data.(int64)
	if !ok {
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}

	// 散功公告
	npcId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFeiShengSanGongNPC)
	plName := chatlogic.FormatModuleNameNoticeStr(pl.GetName())
	mapName := chatlogic.FormatModuleNameNoticeStr(s.MapTemplate().Name)
	args := []int64{int64(chattypes.ChatLinkTypeNpc), int64(npcId)}
	linkText := coreutils.FormatLink(chattypes.ButtonTypeToSanGong, args)
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FeiShengSanGongNotice), plName, mapName, linkText)
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))

	//散功：玩家所在场景的其他玩家获得经验
	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	feiTemp := feishengtemplate.GetFeiShengTemplateService().GetFeiShengTemplate(feiManager.GetFeiShengLevel())
	if feiTemp == nil {
		return
	}
	giveExp := costExp * int64(feiTemp.GiveExpRatio) / common.MAX_RATE
	if giveExp <= 0 {
		return
	}

	for _, tpl := range s.GetAllPlayers() {
		if tpl.GetId() == pl.GetId() {
			continue
		}

		sanGongData := &sanGongData{
			plName:  pl.GetName(),
			giveExp: giveExp,
		}
		ctx := scene.WithPlayer(context.Background(), tpl)
		msg := message.NewScheduleMessage(onGiveExp, ctx, sanGongData, nil)
		tpl.Post(msg)
	}

	return
}

func onGiveExp(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	data := result.(*sanGongData)
	pl, ok := tpl.(player.Player)
	if !ok {
		return nil
	}
	feiShengManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	feiShengNum := feiShengManager.GetFeiShengReceiveNum()
	if feiShengNum >= constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFeiShengLimit) {
		playerlogic.SendSystemMessage(pl, lang.FeiShengLimit)
		return nil
	}
	feiShengManager.ReceiveFeiSheng()
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	addLevelReason := commonlog.LevelLogReasonFeiShengSanGongGive
	propertyManager.AddExp(data.giveExp, addLevelReason, addLevelReason.String())
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCFeiShengSanGongBroadcast(data.plName, data.giveExp)
	pl.SendMsg(scMsg)
	return nil
}

func init() {
	gameevent.AddEventListener(feishengeventtypes.EventTypePlayerSanGong, event.EventListenerFunc(playerSanGong))
}
