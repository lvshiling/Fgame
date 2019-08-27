package cli

import (
	"encoding/json"
	fredis "fgame/fgame/core/redis"
	wxcache "fgame/fgame/gm/gzhdemo/cache"
	wxapi "fgame/fgame/gm/gzhdemo/wxapi"
	wxrouter "fgame/fgame/gm/gzhdemo/wxapi/api"
	"fgame/fgame/pkg/osutils"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	wxpay "fgame/fgame/gm/gzhdemo/wxpay"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rifflock/lfshook"
	wechat "github.com/silenceper/wechat"
)

var (
	debug      = true //暂时先改一下
	configFile = ""
)

func Start() {

	app := cli.NewApp()
	app.Name = "user"
	app.Usage = "user [global options]"

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "debug ",
			Destination: &debug,
			EnvVar:      "USER_DEBUG",
		},
		cli.StringFlag{
			Name:        "config,c",
			Value:       configFile,
			Usage:       "config file",
			Destination: &configFile,
		},
	}

	app.Before = before
	app.Action = start
	app.Run(os.Args)
}

func before(c *cli.Context) error {
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.DebugLevel: "./logs/info.log",
		log.ErrorLevel: "./logs/error.log",
		log.WarnLevel:  "./logs/warn.log",
	}))

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
	return nil
}

type serverConfig struct {
	Cert  string              `json:"cert'`
	Key   string              `json:"key"`
	Host  string              `json:"host"`
	Port  int                 `json:"port"`
	Redis *fredis.RedisConfig `json:"redis"`
	GZH   *gongzhConfig       `json:"gzh"`
	WxPay *wxpay.WxPayConfig  `json:"wxredbagpay"`
}

func newServerConfig(configFile string) (sc *serverConfig, err error) {
	c, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	sc = &serverConfig{}
	err = json.Unmarshal(c, sc)
	if err != nil {
		return nil, err
	}
	return
}

type gongzhConfig struct {
	AppID          string `json:"appid"`
	AppSecret      string `json:"appsecret"`
	Token          string `json:"token"`
	EncodingAESKey string `json:"encodingaeskey"`
}

var (
	apiPath = "/wxapi"
)

func start(ctx *cli.Context) {
	var sc *serverConfig

	config, err := filepath.Abs(configFile)
	if err != nil {
		log.Fatalln("filepath abs failed:", err)
	}
	tsc, err := newServerConfig(config)
	if err != nil {
		log.Fatalln("read config file failed:", err)
	}

	if sc.WxPay != nil {
		if sc.WxPay.CertPath != "" {
			redbagcertpath, redcerterr := filepath.Abs(sc.WxPay.CertPath)
			if redcerterr != nil {
				log.Println("RedBagCertpath error :", redcerterr)
			} else {
				sc.WxPay.CertPath = redbagcertpath
			}
		} else {
			log.Println("RedBag CertPath Is Nil")
		}
	}

	sc = tsc

	n := negroni.Classic()
	//初始化db

	redisService, err := fredis.NewRedisService(sc.Redis)
	if err != nil {
		log.Fatalln("init redis service failed:", err)
	}
	mycache := wxcache.NewMyRedisCache(redisService)
	//配置微信参数
	wxconfig := &wechat.Config{
		AppID:          sc.GZH.AppID,
		AppSecret:      sc.GZH.AppSecret,
		Token:          sc.GZH.Token,
		EncodingAESKey: sc.GZH.EncodingAESKey,
		Cache:          mycache,
	}

	wxservice := wxapi.NewWeChatService(wxconfig)

	addr := fmt.Sprintf("%s:%d", sc.Host, sc.Port)
	router := mux.NewRouter()
	router.PathPrefix(apiPath).Methods(http.MethodOptions).HandlerFunc(handleCrossOrigin)
	subrouter := router.PathPrefix(apiPath).Subrouter()
	wxrouter.Router(subrouter)

	n.UseFunc(setupWxServiceHandler(wxservice))
	n.UseHandler(router)

	//初始化微信到账服务单例
	wxpay.InitInstance(sc.WxPay)

	//register interruput handler
	hooker := osutils.NewInterruptHooker()
	hooker.AddHandler(osutils.InterruptHandlerFunc(stop))
	log.Println("trying to listen ", addr)
	if sc.Port == 443 {
		err = http.ListenAndServeTLS(addr, sc.Cert, sc.Key, n)
	} else {
		http.ListenAndServe(addr, n)
	}

	if err != nil {
		log.Fatalln("error ", err)
	}
}

func stop() {
	log.Println("stop server")
}

func handleCrossOrigin(rw http.ResponseWriter, req *http.Request) {
	log.Debug("跨源")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "content-type")
	rw.Header().Set("Access-Control-Allow-Credentials", "true")
	rw.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	rw.Header().Set("Access-Control-Max-Age", "3600")
	rw.WriteHeader(http.StatusOK)
	log.Debug("跨源,成功")
}

func setupWxServiceHandler(ls *wxapi.WeChatService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := wxapi.WithWeChatService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
