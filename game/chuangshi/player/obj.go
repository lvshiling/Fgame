package player

import (
	"fgame/fgame/core/storage"
	chuangshientity "fgame/fgame/game/chuangshi/entity"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

type PlayerChuangShiYuGaoObject struct {
	id         int64
	player     player.Player
	isJoin     int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerChuangShiYuGaoObject(pl player.Player) *PlayerChuangShiYuGaoObject {
	o := &PlayerChuangShiYuGaoObject{
		player: pl,
	}
	return o
}

func convertPlayerChuangShiYuGaoObjectToEntity(o *PlayerChuangShiYuGaoObject) (*chuangshientity.PlayerChuangShiYuGaoEntity, error) {
	e := &chuangshientity.PlayerChuangShiYuGaoEntity{
		Id:         o.id,
		IsJoin:     o.isJoin,
		PlayerId:   o.player.GetId(),
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerChuangShiYuGaoObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerChuangShiYuGaoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerChuangShiYuGaoObjectToEntity(o)
	return e, err
}

func (o *PlayerChuangShiYuGaoObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.PlayerChuangShiYuGaoEntity)

	o.id = pse.Id
	o.isJoin = pse.IsJoin
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerChuangShiYuGaoObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ChuangShiYuGao"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)

	return
}

//玩家创世信息
type PlayerChuangShiObject struct {
	id            int64
	player        player.Player
	campType      chuangshitypes.ChuangShiCampType
	pos           chuangshitypes.ChuangShiGuanZhi
	jifen         int64
	diamonds      int64
	lastMyPayTime int64 //上次个人工资时间
	weiWang       int64
	joinCampTime  int64 //加入阵营时间
	updateTime    int64
	createTime    int64
	deleteTime    int64
}

func NewPlayerChuangShiObject(pl player.Player) *PlayerChuangShiObject {
	o := &PlayerChuangShiObject{
		player: pl,
	}
	return o
}

func convertPlayerChuangShiObjectToEntity(o *PlayerChuangShiObject) (*chuangshientity.PlayerChuangShiEntity, error) {
	e := &chuangshientity.PlayerChuangShiEntity{
		Id:            o.id,
		Pos:           int32(o.pos),
		CampType:      int32(o.campType),
		PlayerId:      o.player.GetId(),
		Jifen:         o.jifen,
		Diamonds:      o.diamonds,
		LastMyPayTime: o.lastMyPayTime,
		JoinCampTime:  o.joinCampTime,
		WeiWang:       o.weiWang,
		UpdateTime:    o.updateTime,
		CreateTime:    o.createTime,
		DeleteTime:    o.deleteTime,
	}
	return e, nil
}

func (o *PlayerChuangShiObject) IfEnoughJiFen(num int64) bool {
	if num < 0 {
		return false
	}
	if o.jifen >= num {
		return true
	}
	return false
}

func (o *PlayerChuangShiObject) GetLastMyPayTime() int64 {
	return o.lastMyPayTime
}

func (o *PlayerChuangShiObject) GetJifen() int64 {
	return o.jifen
}

func (o *PlayerChuangShiObject) GetDiamonds() int64 {
	return o.diamonds
}

func (o *PlayerChuangShiObject) GetWeiWang() int64 {
	return o.weiWang
}

func (o *PlayerChuangShiObject) GetCampType() chuangshitypes.ChuangShiCampType {
	return o.campType
}

func (o *PlayerChuangShiObject) SetCampType(campType chuangshitypes.ChuangShiCampType) {
	o.campType = campType
}

func (o *PlayerChuangShiObject) GetPos() chuangshitypes.ChuangShiGuanZhi {
	return o.pos
}

func (o *PlayerChuangShiObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerChuangShiObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerChuangShiObjectToEntity(o)
	return e, err
}

func (o *PlayerChuangShiObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.PlayerChuangShiEntity)

	o.id = pse.Id
	o.pos = chuangshitypes.ChuangShiGuanZhi(pse.Pos)
	o.campType = chuangshitypes.ChuangShiCampType(pse.CampType)
	o.jifen = pse.Jifen
	o.diamonds = pse.Diamonds
	o.lastMyPayTime = pse.LastMyPayTime
	o.joinCampTime = pse.JoinCampTime
	o.weiWang = pse.WeiWang
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerChuangShiObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerChuangShi"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerChuangShiObject) IfShenWang() bool {
	return o.pos == chuangshitypes.ChuangShiGuanZhiShenWang
}

func (o *PlayerChuangShiObject) IfChengZhu() bool {
	return o.pos == chuangshitypes.ChuangShiGuanZhiChengZhu
}

func (o *PlayerChuangShiObject) IfJoinCamp() bool {
	return o.pos != chuangshitypes.ChuangShiGuanZhiNone
}

//玩家报名信息
type PlayerChuangShiSignObject struct {
	id         int64
	player     player.Player
	status     chuangshitypes.ShenWangSignUpType
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerChuangShiSignObject(pl player.Player) *PlayerChuangShiSignObject {
	o := &PlayerChuangShiSignObject{
		player: pl,
	}
	return o
}

func convertPlayerChuangShiSignObjectToEntity(o *PlayerChuangShiSignObject) (*chuangshientity.PlayerChuangShiSignEntity, error) {
	e := &chuangshientity.PlayerChuangShiSignEntity{
		Id:         o.id,
		Status:     int32(o.status),
		PlayerId:   o.player.GetId(),
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerChuangShiSignObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerChuangShiSignObject) GetStatus() chuangshitypes.ShenWangSignUpType {
	return o.status
}

func (o *PlayerChuangShiSignObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerChuangShiSignObjectToEntity(o)
	return e, err
}

func (o *PlayerChuangShiSignObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.PlayerChuangShiSignEntity)

	o.id = pse.Id
	o.status = chuangshitypes.ShenWangSignUpType(pse.Status)
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerChuangShiSignObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerChuangShiSign"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerChuangShiSignObject) IfShenWangSignUp() bool {
	return o.status != chuangshitypes.ShenWangSignUpTypeNone
}

//玩家投票信息
type PlayerChuangShiVoteObject struct {
	id           int64
	player       player.Player
	status       chuangshitypes.ShenWangVoteType
	lastVoteTime int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewPlayerChuangShiVoteObject(pl player.Player) *PlayerChuangShiVoteObject {
	o := &PlayerChuangShiVoteObject{
		player: pl,
	}
	return o
}

func converPlayerChuangShiVoteObjectToEntity(o *PlayerChuangShiVoteObject) (*chuangshientity.PlayerChuangShiVoteEntity, error) {
	e := &chuangshientity.PlayerChuangShiVoteEntity{
		Id:           o.id,
		Status:       int32(o.status),
		PlayerId:     o.player.GetId(),
		LastVoteTime: o.lastVoteTime,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *PlayerChuangShiVoteObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerChuangShiVoteObject) GetStatus() chuangshitypes.ShenWangVoteType {
	return o.status
}

func (o *PlayerChuangShiVoteObject) ToEntity() (e storage.Entity, err error) {
	e, err = converPlayerChuangShiVoteObjectToEntity(o)
	return e, err
}

func (o *PlayerChuangShiVoteObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.PlayerChuangShiVoteEntity)

	o.id = pse.Id
	o.status = chuangshitypes.ShenWangVoteType(pse.Status)
	o.lastVoteTime = pse.LastVoteTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerChuangShiVoteObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerChuangShiVote"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerChuangShiVoteObject) IfVote() bool {
	return o.status != chuangshitypes.ShenWangVoteTypeNone
}

// 官职数据
type PlayerChuangShiGuanZhiObject struct {
	id              int64
	player          player.Player
	receiveRewLevel int32
	level           int32
	times           int32
	weiWang         int32
	updateTime      int64
	createTime      int64
	deleteTime      int64
}

func NewPlayerChuangShiGuanZhiObject(pl player.Player) *PlayerChuangShiGuanZhiObject {
	o := &PlayerChuangShiGuanZhiObject{
		player: pl,
	}
	return o
}

func converPlayerChuangShiGuanZhiObjectToEntity(o *PlayerChuangShiGuanZhiObject) (*chuangshientity.PlayerChuangShiGuanZhiEntity, error) {
	e := &chuangshientity.PlayerChuangShiGuanZhiEntity{
		Id:              o.id,
		PlayerId:        o.player.GetId(),
		ReceiveRewLevel: o.receiveRewLevel,
		Level:           o.level,
		Times:           o.times,
		WeiWang:         o.weiWang,
		UpdateTime:      o.updateTime,
		CreateTime:      o.createTime,
		DeleteTime:      o.deleteTime,
	}
	return e, nil
}

func (o *PlayerChuangShiGuanZhiObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerChuangShiGuanZhiObject) GetLevel() int32 {
	return o.level
}

func (o *PlayerChuangShiGuanZhiObject) GetRewLevel() int32 {
	return o.receiveRewLevel
}

func (o *PlayerChuangShiGuanZhiObject) GetTimes() int32 {
	return o.times
}

func (o *PlayerChuangShiGuanZhiObject) GetWeiWang() int32 {
	return o.weiWang
}

func (o *PlayerChuangShiGuanZhiObject) IsCanReceive(rewLevel int32) bool {
	return o.receiveRewLevel < rewLevel
}

func (o *PlayerChuangShiGuanZhiObject) ToEntity() (e storage.Entity, err error) {
	e, err = converPlayerChuangShiGuanZhiObjectToEntity(o)
	return e, err
}

func (o *PlayerChuangShiGuanZhiObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.PlayerChuangShiGuanZhiEntity)

	o.id = pse.Id
	o.receiveRewLevel = pse.ReceiveRewLevel
	o.level = pse.Level
	o.weiWang = pse.WeiWang
	o.times = pse.Times
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerChuangShiGuanZhiObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerChuangShiGuanZhi"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
