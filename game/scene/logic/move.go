package logic

import (
	"fgame/fgame/common/lang"
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	playerlogic "fgame/fgame/game/player/logic"
	propertytypes "fgame/fgame/game/property/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

const (
	MAX_MOVE = 50
)

func HandleObjectMove(pl scene.Player, uid int64, pos coretypes.Position, curPos coretypes.Position, moveSpeed float64, angle float64, moveType scenetypes.MoveType, flag bool) {
	if uid == pl.GetId() {
		HandlePlayerMove(pl, pos, curPos, moveSpeed, angle, moveType, flag)
		return
	}
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"uId":       uid,
				"pos":       pos,
				"angle":     angle,
				"moveSpeed": moveSpeed,
				"moveType":  moveType,
			}).Warn("scene:处理对象移动消息,场景为空")
		return
	}
	lingTong := s.GetLingTong(uid)
	if lingTong == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"uId":       uid,
				"pos":       pos,
				"angle":     angle,
				"moveSpeed": moveSpeed,
				"moveType":  moveType,
			}).Warn("scene:处理对象移动消息,灵童不存在")
		return
	}
	if lingTong.GetOwner() != pl {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"uId":       uid,
				"pos":       pos,
				"angle":     angle,
				"moveSpeed": moveSpeed,
				"moveType":  moveType,
			}).Warn("scene:处理对象移动消息,灵童不是他的")
		return
	}
	HandleLingTongMove(lingTong, pos, curPos, moveSpeed, angle, moveType, flag)
}

var (
	clientFixPosition = coretypes.Position{}
)

func HandlePlayerMove(pl scene.Player, pos coretypes.Position, curPos coretypes.Position, moveSpeed float64, angle float64, moveType scenetypes.MoveType, moveFlag bool) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"pos":       pos,
				"curPos":    curPos,
				"angle":     angle,
				"moveSpeed": moveSpeed,
				"moveType":  moveType,
			}).Warn("scene:处理对象移动消息,场景为空")
		return
	}

	if !s.MapTemplate().GetMap().IsWalkable(curPos.X, curPos.Z) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"pos":       pos,
				"curPos":    curPos,
				"angle":     angle,
				"moveSpeed": moveSpeed,
				"moveType":  moveType,
				"mapId":     s.MapId(),
			}).Warn("scene:处理对象移动消息,移动到掩码外")
		if pos.IsEqual(clientFixPosition) {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"pos":       pos,
					"curPos":    curPos,
					"angle":     angle,
					"moveSpeed": moveSpeed,
					"moveType":  moveType,
					"mapId":     s.MapId(),
				}).Info("客户端跳跃动画拉回")
			FixPosition(pl, pl.GetPosition())
		}

		return
	}

	moveSpeed = float64(pl.GetBattleProperty(propertytypes.BattlePropertyTypeMoveSpeed)) / float64(common.MILL_METER)
	//验证移动
	if coreutils.Distance(curPos, pl.GetPosition()) > MAX_MOVE {
		log.WithFields(log.Fields{
			"id":        pl.GetId(),
			"pos":       pos,
			"curPos":    curPos,
			"angle":     angle,
			"moveSpeed": moveSpeed,
			"moveType":  moveType,
			"mapId":     s.MapId(),
		}).Warn("scene: 移动超过验证")
		//送回原位置
		FixPosition(pl, pl.GetPosition())
		return
	}
	//判断是否在限制区内
	if !s.MapTemplate().IsInLimitArea(curPos) {
		log.WithFields(log.Fields{
			"id":        pl.GetId(),
			"pos":       pos,
			"curPos":    curPos,
			"angle":     angle,
			"moveSpeed": moveSpeed,
			"moveType":  moveType,
			"mapId":     s.MapId(),
		}).Warn("scene: 移动超过限制区")
		playerlogic.SendSystemMessage(pl, lang.SceneMoveOutside)
		FixPosition(pl, pl.GetPosition())
		return
	}
	h := scene.GetCheckMoveHandler(s.MapTemplate().GetMapType())
	if h != nil {
		flag, fixPosition := h.CheckMove(pl, curPos)
		if !flag {
			FixPosition(pl, fixPosition)
			return
		}
	}
	MoveInternal(pl, pos, curPos, angle, moveSpeed, moveType, false, moveFlag)

	lingTong := pl.GetLingTong()
	if lingTong != nil && !pl.IsLingTongHidden() {
		if !CheckIfLingTongAndPlayerSameScene(lingTong) {
			return
		}
		exitDistance := s.MapTemplate().GetMapType().GetExitDistance()
		if coreutils.Distance(lingTong.GetPosition(), pl.GetPosition()) > exitDistance {
			FixPosition(lingTong, pl.GetPosition())
		}
	}
}

