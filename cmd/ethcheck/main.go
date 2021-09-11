package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/mysteriumnetwork/ethcheck/probe"
	"github.com/mysteriumnetwork/ethcheck/rpc"
)

var version = "undefined"

var (
	endpoint     = flag.String("url", "", "RPC endpoint URL")
	reqTimeout   = flag.Duration("req-timeout", 5*time.Second, "timeout for single request")
	totalTimeout = flag.Duration("total-timeout", 20*time.Second, "whole operation timeout")
	lagTreshold  = flag.Duration("lag", 1*time.Minute, "allowed lag treshold")
	showVersion  = flag.Bool("version", false, "show program version and exit")
)

func run() int {
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
		return 0
	}

	if *endpoint == "" {
		log.Fatalln("url flag must be specified")
	}

	parsedURL, err := url.ParseRequestURI(*endpoint)
	if err != nil {
		log.Fatalf("bad URL specified: %v", err)
	}

	rpcClient := rpc.NewHTTPRPCClient(parsedURL, nil)

	ctx, cl := context.WithTimeout(context.Background(), *totalTimeout)
	defer cl()

	err = probe.ComplexProbe(ctx, rpcClient, *reqTimeout, *lagTreshold)
	if err != nil {
		log.Fatalf("complex probe failed: %v", err)
	}

	return 0
}

func main() {
	log.Default().SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	os.Exit(run())
}
