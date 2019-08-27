package rank

import (
	"context"
	centertypes "fgame/fgame/center/types"
	"fgame/fgame/game/center/center"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	rankeventtypes "fgame/fgame/game/rank/event/types"
	rankobj "fgame/fgame/game/rank/obj"
	ranktypes "fgame/fgame/game/rank/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	rankclient "fgame/fgame/rank/client"
	rankpb "fgame/fgame/rank/protocol/pb"
	"fmt"
	"sync"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//排行榜接口处理
type RankService interface {
	Heartbeat()
	//启动排行
	Star() (err error)
	//获取我的排名位置
	GetMyRankPos(rankClassType ranktypes.RankClassType, groupId int32, typ ranktypes.RankType, id int64) (pos int32)
	//排行榜第一
	GetRankFirstId(rankClassType ranktypes.RankClassType, groupId int32, rankType ranktypes.RankType) (playerId int64)
	//获取战力排名
	GetForceListByPage(rankClassType ranktypes.RankClassType, groupId int32, page int32) ([]*rankentity.PlayerForceData, int64)
	//获取兵魂排名
	GetWeaponListByPage(rankClassType ranktypes.RankClassType, groupId int32, page int32) ([]*rankentity.PlayerWeaponData, int64)
	//获取帮派排名
	GetGangListByPage(rankClassType ranktypes.RankClassType, groupId int32, page int32) ([]*rankentity.PlayerGangData, int64)
	//获取仙盟区排名前size
	GetRankGangFew(size int32) []*rankentity.PlayerGangData
	//属性排行榜列表
	GetPropertyListByPage(rankType ranktypes.RankType, rankClassType ranktypes.RankClassType, groupId int32, page int32) ([]*rankentity.PlayerPropertyData, int64)
	//进阶系统排行榜列表
	GetOrderListByPage(rankType ranktypes.RankType, rankClassType ranktypes.RankClassType, groupId int32, page int32) ([]*rankentity.PlayerOrderData, int64)

	//获取榜单所有排行
	GetRankingInfoList(rankClassType ranktypes.RankClassType, rankType ranktypes.RankType, groupId int32) []*ranktypes.RankingInfo
	//注册活动排行榜
	RegisterActivityRank(rankType ranktypes.RankType, config *ranktypes.RankConfig) (err error)
	//立即刷新榜单
	UpdateRankData(rankClassType ranktypes.RankClassType, typ ranktypes.RankType, groupId int32) (err error)
}

type rankService struct {
	rankClient rankclient.RankClient
	//本服排行数据
	rankData *rankobj.RankMap
	//本区排行榜数据
	areaRankData *rankobj.RankMap
	//本服活动排行数据
	activityRankData *rankobj.RankActivityMap
	//读写锁
	rwm sync.RWMutex
	//排行榜时间
	rankTimeMap map[ranktypes.RankClassType]map[ranktypes.RankType]*RankTimeObject
}

//初始化
func (rs *rankService) init() (err error) {
	rs.rankTimeMap = make(map[ranktypes.RankClassType]map[ranktypes.RankType]*RankTimeObject)

	// //排行榜周期重置
	// err = rs.loadRankTime()
	// if err != nil {
	// 	return
	// }

	err = rs.initLocal()
	if err != nil {
		return
	}
	err = rs.initArea()
	if err != nil {
		return
	}

	return
}

//初始化本服的
func (rs *rankService) initLocal() (err error) {

	rankTypeMap := ranktypes.RankClassTypeLocal.GetRankTypeMap()
	for rankType, _ := range rankTypeMap {
		config := ranktypes.NewLocalDefaultConfig()
		if rs.rankData == nil {
			rs.rankData = rankobj.NewRankMap(config.RefreshTime)
		}

		// 注册普通榜单
		rs.rankData.RegisterRank(rankType, config)

		// //注册排行榜重置检查
		// rs.registRankTimeObject(config.ClassType, rankType)
	}

	err = rs.rankData.Init()
	if err != nil {
		return
	}

	return
}

//初始化本区的
func (rs *rankService) initArea() (err error) {
	// 注册区普通榜单
	rankTypeMap := ranktypes.RankClassTypeArea.GetRankTypeMap()
	for rankType, _ := range rankTypeMap {
		config := ranktypes.NewAreaDefaultConfig()
		if rs.areaRankData == nil {
			rs.areaRankData = rankobj.NewRankMap(config.RefreshTime)
		}

		rs.areaRankData.RegisterRank(rankType, config)
	}

	err = rs.areaRankData.Init()
	if err != nil {
		return
	}
	return
}

//加载排行刷新时间
func (rs *rankService) loadRankTime() (err error) {
	entity, err := dao.GetRankDao().GetRankTimeEntity(global.GetGame().GetServerIndex())
	if err != nil {
		return
	}
	if entity != nil {
		timeObj := NewRankTimeObject()
		timeObj.FromEntity(entity)
		subMap, ok := rs.rankTimeMap[timeObj.classRankType]
		if !ok {
			subMap = make(map[ranktypes.RankType]*RankTimeObject)
			rs.rankTimeMap[timeObj.classRankType] = subMap
		}
		subMap[timeObj.rankType] = timeObj
	}

	return
}

//第一次初始化
func (rs *rankService) registRankTimeObject(classRankType ranktypes.RankClassType, rankType ranktypes.RankType) {

	po := rs.getRankTimeObject(classRankType, rankType)
	if po != nil {
		return
	}

	po = NewRankTimeObject()
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	po.id = id
	po.serverId = global.GetGame().GetServerIndex()
	po.classRankType = classRankType
	po.rankType = rankType
	po.thisTime = 0
	po.createTime = now
	po.SetModified()

	subMap, ok := rs.rankTimeMap[classRankType]
	if !ok {
		subMap = make(map[ranktypes.RankType]*RankTimeObject)
		rs.rankTimeMap[classRankType] = subMap
	}
	subMap[rankType] = po
}

func (rs *rankService) getRankTimeObject(classRankType ranktypes.RankClassType, rankType ranktypes.RankType) *RankTimeObject {
	subMap, ok := rs.rankTimeMap[classRankType]
	if !ok {
		return nil
	}
	obj, ok := subMap[rankType]
	if !ok {
		return nil
	}

	return obj
}

func (rs *rankService) startArea() (err error) {
	conn := center.GetCenterService().GetCross(centertypes.GameServerTypeGroup)
	if conn == nil {
		return fmt.Errorf("rank:跨服连接不存在")
	}
	//TODO 修改可能连接变化了
	rs.rankClient = rankclient.NewRankClient(conn)
	err = rs.syncRemoteRankList()
	if err != nil {
		return
	}
	return
}

//注册排行榜
func (rs *rankService) RegisterActivityRank(rankType ranktypes.RankType, config *ranktypes.RankConfig) (err error) {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()

	if rs.activityRankData == nil {
		rs.activityRankData = rankobj.NewRankActivityMap(config.RefreshTime)
		err = rs.activityRankData.Init()
		if err != nil {
			return
		}

	}

	flag := rs.activityRankData.RegisterRank(rankType, config)
	if !flag {
		return fmt.Errorf("排行榜已存在")
	}

	return
}

//启动排行
func (rs *rankService) Star() (err error) {
	//本服排行榜启动
	err = rs.rankData.Star()
	if err != nil {
		return
	}

	//本区排行榜启动
	err = rs.startArea()
	if err != nil {
		return
	}

	//合服
	isMerge := merge.GetMergeService().IsMerge()
	if isMerge {
		rs.mergeServerLocalRank()
	}

	return
}

func (rs *rankService) mergeServerLocalRank() (err error) {
	//本服排行榜重新排 覆盖redis
	rankTypeMap := ranktypes.RankClassTypeLocal.GetRankTypeMap()
	for rankType, _ := range rankTypeMap {
		dataLocal := cs.rankData.GetRankTypeData(rankType)
		dataLocal.ResetRankTime()
	}
	err = rs.rankData.UpdateRank()
	if err != nil {
		return
	}
	return
}

//获取仙盟区排名前size
func (rs *rankService) GetRankGangFew(size int32) []*rankentity.PlayerGangData {
	rankTypeData := rs.getRankData(ranktypes.RankClassTypeArea, ranktypes.RankTypeGang, 0)
	gangRank, ok := rankTypeData.(*rankobj.GangRank)
	if !ok {
		return nil
	}
	rankList, _ := gangRank.GetListAndTime()
	len := len(rankList)
	if len == 0 {
		return nil
	}
	addLen := size
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[0:addLen]
}

func (rs *rankService) convertFromGrpc(resp *rankpb.RankListResponse) {
	//战力转换
	rankForceData := rs.areaRankData.GetRankTypeData(ranktypes.RankTypeForce)
	forceRank, ok := rankForceData.(*rankobj.ForceRank)
	if !ok {
		return
	}
	if resp.Force != nil {
		forceRank.ConvertFromForceInfo(resp.Force)
	}
	//帮派转换
	rankGangData := rs.areaRankData.GetRankTypeData(ranktypes.RankTypeGang)
	gangRank, ok := rankGangData.(*rankobj.GangRank)
	if !ok {
		return
	}
	if resp.Gang != nil {
		gangRank.ConvertFromGangInfo(resp.Gang)
	}
	//坐骑转换
	rankMountData := rs.areaRankData.GetRankTypeData(ranktypes.RankTypeMount)
	mountRank, ok := rankMountData.(*rankobj.MountRank)
	if !ok {
		return
	}
	if resp.Mount != nil {
		mountRank.ConvertFromMountInfo(resp.Mount)
	}

	//转换战翼
	rankWingData := rs.areaRankData.GetRankTypeData(ranktypes.RankTypeWing)
	wingRank, ok := rankWingData.(*rankobj.WingRank)
	if !ok {
		return
	}
	if resp.Wing != nil {
		wingRank.ConvertFromWingInfo(resp.Wing)
	}

	//转换兵魂
	rankWeaponData := rs.areaRankData.GetRankTypeData(ranktypes.RankTypeWeapon)
	weaponRank, ok := rankWeaponData.(*rankobj.WeaponRank)
	if !ok {
		return
	}
	if resp.Weapon != nil {
		weaponRank.ConvertFromWeaponInfo(resp.Weapon)
	}

	//转换护体盾
	rankBodyShieldData := rs.areaRankData.GetRankTypeData(ranktypes.RankTypeBodyShield)
	bodyShieldRank, ok := rankBodyShieldData.(*rankobj.BodyShieldRank)
	if !ok {
		return
	}
	if resp.BodyShield != nil {
		bodyShieldRank.ConvertFromBodyShieldInfo(resp.BodyShield)
	}

	//转换身法
	rankShenFaData := rs.areaRankData.GetRankTypeData(ranktypes.RankTypeShenFa)
	shenFaRank, ok := rankShenFaData.(*rankobj.ShenFaRank)
	if !ok {
		return
	}
	if resp.ShenFa != nil {
		shenFaRank.ConvertFromShenFaInfo(resp.ShenFa)
	}

	//转换领域
	rankLingYuData := rs.areaRankData.GetRankTypeData(ranktypes.RankTypeLingYu)
	lingYuRank, ok := rankLingYuData.(*rankobj.LingYuRank)
	if !ok {
		return
	}
	if resp.LingYu != nil {
		lingYuRank.ConvertFromLingYuInfo(resp.LingYu)
	}

	//转换护体仙羽
	rankFeatherData := rs.areaRankData.GetRankTypeData(ranktypes.RankTypeFeather)
	featherRank, ok := rankFeatherData.(*rankobj.FeatherRank)
	if !ok {
		return
	}
	if resp.Feather != nil {
		featherRank.ConvertFromFeatherInfo(resp.Feather)
	}

	//转换神盾尖刺
	rankShieldData := rs.areaRankData.GetRankTypeData(ranktypes.RankTypeShield)
	shieldRank, ok := rankShieldData.(*rankobj.ShieldRank)
	if !ok {
		return
	}
	if resp.Shield != nil {
		shieldRank.ConvertFromShieldInfo(resp.Shield)
	}

	//转换暗器
	rankAnQiData := rs.areaRankData.GetRankTypeData(ranktypes.RankTypeAnQi)
	anQiRank, ok := rankAnQiData.(*rankobj.AnQiRank)
	if !ok {
		return
	}
	if resp.AnQi != nil {
		anQiRank.ConvertFromShieldInfo(resp.AnQi)
	}

}

func (rs *rankService) resetClient() (err error) {
	conn := center.GetCenterService().GetCross(centertypes.GameServerTypeGroup)
	if conn == nil {
		return fmt.Errorf("rank:跨服连接不存在")
	}

	rs.rankClient = rankclient.NewRankClient(conn)

	return
}

//定时同步排行榜列表
func (rs *rankService) syncRemoteRankList() (err error) {
	//TODO 超时
	ctx := context.TODO()
	resp, err := rs.rankClient.GetRankList(ctx)
	if err != nil {
		return
	}
	rs.convertFromGrpc(resp)
	return nil
}

func (rs *rankService) getRankData(rankClassType ranktypes.RankClassType, rankType ranktypes.RankType, groupId int32) (rankData rankobj.RankTypeData) {
	switch rankClassType {
	case ranktypes.RankClassTypeArea:
		return rs.areaRankData.GetRankTypeData(rankType)
	case ranktypes.RankClassTypeLocal:
		return rs.rankData.GetRankTypeData(rankType)
	case ranktypes.RankClassTypeLocalActivity:
		if rs.activityRankData == nil {
			return
		}
		return rs.activityRankData.GetRankData(groupId)
	}

	return
}

//获取我的排名位置
func (rs *rankService) GetMyRankPos(rankClassType ranktypes.RankClassType, groupId int32, typ ranktypes.RankType, id int64) (pos int32) {
	rs.rwm.RLock()
	defer rs.rwm.RUnlock()

	rankData := rs.getRankData(rankClassType, typ, groupId)
	if rankData == nil {
		return
	}
	return rankData.GetPos(id)
}

//获取所有排名
func (rs *rankService) GetRankingInfoList(rankClassType ranktypes.RankClassType, typ ranktypes.RankType, groupId int32) []*ranktypes.RankingInfo {
	rs.rwm.RLock()
	defer rs.rwm.RUnlock()

	rankData := rs.getRankData(rankClassType, typ, groupId)
	if rankData == nil {
		return nil
	}
	return rankData.GetRankingInfoList()
}

//立即刷新榜单
func (rs *rankService) UpdateRankData(rankClassType ranktypes.RankClassType, typ ranktypes.RankType, groupId int32) (err error) {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()

	rankData := rs.getRankData(rankClassType, typ, groupId)
	if rankData == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	timestamp, flag := timeutils.GetIntervalTimeStampMs(now, rankClassType.RankRefreshTime())
	if !flag {
		return fmt.Errorf("rank:刷新时间应该是能被24小时整除的")
	}
	rankData.UpdateRankDataList(timestamp)
	return
}

//获取战力排行
func (rs *rankService) GetForceListByPage(rankClassType ranktypes.RankClassType, groupId int32, page int32) ([]*rankentity.PlayerForceData, int64) {
	rankData := rs.getRankData(rankClassType, ranktypes.RankTypeForce, groupId)
	if rankData == nil {
		return nil, 0
	}

	forceRank, ok := rankData.(*rankobj.ForceRank)
	if !ok {
		return nil, 0
	}

	pageIndex := page * ranktypes.PageLimit
	forceList, forceTime := forceRank.GetListAndTime()
	len := len(forceList)
	if pageIndex >= int32(len) {
		return nil, forceTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return forceList[pageIndex:addLen], forceTime
}

//获取兵魂排名
func (rs *rankService) GetWeaponListByPage(rankClassType ranktypes.RankClassType, groupId int32, page int32) ([]*rankentity.PlayerWeaponData, int64) {

	rankData := rs.getRankData(rankClassType, ranktypes.RankTypeWeapon, groupId)
	if rankData == nil {
		return nil, 0
	}

	weaponRank, ok := rankData.(*rankobj.WeaponRank)
	if !ok {
		return nil, 0
	}

	pageIndex := page * ranktypes.PageLimit
	rankList, rankTime := weaponRank.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

//获取帮派排名
func (rs *rankService) GetGangListByPage(rankClassType ranktypes.RankClassType, groupId int32, page int32) ([]*rankentity.PlayerGangData, int64) {

	rankData := rs.getRankData(rankClassType, ranktypes.RankTypeGang, groupId)
	if rankData == nil {
		return nil, 0
	}

	gangRank, ok := rankData.(*rankobj.GangRank)
	if !ok {
		return nil, 0
	}

	pageIndex := page * ranktypes.PageLimit
	rankList, rankTime := gangRank.GetListAndTime()
	len := len(rankList)
	if pageIndex >= int32(len) {
		return nil, rankTime
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return rankList[pageIndex:addLen], rankTime
}

func (rs *rankService) GetRankFirstId(rankClassType ranktypes.RankClassType, groupId int32, rankType ranktypes.RankType) (playerId int64) {
	rankData := rs.getRankData(rankClassType, rankType, groupId)
	if rankData == nil {
		return
	}

	return rankData.GetFirstId()
}

//获取PropertyData排名列表
func (rs *rankService) GetPropertyListByPage(rankType ranktypes.RankType, rankClassType ranktypes.RankClassType, groupId int32, page int32) ([]*rankentity.PlayerPropertyData, int64) {
	rankData := rs.getRankData(rankClassType, rankType, groupId)
	if rankData == nil {
		return nil, 0
	}

	h := GetPropertyRankListHandler(rankType)
	if h == nil {
		log.WithFields(
			log.Fields{
				"rankType": rankType,
			}).Warn("rank_property:分页获取排行榜列表处理器不存在")
		return nil, 0
	}

	return h.GetPropertyListByPage(rankData, page)

}

//获取OrderData排名列表
func (rs *rankService) GetOrderListByPage(rankType ranktypes.RankType, rankClassType ranktypes.RankClassType, groupId int32, page int32) ([]*rankentity.PlayerOrderData, int64) {
	rankData := rs.getRankData(rankClassType, rankType, groupId)
	if rankData == nil {
		return nil, 0
	}

	h := GetOrderRankListHandler(rankType)
	if h == nil {
		log.WithFields(
			log.Fields{
				"rankType": rankType,
			}).Warn("rank_order:分页获取排行榜列表处理器不存在")
		return nil, 0
	}

	return h.GetOrderListByPage(rankData, page)
}

//心跳
func (rs *rankService) Heartbeat() {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()
	//本服排行榜
	err := rs.rankData.UpdateRank()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("ranktask:整点更新排行榜列表,错误")
		return
	}
	//本服活动排行
	if rs.activityRankData != nil {
		err = rs.activityRankData.UpdateRank()
		if err != nil {
			log.WithFields(
				log.Fields{
					"error": err,
				}).Error("ranktask:整点更新活动排行榜列表,错误")
			return
		}
	}

	//定时请求远程本区排行
	err = rs.syncRemoteRankList()
	if err != nil {
		sta := status.Convert(err)
		if sta.Code() == codes.Canceled {
			//重新获取
			log.WithFields(
				log.Fields{
					"err": err,
				}).Warn("rank:同步跨服,重新获取客户端")
			err = rs.resetClient()
			if err != nil {
				log.WithFields(
					log.Fields{
						"err": err,
					}).Warn("rank:同步跨服,失败")
				return
			}
			err = rs.syncRemoteRankList()
			if err != nil {
				log.WithFields(
					log.Fields{
						"err": err,
					}).Warn("rank:同步跨服,失败")
				return
			}
		}
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("rank:同步跨服,失败")
	}

	//
	rs.checkRefreshRankTime()
}

//检查排行榜重置
func (rs *rankService) checkRefreshRankTime() {
	for _, subMap := range rs.rankTimeMap {
		for _, rankTimeObj := range subMap {
			if rankTimeObj.GetThisTime() == 0 {
				rankTimeObj.initRankTime()
				continue
			}

			if !rankTimeObj.ifRefreshRankTime() {
				continue
			}

			rankTimeObj.initRankTime()
			gameevent.Emit(rankeventtypes.RankEventTypeRankReset, rankTimeObj, nil)
		}
	}
}

//仅gm使用
func GMRankUpdate() {
	cs.rwm.Lock()
	defer cs.rwm.Unlock()
	rankTypeMap := ranktypes.RankClassTypeLocal.GetRankTypeMap()
	for rankType, _ := range rankTypeMap {
		dataLocal := cs.rankData.GetRankTypeData(rankType)
		if dataLocal == nil {
			continue
		}
		dataLocal.ResetRankTime()
	}

	for _, dataLocalActivity := range cs.activityRankData.GetRankDataMap() {
		dataLocalActivity.ResetRankTime()
	}

	//TODO 超时
	ctx := context.TODO()
	cs.rankClient.RefreshRank(ctx)
}

var (
	once sync.Once
	cs   *rankService
)

func Init() (err error) {
	once.Do(func() {
		cs = &rankService{}
		err = cs.init()
	})
	return err
}

func GetRankService() RankService {
	return cs
}
