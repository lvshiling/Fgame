package marry

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	marryentity "fgame/fgame/game/marry/entity"
	marrytypes "fgame/fgame/game/marry/types"
)

//求婚婚戒数据
type MarryRingObject struct {
	Id           int64
	ServerId     int32
	PlayerId     int64
	PeerId       int64
	PeerName     string
	Ring         marrytypes.MarryRingType
	Status       marrytypes.MarryRingStatusType
	ProposalTime int64
	UpdateTime   int64
	CreateTime   int64
	DeleteTime   int64
}

func NewMarryRingObject() *MarryRingObject {
	pso := &MarryRingObject{}
	return pso
}

func (mro *MarryRingObject) GetDBId() int64 {
	return mro.Id
}

func (mro *MarryRingObject) ToEntity() (e storage.Entity, err error) {
	pre := &marryentity.MarryRingEntity{}
	pre.Id = mro.Id
	pre.ServerId = mro.ServerId
	pre.PlayerId = mro.PlayerId
	pre.PeerId = mro.PeerId
	pre.PeerName = mro.PeerName
	pre.Ring = int32(mro.Ring)
	pre.Status = int32(mro.Status)
	pre.ProposalTime = mro.ProposalTime
	pre.UpdateTime = mro.UpdateTime
	pre.CreateTime = mro.CreateTime
	pre.DeleteTime = mro.DeleteTime
	e = pre
	return
}

func (mro *MarryRingObject) FromEntity(e storage.Entity) (err error) {
	pre, _ := e.(*marryentity.MarryRingEntity)
	mro.Id = pre.Id
	mro.ServerId = pre.ServerId
	mro.PlayerId = pre.PlayerId
	mro.PeerId = pre.PeerId
	mro.PeerName = pre.PeerName
	mro.Ring = marrytypes.MarryRingType(pre.Ring)
	mro.Status = marrytypes.MarryRingStatusType(pre.Status)
	mro.ProposalTime = pre.ProposalTime
	mro.UpdateTime = pre.UpdateTime
	mro.CreateTime = pre.CreateTime
	mro.DeleteTime = pre.DeleteTime
	return
}

func (mro *MarryRingObject) SetModified() {
	e, err := mro.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
