package player

import (
	buffcommon "fgame/fgame/game/buff/common"
	"fgame/fgame/game/buff/dao"
	bufftemplate "fgame/fgame/game/buff/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/idutil"
)

//玩家buff管理器
type PlayerBuffDataManager struct {
	p                player.Player
	playerBuffObject *PlayerBuffObject
}

func (m *PlayerBuffDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerBuffDataManager) Load() (err error) {
	e, err := dao.GetBuffDao().GetBuff(m.p.GetId())
	if err != nil {
		return
	}
	if e == nil {
		m.initPlayerBuffObject()

	} else {
		m.playerBuffObject = newPlayerBuffObject(m.p)
		m.playerBuffObject.FromEntity(e)
	}

	return nil
}

func (m *PlayerBuffDataManager) initPlayerBuffObject() {
	playerBuffObject := newPlayerBuffObject(m.p)
	id, _ := idutil.GetId()

	now := global.GetGame().GetTimeService().Now()
	playerBuffObject.id = id
	playerBuffObject.buffMap = make(map[int32]*buffObject)
	playerBuffObject.createTime = now
	playerBuffObject.SetModified()

	m.playerBuffObject = playerBuffObject
	return
}

//加载后
func (m *PlayerBuffDataManager) AfterLoad() (err error) {
	//刷新buff
	m.refreshBuff()

	return nil
}

//刷新buff
func (m *PlayerBuffDataManager) refreshBuff() {
	now := global.GetGame().GetTimeService().Now()
	dirty := false
	for _, b := range m.playerBuffObject.GetBuffMap() {
		buffId := b.GetBuffId()
		buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
		if buffTemplate == nil {
			continue
		}

		if buffTemplate.GetOfflineSaveType() == scenetypes.BuffOfflineSaveTypeTimeStop {
			useTime := m.p.GetLastLogoutTime() - b.StartTime
			b.UseTime += useTime
			b.StartTime = now
			dirty = true
		}

		//判断是否到期了
		if b.IsExpired() {
			dirty = true
			//移除buff
			m.playerBuffObject.RemoveBuff(buffTemplate.Group)
			continue
		}
	}
	if dirty {
		m.playerBuffObject.SetModified()
	}
	return
}

//移除buff
func (m *PlayerBuffDataManager) RemoveBuff(buffId int32) {
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	if buffTemplate == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerBuffObject.RemoveBuff(buffTemplate.Group)
	m.playerBuffObject.updateTime = now
	m.playerBuffObject.SetModified()
	return
}

func (m *PlayerBuffDataManager) UpdateBuff(b buffcommon.BuffObject) {
	buffId := b.GetBuffId()
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	if buffTemplate == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerBuffObject.RemoveBuff(buffTemplate.Group)

	bo := copyFromBuffObject(b)
	m.playerBuffObject.AddBuff(bo)

	m.playerBuffObject.updateTime = now
	m.playerBuffObject.SetModified()

	return
}

//添加buff
func (m *PlayerBuffDataManager) AddBuff(b buffcommon.BuffObject) {
	buffId := b.GetBuffId()
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	if buffTemplate == nil {
		return
	}
	if !buffTemplate.IsSave() {
		return
	}
	previousB := m.playerBuffObject.GetBuff(buffTemplate.Group)
	if previousB != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	bo := copyFromBuffObject(b)

	m.playerBuffObject.AddBuff(bo)
	m.playerBuffObject.updateTime = now
	m.playerBuffObject.SetModified()

	return
}

//获取buff
func (m *PlayerBuffDataManager) GetBuffs() map[int32]buffcommon.BuffObject {
	buffs := make(map[int32]buffcommon.BuffObject)
	for _, b := range m.playerBuffObject.GetBuffMap() {
		buffs[b.GroupId] = b
	}
	return buffs
}

//心跳
func (m *PlayerBuffDataManager) Heartbeat() {

}

func CreatePlayerBuffDataManager(p player.Player) player.PlayerDataManager {
	pddm := &PlayerBuffDataManager{}
	pddm.p = p

	return pddm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerBuffDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerBuffDataManager))
}
