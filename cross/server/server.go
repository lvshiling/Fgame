package server

import (
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/osutils"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"time"

	log "github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
)

//跨服服务器配置
type CrossServerOptions struct {
	Debug      bool   `json:"debug"`
	Server     string `json:"server"`
	Id         int32  `json:"id"`
	Platform   int32  `json:"platform"`
	ServerType int32  `json:"serverType"`
	Host       string `json:"host"`
	Port       int32  `json:"port"`
	Path       string `json:"path"`
	PprofPort  int    `json:"pprofPort"`
	StatsPort  int    `json:"statsPort"`
}

//游戏服务器配置
type CrossGameServerOptions struct {
	Debug  bool                `json:"debug"`
	Game   *CrossGameOptions   `json:"game"`
	Server *CrossServerOptions `json:"server"`
}

type CrossGameServer struct {
	serverId string
	options  *CrossGameServerOptions
	game     *CrossGame
}

func (gs *CrossGameServer) init() (err error) {
	global.SetupGame(gs.game)
	err = gs.game.init()
	return
}

//重载
func (gs *CrossGameServer) Reload() (err error) {
	return nil
}

func (gs *CrossGameServer) Start() (err error) {

	//设置随机种子
	now := time.Now().UnixNano()
	rand.Seed(now)
	//服务器启动
	log.Infoln("跨服服务器正在启动")
	err = gs.game.start()
	if err != nil {
		return err
	}
	err = gs.startGrpcServer()
	if err != nil {
		return err
	}
	return
}

func (gs *CrossGameServer) Stop() {

	//服务器关闭
	log.Infoln("跨服服务器正在关闭")
	//停止接收消息
	//游戏数据停止
	err := gs.game.stop()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Errorln("跨服服务器关闭,错误")
	}

	//退出
	log.Infoln("跨服服务器关闭")
	os.Exit(0)
}

func (gs *CrossGameServer) startGrpcServer() (err error) {
	host := gs.options.Server.Host
	port := gs.options.Server.Port
	addr := fmt.Sprintf("%s:%d", host, port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	//注册钩子
	gs.initHook()
	log.WithFields(
		log.Fields{
			"addr": addr,
		}).Infoln("跨服服务器启动grpc服务成功")
	s := gs.game.server()

	err = s.Serve(ln)
	return
}

func (gs *CrossGameServer) initHook() {
	hook := osutils.NewInterruptHooker()
	hook.AddHandler(osutils.InterruptHandlerFunc(stop))
	go func() {
		hook.Run()
	}()
}

func (gs *CrossGameServer) readServer(serverIdFile string) (serverId string, err error) {

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
	gs *CrossGameServer
)

//启动跨服游戏服务器
func InitCrossGameServer(options *CrossGameServerOptions) (tgs *CrossGameServer, err error) {
	log.Infoln("正在初始化跨服游戏服务器")

	if gs != nil {
		panic("repeat init game server")
	}

	tgs = &CrossGameServer{
		options: options,
		game:    NewCrossGame(options.Server, options.Game),
	}
	err = tgs.init()
	if err != nil {
		return
	}
	gs = tgs
	log.Infoln("游戏跨服服务器初始化成功")

	return
}

func stop() {
	dumpStacks()
	gs.Stop()
}

var (
	stdFile = "./stack.log"
)

func dumpStacks() {
	buf := make([]byte, 163840000)
	buf = buf[:runtime.Stack(buf, true)]
	writeStack(buf)
}

func writeStack(buf []byte) {
	fd, _ := os.OpenFile(stdFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	fd.WriteString("\n\n\n\n\n")
	fd.WriteString("stdout:" + "\n\n")
	fd.Write(buf)
	fd.Close()
}
