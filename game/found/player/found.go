package player

import (
	"fgame/fgame/game/found/dao"
	"fgame/fgame/game/found/found"
	foundtemplate "fgame/fgame/game/found/template"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/idutil"
)

//玩家资源回收管理器
type PlayerFoundDataManager struct {
	p player.Player
	//当天资源找回记录
	curDayResRecordList []*PlayerFoundObject
	//上一天资源找回列表
	preDayFoundList []*PlayerFoundBackObject
}

func (m *PlayerFoundDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerFoundDataManager) Load() (err error) {
	//加载玩家资源回收列表
	foundResEntityList, err := dao.GetFoundResDao().GetFoundResList(m.p.GetId())
	if err != nil {
		return
	}
	for _, e := range foundResEntityList {
		obj := newPlayerFoundObject(m.p)
		obj.FromEntity(e)

		m.curDayResRecordList = append(m.curDayResRecordList, obj)
	}

	//加载上一天资源找回列表
	foundBackEntityList, err := dao.GetFoundResDao().GetFoundBackList(m.p.GetId())
	if err != nil {
		return
	}
	for _, e := range foundBackEntityList {
		backObj := newPlayerFoundBackObject(m.p)
		backObj.FromEntity(e)

		m.preDayFoundList = append(m.preDayFoundList, backObj)
	}
	return nil
}

func (m *PlayerFoundDataManager) newCurDayRecordObj() {
	resTypeMap := foundtypes.GetPlayModeTypeMap()
	for typ, playModel := range resTypeMap {
		obj := m.getCurDayResRecord(typ)
		if obj != nil {
			continue
		}

		id, _ := idutil.GetId()
		now := global.GetGame().GetTimeService().Now()
		obj = newPlayerFoundObject(m.p)
		obj.id = id
		obj.playModeType = playModel
		obj.joinTimes = 0
		obj.resType = typ
		obj.createTime = now
		obj.updateTime = now

		obj.SetModified()
		m.curDayResRecordList = append(m.curDayResRecordList, obj)

	}
}

func (m *PlayerFoundDataManager) newPreDayFoundObj(resType foundtypes.FoundResourceType) *PlayerFoundBackObject {
	backObj := newPlayerFoundBackObject(m.p)
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	backObj.id = id
	backObj.resType = resType
	backObj.status = foundtypes.FoundBackStatusWaitReceive
	backObj.foundTimes = 0
	backObj.createTime = now
	backObj.updateTime = now
	backObj.SetModified()

	m.preDayFoundList = append(m.preDayFoundList, backObj)

	return backObj
}

//加载后
func (m *PlayerFoundDataManager) AfterLoad() (err error) {
	m.newCurDayRecordObj()
	return
}

//心跳
func (m *PlayerFoundDataManager) Heartbeat() {
}

//更新资源找回列表
func (m *PlayerFoundDataManager) RefreshPreDayFoundList() {

	for _, obj := range m.curDayResRecordList {
		now := global.GetGame().GetTimeService().Now()
		var checkH found.FoundCheckHandler
		var h found.FoundObjDataHandler
		var resLevel int32
		var maxTimes int32
		var group int32
		var foundTem *gametemplate.FoundTemplate

		if !obj.isFoundBackDay() {
			continue
		}

		//找回条件判断
		checkH = found.GetFoundCheckHandler(obj.resType)
		if !checkH.IsCanFoundBack(obj.player) {
			goto NOT_FOUND
		}

		// 找回资源信息
		h = found.GetFoundDataHandler(obj.resType)
		if h == nil {
			goto NOT_FOUND
		}

		resLevel, maxTimes, group = h.GetFoundParam(m.p)
		foundTem = foundtemplate.GetFoundTemplateService().GetFoundTemplateByType(obj.resType, resLevel)
		if foundTem == nil {
			goto NOT_FOUND
		}

		// 功能开启
		if !m.p.IsFuncOpen(foundTem.GetFuncOpenType()) {
			goto NOT_FOUND
		}

		m.resetFoundBackObj(obj.resType, resLevel, group, maxTimes, obj.joinTimes)
		obj.joinTimes = 0
		obj.updateTime = now
		obj.SetModified()
		continue

	NOT_FOUND:
		m.notFoundBack(obj.resType)
	}
}

func (m *PlayerFoundDataManager) notFoundBack(resType foundtypes.FoundResourceType) {
	foundBackObj := m.getPreDayFoundObj(resType)
	if foundBackObj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	foundBackObj.status = foundtypes.FoundBackStatusNotFound
	foundBackObj.updateTime = now
	foundBackObj.SetModified()

}

