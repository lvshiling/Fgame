package logic

import (
	teamcopyscene "fgame/fgame/cross/teamcopy/scene"
	teamcopy "fgame/fgame/cross/teamcopy/teamcopy"
	robotlogic "fgame/fgame/game/robot/logic"
	"fgame/fgame/game/robot/robot"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	teamcopytemplate "fgame/fgame/game/teamcopy/template"

	"github.com/golang/protobuf/proto"
)

func OnTeamCopyRobotEnterScene(sd teamcopyscene.TeamCopySceneData) {
	teamObj := sd.GetTeamObj()
	s := sd.GetScene()
	for _, mem := range teamObj.GetMemberList() {
		if !mem.IsRobot() {
			continue
		}
		//获取机器人
		robotPl := robot.GetRobotService().GetRobot(mem.GetPlayerId())
		if robotPl == nil {
			continue
		}
		robotPl.SetArenaTeam(teamObj.GetTeamId(), teamObj.GetTeamName(), teamObj.GetTeamPurpose())
		PlayerEnterTeamCopyScene(s, robotPl)
	}
}

//进入组队副本
func PlayerEnterTeamCopyScene(s scene.Scene, p scene.Player) {
	// 获取出生位置
	sceneData := s.SceneDelegate().(teamcopyscene.TeamCopySceneData)
	teamObj := sceneData.GetTeamObj()
	purpose := teamObj.GetTeamPurpose()
	bornPos, flag := teamcopytemplate.GetTeamCopyTemplateService().GetTeamCopyBorn(purpose)
	if !flag {
		return
	}
	scenelogic.AsyncPlayerEnterScene(p, s, bornPos)
}

func BroadcastTeamCopy(sd teamcopyscene.TeamCopySceneData, msg proto.Message) {
	t := sd.GetTeamObj()
	for _, mem := range t.GetMemberList() {
		if mem.IsRobot() {
			continue
		}
		if mem.GetStatus() == teamcopyscene.MemberStatusOnline {
			pl := sd.GetScene().GetPlayer(mem.GetPlayerId())
			if pl == nil {
				continue
			}
			pl.SendMsg(msg)
		}
	}
}

func BroadcastTeamCopyExclude(sd teamcopyscene.TeamCopySceneData, p scene.Player, msg proto.Message) {
	t := sd.GetTeamObj()
	for _, mem := range t.GetMemberList() {
		if mem.IsRobot() {
			continue
		}
		if mem.GetPlayerId() == p.GetId() {
			continue
		}
		if mem.GetStatus() == teamcopyscene.MemberStatusOnline {
			pl := sd.GetScene().GetPlayer(mem.GetPlayerId())
			if pl == nil {
				continue
			}
			pl.SendMsg(msg)
		}
	}
}

//机器人退出
func RobotExit(pl scene.RobotPlayer) {
	teamcopy.GetTeamCopyService().TeamCopyMemeberExit(pl)
	robotlogic.RemoveRobot(pl)
}

//复活
func Reborn(sd teamcopyscene.TeamCopySceneData, pl scene.Player) {
	teamObj := sd.GetTeamObj()
	purpose := teamObj.GetTeamPurpose()
	pos, flag := teamcopytemplate.GetTeamCopyTemplateService().GetTeamCopyBorn(purpose)
	if !flag {
		return
	}
	pl.Reborn(pos)
	sd.AddReliveTime(pl)
}
