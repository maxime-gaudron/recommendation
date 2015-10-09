package main

import "github.com/deckarep/golang-set"

type rule struct {
	Antecedents mapset.Set
	Consequents mapset.Set
	Support     float64
	Confidence  float64
	Lift        float64
}

// Generate association rules from an array of frequent item sets
// minConfidence: set the lower limit under which the rules are not considerated relevant
func generateRules(frequentItemSets []candidate, dbsize int, minConfidence float64) (rules []rule) {
	for _, set := range frequentItemSets {
		for _, s := range set.Items.PowerSet().ToSlice() {
			if antecedentsSet, ok := s.(mapset.Set); ok && antecedentsSet.Cardinality() != 0 && antecedentsSet.Cardinality() != set.Items.Cardinality() {
				consequentsSet := set.Items.Difference(antecedentsSet)
				antecedents, consequents := findSets(frequentItemSets, antecedentsSet, consequentsSet)

				// Calculate the metrics
				ruleSupport := float64(set.Count) / float64(dbsize)
				antecedentSupport := float64(antecedents.Count) / float64(dbsize)
				consequentSupport := float64(consequents.Count) / float64(dbsize)
				rule := rule{
					Antecedents: antecedentsSet,
					Consequents: consequentsSet,
					Support:     ruleSupport,
					Confidence:  ruleSupport / antecedentSupport,
					Lift:        ruleSupport / antecedentSupport * consequentSupport,
				}

				if rule.Confidence >= minConfidence {
					rules = append(rules, rule)
				}
			}
		}
	}

	return
}

// Find the antecedent and consequent frequent itemSets
func findSets(frequentItemSets []candidate, antecedents, consequents mapset.Set) (a, c candidate) {
	for _, cs := range frequentItemSets {
		if cs.Items.Equal(consequents) {
			c = cs
		}
		if cs.Items.Equal(antecedents) {
			a = cs
		}
	}

	return
}
