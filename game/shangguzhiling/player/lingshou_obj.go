package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	shangguzhilingentity "fgame/fgame/game/shangguzhiling/entity"
	shangguzhilingtemplate "fgame/fgame/game/shangguzhiling/template"
	shangguzhilingtypes "fgame/fgame/game/shangguzhiling/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

type PlayerShangguzhilingObject struct {
	player        player.Player
	id            int64
	lingShouType  shangguzhilingtypes.LingshouType
	level         int32                                                                     //喂养等级
	experience    int64                                                                     //喂养经验条
	lingwen       map[shangguzhilingtypes.LingwenType]*shangguzhilingtypes.LingwenInfo      //灵纹
	uprankLevel   int32                                                                     //进阶等级
	uprankBless   int64                                                                     //进阶祝福值
	uprankTimes   int32                                                                     //进阶已尝试次数
	linglian      map[shangguzhilingtypes.LinglianPosType]*shangguzhilingtypes.LinglianInfo //灵炼
	linglianTimes int32
	receiveTime   int64
	updateTime    int64
	createTime    int64
	deleteTime    int64
}

// 用来刚解锁灵炼部位的时候，初始随机部位属性
func (o *PlayerShangguzhilingObject) refreshLingLianStatus() {
	lingshouTempService := shangguzhilingtemplate.GetShangguzhilingTemplateService()
	for pos, info := range o.linglian {
		linglianTemp := lingshouTempService.GetLingLianTemplate(o.lingShouType, pos)
		if linglianTemp == nil {
			continue
		}
		// 未解除限制（灵兽等级不够）
		if !o.isUnlock(linglianTemp.NeedSgzlLevel) {
			continue
		}

		// 正常执行下来的未初始随机属性的部位的是不会Lock的
		// if info.IsLock {
		// 	continue
		// }

		// 未初始随机属性
		if info.PoolMark == int32(0) {
			info.PoolMark = linglianTemp.GetBeginRandomPoolTempMark()
		}
	}
	// 局部函数暂时不需要保存
	// o.SetModified()
}

func (o *PlayerShangguzhilingObject) isUnlock(needLevel int32) bool {
	curLevel := o.level
	if needLevel <= curLevel {
		return true
	}
	return false
}

// 获取灵炼属性池模板
func (o *PlayerShangguzhilingObject) GetLingLianPoolTemplate(typ shangguzhilingtypes.LinglianPosType) *gametemplate.ShangguzhilingLinglianPoolTemplate {
	info := o.linglian[typ]
	linglianTemp := shangguzhilingtemplate.GetShangguzhilingTemplateService().GetLingLianTemplate(o.lingShouType, typ)
	if linglianTemp == nil {
		return nil
	}
	return linglianTemp.GetPoolTempById(info.PoolMark)
}

//灵兽类型
func (o *PlayerShangguzhilingObject) GetLingShouType() shangguzhilingtypes.LingshouType {
	return o.lingShouType
}

//灵纹信息
func (o *PlayerShangguzhilingObject) GetLingWenInfo(typ shangguzhilingtypes.LingwenType) *shangguzhilingtypes.LingwenInfo {
	return o.lingwen[typ]
}

//灵炼信息
func (o *PlayerShangguzhilingObject) GetLingLianInfo(typ shangguzhilingtypes.LinglianPosType) *shangguzhilingtypes.LinglianInfo {
	return o.linglian[typ]
}

//所有灵炼信息
func (o *PlayerShangguzhilingObject) GetLingLian() map[shangguzhilingtypes.LinglianPosType]*shangguzhilingtypes.LinglianInfo {
	return o.linglian
}

//喂养等级
func (o *PlayerShangguzhilingObject) GetLevel() int32 {
	return o.level
}

//喂养经验
func (o *PlayerShangguzhilingObject) GetExperience() int64 {
	return o.experience
}

// 进阶等级
func (o *PlayerShangguzhilingObject) GetUprankLevel() int32 {
	return o.uprankLevel
}

func (o *PlayerShangguzhilingObject) GetUprankBless() int64 {
	return o.uprankBless
}

func (o *PlayerShangguzhilingObject) GetUprankTimes() int32 {
	return o.uprankTimes
}

//灵炼次数
func (o *PlayerShangguzhilingObject) GetLingLianTimes() int32 {
	return o.linglianTimes
}

