package center

import (
	"context"
	"encoding/json"
	centerpb "fgame/fgame/center/pb"
	"fgame/fgame/center/store"
	"fgame/fgame/center/types"
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	redis "github.com/chasex/redis-go-cluster"
	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerInfo struct {
	id                int32
	platform          int32
	serverId          int32
	serverType        types.GameServerType
	serverIp          string
	serverPort        string
	remoteIp          string
	remotePort        string
	serverTag         types.ServerTag
	serverStatus      types.ServerStatus
	name              string
	startTime         int64
	parentServerId    int32
	updateTime        int64
	createTime        int64
	deleteTime        int64
	preShow           int32
	pingTaiFuServerId int32
	//服务器状态
	status types.GameServerStatus
	ip     string
	port   int32
	//上次心跳时间
	lastHeartbeatTime int64
}

func (info *ServerInfo) String() string {
	if info == nil {
		return ""
	}
	return fmt.Sprintf("serverId:%d,serverIp:%s,serverPort:%s", info.serverId, info.serverIp, info.serverPort)
}

func (info *ServerInfo) GetId() int32 {
	return info.id
}

func (info *ServerInfo) GetPlatform() int32 {
	return info.platform
}

func (info *ServerInfo) GetParentServerId() int32 {
	return info.parentServerId
}

func (info *ServerInfo) GetPingTaiFuServerId() int32 {
	return info.pingTaiFuServerId
}

func (info *ServerInfo) GetServerId() int32 {
	return info.serverId
}

func (info *ServerInfo) GetServerType() types.GameServerType {
	return info.serverType
}

func (info *ServerInfo) GetServerIp() string {
	return info.serverIp
}

func (info *ServerInfo) GetServerPort() string {
	return info.serverPort
}

func (info *ServerInfo) GetRemoteIp() string {
	return info.remoteIp
}

func (info *ServerInfo) GetRemotePort() string {
	return info.remotePort
}

func (info *ServerInfo) GetServerTag() types.ServerTag {
	return info.serverTag
}

func (info *ServerInfo) GetServerStatus() types.ServerStatus {
	return info.serverStatus
}

func (info *ServerInfo) GetName() string {
	return info.name
}

func (info *ServerInfo) GetStartTime() int64 {
	return info.startTime
}

const (
	maxHearbeatTime = int64(120 * time.Second / time.Millisecond)
)

func (info *ServerInfo) GetStatus() types.GameServerStatus {

	if info.status == types.GameServerStatusNormal {
		now := timeutils.TimeToMillisecond(time.Now())
		if now-info.lastHeartbeatTime > maxHearbeatTime {
			return types.GameServerStatusMaintain
		}
		return info.status
	}
	return info.status
}

func (info *ServerInfo) GetIp() string {
	return info.ip
}

func (info *ServerInfo) GetPort() int32 {
	return info.port
}

func (info *ServerInfo) FromEntity(e *store.ServerEntity) {
	info.id = e.Id
	info.platform = e.Platform
	info.serverId = e.ServerId
	info.parentServerId = e.ParentServerId
	info.name = e.Name
	info.serverType = types.GameServerType(e.ServerType)
	info.status = types.GameServerStatusNormal
	info.serverTag = types.ServerTag(e.ServerTag)
	info.serverStatus = types.ServerStatus(e.ServerStatus)
	info.serverIp = e.ServerIp
	info.serverPort = e.ServerPort
	info.ip = e.ServerIp
	info.serverPort = e.ServerPort
	info.remoteIp = e.ServerRemoteIp
	info.remotePort = e.ServerRemotePort
	info.startTime = e.StartTime
	info.parentServerId = e.ParentServerId
	info.pingTaiFuServerId = e.PingTaiFuServerId
	info.preShow = e.PreShow
	info.updateTime = e.UpdateTime
	info.deleteTime = e.DeleteTime
	info.createTime = e.CreateTime
}

func newServerInfo() *ServerInfo {
	info := &ServerInfo{}
	return info
}

func (info *ServerInfo) register(ip string, port int32) {
	info.ip = ip
	info.port = port
	info.status = types.GameServerStatusNormal
	now := timeutils.TimeToMillisecond(time.Now())
	info.lastHeartbeatTime = now
}

func (info *ServerInfo) maintain() {
	info.status = types.GameServerStatusMaintain
}

type MergeRecordInfo struct {
	id            int32
	platform      int32
	fromServerId  int32
	toServerId    int32
	finalServerId int32
	mergeTime     int64
	updateTime    int64
	createTime    int64
	deleteTime    int64
}

func (info *MergeRecordInfo) String() string {
	if info == nil {
		return ""
	}
	return fmt.Sprintf("platform:%d,fromServer:%d,toServer:%d,finalServer:%d", info.platform, info.fromServerId, info.toServerId, info.finalServerId)
}

func (info *MergeRecordInfo) GetId() int32 {
	return info.id
}

func (info *MergeRecordInfo) GetPlatform() int32 {
	return info.platform
}

func (info *MergeRecordInfo) GetFromServerId() int32 {
	return info.fromServerId
}

func (info *MergeRecordInfo) GetFinalserverId() int32 {
	return info.finalServerId
}

func (info *MergeRecordInfo) FromEntity(e *store.MergeRecordEntity) {
	info.id = e.Id
	info.platform = e.Platform
	info.fromServerId = e.FromServerId
	info.toServerId = e.ToServerId
	info.finalServerId = e.FinalServerId
	info.mergeTime = e.MergeTime
	info.updateTime = e.UpdateTime
	info.deleteTime = e.DeleteTime
	info.createTime = e.CreateTime
}

func newMergeRecordInfo() *MergeRecordInfo {
	info := &MergeRecordInfo{}
	return info
}

type SettingInfo struct {
	MarrySet              int32 `json:"marrySet"`              //结婚配置 1:当前版本 2:廉价版本
	AllianceWarehouseFlag int32 `json:"allianceWarehouseFlag"` //仙盟仓库 0:关 1:开
	JiaoYiHangFlag        int32 `json:"jiaoYiHangFlag"`        //交易行 0:关 1:开
	XianJinFlag           int32 `json:"cashTiXianFlag"`        //现金提现 0:关 1:开
	NeiWanJiaoYiFlag      int32 `json:"neiWanJiaoYiFlag"`      //内玩交易提现 0:关 1:开
	ZhiZuanFlag           int32 `json:"zhiZuanFlag"`           //至尊会员 1:当前版本 2:贵价版本
}

type PlatformSetting struct {
	id          int64
	platform    int32
	settingInfo *SettingInfo
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func (info *PlatformSetting) GetId() int64 {
	return info.id
}

func (info *PlatformSetting) GetPlatform() int32 {
	return info.platform
}

func (info *PlatformSetting) GetSettingInfo() *SettingInfo {
	return info.settingInfo
}

func (info *PlatformSetting) FromEntity(e *store.PlatformSettingEntity) (err error) {
	info.id = e.Id
	info.platform = e.PlatformId
	settingInfo := &SettingInfo{}
	err = json.Unmarshal([]byte(e.SettingContent), settingInfo)
	if err != nil {
		return
	}
	info.settingInfo = settingInfo
	info.updateTime = e.UpdateTime
	info.deleteTime = e.DeleteTime
	info.createTime = e.CreateTime
	return
}

func newPlatformSetting() *PlatformSetting {
	info := &PlatformSetting{}
	return info
}

type PlatformChatset struct {
	id             int64
	platform       int32
	minVip         int32
	minPlayerLevel int32

	worldVip         int32
	worldPlayerLevel int32
	pChatVip         int32
	pChatPlayerLevel int32
	guildVip         int32
	guildPlayerLevel int32
	teamVip          int32
	teamPlayerLevel  int32
	updateTime       int64
	createTime       int64
	deleteTime       int64
}

func (info *PlatformChatset) GetId() int64 {
	return info.id
}

func (info *PlatformChatset) GetPlatform() int32 {
	return info.platform
}

func (info *PlatformChatset) FromEntity(e *store.PlatformChatSetEntity) (err error) {
	info.id = e.Id
	info.platform = e.PlatformId
	info.minVip = e.MinVip
	info.minPlayerLevel = e.MinPlayerlevel

	info.worldVip = e.WorldVip
	info.worldPlayerLevel = e.WorldPlayerLevel
	info.pChatVip = e.PChatVip
	info.pChatPlayerLevel = e.PChatPlayerLevel
	info.guildVip = e.GuildVip
	info.guildPlayerLevel = e.GuildPlayerLevel
	info.teamVip = e.TeamVip
	info.teamPlayerLevel = e.TeamPlayerLevel

	info.updateTime = e.UpdateTime
	info.deleteTime = e.DeleteTime
	info.createTime = e.CreateTime
	return
}

func newPlatformChatset() *PlatformChatset {
	info := &PlatformChatset{}
	return info
}

type CenterOptions struct {
	GMDB        *coredb.DbConfig       `json:"gmDb"`
	DB          *coredb.DbConfig       `json:"db"`
	Redis       *coreredis.RedisConfig `json:"redis"`
	ExpiredTime int64                  `json:"expiredTime"`
	Key         string                 `json:"key"`
}

type CenterServer struct {
	rwm           sync.RWMutex
	options       *CenterOptions
	rs            coreredis.RedisService
	key           []byte
	serverStore   store.ServerStore
	playerStore   store.PlayerStore
	userStore     store.UserStore
	platformStore store.PlatformStore
	//全平台服
	allPlatformServerInfo *ServerInfo
	//平台所有游戏服
	serverCategoryMap map[int32]map[types.GameServerType]map[int32]*ServerInfo
	serverMap         map[int32]*ServerInfo
	//加载sdk对应中心服
	sdkPlatformMap map[int32]int32
	noticeStore    store.NoticeStore
	//获取所有平台公告
	noticeMap map[int32]*NoticeInfo

	//合服记录
	platformMergeRecordMap map[int32]map[int32]*MergeRecordInfo
	marryPriceStore        store.MarryPriceStore
	//结婚配置
	marryPriceMap map[int32]int32

	clientVersionStore store.ClientVersionStore
	//客户端版本
	clientVersion     *ClientVersion
	serverConfigStore store.ServerConfigStore
	//服务器配置
	serverConfig         *ServerConfig
	platformSettingStore store.PlatformSettingStore
	//平台设置
	platformSettingMap   map[int32]*PlatformSetting
	platformChatsetStore store.PlatformChatsetStore
	//聊天配置
	platformChatSetMap map[int32]*PlatformChatset
}

func (s *CenterServer) isValidPlatform(platform int32) bool {
	return true
}

func (s *CenterServer) init() (err error) {
	s.serverMap = make(map[int32]*ServerInfo)
	s.serverCategoryMap = make(map[int32]map[types.GameServerType]map[int32]*ServerInfo)
	s.platformMergeRecordMap = make(map[int32]map[int32]*MergeRecordInfo)
	keyFile, err := filepath.Abs(s.options.Key)
	if err != nil {
		return
	}
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return
	}
	s.key = key
	gmdb, err := coredb.NewDBService(s.options.GMDB)
	if err != nil {
		return
	}
	db, err := coredb.NewDBService(s.options.DB)
	if err != nil {
		return
	}
	s.rs, err = coreredis.NewRedisService(s.options.Redis)
	if err != nil {
		return
	}

	s.serverStore = store.NewServerStore(db)
	err = s.loadAllMergeRecords()
	if err != nil {
		return
	}

	err = s.loadAllServers()
	if err != nil {
		return
	}

	s.playerStore = store.NewPlayerStore(db)
	s.userStore = store.NewUserStore(db)
	s.noticeStore = store.NewNoticeStore(db)
	err = s.loadAllNotices()
	if err != nil {
		return
	}

	s.platformStore = store.NewPlatformStore(gmdb)

	err = s.loadAllSDK()
	if err != nil {
		return
	}

	s.marryPriceStore = store.NewMarryPriceStore(db)

	err = s.loadMarryPirce()
	if err != nil {
		return
	}
	s.clientVersionStore = store.NewClientVersionStore(db)
	err = s.loadClientVersion()
	if err != nil {
		return
	}
	s.serverConfigStore = store.NewServerConfigStore(db)
	err = s.loadServerConfig()
	if err != nil {
		return
	}
	s.platformChatsetStore = store.NewPlatformChatsetStore(db)
	s.platformSettingStore = store.NewPlatformSettingStore(db)
	err = s.loadPlatformSettings()
	if err != nil {
		return
	}

	return
}

//加载所有公告
func (s *CenterServer) loadAllNotices() (err error) {
	s.noticeMap = make(map[int32]*NoticeInfo)
	noticeEntityList, err := s.noticeStore.GetAll()
	if err != nil {
		return
	}
	for _, noticeEntity := range noticeEntityList {
		noticeInfo := newNoticeInfo()
		noticeInfo.FromEntity(noticeEntity)
		s.noticeMap[noticeInfo.platformId] = noticeInfo
	}
	return
}

//加载所有服务器
func (s *CenterServer) loadAllMergeRecords() (err error) {
	mergeRecordEntityList, err := s.serverStore.GetAllMergeList()
	if err != nil {
		return
	}
	for _, mergeRecordEntity := range mergeRecordEntityList {
		mergeRecordInfo := newMergeRecordInfo()
		mergeRecordInfo.FromEntity(mergeRecordEntity)
		flag := s.addMergeRecord(mergeRecordInfo)
		if !flag {
			return fmt.Errorf("center:服务器[%d]加入失败,请检查数据库", mergeRecordInfo.id)
		}
	}

	return
}

//添加合服记录
func (s *CenterServer) addMergeRecord(info *MergeRecordInfo) bool {
	mergeRecordMap, ok := s.platformMergeRecordMap[info.GetPlatform()]
	if !ok {
		mergeRecordMap = make(map[int32]*MergeRecordInfo)
		s.platformMergeRecordMap[info.GetPlatform()] = mergeRecordMap
	}

	_, ok = mergeRecordMap[info.GetFromServerId()]
	if ok {
		return false
	}
	mergeRecordMap[info.GetFromServerId()] = info
	return true
}

//获取服务器
func (s *CenterServer) getMergeRecord(platform int32, fromServerId int32) (info *MergeRecordInfo) {

	infoMap, ok := s.platformMergeRecordMap[platform]
	if !ok {
		return nil
	}
	info, ok = infoMap[fromServerId]
	if !ok {
		return nil
	}

	return info
}

func (s *CenterServer) getMergeServerInfo(platform int32, fromServerId int32) (serverInfo *ServerInfo) {
	mergeInfo := s.getMergeRecord(platform, fromServerId)
	if mergeInfo == nil {
		return
	}
	return s.getServer(platform, types.GameServerTypeSingle, mergeInfo.GetFinalserverId())
}

//获取服务器
func (s *CenterServer) getMergeRecords(platform int32) (infoMap map[int32]*MergeRecordInfo) {
	infoMap, ok := s.platformMergeRecordMap[platform]
	if !ok {
		return nil
	}
	return infoMap
}

func (s *CenterServer) clearMergeRecords(platform int32) {
	_, ok := s.platformMergeRecordMap[platform]
	if !ok {
		return
	}
	delete(s.platformMergeRecordMap, platform)
}

//加载所有服务器
func (s *CenterServer) loadAllServers() (err error) {
	serverEntityList, err := s.serverStore.GetAll()
	if err != nil {
		return
	}
	now := timeutils.TimeToMillisecond(time.Now())
	for _, serverEntity := range serverEntityList {
		serverInfo := newServerInfo()
		serverInfo.FromEntity(serverEntity)
		serverInfo.lastHeartbeatTime = now
		flag := s.addServer(serverInfo)
		if !flag {
			return fmt.Errorf("center:服务器[%d]加入失败,请检查数据库", serverInfo.GetId())
		}
	}
	return
}

//加载所有sdk
func (s *CenterServer) loadAllSDK() (err error) {
	if err = s.refreshSDK(); err != nil {
		return
	}
	return
}

//刷新sdk
func (s *CenterServer) refreshSDK() (err error) {
	sdkPlatformMap := make(map[int32]int32)
	platformEntityList, err := s.platformStore.GetAll()
	if err != nil {
		return
	}
	for _, platformEntity := range platformEntityList {
		sdkPlatformMap[platformEntity.SdkType] = platformEntity.CenterPlatformId
	}
	s.sdkPlatformMap = sdkPlatformMap
	return nil
}

//加载所有sdk
func (s *CenterServer) loadMarryPirce() (err error) {
	if err = s.refreshMarryPrice(); err != nil {
		return
	}
	return
}

//刷新sdk
func (s *CenterServer) refreshMarryPrice() (err error) {
	marryPriceMap := make(map[int32]int32)
	marryPriceEntityList, err := s.marryPriceStore.GetAll()
	if err != nil {
		return
	}
	for _, marryPriceEntity := range marryPriceEntityList {
		marryPriceMap[marryPriceEntity.PlatformId] = marryPriceEntity.KindType
	}
	s.marryPriceMap = marryPriceMap
	return nil
}

//添加服务器
func (s *CenterServer) addServer(info *ServerInfo) bool {
	tServerMap, ok := s.serverCategoryMap[info.GetPlatform()]
	if !ok {
		tServerMap = make(map[types.GameServerType]map[int32]*ServerInfo)
		s.serverCategoryMap[info.platform] = tServerMap
	}
	ttServerMap, ok := tServerMap[info.GetServerType()]
	if !ok {
		ttServerMap = make(map[int32]*ServerInfo)
		tServerMap[info.GetServerType()] = ttServerMap
	}
	_, ok = ttServerMap[info.GetServerId()]
	if ok {
		return false
	}
	ttServerMap[info.GetServerId()] = info
	s.serverMap[info.GetId()] = info
	// if info.GetServerType() == types.GameServerTypeAll {
	// 	s.allPlatformServerInfo = info
	// }
	return true
}

//检测过期服务器
func (s *CenterServer) tick() {

}

func (s *CenterServer) getServerById(serverId int32) (info *ServerInfo) {
	info, ok := s.serverMap[serverId]
	if !ok {
		return nil
	}
	return info
}

//获取服务器
func (s *CenterServer) getServer(platform int32, serverType types.GameServerType, serverId int32) (info *ServerInfo) {
	platformServerMap, ok := s.serverCategoryMap[platform]
	if !ok {
		return
	}
	serverInfoMap := platformServerMap[serverType]
	if len(serverInfoMap) == 0 {
		return
	}
	info, ok = serverInfoMap[serverId]
	if !ok {
		return nil
	}
	return info
}

type serverInfoList []*ServerInfo

func (adl serverInfoList) Len() int {
	return len(adl)
}

func (adl serverInfoList) Less(i, j int) bool {
	return adl[i].serverId < adl[j].serverId
}

func (adl serverInfoList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

const (
	preShowTime = int64(30 * time.Minute / time.Millisecond)
)

//获取服务器
func (s *CenterServer) getGameServerList(platform int32) (infoList []*ServerInfo) {

	platformServerMap, ok := s.serverCategoryMap[platform]
	if !ok {
		return
	}
	serverInfoMap := platformServerMap[types.GameServerTypeSingle]
	if len(serverInfoMap) == 0 {
		return
	}
	infoList = make([]*ServerInfo, 0, len(serverInfoMap))
	for _, info := range serverInfoMap {
		infoList = append(infoList, info)
	}
	sort.Sort(serverInfoList(infoList))
	return infoList
}

func (s *CenterServer) clearServerList(platform int32, serverType types.GameServerType) {
	platformServerMap, ok := s.serverCategoryMap[platform]
	if !ok {
		return
	}
	delete(platformServerMap, serverType)
}

func (s *CenterServer) clearCrossList(serverType types.GameServerType) {
	crossServerMap, ok := s.serverCategoryMap[0]
	if !ok {
		return
	}
	delete(crossServerMap, serverType)
}

const (
	groupCapacity = 8
)

func (s *CenterServer) GetServerList(ctx context.Context, req *centerpb.ServerInfoListRequest) (res *centerpb.ServerInfoListResponse, err error) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	platform := req.GetPlatform()
	gm := req.GetGm()
	centerPlatformId := s.sdkPlatformMap[platform]
	res = &centerpb.ServerInfoListResponse{}
	serverInfoList := s.getGameServerList(centerPlatformId)
	mergeServerInfoList := make([]*ServerInfo, 0, len(serverInfoList))
	for _, serverInfo := range serverInfoList {
		mergeServerInfo := s.getMergeServerInfo(centerPlatformId, serverInfo.GetServerId())
		mergeServerInfoList = append(mergeServerInfoList, mergeServerInfo)
	}

	res.ServerInfoList = buildServerInfoList(gm, serverInfoList, mergeServerInfoList)

	return
}

func (s *CenterServer) GetServerInfo(ctx context.Context, req *centerpb.ServerInfoRequest) (res *centerpb.ServerInfoResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	sdk := req.GetPlatform()
	serverId := req.GetServerId()
	log.WithFields(
		log.Fields{
			"sdk":      sdk,
			"serverId": serverId,
		}).Debug("server:获取单服数据")
	platform := s.sdkPlatformMap[sdk]
	info := s.getServer(platform, types.GameServerTypeSingle, serverId)

	res = &centerpb.ServerInfoResponse{}
	if info == nil {
		return
	}

	mergeInfo := s.getMergeServerInfo(platform, serverId)

	res.ServerInfo = buildServerInfo(info, mergeInfo)
	log.Debug("server:获取单服数据成功")
	return
}

func (s *CenterServer) GetServerInfoByPlatform(ctx context.Context, req *centerpb.ServerInfoByPlatformRequest) (res *centerpb.ServerInfoByPlatformResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	platform := req.GetPlatform()
	serverId := req.GetServerId()
	log.WithFields(
		log.Fields{
			"platform": platform,
			"serverId": serverId,
		}).Debug("server:获取单服数据")

	info := s.getServer(platform, types.GameServerTypeSingle, serverId)

	res = &centerpb.ServerInfoByPlatformResponse{}
	if info == nil {
		return
	}

	mergeInfo := s.getMergeServerInfo(platform, serverId)

	res.ServerInfo = buildServerInfo(info, mergeInfo)
	log.Debug("server:获取单服数据成功")
	return
}

func (s *CenterServer) GetCrossList(ctx context.Context, req *centerpb.ServerCrossListRequest) (res *centerpb.ServerCrossListResponse, err error) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	log.WithFields(
		log.Fields{}).Debug("server:获取跨服列表")
	serverId := req.GetServerId()
	serverInfo := s.getServerById(serverId)
	if serverInfo == nil {
		log.WithFields(
			log.Fields{
				"serverId": serverId,
			}).Warn("server:获取跨服列表,失败")
		err = status.Errorf(codes.InvalidArgument, "无效服务器id: %d", serverId)
		return
	}
	serverType := serverInfo.GetServerType()
	if serverType != types.GameServerTypeSingle {
		log.WithFields(
			log.Fields{
				"serverId":   serverId,
				"serverType": serverType.String(),
			}).Warn("server:获取跨服列表,失败")
		err = status.Errorf(codes.InvalidArgument, "无效的服务器类型: %s", serverType.String())
		return
	}
	res = &centerpb.ServerCrossListResponse{}
	//获取组跨服 根据服务器组id
	groupId := int32(math.Ceil(float64(serverInfo.GetServerId()) / float64(groupCapacity)))
	groupServer := s.getServer(serverInfo.GetPlatform(), types.GameServerTypeGroup, groupId)

	if groupServer != nil && groupServer.GetStatus() == types.GameServerStatusNormal {
		res.ServerInfoList = append(res.ServerInfoList, ConvertCrossServerInfo(groupServer))
	}

	//获取区跨服 根据战区id
	regionServer := s.getServer(serverInfo.GetPlatform(), types.GameServerTypeRegion, serverInfo.GetParentServerId())
	if regionServer != nil && regionServer.GetStatus() == types.GameServerStatusNormal {
		res.ServerInfoList = append(res.ServerInfoList, ConvertCrossServerInfo(regionServer))
	}
	// 获取平台跨服 根据平台id
	platformServer := s.getServer(serverInfo.GetPlatform(), types.GameServerTypePlatform, 1)
	if platformServer != nil && platformServer.GetStatus() == types.GameServerStatusNormal {
		res.ServerInfoList = append(res.ServerInfoList, ConvertCrossServerInfo(platformServer))
	}
	crossServer := s.getServer(0, types.GameServerTypeAll, serverInfo.GetPingTaiFuServerId())
	if crossServer == nil {
		crossServer = s.getServer(0, types.GameServerTypeAll, 1)
	}
	if crossServer != nil && crossServer.GetStatus() == types.GameServerStatusNormal {
		res.ServerInfoList = append(res.ServerInfoList, ConvertCrossServerInfo(crossServer))
	}

	return
}

