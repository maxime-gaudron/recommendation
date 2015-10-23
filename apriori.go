package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/deckarep/golang-set"
)

type candidate struct {
	Items     mapset.Set
	Count     int
	UpdatedAt int
}

// Apriori algorithm, mining frequent item sets out of an array of transactions
// minSupport: set the lower limit under which the items' sets are considerated not frequent
func apriori(transactions []mapset.Set, minSupport float64) (large []candidate) {
	// Initial empty frontier set to start the first iteration
	frontier := []candidate{
		candidate{
			Items: mapset.NewThreadUnsafeSet(),
			Count: len(transactions),
		},
	}

	run := 0
	for len(frontier) > 0 {
		candidates := map[string]candidate{}

		run++
		fmt.Printf("Pass #%d with %d frontier sets\n", run, len(frontier))

		for k, t := range transactions {
			for _, f := range frontier {
				if t.IsSuperset(f.Items) {
					for _, c := range extend(f, t, candidates) {
						candidates = upsert(candidates, c, k)
					}
				}
			}
		}

		frontier = []candidate{}
		for _, c := range candidates {
			if float64(c.Count)/float64(len(transactions)) >= minSupport {
				large = append(large, c)

				// TODO: As extend should be recursive this should be more complex than just checking for the minSupport.
				frontier = append(frontier, c)
			}
		}
	}

	return
}

// Update the count of the set to insert if found and not updated during the current iteration or insert it
func upsert(large map[string]candidate, c mapset.Set, position int) map[string]candidate {
	keyArr := []string{}
	for _, s := range c.ToSlice() {
		if str, ok := s.(string); ok {
			keyArr = append(keyArr, str)
		}
	}
	sort.Strings(keyArr)
	key := strings.Join(keyArr, "")

	data, ok := large[key]
	if !ok {
		// Not found ? insert with count = 1
		large[key] = candidate{Items: c, Count: 1, UpdatedAt: position}
	} else {
		if data.UpdatedAt < position {
			element := large[key]
			element.Count++
			element.UpdatedAt = position
			large[key] = element
		}
	}

	return large
}

// Extend the candidate set with items from the transaction
// TODO: Normally uses statistical independence assumption to extend further
func extend(c candidate, transaction mapset.Set, candidates map[string]candidate) (extended []mapset.Set) {
	for _, i := range transaction.Difference(c.Items).ToSlice() {
		extendedItems := mapset.NewThreadUnsafeSetFromSlice(c.Items.ToSlice())
		extendedItems.Add(i)
		extended = append(extended, extendedItems)
	}

	return
}
