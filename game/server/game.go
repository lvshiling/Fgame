package server

import (
	centertypes "fgame/fgame/center/types"
	_ "fgame/fgame/common/codec"
	_ "fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	fgamedb "fgame/fgame/core/db"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/module"
	fgameredis "fgame/fgame/core/redis"
	"fgame/fgame/core/session"
	"fgame/fgame/core/template"
	coretime "fgame/fgame/core/time"
	_ "fgame/fgame/game/activity"
	_ "fgame/fgame/game/additionsys"
	_ "fgame/fgame/game/alliance"
	_ "fgame/fgame/game/anqi"
	_ "fgame/fgame/game/arena"
	_ "fgame/fgame/game/arenapvp"
	_ "fgame/fgame/game/baby"
	_ "fgame/fgame/game/bagua"
	_ "fgame/fgame/game/battle"
	_ "fgame/fgame/game/bodyshield"
	_ "fgame/fgame/game/buff"
	_ "fgame/fgame/game/cache"
	_ "fgame/fgame/game/cangjingge"
	_ "fgame/fgame/game/cd"
	_ "fgame/fgame/game/center"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/charge"
	chargecharge "fgame/fgame/game/charge/charge"
	_ "fgame/fgame/game/chat"
	_ "fgame/fgame/game/chess"
	_ "fgame/fgame/game/christmas"
	_ "fgame/fgame/game/chuangshi"
	_ "fgame/fgame/game/click"
	gamecodec "fgame/fgame/game/codec"
	_ "fgame/fgame/game/collect"
	_ "fgame/fgame/game/common"
	_ "fgame/fgame/game/compensate"
	"fgame/fgame/game/constant/constant"
	"fgame/fgame/game/coupon"
	couponcoupon "fgame/fgame/game/coupon/coupon"
	_ "fgame/fgame/game/cross"
	_ "fgame/fgame/game/dan"
	_ "fgame/fgame/game/densewat"
	_ "fgame/fgame/game/dianxing"
	_ "fgame/fgame/game/dingshi"
	_ "fgame/fgame/game/dragon"
	_ "fgame/fgame/game/drop"
	_ "fgame/fgame/game/dummy"
	_ "fgame/fgame/game/email"
	_ "fgame/fgame/game/emperor"
	_ "fgame/fgame/game/equipbaoku"
	_ "fgame/fgame/game/fabao"
	_ "fgame/fgame/game/fashion"
	_ "fgame/fgame/game/feedbackfee"
	_ "fgame/fgame/game/feisheng"
	_ "fgame/fgame/game/fireworks"
	_ "fgame/fgame/game/foe"
	_ "fgame/fgame/game/found"
	_ "fgame/fgame/game/fourgod"
	_ "fgame/fgame/game/friend"
	_ "fgame/fgame/game/funcopen"
	_ "fgame/fgame/game/gem"
	"fgame/fgame/game/global"
	_ "fgame/fgame/game/gm"
	_ "fgame/fgame/game/godsiege"
	_ "fgame/fgame/game/goldequip"
	_ "fgame/fgame/game/guaji"
	_ "fgame/fgame/game/guidereplica"
	_ "fgame/fgame/game/hongbao"
	_ "fgame/fgame/game/jieyi"
	_ "fgame/fgame/game/lineup"
	_ "fgame/fgame/game/qixue"
	_ "fgame/fgame/game/ring"
	_ "fgame/fgame/game/shangguzhiling"
	"fgame/fgame/game/trade"
	_ "fgame/fgame/game/wushuangweapon"
	_ "fgame/fgame/game/xianzuncard"
	_ "fgame/fgame/game/yuxi"
	_ "fgame/fgame/game/zhenxi"
	"time"

	_ "fgame/fgame/game/house"
	_ "fgame/fgame/game/huiyuan"
	_ "fgame/fgame/game/shenyu"
	_ "fgame/fgame/game/systemcompensate"

	tradetrade "fgame/fgame/game/trade/trade"

	_ "fgame/fgame/game/fushi"
	_ "fgame/fgame/game/hunt"
	_ "fgame/fgame/game/inventory"
	_ "fgame/fgame/game/juexue"
	_ "fgame/fgame/game/lianyu"
	_ "fgame/fgame/game/lingtong"
	_ "fgame/fgame/game/lingtongdev"
	_ "fgame/fgame/game/lingyu"
	_ "fgame/fgame/game/liveness"
	_ "fgame/fgame/game/log"
	gamelog "fgame/fgame/game/log/log"
	_ "fgame/fgame/game/login"
	_ "fgame/fgame/game/longgong"
	_ "fgame/fgame/game/lucky"
	_ "fgame/fgame/game/major"
	_ "fgame/fgame/game/marry"
	_ "fgame/fgame/game/massacre"
	_ "fgame/fgame/game/material"
	"fgame/fgame/game/merge/merge"

	_ "fgame/fgame/game/arenaboss"
	_ "fgame/fgame/game/mingge"
	_ "fgame/fgame/game/misc"
	_ "fgame/fgame/game/moonlove"
	_ "fgame/fgame/game/mount"
	_ "fgame/fgame/game/myboss"
	_ "fgame/fgame/game/notice"
	_ "fgame/fgame/game/npc"
	_ "fgame/fgame/game/onearena"
	_ "fgame/fgame/game/outlandboss"
	_ "fgame/fgame/game/pk"
	"fgame/fgame/game/player"
	_ "fgame/fgame/game/player/module"
	"fgame/fgame/game/processor"
	_ "fgame/fgame/game/processor"
	_ "fgame/fgame/game/property"
	_ "fgame/fgame/game/quest"
	_ "fgame/fgame/game/quiz"
	_ "fgame/fgame/game/rank"
	_ "fgame/fgame/game/realm"
	_ "fgame/fgame/game/reddot"
	_ "fgame/fgame/game/register"
	_ "fgame/fgame/game/relive"
	_ "fgame/fgame/game/remote"
	remoteapi "fgame/fgame/game/remote/api"
	_ "fgame/fgame/game/robot"
	_ "fgame/fgame/game/scene"
	"fgame/fgame/game/scene/scene"
	_ "fgame/fgame/game/secretcard"
	gamesession "fgame/fgame/game/session"
	_ "fgame/fgame/game/shareboss"
	_ "fgame/fgame/game/shenfa"
	_ "fgame/fgame/game/shengtan"
	_ "fgame/fgame/game/shenmo"

	_ "fgame/fgame/game/itemskill"
	_ "fgame/fgame/game/shenqi"
	_ "fgame/fgame/game/shihunfan"
	_ "fgame/fgame/game/shop"
	_ "fgame/fgame/game/shopdiscount"
	_ "fgame/fgame/game/skill"
	_ "fgame/fgame/game/songbuting"
	_ "fgame/fgame/game/soul"
	_ "fgame/fgame/game/soulruins"
	_ "fgame/fgame/game/supremetitle"
	_ "fgame/fgame/game/synthesis"
	_ "fgame/fgame/game/systemskill"
	_ "fgame/fgame/game/team"
	_ "fgame/fgame/game/teamcopy"
	_ "fgame/fgame/game/template" //消息结构注册
	_ "fgame/fgame/game/tianmo"
	_ "fgame/fgame/game/tianshu"
	_ "fgame/fgame/game/title"
	_ "fgame/fgame/game/tower"
	_ "fgame/fgame/game/transportation"
	_ "fgame/fgame/game/treasurebox"
	_ "fgame/fgame/game/tulong"

	_ "fgame/fgame/game/daliwan"
	_ "fgame/fgame/game/feedbackfee"
	_ "fgame/fgame/game/trade"
	_ "fgame/fgame/game/tulongequip"
	_ "fgame/fgame/game/unrealboss"
	_ "fgame/fgame/game/vip"
	_ "fgame/fgame/game/wardrobe"
	_ "fgame/fgame/game/weapon"
	_ "fgame/fgame/game/week"
	_ "fgame/fgame/game/welfare"
	_ "fgame/fgame/game/welfarescene"
	_ "fgame/fgame/game/wing"
	_ "fgame/fgame/game/world"
	_ "fgame/fgame/game/worldboss"
	_ "fgame/fgame/game/xianfu"
	_ "fgame/fgame/game/xiantao"
	_ "fgame/fgame/game/xianti"
	_ "fgame/fgame/game/xinfa"
	_ "fgame/fgame/game/xuechi" //模板数据注册
	_ "fgame/fgame/game/xuedun"
	_ "fgame/fgame/game/xuekuang"
	_ "fgame/fgame/game/yinglingpu"
	_ "fgame/fgame/game/zhenfa"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
	"sync/atomic"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/xozrc/pkg/httputils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes" //注册模块
)

