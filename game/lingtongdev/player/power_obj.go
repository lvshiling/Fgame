package player

import (
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lingtongdev/dao"
	"fgame/fgame/game/lingtongdev/entity"
	lingtongdevevent "fgame/fgame/game/lingtongdev/event/types"
	"fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fmt"
)

//灵童养成战力对象
type PlayerLingTongPowerObject struct {
	player     player.Player
	id         int64
	playerId   int64
	classType  types.LingTongDevSysType
	power      int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerLingTongPowerObject(pl player.Player) *PlayerLingTongPowerObject {
	pwo := &PlayerLingTongPowerObject{
		player: pl,
	}
	return pwo
}

func convertLingTongPowerObjectToEntity(pwo *PlayerLingTongPowerObject) (*entity.PlayerLingTongPowerEntity, error) {
	e := &entity.PlayerLingTongPowerEntity{
		Id:         pwo.id,
		PlayerId:   pwo.playerId,
		ClassType:  int32(pwo.classType),
		Power:      pwo.power,
		UpdateTime: pwo.updateTime,
		CreateTime: pwo.createTime,
		DeleteTime: pwo.deleteTime,
	}
	return e, nil
}

func (pwo *PlayerLingTongPowerObject) GetPlayerId() int64 {
	return pwo.playerId
}

func (pwo *PlayerLingTongPowerObject) GetDBId() int64 {
	return pwo.id
}

func (pwo *PlayerLingTongPowerObject) GetClassType() types.LingTongDevSysType {
	return pwo.classType
}

func (pwo *PlayerLingTongPowerObject) GetPower() int64 {
	return pwo.power
}

func (pwo *PlayerLingTongPowerObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertLingTongPowerObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerLingTongPowerObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerLingTongPowerEntity)

	pwo.id = pse.Id
	pwo.playerId = pse.PlayerId
	pwo.classType = types.LingTongDevSysType(pse.ClassType)
	pwo.power = pse.Power
	pwo.updateTime = pse.UpdateTime
	pwo.createTime = pse.CreateTime
	pwo.deleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerLingTongPowerObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("lingtongpower: err[%s]", err.Error()))
	}
	obj, ok := e.(playertypes.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}

func (m *PlayerLingTongDevDataManager) loadPower() (err error) {
	m.playerPowerMap = make(map[types.LingTongDevSysType]*PlayerLingTongPowerObject)
	powerList, err := dao.GetLingTongDevDao().GetLingTongPowerList(m.p.GetId())
	if err != nil {
		return
	}
	for _, obj := range powerList {
		pwo := NewPlayerLingTongPowerObject(m.p)
		pwo.FromEntity(obj)
		m.playerPowerMap[pwo.GetClassType()] = pwo
	}
	return
}

func (m *PlayerLingTongDevDataManager) initPlayerLingTongPowerObj(classType types.LingTongDevSysType) (obj *PlayerLingTongPowerObject) {
	obj = m.getLingTongPowerInfo(classType)
	if obj != nil {
		return
	}
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj = NewPlayerLingTongPowerObject(m.p)
	obj.playerId = m.p.GetId()
	obj.id = id
	obj.classType = classType
	obj.power = 0
	obj.createTime = now
	obj.SetModified()
	m.playerPowerMap[classType] = obj
	return obj
}

func (m *PlayerLingTongDevDataManager) getLingTongPowerInfo(classType types.LingTongDevSysType) *PlayerLingTongPowerObject {
	obj, ok := m.playerPowerMap[classType]
	if !ok {
		return nil
	}
	return obj
}

//灵童养成战斗力
func (m *PlayerLingTongDevDataManager) LingTongDevPower(classType types.LingTongDevSysType, power int64) {
	if power <= 0 {
		return
	}
	lingTongObj := m.getLingTongInfo(classType)
	if lingTongObj == nil {
		return
	}
	obj := m.getLingTongPowerInfo(classType)
	if obj == nil {
		obj = m.initPlayerLingTongPowerObj(classType)
	}

	if obj.power == power {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj.power = power
	obj.updateTime = now
	obj.SetModified()
	gameevent.Emit(lingtongdevevent.EventTypeLingTongDevPowerChanged, m.p, obj)
	return
}
