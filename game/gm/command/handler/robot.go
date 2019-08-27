package handler

import (
	"fgame/fgame/common/lang"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/merge/merge"
	playerlogic "fgame/fgame/game/player/logic"
	propertylogic "fgame/fgame/game/property/logic"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/robot/robot"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"math"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeRobot, command.CommandHandlerFunc(handleRobot))
}

func handleRobot(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:添加机器人")
	if len(c.Args) <= 0 {
		log.Warn("gm:添加机器人,参数少于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	numStr := c.Args[0]
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"num":   numStr,
			}).Warn("gm:添加机器人,num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	showServerId := merge.GetMergeService().GetMergeTime() != 0
	//TODO 修改物品数量
	err = addRobot(pl, num, showServerId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"num":   numStr,
			}).Warn("gm:添加机器人,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":  pl.GetId(),
			"num": numStr,
		}).Debug("gm:添加机器人,完成")
	return
}

func addRobot(pl scene.Player, num int64, showServerId bool) (err error) {
	if pl.GetScene() == nil {
		return
	}
	mapTemplate := pl.GetScene().MapTemplate()
	index := int32(0)
	// sdkList := center.GetCenterService().GetSdkList()
	// robotQuestTemplate := scenetemplate.GetSceneTemplateService().GetRobotQuestTemplate(sdkList, pl.GetScene().MapId())

	for i := 0; i < int(num); i++ {
		// properties := robotQuestTemplate.RandomProperty()
		// hit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeInitalHit)
		// properties[propertytypes.BattlePropertyTypeHit] = int64(hit)
		// moveSpeed := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeInitMoveSpeed)
		// properties[propertytypes.BattlePropertyTypeMoveSpeed] = int64(moveSpeed)
		tempProperties := pl.GetAllBattleProperties()
		properties := make(map[propertytypes.BattlePropertyType]int64)
		for k, v := range tempProperties {
			properties[propertytypes.BattlePropertyType(k)] = v
		}
		power := propertylogic.CulculateAllForce(properties)
		showServerId := merge.GetMergeService().GetMergeTime() != 0
		//添加机器人
		// p := robot.GetRobotService().CreateQuestRobot(robotQuestTemplate.QuestBeginId, robotQuestTemplate.QuestOverId, properties, power, showServerId)
		// if p == nil {
		// 	return
		// }
		nextIndex, pos := getRobotPosition(mapTemplate, pl.GetPos(), index)
		robotPlayer := robot.GetRobotService().CreateClientTestRobot(properties, power, showServerId)
		robotPlayer.SetEnterPos(pos)
		pl.GetScene().AddSceneObject(robotPlayer)
		index = nextIndex
	}
	return
}

const (
	maxIndex     = 9
	elapseRadius = 1
	maxCycle     = 100
)

func getRobotPosition(mapTemplate *gametemplate.MapTemplate, centerPosition coretypes.Position, index int32) (nextIndex int32, pos coretypes.Position) {
	nextIndex = index
	//以防卡死
	for i := 0; i < maxCycle; i++ {
		elapase := nextIndex/maxIndex + 1
		tempIndex := index % maxIndex
		radius := float64(elapseRadius * elapase)
		//获取
		angle := float64(tempIndex) / float64(maxIndex) * math.Pi * 2
		destPos := coretypes.Position{
			X: centerPosition.X + math.Cos(angle)*radius,
			Y: centerPosition.Y,
			Z: centerPosition.Z + math.Sin(angle)*radius,
		}

		if mapTemplate.GetMap().IsMask(destPos.X, destPos.Z) {
			destPos.Y = mapTemplate.GetMap().GetHeight(destPos.X, destPos.Z)
			return nextIndex + 1, destPos
		}
		nextIndex += 1
	}
	return nextIndex, centerPosition
}
