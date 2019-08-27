package server

import (
	centertypes "fgame/fgame/center/types"
	"fgame/fgame/common/message"
	fgamedb "fgame/fgame/core/db"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/module"
	fgameredis "fgame/fgame/core/redis"
	"fgame/fgame/core/session"
	"fmt"
	"net/http"

	grpcsession "fgame/fgame/core/session/grpc"
	grpcpb "fgame/fgame/core/session/grpc/pb"
	"fgame/fgame/core/template"
	coretime "fgame/fgame/core/time"
	_ "fgame/fgame/cross/activity"
	_ "fgame/fgame/cross/arena"
	_ "fgame/fgame/cross/arenaboss"
	_ "fgame/fgame/cross/arenapvp"
	arenapvpapi "fgame/fgame/cross/arenapvp/api"
	_ "fgame/fgame/cross/buff"
	crosscenter "fgame/fgame/cross/center"
	_ "fgame/fgame/cross/chat"
	_ "fgame/fgame/cross/chuangshi"
	crosscodec "fgame/fgame/cross/codec"
	_ "fgame/fgame/cross/common"
	_ "fgame/fgame/cross/drop"
	_ "fgame/fgame/cross/fashion"
	_ "fgame/fgame/cross/gm"
	_ "fgame/fgame/cross/goldequip"
	_ "fgame/fgame/cross/inventory"
	_ "fgame/fgame/cross/jieyi"
	_ "fgame/fgame/cross/lineup"
	_ "fgame/fgame/cross/lingyu"
	_ "fgame/fgame/cross/login"
	_ "fgame/fgame/cross/notice"
	_ "fgame/fgame/cross/pk"
	_ "fgame/fgame/cross/player"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	_ "fgame/fgame/cross/property"
	_ "fgame/fgame/cross/qixue"
	_ "fgame/fgame/cross/scene"
	settimeapi "fgame/fgame/cross/settime/api"
	_ "fgame/fgame/cross/shareboss"
	sharebossapi "fgame/fgame/cross/shareboss/api"
	_ "fgame/fgame/cross/shenfa"
	_ "fgame/fgame/cross/shenmo"
	shenmoapi "fgame/fgame/cross/shenmo/api"
	_ "fgame/fgame/cross/skill"
	_ "fgame/fgame/cross/soul"
	_ "fgame/fgame/cross/title"
	_ "fgame/fgame/cross/treasurebox"
	treasureboxapi "fgame/fgame/cross/treasurebox/api"
	tulongapi "fgame/fgame/cross/tulong/api"
	_ "fgame/fgame/cross/weapon"
	_ "fgame/fgame/cross/wing"
	_ "fgame/fgame/cross/worldboss"
	_ "fgame/fgame/cross/zhenxi"
	_ "fgame/fgame/game/dummy"
	"fgame/fgame/game/global"
	_ "fgame/fgame/game/robot"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/idutil"
	rankapi "fgame/fgame/rank/api"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc"

	//注册模块
	"fgame/fgame/game/constant/constant"
	_ "net/http/pprof"

	//模板数据注册
	_ "fgame/fgame/game/npc"
	_ "fgame/fgame/game/template"

	//消息结构注册
	_ "fgame/fgame/common/codec"

	_ "fgame/fgame/common/lang"
	_ "fgame/fgame/cross/godsiege"
	_ "fgame/fgame/cross/massacre"
	_ "fgame/fgame/cross/relive"

	_ "fgame/fgame/cross/tulong"
	_ "fgame/fgame/cross/xuechi"

	_ "fgame/fgame/cross/alliance"
	_ "fgame/fgame/cross/battle"
	_ "fgame/fgame/cross/collect"
	_ "fgame/fgame/cross/guaji"
	_ "fgame/fgame/cross/lianyu"
	_ "fgame/fgame/cross/lingtong"
	_ "fgame/fgame/cross/lingtongdev"
	_ "fgame/fgame/cross/mount"
	_ "fgame/fgame/cross/npc"
	_ "fgame/fgame/cross/team"
	_ "fgame/fgame/cross/teamcopy"
	_ "fgame/fgame/cross/xuekuang"
	"fgame/fgame/rank/rank"
)

//游戏配置
type CrossGameOptions struct {
	Map      string `json:"map"`
	Template string `json:"template"`
	//各种服务的配置
	Db     *fgamedb.DbConfig         `json:"db"`
	Redis  *fgameredis.RedisConfig   `json:"redis"`
	Center *crosscenter.CenterConfig `json:"center"`
	GameDb *fgamedb.DbConfig         `json:"gameDb"`
	GmOpen bool                      `json:"gmOpen"`
}

type CrossGame struct {
	serverOptions *CrossServerOptions
	//配置选项
	options     *CrossGameOptions
	timeService coretime.TimeService
	//本地日志服务
	cs *crosscenter.CenterService
	//数据库
	db fgamedb.DBService
	//游戏数据库 排行版使用
	gameDb fgamedb.DBService
	// 缓存
	rs fgameredis.RedisService
	//异步操作服务
	operationService global.OpeartionService
	//消息处理器
	msgHandler message.Handler
	//全局业务
	globalRunner *global.GlobalRunner

	//服务器时间
	serverTime int64
	open       bool
}

