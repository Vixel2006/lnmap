package config

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Target  string
	Ports   []int
	Timeout time.Duration
}

func ParsePorts(raw string) ([]int, error) {
	var ports []int
	for _, part := range strings.Split(raw, ",") {
		part = strings.TrimSpace(part)
		if strings.Contains(part, "-") {
			bounds := strings.SplitN(part, "-", 2)
			if len(bounds) != 2 {
				return nil, fmt.Errorf("invalid port range: %s", part)
			}
			start, err := strconv.Atoi(strings.TrimSpace(bounds[0]))
			if err != nil {
				return nil, fmt.Errorf("invalid port start: %s", bounds[0])
			}
			end, err := strconv.Atoi(strings.TrimSpace(bounds[1]))
			if err != nil {
				return nil, fmt.Errorf("invalid port end: %s", bounds[1])
			}
			if start > end || start < 1 || end > 65535 {
				return nil, fmt.Errorf("invalid port range: %d-%d", start, end)
			}
			for p := start; p <= end; p++ {
				ports = append(ports, p)
			}
		} else {
			p, err := strconv.Atoi(part)
			if err != nil || p < 1 || p > 65535 {
				return nil, fmt.Errorf("invalid port: %s", part)
			}
			ports = append(ports, p)
		}
	}
	return ports, nil
}
