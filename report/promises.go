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
	rdr.FieldsPerRecord = 4
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
				Outcome:  record[3],
			})
	}

	return nil
}
