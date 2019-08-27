package processor

import (
	"fgame/fgame/common/codec"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
	gamecodec "fgame/fgame/game/codec"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

//TODO 统计 开关
//处理器
type Processor struct {
	h message.Handler
}

//处理跨服消息
func (p *Processor) ProcessCross(s session.Session, msgBytes []byte) (err error) {
	//解析
	msg, err := gamecodec.GetCodec().Decode(msgBytes)
	if err != nil {
		return err
	}
	log.WithFields(
		log.Fields{
			"sessionId": s.Id(),
			"msgType":   msg.MessageType,
		}).Debug("对话处理器,接收跨服消息")

	ps := gamesession.SessionInContext(s.Context())
	smsg := message.NewCrossSessionMessage(s, msg)
	//用户在场景内了
	tpl := ps.Player()
	pl := tpl.(player.Player)
	pl.Post(smsg)

	return
}

//处理外部消息
func (p *Processor) Process(s session.Session, msgBytes []byte) (err error) {
	//解析
	msg, err := gamecodec.GetCodec().Decode(msgBytes)
	if err != nil {
		return err
	}
	log.WithFields(
		log.Fields{
			"sessionId": s.Id(),
			"msgType":   msg.MessageType,
		}).Debug("对话处理器,接收消息")

	ps := gamesession.SessionInContext(s.Context())
	smsg := message.NewSessionMessage(s, msg)
	//还未登陆
	tpl := ps.Player()
	if tpl == nil {
		if !codec.IsLoginMsg(msg.MessageType) {
			log.WithFields(
				log.Fields{
					"sessionId": s.Id(),
					"msgType":   msg.MessageType,
				}).Warn("processor:还未登陆,不能处理消息")
			return
		}
		err = p.h.HandleMessage(smsg)
		if err != nil {
			return
		}
		return
	}
	//用户在场景内了
	pl := tpl.(player.Player)
	//判断玩家是否可以处理当前消息
	if !pl.CanProcess(smsg) {
		log.WithFields(
			log.Fields{
				"sessionId": s.Id(),
				"playerId":  pl.GetId(),
				"state":     pl.CurrentState(),
				"msgType":   msg.MessageType,
			}).Warn("processor:不能处理消息")
		return
	}
	//跨服
	if pl.IsCross() {
		pl.Post(smsg)
	} else {
		sc := pl.GetScene()
		if sc != nil {
			pl.Post(smsg)
		} else {
			//用户不在场景内
			//放入全局队列
			global.GetGame().GetGlobalRunner().Post(smsg)
		}
	}
	return
}

//处理内部消息
func (p *Processor) ProcessInternal(msg message.ScheduleMessage) {

	ps := gamesession.SessionInContext(msg.Context())
	if ps == nil {
		//全局数据
		global.GetGame().GetGlobalRunner().Post(msg)
		return
	}
	//玩家数据
	tpl := ps.Player()
	pl := tpl.(player.Player)
	sc := pl.GetScene()
	if sc != nil {
		pl.Post(msg)
	} else {
		//用户不在场景内
		//放入全局队列
		global.GetGame().GetGlobalRunner().Post(msg)
	}
}

var (
	p *Processor
)

func InitProcessor(mh message.Handler) {
	p = &Processor{
		h: mh,
	}
}

func GetMessageProcessor() *Processor {
	return p
}
