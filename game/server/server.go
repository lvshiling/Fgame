package server

import (
	"fgame/fgame/core/session"
	tcpsession "fgame/fgame/core/session/tcp"
	websocketsession "fgame/fgame/core/session/websocket"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/osutils"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
)

//服务器配置
type ServerOptions struct {
	Debug     bool   `json:"debug"`
	Server    string `json:"server"`
	Host      string `json:"host"`
	Port      int32  `json:"port"`
	Id        int32  `json:"id"`
	Platform  int32  `json:"platform"`
	Path      string `json:"path"`
	PprofPort int32  `json:"pprofPort"`
	StatsPort int32  `json:"statsPort"`
}

//游戏服务器配置
type GameServerOptions struct {
	Debug  bool           `json:"debug"`
	Game   *GameOptions   `json:"game"`
	Server *ServerOptions `json:"server"`
}

type GameServer struct {
	serverId string
	options  *GameServerOptions
	game     *Game
	ln       net.Listener
}

func (gs *GameServer) init() (err error) {
	global.SetupGame(gs.game)
	err = gs.game.init()
	return
}

//重载
func (gs *GameServer) Reload() (err error) {
	return nil
}

func (gs *GameServer) Start() (err error) {
	//设置随机种子
	now := time.Now().UnixNano()
	rand.Seed(now)

	//启动统计服务

	debug := gs.options.Server.Debug
	//启动统计
	if debug {
		host := gs.options.Server.Host
		statsPort := gs.options.Server.StatsPort
		statsAddr := fmt.Sprintf("%s:%d", host, statsPort)
		statsMux := mux.NewRouter()
		statsR := statsMux.PathPrefix("/api/stats").Subrouter()
		gs.game.statsRouter(statsR)
		statsN := negroni.Classic()
		statsN.UseHandler(statsMux)
		go func() {
			statsN.Run(statsAddr)
		}()
		pprofPort := gs.options.Server.PprofPort
		pprofAddr := fmt.Sprintf("%s:%d", host, pprofPort)
		go func() {
			http.ListenAndServe(pprofAddr, nil)
		}()
	}

	//服务器启动
	log.Infoln("服务器正在启动")
	err = gs.game.start()
	if err != nil {
		return err
	}
	//启动远程服务
	err = gs.startTCPServer()
	if err != nil {
		return err
	}

	return
}

func (gs *GameServer) Stop() {
	//服务器关闭
	log.Infoln("服务器正在关闭")
	//停止接收消息
	// gs.ln.Close()

	//游戏数据停止
	err := gs.game.stop()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Errorln("服务器关闭,错误")
	}

	//退出
	log.Infoln("服务器关闭")
	os.Exit(0)
}

func (gs *GameServer) startTCPServer() (err error) {
	log.Infoln("服务器正在tcp服务")
	serverIdFile := gs.options.Server.Server
	serverId, err := gs.readServer(serverIdFile)
	if err != nil {
		return
	}
	gs.serverId = serverId
	sessionOpener := session.SessionHandlerFunc(gs.game.SessionOpen)
	sessionCloser := session.SessionHandlerFunc(gs.game.SessionClose)
	sessionReceiver := session.HandlerFunc(gs.game.SessionReceive)
	sessionSender := session.HandlerFunc(gs.game.SessionSend)
	handler := tcpsession.NewTCPHandler(sessionOpener, sessionCloser, sessionReceiver, sessionSender)
	host := gs.options.Server.Host
	port := gs.options.Server.Port
	addr := fmt.Sprintf("%s:%d", host, port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	gs.ln = ln
	//注册钩子
	gs.initHook()

	log.WithFields(
		log.Fields{
			"addr": addr,
		}).Infoln("服务器启动tcp服务成功")

	for {
		//TODO nats源码 可能接受失败 需要测试
		conn, err := ln.Accept()
		if err != nil {
			if err == syscall.EINVAL {
				break
			}
			log.WithFields(log.Fields{
				"error": err,
			}).Warn("tcp accept error")
			continue
		}
		tconn := tcpsession.NewTCPConnection(conn)
		go handler.Handle(tconn)
	}
	log.Info("tcp closed")
	return nil
}

//开始监听端口
func (gs *GameServer) startWebsocketServer() (err error) {
	serverIdFile := gs.options.Server.Server
	serverId, err := gs.readServer(serverIdFile)
	if err != nil {
		return
	}
	gs.serverId = serverId
	sessionOpener := session.SessionHandlerFunc(gs.game.SessionOpen)
	sessionCloser := session.SessionHandlerFunc(gs.game.SessionClose)
	sessionReceiver := session.HandlerFunc(gs.game.SessionReceive)
	sessionSender := session.HandlerFunc(gs.game.SessionSend)
	path := gs.options.Server.Path
	http.Handle(path, websocket.Handler(websocketsession.NewWebsocketHandler(sessionOpener, sessionCloser, sessionReceiver, sessionSender).Handle))
	host := gs.options.Server.Host
	port := gs.options.Server.Port
	addr := fmt.Sprintf("%s:%d", host, port)

	//TODO 做性能分析使用
	// debug := gs.options.Server.Debug
	// //启动统计
	// if debug {
	// 	statsPort := gs.options.Server.StatsPort
	// 	statsAddr := fmt.Sprintf("%s:%d", sc.Host, statsPort)
	// 	statsMux := mux.NewRouter()
	// 	statsR := statsMux.PathPrefix("/api/stats").Subrouter()
	// 	qipai.StatsRouter(statsR)
	// 	statsN := negroni.Classic()
	// 	statsN.Use(qipai.SetupQipaiServiceHandler(qps))
	// 	statsN.UseHandler(statsMux)
	// 	go func() {
	// 		statsN.Run(statsAddr)
	// 	}()

	// 	pprofAddr := fmt.Sprintf("%s:%d", sc.Host, 6060)
	// 	go func() {
	// 		http.ListenAndServe(pprofAddr, nil)
	// 	}()
	// }

	log.WithFields(log.Fields{}).Infoln("服务器启动")
	//注册钩子
	gs.initHook()
	err = http.ListenAndServe(addr, nil)
	return
}

func (gs *GameServer) initHook() {
	hook := osutils.NewInterruptHooker()
	hook.AddHandler(osutils.InterruptHandlerFunc(stop))
	go func() {
		hook.Run()
	}()
}

func (gs *GameServer) readServer(serverIdFile string) (serverId string, err error) {

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
	gs *GameServer
)

//启动游戏服务器
func InitGameServer(options *GameServerOptions) (tgs *GameServer, err error) {
	log.Infoln("正在初始化游戏服务器")

	if gs != nil {
		panic("repeat init game server")
	}

	tgs = &GameServer{
		options: options,
		game:    NewGame(options.Server, options.Game),
	}
	err = tgs.init()
	if err != nil {
		return
	}
	gs = tgs
	log.Infoln("游戏服务器初始化成功")

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
