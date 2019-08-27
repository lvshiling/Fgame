package types

type TuLongBossType int32

const (
	//大boss
	TuLongBossTypeBig TuLongBossType = iota
	//小boss
	TuLongBossTypeSmall
)

func (t TuLongBossType) Valid() bool {
	switch t {
	case TuLongBossTypeBig,
		TuLongBossTypeSmall:
		return true
	}
	return false
}

type TuLongPosType int32

const (
	//boss出生
	TuLongPosTypeBoss TuLongPosType = iota
	//玩家出生
	TuLongPosTypePlayer
)

func (t TuLongPosType) Valid() bool {
	switch t {
	case TuLongPosTypeBoss,
		TuLongPosTypePlayer:
		return true
	}
	return false
}
