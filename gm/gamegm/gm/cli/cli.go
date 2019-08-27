package cli

import (
	"encoding/json"
	api "fgame/fgame/gm/gamegm/gm/api"
	"fgame/fgame/gm/gamegm/gm/middleware"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	gmdb "fgame/fgame/gm/gamegm/db"
	gmredis "fgame/fgame/gm/gamegm/redis"

	gmchannel "fgame/fgame/gm/gamegm/gm/channel/service"
	recycleservice "fgame/fgame/gm/gamegm/gm/game/recycle/service"
	gmuserservice "fgame/fgame/gm/gamegm/gm/user/service"

	gmcenterplatform "fgame/fgame/gm/gamegm/gm/center/service"
	playerservice "fgame/fgame/gm/gamegm/gm/game/player/service"
	gmplatform "fgame/fgame/gm/gamegm/gm/platform/service"
	gmsensitive "fgame/fgame/gm/gamegm/gm/sensitive/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	loginhandler "fgame/fgame/gm/gamegm/monitor/login/handler"
	session "fgame/fgame/gm/gamegm/session"

	centerorder "fgame/fgame/gm/gamegm/gm/center/order/service"
	centerplatform "fgame/fgame/gm/gamegm/gm/center/platform/service"
	centerserver "fgame/fgame/gm/gamegm/gm/center/server/service"

	mongo "fgame/fgame/core/mongo"
	chatsetservice "fgame/fgame/gm/gamegm/gm/center/chatset/service"
	mongoservice "fgame/fgame/gm/gamegm/mglog/service"
	chathandler "fgame/fgame/gm/gamegm/monitor/chatmonitor/handler"
	userremote "fgame/fgame/gm/gamegm/remote/service"
	websocketsession "fgame/fgame/gm/gamegm/session/websocket"

	_ "fgame/fgame/game/remote/cmd/handler"
	alliservice "fgame/fgame/gm/gamegm/gm/game/alliance/service"
	singleserverservice "fgame/fgame/gm/gamegm/gm/game/singleserver/service"
	mailservice "fgame/fgame/gm/gamegm/gm/manage/mail/service"
	_ "fgame/fgame/gm/gamegm/gm/mglog/metadata/msg"
	nats "fgame/fgame/gm/gamegm/nats"

	serversupp "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/service"
	supportplayer "fgame/fgame/gm/gamegm/gm/manage/supportplayer/service"

	redeem "fgame/fgame/gm/gamegm/gm/center/redeem/service"
	rpservice "fgame/fgame/gm/gamegm/gm/center/staticreport/service"
	ntservice "fgame/fgame/gm/gamegm/gm/manage/notice/service"

	centeruser "fgame/fgame/gm/gamegm/gm/center/user/service"

	centernotice "fgame/fgame/gm/gamegm/gm/center/notice/service"
	centerset "fgame/fgame/gm/gamegm/gm/center/set/service"
	playerstat "fgame/fgame/gm/gamegm/gm/mglog/service"
	organizeservice "fgame/fgame/gm/gamegm/gm/organize/service"
	"fgame/fgame/gm/gamegm/gm/tick"

	tempservice "fgame/fgame/gm/gamegm/gm/template"

	corerunner "fgame/fgame/core/runner"

	yizqservice "fgame/fgame/gm/gamegm/gm/center/jiaoyizhanqu/service"
	dailyservice "fgame/fgame/gm/gamegm/gm/manage/serverdaily/service"

	feedbackfee "fgame/fgame/gm/gamegm/gm/feedbackfee/service"
	"fgame/fgame/gm/gamegm/gm/openapi"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rifflock/lfshook"
	"github.com/xozrc/pkg/osutils"
	"golang.org/x/net/websocket"
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
	app.Name = "gm"
	app.Usage = "gm [global options]"

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
		log.WarnLevel:  "./logs/info.log",
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
	Cert         string                     `json:"cert'`
	Key          string                     `json:"key"`
	Host         string                     `json:"host"`
	Port         int                        `json:"port"`
	Redis        *gmredis.RedisConfig       `json:"redis"`
	DB           *gmdb.DbConfig             `json:"db"`
	CenterDB     *gmdb.DbConfig             `json:"centerdb"`
	Login        *gmuserservice.LoginConfig `json:"gmUser"`
	GameDBArray  []*gamedbConfig            `json:"gamedb"`
	Nats         *nats.NatsOptions          `json:"nats"`
	Mongo        *mongo.MongoConfig         `json:"mongo"`
	Center       *centerConfig              `json:"center"`
	TemplatePath string                     `json:"template"`
}

