package battle

import (
	"fgame/fgame/core/nav"
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"

	astar "github.com/beefsack/go-astar"
)

const (
	moveTime = 120
)

type MoveAction struct {
	bo                scene.BattleObject
	destPos           coretypes.Position
	lastTime          int64
	pause             bool
	paths             []astar.Pather
	currentIndex      int32
	portalIndex       int32
	lastHeartbeatTime int64
}

func (m *MoveAction) SetDestPosition(destPos coretypes.Position) bool {
	s := m.bo.GetScene()
	if s == nil {
		return false
	}

	found, portalIndex, paths := generateNavThroughPortal(s, m.bo.GetPosition(), destPos)
	if !found {
		return false
	}
	m.paths = paths
	m.currentIndex = 1
	m.portalIndex = portalIndex
	m.lastTime = global.GetGame().GetTimeService().Now()
	m.destPos = destPos
	m.pause = false
	return true
}

func (m *MoveAction) PauseMove() {
	m.pause = true
}

func (m *MoveAction) Heartbeat() {
	now := global.GetGame().GetTimeService().Now()
	if now-m.lastHeartbeatTime < moveTime {
		return
	}
	m.lastHeartbeatTime = now
	if m.pause {
		return
	}

	defer func() {
		m.lastTime = now
	}()
	if m.bo.GetBattleLimit()&scenetypes.BattleLimitTypeMove.Mask() != 0 {
		return
	}
	speed := float64(0.0)
	switch tbo := m.bo.(type) {
	case scene.LingTong:
		if tbo.GetOwner() != nil {
			speed = float64(tbo.GetOwner().GetBattleProperty(propertytypes.BattlePropertyTypeMoveSpeed)) / common.MILL_METER
		}
		break
	default:
		speed = float64(m.bo.GetBattleProperty(propertytypes.BattlePropertyTypeMoveSpeed)) / common.MILL_METER

		break
	}

	if speed <= 0 {
		return
	}

	s := m.bo.GetScene()

	//TODO 修改
	remainTime := float64(now - m.lastTime)

	boPosition := m.bo.GetPosition()
	var nextPos coretypes.Position
	useDest := false
	if len(m.paths) == 0 {
		nextPos = m.destPos
		useDest = true
	} else {
		previousPos := m.bo.GetPosition()

		for {
			if int(m.currentIndex+1) >= len(m.paths) {
				useDest = true
				break
			}
			index := int32(len(m.paths) - int(m.currentIndex) - 1)
			tile := m.paths[index].(*nav.Tile)

			if index == m.portalIndex {
				// fmt.Println("瞬间移动")
				nextTile := m.paths[index-1].(*nav.Tile)
				destX, destZ := s.GetWorld().GetPositionForTile(nextTile)
				nextPos = coretypes.Position{X: float64(destX), Z: float64(destZ)}
				nextPos.Y = s.MapTemplate().GetMap().GetHeight(nextPos.X, nextPos.Z)
				gameevent.Emit(battleeventtypes.EventTypeBattleObjectMoveFix, m.bo, nextPos)
				m.currentIndex += 2
				return
			}

			destX, destZ := s.GetWorld().GetPositionForTile(tile)
			nextPos = coretypes.Position{X: float64(destX), Z: float64(destZ)}

			originDistance := coreutils.Distance(previousPos, nextPos)
			needTime := originDistance / speed * float64(common.SECOND)
			if remainTime <= needTime {
				break
			} else {
				previousPos = nextPos
				remainTime -= needTime
				m.currentIndex++
			}
		}
		boPosition = previousPos
	}
	if useDest {
		nextPos = m.destPos
	}
	//移动的距离
	distance := (float64(remainTime) / float64(common.SECOND)) * speed

	//设置角度
	angle := coreutils.GetAngle(m.bo.GetPosition(), nextPos)
	originDistance := coreutils.Distance(boPosition, nextPos)

	if distance > originDistance {
		nextPos.Y = s.MapTemplate().GetMap().GetHeight(nextPos.X, nextPos.Z)
	}
	newPos := nextPos
	lerp := distance / originDistance
	if lerp < 1 {
		//获取差值
		newPos = coreutils.Lerp(boPosition, nextPos, lerp)
		angle = coreutils.GetAngle(m.bo.GetPosition(), newPos)
	}
	//获取位置
	newPos.Y = s.MapTemplate().GetMap().GetHeight(newPos.X, newPos.Z)

	eventData := battleeventtypes.CreateBattleObjectMoveTriggerEventData(newPos, angle, speed)
	gameevent.Emit(battleeventtypes.EventTypeBattleObjectMoveTrigger, m.bo, eventData)

	if coreutils.DistanceSquare(m.bo.GetPosition(), m.destPos) <= common.MIN_DISTANCE_SQUARE_ERROR {
		m.pause = true
	}
}

