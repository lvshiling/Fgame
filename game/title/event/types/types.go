package types

type TitleEventType string

const (
	//称号改变事件
	EventTypeTitleChanged TitleEventType = "TitleChanged"
	//活动称号失效
	EventTypeTitleActivityOverdue TitleEventType = "TitleActivityFail"
	//称号激活
	EventTypeTitleActivate TitleEventType = "TitleActivate"
	//称号过期
	EventTypeTitleTimeExpire TitleEventType = "TitleTimeExpire"
)

type PlayerTitleTimeExpireEventData struct {
	titleId    int32
	expireTime int64
}

func CreatePlayerTitleTimeExpireEventData(titleId int32, expireTime int64) *PlayerTitleTimeExpireEventData {
	data := &PlayerTitleTimeExpireEventData{
		titleId:    titleId,
		expireTime: expireTime,
	}
	return data
}

func (d *PlayerTitleTimeExpireEventData) GetTitleId() int32 {
	return d.titleId
}

func (d *PlayerTitleTimeExpireEventData) GetExpireTime() int64 {
	return d.expireTime
}
