package scan

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type isopen bool

func (o isopen) String() string {
	if o {
		return "open"
	}
	return "closed"
}

type Port struct {
	Num  int
	Open isopen
}

func validatePorts(ports []int) error {
	invalidPorts := make([]int, 0)
	for _, p := range ports {
		if p < 0 || p > 65535 {
			invalidPorts = append(invalidPorts, p)
		}
	}
	if len(invalidPorts) > 0 {
		return fmt.Errorf("invalid ports: %v", invalidPorts)
	}
	return nil
}

func ToPortList(rawPorts, rawRange string) ([]Port, error) {
	ps := strings.Split(rawPorts, ",")
	rs := strings.Split(rawRange, "-")
	ports := make([]int, 0)
	if len(rs) == 2 {
		from, err := strconv.Atoi(rs[0])
		if err != nil {
			return nil, err
		}
		to, err := strconv.Atoi(rs[1])
		if err != nil {
			return nil, err
		}
		for i := from; i <= to; i++ {
			ports = append(ports, i)
		}
	} else if len(rs) > 0 {
		return nil, fmt.Errorf("invalid range: %s", rawRange)
	}

	if len(ps) > 0 {
		for _, p := range ps {
			p, err := strconv.Atoi(p)
			if err != nil {
				return nil, err
			}
			index := slices.Index(ports, p)
			if index == -1 {
				ports = append(ports, p)
			}
		}
	}
	slices.Sort(ports)
	ports = slices.Compact(ports)
	if err := validatePorts(ports); err != nil {
		return nil, err
	}
	result := make([]Port, 0, len(ports))
	for _, p := range ports {
		result = append(result, Port{Num: p})
	}
	return result, nil
}

func (p Port) String() string {
	return fmt.Sprintf("%d\t%s", p.Num, p.Open)
}
