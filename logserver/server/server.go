package server

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

type LogServerOptions struct {
	GlobalOptions *Options       `json:"global"`
	ServerOptions *ServerOptions `json:"server"`
}

//服务器配置
type ServerOptions struct {
	Server    string `json:"server"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Path      string `json:"path"`
	PprofPort int    `json:"pprofPort"`
	StatsPort int    `json:"statsPort"`
}

type Server struct {
	serverId string
	options  *ServerOptions
}

func (s *Server) init() (err error) {

	return
}

func (s *Server) Start() (err error) {
	//服务器启动
	log.Infoln("服务器正在启动")

	return
}

func (as *Server) Stop() {
	//服务器关闭
	log.Infoln("服务器正在关闭")

	log.Infoln("服务器关闭")
	os.Exit(0)
}

//单例
var (
	gs *Server
)

//启动服务器
func Init(options *LogServerOptions) (err error) {
	log.Infoln("正在初始化服务器")

	if gs != nil {
		panic("repeat init  server")
	}

	gs = &Server{
		options: options.ServerOptions,
	}

	err = gs.init()
	if err != nil {
		return
	}

	g := newGlobal(options.GlobalOptions)
	err = g.Init()
	if err != nil {
		return
	}
	g.Start()
	log.Infoln("服务器初始化成功")

	return
}

func stop() {
	gs.Stop()
}
