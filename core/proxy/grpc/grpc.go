package grpc

import (
	"fgame/fgame/core/proxy"
	proxypb "fgame/fgame/core/proxy/grpc/pb"
	"fgame/fgame/core/session"
	"io"

	log "github.com/Sirupsen/logrus"
)

type GrpcProxy struct {
	client proxypb.Proxy_ForwardClient
	//sesion
	session session.Session
}

func (gp *GrpcProxy) Session() session.Session {
	return gp.session
}

func (gp *GrpcProxy) Forword(msg []byte) (err error) {
	log.WithFields(log.Fields{
		"msg": msg,
	}).Debug("代理转发消息")

	pMsg := &proxypb.Message{}
	pMsg.Body = msg
	//TODO 重新连接
	err = gp.client.Send(pMsg)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("代理转发消息错误")
		return
	}
	return
}

//接收信息
func (gp *GrpcProxy) reverse() {
	go func() {
		for {
			//转发失败
			msg, err := gp.client.Recv()
			if err != nil {
				if err == io.EOF {
					//TODO 客户端断开
					break
				}
				//TODO 读取错误
				break
			}
			//发送session
			msgBytes := msg.Body
			err = gp.session.Send(msgBytes)
			if err != nil {
				//TODO 转发失败
				break
			}
		}
	}()
}

//session 关闭
func (gp *GrpcProxy) onSessionClose() error {
	return nil
}

func NewGrpcProxy(s session.Session) proxy.Proxy {
	gp := &GrpcProxy{}
	gp.session = s

	return nil
}
