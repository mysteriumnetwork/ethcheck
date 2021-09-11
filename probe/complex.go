package probe

import (
	"context"
	"fmt"
	"time"

	"github.com/mysteriumnetwork/ethcheck/rpc"
)

func ComplexProbe(ctx context.Context, rpcClient rpc.Client, reqTimeout, lagTreshold time.Duration) error {
	ctx, cl := context.WithCancel(ctx)
	defer cl()

	responses := make(chan error)

	go func() {
		err := LagProbe(ctx, rpcClient, reqTimeout, lagTreshold)
		if err != nil {
			err = fmt.Errorf("lag probe failed: %w", err)
		}
		responses <- err
	}()

	go func() {
		err := RandomProbe(ctx, rpcClient, reqTimeout)
		if err != nil {
			err = fmt.Errorf("random probe failed: %w", err)
		}
		responses <- err
	}()

	for i := 0; i < 2; i++ {
		err := <-responses
		if err != nil {
			return err
		}
	}

	return nil
}
