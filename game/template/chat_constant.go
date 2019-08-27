package template

import (
	"fgame/fgame/core/template"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

type ChatConstantTemplate struct {
	*ChatConstantTemplateVO
	constantType chattypes.ChatConstantType
}

func (ct *ChatConstantTemplate) GetConstantType() chattypes.ChatConstantType {
	return ct.constantType
}

func (ct *ChatConstantTemplate) TemplateId() int {
	return ct.Id
}

//组合数据
func (ct *ChatConstantTemplate) Patch() (err error) {
	//统一处理错误方式
	defer func() {
		if err != nil {
			err = template.NewTemplateError(ct.FileName(), ct.TemplateId(), err)
			return
		}
	}()
	ct.constantType = chattypes.ChatConstantType(ct.Type)

	return nil
}

//检验数值
func (ct *ChatConstantTemplate) Check() (err error) {
	//统一处理错误方式
	defer func() {
		if err != nil {
			err = template.NewTemplateError(ct.FileName(), ct.TemplateId(), err)
			return
		}
	}()
	switch ct.constantType {
	case chattypes.ChatConstantTypeStopChatStartTime,
		chattypes.ChatConstantTypeStopChatEndTime:
		_, err = timeutils.ParseDayOfHHMM(ct.Value)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", ct.Value)
			return template.NewTemplateFieldError(ct.constantType.String(), err) 
		}
		break
	}

	return
}

func (ct *ChatConstantTemplate) PatchAfterCheck() {

}

func (ct *ChatConstantTemplate) FileName() string {
	return "tb_chat_constant.json"
}

func init() {
	template.Register((*ChatConstantTemplate)(nil))
}
