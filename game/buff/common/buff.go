package common

import "fgame/fgame/game/global"
import bufftemplate "fgame/fgame/game/buff/template"
import scenetypes "fgame/fgame/game/scene/types"

type BuffObject interface {
	//获取上buff的
	GetOwnerId() int64
	//buff id
	GetBuffId() int32
	//获取组id
	GetGroupId() int32
	GetUseTime() int64
	//开始时间
	GetStartTime() int64
	//上次触发时间
	GetLastTouchTime() int64
	//获取持续时间
	GetDuration() int64
	//累加次数
	GetCulTime() int32
	//获取剩余时间
	GetRemainTime() int64
	//过期了
	IsExpired() bool
	GetTianFuList() []int32
}

type buffObject struct {
	OwnerId       int64   `json:"ownerId"`
	BuffId        int32   `json:"buffId"`
	GroupId       int32   `json:"groupId"`
	StartTime     int64   `json:"startTime"`
	UseTime       int64   `json:"useTime"`
	CulTime       int32   `json:"culTime"`
	LastTouchTime int64   `json:"lastTouchTime"`
	Duration      int64   `json:"duration"`
	TianFuList    []int32 `json:"tianFuList"`
}

func (bo *buffObject) GetOwnerId() int64 {
	return bo.OwnerId
}

func (bo *buffObject) GetBuffId() int32 {
	return bo.BuffId
}

func (bo *buffObject) GetGroupId() int32 {
	return bo.GroupId
}

func (bo *buffObject) GetStartTime() int64 {
	return bo.StartTime
}

func (bo *buffObject) GetLastTouchTime() int64 {
	return bo.LastTouchTime
}

func (bo *buffObject) GetCulTime() int32 {
	return bo.CulTime
}

func (bo *buffObject) GetDuration() int64 {
	return bo.Duration
}

func (bo *buffObject) GetRemainTime() int64 {
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(bo.BuffId)
	if buffTemplate == nil {
		//TODO 优化
		return int64(0)
	}

	now := global.GetGame().GetTimeService().Now()
	totalTime := bo.GetDuration()

	if buffTemplate.GetOfflineSaveType() == scenetypes.BuffOfflineSaveTypeTimeStop {
		return totalTime - (now - bo.StartTime) - bo.UseTime
	} else {
		return totalTime - (now - bo.StartTime)
	}
}

func (bo *buffObject) IsExpired() bool {
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(bo.BuffId)
	if buffTemplate == nil {
		//TODO 优化
		return true
	}
	if buffTemplate.TimeDuration == 0 {
		return false
	}
	if bo.GetRemainTime() <= 0 {
		return true
	}
	return false
}

func (bo *buffObject) GetTianFuList() []int32 {
	return bo.TianFuList
}

func (bo *buffObject) GetUseTime() int64 {
	return bo.UseTime
}

func (bo *buffObject) GetTimes() int32 {
	return bo.CulTime
}

func NewBuffObject(ownerId int64, buffId int32, groupId int32, startTime int64, useTime int64, culTime int32, lastTouchTime int64, duration int64, tianFuList []int32) BuffObject {
	bo := &buffObject{
		OwnerId:       ownerId,
		BuffId:        buffId,
		GroupId:       groupId,
		StartTime:     startTime,
		UseTime:       useTime,
		CulTime:       culTime,
		LastTouchTime: lastTouchTime,
		Duration:      duration,
	}
	return bo
}
