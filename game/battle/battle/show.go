package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/mount/mount"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
)

type PlayerShowObject struct {
	fashionId      int32
	weaponId       int32
	weaponState    int32
	titleId        int32
	titleHidden    bool
	wingId         int32
	mountId        int32
	mountAdvanceId int32
	mountHidden    bool
	shenFaId       int32
	lingYuId       int32
	fourGodKey     int32
	realm          int32
	spouse         string
	spouseId       int64
	weddingStatus  int32
	model          int32
	ringType       int32
	ringLevel      int32
	faBaoId        int32
	petId          int32
	xianTiId       int32
	baGua          int32
	flyPetId       int32
	developLevel   int32
	shenYuKey      int32
}

func CreatePlayerShowObject(
	fashionId int32,
	weaponId int32,
	weaponState int32,
	titleId int32,
	wingId int32,
	mountId int32,
	mountAdvanceId int32,
	mountHidden bool,
	shenFaId int32,
	lingYuId int32,
	fourGodKey int32,
	realm int32,
	spouse string,
	spouseId int64,
	weddingStatus int32,
	ringType int32,
	ringLevel int32,
	faBaoId int32,
	petId int32,
	xianTiId int32,
	baGua int32,
	flyPetId int32,
	developLevel int32,
	shenYuKey int32,
) *PlayerShowObject {
	obj := &PlayerShowObject{}
	obj.fashionId = fashionId
	obj.weaponId = weaponId
	obj.weaponState = weaponState
	obj.titleId = titleId
	obj.titleHidden = false
	obj.wingId = wingId
	obj.mountAdvanceId = mountAdvanceId
	obj.mountId = mountId
	obj.shenFaId = shenFaId
	obj.lingYuId = lingYuId
	obj.fourGodKey = fourGodKey
	obj.realm = realm
	obj.spouse = spouse
	obj.spouseId = spouseId
	obj.weddingStatus = weddingStatus
	obj.model = 0
	obj.mountHidden = mountHidden
	obj.faBaoId = faBaoId
	obj.ringType = ringType
	obj.ringLevel = ringLevel
	obj.faBaoId = faBaoId
	obj.petId = petId
	obj.xianTiId = xianTiId
	obj.baGua = baGua
	obj.flyPetId = flyPetId
	obj.developLevel = developLevel
	obj.shenYuKey = shenYuKey
	return obj
}

type PlayerShowManager struct {
	p       scene.Player
	showObj *PlayerShowObject
}

func (m *PlayerShowManager) GetFashionId() int32 {
	return m.showObj.fashionId
}
func (m *PlayerShowManager) SetFashionId(fashionId int32) {
	m.showObj.fashionId = fashionId

	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowFashionChanged, m.p, nil)
}

func (m *PlayerShowManager) GetWeaponId() int32 {
	return m.showObj.weaponId
}

func (m *PlayerShowManager) SetWeapon(weaponId int32, weaponState int32) {
	m.showObj.weaponId = weaponId
	m.showObj.weaponState = weaponState
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowWeaponChanged, m.p, nil)

}

func (m *PlayerShowManager) GetWeaponState() int32 {
	return m.showObj.weaponState
}

func (m *PlayerShowManager) SetWeaponState(weaponState int32) {
	m.showObj.weaponState = weaponState
}

func (m *PlayerShowManager) GetTitleId() int32 {
	if m.showObj.titleHidden {
		return 0
	}
	return m.showObj.titleId
}

func (m *PlayerShowManager) SetTitleId(titleId int32) {
	if m.showObj.titleId == titleId {
		return
	}
	m.showObj.titleId = titleId
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowTitleChanged, m.p, nil)
}

func (m *PlayerShowManager) TitleHidden(hidden bool) {
	if m.showObj.titleHidden == hidden {
		return
	}
	m.showObj.titleHidden = hidden
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowTitleChanged, m.p, nil)
}

func (m *PlayerShowManager) GetWingId() int32 {
	return m.showObj.wingId
}

func (m *PlayerShowManager) SetWingId(wingId int32) {
	m.showObj.wingId = wingId
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowWingChanged, m.p, nil)

}

func (m *PlayerShowManager) GetMountId() int32 {
	return m.showObj.mountId
}

func (m *PlayerShowManager) GetMountAdvanceId() int32 {
	return m.showObj.mountAdvanceId
}

func (m *PlayerShowManager) SetMountId(mountId int32, advanceId int32) {
	m.showObj.mountId = mountId
	m.showObj.mountAdvanceId = advanceId
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowMountChanged, m.p, nil)
}

func (m *PlayerShowManager) MountHidden(hidden bool) {
	m.showObj.mountHidden = hidden
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowMountHidden, m.p, nil)
}

func (m *PlayerShowManager) IsMountHidden() bool {
	return m.showObj.mountHidden
}

func (m *PlayerShowManager) MountSync(hidden bool) {
	m.showObj.mountHidden = hidden
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowMountSync, m.p, nil)
}

func (m *PlayerShowManager) GetShenFaId() int32 {
	return m.showObj.shenFaId
}

func (m *PlayerShowManager) SetShenFaId(shenFaId int32) {
	m.showObj.shenFaId = shenFaId
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowShenFaChanged, m.p, nil)

}

