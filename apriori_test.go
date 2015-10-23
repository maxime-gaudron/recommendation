package main

import (
	"fmt"
	"testing"

	"github.com/deckarep/golang-set"
)

// TODO: Write func to compare the content with fixtures, not only the number of frequent item sets
func TestApriori(t *testing.T) {
	transactions := []mapset.Set{
		mapset.NewThreadUnsafeSetFromSlice([]interface{}{"A", "C", "T", "W"}),
		mapset.NewThreadUnsafeSetFromSlice([]interface{}{"C", "D", "W"}),
		mapset.NewThreadUnsafeSetFromSlice([]interface{}{"A", "C", "T", "W"}),
		mapset.NewThreadUnsafeSetFromSlice([]interface{}{"A", "C", "D", "W"}),
		mapset.NewThreadUnsafeSetFromSlice([]interface{}{"A", "C", "D", "T", "W"}),
		mapset.NewThreadUnsafeSetFromSlice([]interface{}{"C", "D", "T"}),
	}

	fmt.Printf("Test apriori with minSupport=1\n")
	fSets := apriori(transactions, 1)
	if len(fSets) != 1 || fSets[0].Count != 6 {
		fmt.Printf("%+v\n", fSets)
		t.Fail()
	}
	expected := []interface{}{"C"}
	if !fSets[0].Items.Equal(mapset.NewThreadUnsafeSetFromSlice(expected)) {
		fmt.Printf("%+v\n", fSets)

		t.Fail()
	}

	fmt.Printf("Test apriori with minSupport=0.83\n")
	fSets = apriori(transactions, 0.83)
	if len(fSets) != 3 {
		fmt.Printf("%+v\n", fSets)

		t.Fail()
	}

	fmt.Printf("Test apriori with minSupport=0.66\n")
	fSets = apriori(transactions, 0.66)
	if len(fSets) != 11 {
		fmt.Printf("%+v\n", fSets)

		t.Fail()
	}

	fmt.Printf("Test apriori with minSupport=0.49\n")
	fSets = apriori(transactions, 0.49)
	if len(fSets) != 19 {
		fmt.Printf("%+v\n", fSets)

		t.Fail()
	}
}
