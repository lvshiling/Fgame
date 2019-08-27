package teamcopy

import (
	scene "fgame/fgame/cross/teamcopy/scene"
	arenatemplate "fgame/fgame/game/arena/template"
	"fgame/fgame/game/common/common"
	propertylogic "fgame/fgame/game/property/logic"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/robot/robot"
	teamtypes "fgame/fgame/game/team/types"
	teamcopytemplate "fgame/fgame/game/teamcopy/template"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"math"
	"math/rand"
)

func createTeamCopyTeamWithRobot(memList []*BattleTeamMember) (t *scene.TeamObject) {
	num := len(memList)
	if num <= 0 || num > teamtypes.TeamMaxNum {
		return
	}
	needComplement := teamtypes.TeamMaxNum - num
	robotMemList := make([]*BattleTeamMember, 0, needComplement)

	allProperties := make(map[int32]int64)

	teamPurpose := memList[0].GetTeamPurpose()
	for _, mem := range memList {
		for typ, val := range mem.GetBattleProperties() {
			allProperties[typ] += val
		}

	}

	//随机属性
	avgProperties := make(map[int32]int64)
	for typ, val := range allProperties {
		avgProperties[typ] = int64(math.Ceil(float64(val) / float64(num)))
	}
	randonServer := rand.Intn(maxServer) + 1
	serverId := int32(randonServer)

	teamCopyTemplate := teamcopytemplate.GetTeamCopyTemplateService().GetTeamCopyTempalte(teamPurpose)
	for i := 0; i < needComplement; i++ {
		robotProperties := make(map[propertytypes.BattlePropertyType]int64)
		//TODO 假人属性要单独弄一套
		percent := float64(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RandomPropertyPercent()) / float64(common.MAX_RATE)
		for typ, val := range avgProperties {
			if typ != int32(propertytypes.BattlePropertyTypeMaxHP) && typ != int32(propertytypes.BattlePropertyTypeAttack) && typ != int32(propertytypes.BattlePropertyTypeDefend) {
				robotProperties[propertytypes.BattlePropertyType(typ)] = val
			} else {
				robotProperties[propertytypes.BattlePropertyType(typ)] = int64(math.Ceil(float64(val) * percent))
			}

		}
		force := propertylogic.CulculateAllForce(robotProperties)

		reliveTime := teamCopyTemplate.RandomRevive()
		robotPlayer := robot.GetRobotService().CreateTeamCopyRobot(serverId, robotProperties, reliveTime, force)
		robotMemList = append(robotMemList, convertMemberFromRobotPlayer(robotPlayer, teamPurpose))
	}
	t = combinePlayerList(teamPurpose, memList, robotMemList)
	return
}

func combinePlayerList(purpose teamtypes.TeamPurposeType, aMemListOfList ...[]*BattleTeamMember) (t *scene.TeamObject) {
	totalNum := 0
	memListOfList := make([][]*scene.TeamMemberObject, 0, len(aMemListOfList))
	for _, aMemList := range aMemListOfList {
		totalNum += len(aMemList)
		memListOfList = append(memListOfList, convertToTeamMemberObjectList(aMemList))
	}
	if totalNum != teamtypes.TeamMaxNum {
		panic(fmt.Errorf("teamcopy:组队成员不等于3"))
	}
	teamId, _ := idutil.GetId()
	return scene.CreateTeamWithMemberLists(teamId, purpose, memListOfList...)
}
