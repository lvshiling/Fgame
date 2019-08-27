package server

import (
	"fgame/fgame/pkg/osutils"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"time"

	

	log "github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
)

//排行榜服务器配置
type RankServerOptions struct {
	Server    string `json:"server"`
	Id        int32  `json:"id"`
	Host      string `json:"host"`
	Port      int32  `json:"port"`
	Path      string `json:"path"`
	PprofPort int    `json:"pprofPort"`
	StatsPort int    `json:"statsPort"`
}

//游戏服务器配置
type RankGameServerOptions struct {
	Debug  bool               `json:"debug"`
	Game   *RankGameOptions   `json:"game"`
	Server *RankServerOptions `json:"server"`
}

type RankGameServer struct {
	serverId string
	options  *RankGameServerOptions
	game     *RankGame
}

func (gs *RankGameServer) init() (err error) {
	err = gs.game.init()
	return
}

//重载
func (gs *RankGameServer) Reload() (err error) {
	return nil
}

func (gs *RankGameServer) Start() (err error) {
	//设置随机种子
	now := time.Now().UnixNano()
	rand.Seed(now)
	//服务器启动
	log.Infoln("排行榜服务器正在启动")
	err = gs.startGrpcServer()
	if err != nil {
		return err
	}
	return
}

func (gs *RankGameServer) Stop() {

	//服务器关闭
	log.Infoln("排行榜服务器正在关闭")

	//退出
	log.Infoln("排行榜服务器关闭")
	os.Exit(0)
}

func (gs *RankGameServer) startGrpcServer() (err error) {
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
		}).Infoln("排行榜服务器启动grpc服务成功")
	s := gs.game.server()

	err = s.Serve(ln)
	return
}

func (gs *RankGameServer) initHook() {
	hook := osutils.NewInterruptHooker()
	hook.AddHandler(osutils.InterruptHandlerFunc(stop))
	go func() {
		hook.Run()
	}()
}

func (gs *RankGameServer) readServer(serverIdFile string) (serverId string, err error) {

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
	gs *RankGameServer
)

//启动排行榜游戏服务器
func InitRankGameServer(options *RankGameServerOptions) (tgs *RankGameServer, err error) {
	log.Infoln("正在初始化排行榜游戏服务器")

	if gs != nil {
		panic("repeat init game server")
	}

	tgs = &RankGameServer{
		options: options,
		game:    NewRankGame(options.Server, options.Game),
	}
	err = tgs.init()
	if err != nil {
		return
	}
	gs = tgs
	log.Infoln("游戏排行榜服务器初始化成功")

	return
}

func stop() {
	gs.Stop()
}
