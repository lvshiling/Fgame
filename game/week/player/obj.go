package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	weekentity "fgame/fgame/game/week/entity"
	weektypes "fgame/fgame/game/week/types"
	"fgame/fgame/pkg/timeutils"

	"github.com/pkg/errors"
)

//周卡
type PlayerWeekObject struct {
	player      player.Player
	id          int64
	updateTime  int64
	createTime  int64
	deleteTime  int64
	weekDataMap map[weektypes.WeekType]*WeekData
}

func newPlayerWeekObject(pl player.Player) *PlayerWeekObject {
	o := &PlayerWeekObject{
		player: pl,
	}
	return o
}

func convertPlayerWeekObjectToEntity(o *PlayerWeekObject) (e *weekentity.PlayerWeekEntity, err error) {
	e = &weekentity.PlayerWeekEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}

	for weekType, data := range o.weekDataMap {
		switch weekType {
		case weektypes.WeekTypeJunior:
			{
				e.JuniorCycDay = data.cycDay
				e.JuniorExpireTime = data.expireTime
				e.JuniorLastDayRewTime = data.lastDayRewTime
			}
		case weektypes.WeekTypeSenior:
			{
				e.SeniorCycDay = data.cycDay
				e.SeniorExpireTime = data.expireTime
				e.SeniorLastDayRewTime = data.lastDayRewTime
			}
		}
	}

	return e, nil
}

func (o *PlayerWeekObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerWeekObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerWeekObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerWeekObjectToEntity(o)
	return e, err
}

func (o *PlayerWeekObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*weekentity.PlayerWeekEntity)

	o.id = te.Id
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime

	o.weekDataMap = make(map[weektypes.WeekType]*WeekData)
	juniorData := newWeekData(weektypes.WeekTypeJunior, te.JuniorCycDay, te.JuniorExpireTime, te.JuniorLastDayRewTime)
	o.weekDataMap[juniorData.weekType] = juniorData

	seniorData := newWeekData(weektypes.WeekTypeSenior, te.SeniorCycDay, te.SeniorExpireTime, te.SeniorLastDayRewTime)
	o.weekDataMap[seniorData.weekType] = seniorData
	return nil
}

func (o *PlayerWeekObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Week"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

// 封装周卡数据
type WeekData struct {
	weekType       weektypes.WeekType
	cycDay         int32
	expireTime     int64
	lastDayRewTime int64
}

func newWeekData(weekType weektypes.WeekType, cycDay int32, expireTime, lastDayRewTime int64) *WeekData {
	d := &WeekData{
		weekType:       weekType,
		cycDay:         cycDay,
		expireTime:     expireTime,
		lastDayRewTime: lastDayRewTime,
	}
	return d
}

func initWeekData(weekType weektypes.WeekType) *WeekData {
	return newWeekData(weekType, 0, 0, 0)
}

func (o *WeekData) IsWeek(now int64) bool {
	return now < o.expireTime
}

func (o *WeekData) IsReceiveRewards(now int64) bool {
	diff, _ := timeutils.DiffDay(now, o.lastDayRewTime)
	if diff > 0 {
		return false
	}
	return true
}

func (o *WeekData) GetExpireTime() int64 {
	return o.expireTime
}

func (o *WeekData) GetCycleDay() int32 {
	return o.cycDay
}

func (o *WeekData) GetNextCycleDay() int32 {
	return o.GetCycleDay() + 1
}

func (o *WeekData) setExpireTime(expireTime int64) {
	o.expireTime = expireTime
}

func (o *WeekData) cycleDayIncr(now int64) {
	o.cycDay += 1
	o.lastDayRewTime = now
}