func (s *CenterServer) Refresh(ctx context.Context, req *centerpb.RefreshServerInfoListRequest) (res *centerpb.RefreshServerInfoListResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	log.WithFields(
		log.Fields{}).Info("server:刷新游戏服务器")
	platform := req.GetPlatform()
	mergeRecordList, err := s.serverStore.GetMergeList(platform)

	//清空合服记录
	s.clearMergeRecords(platform)
	for _, mergeRecordEntity := range mergeRecordList {
		mergeRecordInfo := newMergeRecordInfo()
		mergeRecordInfo.FromEntity(mergeRecordEntity)
		s.addMergeRecord(mergeRecordInfo)
	}
	now := timeutils.TimeToMillisecond(time.Now())
	for i := types.GameServerTypeSingle; i <= types.GameServerTypeRegion; i++ {
		serverEntityList, err := s.serverStore.GetServerList(platform, i)
		if err != nil {
			return nil, err
		}
		//清空服务器列表
		s.clearServerList(platform, i)
		for _, serverEntity := range serverEntityList {
			serverInfo := newServerInfo()
			serverInfo.FromEntity(serverEntity)
			serverInfo.lastHeartbeatTime = now
			s.addServer(serverInfo)
		}
	}

	for i := types.GameServerTypeAll; i <= types.GameServerTypeChenZhan; i++ {
		serverEntityList, err := s.serverStore.GetServerList(0, i)
		if err != nil {
			return nil, err
		}
		//清空服务器列表
		s.clearCrossList(i)
		for _, serverEntity := range serverEntityList {
			serverInfo := newServerInfo()
			serverInfo.FromEntity(serverEntity)
			serverInfo.lastHeartbeatTime = now
			s.addServer(serverInfo)
		}
	}

	res = &centerpb.RefreshServerInfoListResponse{}
	res.Platform = platform
	log.WithFields(
		log.Fields{
			"platform": platform,
		}).Info("server:刷新游戏服务器成功")
	return
}

