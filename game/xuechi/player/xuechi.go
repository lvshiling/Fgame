package player

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/xuechi/dao"
	xuechientity "fgame/fgame/game/xuechi/entity"
	"fgame/fgame/pkg/idutil"
	"fmt"
)

//血池对象
type PlayerXueChiObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	BloodLine  int32
	Blood      int64
	LastTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerXueChiObject(pl player.Player) *PlayerXueChiObject {
	pso := &PlayerXueChiObject{
		player: pl,
	}
	return pso
}

func (pxco *PlayerXueChiObject) GetPlayerId() int64 {
	return pxco.PlayerId
}

func (pxco *PlayerXueChiObject) GetDBId() int64 {
	return pxco.Id
}

func (pxco *PlayerXueChiObject) ToEntity() (e storage.Entity, err error) {
	e = &xuechientity.PlayerXueChiEntity{
		Id:         pxco.Id,
		PlayerId:   pxco.PlayerId,
		BloodLine:  pxco.BloodLine,
		Blood:      pxco.Blood,
		LastTime:   pxco.LastTime,
		UpdateTime: pxco.UpdateTime,
		CreateTime: pxco.CreateTime,
		DeleteTime: pxco.DeleteTime,
	}
	return e, err
}

func (pxco *PlayerXueChiObject) FromEntity(e storage.Entity) error {
	pxce, _ := e.(*xuechientity.PlayerXueChiEntity)
	pxco.Id = pxce.Id
	pxco.PlayerId = pxce.PlayerId
	pxco.BloodLine = pxce.BloodLine
	pxco.Blood = pxce.Blood
	pxco.LastTime = pxce.LastTime
	pxco.UpdateTime = pxce.UpdateTime
	pxco.CreateTime = pxce.CreateTime
	pxco.DeleteTime = pxce.DeleteTime
	return nil
}

func (pxco *PlayerXueChiObject) SetModified() {
	e, err := pxco.ToEntity()
	if err != nil {
		panic(fmt.Errorf("xuechi: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pxco.player.AddChangedObject(obj)
	return
}

//玩家血池管理器
type PlayerXueChiDataManager struct {
	p player.Player
	//玩家血池
	playerXueChiObject *PlayerXueChiObject
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

func (pxcdm *PlayerXueChiDataManager) Player() player.Player {
	return pxcdm.p
}

//加载
func (pxcdm *PlayerXueChiDataManager) Load() (err error) {
	//加载玩家血池
	xueChiEntity, err := dao.GetXueChiDao().GetXueChiEntity(pxcdm.p.GetId())
	if err != nil {
		return
	}
	if xueChiEntity == nil {
		pxcdm.initPlayerXueChiObject()
	} else {
		pxcdm.playerXueChiObject = NewPlayerXueChiObject(pxcdm.p)
		pxcdm.playerXueChiObject.FromEntity(xueChiEntity)
	}
	return nil
}

//第一次初始化
func (pxcdm *PlayerXueChiDataManager) initPlayerXueChiObject() {
	pxco := NewPlayerXueChiObject(pxcdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pxco.Id = id

	initBloodLine := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeXueChiInitBloodLine)
	initBlood := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeXueChiInitBlood))
	//生成id
	pxco.PlayerId = pxcdm.p.GetId()
	pxco.BloodLine = initBloodLine
	pxco.Blood = initBlood
	pxco.LastTime = 0
	pxco.CreateTime = now
	pxcdm.playerXueChiObject = pxco
	pxco.SetModified()
}

//加载后
func (pxcdm *PlayerXueChiDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (pxcdm *PlayerXueChiDataManager) Heartbeat() {

}

func (pxcdm *PlayerXueChiDataManager) Save() {
	now := global.GetGame().GetTimeService().Now()
	pxcdm.playerXueChiObject.BloodLine = pxcdm.p.GetBloodLine()
	pxcdm.playerXueChiObject.Blood = pxcdm.p.GetBlood()
	pxcdm.playerXueChiObject.UpdateTime = now
	pxcdm.playerXueChiObject.SetModified()
}

func (pxcdm *PlayerXueChiDataManager) GetXueChi() *PlayerXueChiObject {
	return pxcdm.playerXueChiObject
}

//仅Gm 使用
func (pxcdm *PlayerXueChiDataManager) GmSetBlood(blood int64) {
	now := global.GetGame().GetTimeService().Now()
	pxcdm.playerXueChiObject.Blood = blood
	pxcdm.playerXueChiObject.UpdateTime = now
	pxcdm.playerXueChiObject.SetModified()
}

func CreatePlayerXueChiDataManager(p player.Player) player.PlayerDataManager {
	pxcdm := &PlayerXueChiDataManager{}
	pxcdm.p = p
	pxcdm.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	return pxcdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerXueChiDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerXueChiDataManager))
}
