package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/wushuangweapon/dao"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"
)

type PlayerWushuangWeaponDataManager struct {
	p                                 player.Player
	playerWushuangWeaponSlotObjectMap map[wushuangweapontypes.WushuangWeaponPart]*PlayerWushuangWeaponSlotObject
	wushuangSettingsMap               map[int32]*PlayerWushuangSettingsObject
	buchangObj                        *PlayerWushuangBuchangObject
}

func (m *PlayerWushuangWeaponDataManager) Player() player.Player {
	return m.p
}

func (m *PlayerWushuangWeaponDataManager) GetWushuangSettings(itemId int32) *PlayerWushuangSettingsObject {
	obj, ok := m.wushuangSettingsMap[itemId]
	if !ok {
		return nil
	}
	return obj
}

func (m *PlayerWushuangWeaponDataManager) SaveWushuangSettings(itemId int32, level int32) {
	settingsObj := m.GetWushuangSettings(itemId)
	now := global.GetGame().GetTimeService().Now()
	if settingsObj != nil {
		settingsObj.level = level
		settingsObj.updateTime = now
		settingsObj.SetModified()
	} else {
		obj := createPlayerWushuangSettingsObject(m.p, itemId, level)
		obj.updateTime = now
		obj.SetModified()
		//记得加载到内存！
		m.wushuangSettingsMap[itemId] = obj
	}
}

func (m *PlayerWushuangWeaponDataManager) GetSlotObjectFromBodyPos(bodyPos wushuangweapontypes.WushuangWeaponPart) (obj *PlayerWushuangWeaponSlotObject) {
	obj, ok := m.playerWushuangWeaponSlotObjectMap[bodyPos]
	if !ok {
		return nil
	}
	return
}

func (m *PlayerWushuangWeaponDataManager) GetSlotObjectMap() map[wushuangweapontypes.WushuangWeaponPart]*PlayerWushuangWeaponSlotObject {
	return m.playerWushuangWeaponSlotObjectMap
}

//仅gm 修改部位经验
func (m *PlayerWushuangWeaponDataManager) GmSetSlotExperience(bodyPos wushuangweapontypes.WushuangWeaponPart, experience int64) {
	obj := m.GetSlotObjectFromBodyPos(bodyPos)
	if obj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()

	obj.experience = experience
	obj.updateLevel(0)
	obj.updateTime = now
	obj.SetModified()
}

func (m *PlayerWushuangWeaponDataManager) ToWushuangListInfo() []*wushuangweapontypes.WushuangInfo {
	wushuangList := make([]*wushuangweapontypes.WushuangInfo, 0, wushuangweapontypes.MaxBodyType)
	for i := wushuangweapontypes.MinBodyType; i <= wushuangweapontypes.MaxBodyType; i++ {
		obj := m.GetSlotObjectFromBodyPos(i)
		info := &wushuangweapontypes.WushuangInfo{}
		bp := int32(i)
		info.BodyPos = bp
		info.Level = obj.GetLevel()
		info.ItemId = obj.GetItemId()
		info.Exp = obj.GetExperience()
		wushuangList = append(wushuangList, info)
	}
	return wushuangList
}

func (m *PlayerWushuangWeaponDataManager) IsSendBuchangEmail() bool {
	obj := m.buchangObj
	if obj == nil {
		return false
	}
	if obj.isSendEmail == int32(1) {
		return true
	}
	return false
}

func (m *PlayerWushuangWeaponDataManager) SendBuchangEmail() {
	obj := m.buchangObj
	if obj == nil {
		obj = createPlayerWushuangBuchangObject(m.p)
		m.buchangObj = obj
	}
	now := global.GetGame().GetTimeService().Now()
	obj.isSendEmail = int32(1)
	obj.updateTime = now
	obj.SetModified()
}

//加载
func (m *PlayerWushuangWeaponDataManager) Load() (err error) {
	wushuangWeaponSlotList, err := dao.GetWushuangWeaponDao().GetAllWushuangWeaponSlotEntity(m.p.GetId())
	if err != nil {
		return err
	}
	wushuangSettingsList, err := dao.GetWushuangWeaponDao().GetAllWushuangSettings(m.p.GetId())
	if err != nil {
		return err
	}
	wushuangBuchangEntity, err := dao.GetWushuangWeaponDao().GetAllWushuangBuchang(m.p.GetId())
	if err != nil {
		return err
	}

	if wushuangBuchangEntity != nil {
		newObj := NewPlayerWushuangBuchangObject(m.p)
		err = newObj.FromEntity(wushuangBuchangEntity)
		if err != nil {
			return
		}
		m.buchangObj = newObj
	}

	m.playerWushuangWeaponSlotObjectMap = make(map[wushuangweapontypes.WushuangWeaponPart]*PlayerWushuangWeaponSlotObject)
	m.wushuangSettingsMap = make(map[int32]*PlayerWushuangSettingsObject)

	// 加载已有数据
	for _, e := range wushuangWeaponSlotList {
		pwwso := NewPlayerWushuangWeaponSlotObject(m.p)
		err = pwwso.FromEntity(e)
		if err != nil {
			return
		}
		m.playerWushuangWeaponSlotObjectMap[pwwso.GetBodyPart()] = pwwso
	}

	// 初始化未有数据
	for i := wushuangweapontypes.MinBodyType; i <= wushuangweapontypes.MaxBodyType; i++ {
		_, ok := m.playerWushuangWeaponSlotObjectMap[i]
		if !ok {
			newObj := NewPlayerWushuangWeaponSlotObject(m.p)
			initNewPlayerWushuangWeaponSlotObject(newObj, i)
			m.playerWushuangWeaponSlotObjectMap[i] = newObj
		}
	}

	// 加载已有设置
	for _, e := range wushuangSettingsList {
		obj := NewPlayerWushuangSettingsObject(m.p)
		err = obj.FromEntity(e)
		if err != nil {
			return
		}
		m.wushuangSettingsMap[obj.itemId] = obj
	}

	return nil
}

//加载后
func (m *PlayerWushuangWeaponDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerWushuangWeaponDataManager) Heartbeat() {
}

func createPlayerWushuangWeaponDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerWushuangWeaponDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerWushuangWeaponDataManagerType, player.PlayerDataManagerFactoryFunc(createPlayerWushuangWeaponDataManager))
}
