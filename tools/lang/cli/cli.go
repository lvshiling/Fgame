package cli

import (
	"encoding/json"
	"fgame/fgame/common/lang"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	log "github.com/Sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"github.com/urfave/cli"
)

var (
	debug      = false
	sourceFile = ""
	destFile   = "./export/lang.xlsx"
)

func Start() {
	app := cli.NewApp()
	app.Name = "lang"
	app.Usage = "lang [global options] command [command options] [arguments...]."

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "debug",
			Destination: &debug,
		},
		cli.StringFlag{
			Name:        "sourceFile",
			Value:       sourceFile,
			Usage:       "source file",
			Destination: &sourceFile,
		},
		cli.StringFlag{
			Name:        "destFile",
			Value:       destFile,
			Usage:       "dest file",
			Destination: &destFile,
		},
	}

	app.Action = start
	app.Run(os.Args)
}

type Lang struct {
	Id        int32  `json:"id"`
	ServerMsg string `json:"server_msg"`
	ClientMsg string `json:"client_msg"`
}

func start(c *cli.Context) {
	langList := make([]*Lang, 0, 8)
	if len(sourceFile) != 0 {
		tLangList, err := readSource(sourceFile)
		if err != nil {
			log.WithFields(
				log.Fields{
					"error": err,
				}).Fatalln("读取文件失败")

		}
		langList = tLangList
	}
	langMap := make(map[int32]*Lang)
	for _, lang := range langList {
		langMap[lang.Id] = lang
	}
	//添加不存在的语言
	for code, str := range lang.GetLangMap() {
		_, ok := langMap[int32(code)]
		if ok {
			continue
		}
		langMap[int32(code)] = &Lang{
			Id:        int32(code),
			ServerMsg: str,
		}
	}
	err := writeDest(destFile, langMap)
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Fatalln("写入文件失败")
	}
}

func readSource(sourceFile string) (langList []*Lang, err error) {
	sourceFile, err = filepath.Abs(sourceFile)
	if err != nil {
		return

	}
	content, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return

	}
	langList = make([]*Lang, 0, 8)
	err = json.Unmarshal(content, &langList)
	if err != nil {
		return
	}
	return
}

func writeDest(destFile string, langMap map[int32]*Lang) (err error) {
	destFile, err = filepath.Abs(destFile)
	if err != nil {
		return
	}
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("tb_lang")
	if err != nil {
		return err
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "id"
	cell = row.AddCell()
	cell.Value = "server_msg"
	cell = row.AddCell()
	cell.Value = "client_msg"

	ids := make([]int, 0, 8)
	for code, _ := range langMap {
		ids = append(ids, int(code))
	}

	sort.Sort(sort.IntSlice(ids))
	for _, id := range ids {
		lang := langMap[int32(id)]
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.SetInt(id)
		cell = row.AddCell()
		cell.Value = lang.ServerMsg
		cell = row.AddCell()
		cell.Value = lang.ClientMsg
	}
	err = file.Save(destFile)
	if err != nil {
		return err
	}

	return
}
