package scanner

import (
	"net"
	"strconv"
	"time"
)

type Result struct {
	Port   int
	Open   bool
	Banner string
}

type Scanner struct {
	target  string
	timeout time.Duration
}

func New(target string, timeout time.Duration) *Scanner {
	return &Scanner{target: target, timeout: timeout}
}

func (s *Scanner) ScanPort(port int) Result {
	addr := net.JoinHostPort(s.target, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", addr, s.timeout)
	if err != nil {
		return Result{Port: port, Open: false}
	}
	defer conn.Close()

	banner, err := readBanner(conn, 200*time.Millisecond)
	if err == nil {
		return Result{Port: port, Open: true, Banner: banner}
	}

	response, err := activeProbe(conn)
	if err == nil {
		return Result{Port: port, Open: true, Banner: "Active Response: " + response}
	}

	return Result{Port: port, Open: true, Banner: "Service remained completely silent"}
}

func readBanner(conn net.Conn, timeout time.Duration) (string, error) {
	conn.SetReadDeadline(time.Now().Add(timeout))
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func activeProbe(conn net.Conn) (string, error) {
	conn.SetReadDeadline(time.Time{})
	_, err := conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	if err != nil {
		return "", err
	}
	return readBanner(conn, 2*time.Second)
}