type gamedbConfig struct {
	GroupId  int32          `json:"groupid"`
	GrpcHost string         `json:"grpcHost"`
	DB       *gmdb.DbConfig `json:"db"`
}

type centerConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
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
	natsService *nats.NatsGlobal
)

var (
	tickService tick.IGmTickService
)
var (
	supportPoolRunner     corerunner.Runner
	supportPoolRunnerTime time.Duration = 10 * time.Minute
)

type fmtLog struct{}

func (m *fmtLog) Output(calldepth int, s string) error {
	log.Info(s)
	return nil
}

func start(ctx *cli.Context) {
	// mgo.SetDebug(true)
	// mgo.SetLogger(&fmtLog{})
	config, err := filepath.Abs(configFile)
	if err != nil {
		log.Fatalln("filepath abs failed:", err)
	}

	sc, err := newServerConfig(config)
	if err != nil {
		log.Fatalln("read config file failed:", err)
	}
	n := negroni.Classic()
	//初始化db
	db, err := gmdb.NewDBService(sc.DB)
	if err != nil {
		log.Fatalln("init db service failed:", err)
	}
	centerdb, err := gmdb.NewDBService(sc.CenterDB)
	if err != nil {
		log.Fatalln("init centerdb service failed:", err)
	}

	gmdb.SetInstanceDB(db)
	gmuserservice.InitGmUserService()

	rs, err := gmredis.NewRedisService(sc.Redis)
	if err != nil {
		log.Fatalln("init redis service failed:", err)
	}

	los, err := gmuserservice.NewLoginService(sc.Login, db, rs)
	if err != nil {
		log.Fatalln("init login service failed:", err)
	}

	mgdb, err := mongo.NewMongoService(sc.Mongo)
	if err != nil {
		log.Fatalln("init mongodb failed:", err)
	}
	err = tempservice.GetGmTemplateService().Init(sc.TemplatePath)
	if err != nil {
		log.Fatalln("init template failed:", err)
	}
	supportPoolRunner = corerunner.NewRunner(supportPoolRunnerTime)

	mgService := mongoservice.NewMgLogService(mgdb, sc.Mongo)

	//设置全局跨域
	channel := gmchannel.NewChannelService(db)
	platform := gmplatform.NewPlatformService(db)
	sensitiveService := gmsensitive.NewSensitiveService(db)
	centerservice := gmcenterplatform.NewCenterPlatformService(centerdb)

	centerplatformService := centerplatform.NewCenterPlatformService(centerdb, db)
	centerServerService := centerserver.NewCenterServerService(centerdb)
	centerChatSetService := chatsetservice.NewChatSetService(centerdb)
	allianceService := alliservice.NewAllianceService()
	singleServerService := singleserverservice.NewSingleServerService()
	mailService := mailservice.NewMailService(db)
	serverSupportPool := serversupp.NewServerSupportPool(db, rs, centerdb)
	supportplayerService := supportplayer.NewSupportPlayerService(db)
	os := centerorder.NewOrderService(centerdb)
	reportService := rpservice.NewReportStatic(mgdb, sc.Mongo)
	playerStaticService := rpservice.NewPlayerStatic(db, centerdb)
	ns := ntservice.NewNoticeInfo(db)
	redeemService := redeem.NewRedeemService(centerdb)
	centerUserService := centeruser.NewCenterUserService(centerdb)
	tickService = tick.NewTickService(centerdb, db)
	serverList, err := centerServerService.GetAllCenterServerList()
	centerNoticeService := centernotice.NewLoginNoticeService(centerdb)
	organizeService := organizeservice.NewOrganizeService(centerServerService, platform)
	playerStatsService := playerstat.NewplayerStatsService(db)
	centerSetService := centerset.NewCenterSetService(centerdb)
	dailyServerService := dailyservice.NewServerDailyStatService(db)
	recycleService := recycleservice.NewRecycleService()
	jiaoYiZhanQuService := yizqservice.NewJiaoYiZhanQuService(centerdb)
	feedbackfeeService := feedbackfee.NewFeedBackFeeService(centerdb)
	supportPoolRunner.AddTask(serverSupportPool)
	err = supportPoolRunner.Start()
	if err != nil {
		log.Fatalln("启动定时器失败，supportPoolRunner:", err)
	}
	//初始化游戏DB、
	if serverList != nil && len(serverList) > 0 {
		for _, value := range serverList {
			dbhost := fmt.Sprintf("%s:%s", value.ServerDbIp, value.ServerDbPort)
			dbconfig := &gmdb.DbConfig{
				Debug:       true,
				Dialect:     "mysql",
				User:        value.ServerDBUser,
				Password:    value.ServerDBPassword,
				Host:        dbhost,
				DBName:      value.ServerDBName,
				ParseTime:   true,
				Charset:     "utf8mb4",
				MaxIdle:     50,
				MaxActive:   100,
				MaxLifeTime: 240,
			}
			gmdb.AddDbConfig(gmdb.GameDbLink(value.Id), dbconfig)

			//注册grpc服务
			grpcHost := fmt.Sprintf("%s:%s", value.ServerRemoteIp, value.ServerRemotePort)
			_, grpcerr := userremote.RegisterGrpc(int32(value.Id), grpcHost)
			if grpcerr != nil {
				log.Fatalln("init grpc service failed:", grpcerr)
				break
			}
		}
	}

	gmdb.InitDBManager(centerdb)

	cenRemoteService, cenerr := userremote.NewCenterService(sc.Center.Host, sc.Center.Port)
	if cenerr != nil {
		log.Fatalln("init center grpc service failed:", err)

	}
	cenRemoteNotice, cenerr := userremote.NewNotice(sc.Center.Host, sc.Center.Port)
	if cenerr != nil {
		log.Fatalln("init centerNotice grpc service failed:", err)

	}
	//远程服务
	userRemoteService := userremote.NewUserRemoteServer()

	//玩家
	pl := playerservice.NewPlayService()
	plMongoLogService := playerservice.NewPlayerMongoLogService(mgdb, sc.Mongo)

	n.UseHandlerFunc(handleCrossOrigin_all)

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix(apiPath).Methods(http.MethodOptions).HandlerFunc(handleCrossOrigin)

	subrouter := router.PathPrefix(apiPath).Subrouter()
	api.Router(subrouter)
	openapi.Router(subrouter)

	/**********************websocket******************/
	d := initDispatch()

	cs := monitor.NewCenterServer(centerdb)
	cs.SyncServer() //同步中心服务器信息到内存
	ms := monitor.NewMonitorService(d, los, cs)
	sessionOpener := session.SessionHandlerFunc(ms.SessionOpen)
	sessionCloser := session.SessionHandlerFunc(ms.SessionClose)
	sessionReceiver := session.HandlerFunc(ms.SessionReceive)
	sessionSender := session.HandlerFunc(ms.SessionSend)

	ms.StartTestTick()

	router.Handle("/websocket", websocket.Handler(websocketsession.NewWebsocketHandler(sessionOpener, sessionCloser, sessionReceiver, sessionSender).Handle))
	/**********************websocket******************/

	n.Use(negroni.NewRecovery())
	n.Use(SetupDealerLoginServiceHandler(los))
	n.Use(gmchannel.SetupChannelServiceHandler(channel))
	n.Use(gmplatform.SetupPlatformServiceHandler(platform))
	n.Use(gmcenterplatform.SetupCenterPlatformServiceHandler(centerservice))
	n.Use(playerservice.SetupPlayerServiceHandler(pl))
	n.Use(playerservice.SetupPlayerMongoLogServiceHandler(plMongoLogService))
	n.Use(gmsensitive.SetupSensitiveServiceHandler(sensitiveService))
	n.Use(gmuserservice.SetupUserServiceHandler(gmuserservice.GetGmUserServiceInstance()))
	n.Use(userremote.SetupUserRemoteServiceHandler(userRemoteService))
	n.Use(monitor.SetupCenterServerServiceHandler(cs))
	n.Use(chatsetservice.SetupChatSetServiceHandler(centerChatSetService))
	n.Use(alliservice.SetupAllianceServiceHandler(allianceService))
	n.Use(singleserverservice.SetupSingleServerServiceHandler(singleServerService))
	n.Use(mailservice.SetupMailServiceHandler(mailService))
	n.Use(serversupp.SetupServerSupportPoolHandler(serverSupportPool))
	n.Use(supportplayer.SetupPlayerServiceHandler(supportplayerService))
	n.Use(ntservice.SetupNoticeServiceHandler(ns))
	n.Use(recycleservice.SetupRecycleServiceHandler(recycleService))
	//日志服务配置
	n.Use(mongoservice.SetupMgLogServiceHandler(mgService))
	n.Use(playerstat.SetupPlayerStatsServiceHandler(playerStatsService))

	//中心平台配置
	n.Use(centerplatform.SetupCenterPlatformServiceHandler(centerplatformService))
	n.Use(centerserver.SetupCenterServerServiceHandler(centerServerService))
	n.Use(centerorder.SetupOrderServiceHandler(os))
	n.Use(rpservice.SetupStaticReportServiceHandler(reportService))
	n.Use(rpservice.SetupPlayerStaticHandler(playerStaticService))
	n.Use(redeem.SetupRedeemServiceHandler(redeemService))
	n.Use(centeruser.SetupCenterUserServiceHandler(centerUserService))
	n.Use(centernotice.SetupLoginNoticeServiceHandler(centerNoticeService))
	n.Use(centerset.SetupCenterSetServiceHandler(centerSetService))
	n.Use(dailyservice.SetupserverDailyStatHandler(dailyServerService))
	n.Use(yizqservice.SetupJiaoYiZhanQuServiceHandler(jiaoYiZhanQuService))
	n.Use(feedbackfee.SetupFeedBackFeeServiceHandler(feedbackfeeService))

	//中心remote
	n.Use(userremote.SetupCenterServiceHandler(cenRemoteService))
	n.Use(userremote.SetupNoticeHandler(cenRemoteNotice))

	//基础服务
	n.Use(organizeservice.SetupOrganizeServiceHandler(organizeService))

	n.Use(middleware.AuthHandlerMiddleware())
	n.Use(middleware.OpenApiHandlerMiddleware())

	n.UseHandler(router)
	// n.UseHandler(websocketRouter)

	//register interruput handler
	addr := fmt.Sprintf("%s:%d", sc.Host, sc.Port)
	hooker := osutils.NewInterruptHooker()
	hooker.AddHandler(osutils.InterruptHandlerFunc(stop))

	//日志监测服务,并注册日志服务
	natsService = nats.NewNatsGlobal(sc.Nats)
	natsService.OnReceiveLog(ms)
	//监测棋牌服务
	err = natsService.Init()
	if err != nil {
		log.Fatalln("init nats service init failed:", err)
	}
	err = natsService.Start()
	if err != nil {
		log.Fatalln("init nats service failed:", err)
	}

	log.Println("trying to listen ", addr)

	// err = http.ListenAndServeTLS(addr, sc.Cert, sc.Key, n)
	err = http.ListenAndServe(addr, n)
	if err != nil {
		log.Fatalln("error ", err)
	}
}

