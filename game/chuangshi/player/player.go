package player

import (
	"fgame/fgame/game/chuangshi/dao"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
)

type PlayerChuangShiDataManager struct {
	pl                         player.Player
	playerChuangShiYuGaoObject *PlayerChuangShiYuGaoObject
	// //玩家神王报名
	// playerChuangShiSignObject *PlayerChuangShiSignObject
	// // 玩家神王投票
	// playerChuangShiVoteObject *PlayerChuangShiVoteObject
	// // 玩家创世信息
	// playerChuangShiObject *PlayerChuangShiObject
	// // 玩家官职信息
	// playerChuangShiGuanZhiObject *PlayerChuangShiGuanZhiObject
}

//加载
func (m *PlayerChuangShiDataManager) Load() (err error) {
	err = m.loadChuangShiYuGao()
	if err != nil {
		return
	}

	// err = m.loadChuangShiSign()
	// if err != nil {
	// 	return
	// }

	// err = m.loadChuangShiVote()
	// if err != nil {
	// 	return
	// }

	// err = m.loadChuangShi()
	// if err != nil {
	// 	return
	// }

	// err = m.loadChuangShiGuanZhi()
	// if err != nil {
	// 	return
	// }

	return nil
}

// func (m *PlayerChuangShiDataManager) loadChuangShiGuanZhi() (err error) {
// 	guanZhiEntity, err := dao.GetChuangShiDao().GetPlayerChuangShiGuanZhiEntity(m.pl.GetId())
// 	if err != nil {
// 		return
// 	}
// 	now := global.GetGame().GetTimeService().Now()
// 	if guanZhiEntity == nil {
// 		obj := NewPlayerChuangShiGuanZhiObject(m.pl)
// 		id, _ := idutil.GetId()
// 		obj.id = id
// 		obj.createTime = now
// 		obj.SetModified()
// 		m.playerChuangShiGuanZhiObject = obj
// 	} else {
// 		obj := NewPlayerChuangShiGuanZhiObject(m.pl)
// 		obj.FromEntity(guanZhiEntity)
// 		m.playerChuangShiGuanZhiObject = obj
// 	}
// 	return
// }

func (m *PlayerChuangShiDataManager) loadChuangShiYuGao() (err error) {
	//加载玩家创世之战预告信息
	chuangShiEntity, err := dao.GetChuangShiDao().GetPlayerChuangShiYuGaoEntity(m.pl.GetId())
	if err != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	obj := NewPlayerChuangShiYuGaoObject(m.pl)
	if chuangShiEntity == nil {
		id, _ := idutil.GetId()
		obj.id = id
		obj.createTime = now
		obj.SetModified()
	} else {
		obj.FromEntity(chuangShiEntity)
	}
	m.playerChuangShiYuGaoObject = obj
	return
}

// func (m *PlayerChuangShiDataManager) loadChuangShiSign() (err error) {
// 	entity, err := dao.GetChuangShiDao().GetPlayerChuangShiSignEntity(m.pl.GetId())
// 	if err != nil {
// 		return
// 	}
// 	now := global.GetGame().GetTimeService().Now()
// 	if entity == nil {
// 		obj := NewPlayerChuangShiSignObject(m.pl)
// 		id, _ := idutil.GetId()
// 		obj.id = id
// 		obj.status = chuangshitypes.ShenWangSignUpTypeNone
// 		obj.createTime = now
// 		obj.SetModified()
// 		m.playerChuangShiSignObject = obj
// 	} else {
// 		obj := NewPlayerChuangShiSignObject(m.pl)
// 		obj.FromEntity(entity)
// 		m.playerChuangShiSignObject = obj
// 	}
// 	return
// }

// func (m *PlayerChuangShiDataManager) loadChuangShiVote() (err error) {
// 	entity, err := dao.GetChuangShiDao().GetPlayerChuangShiVoetEntity(m.pl.GetId())
// 	if err != nil {
// 		return
// 	}
// 	now := global.GetGame().GetTimeService().Now()
// 	if entity == nil {
// 		obj := NewPlayerChuangShiVoteObject(m.pl)
// 		id, _ := idutil.GetId()
// 		obj.id = id
// 		obj.status = chuangshitypes.ShenWangVoteTypeNone
// 		obj.createTime = now
// 		obj.SetModified()
// 		m.playerChuangShiVoteObject = obj
// 	} else {
// 		obj := NewPlayerChuangShiVoteObject(m.pl)
// 		obj.FromEntity(entity)
// 		m.playerChuangShiVoteObject = obj
// 	}
// 	return
// }

