package cli

import (
	"encoding/json"
	"fgame/fgame/pkg/mathutils"
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
	debug  = false
	config = ""
	jieShu = 1
	times  = 1
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
			Name:        "jie",
			Value:       jieShu,
			Usage:       "jie",
			Destination: &jieShu,
		},
		cli.IntFlag{
			Name:        "times",
			Value:       times,
			Usage:       "times",
			Destination: &times,
		},
	}

	app.Action = start
	app.Run(os.Args)
}

type Jie struct {
	Id         int32 `json:"id"`
	Rate       int32 `json:"rate"`
	HuituiRate int32 `json:"huitui_rate"`
}

var (
	jieMap = make(map[int32]*Jie)
	jieArr = make([]*Jie, 0, 16)
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

	err = json.Unmarshal(content, &jieArr)
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Fatalln("解析配置,失败")
	}
	for _, jie := range jieArr {
		jieMap[jie.Id] = jie
	}
	fmt.Printf("正在模拟升阶[%d],总次数[%d]\n", jieShu, times)
	flag, totalNumMap := simulateTimes(int32(jieShu), int32(times))
	if !flag {
		fmt.Println("模拟失败,参数错误")
		return
	}
	// totalNum := int32(0)
	// for _, num := range totalNumMap {
	// 	totalNum += num
	// }
	for i := 1; i < jieShu; i++ {
		num := totalNumMap[int32(i)]
		// avg := float64(num) / float64(totalNum)
		fmt.Printf("[%d]阶升级到[%d]阶,升阶平均值[%d]\n", i, i+1, num)
	}

}

func simulateTimes(jie int32, times int32) (flag bool, totalNumMap map[int32]int32) {

	totalNumMap = make(map[int32]int32)
	for i := 0; i < int(times); i++ {
		flag, numMap := simulate(jie)
		if !flag {
			return false, totalNumMap
		}
		for j, num := range numMap {
			totalNumMap[j] += int32(num)
		}
	}
	// avg = float64(totalNum) / float64(times)
	flag = true
	return
}

func simulate(jie int32) (flag bool, numMap map[int32]int32) {
	_, ok := jieMap[jie]
	if !ok {
		return
	}
	numMap = make(map[int32]int32)

	i := int32(1)
	for i < jie {
		numMap[i] += 1
		nextJie := jieMap[i+1]
		jinJie := mathutils.RandomHit(10000, int(nextJie.Rate))
		if jinJie {
			i++
			continue
		}
		currentJie := jieMap[i]
		huiTui := mathutils.RandomHit(10000, int(currentJie.HuituiRate))
		if huiTui {
			i--
		}
	}
	flag = true
	return
}
