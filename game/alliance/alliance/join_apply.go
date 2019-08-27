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

type AllianceJoinApplyObject struct {
	id         int64
	allianceId int64
	joinId     int64
	level      int32
	name       string
	role       playertypes.RoleType
	sex        playertypes.SexType
	force      int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func createAllianceJoinApplyObject() *AllianceJoinApplyObject {
	o := &AllianceJoinApplyObject{}
	return o
}

func convertAllianceJoinApplyObjectToEntity(o *AllianceJoinApplyObject) (*allianceentity.AllianceJoinApplyEntity, error) {
	e := &allianceentity.AllianceJoinApplyEntity{
		Id:         o.id,
		AllianceId: o.allianceId,
		JoinId:     o.joinId,
		Name:       o.name,
		Role:       int32(o.role),
		Sex:        int32(o.sex),
		Force:      o.force,
		Level:      o.level,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *AllianceJoinApplyObject) GetId() int64 {
	return o.id
}

func (o *AllianceJoinApplyObject) GetDBId() int64 {
	return o.id
}

func (o *AllianceJoinApplyObject) GetAllianceId() int64 {
	return o.allianceId
}

func (o *AllianceJoinApplyObject) GetJoinId() int64 {
	return o.joinId
}

func (o *AllianceJoinApplyObject) GetLevel() int32 {
	return o.level
}

func (o *AllianceJoinApplyObject) GetRole() playertypes.RoleType {
	return o.role
}

func (o *AllianceJoinApplyObject) GetSex() playertypes.SexType {
	return o.sex
}

func (o *AllianceJoinApplyObject) GetName() string {
	return o.name
}

func (o *AllianceJoinApplyObject) GetForce() int64 {
	return o.force
}

func (o *AllianceJoinApplyObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *AllianceJoinApplyObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertAllianceJoinApplyObjectToEntity(o)
	return e, err
}

func (o *AllianceJoinApplyObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*allianceentity.AllianceJoinApplyEntity)
	o.id = ae.Id
	o.allianceId = ae.AllianceId
	o.joinId = ae.JoinId
	o.name = ae.Name
	o.sex = playertypes.SexType(ae.Sex)
	o.role = playertypes.RoleType(ae.Role)
	o.force = ae.Force
	o.level = ae.Level
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *AllianceJoinApplyObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AllianceJoinApply"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

func (o *AllianceJoinApplyObject) isApplyCD() bool {
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
