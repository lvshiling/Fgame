package login

import (
	"context"
	fgamedb "fgame/fgame/core/db"
	fgameredis "fgame/fgame/core/redis"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
)

//用户token
const (
	userTokenRedisKey = "user"
)

func getUserTokenRedisKey(userId int64) string {
	return fgameredis.Combine(userTokenRedisKey, fmt.Sprintf("%d", userId))
}

type LoginConfig struct {
	//毫秒
	Key         string `json:"key"`
	ExpiredTime int64  `json:"expiredTime"`
	AppId       string `json:"appId"`
	Secret      string `json:"secret"`
}

type LoginService interface {
	VisitLogin(deviceMac string) (t string, expiredTime int64, err error)

	Verify(tokenStr string) (int64, error)
	Register(userName string, password string) (t string, expiredTime int64, err error)
	Login(userName string, password string) (t string, expiredTime int64, err error)
}

type loginService struct {
	userConfig *LoginConfig
	key        []byte
	rs         fgameredis.RedisService
	db         fgamedb.DBService
	us         UserService
}

func (ls *loginService) Logout(id int64) (err error) {
	conn := ls.rs.Pool().Get()
	if conn.Err() != nil {
		err = conn.Err()
		return
	}
	defer conn.Close()

	userTokenKey := getUserTokenRedisKey(id)
	_, err = conn.Do("del", userTokenKey)
	if err != nil {
		return err
	}
	return nil
}

func (us *loginService) login(id int64) (t string, expiredTime int64, err error) {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	expiredTime += now + us.userConfig.ExpiredTime
	claims := &jwt.StandardClaims{}
	claims.ExpiresAt = expiredTime
	claims.Issuer = fmt.Sprintf("%d", id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	t, err = token.SignedString(us.key)
	if err != nil {
		return
	}
	//保存redis
	conn := us.rs.Pool().Get()
	err = conn.Err()
	if err != nil {
		return
	}
	defer conn.Close()

	userTokenKey := getUserTokenRedisKey(id)

	ok, err := redis.String(conn.Do("setex", userTokenKey, us.userConfig.ExpiredTime/1000, t))
	if err != nil {
		return
	}
	if ok != fgameredis.OK {
		err = fmt.Errorf("redis set failed %s", ok)
		return
	}
	return
}

func (us *loginService) Verify(tokenStr string) (id int64, err error) {
	claims := &jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) { return us.key, nil })
	if err != nil {
		return 0, err
	}

	idStr := claims.Issuer

	if len(idStr) == 0 {
		return 0, nil
	}

	id, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	//保存redis
	conn := us.rs.Pool().Get()
	err = conn.Err()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	userTokenKey := getUserTokenRedisKey(id)

	cacheToken, err := redis.String(conn.Do("get", userTokenKey))
	if err != nil {
		return
	}
	if cacheToken != tokenStr {
		return 0, nil
	}

	return id, nil
}

func (ls *loginService) VisitLogin(deviceMac string) (t string, expiredTime int64, err error) {

	user, err := ls.us.GetUserByDeviceMac(deviceMac)
	if err != nil {
		return
	}
	if user == nil {
		user, err = ls.us.Register(deviceMac, "")
		if err != nil {
			return
		}
	}

	t, expiredTime, err = ls.login(user.Id)
	return
}

func (ls *loginService) Register(name string, password string) (t string, expiredTime int64, err error) {

	user, err := ls.us.GetUserByName(name)
	if err != nil {
		return
	}
	if user == nil {
		user, err = ls.us.RegisterUser(name, password)
		if err != nil {
			return
		}
	}

	t, expiredTime, err = ls.login(user.Id)
	return
}

func (ls *loginService) Login(name string, password string) (t string, expiredTime int64, err error) {

	user, err := ls.us.GetUserByName(name)
	if err != nil {
		return
	}
	if user == nil {
		return
	}
	if !strings.EqualFold(user.Password, password) {
		return
	}
	t, expiredTime, err = ls.login(user.Id)
	return
}

func NewLoginService(uc *LoginConfig, db fgamedb.DBService, rs fgameredis.RedisService, us UserService) (ls LoginService, err error) {

	//读取key
	keyFile, err := filepath.Abs(uc.Key)
	if err != nil {
		return
	}
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return
	}
	ls = &loginService{
		userConfig: uc,
		key:        key,
		rs:         rs,
		db:         db,
		us:         us,
	}
	return
}

type contextKey string

const (
	loginServiceKey = "LoginService"
)

func WithLoginService(ctx context.Context, us LoginService) context.Context {
	return context.WithValue(ctx, loginServiceKey, us)
}

func LoginServiceInContext(ctx context.Context) LoginService {
	us, ok := ctx.Value(loginServiceKey).(LoginService)
	if !ok {
		return nil
	}
	return us
}

const (
	userContextKey = contextKey("User")
)

func WithUser(ctx context.Context, userId int64) context.Context {
	return context.WithValue(ctx, userContextKey, userId)
}

func UserInContext(ctx context.Context) int64 {
	userId, ok := ctx.Value(userContextKey).(int64)
	if !ok {
		return 0
	}
	return userId
}

func AuthHandler() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		us := LoginServiceInContext(req.Context())
		if us == nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		if ah := req.Header.Get("Authorization"); ah != "" {

			// Should be a bearer token
			if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
				id, err := us.Verify(ah[7:])
				if err != nil {
					rw.WriteHeader(http.StatusUnauthorized)
					return
				}
				if id == 0 {
					rw.WriteHeader(http.StatusUnauthorized)
					return
				}
				ctx := WithUser(req.Context(), id)
				nreq := req.WithContext(ctx)
				hf.ServeHTTP(rw, nreq)
				return
			}
		}
		rw.WriteHeader(http.StatusUnauthorized)
	})
}

func SetupLoginServiceHandler(us LoginService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithLoginService(ctx, us)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