func (s *CenterServer) RefreshSDK(ctx context.Context, req *centerpb.RefreshSDKListRequest) (res *centerpb.RefreshSDKListResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	log.Info("server:刷新sdk")
	err = s.refreshSDK()
	if err != nil {
		return
	}
	res = &centerpb.RefreshSDKListResponse{}
	log.Info("server:刷新sdk成功")
	return
}

func (s *CenterServer) RefreshMarryPrice(ctx context.Context, req *centerpb.RefreshMarryPriceListRequest) (res *centerpb.RefreshMarryPriceListResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	platform := req.GetPlatform()
	log.WithFields(
		log.Fields{
			"platform": platform,
		}).Info("server:刷新结婚配置")
	e, err := s.marryPriceStore.Get(platform)
	if err != nil {
		return
	}
	if e == nil {
		delete(s.marryPriceMap, platform)
	} else {
		s.marryPriceMap[platform] = e.KindType
	}

	res = &centerpb.RefreshMarryPriceListResponse{}
	res.Platform = platform
	log.WithFields(
		log.Fields{
			"platform": platform,
		}).Info("server:刷新结婚配置成功")
	return
}

func (s *CenterServer) Register(ctx context.Context, req *centerpb.ServerRegisterRequest) (res *centerpb.ServerRegisterResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	log.WithFields(
		log.Fields{}).Debug("server:注册服务器")
	serverType := req.GetServerType()
	gameServerType := types.GameServerType(serverType)
	platform := req.GetPlatform()
	serverId := req.GetServerId()
	serverIp := req.GetServerIp()
	serverPort := req.GetServerPort()

	//参数不对
	if !gameServerType.Valid() {
		log.WithFields(
			log.Fields{
				"serverType": gameServerType.String(),
				"platform":   platform,
				"serverId":   serverId,
				"serverIp":   serverIp,
				"serverPort": serverPort,
			}).Warn("server:注册服务器,失败")
		err = status.Errorf(codes.InvalidArgument, "无效服务器类型: %d", serverType)
		return
	}

	if !s.isValidPlatform(platform) {
		log.WithFields(
			log.Fields{
				"serverType": gameServerType.String(),
				"platform":   platform,
				"serverId":   serverId,
				"serverIp":   serverIp,
				"serverPort": serverPort,
			}).Warn("server:注册服务器,失败")
		//TODO 返回参数不对
		err = status.Errorf(codes.InvalidArgument, "无效平台: %d", platform)
		return
	}

	serverInfo := s.getServer(platform, gameServerType, serverId)
	if serverInfo == nil {
		log.WithFields(
			log.Fields{
				"serverType": gameServerType.String(),
				"platform":   platform,
				"serverId":   serverId,
				"serverIp":   serverIp,
				"serverPort": serverPort,
			}).Warn("server:注册服务器,失败")
		err = status.Errorf(codes.InvalidArgument, "无效服务器id: %d", serverId)
		return
	}

	serverInfo.register(serverIp, serverPort)

	// if gameServerType != types.GameServerTypeSingle {
	// 	//注册
	// 	serverInfo.register(serverIp, serverPort)
	// } else {

	// 	//所有合服的
	// 	infoList := s.getGameServerList(platform)
	// 	for _, info := range infoList {
	// 		if info.serverIp == serverInfo.serverIp && info.serverPort == serverInfo.serverPort {
	// 			info.register(serverIp, serverPort)
	// 		}
	// 	}
	// }
	marryKindType := s.marryPriceMap[platform]
	res = &centerpb.ServerRegisterResponse{}
	res.Id = serverInfo.GetId()
	res.StartTime = serverInfo.GetStartTime()
	res.MarryKindType = marryKindType
	tradeServerIp := s.serverConfig.tradeIp
	res.TradeServerIp = tradeServerIp
	//获取所有的sdk
	for sdkType, tempPlatform := range s.sdkPlatformMap {
		if platform == tempPlatform {
			res.SdkList = append(res.SdkList, sdkType)
		}
	}
	platformSetting, ok := s.platformSettingMap[platform]
	if ok {
		res.PlatformSetting = &centerpb.PlatformSetting{}
		res.PlatformSetting.MarryKindType = platformSetting.settingInfo.MarrySet
		res.PlatformSetting.TradeFlag = platformSetting.settingInfo.JiaoYiHangFlag
		res.PlatformSetting.AllianceFlag = platformSetting.settingInfo.AllianceWarehouseFlag
		res.PlatformSetting.XianJinFlag = platformSetting.settingInfo.XianJinFlag
		res.PlatformSetting.NeiWanJiaoYiFlag = platformSetting.settingInfo.NeiWanJiaoYiFlag
		res.PlatformSetting.ZhiZuanFlag = platformSetting.settingInfo.ZhiZuanFlag
	}
	platformChatset, ok := s.platformChatSetMap[platform]
	if ok {
		res.PlatformChatSetting = &centerpb.PlatformChatSetting{}
		res.PlatformChatSetting.MinVip = platformChatset.minVip
		res.PlatformChatSetting.MinPlayerLevel = platformChatset.minPlayerLevel
		res.PlatformChatSetting.WorldVip = platformChatset.worldVip
		res.PlatformChatSetting.WorldPlayerLevel = platformChatset.worldPlayerLevel
		res.PlatformChatSetting.TeamPlayerLevel = platformChatset.teamPlayerLevel
		res.PlatformChatSetting.TeamVip = platformChatset.teamVip
		res.PlatformChatSetting.AllianceVip = platformChatset.guildVip
		res.PlatformChatSetting.AlliancePlayerLevel = platformChatset.guildPlayerLevel
		res.PlatformChatSetting.PVip = platformChatset.pChatVip
		res.PlatformChatSetting.PPlayerLevel = platformChatset.pChatPlayerLevel
	}
	log.WithFields(
		log.Fields{
			"serverType": gameServerType.String(),
			"platform":   platform,
			"serverId":   serverId,
			"name":       serverInfo.GetName(),
			"serverIp":   serverIp,
			"serverPort": serverPort,
		}).Debug("server:注册服务器,成功")
	return
}

