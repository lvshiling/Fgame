package log

type BabyLogReason int32

const (
	BabyLogReasonGM BabyLogReason = iota + 1
	BabyLogReasonTalentActivate
	BabyLogReasonTalentRefresh
	BabyLogReasonLearn
)

func (zslr BabyLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	babyLogReasonMap = map[BabyLogReason]string{
		BabyLogReasonGM:             "gm修改",
		BabyLogReasonTalentActivate: "宝宝天赋激活,babyId:%d",
		BabyLogReasonTalentRefresh:  "宝宝天赋洗练,babyId:%d",
		BabyLogReasonLearn:          "宝宝读书升级,babyId:%d",
	}
)

func (ar BabyLogReason) String() string {
	return babyLogReasonMap[ar]
}
