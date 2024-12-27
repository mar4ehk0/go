package server

import (
	"fmt"
	"strconv"

	"github.com/spf13/pflag"
)

type Addr struct {
	ip   string
	port int64
}

func NewAddr() *Addr {
	ip, port := readFlags()

	return &Addr{ip: ip, port: port}
}

func (a *Addr) Connection() string {
	addr := fmt.Sprintf("%s:%s", a.ip, strconv.Itoa(int(a.port)))

	return addr
}

func readFlags() (string, int64) {
	ip := pflag.StringP("ip", "i", "127.0.0.1", "Listen IP addr for server")
	port := pflag.Int64P("port", "p", 8080, "Listen Port for server")

	pflag.Parse()

	return *ip, *port
}
