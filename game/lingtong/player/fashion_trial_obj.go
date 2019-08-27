package player

import (
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/lingtong/dao"
	"fgame/fgame/game/lingtong/entity"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	"fgame/fgame/game/lingtong/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fmt"
)

//灵童时装试用
type PlayerLingTongFashionTrialObject struct {
	player         player.Player
	id             int64
	playerId       int64
	trialFashionId int32
	activateTime   int64
	durationTime   int64
	isExpire       int32
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func NewPlayerLingTongFashionTrialObject(pl player.Player) *PlayerLingTongFashionTrialObject {
	pwo := &PlayerLingTongFashionTrialObject{
		player: pl,
	}
	return pwo
}

func convertLingTongFashionTrialObjectToEntity(pwo *PlayerLingTongFashionTrialObject) (*entity.PlayerLingTongFashionTrialEntity, error) {
	e := &entity.PlayerLingTongFashionTrialEntity{
		Id:             pwo.id,
		PlayerId:       pwo.playerId,
		TrialFashionId: pwo.trialFashionId,
		ActivateTime:   pwo.activateTime,
		DurationTime:   pwo.durationTime,
		IsExpire:       pwo.isExpire,
		UpdateTime:     pwo.updateTime,
		CreateTime:     pwo.createTime,
		DeleteTime:     pwo.deleteTime,
	}
	return e, nil
}

func (pwo *PlayerLingTongFashionTrialObject) GetPlayerId() int64 {
	return pwo.playerId
}

func (pwo *PlayerLingTongFashionTrialObject) GetDBId() int64 {
	return pwo.id
}

func (pwo *PlayerLingTongFashionTrialObject) GetTrialFashionId() int32 {
	return pwo.trialFashionId
}

func (pwo *PlayerLingTongFashionTrialObject) GetActivateTime() int64 {
	return pwo.activateTime
}

func (pwo *PlayerLingTongFashionTrialObject) GetDurationTime() int64 {
	return pwo.durationTime
}

func (pwo *PlayerLingTongFashionTrialObject) GetIsExpire() bool {
	return pwo.isExpire == 1
}

func (pwo *PlayerLingTongFashionTrialObject) refresh() {
	now := global.GetGame().GetTimeService().Now()
	if pwo.GetIsExpire() {
		return
	}
	diffTime := now - pwo.activateTime
	if diffTime >= pwo.durationTime {
		pwo.isExpire = 1
		pwo.activateTime = 0
		pwo.durationTime = 0
		pwo.SetModified()
	}
}

func (pwo *PlayerLingTongFashionTrialObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertLingTongFashionTrialObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerLingTongFashionTrialObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerLingTongFashionTrialEntity)

	pwo.id = pse.Id
	pwo.playerId = pse.PlayerId
	pwo.trialFashionId = pse.TrialFashionId
	pwo.activateTime = pse.ActivateTime
	pwo.durationTime = pse.DurationTime
	pwo.isExpire = pse.IsExpire
	pwo.updateTime = pse.UpdateTime
	pwo.createTime = pse.CreateTime
	pwo.deleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerLingTongFashionTrialObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("lingtongfashiontrial: err[%s]", err.Error()))
	}
	obj, ok := e.(playertypes.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}

func (m *PlayerLingTongDataManager) loadLingTongFashionTrial() (err error) {
	//加载玩家灵童时装试用
	lingTongFashionTrialEntity, err := dao.GetLingTongDao().GetLingTongFashionTrialEntity(m.p.GetId())
	if err != nil {
		return
	}

	if lingTongFashionTrialEntity == nil {
		m.initPlayerLingTongFashionTrialObject()
	} else {
		m.trialFashionObj = NewPlayerLingTongFashionTrialObject(m.p)
		m.trialFashionObj.FromEntity(lingTongFashionTrialEntity)
		m.trialFashionObj.refresh()
	}
	return
}

func (m *PlayerLingTongDataManager) initPlayerLingTongFashionTrialObject() {
	obj := NewPlayerLingTongFashionTrialObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	//生成id
	obj.playerId = m.p.GetId()
	obj.trialFashionId = 0
	obj.activateTime = 0
	obj.durationTime = 0
	obj.isExpire = 1
	obj.createTime = now
	m.trialFashionObj = obj
	obj.SetModified()
}

func (m *PlayerLingTongDataManager) getFashionTrailObject() *PlayerLingTongFashionTrialObject {
	return m.trialFashionObj
}

func (m *PlayerLingTongDataManager) IsFashionTrial(fashionId int32) (flag bool) {
	if !m.trialFashionObj.GetIsExpire() {
		flag = true
		return
	}
	return
}

func (m *PlayerLingTongDataManager) GetFashionTrialObject() *PlayerLingTongFashionTrialObject {
	return m.getFashionTrailObject()
}

// 时装试用
func (m *PlayerLingTongDataManager) UseFashionTrialCard(trialCardItemId int32) (expireTime int64) {
	itemTemplate := item.GetItemService().GetItem(int(trialCardItemId))
	if itemTemplate == nil {
		return
	}
	trialFashionId := itemTemplate.TypeFlag1
	durationTime := int64(itemTemplate.TypeFlag2)
	now := global.GetGame().GetTimeService().Now()
	m.trialFashionObj.trialFashionId = trialFashionId
	m.trialFashionObj.durationTime = durationTime
	m.trialFashionObj.activateTime = now
	m.trialFashionObj.isExpire = 0
	m.trialFashionObj.SetModified()
	expireTime = now + durationTime
	m.trialFashionUse(trialFashionId)
	return
}

// 时装试用过期
func (m *PlayerLingTongDataManager) TrialFashionOverdue(trialId int32, overdueType types.LingTongFashionTrialOverdueType) {
	trialObj := m.getFashionTrailObject()
	if trialObj == nil {
		return
	}
	if trialObj.GetIsExpire() {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	trialObj.activateTime = 0
	trialObj.durationTime = 0
	trialObj.isExpire = 1
	trialObj.updateTime = now
	trialObj.SetModified()

	m.refreshLingTongFashion()

	eventData := lingtongeventtypes.CreateLingTongFashionTrialOverdueEventData(trialId, overdueType)
	gameevent.Emit(lingtongeventtypes.EventTypeLingTongFashionTrialOverdue, m.p, eventData)
}