func (s *CenterServer) Unregister(ctx context.Context, req *centerpb.ServerUnregisterRequest) (res *centerpb.ServerUnregisterResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	serverId := req.GetServerId()
	log.WithFields(
		log.Fields{
			"serverId": serverId,
		}).Debug("server:取消注册服务器")

	serverInfo := s.getServerById(serverId)
	if serverInfo == nil {
		log.WithFields(
			log.Fields{
				"serverId": serverId,
			}).Warn("server:取消注册服务器,服务器不存在")
		err = status.Errorf(codes.NotFound, "无效服务器id: %d", serverId)
		return
	}

	serverInfo.maintain()
	// gameServerType := serverInfo.GetServerType()
	// if gameServerType != types.GameServerTypeSingle {
	// 	//注册
	// 	serverInfo.maintain()
	// } else {
	// 	//所有合服的
	// 	infoList := s.getGameServerList(serverInfo.GetPlatform())
	// 	for _, info := range infoList {
	// 		if info.serverIp == serverInfo.serverIp && info.serverPort == serverInfo.serverPort {
	// 			info.maintain()
	// 		}
	// 	}
	// }
	res = &centerpb.ServerUnregisterResponse{}
	res.ServerId = serverId

	log.WithFields(
		log.Fields{
			"serverId": serverId,
			"name":     serverInfo.GetName(),
		}).Debug("server:取消注册服务器,成功")
	return
}

