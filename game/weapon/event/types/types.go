package types

type WeaponEventType string

const (
	//武器改变事件
	EventTypeWeaponChanged WeaponEventType = "WeaponChanged"
	//兵魂激活
	EventTypeWeaponActivate WeaponEventType = "WeaponActivate"
	//兵魂觉醒
	EventTypeWeaponAwaken WeaponEventType = "WeaponAwaken"
	//兵魂失效
	EventTypeWeaponRemove WeaponEventType = "WeaponRemove"
)
