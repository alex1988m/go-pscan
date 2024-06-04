package scan

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

type result struct {
	host string
	port Port
	err  error
}

func (r result) String() string {
	if r.err != nil {
		return fmt.Sprintf("%s\t%s\t%s", r.host, r.port, r.err.Error())
	}
	return fmt.Sprintf("%s\t%s", r.host, r.port)
}

type results []result

func (r results) Len() int      { return len(r) }
func (r results) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r results) Less(i, j int) bool {
	if r[i].host == r[j].host {
		return r[i].port.Num < r[j].port.Num
	}
	return r[i].host < r[j].host
}

type PortScanner struct {
	Hosts   []string
	Ports   []Port
	W       io.Writer
	results results
}

func (ps *PortScanner) ValidateHosts() error {
	var wg sync.WaitGroup
	errors := make([]string, 0)
	hostCh := make(chan string, len(ps.Hosts))
	errCh := make(chan error, len(ps.Hosts))
	wg.Add(1)
	go ps.generateHosts(hostCh, &wg)
	cpuNum := runtime.NumCPU()
	wg.Add(cpuNum)
	for i := 0; i < cpuNum; i++ {
		go ps.validateHostWorker(&wg, hostCh, errCh)
	}

	go ps.awaitThreads(&wg, errCh)
	errors = ps.collectErrors(errCh, errors)
	if len(errors) > 0 {
		return fmt.Errorf("hosts validation error: \n%s", strings.Join(errors, "\n"))
	}
	return nil
}

func (ps *PortScanner) PrintResults() error {
	errors := make([]string, 0)
	for _, result := range ps.results {
		_, err := fmt.Fprintln(ps.W, result)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("print results error: \n%s", strings.Join(errors, "\n"))
	}
	return nil
}

func (ps *PortScanner) SortResults() {
	sort.Slice(ps.results, func(i, j int) bool {
		if ps.results[i].host == ps.results[j].host {
			return ps.results[i].port.Num < ps.results[j].port.Num
		}
		return ps.results[i].host < ps.results[j].host
	})
}

func (ps *PortScanner) ScanPorts() {
	var wg sync.WaitGroup
	resCh := make(chan result, len(ps.Hosts)*len(ps.Ports))
	dataCh := make(chan result, len(ps.Hosts)*len(ps.Ports))
	wg.Add(1)
	go ps.generateHostPort(&wg, dataCh)

	cpuNum := runtime.NumCPU()
	wg.Add(cpuNum)
	for i := 0; i < cpuNum; i++ {
		go ps.scanPortWorker(&wg, dataCh, resCh)
	}
	go func() {
		wg.Wait()
		close(resCh)
	}()
	for result := range resCh {
		ps.results = append(ps.results, result)
	}
}

func (*PortScanner) scanPortWorker(wg *sync.WaitGroup, dataCh chan result, resCh chan result) {
	defer wg.Done()
	for data := range dataCh {
		host := data.host
		port := data.port
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port.Num), time.Second)
		if err == nil {
			port.Open = true
			_ = conn.Close()
		}
		resCh <- result{host, port, err}
	}
}

func (ps *PortScanner) generateHostPort(wg *sync.WaitGroup, dataCh chan result) {
	defer close(dataCh)
	defer wg.Done()
	for _, host := range ps.Hosts {
		for _, port := range ps.Ports {
			dataCh <- result{host, port, nil}
		}
	}
}

func (*PortScanner) collectErrors(errCh chan error, errors []string) []string {
	for err := range errCh {
		errors = append(errors, err.Error())
	}
	return errors
}

func (*PortScanner) awaitThreads(wg *sync.WaitGroup, errCh chan error) {
	wg.Wait()
	close(errCh)
}

func (ps *PortScanner) generateHosts(hostCh chan string, wg *sync.WaitGroup) {
	defer close(hostCh)
	defer wg.Done()
	for _, host := range ps.Hosts {
		hostCh <- host
	}
}

func (ps *PortScanner) validateHostWorker(wg *sync.WaitGroup, hostCh chan string, errCh chan error) {
	defer wg.Done()
	for host := range hostCh {
		if _, err := net.LookupHost(host); err != nil {
			errCh <- err
		}
	}
}
