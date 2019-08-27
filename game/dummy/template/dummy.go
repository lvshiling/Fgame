package template

import (
	"fgame/fgame/core/template"
	dummytypes "fgame/fgame/game/dummy/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	playertypes "fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"math/rand"
	"sync"
)

//假名模板
type DummyTemplateService interface {
	GetRandomDummyName() string //获取随机假名
	GetRandomDummyNameBySex(sexType playertypes.SexType) string
	GetGameRandomDummyName() string
}
type dummyTemplateService struct {
	dummyListMap map[dummytypes.DummyType][]*gametemplate.DummyNameTemplate //假名
}

func (s *dummyTemplateService) init() (err error) {
	s.dummyListMap = make(map[dummytypes.DummyType][]*gametemplate.DummyNameTemplate)

	tmepMap := template.GetTemplateService().GetAll((*gametemplate.DummyNameTemplate)(nil))
	for _, tem := range tmepMap {
		temp, _ := tem.(*gametemplate.DummyNameTemplate)
		s.dummyListMap[temp.GetDummyType()] = append(s.dummyListMap[temp.GetDummyType()], temp)
	}

	return
}

func (dt *dummyTemplateService) GetGameRandomDummyName() string {
	if merge.GetMergeService().GetMergeTime() != 0 {
		serverId := global.GetGame().GetServerIndex()
		sexType := playertypes.RandomSex()
		name := dt.GetRandomDummyNameBySex(sexType)
		return fmt.Sprintf("s%d.%s", serverId, name)
	} else {
		return dt.GetRandomDummyName()
	}
}

func (dt *dummyTemplateService) GetRandomDummyName() string {
	sexType := playertypes.RandomSex()
	return dt.GetRandomDummyNameBySex(sexType)
}

func (dt *dummyTemplateService) GetRandomDummyNameBySex(sexType playertypes.SexType) string {
	surName := dt.getRandomDummyName(dummytypes.DummyTypeSurname)
	nameStr := dt.getRandomDummyName(sexType.DummyType())
	return surName + nameStr
}

//随机假名字
func (dt *dummyTemplateService) getRandomDummyName(dummyType dummytypes.DummyType) string {
	surnameList := dt.dummyListMap[dummyType]
	len := len(surnameList)
	if len == 0 {
		return ""
	}
	randIndex := rand.Intn(len)

	return surnameList[randIndex].Name
}

var (
	once sync.Once
	s    *dummyTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &dummyTemplateService{}
		err = s.init()
		if err != nil {
			return
		}
	})
	return
}

func GetDummyTemplateService() DummyTemplateService {
	return s
}
