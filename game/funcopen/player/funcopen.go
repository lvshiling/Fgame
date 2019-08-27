package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/funcopen/dao"
	funcopenentity "fgame/fgame/game/funcopen/entity"
	funcopeneventtypes "fgame/fgame/game/funcopen/event/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	funcopentypes "fgame/fgame/game/funcopen/types"

	"github.com/pkg/errors"
)

//功能开启对象
type PlayerFuncOpenObject struct {
	player       player.Player
	Id           int64
	FuncOpenList []funcopentypes.FuncOpenType
	UpdateTime   int64
	CreateTime   int64
	DeleteTime   int64
}

func NewPlayerFuncOpenObject(pl player.Player) *PlayerFuncOpenObject {
	pdo := &PlayerFuncOpenObject{
		player: pl,
	}
	return pdo
}

func convertNewPlayerFuncOpenObjectToEntity(pdo *PlayerFuncOpenObject) (*funcopenentity.PlayerFuncOpenEntity, error) {
	funcOpenList, err := json.Marshal(pdo.FuncOpenList)
	if err != nil {
		return nil, err
	}

	e := &funcopenentity.PlayerFuncOpenEntity{
		Id:           pdo.Id,
		PlayerId:     pdo.player.GetId(),
		FuncOpenList: string(funcOpenList),
		UpdateTime:   pdo.UpdateTime,
		CreateTime:   pdo.CreateTime,
		DeleteTime:   pdo.DeleteTime,
	}
	return e, err
}

func (pdo *PlayerFuncOpenObject) GetPlayerId() int64 {
	return pdo.player.GetId()
}

func (pdo *PlayerFuncOpenObject) GetDBId() int64 {
	return pdo.Id
}

func (pdo *PlayerFuncOpenObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerFuncOpenObjectToEntity(pdo)
	return e, err
}

func (pdo *PlayerFuncOpenObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*funcopenentity.PlayerFuncOpenEntity)
	var funcOpenList []funcopentypes.FuncOpenType
	if err := json.Unmarshal([]byte(pse.FuncOpenList), &funcOpenList); err != nil {
		return err
	}
	pdo.Id = pse.Id
	pdo.FuncOpenList = funcOpenList
	pdo.UpdateTime = pse.UpdateTime
	pdo.CreateTime = pse.CreateTime
	pdo.DeleteTime = pse.DeleteTime
	return nil
}

func (pdo *PlayerFuncOpenObject) SetModified() {
	e, err := pdo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FuncOpen"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pdo.player.AddChangedObject(obj)
	return
}

//玩家功能开启管理器
type PlayerFuncOpenDataManager struct {
	p player.Player
	//玩家功能开启对象
	playerFuncOpenObject *PlayerFuncOpenObject
}

func (m *PlayerFuncOpenDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerFuncOpenDataManager) Load() (err error) {
	//加载玩家功能开启列表
	funcOpenEntity, err := dao.GetFuncOpenDao().GetFuncOpenEntity(m.p.GetId())
	if err != nil {
		return
	}
	if funcOpenEntity == nil {
		m.initPlayerFuncOpenObject()
	} else {
		m.playerFuncOpenObject = NewPlayerFuncOpenObject(m.p)
		err = m.playerFuncOpenObject.FromEntity(funcOpenEntity)
		if err != nil {
			return
		}
	}

	return
}

//第一次初始化
func (m *PlayerFuncOpenDataManager) initPlayerFuncOpenObject() {
	pdo := NewPlayerFuncOpenObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pdo.Id = id

	pdo.FuncOpenList = make([]funcopentypes.FuncOpenType, 0, 8)
	pdo.CreateTime = now
	pdo.SetModified()

	m.playerFuncOpenObject = pdo
	return
}

//加载后
func (m *PlayerFuncOpenDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerFuncOpenDataManager) Heartbeat() {
}

//获取功能列表
func (m *PlayerFuncOpenDataManager) GetOpenFuncList() []funcopentypes.FuncOpenType {
	return m.playerFuncOpenObject.FuncOpenList
}

//获取功能
func (m *PlayerFuncOpenDataManager) IsOpen(funcOpenType funcopentypes.FuncOpenType) bool {
	for _, f := range m.playerFuncOpenObject.FuncOpenList {
		if f == funcOpenType {
			return true
		}
	}
	return false
}

//添加功能
func (m *PlayerFuncOpenDataManager) AddOpenFunc(funcOpenType funcopentypes.FuncOpenType) (flag bool) {
	if m.IsOpen(funcOpenType) {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerFuncOpenObject.FuncOpenList = append(m.playerFuncOpenObject.FuncOpenList, funcOpenType)
	m.playerFuncOpenObject.UpdateTime = now
	m.playerFuncOpenObject.SetModified()

	flag = true

	gameevent.Emit(funcopeneventtypes.EventTypeFuncOpen, m.p, funcOpenType)
	return
}

func CreatePlayerFuncOpenDataManager(p player.Player) player.PlayerDataManager {
	pddm := &PlayerFuncOpenDataManager{}
	pddm.p = p
	return pddm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerFuncOpenDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerFuncOpenDataManager))
}
