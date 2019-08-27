package server

import (
	"fgame/fgame/common/message"
	"fgame/fgame/sdk"

	"fgame/fgame/account/center/center"
	accountcodec "fgame/fgame/account/codec"
	_ "fgame/fgame/account/common"
	commonlogic "fgame/fgame/account/common/logic"
	_ "fgame/fgame/account/login"
	_ "fgame/fgame/account/notice"
	_ "fgame/fgame/account/player"
	"fgame/fgame/account/processor"
	_ "fgame/fgame/account/serverlist"
	accountsession "fgame/fgame/account/session"
	"fgame/fgame/common/lang"
	_ "fgame/fgame/common/lang"
	coremodule "fgame/fgame/core/module"
	"fgame/fgame/core/session"
	_ "fgame/fgame/sdk/sdk"

	log "github.com/Sirupsen/logrus"
)

//账户配置
type AccountOptions struct {
	Center   *center.CenterConfig `json:"center"`
	Sdk      string               `json:"sdk"`
	EnablePC bool                 `json:"enablePC"`
}

type Account struct {
	serverOptions *ServerOptions
	//配置选项
	options *AccountOptions
	//消息处理器
	msgHandler message.Handler
}

func (g *Account) init() error {
	log.WithFields(
		log.Fields{}).Infoln("server:账户数据初始化")
	g.initMessageHandler()

	cod := accountcodec.GetCodec()
	//初始化处理器
	processor.InitProcessor(g.msgHandler, cod)
	//初始化sdk
	err := g.initSDK()
	if err != nil {
		return err
	}
	//初始化中心服务器
	err = g.initCenter()
	if err != nil {
		return err
	}

	//初始化模块服务
	err = g.initModules()
	if err != nil {
		return err
	}

	return nil
}

func (g *Account) initSDK() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("server:初始化sdk")

	err = sdk.Init(g.options.Sdk)
	if err != nil {
		return err
	}
	log.WithFields(
		log.Fields{}).Infoln("server:初始化sdk成功")
	return
}

func (g *Account) initMessageHandler() {
	log.WithFields(
		log.Fields{}).Infoln("server:账户正在初始化消息处理器")

	g.msgHandler = message.HandlerFunc(handleMessage)
	log.WithFields(
		log.Fields{}).Infoln("server:账户初始化消息处理器成功")

	return
}

func (g *Account) initCenter() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("server:账户正在初始化中心")

	err = center.Init(g.options.Center)
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{}).Infoln("server:账户初始化中心")
	return
}

func (g *Account) initModules() (err error) {
	for _, m := range coremodule.GetModules() {
		err = m.InitTemplate()
		if err != nil {
			return
		}
	}
	for _, m := range coremodule.GetModules() {
		err = m.Init()
		if err != nil {
			return
		}
	}
	return
}

func (g *Account) start() (err error) {
	defer func() {
		if err != nil {
			g.stop()
		}
	}()
	//中心服启动
	center.GetCenterService().Start()

	for _, m := range coremodule.GetModules() {
		log.WithFields(
			log.Fields{
				"module": m.String(),
			}).Infoln("server:账户服务器初始化模块")
		m.Start()

	}

	return
}

func (g *Account) stop() (err error) {
	//中心服启动了
	for _, m := range coremodule.GetModules() {
		m.Stop()
		log.WithFields(
			log.Fields{
				"module": m.String(),
			}).Infoln("server:账户服务器停止模块")
	}
	center.GetCenterService().Stop()
	return
}

//对话开启
func (g *Account) SessionOpen(s session.Session) error {
	ps := accountsession.NewPlayerSession(s, accountcodec.GetCodec(), nil)
	nctx := accountsession.WithSession(s.Context(), ps)
	s.SetContext(nctx)
	// 添加统计
	return nil
}

//对话关闭
func (g *Account) SessionClose(s session.Session) error {
	// 添加统计
	ps := accountsession.SessionInContext(s.Context())
	if ps == nil {
		panic("server: 对话应该可以获取")
	}
	//关闭写循环 以防内存泄露
	defer ps.Close(false)

	return nil
}

func (g *Account) SessionReceive(s session.Session, msg []byte) error {
	// 添加统计
	err := processor.GetMessageProcessor().Process(s, msg)
	if err != nil {
		sess := accountsession.SessionInContext(s.Context())
		commonlogic.SendSessionSystemMessage(sess, lang.AccountLoginMessageCanNotHandle)
		return err
	}
	return nil
}

func (g *Account) SessionSend(s session.Session, msg []byte) error {
	// 添加统计
	return nil
}

func (g *Account) EnablePC() bool {
	return g.options.EnablePC
}

//全局对象
func (g *Account) GetMessageHandler() message.Handler {
	return g.msgHandler
}

func NewAccount(serverOptions *ServerOptions, options *AccountOptions) *Account {
	g := &Account{
		serverOptions: serverOptions,
		options:       options,
	}
	return g
}
