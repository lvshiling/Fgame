package api

import (
	"context"
	"fgame/fgame/trade_server/pb"
	"fgame/fgame/trade_server/remote"
	"fgame/fgame/trade_server/trade"

	log "github.com/Sirupsen/logrus"
)

//交易服务器
type TradeManagerServer struct {
	m *trade.TradeServer
	r remote.RemoteService
}

// 上架
func (s *TradeManagerServer) GetTradeList(ctx context.Context, in *pb.TradeItemListRequest) (res *pb.TradeItemListResponse, err error) {
	platform := in.GetPlatform()
	serverId := in.GetServerId()

	tradeObjList := s.m.GetTradeList(platform, serverId)
	res = &pb.TradeItemListResponse{}
	res.TradeItemList = BuildTradeItemList(tradeObjList)
	return
}

// 上架
func (s *TradeManagerServer) Upload(ctx context.Context, in *pb.TradeUploadItemRequest) (res *pb.TradeUploadItemResponse, err error) {
	platform := in.GetPlatform()
	serverId := in.GetServerId()
	tradeId := in.GetTradeId()
	playerId := in.GetPlayerId()
	playerName := in.GetPlayerName()
	itemId := in.GetItemId()
	itemNum := in.GetItemNum()
	gold := in.GetGold()
	level := in.GetLevel()
	propertyData := in.GetPropertyData()
	tradeObj, err := s.m.Upload(platform, serverId, tradeId, playerId, playerName, itemId, itemNum, propertyData, gold, level)
	if err != nil {
		log.WithFields(
			log.Fields{
				"platform":     platform,
				"serverId":     serverId,
				"tradeId":      tradeId,
				"playerId":     playerId,
				"playerName":   playerName,
				"itemId":       itemId,
				"itemNum":      itemNum,
				"gold":         gold,
				"propertyData": propertyData,
				"err":          err,
			}).Error("trade:上传错误")
		return
	}
	log.WithFields(
		log.Fields{
			"platform":     platform,
			"serverId":     serverId,
			"tradeId":      tradeId,
			"playerId":     playerId,
			"playerName":   playerName,
			"itemId":       itemId,
			"itemNum":      itemNum,
			"gold":         gold,
			"propertyData": propertyData,
		}).Info("trade:上传成功")
	res = &pb.TradeUploadItemResponse{}
	globalTradeItem := BuildTradeItem(tradeObj)
	res.TradeItem = globalTradeItem
	return
}

// 下架
func (s *TradeManagerServer) Withdraw(ctx context.Context, in *pb.TradeWithdrawItemRequest) (res *pb.TradeWithdrawItemResponse, err error) {
	tradeId := in.GetGlobalTradeId()
	platform := in.GetPlatform()
	serverId := in.GetServerId()
	_, err = s.m.Withdraw(platform, serverId, tradeId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"platform": platform,
				"serverId": serverId,
				"tradeId":  tradeId,
				"err":      err,
			}).Error("trade:下架错误")
		return
	}
	log.WithFields(
		log.Fields{
			"platform": platform,
			"serverId": serverId,
			"tradeId":  tradeId,
		}).Info("trade:下架成功")
	res = &pb.TradeWithdrawItemResponse{}
	res.GlobalTradeId = tradeId
	res.Platform = platform
	res.ServerId = serverId
	return
}

// 交易
func (s *TradeManagerServer) Trade(ctx context.Context, in *pb.TradeItemRequest) (res *pb.TradeItemResponse, err error) {
	platform := in.GetPlatform()
	serverId := in.GetServerId()
	buyPlayerId := in.GetBuyPlayerId()
	buyPlayerName := in.GetBuyPlayerName()
	tradeId := in.GetGlobalTradeId()
	tradeObj, err := s.m.TradeItem(tradeId, platform, serverId, buyPlayerId, buyPlayerName)
	if err != nil {
		log.WithFields(
			log.Fields{
				"platform":      platform,
				"serverId":      serverId,
				"tradeId":       tradeId,
				"buyPlayerId":   buyPlayerId,
				"buyPlayerName": buyPlayerName,
				"err":           err,
			}).Error("trade:交易错误")
		return
	}
	log.WithFields(
		log.Fields{
			"platform":      platform,
			"serverId":      serverId,
			"tradeId":       tradeId,
			"buyPlayerId":   buyPlayerId,
			"buyPlayerName": buyPlayerName,
		}).Info("trade:交易成功")
	s.r.Sell(tradeObj)
	res = &pb.TradeItemResponse{}
	globalTradeItem := BuildTradeItem(tradeObj)
	res.TradeItem = globalTradeItem
	return
}

func NewTradeManagerServer(m *trade.TradeServer, r remote.RemoteService) *TradeManagerServer {
	ss := &TradeManagerServer{
		m: m,
		r: r,
	}
	return ss
}
