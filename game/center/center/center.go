package center

import (
	"context"
	acccounttypes "fgame/fgame/account/login/types"
	logintypes "fgame/fgame/account/login/types"
	centerclient "fgame/fgame/center/client"
	centerpb "fgame/fgame/center/pb"
	centertypes "fgame/fgame/center/types"
	gamecentertypes "fgame/fgame/game/center/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/robot/robot"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
)

type CenterConfig struct {
	Host         string `json:"host"`
	Port         int32  `json:"port"`
	SyncInterval int64  `json:"syncInterval"`
}

type ServerInfo struct {
	serverIp   string
	serverPort int32
	serverType centertypes.GameServerType
}

func (info *ServerInfo) GetServerIp() string {
	return info.serverIp
}

func (info *ServerInfo) GetServerPort() int32 {
	return info.serverPort
}

func (info *ServerInfo) GetServerType() centertypes.GameServerType {
	return info.serverType
}

func (info *ServerInfo) GetHost() string {
	return fmt.Sprintf("%s:%d", info.serverIp, info.serverPort)
}

func (info *ServerInfo) String() string {
	return fmt.Sprintf("服务器类型:%s,IP地址:%s:%d", info.serverType.String(), info.serverIp, info.serverPort)
}

func (info *ServerInfo) Equal(info2 *ServerInfo) bool {
	if info.serverType != info2.serverType {
		return false
	}
	if info.serverIp != info2.serverIp {
		return false
	}
	if info.serverPort != info2.serverPort {
		return false
	}
	return true
}

type ChatSet struct {
	minVip              int32
	minPlayerLevel      int32
	worldVip            int32
	worldPlayerLevel    int32
	teamVip             int32
	teamPlayerLevel     int32
	allianceVip         int32
	alliancePlayerLevel int32
	pVip                int32
	pPlayerLevel        int32
}

func (s *ChatSet) GetWorldVip() int32 {
	return s.worldVip
}

func (s *ChatSet) GetWorldPlayerLevel() int32 {
	return s.worldPlayerLevel
}

func (s *ChatSet) GetTeamVip() int32 {
	return s.teamVip
}
func (s *ChatSet) GetTeamPlayerLevel() int32 {
	return s.teamPlayerLevel
}

func (s *ChatSet) GetAlliancePlayerLevel() int32 {
	return s.alliancePlayerLevel
}

func (s *ChatSet) GetAllianceVip() int32 {
	return s.allianceVip
}

func (s *ChatSet) GetPrivatePlayerLevel() int32 {
	return s.pPlayerLevel
}

func (s *ChatSet) GetPrivateVip() int32 {
	return s.pVip
}

func NewChatSet(minVip, minPlayerLevel, worldVip, worldPlayerLevel, teamVip, teamPlayerLevel, allianceVip, alliancePlayerLevel, pVip, pPlayerLevel int32) *ChatSet {
	s := &ChatSet{
		minVip:              minVip,
		minPlayerLevel:      minPlayerLevel,
		worldVip:            worldVip,
		worldPlayerLevel:    worldPlayerLevel,
		teamVip:             teamVip,
		teamPlayerLevel:     teamPlayerLevel,
		allianceVip:         allianceVip,
		alliancePlayerLevel: alliancePlayerLevel,
		pVip:                pVip,
		pPlayerLevel:        pPlayerLevel,
	}
	return s
}

type CenterService struct {
	rwm        sync.RWMutex
	config     *CenterConfig
	client     *centerclient.Client
	syncTicker *time.Ticker
	done       chan struct{}
	//跨服列表
	serverInfoMap map[centertypes.GameServerType]*ServerInfo
	//跨服连接
	connMap map[centertypes.GameServerType]*grpc.ClientConn
	//服务器id
	serverId int32
	//开服时间
	startTime int64
	//交易服ip
	tradeServerIp string
	sdkList       []acccounttypes.SDKType
	//结婚类型
	marryPriceType   gamecentertypes.MarryPriceType
	tradeFlag        int32
	allianceFlag     int32
	xianJinFlag      int32
	neiWanJiaoYiFlag int32
	zhiZuanType      gamecentertypes.ZhiZunType
	chatSet          *ChatSet
}