func (s *CenterServer) SyncPlayerServerInfo(ctx context.Context, req *centerpb.PlayerInfoSyncRequest) (res *centerpb.PlayerInfoSyncResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerEntity, err := s.playerStore.GetPlayerServerEntity(req.PlayerInfo.UserId, req.PlayerInfo.ServerId)
	if err != nil {
		return
	}

	now := timeutils.TimeToMillisecond(time.Now())
	if playerEntity == nil {
		playerEntity = store.NewPlayerEntity()
		playerEntity.CreateTime = now
	}

	playerEntity.ServerId = req.PlayerInfo.ServerId
	playerEntity.UserId = req.PlayerInfo.UserId
	playerEntity.PlayerId = req.PlayerInfo.PlayerId

	playerEntity.PlayerName = req.PlayerInfo.PlayerName
	playerEntity.Role = req.PlayerInfo.Role
	playerEntity.Sex = req.PlayerInfo.Sex
	playerEntity.Level = req.PlayerInfo.Level
	playerEntity.ZhuanShu = req.PlayerInfo.ZhuanShu
	playerEntity.UpdateTime = now
	s.playerStore.Save(playerEntity)
	res = &centerpb.PlayerInfoSyncResponse{}
	return
}

func (s *CenterServer) GetPlayerServerList(ctx context.Context, req *centerpb.PlayerListRequest) (res *centerpb.PlayerListResponse, err error) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	playerEntityList, err := s.playerStore.GetPlayerServerList(req.UserId)
	if err != nil {
		return
	}
	res = &centerpb.PlayerListResponse{}
	for _, playerEntity := range playerEntityList {
		playerInfo := playerEntity.ConvertToGrpcFormat()
		res.PlayerList = append(res.PlayerList, playerInfo)
	}
	return
}

