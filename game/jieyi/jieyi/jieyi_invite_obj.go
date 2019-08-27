package jieyi

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	jieyientity "fgame/fgame/game/jieyi/entity"
	jieyitypes "fgame/fgame/game/jieyi/types"

	"github.com/pkg/errors"
)

type JieYiInviteObject struct {
	id             int64
	serverId       int32
	state          jieyitypes.InviteState
	daoJu          jieyitypes.JieYiDaoJuType
	inviteDaoJu    jieyitypes.JieYiDaoJuType
	inviteToken    jieyitypes.JieYiTokenType
	inviteTokenLev int32
	nameLev        int32
	inviteId       int64
	inviteeId      int64
	name           string
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func NewJieYiInviteObject() *JieYiInviteObject {
	o := &JieYiInviteObject{}
	return o
}

func convertJieYiInviteObjectToEntity(o *JieYiInviteObject) (*jieyientity.JieYiInviteEntity, error) {
	e := &jieyientity.JieYiInviteEntity{
		Id:             o.id,
		ServerId:       o.serverId,
		State:          int32(o.state),
		DaoJu:          int32(o.daoJu),
		InviteDaoJu:    int32(o.inviteDaoJu),
		InviteToken:    int32(o.inviteToken),
		InviteTokenLev: o.inviteTokenLev,
		NameLev:        o.nameLev,
		InviteId:       o.inviteId,
		InviteeId:      o.inviteeId,
		Name:           o.name,
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return e, nil
}

func (o *JieYiInviteObject) GetDBId() int64 {
	return o.id
}

func (o *JieYiInviteObject) GetInviteId() int64 {
	return o.inviteId
}

func (o *JieYiInviteObject) GetInviteeId() int64 {
	return o.inviteeId
}

func (o *JieYiInviteObject) GetJieYiName() string {
	return o.name
}

func (o *JieYiInviteObject) GetJieYiDaoJu() jieyitypes.JieYiDaoJuType {
	return o.daoJu
}

func (o *JieYiInviteObject) GetInviteDaoJu() jieyitypes.JieYiDaoJuType {
	return o.inviteDaoJu
}

func (o *JieYiInviteObject) GetInviteState() jieyitypes.InviteState {
	return o.state
}

func (o *JieYiInviteObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertJieYiInviteObjectToEntity(o)
	return e, err
}

func (o *JieYiInviteObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*jieyientity.JieYiInviteEntity)

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.daoJu = jieyitypes.JieYiDaoJuType(pse.DaoJu)
	o.inviteDaoJu = jieyitypes.JieYiDaoJuType(pse.InviteDaoJu)
	o.inviteToken = jieyitypes.JieYiTokenType(pse.InviteToken)
	o.inviteTokenLev = pse.InviteTokenLev
	o.nameLev = pse.NameLev
	o.state = jieyitypes.InviteState(pse.State)
	o.name = pse.Name
	o.inviteId = pse.InviteId
	o.inviteeId = pse.InviteeId
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *JieYiInviteObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "JieYiInvite"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}

func (o *JieYiInviteObject) Fail(now int64) bool {
	if o.state != jieyitypes.InviteStateInit {
		return false
	}
	o.state = jieyitypes.InviteStateFail
	o.updateTime = now
	return true
}

func (o *JieYiInviteObject) Success(now int64) bool {
	if o.state != jieyitypes.InviteStateInit {
		return false
	}
	o.state = jieyitypes.InviteStateSuccess
	o.updateTime = now
	return true
}

func (o *JieYiInviteObject) IfCanSuccess() bool {
	if o.state != jieyitypes.InviteStateInit {
		return false
	}

	return true
}
