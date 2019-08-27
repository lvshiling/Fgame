package cli

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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
	app.Name = "template"
	app.Usage = "template [global options] command [command options] [arguments...]."

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
}

type Class struct {
	File      string      `toml:"file"`
	Class     string      `toml:"class"`
	Comment   string      `toml:"comment"`
	Variables []*Variable `toml:"variables"`
	Pkg       string      `toml:"pkg"`
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
			}).Fatalln("生成模板文件失败")
	}
	log.WithFields(log.Fields{
		"dir": absolultDir,
	}).Infoln("生成模板文件成功")

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

	err = t.Execute(buff, sc)
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

const strTmp = "/*此类自动生成,请勿修改*/\n" +
	"package {{.Pkg}}\n" +
	"/*{{.Comment}}*/\n" +
	"type {{.Class}} struct{\n" +
	"	{{range $variable :=.Variables}}\n" +
	"	//{{$variable.Comment}}\n" +
	"	{{$variable.Name}} {{$variable.Type}} `json:" + `"{{$variable.Tag}}"` + "`\n" +
	"	{{end}}\n" +
	"}"
