package data

import (
	chuangshipb "fgame/fgame/cross/chuangshi/pb"
	alliancetypes "fgame/fgame/game/alliance/types"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
)

// 阵营
type CampData struct {
	Jifen          int64                             // 积分
	Diamonds       int64                             // 钻石
	CampType       chuangshitypes.ChuangShiCampType  // 阵营
	Force          int64                             // 总战力
	KingMem        *MemberInfo                       // 神王
	ShenWangStatus chuangshitypes.ShenWangStatusType // 神王选举阶段
	SignUpList     []*MemberInfo                     // 报名列表
	MemberList     []*MemberInfo                     // 成员列表
	VoteList       []*VoteInfo                       // 投票列表
	CityList       []*CityInfo                       // 城池列表
}

func (c *CampData) GetCityByChengZhuId(chengzhuId int64) *CityInfo {
	for _, city := range c.CityList {
		if city.Member.PlayerId != chengzhuId {
			continue
		}

		return city
	}

	return nil
}

func (c *CampData) GetCityById(cityId int64) *CityInfo {
	for _, city := range c.CityList {
		if city.CityId != cityId {
			continue
		}

		return city
	}

	return nil
}

//跨服数据转换
func ConvertToCampList(campList []*chuangshipb.ChuangShiCamp) (infoList []*CampData, memberMap map[int64]*MemberInfo) {
	memberMap = make(map[int64]*MemberInfo)
	for _, camp := range campList {
		infoList = append(infoList, ConvertToCamp(camp))

		for _, mem := range camp.MemberList {
			memberMap[mem.PlayerId] = convertToMember(mem)
		}
	}
	return infoList, memberMap
}

func ConvertToCamp(camp *chuangshipb.ChuangShiCamp) (info *CampData) {
	data := &CampData{}
	data.Jifen = camp.Jifen
	data.Diamonds = camp.Diamonds
	data.CampType = chuangshitypes.ChuangShiCampType(camp.CampType)
	data.ShenWangStatus = chuangshitypes.ShenWangStatusType(camp.ShenWangStatus)
	data.Force = camp.Power
	data.KingMem = convertToMember(camp.KingMember)
	data.SignUpList = ConvertToMemberList(camp.SignList)
	data.MemberList = ConvertToMemberList(camp.MemberList)
	data.VoteList = ConvertToVoteList(camp.VoteList)
	data.CityList = ConvertToCityList(camp.CityList)
	return data
}

// 成员
type MemberInfo struct {
	Platform     int32                            //平台
	ServerId     int32                            //服务器
	PlayerId     int64                            //玩家id
	PlayerName   string                           //玩家名字
	AllianceId   int64                            //仙盟id
	AllianceName string                           //仙盟名字
	Pos          chuangshitypes.ChuangShiGuanZhi  //官职
	AlPos        alliancetypes.AlliancePosition   //仙盟职位
	Force        int64                            //战力
	CampType     chuangshitypes.ChuangShiCampType //阵营
}

func (m *MemberInfo) IfMengZhu() bool {
	return m.AlPos == alliancetypes.AlliancePositionMengZhu
}

func (mem *MemberInfo) IfShenWang() bool {
	return mem.Pos == chuangshitypes.ChuangShiGuanZhiShenWang
}

func (mem *MemberInfo) IfChengZhu() bool {
	return mem.Pos == chuangshitypes.ChuangShiGuanZhiChengZhu
}

func CareteMemberInfo(pl player.Player) *MemberInfo {
	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerIndex()
	playerId := pl.GetId()
	playerName := pl.GetName()
	force := pl.GetForce()
	allianceId := pl.GetAllianceId()
	allianceName := pl.GetAllianceName()
	alPos := pl.GetMemPos()

	data := &MemberInfo{}
	data.Platform = platform
	data.ServerId = serverId
	data.PlayerId = playerId
	data.PlayerName = playerName
	data.AllianceId = allianceId
	data.AllianceName = allianceName
	data.AlPos = alliancetypes.AlliancePosition(alPos)
	data.Force = force
	return data
}

//跨服数据转换
func ConvertToMemberList(memList []*chuangshipb.ChuangShiMemberInfo) (infoList []*MemberInfo) {
	for _, mem := range memList {
		infoList = append(infoList, convertToMember(mem))
	}
	return infoList
}

