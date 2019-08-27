package hongbao

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	hongbaoentity "fgame/fgame/game/hongbao/entity"
	itemtypes "fgame/fgame/game/item/types"
	playertypes "fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//奖品信息
type AwardInfo struct {
	ItemId  int32 `json:"itemId"`
	ItemCnt int32 `json:"itemCnt"`
	Level   int32 `json:"level"`
}

//抢红包的玩家信息
type SnatcherInfo struct {
	PlayerId int64                `json:"playerId"`
	Name     string               `json:"name"`
	Role     playertypes.RoleType `json:"role"`
	Sex      playertypes.SexType  `json:"sex"`
	Level    int32                `json:"level"`
}

//红包对象
type HongBaoObject struct {
	id          int64
	serverId    int32
	hongBaoType itemtypes.ItemHongBaoSubType
	sendId      int64
	awardList   []*AwardInfo
	snatchLog   []*SnatcherInfo
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func NewHongBaoObject() *HongBaoObject {
	o := &HongBaoObject{}
	return o
}

func convertNewHongBaoObjectToEntity(o *HongBaoObject) (*hongbaoentity.HongBaoEntity, error) {
	awardBytes, err := json.Marshal(o.awardList)
	if err != nil {
		return nil, err
	}
	snatchBytes, err := json.Marshal(o.snatchLog)
	if err != nil {
		return nil, err
	}
	e := &hongbaoentity.HongBaoEntity{
		Id:          o.id,
		ServerId:    o.serverId,
		HongBaoType: int32(o.hongBaoType),
		SendId:      o.sendId,
		AwardList:   string(awardBytes),
		SnatchLog:   string(snatchBytes),
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}
	return e, nil
}

func (o *HongBaoObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *HongBaoObject) GetSendId() int64 {
	return o.sendId
}

func (o *HongBaoObject) GetHongBaoType() itemtypes.ItemHongBaoSubType {
	return o.hongBaoType
}

func (o *HongBaoObject) GetAwardList() []*AwardInfo {
	return o.awardList
}

func (o *HongBaoObject) GetSnatchLog() []*SnatcherInfo {
	return o.snatchLog
}

func (o *HongBaoObject) GetDBId() int64 {
	return o.id
}

func (o *HongBaoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewHongBaoObjectToEntity(o)
	return e, err
}

func (o *HongBaoObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*hongbaoentity.HongBaoEntity)

	var awardArr []*AwardInfo
	if err := json.Unmarshal([]byte(pse.AwardList), &awardArr); err != nil {
		return err
	}
	var snatchArr []*SnatcherInfo
	if err := json.Unmarshal([]byte(pse.SnatchLog), &snatchArr); err != nil {
		return err
	}

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.hongBaoType = itemtypes.ItemHongBaoSubType(pse.HongBaoType)
	o.sendId = pse.SendId
	o.awardList = awardArr
	o.snatchLog = snatchArr
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *HongBaoObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "HongBao"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}