const (
	operationPoolSize = 50
)

func (g *CrossGame) init() error {
	log.WithFields(
		log.Fields{}).Infoln("跨服游戏数据初始化")
	//初始化时间
	g.initTimeService()

	//初始化中心服
	cs, err := crosscenter.NewCenterService(g.options.Center)
	if err != nil {
		return err
	}
	g.cs = cs
	idutil.SetupWorker(int64(cs.GetServerId()))

	g.serverTime = g.timeService.Now()
	//初始化数据库
	err = g.initDb()
	if err != nil {
		return err
	}
	//初始化redis
	err = g.initRedis()
	if err != nil {
		return err
	}

	//特殊处理
	if g.GetServerType() == centertypes.GameServerTypeGroup {
		err = g.initGameDb()
		if err != nil {
			return err
		}
	}
	//初始化异步操作服务
	g.operationService = global.NewOperationService(operationPoolSize)

	//初始化模板服务
	err = g.initTemplate()
	if err != nil {
		return err
	}
	//初始化常量
	err = constant.Init()
	if err != nil {
		return err
	}

	//初始化外部消息处理器
	if err = g.initMessageHandler(); err != nil {
		return err
	}

	//初始化全局业务
	g.globalRunner = global.NewGlobalRunner(g.msgHandler)
	//初始化排行版

	//TODO zrc:优化不同跨服加载

	switch g.GetServerType() {
	case centertypes.GameServerTypeGroup:
		{
			err = rank.Init(g.gameDb, g.rs)
			if err != nil {
				return err
			}
			break
		}
	}

	err = g.initModule()
	if err != nil {
		return err
	}

	return nil
}

func (g *CrossGame) initTimeService() {
	ts := coretime.NewTimeService()
	heartbeat.SetupTimeService(ts)
	g.timeService = ts
}

func (g *CrossGame) initDb() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("跨服游戏正在初始化数据库")
	g.db, err = fgamedb.NewDBService(g.options.Db)
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{}).Infoln("跨服游戏初始化数据库成功")
	return
}

func (g *CrossGame) initGameDb() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("跨服游戏正在初始化游戏数据库")
	g.gameDb, err = fgamedb.NewDBService(g.options.GameDb)
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{}).Infoln("跨服游戏初始化游戏数据库成功")
	return
}

func (g *CrossGame) initRedis() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("跨服游戏正在初始化redis")
	g.rs, err = fgameredis.NewRedisService(g.options.Redis)
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{}).Infoln("跨服游戏初始化redis成功")
	return
}

func (g *CrossGame) initTemplate() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("跨服游戏正在初始化模板数据")
	templateDir := g.options.Template
	mapDir := g.options.Map
	//初始化模板服务
	_, err = template.InitTemplateService(templateDir, mapDir)
	if err != nil {
		return err
	}

	log.WithFields(
		log.Fields{}).Infoln("跨服游戏初始化模板数据成功")
	return
}

func (g *CrossGame) initModule() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("跨服游戏正在初始化模块")
	for _, m := range module.GetBaseModules() {
		err = m.InitTemplate()
		if err != nil {
			return
		}
	}
	for _, m := range module.GetBaseModules() {
		err = m.Init()
		if err != nil {
			return
		}
	}

	for _, m := range module.GetModules() {
		err = m.InitTemplate()
		if err != nil {
			return
		}
	}
	for _, m := range module.GetModules() {
		err = m.Init()
		if err != nil {
			return
		}
	}
	log.WithFields(
		log.Fields{}).Infoln("跨服游戏正在初始化模块成功")
	return
}

func (g *CrossGame) initMessageHandler() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("跨服游戏正在初始化消息处理器")

	g.msgHandler = message.HandlerFunc(handleMessage)
	log.WithFields(
		log.Fields{}).Infoln("跨服游戏初始化消息处理器成功")

	return
}

func (g *CrossGame) start() (err error) {

	debug := gs.options.Server.Debug
	//启动统计
	if debug {
		host := gs.options.Server.Host
		pprofPort := gs.options.Server.PprofPort
		pprofAddr := fmt.Sprintf("%s:%d", host, pprofPort)
		go func() {
			http.ListenAndServe(pprofAddr, nil)
		}()
	}
	//启动模块
	err = g.startModule()
	if err != nil {
		return err
	}
	//启动全局业务
	g.startGlobal()

	if g.GetServerType() == centertypes.GameServerTypeGroup {
		//初始化排行榜
		err = rank.Start()
		if err != nil {
			return err
		}
	}
	//启动场景
	g.startScene()

	g.cs.Start()
	g.open = true
	return
}

func (g *CrossGame) startModule() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在启动模块")
	for _, m := range module.GetBaseModules() {
		m.Start()

	}
	for _, m := range module.GetModules() {
		m.Start()

	}
	log.WithFields(
		log.Fields{}).Infoln("游戏启动模块成功")
	return
}

func (g *CrossGame) startGlobal() {
	g.globalRunner.Start()
	return
}