type RemoteServerOptions struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type RegisterServerOptions struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

//万世剑仙
//游戏配置
type GameOptions struct {
	Map      string `json:"map"`
	Template string `json:"template"`
	//各种服务的配置
	Db           *fgamedb.DbConfig          `json:"db"`
	Redis        *fgameredis.RedisConfig    `json:"redis"`
	Center       *center.CenterConfig       `json:"center"`
	Log          *gamelog.LogConfig         `json:"log"`
	Remote       *RemoteServerOptions       `json:"remote"`
	Register     *RegisterServerOptions     `json:"register"`
	Charge       *chargecharge.ChargeConfig `json:"charge"`
	Coupon       *couponcoupon.CouponConfig `json:"coupon"`
	Trade        *tradetrade.TradeOptions   `json:"trade"`
	GmOpen       bool                       `json:"gmOpen"`
	CrossDisable bool                       `json:"crossDisable"`
}

type Game struct {
	serverOptions *ServerOptions
	//配置选项
	options     *GameOptions
	timeService coretime.TimeService
	//TODO 本地日志服务

	//数据库
	db fgamedb.DBService
	//缓存
	rs fgameredis.RedisService
	//异步操作服务
	operationService global.OpeartionService
	//全局业务
	globalRunner *global.GlobalRunner
	//消息处理器
	msgHandler message.Handler
	//统计
	stats *stats

	//开启
	open bool
}

