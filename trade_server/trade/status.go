package trade

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	statusTradeItemAreadyNoExistOrSell = status.New(codes.FailedPrecondition, "商品已经下架或出售")
	statusTradeUploadMax               = status.New(codes.FailedPrecondition, "商品上架已达最大上限")
)
