package chuangshi

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	chuangshientity "fgame/fgame/cross/chuangshi/entity"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

const (
	voteMaxLen = 10 //神王候选人数
)

// 成员战力排序
type signList []*ChuangShiSignInfo

func (a signList) Len() int           { return len(a) }
func (a signList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a signList) Less(i, j int) bool { return a[i].Member.force < a[j].Member.force }

// 神王票数排序
type voteList []*ChuangShiVoteInfo

func (a voteList) Len() int           { return len(a) }
func (a voteList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a voteList) Less(i, j int) bool { return a[i].Vote.ticketNum < a[j].Vote.ticketNum }

//神王报名
type ChuangShiSignInfo struct {
	Member *ChuangShiMemberObject
	Sign   *ChuangShiShenWangSignUpObject
}

//神王投票
type ChuangShiVoteInfo struct {
	Member *ChuangShiMemberObject
	Vote   *ChuangShiShenWangVoteObject
}

type Camp struct {
	campObj    *ChuangShiCampObject     //阵营信息
	memberList []*ChuangShiMemberObject //阵营成员
	signUpList []*ChuangShiSignInfo     //神王报名信息
	voteList   []*ChuangShiVoteInfo     //神王投票列表
	cityList   []*CityData              //城市列表
}

func createCamp(campObj *ChuangShiCampObject) *Camp {
	camp := &Camp{
		campObj: campObj,
	}
	return camp
}

func (c *Camp) ifCanTarget(campType chuangshitypes.ChuangShiCampType) bool {
	targetMap := c.campObj.targetCityMap
	if len(targetMap) >= 2 {
		return false
	}

	_, ok := targetMap[campType]
	if ok {
		return false
	}

	return true
}

func (c *Camp) getShenWangVote(supportId int64) *ChuangShiVoteInfo {
	for _, vote := range c.voteList {
		if vote.Member.playerId != supportId {
			continue
		}

		return vote
	}

	return nil
}

func (c *Camp) getShenWangSign(playerId int64) *ChuangShiSignInfo {
	for _, sign := range c.signUpList {
		if sign.Member.playerId != playerId {
			continue
		}

		return sign
	}

	return nil
}
func (c *Camp) getMember(memberId int64) *ChuangShiMemberObject {
	for _, mem := range c.memberList {
		if mem.playerId != memberId {
			continue
		}

		return mem
	}

	return nil
}

func (c *Camp) getCityByChengZhuId(chengzhuId int64) *CityData {
	for _, city := range c.cityList {
		if city.city.ownerId != chengzhuId {
			continue
		}

		return city
	}

	return nil
}

func (c *Camp) getCity(cityId int64) *CityData {
	for _, city := range c.cityList {
		if city.city.id != cityId {
			continue
		}

		return city
	}

	return nil
}

func (c *Camp) addShenWangSign(signObj *ChuangShiShenWangSignUpObject) {
	mem := c.getMember(signObj.playerId)

	sign := &ChuangShiSignInfo{}
	sign.Member = mem
	sign.Sign = signObj

	c.signUpList = append(c.signUpList, sign)
}

func (c *Camp) addShenWangVote(voteObj *ChuangShiShenWangVoteObject) {
	mem := c.getMember(voteObj.playerId)

	vote := &ChuangShiVoteInfo{}
	vote.Member = mem
	vote.Vote = voteObj

	c.voteList = append(c.voteList, vote)
}

func (c *Camp) addMember(mem *ChuangShiMemberObject) (flag bool) {
	tmemObj := c.getMember(mem.playerId)
	if tmemObj != nil {
		return
	}

	flag = true
	c.memberList = append(c.memberList, mem)
	return
}

func (c *Camp) addCity(city *CityData) {
	c.cityList = append(c.cityList, city)
	return
}

func (c *Camp) GetShenWangVoteList() []*ChuangShiVoteInfo {
	return c.voteList
}

func (c *Camp) GetCityList() []*CityData {
	return c.cityList
}

func (c *Camp) GetMember(memId int64) *ChuangShiMemberObject {
	return c.getMember(memId)
}

func (c *Camp) GetShenWangSignList() (signList []*ChuangShiSignInfo) {
	return c.signUpList
}

func (c *Camp) GetCampObj() *ChuangShiCampObject {
	return c.campObj
}

func (c *Camp) GetMemberList() (memList []*ChuangShiMemberObject) {
	return c.memberList
}

type ChuangShiCampObject struct {
	id             int64
	platform       int32
	serverId       int32
	campType       chuangshitypes.ChuangShiCampType           //
	kingId         int64                                      //神王id
	force          int64                                      //阵营战力
	shenWangStatus chuangshitypes.ShenWangStatusType          //神王选举阶段
	jifen          int64                                      //积分
	diamonds       int64                                      //钻石
	lastShouYiTime int64                                      //上次工资时间
	payJifen       int64                                      //积分工资（待领取）
	payDiamonds    int64                                      //钻石工资（待领取）
	targetCityMap  map[chuangshitypes.ChuangShiCampType]int64 //攻城目标
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func NewChuangShiCampObject() *ChuangShiCampObject {
	o := &ChuangShiCampObject{}
	return o
}

func convertChuangShiCampObjectToEntity(o *ChuangShiCampObject) (*chuangshientity.ChuangShiCampEntity, error) {
	data, err := json.Marshal(o.targetCityMap)
	if err != nil {
		return nil, err
	}

	e := &chuangshientity.ChuangShiCampEntity{
		Id:             o.id,
		Platform:       o.platform,
		ServerId:       o.serverId,
		CampType:       int32(o.campType),
		ShenWangStatus: int32(o.shenWangStatus),
		Jifen:          o.jifen,
		Diamonds:       o.diamonds,
		PayJifen:       o.payJifen,
		PayDiamonds:    o.payDiamonds,
		LastShouYiTime: o.lastShouYiTime,
		TargetMap:      string(data),
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return e, nil
}

func (o *ChuangShiCampObject) GetDeamonds() int64 {
	return o.diamonds
}

func (o *ChuangShiCampObject) GetJifen() int64 {
	return o.jifen
}

func (o *ChuangShiCampObject) GetId() int64 {
	return o.id
}

func (o *ChuangShiCampObject) GetCampType() chuangshitypes.ChuangShiCampType {
	return o.campType
}

func (o *ChuangShiCampObject) GetShenWangStatus() chuangshitypes.ShenWangStatusType {
	return o.shenWangStatus
}

func (o *ChuangShiCampObject) GetDBId() int64 {
	return o.id
}

func (o *ChuangShiCampObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertChuangShiCampObjectToEntity(o)
	return e, err
}

func (o *ChuangShiCampObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.ChuangShiCampEntity)

	dataMap := make(map[chuangshitypes.ChuangShiCampType]int64)
	err := json.Unmarshal([]byte(pse.TargetMap), &dataMap)
	if err != nil {
		return err
	}

	o.id = pse.Id
	o.platform = pse.Platform
	o.serverId = pse.ServerId
	o.campType = chuangshitypes.ChuangShiCampType(pse.CampType)
	o.shenWangStatus = chuangshitypes.ShenWangStatusType(pse.ShenWangStatus)
	o.jifen = pse.Jifen
	o.diamonds = pse.Diamonds
	o.payJifen = pse.PayJifen
	o.payDiamonds = pse.PayDiamonds
	o.lastShouYiTime = pse.LastShouYiTime
	o.targetCityMap = dataMap
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *ChuangShiCampObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ChuangShiCity"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}
