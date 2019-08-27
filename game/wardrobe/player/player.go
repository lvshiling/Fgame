package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/wardrobe/dao"
	wardrobeeventtypes "fgame/fgame/game/wardrobe/event/types"
	wardrobetemplate "fgame/fgame/game/wardrobe/template"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	"fgame/fgame/pkg/idutil"
)

//玩家衣橱管理器
type PlayerWardrobeDataManager struct {
	p player.Player
	//玩家衣橱map
	playerWardrobeMap map[int32]map[int32]*PlayerWardrobeObject
	//玩家衣橱套装培养
	playerPeiYangMap map[int32]*PlayerWardrobePeiYangObject
}

func (m *PlayerWardrobeDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerWardrobeDataManager) Load() (err error) {
	m.playerWardrobeMap = make(map[int32]map[int32]*PlayerWardrobeObject)
	m.playerPeiYangMap = make(map[int32]*PlayerWardrobePeiYangObject)
	m.loadWardrobe()
	m.loadPeiYang()
	return nil
}

func (m *PlayerWardrobeDataManager) loadWardrobe() (err error) {
	//加载玩家衣橱信息
	wardrobeList, err := dao.GetWardrobeDao().GetWardrobeList(m.p.GetId())
	if err != nil {
		return
	}
	for _, wardrobeObj := range wardrobeList {
		obj := NewPlayerWardrobeObject(m.p)
		obj.FromEntity(wardrobeObj)
		m.addSuit(obj)
	}
	return
}

func (m *PlayerWardrobeDataManager) loadPeiYang() (err error) {
	//加载玩家衣橱信息
	wardrobeList, err := dao.GetWardrobeDao().GetWardrobePeiYangList(m.p.GetId())
	if err != nil {
		return
	}
	for _, wardrobeObj := range wardrobeList {
		obj := NewPlayerWardrobePeiYangObject(m.p)
		obj.FromEntity(wardrobeObj)
		m.playerPeiYangMap[obj.GetType()] = obj
	}
	return
}

func (m *PlayerWardrobeDataManager) addSuit(obj *PlayerWardrobeObject) {
	typ := obj.GetType()
	subType := obj.GetSubType()
	suitMap, ok := m.playerWardrobeMap[typ]
	if !ok {
		suitMap = make(map[int32]*PlayerWardrobeObject)
		m.playerWardrobeMap[typ] = suitMap
	}
	suitMap[subType] = obj
}

//加载后
func (m *PlayerWardrobeDataManager) AfterLoad() (err error) {

	return nil
}

//心跳
func (m *PlayerWardrobeDataManager) Heartbeat() {

}

func (m *PlayerWardrobeDataManager) initMap() {
	max := wardrobetemplate.GetWardrobeTemplateService().GetYiChuMaxType()
	for typ := int32(0); typ <= max; typ++ {
		m.CheckActive(typ, true)
	}
}

func (m *PlayerWardrobeDataManager) initPlayerWardrobeObj(typ int32,
	subType int32) (obj *PlayerWardrobeObject) {

	suitMap, ok := m.playerWardrobeMap[typ]
	if !ok {
		suitMap = make(map[int32]*PlayerWardrobeObject)
		m.playerWardrobeMap[typ] = suitMap
	}
	obj, ok = suitMap[subType]
	if ok {
		return obj
	}
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj = NewPlayerWardrobeObject(m.p)
	obj.playerId = m.p.GetId()
	obj.id = id
	obj.typ = typ
	obj.subType = subType
	obj.createTime = now
	obj.SetModified()
	m.addSuit(obj)
	return obj
}

func (m *PlayerWardrobeDataManager) initPlayerPeiYangObj(typ int32) (obj *PlayerWardrobePeiYangObject) {
	obj, ok := m.playerPeiYangMap[typ]
	if ok {
		return
	}
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj = NewPlayerWardrobePeiYangObject(m.p)
	obj.playerId = m.p.GetId()
	obj.id = id
	obj.typ = typ
	obj.peiYangNum = 0
	obj.createTime = now
	obj.SetModified()
	m.playerPeiYangMap[typ] = obj
	return obj
}

//登录下发刷新
func (m *PlayerWardrobeDataManager) RefreshAfterLoad() {
	m.refreshSuitMap()
	m.initMap()
}

