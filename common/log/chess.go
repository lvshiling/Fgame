package log

type ChessLogReason int32

const (
	ChessLogReasonGM ChessLogReason = iota + 1
	ChessLogReasonAttendLog
)

func (zslr ChessLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	chessLogReasonMap = map[ChessLogReason]string{
		ChessLogReasonGM:        "gm修改",
		ChessLogReasonAttendLog: "参与棋局，类型：%d,奖励id：%v",
	}
)

func (r ChessLogReason) String() string {
	return chessLogReasonMap[r]
}
