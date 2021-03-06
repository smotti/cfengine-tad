package config

import (
	"flag"
	"time"

	"github.com/vharitonsky/iniflags"
)

var (
	Name     = flag.String("name", "tad", "Nick to use in IRC")
	Server   = flag.String("server", "127.0.0.1:6668", "Host:Port to connect to")
	Channels = flag.String("chan", "#tad", "Channels to join")
	Ssl      = flag.Bool("ssl", false, "Use SSL/TLS")
	Listen   = flag.Bool("listenChannel", false, "Listen for command on public channels")
	HostInfo = flag.String("hostInfo", "./data/va_host_info_report.json", "Path to host info report")
	Promises = flag.String("promises", "./data/promises_outcome.log", "Path to promises report")
	Classes  = flag.String("classes", "./data/classes.log", "Path to classes report")
	//	WatchDir      = flag.String("watchDir", "./data/", "Path to dir whose files should be watch for changes")
	WatchInterval  = flag.Duration("watchInterval", time.Duration(1)*time.Second, "Interval when to check files in watch dir (in seconds)")
	WatchReports   = flag.Bool("watchReports", false, "Watch report files for change every watchInterval and notify if changed")
	NotifyInterval = flag.Duration("notifyInterval", time.Duration(300)*time.Second, "Interval on when to notify about repaired and/or failed promises")
	NotifyRepaired = flag.Bool("notifyRepaired", true, "Notify about repaired promises")
	NotifyFailed   = flag.Bool("notifyFailed", true, "Notify about failed promises")
)

func init() {
	iniflags.Parse()
}
