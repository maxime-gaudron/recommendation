package main

import (
	"fmt"

	"github.com/deckarep/golang-set"
)

type candidate struct {
	Items     []int
	Count     int
}

type hashTreeNode struct {
	Items     []int
	Count     int
}

// Apriori algorithm, mining frequent item sets out of an array of transactions
// minSupport: set the lower limit under which the items' sets are considerated not frequent
func apriori(transactions []mapset.Set, minSupport float64) ([]candidate) {
  large := map[int]map[string]int{
    1: apriori_first_pass(transactions),
  }

  for k := 2; len(large[k - 1]) > 0; k++ {
    candidates := apriori_gen(large[k - 1])
    fmt.Printf("%+v\n", large[k - 1])
  }

  return []candidate{}
}

func apriori_first_pass(transactions []mapset.Set) (large map[string]int) {
  minSupport := 100

  for _, t := range transactions {
    for _, i := range t.ToSlice() {
        large[1][i.(string)]++
    }
  }

  for k, v := range large {
    if v < minSupport {
      delete(large, k)
    }
  }
}

func apriori_gen(map[string]int) {
  apriori_join()
  apriori_prune()
}
