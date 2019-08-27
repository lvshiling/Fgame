package player

import (
	"fgame/fgame/core/utils"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	ylpdao "fgame/fgame/game/yinglingpu/dao"
	yinglingputemplate "fgame/fgame/game/yinglingpu/template"
	ylptypes "fgame/fgame/game/yinglingpu/types"
	"fgame/fgame/pkg/idutil"
)

type PlayerYingLingPuManager struct {
	player     player.Player
	ylpMap     map[ylptypes.YingLingPuTuJianType]map[int32]*YingLingPuObject          //英灵普图鉴
	ylpSpMap   map[ylptypes.YingLingPuTuJianType]map[int32][]*YingLingPuSuiPianObject //英灵普图鉴碎片
	yipSuitMap map[int32][]int32                                                      //英灵普姻缘
}

//***********接口开始****************

//玩家
func (m *PlayerYingLingPuManager) Player() player.Player {
	return m.player
}

//加载
func (m *PlayerYingLingPuManager) Load() error {
	playerId := m.player.GetId()
	ylpList, err := ylpdao.GetYingLingPuDao().GetYingLingPu(playerId)
	if err != nil {
		return err
	}
	ylpSpList, err := ylpdao.GetYingLingPuDao().GetYingLingPuSuiPian(playerId)
	if err != nil {
		return err
	}
	m.ylpMap = make(map[ylptypes.YingLingPuTuJianType]map[int32]*YingLingPuObject)
	m.ylpSpMap = make(map[ylptypes.YingLingPuTuJianType]map[int32][]*YingLingPuSuiPianObject)

	for i := ylptypes.YingLingPuTuJianTypeCommon; i <= ylptypes.GetMaxYingLingPuType(); i++ {
		m.ylpMap[i] = make(map[int32]*YingLingPuObject)
		m.ylpSpMap[i] = make(map[int32][]*YingLingPuSuiPianObject)
	}

	for _, value := range ylpList {
		ylpObject := NewYingLingPuObject(m.player)
		err = ylpObject.FromEntity(value)
		if err != nil {
			return err
		}
		m.ylpMap[ylpObject.TuJianType][ylpObject.TuJianId] = ylpObject
	}
	for _, value := range ylpSpList {
		ylpSpObject := NewYingLingPuSuiPianObject(m.player)
		err = ylpSpObject.FromEntity(value)
		if err != nil {
			return err
		}
		m.ylpSpMap[ylpSpObject.TuJianType][ylpSpObject.TuJianId] = append(m.ylpSpMap[ylpSpObject.TuJianType][ylpSpObject.TuJianId], ylpSpObject)
	}
	return nil
}

//加载后
func (m *PlayerYingLingPuManager) AfterLoad() error {
	return nil
}

//心跳
func (m *PlayerYingLingPuManager) Heartbeat() {

}

func NewPlaerYingLingPuManager(pl player.Player) player.PlayerDataManager {
	rst := &PlayerYingLingPuManager{
		player: pl,
	}
	return rst
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerYingLingPuManagerType, player.PlayerDataManagerFactoryFunc(NewPlaerYingLingPuManager))
}

//***********自定义开始****************/

func (m *PlayerYingLingPuManager) GetAllYingLingPu() []*YingLingPuObject {
	rst := make([]*YingLingPuObject, 0)
	for _, value := range m.ylpMap {
		for _, valueObj := range value {
			rst = append(rst, valueObj)
		}
	}
	return rst
}

// 获取英灵普已收集套装
func (m *PlayerYingLingPuManager) GetYingLingPuSuitMap() map[int32]*gametemplate.YingLingPuSuitTemplate {
	var hadCollectIdList []int32
	for _, value := range m.ylpMap {
		for _, valueObj := range value {
			ylpTemplate := yinglingputemplate.GetYingLingPuTemplateService().GetYingLingPuById(valueObj.TuJianId, valueObj.TuJianType)
			hadCollectIdList = append(hadCollectIdList, int32(ylpTemplate.Id))
		}
	}

	percentSuitMap := make(map[int32]*gametemplate.YingLingPuSuitTemplate)
	suitTempMap := yinglingputemplate.GetYingLingPuTemplateService().GetYingLingPuSuitMap()
	for _, suitTemp := range suitTempMap {
		isGroup := true
		for _, needYlpId := range suitTemp.GetSuitCondition() {
			if !utils.ContainInt32(hadCollectIdList, needYlpId) {
				isGroup = false
				break
			}
		}

		if !isGroup {
			continue
		}

		for _, needYlpId := range suitTemp.GetSuitCondition() {
			percentSuitMap[needYlpId] = suitTemp
		}
	}

	return percentSuitMap
}

func (m *PlayerYingLingPuManager) GetAllYingLingPuSuiPian() []*YingLingPuSuiPianObject {
	rst := make([]*YingLingPuSuiPianObject, 0)
	for _, value := range m.ylpSpMap {
		for _, valueObj := range value {
			rst = append(rst, valueObj...)
		}
	}
	return rst
}

