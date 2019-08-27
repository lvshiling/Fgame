package alliance

import (
	"fgame/fgame/game/common/common"

	"fgame/fgame/core/storage"
	allianceentity "fgame/fgame/game/alliance/entity"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

type AllianceMemberObject struct {
	id             int64
	alliance       *Alliance
	memberId       int64
	lingyuId       int32
	role           playertypes.RoleType
	sex            playertypes.SexType
	level          int32
	name           string
	force          int64
	zhuanSheng     int32
	vip            int32
	gongXian       int64
	position       alliancetypes.AlliancePosition
	joinTime       int64
	onlineStatus   playertypes.PlayerOnlineState
	lastLogoutTime int64
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func createAllianceMemberObject(al *Alliance) *AllianceMemberObject {
	o := &AllianceMemberObject{
		alliance: al,
	}
	return o
}

func convertAllianceMemberObjectToEntity(o *AllianceMemberObject) (*allianceentity.AllianceMemberEntity, error) {
	e := &allianceentity.AllianceMemberEntity{
		Id:             o.id,
		AllianceId:     o.GetAllianceId(),
		MemberId:       o.memberId,
		LingyuId:       o.lingyuId,
		Position:       int32(o.position),
		Role:           int32(o.role),
		Sex:            int32(o.sex),
		JoinTime:       o.joinTime,
		Level:          o.level,
		Name:           o.name,
		Force:          o.force,
		Vip:            o.vip,
		ZhuanSheng:     o.zhuanSheng,
		GongXian:       o.gongXian,
		LastLogoutTime: o.lastLogoutTime,
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return e, nil
}

func (o *AllianceMemberObject) GetId() int64 {
	return o.id
}

func (o *AllianceMemberObject) GetDBId() int64 {
	return o.id
}

func (o *AllianceMemberObject) GetAllianceId() int64 {
	return o.alliance.GetAllianceId()
}

func (o *AllianceMemberObject) GetAlliance() *Alliance {
	return o.alliance
}

func (o *AllianceMemberObject) GetMemberId() int64 {
	return o.memberId
}

func (o *AllianceMemberObject) GetPosition() alliancetypes.AlliancePosition {
	return o.position
}

func (o *AllianceMemberObject) GetVip() int32 {
	return o.vip
}

func (o *AllianceMemberObject) GetLevel() int32 {
	return o.level
}

func (o *AllianceMemberObject) GetSex() playertypes.SexType {
	return o.sex
}
func (o *AllianceMemberObject) GetRole() playertypes.RoleType {
	return o.role
}

func (o *AllianceMemberObject) GetName() string {
	return o.name
}

func (o *AllianceMemberObject) GetForce() int64 {
	return o.force
}
func (o *AllianceMemberObject) GetZhuanSheng() int32 {
	return o.zhuanSheng
}

func (o *AllianceMemberObject) GetLingyuId() int32 {
	return o.lingyuId
}

func (o *AllianceMemberObject) GetGongXian() int64 {
	return o.gongXian
}

func (o *AllianceMemberObject) GetJoinTime() int64 {
	return o.joinTime
}

func (o *AllianceMemberObject) GetOnlineStatus() playertypes.PlayerOnlineState {
	return o.onlineStatus
}

func (o *AllianceMemberObject) GetLastLogoutTime() int64 {
	return o.lastLogoutTime
}

func (o *AllianceMemberObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertAllianceMemberObjectToEntity(o)
	return e, err
}

func (o *AllianceMemberObject) FromEntity(e storage.Entity) error {
	ame, _ := e.(*allianceentity.AllianceMemberEntity)
	o.id = ame.Id
	o.memberId = ame.MemberId
	o.lingyuId = ame.LingyuId
	o.role = playertypes.RoleType(ame.Role)
	o.sex = playertypes.SexType(ame.Sex)
	o.position = alliancetypes.AlliancePosition(ame.Position)
	o.joinTime = ame.JoinTime
	o.name = ame.Name
	o.level = ame.Level
	o.force = ame.Force
	o.vip = ame.Vip
	o.zhuanSheng = ame.ZhuanSheng
	o.gongXian = ame.GongXian
	o.joinTime = ame.JoinTime
	o.lastLogoutTime = ame.LastLogoutTime
	o.onlineStatus = playertypes.PlayerOnlineStateOffline
	now := global.GetGame().GetTimeService().Now()
	if o.lastLogoutTime == 0 {
		o.lastLogoutTime = now
	}
	o.updateTime = ame.UpdateTime
	o.createTime = ame.CreateTime
	o.deleteTime = ame.DeleteTime
	return nil
}

func (o *AllianceMemberObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AllianceMember"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

func (o *AllianceMemberObject) Priority(other *AllianceMemberObject) bool {
	return o.position.Priority(other.position)
}

const (
	impeachTime = int64(common.DAY * 7)
)

func (o *AllianceMemberObject) IsOfflineOneWeek() bool {
	now := global.GetGame().GetTimeService().Now()
	if o.onlineStatus == playertypes.PlayerOnlineStateOnline {
		return false
	}
	if o.lastLogoutTime <= 0 {
		return false
	}
	offLineTime := now - o.lastLogoutTime
	return offLineTime > impeachTime
}

func (o *AllianceMemberObject) IsPositionMember() bool {
	return o.position == alliancetypes.AlliancePositionMember
}

func (o *AllianceMemberObject) IsMengZhu() bool {
	return o.position == alliancetypes.AlliancePositionMengZhu
}

func (o *AllianceMemberObject) IsFuMengZhu() bool {
	return o.position == alliancetypes.AlliancePositionFuMengZhu
}