func stop() {
	log.Println("server stop")
	if natsService != nil {
		err := natsService.Stop()
		if err != nil {
			log.Println("nats stop error :", err)
		}
	}
	if tickService != nil {
		tickService.StopTick()
	}
	if supportPoolRunner != nil {
		supportPoolRunner.Stop()
	}
}

func initDispatch() *monitor.Dispatcher {
	d := monitor.NewDispatch()
	loginhandler.InitDispatcher(d)
	chathandler.InitDispatcher(d)
	return d
}

func SetupDealerLoginServiceHandler(ls gmuserservice.ILoginService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := gmuserservice.WithLoginService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}

func handleCrossOrigin(rw http.ResponseWriter, req *http.Request) {
	log.Debug("跨源")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "content-type,authorization")
	rw.Header().Set("Access-Control-Allow-Credentials", "true")
	rw.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	rw.Header().Set("Access-Control-Max-Age", "3600")
	rw.WriteHeader(http.StatusOK)
	log.Debug("跨源,成功")
}

func handleCrossOrigin_all(rw http.ResponseWriter, req *http.Request) {
	log.Debug("全局跨源")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "content-type,authorization")
	rw.Header().Set("Access-Control-Allow-Credentials", "true")
	rw.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	rw.Header().Set("Access-Control-Max-Age", "3600")
	if req.Method == http.MethodOptions {
		rw.WriteHeader(http.StatusOK)
	}

	log.Debug("全局跨源,成功")
}
