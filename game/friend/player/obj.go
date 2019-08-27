package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	friendentity "fgame/fgame/game/friend/entity"
	friendtypes "fgame/fgame/game/friend/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//黑名单对象
type PlayerFriendBlackObject struct {
	player     player.Player
	Id         int64
	FriendId   int64
	Black      int32
	RevBlack   int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func newPlayerFriendBlackObject(pl player.Player) *PlayerFriendBlackObject {
	obj := &PlayerFriendBlackObject{
		player: pl,
	}
	return obj
}

func convertPlayerFriendBlackObjectToEntity(obj *PlayerFriendBlackObject) (e *friendentity.PlayerFriendBlackEntity, err error) {

	e = &friendentity.PlayerFriendBlackEntity{
		Id:         obj.Id,
		PlayerId:   obj.player.GetId(),
		FriendId:   obj.FriendId,
		Black:      obj.Black,
		RevBlack:   obj.RevBlack,
		UpdateTime: obj.UpdateTime,
		CreateTime: obj.CreateTime,
		DeleteTime: obj.DeleteTime,
	}
	return e, nil
}

func (obj *PlayerFriendBlackObject) GetPlayerId() int64 {
	return obj.player.GetId()
}

func (obj *PlayerFriendBlackObject) GetDBId() int64 {
	return obj.Id
}

func (obj *PlayerFriendBlackObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFriendBlackObjectToEntity(obj)
	return e, err
}

func (obj *PlayerFriendBlackObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*friendentity.PlayerFriendBlackEntity)
	obj.Id = pe.Id
	obj.FriendId = pe.FriendId
	obj.Black = pe.Black
	obj.RevBlack = pe.RevBlack
	obj.UpdateTime = pe.UpdateTime
	obj.CreateTime = pe.CreateTime
	obj.DeleteTime = pe.DeleteTime
	return nil
}

func (obj *PlayerFriendBlackObject) SetModified() {
	e, err := obj.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Friend_Black"))
	}
	pe, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	obj.player.AddChangedObject(pe)
	return
}

//玩家接受到的添加好友的邀请
type PlayerFriendInviteObject struct {
	player     player.Player
	Id         int64
	InviteId   int64
	Name       string
	Role       int32
	Sex        int32
	Force      int64
	Level      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func newPlayerFriendInviteObject(pl player.Player) *PlayerFriendInviteObject {
	obj := &PlayerFriendInviteObject{
		player: pl,
	}
	return obj
}

func convertPlayerFriendInviteObjectToEntity(obj *PlayerFriendInviteObject) (e *friendentity.PlayerFriendInviteEntity, err error) {

	e = &friendentity.PlayerFriendInviteEntity{
		Id:         obj.Id,
		PlayerId:   obj.player.GetId(),
		InviteId:   obj.InviteId,
		Name:       obj.Name,
		Role:       obj.Role,
		Sex:        obj.Sex,
		Force:      obj.Force,
		Level:      obj.Level,
		UpdateTime: obj.UpdateTime,
		CreateTime: obj.CreateTime,
		DeleteTime: obj.DeleteTime,
	}
	return e, nil
}

func (obj *PlayerFriendInviteObject) GetPlayerId() int64 {
	return obj.player.GetId()
}

func (obj *PlayerFriendInviteObject) GetDBId() int64 {
	return obj.Id
}

func (obj *PlayerFriendInviteObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFriendInviteObjectToEntity(obj)
	return e, err
}

func (obj *PlayerFriendInviteObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*friendentity.PlayerFriendInviteEntity)
	obj.Id = pe.Id
	obj.InviteId = pe.InviteId
	obj.Name = pe.Name
	obj.Role = pe.Role
	obj.Sex = pe.Sex
	obj.Force = pe.Force
	obj.Level = pe.Level
	obj.UpdateTime = pe.UpdateTime
	obj.CreateTime = pe.CreateTime
	obj.DeleteTime = pe.DeleteTime
	return nil
}

func (obj *PlayerFriendInviteObject) SetModified() {
	e, err := obj.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Friend_Invite"))
	}
	pe, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	obj.player.AddChangedObject(pe)
	return
}

//玩家赞赏数据
type PlayerFriendFeedbackObject struct {
	player       player.Player
	id           int64
	friendId     int64
	friendName   string
	noticeType   friendtypes.FriendNoticeType
	feedbackType friendtypes.FriendFeedbackType
	condition    int32
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func newPlayerFriendFeedbackObject(pl player.Player) *PlayerFriendFeedbackObject {
	obj := &PlayerFriendFeedbackObject{
		player: pl,
	}
	return obj
}

func convertPlayerFriendFeedbackObjectToEntity(obj *PlayerFriendFeedbackObject) (e *friendentity.PlayerFriendFeedbackEntity, err error) {

	e = &friendentity.PlayerFriendFeedbackEntity{
		Id:           obj.id,
		PlayerId:     obj.player.GetId(),
		FriendId:     obj.friendId,
		FriendName:   obj.friendName,
		NoticeType:   int32(obj.noticeType),
		FeedbackType: int32(obj.feedbackType),
		Condition:    obj.condition,
		UpdateTime:   obj.updateTime,
		CreateTime:   obj.createTime,
		DeleteTime:   obj.deleteTime,
	}
	return e, nil
}

func (obj *PlayerFriendFeedbackObject) GetPlayerId() int64 {
	return obj.player.GetId()
}

func (obj *PlayerFriendFeedbackObject) GetFriendName() string {
	return obj.friendName
}

func (obj *PlayerFriendFeedbackObject) GetFriendId() int64 {
	return obj.friendId
}

func (obj *PlayerFriendFeedbackObject) GetNoticeType() friendtypes.FriendNoticeType {
	return obj.noticeType
}

func (obj *PlayerFriendFeedbackObject) GetFeedbackType() friendtypes.FriendFeedbackType {
	return obj.feedbackType
}

func (obj *PlayerFriendFeedbackObject) GetCondition() int32 {
	return obj.condition
}

func (obj *PlayerFriendFeedbackObject) GetDBId() int64 {
	return obj.id
}

func (obj *PlayerFriendFeedbackObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFriendFeedbackObjectToEntity(obj)
	return e, err
}

func (obj *PlayerFriendFeedbackObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*friendentity.PlayerFriendFeedbackEntity)
	obj.id = pe.Id
	obj.friendId = pe.FriendId
	obj.friendName = pe.FriendName
	obj.noticeType = friendtypes.FriendNoticeType(pe.NoticeType)
	obj.feedbackType = friendtypes.FriendFeedbackType(pe.FeedbackType)
	obj.condition = pe.Condition
	obj.updateTime = pe.UpdateTime
	obj.createTime = pe.CreateTime
	obj.deleteTime = pe.DeleteTime
	return nil
}

