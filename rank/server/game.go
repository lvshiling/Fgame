package server

import (
	fgamedb "fgame/fgame/core/db"
	"fgame/fgame/core/module"
	fgameredis "fgame/fgame/core/redis"

	"fgame/fgame/rank/api"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
)

//游戏配置
type RankGameOptions struct {
	//各种服务的配置
	Db    *fgamedb.DbConfig       `json:"db"`
	Redis *fgameredis.RedisConfig `json:"redis"`
}

type RankGame struct {
	serverOptions *RankServerOptions
	//配置选项
	options *RankGameOptions
	//数据库
	db fgamedb.DBService
	// 缓存
	rs fgameredis.RedisService
}

func (g *RankGame) init() error {
	log.WithFields(
		log.Fields{}).Infoln("排行榜服务器游戏数据初始化")

	//初始化数据库
	err := g.initDb()
	if err != nil {
		return err
	}
	//初始化redis
	err = g.initRedis()
	if err != nil {
		return err
	}

	err = g.initModule()
	if err != nil {
		return err
	}

	return nil
}

func (g *RankGame) initDb() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("排行榜服务器游戏正在初始化数据库")
	g.db, err = fgamedb.NewDBService(g.options.Db)
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{}).Infoln("排行榜服务器游戏初始化数据库成功")
	return
}

func (g *RankGame) initRedis() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("排行榜服务器游戏正在初始化redis")
	g.rs, err = fgameredis.NewRedisService(g.options.Redis)
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{}).Infoln("排行榜服务器游戏初始化redis成功")
	return
}

func (g *RankGame) initModule() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化模块")
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
		log.Fields{}).Infoln("游戏正在初始化模块成功")
	return
}

func (g *RankGame) GetDB() fgamedb.DBService {
	return g.db
}

func (g *RankGame) GetRedisService() fgameredis.RedisService {
	return g.rs
}

func (g *RankGame) GetOptions() *RankGameOptions {
	return g.options
}

func (g *RankGame) GetServerIndex() int32 {
	return g.serverOptions.Id
}

func (g *RankGame) GetServerIp() string {
	return g.serverOptions.Host
}

func (g *RankGame) GetServerPort() int32 {
	return g.serverOptions.Port
}

func (g *RankGame) server() *grpc.Server {
	s := grpc.NewServer()
	//注册排行榜服务器查询
	api.Server(s)
	return s
}

func NewRankGame(serverOptions *RankServerOptions, options *RankGameOptions) *RankGame {
	g := &RankGame{
		serverOptions: serverOptions,
		options:       options,
	}
	return g
}
