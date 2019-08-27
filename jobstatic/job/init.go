package job

import (
	"fgame/fgame/core/db"
	mongo "fgame/fgame/core/mongo"
	"fgame/fgame/core/redis"
	"fgame/fgame/jobstatic/job/onlinestat"
	"fgame/fgame/jobstatic/job/playerstat"
	"fgame/fgame/jobstatic/job/serverdailystat"
)

var (
	jobManager *JobManager
)

func Init(ds db.DBService, rs redis.RedisService, mg mongo.MongoService, mgconfig *mongo.MongoConfig, centerDs db.DBService) {
	jobManager = NewJobManager()
	playerStats := playerstat.NewPlayerStatJob(ds, rs, mg, mgconfig)
	jobManager.AddJob(playerStats)
	serverStats := onlinestat.NewServerOnLineJob(ds, rs, mg, mgconfig)
	jobManager.AddJob(serverStats)
	// poolStats := supportpool.NewSupportPoolJob(ds, rs, centerDs)
	// jobManager.AddJob(poolStats)
	serverDailJob := serverdailystat.NewServerDailyStatJob(ds, rs, mg, mgconfig, centerDs)
	jobManager.AddJob(serverDailJob)
}

func Start() {
	jobManager.Start()
}

func Stop() {
	jobManager.Stop()
}