func (obj *PlayerFriendFeedbackObject) SetModified() {
	e, err := obj.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FriendFeedback"))
	}
	pe, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	obj.player.AddChangedObject(pe)
	return
}

//玩家添加好友数据
type PlayerFriendAddRewObject struct {
	player               player.Player
	id                   int64
	frDummyNum           int32
	lastAddDummyTime     int64
	rewRecord            map[int32]int32
	congratulateTimes    int32
	lastCongratulateTime int64

	updateTime int64
	createTime int64
	deleteTime int64
}

func newPlayerFriendAddRewObject(pl player.Player) *PlayerFriendAddRewObject {
	obj := &PlayerFriendAddRewObject{
		player: pl,
	}
	return obj
}

func convertPlayerFriendAddRewObjectToEntity(obj *PlayerFriendAddRewObject) (e *friendentity.PlayerFriendAddRewEntity, err error) {

	data, err := json.Marshal(obj.rewRecord)
	if err != nil {
		return
	}

	e = &friendentity.PlayerFriendAddRewEntity{
		Id:                   obj.id,
		PlayerId:             obj.player.GetId(),
		FrDummyNum:           obj.frDummyNum,
		LastAddDummyTime:     obj.lastAddDummyTime,
		RewRecord:            string(data),
		CongratulateTimes:    obj.congratulateTimes,
		LastCongratulateTime: obj.lastCongratulateTime,
		UpdateTime:           obj.updateTime,
		CreateTime:           obj.createTime,
		DeleteTime:           obj.deleteTime,
	}
	return e, nil
}

func (obj *PlayerFriendAddRewObject) GetPlayerId() int64 {
	return obj.player.GetId()
}

func (obj *PlayerFriendAddRewObject) GetDBId() int64 {
	return obj.id
}

func (obj *PlayerFriendAddRewObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFriendAddRewObjectToEntity(obj)
	return e, err
}

func (obj *PlayerFriendAddRewObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*friendentity.PlayerFriendAddRewEntity)

	recordMap := make(map[int32]int32)
	err = json.Unmarshal([]byte(pe.RewRecord), &recordMap)
	if err != nil {
		return
	}

	obj.id = pe.Id
	obj.frDummyNum = pe.FrDummyNum
	obj.rewRecord = recordMap
	obj.lastAddDummyTime = pe.LastAddDummyTime
	obj.congratulateTimes = pe.CongratulateTimes
	obj.lastCongratulateTime = pe.LastCongratulateTime
	obj.updateTime = pe.UpdateTime
	obj.createTime = pe.CreateTime
	obj.deleteTime = pe.DeleteTime
	return nil
}

func (obj *PlayerFriendAddRewObject) SetModified() {
	e, err := obj.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FriendAddRew"))
	}
	pe, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	obj.player.AddChangedObject(pe)
	return
}

//玩家赞赏记录数据
type PlayerFriendAdmireObject struct {
	player      player.Player
	id          int64
	friId       int64
	admireTimes int32
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func newPlayerFriendAdmireObject(pl player.Player) *PlayerFriendAdmireObject {
	obj := &PlayerFriendAdmireObject{
		player: pl,
	}
	return obj
}

func convertPlayerFriendAdmireObjectToEntity(obj *PlayerFriendAdmireObject) (e *friendentity.PlayerFriendAdmireEntity, err error) {
	e = &friendentity.PlayerFriendAdmireEntity{
		Id:          obj.id,
		PlayerId:    obj.player.GetId(),
		FriId:       obj.friId,
		AdmireTimes: obj.admireTimes,
		UpdateTime:  obj.updateTime,
		CreateTime:  obj.createTime,
		DeleteTime:  obj.deleteTime,
	}
	return e, nil
}

func (obj *PlayerFriendAdmireObject) GetPlayerId() int64 {
	return obj.player.GetId()
}

func (obj *PlayerFriendAdmireObject) GetDBId() int64 {
	return obj.id
}

func (obj *PlayerFriendAdmireObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFriendAdmireObjectToEntity(obj)
	return e, err
}

func (obj *PlayerFriendAdmireObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*friendentity.PlayerFriendAdmireEntity)

	obj.id = pe.Id
	obj.friId = pe.FriId
	obj.admireTimes = pe.AdmireTimes
	obj.updateTime = pe.UpdateTime
	obj.createTime = pe.CreateTime
	obj.deleteTime = pe.DeleteTime
	return nil
}

func (obj *PlayerFriendAdmireObject) SetModified() {
	e, err := obj.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FriendAdmire"))
	}
	pe, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	obj.player.AddChangedObject(pe)
	return
}
