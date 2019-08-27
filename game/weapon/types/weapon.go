package types

//兵魂标识类型
type WeaponTagType int32

const (
	//永久兵魂
	WeaponTagTypePermanent WeaponTagType = 1 + iota
	//其它途径获取或失去
	WeaponTagTypeTemp
)

func (wtt WeaponTagType) Valid() bool {
	switch wtt {
	case WeaponTagTypePermanent,
		WeaponTagTypeTemp:
		break
	default:
		return false
	}
	return true
}

//兵魂类型
type WeaponType int32

const (
	//普通兵魂
	WeaponTypeNormal WeaponType = 1 + iota
	//运营兵魂
	WeaponTypeYunYing
	//特殊兵魂
	WeaponTypeSpecial
	//定制兵魂
	WeaponTypeDingZhi
)

func (wt WeaponType) Valid() bool {
	switch wt {
	case WeaponTypeNormal,
		WeaponTypeYunYing,
		WeaponTypeSpecial,
		WeaponTypeDingZhi:
		break
	default:
		return false
	}
	return true
}

//觉醒
type WeaponAwakenType int32

const (
	//不可觉醒
	WeaponAwakenTypeNo WeaponAwakenType = iota
	//可觉醒
	WeaponAwakenTypeOk
)

func (wat WeaponAwakenType) Valid() bool {
	switch wat {
	case WeaponAwakenTypeNo,
		WeaponAwakenTypeOk:
		break
	default:
		return false
	}
	return true
}

//兵魂觉醒状态
type WeaponAwakenStatusType int32

const (
	//未觉醒
	WeaponAwakenStatusTypeNo = iota
	//已觉醒
	WeaponAwakenStatusTypeOk
)
