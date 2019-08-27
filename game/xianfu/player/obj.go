package player

import (
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	xianfuentity "fgame/fgame/game/xianfu/entity"
	xianfueventtypes "fgame/fgame/game/xianfu/event/types"
	xianfutemplate "fgame/fgame/game/xianfu/template"
	xianfutypes "fgame/fgame/game/xianfu/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//仙府对象
type PlayerXianfuObject struct {
	player     player.Player
	id         int64
	xianfuId   int32
	xianfuType xianfutypes.XianfuType
	useTimes   int32
	startTime  int64
	state      xianfutypes.XianfuState
	group      int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func CreateNewPlayerXianfuObject(pl player.Player) *PlayerXianfuObject {
	newObj := &PlayerXianfuObject{
		player: pl,
	}
	return newObj
}

//数据库id
func (o *PlayerXianfuObject) GetDBId() int64 {
	return o.id
}

//对象转换为数据库实体
func (o *PlayerXianfuObject) ToEntity() (e storage.Entity, err error) {
	e = &xianfuentity.PlayerXianFuEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		XianfuId:   o.xianfuId,
		XianfuType: int32(o.xianfuType),
		UseTimes:   o.useTimes,
		StartTime:  o.startTime,
		State:      int32(o.state),
		Group:      o.group,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

//数据库实体转对象
func (o *PlayerXianfuObject) FromEntity(e storage.Entity) (err error) {
	pxfe, _ := e.(*xianfuentity.PlayerXianFuEntity)
	o.id = pxfe.Id
	o.xianfuId = pxfe.XianfuId
	o.xianfuType = xianfutypes.XianfuType(pxfe.XianfuType)
	o.useTimes = pxfe.UseTimes
	o.startTime = pxfe.StartTime
	o.state = xianfutypes.XianfuState(pxfe.State)
	o.group = pxfe.Group
	o.updateTime = pxfe.UpdateTime
	o.createTime = pxfe.CreateTime
	o.deleteTime = pxfe.DeleteTime
	return nil
}

//提交修改
func (o *PlayerXianfuObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(fmt.Errorf("xianfu: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//更新完成仙府升级
func (o *PlayerXianfuObject) upgradeDone(nowTime int64) (err error) {
	xfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(o.xianfuId, o.xianfuType)
	// beforeLevel := o.xianfuId
	o.xianfuId = xfTemplate.GetNextId()
	o.state = xianfutypes.XianfuStateWaitedToUpgrade
	o.startTime = 0
	o.updateTime = nowTime
	o.SetModified()

	gameevent.Emit(xianfueventtypes.EventTypeXianFuUpgradeSuccess, o.player, o.xianfuType)
	//TODO 临时删除引起登陆问题
	// upReason := commonlog.XianFuLogReasonUpgrade
	// upReasonText := fmt.Sprintf(upReason.String(), o.xianfuType)
	// uplevel := xfTemplate.GetNextId() - beforeLevel
	// logData := xianfueventtypes.CreatePlayerXianFuLogEventData(beforeLevel, uplevel, o.xianfuType, upReason, upReasonText)
	// gameevent.Emit(xianfueventtypes.EventTypeXianFuLog, o.player, logData)
	return
}

//是否完成升级
func (o *PlayerXianfuObject) isUpgradeDone() bool {
	if o.state == xianfutypes.XianfuStateUpgrading {
		now := global.GetGame().GetTimeService().Now()
		nextXfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(o.xianfuId+1, o.xianfuType)

		needTime := nextXfTemplate.GetUpgradeTime()
		costTime := now - o.startTime
		return costTime >= needTime
	}
	return false
}

//刷新xianfuObject挑战次数
func (o *PlayerXianfuObject) refreshUseTimes(nowTime int64) error {
	isSame, err := timeutils.IsSameFive(o.updateTime, nowTime)
	if err != nil {
		return err
	}

	if !isSame {
		o.useTimes = 0
		o.updateTime = nowTime
		o.SetModified()
	}
	return nil
}

func (o *PlayerXianfuObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerXianfuObject) GetXianfuId() int32 {
	return o.xianfuId
}
func (o *PlayerXianfuObject) GetXianfuType() xianfutypes.XianfuType {
	return o.xianfuType
}
func (o *PlayerXianfuObject) GetUseTimes() int32 {
	return o.useTimes
}
func (o *PlayerXianfuObject) GetState() xianfutypes.XianfuState {
	return o.state
}
func (o *PlayerXianfuObject) GetStartTime() int64 {
	return o.startTime
}

func (o *PlayerXianfuObject) SetUseTimes(num int32) {
	o.useTimes = num
}

func (o *PlayerXianfuObject) GetGroup() int32 {
	return o.group
}
