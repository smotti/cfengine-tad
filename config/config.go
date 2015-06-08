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
	HostInfo = flag.String("hostInfo", "./data/va_host_info_report.json", "Path to host info report")
	Promises = flag.String("promises", "./data/promises.csv", "Path to promises report")
	Classes  = flag.String("classes", "./data/classes.txt", "Path to classes report")
)

func init() {
	iniflags.Parse()
}
