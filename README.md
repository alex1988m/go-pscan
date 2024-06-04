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
    ./pscan scan --hosts-file pscan.hosts --ports 80,443     

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

## pscan cli

Fast TCP port scanner

### Synopsis

pscan - short for Port Scanner - executes TCP port scan on a list of hosts.
	pScan allows you to add, list, and delete hosts from the list.
	pScan executes a port scan on specified TCP ports. You can customize the
	target ports using a command line flag.

### Options

```
  --config string       config file (default is $HOME/.pscan.yaml)
  -h, --help                help for pscan
  -f, --hosts-file string   file to store hosts (default "pscan.hosts")
```

### SEE ALSO

* [pscan completion](./docs/pscan_completion.md)	 - Generate bash completion for your command
* [pscan docs](./docs/pscan_docs.md)	 - Generate documentation for your command
* [pscan hosts](./docs/pscan_hosts.md)	 - Manage the hosts list
* [pscan scan](./docs/pscan_scan.md)	 - scan hosts ports

