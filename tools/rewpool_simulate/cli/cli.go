package cli

import (
	"encoding/json"
	commonlogic "fgame/fgame/game/common/logic"
	commontypes "fgame/fgame/game/common/types"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	debug               = false
	config              = ""
	position            = 0
	times               = 100
	poolLen             = 20
	maxBackTimesForward = 10
	defaultBackTimes    = 0
)

func Start() {
	app := cli.NewApp()
	app.Name = "jie_simulate"
	app.Usage = "jie_simulate [global options] command [command options] [arguments...]."

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "debug",
			Destination: &debug,
		},
		cli.StringFlag{
			Name:        "config",
			Value:       config,
			Usage:       "config",
			Destination: &config,
		},
		cli.IntFlag{
			Name:        "position",
			Value:       position,
			Usage:       "position",
			Destination: &position,
		},
		cli.IntFlag{
			Name:        "times",
			Value:       times,
			Usage:       "times",
			Destination: &times,
		},
		cli.IntFlag{
			Name:        "maxBackTimesForward",
			Value:       maxBackTimesForward,
			Usage:       "maxBackTimes",
			Destination: &maxBackTimesForward,
		},
		cli.IntFlag{
			Name:        "defaultBackTimes",
			Value:       defaultBackTimes,
			Usage:       "defaultBackTimes",
			Destination: &defaultBackTimes,
		},
	}

	app.Action = start
	app.Run(os.Args)
}

type RewNode struct {
	Level       int32 `json:"level"`
	ForwardRate int32 `json:"percent1"`
	BackRate    int32 `json:"percent2"`
	StillRate   int32 `json:"rate"`
}

var (
	PoolMap = make(map[int32]*RewNode)
	PoolArr = make([]*RewNode, 0, poolLen)
)

func start(c *cli.Context) {
	rand.Seed(time.Now().UnixNano())
	tempConfig, err := filepath.Abs(config)
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Fatalln("获取配置目录失败")
	}

	content, err := ioutil.ReadFile(tempConfig)
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Fatalln("读取配置,失败")
	}

	err = json.Unmarshal(content, &PoolArr)
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Fatalln("解析配置,失败")
	}
	for _, node := range PoolArr {
		PoolMap[node.Level] = node
	}
	fmt.Printf("正在模拟奖池抽奖起始位置[%d],总次数[%d],最大回退次数[%d]后前进,起始回退次数[%d]\n", position, times, maxBackTimesForward, defaultBackTimes)
	err = simulateTimes(int32(position), int32(times), int32(maxBackTimesForward), int32(defaultBackTimes))
	if err != nil {
		fmt.Println("模拟失败,参数错误,%s", err)
		return
	}
}

func simulateTimes(position int32, times int32, maxBackTimesForward int32, defaultBackTimes int32) (err error) {
	rewlist := []*commontypes.RewNode{}
	recordMap := make(map[int32]int32)
	for _, node := range PoolMap {
		rewNode := commontypes.CreateRewNode(node.Level, node.ForwardRate, node.BackRate, node.StillRate)
		rewlist = append(rewlist, rewNode)
		recordMap[node.Level] = 0
	}
	rewpools := commontypes.CreateRewPools(rewlist)
	forwardTimes, backTimes, path, err := commonlogic.SimulateRewPools(rewpools, position, times, defaultBackTimes, maxBackTimesForward)
	pos := position
	for _, p := range path {
		switch p {
		case commontypes.PathTypeBack:
			recordMap[pos]++
			pos--
		case commontypes.PathTypeForward:
			recordMap[pos]++
			pos++
		case commontypes.PathTypeBackTimesEnoughForward:
			recordMap[pos]++
			pos++
		case commontypes.PathTypeStill:
			recordMap[pos]++
		}
	}
	if err != nil {
		return
	}
	fmt.Println("总前进次数:", forwardTimes)
	fmt.Println("总后退次数:", backTimes)
	for pos := int32(0); pos < int32(len(recordMap)); pos++ {
		fmt.Println(fmt.Sprintf("位置[%d]总抽奖次数[%d]:", pos, recordMap[pos]))
	}
	return
}
