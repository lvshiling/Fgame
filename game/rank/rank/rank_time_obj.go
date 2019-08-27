package rank

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	rankentity "fgame/fgame/game/rank/entity"
	ranktypes "fgame/fgame/game/rank/types"
	"fgame/fgame/pkg/timeutils"
)

//神魔战场时间戳数据
type RankTimeObject struct {
	id            int64
	serverId      int32
	classRankType ranktypes.RankClassType
	rankType      ranktypes.RankType
	thisTime      int64
	updateTime    int64
	createTime    int64
	deleteTime    int64
}

func NewRankTimeObject() *RankTimeObject {
	po := &RankTimeObject{}
	return po
}

func (o *RankTimeObject) GetDBId() int64 {
	return o.id
}

func (o *RankTimeObject) GetThisTime() int64 {
	return o.thisTime
}

func (o *RankTimeObject) GetRankType() ranktypes.RankType {
	return o.rankType
}

func (o *RankTimeObject) GetClassRankType() ranktypes.RankClassType {
	return o.classRankType
}

func (o *RankTimeObject) ToEntity() (e storage.Entity, err error) {
	re := &rankentity.RankTimeEntity{}
	re.Id = o.id
	re.ServerId = o.serverId
	re.ClassRankType = int32(o.classRankType)
	re.RankType = int32(o.rankType)
	re.ThisTime = o.thisTime
	re.UpdateTime = o.updateTime
	re.CreateTime = o.createTime
	re.DeleteTime = o.deleteTime
	e = re
	return
}

func (o *RankTimeObject) FromEntity(e storage.Entity) (err error) {
	re, _ := e.(*rankentity.RankTimeEntity)
	o.id = re.Id
	o.serverId = re.ServerId
	o.thisTime = re.ThisTime
	o.classRankType = ranktypes.RankClassType(re.ClassRankType)
	o.rankType = ranktypes.RankType(re.RankType)
	o.updateTime = re.UpdateTime
	o.createTime = re.CreateTime
	o.deleteTime = re.DeleteTime
	return
}

func (o *RankTimeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

func (o *RankTimeObject) initRankTime() {
	now := global.GetGame().GetTimeService().Now()
	timeStamp, _ := timeutils.MondayFivePointTime(now)
	o.thisTime = timeStamp
	o.updateTime = now
	o.SetModified()
}

func (o *RankTimeObject) ifRefreshRankTime() bool {
	now := global.GetGame().GetTimeService().Now()
	diffTime := now - o.thisTime
	if diffTime < 0 {
		diffTime *= -1 //绝对值
	}
	if diffTime < int64(7*timeutils.DAY) {
		return false
	}

	return true
}
