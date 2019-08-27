package chuangshi

import (
	chuangshitypes "fgame/fgame/game/chuangshi/types"
)

type ChuangShiCityObject struct {
	campType   chuangshitypes.ChuangShiCampType
	cityType   chuangshitypes.ChuangShiCityType
	index      int32
	kingMember *ChuangShiMemberInfoObject
}

func (o *ChuangShiCityObject) GetCampType() chuangshitypes.ChuangShiCampType {
	return o.campType
}

func (o *ChuangShiCityObject) GetCityType() chuangshitypes.ChuangShiCityType {
	return o.cityType
}

func (o *ChuangShiCityObject) GetIndex() int32 {
	return o.index
}

func (o *ChuangShiCityObject) GetKingMember() *ChuangShiMemberInfoObject {
	return o.kingMember
}

func newChuangShiCityObject() *ChuangShiCityObject {
	o := &ChuangShiCityObject{}
	return o
}