func HandleLingTongMove(lingTong scene.LingTong, pos coretypes.Position, curPos coretypes.Position, moveSpeed float64, angle float64, moveType scenetypes.MoveType, moveFlag bool) {
	s := lingTong.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId":  lingTong.GetOwner().GetId(),
				"pos":       pos,
				"curPos":    curPos,
				"angle":     angle,
				"moveSpeed": moveSpeed,
				"moveType":  moveType,
			}).Warn("scene:处理灵童移动消息,场景为空")
		return
	}
	p := lingTong.GetOwner()
	if p == nil {
		log.WithFields(
			log.Fields{
				"playerId":  lingTong.GetOwner().GetId(),
				"pos":       pos,
				"curPos":    curPos,
				"angle":     angle,
				"moveSpeed": moveSpeed,
				"moveType":  moveType,
			}).Warn("scene:处理灵童移动消息,场景为空")
		return
	}
	//验证移动
	if coreutils.Distance(curPos, lingTong.GetPosition()) > MAX_MOVE {
		log.WithFields(log.Fields{
			"playerId":  lingTong.GetOwner().GetId(),
			"pos":       pos,
			"curPos":    curPos,
			"angle":     angle,
			"moveSpeed": moveSpeed,
			"moveType":  moveType,
			"mapId":     s.MapId(),
		}).Warn("scene: 移动超过验证")
		//送回原位置
		FixPosition(lingTong, lingTong.GetPosition())
		return
	}
	moveSpeed = float64(p.GetBattleProperty(propertytypes.BattlePropertyTypeMoveSpeed)) * float64(1.5) / float64(common.MILL_METER)
	MoveInternal(lingTong, pos, curPos, angle, moveSpeed, moveType, false, moveFlag)
	log.WithFields(
		log.Fields{
			"playerId":  lingTong.GetOwner().GetId(),
			"pos":       pos,
			"curPos":    curPos,
			"angle":     angle,
			"moveSpeed": moveSpeed,
			"moveType":  moveType,
		}).Debug("scene:处理灵童移动消息")
}

func Move(bo scene.BattleObject, pos coretypes.Position, angle float64, moveSpeed float64, moveType scenetypes.MoveType, attackMove bool, moveFlag bool) {
	MoveInternal(bo, pos, pos, angle, moveSpeed, moveType, attackMove, moveFlag)
}

func MoveInternal(bo scene.BattleObject, pos coretypes.Position, curPos coretypes.Position, angle float64, moveSpeed float64, moveType scenetypes.MoveType, attackMove bool, moveFlag bool) {
	log.WithFields(log.Fields{
		"id":              bo.GetId(),
		"sceneObjectType": bo.GetSceneObjectType(),
		"pos":             pos,
		"curPos":          curPos,
		"angle":           angle,
		"moveSpeed":       moveSpeed,
		"moveType":        moveType,
	}).Debug("scene: 移动")

	if bo.GetScene() == nil {
		log.WithFields(log.Fields{
			"id":              bo.GetId(),
			"sceneObjectType": bo.GetSceneObjectType(),
			"pos":             pos,
			"angle":           angle,
			"moveSpeed":       moveSpeed,
			"moveType":        moveType,
		}).Error("scene: 移动")
		return
	}
	//保存数据
	bo.Move(curPos, angle)
	//保存aoi数据
	bo.GetScene().Move(bo, curPos)
	//TODO 优化
	for _, obj := range bo.GetNeighbors() {
		switch nei := obj.(type) {
		case scene.BattleObject:
			nei.OnMove(bo, curPos, angle)
			break
		}
	}

	//发送事件
	if !attackMove {
		msg := pbutil.BuildSCObjectMove(bo, pos, moveSpeed, angle, moveType, moveFlag)
		BroadcastNeighborIncludeSelf(bo, msg)
		gameevent.Emit(sceneeventtypes.EventTypeBattleObjectMove, bo, nil)
	}
}

