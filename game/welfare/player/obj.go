package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	welfareentity "fgame/fgame/game/welfare/entity"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	"github.com/pkg/errors"
)

//开服活动
type PlayerOpenActivityObject struct {
	player          player.Player
	id              int64
	groupId         int32
	activityType    welfaretypes.OpenActivityType
	activitySubType welfaretypes.OpenActivitySubType
	activityData    welfaretypes.OpenActivityData
	startTime       int64
	endTime         int64
	updateTime      int64
	createTime      int64
	deleteTime      int64
	isReddot        bool //是否红点提示
}

func newPlayerOpenActivityObject(pl player.Player) *PlayerOpenActivityObject {
	o := &PlayerOpenActivityObject{
		player: pl,
	}
	return o
}

func convertPlayerOpenActivityObjectToEntity(o *PlayerOpenActivityObject) (e *welfareentity.PlayerOpenActivityEntity, err error) {

	activityData, err := json.Marshal(o.activityData)
	if err != nil {
		return
	}

	e = &welfareentity.PlayerOpenActivityEntity{
		Id:              o.id,
		PlayerId:        o.player.GetId(),
		GroupId:         o.groupId,
		ActivityType:    int32(o.activityType),
		ActivitySubType: o.activitySubType.SubType(),
		ActivityData:    string(activityData),
		StartTime:       o.startTime,
		EndTime:         o.endTime,
		UpdateTime:      o.updateTime,
		CreateTime:      o.createTime,
		DeleteTime:      o.deleteTime,
	}

	return e, nil
}

func (o *PlayerOpenActivityObject) GetGroupId() int32 {
	return o.groupId
}

func (o *PlayerOpenActivityObject) GetActivityData() welfaretypes.OpenActivityData {
	return o.activityData
}

func (o *PlayerOpenActivityObject) GetUpdateTime() int64 {
	return o.updateTime
}

func (o *PlayerOpenActivityObject) GetStartTime() int64 {
	return o.startTime
}

func (o *PlayerOpenActivityObject) GetEndTime() int64 {
	return o.endTime
}

func (o *PlayerOpenActivityObject) GetPlayer() player.Player {
	return o.player
}

func (o *PlayerOpenActivityObject) GetActivityType() welfaretypes.OpenActivityType {
	return o.activityType
}

func (o *PlayerOpenActivityObject) GetActivitySubType() welfaretypes.OpenActivitySubType {
	return o.activitySubType
}

func (o *PlayerOpenActivityObject) GetIsReddot() bool {
	return o.isReddot
}

func (o *PlayerOpenActivityObject) SetIsReddot(isReddot bool) {
	o.isReddot = isReddot
}

func (o *PlayerOpenActivityObject) SetUpdateTime(time int64) {
	o.updateTime = time
}

func (o *PlayerOpenActivityObject) IsFeedbackCharge() bool {
	if o.activityType == welfaretypes.OpenActivityTypeFeedback && o.activitySubType == welfaretypes.OpenActivityFeedbackSubTypeCharge {
		return true
	}

	return false
}

func (o *PlayerOpenActivityObject) IsAdvanced() bool {
	if o.activityType == welfaretypes.OpenActivityTypeAdvanced && o.activitySubType == welfaretypes.OpenActivityAdvancedSubTypeFeedback {
		return true
	}

	return false
}

func (o *PlayerOpenActivityObject) IsSingleCharge() bool {
	if o.activityType == welfaretypes.OpenActivityTypeFeedback && o.activitySubType == welfaretypes.OpenActivityFeedbackSubTypeSingleChagre {
		return true
	}

	return false
}

func (o *PlayerOpenActivityObject) IsGoldBow() bool {
	if o.activityType == welfaretypes.OpenActivityTypeFeedback && o.activitySubType == welfaretypes.OpenActivityFeedbackSubTypeGoldBowl {
		return true
	}

	return false
}

