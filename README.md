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
	./pscan scan --hosts-file <path-to-hosts-file> --ports <ports> --range <port-range> --filter <open|closed> --timeout <timeout-in-ms>

You can use `.pscan.yaml` to setup port scan
### SEE ALSO

* [pscan docs](./docs/pscan_docs.md)	 - generate documentation for your command
* [pscan hosts](./docs/pscan_hosts.md)	 - manage the hosts list
* [pscan scan](./docs/pscan_scan.md)	 - scan hosts ports
  
## Installation

To use this package in your Go project, you can simply import it:

## Usage
Here's a basic example demonstrating how to use the PortScanner package:

```go
package main

import (
	"fmt"
	import "github.com/alex1988m/go-pscan"
)

func main() {
	// Initialize PortScanner with hosts and ports
	ps := portscanner.PortScanner{
		Hosts: []string{"example.com", "example.org"},
		Ports: []portscanner.Port{{Num: 80}, {Num: 443}},
		W:     os.Stdout, // Output writer
	}

	// Validate hosts
	err := ps.ValidateHosts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Scan ports
	ps.ScanPorts()

	// Sort results
	ps.SortResults()

	// Print results
	err = ps.PrintResults()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
```