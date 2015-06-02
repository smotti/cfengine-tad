package report

type (
	// Report defines an interface that reports must implement.
	Report interface {
		Read() error // Read contents of Filename
	}
)
