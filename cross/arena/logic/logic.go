package logic

import (
	"fgame/fgame/cross/arena/arena"
	"fgame/fgame/cross/arena/pbutil"
	arenascene "fgame/fgame/cross/arena/scene"
	"fgame/fgame/cross/player/player"
	arenatemplate "fgame/fgame/game/arena/template"
	robotlogic "fgame/fgame/game/robot/logic"
	"fgame/fgame/game/robot/robot"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	teamtypes "fgame/fgame/game/team/types"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

//竞技场匹配成功
func OnArenaMatched(arenaScene scene.Scene, team1 *arenascene.ArenaTeam, team2 *arenascene.ArenaTeam) (err error) {
	if team1.GetCurrent() == 0 {
		team1Id := team1.GetTeam().GetTeamId()
		team1Name := team1.GetTeam().GetTeamName()
		//机器人
		for _, mem := range team1.GetTeam().GetMemberList() {
			//获取机器人
			robotPl := robot.GetRobotService().GetRobot(mem.GetPlayerId())
			if robotPl == nil {
				continue
			}
			robotPl.SetArenaTeam(team1Id, team1Name, teamtypes.TeamPurposeTypeNormal)
			PlayerEnterArenaScene(arenaScene, robotPl)
		}
	} else {
		//重置复活次数
		isMsg := pbutil.BuildISArenaResetReliveTimes()
		//直接跳转场景
		for _, mem := range team1.GetTeam().GetMemberList() {
			//获取在线玩家
			memPl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
			if memPl != nil {
				PlayerEnterArenaScene(arenaScene, memPl)
				memPl.SendMsg(isMsg)
				continue
			}
			//获取机器人
			robotPl := robot.GetRobotService().GetRobot(mem.GetPlayerId())
			if robotPl != nil {
				PlayerEnterArenaScene(arenaScene, robotPl)
				continue
			}
		}
	}

	if team2.GetCurrent() == 0 {
		team2Id := team2.GetTeam().GetTeamId()
		team2Name := team2.GetTeam().GetTeamName()
		//机器人
		for _, mem := range team2.GetTeam().GetMemberList() {
			//获取机器人
			robotPl := robot.GetRobotService().GetRobot(mem.GetPlayerId())
			if robotPl == nil {
				continue
			}
			robotPl.SetArenaTeam(team2Id, team2Name, teamtypes.TeamPurposeTypeNormal)
			PlayerEnterArenaScene(arenaScene, robotPl)
		}
	} else {
		//重置复活次数
		isMsg := pbutil.BuildISArenaResetReliveTimes()
		for _, mem := range team2.GetTeam().GetMemberList() {
			//获取在线玩家
			memPl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
			if memPl != nil {
				PlayerEnterArenaScene(arenaScene, memPl)
				memPl.SendMsg(isMsg)
				continue
			}
			//获取机器人
			robotPl := robot.GetRobotService().GetRobot(mem.GetPlayerId())
			if robotPl != nil {
				PlayerEnterArenaScene(arenaScene, robotPl)
				continue
			}
		}
	}

	return nil
}

//进入竞技场
func PlayerEnterArenaScene(arenaScene scene.Scene, p scene.Player) {
	arenaSceneData := arenaScene.SceneDelegate().(arenascene.ArenaSceneData)
	//竞技场景
	bornPos := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().GetTeam1BornPos()
	if arenaSceneData.GetTeam2().GetTeam().GetTeamId() == p.GetTeamId() {
		bornPos = arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().GetTeam2BornPos()
	}

	scenelogic.AsyncPlayerEnterScene(p, arenaScene, bornPos)
}

//竞技场匹配下一场
func OnArenaNextMatch(t *arenascene.ArenaTeam) (err error) {
	scArenaNextMatchBroadcast := pbutil.BuildSCArenaNextMatchBroadcast()
	for _, mem := range t.GetTeam().GetMemberList() {
		memPl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if memPl == nil {
			continue
		}
		memPl.SendMsg(scArenaNextMatchBroadcast)
	}

	return nil
}

//四神进入
func OnTeamFourGodEnter(s scene.Scene, teamObject *arenascene.ArenaTeam) {
	log.WithFields(
		log.Fields{
			"teamId": teamObject.GetTeam().GetTeamId(),
		}).Info("arena:四神进入")

	bornPos := s.MapTemplate().GetBornPos()
	//进入四圣兽
	for _, mem := range teamObject.GetTeam().GetMemberList() {
		pl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if pl == nil {
			continue
		}
		scenelogic.AsyncPlayerEnterScene(pl, s, bornPos)
	}
}

func BroadcastArenaTeam(t *arenascene.ArenaTeam, msg proto.Message) {
	for _, mem := range t.GetTeam().GetMemberList() {
		if mem.GetStatus() == arenascene.MemberStatusOnline {
			pl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
			if pl == nil {
				continue
			}
			pl.SendMsg(msg)
		}
	}
}

//采集经验数打断
func CollectInterrupt(pl scene.Player) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeArenaShengShou {
		return
	}
	sd := s.SceneDelegate()
	if sd == nil {
		return
	}

	fourGodSceneData, ok := sd.(arenascene.FourGodSceneData)
	if !ok {
		return
	}
	fourGodSceneData.ClearCollect(pl)
}

//复活
func Reborn(pl scene.Player) {
	pl.Reborn(pl.GetPos())
	reliveTime := pl.GetArenaReliveTime() + 1
	//扣除复活次数
	pl.SetArenaReliveTime(reliveTime)
}

//机器人退出
func RobotExit(pl scene.RobotPlayer) {
	arena.GetArenaService().ArenaMemeberExit(pl)
	robotlogic.RemoveRobot(pl)
}