func (m *PlayerYingLingPuManager) GetAllYingLingPuSuiPianByTuJian(ylpId int32, ylpType ylptypes.YingLingPuTuJianType) []*YingLingPuSuiPianObject {
	rst := make([]*YingLingPuSuiPianObject, 0)
	for _, valueObj := range m.ylpSpMap[ylpType][ylpId] {
		rst = append(rst, valueObj)
	}
	return rst
}

func (m *PlayerYingLingPuManager) GetYlpInfo(ylpId int32, ylpType ylptypes.YingLingPuTuJianType) *YingLingPuObject {
	ylpMap := m.ylpMap[ylpType]
	if ylpMap != nil {
		_, ok := ylpMap[ylpId]
		if ok {
			return ylpMap[ylpId]
		}
	}
	return nil
}

//获取英灵普碎片集合
func (m *PlayerYingLingPuManager) GetYlpSpList(ylpId int32, ylpType ylptypes.YingLingPuTuJianType) []*YingLingPuSuiPianObject {
	ylpSpMap := m.ylpSpMap[ylpType]
	if ylpSpMap == nil {
		return nil
	}
	return ylpSpMap[ylpId]
}

func (m *PlayerYingLingPuManager) ExistsYlp(p_yluId int32, ylpType ylptypes.YingLingPuTuJianType) bool {
	rst := false
	ylpMap := m.ylpMap[ylpType]
	if ylpMap != nil {
		_, ok := ylpMap[p_yluId]
		if ok {
			rst = true
		}
	}
	return rst
}

func (m *PlayerYingLingPuManager) ExistsYlpSp(ylpId int32, ylpType ylptypes.YingLingPuTuJianType, suiPianId int32) bool {
	rst := false
	ylpSpMap := m.ylpSpMap[ylpType]
	if ylpSpMap != nil {
		_, ok := ylpSpMap[ylpId]
		if !ok {
			return false
		}
		spList := ylpSpMap[ylpId]
		for _, value := range spList {
			if value.SuiPianId == suiPianId {
				return true
			}
		}
	}
	return rst
}

//添加英灵普碎片
func (m *PlayerYingLingPuManager) AddYingLingPuSuiPian(ylpId int32, ylpType ylptypes.YingLingPuTuJianType, suiPianId int32) bool {
	if m.ExistsYlpSp(ylpId, ylpType, suiPianId) {
		return false
	}
	yingLingPuTemplate := yinglingputemplate.GetYingLingPuTemplateService().GetYingLingPuById(ylpId, ylpType)
	if yingLingPuTemplate == nil {
		return false
	}

	spObj := NewYingLingPuSuiPianObject(m.player)
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	spObj.Id = id
	spObj.SuiPianId = suiPianId
	spObj.CreateTime = now
	spObj.TuJianId = ylpId
	spObj.TuJianType = ylpType
	spObj.UpdateTime = now
	spObj.SetModified()
	m.ylpSpMap[ylpType][ylpId] = append(m.ylpSpMap[ylpType][ylpId], spObj)

	//检查是否集齐
	allFlag := true //是否全部部位已经镶嵌
	templateSuiPianMap := yingLingPuTemplate.GetSuiPianMap()
	for _, value := range templateSuiPianMap {
		if !m.ExistsYlpSp(ylpId, ylpType, value.SuipianId) { //不包含
			allFlag = false
			break
		}
	}

	if allFlag { //如果全部集齐，自动合成一个等级为0的英灵普
		m.AddYingLingPu(ylpId, ylpType)
	}
	return true
}

//升级英灵普
func (m *PlayerYingLingPuManager) UpYingLingPu(ylpId int32, ylpType ylptypes.YingLingPuTuJianType) bool {
	if !m.ExistsYlp(ylpId, ylpType) {
		return false
	}

	ylpObj := m.GetYlpInfo(ylpId, ylpType)
	ylpObj.Level = ylpObj.Level + 1
	now := global.GetGame().GetTimeService().Now()
	ylpObj.UpdateTime = now
	ylpObj.SetModified()
	return true
}

//添加英灵普
func (m *PlayerYingLingPuManager) AddYingLingPu(ylpId int32, ylpType ylptypes.YingLingPuTuJianType) bool {
	if m.ExistsYlp(ylpId, ylpType) {
		return false
	}
	ylpObj := NewYingLingPuObject(m.player)
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	ylpObj.Id = id
	ylpObj.Level = 0
	ylpObj.TuJianId = ylpId
	ylpObj.TuJianType = ylpType
	ylpObj.UpdateTime = now
	ylpObj.CreateTime = now
	ylpObj.SetModified()
	m.ylpMap[ylpType][ylpId] = ylpObj
	return true
}
