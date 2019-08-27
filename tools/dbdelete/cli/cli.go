package cli

import (
	"encoding/json"
	merge "fgame/fgame/tools/dbdelete/merge"
	model "fgame/fgame/tools/dbdelete/model"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/rifflock/lfshook"
)

var (
	debug      = false
	configFile = "./config/config.json"
)

type dbDeleteConfig struct {
	FromPlatformId int32 `json:"fromPlatformId"`
	FromServerId   int32 `json:"fromServerId"`
}

type deleteConfig struct {
	DumpPath      string                `json:"dumpPath"`
	MySqlPath     string                `json:"mysqlPath"`
	DBMergeConfig *dbDeleteConfig       `json:"dbDelete"`
	DBServerArray []*model.DBConfigInfo `json:"dbMap"`
}

func readMergeConfig(p_filePath string) (config *deleteConfig, err error) {
	c, err := ioutil.ReadFile(p_filePath)
	if err != nil {
		return nil, err
	}
	sc := &deleteConfig{}
	err = json.Unmarshal(c, sc)
	if err != nil {
		return nil, err
	}
	return sc, nil
}

func Start() {
	app := cli.NewApp()
	app.Name = "dbmerge"
	app.Usage = "dbmerge [global options]"

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config,c",
			Value:       configFile,
			Usage:       "config file",
			Destination: &configFile,
		},
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "debug",
			Destination: &debug,
		},
	}
	app.Before = before
	app.Action = start
	app.Run(os.Args)
}

func before(ctx *cli.Context) error {
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.DebugLevel: "./logs/info.log",
		log.InfoLevel:  "./logs/info.log",
		log.WarnLevel:  "./logs/warn.log",
		log.ErrorLevel: "./logs/error.log",
	}))

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	return nil
}

func start(ctx *cli.Context) {
	config, err := filepath.Abs(configFile)
	if err != nil {
		log.Fatalln("filepath abs failed:", err)
	}

	sc, err := readMergeConfig(config)
	if err != nil {
		log.Fatalln("read config file failed:", err)
	}

	fromPlatFormId := int64(sc.DBMergeConfig.FromPlatformId)
	fromServerId := int(sc.DBMergeConfig.FromServerId)

	dbconfigManage := merge.NewDbConfigManage()
	for _, item := range sc.DBServerArray {
		fmt.Println("resigter ", item.PlatformId, item.ServerId)
		dbconfigManage.RegisterDbConfigInfo(item.PlatformId, item.ServerId, item)
	}

	mers := merge.NewMergeService(sc.DumpPath, sc.MySqlPath)
	cbs := merge.NewCombinService(mers, dbconfigManage)
	err = cbs.CombinService(fromPlatFormId, fromServerId)
	if err != nil {
		log.Fatalln("error of merge db :", err)
	}
}