func (s *CenterService) init(config *CenterConfig) (err error) {
	s.config = config
	s.serverInfoMap = make(map[centertypes.GameServerType]*ServerInfo)
	s.connMap = make(map[centertypes.GameServerType]*grpc.ClientConn)
	//TODO 修改配置
	syncInterValNano := s.config.SyncInterval * int64(time.Millisecond)
	s.syncTicker = time.NewTicker(time.Duration(syncInterValNano))
	s.done = make(chan struct{}, 1)
	cfg := &centerclient.Config{
		Host: s.config.Host,
		Port: s.config.Port,
	}
	s.client, err = centerclient.NewClient(cfg)
	if err != nil {
		return
	}
	//同步注册
	id, startTime, marryPriceType, allianceFlag, tradeFlag, xianJinFlag, neiWanJiaoYiFlag, zhiZuanType, tradeServerIp, sdkList, chatSet, err := s.register()
	if err != nil {
		return
	}
	s.serverId = id
	s.startTime = startTime
	s.marryPriceType = marryPriceType
	s.tradeServerIp = tradeServerIp
	s.allianceFlag = allianceFlag
	s.tradeFlag = tradeFlag
	s.xianJinFlag = xianJinFlag
	s.neiWanJiaoYiFlag = neiWanJiaoYiFlag
	s.zhiZuanType = zhiZuanType
	s.chatSet = chatSet
	for _, sdk := range sdkList {
		s.sdkList = append(s.sdkList, acccounttypes.SDKType(sdk))
	}
	err = s.updateCrossList()
	if err != nil {
		return
	}
	return
}

func (s *CenterService) Start() {
	log.WithFields(
		log.Fields{}).Info("center:中心服服务开启")
	go func() {
	Loop:
		for {
			select {
			case <-s.syncTicker.C:
				s.ping()
				s.updateCrossList()
				break
			case <-s.done:
				break Loop
			}
		}
		s.unregister()
		log.WithFields(
			log.Fields{}).Info("center:中心服服务结束")
	}()
}

func (s *CenterService) Stop() {
	s.syncTicker.Stop()
	s.done <- struct{}{}
}

func (s *CenterService) GetServerId() int32 {
	return s.serverId
}

func (s *CenterService) GetStartTime() int64 {
	return s.startTime
}

func (s *CenterService) GetMarryKindType() gamecentertypes.MarryPriceType {
	return s.marryPriceType
}

func (s *CenterService) GetZhiZunType() gamecentertypes.ZhiZunType {
	if !s.zhiZuanType.Valid() {
		return defaultZhiZunType
	}
	return s.zhiZuanType
}

func (s *CenterService) IsTradeOpen() bool {
	return s.tradeFlag != 0
}

func (s *CenterService) IsAllianceOpen() bool {
	return s.allianceFlag != 0
}

func (s *CenterService) IsXianJinOpen() bool {
	return s.xianJinFlag != 0
}

func (s *CenterService) IsNeiWanJiaoYiOpen() bool {
	return s.neiWanJiaoYiFlag != 0
}

func (s *CenterService) GetChatSet() *ChatSet {
	return s.chatSet
}

func (s *CenterService) SetStartTime(time int64) {
	s.startTime = time
}

const (
	defaultMarryPriceType   = gamecentertypes.MarryPriceTypeNormal
	defaultAllianceFlag     = 1
	defaultTradeFlag        = 0
	defaultXianJinFlag      = 0
	defaultNeiWanJiaoYiFlag = 0
	defaultZhiZunType       = gamecentertypes.ZhiZunTypeNormal
)

