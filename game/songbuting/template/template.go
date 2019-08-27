package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type SongBuTingTemplateService interface {
	GetSongBuTingTemplate() *gametemplate.YuanBaoSongBuTingTemplate
}

type songBuTingTemplateService struct {
	songBuTingTemplate *gametemplate.YuanBaoSongBuTingTemplate
}

func (ss *songBuTingTemplateService) init() error {
	songBuTingTemplateList := template.GetTemplateService().GetAll((*gametemplate.YuanBaoSongBuTingTemplate)(nil))

	for _, to := range songBuTingTemplateList {
		ss.songBuTingTemplate = to.(*gametemplate.YuanBaoSongBuTingTemplate)
		break
	}
	return nil
}

func (ss *songBuTingTemplateService) GetSongBuTingTemplate() *gametemplate.YuanBaoSongBuTingTemplate {
	return ss.songBuTingTemplate
}

var (
	once sync.Once
	cs   *songBuTingTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &songBuTingTemplateService{}
		err = cs.init()
	})
	return err
}

func GetSongBuTingTemplateService() SongBuTingTemplateService {
	return cs
}
