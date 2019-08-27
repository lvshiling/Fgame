package center

import (
	"context"
	centerpb "fgame/fgame/center/pb"

	log "github.com/Sirupsen/logrus"
	redis "github.com/chasex/redis-go-cluster"
)

const (
	forbidIpKey = "fgame.center.forbidip"
)

func (s *CenterServer) GetForbidIpList(ctx context.Context, req *centerpb.ForbidIpListRequest) (res *centerpb.ForbidIpListResponse, err error) {
	log.WithFields(
		log.Fields{}).Info("server:获取封禁ip")
	conn := s.rs.Pool().Get()
	err = conn.Err()
	if err != nil {
		return
	}
	defer conn.Close()
	ipList, err := redis.Strings(conn.Do("smembers", forbidIpKey))
	if err != nil {
		return
	}
	for _, ip := range ipList {
		res.IpList = append(res.IpList, ip)
	}
	log.WithFields(
		log.Fields{}).Info("server:获取封禁ip成功")
	return
}

func (s *CenterServer) ForbidIp(ctx context.Context, req *centerpb.ForbidIpRequest) (res *centerpb.ForbidIpResponse, err error) {
	forbidIp := req.GetIp()
	forbid := req.GetForbid()
	log.WithFields(
		log.Fields{
			"forbidIp": forbidIp,
			"forbid":   forbid,
		}).Info("server:封禁ip")
	conn := s.rs.Pool().Get()
	err = conn.Err()
	if err != nil {
		return
	}
	defer conn.Close()
	if forbid {
		_, err = conn.Do("sadd", forbidIpKey, forbidIp)
		if err != nil {
			return
		}
	} else {
		_, err = conn.Do("srem", forbidIpKey, forbidIp)
		if err != nil {
			return
		}
	}

	res = &centerpb.ForbidIpResponse{}
	res.Ip = forbidIp
	res.Forbid = forbid

	log.WithFields(
		log.Fields{
			"forbidIp": forbidIp,
			"forbid":   forbid,
		}).Info("server:封禁ip成功")

	return
}

func (s *CenterServer) isIpForbid(ip string) (flag bool, err error) {

	conn := s.rs.Pool().Get()
	err = conn.Err()
	if err != nil {
		return
	}
	defer conn.Close()

	intFlag, err := redis.Int(conn.Do("sismember", forbidIpKey, ip))
	if err != nil {
		return
	}
	flag = true
	if intFlag == 0 {
		flag = false
	}

	return
}
func (s *CenterServer) ForbidSearch(ctx context.Context, req *centerpb.ForbidIpSearchRequest) (res *centerpb.ForbidIpSearchResponse, err error) {
	forbidIp := req.GetIp()

	log.WithFields(
		log.Fields{
			"forbidIp": forbidIp,
		}).Info("server:查找封禁ip")
	conn := s.rs.Pool().Get()
	err = conn.Err()
	if err != nil {
		return
	}
	defer conn.Close()

	forbid, err := s.isIpForbid(forbidIp)
	if err != nil {
		return
	}
	res = &centerpb.ForbidIpSearchResponse{}
	res.Ip = forbidIp
	res.Forbid = forbid
	log.WithFields(
		log.Fields{
			"forbidIp": forbidIp,
		}).Info("server:查找封禁ip")
	return
}

func (s *CenterServer) ForbidUser(ctx context.Context, req *centerpb.ForbidUserRequest) (res *centerpb.ForbidUserResponse, err error) {
	forbid := req.GetForbid()
	userId := req.GetUserId()
	res = &centerpb.ForbidUserResponse{}
	res.UserId = userId
	res.Forbid = forbid
	return
}

func (s *CenterServer) ForbidUserSearch(ctx context.Context, req *centerpb.ForbidUserSearchRequest) (res *centerpb.ForbidUserSearchResponse, err error) {
	userId := req.GetUserId()

	res = &centerpb.ForbidUserSearchResponse{}
	res.UserId = userId
	res.Forbid = true
	return
}
