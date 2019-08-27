package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	propertylogic "fgame/fgame/game/property/logic"
	propertytypes "fgame/fgame/game/property/types"
	robotlogic "fgame/fgame/game/robot/logic"
	"fgame/fgame/game/robot/robot"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"

	log "github.com/Sirupsen/logrus"
)

//机器人检测
func sceneRobotCheck(target event.EventTarget, data event.EventData) (err error) {
	s := target.(scene.Scene)
	sdkList := center.GetCenterService().GetSdkList()
	robotQuestTemplate := scenetemplate.GetSceneTemplateService().GetRobotQuestTemplate(sdkList, s.MapId())
	if robotQuestTemplate == nil {
		return
	}
	//玩家数量
	playerNum := len(s.GetAllPlayers())
	//不加机器人
	if playerNum <= 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	elapse := now - s.GetLastRobotTime()
	if elapse < robotQuestTemplate.GetRefreshTime() {
		return
	}
	s.SetLastRobotTime(now)
	//获取机器人数量
	robotNum := s.GetNumOfRobot()
	playerLimitCount := robotQuestTemplate.GetPlayerLimitCount()
	if playerNum >= int(playerLimitCount) || robotNum >= int32(playerNum) {
		if robotNum <= 0 {
			return
		}
		needRemove := robotNum
		if robotNum < int32(playerNum) {
			needRemove = playerLimitCount - int32(playerNum)
			if robotNum <= needRemove {
				needRemove = robotNum
			}
		}

		log.WithFields(
			log.Fields{
				"mapId": s.MapId(),
			}).Info("scene:移除机器人")
		i := int32(0)
		for _, r := range s.GetAllQuestRobots() {
			rp := r.(scene.RobotPlayer)
			robotlogic.RemoveRobot(rp)
			i++
			if i >= needRemove {
				return
			}
		}
		return
	}

	properties := robotQuestTemplate.RandomProperty()

	hit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeInitalHit)
	properties[propertytypes.BattlePropertyTypeHit] = int64(hit)
	moveSpeed := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeInitMoveSpeed)
	properties[propertytypes.BattlePropertyTypeMoveSpeed] = int64(moveSpeed)
	power := propertylogic.CulculateAllForce(properties)
	showServerId := merge.GetMergeService().GetMergeTime() != 0
	questBeginId := robotQuestTemplate.GetQuestBeginId()
	questOverId := robotQuestTemplate.GetQuestOverId()
	//添加机器人
	p := robot.GetRobotService().CreateQuestRobot(questBeginId, questOverId, properties, power, showServerId)
	if p == nil {
		return
	}
	log.WithFields(
		log.Fields{
			"mapId": s.MapId(),
		}).Info("scene:添加机器人")
	p.SetEnterPos(s.MapTemplate().GetBornPos())
	s.AddSceneObject(p)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeSceneRobotCheck, event.EventListenerFunc(sceneRobotCheck))
}