func (s *CenterService) register() (serverId int32, startTime int64, marryPriceType gamecentertypes.MarryPriceType, allianceFlag int32, tradeFlag int32, xianJinFlag int32, neiWanJiaoYiFlag int32, zhiZuanType gamecentertypes.ZhiZunType, tradeServerIp string, sdkList []int32, chatSet *ChatSet, err error) {
	//TODO 添加超时机制
	ctx := context.TODO()
	serverType := global.GetGame().GetServerType()
	platform := global.GetGame().GetPlatform()
	serverIndex := global.GetGame().GetServerIndex()
	serverIp := global.GetGame().GetServerIp()
	serverPort := global.GetGame().GetServerPort()

	resp, err := s.client.Register(ctx, serverType, platform, serverIndex, serverIp, serverPort)

	if err != nil {
		return
	}
	startTime = resp.GetStartTime()
	marryPriceType = gamecentertypes.MarryPriceTypeNormal
	allianceFlag = defaultAllianceFlag
	tradeFlag = defaultTradeFlag
	xianJinFlag = defaultXianJinFlag
	neiWanJiaoYiFlag = defaultNeiWanJiaoYiFlag
	zhiZuanType = defaultZhiZunType
	if resp.PlatformSetting != nil {
		marryPriceType = gamecentertypes.MarryPriceType(resp.GetPlatformSetting().GetMarryKindType())
		allianceFlag = resp.GetPlatformSetting().GetAllianceFlag()
		tradeFlag = resp.GetPlatformSetting().GetTradeFlag()
		xianJinFlag = resp.GetPlatformSetting().GetXianJinFlag()
		neiWanJiaoYiFlag = resp.GetPlatformSetting().GetNeiWanJiaoYiFlag()
		zhiZuanType = gamecentertypes.ZhiZunType(resp.GetPlatformSetting().GetZhiZuanFlag())
	}

	if resp.PlatformChatSetting != nil {
		chatSet = NewChatSet(
			resp.PlatformChatSetting.MinVip,
			resp.PlatformChatSetting.MinPlayerLevel,
			resp.PlatformChatSetting.WorldVip,
			resp.PlatformChatSetting.WorldPlayerLevel,
			resp.PlatformChatSetting.TeamVip,
			resp.PlatformChatSetting.TeamPlayerLevel,
			resp.PlatformChatSetting.AllianceVip,
			resp.PlatformChatSetting.AlliancePlayerLevel,
			resp.PlatformChatSetting.PVip,
			resp.PlatformChatSetting.PPlayerLevel,
		)
	}
	onlineNum := player.GetOnlinePlayerManager().Count()
	//TODO:zrc 临时修改
	robotNum := int32(0)
	rs := robot.GetRobotService()
	if rs != nil {
		robotNum = robot.GetRobotService().GetNumOfRobot()
	}
	tradeServerIp = resp.GetTradeServerIp()
	sdkList = resp.GetSdkList()
	log.WithFields(
		log.Fields{
			"serverType":  serverType.String(),
			"platform":    platform,
			"serverIndex": serverIndex,
			"serverIp":    serverIp,
			"serverPort":  serverPort,
			"onlineNum":   onlineNum,
			"robotNum":    robotNum,
			"sdkList":     sdkList,
		}).Info("center:同步服务器")
	if chatSet != nil {
		log.WithFields(
			log.Fields{
				"privateLevel":        chatSet.GetPrivatePlayerLevel(),
				"privateVip":          chatSet.GetPrivateVip(),
				"alliancePlayerLevel": chatSet.GetAlliancePlayerLevel(),
				"allianceVip":         chatSet.GetAllianceVip(),
				"teamPlayerLevel":     chatSet.GetTeamPlayerLevel(),
				"teamVip":             chatSet.GetTeamVip(),
				"worldPlayerLevel":    chatSet.GetWorldPlayerLevel(),
				"worldVip":            chatSet.GetWorldVip(),
			}).Info("center:同步服务器")
	}
	return resp.GetId(), startTime, marryPriceType, allianceFlag, tradeFlag, xianJinFlag, neiWanJiaoYiFlag, zhiZuanType, tradeServerIp, sdkList, chatSet, nil
}

func (s *CenterService) ping() {
	//TODO 添加超时机制
	_, _, marryPriceType, allianceFlag, tradeFlag, xianJinFlag, neiWanJiaoYiFlag, zhiZuanType, _, _, chatSet, err := s.register()
	if err != nil {
		log.WithFields(
			log.Fields{}).Warn("center:同步服务器,失败")
	}

	s.rwm.Lock()
	defer func() {
		s.rwm.Unlock()

	}()
	s.marryPriceType = marryPriceType
	s.allianceFlag = allianceFlag
	s.tradeFlag = tradeFlag
	s.xianJinFlag = xianJinFlag
	s.neiWanJiaoYiFlag = neiWanJiaoYiFlag
	s.zhiZuanType = zhiZuanType
	s.chatSet = chatSet
	return
}

func (s *CenterService) updateCrossList() (err error) {
	//TODO 添加超时机制
	ctx := context.TODO()

	serverId := global.GetGame().GetServerId()
	resp, err := s.client.GetCrossList(ctx, serverId)
	if err != nil {

		return
	}
	tempServerInfoList := resp.GetServerInfoList()
	serverInfoList := ConvertFromCrossServerInfoList(tempServerInfoList)

	newServerMap := getServerInfoMap(serverInfoList)

	s.rwm.Lock()
	defer s.rwm.Unlock()

	//关闭不存在的跨服
	for _, serverInfo := range s.serverInfoMap {
		newServerInfo, ok := newServerMap[serverInfo.GetServerType()]
		if ok && newServerInfo.Equal(serverInfo) {
			continue
		}
		delete(s.serverInfoMap, serverInfo.GetServerType())
		conn, ok := s.connMap[serverInfo.GetServerType()]
		if !ok {
			continue
		}
		log.WithFields(
			log.Fields{
				"serverInfo": serverInfo,
			}).Debug("center:关闭不存在的跨服连接")

		conn.Close()
		delete(s.connMap, serverInfo.GetServerType())
	}

	//添加不存在的跨服
	for _, serverInfo := range newServerMap {
		_, ok := s.serverInfoMap[serverInfo.GetServerType()]
		if ok {
			continue
		}
		crossIp := serverInfo.GetHost()
		//添加连接
		conn, err := grpc.Dial(crossIp, grpc.WithInsecure())
		if err != nil {
			log.WithFields(
				log.Fields{
					"serverInfo": serverInfo.String(),
				}).Warn("center:连接跨服失败")
			continue
		}
		log.WithFields(
			log.Fields{
				"serverInfo": serverInfo,
			}).Debug("center:创建新的跨服连接")
		s.connMap[serverInfo.GetServerType()] = conn
		s.serverInfoMap[serverInfo.GetServerType()] = serverInfo
	}

	log.WithFields(
		log.Fields{
			"serverInfoList": serverInfoList,
		}).Info("center:获取跨服列表")
	return
}

