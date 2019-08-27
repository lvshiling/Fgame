package merge

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/dao"
	"fgame/fgame/pkg/idutil"

	log "github.com/Sirupsen/logrus"
)

type MergeService interface {
	Init() error
	IsMerge() bool
	GetMergeTime() int64
	FinishMerge()
	GMSetMergeTime(time int64)
}

type mergeService struct {
	mergeObj *MergeObject
}

func (s *mergeService) Init() (err error) {
	log.Infoln("merge:初始化")
	db := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(db, rs)
	if err != nil {
		return
	}
	serverId := global.GetGame().GetServerIndex()
	mergeEntity, err := dao.GetMergeDao().GetMergeEntity(serverId)
	if err != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	mergeObj := newMergeObject()
	if mergeEntity == nil {
		mergeObj.id, _ = idutil.GetId()
		mergeObj.merge = 0
		mergeObj.serverId = serverId
		mergeObj.createTime = now
		mergeObj.SetModified()
	} else {
		mergeObj.FromEntity(mergeEntity)
	}
	s.mergeObj = mergeObj
	return
}

func (s *mergeService) IsMerge() bool {
	return s.mergeObj.merge != 0
}

func (s *mergeService) FinishMerge() {
	if s.mergeObj.merge == 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	s.mergeObj.merge = 0
	s.mergeObj.updateTime = now
	s.mergeObj.SetModified()
	log.Infoln("merge:完成合服")
}

func (s *mergeService) GMSetMergeTime(time int64) {
	now := global.GetGame().GetTimeService().Now()
	s.mergeObj.mergeTime = time
	s.mergeObj.updateTime = now
	s.mergeObj.SetModified()
}

func (s *mergeService) GetMergeTime() int64 {
	return s.mergeObj.mergeTime
}

var (
	s *mergeService
)

func GetMergeService() MergeService {
	return s
}

func Init() (err error) {
	s = &mergeService{}
	err = s.Init()
	return
}
