package types

import (
	commonlog "fgame/fgame/common/log"
)

type ChessEventType string

const (
	EventTypeAttendChess    ChessEventType = "AttendChess"    //参与棋局
	EventTypeAttendChessLog                = "AttendChessLog" //参与棋局后台日志
)

//棋局后台日志
type PlayerAttendChessLogEventData struct {
	attendTimes int32
	reason      commonlog.ChessLogReason
	reasonText  string
}

func CreatePlayerAttendChessLogEventData(attendTimes int32, reason commonlog.ChessLogReason, reasonText string) *PlayerAttendChessLogEventData {
	d := &PlayerAttendChessLogEventData{
		attendTimes: attendTimes,
		reason:      reason,
		reasonText:  reasonText,
	}
	return d
}

func (d *PlayerAttendChessLogEventData) GetAttendTimes() int32 {
	return d.attendTimes
}

func (d *PlayerAttendChessLogEventData) GetReason() commonlog.ChessLogReason {
	return d.reason
}

func (d *PlayerAttendChessLogEventData) GetReasonText() string {
	return d.reasonText
}