func (g *CrossGame) startScene() {
	scene.GetSceneService().Start()
}

func (g *CrossGame) stopModule() {
	for _, m := range module.GetModules() {
		m.Stop()
	}
}

func (g *CrossGame) stop() (err error) {
	if g.GetServerType() == centertypes.GameServerTypeGroup {
		//初始化排行榜
		rank.Stop()
	}
	g.cs.Stop()
	g.stopModule()
	//停止接收数据
	//完成剩余的场景消息
	scene.GetSceneService().Stop()
	//完成剩余的全局业务消息
	g.globalRunner.Stop()
	return
}

func (g *CrossGame) Connect(s grpcpb.Connection_ConnectServer) error {
	sessionOpener := session.SessionHandlerFunc(g.SessionOpen)
	sessionCloser := session.SessionHandlerFunc(g.SessionClose)
	sessionReceiver := session.HandlerFunc(g.SessionReceive)
	sessionSender := session.HandlerFunc(g.SessionSend)
	h := grpcsession.NewGrpcHandler(sessionOpener, sessionCloser, sessionReceiver, sessionSender)
	h.Handle(s)
	return nil
}

//对话开启
func (g *CrossGame) SessionOpen(s session.Session) error {
	options := &gamesession.SessionOptions{
		AuthTimeout:  30,
		PingTimeout:  210,
		SendMsgQueue: 2000,
	}
	//设置玩家
	ps := gamesession.NewPlayerSession(s, crosscodec.GetCodec(), options)
	nctx := gamesession.WithSession(s.Context(), ps)
	s.SetContext(nctx)
	return nil
}

//对话关闭
func (g *CrossGame) SessionClose(s session.Session) error {

	ps := gamesession.SessionInContext(s.Context())
	if ps == nil {
		panic("SessionClose: never reach here")
	}
	defer ps.Close(false)
	p := ps.Player()
	if p == nil {
		return nil
	}

	tp := p.(*player.Player)
	tp.Logout()

	return nil
}

func (g *CrossGame) SessionReceive(s session.Session, msg []byte) error {
	// handleSessionRecvStats(s, msg)
	log.WithFields(
		log.Fields{
			"sessionId": s.Id(),
			"msg":       msg,
		}).Debug("对话处理器,接收消息")

	//消息处理器处理
	err := processor.GetMessageCrossProcessor().Process(s, msg)
	if err != nil {
		return err
	}

	return nil
}

func (g *CrossGame) SessionSend(s session.Session, msg []byte) error {
	// return handleSessionSendStats(s, msg)
	return nil
}

func (g *CrossGame) GetGlobalRunner() *global.GlobalRunner {
	return g.globalRunner
}

func (g *CrossGame) GetGlobalUpdater() *global.GlobalUpdater {
	return g.globalRunner.GetUpdater()
}

//全局对象
func (g *CrossGame) GetMessageHandler() message.Handler {
	return g.msgHandler
}

func (g *CrossGame) GetDB() fgamedb.DBService {
	return g.db
}

func (g *CrossGame) GetRedisService() fgameredis.RedisService {
	return g.rs
}

func (g *CrossGame) GetOperationService() global.OpeartionService {
	return g.operationService
}

func (g *CrossGame) GetTimeService() coretime.TimeService {
	return g.timeService
}

func (g *CrossGame) GetOptions() *CrossGameOptions {
	return g.options
}

func (g *CrossGame) GetServerType() centertypes.GameServerType {
	return centertypes.GameServerType(g.serverOptions.ServerType)
}

func (g *CrossGame) GetServerIndex() int32 {
	return g.serverOptions.Id
}

func (g *CrossGame) GetServerId() int32 {
	return g.cs.GetServerId()
}

func (g *CrossGame) GetPlatform() int32 {
	return g.serverOptions.Platform
}

func (g *CrossGame) GetServerIp() string {
	return g.serverOptions.Host
}

func (g *CrossGame) GetServerPort() int32 {
	return g.serverOptions.Port
}

func (g *CrossGame) GetServerTime() int64 {
	return g.serverTime
}

func (g *CrossGame) GetMergeServerTime() int64 {
	return g.serverTime
}

func (g *CrossGame) Open() bool {
	return g.open
}

func (g *CrossGame) GMOpen() bool {
	return g.options.GmOpen
}

func (g *CrossGame) CrossDisable() bool {
	return false
}

func (g *CrossGame) server() *grpc.Server {
	var options = []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}

	s := grpc.NewServer(options...)
	//注册跨服连接
	grpcpb.RegisterConnectionServer(s, gs.game)
	//TODO 注册服务
	//注册跨服boss查询
	sharebossapi.Server(s)
	rankapi.Server(s)
	tulongapi.Server(s)
	shenmoapi.Server(s)
	treasureboxapi.Server(s)
	settimeapi.Server(s)
	arenapvpapi.Server(s)
	// chuangshiapi.Server(s)
	return s
}

func NewCrossGame(serverOptions *CrossServerOptions, options *CrossGameOptions) *CrossGame {
	g := &CrossGame{
		serverOptions: serverOptions,
		options:       options,
	}
	return g
}
