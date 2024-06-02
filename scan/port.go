package scan

import (
	"fmt"
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

func ToPortList(raw string) ([]Port, error) {
	ports := make([]Port, 0, len(raw))
	for _, p := range strings.Split(raw, ",") {
		num, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		ports = append(ports, Port{Num: num})
	}
	return ports, nil
}

func (p Port) String() string {
	return fmt.Sprintf("%d: %s", p.Num, p.Open)
}
