package data

import (
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	"fgame/fgame/game/scene/scene"
)

type BattlePlayerStatus int32

const (
	BattlePlayerStatusOffline BattlePlayerStatus = iota //离线
	BattlePlayerStatusOnline                            //在线
)

//对战信息
type PvpPlayerInfo struct {
	*BattlePlayerBasicInfo
	state          arenapvptypes.ArenapvpState
	BattleDataList []*BattleResultData //对战记录
}

func (info *PvpPlayerInfo) Failed() bool {
	if info.state != arenapvptypes.ArenapvpStateInit {
		return false
	}
	info.state = arenapvptypes.ArenapvpStateFailed
	return true
}

func (info *PvpPlayerInfo) Exit() bool {
	if info.state != arenapvptypes.ArenapvpStateInit {
		return false
	}
	info.state = arenapvptypes.ArenapvpStateExit
	return true
}

func (info *PvpPlayerInfo) GetState() arenapvptypes.ArenapvpState {
	return info.state
}

func (info *PvpPlayerInfo) GetBattleData(pvpType arenapvptypes.ArenapvpType) *BattleResultData {
	for _, result := range info.BattleDataList {
		if result.PvpType != pvpType {
			continue
		}

		return result
	}

	return nil
}

func (info *PvpPlayerInfo) GetCurBatlleData() *BattleResultData {
	battleLen := len(info.BattleDataList)
	if battleLen == 0 {
		return nil
	}

	return info.BattleDataList[battleLen-1]
}

func NewPvpPlayerInfo() *PvpPlayerInfo {
	d := &PvpPlayerInfo{}
	d.BattlePlayerBasicInfo = &BattlePlayerBasicInfo{}
	return d
}

func ConvertToPvpPlayerInfo(pl scene.Player) *PvpPlayerInfo {
	info := &PvpPlayerInfo{}
	info.BattlePlayerBasicInfo = CreateBattlePlayerBasicInfo(pl)
	return info
}

// 玩家基础信息
type BattlePlayerBasicInfo struct {
	Platform    int32  //平台
	ServerId    int32  //服id
	PlayerId    int64  //玩家id
	PlayerName  string //姓名
	Sex         int32  //
	Role        int32  //
	IsRobot     bool   //是否机器人
	Force       int64  //
	WingId      int32
	WeaponId    int32
	FashionId   int32
	WeaponState int32
	XianTiId    int32
	LingYuId    int32
	FaBaoId     int32
}

func CreateBattlePlayerBasicInfo(pl scene.Player) *BattlePlayerBasicInfo {
	info := &BattlePlayerBasicInfo{}
	info.Platform = pl.GetPlatform()
	info.ServerId = pl.GetServerId()
	info.PlayerId = pl.GetId()
	info.PlayerName = pl.GetName()
	info.Sex = int32(pl.GetSex())
	info.Role = int32(pl.GetRole())
	info.IsRobot = pl.IsRobot()
	info.Force = pl.GetForce()
	info.WingId = pl.GetWingId()
	info.WeaponId = pl.GetWeaponId()
	info.WeaponState = pl.GetWeaponState()
	info.FashionId = pl.GetFashionId()
	info.XianTiId = pl.GetXianTiId()
	info.LingYuId = pl.GetLingYuId()
	info.FaBaoId = pl.GetFaBaoId()
	return info
}

//对战记录
type BattleResultData struct {
	PvpType   arenapvptypes.ArenapvpType
	WinnerId  int64 //获胜id
	BattleId1 int64 //对战id1
	BattleId2 int64 //对战id2
	Index     int32 //场次
}

func (d *BattleResultData) GetBattleId(playerId int64) int64 {
	if d.BattleId1 == playerId {
		return d.BattleId2
	}

	return d.BattleId1
}

func CreateBattleResultData(pvpType arenapvptypes.ArenapvpType, battleId1, battleId2 int64, index int32) *BattleResultData {
	info := &BattleResultData{}
	info.PvpType = pvpType
	info.BattleId1 = battleId1
	info.BattleId2 = battleId2
	info.Index = index
	return info
}

//竞猜对象
type GuessData struct {
	PvpType    arenapvptypes.ArenapvpType //竞猜类型
	RaceNumber int32                      //届数
	PlayerList []*PvpPlayerInfo
}

func CreateGuessData(pvpType arenapvptypes.ArenapvpType, raceNum int32, playerList []*PvpPlayerInfo) *GuessData {
	info := &GuessData{}
	info.PvpType = pvpType
	info.RaceNumber = raceNum
	info.PlayerList = playerList
	return info
}

func (d *GuessData) GetWinnerId() int64 {
	if len(d.PlayerList) == 0 {
		return 0
	}

	pvpType := d.PvpType
	battlePlayer := d.PlayerList[0]
	battleResult := battlePlayer.GetBattleData(pvpType)
	return battleResult.WinnerId
}

//海选会场信息
type ElectionData struct {
	ElectionIndex int32  //会场编号哦
	PlNumber      int32  //会场人数
	LastLuckyTime int64  //上次幸运奖时间
	LuckyNameText string //中奖名单
}

//霸主信息
type BaZhuData struct {
	RaceNumber  int32
	Platform    int32  //平台
	ServerId    int32  //服id
	PlayerId    int64  //玩家id
	PlayerName  string //姓名
	Sex         int32  //
	Role        int32  //
	WingId      int32
	WeaponId    int32
	FashionId   int32
	WeaponState int32
	XianTiId    int32
	LingYuId    int32
	FaBaoId     int32
}
