ethcheck
========

Checks Ethereum full node RPC API to determine whether it is available, up to date and responsive within certain time constraints.

## Installation

#### Binaries

Pre-built binaries are available [here](https://github.com/mysteriumnetwork/ethcheck/releases/latest).

#### Build from source

Alternatively, you may install application from source. Run the following within the source directory:

```
make install
```

## Usage

Program is intended to work as an external health check command for load balancers and reverse proxies. It returns zero exit code if remote destination is healthy and non-zero exit code otherwise.

## Synopsis

```
$ ethcheck -h
Usage of ethcheck:
  -address-override string
    	force remote host address
  -lag duration
    	allowed lag treshold (default 1m0s)
  -port-override string
    	force remote host port
  -req-timeout duration
    	timeout for single request (default 5s)
  -total-timeout duration
    	whole operation timeout (default 20s)
  -url string
    	RPC endpoint URL
  -version
    	show program version and exit
```
