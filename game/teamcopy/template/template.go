package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	teamtypes "fgame/fgame/game/team/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//组队副本接口处理
type TeamCopyTemplateService interface {
	//获取组队副本模板
	GetTeamCopyTempalte(purpose teamtypes.TeamPurposeType) *gametemplate.TeamCopyTemplate
	//获取出生位置
	GetTeamCopyBorn(purpose teamtypes.TeamPurposeType) (pos coretypes.Position, flag bool)
	//获取组队副本奖励次数
	GetTeamCopyRewardNumber(purpose teamtypes.TeamPurposeType) (num int32)
}

type teamCopyTemplateService struct {
	//组队副本模板
	teamCopyTemplateMap map[teamtypes.TeamPurposeType]*gametemplate.TeamCopyTemplate
}

//初始化
func (ts *teamCopyTemplateService) init() error {
	ts.teamCopyTemplateMap = make(map[teamtypes.TeamPurposeType]*gametemplate.TeamCopyTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.TeamCopyTemplate)(nil))
	for _, templateObject := range templateMap {
		teamCopyTemplate, _ := templateObject.(*gametemplate.TeamCopyTemplate)

		purposeType := teamCopyTemplate.GetPurposeType()
		ts.teamCopyTemplateMap[purposeType] = teamCopyTemplate
	}

	return nil
}

//获取组队副本模板
func (ts *teamCopyTemplateService) GetTeamCopyTempalte(purpose teamtypes.TeamPurposeType) *gametemplate.TeamCopyTemplate {
	return ts.teamCopyTemplateMap[purpose]
}

//获取组队副本奖励次数
func (ts *teamCopyTemplateService) GetTeamCopyRewardNumber(purpose teamtypes.TeamPurposeType) (num int32) {
	temp := ts.GetTeamCopyTempalte(purpose)
	if temp == nil {
		return
	}
	num = temp.RewardNumber
	return
}

//获取出生位置
func (ts *teamCopyTemplateService) GetTeamCopyBorn(purpose teamtypes.TeamPurposeType) (pos coretypes.Position, flag bool) {
	teamCopyTemplate := ts.GetTeamCopyTempalte(purpose)
	if teamCopyTemplate == nil {
		return
	}
	pos = teamCopyTemplate.GetBornPos()
	flag = true
	return
}

var (
	once sync.Once
	cs   *teamCopyTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &teamCopyTemplateService{}
		err = cs.init()
	})
	return err
}

func GetTeamCopyTemplateService() TeamCopyTemplateService {
	return cs
}
