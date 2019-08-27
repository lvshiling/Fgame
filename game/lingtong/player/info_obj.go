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
	xiuxianbookeventtypes "fgame/fgame/game/welfare/xiuxianbook/event/types"
	"fgame/fgame/pkg/idutil"
	"fmt"
)

//灵童激活
type PlayerLingTongInfoObject struct {
	player       player.Player
	id           int64
	playerId     int64
	lingTongId   int32
	lingTongName string
	upgradeLevel int32
	upgradeNum   int32
	upgradePro   int32
	peiYangLevel int32
	peiYangNum   int32
	peiYangPro   int32
	starLevel    int32
	starNum      int32
	starPro      int32
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewPlayerLingTongInfoObject(pl player.Player) *PlayerLingTongInfoObject {
	pwo := &PlayerLingTongInfoObject{
		player: pl,
	}
	return pwo
}

func convertLingTongInfoObjectToEntity(pwo *PlayerLingTongInfoObject) (*entity.PlayerLingTongInfoEntity, error) {
	e := &entity.PlayerLingTongInfoEntity{
		Id:           pwo.id,
		PlayerId:     pwo.playerId,
		LingTongId:   pwo.lingTongId,
		LingTongName: pwo.lingTongName,
		UpgradeLevel: pwo.upgradeLevel,
		UpgradeNum:   pwo.upgradeNum,
		UpgradePro:   pwo.upgradePro,
		PeiYangLevel: pwo.peiYangLevel,
		PeiYangNum:   pwo.peiYangNum,
		PeiYangPro:   pwo.peiYangPro,
		StarLevel:    pwo.starLevel,
		StarNum:      pwo.starNum,
		StarPro:      pwo.starPro,
		UpdateTime:   pwo.updateTime,
		CreateTime:   pwo.createTime,
		DeleteTime:   pwo.deleteTime,
	}
	return e, nil
}

func (pwo *PlayerLingTongInfoObject) GetPlayerId() int64 {
	return pwo.playerId
}

func (pwo *PlayerLingTongInfoObject) GetDBId() int64 {
	return pwo.id
}

func (pwo *PlayerLingTongInfoObject) GetLingTongId() int32 {
	return pwo.lingTongId
}

func (pwo *PlayerLingTongInfoObject) GetLingTongName() string {
	return pwo.lingTongName
}

func (pwo *PlayerLingTongInfoObject) GetLevel() int32 {
	return pwo.upgradeLevel
}

func (pwo *PlayerLingTongInfoObject) GetNum() int32 {
	return pwo.upgradeNum
}

func (pwo *PlayerLingTongInfoObject) GetPro() int32 {
	return pwo.upgradePro
}

func (pwo *PlayerLingTongInfoObject) GetPeiYangLevel() int32 {
	return pwo.peiYangLevel
}

func (pwo *PlayerLingTongInfoObject) GetPeiYangNum() int32 {
	return pwo.peiYangNum
}

func (pwo *PlayerLingTongInfoObject) GetPeiYangPro() int32 {
	return pwo.peiYangPro
}

func (pwo *PlayerLingTongInfoObject) GetStarLevel() int32 {
	return pwo.starLevel
}

func (pwo *PlayerLingTongInfoObject) GetStarNum() int32 {
	return pwo.starNum
}

func (pwo *PlayerLingTongInfoObject) GetStarPro() int32 {
	return pwo.starPro
}

func (pwo *PlayerLingTongInfoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertLingTongInfoObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerLingTongInfoObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerLingTongInfoEntity)

	pwo.id = pse.Id
	pwo.playerId = pse.PlayerId
	pwo.lingTongId = pse.LingTongId
	pwo.lingTongName = pse.LingTongName
	pwo.upgradeLevel = pse.UpgradeLevel
	pwo.upgradeNum = pse.UpgradeNum
	pwo.upgradePro = pse.UpgradePro
	pwo.peiYangLevel = pse.PeiYangLevel
	pwo.peiYangNum = pse.PeiYangNum
	pwo.peiYangPro = pse.PeiYangPro
	pwo.starLevel = pse.StarLevel
	pwo.starNum = pse.StarNum
	pwo.starPro = pse.StarPro
	pwo.updateTime = pse.UpdateTime
	pwo.createTime = pse.CreateTime
	pwo.deleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerLingTongInfoObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("lingtonginfo: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}

func (m *PlayerLingTongDataManager) loadLingTongInfo() (err error) {
	m.lingTongMap = make(map[int32]*PlayerLingTongInfoObject)
	//加载玩家灵童信息
	lingTongInfoList, err := dao.GetLingTongDao().GetLingTongInfoList(m.p.GetId())
	if err != nil {
		return
	}
	for _, lingTongInfoObj := range lingTongInfoList {
		pwo := NewPlayerLingTongInfoObject(m.p)
		pwo.FromEntity(lingTongInfoObj)
		m.lingTongMap[pwo.lingTongId] = pwo
	}
	return
}

func (m *PlayerLingTongDataManager) initPlayerLingTongInfoObject(lingTongId int32) (obj *PlayerLingTongInfoObject) {
	obj, ok := m.lingTongMap[lingTongId]
	if ok {
		return
	}
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return
	}
	obj = NewPlayerLingTongInfoObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	//生成id
	obj.playerId = m.p.GetId()
	obj.lingTongId = lingTongId
	obj.lingTongName = lingTongTemplate.Name
	obj.upgradeLevel = 0
	obj.upgradeNum = 0
	obj.upgradePro = 0
	obj.peiYangLevel = 0
	obj.peiYangNum = 0
	obj.peiYangPro = 0
	obj.starLevel = 0
	obj.starNum = 0
	obj.starPro = 0
	obj.createTime = now
	m.lingTongMap[lingTongId] = obj
	obj.SetModified()
	return
}

func (m *PlayerLingTongDataManager) GetLingTongMap() map[int32]*PlayerLingTongInfoObject {
	return m.lingTongMap
}

func (m *PlayerLingTongDataManager) GetLingTongInfo(lingTongId int32) (obj *PlayerLingTongInfoObject, flag bool) {
	obj, ok := m.lingTongMap[lingTongId]
	if !ok {
		return
	}
	flag = true
	return
}

func (m *PlayerLingTongDataManager) LingTongActivate(lingTongId int32) (obj *PlayerLingTongInfoObject) {
	_, flag := m.GetLingTongInfo(lingTongId)
	if flag {
		return
	}
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return
	}
	obj = m.initPlayerLingTongInfoObject(lingTongId)
	gameevent.Emit(lingtongeventtypes.EventTypeLingTongActivate, m.p, lingTongId)
	return obj
}