func (s *CenterService) unregister() {
	//TODO 添加超时机制
	ctx := context.TODO()
	serverId := s.serverId
	_, err := s.client.Unregister(ctx, serverId)

	if err != nil {
		log.WithFields(
			log.Fields{
				"serverId": serverId,
			}).Warn("center:取消注册失败")
		return
	}
	log.WithFields(
		log.Fields{
			"serverId": serverId,
		}).Info("center:取消注册服务器")
	return
}

func (s *CenterService) GetCross(serverType centertypes.GameServerType) *grpc.ClientConn {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	conn, ok := s.connMap[serverType]
	if !ok {
		return nil
	}
	return conn
}

func (s *CenterService) GetAllCross() (connList []*grpc.ClientConn) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	crossServerTypeList := centertypes.GetCrossServerTypeList()
	for _, crossServerType := range crossServerTypeList {
		conn, ok := s.connMap[crossServerType]
		if !ok {
			continue
		}
		connList = append(connList, conn)
	}
	return
}

func (s *CenterService) SyncPlayerInfo(pl player.Player) {
	go func() {
		ctx := context.TODO()
		userId := pl.GetUserId()
		serverId := pl.GetServerId()
		playerId := pl.GetId()
		playerName := pl.GetName()
		role := int32(pl.GetRole())
		sex := int32(pl.GetSex())
		level := pl.GetLevel()
		zhuanShu := pl.GetZhuanSheng()

		req := &centerpb.PlayerInfoSyncRequest{}
		req.PlayerInfo = &centerpb.PlayerServerInfo{}
		req.PlayerInfo.UserId = userId
		req.PlayerInfo.ServerId = serverId
		req.PlayerInfo.PlayerId = playerId
		req.PlayerInfo.PlayerName = playerName
		req.PlayerInfo.Role = role
		req.PlayerInfo.Sex = sex
		req.PlayerInfo.Level = level
		req.PlayerInfo.ZhuanShu = zhuanShu

		_, err := s.client.SyncPlayerServerInfo(ctx, req)
		if err != nil {
			return
		}
	}()
}

func (s *CenterService) Login(ctx context.Context, token string, serverId int32, originServerId int32) (userId int64, platformUserId string, sdkType logintypes.SDKType, devicePlatformType logintypes.DevicePlatformType, gm int32, iosVersion string, androidVersion string, err error) {
	req := &centerpb.LoginVerifyRequest{}
	req.Token = token
	req.ServerId = serverId
	req.OriginServerId = originServerId
	resp, err := s.client.LoginVerify(ctx, req)
	if err != nil {
		return
	}
	userId = resp.UserId
	sdkType = logintypes.SDKType(resp.SdkType)
	devicePlatformType = logintypes.DevicePlatformType(resp.DevicePlatformType)
	platformUserId = resp.PlatformUserId
	gm = resp.Gm
	iosVersion = resp.IosVersion
	androidVersion = resp.AndroidVersion
	return
}

func (s *CenterService) GetTradeServer() (ip string, port int32, err error) {
	tempArr := strings.Split(s.tradeServerIp, ":")
	if len(tempArr) != 2 {
		err = fmt.Errorf("交易服ip[%s]格式不对", s.tradeServerIp)
		return
	}
	portInt, err := strconv.ParseInt(tempArr[1], 10, 64)
	if err != nil {
		return
	}
	ip = tempArr[0]
	port = int32(portInt)
	return
}

func (s *CenterService) GetSdkList() []acccounttypes.SDKType {
	return s.sdkList
}

func getServerInfoMap(infoList []*ServerInfo) map[centertypes.GameServerType]*ServerInfo {
	infoMap := make(map[centertypes.GameServerType]*ServerInfo)
	for _, info := range infoList {
		infoMap[info.GetServerType()] = info
	}
	return infoMap
}

var (
	once sync.Once
	cs   *CenterService
)

func Init(config *CenterConfig) (err error) {
	once.Do(func() {
		cs = &CenterService{}
		err = cs.init(config)
	})
	return err
}

func GetCenterService() *CenterService {
	return cs
}
