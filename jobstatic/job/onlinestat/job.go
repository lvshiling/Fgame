package onlinestat

import (
	"fgame/fgame/core/db"
	mongo "fgame/fgame/core/mongo"
	"fgame/fgame/core/redis"
	"fgame/fgame/pkg/timeutils"
	"time"

	"github.com/jinzhu/gorm"

	"fgame/fgame/gm/gamegm/gm/center/staticreport/model"
	stservice "fgame/fgame/gm/gamegm/gm/center/staticreport/service"

	log "github.com/Sirupsen/logrus"
	redigo "github.com/garyburd/redigo/redis"
)

const (
	initMaxDays = 30
)

/*******接口开始*************/

type ServerOnLineJob struct {
	ds       db.DBService
	rs       redis.RedisService
	ms       mongo.MongoService
	msConfig *mongo.MongoConfig
}

func (m *ServerOnLineJob) GetId() string {
	return "serverOnLine"
}

func (m *ServerOnLineJob) Run() error {
	log.Debug("作业ServerOnLineJob开始运行...")
	service := stservice.NewReportStatic(m.ms, m.msConfig)
	startTime, err := m.getLastTime()
	if err != nil {
		return err
	}
	now := timeutils.TimeToMillisecond(time.Now())
	beginNow, _ := timeutils.BeginOfNow(now)
	if startTime == 0 {
		firstMongoTime, err := service.GetOnLineServerFirstTime()
		if err != nil {
			return err
		}
		beginFirstTime, _ := timeutils.BeginOfNow(firstMongoTime)
		startTime = beginFirstTime
	}
	//mongo未开始记录数据
	if startTime == 0 {
		startTime = beginNow
		err = m.setLastTime(beginNow)
		if err != nil {
			return err
		}
	}
	if startTime >= beginNow {
		return nil
	}

	endTime := beginNow
	for exStart := startTime; exStart < endTime; exStart = exStart + int64(timeutils.DAY) {
		exEnd := exStart + int64(timeutils.DAY)
		log.WithFields(log.Fields{
			"start": exStart,
			"end":   exEnd,
		}).Debug("开始统计在线")
		onLineData, err := service.GetOnLinePlayerNumGroupServer(exStart, exEnd)
		if err != nil {
			return err
		}
		err = m.saveDB(onLineData)
		if err != nil {
			return err
		}
		err = m.setLastTime(beginNow)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *ServerOnLineJob) GetTickSecond() int64 {
	// return 60 * 60
	// return 10 * 60
	return 25
}

/*******接口结束*************/

func (m *ServerOnLineJob) saveDB(data []*model.OnLinePlayerStaticServer) error {
	for _, value := range data {
		now := timeutils.TimeToMillisecond(time.Now())
		firstLoginInfo := &ServerOnLineStatic{}
		exdb := m.ds.DB().Where("playerId=? and onLineIndex=0 and platformId=? and serverId=?", value.Id.PlayerId, value.Id.Platform, value.Id.ServerId).First(firstLoginInfo)
		if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
			return exdb.Error
		}
		beginTimeLogin, _ := timeutils.BeginOfNow(value.MaxLoginTime)
		if firstLoginInfo.Id == 0 { //表示未存在
			firstLoginInfo.PlayerId = value.Id.PlayerId
			firstLoginInfo.ServerId = value.Id.ServerId
			firstLoginInfo.PlatformId = value.Id.Platform
			firstLoginInfo.OnLineIndex = 0
			firstLoginInfo.OnLineTime = value.MaxLoginTime
			firstLoginInfo.UpdateTime = now
			firstLoginInfo.CreateTime = now
			firstLoginInfo.OnLineDate = beginTimeLogin
			exdb = m.ds.DB().Save(firstLoginInfo)
			if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
				return exdb.Error
			}
			continue
		}

		entityInfo := &ServerOnLineStatic{}
		exdb = m.ds.DB().Where("onLineTime = ? and playerId = ? and platformId=? and serverId=?", beginTimeLogin, value.Id.PlayerId, value.Id.Platform, value.Id.ServerId).First(entityInfo)
		if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
			return exdb.Error
		}
		if entityInfo.OnLineDate == firstLoginInfo.OnLineDate { //拿到的是起始的那条
			continue
		}
		index := int((beginTimeLogin - firstLoginInfo.OnLineDate) / int64(timeutils.DAY))
		if index > initMaxDays { //大于最大天数则更改最大天数记录
			exdb = m.ds.DB().Where("playerId = ? and onLineIndex=? and platformId=? and serverId=?", value.Id.PlayerId, initMaxDays, value.Id.Platform, value.Id.ServerId).First(entityInfo)
			if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
				return exdb.Error
			}
			index = initMaxDays
		}
		entityInfo.PlayerId = value.Id.PlayerId
		entityInfo.ServerId = value.Id.ServerId
		entityInfo.PlatformId = value.Id.Platform
		entityInfo.OnLineIndex = index
		entityInfo.OnLineTime = value.MaxLoginTime
		entityInfo.UpdateTime = now
		entityInfo.CreateTime = now
		entityInfo.OnLineDate = beginTimeLogin
		exdb = m.ds.DB().Save(entityInfo)
		if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
			return exdb.Error
		}
	}
	return nil
}

func (m *ServerOnLineJob) getLastTime() (int64, error) {
	key := getLastTimeKey()
	conn := m.rs.Pool().Get()
	defer conn.Close()
	result, err := redigo.Int64(conn.Do("GET", key))
	if err != nil && err != redigo.ErrNil {
		return 0, err
	}
	return result, nil
}

func (m *ServerOnLineJob) setLastTime(lastTime int64) error {
	key := getLastTimeKey()
	conn := m.rs.Pool().Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, lastTime)
	if err != nil {
		return err
	}
	return nil
}

func getLastTimeKey() string {
	return "serverOnLineStats_lastTime_new"
}

func NewServerOnLineJob(ds db.DBService, rs redis.RedisService, ms mongo.MongoService, msc *mongo.MongoConfig) *ServerOnLineJob {
	rst := &ServerOnLineJob{
		ds:       ds,
		rs:       rs,
		ms:       ms,
		msConfig: msc,
	}
	return rst
}
