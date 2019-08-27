package api

import (
	"net/http"
	"sort"

	centerserver "fgame/fgame/gm/gamegm/gm/center/server/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	centerServermodel "fgame/fgame/gm/gamegm/gm/center/model"
	orderservice "fgame/fgame/gm/gamegm/gm/center/order/service"
	staticservice "fgame/fgame/gm/gamegm/gm/center/staticreport/service"
	platformservice "fgame/fgame/gm/gamegm/gm/platform/service"
	userservice "fgame/fgame/gm/gamegm/gm/user/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type goldChangeRequest struct {
	PlatformId int   `json:"platformId"` //gm平台的id
	ServerId   int   `json:"serverId"`   //中心平台的服务主键id
	StartMoney int   `json:"startMoney"`
	EndMoney   int   `json:"endMoney"`
	StartTime  int64 `json:"startTime"`
	EndTime    int64 `json:"endTime"`
	GoldType   int   `json:"goldType"` //类型，1增加2减少0全部
}

type goldChangeRespon struct {
	ItemArray      []*goldChangeResponItem `json:"itemArray"`
	TotalChangeNum int64                   `json:"totalChangeNum"` //总变化数

	ServerList []*goldChangeResponServer `json:"serverList"` //服务器列表
}

type goldChangeResponItem struct {
	Reason    int                                 `json:"reason"`    //中心平台的服务器Id
	ServerMap map[int]*goldChangeResponServerItem `json:"serverMap"` //服务器数据
}

type goldChangeResponServerItem struct {
	PlayerCount int   `json:"playerCount"`
	ChangeNum   int64 `json:"changeNum"`
}
type goldChangeResponServer struct {
	ServerId         int    `json:"serverId"`
	ServerName       string `json:"serverName"`
	ServerIndex      int    `json:"serverIndex"`
	TotalPlayerCount int    `json:"playerCount"`
	TotalChangeNum   int64  `json:"changeNum"`
}