func (m *PlayerShowManager) GetLingYuId() int32 {
	return m.showObj.lingYuId
}

func (m *PlayerShowManager) SetLingYuId(lingYuId int32) {
	m.showObj.lingYuId = lingYuId
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowLingYuChanged, m.p, nil)

}

func (m *PlayerShowManager) SetFourGodKey(fourGodKey int32) {
	m.showObj.fourGodKey = fourGodKey
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowFourGodKeyChanged, m.p, nil)
}

func (m *PlayerShowManager) GetFourGodKey() int32 {
	return m.showObj.fourGodKey
}

func (m *PlayerShowManager) SetRealm(realm int32) {
	m.showObj.realm = realm
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowRealmChanged, m.p, nil)
}

func (m *PlayerShowManager) GetRealm() int32 {
	return m.showObj.realm
}

func (m *PlayerShowManager) SetSpouse(name string) {
	m.showObj.spouse = name
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowSpouseChanged, m.p, nil)
}

func (m *PlayerShowManager) GetSpouse() string {
	return m.showObj.spouse
}

func (m *PlayerShowManager) SetSpouseId(spouseId int64) {
	m.showObj.spouseId = spouseId

}

func (m *PlayerShowManager) GetSpouseId() int64 {
	return m.showObj.spouseId
}

func (m *PlayerShowManager) SetWeddingStatus(status int32) {
	m.showObj.weddingStatus = status
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowWeddingStatusChanged, m.p, nil)
}

func (m *PlayerShowManager) GetWeddingStatus() int32 {
	return m.showObj.weddingStatus
}

func (m *PlayerShowManager) SetModel(model int32) {
	m.showObj.model = model
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowModelChanged, m.p, nil)
}

func (m *PlayerShowManager) GetModel() int32 {
	return m.showObj.model
}

func (m *PlayerShowManager) SetRingType(ringType int32) {
	m.showObj.ringType = ringType
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowRingTypeChanged, m.p, nil)
}

func (m *PlayerShowManager) GetRingType() int32 {
	return m.showObj.ringType
}

func (m *PlayerShowManager) SetRingLevel(ringLevel int32) {
	m.showObj.ringLevel = ringLevel
}

func (m *PlayerShowManager) GetRingLevel() int32 {
	return m.showObj.ringLevel
}

func (m *PlayerShowManager) SetFaBaoId(faBaoId int32) {
	m.showObj.faBaoId = faBaoId
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowFaBaoChanged, m.p, nil)
}

func (m *PlayerShowManager) GetFaBaoId() int32 {
	return m.showObj.faBaoId
}

func (m *PlayerShowManager) GetXianTiId() int32 {
	return m.showObj.xianTiId
}

func (m *PlayerShowManager) SetXianTiId(xianTiId int32) {
	m.showObj.xianTiId = xianTiId
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowXianTiChanged, m.p, nil)
}

func (m *PlayerShowManager) SetBaGua(level int32) {
	m.showObj.baGua = level
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowBaGuaChanged, m.p, nil)
}

func (m *PlayerShowManager) GetBaGua() int32 {
	return m.showObj.baGua
}

func (m *PlayerShowManager) SetPetId(petId int32) {
	m.showObj.petId = petId
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowPetChanged, m.p, nil)
}

func (m *PlayerShowManager) GetPetId() int32 {
	return m.showObj.petId
}

func (m *PlayerShowManager) SetMarryDevelopLevel(developLevel int32) {
	m.showObj.developLevel = developLevel
}

func (m *PlayerShowManager) GetMarryDevelopLevel() int32 {
	return m.showObj.developLevel
}

func (m *PlayerShowManager) SetFlyPetId(flyPetId int32) {
	m.showObj.flyPetId = flyPetId
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowFlyPetChanged, m.p, nil)
}

func (m *PlayerShowManager) GetFlyPetId() int32 {
	return m.showObj.flyPetId
}

func (m *PlayerShowManager) GetExtraSpeed() int64 {
	if m.showObj.mountHidden {
		return 0
	} else {
		mountAdvanceId := m.showObj.mountAdvanceId
		if mountAdvanceId != 0 {
			mountTemplate := mount.GetMountService().GetMountNumber(mountAdvanceId)
			moveSpeed := mountTemplate.GetBattleAttrTemplate().GetAllBattleProperty()[propertytypes.BattlePropertyTypeMoveSpeed]
			return moveSpeed
		}
		mountId := m.showObj.mountId
		if mountId != 0 {
			mountTemplate := mount.GetMountService().GetMount(int(mountId))
			moveSpeed := mountTemplate.GetBattleAttrTemplate().GetAllBattleProperty()[propertytypes.BattlePropertyTypeMoveSpeed]
			return moveSpeed
		}

		return 0

	}
}

func (m *PlayerShowManager) SetShenYuKey(keyNum int32) {
	m.showObj.shenYuKey = keyNum
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShowShenYuKeyChanged, m.p, nil)
}

func (m *PlayerShowManager) GetShenYuKey() int32 {
	return m.showObj.shenYuKey
}

func CreatePlayerShowManagerWithObject(p scene.Player, showObj *PlayerShowObject) *PlayerShowManager {
	m := &PlayerShowManager{
		p:       p,
		showObj: showObj,
	}
	return m
}
