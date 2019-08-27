package game

import (
	"context"
	"fgame/fgame/core/template"
	"fgame/fgame/pkg/osutils"
	"fmt"
	"sync"

	log "github.com/Sirupsen/logrus"

	"fgame/fgame/client/client"
	_ "fgame/fgame/client/common"

	_ "fgame/fgame/client/funcopen"

	_ "fgame/fgame/client/login"

	_ "fgame/fgame/client/player/handler"

	_ "fgame/fgame/client/scene"
)

type GameOptions struct {
	//各种服务的配置
	Template string `json:"template"`
	Map      string `json:"map"`
	Robot    int32  `json:"robot"`
	From     int32  `json:"from"`
}

//客户端配置
type ClientOptions struct {
	Game   *GameOptions          `json:"game"`
	Client *client.ClientOptions `json:"client"`
}

type Game struct {
	m         sync.Mutex
	options   *ClientOptions
	wg        sync.WaitGroup
	clientMap map[int64]*client.Client
}

func (g *Game) init() error {
	err := g.initTemplate()
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) start() (err error) {
	robot := int32(0)
	if g.options.Game.Robot > robot {
		robot = g.options.Game.Robot
	}

	for i := 0; i < int(robot); i++ {
		g.wg.Add(1)
		userName := fmt.Sprintf("test_%d", int32(i)+g.options.Game.From)
		go func(index int32, userName string) {
			c, err := client.NewClient(g.options.Client, userName)

			if err != nil {
				log.WithFields(
					log.Fields{
						"index": index,
						"err":   err,
					}).Warn("game:连接错误")
				g.wg.Done()
				return
			}
			g.clientMap[int64(index)] = c
			<-c.Done()
			g.wg.Done()
		}(int32(i), userName)

	}

	g.wg.Wait()
	return
}

func (g *Game) stop() {

	g.m.Lock()
	defer g.m.Unlock()

	for _, c := range g.clientMap {
		c.Close()
	}
}

func (g *Game) initTemplate() (err error) {
	log.WithFields(
		log.Fields{}).Infoln("游戏正在初始化模板数据")
	templateDir := g.options.Game.Template
	mapDir := g.options.Game.Map
	//初始化模板服务
	_, err = template.InitTemplateService(templateDir, mapDir)
	if err != nil {
		return err
	}
	log.WithFields(
		log.Fields{}).Infoln("游戏初始化模板数据成功")
	return
}

func NewGame(options *ClientOptions) *Game {
	g := &Game{
		options:   options,
		clientMap: make(map[int64]*client.Client),
	}
	return g
}

type contextKey string

const (
	gameKey = contextKey("fgame.client.game")
)

func GameInContext(ctx context.Context) *Game {
	g, ok := ctx.Value(gameKey).(*Game)
	if !ok {
		return nil
	}
	return g
}

func WithGame(ctx context.Context, g *Game) context.Context {
	return context.WithValue(ctx, gameKey, g)
}

var (
	g *Game
)

func InitGame(options *ClientOptions) (err error) {
	log.Infoln("游戏客户端初始化")

	tg := NewGame(options)
	err = tg.init()
	if err != nil {
		return
	}
	g = tg
	initHook()
	err = tg.start()
	if err != nil {
		return
	}
	return
}

func initHook() {
	hook := osutils.NewInterruptHooker()
	hook.AddHandler(osutils.InterruptHandlerFunc(stop))
	go func() {
		hook.Run()
	}()
}

func stop() {
	g.stop()
}
