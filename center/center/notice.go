package center

import (
	"fgame/fgame/center/store"
)

type NoticeInfo struct {
	id         int64
	platformId int32
	content    string
	updateTime int64
	createTime int64
	deleteTime int64
}

func (info *NoticeInfo) GetId() int64 {
	return info.id
}

func (info *NoticeInfo) GetContent() string {
	return info.content
}

func (info *NoticeInfo) FromEntity(e *store.NoticeEntity) {
	info.id = e.Id
	info.platformId = e.PlatformId
	info.content = e.Content
	info.updateTime = e.UpdateTime
	info.deleteTime = e.DeleteTime
	info.createTime = e.CreateTime
}
func newNoticeInfo() *NoticeInfo {
	info := &NoticeInfo{}
	return info
}
