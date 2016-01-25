package main

import (
	"encoding/csv"
	"io"
	"os"
	"sort"

	"github.com/deckarep/golang-set"
)

func parseCsv(file string) ([]mapset.Set, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	csvr.FieldsPerRecord = -1

	data := []mapset.Set{}
	for {
		row, err := csvr.Read()

		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return data, err
		}

		t := mapset.NewThreadUnsafeSet()
		sort.Strings(row)
		for _, r := range row {
			t.Add(r)
		}

		data = append(data, t)
	}
}