func handleGoldChangeStatic(rw http.ResponseWriter, req *http.Request) {
	form := &goldChangeRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取元宝列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &goldChangeRespon{}

	cs := centerserver.CenterServerServiceInContext(req.Context())
	serverindex := 0
	if form.ServerId > 0 {
		serverinfo, servererr := cs.GetCenterServer(int64(form.ServerId))
		if servererr != nil {
			log.WithFields(log.Fields{
				"error":    err,
				"serverid": form.ServerId,
			}).Error("获取元宝列表，获取服务器信息异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		serverindex = serverinfo.ServerId
	}
	allServer, err := cs.GetAllCenterServerList()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取元宝列表，获取全部服务器信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	centerplatform := 0
	skdType := 0
	if form.PlatformId > 0 {
		ps := platformservice.PlatformServiceInContext(req.Context())
		plInfo, plerr := ps.GetPlatformInfo(int64(form.PlatformId))
		if plerr != nil {
			log.WithFields(log.Fields{
				"error":    err,
				"serverid": form.ServerId,
			}).Error("获取元宝列表，获取平台信息异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		centerplatform = int(plInfo.CenterPlatformID)
		skdType = plInfo.SdkType
	}

	//获取权限sdktype
	sdklist, err := userservice.GetUserSdkList(req.Context())
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取用户sdk权限列表失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	//获取满足条件的玩家列表
	playerArray := make([]int64, 0)
	if form.StartMoney > 0 || form.EndMoney > 0 {
		os := orderservice.OrderServiceInContext(req.Context())
		playerArray, err = os.GetOrderAmountPlayerList(form.StartTime, form.EndTime, form.StartMoney, form.EndMoney, skdType, serverindex, sdklist)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("获取元宝列表异常，获取玩家列表异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(playerArray) == 0 {

			rr := gmhttp.NewSuccessResult(respon)
			httputils.WriteJSON(rw, http.StatusOK, rr)
		}
	}

	ss := staticservice.StaticReportServiceInContext(req.Context())
	if ss == nil {
		log.Error("staticreport服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := ss.GetGoldChangeStatic(form.StartTime, form.EndTime, centerplatform, serverindex, skdType, playerArray, sdklist)
	log.Debug("mongo查询结果记录：", len(rst))
	if err != nil {
		log.WithFields(log.Fields{
			"startTime":      form.StartTime,
			"endTime":        form.EndTime,
			"startMoney":     form.StartMoney,
			"endMoney":       form.EndMoney,
			"centerplatform": centerplatform,
			"serverindex":    serverindex,
			"sdktype":        skdType,
			"playerArray":    playerArray,
			"sdklist":        sdklist,
			"error":          err,
		}).Error("获取元宝列表异常，获取mongo统计数据异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	allServerMap := make(map[int]map[int]*centerServermodel.CenterServer)
	if len(allServer) > 0 {
		for _, value := range allServer {
			if value.ServerType != 0 {
				continue
			}
			if _, ok := allServerMap[int(value.Platform)]; !ok {
				allServerMap[int(value.Platform)] = make(map[int]*centerServermodel.CenterServer)
			}
			platMap := allServerMap[int(value.Platform)]
			platMap[value.ServerId] = value
		}
	}

	//组合转换结果
	// respon := &goldChangeRespon{}
	respon.ServerList = make([]*goldChangeResponServer, 0)
	respon.ItemArray = make([]*goldChangeResponItem, 0)
	totalGold := int64(0)
	reasonMap := make(map[int]map[int]*goldChangeResponServerItem) //key：原因id，value：{key:服务器id，value内容}
	serverMap := make(map[int]*goldChangeResponServer)             //key：服务器id
	totalServerName := &goldChangeResponServer{
		ServerId:    -1,
		ServerName:  "合计",
		ServerIndex: -1,
	}
	// respon.ServerList = append(respon.ServerList, totalServerName)
	serverMap[-1] = totalServerName

	goldTypeMap := getGoldChangeType(form.GoldType)
	for _, value := range rst {
		serverInfo := getPlatformInfo(allServerMap, value.Id.PlatformId, value.Id.ServerId)
		if serverInfo == nil {
			log.WithFields(log.Fields{
				"platformid": value.Id.PlatformId,
				"serverid":   value.Id.ServerId,
			}).Debug("检查服务器为空")
			continue
		}
		reason := int32(value.Id.Reason)
		if _, ok := goldTypeMap[reason]; !ok {
			continue
		}

		totalGold += value.ChangedNum
		if _, ok := serverMap[int(serverInfo.Id)]; !ok {
			serverItem := &goldChangeResponServer{
				ServerId:    int(serverInfo.Id),
				ServerName:  serverInfo.ServerName,
				ServerIndex: serverInfo.ServerId,
			}
			serverMap[int(serverInfo.Id)] = serverItem
		}
		serverMap[int(serverInfo.Id)].TotalChangeNum += value.ChangedNum
		serverMap[int(serverInfo.Id)].TotalPlayerCount += value.PlayerCount
		serverMap[-1].TotalChangeNum += value.ChangedNum
		serverMap[-1].TotalPlayerCount += value.PlayerCount

		if _, ok := reasonMap[value.Id.Reason]; !ok {
			reasonMap[value.Id.Reason] = make(map[int]*goldChangeResponServerItem)
		}
		reasonMapValue := reasonMap[value.Id.Reason]
		goldItem := &goldChangeResponServerItem{
			PlayerCount: value.PlayerCount,
			ChangeNum:   value.ChangedNum,
		}
		reasonMapValue[int(serverInfo.Id)] = goldItem
		if _, ok := reasonMapValue[-1]; !ok {
			reasonMapValue[-1] = &goldChangeResponServerItem{
				PlayerCount: 0,
				ChangeNum:   0,
			}
		}
		reasonMapValue[-1].ChangeNum += value.ChangedNum
		reasonMapValue[-1].PlayerCount += value.PlayerCount
	}
	for _, value := range serverMap {
		respon.ServerList = append(respon.ServerList, value)
	}

	for key, value := range reasonMap {
		myItem := &goldChangeResponItem{
			Reason:    key,
			ServerMap: value,
		}
		respon.ItemArray = append(respon.ItemArray, myItem)
	}

	sort.Sort(goldChangeResponItemSlice(respon.ItemArray))
	sort.Sort(goldChangeResponServerSlice(respon.ServerList))
	respon.TotalChangeNum = totalGold
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}

func getPlatformInfo(p_map map[int]map[int]*centerServermodel.CenterServer, p_platformid int, p_serverid int) *centerServermodel.CenterServer {
	if platformmap, ok := p_map[p_platformid]; ok {
		if value, vok := platformmap[p_serverid]; vok {
			return value
		}
	}
	return nil
}

type goldChangeResponItemSlice []*goldChangeResponItem //用于排序用的
func (a goldChangeResponItemSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a goldChangeResponItemSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}

func (a goldChangeResponItemSlice) Less(i, j int) bool { //
	return a[i].Reason < a[j].Reason
}

type goldChangeResponServerSlice []*goldChangeResponServer //用于排序用的
func (a goldChangeResponServerSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a goldChangeResponServerSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}

func (a goldChangeResponServerSlice) Less(i, j int) bool { //
	return a[i].ServerIndex < a[j].ServerIndex
}
