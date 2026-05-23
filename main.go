package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"lnmap/config"
	"lnmap/scanner"
)

func run() error {
	target := flag.String("target", "", "Target host or IP to scan")
	portsRaw := flag.String("ports", "22,80,443", "Comma-separated ports or ranges (e.g. 22,80,1-1000)")
	timeout := flag.Duration("timeout", 200*time.Millisecond, "TCP connection timeout (e.g. 200ms, 1s)")

	flag.Parse()

	if *target == "" {
		if flag.NArg() > 0 {
			*target = flag.Arg(0)
		} else {
			return fmt.Errorf("no target specified; use -target <host> or pass it as an argument")
		}
	}

	ports, err := config.ParsePorts(*portsRaw)
	if err != nil {
		return err
	}

	cfg := config.Config{
		Target:  *target,
		Ports:   ports,
		Timeout: *timeout,
	}

	sc := scanner.New(cfg.Target, cfg.Timeout)

	fmt.Printf("Scanning %s on %d port(s)...\n", cfg.Target, len(cfg.Ports))

	var wg sync.WaitGroup
	for _, port := range cfg.Ports {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()

			result := sc.ScanPort(port)
			if !result.Open {
				fmt.Printf("[-] tcp/%d Closed\n", result.Port)
			} else {
				fmt.Printf("[+] tcp/%d Open -> %s\n", result.Port, result.Banner)
			}
		}(port)
	}

	wg.Wait()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
