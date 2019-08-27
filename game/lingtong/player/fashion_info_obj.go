package player

import (
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lingtong/dao"
	"fgame/fgame/game/lingtong/entity"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/lingtong/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	"fmt"
)

//灵童时装信息
type PlayerLingTongFashionInfoObject struct {
	player       player.Player
	id           int64
	playerId     int64
	fashionId    int32
	upgradeLevel int32
	upgradeNum   int32
	upgradePro   int32
	isExpire     int32
	activateTime int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewPlayerLingTongFashionInfoObject(pl player.Player) *PlayerLingTongFashionInfoObject {
	pwo := &PlayerLingTongFashionInfoObject{
		player: pl,
	}
	return pwo
}

func convertLingTongFashionInfoObjectToEntity(pwo *PlayerLingTongFashionInfoObject) (*entity.PlayerLingTongFashionInfoEntity, error) {
	e := &entity.PlayerLingTongFashionInfoEntity{
		Id:           pwo.id,
		PlayerId:     pwo.playerId,
		FashionId:    pwo.fashionId,
		UpgradeLevel: pwo.upgradeLevel,
		UpgradeNum:   pwo.upgradeNum,
		UpgradePro:   pwo.upgradePro,
		IsExpire:     pwo.isExpire,
		ActivateTime: pwo.activateTime,
		UpdateTime:   pwo.updateTime,
		CreateTime:   pwo.createTime,
		DeleteTime:   pwo.deleteTime,
	}
	return e, nil
}

func (pwo *PlayerLingTongFashionInfoObject) GetPlayerId() int64 {
	return pwo.playerId
}

func (pwo *PlayerLingTongFashionInfoObject) GetDBId() int64 {
	return pwo.id
}

func (pwo *PlayerLingTongFashionInfoObject) GetFashionId() int32 {
	return pwo.fashionId
}

func (pwo *PlayerLingTongFashionInfoObject) GetUpgradeLevel() int32 {
	return pwo.upgradeLevel
}

func (pwo *PlayerLingTongFashionInfoObject) GetUpradeNum() int32 {
	return pwo.upgradeNum
}

func (pwo *PlayerLingTongFashionInfoObject) GetUpgradePro() int32 {
	return pwo.upgradePro
}

func (pwo *PlayerLingTongFashionInfoObject) GetIsExpire() bool {
	return pwo.isExpire == 1
}

func (pwo *PlayerLingTongFashionInfoObject) GetActivateTime() int64 {
	return pwo.activateTime
}

func (pwo *PlayerLingTongFashionInfoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertLingTongFashionInfoObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerLingTongFashionInfoObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerLingTongFashionInfoEntity)

	pwo.id = pse.Id
	pwo.playerId = pse.PlayerId
	pwo.fashionId = pse.FashionId
	pwo.upgradeLevel = pse.UpgradeLevel
	pwo.upgradeNum = pse.UpgradeNum
	pwo.upgradePro = pse.UpgradePro
	pwo.isExpire = pse.IsExpire
	pwo.activateTime = pse.ActivateTime
	pwo.updateTime = pse.UpdateTime
	pwo.createTime = pse.CreateTime
	pwo.deleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerLingTongFashionInfoObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("lingtongfashioninfo: err[%s]", err.Error()))
	}
	obj, ok := e.(playertypes.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}

func (pwo *PlayerLingTongFashionInfoObject) refresh() (activeFlag bool) {
	now := global.GetGame().GetTimeService().Now()
	activeFlag = true
	fashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(pwo.GetFashionId())
	if fashionTemplate == nil {
		activeFlag = false
		return
	}
	//时装判断
	fashionType := fashionTemplate.GetFashionType()
	if fashionType == types.LingTongFashionTypeNormal {
		return
	}

	if pwo.GetIsExpire() {
		activeFlag = false
		return
	}

	existTime := int64(fashionTemplate.Time)
	diffTime := now - pwo.GetActivateTime()
	if diffTime >= existTime {
		pwo.isExpire = 1
		pwo.activateTime = 0
		pwo.SetModified()
		activeFlag = false
	}
	return
}

