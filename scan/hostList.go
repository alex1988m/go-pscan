package scan

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	ErrExists    = errors.New("host already in the list")
	ErrNotExists = errors.New("host not in the list")
)

type HostsList struct {
	Filename string
	W        io.Writer
	Hosts    []string
}

func (h *HostsList) Print() error {
	if err := h.Load(); err != nil {
		return nil
	}
	for _, host := range h.Hosts {
		if _, err := fmt.Fprintln(os.Stdout, host); err != nil {
			return err
		}
	}
	return nil
}

func (h *HostsList) search(host string) (bool, int) {
	for i, h := range h.Hosts {
		if h == host {
			return true, i
		}
	}
	return false, -1
}

func (h *HostsList) Add(hosts []string) error {
	if err := h.Load(); err != nil {
		return err
	}
	for _, host := range hosts {
		if found, _ := h.search(host); found {
			return fmt.Errorf("%w: %s", ErrExists, host)
		}
	}
	h.Hosts = append(h.Hosts, hosts...)
	if err := h.Save(); err != nil {
		return err
	}
	for _, host := range hosts {
		fmt.Fprintln(h.W, "Added", host)
	}
	return nil
}

func (h *HostsList) Remove(hosts []string) error {
	if err := h.Load(); err != nil {
		return err
	}
	for _, host := range hosts {
		if found, i := h.search(host); !found {
			return fmt.Errorf("%w: %s", ErrNotExists, host)
		} else {
			h.Hosts = append(h.Hosts[:i], h.Hosts[i+1:]...)
		}
	}
	if err := h.Save(); err != nil {
		return err
	}
	for _, host := range hosts {
		fmt.Fprintln(h.W, "Removed", host)
	}
	return nil
}

func (h *HostsList) Load() error {
	file, err := os.Open(h.Filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	defer file.Close()
	var hosts []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hosts = append(hosts, scanner.Text())
	}
	h.Hosts = hosts
	return nil
}

func (h *HostsList) Save() error {
	if err := os.WriteFile(h.Filename, []byte(strings.Join(h.Hosts, "\n")), 0644); err != nil {
		return err
	}
	return nil
}
