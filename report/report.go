package report

import (
	"bufio"
	"crypto/md5"
	"io"
	"log"
	"os"

	"github.com/sorcix/irc"
)

var (
	Reports = make(map[string]Report)
)

type (
	// Report defines an interface that reports must implement.
	Report interface {
		Read() error             // Read contents of Filename
		Watch(chan *irc.Message) // Launches a go routine to watch for changes in report file
	}
)

const FILE_TYPES = ".*\\.(json|txt|csv|log)"

// Register registers a report with the application and reads it into memory.
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

// calcHashSum uses the MD5 hash algorithm to calculate a hash sum of the
// content for the given file. It returns the sum as a byte slice.
func calcHashSum(s string) ([]byte, error) {
	// Open the file.
	file, err := os.Open(s)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get a new hash value.
	h := md5.New()

	// Read line by line and add to the hash sum.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		io.WriteString(h, scanner.Text())
	}

	return h.Sum(nil), nil
}
