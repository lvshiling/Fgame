package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
)

func BuildSCNoticeNum(content []byte, intervalTime int64, num int32) *uipb.SCNotice {
	notice := &uipb.SCNotice{}
	notice.Content = content
	notice.IntervalTime = &intervalTime
	notice.Num = &num
	return notice
}

func NoticeTimeBroadcast(content []byte, intervalTime int64, startTime int64, endTime int64) *uipb.SCNotice {
	notice := &uipb.SCNotice{}
	notice.Content = content
	notice.IntervalTime = &intervalTime
	notice.StartTime = &startTime
	notice.EndTime = &endTime
	return notice
}