const (
	operationPoolSize = 50
)

func (g *Game) init() error {
	log.WithFields(
		log.Fields{}).Infoln("游戏数据初始化")
	//初始化时间
	ts := coretime.NewTimeService()
	heartbeat.SetupTimeService(ts)
	g.timeService = ts

	//初始化统计
	g.initStats()

	//初始化中心服务器
	err := g.initCenter()
	if err != nil {
		return err
	}
	//设置id服务
	idutil.SetupWorker(int64(center.GetCenterService().GetServerId()))

	//初始化日志服务
	err = g.initLogService()
	if err != nil {
		return err
	}
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
	err = g.initMessageHandler()
	if err != nil {
		return err
	}

	//初始化处理器
	processor.InitProcessor(g.msgHandler)
	//初始化全局业务
	g.globalRunner = global.NewGlobalRunner(g.msgHandler)

	//初始化合服
	err = merge.Init()
	if err != nil {
		return err
	}

	//初始化充值服务
	err = charge.Init(g.options.Charge)
	if err != nil {
		return err
	}
	//初始化兑换码服务
	err = coupon.Init(g.options.Coupon)
	if err != nil {
		return err
	}

	err = g.initBaseModule()
	if err != nil {
		return err
	}

	err = g.initModule()
	if err != nil {
		return err
	}
	ip, port, err := center.GetCenterService().GetTradeServer()
	if err != nil {
		return err
	}
	tradeOps := &tradetrade.TradeOptions{}
	tradeOps.Host = ip
	tradeOps.Port = port
	//初始化交易服务
	err = trade.Init(tradeOps)
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) initStats() {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化统计服务")
	g.stats = newStats()
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化统计服务")
}

