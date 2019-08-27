package player

import (
	"fgame/fgame/core/storage"
	arenaentity "fgame/fgame/game/arena/entity"

	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//竞技场对象
type PlayerArenaObject struct {
	player          player.Player
	id              int64
	endTime         int64
	culRewardTime   int32 //废弃
	totalRewardTime int32
	reliveTime      int32
	jiFenCount      int32
	jiFenDay        int32
	arenaTime       int64
	winCount        int32
	failCount       int32
	dayWinCount     int32
	dayMaxWinCount  int32
	rankRewTime     int64
	updateTime      int64
	createTime      int64
	deleteTime      int64
}

func (o *PlayerArenaObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerArenaObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerArenaObject) FromEntity(e storage.Entity) error {
	te := e.(*arenaentity.PlayerArenaEntity)
	o.id = te.Id
	o.endTime = te.EndTime
	o.culRewardTime = te.CulRewardTime
	o.totalRewardTime = te.TotalRewardTime
	o.jiFenCount = te.JiFenCount
	o.jiFenDay = te.JiFenDay
	o.arenaTime = te.ArenaTime
	o.winCount = te.WinCount
	o.failCount = te.FailCount
	o.reliveTime = te.ReliveTime
	o.dayWinCount = te.DayMaxWinCount
	o.dayMaxWinCount = te.DayMaxWinCount
	o.rankRewTime = te.RankRewTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}
func (o *PlayerArenaObject) ToEntity() (e storage.Entity, err error) {

	e = &arenaentity.PlayerArenaEntity{
		Id:              o.id,
		PlayerId:        o.player.GetId(),
		EndTime:         o.endTime,
		CulRewardTime:   o.culRewardTime,
		TotalRewardTime: o.totalRewardTime,
		JiFenCount:      o.jiFenCount,
		JiFenDay:        o.jiFenDay,
		ArenaTime:       o.arenaTime,
		WinCount:        o.winCount,
		FailCount:       o.failCount,
		DayWinCount:     o.dayWinCount,
		DayMaxWinCount:  o.dayMaxWinCount,
		RankRewTime:     o.rankRewTime,
		ReliveTime:      o.reliveTime,
		UpdateTime:      o.updateTime,
		CreateTime:      o.createTime,
		DeleteTime:      o.deleteTime,
	}
	return e, nil
}

func (o *PlayerArenaObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Arena"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)

}

func (o *PlayerArenaObject) GetCulRewardTime() int32 {
	return o.culRewardTime
}

func (o *PlayerArenaObject) GetReliveTime() int32 {
	return o.reliveTime
}

func (o *PlayerArenaObject) GetEndTime() int64 {
	return o.endTime
}

func (o *PlayerArenaObject) GetTotalRewardTime() int32 {
	return o.totalRewardTime
}

func (o *PlayerArenaObject) GetJiFenDay() int32 {
	return o.jiFenDay
}

func (o *PlayerArenaObject) GetJiFenCount() int32 {
	return o.jiFenCount
}

func (o *PlayerArenaObject) GetWinCount() int32 {
	return o.winCount
}

func (o *PlayerArenaObject) GetFailCount() int32 {
	return o.failCount
}

func (o *PlayerArenaObject) GetDayMaxWinCount() int32 {
	return o.dayMaxWinCount
}

func (o *PlayerArenaObject) GetDayWinCount() int32 {
	return o.dayWinCount
}

func (o *PlayerArenaObject) GetRankRewTime() int64 {
	return o.rankRewTime
}

func (o *PlayerArenaObject) IfEnoughPoint(num int32) bool {
	if num < 0 {
		return false
	}

	if o.jiFenCount >= num {
		return true
	}

	return false
}
func (o *PlayerArenaObject) IfExtralWinRew() bool {
	nextWinCount := o.dayWinCount + 1
	if o.dayMaxWinCount >= nextWinCount {
		return false
	}

	return true
}

func CreatePlayerArenaObject(pl player.Player) *PlayerArenaObject {
	o := &PlayerArenaObject{
		player: pl,
	}
	return o
}
