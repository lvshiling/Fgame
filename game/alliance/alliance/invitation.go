package alliance

import (
	"fgame/fgame/core/storage"
	allianceentity "fgame/fgame/game/alliance/entity"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

type AllianceInvitationObject struct {
	id           int64
	allianceId   int64
	invitationId int64
	level        int32
	name         string
	role         playertypes.RoleType
	sex          playertypes.SexType
	force        int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func createAllianceInvitationObject() *AllianceInvitationObject {
	o := &AllianceInvitationObject{}
	return o
}

func convertAllianceInvitationObjectToEntity(o *AllianceInvitationObject) (*allianceentity.AllianceInvitationEntity, error) {
	e := &allianceentity.AllianceInvitationEntity{
		Id:           o.id,
		AllianceId:   o.allianceId,
		InvitationId: o.invitationId,
		Name:         o.name,
		Role:         int32(o.role),
		Sex:          int32(o.sex),
		Force:        o.force,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *AllianceInvitationObject) GetId() int64 {
	return o.id
}

func (o *AllianceInvitationObject) GetDBId() int64 {
	return o.id
}

func (o *AllianceInvitationObject) GetAllianceId() int64 {
	return o.allianceId
}

func (o *AllianceInvitationObject) GetInvitationId() int64 {
	return o.invitationId
}

func (o *AllianceInvitationObject) GetLevel() int32 {
	return o.level
}

func (o *AllianceInvitationObject) GetRole() playertypes.RoleType {
	return o.role
}

func (o *AllianceInvitationObject) GetSex() playertypes.SexType {
	return o.sex
}

func (o *AllianceInvitationObject) GetName() string {
	return o.name
}

func (o *AllianceInvitationObject) GetForce() int64 {
	return o.force
}
func (o *AllianceInvitationObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *AllianceInvitationObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertAllianceInvitationObjectToEntity(o)
	return e, err
}

func (o *AllianceInvitationObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*allianceentity.AllianceInvitationEntity)
	o.id = ae.Id
	o.allianceId = ae.AllianceId
	o.invitationId = ae.InvitationId
	o.name = ae.Name
	o.sex = playertypes.SexType(ae.Sex)
	o.role = playertypes.RoleType(ae.Role)
	o.force = ae.Force
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *AllianceInvitationObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AllianceInvitation"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

func (o *AllianceInvitationObject) IsApplyCD() bool {
	if o.updateTime == 0 {
		return false
	}

	cd := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceJoinCoolingTime))
	now := global.GetGame().GetTimeService().Now()
	dif := now - o.updateTime

	if dif > cd {
		return false
	}
	return true
}