func convertToMember(mem *chuangshipb.ChuangShiMemberInfo) (info *MemberInfo) {
	if mem == nil {
		return
	}

	data := &MemberInfo{}
	data.Platform = mem.Platform
	data.ServerId = mem.ServerId
	data.PlayerId = mem.PlayerId
	data.PlayerName = mem.Name
	data.AllianceName = mem.AllianceName
	data.AllianceId = mem.AllianceId
	data.Pos = chuangshitypes.ChuangShiGuanZhi(mem.Pos)
	data.AlPos = alliancetypes.AlliancePosition(mem.AlPos)
	data.Force = mem.Force

	return data
}

// 投票
type VoteInfo struct {
	Member    *MemberInfo //玩家信息
	TicketNum int32       //票数
}

//跨服数据转换
func ConvertToVoteList(voteList []*chuangshipb.ChuangShiVote) (infoList []*VoteInfo) {
	for _, vote := range voteList {
		data := &VoteInfo{}
		data.TicketNum = vote.TicketNum
		data.Member = convertToMember(vote.Member)

		infoList = append(infoList, data)
	}
	return infoList
}

// 城池
type CityInfo struct {
	CityId      int64                            //城池id
	Camp        chuangshitypes.ChuangShiCampType //阵营
	OrignalCamp chuangshitypes.ChuangShiCampType //初始阵营
	CityType    chuangshitypes.ChuangShiCityType //城市类型
	Index       int32                            //城市索引
	Member      *MemberInfo                      //城主信息
	Jifen       int64                            //积分
	Diamonds    int64                            //钻石
	JianSheList []*JianSheData                   //城池建设
}

func (c *CityInfo) IfFuShu() bool {
	return c.CityType == chuangshitypes.ChuangShiCityTypeFushu
}

func (c *CityInfo) IfChengZhu(playerId int64) bool {
	return c.Member.PlayerId == playerId
}

func (c *CityInfo) IfActivateJianSheSkill(level int32) bool {
	jianShe := c.GetJianShe(chuangshitypes.ChuangShiCityJianSheTypeTianQi)
	_, ok := jianShe.SkillMap[level]
	return ok
}

func (c *CityInfo) GetJianShe(jianSheType chuangshitypes.ChuangShiCityJianSheType) *JianSheData {
	for _, jianShe := range c.JianSheList {
		if jianShe.JianSheType != jianSheType {
			continue
		}

		return jianShe
	}

	return nil
}

//跨服数据转换
func ConvertToCityList(cityList []*chuangshipb.ChuangShiCity) (infoList []*CityInfo) {
	for _, city := range cityList {
		data := &CityInfo{}
		data.Camp = chuangshitypes.ChuangShiCampType(city.Camp)
		data.OrignalCamp = chuangshitypes.ChuangShiCampType(city.OrignalCamp)
		data.CityType = chuangshitypes.ChuangShiCityType(city.CityType)
		data.Index = city.Index
		data.CityId = city.CityId
		data.Member = convertToMember(city.KingMember)
		data.Jifen = city.Jifen
		data.Diamonds = city.Diamonds
		data.JianSheList = ConvertToJianSheList(city.JianSheList)

		infoList = append(infoList, data)
	}
	return infoList
}

// 城池建设
type JianSheData struct {
	JianSheType   chuangshitypes.ChuangShiCityJianSheType //建筑类型
	Level         int32                                   //等级
	Exp           int32                                   //经验
	SkillMap      map[int32]int64                         //技能激活记录（天气台专用）
	SkillLevelSet int32                                   //当前使用技能（天气台专用）
}

//跨服数据转换
func ConvertToJianSheList(jianSheList []*chuangshipb.CityJianShe) (infoList []*JianSheData) {
	for _, jianShe := range jianSheList {
		data := &JianSheData{}
		data.JianSheType = chuangshitypes.ChuangShiCityJianSheType(jianShe.JianSheType)
		data.Level = jianShe.Level
		data.Exp = jianShe.Exp
		data.SkillLevelSet = jianShe.SkillLevelSet
		for _, skill := range jianShe.SkillMap {
			data.SkillMap = make(map[int32]int64)
			data.SkillMap[skill.Level] = skill.ActivateTime
		}

		infoList = append(infoList, data)
	}
	return infoList
}
