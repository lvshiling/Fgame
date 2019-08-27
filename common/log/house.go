package log

type HouseLogReason int32

const (
	HouseLogReasonGM HouseLogReason = iota + 1
)

func (r HouseLogReason) Reason() int32 {
	return int32(r)
}

var (
	houseLogReasonMap = map[HouseLogReason]string{
		HouseLogReasonGM: "gm修改",
	}
)

func (r HouseLogReason) String() string {
	return houseLogReasonMap[r]
}
