package server

import (
	"fgame/fgame/account/global"
	"fgame/fgame/core/session"
	tcpsession "fgame/fgame/core/session/tcp"
	"fgame/fgame/pkg/osutils"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
)

//服务器配置
type ServerOptions struct {
	Debug     bool   `json:"debug"`
	Server    string `json:"server"`
	Host      string `json:"host"`
	Port      int32  `json:"port"`
	Path      string `json:"path"`
	PprofPort int32  `json:"pprofPort"`
	StatsPort int32  `json:"statsPort"`
}

//游戏服务器配置
type AccountServerOptions struct {
	Debug   bool            `json:"debug"`
	Account *AccountOptions `json:"account"`
	Server  *ServerOptions  `json:"server"`
}

type AccountServer struct {
	serverId string
	options  *AccountServerOptions
	account  *Account
	ln       net.Listener
}

func (gs *AccountServer) init() (err error) {
	global.SetupAccount(gs.account)
	err = gs.account.init()
	return
}

func (gs *AccountServer) start() (err error) {
	//设置随机种子
	now := time.Now().UnixNano()
	rand.Seed(now)

	//启动统计服务

	// debug := gs.options.Server.Debug
	// //启动统计
	// if debug {
	// 	host := gs.options.Server.Host
	// 	statsPort := gs.options.Server.StatsPort
	// 	statsAddr := fmt.Sprintf("%s:%d", host, statsPort)
	// 	statsMux := mux.NewRouter()
	// 	statsR := statsMux.PathPrefix("/api/stats").Subrouter()
	// 	gs.game.statsRouter(statsR)
	// 	statsN := negroni.Classic()
	// 	statsN.UseHandler(statsMux)
	// 	go func() {
	// 		statsN.Run(statsAddr)
	// 	}()
	// 	pprofPort := gs.options.Server.PprofPort
	// 	pprofAddr := fmt.Sprintf("%s:%d", host, pprofPort)
	// 	go func() {
	// 		http.ListenAndServe(pprofAddr, nil)
	// 	}()
	// }

	//服务器启动
	log.Infoln("server:服务器正在启动")
	err = gs.account.start()
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

//记录进程id
func (s *AccountServer) logPid() error {
	// pidStr := strconv.Itoa(os.Getpid())
	return nil
	// return ioutil.WriteFile(s.getOpts().PidFile, []byte(pidStr), 0660)
}

func (gs *AccountServer) stop() {
	//服务器关闭
	log.Infoln("server:服务器正在关闭")
	//停止接收消息
	// gs.ln.Close()

	//游戏数据停止
	err := gs.account.stop()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Errorln("server:服务器关闭,错误")
	}

	//退出
	log.Infoln("server:服务器关闭")
	os.Exit(0)
}

func (gs *AccountServer) startTCPServer() (err error) {
	log.Infoln("server:服务器正在tcp服务")
	serverIdFile := gs.options.Server.Server
	serverId, err := gs.readServer(serverIdFile)
	if err != nil {
		return
	}
	gs.serverId = serverId
	sessionOpener := session.SessionHandlerFunc(gs.account.SessionOpen)
	sessionCloser := session.SessionHandlerFunc(gs.account.SessionClose)
	sessionReceiver := session.HandlerFunc(gs.account.SessionReceive)
	sessionSender := session.HandlerFunc(gs.account.SessionSend)
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
		}).Infoln("server:服务器启动tcp服务成功")

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

func (gs *AccountServer) initHook() {
	hook := osutils.NewInterruptHooker()
	hook.AddHandler(osutils.InterruptHandlerFunc(stop))
	go func() {
		hook.Run()
	}()
}

func (gs *AccountServer) readServer(serverIdFile string) (serverId string, err error) {

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
	gs   *AccountServer
)

//启动账户服务器
func InitAccountServer(options *AccountServerOptions) (err error) {
	once.Do(func() {
		log.Info("server:正在初始化账户服务器")
		if gs != nil {
			panic(fmt.Errorf("server:重复初始化账户服务器"))
		}
		gs = &AccountServer{
			options: options,
			account: NewAccount(options.Server, options.Account),
		}
		err = gs.init()
		if err != nil {
			return
		}
		err = gs.start()
		if err != nil {
			return
		}
		return
	})
	return
}

func stop() {
	gs.stop()
}