type LingTongFashionContainer struct {
	p player.Player
	//激活时装
	activateMap map[int32]*PlayerLingTongFashionInfoObject
	//未激活时效时装
	noExpireFahionMap map[int32]*PlayerLingTongFashionInfoObject
}

func NewLingTongFashionContainer(p player.Player) *LingTongFashionContainer {
	d := &LingTongFashionContainer{
		p:                 p,
		activateMap:       make(map[int32]*PlayerLingTongFashionInfoObject),
		noExpireFahionMap: make(map[int32]*PlayerLingTongFashionInfoObject),
	}
	return d
}

func (m *LingTongFashionContainer) GetActivateFashionMap() map[int32]*PlayerLingTongFashionInfoObject {
	return m.activateMap
}

func (m *LingTongFashionContainer) GetFashionObj(fashionId int32) *PlayerLingTongFashionInfoObject {
	return m.getFashionObj(fashionId)
}

func (m *LingTongFashionContainer) getFashionObj(fashionId int32) *PlayerLingTongFashionInfoObject {
	fashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
	if fashionTemplate == nil {
		return nil
	}
	obj, ok := m.activateMap[fashionId]
	if !ok {
		return nil
	}
	return obj
}

func (m *PlayerLingTongDataManager) loadFashion() (err error) {
	m.fashionMapObj = NewLingTongFashionContainer(m.p)
	//加载玩家灵童时装信息
	fashionInfoList, err := dao.GetLingTongDao().GetLingTongFashionInfoList(m.p.GetId())
	if err != nil {
		return
	}
	for _, fashionObj := range fashionInfoList {
		pwo := NewPlayerLingTongFashionInfoObject(m.p)
		pwo.FromEntity(fashionObj)
		pwo.refresh()
		m.addFashion(pwo)
	}
	return
}

func (m *PlayerLingTongDataManager) addFashion(pwo *PlayerLingTongFashionInfoObject) {
	if !pwo.GetIsExpire() {
		m.fashionMapObj.activateMap[pwo.GetFashionId()] = pwo
	} else {
		m.fashionMapObj.noExpireFahionMap[pwo.GetFashionId()] = pwo
	}
}

func (m *PlayerLingTongDataManager) initPlayerLingTongFashionInfoObject(fashionId int32, now int64) (obj *PlayerLingTongFashionInfoObject) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
	if lingTongTemplate == nil {
		return
	}
	flag := lingtongtemplate.GetLingTongTemplateService().IsBornFashion(fashionId)
	if flag {
		return
	}
	obj = m.getActivateFashionById(fashionId)
	if obj != nil {
		return
	}
	obj = m.getNoExpireFahionById(fashionId)
	if obj != nil {
		return
	}
	obj = NewPlayerLingTongFashionInfoObject(m.p)

	id, _ := idutil.GetId()
	obj.id = id
	//生成id
	obj.playerId = m.p.GetId()
	obj.fashionId = fashionId
	obj.upgradeLevel = 0
	obj.upgradeNum = 0
	obj.upgradePro = 0
	obj.activateTime = now
	obj.isExpire = 0
	obj.createTime = now
	m.fashionMapObj.activateMap[fashionId] = obj
	obj.SetModified()
	return
}

func (m *PlayerLingTongDataManager) getActivateFashionMap() map[int32]*PlayerLingTongFashionInfoObject {
	if m.fashionMapObj == nil {
		return nil
	}
	return m.fashionMapObj.activateMap
}

func (m *PlayerLingTongDataManager) getNoExpireFashionMap() map[int32]*PlayerLingTongFashionInfoObject {
	if m.fashionMapObj == nil {
		return nil
	}
	return m.fashionMapObj.noExpireFahionMap
}

func (m *PlayerLingTongDataManager) getActivateFashionById(fashionId int32) *PlayerLingTongFashionInfoObject {
	activateMap := m.getActivateFashionMap()
	if activateMap == nil {
		return nil
	}
	obj, ok := activateMap[fashionId]
	if !ok {
		return nil
	}
	return obj
}

