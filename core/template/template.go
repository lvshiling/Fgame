package template

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	log "github.com/Sirupsen/logrus"
)

//数据模板 interface
type TemplateObject interface {
	TemplateId() int
	FileName() string
	//检查有效性
	Check() error
	//组合成需要的数据
	Patch() error
	//检验后组合
	PatchAfterCheck()
}

//模板服务
type TemplateService interface {

	//获取模板
	Get(id int, to TemplateObject) (o TemplateObject)
	GetAll(to TemplateObject) (os map[int]TemplateObject)
	ReadMap(mapFile string) (*Map, error)
}

type templateService struct {
	templateMapOfMap map[reflect.Type]map[int]TemplateObject
	mapDir           string
}

func (ts *templateService) check() error {
	for _, tos := range ts.templateMapOfMap {
		for _, to := range tos {
			err := to.Check()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (ts *templateService) patch() error {
	for _, tos := range ts.templateMapOfMap {
		for _, to := range tos {
			err := to.Patch()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (ts *templateService) patchAfterCheck() {
	for _, tos := range ts.templateMapOfMap {
		for _, to := range tos {
			to.PatchAfterCheck()
		}
	}
}

func (ts *templateService) Get(id int, to TemplateObject) (o TemplateObject) {
	typ := reflect.TypeOf(to)
	tos, exist := ts.templateMapOfMap[typ]
	if !exist {
		return nil
	}
	o, exist = tos[id]
	if !exist {
		return nil
	}
	return
}

func (ts *templateService) GetAll(to TemplateObject) (tos map[int]TemplateObject) {
	typ := reflect.TypeOf(to)
	tos, exist := ts.templateMapOfMap[typ]
	if !exist {
		return nil
	}
	return
}

const (
	jsonExt = ".json"
)

//TOOD 优化实现过程
func (ts *templateService) init(dir string, mapDir string) (err error) {
	ts.mapDir = mapDir
	ts.templateMapOfMap = make(map[reflect.Type]map[int]TemplateObject)

	absolultDir, err := filepath.Abs(dir)
	if err != nil {
		return
	}
	log.WithFields(log.Fields{
		"dir": absolultDir,
	}).Infoln("正在初始化模板服务")

	err = filepath.Walk(absolultDir, func(path string, info os.FileInfo, err error) error {
		ext := filepath.Ext(path)
		if !strings.EqualFold(jsonExt, ext) {
			return nil
		}
		log.WithFields(log.Fields{
			"name": info.Name(),
		}).Infoln("读取文件")
		typ, ok := templateObjectMap[info.Name()]
		if !ok {
			log.WithFields(log.Fields{
				"name": info.Name(),
			}).Warnln("结构不存在")
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		sliceType := reflect.SliceOf(typ)

		ls := reflect.MakeSlice(sliceType, 0, 16)
		x := reflect.New(ls.Type())

		err = json.Unmarshal(content, x.Interface())
		if err != nil {
			return err
		}

		_, exist := ts.templateMapOfMap[typ]
		if !exist {
			ts.templateMapOfMap[typ] = make(map[int]TemplateObject)
		}
		ex := x.Elem()
		len := ex.Len()
		for i := 0; i < len; i++ {
			to := ex.Index(i).Interface().(TemplateObject)
			_, exist := ts.templateMapOfMap[typ][to.TemplateId()]
			if exist {
				return fmt.Errorf("%s templateId:%d 应该是唯一的", info.Name(), to.TemplateId())
			}
			ts.templateMapOfMap[typ][to.TemplateId()] = to
		}
		return nil
	})

	if err != nil {
		return
	}

	//组合数据
	if err = ts.patch(); err != nil {
		return
	}

	//检查数据有效性
	if err = ts.check(); err != nil {
		return
	}
	//组合
	ts.patchAfterCheck()
	return
}

var (
	ts   TemplateService
	once sync.Once
)

func GetTemplateService() TemplateService {
	return ts
}

//初始化
func InitTemplateService(dir string, mapDir string) (TemplateService, error) {
	var err error
	once.Do(func() {
		tts := &templateService{}
		//先赋值由于init使用到
		ts = tts

		err = tts.init(dir, mapDir)
		if err != nil {
			return
		}

	})
	return ts, err
}

var (
	templateObjectMap map[string]reflect.Type
)

func init() {
	templateObjectMap = make(map[string]reflect.Type)
}

func Register(to TemplateObject) {
	_, exist := templateObjectMap[to.FileName()]
	if exist {
		panic(fmt.Sprintf("repeate register %s template object", to.FileName()))
	}
	templateObjectMap[to.FileName()] = reflect.TypeOf(to)
}
