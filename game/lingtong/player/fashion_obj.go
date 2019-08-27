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
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fmt"
)

//灵童时装
type PlayerLingTongFashionObject struct {
	player     player.Player
	id         int64
	playerId   int64
	lingTongId int32
	fashionId  int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerLingTongFashionObject(pl player.Player) *PlayerLingTongFashionObject {
	pwo := &PlayerLingTongFashionObject{
		player: pl,
	}
	return pwo
}

func convertLingTongFashionObjectToEntity(pwo *PlayerLingTongFashionObject) (*entity.PlayerLingTongFashionEntity, error) {
	e := &entity.PlayerLingTongFashionEntity{
		Id:         pwo.id,
		PlayerId:   pwo.playerId,
		LingTongId: pwo.lingTongId,
		FashionId:  pwo.fashionId,
		UpdateTime: pwo.updateTime,
		CreateTime: pwo.createTime,
		DeleteTime: pwo.deleteTime,
	}
	return e, nil
}

func (pwo *PlayerLingTongFashionObject) GetPlayerId() int64 {
	return pwo.playerId
}

func (pwo *PlayerLingTongFashionObject) GetDBId() int64 {
	return pwo.id
}

func (pwo *PlayerLingTongFashionObject) GetLingTongId() int32 {
	return pwo.lingTongId
}

func (pwo *PlayerLingTongFashionObject) GetFashionId() int32 {
	return pwo.fashionId
}

func (pwo *PlayerLingTongFashionObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertLingTongFashionObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerLingTongFashionObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerLingTongFashionEntity)

	pwo.id = pse.Id
	pwo.playerId = pse.PlayerId
	pwo.lingTongId = pse.LingTongId
	pwo.fashionId = pse.FashionId
	pwo.updateTime = pse.UpdateTime
	pwo.createTime = pse.CreateTime
	pwo.deleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerLingTongFashionObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("lingtongfashion: err[%s]", err.Error()))
	}
	obj, ok := e.(playertypes.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}

func (m *PlayerLingTongDataManager) loadLingTongFashion() (err error) {
	m.lingTongFashionMap = make(map[int32]*PlayerLingTongFashionObject)
	//加载玩家灵童时装
	fashionList, err := dao.GetLingTongDao().GetLingTongFashionList(m.p.GetId())
	if err != nil {
		return
	}
	for _, fashionObj := range fashionList {
		pwo := NewPlayerLingTongFashionObject(m.p)
		pwo.FromEntity(fashionObj)
		m.lingTongFashionMap[pwo.lingTongId] = pwo
	}
	return
}

func (m *PlayerLingTongDataManager) initPlayerLingTongFashionObject(lingTongId int32) (obj *PlayerLingTongFashionObject) {
	obj, ok := m.lingTongFashionMap[lingTongId]
	if ok {
		return
	}
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return
	}
	obj = NewPlayerLingTongFashionObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	//生成id
	obj.playerId = m.p.GetId()
	obj.lingTongId = lingTongId
	obj.fashionId = lingTongTemplate.LingTongFashionId
	obj.createTime = now
	m.lingTongFashionMap[lingTongId] = obj
	obj.SetModified()
	return
}

func (m *PlayerLingTongDataManager) LingTongActivateFashionInit(lingtTongId int32) {
	m.initPlayerLingTongFashionObject(lingtTongId)
}

func (m *PlayerLingTongDataManager) GetLingTongFashion() map[int32]*PlayerLingTongFashionObject {
	return m.lingTongFashionMap
}

func (m *PlayerLingTongDataManager) GetLingTongFashionById(lingTongId int32) *PlayerLingTongFashionObject {
	obj, ok := m.lingTongFashionMap[lingTongId]
	if !ok {
		return nil
	}
	return obj
}

