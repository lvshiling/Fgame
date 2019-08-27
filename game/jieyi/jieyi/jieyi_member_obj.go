package jieyi

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	jieyientity "fgame/fgame/game/jieyi/entity"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/idutil"

	playertypes "fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

type JieYiMemberObject struct {
	id           int64
	serverId     int32
	jieYiId      int64
	jieYi        *JieYi
	playerId     int64
	name         string
	level        int32
	role         playertypes.RoleType
	sex          playertypes.SexType
	zhuanSheng   int32
	force        int64
	tokenType    jieyitypes.JieYiTokenType
	tokenLev     int32
	tokenPro     int32
	tokenNum     int32
	jieYiDaoJu   jieyitypes.JieYiDaoJuType
	onlineStatus playertypes.PlayerOnlineState
	shengWeiZhi  int32
	nameLev      int32
	jieYiTime    int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
	rank         int32
}

func newJieYiMemberObject(jieYi *JieYi) *JieYiMemberObject {
	o := &JieYiMemberObject{
		jieYi: jieYi,
	}
	return o
}

func convertJieYiMemberObjectToEntity(o *JieYiMemberObject) (*jieyientity.JieYiMemberEntity, error) {
	e := &jieyientity.JieYiMemberEntity{
		Id:          o.id,
		ServerId:    o.serverId,
		JieYiId:     o.jieYiId,
		PlayerId:    o.playerId,
		Name:        o.name,
		Level:       o.level,
		Role:        int32(o.role),
		Sex:         int32(o.sex),
		ZhuanSheng:  o.zhuanSheng,
		Force:       o.force,
		TokenType:   int32(o.tokenType),
		TokenLev:    o.tokenLev,
		TokenPro:    o.tokenPro,
		TokenNum:    o.tokenNum,
		JieYiDaoJu:  int32(o.jieYiDaoJu),
		ShengWeiZhi: o.shengWeiZhi,
		NameLev:     o.nameLev,
		JieYiTime:   o.jieYiTime,
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}
	return e, nil
}

func (o *JieYiMemberObject) GetDBId() int64 {
	return o.id
}

func (o *JieYiMemberObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *JieYiMemberObject) GetJieYiId() int64 {
	return o.jieYiId
}

func (o *JieYiMemberObject) GetPlayerName() string {
	return o.name
}

func (o *JieYiMemberObject) GetRole() playertypes.RoleType {
	return o.role
}

func (o *JieYiMemberObject) GetSex() playertypes.SexType {
	return o.sex
}

func (o *JieYiMemberObject) GetLevel() int32 {
	return o.level
}

func (o *JieYiMemberObject) GetForce() int64 {
	return o.force
}

func (o *JieYiMemberObject) GetZhuanSheng() int32 {
	return o.zhuanSheng
}

func (o *JieYiMemberObject) GetJieYiTime() int64 {
	return o.jieYiTime
}

func (o *JieYiMemberObject) GetTokenType() jieyitypes.JieYiTokenType {
	return o.tokenType
}
func (o *JieYiMemberObject) GetRank() int32 {
	return o.rank
}

func (o *JieYiMemberObject) HasToken() bool {
	return o.tokenType != -1
}

func (o *JieYiMemberObject) GetTokenLev() int32 {
	return o.tokenLev
}

func (o *JieYiMemberObject) GetDaoJuType() jieyitypes.JieYiDaoJuType {
	return o.jieYiDaoJu
}

func (o *JieYiMemberObject) GetOnLineState() playertypes.PlayerOnlineState {
	return o.onlineStatus
}

func (o *JieYiMemberObject) GetShengWeiZhi() int32 {
	return o.shengWeiZhi
}

func (o *JieYiMemberObject) GetNameLev() int32 {
	return o.nameLev
}

func (o *JieYiMemberObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertJieYiMemberObjectToEntity(o)
	return e, err
}

func (o *JieYiMemberObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*jieyientity.JieYiMemberEntity)

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.jieYiId = pse.JieYiId
	o.playerId = pse.PlayerId
	o.name = pse.Name
	o.level = pse.Level
	o.role = playertypes.RoleType(pse.Role)
	o.sex = playertypes.SexType(pse.Sex)
	o.zhuanSheng = pse.ZhuanSheng
	o.force = pse.Force
	o.tokenType = jieyitypes.JieYiTokenType(pse.TokenType)
	o.tokenLev = pse.TokenLev
	o.tokenPro = pse.TokenPro
	o.tokenNum = pse.TokenNum
	o.jieYiDaoJu = jieyitypes.JieYiDaoJuType(pse.JieYiDaoJu)
	o.onlineStatus = playertypes.PlayerOnlineStateOffline
	o.jieYiTime = pse.JieYiTime
	o.nameLev = pse.NameLev
	o.shengWeiZhi = pse.ShengWeiZhi
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *JieYiMemberObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "JieYiMember"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}

func (o *JieYiMemberObject) GetJieYi() *JieYi {
	return o.jieYi
}

func newJieYiMemberObjectWithPlayer(jieYi *JieYi, pl player.Player, daoJu jieyitypes.JieYiDaoJuType, token jieyitypes.JieYiTokenType, tokenLev, nameLev int32, now int64) *JieYiMemberObject {
	o := &JieYiMemberObject{
		jieYi: jieYi,
	}
	playerId := pl.GetId()

	id, _ := idutil.GetId()
	o.id = id
	o.jieYiId = jieYi.getJieYiId()
	o.serverId = global.GetGame().GetServerIndex()
	o.playerId = playerId
	o.tokenType = token
	o.tokenLev = tokenLev
	o.nameLev = nameLev
	o.createTime = now
	o.jieYiTime = now - 1
	o.jieYiDaoJu = daoJu
	o.onlineStatus = playertypes.PlayerOnlineStateOnline
	o.name = pl.GetName()
	o.level = pl.GetLevel()
	o.role = pl.GetRole()
	o.sex = pl.GetSex()
	o.zhuanSheng = pl.GetZhuanSheng()
	o.force = pl.GetForce()
	o.createTime = now
	return o
}
