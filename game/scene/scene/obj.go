package scene

import (
	"fgame/fgame/core/aoi"
	coretypes "fgame/fgame/core/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//场景对象
type SceneObject interface {
	aoi.AOI
	//获取角度
	GetAngle() float64
	//设置角度
	SetAngle(angle float64)
	//获取场景对象类型
	GetSceneObjectType() scenetypes.BiologyType
	//获取aoi邻居
	GetNeighbors() map[int64]aoi.AOI
	//获取邻居
	IsNeighbor(id int64) bool
	//进入场景
	EnterScene(s Scene)
	//获取场景
	GetScene() Scene
	//退出场景
	ExitScene(active bool)

	//移动
	Move(pos coretypes.Position, angle float64)
	String() string
}

//场景单位
type SceneObjectBase struct {
	so              SceneObject
	pos             coretypes.Position
	angle           float64
	sceneObjectType scenetypes.BiologyType
	//周围物体
	neighbors map[int64]aoi.AOI

	//场景
	s Scene
}

func (sub *SceneObjectBase) String() string {
	return fmt.Sprintf("id:%d,position:%s,angle:%.2f,type:%s", sub.so.GetId(), sub.pos.String(), sub.angle, sub.sceneObjectType.String())
}

func (sub *SceneObjectBase) EnterScene(s Scene) {
	sub.ResetNeighbor()
	sub.s = s
}

func (sub *SceneObjectBase) GetScene() Scene {
	return sub.s
}

func (sub *SceneObjectBase) ExitScene(active bool) {
	sub.s = nil
}

func (sub *SceneObjectBase) GetNeighbors() map[int64]aoi.AOI {
	return sub.neighbors
}

func (sub *SceneObjectBase) IsNeighbor(id int64) bool {
	_, ok := sub.neighbors[id]
	if !ok {
		return false
	}
	return true
}

func (sub *SceneObjectBase) GetPosition() coretypes.Position {
	return sub.pos
}

func (sub *SceneObjectBase) SetPosition(pos coretypes.Position) {
	sub.pos = pos
}

func (sub *SceneObjectBase) GetAngle() float64 {
	return sub.angle
}

func (sub *SceneObjectBase) SetAngle(angle float64) {
	sub.angle = angle
}

func (sub *SceneObjectBase) GetSceneObjectType() scenetypes.BiologyType {
	return sub.sceneObjectType
}

func (sub *SceneObjectBase) OnEnterAOI(other aoi.AOI) {
	// _, ok := sub.neighbors[other.GetId()]
	// if ok {
	// 	panic(fmt.Errorf("scene:同样的对象[%d]", other.GetId()))
	// }
	sub.neighbors[other.GetId()] = other
}

func (sub *SceneObjectBase) OnLeaveAOI(other aoi.AOI, complete bool) {
	// _, ok := sub.neighbors[other.GetId()]
	// if !ok {
	// 	panic(fmt.Errorf("scene:不存在[%d]", other.GetId()))
	// }

	delete(sub.neighbors, other.GetId())
}

func (sub *SceneObjectBase) ResetNeighbor() {
	for id, _ := range sub.neighbors {
		delete(sub.neighbors, id)
	}
}

func (sub *SceneObjectBase) Move(pos coretypes.Position, angle float64) {
	// sub.SetPosition(pos)
	sub.SetAngle(angle)

}

func NewSceneObjectBase(so SceneObject, pos coretypes.Position, angle float64, sceneObjectType scenetypes.BiologyType) *SceneObjectBase {
	sob := &SceneObjectBase{
		so:              so,
		pos:             pos,
		angle:           angle,
		sceneObjectType: sceneObjectType,
		neighbors:       make(map[int64]aoi.AOI),
	}

	return sob
}
