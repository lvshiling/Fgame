package service

import (
	"errors"
	fclient "fgame/fgame/game/remote/client"
	gmdb "fgame/fgame/gm/gamegm/db"

	"google.golang.org/grpc"
)

var grpcMap map[int32]*grpc.ClientConn
var clientMap map[int32]fclient.RemoteClient

func RegisterGrpc(p_serverId int32, p_host string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(p_host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	grpcMap[p_serverId] = conn
	clientMap[p_serverId] = fclient.NewRemoteClient(conn)
	return conn, err
}

func GetGrpcConn(p_serverId int32) (*grpc.ClientConn, error) {
	if len(grpcMap) > 0 {
		finalServerId := gmdb.GetFinnalServerId(p_serverId)
		if value, ok := grpcMap[finalServerId]; ok {
			return value, nil
		}
		// for _, value := range grpcMap {
		// 	return value, nil
		// }
	}
	return nil, errors.New("无服务")
}

func GetFClientRemote(p_service int32) (fclient.RemoteClient, error) {
	if len(clientMap) > 0 {
		finalServerId := gmdb.GetFinnalServerId(p_service)

		if value, ok := clientMap[finalServerId]; ok {
			return value, nil
		}
		// for _, value := range clientMap {
		// 	return value, nil
		// }
	}
	return nil, errors.New("无服务")
}

func init() {
	grpcMap = make(map[int32]*grpc.ClientConn)
	clientMap = make(map[int32]fclient.RemoteClient)
}
