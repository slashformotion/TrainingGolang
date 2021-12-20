package cert

import (
	"encoding/csv"
	"os"
)

func ParseCsv(filename string) ([]*Cert, error) {
	certs := make([]*Cert, 0)
	f, err := os.Open(filename)
	if err != nil {
		return certs, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return certs, err
	}
	for _, rec := range records {
		course := rec[0]
		n := rec[1]
		d := rec[2]
		c, err := New(course, n, d)
		if err != nil {
			return certs, err
		}
		certs = append(certs, c)
	}
	return certs, nil
}
