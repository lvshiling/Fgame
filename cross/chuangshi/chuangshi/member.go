package chuangshi

import (
	"fgame/fgame/core/storage"
	chuangshientity "fgame/fgame/cross/chuangshi/entity"
	alliancetypes "fgame/fgame/game/alliance/types"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//
type ChuangShiMemberObject struct {
	camp             *Camp
	id               int64
	platform         int32
	serverId         int32
	campType         chuangshitypes.ChuangShiCampType
	playerPlatform   int32
	playerServerId   int32
	playerId         int64
	playerName       string
	allianceId       int64
	allianceName     string
	force            int64
	scheduleJifen    int64                           //分配的积分
	scheduleDiamonds int64                           //分配的钻石
	pos              chuangshitypes.ChuangShiGuanZhi //官职
	alPos            alliancetypes.AlliancePosition  //仙盟职位
	updateTime       int64
	createTime       int64
	deleteTime       int64
}

func newChuangShiMemberObject(camp *Camp) *ChuangShiMemberObject {
	o := &ChuangShiMemberObject{}
	o.camp = camp
	return o
}

func convertChuangShiMemberObjectToEntity(o *ChuangShiMemberObject) (*chuangshientity.ChuangShiMemberEntity, error) {
	e := &chuangshientity.ChuangShiMemberEntity{
		Id:             o.id,
		Platform:       o.platform,
		ServerId:       o.serverId,
		PlayerPlatform: o.playerPlatform,
		PlayerServerId: o.playerServerId,
		PlayerId:       o.playerId,
		PlayerName:     o.playerName,
		AllianceId:     o.allianceId,
		AllianceName:   o.allianceName,
		Pos:            int32(o.pos),
		CampType:       int32(o.campType),
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return e, nil
}

func (o *ChuangShiMemberObject) GetCamp() *Camp {
	return o.camp
}

func (o *ChuangShiMemberObject) GetId() int64 {
	return o.id
}

func (o *ChuangShiMemberObject) GetDBId() int64 {
	return o.id
}

func (o *ChuangShiMemberObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertChuangShiMemberObjectToEntity(o)
	return e, err
}

func (o *ChuangShiMemberObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.ChuangShiMemberEntity)
	o.id = pse.Id
	o.platform = pse.Platform
	o.serverId = pse.ServerId
	o.playerPlatform = pse.PlayerPlatform
	o.playerServerId = pse.PlayerServerId
	o.playerId = pse.PlayerId
	o.playerName = pse.PlayerName
	o.allianceId = pse.AllianceId
	o.allianceName = pse.AllianceName
	o.pos = chuangshitypes.ChuangShiGuanZhi(pse.Pos)
	o.campType = chuangshitypes.ChuangShiCampType(pse.CampType)
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *ChuangShiMemberObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ChuangShiMember"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}

func (o *ChuangShiMemberObject) GetPlatform() int32 {
	return o.platform
}

func (o *ChuangShiMemberObject) GetServerId() int32 {
	return o.serverId
}

func (o *ChuangShiMemberObject) GetPlayerPlatform() int32 {
	return o.playerPlatform
}

func (o *ChuangShiMemberObject) GetPlayerServerId() int32 {
	return o.playerServerId
}

func (o *ChuangShiMemberObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *ChuangShiMemberObject) GetPlayerName() string {
	return o.playerName
}

func (o *ChuangShiMemberObject) GetAllianceId() int64 {
	return o.allianceId
}

func (o *ChuangShiMemberObject) GetAlPos() alliancetypes.AlliancePosition {
	return o.alPos
}

func (o *ChuangShiMemberObject) GetForce() int64 {
	return o.force
}

func (o *ChuangShiMemberObject) GetCampType() chuangshitypes.ChuangShiCampType {
	return o.camp.campObj.campType
}

func (o *ChuangShiMemberObject) GetAllianceName() string {
	return o.allianceName
}

func (o *ChuangShiMemberObject) GetPos() chuangshitypes.ChuangShiGuanZhi {
	return o.pos
}

func (o *ChuangShiMemberObject) IfShenWang() bool {
	return o.pos == chuangshitypes.ChuangShiGuanZhiShenWang
}
