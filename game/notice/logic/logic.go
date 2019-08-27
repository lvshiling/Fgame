package logic

import (
	"fgame/fgame/game/notice/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

//间隔时间+持续时间
func NoticeTimeBroadcast(content []byte, intervalTime int64, startTime int64, endTime int64) {
	noticeNum := pbutil.NoticeTimeBroadcast(content, intervalTime, startTime, endTime)
	player.GetOnlinePlayerManager().BroadcastMsg(noticeNum)
}

//间隔时间+次数
func NoticeNumBroadcast(content []byte, intervalTime int64, num int32) {
	noticeNum := pbutil.BuildSCNoticeNum(content, intervalTime, num)
	player.GetOnlinePlayerManager().BroadcastMsg(noticeNum)
}

//场景内公告 间隔时间+持续时间
func NoticeTimeBroadcastScene(s scene.Scene, content []byte, intervalTime int64, startTime int64, endTime int64) {
	if s == nil {
		return
	}
	noticeNum := pbutil.NoticeTimeBroadcast(content, intervalTime, startTime, endTime)
	s.BroadcastMsg(noticeNum)
}

//场景内公告 间隔时间+次数
func NoticeNumBroadcastScene(s scene.Scene, content []byte, intervalTime int64, num int32) {
	if s == nil {
		return
	}
	noticeNum := pbutil.BuildSCNoticeNum(content, intervalTime, num)
	s.BroadcastMsg(noticeNum)
}