// func (m *PlayerChuangShiDataManager) loadChuangShi() (err error) {
// 	entity, err := dao.GetChuangShiDao().GetPlayerChuangShiEntity(m.pl.GetId())
// 	if err != nil {
// 		return
// 	}
// 	now := global.GetGame().GetTimeService().Now()
// 	if entity == nil {
// 		obj := NewPlayerChuangShiObject(m.pl)
// 		id, _ := idutil.GetId()
// 		obj.id = id
// 		obj.pos = chuangshitypes.ChuangShiGuanZhiNone
// 		obj.createTime = now
// 		obj.SetModified()
// 		m.playerChuangShiObject = obj
// 	} else {
// 		obj := NewPlayerChuangShiObject(m.pl)
// 		obj.FromEntity(entity)
// 		m.playerChuangShiObject = obj
// 	}
// 	return
// }

func (m *PlayerChuangShiDataManager) Player() player.Player {
	return m.pl
}

func (m *PlayerChuangShiDataManager) IsJoin() bool {
	return m.playerChuangShiYuGaoObject.isJoin != 0
}

func (m *PlayerChuangShiDataManager) BaoMingSucess() bool {
	if m.playerChuangShiYuGaoObject.isJoin != 0 {
		return false
	}
	m.playerChuangShiYuGaoObject.isJoin = 1
	now := global.GetGame().GetTimeService().Now()
	m.playerChuangShiYuGaoObject.updateTime = now
	m.playerChuangShiYuGaoObject.SetModified()

	return true
}

//加载后
func (m *PlayerChuangShiDataManager) AfterLoad() (err error) {
	// m.refreshVote()
	return
}

//心跳
func (m *PlayerChuangShiDataManager) Heartbeat() {
}

// //投票
// func (m *PlayerChuangShiDataManager) AttendVote() (flag bool) {
// 	if m.playerChuangShiVoteObject.IfVote() {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerChuangShiVoteObject.status = chuangshitypes.ShenWangVoteTypeVoting
// 	m.playerChuangShiVoteObject.updateTime = now
// 	m.playerChuangShiVoteObject.SetModified()
// 	flag = true
// 	return
// }

// //神王投票结果
// func (m *PlayerChuangShiDataManager) VoteResult(status chuangshitypes.ShenWangVoteType) (flag bool) {
// 	if m.playerChuangShiVoteObject.status != chuangshitypes.ShenWangVoteTypeVoting {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerChuangShiVoteObject.status = status
// 	m.playerChuangShiVoteObject.updateTime = now
// 	m.playerChuangShiVoteObject.SetModified()
// 	flag = true
// 	return
// }

// func (m *PlayerChuangShiDataManager) refreshVote() {
// 	now := global.GetGame().GetTimeService().Now()
// 	isSame, _ := timeutils.IsSameDay(now, m.playerChuangShiVoteObject.lastVoteTime)
// 	if !isSame {
// 		m.playerChuangShiVoteObject.status = chuangshitypes.ShenWangVoteTypeNone
// 		m.playerChuangShiVoteObject.lastVoteTime = now
// 		m.playerChuangShiGuanZhiObject.updateTime = now
// 		m.playerChuangShiGuanZhiObject.SetModified()
// 	}
// }

// // 获取玩家官职等级
// func (m *PlayerChuangShiDataManager) GetPlayerChuangShiGuanZhiInfo() *PlayerChuangShiGuanZhiObject {
// 	return m.playerChuangShiGuanZhiObject
// }

// // 获取玩家创世信息
// func (m *PlayerChuangShiDataManager) GetPlayerChuangShiInfo() *PlayerChuangShiObject {
// 	return m.playerChuangShiObject
// }

// // 获取玩家神王投票信息
// func (m *PlayerChuangShiDataManager) GetPlayerChuangShiVoteInfo() *PlayerChuangShiVoteObject {
// 	m.refreshVote()
// 	return m.playerChuangShiVoteObject
// }

// // 获取玩家神王报名信息
// func (m *PlayerChuangShiDataManager) GetPlayerChuangShiSignInfo() *PlayerChuangShiSignObject {
// 	return m.playerChuangShiSignObject
// }

// // 使用积分
// func (m *PlayerChuangShiDataManager) UseJiFen(num int64) bool {
// 	if !m.playerChuangShiObject.IfEnoughJiFen(num) {
// 		return false
// 	}
// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerChuangShiObject.jifen -= num
// 	m.playerChuangShiObject.updateTime = now
// 	m.playerChuangShiObject.SetModified()
// 	return true
// }

