package cli

import (
	"encoding/json"
	"fgame/fgame/game/server"
	"io/ioutil"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func readServerConfig(configFile string) (sc *server.GameServerOptions, err error) {
	log.Infoln("正在读取配置")
	abs, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, err
	}
	sc = &server.GameServerOptions{}
	if err = json.Unmarshal(bs, sc); err != nil {
		return nil, err
	}
	log.Infoln("读取配置成功")
	return sc, nil
}

func serverStart(c *cli.Context) {
	log.Infoln("正在启动服务器")

	sc, err := readServerConfig(configFile)
	if err != nil {
		log.Fatalln("读取服务器配置失败", err)
	}

	gs, err := server.InitGameServer(sc)
	if err != nil {
		log.Fatalln("初始化服务器失败", err)
	}
	err = gs.Start()
	if err != nil {
		log.Fatalln("启动服务器", err)
	}
}
