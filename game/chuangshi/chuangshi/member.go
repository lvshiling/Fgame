package chuangshi

import (
	chuangshitypes "fgame/fgame/game/chuangshi/types"
)

type ChuangShiMemberInfoObject struct {
	platform int32
	serverId int32
	name     string
	pos      chuangshitypes.ChuangShiGuanZhi
}

func (o *ChuangShiMemberInfoObject) GetPlatform() int32 {
	return o.platform
}

func (o *ChuangShiMemberInfoObject) GetServerId() int32 {
	return o.serverId
}

func (o *ChuangShiMemberInfoObject) GetName() string {
	return o.name
}

func (o *ChuangShiMemberInfoObject) GetPos() chuangshitypes.ChuangShiGuanZhi {
	return o.pos
}

func newChuangShiMemberInfoObject() *ChuangShiMemberInfoObject {
	o := &ChuangShiMemberInfoObject{}
	return o
}
