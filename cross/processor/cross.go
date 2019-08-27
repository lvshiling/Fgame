package processor

import (
	"fgame/fgame/common/codec"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
	crosscodec "fgame/fgame/cross/codec"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/game/global"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

//处理器
type CrossProcessor struct {
}

//处理外部消息
func (p *CrossProcessor) Process(s session.Session, msgBytes []byte) (err error) {
	//解析
	msg, err := crosscodec.Decode(msgBytes)
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
		//判断是否是跨服登陆消息
		if !codec.IsCrossLoginMsg(msg.MessageType) {
			log.WithFields(
				log.Fields{
					"sessionId": s.Id(),
					"msgType":   msg.MessageType,
				}).Warn("processor:还未登陆,不能处理消息")
			return
		}
		global.GetGame().GetGlobalRunner().Post(smsg)
		return
	}
	//用户在场景内了
	pl := tpl.(*player.Player)
	//判断玩家是否可以处理当前消息
	// if !pl.CanProcess(smsg) {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"sessionId": s.Id(),
	// 			"playerId":  pl.GetId(),
	// 			"state":     pl.CurrentState(),
	// 			"msgType":   msg.MessageType,
	// 		}).Warn("processor:不能处理消息")
	// 	return
	// }
	sc := pl.GetScene()
	if sc != nil {
		pl.Post(smsg)
	} else {
		//用户不在场景内
		//放入全局队列
		global.GetGame().GetGlobalRunner().Post(smsg)
	}

	return
}

//处理内部消息
func (p *CrossProcessor) ProcessInternal(msg message.ScheduleMessage) {
	ps := gamesession.SessionInContext(msg.Context())
	if ps == nil {
		//全局数据
		global.GetGame().GetGlobalRunner().Post(msg)
		return
	}
	//玩家数据
	tpl := ps.Player()
	pl := tpl.(*player.Player)
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
	cp *CrossProcessor
)

func init() {
	cp = &CrossProcessor{}
}

func GetMessageCrossProcessor() *CrossProcessor {
	return cp
}
