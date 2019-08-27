package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gametemplate "fgame/fgame/game/template"
	"strconv"
	"sync"
)

//境界接口处理
type RealmTemplateService interface {
	//获取天劫塔配置通过等级
	GetTianJieTaTemplateByLevel(level int32) *gametemplate.TianJieTaTemplate
	//夫妻助战邀请cd时间
	GetInvitePairCdTime() int64
	//获取夫妻助战cd时间
	GetInvitePairCdTimeStr() string
	//获取技能
	GetSkillId(level int32) (skillId int32)
	//获取天劫塔补偿配置通过等级
	GetTianJieTaBuChangTemplateByLevel(level int32) *gametemplate.TianJieTaBuChangTemplate
}

type realmTemplateService struct {
	//天劫塔
	tianJieTaTemplateMap        map[int32]*gametemplate.TianJieTaTemplate
	cdTime                      string
	tianJieTaBuChangTemplateMap map[int32]*gametemplate.TianJieTaBuChangTemplate
}

//初始化
func (rs *realmTemplateService) init() (err error) {
	rs.tianJieTaTemplateMap = make(map[int32]*gametemplate.TianJieTaTemplate)
	rs.tianJieTaBuChangTemplateMap = make(map[int32]*gametemplate.TianJieTaBuChangTemplate)

	//赋值tianJieTaTemplateMap
	tjtTemplateMap := template.GetTemplateService().GetAll((*gametemplate.TianJieTaTemplate)(nil))
	for _, templateObject := range tjtTemplateMap {
		tjtTemplate, _ := templateObject.(*gametemplate.TianJieTaTemplate)

		_, ok := rs.tianJieTaTemplateMap[tjtTemplate.Level]
		if !ok {
			rs.tianJieTaTemplateMap[tjtTemplate.Level] = tjtTemplate
		}
	}

	//赋值tianJieTaBuChangTemplateMap
	tjtbcTemplateMap := template.GetTemplateService().GetAll((*gametemplate.TianJieTaBuChangTemplate)(nil))
	for _, templateObject := range tjtbcTemplateMap {
		tjtbcTemplate, _ := templateObject.(*gametemplate.TianJieTaBuChangTemplate)

		_, ok := rs.tianJieTaBuChangTemplateMap[tjtbcTemplate.Number]
		if !ok {
			rs.tianJieTaBuChangTemplateMap[tjtbcTemplate.Number] = tjtbcTemplate
		}
	}

	cdTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeRealmInvitePairCdTime) / 1000
	rs.cdTime = strconv.Itoa(int(cdTime))
	return nil
}

//获取技能
func (rs *realmTemplateService) GetSkillId(level int32) (skillId int32) {
	skillId = int32(0)
	to := rs.GetTianJieTaTemplateByLevel(level)
	if to != nil {
		skillId = to.SkillId
	}
	return
}

//获取天劫塔配置通过等级
func (rs *realmTemplateService) GetTianJieTaTemplateByLevel(level int32) *gametemplate.TianJieTaTemplate {
	to, ok := rs.tianJieTaTemplateMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取天劫塔配置通过等级
func (rs *realmTemplateService) GetTianJieTaBuChangTemplateByLevel(level int32) *gametemplate.TianJieTaBuChangTemplate {
	to, ok := rs.tianJieTaBuChangTemplateMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取夫妻助战cd时间
func (rs *realmTemplateService) GetInvitePairCdTimeStr() string {
	return rs.cdTime
}

//天劫塔助战按钮邀请成功的冷却时间(毫秒)
func (rs *realmTemplateService) GetInvitePairCdTime() int64 {
	return int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeRealmInvitePairCdTime))
}

var (
	once sync.Once
	cs   *realmTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &realmTemplateService{}
		err = cs.init()
	})
	return err
}

func GetRealmTemplateService() RealmTemplateService {
	return cs
}
