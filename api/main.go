package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/deckarep/golang-set"
	"github.com/gorilla/mux"
)

type byConfidence []rule

func (s byConfidence) Len() int {
	return len(s)
}
func (s byConfidence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byConfidence) Less(i, j int) bool {
	return s[i].Confidence < s[j].Confidence
}

var rules map[string][]rule

func main() {
	fmt.Printf("Parsing CSV\n")
	data, err := parseCsv("./output.csv")
	if err != nil {
		log.Fatal(err)
	}

	for k := range data {
		sort.Sort(sort.Reverse(byConfidence(data[k])))
	}

	rules = data

	fmt.Printf("Loaded %d transactions\n", len(rules))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/recommand/{set}", recommand)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func stringsToSet(strings []string) mapset.Set {
	set := mapset.NewThreadUnsafeSet()
	for _, s := range strings {
		set.Add(s)
	}

	return set
}

func recommand(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	query := strings.Split(mux.Vars(r)["set"], ",")
	sort.Strings(query)

	data, ok := rules[strings.Join(query, "")]
	if !ok {
		data = []rule{}
	}

	limit, err := strconv.ParseInt(r.FormValue("limit"), 10, 64)
	if err != nil {
		limit = 10
	}
	if int(limit) > len(data) {
		limit = int64(len(data))
	}

	js, err := json.Marshal(data[:limit])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	log.Printf("%s\t%s\t%s", time.Since(start), r.Method, r.RequestURI)
}

type rule struct {
	Antecedents []string
	Consequents []string
	Confidence  float64
	Support     float64
	Lift        float64
}

func (w rule) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"lift":            w.Lift,
		"support":         w.Support,
		"confidence":      w.Confidence,
		"recommendations": w.Consequents,
	})
}

func parseCsv(file string) (map[string][]rule, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	csvr.FieldsPerRecord = 5

	headers := true
	rules := map[string][]rule{}
	for {
		row, err := csvr.Read()

		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return rules, err
		}

		if headers {
			headers = false
			continue
		}

		antecedents := strings.Split(row[0], ";")
		consequents := strings.Split(row[1], ";")
		sort.Strings(antecedents)
		sort.Strings(consequents)

		r := rule{
			Lift:        mustParseFloat(row[2]),
			Support:     mustParseFloat(row[3]),
			Confidence:  mustParseFloat(row[4]),
			Antecedents: antecedents,
			Consequents: consequents,
		}

		key := strings.Join(antecedents, "")
		rules[key] = append(rules[key], r)
	}
}

func mustParseFloat(value string) float64 {
	i, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return i
}