type userCache struct {
	Token          string `json:"token"`
	Platform       int32  `json:"platform"`
	DevicePlatform int32  `json:"devicePlatform"`
	PlatformUserId string `json:"platformUserId"`
	Gm             int32  `json:"gm"`
}

func (s *CenterServer) login(platform int32, devicePlatform int32, userId int64, platformUserId string, gm int32) (t string, expiredTime int64, err error) {
	now := timeutils.TimeToMillisecond(time.Now())
	expiredTime += now + s.options.ExpiredTime
	claims := &jwt.StandardClaims{}
	claims.ExpiresAt = expiredTime
	claims.Issuer = fmt.Sprintf("%d", userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	t, err = token.SignedString(s.key)
	if err != nil {
		return
	}
	//保存redis
	conn := s.rs.Pool().Get()
	err = conn.Err()
	if err != nil {
		return
	}
	defer conn.Close()

	userTokenKey := getUserTokenRedisKey(userId)
	uCache := &userCache{
		Token:          t,
		Platform:       platform,
		DevicePlatform: devicePlatform,
		PlatformUserId: platformUserId,
		Gm:             gm,
	}
	cacheContent, err := json.Marshal(uCache)
	if err != nil {
		return
	}
	ok, err := redis.String(conn.Do("setex", userTokenKey, s.options.ExpiredTime/1000, cacheContent))
	if err != nil {
		return
	}
	if ok != coreredis.OK {
		err = fmt.Errorf("redis set failed %s", ok)
		return
	}
	return
}

var (
	userKey = "fgame.user"
)

func getUserTokenRedisKey(userId int64) string {
	return coreredis.Combine(userKey, fmt.Sprintf("%d", userId))
}

func (s *CenterServer) SelfLogin(ctx context.Context, req *centerpb.SelfLoginRequest) (res *centerpb.SelfLoginResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	name := req.GetName()
	password := req.GetPassword()
	if len(name) <= 0 || len(password) <= 0 {
		res = &centerpb.SelfLoginResponse{}
		res.PlatformUserId = ""
		return
	}
	userEntity, err := s.userStore.GetUserByNameAndPassword(name, password)
	if err != nil {
		return
	}
	if userEntity == nil {
		res = &centerpb.SelfLoginResponse{}
		res.PlatformUserId = ""
		return
	}
	res = &centerpb.SelfLoginResponse{}
	res.PlatformUserId = userEntity.PlatformUserId
	res.Platform = userEntity.Platform
	return
}

func (s *CenterServer) Login(ctx context.Context, req *centerpb.LoginRequest) (res *centerpb.LoginResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	platform := req.GetPlatform()
	platformUserId := req.GetUserId()
	devicePlatform := req.GetDevicePlatform()
	ip := req.GetIp()
	if len(ip) > 0 {
		flag, err := s.isIpForbid(ip)
		if err != nil {
			return nil, err
		}
		if flag {
			res = &centerpb.LoginResponse{}
			res.Token = ""
			res.UserId = 0
			res.ExpiredTime = 0
			res.Gm = 0
			return res, nil
		}
	}
	userEntity, err := s.userStore.GetUserByPlatform(platform, platformUserId)
	if err != nil {
		return
	}
	if userEntity == nil {
		userEntity, err = s.userStore.RegisterPlatformUser(platform, platformUserId, "", "")
		if err != nil {
			return
		}
	}
	userId := userEntity.Id
	gm := userEntity.Gm
	t, expiredTime, err := s.login(platform, devicePlatform, userId, platformUserId, gm)
	if err != nil {
		return
	}
	res = &centerpb.LoginResponse{}
	res.Token = t
	res.UserId = userId
	res.ExpiredTime = expiredTime
	res.Gm = gm
	return
}

func (s *CenterServer) LoginVerify(ctx context.Context, req *centerpb.LoginVerifyRequest) (res *centerpb.LoginVerifyResponse, err error) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	token := req.GetToken()
	claims := &jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) { return s.key, nil })
	if err != nil {
		log.WithFields(
			log.Fields{
				"token": token,
			}).Error("server:登陆验证,解析错误")
		return
	}

	idStr := claims.Issuer
	res = &centerpb.LoginVerifyResponse{}
	if len(idStr) == 0 {
		log.WithFields(
			log.Fields{
				"token": token,
			}).Error("server:登陆验证,id是空")
		res.UserId = 0
		return
	}

	userId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"token": token,
				"id":    idStr,
			}).Error("server:登陆验证,id不是数字")
		return
	}

	//保存redis
	conn := s.rs.Pool().Get()
	err = conn.Err()
	if err != nil {
		log.WithFields(
			log.Fields{
				"token": token,
				"id":    idStr,
				"err":   err,
			}).Error("server:登陆验证,redis连接错误")
		return
	}
	defer conn.Close()

	userTokenKey := getUserTokenRedisKey(userId)

	cUserBytes, err := redis.Bytes(conn.Do("get", userTokenKey))
	if err != nil {
		if err == redis.ErrNil {
			log.WithFields(
				log.Fields{
					"token": token,
					"id":    idStr,
				}).Warn("server:登陆验证,登陆过期")
			res.UserId = 0
			return
		}
		log.WithFields(
			log.Fields{
				"token": token,
				"id":    idStr,
				"err":   err,
			}).Warn("server:登陆验证,redis连接错误")
		return
	}
	uCache := &userCache{}
	err = json.Unmarshal(cUserBytes, uCache)
	if err != nil {
		log.WithFields(
			log.Fields{
				"token": token,
				"id":    idStr,
				"err":   err,
				"cache": string(cUserBytes),
			}).Error("server:登陆验证,解析错误")
		return
	}

	if uCache.Token != token {
		log.WithFields(
			log.Fields{
				"token":      token,
				"id":         idStr,
				"cacheToken": uCache.Token,
			}).Warn("server:登陆验证,token不一致")
		res = &centerpb.LoginVerifyResponse{}
		res.UserId = 0
		return
	}

	originServerId := req.GetOriginServerId()
	serverId := req.GetServerId()

	if originServerId != serverId {
		centerPlatformId := s.sdkPlatformMap[uCache.Platform]
		mergeRecordInfo := s.getMergeRecord(centerPlatformId, originServerId)
		if mergeRecordInfo == nil {
			log.WithFields(
				log.Fields{
					"token":            token,
					"id":               idStr,
					"centerPlatformId": centerPlatformId,
					"originServerId":   originServerId,
					"serverId":         serverId,
				}).Warn("server:登陆验证,登陆服务器不对")
			res = &centerpb.LoginVerifyResponse{}
			res.UserId = 0
			return
		}
		if mergeRecordInfo.GetFinalserverId() != serverId {
			log.WithFields(
				log.Fields{
					"token":              token,
					"id":                 idStr,
					"originServerId":     originServerId,
					"serverId":           serverId,
					"mergeFinalServerId": mergeRecordInfo.GetFinalserverId(),
				}).Warn("server:登陆验证,登陆服务器不对")
			res = &centerpb.LoginVerifyResponse{}
			res.UserId = 0
			return
		}
	}
	res = &centerpb.LoginVerifyResponse{}
	res.UserId = userId
	res.SdkType = uCache.Platform
	res.DevicePlatformType = uCache.DevicePlatform
	res.PlatformUserId = uCache.PlatformUserId
	res.Gm = uCache.Gm
	iosVersion := s.clientVersion.iosVersion
	res.IosVersion = iosVersion
	androidVersion := s.clientVersion.androidVersion
	res.AndroidVersion = androidVersion
	return
}