// //神王选举报名
// func (m *PlayerChuangShiDataManager) ShenWangSignUp() (flag bool) {
// 	if m.playerChuangShiSignObject.IfShenWangSignUp() {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerChuangShiSignObject.status = chuangshitypes.ShenWangSignUpTypeSigning
// 	m.playerChuangShiSignObject.updateTime = now
// 	m.playerChuangShiSignObject.SetModified()
// 	flag = true
// 	return
// }

// //神王报名结果
// func (m *PlayerChuangShiDataManager) SignUpResult(status chuangshitypes.ShenWangSignUpType) (flag bool) {
// 	if m.playerChuangShiSignObject.status != chuangshitypes.ShenWangSignUpTypeSigning {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerChuangShiSignObject.status = status
// 	m.playerChuangShiSignObject.updateTime = now
// 	m.playerChuangShiSignObject.SetModified()
// 	flag = true
// 	return
// }

// // 官职数据更新
// func (m *PlayerChuangShiDataManager) UpdatePlayerGuanZhiData(success bool, useWeiWang int32) {
// 	now := global.GetGame().GetTimeService().Now()
// 	if success {
// 		m.playerChuangShiGuanZhiObject.level++
// 		m.playerChuangShiGuanZhiObject.times = 0
// 	} else {
// 		m.playerChuangShiGuanZhiObject.times++
// 	}
// 	m.playerChuangShiGuanZhiObject.weiWang -= useWeiWang
// 	m.playerChuangShiGuanZhiObject.updateTime = now
// 	m.playerChuangShiGuanZhiObject.SetModified()
// }

// //领取时装
// func (m *PlayerChuangShiDataManager) ReceiveGuanZhiRew(rewLevel int32) bool {
// 	if m.playerChuangShiGuanZhiObject.receiveRewLevel >= rewLevel {
// 		return false
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerChuangShiGuanZhiObject.receiveRewLevel = rewLevel
// 	m.playerChuangShiGuanZhiObject.updateTime = now
// 	m.playerChuangShiGuanZhiObject.SetModified()

// 	return true
// }

// //领取个人工资
// func (m *PlayerChuangShiDataManager) ReceiveMyPay(cityList []*chuangshidata.CityInfo) (flag bool) {
// 	if m.playerChuangShiObject.pos == chuangshitypes.ChuangShiGuanZhiNone {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	constantTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangshiConstantTemp()

// 	lastReceiveTime := m.playerChuangShiObject.lastMyPayTime
// 	if lastReceiveTime == 0 {
// 		lastReceiveTime = m.playerChuangShiObject.joinCampTime
// 	}
// 	rewCount, newLastRewTime := constantTemp.RewCout(lastReceiveTime, now)
// 	if rewCount <= 0 {
// 		return
// 	}

// 	addJifen := int32(0)
// 	addDiamonds := int32(0)
// 	for _, city := range cityList {
// 		cityTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiCityTemp(city.Camp, city.CityType, city.Index)
// 		addJifen += cityTemp.PlayerRewJifen * rewCount
// 		addDiamonds += cityTemp.PlayerRewZuanshi * rewCount
// 	}

// 	m.playerChuangShiObject.jifen += int64(addJifen)
// 	m.playerChuangShiObject.diamonds += int64(addDiamonds)
// 	m.playerChuangShiObject.lastMyPayTime = newLastRewTime
// 	m.playerChuangShiObject.updateTime = now
// 	m.playerChuangShiObject.SetModified()

// 	flag = true
// 	return
// }

// //城池建设
// func (m *PlayerChuangShiDataManager) AddJiFen(num int32) (flag bool) {
// 	if num <= 0 {
// 		panic(fmt.Errorf("积分不能小于1，%d", num))
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerChuangShiObject.jifen += int64(num)
// 	m.playerChuangShiObject.updateTime = now
// 	m.playerChuangShiObject.SetModified()

// 	flag = true
// 	return
// }

// //加入阵营
// func (m *PlayerChuangShiDataManager) JoinCamp(campType chuangshitypes.ChuangShiCampType) (flag bool) {
// 	if m.playerChuangShiObject.IfJoinCamp() {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerChuangShiObject.pos = chuangshitypes.ChuangShiGuanZhiPingMing
// 	m.playerChuangShiObject.campType = campType
// 	m.playerChuangShiObject.joinCampTime = now
// 	m.playerChuangShiObject.updateTime = now
// 	m.playerChuangShiObject.SetModified()

// 	flag = true
// 	return
// }

// 退出阵营

// 阵营改变

func CreatePlayerChuangShiDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerChuangShiDataManager{}
	m.pl = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerChuangShiDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerChuangShiDataManager))
}