func (o *PlayerOpenActivityObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerOpenActivityObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerOpenActivityObjectToEntity(o)
	return e, err
}

func (o *PlayerOpenActivityObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*welfareentity.PlayerOpenActivityEntity)
	typ := welfaretypes.OpenActivityType(te.ActivityType)
	subType := welfaretypes.CreateOpenActivitySubType(typ, te.ActivitySubType)

	data, err := CreateOpenActivityData(typ, subType, te.ActivityData)
	if err != nil {
		return err
	}
	if data == nil {
		panic(fmt.Errorf("welfare:加载活动数据错误，typ:%d,subType:%d", typ, subType))
	}

	o.id = te.Id
	o.groupId = te.GroupId
	o.activityType = typ
	o.activitySubType = subType
	o.activityData = data
	o.startTime = te.StartTime
	o.endTime = te.EndTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime

	return nil
}

func (o *PlayerOpenActivityObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Welfare"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)

	return
}

//活动充值
type PlayerOpenActivityChargeObject struct {
	player     player.Player
	id         int64
	groupId    int32
	goldNum    int32
	startTime  int64
	endTime    int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func newPlayerOpenActivityChargeObject(pl player.Player) *PlayerOpenActivityChargeObject {
	o := &PlayerOpenActivityChargeObject{
		player: pl,
	}
	return o
}

func convertPlayerOpenActivityChargeObjectToEntity(o *PlayerOpenActivityChargeObject) (e *welfareentity.PlayerOpenActivityChargeEntity, err error) {

	e = &welfareentity.PlayerOpenActivityChargeEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		GroupId:    o.groupId,
		GoldNum:    o.goldNum,
		StartTime:  o.startTime,
		EndTime:    o.endTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerOpenActivityChargeObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerOpenActivityChargeObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerOpenActivityChargeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerOpenActivityChargeObjectToEntity(o)
	return e, err
}

func (o *PlayerOpenActivityChargeObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*welfareentity.PlayerOpenActivityChargeEntity)

	o.id = te.Id
	o.groupId = te.GroupId
	o.goldNum = te.GoldNum
	o.startTime = te.StartTime
	o.endTime = te.EndTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerOpenActivityChargeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OpenActivityCharge"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//活动消费
type PlayerOpenActivityCostObject struct {
	player     player.Player
	id         int64
	groupId    int32
	goldNum    int64
	startTime  int64
	endTime    int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func newPlayerOpenActivityCostObject(pl player.Player) *PlayerOpenActivityCostObject {
	o := &PlayerOpenActivityCostObject{
		player: pl,
	}
	return o
}

func convertPlayerOpenActivityCostObjectToEntity(o *PlayerOpenActivityCostObject) (e *welfareentity.PlayerOpenActivityCostEntity, err error) {

	e = &welfareentity.PlayerOpenActivityCostEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		GroupId:    o.groupId,
		GoldNum:    o.goldNum,
		StartTime:  o.startTime,
		EndTime:    o.endTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerOpenActivityCostObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerOpenActivityCostObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerOpenActivityCostObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerOpenActivityCostObjectToEntity(o)
	return e, err
}

func (o *PlayerOpenActivityCostObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*welfareentity.PlayerOpenActivityCostEntity)

	o.id = te.Id
	o.groupId = te.GroupId
	o.goldNum = te.GoldNum
	o.startTime = te.StartTime
	o.endTime = te.EndTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerOpenActivityCostObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OpenActivityCost"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//活动次数数据记录
type PlayerActivityNumRecordObject struct {
	player     player.Player
	id         int64
	groupId    int32
	times      int32
	startTime  int64
	endTime    int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func newPlayerActivityNumRecordObject(pl player.Player) *PlayerActivityNumRecordObject {
	o := &PlayerActivityNumRecordObject{
		player: pl,
	}
	return o
}

func convertPlayerActivityNumRecordObjectToEntity(o *PlayerActivityNumRecordObject) (e *welfareentity.PlayerActivityNumRecordEntity, err error) {

	e = &welfareentity.PlayerActivityNumRecordEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		GroupId:    o.groupId,
		Times:      o.times,
		StartTime:  o.startTime,
		EndTime:    o.endTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerActivityNumRecordObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerActivityNumRecordObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerActivityNumRecordObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerActivityNumRecordObjectToEntity(o)
	return e, err
}

func (o *PlayerActivityNumRecordObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*welfareentity.PlayerActivityNumRecordEntity)

	o.id = te.Id
	o.groupId = te.GroupId
	o.times = te.Times
	o.startTime = te.StartTime
	o.endTime = te.EndTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerActivityNumRecordObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ActivityNumRecord"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//首充对象
type PlayerFirstChargeObject struct {
	player     player.Player
	id         int64
	isReceive  bool
	updateTime int64
	createTime int64
	deleteTime int64
}

func newPlayerFirstChargeObject(pl player.Player) *PlayerFirstChargeObject {
	o := &PlayerFirstChargeObject{
		player: pl,
	}
	return o
}

func convertPlayerFirstChargeObjectToEntity(o *PlayerFirstChargeObject) (e *welfareentity.PlayerFirstChargeEntity, err error) {

	e = &welfareentity.PlayerFirstChargeEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		IsReceive:  o.isReceive,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerFirstChargeObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerFirstChargeObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerFirstChargeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFirstChargeObjectToEntity(o)
	return e, err
}

func (o *PlayerFirstChargeObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*welfareentity.PlayerFirstChargeEntity)

	o.id = te.Id
	o.isReceive = te.IsReceive
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerFirstChargeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OpenActivityFirstCharge"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//活动开启邮件记录
type PlayerActivityOpenMailObject struct {
	player     player.Player
	id         int64
	group      int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func newPlayerActivityOpenMailObject(pl player.Player) *PlayerActivityOpenMailObject {
	o := &PlayerActivityOpenMailObject{
		player: pl,
	}
	return o
}

func convertPlayerActivityOpenMailObjectToEntity(o *PlayerActivityOpenMailObject) (e *welfareentity.PlayerActivityOpenMailEntity, err error) {

	e = &welfareentity.PlayerActivityOpenMailEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		Group:      o.group,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerActivityOpenMailObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerActivityOpenMailObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerActivityOpenMailObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerActivityOpenMailObjectToEntity(o)
	return e, err
}

func (o *PlayerActivityOpenMailObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*welfareentity.PlayerActivityOpenMailEntity)

	o.id = te.Id
	o.group = te.Group
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerActivityOpenMailObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OpenActivityOpenMail"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//活动次数数据记录
type PlayerActivityAddNumObject struct {
	player     player.Player
	id         int64
	groupId    int32
	addNum     int32
	startTime  int64
	endTime    int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func newPlayerActivityAddNumObject(pl player.Player) *PlayerActivityAddNumObject {
	o := &PlayerActivityAddNumObject{
		player: pl,
	}
	return o
}

func convertPlayerActivityAddNumObjectToEntity(o *PlayerActivityAddNumObject) (e *welfareentity.PlayerActivityAddNumEntity, err error) {

	e = &welfareentity.PlayerActivityAddNumEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		GroupId:    o.groupId,
		AddNum:     o.addNum,
		StartTime:  o.startTime,
		EndTime:    o.endTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerActivityAddNumObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerActivityAddNumObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerActivityAddNumObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerActivityAddNumObjectToEntity(o)
	return e, err
}

func (o *PlayerActivityAddNumObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*welfareentity.PlayerActivityAddNumEntity)

	o.id = te.Id
	o.groupId = te.GroupId
	o.addNum = te.AddNum
	o.startTime = te.StartTime
	o.endTime = te.EndTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerActivityAddNumObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ActivityAddNum"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