func (s *CenterServer) GMLogin(ctx context.Context, req *centerpb.GMLoginRequest) (res *centerpb.GMLoginResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	platform := req.GetSdkType()
	platformUserId := req.GetUserId()
	name := req.GetName()
	password := req.GetPassword()
	userEntity, err := s.userStore.GetUserByPlatform(platform, platformUserId)
	if err != nil {
		return
	}
	if userEntity == nil {
		userEntity, err = s.userStore.RegisterPlatformUser(platform, platformUserId, name, password)
		if err != nil {
			return
		}
	}
	userId := userEntity.Id
	gm := userEntity.Gm

	t, expiredTime, err := s.login(platform, 0, userId, platformUserId, gm)
	if err != nil {
		return
	}
	res = &centerpb.GMLoginResponse{}
	res.Token = t
	res.UserId = userId
	res.ExpiredTime = expiredTime

	return
}
func (s *CenterServer) RefreshNotice(ctx context.Context, req *centerpb.RefreshNoticeRequest) (res *centerpb.RefreshNoticeResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	noticeMap := make(map[int32]*NoticeInfo)
	noticeEntityList, err := s.noticeStore.GetAll()
	if err != nil {
		return
	}
	for _, noticeEntity := range noticeEntityList {
		noticeInfo := newNoticeInfo()
		noticeInfo.FromEntity(noticeEntity)
		noticeMap[noticeInfo.platformId] = noticeInfo
	}
	s.noticeMap = noticeMap
	//重新加载
	res = &centerpb.RefreshNoticeResponse{}
	return
}

