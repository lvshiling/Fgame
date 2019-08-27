package service

import (
	"context"
	gmdb "fgame/fgame/gm/gamegm/db"
	"fgame/fgame/gm/gamegm/gm/types"
	gmredis "fgame/fgame/gm/gamegm/redis"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	usmodel "fgame/fgame/gm/gamegm/gm/user/model"

	redis "github.com/chasex/redis-go-cluster"
	jwt "github.com/dgrijalva/jwt-go"
)

type GmUser struct {
	UserId      int64    `json:"user_id"`
	UserName    string   `json:"name"`
	Access      []string `json:"access"`
	Token       string   `json:"token"`
	Avator      string   `json:"avator"`
	ExpiredTime int64    `json:"expiredTime"`
}

type ILoginService interface {
	VerifyToken(token string) (dealerId int64, privilege int32, err error)
	Login(p_username string, p_password string) (userInfo *GmUser, err error)
	LoginOut(p_userid int64) error

	GetUserInfo(p_userid int64) (userInfo *GmUser, err error)
}

type LoginConfig struct {
	Key         string `json:"key"`
	ExpiredTime int64  `json:"expiredTime"`
}

type loginService struct {
	_cfg *LoginConfig
	_db  gmdb.DBService
	_rs  gmredis.RedisService
	_key []byte
}

func (ds *loginService) VerifyToken(tokenStr string) (dealerId int64, privilege int32, err error) {
	claims := &jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) { return ds._key, nil })
	if err != nil {
		return 0, 0, err
	}

	idStr := claims.Issuer

	if len(idStr) == 0 {
		return 0, 0, nil
	}
	//TODO 为什么不能类型转换int
	dealerId, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	//保存redis
	conn := ds._rs.Pool().Get()
	if conn.Err() != nil {
		err = conn.Err()
		return 0, 0, err
	}
	defer conn.Close()

	userTokenKey := GetGmUserKey(dealerId)

	cacheToken, err := redis.String(conn.Do("get", userTokenKey))
	if err != nil {
		return
	}
	if cacheToken != tokenStr {
		return 0, 0, nil
	}

	pKey := getGmUserPrivilegeKey(dealerId)
	privilegeInt, err := redis.Int(conn.Do("get", pKey))
	if err != nil {
		return
	}
	privilege = int32(privilegeInt)
	return
}

//登陆
func (ds *loginService) Login(p_username string, p_password string) (userInfo *GmUser, err error) {
	userInfo = &GmUser{}
	dbInfo := &usmodel.DBGmUserInfo{}

	tdb := ds._db.DB().Where("userName=? and psd=? and deleteTime=0", p_username, p_password).First(dbInfo)
	if tdb.Error != nil && tdb.Error != gorm.ErrRecordNotFound {
		err = tdb.Error
		return
	}
	if dbInfo.UserId <= 0 {
		return
	}

	plevel := types.PrivilegeLevel(int32(dbInfo.PrivilegeLevel))
	token, expiredTime, err := ds.login(dbInfo.UserId, plevel)
	if err != nil {
		return
	}
	userInfo.UserId = dbInfo.UserId
	userInfo.UserName = dbInfo.UserName
	userInfo.Token = token
	userInfo.ExpiredTime = expiredTime
	userInfo.Avator = dbInfo.Avator
	userInfo.Access = make([]string, 0)
	userInfo.Access = append(userInfo.Access, plevel.Code())
	return
}

func (ds *loginService) LoginOut(p_userid int64) (err error) {
	key := GetGmUserKey(p_userid)
	conn := ds._rs.Pool().Get()
	if conn.Err() != nil {
		err = conn.Err()
		return
	}
	defer conn.Close()

	_, err = redis.Int64(conn.Do("DEL", key))

	return
}

func (ds *loginService) login(id int64, privilegeLevel types.PrivilegeLevel) (t string, expiredTime int64, err error) {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	expiredTime += now + ds._cfg.ExpiredTime
	claims := &jwt.StandardClaims{}
	claims.ExpiresAt = expiredTime
	claims.Issuer = fmt.Sprintf("%d", id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	t, err = token.SignedString(ds._key)
	if err != nil {
		return
	}

	//保存redis
	conn := ds._rs.Pool().Get()
	if conn.Err() != nil {
		err = conn.Err()
		return
	}
	defer conn.Close()
	//TODO 合并成一个
	privilegeKey := getGmUserPrivilegeKey(id)
	ok, err := redis.String(conn.Do("setex", privilegeKey, ds._cfg.ExpiredTime/1000, int32(privilegeLevel)))
	if err != nil {
		return
	}

	if ok != gmredis.OK {
		err = fmt.Errorf("redis set failed %s", ok)
		return
	}

	// userSessionKey := getGmUserSessionKey(id)
	// ok, err = redis.String(conn.Do("setex", userSessionKey, ds._cfg.ExpiredTime/1000, sessionKey))
	// if err != nil {
	// 	return
	// }

	if ok != gmredis.OK {
		err = fmt.Errorf("redis set failed %s", ok)
		return
	}

	userTokenKey := GetGmUserKey(id)

	ok, err = redis.String(conn.Do("setex", userTokenKey, ds._cfg.ExpiredTime/1000, t))
	if err != nil {
		return
	}

	if ok != gmredis.OK {
		err = fmt.Errorf("redis set failed %s", ok)
		return
	}
	return
}

func (ds *loginService) GetUserInfo(p_userid int64) (userInfo *GmUser, err error) {
	userInfo = &GmUser{}
	dbInfo := &usmodel.DBGmUserInfo{}

	tdb := ds._db.DB().Where("id = ?", p_userid).First(dbInfo)
	if tdb.Error != nil && tdb.Error != gorm.ErrRecordNotFound {
		err = tdb.Error
		return
	}
	if dbInfo.UserId <= 0 {
		return
	}

	plevel := types.PrivilegeLevel(int32(dbInfo.PrivilegeLevel))

	userInfo.UserId = dbInfo.UserId
	userInfo.UserName = dbInfo.UserName
	userInfo.Avator = dbInfo.Avator
	userInfo.Access = make([]string, 0)
	userInfo.Access = append(userInfo.Access, plevel.Code())
	return
}

const (
	loginServiceKey = contextKey("LoginService")
)

func WithLoginService(ctx context.Context, ls ILoginService) context.Context {
	return context.WithValue(ctx, loginServiceKey, ls)
}

func LoginServiceInContext(ctx context.Context) ILoginService {
	us, ok := ctx.Value(loginServiceKey).(ILoginService)
	if !ok {
		return nil
	}
	return us
}

func NewLoginService(cfg *LoginConfig, db gmdb.DBService, rs gmredis.RedisService) (tlos ILoginService, err error) {
	//读取key
	keyFile, err := filepath.Abs(cfg.Key)
	if err != nil {
		return
	}
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return
	}
	los := &loginService{}
	los._cfg = cfg
	los._db = db
	los._rs = rs
	los._key = key
	tlos = los
	return
}
