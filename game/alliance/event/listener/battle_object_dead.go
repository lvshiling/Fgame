package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancetemplate "fgame/fgame/game/alliance/template"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptempalte "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

//生物死亡
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	//不是玩家
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	s := p.GetScene()
	if s == nil {
		return
	}

	attackId := data.(int64)
	so := s.GetSceneObject(attackId)
	if so == nil {
		return
	}
	attackPl, ok := so.(player.Player)
	if !ok {
		return
	}

	//仙盟成员死亡推送
	diedAllianceNotice(p, attackPl)

	// 城战死亡处理-腰牌
	diedOnAllianceWarYaoPai(p, attackId)
	// 城战死亡处理-积分
	diedOnAllianceWarPoint(p, attackPl)

	return
}

//广播帮派
func diedAllianceNotice(pl player.Player, attackPl player.Player) {
	al := alliance.GetAllianceService().GetAlliance(pl.GetAllianceId())
	if al == nil {
		return
	}

	noticeCd := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceMemberDiedNoticeCD))
	allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	lastNoticeTime := allianceManager.GetLastBeKilledNoticeTime()
	now := global.GetGame().GetTimeService().Now()
	elapse := now - lastNoticeTime
	//cd中
	if elapse < noticeCd {
		return
	}

	memName := coreutils.FormatColor(chattypes.ColorTypePlayerName, pl.GetName())
	posStr := coreutils.FormatColor(chattypes.ColorTypeBossMap, coreutils.FormatStrPosiotn(fmt.Sprintf("%.0f", pl.GetPos().X), fmt.Sprintf("%.0f", pl.GetPos().Z)))
	mapName := coreutils.FormatColor(chattypes.ColorTypeBossMap, pl.GetScene().MapTemplate().Name)
	attackName := coreutils.FormatColor(chattypes.ColorTypePlayerName, attackPl.GetName())
	pos := pl.GetPosition()
	args := []int64{int64(chattypes.ChatAllianceRescue), int64(funcopentypes.FuncOpenTypeAlliance), int64(pl.GetScene().MapId()), int64(pos.X), int64(pos.Y), int64(pos.Z)}
	link := coreutils.FormatLink(chattypes.ButtonTypeToRescue, args)

	format := lang.GetLangService().ReadLang(lang.AllianceMemberDeadNotice)
	content := fmt.Sprintf(format, memName, mapName, posStr, attackName, attackName, link)
	chatlogic.SystemBroadcastAlliance(al, chattypes.MsgTypeText, []byte(content))

	allianceManager.UpdateBeKillNoticeTime()
}

func diedOnAllianceWarYaoPai(p player.Player, attackId int64) {
	s := p.GetScene()
	//不是城战
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeChengZhan && s.MapTemplate().GetMapType() != scenetypes.SceneTypeHuangGong {
		return
	}

	allianceManager := p.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	lastKilledTime := allianceManager.GetLastBeKilledTime(attackId)
	now := global.GetGame().GetTimeService().Now()
	elapse := now - lastKilledTime
	warTemplate := alliancetemplate.GetAllianceTemplateService().GetWarTemplate()
	//cd中
	if elapse < int64(warTemplate.PlateTime) {
		return
	}

	//随机掉落
	flag := mathutils.RandomHit(common.MAX_RATE, int(warTemplate.PalteOdd))
	if !flag {
		return
	}
	allianceManager.KilledByAlliancePlayer(attackId)

	dropId := warTemplate.DropId
	dropTemplate := droptempalte.GetDropTemplateService().GetDropFromGroup(dropId)
	scenelogic.CustomItemDrop(s, p.GetPosition(), attackId, dropTemplate.ItemId, dropTemplate.RandomNum(), dropTemplate.RandomStack(), dropTemplate.ProtectedTime, dropTemplate.ExistTime)
	// //随机
	// randomYaoPai := int32(mathutils.RandomRange(int(warTemplate.PlateEachMin), int(warTemplate.PlateEachMax+1)))
	// //TODO
	// stack := int32(1)
	// protectTime := int32(common.MINUTE * 2)
	// existTime := int32(common.MINUTE * 5)
	// scenelogic.CustomItemDrop(s, p.GetPosition(), attackId, int32(constanttypes.YaoPai), randomYaoPai, stack, protectTime, existTime)
}

func diedOnAllianceWarPoint(p player.Player, attackPl player.Player) {
	s := p.GetScene()
	//不是城战
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeChengZhan && s.MapTemplate().GetMapType() != scenetypes.SceneTypeHuangGong {
		return
	}

	ctx := scene.WithPlayer(context.Background(), attackPl)
	msg := message.NewScheduleMessage(onAllianceWarKillPlayer, ctx, nil, nil)
	attackPl.Post(msg)
}

func onAllianceWarKillPlayer(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)

	warTemplate := alliancetemplate.GetAllianceTemplateService().GetWarTemplate()
	allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	curPoint := allianceManager.AddWarPoint(warTemplate.KillJiFen)

	scMsg := pbutil.BuildSCAllianceSceneWarPointChanged(curPoint)
	pl.SendMsg(scMsg)
	return nil
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
