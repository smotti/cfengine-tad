package report

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"time"

	"github.com/smotti/tad/config"
	"github.com/sorcix/irc"
)

type (
	// Context defines a report about CFEngine defined classes.
	Context struct {
		Filename   string
		Classes    []string
		Checksum   []byte
		OldModTime time.Time
	}
)

// init registers the report at application startup.
func init() {
	h, err := calcHashSum(*config.Classes)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	var r Report
	r = &Context{
		Filename: *config.Classes,
		Checksum: h,
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

	c.Classes = nil
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c.Classes = append(c.Classes, scanner.Text())
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Watch starts a go routine for the report and watches its source file for
// changes. It implements the Report interface.
func (co *Context) Watch(c chan *irc.Message) {
	go func() {
		for {
			// Continuously watch file but only check hash sum if ModTime
			// changed.
			fi, err := os.Stat(co.Filename)
			if os.IsNotExist(err) {
				continue
			}

			if fi.ModTime().After(co.OldModTime) {
				// Calc new hash sum.
				newSum, err := calcHashSum(co.Filename)
				if err != nil {
					log.Println("Error:", err)
				}
				// Check if new and old sum differ.
				if !bytes.Equal(newSum, co.Checksum) {
					// Reread the report.
					if err := co.Read(); err != nil {
						log.Println("Error:", err)
					} else {
						// Set new checksum.
						co.Checksum = newSum

						log.Println("Checksum changed for", co.Filename)
						c <- &irc.Message{ // Send message to irc server.
							Command:  irc.PRIVMSG,
							Params:   []string{*config.Channels},
							Trailing: "Checksum changed for " + co.Filename,
						}
					}
				}
			}

			time.Sleep(*config.WatchInterval)
		}
	}()
}
