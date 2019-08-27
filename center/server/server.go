package server

import (
	"fgame/fgame/center/api"
	"fgame/fgame/center/center"
	"fgame/fgame/pkg/osutils"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
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

//中心服务器配置
type CenterServerOptions struct {
	Debug  bool                  `json:"debug"`
	Server *ServerOptions        `json:"server"`
	Center *center.CenterOptions `json:"center"`
}

type Server struct {
	serverId   string
	options    *CenterServerOptions
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
	//设置随机种子
	now := time.Now().UnixNano()
	rand.Seed(now)
	//服务器启动
	log.Infoln("中心服务器正在启动")
	err = s.startGRPCServer()
	if err != nil {
		return err
	}
	return
}

func (s *Server) Stop() {

	//服务器关闭
	log.Infoln("中心服务器正在关闭")
	//停止接收消息
	//数据停止
	s.grpcServer.Stop()
	//退出
	log.Infoln("中心服务器关闭")
	os.Exit(0)
}

func (s *Server) startGRPCServer() (err error) {
	log.Infoln("服务器正在grpc服务")
	grpc.EnableTracing = true

	serverIdFile := s.options.Server.Server
	serverId, err := ReadServer(serverIdFile)
	if err != nil {
		return
	}

	s.serverId = serverId

	centerServer, err := center.NewCenterServer(s.options.Center)
	if err != nil {
		return
	}

	var options = []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}

	s.grpcServer = api.Server(centerServer, options...)

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
func StartServer(options *CenterServerOptions) (err error) {
	once.Do(func() {
		log.Infoln("正在初始化中心服务器")

		ts := &Server{
			options: options,
		}
		s = ts
		err = s.init()
		if err != nil {
			return
		}
		err = s.Start()
		log.Infoln("中心服务器初始化成功")
	})

	return
}

func stop() {
	s.Stop()
}
