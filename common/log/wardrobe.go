package log

type WardrobeLogReason int32

const (
	WardrobeLogReasonGM WardrobeLogReason = iota + 1
	WardrobeLogReasonUpgrade
)

func (r WardrobeLogReason) Reason() int32 {
	return int32(r)
}

var (
	wardrobeLogReasonMap = map[WardrobeLogReason]string{
		WardrobeLogReasonGM:      "gm修改",
		WardrobeLogReasonUpgrade: "衣橱套装:%d(%s),喂养资质丹",
	}
)

func (r WardrobeLogReason) String() string {
	return wardrobeLogReasonMap[r]
}
