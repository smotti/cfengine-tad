package report

import "log"

var (
	Reports = make(map[string]Report)
)

type (
	// Report defines an interface that reports must implement.
	Report interface {
		Read() error // Read contents of Filename
	}
)

// register registers a report with the application and reads it into memory.
func Register(reportType string, report Report) {
	if _, exists := Reports[reportType]; exists {
		log.Fatalln(reportType, "Report already registered")
	}

	if err := report.Read(); err != nil {
		log.Println("Error:", err)
	}

	log.Println("Register", reportType, "report")
	Reports[reportType] = report
}
