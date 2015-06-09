package report

import (
	"encoding/csv"
	"os"

	"github.com/smotti/tad/config"
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
	}
)

// init registers the report at application startup.
func init() {
	var r Report
	r = &Promises{
		Filename: *config.Promises,
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

// ToString returns the Promise struct as a string.
func (p *Promise) ToString() string {
	return "class: " + p.Class + ", handler: " + p.Handler + ", promiser: " + p.Promiser + ", promisee: " + p.Promisee + ", outcome: " + p.Outcome
}