func (m *PlayerWardrobeDataManager) refreshSuitMap() {
	now := global.GetGame().GetTimeService().Now()
	//校验是否有失效
	for _, suitMap := range m.playerWardrobeMap {
		for _, wardrobeObj := range suitMap {
			//可能策划改配置
			// if wardrobeObj.GetIsPermanent() {
			// 	continue
			// }
			if !wardrobeObj.GetIsActive() {
				continue
			}
			typ := wardrobeObj.GetType()
			subType := wardrobeObj.GetSubType()
			wardrobeTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuTemplate(typ, subType)
			if wardrobeTemplate == nil {
				continue
			}
			_, totalNum := wardrobetemplate.GetWardrobeTemplateService().GetYiChuActiveNum(m.p, typ)
			if totalNum < wardrobeTemplate.Number {
				wardrobeObj.activeFlag = 0
				wardrobeObj.updateTime = now
				wardrobeObj.SetModified()
				gameevent.Emit(wardrobeeventtypes.EventTypeWardrobeRemove, m.p, wardrobeObj)
			}
		}
	}
}

func (m *PlayerWardrobeDataManager) GetWardrobeMap() map[int32]map[int32]*PlayerWardrobeObject {
	return m.playerWardrobeMap
}

func (m *PlayerWardrobeDataManager) GetWardrobeMapByType(typ int32) map[int32]*PlayerWardrobeObject {
	suitMap, ok := m.playerWardrobeMap[typ]
	if !ok {
		return nil
	}
	return suitMap
}

func (m *PlayerWardrobeDataManager) GetWardrobeActivateNumByType(typ int32) (num int32) {
	suitMap, ok := m.playerWardrobeMap[typ]
	if !ok {
		return
	}
	for _, obj := range suitMap {
		if obj.GetIsActive() {
			num++
		}
	}
	return
}

func (m *PlayerWardrobeDataManager) GetWardrobeActivateTypeNumByNum(subTypeInt int32) (totalNum int32) {
	for _, suitMap := range m.playerWardrobeMap {
		for curSubType, _ := range suitMap {
			if curSubType == subTypeInt {
				totalNum++
			}
		}
	}
	return
}

func (m *PlayerWardrobeDataManager) GetWardrobeByType(typ int32, subType int32) *PlayerWardrobeObject {
	suitMap := m.GetWardrobeMapByType(typ)
	if suitMap == nil {
		return nil
	}
	obj, ok := suitMap[subType]
	if !ok {
		return nil
	}
	return obj
}

func (m *PlayerWardrobeDataManager) IsWardrobeActive(typ int32) (flag bool) {
	suitMap := m.GetWardrobeMapByType(typ)
	if suitMap == nil {
		return
	}

	for _, wardrobeObj := range suitMap {
		if wardrobeObj.GetIsActive() {
			return true
		}
	}
	return
}

func (m *PlayerWardrobeDataManager) GetWardrobePeiYangNum(typ int32) int32 {
	peiYangObj, ok := m.playerPeiYangMap[typ]
	if !ok {
		return 0
	}
	return peiYangObj.peiYangNum
}

func (m *PlayerWardrobeDataManager) GetWardrobePeiYang(typ int32) *PlayerWardrobePeiYangObject {
	peiYangObj, ok := m.playerPeiYangMap[typ]
	if !ok {
		return nil
	}
	return peiYangObj
}

func (m *PlayerWardrobeDataManager) ActiveSeqId(sysType wardrobetypes.WardrobeSysType, seqId int32) {
	max := wardrobetemplate.GetWardrobeTemplateService().GetYiChuMaxType()
	for typ := int32(0); typ <= max; typ++ {
		suitTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSuitTemplate(typ)
		if !suitTemplate.IfExist(sysType, seqId) {
			continue
		}
		m.CheckActive(typ, true)
	}
}

