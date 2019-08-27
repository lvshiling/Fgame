package template

import (
	"fgame/fgame/core/template"
	chattypes "fgame/fgame/game/chat/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sync"
)

// 聊天常量配置
type ChatConstantService interface {
	GetChatConstant(ct chattypes.ChatConstantType) string
	IfRedirect(msg string) bool
}

//快捷缓存
//常量配置的整合
type chatConstantService struct {
	//常量表
	constantTemplateMap map[chattypes.ChatConstantType]*gametemplate.ChatConstantTemplate
	//屏蔽
	houTaiMsgMap map[string]*gametemplate.HoutaiMsgTemplate
}

func (cs *chatConstantService) init() (err error) {
	cs.constantTemplateMap = make(map[chattypes.ChatConstantType]*gametemplate.ChatConstantTemplate)

	templateChatConstantMap := template.GetTemplateService().GetAll((*gametemplate.ChatConstantTemplate)(nil))
	for _, templateObject := range templateChatConstantMap {
		tempChatConstant, _ := templateObject.(*gametemplate.ChatConstantTemplate)
		cs.constantTemplateMap[tempChatConstant.GetConstantType()] = tempChatConstant
	}

	for ct := chattypes.ChatConstantTypeMin; ct <= chattypes.ChatConstantTypeMax; ct++ {
		_, exist := cs.constantTemplateMap[ct]
		if !exist {
			return fmt.Errorf("constant:%d no exist", ct)
		}
	}
	cs.houTaiMsgMap = make(map[string]*gametemplate.HoutaiMsgTemplate)
	templateHoutaiMsgConstantMap := template.GetTemplateService().GetAll((*gametemplate.HoutaiMsgTemplate)(nil))
	for _, templateObject := range templateHoutaiMsgConstantMap {
		tempHoutaiMsgTemplate, _ := templateObject.(*gametemplate.HoutaiMsgTemplate)
		cs.houTaiMsgMap[tempHoutaiMsgTemplate.MsgText] = tempHoutaiMsgTemplate
	}
	return nil
}

func (cs *chatConstantService) GetChatConstant(ct chattypes.ChatConstantType) (val string) {
	tem, exist := cs.constantTemplateMap[ct]
	if !exist {
		return ""
	}
	return tem.Value
}

func (cs *chatConstantService) IfRedirect(msg string) bool {
	_, ok := cs.houTaiMsgMap[msg]
	if !ok {
		return true
	}
	return false
}

var (
	once sync.Once
	cs   *chatConstantService
)

func Init() (err error) {
	once.Do(func() {
		cs = &chatConstantService{}
		err = cs.init()
	})
	return err
}

func GetChatConstantService() ChatConstantService {
	return cs
}