func (s *CenterServer) GetNotice(ctx context.Context, req *centerpb.NoticeRequest) (res *centerpb.NoticeResponse, err error) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	log.WithFields(
		log.Fields{}).Info("server:获取公告")
	platform := req.Platform
	centerPlatformId := s.sdkPlatformMap[platform]
	notice, ok := s.noticeMap[centerPlatformId]
	//重新加载
	res = &centerpb.NoticeResponse{}
	if !ok {
		defaultNotice, ok := s.noticeMap[0]
		if !ok {
			res.Notice = ""
			return
		}
		res.Notice = defaultNotice.GetContent()
		return
	}
	res.Notice = notice.GetContent()
	return
}

//获取交易服务器列表
func (s *CenterServer) GetTradeServerList(ctx context.Context, req *centerpb.TradeServerListRequest) (res *centerpb.TradeServerListResponse, err error) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	log.WithFields(
		log.Fields{}).Info("server:获取交易服务器列表")

	res = &centerpb.TradeServerListResponse{}

	for _, serverMap := range s.serverCategoryMap {
		gameServerMap := serverMap[types.GameServerTypeSingle]
		for _, serverInfo := range gameServerMap {
			finalServerInfo := s.getMergeServerInfo(serverInfo.platform, serverInfo.serverId)
			tradeServerInfo := &centerpb.TradeServerInfo{}
			tradeServerInfo.Platform = serverInfo.platform
			tradeServerInfo.ServerId = serverInfo.serverId
			if finalServerInfo == nil {
				tradeServerInfo.RegionId = serverInfo.parentServerId
			} else {
				tradeServerInfo.RegionId = finalServerInfo.parentServerId
			}

			res.TradeServerInfoList = append(res.TradeServerInfoList, tradeServerInfo)
		}
	}

	return
}

func NewCenterServer(options *CenterOptions) (ss *CenterServer, err error) {
	ss = &CenterServer{}
	ss.options = options
	err = ss.init()
	if err != nil {
		return
	}
	return
}
