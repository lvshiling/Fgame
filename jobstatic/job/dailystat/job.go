package dailystat

import (
	"fgame/fgame/core/db"
	"fgame/fgame/core/mongo"
	"fgame/fgame/core/redis"

	log "github.com/Sirupsen/logrus"
)

type ServerDailyJob struct {
	ds       db.DBService
	centerDs db.DBService
	rs       redis.RedisService
	ms       mongo.MongoService
	msConfig *mongo.MongoConfig
}

func (m *ServerDailyJob) GetId() string {
	return "serverDailyJob"
}

func (m *ServerDailyJob) Run() error {
	log.Debug("作业ServerDailyJob开始运行...")

	return nil
}

func (m *ServerDailyJob) GetTickSecond() int64 {
	return 25
}
