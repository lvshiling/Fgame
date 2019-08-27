package player

import (
	"fgame/fgame/core/storage"
	jieyientity "fgame/fgame/game/jieyi/entity"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

type PlayerJieYiObject struct {
	id              int64
	player          player.Player
	jieYiId         int64
	name            string
	rank            int32
	jieYiDaoJu      jieyitypes.JieYiDaoJuType
	tokenType       jieyitypes.JieYiTokenType
	tokenLev        int32
	tokenPro        int32
	tokenNum        int32
	shengWeiZhi     int32
	nameLev         int32
	namePro         int32
	nameNum         int32
	lastQiuYuanTime int64
	lastDropTime    int64
	lastInviteTime  int64
	lastPostTime    int64
	lastLeaveTime   int64
	updateTime      int64
	createTime      int64
	deleteTime      int64
}

func NewJieYiObject(pl player.Player) *PlayerJieYiObject {
	o := &PlayerJieYiObject{
		player: pl,
	}
	return o
}

func convertPlayerJieYiObjectToEntity(o *PlayerJieYiObject) (*jieyientity.PlayerJieYiEntity, error) {
	e := &jieyientity.PlayerJieYiEntity{
		Id:              o.id,
		PlayerId:        o.player.GetId(),
		JieYiId:         o.jieYiId,
		Name:            o.name,
		Rank:            o.rank,
		JieYiDaoJu:      int32(o.jieYiDaoJu),
		TokenType:       int32(o.tokenType),
		TokenLev:        o.tokenLev,
		TokenPro:        o.tokenPro,
		TokenNum:        o.tokenNum,
		ShengWeiZhi:     o.shengWeiZhi,
		NameLev:         o.nameLev,
		NamePro:         o.namePro,
		NameNum:         o.nameNum,
		LastQiuYuanTime: o.lastQiuYuanTime,
		LastDropTime:    o.lastDropTime,
		LastInviteTime:  o.lastInviteTime,
		LastPostTime:    o.lastPostTime,
		LastLeaveTime:   o.lastLeaveTime,
		UpdateTime:      o.updateTime,
		CreateTime:      o.createTime,
		DeleteTime:      o.deleteTime,
	}
	return e, nil
}

func (o *PlayerJieYiObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerJieYiObject) GetJieYiId() int64 {
	return o.jieYiId
}

func (o *PlayerJieYiObject) GetTokenNum() int32 {
	return o.tokenNum
}

func (o *PlayerJieYiObject) GetTokenPro() int32 {
	return o.tokenPro
}

func (o *PlayerJieYiObject) GetLastInviteTime() int64 {
	return o.lastInviteTime
}

func (o *PlayerJieYiObject) GetLastQiuYuanTime() int64 {
	return o.lastQiuYuanTime
}

func (o *PlayerJieYiObject) GetTokenLevel() int32 {
	return o.tokenLev
}

func (o *PlayerJieYiObject) GetJieYiName() string {
	return o.name
}

func (o *PlayerJieYiObject) GetTokenType() jieyitypes.JieYiTokenType {
	return o.tokenType
}

func (o *PlayerJieYiObject) HasToken() bool {
	return o.tokenType != -1
}

func (o *PlayerJieYiObject) GetLastPostTime() int64 {
	return o.lastPostTime
}

func (o *PlayerJieYiObject) GetDaoJuType() jieyitypes.JieYiDaoJuType {
	return o.jieYiDaoJu
}

func (o *PlayerJieYiObject) GetNameLev() int32 {
	return o.nameLev
}

func (o *PlayerJieYiObject) GetNamePro() int32 {
	return o.namePro
}

func (o *PlayerJieYiObject) GetNameNum() int32 {
	return o.nameNum
}

func (o *PlayerJieYiObject) GetShengWeiZhi() int32 {
	return o.shengWeiZhi
}

func (o *PlayerJieYiObject) GetLastLeaveTime() int64 {
	return o.lastLeaveTime
}

func (o *PlayerJieYiObject) GetJieYiRank() int32 {
	return o.rank
}

func (o *PlayerJieYiObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerJieYiObjectToEntity(o)
	return e, err
}

func (o *PlayerJieYiObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*jieyientity.PlayerJieYiEntity)

	o.id = pse.Id
	o.jieYiId = pse.JieYiId
	o.name = pse.Name
	o.rank = pse.Rank
	o.jieYiDaoJu = jieyitypes.JieYiDaoJuType(pse.JieYiDaoJu)
	o.tokenType = jieyitypes.JieYiTokenType(pse.TokenType)
	o.tokenPro = pse.TokenPro
	o.tokenNum = pse.TokenNum
	o.tokenLev = pse.TokenLev
	o.shengWeiZhi = pse.ShengWeiZhi
	o.nameLev = pse.NameLev
	o.nameNum = pse.NameNum
	o.namePro = pse.NamePro
	o.lastQiuYuanTime = pse.LastQiuYuanTime
	o.lastDropTime = pse.LastDropTime
	o.lastInviteTime = pse.LastInviteTime
	o.lastPostTime = pse.LastPostTime
	o.lastLeaveTime = pse.LastLeaveTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerJieYiObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerJieYi"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)

	return
}