//被攻击位移
func AttackedMove(bo scene.BattleObject, pos coretypes.Position, angle float64, moveSpeed float64, stopTime float64) {
	bo.AttackedMove(pos, angle, moveSpeed, stopTime)
	//保存aoi数据
	bo.GetScene().Move(bo, pos)

	for _, obj := range bo.GetNeighbors() {
		switch nei := obj.(type) {
		case scene.BattleObject:
			nei.OnMove(bo, pos, angle)
			break
		}
	}
	msg := pbutil.BuildSCObjectMove(bo, pos, moveSpeed, angle, scenetypes.MoveTypeHit, false)
	BroadcastNeighborIncludeSelf(bo, msg)
}

//移动到
func PlayerMoveThroughPortal(pl scene.Player, portal *gametemplate.PortalTemplate, mapId int32, destPos coretypes.Position) bool {
	//玩家场景为空
	s := pl.GetScene()
	if s == nil {
		return false
	}

	playerMapId := s.MapId()
	//同一个场景
	if playerMapId == mapId {
		// if IfMoveThroughPortal(s, pl.GetPosition(), portal, destPos) {
		// 	pl.SetDestPosition(destPos)
		// 	return true
		// }
		flag := pl.SetDestPosition(destPos)
		if flag {
			return true
		}
		FixPosition(pl, destPos)
		return true
	}
	//传送阵是空的
	if portal != nil {
		portalSceneTemplate := scenetemplate.GetSceneTemplateService().GetPortalSceneTemplate(int32(portal.TemplateId()))
		if portalSceneTemplate != nil {
			//玩家和传送阵同一个地图
			if playerMapId == portalSceneTemplate.SceneID {
				if coreutils.Distance(pl.GetPosition(), portalSceneTemplate.GetPos()) <= float64(common.MIN_DISTANCE_ERROR) {
					return PlayerEnterPortal(pl, int32(portal.TemplateId()))
				}
				//设置目的地传送点
				pl.SetDestPosition(portalSceneTemplate.GetPos())
				return true
			}
		}
	}

	//进入场景
	targetScene := scene.GetSceneService().GetWorldSceneByMapId(mapId)
	if targetScene == nil {
		return false
	}
	//直接飞
	return PlayerEnterScene(pl, targetScene, targetScene.MapTemplate().GetBornPos())

}

func PlayerEnterPortal(pl scene.Player, portalId int32) bool {
	portalTemplate := scenetemplate.GetSceneTemplateService().GetPortal(portalId)
	if portalTemplate == nil {
		return false
	}
	originS := pl.GetScene()
	if originS != nil {
		//同一个场景
		if int32(originS.MapTemplate().TemplateId()) == portalTemplate.MapId {
			FixPosition(pl, portalTemplate.GetPosition())
			return true
		}
	}
	PlayerEnterMapWithPortal(pl, portalTemplate.MapId, portalTemplate.GetPosition())
	return true
}

//移动到
func MoveToPortal(pl scene.Player, portal *gametemplate.PortalTemplate) bool {
	s := pl.GetScene()
	if s == nil {
		return false
	}

	portalSceneTemplate := scenetemplate.GetSceneTemplateService().GetPortalSceneTemplate(int32(portal.TemplateId()))
	if portalSceneTemplate == nil {
		return false
	}

	portalPos := portalSceneTemplate.GetPos()
	if coreutils.Distance(pl.GetPosition(), portalPos) <= float64(common.MIN_DISTANCE_ERROR) {
		return PlayerEnterPortal(pl, int32(portal.TemplateId()))
	}
	return pl.SetDestPosition(portalPos)
}
