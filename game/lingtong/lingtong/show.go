package lingtong

import (
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	"fgame/fgame/game/scene/scene"
)

type LingTongShowObject struct {
	fashionId   int32
	weaponId    int32
	weaponState int32
	titleId     int32
	wingId      int32
	mountId     int32
	mountHidden bool
	shenFaId    int32
	lingYuId    int32
	faBaoId     int32
	xianTiId    int32
}

func CreateLingTongShowObject(
	fashionId int32,
	weaponId int32,
	weaponState int32,
	titleId int32,
	wingId int32,
	mountId int32,
	mountHidden bool,
	shenFaId int32,
	lingYuId int32,
	faBaoId int32,
	xianTiId int32,
) *LingTongShowObject {
	obj := &LingTongShowObject{}
	obj.fashionId = fashionId
	obj.weaponId = weaponId
	obj.weaponState = weaponState
	obj.titleId = titleId
	obj.wingId = wingId
	obj.mountId = mountId
	obj.shenFaId = shenFaId
	obj.lingYuId = lingYuId
	obj.mountHidden = mountHidden
	obj.faBaoId = faBaoId
	obj.xianTiId = xianTiId
	return obj
}

type LingTongShowManager struct {
	p       scene.LingTong
	showObj *LingTongShowObject
}

func (m *LingTongShowManager) GetLingTongFashionId() int32 {
	return m.showObj.fashionId
}
func (m *LingTongShowManager) SetLingTongFashionId(fashionId int32) {
	m.showObj.fashionId = fashionId

	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongShowFashionChanged, m.p, nil)
}

func (m *LingTongShowManager) GetLingTongWeaponId() int32 {
	return m.showObj.weaponId
}

func (m *LingTongShowManager) SetLingTongWeapon(weaponId int32, weaponState int32) {
	m.showObj.weaponId = weaponId
	m.showObj.weaponState = weaponState
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongShowWeaponChanged, m.p, nil)

}

func (m *LingTongShowManager) GetLingTongWeaponState() int32 {
	return m.showObj.weaponState
}

func (m *LingTongShowManager) SetLingTongWeaponState(weaponState int32) {
	m.showObj.weaponState = weaponState
}

func (m *LingTongShowManager) GetLingTongTitleId() int32 {
	return m.showObj.titleId
}

func (m *LingTongShowManager) SetLingTongTitleId(titleId int32) {
	m.showObj.titleId = titleId
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongShowTitleChanged, m.p, nil)

}

func (m *LingTongShowManager) GetLingTongWingId() int32 {
	return m.showObj.wingId
}

func (m *LingTongShowManager) SetLingTongWingId(wingId int32) {
	m.showObj.wingId = wingId
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongShowWingChanged, m.p, nil)

}

func (m *LingTongShowManager) GetLingTongMountId() int32 {
	return m.showObj.mountId
}

func (m *LingTongShowManager) SetLingTongMountId(mountId int32) {
	m.showObj.mountId = mountId
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongShowMountChanged, m.p, nil)
}

func (m *LingTongShowManager) LingTongMountHidden(hidden bool) {
	m.showObj.mountHidden = hidden
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongShowMountHidden, m.p, nil)
}

func (m *LingTongShowManager) IsLingTongMountHidden() bool {
	return m.showObj.mountHidden
}

func (m *LingTongShowManager) GetLingTongShenFaId() int32 {
	return m.showObj.shenFaId
}

func (m *LingTongShowManager) SetLingTongShenFaId(shenFaId int32) {
	m.showObj.shenFaId = shenFaId
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongShowShenFaChanged, m.p, nil)

}

func (m *LingTongShowManager) GetLingTongLingYuId() int32 {
	return m.showObj.lingYuId
}

func (m *LingTongShowManager) SetLingTongLingYuId(lingYuId int32) {
	m.showObj.lingYuId = lingYuId
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongShowLingYuChanged, m.p, nil)

}

func (m *LingTongShowManager) SetLingTongFaBaoId(faBaoId int32) {
	m.showObj.faBaoId = faBaoId
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongShowFaBaoChanged, m.p, nil)
}

func (m *LingTongShowManager) GetLingTongFaBaoId() int32 {
	return m.showObj.faBaoId
}

func (m *LingTongShowManager) GetLingTongXianTiId() int32 {
	return m.showObj.xianTiId
}

func (m *LingTongShowManager) SetLingTongXianTiId(xianTiId int32) {
	m.showObj.xianTiId = xianTiId
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongShowXianTiChanged, m.p, nil)
}

func (m *LingTongShowManager) GetExtraSpeed() int64 {
	return 0
}

func CreateLingTongShowManagerWithObject(p scene.LingTong, showObj *LingTongShowObject) *LingTongShowManager {
	m := &LingTongShowManager{
		p:       p,
		showObj: showObj,
	}
	return m
}