func (m *MoveAction) IsMove() bool {
	return !m.pause
}

func CreateMoveAction(bo scene.BattleObject) *MoveAction {
	m := &MoveAction{}
	m.bo = bo
	m.pause = true
	return m
}

func generateNavThroughPortal(s scene.Scene, pos coretypes.Position, destPos coretypes.Position) (flag bool, portalIndex int32, paths []astar.Pather) {
	flag, paths = generateNav(s, pos, destPos)
	if flag {
		return
	}
	portalIndex = -1
	allPortal := scenetemplate.GetSceneTemplateService().GetPortalTemplateMapByMapId(s.MapId())
	for _, portal := range allPortal {
		sourceMapId := s.MapId()
		var portalSceneTemplate *gametemplate.SceneTemplate

		//不同地图
		if portal.MapId != sourceMapId {
			continue
		}

		portalSceneTemplate = scenetemplate.GetSceneTemplateService().GetPortalSceneTemplate(int32(portal.TemplateId()))
		if portalSceneTemplate == nil {
			continue
		}

		portalPos := portalSceneTemplate.GetPos()
		//玩家不在传送阵场景,而且不在目的地地图
		if sourceMapId != portalSceneTemplate.SceneID {
			continue
		}
		toPortal, toPortalPath := generateNav(s, pos, portalPos)
		if !toPortal {
			continue
		}
		fromPortal, fromPortalPath := generateNav(s, portal.GetPosition(), destPos)
		if !fromPortal {
			continue
		}
		paths = append(paths, fromPortalPath...)
		paths = append(paths, toPortalPath...)
		portalIndex = int32(len(fromPortalPath))
		flag = true
		return
	}
	return
}

func generateNav(s scene.Scene, pos coretypes.Position, destPos coretypes.Position) (flag bool, paths []astar.Pather) {
	from := nav.TileFromWorld(s.GetWorld(), pos.X, pos.Z)
	to := nav.TileFromWorld(s.GetWorld(), destPos.X, destPos.Z)
	//TODO to可能为空
	if to == nil {
		return
	}

	x, z := to.GetXZ()
	if !to.IsMask() {
		//附近相邻的位置
		leftTile := s.GetWorld().Tile(z, x-1)
		if leftTile != nil && leftTile.IsMask() {
			to = leftTile
			goto AfterFind
		}
		leftUpTile := s.GetWorld().Tile(z+1, x-1)
		if leftUpTile != nil && leftUpTile.IsMask() {

			to = leftUpTile
			goto AfterFind
		}
		leftDownTile := s.GetWorld().Tile(z-1, x-1)
		if leftDownTile != nil && leftDownTile.IsMask() {
			to = leftDownTile
			goto AfterFind
		}
		rightTile := s.GetWorld().Tile(z, x+1)
		if rightTile != nil && rightTile.IsMask() {

			to = rightTile
			goto AfterFind
		}
		rightUpTile := s.GetWorld().Tile(z+1, x+1)
		if rightUpTile != nil && rightUpTile.IsMask() {

			to = rightUpTile
			goto AfterFind
		}
		rightDownTile := s.GetWorld().Tile(z-1, x+1)
		if rightDownTile != nil && rightDownTile.IsMask() {

			to = rightDownTile
			goto AfterFind
		}
		upTile := s.GetWorld().Tile(z+1, x)
		if upTile != nil && upTile.IsMask() {

			to = upTile
			goto AfterFind
		}
		downTile := s.GetWorld().Tile(z-1, x)
		if downTile != nil && downTile.IsMask() {

			to = downTile
			goto AfterFind
		}
	}
AfterFind:
	//不能移动
	if to == nil {
		return
	}

	if !to.IsMask() {
		return
	}

	if from == to {
		return true, nil
	} else {
		//获取寻路
		paths, _, found := astar.Path(from, to)
		if !found {
			return false, nil
		}
		if len(paths) <= 1 {
			return false, nil
		}
		return true, paths
	}
}