func (m *PlayerLingTongDataManager) refreshLingTongFashion() {
	now := global.GetGame().GetTimeService().Now()
	lingTongObj := m.GetLingTong()
	for lingTongId, fashionObj := range m.lingTongFashionMap {
		fashionId := fashionObj.GetFashionId()
		flag := lingtongtemplate.GetLingTongTemplateService().IsBornFashion(fashionId)
		if flag {
			continue
		}
		if m.getActivateFashionById(fashionId) != nil {
			continue
		}
		trialObj := m.getFashionTrailObject()
		if trialObj.GetTrialFashionId() == fashionId && trialObj.GetIsExpire() {
			lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
			if lingTongTemplate == nil {
				continue
			}
			fashionObj.fashionId = lingTongTemplate.LingTongFashionId
			fashionObj.updateTime = now
			fashionObj.SetModified()

			if lingTongObj.GetLingTongId() == lingTongId {
				gameevent.Emit(lingtongeventtypes.EventTypeLingTongFashionChanged, m.p, fashionObj)
			}
		}
	}
}

func (m *PlayerLingTongDataManager) trialFashionUse(trialFashionId int32) {
	lingTongObj := m.GetLingTong()
	if lingTongObj == nil {
		return
	}
	lingTongId := lingTongObj.GetLingTongId()
	lingTongFashionObj := m.GetLingTongFashionById(lingTongId)
	if lingTongFashionObj == nil {
		return
	}
	fashionId := lingTongFashionObj.GetFashionId()
	lingTongFashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
	if lingTongFashionTemplate == nil {
		return
	}
	flag := lingtongtemplate.GetLingTongTemplateService().IsBornFashion(fashionId)
	if !flag {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	lingTongFashionObj.fashionId = trialFashionId
	lingTongFashionObj.updateTime = now
	lingTongFashionObj.SetModified()
	gameevent.Emit(lingtongeventtypes.EventTypeLingTongFashionChanged, m.p, lingTongFashionObj)
}

func (m *PlayerLingTongDataManager) fashionExipre(exipreFashionId int32) {
	lingTongObj := m.GetLingTong()
	if lingTongObj == nil {
		return
	}
	lingTongId := lingTongObj.GetLingTongId()
	lingTongFashionObj := m.GetLingTongFashionById(lingTongId)
	if lingTongFashionObj == nil {
		return
	}
	fashionId := lingTongFashionObj.GetFashionId()
	lingTongFashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
	if lingTongFashionTemplate == nil {
		return
	}
	if fashionId != exipreFashionId {
		return
	}
	m.Unload()
}

func (m *PlayerLingTongDataManager) FashionWear(fashionId int32) (flag bool) {
	lingTongFashiongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
	if lingTongFashiongTemplate == nil {
		return
	}
	isBorn := lingtongtemplate.GetLingTongTemplateService().IsBornFashion(fashionId)
	if isBorn {
		return
	}

	lingTongObj := m.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		return
	}

	fashionInfo := m.GetFashionInfoById(fashionId)
	if fashionInfo == nil {
		return
	}

	lingTongId := lingTongObj.GetLingTongId()
	lingTongFashionObj := m.GetLingTongFashionById(lingTongId)
	if lingTongFashionObj == nil {
		return
	}

	if lingTongFashionObj.GetFashionId() == fashionId {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	lingTongFashionObj.fashionId = fashionId
	lingTongFashionObj.updateTime = now
	lingTongFashionObj.SetModified()
	gameevent.Emit(lingtongeventtypes.EventTypeLingTongFashionChanged, m.p, lingTongFashionObj)
	flag = true
	return
}

func (m *PlayerLingTongDataManager) Unload() (wearFashionId int32, flag bool) {
	lingTongObj := m.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		return
	}

	lingTongId := lingTongObj.GetLingTongId()
	lingTongFashionObj := m.GetLingTongFashionById(lingTongId)
	if lingTongFashionObj == nil {
		return
	}
	fashionId := lingTongFashionObj.GetFashionId()
	isBorn := lingtongtemplate.GetLingTongTemplateService().IsBornFashion(fashionId)
	if isBorn {
		return
	}

	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	lingTongFashionObj.fashionId = lingTongTemplate.LingTongFashionId
	lingTongFashionObj.updateTime = now
	lingTongFashionObj.SetModified()
	gameevent.Emit(lingtongeventtypes.EventTypeLingTongFashionChanged, m.p, lingTongFashionObj)
	flag = true
	wearFashionId = lingTongTemplate.LingTongFashionId
	return
}
