package probe

import (
	"context"
	"fmt"
	"time"

	"github.com/mysteriumnetwork/ethcheck/rpc"
)

func LagProbe(ctx context.Context, rpcClient rpc.Client, reqTimeout time.Duration, lagTreshold time.Duration, blockTolerance int) error {
	now := time.Now().Truncate(0).UTC()

	blockNumber, err := getInspectedBlockNumber(ctx, rpcClient, reqTimeout, blockTolerance)
	if err != nil {
		return err
	}

	ctx1, cl := context.WithTimeout(ctx, reqTimeout)
	defer cl()

	var blockInfo *ShortBlockInfo
	err = rpcClient.Call(ctx1, &blockInfo, "eth_getBlockByNumber", blockNumber, false)
	if err != nil {
		return fmt.Errorf("eth_getBlockByNumber(): %w", err)
	}

	if blockInfo == nil {
		return fmt.Errorf("eth_getBlockByNumber(): block %v not found.", blockNumber)
	}

	blockTime := time.Unix(int64(blockInfo.Timestamp), 0).UTC()
	lag := now.Sub(blockTime)

	if lag >= lagTreshold {
		return fmt.Errorf("lag %s above treshold %s", lag, lagTreshold)
	}
	return nil
}

func getInspectedBlockNumber(ctx context.Context, rpcClient rpc.Client, reqTimeout time.Duration, blockTolerance int) (interface{}, error) {
	if blockTolerance == -1 {
		return "latest", nil
	}

	var blockNumberResult BlockNumberResult
	ctx1, cl := context.WithTimeout(ctx, reqTimeout)
	defer cl()

	err := rpcClient.Call(ctx1, &blockNumberResult, "eth_blockNumber")
	if err != nil {
		return nil, err
	}

	return blockNumberResult.Sub(blockTolerance), nil
}
