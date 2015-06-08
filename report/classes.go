package report

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/smotti/tad/config"
)

type (
	// Context defines a report about CFEngine defined classes.
	Context struct {
		Filename string
		Classes  []string
	}
)

// init registers the report at application startup.
func init() {
	var r Report
	r = &Context{
		Filename: *config.Classes,
	}

	Register("classes", r)
}

// Read implements the Report interface for Context.
func (c *Context) Read() error {
	file, err := os.Open(c.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c.Classes = append(c.Classes, scanner.Text())
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		log.Println("Error:", err)
	}

	return nil
}