func (g *Game) initLogService() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化日志服务")
	err = gamelog.Init(g.options.Log)
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{}).Infoln("游戏初始化日志服务")
	return
}

func (g *Game) initCenter() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化注册服务器")
	err = center.Init(g.options.Center)
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{}).Infoln("游戏初始化注册服务器成功")
	return
}

func (g *Game) initDb() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化数据库")
	g.db, err = fgamedb.NewDBService(g.options.Db)
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{}).Infoln("游戏初始化数据库成功")
	return
}

func (g *Game) initRedis() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化redis")
	g.rs, err = fgameredis.NewRedisService(g.options.Redis)
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{}).Infoln("游戏初始化redis成功")
	return
}

func (g *Game) initTemplate() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化模板数据")
	templateDir := g.options.Template
	mapDir := g.options.Map
	//初始化模板服务
	_, err = template.InitTemplateService(templateDir, mapDir)
	if err != nil {
		return err
	}

	log.WithFields(
		log.Fields{}).Infoln("游戏初始化模板数据成功")
	return
}

func (g *Game) initBaseModule() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化基础模块")

	for _, m := range module.GetBaseModules() {
		log.Infof("游戏正在初始化基础模块[%s]模板", m.String())
		err = m.InitTemplate()
		if err != nil {
			return
		}
	}
	for _, m := range module.GetBaseModules() {
		log.Infof("游戏正在初始化基础模块[%s]", m.String())
		err = m.Init()
		if err != nil {
			return
		}
	}
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化基础模块成功")
	return
}

func (g *Game) initModule() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化模块")

	for _, m := range module.GetModules() {
		log.Infof("游戏正在初始化模块[%s]模板", m.String())
		err = m.InitTemplate()
		if err != nil {
			return
		}
	}
	for _, m := range module.GetModules() {
		log.Infof("游戏正在初始化模块[%s]", m.String())
		err = m.Init()
		if err != nil {
			return
		}
	}
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化模块成功")
	return
}

func (g *Game) initMessageHandler() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化消息处理器")

	g.msgHandler = message.HandlerFunc(handleMessage)
	log.WithFields(
		log.Fields{}).Infoln("游戏初始化消息处理器成功")

	return
}

func (g *Game) start() (err error) {
	g.startStats()
	err = g.startLogService()
	if err != nil {
		return
	}
	//启动模块
	g.startModule()

	//启动交易
	g.startTrade()

	//启动全局业务
	g.startGlobal()

	//完成合并
	g.finishMerge()
	//启动场景
	g.startScene()

	//开始同步到中心服
	g.startCenter()

	//启动远程服务
	err = g.startRemoteServer()
	if err != nil {
		return
	}

	//开启
	g.open = true
	return
}

func (g *Game) startRemoteServer() (err error) {
	//启动远程服务
	remoteServer := g.remoteServer()
	host := g.options.Remote.Host
	port := g.options.Remote.Port
	addr := fmt.Sprintf("%s:%d", host, port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	errChan := make(chan error)

	go func() {
		err = remoteServer.Serve(ln)
		//TODO 重新启动
		if err != nil {
			errChan <- err
		}
	}()
	select {
	case terr := <-errChan:
		{
			err = terr
			return err
		}
	case <-time.After(time.Second * 5):
		{
			break
		}
	}
	log.WithFields(
		log.Fields{
			"addr": addr,
		}).Infoln("game:启动远程服务监听")

	return
}

func (g *Game) remoteServer() *grpc.Server {
	var options = []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(remoteRecovery)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(remoteRecovery)),
		)),
	}
	s := grpc.NewServer(options...)
	//注册跨服连接
	remoteapi.Server(s)
	return s
}

func remoteRecovery(p interface{}) (err error) {
	debug.PrintStack()
	return grpc.Errorf(codes.Internal, "%s", p)
}

