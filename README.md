# Port Scanner

Port Scanner is a Go package for scanning network ports on multiple hosts concurrently.

## Introduction

This package provides functionality to scan ports on multiple hosts concurrently using Go routines. It supports scanning TCP ports and provides features such as validation of hostnames, printing results, and sorting results.

## Features

- Concurrent scanning of TCP ports on multiple hosts
- Validation of hostnames
- Printing scan results
- Sorting scan results by host and port number

## Build from Source

To run `pscan` as a binary:

    go build -o pscan
    ./pscan hosts add example.com
    ./pscan hosts list
	./pscan docs --dir ./docs
	./pscan scan --hosts-file <path-to-hosts-file> --ports <ports> --range <port-range> --filter <open|closed> --timeout <timeout-in-ms>

## Environment Variables:
	PSCAN_PORTS: Ports to scan within hosts.
	PSCAN_RANGE: Port range to scan within hosts.
	PSCAN_FILTER: Filter open or closed ports.
	PSCAN_TIMEOUT: Timeout in milliseconds for each port scan.

### You can use `.pscan.yaml` to setup port scan:
	ports: "22,80,443"
	range: "1000-1100"
	filter: "open"
	timeout: "500"

### SEE ALSO

* [pscan docs](./docs/pscan_docs.md)	 - generate documentation for your command
* [pscan hosts](./docs/pscan_hosts.md)	 - manage the hosts list
* [pscan scan](./docs/pscan_scan.md)	 - scan hosts ports
  
## Installation

To use this package in your Go project, you can simply import it:
```go get github.com/alex1988m/go-pscan ```
## Usage
Here's a basic example demonstrating how to use the PortScanner package:

```go
package main

import (
	"fmt"
	"os"

	"github.com/alex1988m/go-pscan/scan"
)

func main() {
	hosts := []string{"example.com", "localhost"}
	ports, err := scan.ToPortList("22,80", "1000-1010")
	if err != nil {
		fmt.Println("Error parsing ports:", err)
		return
	}

	portScanner := &scan.PortScanner{
		Hosts:   hosts,
		Ports:   ports,
		W:       os.Stdout,
		Filter:  "open", // Can be "open", "closed", or ""
		Timeout: 1000,   // Timeout in milliseconds
	}

	if err := portScanner.ValidateHosts(); err != nil {
		fmt.Println("Error validating hosts:", err)
		return
	}

	portScanner.ScanPorts()
	portScanner.SortResults()

	if err := portScanner.PrintResults(); err != nil {
		fmt.Println("Error printing results:", err)
	}
}
```