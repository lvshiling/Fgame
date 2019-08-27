package template

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"fgame/fgame/core/utils"

	log "github.com/Sirupsen/logrus"
)

const (
	jsonExt   = ".json"
	questFile = `tb_quest.json`
)

type TemplateService interface {
	Init(mkdir string) error //初始化
	GetAllQuest() map[int]*QuestTemplateVO
	GetMainQuest() []*QuestTemplateVO //获取主线剧情
}

type templateService struct {
	questMap       map[int]*QuestTemplateVO
	mainQuestArray []*QuestTemplateVO
}

func (m *templateService) Init(dir string) error {
	absolultDir, err := filepath.Abs(dir)
	if err != nil {
		return nil
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
		fileName := info.Name()
		if fileName == questFile {
			rdErr := m.initQuest(path)
			return rdErr
		}
		return nil
	})
	return nil
}

func (m *templateService) GetAllQuest() map[int]*QuestTemplateVO {
	return m.questMap
}

func (m *templateService) GetMainQuest() []*QuestTemplateVO {
	return m.mainQuestArray
}

func (m *templateService) initQuest(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	tempArray := make([]*QuestTemplateVO, 0)
	err = json.Unmarshal(content, &tempArray)
	if err != nil {
		return err
	}
	firstQuest := &QuestTemplateVO{}
	m.questMap = make(map[int]*QuestTemplateVO)
	for _, value := range tempArray {
		m.questMap[value.Id] = value
		if value.Id == 1 {
			firstQuest = value
		}
	}

	m.mainQuestArray = append(m.mainQuestArray, firstQuest)
	tempQuest := firstQuest
	for {
		if len(tempQuest.NextQuest) == 0 {
			break
		}
		nextArray, err := utils.SplitAsIntArray(tempQuest.NextQuest)
		if err != nil {
			return err
		}
		for _, value := range nextArray {
			mainQuest, exists := m.questMap[int(value)]
			if !exists {
				continue
			}
			hasMain := false
			if mainQuest.IsOnceQuest() {
				m.mainQuestArray = append(m.mainQuestArray, mainQuest)
				tempQuest = mainQuest
				hasMain = true
				break
			}
			if !hasMain {
				break
			}
		}
	}

	return nil
}

var (
	tempService TemplateService = &templateService{}
)

func GetGmTemplateService() TemplateService {
	return tempService
}
