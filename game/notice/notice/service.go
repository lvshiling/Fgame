package notice

import (
	"fgame/fgame/core/runner"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	noticeeventtypes "fgame/fgame/game/notice/event/types"
	"sync"
)

type NoticeService interface {
	runner.Task
	// 添加GM公告
	AddGmNotice(content string, beginTime, endTime, intervalTime int64)
}

type noticeService struct {
	rwm          sync.RWMutex
	gmNoticeList []*GmNoticeData //后台公告列表
}

func (s *noticeService) init() (err error) {
	return
}

func (s *noticeService) Heartbeat() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.checkBroadcastNotice()
}

func (s *noticeService) AddGmNotice(content string, beginTime, endTime, intervalTime int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	notice := &GmNoticeData{
		content:   content,
		beginTime: beginTime,
		endTime:   endTime,
		interval:  intervalTime,
	}
	s.gmNoticeList = append(s.gmNoticeList, notice)
}

func (s *noticeService) checkBroadcastNotice() {
	now := global.GetGame().GetTimeService().Now()
	var newNoticeList []*GmNoticeData
	for _, noticeData := range s.gmNoticeList {
		if now > noticeData.endTime {
			continue
		}

		// 剔除过期的
		newNoticeList = append(newNoticeList, noticeData)

		if now < noticeData.beginTime {
			continue
		}
		if !noticeData.isNotice() {
			continue
		}

		gameevent.Emit(noticeeventtypes.EventTypeBroadcastNotice, noticeData, nil)
		gameevent.Emit(noticeeventtypes.EventTypeBroadcastSystem, noticeData, nil)
		noticeData.lastNoticeTime = now
	}
	s.gmNoticeList = newNoticeList
}

type GmNoticeData struct {
	content        string
	beginTime      int64
	endTime        int64
	interval       int64
	lastNoticeTime int64
}

func (d *GmNoticeData) isNotice() bool {
	now := global.GetGame().GetTimeService().Now()
	diff := now - d.lastNoticeTime
	if diff < d.interval {
		return false
	}

	return true
}

func (d *GmNoticeData) GetContent() string {
	return d.content
}

var (
	s    *noticeService
	once sync.Once
)

func Init() (err error) {
	once.Do(func() {
		s = &noticeService{}
		err = s.init()
	})

	return
}

func GetNoticeService() NoticeService {
	return s
}
