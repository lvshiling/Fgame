package cli

import (
	"encoding/json"
	"fgame/fgame/rank/server"
	"io/ioutil"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func readRankServerConfig(configFile string) (sc *server.RankGameServerOptions, err error) {
	log.Infoln("正在读取配置")
	abs, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, err
	}
	sc = &server.RankGameServerOptions{}
	if err = json.Unmarshal(bs, sc); err != nil {
		return nil, err
	}
	log.Infoln("读取配置成功")
	return sc, nil
}

func rankServerStart(c *cli.Context) {
	log.Infoln("正在启动排行榜服务器")

	sc, err := readRankServerConfig(configFile)
	if err != nil {
		log.Fatalln("读取排行榜服务器配置失败", err)
	}

	gs, err := server.InitRankGameServer(sc)
	if err != nil {
		log.Fatalln("初始化排行榜服务器失败", err)
	}
	err = gs.Start()
	if err != nil {
		log.Fatalln("启动排行榜服务器", err)
	}
}
