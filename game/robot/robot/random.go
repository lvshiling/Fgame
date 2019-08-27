package robot

import (
	"fgame/fgame/game/fashion/fashion"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	"fgame/fgame/game/scene/scene"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	"fgame/fgame/game/title/title"
	"fgame/fgame/game/weapon/weapon"
	"fgame/fgame/game/wing/wing"
	"time"
)

const (
	randomTaskTime = 10 * time.Second
)

type RandomRobotTask struct {
	p scene.RobotPlayer
}

func (t *RandomRobotTask) Run() {

	//随机时装
	randomFashionTemplate := fashion.GetFashionService().RandomFashionTemplate()
	fashionId := int32(randomFashionTemplate.TemplateId())
	t.p.SetFashionId(fashionId)
	weaponTemplate := weapon.GetWeaponService().RandomWeaponTemplate()
	weaponId := int32(weaponTemplate.TemplateId())

	weaponState := int32(0)
	if weaponTemplate.IsAwaken > 0 {
		weaponState = int32(1)
	}
	t.p.SetWeapon(weaponId, weaponState)
	titleTemplate := title.GetTitleService().RandomTitleTemplate()
	titleId := int32(titleTemplate.TemplateId())
	t.p.SetTitleId(titleId)
	wingTemplate := wing.GetWingService().RandomWingTemplate()
	wingId := int32(wingTemplate.TemplateId())
	t.p.SetWingId(wingId)

	lingYuTemplate := lingyutemplate.GetLingyuTemplateService().RandomLingYuTemplate()
	lingYuId := int32(lingYuTemplate.TemplateId())
	t.p.SetLingYuId(lingYuId)

	shenFaTemplate := shenfatemplate.GetShenfaTemplateService().RandomShenFaTemplate()
	shenFaId := int32(shenFaTemplate.TemplateId())
	t.p.SetShenFaId(shenFaId)

}

func (t *RandomRobotTask) ElapseTime() time.Duration {
	return randomTaskTime
}

func CreateRandomRobotTask(p scene.RobotPlayer) *RandomRobotTask {
	t := &RandomRobotTask{
		p: p,
	}
	return t
}
