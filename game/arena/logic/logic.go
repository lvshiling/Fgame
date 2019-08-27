package logic

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	arenaeventtypes "fgame/fgame/game/arena/event/types"
	"fgame/fgame/game/arena/pbutil"
	playerarena "fgame/fgame/game/arena/player"
	arenatemplate "fgame/fgame/game/arena/template"
	arenatypes "fgame/fgame/game/arena/types"
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

//TODO 修改为注册活动参加
//func ArenaMatch(pl player.Player, activityTemplate *gametemplate.ActivityTemplate) (flag bool, err error)
//3v3竞技开始匹配
func ArenaMatch(pl player.Player) (canMatch bool, err error) {
	teamId := pl.GetTeamId()
	if teamId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arena:玩家还没组队")
		playerlogic.SendSystemMessage(pl, lang.PlayerNotInTeam)
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"teamId":   teamId,
		}).Info("arena:玩家开始匹配")

	//判断活动时间
	now := global.GetGame().GetTimeService().Now()
	arenaConstantTemp := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate()
	if !arenaConstantTemp.IsOnArenaTime(now) {
		err = team.ErrorTeamInMatchNoActivityTime
		return
	}
	//组队匹配
	canMatch, err = team.GetTeamService().ArenaMatch(pl)
	if err != nil {
		return
	}
	return
}

//竞技发送匹配
func ArenaMatchSend(pl player.Player) {
	teamData := team.GetTeamService().GetTeam(pl.GetTeamId())
	if teamData == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Info("arena:竞技匹配队伍不存在")
		//退出竞技
		crosslogic.PlayerExitCross(pl)
		return
	}
	pList := make([]player.Player, 0, len(teamData.GetMemberList()))
	for _, mem := range teamData.GetMemberList() {
		p := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if p == nil {
			continue
		}
		pList = append(pList, p)
	}
	siArenaMatch := pbutil.BuildSIArenaMatch(pList)
	pl.SendCrossMsg(siArenaMatch)
}

//队伍开始匹配
func OnTeamArenaMatch(pl player.Player, teamObject *team.TeamObject) {
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"teamId":   teamObject.GetTeamId(),
		}).Info("arena:队伍开始匹配")

	//进入跨服
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeArena)

	//广播消息
	for _, mem := range teamObject.GetMemberList() {
		memPl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if memPl == nil {
			continue
		}

		scArenaMatchBroadcast := pbutil.BuildSCArenaMatchBroadcast()
		memPl.SendMsg(scArenaMatchBroadcast)
	}
}

//竞技停止匹配发送
func ArenaStopMatchSend(pl player.Player) {
	siArenaStopMatch := pbutil.BuildSIArenaStopMatch()
	pl.SendCrossMsg(siArenaStopMatch)
}

//3v3竞技停止匹配
func ArenaStopMatch(pl player.Player) (err error) {
	//组队停止匹配
	err = team.GetTeamService().ArenaStopMatch(pl)
	if err != nil {
		return
	}
	return
}

//队伍停止匹配
func OnTeamArenaStopMatch(pl player.Player, teamObject *team.TeamObject) {
	teamId := pl.GetTeamId()
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"teamId":   teamId,
		}).Info("arena:正在停止匹配中")

	//广播消息
	for _, mem := range teamObject.GetMemberList() {
		memPl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if memPl == nil {
			continue
		}

		scArenaStopMatchBroadcast := pbutil.BuildSCArenaStopMatchBroadcast()
		memPl.SendMsg(scArenaStopMatchBroadcast)
	}

}

//匹配到了
func OnTeamArenaMatched(pl player.Player, teamObject *team.TeamObject) {
	log.WithFields(
		log.Fields{
			"teamId": teamObject.GetTeamId(),
		}).Info("arena:匹配到了")

	scArenaMatchResult := pbutil.BuildSCArenaMatchResult(true)
	pl.SendMsg(scArenaMatchResult)
	//匹配成功
	for _, mem := range teamObject.GetMemberList() {
		tpl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if tpl == nil {
			continue
		}
		if tpl == pl {
			continue
		}
		ctx := scene.WithPlayer(context.Background(), tpl)
		scheduleMsg := message.NewScheduleMessage(onTeamArenaMatched, ctx, nil, nil)
		tpl.Post(scheduleMsg)
	}

	ctx := scene.WithPlayer(context.Background(), pl)
	scheduleMsg := message.NewScheduleMessage(onTeamCaptainArenaMatched, ctx, nil, nil)
	pl.Post(scheduleMsg)

}

//队伍匹配到了
func onTeamCaptainArenaMatched(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)

	now := global.GetGame().GetTimeService().Now()
	arenaConstantTemp := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate()
	endTime := arenaConstantTemp.GetEndTime(now)
	arenaManager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	arenaManager.EnterArena(endTime)
	crosslogic.CrossPlayerDataLogin(pl)

	//参加3v3活动
	gameevent.Emit(arenaeventtypes.EventTypeArenaJoin, pl, nil)
	return nil
}

//队伍匹配到了
func onTeamArenaMatched(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	scArenaMatchResult := pbutil.BuildSCArenaMatchResult(true)
	tpl.SendMsg(scArenaMatchResult)

	//进入竞技场
	arenaManager := tpl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	endTime := int64(0)
	arenaManager.EnterArena(endTime)
	//进入跨服
	crosslogic.PlayerEnterCross(tpl, crosstypes.CrossTypeArena)

	//参加3v3活动
	gameevent.Emit(arenaeventtypes.EventTypeArenaJoin, tpl, nil)
	return nil
}

//匹配失败
func OnTeamArenaMatchFailed(pl player.Player, teamObject *team.TeamObject) {
	log.WithFields(
		log.Fields{
			"teamId": teamObject.GetTeamId(),
		}).Info("arena:匹配失败")
	//匹配成功
	for _, mem := range teamObject.GetMemberList() {
		memPl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if memPl == nil {
			continue
		}

		scArenaMatchResult := pbutil.BuildSCArenaMatchResult(false)
		memPl.SendMsg(scArenaMatchResult)
	}
}

//匹配失败
func OnTeamArenaStopMatchOther(pl player.Player, teamObject *team.TeamObject) {
	log.WithFields(
		log.Fields{
			"teamId": teamObject.GetTeamId(),
		}).Info("arena:停止匹配其它原因")
	//匹配成功
	for _, mem := range teamObject.GetMemberList() {
		memPl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if memPl == nil {
			continue
		}

		scArenaMatchResult := pbutil.BuildSCArenaMatchResult(false)
		memPl.SendMsg(scArenaMatchResult)
	}

}

//gm选择四神
func GMSelectFourGod(pl player.Player, fourGodType arenatypes.FourGodType) (err error) {
	if !pl.IsCross() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arena:玩家没在跨服")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("arena:发送选择四神")
	csArenaSelectFourGod := pbutil.BuildCSArenaSelectFourGod(fourGodType)
	pl.SendCrossMsg(csArenaSelectFourGod)
	return nil
}

//竞技停止匹配发送
func GMArenaNextSend(pl player.Player) (err error) {
	csArenaNextMatch := pbutil.BuildCSArenaNextMatch()
	pl.SendCrossMsg(csArenaNextMatch)
	return nil
}
