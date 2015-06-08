package config

import (
	"flag"

	"github.com/vharitonsky/iniflags"
)

var (
	Name     = flag.String("name", "tad", "Nick to use in IRC")
	Server   = flag.String("server", "127.0.0.1:6668", "Host:Port to connect to")
	Channels = flag.String("chan", "#tad", "Channels to join")
	Ssl      = flag.Bool("ssl", false, "Use SSL/TLS")
	Listen   = flag.Bool("listenChannel", false, "Listen for command on public channels")
)

const (
	HostInfoReport = "./data/va_host_info_report.json"
)

func init() {
	iniflags.Parse()
}
