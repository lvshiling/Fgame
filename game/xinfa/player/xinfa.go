package player

import (
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/xinfa/dao"
	xinfaentity "fgame/fgame/game/xinfa/entity"
	xinfaeventtypes "fgame/fgame/game/xinfa/event/types"
	xinfatypes "fgame/fgame/game/xinfa/types"
	"fgame/fgame/game/xinfa/xinfa"
	"fgame/fgame/pkg/idutil"
	"fmt"
)

//心法对象
type PlayerXinFaObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Type       xinfatypes.XinFaType
	Level      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerXinFaObject(pl player.Player) *PlayerXinFaObject {
	pso := &PlayerXinFaObject{
		player: pl,
	}
	return pso
}

func (pxfo *PlayerXinFaObject) GetPlayerId() int64 {
	return pxfo.PlayerId
}

func (pxfo *PlayerXinFaObject) GetDBId() int64 {
	return pxfo.Id
}

func (pxfo *PlayerXinFaObject) ToEntity() (e storage.Entity, err error) {
	e = &xinfaentity.PlayerXinFaEntity{
		Id:         pxfo.Id,
		PlayerId:   pxfo.PlayerId,
		Type:       int32(pxfo.Type),
		Level:      pxfo.Level,
		UpdateTime: pxfo.UpdateTime,
		CreateTime: pxfo.CreateTime,
		DeleteTime: pxfo.DeleteTime,
	}
	return e, err
}

func (pxfo *PlayerXinFaObject) FromEntity(e storage.Entity) error {
	pxfe, _ := e.(*xinfaentity.PlayerXinFaEntity)
	pxfo.Id = pxfe.Id
	pxfo.PlayerId = pxfe.PlayerId
	pxfo.Type = xinfatypes.XinFaType(pxfe.Type)
	pxfo.Level = pxfe.Level
	pxfo.UpdateTime = pxfe.UpdateTime
	pxfo.CreateTime = pxfe.CreateTime
	pxfo.DeleteTime = pxfe.DeleteTime
	return nil
}

func (pxfo *PlayerXinFaObject) SetModified() {
	e, err := pxfo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("xinfa: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pxfo.player.AddChangedObject(obj)
	return
}

//玩家心法管理器
type PlayerXinFaDataManager struct {
	p player.Player
	//玩家心法map
	playerXinFaObjectMap map[xinfatypes.XinFaType]*PlayerXinFaObject
}

func (pxfdm *PlayerXinFaDataManager) Player() player.Player {
	return pxfdm.p
}

func (pxfdm *PlayerXinFaDataManager) GetAllXinfaLevel() int32 {
	totalLevel := int32(0)
	for _, obj := range pxfdm.playerXinFaObjectMap {
		totalLevel += obj.Level
	}
	return totalLevel
}

//加载
func (pxfdm *PlayerXinFaDataManager) Load() (err error) {
	pxfdm.playerXinFaObjectMap = make(map[xinfatypes.XinFaType]*PlayerXinFaObject)
	//加载玩家心法
	xinFas, err := dao.GetXinFaDao().GetXinFaList(pxfdm.p.GetId())
	if err != nil {
		return
	}
	//心法信息
	for _, xinFa := range xinFas {
		pxfo := NewPlayerXinFaObject(pxfdm.p)
		pxfo.FromEntity(xinFa)
		pxfdm.playerXinFaObjectMap[pxfo.Type] = pxfo
	}

	return nil
}

//加载后
func (pxfdm *PlayerXinFaDataManager) AfterLoad() (err error) {
	pxfdm.resetMaxLevel()
	return nil
}

func (pxfdm *PlayerXinFaDataManager) resetMaxLevel() (err error) {
	now := global.GetGame().GetTimeService().Now()
	for typ, obj := range pxfdm.playerXinFaObjectMap {
		temp := xinfa.GetXinFaService().GetXinFaMaxLevel(typ)
		if obj.Level > temp.Level {
			obj.Level = temp.Level
			obj.UpdateTime = now
			obj.SetModified()
		}
	}
	return
}

//心跳
func (pxfdm *PlayerXinFaDataManager) Heartbeat() {

}

//获取玩家心法map
func (pxfdm *PlayerXinFaDataManager) GetXinFaMap() map[xinfatypes.XinFaType]*PlayerXinFaObject {
	return pxfdm.playerXinFaObjectMap
}

//获取心法通过心法类型
func (pxfdm *PlayerXinFaDataManager) GetXinFaByTyp(typ xinfatypes.XinFaType) *PlayerXinFaObject {
	xinFaObj, exist := pxfdm.playerXinFaObjectMap[typ]
	if !exist {
		return nil
	}
	return xinFaObj
}

//获取等级通过心法类型
func (pxfdm *PlayerXinFaDataManager) GetXinFaLevelByTyp(typ xinfatypes.XinFaType) int32 {
	obj := pxfdm.GetXinFaByTyp(typ)
	if obj == nil {
		return 0
	}
	return obj.Level
}

//是否已激活
func (pxfdm *PlayerXinFaDataManager) IfXinFaExist(typ xinfatypes.XinFaType) bool {
	obj := pxfdm.GetXinFaByTyp(typ)
	if obj == nil {
		return false
	}
	return true
}

//是否达到满级
func (pxfdm *PlayerXinFaDataManager) ifFullLevel(typ xinfatypes.XinFaType) bool {
	level := pxfdm.GetXinFaLevelByTyp(typ)
	to := xinfa.GetXinFaService().GetXinFaByTypeAndLevel(xinfatypes.XinFaType(typ), level)
	if to.NextId == 0 {
		return true
	}
	return false
}

//能否升级
func (pxfdm *PlayerXinFaDataManager) IfCanUpgrade(typ xinfatypes.XinFaType) bool {
	flag := pxfdm.IfXinFaExist(typ)
	if !flag {
		return false
	}
	flag = pxfdm.ifFullLevel(typ)
	if flag {
		return false
	}
	return true
}

//心法激活
func (pxfdm *PlayerXinFaDataManager) XinFaActive(typ xinfatypes.XinFaType) bool {
	flag := typ.Valid()
	if !flag {
		return false
	}
	flag = pxfdm.IfXinFaExist(typ)
	if flag {
		return false
	}
	id, err := idutil.GetId()
	if err != nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	pxfo := NewPlayerXinFaObject(pxfdm.p)
	pxfo.Id = id
	pxfo.PlayerId = pxfdm.p.GetId()
	pxfo.Type = typ
	pxfo.Level = int32(1)
	pxfo.CreateTime = now
	pxfo.SetModified()

	pxfdm.playerXinFaObjectMap[typ] = pxfo
	gameevent.Emit(xinfaeventtypes.EventTypeXinFaActive, pxfdm.p, typ)
	return true
}

//心法升级
func (pxfdm *PlayerXinFaDataManager) Upgrade(typ xinfatypes.XinFaType) bool {
	flag := pxfdm.IfCanUpgrade(typ)
	if !flag {
		return false
	}

	obj := pxfdm.GetXinFaByTyp(typ)
	now := global.GetGame().GetTimeService().Now()
	obj.UpdateTime = now
	obj.Level += 1
	obj.SetModified()

	gameevent.Emit(xinfaeventtypes.EventTypeXinFaUpgrade, pxfdm.p, typ)
	return true
}

func CreatePlayerXinFaDataManager(p player.Player) player.PlayerDataManager {
	pxfdm := &PlayerXinFaDataManager{}
	pxfdm.p = p
	return pxfdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerXinFaDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerXinFaDataManager))
}
