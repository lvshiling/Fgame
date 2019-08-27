package cli

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	debug   = false
	dir     = ""
	destDir = ""
)

func Start() {
	app := cli.NewApp()
	app.Name = "log"
	app.Usage = "log [global options] command [command options] [arguments...]."

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "debug",
			Destination: &debug,
		},
		cli.StringFlag{
			Name:        "dir",
			Value:       dir,
			Usage:       "dir",
			Destination: &dir,
		},
		cli.StringFlag{
			Name:        "destDir",
			Value:       destDir,
			Usage:       "destDir",
			Destination: &destDir,
		},
	}

	app.Action = start
	app.Run(os.Args)
}

type Variable struct {
	Name    string `toml:"name"`
	Type    string `toml:"type"`
	Comment string `toml:"comment"`
	Tag     string `toml:"tag"`
	Show    string `toml:"show"`
}

type Class struct {
	File      string `toml:"file"`
	Class     string `toml:"class"`
	Base      string `toml:"base"`
	BaseTag   string
	Comment   string      `toml:"comment"`
	Variables []*Variable `toml:"variables"`
	Pkg       string      `toml:"pkg"`
	Name      string      `toml:"name"`
	MsgType   string
}

func readClass(configFile string) (sc *Class, err error) {

	abs, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, err
	}
	sc = &Class{}
	if _, err = toml.Decode(string(bs), sc); err != nil {
		return nil, err
	}

	return sc, nil
}

const (
	tomlExt = ".toml"
)

var (
	destAbsoluteDir string
)

func start(c *cli.Context) {
	tempDestAbsoluteDir, err := filepath.Abs(destDir)
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Fatalln("获取目的目录失败")
	}
	destAbsoluteDir = tempDestAbsoluteDir
	absolultDir, err := filepath.Abs(dir)
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Fatalln("获取目录失败")
	}
	log.WithFields(log.Fields{
		"dir": absolultDir,
	}).Infoln("正在读取配置目录")

	err = filepath.Walk(absolultDir, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		ext := filepath.Ext(path)
		if !strings.EqualFold(tomlExt, ext) {
			return nil
		}
		log.WithFields(log.Fields{
			"name": info.Name(),
		}).Infoln("读取文件")

		err = Generate(path)
		if err != nil {
			return err
		}
		log.WithFields(log.Fields{
			"path": path,
		}).Infoln("生成成功")
		return nil
	})
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Fatalln("生成日志文件失败")
	}
	log.WithFields(log.Fields{
		"dir": absolultDir,
	}).Infoln("生成日志文件成功")
}

func convertClass(sc *Class) *Class {
	newSc := &Class{}
	newSc.Class = sc.Class
	newSc.Comment = sc.Comment
	newSc.File = sc.File
	newSc.Base = sc.Base
	newSc.Name = sc.Name
	newSc.Pkg = sc.Pkg
	if newSc.Base == "PlayerLogMsg" {
		newSc.MsgType = "MsgTypePlayerLog"
	} else {
		if newSc.Base == "AllianceLogMsg" {
			newSc.MsgType = "MsgTypeAllianceLog"
		} else {
			if newSc.Base == "PlayerTradeLogMsg" {
				newSc.MsgType = "MsgTypeJiaoYiLog"
			} else {
				newSc.MsgType = "MsgTypeServerLog"
			}
		}
	}
	newSc.Variables = make([]*Variable, 0, len(sc.Variables))
	for _, variable := range sc.Variables {
		newVar := &Variable{}
		newVar.Show = variable.Show
		newVar.Comment = variable.Comment
		spilitComment := strings.Split(variable.Comment, "(")
		if len(spilitComment) > 0 {
			newVar.Comment = spilitComment[0]
		}
		newVar.Name = variable.Name
		newVar.Tag = variable.Tag
		// newTag := `"` + variable.Tag + `"`
		// newVar.Tag = fmt.Sprintf("`json:%s`", newTag)
		newVar.Type = variable.Type
		newSc.Variables = append(newSc.Variables, newVar)
	}
	bsonTag := `bson:",inline"`
	newSc.BaseTag = fmt.Sprintf("`%s`", bsonTag)
	fmt.Println(newSc.BaseTag)
	return newSc
}

func Generate(configFile string) (err error) {

	sc, err := readClass(configFile)
	if err != nil {
		return
	}
	t, err := template.New("").Parse(strTmp)
	if err != nil {
		return
	}
	buff := bytes.NewBufferString("")
	newSc := convertClass(sc)
	err = t.Execute(buff, newSc)
	if err != nil {
		return
	}
	//进行格式化
	src, err := format.Source(buff.Bytes())
	if err != nil {
		return
	}
	destPath := filepath.Join(destAbsoluteDir, sc.File)
	err = ioutil.WriteFile(destPath, src, os.ModePerm)
	if err != nil {
		return
	}
	return
}

// const strTmpold = `
// /*此类自动生成,请勿修改*/
// 	package {{.Pkg}}
// 	import logserverlog "fgame/fgame/logserver/log"
// 	func init(){
// 		logserverlog.RegisterLogMsg((*{{.Class}})(nil))
// 	}
// 	/*{{.Comment}}*/
// 	type {{.Class}} struct{
// 		{{.Base}} {{.BaseTag}}
// 		{{range $variable :=.Variables}}
// 		//{{$variable.Comment}}
// 		{{$variable.Name}} {{$variable.Type}} {{$variable.Tag}}
// 		{{end}}
// 	}
// 	func(c *{{.Class}}) LogName() string{
// 		return "{{.Name}}"
// 	}
// `

const strTmp = `
package msg

import (
	mglog "fgame/fgame/gm/gamegm/gm/mglog/metadata"
)

func init() {
	rst := make([]*mglog.MsgItemInfo, 0)
	var item *mglog.MsgItemInfo

	{{range $variable :=.Variables}}
	item = &mglog.MsgItemInfo{
		Label:      "{{$variable.Comment}}",
		DataColumn: "{{$variable.Tag}}",
		ShowType:   "{{$variable.Show}}",
	}
	rst = append(rst, item)
	{{end}}

	mglog.RegisterMsgItemInfo("{{.Name}}", "{{.Comment}}", mglog.{{.MsgType}}, rst)
}
`
