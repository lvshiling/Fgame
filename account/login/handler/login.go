package handler

import (
	"context"
	commonlogic "fgame/fgame/account/common/logic"
	"fgame/fgame/account/global"
	"fgame/fgame/account/login/login"
	"fgame/fgame/account/login/pbutil"
	"fgame/fgame/account/login/types"
	"fgame/fgame/account/notice/notice"
	"fgame/fgame/account/player/player"
	"fgame/fgame/account/processor"
	"fgame/fgame/account/serverlist/serverlist"
	accountsession "fgame/fgame/account/session"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/exception"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"strings"

	"time"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/status"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ACCOUNT_LOGIN_TYPE), dispatch.HandlerFunc(handleAccountLogin))
}

//登陆
func handleAccountLogin(s session.Session, msg interface{}) error {
	log.Debug("login:处理账户登陆")
	//不管请求怎么样 直接关闭

	csAccountLogin := msg.(*uipb.CSAccountLogin)
	devicePlatformInt := csAccountLogin.GetDevicePlatform()
	devicePlatform := types.DevicePlatformType(devicePlatformInt)
	sess := accountsession.SessionInContext(s.Context())
	if !devicePlatform.Valid() {
		log.WithFields(
			log.Fields{
				"devicePlatform": devicePlatformInt,
			}).Warn("login:设备平台无效")
		commonlogic.SendSessionSystemMessage(sess, lang.CommonArgumentInvalid)
		s.Close()
		return nil
	}

	platformInt := csAccountLogin.GetPlatform()
	platform := types.SDKType(platformInt)
	if !global.GetAccount().EnablePC() && platform == types.SDKTypePC {
		log.WithFields(
			log.Fields{
				"devicePlatform": devicePlatformInt,
			}).Warn("login:不支持pc登陆")
		commonlogic.SendSessionSystemMessage(sess, lang.CommonArgumentInvalid)
		s.Close()
		return nil
	}
	err := accountLogin(sess, platform, devicePlatform, msg)
	if err != nil {
		log.WithFields(
			log.Fields{
				"devicePlatform": devicePlatformInt,
				"platform":       platformInt,
				"err":            err,
			}).Error("login:登陆,错误")
		return err
	}
	log.WithFields(
		log.Fields{
			"devicePlatform": devicePlatformInt,
			"platform":       platformInt,
		}).Debug("login:处理账户登陆完成")

	return nil
}

func accountLogin(sess accountsession.Session, platform types.SDKType, devicePlatform types.DevicePlatformType, data interface{}) (err error) {
	defer sess.Close(true)

	h := login.GetLoginHandler(platform)
	if h == nil {
		log.WithFields(
			log.Fields{
				"platform": platform,
			}).Warn("login:登陆处理器没有注册")
		commonlogic.SendSessionSystemMessage(sess, lang.AccountLoginPlatformInvalid)

		return
	}
	flag, returnPlatform, platformUserId, err := h.Login(devicePlatform, data)
	if err != nil {
		return
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"platform": platform,
			}).Warn("login:登陆验证,失败")
		commonlogic.SendSessionSystemMessage(sess, lang.AccountLoginPlatformVerifyFailed)
		return
	}

	if len(platformUserId) == 0 {
		log.WithFields(
			log.Fields{
				"platform": platform,
			}).Warn("login:登陆验证,失败")
		commonlogic.SendSessionSystemMessage(sess, lang.AccountLoginPlatformVerifyFailed)
		return
	}

	//TODO 是否放在底层
	defer func() {
		if err != nil {
			if _, ok := status.FromError(err); ok {
				log.WithFields(
					log.Fields{
						"err": err,
					}).Warn("login:登陆验证,失败")
				commonlogic.SendSessionExceptionMessage(sess, exception.ExceptionCodeServerBusy)
				err = nil
			}
		}
	}()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	ip := sess.Ip()
	realIpList := strings.Split(ip, ":")
	if len(realIpList) >= 1 {
		ip = realIpList[0]
	}
	//发送登陆
	userId, token, expiredTime, gm, err := login.GetLoginService().Login(timeoutCtx, devicePlatform, types.SDKType(returnPlatform), platformUserId, ip)
	if err != nil {
		return
	}
	cancel()
	if userId == 0 {
		log.WithFields(
			log.Fields{
				"ip": ip,
			}).Warn("login:登陆验证,账户或ip被封")
		commonlogic.SendSessionExceptionMessage(sess, exception.ExceptionCodeAccountForbid)
		return
	}

	serverPlatform := int32(returnPlatform)

	eg, egCtx := errgroup.WithContext(context.Background())
	var playerInfoList []*player.PlayerInfo
	var serverInfoList []*serverlist.ServerInfo
	var content string
	eg.Go(getPlayerList(egCtx, userId, &playerInfoList))
	eg.Go(getServerList(egCtx, serverPlatform, gm, &serverInfoList))
	eg.Go(getNotice(egCtx, serverPlatform, &content))
	err = eg.Wait()
	if err != nil {
		return
	}
	scAccountLogin := pbutil.BuildSCAccountLogin(userId, token, expiredTime, playerInfoList, serverInfoList, content, platformUserId)
	sess.Send(scAccountLogin)
	return
}

//TODO 调整超时参数
var (
	rpcTimeout = time.Second * 10
)

func getPlayerList(ctx context.Context, userId int64, playerInfoListPtr *[]*player.PlayerInfo) func() error {
	return func() error {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
		defer cancel()
		playerInfoList, err := player.GetPlayerService().GetPlayerList(timeoutCtx, userId)
		if err != nil {
			return err
		}
		*playerInfoListPtr = playerInfoList
		return nil
	}
}

func getServerList(ctx context.Context, platform int32, gm int32, sererInfoListPtr *[]*serverlist.ServerInfo) func() error {
	return func() error {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
		defer cancel()
		serverInfoList, err := serverlist.GetServerService().GetServerList(timeoutCtx, platform, gm)
		if err != nil {
			return err
		}
		*sererInfoListPtr = serverInfoList
		return nil
	}
}

func getNotice(ctx context.Context, platform int32, contentPtr *string) func() error {
	return func() error {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
		defer cancel()
		n, err := notice.GetNoticeService().GetNotice(timeoutCtx, platform)
		if err != nil {
			return err
		}
		*contentPtr = n
		return nil
	}
}
