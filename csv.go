package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

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
		for _, r := range row {
			t.Add(r)
		}

		data = append(data, t)
	}
}

func toCsv(rules []rule) {
	f, err := os.Create("./output.csv")
	if err != nil {
		log.Fatalln("error opening file:", err)
	}

	w := csv.NewWriter(f)

	if err := w.Write([]string{"antecedent", "consequent", "support", "confidence", "lift"}); err != nil {
		log.Fatalln("error writing headers to csv:", err)
	}

	for _, rule := range rules {
		if err := w.Write(getRecord(rule)); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func getRecord(rule rule) []string {
	antecedents := []string{}
	for _, item := range rule.Antecedents.ToSlice() {
		if str, ok := item.(string); ok {
			antecedents = append(antecedents, str)
		}
	}

	consequents := []string{}
	for _, item := range rule.Consequents.ToSlice() {
		if str, ok := item.(string); ok {
			consequents = append(consequents, str)
		}
	}

	return []string{
		strings.Join(antecedents, ";"),
		strings.Join(consequents, ";"),
		strconv.FormatFloat(rule.Support, 'f', 6, 64),
		strconv.FormatFloat(rule.Confidence, 'f', 6, 64),
		strconv.FormatFloat(rule.Lift, 'f', 6, 64),
	}
}
