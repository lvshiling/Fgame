package types

import (
	itemtypes "fgame/fgame/game/item/types"
)

type WushuangWeaponPart int32

const (
	WushuangWeaponPartWeapon   WushuangWeaponPart = iota + 1 //武器
	WushuangWeaponPartClothes                                //衣服
	WushuangWeaponPartHead                                   //头盔
	WushuangWeaponPartShoes                                  //鞋子
	WushuangWeaponPartNecklace                               //项链
	WushuangWeaponPartPendant                                //挂饰
)

func (t WushuangWeaponPart) Valid() bool {
	switch t {
	case
		WushuangWeaponPartWeapon,
		WushuangWeaponPartClothes,
		WushuangWeaponPartHead,
		WushuangWeaponPartShoes,
		WushuangWeaponPartNecklace,
		WushuangWeaponPartPendant:
		return true
	default:
		return false
	}
}

var (
	wushuangWeaponPartMap = map[WushuangWeaponPart]string{
		WushuangWeaponPartWeapon:   "武器",
		WushuangWeaponPartClothes:  "衣服",
		WushuangWeaponPartHead:     "头盔",
		WushuangWeaponPartShoes:    "鞋子",
		WushuangWeaponPartNecklace: "项链",
		WushuangWeaponPartPendant:  "挂饰",
	}
)

func (t WushuangWeaponPart) String() string {
	return wushuangWeaponPartMap[t]
}

var (
	transformItemSubTypeToBodyPos = map[itemtypes.ItemWushuangWeaponSubType]WushuangWeaponPart{
		itemtypes.ItemWushuangWeaponSubTypeWeapon:   WushuangWeaponPartWeapon,
		itemtypes.ItemWushuangWeaponSubTypeCloths:   WushuangWeaponPartClothes,
		itemtypes.ItemWushuangWeaponSubTypeHead:     WushuangWeaponPartHead,
		itemtypes.ItemWushuangWeaponSubTypeShoes:    WushuangWeaponPartShoes,
		itemtypes.ItemWushuangWeaponSubTypeNecklace: WushuangWeaponPartNecklace,
		itemtypes.ItemWushuangWeaponSubTypePendant:  WushuangWeaponPartPendant,
	}
)

func GetBodyPosFromItemSubType(subtype itemtypes.ItemWushuangWeaponSubType) WushuangWeaponPart {
	return transformItemSubTypeToBodyPos[subtype]
}

func (part WushuangWeaponPart) GetItemSubType() itemtypes.ItemWushuangWeaponSubType {
	return itemtypes.ItemWushuangWeaponSubType(int32(part) - 1)
}

var (
	MinBodyType = WushuangWeaponPartWeapon
	MaxBodyType = WushuangWeaponPartPendant
)

type BodyPosWearType int32

const (
	// 兵魂
	BodyPosWearTypeWeapon BodyPosWearType = iota + 1
	// 时装
	BodyPosWearTypeCloths
)

func (t BodyPosWearType) Valid() bool {
	switch t {
	case BodyPosWearTypeWeapon,
		BodyPosWearTypeCloths:
		return true
	default:
		return false
	}
}
