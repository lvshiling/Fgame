package player

import (
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"

	"fgame/fgame/game/lingtong/dao"
	"fgame/fgame/game/lingtong/entity"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fmt"
)

//灵童信息
type PlayerLingTongObject struct {
	player     player.Player
	id         int64
	playerId   int64
	lingTongId int32
	level      int32
	basePower  int64
	power      int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerLingTongObject(pl player.Player) *PlayerLingTongObject {
	pwo := &PlayerLingTongObject{
		player: pl,
	}
	return pwo
}

func convertLingTongObjectToEntity(pwo *PlayerLingTongObject) (*entity.PlayerLingTongEntity, error) {
	e := &entity.PlayerLingTongEntity{
		Id:         pwo.id,
		PlayerId:   pwo.playerId,
		LingTongId: pwo.lingTongId,
		Level:      pwo.level,
		Power:      pwo.power,
		BasePower:  pwo.basePower,
		UpdateTime: pwo.updateTime,
		CreateTime: pwo.createTime,
		DeleteTime: pwo.deleteTime,
	}
	return e, nil
}

func (pwo *PlayerLingTongObject) GetPlayerId() int64 {
	return pwo.playerId
}

func (pwo *PlayerLingTongObject) GetDBId() int64 {
	return pwo.id
}

func (pwo *PlayerLingTongObject) GetLingTongId() int32 {
	return pwo.lingTongId
}

func (pwo *PlayerLingTongObject) IsActivateSys() bool {
	return pwo.lingTongId != 0
}

func (pwo *PlayerLingTongObject) GetPower() int64 {
	return pwo.power
}

func (pwo *PlayerLingTongObject) GetBasePower() int64 {
	return pwo.basePower
}

func (pwo *PlayerLingTongObject) GetLevel() int32 {
	return pwo.level
}

func (pwo *PlayerLingTongObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertLingTongObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerLingTongObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerLingTongEntity)

	pwo.id = pse.Id
	pwo.playerId = pse.PlayerId
	pwo.lingTongId = pse.LingTongId
	pwo.level = pse.Level
	pwo.power = pse.Power
	pwo.basePower = pse.BasePower
	pwo.updateTime = pse.UpdateTime
	pwo.createTime = pse.CreateTime
	pwo.deleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerLingTongObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("lingtong: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}

func (m *PlayerLingTongDataManager) loadLingTong() (err error) {
	//加载玩家灵童
	lingTongEntity, err := dao.GetLingTongDao().GetLingTongEntity(m.p.GetId())
	if err != nil {
		return
	}

	if lingTongEntity == nil {
		m.initPlayerLingTongObject()
	} else {
		m.lingTongObj = NewPlayerLingTongObject(m.p)
		m.lingTongObj.FromEntity(lingTongEntity)
	}
	return
}

func (m *PlayerLingTongDataManager) initPlayerLingTongObject() {
	obj := NewPlayerLingTongObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	//生成id
	obj.playerId = m.p.GetId()
	obj.lingTongId = 0
	obj.power = 0
	obj.basePower = 0
	obj.level = 0
	obj.createTime = now
	m.lingTongObj = obj
	obj.SetModified()
}

func (m *PlayerLingTongDataManager) GetLingTong() *PlayerLingTongObject {
	return m.lingTongObj
}

func (m *PlayerLingTongDataManager) LingTongChuZhan(lingTongId int32) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return
	}
	if m.lingTongObj.GetLingTongId() == lingTongId {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.lingTongObj.lingTongId = lingTongId
	m.lingTongObj.updateTime = now
	m.lingTongObj.SetModified()
	gameevent.Emit(lingtongeventtypes.EventTypeLingTongChuZhanChanged, m.p, lingTongId)
}

func (m *PlayerLingTongDataManager) AddLevel() {
	now := global.GetGame().GetTimeService().Now()
	m.lingTongObj.level++
	m.lingTongObj.updateTime = now
	m.lingTongObj.SetModified()
	gameevent.Emit(lingtongeventtypes.EventTypeLingTongLevelChanged, m.p, m.lingTongObj)
}

func (m *PlayerLingTongDataManager) Power(power int64) {
	if power < 0 {
		return
	}
	if power == m.lingTongObj.power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.lingTongObj.power = power
	m.lingTongObj.updateTime = now
	m.lingTongObj.SetModified()
}

func (m *PlayerLingTongDataManager) BasePower(basePower int64) bool {
	if basePower < 0 {
		return false
	}
	//减少数据库操作
	if basePower == m.lingTongObj.basePower {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	m.lingTongObj.basePower = basePower
	m.lingTongObj.updateTime = now
	m.lingTongObj.SetModified()
	return true
}

func (m *PlayerLingTongDataManager) GetBasePower() int64 {
	return m.lingTongObj.basePower
}