func (m *PlayerWardrobeDataManager) CheckActive(typ int32, sendEvent bool) {
	now := global.GetGame().GetTimeService().Now()
	_, totalNum := wardrobetemplate.GetWardrobeTemplateService().GetYiChuActiveNum(m.p, typ)
	subTypeList := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSubTypeList(typ)
	for _, subType := range subTypeList {
		obj := m.GetWardrobeByType(typ, subType)
		//可能策划改配置
		if obj != nil && obj.GetIsActive() {
			continue
		}
		// if obj != nil && obj.GetIsPermanent() {
		// 	continue
		// }
		wardrobeTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuTemplate(typ, subType)
		if wardrobeTemplate == nil {
			continue
		}

		oldActive := false
		//找到第一件不满足的就可以退出了
		if totalNum < wardrobeTemplate.Number {
			break
		}
		newActive := true
		//激活过
		if obj != nil {
			oldActive = obj.GetIsActive()
		} else {
			obj = m.initPlayerWardrobeObj(typ, subType)
		}

		//可能策划改配置
		//永久激活
		// if permanentNum >= wardrobeTemplate.Number {
		// 	obj.activeFlag = 1
		// 	obj.permanent = 1
		// 	obj.updateTime = now
		// 	obj.SetModified()
		// } else {
		// 	obj.activeFlag = 1
		// 	obj.updateTime = now
		// 	obj.SetModified()
		// }

		obj.activeFlag = 1
		obj.updateTime = now
		obj.SetModified()

		//激活衣橱套装了
		if sendEvent && !oldActive && newActive {
			gameevent.Emit(wardrobeeventtypes.EventTypeWardrobeActive, m.p, obj)
		}
	}
}

func (m *PlayerWardrobeDataManager) RemoveSeqId(sysType wardrobetypes.WardrobeSysType, seqId int32) {
	now := global.GetGame().GetTimeService().Now()
	for typ, suitMap := range m.playerWardrobeMap {
		suitTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSuitTemplate(typ)
		if suitTemplate == nil {
			continue
		}
		if !suitTemplate.IfExist(sysType, seqId) {
			continue
		}
		_, totalNum := wardrobetemplate.GetWardrobeTemplateService().GetYiChuActiveNum(m.p, typ)
		for subType, wardrobeObj := range suitMap {
			//可能策划改配置
			// if wardrobeObj.GetIsPermanent() {
			// 	continue
			// }
			if !wardrobeObj.GetIsActive() {
				continue
			}
			typ := wardrobeObj.GetType()
			wardrobeTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuTemplate(typ, subType)
			if wardrobeTemplate == nil {
				continue
			}
			if totalNum >= wardrobeTemplate.Number {
				continue
			}
			wardrobeObj.activeFlag = 0
			wardrobeObj.updateTime = now
			wardrobeObj.SetModified()
			gameevent.Emit(wardrobeeventtypes.EventTypeWardrobeRemove, m.p, wardrobeObj)
		}
	}
}

//是否能培养
func (m *PlayerWardrobeDataManager) IfCanPeiYang(typ int32) (eatNum int32, flag bool) {
	suitMap := m.GetWardrobeMapByType(typ)
	if suitMap == nil {
		return
	}
	curNum := m.GetWardrobePeiYangNum(typ)
	maxNum := int32(0)
	for subType, wardrobeObj := range suitMap {
		if !wardrobeObj.GetIsActive() {
			continue
		}
		wardrobeTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuTemplate(typ, subType)
		if wardrobeTemplate == nil {
			continue
		}
		if curNum < wardrobeTemplate.ShiDanLimit {
			flag = true
			eatNum = wardrobeTemplate.ShiDanLimit - curNum
		}

		if eatNum > maxNum {
			maxNum = eatNum
		}
	}
	eatNum = maxNum
	return
}

func (m *PlayerWardrobeDataManager) EatCulDan(typ int32, num int32) {
	if num <= 0 {
		return
	}

	to := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSuitTemplate(typ)
	if to == nil {
		return
	}
	culTemplate := to.GetPeiYangByLevel(num)
	if culTemplate == nil {
		return
	}

	oldLevel := int32(0)
	obj := m.GetWardrobePeiYang(typ)
	if obj == nil {
		obj = m.initPlayerPeiYangObj(typ)
	} else {
		oldLevel = obj.GetPeiYangNum()
	}

	obj.peiYangNum = num
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	eventData := wardrobeeventtypes.CreateWardrobePeiYangEventData(typ, oldLevel, num)
	gameevent.Emit(wardrobeeventtypes.EventTypeWardrobePeiYangUpgrade, m.p, eventData)
	return
}

//仅gm使用
func (m *PlayerWardrobeDataManager) GmSetPeiYangLevel(typ int32, num int32) {
	now := global.GetGame().GetTimeService().Now()
	obj, ok := m.playerPeiYangMap[typ]
	if !ok {
		return
	}
	obj.peiYangNum = num
	obj.updateTime = now
	obj.SetModified()
}

func CreatePlayerWardrobeDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerWardrobeDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerWardrobeDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerWardrobeDataManager))
}