func (g *Game) startStats() {
	g.stats.Start()
}

func (g *Game) startLogService() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在启动日志服务")
	gamelog.GetLogService().Start()

	log.WithFields(
		log.Fields{}).Infoln("游戏启动日志服务成功")
	return
}

func (g *Game) startModule() {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在启动模块")

	for _, m := range module.GetBaseModules() {
		log.Infof("游戏正在启动基础模块[%s]", m.String())
		m.Start()

	}

	for _, m := range module.GetModules() {
		log.Infof("游戏正在启动模块[%s]", m.String())
		m.Start()

	}
	log.WithFields(
		log.Fields{}).Infoln("游戏启动模块成功")

}

func (g *Game) startTrade() {
	trade.Start()
}

func (g *Game) startGlobal() {
	g.globalRunner.Start()

}

func (g *Game) startScene() {
	scene.GetSceneService().Start()
}

func (g *Game) startCenter() {
	center.GetCenterService().Start()
}

func (g *Game) finishMerge() {
	merge.GetMergeService().FinishMerge()
}

func (g *Game) stopModule() {
	for _, m := range module.GetModules() {
		log.Infof("游戏正在停止基础模块[%s]", m.String())
		m.Stop()
	}
	for _, m := range module.GetBaseModules() {
		log.Infof("游戏正在停止基础模块[%s]", m.String())
		m.Stop()
	}
}

func (g *Game) stop() (err error) {
	//停止接收玩家数据
	g.open = false
	//下线所有玩家
	player.GetOnlinePlayerManager().OfflineAllPlayers()
	//TODO 等候玩家下线
	//停止统计
	g.stats.Stop()
	//停止中心服
	center.GetCenterService().Stop()
	//交易
	trade.Stop()
	g.stopModule()

	//停止接收数据
	//完成剩余的场景消息
	scene.GetSceneService().Stop()
	//完成剩余的全局业务消息
	g.globalRunner.Stop()
	//操作服务关闭
	g.operationService.Stop()
	return
}

//对话开启
func (g *Game) SessionOpen(s session.Session) error {
	log.WithFields(
		log.Fields{
			"sessionId": s.Id(),
		}).Debug("game:对话开启")
	//设置玩家
	if !g.open {
		log.WithFields(
			log.Fields{
				"sessionId": s.Id(),
			}).Debug("game:服务器已经关闭")
		s.Close()
		return nil
	}

	defaultSessionOptions := &gamesession.SessionOptions{
		AuthTimeout:  30,
		PingTimeout:  90,
		SendMsgQueue: 2000,
	}
	//TODO 设置session配置
	ps := gamesession.NewPlayerSession(s, gamecodec.GetCodec(), defaultSessionOptions)
	nctx := gamesession.WithSession(s.Context(), ps)
	s.SetContext(nctx)

	//TODO 添加统计
	g.handleSessionOpenStats(s)
	return nil
}

//对话关闭
func (g *Game) SessionClose(s session.Session) error {
	// 添加统计
	g.handleSessionCloseStats(s)

	ps := gamesession.SessionInContext(s.Context())
	if ps == nil {
		log.WithFields(
			log.Fields{
				"sessionId": s.Id(),
			}).Debug("game:对话关闭")
		return nil
	}
	//关闭写循环 以防内存泄露
	defer ps.Close(false)

	p := ps.Player()
	if p == nil {
		return nil
	}
	//关闭
	tp := p.(player.Player)
	tp.Logout()

	return nil
}

func (g *Game) SessionReceive(s session.Session, msg []byte) error {
	g.handleSessionRecvStats(s, msg)
	log.WithFields(
		log.Fields{
			"sessionId": s.Id(),
			"msg":       msg,
		}).Debug("对话处理器,接收消息")

	//TODO
	//消息处理器处理
	err := processor.GetMessageProcessor().Process(s, msg)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) SessionSend(s session.Session, msg []byte) error {
	g.handleSessionSendStats(s, msg)
	return nil
}

