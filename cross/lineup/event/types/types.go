package types

type LineupEventType string

const (
	//玩家取消排队
	EventTypeLineupCancleLineUp       LineupEventType = "LineupCancleLineUp"       //玩家取消排队
	EventTypeLineupPlayerLineUpFinish                 = "LineupPlayerLineUpFinish" //玩家排队完成
	EventTypeLineupPlayerExit                         = "LineupPlayerExit"         //玩家退出无间炼狱场景
)

type CancleLineUpEventData struct {
	crossType int32
	index     int32
}

func CreateCancleLineUpEventData(crossType, index int32) *CancleLineUpEventData {
	d := &CancleLineUpEventData{
		crossType: crossType,
		index:     index,
	}

	return d
}

func (d *CancleLineUpEventData) GetCrossType() int32 {
	return d.crossType
}

func (d *CancleLineUpEventData) GetIndex() int32 {
	return d.index
}

type PlayerLineUpFinishEventData struct {
	crossType int32
	playerId  int64
}

func CreatePlayerLineUpFinishEventData(crossType int32, playerId int64) *PlayerLineUpFinishEventData {
	d := &PlayerLineUpFinishEventData{
		crossType: crossType,
		playerId:  playerId,
	}

	return d
}

func (d *PlayerLineUpFinishEventData) GetCrossType() int32 {
	return d.crossType
}

func (d *PlayerLineUpFinishEventData) GetPlayerId() int64 {
	return d.playerId
}
