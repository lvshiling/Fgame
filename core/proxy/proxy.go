package proxy

import "fgame/fgame/core/session"

type Proxy interface {
	//服务
	Service() string
	//服务id
	Id() string
	//一个代理一个session
	Session() session.Session
	//转发
	Forword(msg []byte) (err error)
}
