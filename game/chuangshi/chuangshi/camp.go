package chuangshi

import (
	chuangshitypes "fgame/fgame/game/chuangshi/types"
)

type ChuangShiCampObject struct {
	campType   chuangshitypes.ChuangShiCampType
	kingMember *ChuangShiMemberInfoObject
	power      int64
}

func (o *ChuangShiCampObject) GetCampType() chuangshitypes.ChuangShiCampType {
	return o.campType
}

func (o *ChuangShiCampObject) GetKingMember() *ChuangShiMemberInfoObject {
	return o.kingMember
}

func (o *ChuangShiCampObject) GetPower() int64 {
	return o.power
}

func newChuangShiCampObject() *ChuangShiCampObject {
	o := &ChuangShiCampObject{}
	return o
}
