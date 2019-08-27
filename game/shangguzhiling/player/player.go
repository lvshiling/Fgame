package player

import (
	commonlogic "fgame/fgame/game/common/logic"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/shangguzhiling/dao"
	shangguzhilingtemplate "fgame/fgame/game/shangguzhiling/template"
	shangguzhilingtypes "fgame/fgame/game/shangguzhiling/types"
	"fgame/fgame/pkg/mathutils"
)

type PlayerShangguzhilingDataManager struct {
	p           player.Player
	lingShouMap map[shangguzhilingtypes.LingshouType]*PlayerShangguzhilingObject
}

func (m *PlayerShangguzhilingDataManager) Player() player.Player {
	return m.p
}

// 领取奖励
func (m *PlayerShangguzhilingDataManager) UpdateLastReceiveTime(typ shangguzhilingtypes.LingshouType) {
	obj, ok := m.lingShouMap[typ]
	if !ok {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	obj.receiveTime = now
	obj.updateTime = now
	obj.SetModified()
}

// 获取当前存于内存的灵兽对象
func (m *PlayerShangguzhilingDataManager) GetCurrentLingShouObjectList() []*PlayerShangguzhilingObject {
	lingshouList := []*PlayerShangguzhilingObject{}
	for _, obj := range m.lingShouMap {
		lingshouList = append(lingshouList, obj)
	}
	return lingshouList
}

// 需要刷新的情况（灵兽升级的时候，初始化的时候）
func (m *PlayerShangguzhilingDataManager) refreshLingShouInitData() {
	for typ := shangguzhilingtypes.MinLingshouType; typ <= shangguzhilingtypes.MaxLingshouType; typ++ {
		obj, ok := m.lingShouMap[typ]
		if ok {
			//以防策划数据修改
			obj.refreshLingLianStatus()
			obj.SetModified()
			continue
		}
		if m.IsLingShouUnlock(typ) {
			obj = createNewPlayerShangguzhilingObject(m.p, typ)
			obj.refreshLingLianStatus()
			obj.SetModified()
			// 保存到内存
			m.lingShouMap[typ] = obj
		}
	}
}

// 获取灵兽信息对象
func (m *PlayerShangguzhilingDataManager) GetLingShouObj(typ shangguzhilingtypes.LingshouType) *PlayerShangguzhilingObject {
	obj, ok := m.lingShouMap[typ]
	if !ok {
		return nil
	}
	return obj
}

// 灵兽是否已经解锁
func (m *PlayerShangguzhilingDataManager) IsLingShouUnlock(typ shangguzhilingtypes.LingshouType) bool {
	baseTemp := shangguzhilingtemplate.GetShangguzhilingTemplateService().GetLingShouTemplate(typ)
	if baseTemp == nil {
		return false
	}

	needTpy := baseTemp.GetNeedLingShouType()
	// 需要灵兽为自己即默认开启
	if needTpy == typ {
		return true
	}
	needObj := m.GetLingShouObj(needTpy)
	if needObj == nil {
		return false
	}
	return needObj.isUnlock(baseTemp.NeedSgzlLevel)
}

// 灵兽增加经验
func (m *PlayerShangguzhilingDataManager) AddExp(typ shangguzhilingtypes.LingshouType, exp int64) bool {
	obj, ok := m.lingShouMap[typ]
	if !ok {
		return false
	}
	baseTemp := shangguzhilingtemplate.GetShangguzhilingTemplateService().GetLingShouTemplate(typ)
	if baseTemp == nil {
		return false
	}
	temp := baseTemp.GetLevelTemp(obj.GetLevel() + 1) //取下一级别模板
	if temp == nil {
		return false
	}
	maxExp := temp.GeExperience()
	curExp := obj.experience + exp
	if curExp >= maxExp {
		obj.experience = curExp - maxExp
		obj.level++
		obj.refreshLingLianStatus()
		m.refreshLingShouInitData()
	} else {
		obj.experience = curExp
	}
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	return true
}

// 灵纹是否已经解锁
func (m *PlayerShangguzhilingDataManager) IsLingWenUnlock(typ shangguzhilingtypes.LingshouType, lingwenTyp shangguzhilingtypes.LingwenType) bool {
	baseTemp := shangguzhilingtemplate.GetShangguzhilingTemplateService().GetLingWenTemplate(typ, lingwenTyp)
	if baseTemp == nil {
		return false
	}

	needObj := m.GetLingShouObj(typ)
	if needObj == nil {
		return false
	}
	return needObj.isUnlock(baseTemp.NeedSgzlLevel)
}

//已解锁的灵纹列表
func (m *PlayerShangguzhilingDataManager) GetLingWenUnlockList(typ shangguzhilingtypes.LingshouType) []shangguzhilingtypes.LingwenType {
	subTypeList := []shangguzhilingtypes.LingwenType{}
	for subType := shangguzhilingtypes.MinLingwenType; subType <= shangguzhilingtypes.MaxLingwenType; subType++ {
		if m.IsLingWenUnlock(typ, subType) {
			subTypeList = append(subTypeList, subType)
		}
	}
	return subTypeList
}

// 灵纹增加经验
func (m *PlayerShangguzhilingDataManager) AddLingWenExp(typ shangguzhilingtypes.LingshouType, lingwenTyp shangguzhilingtypes.LingwenType, exp int64) bool {
	obj, ok := m.lingShouMap[typ]
	if !ok {
		return false
	}
	lingwenInfo, ok := obj.lingwen[lingwenTyp]
	if !ok {
		return false
	}

	lingwenTemp := shangguzhilingtemplate.GetShangguzhilingTemplateService().GetLingWenTemplate(typ, lingwenTyp)
	if lingwenTemp == nil {
		return false
	}
	temp := lingwenTemp.GetLevelTemp(lingwenInfo.Level + 1) //取下一级模板
	if temp == nil {
		return false
	}
	maxExp := temp.GetExperience()
	curExp := lingwenInfo.Experience + exp
	if curExp >= maxExp {
		lingwenInfo.Experience = curExp - maxExp
		lingwenInfo.Level++
	} else {
		lingwenInfo.Experience = curExp
	}
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	return true
}

// 灵兽进阶
func (m *PlayerShangguzhilingDataManager) UpRank(typ shangguzhilingtypes.LingshouType) bool {
	obj, ok := m.lingShouMap[typ]
	if !ok {
		return false
	}

	// ---
	baseTemp := shangguzhilingtemplate.GetShangguzhilingTemplateService().GetLingShouTemplate(typ)
	if baseTemp == nil {
		return false
	}
	upRankTemp := baseTemp.GetRankTemp(obj.uprankLevel + 1) //取下一级别模板
	if upRankTemp == nil {
		return false
	}
	updateRate := upRankTemp.UpdateWfb
	blessMax := upRankTemp.ZhufuMax
	addMin := upRankTemp.AddMin
	addMax := upRankTemp.AddMax + 1

	randBless := int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes := int32(1)
	curTimesNum := obj.uprankTimes
	curTimesNum += addTimes
	curBless := int32(obj.uprankBless)

	bless, sucess := commonlogic.AdvancedStatusAndProgress(curTimesNum, curBless, upRankTemp.TimesMin, upRankTemp.TimesMax, randBless, updateRate, blessMax)
	// ---

	if sucess {
		obj.uprankLevel++
		obj.uprankTimes = int32(0)
		obj.uprankBless = int64(0)
	} else {
		obj.uprankTimes++
		obj.uprankBless += int64(bless)
	}
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	return true
}

// 灵炼部位是否解除限制
func (m *PlayerShangguzhilingDataManager) IsLingLianPosJiesuo(typ shangguzhilingtypes.LingshouType, pos shangguzhilingtypes.LinglianPosType) bool {
	obj := m.GetLingShouObj(typ)
	if obj == nil {
		return false
	}
	linglianTemp := shangguzhilingtemplate.GetShangguzhilingTemplateService().GetLingLianTemplate(typ, pos)
	if linglianTemp == nil {
		return false
	}
	return obj.isUnlock(linglianTemp.NeedSgzlLevel)
}

//已解除限制的灵炼部位列表
func (m *PlayerShangguzhilingDataManager) GetLingLianPosJiesuoList(typ shangguzhilingtypes.LingshouType) []shangguzhilingtypes.LinglianPosType {
	subTypeList := []shangguzhilingtypes.LinglianPosType{}
	for subType := shangguzhilingtypes.MinLinglianPosType; subType <= shangguzhilingtypes.MaxLinglianPosType; subType++ {
		if m.IsLingLianPosJiesuo(typ, subType) {
			subTypeList = append(subTypeList, subType)
		}
	}
	return subTypeList
}

// 灵兽灵炼
func (m *PlayerShangguzhilingDataManager) LingLian(typ shangguzhilingtypes.LingshouType, changeLockStatusList []shangguzhilingtypes.LinglianPosType) {
	obj, ok := m.lingShouMap[typ]
	if !ok {
		return
	}

	// for _, pos := range changeLockStatusList {
	// 	info := obj.linglian[pos]
	// 	//前置状态为非锁定
	// 	if !info.IsLock {
	// 		info.LockTimes++
	// 	}
	// 	info.IsLock = !info.IsLock
	// }

	//策划修改简化
	lockMap := make(map[shangguzhilingtypes.LinglianPosType]bool)
	for _, pos := range changeLockStatusList {
		info := obj.linglian[pos]
		if info == nil {
			continue
		}
		lockMap[pos] = true
	}

	lingshouTempService := shangguzhilingtemplate.GetShangguzhilingTemplateService()
	for pos, info := range obj.linglian {
		// // 锁定
		// if info.IsLock {
		// 	continue
		// }
		//策划修改简化
		if lockMap[pos] {
			continue
		}

		linglianTemp := lingshouTempService.GetLingLianTemplate(typ, pos)
		if linglianTemp == nil {
			continue
		}
		// 未解除限制（灵兽等级不够）
		if !obj.isUnlock(linglianTemp.NeedSgzlLevel) {
			continue
		}

		info.PoolMark = linglianTemp.GetRandomPoolTempMark()
	}

	obj.linglianTimes++
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
}

//获取列表里面状态变化【非锁定】----->【锁定】所需要消耗的锁定物品数量
func (m *PlayerShangguzhilingDataManager) GetLingLianLockNeedNum(typ shangguzhilingtypes.LingshouType, changeLockStatusList []shangguzhilingtypes.LinglianPosType) int32 {
	obj := m.GetLingShouObj(typ)
	if obj == nil {
		return 0
	}
	num := int32(0)
	for _, pos := range changeLockStatusList {
		info := obj.linglian[pos]
		// //当前非锁定
		// if !info.IsLock {
		// 	cnt := shangguzhilingtemplate.GetShangguzhilingTemplateService().GetLingLianLockUseItemCount(info.LockTimes)
		// 	num += cnt
		// }
		// 策划修改简化
		if info == nil {
			continue
		}
		num++
	}
	// 策划修改简化
	num = shangguzhilingtemplate.GetShangguzhilingTemplateService().GetLingLianLockUseItemCount(num)
	return num
}

//加载
func (m *PlayerShangguzhilingDataManager) Load() (err error) {
	//加载灵兽数据
	lingshouEntityList, err := dao.GetShangguzhilingDao().GetShangguzhilingEntity(m.p.GetId())
	if err != nil {
		return
	}

	m.lingShouMap = make(map[shangguzhilingtypes.LingshouType]*PlayerShangguzhilingObject)
	for _, e := range lingshouEntityList {
		obj := NewPlayerShangguzhilingObject(m.p)
		obj.FromEntity(e)
		m.lingShouMap[obj.lingShouType] = obj
	}
	// 以防万一
	m.refreshLingShouInitData()

	return nil
}

//加载后
func (m *PlayerShangguzhilingDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerShangguzhilingDataManager) Heartbeat() {
}

func createPlayerShangguzhilingDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerShangguzhilingDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerShangguzhilingDataManagerType, player.PlayerDataManagerFactoryFunc(createPlayerShangguzhilingDataManager))
}
