package cli

import (
	"encoding/json"
	fgamedb "fgame/fgame/core/db"
	fgameredis "fgame/fgame/core/redis"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	loginapi "fgame/fgame/login/api"
	"fgame/fgame/login/login"

	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"github.com/urfave/cli"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/xozrc/pkg/osutils"
)

var (
	debug      = false
	configFile = ""
)

func Start() {

	app := cli.NewApp()
	app.Name = "user"
	app.Usage = "user [global options]"

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "debug ",
			Destination: &debug,
			EnvVar:      "USER_DEBUG",
		},
		cli.StringFlag{
			Name:        "config,c",
			Value:       configFile,
			Usage:       "config file",
			Destination: &configFile,
		},
	}

	app.Before = before
	app.Action = start
	app.Run(os.Args)
}

func before(c *cli.Context) error {
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.DebugLevel: "./logs/info.log",
		log.ErrorLevel: "./logs/error.log",
	}))

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
	return nil
}

type serverConfig struct {
	Host  string                  `json:"host"`
	Port  int                     `json:"port"`
	User  *login.LoginConfig      `json:"user"`
	Redis *fgameredis.RedisConfig `json:"redis"`
	DB    *fgamedb.DbConfig       `json:"db"`
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

var (
	apiPath = "/api"
)

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

	n := negroni.Classic()
	//初始化db
	dbService, err := fgamedb.NewDBService(sc.DB)
	if err != nil {
		log.Fatalf("init db service failed:%#v,%#v", sc.DB, err)
	}
	redisService, err := fgameredis.NewRedisService(sc.Redis)
	if err != nil {
		log.Fatalln("init redis service failed:", err)
	}

	userService := login.NewUserService(dbService, redisService)
	loginService, err := login.NewLoginService(sc.User, dbService, redisService, userService)
	if err != nil {
		log.Fatalln("init user service failed:", err)
	}

	addr := fmt.Sprintf("%s:%d", sc.Host, sc.Port)
	router := mux.NewRouter()

	subrouter := router.PathPrefix(apiPath).Subrouter()

	loginapi.Router(subrouter)

	n.UseFunc(setupUserServiceHandler(userService))
	n.UseFunc(setupLoginServiceHandler(loginService))

	n.UseHandler(router)

	//register interruput handler
	hooker := osutils.NewInterruptHooker()
	hooker.AddHandler(osutils.InterruptHandlerFunc(stop))
	log.Println("trying to listen ", addr)
	n.Run(addr)
}

func stop() {
	log.Println("stop server")
}

//设置用户服务
func setupUserServiceHandler(us login.UserService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := login.WithUserService(ctx, us)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}

func setupLoginServiceHandler(ls login.LoginService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := login.WithLoginService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
