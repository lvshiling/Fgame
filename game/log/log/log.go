package log

import (
	"encoding/json"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/jieyi/jieyi"

	"fgame/fgame/core/messaging"
	messagingnats "fgame/fgame/core/messaging/nats"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	logserverlog "fgame/fgame/logserver/log"
	logservermodel "fgame/fgame/logserver/model"
	logserverpb "fgame/fgame/logserver/pb"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats"
)

type LogService interface {
	Start()
	Stop()
	SendLog(logMsg logserverlog.LogMsg)
	SendChatLog(chatLogObj logserverlog.LogMsg)
}

type LogConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type logService struct {
	p        messaging.Producer
	msgChan  chan proto.Message
	chatChan chan proto.Message
	cfg      *LogConfig
}

func (s *logService) init() (err error) {
	url := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	conn, err := nats.Connect(url)
	if err != nil {
		return
	}
	s.p = messagingnats.NewNatsProducer(conn)
	return
}

func (s *logService) Start() {
	go func() {
		for {
			select {
			case msg := <-s.msgChan:
				logBytes, err := proto.Marshal(msg)
				if err != nil {
					return
				}
				err = s.p.Send(topic, logBytes)
				if err != nil {
					log.WithFields(
						log.Fields{
							"err": err,
						}).Warn("log:发送远程日志,失败")
				}
			case msg := <-s.chatChan:
				logBytes, err := proto.Marshal(msg)
				if err != nil {
					return
				}
				err = s.p.Send(chatTopic, logBytes)
				if err != nil {
					log.WithFields(
						log.Fields{
							"err": err,
						}).Warn("log:发送聊天监控日志,失败")
				}
			}
		}
	}()

}

func (s *logService) Stop() {
	close(s.msgChan)
	//TODO 等候结束
}

var (
	topic     = "log"
	chatTopic = "chat"
)

func (s *logService) SendLog(logObj logserverlog.LogMsg) {
	msg, err := protoMsgFromLogMsg(logObj)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("log:压缩错误")
		return
	}
	s.msgChan <- msg
	return
}

func protoMsgFromLogMsg(logObj logserverlog.LogMsg) (tmsg proto.Message, err error) {
	content, err := json.Marshal(logObj)
	if err != nil {
		return
	}
	msg := &logserverpb.LogMessage{}
	msg.Content = content
	logName := logObj.LogName()
	msg.LogName = &logName
	tmsg = msg
	return
}

func (s *logService) SendChatLog(chatLogObj logserverlog.LogMsg) {
	msg, err := protoMsgFromLogMsg(chatLogObj)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("log:压缩错误")
		return
	}
	s.chatChan <- msg
	s.msgChan <- msg
}

var (
	capacity = 100000
)

func NewLogService(p messaging.Producer) LogService {
	s := &logService{}
	s.p = p

	return s
}

var (
	s *logService
)

func GetLogService() LogService {
	return s
}

func Init(logCfg *LogConfig) (err error) {
	s = &logService{
		cfg: logCfg,
	}
	s.msgChan = make(chan proto.Message, capacity)
	s.chatChan = make(chan proto.Message, capacity)
	err = s.init()
	if err != nil {
		return
	}
	return
}

func PlayerLogMsgFromPlayer(pl player.Player) (msg logservermodel.PlayerLogMsg) {
	sdkUserId := pl.GetPlatformUserId()
	userId := pl.GetUserId()
	playerId := pl.GetId()
	name := pl.GetName()
	role := int32(pl.GetRole())
	sex := int32(pl.GetSex())
	level := pl.GetLevel()
	now := global.GetGame().GetTimeService().Now()
	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerIndex()
	serverType := int32(global.GetGame().GetServerType())
	ip := pl.GetIp()
	vipLevel := pl.GetVip()
	sdkType := int32(pl.GetSDKType())
	deviceType := int32(pl.GetDevicePlatformType())
	msg = logservermodel.PlayerLogMsg{
		LogTime:    now,
		Platform:   platform,
		ServerType: serverType,
		ServerId:   serverId,
		UserId:     userId,
		SdkUserId:  sdkUserId,
		PlayerId:   playerId,
		Ip:         ip,
		Name:       name,
		Role:       role,
		Sex:        sex,
		Level:      level,
		Vip:        vipLevel,
		SdkType:    sdkType,
		DeviceType: deviceType,
	}

	return
}

func SystemLogMsg() (msg logservermodel.SystemLogMsg) {

	now := global.GetGame().GetTimeService().Now()
	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerIndex()
	serverType := int32(global.GetGame().GetServerType())

	msg = logservermodel.SystemLogMsg{
		LogTime:    now,
		Platform:   platform,
		ServerType: serverType,
		ServerId:   serverId,
	}

	return
}

func AllianceLogMsgFromPlayer(al *alliance.Alliance) (msg logservermodel.AllianceLogMsg) {
	now := global.GetGame().GetTimeService().Now()
	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerIndex()
	serverType := int32(global.GetGame().GetServerType())
	allianceId := al.GetAllianceId()
	allianceName := al.GetAllianceName()
	allianceLevel := al.GetAllianceLevel()
	allianceJianShe := al.GetAllianceObject().GetJianShe()
	msg = logservermodel.AllianceLogMsg{
		LogTime:    now,
		Platform:   platform,
		ServerType: serverType,
		ServerId:   serverId,
		AllianceId: allianceId,
		Name:       allianceName,
		Level:      allianceLevel,
		JianShe:    allianceJianShe,
	}

	return
}

func JieYiLogMsgFromPlayer(jieYi *jieyi.JieYi) (msg logservermodel.JieYiLogMsg) {
	now := global.GetGame().GetTimeService().Now()
	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerIndex()
	serverType := int32(global.GetGame().GetServerType())
	jieYiId := jieYi.GetJieYiObject().GetId()
	jieYiName := jieYi.GetJieYiObject().GetName()
	msg = logservermodel.JieYiLogMsg{
		LogTime:    now,
		Platform:   platform,
		ServerType: serverType,
		ServerId:   serverId,
		JieYiId:    jieYiId,
		Name:       jieYiName,
	}
	return
}

func PlayerTradeLogMsgFromTradeObject(playerId int64) (msg logservermodel.PlayerTradeLogMsg) {
	now := global.GetGame().GetTimeService().Now()
	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerIndex()
	serverType := int32(global.GetGame().GetServerType())
	msg = logservermodel.PlayerTradeLogMsg{
		LogTime:    now,
		Platform:   platform,
		ServerType: serverType,
		ServerId:   serverId,
		PlayerId:   playerId,
	}

	return
}
