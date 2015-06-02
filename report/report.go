package report

import (
    "log"
)

var (
    reports = make(map[string]Report)
)

type (
	// Report defines an interface that reports must implement.
	Report interface {
		Read() error // Read contents of Filename
	}
)

// Register registers a report with the application.
func Register(reportType string, report Report) {
    if _, exists := reports[reportType]; exists {
        log.Fatalln(reportType, "Report already registered!")
    }

    log.Println("Register", reportType, "report")
    reports[reportType] = report
}
