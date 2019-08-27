package template

import (
	"fgame/fgame/core/template"
	questtypes "fgame/fgame/game/quest/types"
	"fmt"
)

//运营活动次数配置
type KaiFuMuBiaoQuestTemplate struct {
	*KaiFuMuBiaoQuestTemplateVO
	nextKaiFuMuBiaoQuestTemplate *KaiFuMuBiaoQuestTemplate
}

func (t *KaiFuMuBiaoQuestTemplate) TemplateId() int {
	return t.Id
}

func (t *KaiFuMuBiaoQuestTemplate) PatchAfterCheck() {
}

func (t *KaiFuMuBiaoQuestTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*KaiFuMuBiaoQuestTemplate)(nil))
		if to != nil {
			nextTemplate := to.(*KaiFuMuBiaoQuestTemplate)
			t.nextKaiFuMuBiaoQuestTemplate = nextTemplate
		}
	}

	return nil
}

func (t *KaiFuMuBiaoQuestTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	to := template.GetTemplateService().Get(int(t.Quest), (*QuestTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.Quest)
		err = template.NewTemplateFieldError("Quest", err)
		return
	}
	questTemplate := to.(*QuestTemplate)
	if questTemplate.GetQuestType() != questtypes.QuestTypeKaiFuMuBiao {
		err = fmt.Errorf("[%d] invalid", t.Quest)
		err = template.NewTemplateFieldError("Quest", err)
		return
	}
	if !questTemplate.AutoAccept() {
		err = fmt.Errorf("[%d] invalid", t.Quest)
		err = template.NewTemplateFieldError("Quest", err)
		return
	}

	return nil
}

func (t *KaiFuMuBiaoQuestTemplate) FileName() string {
	return "tb_kaifumubiao_quest.json"
}

func init() {
	template.Register((*KaiFuMuBiaoQuestTemplate)(nil))
}
