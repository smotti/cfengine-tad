package report

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/smotti/tad/config"
	"github.com/sorcix/irc"
)

type (
	// TODO: Don't forget to add Promisee.
	Promise struct {
		Class    string
		Handler  string
		Promiser string
		Promisee string
		Outcome  string
	}

	Promises struct {
		Filename string
		List     []*Promise
		Checksum []byte
	}
)

// init registers the report at application startup.
func init() {
	h, err := calcHashSum(*config.Promises)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	var r Report
	r = &Promises{
		Filename: *config.Promises,
		Checksum: h,
	}

	Register("promises", r)
}

// Read implements the Report interface for Promises.
func (p *Promises) Read() error {
	file, err := os.Open(p.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	rdr := csv.NewReader(file)
	rdr.Comma = ';'
	rdr.FieldsPerRecord = 5
	rdr.TrimLeadingSpace = true

	records, err := rdr.ReadAll()
	if err != nil {
		return err
	}

	p.List = nil
	for _, record := range records {
		p.List = append(
			p.List,
			&Promise{
				Class:    record[0],
				Handler:  record[1],
				Promiser: record[2],
				Promisee: record[3],
				Outcome:  record[4],
			})
	}

	return nil
}

// Watch starts a go routine for the report and watches its source file for
// changes. It implements the Report interface.
func (p *Promises) Watch(c chan *irc.Message) {
	go func() {
		for {
			// Calc new hash sum.
			newSum, err := calcHashSum(p.Filename)
			if err != nil {
				log.Println("Error:", err)
			}
			// Check if new and old sum differ.
			if !bytes.Equal(newSum, p.Checksum) {
				p.Checksum = newSum // Set new sum.
				p.Read()            // Reread the report.
				log.Println("Checksum changed for", p.Filename)
				c <- &irc.Message{ // Send message to irc server.
					Command:  irc.PRIVMSG,
					Params:   []string{*config.Channels},
					Trailing: "Checksum changed for " + p.Filename,
				}
			}

			time.Sleep(*config.WatchInterval * time.Second)
		}
	}()
}

// ToString returns the Promise struct as a string.
func (p *Promise) ToString() string {
	return "class: " + p.Class + ", handler: " + p.Handler + ", promiser: " + p.Promiser + ", promisee: " + p.Promisee + ", outcome: " + p.Outcome
}