func (m *PlayerFoundDataManager) resetFoundBackObj(resType foundtypes.FoundResourceType, resLevel, group, maxTimes, joinTimes int32) *PlayerFoundBackObject {
	foundBackObj := m.getPreDayFoundObj(resType)
	if foundBackObj == nil {
		foundBackObj = m.newPreDayFoundObj(resType)
	}

	foundTimes := found.GetFoundService().CountFoundBackTimes(resType, maxTimes, joinTimes)
	now := global.GetGame().GetTimeService().Now()

	foundBackObj.status = foundtypes.FoundBackStatusWaitReceive
	foundBackObj.foundTimes = foundTimes
	foundBackObj.resLevel = resLevel
	foundBackObj.group = group
	foundBackObj.createTime = now
	foundBackObj.updateTime = now
	foundBackObj.SetModified()

	return foundBackObj
}

//资源使用记录+1
func (m *PlayerFoundDataManager) IncreFoundResJoinTimes(resType foundtypes.FoundResourceType) {
	m.addFoundResJoinTimes(resType, 1)
}

//资源使用记录+1
func (m *PlayerFoundDataManager) IncreFoundResJoinTimesBatch(resType foundtypes.FoundResourceType, joinTimes int32) {
	m.addFoundResJoinTimes(resType, joinTimes)
}

func (m *PlayerFoundDataManager) addFoundResJoinTimes(resType foundtypes.FoundResourceType, joinTimes int32) {
	m.RefreshPreDayFoundList()

	obj := m.getCurDayResRecord(resType)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	obj.joinTimes += joinTimes
	obj.updateTime = now
	obj.SetModified()
}

//获取资源使用记录
func (m *PlayerFoundDataManager) getCurDayResRecord(typ foundtypes.FoundResourceType) *PlayerFoundObject {
	for _, obj := range m.curDayResRecordList {
		if typ == obj.resType {
			return obj
		}
	}
	return nil
}

//获取资源使用记录列表
func (m *PlayerFoundDataManager) GetCurDayResRecordList() []*PlayerFoundObject {
	return m.curDayResRecordList
}

//上一天的资源找回列表
func (m *PlayerFoundDataManager) GetPreDayFoundList() []*PlayerFoundBackObject {
	m.RefreshPreDayFoundList()

	var newList []*PlayerFoundBackObject
	for _, backObj := range m.preDayFoundList {
		if backObj.status == foundtypes.FoundBackStatusNotFound {
			continue
		}
		if backObj.group == 0 {
			continue
		}
		if backObj.foundTimes == 0 {
			continue
		}

		newList = append(newList, backObj)
	}
	return newList
}

//获取资源找回
func (m *PlayerFoundDataManager) getPreDayFoundObj(typ foundtypes.FoundResourceType) *PlayerFoundBackObject {
	for _, res := range m.preDayFoundList {
		if res.resType == typ {
			return res
		}
	}

	return nil
}

//是否存在资源找回
func (m *PlayerFoundDataManager) IsCanFoundBack(typ foundtypes.FoundResourceType) bool {
	res := m.getPreDayFoundObj(typ)
	if res == nil {
		return false
	}

	return res.foundTimes > 0
}

//是否领取资源找回
func (m *PlayerFoundDataManager) IsReceiveFound(typ foundtypes.FoundResourceType) bool {
	res := m.getPreDayFoundObj(typ)
	if res == nil {
		return false
	}
	if res.status == foundtypes.FoundBackStatusWaitReceive {
		return false
	}
	return true
}

//领取资源找回
func (m *PlayerFoundDataManager) ReceiveFound(typ foundtypes.FoundResourceType) {
	obj := m.getPreDayFoundObj(typ)
	if obj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj.status = foundtypes.FoundBackStatusHadReceive
	obj.updateTime = now
	obj.SetModified()
}

//获取找回次数
func (m *PlayerFoundDataManager) GetFoundTimes(resType foundtypes.FoundResourceType) int32 {
	preFoundObj := m.getPreDayFoundObj(resType)
	if preFoundObj == nil {
		return 0
	}

	return preFoundObj.foundTimes
}

//获取找回资源等级
func (m *PlayerFoundDataManager) GetFoundResLevel(resType foundtypes.FoundResourceType) int32 {
	preFoundObj := m.getPreDayFoundObj(resType)
	if preFoundObj == nil {
		return 0
	}

	return preFoundObj.resLevel
}

//获取资源波数
func (m *PlayerFoundDataManager) GetFoundGroup(resType foundtypes.FoundResourceType) int32 {
	preFoundObj := m.getPreDayFoundObj(resType)
	if preFoundObj == nil {
		return 0
	}

	return preFoundObj.group
}

func createPlayerFoundDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerFoundDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerFoundDataManagerType, player.PlayerDataManagerFactoryFunc(createPlayerFoundDataManager))
}
