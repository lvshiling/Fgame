package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arena/arena"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	arenalogic "fgame/fgame/cross/arena/logic"
	"fgame/fgame/cross/arena/pbutil"
	arenascene "fgame/fgame/cross/arena/scene"
	gameevent "fgame/fgame/game/event"

	log "github.com/Sirupsen/logrus"
)

//竞技场场景结束
func arenaSceneEnd(target event.EventTarget, data event.EventData) (err error) {

	sd := target.(arenascene.ArenaSceneData)
	winnerTeam := sd.GetWinnerTeam()
	winnerTeamId := int64(0)
	if winnerTeam != nil {
		winnerTeamId = winnerTeam.GetTeam().GetTeamId()
	}

	log.WithFields(
		log.Fields{
			"team1":      sd.GetTeam1(),
			"team2":      sd.GetTeam2(),
			"winnerTeam": winnerTeamId,
		}).Infoln("arena:竞技场场景结束")
	arena.GetArenaService().ArenaEnd(sd, winnerTeamId)

	team1 := sd.GetTeam1()
	team2 := sd.GetTeam2()
	if winnerTeam != nil {
		team1Level := team1.GetCurrent()
		for _, mem := range team1.GetTeam().GetMemberList() {
			if mem.IsRobot() {
				continue
			}

			pl := sd.GetScene().GetPlayer(mem.GetPlayerId())
			if pl == nil {
				continue
			}

			win := true
			if winnerTeamId != team1.GetTeam().GetTeamId() {
				win = false
			}

			//发送获胜
			isArenaWin := pbutil.BuildISArenaWin(team1Level, win)
			pl.SendMsg(isArenaWin)
		}

		//
		team2Level := team2.GetCurrent()
		for _, mem := range team2.GetTeam().GetMemberList() {
			if mem.IsRobot() {
				continue
			}

			pl := sd.GetScene().GetPlayer(mem.GetPlayerId())
			if pl == nil {
				continue
			}

			win := true
			if winnerTeamId != team2.GetTeam().GetTeamId() {
				win = false
			}

			//发送获胜
			isArenaWin := pbutil.BuildISArenaWin(team2Level, win)
			pl.SendMsg(isArenaWin)
		}
	}
	//TODO 记录数据

	scArenaSceneEnd := pbutil.BuildSCArenaSceneEnd(winnerTeamId)
	arenalogic.BroadcastArenaTeam(team1, scArenaSceneEnd)
	arenalogic.BroadcastArenaTeam(team2, scArenaSceneEnd)

	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaSceneEnd, event.EventListenerFunc(arenaSceneEnd))
}