func (m *PlayerLingTongDataManager) getNoExpireFahionById(fashionId int32) *PlayerLingTongFashionInfoObject {
	noExpireMap := m.getNoExpireFashionMap()
	if noExpireMap == nil {
		return nil
	}
	obj, ok := noExpireMap[fashionId]
	if !ok {
		return nil
	}
	return obj
}

func (m *PlayerLingTongDataManager) activateNoExpireFashion(fashionId int32, now int64) {
	noExpireMap := m.getNoExpireFashionMap()
	if noExpireMap == nil {
		return
	}
	obj, ok := noExpireMap[fashionId]
	if !ok {
		return
	}
	defer delete(noExpireMap, fashionId)
	obj.activateTime = now
	obj.updateTime = now
	m.fashionMapObj.activateMap[fashionId] = obj
}

func (m *PlayerLingTongDataManager) GetActivateFashionMap() map[int32]*PlayerLingTongFashionInfoObject {
	return m.getActivateFashionMap()
}

func (m *PlayerLingTongDataManager) GetFashionInfoById(fashionId int32) *PlayerLingTongFashionInfoObject {
	return m.getActivateFashionById(fashionId)
}

func (m *PlayerLingTongDataManager) fashionRefreshCheck(fashionId int32, now int64) (expireFlag bool) {
	fashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
	if fashionTemplate == nil {
		return
	}
	fashionInfo := m.GetFashionInfoById(fashionId)
	if fashionInfo == nil {
		return
	}
	expireFlag = false
	//时装判断
	existTime := int64(fashionTemplate.Time)
	if existTime == 0 {
		return
	}

	diffTime := now - fashionInfo.GetActivateTime()
	if diffTime >= existTime {
		fashionInfo.isExpire = 1
		fashionInfo.activateTime = 0
		fashionInfo.updateTime = now
		expireFlag = true
		fashionInfo.SetModified()
		delete(m.fashionMapObj.activateMap, fashionId)
		m.fashionMapObj.noExpireFahionMap[fashionId] = fashionInfo
		m.fashionExipre(fashionId)
	}

	return
}

//时装激活
func (m *PlayerLingTongDataManager) FashionActive(fashionId int32) (obj *PlayerLingTongFashionInfoObject, flag bool) {
	fashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
	if fashionTemplate == nil {
		return
	}
	isBorn := lingtongtemplate.GetLingTongTemplateService().IsBornFashion(fashionId)
	if isBorn {
		return
	}
	fashionObj := m.GetFashionInfoById(fashionId)
	if fashionObj != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	fashionObj = m.getNoExpireFahionById(fashionId)
	if fashionObj != nil {
		m.activateNoExpireFashion(fashionId, now)
	} else {
		m.initPlayerLingTongFashionInfoObject(fashionId, now)
	}

	m.TrialFashionOverdue(fashionId, types.LingTongFashionTrialOverdueTypeActivate)
	gameevent.Emit(lingtongeventtypes.EventTypeLingTongFashionActivate, m.p, fashionId)
	flag = true
	obj = m.GetFashionInfoById(fashionId)
	return
}

//是否能升星
func (m *PlayerLingTongDataManager) IfCanUpStar(fashionId int32) (flag bool) {
	fashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
	if fashionTemplate == nil {
		return
	}
	fashionObj := m.GetFashionInfoById(fashionId)
	if fashionObj == nil {
		return
	}

	if fashionTemplate.LingTongUpstarId == 0 {
		return
	}

	star := fashionObj.GetUpgradeLevel()
	if star <= 0 {
		return true
	}
	nextTo := fashionTemplate.GetLingTongFashionUpstarByLevel(star)
	if nextTo.NextId != 0 {
		return true
	}
	return false
}

//灵童时装升星
func (m *PlayerLingTongDataManager) Upstar(fashionId int32, pro int32, sucess bool) (flag bool) {
	flag = m.IfCanUpStar(fashionId)
	if !flag {
		return
	}
	obj := m.GetFashionInfoById(fashionId)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		obj.upgradeLevel += 1
		obj.upgradeNum = 0
		obj.upgradePro = pro
	} else {
		obj.upgradeNum += 1
		obj.upgradePro += pro
	}
	obj.updateTime = now
	obj.SetModified()
	return true
}