//上一次领取的时间
func (o *PlayerShangguzhilingObject) GetLastReceiveTime() int64 {
	return o.receiveTime
}

func NewPlayerShangguzhilingObject(pl player.Player) *PlayerShangguzhilingObject {
	obj := &PlayerShangguzhilingObject{
		player: pl,
	}
	return obj
}

func createNewPlayerShangguzhilingObject(pl player.Player, typ shangguzhilingtypes.LingshouType) *PlayerShangguzhilingObject {
	lingwenmap := make(map[shangguzhilingtypes.LingwenType]*shangguzhilingtypes.LingwenInfo)
	for i := shangguzhilingtypes.MinLingwenType; i <= shangguzhilingtypes.MaxLingwenType; i++ {
		lingwenmap[i] = &shangguzhilingtypes.LingwenInfo{
			Level:      int32(0),
			Experience: int64(0),
		}
	}

	linglianmap := make(map[shangguzhilingtypes.LinglianPosType]*shangguzhilingtypes.LinglianInfo)
	for i := shangguzhilingtypes.MinLinglianPosType; i <= shangguzhilingtypes.MaxLinglianPosType; i++ {
		linglianmap[i] = &shangguzhilingtypes.LinglianInfo{
			PoolMark: int32(0),
			// IsLock:   false,
		}
	}
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj := &PlayerShangguzhilingObject{
		id:            id,
		player:        pl,
		lingShouType:  typ,
		level:         int32(0),
		experience:    int64(0),
		lingwen:       lingwenmap,
		linglian:      linglianmap,
		uprankLevel:   int32(0),
		uprankBless:   int64(0),
		uprankTimes:   int32(0),
		linglianTimes: int32(0),
		receiveTime:   int64(0),
		createTime:    now,
	}
	return obj
}

func (o *PlayerShangguzhilingObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerShangguzhilingObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerShangguzhilingObject) ToEntity() (e storage.Entity, err error) {
	lingwenBytes, err := json.Marshal(o.lingwen)
	if err != nil {
		return nil, err
	}
	linglianBytes, err := json.Marshal(o.linglian)
	if err != nil {
		return nil, err
	}
	e = &shangguzhilingentity.PlayerShangguzhilingEntity{
		Id:            o.id,
		PlayerId:      o.player.GetId(),
		Level:         o.level,
		Experience:    o.experience,
		LingShouType:  int32(o.lingShouType),
		LingWen:       string(lingwenBytes),
		Linglian:      string(linglianBytes),
		UprankLevel:   o.uprankLevel,
		UprankBless:   o.uprankBless,
		UprankTimes:   o.uprankTimes,
		LinglianTimes: o.linglianTimes,
		ReceiveTime:   o.receiveTime,
		UpdateTime:    o.updateTime,
		CreateTime:    o.createTime,
		DeleteTime:    o.deleteTime,
	}
	return e, nil
}

func (o *PlayerShangguzhilingObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*shangguzhilingentity.PlayerShangguzhilingEntity)

	lingwen := make(map[shangguzhilingtypes.LingwenType]*shangguzhilingtypes.LingwenInfo)
	for i := shangguzhilingtypes.MinLingwenType; i <= shangguzhilingtypes.MaxLingwenType; i++ {
		lingwen[i] = &shangguzhilingtypes.LingwenInfo{}
	}
	err := json.Unmarshal([]byte(te.LingWen), &lingwen)
	if err != nil {
		return err
	}

	linglian := make(map[shangguzhilingtypes.LinglianPosType]*shangguzhilingtypes.LinglianInfo)
	for i := shangguzhilingtypes.MinLinglianPosType; i <= shangguzhilingtypes.MaxLinglianPosType; i++ {
		linglian[i] = &shangguzhilingtypes.LinglianInfo{}
	}
	err = json.Unmarshal([]byte(te.Linglian), &linglian)
	if err != nil {
		return err
	}

	o.id = te.Id
	o.level = te.Level
	o.experience = te.Experience
	o.lingShouType = shangguzhilingtypes.LingshouType(te.LingShouType)
	o.lingwen = lingwen
	o.linglian = linglian
	o.uprankLevel = te.UprankLevel
	o.uprankBless = te.UprankBless
	o.uprankTimes = te.UprankTimes
	o.linglianTimes = te.LinglianTimes
	o.receiveTime = te.ReceiveTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerShangguzhilingObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Shangguzhiling"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
