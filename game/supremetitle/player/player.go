package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/supremetitle/dao"
	supremetitleeventtypes "fgame/fgame/game/supremetitle/event/types"
	supremetitletemplate "fgame/fgame/game/supremetitle/template"
)

//玩家称号管理器
type PlayerSupremeTitleDataManager struct {
	p player.Player
	//玩家至尊称号列表
	playerTitleObjectMap map[int32]*PlayerSupremeTitleObject
	//玩家穿戴至尊称号
	playerWearTitleObject *PlayerWearSupremeTitleObject
}

func (m *PlayerSupremeTitleDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerSupremeTitleDataManager) Load() (err error) {
	err = m.loadTitle()
	if err != nil {
		return
	}
	err = m.loadTitleWear()
	if err != nil {
		return
	}
	return nil
}

func (m *PlayerSupremeTitleDataManager) loadTitle() (err error) {
	m.playerTitleObjectMap = make(map[int32]*PlayerSupremeTitleObject)
	//加载玩家至尊称号信息
	supremeTitleList, err := dao.GetSupremeTitleDao().GetSupremeTitleList(m.p.GetId())
	if err != nil {
		return
	}
	//至尊称号信息
	for _, supremeTitle := range supremeTitleList {
		pto := NewPlayerSupremeTitleObject(m.p)
		pto.FromEntity(supremeTitle)
		m.playerTitleObjectMap[pto.GetTitleId()] = pto
	}
	return
}

func (m *PlayerSupremeTitleDataManager) loadTitleWear() (err error) {
	//加载玩家穿戴至尊称号信息
	supremeTitleWearEntity, err := dao.GetSupremeTitleDao().GetSupremeTitleWearEntity(m.p.GetId())
	if err != nil {
		return
	}
	if supremeTitleWearEntity == nil {
		m.initPlayerWearSupremeTitleObject()
	} else {
		m.playerWearTitleObject = NewPlayerWearSupremeTitleObject(m.p)
		m.playerWearTitleObject.FromEntity(supremeTitleWearEntity)
	}
	return
}

//第一次初始化
func (m *PlayerSupremeTitleDataManager) initPlayerWearSupremeTitleObject() {
	pwto := NewPlayerWearSupremeTitleObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwto.id = id
	//生成id
	pwto.titleWear = int32(0)
	pwto.createTime = now
	m.playerWearTitleObject = pwto
	pwto.SetModified()
}

//加载后
func (m *PlayerSupremeTitleDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (m *PlayerSupremeTitleDataManager) Heartbeat() {
}

func (m *PlayerSupremeTitleDataManager) GetTitleMap() map[int32]*PlayerSupremeTitleObject {
	return m.playerTitleObjectMap
}

//称号穿戴信息
func (m *PlayerSupremeTitleDataManager) GetTitleWear() *PlayerWearSupremeTitleObject {
	return m.playerWearTitleObject
}

//称号穿戴信息
func (m *PlayerSupremeTitleDataManager) GetTitleId() int32 {
	return m.playerWearTitleObject.GetTitleWear()
}

//是否已拥有该称号
func (m *PlayerSupremeTitleDataManager) IfTitleExist(titleId int32) bool {
	_, exist := m.playerTitleObjectMap[titleId]
	if exist {
		return true
	}
	return false
}

//是否已穿戴
func (m *PlayerSupremeTitleDataManager) HasedWeared(titleId int32) bool {
	return m.playerWearTitleObject.titleWear == titleId
}

//称号激活
func (m *PlayerSupremeTitleDataManager) TitleActive(titleId int32) (flag bool) {
	titleTemplate := supremetitletemplate.GetTitleDingZhiTemplateService().GetTitleDingZhiTempalte(titleId)
	if titleTemplate == nil {
		return
	}

	flag = m.IfTitleExist(titleId)
	if flag {
		return false
	}

	id, err := idutil.GetId()
	if err != nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	pto := NewPlayerSupremeTitleObject(m.p)
	pto.id = id
	pto.titleId = titleId
	pto.createTime = now
	pto.SetModified()
	m.playerTitleObjectMap[titleId] = pto
	gameevent.Emit(supremetitleeventtypes.EventTypeSupremeTitleActivate, m.p, titleId)
	return true
}

//称号穿戴
func (m *PlayerSupremeTitleDataManager) TitleWear(titleId int32) bool {
	flag := m.IfTitleExist(titleId)
	if !flag {
		return false
	}

	flag = m.HasedWeared(titleId)
	if flag {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerWearTitleObject.titleWear = titleId
	m.playerWearTitleObject.updateTime = now
	m.playerWearTitleObject.SetModified()
	//发送事件
	gameevent.Emit(supremetitleeventtypes.EventTypeSupremeTitleChanged, m.p, nil)
	return true
}

//称号卸下
func (m *PlayerSupremeTitleDataManager) TitleNoWear() {
	titleWear := m.GetTitleId()
	if titleWear == 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerWearTitleObject.titleWear = 0
	m.playerWearTitleObject.updateTime = now
	m.playerWearTitleObject.SetModified()

	//发送事件
	gameevent.Emit(supremetitleeventtypes.EventTypeSupremeTitleChanged, m.p, nil)
	return
}

func CreatePlayerSupremeTitleDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerSupremeTitleDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerSupremeTitleDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerSupremeTitleDataManager))
}
