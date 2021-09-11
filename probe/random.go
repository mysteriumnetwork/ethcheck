package probe

import (
	"context"
	"fmt"
	"time"

	"github.com/mysteriumnetwork/ethcheck/rpc"
)

func RandomProbe(ctx context.Context, rpcClient rpc.Client, reqTimeout time.Duration) error {
	paramObject, err := NewRandomEthCallParamObject()
	if err != nil {
		return fmt.Errorf("ethcall(): can't create param object: %w", err)
	}

	ctx1, cl := context.WithTimeout(ctx, reqTimeout)
	defer cl()

	var result string
	err = rpcClient.Call(ctx1, &result, "eth_call", paramObject, "latest")
	if err != nil {
		return err
	}

	return nil
}
