package types

type WeaponInfo struct {
	WeaponId int32 `json:"weaponId"`
	Level    int32 `json:"level"`
	CulLevel int32 `json:"culLevel"`
	CulPro   int32 `jsom:"culPro"`
	State    int32 `json:"state"`
}

type AllWeaponInfo struct {
	Wear       int32         `json:"wear"`
	WeaponList []*WeaponInfo `json:"weaponList"`
}
