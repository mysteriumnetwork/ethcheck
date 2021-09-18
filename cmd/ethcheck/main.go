package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	fixedDialer "github.com/mysteriumnetwork/ethcheck/dialer"
	"github.com/mysteriumnetwork/ethcheck/probe"
	"github.com/mysteriumnetwork/ethcheck/rpc"
)

var version = "undefined"

var (
	endpoint        = flag.String("url", "", "RPC endpoint URL")
	reqTimeout      = flag.Duration("req-timeout", 5*time.Second, "timeout for single request")
	totalTimeout    = flag.Duration("total-timeout", 20*time.Second, "whole operation timeout")
	lagTreshold     = flag.Duration("lag", 1*time.Minute, "allowed lag treshold")
	showVersion     = flag.Bool("version", false, "show program version and exit")
	addressOverride = flag.String("address-override", "", "force remote host address")
	portOverride    = flag.String("port-override", "", "force remote host port")
	blockTolerance  = flag.Int("block-tolerance", -1, "request metainfo for "+
		"latest_block - block_tolerance or -1 for symbolic name \"latest\"")
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

	var dialer fixedDialer.ContextDialer
	dialer = &net.Dialer{}
	dialer = fixedDialer.NewFixedDialer(*addressOverride, *portOverride, dialer)

	httpTransport := &http.Transport{
		DialContext:       dialer.DialContext,
		ForceAttemptHTTP2: true,
	}

	httpClient := rpc.NewDefaultHTTPClient()
	httpClient.Transport = httpTransport

	rpcClient := rpc.NewHTTPRPCClient(parsedURL, httpClient)

	ctx, cl := context.WithTimeout(context.Background(), *totalTimeout)
	defer cl()

	err = probe.ComplexProbe(ctx, rpcClient, *reqTimeout, *lagTreshold, *blockTolerance)
	if err != nil {
		log.Printf("complex probe failed: %v", err)
		return 1
	}

	return 0
}

func main() {
	log.Default().SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	os.Exit(run())
}
