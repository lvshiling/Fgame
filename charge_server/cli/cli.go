package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	api "fgame/fgame/charge_server/api"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/remote"
	"fgame/fgame/sdk"

	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"

	_ "fgame/fgame/game/remote/cmd/handler"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rifflock/lfshook"
	"github.com/xozrc/pkg/osutils"
)

var (
	debug      = false
	configFile = "./config/config.json"
)

const (
	apiPath = "/api"
)

func Start() {
	app := cli.NewApp()
	app.Name = "charge"
	app.Usage = "charge [global options]"

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
		log.SetLevel(log.InfoLevel)
	}
	return nil
}

type serverConfig struct {
	Host   string                 `json:"host"`
	Port   int                    `json:"port"`
	Redis  *coreredis.RedisConfig `json:"redis"`
	DB     *coredb.DbConfig       `json:"db"`
	Remote *remote.RemoteConfig   `json:"remote"`
	Sdk    string                 `json:"sdk"`
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

func start(ctx *cli.Context) {

	config, err := filepath.Abs(configFile)
	if err != nil {
		log.Fatalln("filepath abs failed:", err)
	}

	sc, err := newServerConfig(config)
	if err != nil {
		log.Fatalln("read config file failed:", err)
	}

	db, err := coredb.NewDBService(sc.DB)
	if err != nil {
		log.Fatalln("init db service failed:", err)
	}

	rs, err := coreredis.NewRedisService(sc.Redis)
	if err != nil {
		log.Fatalln("init redis service failed:", err)
	}
	//初始化sdk
	err = sdk.Init(sc.Sdk)
	if err != nil {
		log.Fatalln("初始化sdk失败:", err)
	}
	//充值服务
	chargeService := charge.NewChargeService(db, rs)

	//远程服务
	remoteService, err := remote.NewRemoteService(sc.Remote, db, rs)
	if err != nil {
		log.Fatalln("初始化远程服务失败:", err)
	}
	remoteService.Start()

	n := negroni.Classic()
	router := mux.NewRouter().StrictSlash(true)
	subrouter := router.PathPrefix(apiPath).Subrouter()
	api.Router(subrouter)
	n.UseFunc(charge.SetupChargeServiceHandler(chargeService))
	n.UseFunc(remote.SetupRemoteServiceHandler(remoteService))
	n.UseHandler(router)

	//register interruput handler
	addr := fmt.Sprintf("%s:%d", sc.Host, sc.Port)
	hooker := osutils.NewInterruptHooker()
	hooker.AddHandler(osutils.InterruptHandlerFunc(stop))

	log.Println("trying to listen ", addr)

	err = http.ListenAndServe(addr, n)
	if err != nil {
		log.Fatalln("charge:监听错误:", err)
	}
}

func stop() {
	log.Println("charge:服务器关闭")
}
