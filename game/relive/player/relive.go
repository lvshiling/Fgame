package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/relive/dao"
	reliveentity "fgame/fgame/game/relive/entity"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

//玩家复活
type PlayerReliveObject struct {
	player         player.Player
	id             int64
	culTime        int32
	lastReliveTime int64
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func NewPlayerReliveObject(pl player.Player) *PlayerReliveObject {
	o := &PlayerReliveObject{
		player: pl,
	}
	return o
}

func (o *PlayerReliveObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerReliveObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerReliveObject) ToEntity() (e storage.Entity, err error) {
	e = &reliveentity.PlayerReliveEntity{
		Id:             o.id,
		PlayerId:       o.GetPlayerId(),
		CulTime:        o.culTime,
		LastReliveTime: o.lastReliveTime,
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return e, nil
}

func (o *PlayerReliveObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*reliveentity.PlayerReliveEntity)

	o.id = te.Id
	o.culTime = te.CulTime
	o.lastReliveTime = te.LastReliveTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerReliveObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "relive"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//玩家复活管理器
type PlayerReliveDataManager struct {
	p player.Player
	//玩家天劫塔对象
	playerReliveObject *PlayerReliveObject
	// runner             heartbeat.HeartbeatTaskRunner
}

func (m *PlayerReliveDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerReliveDataManager) Load() (err error) {
	//加载玩家天劫塔信息
	reliveEntity, err := dao.GetReliveDao().GetPlayerReliveEntity(m.p.GetId())
	if err != nil {
		return
	}
	if reliveEntity == nil {
		m.initPlayerReliveObject()
	} else {
		m.playerReliveObject = NewPlayerReliveObject(m.p)
		m.playerReliveObject.FromEntity(reliveEntity)
	}

	return nil
}

//第一次初始化
func (m *PlayerReliveDataManager) initPlayerReliveObject() {
	o := NewPlayerReliveObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.culTime = 0
	o.lastReliveTime = now
	o.createTime = now
	o.SetModified()
	m.playerReliveObject = o
}

//加载后
func (m *PlayerReliveDataManager) AfterLoad() (err error) {
	now := global.GetGame().GetTimeService().Now()
	elapse := now - m.playerReliveObject.lastReliveTime
	clearTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeReliveTimesClearTime)
	//复活清空数据
	if elapse >= int64(clearTime) {
		m.playerReliveObject.culTime = 0
		m.playerReliveObject.updateTime = now
		m.playerReliveObject.lastReliveTime = now
		m.playerReliveObject.SetModified()
		return
	}
	return nil
}

//获取累计复活次数
func (m *PlayerReliveDataManager) GetCulTime() int32 {
	return m.playerReliveObject.culTime
}

func (m *PlayerReliveDataManager) GetLastReliveTime() int64 {
	return m.playerReliveObject.lastReliveTime
}

func (m *PlayerReliveDataManager) Save() {
	now := global.GetGame().GetTimeService().Now()
	m.playerReliveObject.culTime = m.p.GetCulReliveTime()
	m.playerReliveObject.lastReliveTime = m.p.GetLastReliveTime()
	m.playerReliveObject.updateTime = now
	m.playerReliveObject.SetModified()
}

// func (m *PlayerReliveDataManager) Refresh() bool {
// 	now := global.GetGame().GetTimeService().Now()
// 	elapse := now - m.playerReliveObject.lastReliveTime
// 	clearTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeReliveTimesClearTime)
// 	//复活清空数据
// 	if elapse >= int64(clearTime) {
// 		m.playerReliveObject.culTime = 0
// 		m.playerReliveObject.updateTime = now
// 		m.playerReliveObject.lastReliveTime = now
// 		m.playerReliveObject.SetModified()
// 		return true
// 	}
// 	return false
// }

// //复活
// func (m *PlayerReliveDataManager) Relive() {
// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerReliveObject.culTime += 1
// 	m.playerReliveObject.updateTime = now
// 	m.playerReliveObject.lastReliveTime = now
// 	m.playerReliveObject.SetModified()
// 	//玩家复活
// 	// gameevent.Emit(reliveeventtypes.EventTypePlayerRelive, m.p, nil)
// }

//心跳
func (m *PlayerReliveDataManager) Heartbeat() {
	// m.runner.Heartbeat()
}

func CreatePlayerReliveDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerReliveDataManager{}
	m.p = p
	// m.runner = heartbeat.NewHeartbeatTaskRunner()
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerReliveDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerReliveDataManager))
}
