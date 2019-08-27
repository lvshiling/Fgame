package server

import (
	centerclient "fgame/fgame/center/client"
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	_ "fgame/fgame/game/remote/cmd/handler"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/osutils"
	"fgame/fgame/trade_server/api"
	"fgame/fgame/trade_server/remote"
	"fgame/fgame/trade_server/trade"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
)

//服务器配置
type ServerOptions struct {
	Server    string `json:"server"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Path      string `json:"path"`
	PprofPort int    `json:"pprofPort"`
	TracePort int    `json:"tracePort"`
}

//交易服务器配置
type TradeServerOptions struct {
	Debug  bool                   `json:"debug"`
	Server *ServerOptions         `json:"server"`
	Redis  *coreredis.RedisConfig `json:"redis"`
	DB     *coredb.DbConfig       `json:"db"`
	Center *centerclient.Config   `json:"center"`
}

type Server struct {
	serverId   string
	options    *TradeServerOptions
	grpcServer *grpc.Server
}

func (s *Server) init() (err error) {
	return
}

//重载
func (s *Server) Reload() (err error) {
	return nil
}

func (s *Server) Start() (err error) {
	host := s.options.Server.Host
	pprofPort := s.options.Server.PprofPort
	pprofAddr := fmt.Sprintf("%s:%d", host, pprofPort)
	go func() {
		http.ListenAndServe(pprofAddr, nil)
	}()
	//设置随机种子
	now := time.Now().UnixNano()
	rand.Seed(now)
	//设置服务器id
	err = idutil.SetupWorker(1)
	if err != nil {
		return
	}
	//服务器启动
	log.Infoln("trade:交易服务器正在启动")
	err = s.startGRPCServer()
	if err != nil {
		return err
	}
	return
}

func (s *Server) Stop() {

	//服务器关闭
	log.Infoln("trade:交易服务器正在关闭")
	//停止接收消息
	//数据停止
	s.grpcServer.Stop()
	//退出
	log.Infoln("trade:交易服务器关闭")
	os.Exit(0)
}

func (s *Server) startGRPCServer() (err error) {
	log.Infoln("trade:交易服务器正在grpc服务")
	grpc.EnableTracing = true

	serverIdFile := s.options.Server.Server
	serverId, err := ReadServer(serverIdFile)
	if err != nil {
		return
	}
	s.serverId = serverId

	db, err := coredb.NewDBService(s.options.DB)
	if err != nil {
		log.Fatalln("init db service failed:", err)
	}

	rs, err := coreredis.NewRedisService(s.options.Redis)
	if err != nil {
		log.Fatalln("init redis service failed:", err)
	}

	centerClient, err := centerclient.NewClient(s.options.Center)
	if err != nil {
		log.Fatalln("init center client failed:", err)
	}

	tradeServer, err := trade.NewTradeServer(db, rs, centerClient)
	if err != nil {
		return
	}

	remoteService, err := remote.NewRemoteService(db, rs, centerClient)
	if err != nil {
		return
	}
	//启动
	tradeServer.Start()
	remoteService.Start()

	var options = []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}

	s.grpcServer = api.Server(tradeServer, remoteService, options...)

	addr := fmt.Sprintf("%s:%d", s.options.Server.Host, s.options.Server.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	s.startTrace()
	//注册钩子
	s.initHook()
	log.WithFields(
		log.Fields{
			"addr": addr,
		}).Infoln("服务器启动grpc服务成功")

	err = s.grpcServer.Serve(listener)
	if err != nil {
		return
	}

	log.Info("服务器grpc服务关闭")
	return nil
}

func (s *Server) startTrace() {

	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	tracePort := fmt.Sprintf("%s:%d", s.options.Server.Host, s.options.Server.TracePort)
	go http.ListenAndServe(tracePort, nil)
	log.WithFields(
		log.Fields{
			"addr": tracePort,
		}).Infoln("服务器启动grpc trace")

}

func (s *Server) initHook() {
	hook := osutils.NewInterruptHooker()
	hook.AddHandler(osutils.InterruptHandlerFunc(stop))
	go func() {
		hook.Run()
	}()
}

//TODO 修改通用
func ReadServer(serverIdFile string) (serverId string, err error) {

	abs, err := filepath.Abs(serverIdFile)
	if err != nil {
		return
	}
	_, err = os.Stat(abs)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
		} else {
			//生成新的serverId
			serverUUID := uuid.NewV1()

			serverId = serverUUID.String()
			err = ioutil.WriteFile(abs, []byte(serverId), os.ModePerm)
			if err != nil {
				return "", err
			}
			return serverId, nil
		}
	} else {
		bs, err := ioutil.ReadFile(abs)
		if err != nil {
			return "", err
		}
		serverId = string(bs)
		return serverId, nil
	}
}

//单例
var (
	once sync.Once
	s    *Server
)

//启动游戏服务器
func StartServer(options *TradeServerOptions) (err error) {
	once.Do(func() {
		log.Infoln("正在初始化交易服务器")

		ts := &Server{
			options: options,
		}
		s = ts
		err = s.init()
		if err != nil {
			return
		}
		err = s.Start()
		log.Infoln("交易服务器初始化成功")
	})

	return
}

func stop() {
	s.Stop()
}
