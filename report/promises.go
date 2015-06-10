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
		Filename   string
		Checksum   []byte
		All        []*Promise
		Kept       []*Promise
		Repaired   []*Promise
		Failed     []*Promise
		Unknown    []*Promise
		OldModTime time.Time
	}
)

// init registers the report at application startup.
func init() {
	h, err := calcHashSum(*config.Promises)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	fi, err := os.Stat(*config.Promises)
	if err != nil {
		log.Println("Error:", err)
	}

	var r Report
	r = &Promises{
		Filename:   *config.Promises,
		Checksum:   h,
		OldModTime: fi.ModTime(),
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

	p.All = nil
	p.Unknown = nil
	p.Kept = nil
	p.Repaired = nil
	p.Failed = nil
	for _, record := range records {
		promise := &Promise{
			Class:    record[0],
			Handler:  record[1],
			Promiser: record[2],
			Promisee: record[3],
			Outcome:  record[4],
		}

		p.All = append(p.All, promise)

		switch record[4] {
		default:
			p.Unknown = append(p.Unknown, promise)
		case "kept":
			p.Kept = append(p.Kept, promise)
		case "repaired":
			p.Repaired = append(p.Repaired, promise)
		case "failed":
			p.Failed = append(p.Failed, promise)
		}
	}

	return nil
}

// Watch starts a go routine for the report and watches its source file for
// changes. It implements the Report interface.
func (p *Promises) Watch(c chan *irc.Message) {
	go func() {
		for {
			// Continuously watch file but only check hash sum if ModTime
			// changed.
			fi, err := os.Stat(p.Filename)
			if os.IsNotExist(err) {
				continue
			}

			if fi.ModTime().After(p.OldModTime) {
				// Calc new hash sum.
				newSum, err := calcHashSum(p.Filename)
				if err != nil {
					log.Println("Error:", err)
				}
				// Check if new and old sum differ.
				if !bytes.Equal(newSum, p.Checksum) {
					if err := p.Read(); err != nil {
						log.Println("Error:", err)
					} else {
						p.Checksum = newSum
						log.Println("Checksum changed for", p.Filename)
						c <- &irc.Message{ // Send message to irc server.
							Command:  irc.PRIVMSG,
							Params:   []string{*config.Channels},
							Trailing: "Checksum changed for " + p.Filename,
						}
					}
				}
			}

			time.Sleep(*config.WatchInterval * time.Second)
		}
	}()
}

// Notify starts a go routine to send notifications about repaired and failed
// promises to irc every config.notifyInterval.
func (p *Promises) Notify(c chan *irc.Message) {
}

// ToString returns the Promise struct as a string.
func (p *Promise) ToString() string {
	return "class: " + p.Class + ", handler: " + p.Handler + ", promiser: " + p.Promiser + ", promisee: " + p.Promisee + ", outcome: " + p.Outcome
}