func (m *PlayerLingTongDataManager) ShengJi(lingTongId int32, pro int32, sucess bool) (flag bool) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return
	}
	obj, flag := m.GetLingTongInfo(lingTongId)
	if !flag {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextLevel := obj.GetLevel() + 1
		nextLingTongShengJiTemplate := lingTongTemplate.GetLingTongShengJiByLevel(nextLevel)
		if nextLingTongShengJiTemplate == nil {
			return
		}
		obj.upgradeLevel += 1
		obj.upgradeNum = 0
		obj.upgradePro = pro
	} else {
		obj.upgradeNum += 1
		obj.upgradePro += pro
	}
	obj.updateTime = now
	obj.SetModified()
	flag = true
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, m.p, nil)
	return
}

func (m *PlayerLingTongDataManager) EatCulDan(lingTongId int32, level int32) (flag bool) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return
	}
	obj, flag := m.GetLingTongInfo(lingTongId)
	if !flag {
		return
	}
	if obj.GetPeiYangLevel() == level || level <= 0 {
		return
	}
	peiYangTemplate := lingTongTemplate.GetLingTongPeiYangByLevel(level)
	if peiYangTemplate == nil {
		return
	}

	obj.peiYangLevel = level
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	return
}

func (m *PlayerLingTongDataManager) Rename(lingTongId int32, lingTongName string) (flag bool) {
	obj, flag := m.GetLingTongInfo(lingTongId)
	if !flag {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	obj.lingTongName = lingTongName
	obj.updateTime = now
	obj.SetModified()
	gameevent.Emit(lingtongeventtypes.EventTypeLingTongRename, m.p, obj)
	flag = true
	return
}

func (m *PlayerLingTongDataManager) LingTongUpstar(lingTongId int32, pro int32, sucess bool) (flag bool) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return
	}
	obj, flag := m.GetLingTongInfo(lingTongId)
	if !flag {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextLevel := obj.GetStarLevel() + 1
		nextLingTongUpstarTemplate := lingTongTemplate.GetLingTongUpstarByLevel(nextLevel)
		if nextLingTongUpstarTemplate == nil {
			return
		}
		obj.starLevel += 1
		obj.starNum = 0
		obj.starPro = pro
	} else {
		obj.starNum += 1
		obj.starPro += pro
	}
	obj.updateTime = now
	obj.SetModified()
	flag = true
	return
}
