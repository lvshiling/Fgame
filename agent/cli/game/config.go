package game

import (
	"encoding/json"
	agentsetup "fgame/fgame/agent/agent/setup"
	gameserver "fgame/fgame/game/server"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	configFile = ""
)

type GameConfig struct {
	Dir  string `json:"dir"`
	Base string `json:"base"`
}

var (
	configCommand = cli.Command{
		Name:        "config",
		Usage:       "agent [global options] game config [command options] [arguments...].",
		Description: "game agent config",
		Action:      actionConfig,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "config,c",
				Usage:       "config",
				Destination: &configFile,
			},
		},
	}
)

func actionConfig(c *cli.Context) error {
	args := c.Args()
	if len(args) <= 0 {
		return cli.NewExitError("缺少参数", 1)
	}
	serverIds := make([]int32, 0, 8)

	for _, arg := range args {
		serverId, err := strconv.ParseInt(arg, 10, 32)
		if err != nil {
			return cli.NewMultiError(errors.Wrap(err, "参数需要是整数"))
		}
		serverIds = append(serverIds, int32(serverId))
	}

	cfg, err := initConfig(configFile)
	if err != nil {
		return cli.NewMultiError(errors.Wrap(err, "配置读取失败"))
	}
	baseCfg, err := initBaseConfig(cfg.Base)
	if err != nil {
		return cli.NewMultiError(errors.Wrap(err, "基础配置读取失败"))
	}
	for _, serverId := range serverIds {
		err = agentsetup.SetupGameServer(baseCfg, cfg.Dir, serverId)
		if err != nil {
			return cli.NewMultiError(errors.Wrap(err, "设置游戏服配置失败"))
		}
	}

	return nil
}

func initConfig(configFile string) (sc *GameConfig, err error) {
	abs, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, err
	}
	sc = &GameConfig{}
	if err = json.Unmarshal(bs, sc); err != nil {
		return nil, err
	}

	return sc, nil
}

func initBaseConfig(configFile string) (sc *gameserver.GameServerOptions, err error) {

	abs, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, err
	}
	sc = &gameserver.GameServerOptions{}
	if err = json.Unmarshal(bs, sc); err != nil {
		return nil, err
	}

	return sc, nil
}
