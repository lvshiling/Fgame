package treasurebox

import (
	"context"
	centertypes "fgame/fgame/center/types"
	treasureboxclient "fgame/fgame/cross/treasurebox/client"
	arenatemplate "fgame/fgame/game/arena/template"
	"fgame/fgame/game/center/center"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/global"
	"fmt"
	"sort"
	"sync"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BoxLogInfo struct {
	ServerId   int32
	PlayerName string
	LastTime   int64
	ItemMap    map[int32]int32
}

func newboxLogInfo() *BoxLogInfo {
	d := &BoxLogInfo{}
	return d
}

//记录排序
type BoxLogInfoList []*BoxLogInfo

func (adl BoxLogInfoList) Len() int {
	return len(adl)
}

func (adl BoxLogInfoList) Less(i, j int) bool {
	return adl[i].LastTime < adl[j].LastTime
}

func (adl BoxLogInfoList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

//跨服宝箱接口处理
type TreasureBoxService interface {
	Heartbeat()
	Star() (err error)
	//获取跨服宝箱日志列表
	GetTreasureBoxLogList(logTime int64) []*BoxLogInfo
	//开宝箱
	OpenTreasureBox(serverId int32, playerName string, itemList []*droptemplate.DropItemData)
}

type treasureBoxService struct {
	logBoxList []*BoxLogInfo
	//读写锁
	rwm               sync.RWMutex
	treasureboxClient treasureboxclient.TreasureBoxClient
}

//初始化
func (rs *treasureBoxService) init() (err error) {
	rs.logBoxList = make([]*BoxLogInfo, 0, 8)

	return
}

//获取跨服宝箱日志
func (rs *treasureBoxService) GetTreasureBoxLogList(logTime int64) (boxLogList []*BoxLogInfo) {
	rs.rwm.RLock()
	defer rs.rwm.RUnlock()

	pos := int(-1)
	for index, logBox := range rs.logBoxList {
		if logBox.LastTime > logTime {
			pos = index
			break
		}
	}
	addLen := len(rs.logBoxList)
	if pos == -1 {
		return
	}
	if addLen-pos == 1 {
		boxLogList = append(boxLogList, rs.logBoxList[pos])
	} else {
		boxLogList = append(boxLogList, rs.logBoxList[pos:addLen]...)
	}
	return
}

func (rs *treasureBoxService) OpenTreasureBox(serverId int32, playerName string, itemList []*droptemplate.DropItemData) {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	boxLogInfo := newboxLogInfo()
	boxLogInfo.ServerId = serverId
	boxLogInfo.PlayerName = playerName
	boxLogInfo.LastTime = now
	boxLogInfo.ItemMap = make(map[int32]int32)
	for _, itemData := range itemList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		boxLogInfo.ItemMap[itemId] = num
	}

	rs.logBoxList = append(rs.logBoxList, boxLogInfo)

	sort.Sort(BoxLogInfoList(rs.logBoxList))

	logLen := len(rs.logBoxList)

	riZhiMax := int(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RiZhiMax)
	if logLen > riZhiMax {
		rs.logBoxList = rs.logBoxList[:riZhiMax]
	}

	//TODO 超时
	ctx := context.TODO()
	req := convertFromOpenInfo(serverId, playerName, now, itemList)
	rs.treasureboxClient.OpenTreasureBox(ctx, req)

}

func (rs *treasureBoxService) Star() (err error) {
	conn := center.GetCenterService().GetCross(centertypes.GameServerTypePlatform)
	if conn == nil {
		return fmt.Errorf("treasurebox:跨服连接不存在")
	}
	//TODO 修改可能连接变化了
	rs.treasureboxClient = treasureboxclient.NewTreasureBoxClient(conn)
	err = rs.syncRemoteLogList()
	if err != nil {
		return
	}
	return
}

func (rs *treasureBoxService) syncRemoteLogList() (err error) {
	//TODO 超时
	ctx := context.TODO()
	resp, err := rs.treasureboxClient.GetTreasureBoxLogList(ctx)
	if err != nil {
		return
	}
	rs.logBoxList = convertFromBoxLogInfoList(resp.LogList)
	return
}

func (rs *treasureBoxService) resetClient() (err error) {
	conn := center.GetCenterService().GetCross(centertypes.GameServerTypePlatform)
	if conn == nil {
		return fmt.Errorf("treasurebox:跨服连接不存在")
	}

	rs.treasureboxClient = treasureboxclient.NewTreasureBoxClient(conn)

	return
}

//心跳
func (rs *treasureBoxService) Heartbeat() {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()

	err := rs.syncRemoteLogList()
	if err != nil {
		sta := status.Convert(err)
		if sta.Code() == codes.Canceled {
			//重新获取
			log.WithFields(
				log.Fields{
					"err": err,
				}).Warn("treasurebox:同步跨服,重新获取客户端")
			err = rs.resetClient()
			if err != nil {
				log.WithFields(
					log.Fields{
						"err": err,
					}).Warn("treasurebox:同步跨服,失败")
				return
			}
			err = rs.syncRemoteLogList()
			if err != nil {
				log.WithFields(
					log.Fields{
						"err": err,
					}).Warn("treasurebox:同步跨服,失败")
				return
			}
		}
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("treasurebox:同步跨服,失败")
	}
}

var (
	once sync.Once
	cs   *treasureBoxService
)

func Init() (err error) {
	once.Do(func() {
		cs = &treasureBoxService{}
		err = cs.init()
	})
	return err
}

func GetTreasureBoxService() TreasureBoxService {
	return cs
}
