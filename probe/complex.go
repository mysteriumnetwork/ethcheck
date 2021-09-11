package probe

import (
	"context"
	"fmt"
	"time"

	"github.com/mysteriumnetwork/ethcheck/rpc"
)

func ComplexProbe(ctx context.Context, rpcClient rpc.Client, reqTimeout, lagTreshold time.Duration) error {
	err := LagProbe(ctx, rpcClient, reqTimeout, lagTreshold)
	if err != nil {
		return fmt.Errorf("lag probe failed: %w", err)
	}

	err = RandomProbe(ctx, rpcClient, reqTimeout)
	if err != nil {
		return fmt.Errorf("random probe failed: %w", err)
	}

	return nil
}
