package main

import (
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

	fSets := apriori(transactions, 1)
	if len(fSets) != 1 || fSets[0].Count != 6 {
		t.Fail()
	}
	expected := []interface{}{"C"}
	if !fSets[0].Items.Equal(mapset.NewThreadUnsafeSetFromSlice(expected)) {
		t.Fail()
	}

	fSets = apriori(transactions, 0.83)
	if len(fSets) != 3 {
		t.Fail()
	}

	fSets = apriori(transactions, 0.66)
	if len(fSets) != 11 {
		t.Fail()
	}

	fSets = apriori(transactions, 0.49)
	if len(fSets) != 19 {
		t.Fail()
	}
}
