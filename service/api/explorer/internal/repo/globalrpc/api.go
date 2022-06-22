package globalrpc

import (
	"context"

	"github.com/zecrey-labs/zecrey-legend/common/commonAsset"
	"github.com/zecrey-labs/zecrey-legend/common/model/account"
	"github.com/zecrey-labs/zecrey-legend/common/model/mempool"
	"github.com/zecrey-labs/zecrey-legend/service/api/explorer/internal/svc"
	"github.com/zecrey-labs/zecrey-legend/service/rpc/globalRPC/globalRPCProto"
	"github.com/zecrey-labs/zecrey-legend/service/rpc/globalRPC/globalrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type GlobalRPC interface {
	GetLatestAccountInfo(accountIndex int64) (accountInfo *commonAsset.AccountInfo, err error)
	SendTx(txType uint32, txInfo string) (string, error)
	GetLatestTxsListByAccountIndex(accountIndex uint32, limit, offset uint32) ([]*mempool.MempoolTx, uint32, error)
	GetLatestAssetsListByAccountIndex(accountIndex uint32) ([]*globalrpc.AssetResult, error)
	GetLpValue(pairIndex uint32, lpAmount string) (*globalRPCProto.RespGetLpValue, error)
	GetPairInfo(pairIndex uint32) (*globalRPCProto.RespGetLatestPairInfo, error)
	GetSwapAmount(ctx context.Context, pairIndex, assetId uint64, assetAmount string, isFrom bool) (string, uint32, error)
	GetLatestTxsListByAccountIndexAndTxType(accountIndex uint64, txType uint64, limit uint64, offset uint64) ([]*mempool.MempoolTx, error)
	GetMaxOfferId(accountIndex uint32) (uint64, error)
}

func New(svcCtx *svc.ServiceContext, ctx context.Context) GlobalRPC {
	return &globalRPC{
		AccountModel:        account.NewAccountModel(svcCtx.Conn, svcCtx.Config.CacheRedis, svcCtx.GormPointer),
		AccountHistoryModel: account.NewAccountHistoryModel(svcCtx.Conn, svcCtx.Config.CacheRedis, svcCtx.GormPointer),
		MempoolModel:        mempool.NewMempoolModel(svcCtx.Conn, svcCtx.Config.CacheRedis, svcCtx.GormPointer),
		MempoolDetailModel:  mempool.NewMempoolDetailModel(svcCtx.Conn, svcCtx.Config.CacheRedis, svcCtx.GormPointer),
		RedisConnection:     svcCtx.RedisConn,
		globalRPC:           globalrpc.NewGlobalRPC(zrpc.MustNewClient(svcCtx.Config.GlobalRpc)),
		ctx:                 ctx,
	}
}
