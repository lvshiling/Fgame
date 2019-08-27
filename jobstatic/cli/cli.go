package cli

import (
	"encoding/json"
	"fgame/fgame/core/db"
	"fgame/fgame/core/redis"
	"fgame/fgame/pkg/osutils"
	"io/ioutil"
	"os"
	"path/filepath"

	mongo "fgame/fgame/core/mongo"
	"fgame/fgame/jobstatic/job"

	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"github.com/urfave/cli"
)

var (
	debug                    = false
	configFile               = ""
	done       chan struct{} = make(chan struct{})
)

func Start() {
	app := cli.NewApp()
	app.Name = "jobstatic"
	app.Usage = "jobstatic [global options]"

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config,c",
			Value:       configFile,
			Usage:       "config file",
			Destination: &configFile,
		},
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "debug",
			Destination: &debug,
			EnvVar:      "JOB_DEBUG",
		},
	}
	app.Before = before
	app.Action = start
	app.Run(os.Args)
}

func before(ctx *cli.Context) error {
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.DebugLevel: "./logs/info.log",
		log.InfoLevel:  "./logs/info.log",
		log.WarnLevel:  "./logs/warn.log",
		log.ErrorLevel: "./logs/error.log",
	}))

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
	// log.SetLevel(log.DebugLevel)
	return nil
}

func start(ctx *cli.Context) {
	var sc *serverConfig

	config, err := filepath.Abs(configFile)
	if err != nil {
		log.Fatalln("filepath abs failed:", err)
	}

	tsc, err := newServerConfig(config)
	if err != nil {
		log.Fatalln("read config file failed:", err)
	}
	sc = tsc

	ds, err := db.NewDBService(sc.DB)
	if err != nil {
		log.Fatalf("init db service failed:%#v,%#v", sc.DB, err)
	}
	centerDs, err := db.NewDBService(sc.CenterDB)
	if err != nil {
		log.Fatalf("init db service failed:%#v,%#v", sc.CenterDB, err)
	}
	// jobservice := job.NewJobService(db, sc.JS)
	rs, err := redis.NewRedisService(sc.Redis)
	if err != nil {
		log.Fatalf("init redis service failed:%#v,%#v", sc.DB, err)
	}
	mgdb, err := mongo.NewMongoService(sc.Mongo)
	if err != nil {
		log.Fatalln("init mongodb failed:", err)
	}
	hooker := osutils.NewInterruptHooker()
	hooker.AddHandler(osutils.InterruptHandlerFunc(stop))
	job.Init(ds, rs, mgdb, sc.Mongo, centerDs)
	job.Start()
	log.Println("定时作业运行中... ")
	select {
	case <-done:
		{
			log.Printf("停止...")
		}
	}
}

func stop() {
	job.Stop()
	done <- struct{}{}
}

func newServerConfig(configFile string) (sc *serverConfig, err error) {
	c, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	sc = &serverConfig{}
	err = json.Unmarshal(c, sc)
	if err != nil {
		return nil, err
	}
	return
}

type serverConfig struct {
	Redis    *redis.RedisConfig `json:"redis"`
	DB       *db.DbConfig       `json:"db"`
	Mongo    *mongo.MongoConfig `json:"mongo"`
	CenterDB *db.DbConfig       `json:"centerDb"`
}