func (g *Game) GetGlobalRunner() *global.GlobalRunner {
	return g.globalRunner
}

func (g *Game) GetGlobalUpdater() *global.GlobalUpdater {
	return g.globalRunner.GetUpdater()
}

//全局对象
func (g *Game) GetMessageHandler() message.Handler {
	return g.msgHandler
}

func (g *Game) GetDB() fgamedb.DBService {
	return g.db
}

func (g *Game) GetRedisService() fgameredis.RedisService {
	return g.rs
}

func (g *Game) GetOperationService() global.OpeartionService {
	return g.operationService
}

func (g *Game) GetTimeService() coretime.TimeService {
	return g.timeService
}

func (g *Game) GetOptions() *GameOptions {
	return g.options
}

func (g *Game) GetServerType() centertypes.GameServerType {
	return centertypes.GameServerTypeSingle
}

func (g *Game) GetServerId() int32 {
	return center.GetCenterService().GetServerId()
}

func (g *Game) GetServerIndex() int32 {
	return g.serverOptions.Id
}

func (g *Game) GetPlatform() int32 {
	return g.serverOptions.Platform
}

func (g *Game) GetServerIp() string {
	if g.options.Register == nil {
		return g.serverOptions.Host
	}
	return g.options.Register.Host
}

func (g *Game) GetServerPort() int32 {
	if g.options.Register == nil {
		return g.serverOptions.Port
	}
	return g.options.Register.Port
}

func (g *Game) GetServerTime() int64 {
	return center.GetCenterService().GetStartTime()
}

func (g *Game) Open() bool {
	return g.open
}

func (g *Game) GMOpen() bool {
	return g.options.GmOpen
}

func (g *Game) CrossDisable() bool {
	return g.options.CrossDisable
}

//-------------------统计--------------------------
func (g *Game) handleSessionOpenStats(s session.Session) {
	atomic.AddInt64(&g.stats.conn.CurrentConnections, 1)
	atomic.AddInt64(&g.stats.conn.TotalConnections, 1)
}

//消息接收统计中间件
func (g *Game) handleSessionRecvStats(s session.Session, msg []byte) {
	atomic.AddInt64(&g.stats.msg.InMsgs, 1)
	atomic.AddInt64(&g.stats.msg.InBytes, int64(len(msg)))
	return
}

//消息发送统计中间件
func (g *Game) handleSessionSendStats(s session.Session, msg []byte) {
	msgLen := int64(len(msg))
	atomic.AddInt64(&g.stats.msg.OutMsgs, 1)
	atomic.AddInt64(&g.stats.msg.OutBytes, msgLen)
	return
}

//连接统计中间件 连接关闭
func (g *Game) handleSessionCloseStats(s session.Session) error {
	atomic.AddInt64(&g.stats.conn.CurrentConnections, -1)
	return nil
}

//统计http请求
func (g *Game) handleStats(rw http.ResponseWriter, req *http.Request) {
	sr := convertToStatsResult(g.stats)
	err := httputils.WriteJSON(rw, http.StatusOK, sr)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

//连接统计http请求
func (g *Game) handleConnStats(rw http.ResponseWriter, req *http.Request) {
	err := httputils.WriteJSON(rw, http.StatusOK, g.stats.conn)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

//接收消息和发送消息http统计
func (g *Game) handleMsgStats(rw http.ResponseWriter, req *http.Request) {
	err := httputils.WriteJSON(rw, http.StatusOK, g.stats.msg)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func (g *Game) statsRouter(r *mux.Router) {
	r.Path("/").Handler(http.HandlerFunc(g.handleStats))
	r.Path("/conn").Handler(http.HandlerFunc(g.handleConnStats))
	r.Path("/msg").Handler(http.HandlerFunc(g.handleMsgStats))
}

func NewGame(serverOptions *ServerOptions, options *GameOptions) *Game {
	g := &Game{
		serverOptions: serverOptions,
		options:       options,
	}
	return g
}
