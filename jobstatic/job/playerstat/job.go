package playerstat

import (
	"encoding/json"
	"fgame/fgame/core/db"
	mongo "fgame/fgame/core/mongo"
	"fgame/fgame/core/redis"
	mongomodel "fgame/fgame/logserver/model"
	"time"

	"github.com/jinzhu/gorm"

	"fgame/fgame/pkg/timeutils"

	log "github.com/Sirupsen/logrus"
	redigo "github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2/bson"
)

const (
	eachDataTime  = 5 * 60  //每次捞取多少数据,秒
	maxTimeBefore = 60 * 60 //最迟的数据，即距离现在的秒数为数据的最大时间
)

type PlayerStatJob struct {
	ds       db.DBService
	rs       redis.RedisService
	ms       mongo.MongoService
	msConfig *mongo.MongoConfig
}

/*******接口开始*************/

func (m *PlayerStatJob) GetId() string {
	return "playerStat"
}
func (m *PlayerStatJob) Run() error {
	log.Debug("作业playerStat:开始运行...")
	lastTime, err := m.getLastTime()
	if err != nil {
		return err
	}
	now := timeutils.TimeToMillisecond(time.Now())
	maxTime := now - int64(time.Second/time.Millisecond*maxTimeBefore)
	if lastTime == 0 {
		lastTime, _ = timeutils.BeginOfNow(now)
	}
	if lastTime >= maxTime {
		return nil
	}

	for {
		if lastTime >= maxTime {
			break
		}
		endTime := lastTime + int64(time.Second/time.Millisecond*eachDataTime)
		if endTime > maxTime {
			endTime = maxTime
		}
		err = m.setLastTime(endTime)
		if err != nil {
			return err
		}
		err = m.stats(lastTime, endTime)
		if err != nil {
			m.setLastTime(lastTime)
			return err
		}

		lastTime = endTime
	}

	return nil
}

func (m *PlayerStatJob) GetTickSecond() int64 {
	return 60
	// return 10 * 60
}

/*******接口结束*************/

//统计数据
func (m *PlayerStatJob) stats(beginTime int64, endTime int64) error {
	log.WithFields(log.Fields{
		"beginTime": beginTime,
		"endTime":   endTime,
	}).Debug("开始统计...")
	mogoData, err := m.getMongoData(beginTime, endTime) //读取
	if err != nil {
		return err
	}
	if len(mogoData) == 0 {
		log.WithFields(log.Fields{
			"beginTime": beginTime,
			"endTime":   endTime,
		}).Debug("作业playerStat:获取数据为空...")
		return nil
	}
	statData, err := m.changeMgToDb(mogoData) //统计转换
	if err != nil {
		return err
	}
	err = m.saveDb(statData) //入库
	if err != nil {
		return err
	}
	return nil
}

//获取mongo里的数据
func (m *PlayerStatJob) getMongoData(beginTime int64, endTime int64) ([]*mongomodel.PlayerStats, error) {
	log.WithFields(log.Fields{
		"beginTime": beginTime,
		"endTime":   endTime,
	}).Debug("作业playerStat:获取Mongo数据...")
	rst := make([]*mongomodel.PlayerStats, 0)
	c := m.ms.Session().DB(m.msConfig.Database).C("player_stats")
	condition := bson.M{}
	datecondition := bson.M{}
	datecondition["$gte"] = beginTime
	datecondition["$lt"] = endTime
	condition["logtime"] = datecondition
	err := c.Find(condition).All(&rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func (m *PlayerStatJob) changeMgToDb(mgData []*mongomodel.PlayerStats) ([]*PlayerStatEntity, error) {
	rst := make([]*PlayerStatEntity, 0)
	hashRst := make(map[int64]map[string]int)
	for _, value := range mgData {
		// valueTime := time.Unix(value.LogTime/int64(time.Millisecond), 0)
		beginTime, _ := timeutils.BeginOfNow(value.LogTime)
		_, exists := hashRst[beginTime]
		if !exists {
			hashRst[beginTime] = make(map[string]int)
		}

		dataMap, err := m.changeStatsStringToMap(value.Stats)
		if err != nil {
			return nil, err
		}
		for key, value := range dataMap {
			_, exists = hashRst[beginTime][key]
			if !exists {
				hashRst[beginTime][key] = 0
			}
			hashRst[beginTime][key] += value
		}
	}
	for dayTime, mapValue := range hashRst {
		for kind, value := range mapValue {
			item := &PlayerStatEntity{}
			item.StatCount = value
			item.StatType = kind
			item.BeginTime = dayTime
			rst = append(rst, item)
		}
	}
	return rst, nil
}

func (m *PlayerStatJob) changeStatsStringToMap(dataString string) (map[string]int, error) {
	rst := make(map[string]int)
	byteData := []byte(dataString)
	err := json.Unmarshal(byteData, &rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func (m *PlayerStatJob) saveDb(dataArray []*PlayerStatEntity) error {
	saveDb := m.ds.DB().Begin()
	now := timeutils.TimeToMillisecond(time.Now())
	for _, value := range dataArray {
		beginTime := value.BeginTime
		statType := value.StatType
		dbInfo := &PlayerStatEntity{}
		exdb := saveDb.Where("beginTime=? and statType=?", beginTime, statType).First(dbInfo)
		if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
			saveDb.Rollback()
			return exdb.Error
		}
		if dbInfo.Id == 0 {
			dbInfo = value
			dbInfo.CreateTime = now
		} else {
			dbInfo.StatCount += value.StatCount
		}
		dbInfo.UpdateTime = now
		exdb = saveDb.Save(dbInfo)
		if exdb.Error != nil {
			return exdb.Error
		}
	}
	saveDb.Commit()
	return nil
}

func (m *PlayerStatJob) getLastTime() (int64, error) {
	key := getLastTimeKey()
	conn := m.rs.Pool().Get()
	defer conn.Close()
	result, err := redigo.Int64(conn.Do("GET", key))
	if err != nil && err != redigo.ErrNil {
		return 0, err
	}
	return result, nil
}

func (m *PlayerStatJob) setLastTime(lastTime int64) error {
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
	return "playerStats_lastTime"
}

func NewPlayerStatJob(ds db.DBService, rs redis.RedisService, ms mongo.MongoService, msc *mongo.MongoConfig) *PlayerStatJob {
	rst := &PlayerStatJob{
		ds:       ds,
		rs:       rs,
		ms:       ms,
		msConfig: msc,
	}
	return rst
}
