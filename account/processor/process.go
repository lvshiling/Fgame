package processor

import (
	"fgame/fgame/common/codec"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
)

var (
	d = dispatch.NewDispatch()
)

func Register(msgType codec.MessageType, h dispatch.Handler) {
	d.Register(msgType, h)
}

func GetDispatch() dispatch.Dispatch {
	return d
}

//TODO 统计 开关
//处理器
type Processor struct {
	h   message.Handler
	cod *codec.Codec
}

//处理外部消息
func (p *Processor) Process(s session.Session, msgBytes []byte) (err error) {
	//解析
	msg, err := p.cod.Decode(msgBytes)
	if err != nil {
		return err
	}

	smsg := message.NewSessionMessage(s, msg)
	p.h.HandleMessage(smsg)
	return
}

var (
	p *Processor
)

func InitProcessor(mh message.Handler, cod *codec.Codec) {
	p = &Processor{
		h:   mh,
		cod: cod,
	}
}

func GetMessageProcessor() *Processor {
	return p
}
