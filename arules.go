package main

import "github.com/deckarep/golang-set"

type rule struct {
	Antecedents mapset.Set
	Consequents mapset.Set
	Support     float64
	Confidence  float64
	Lift        float64
}

func generateRules(frequentItemSets []candidate, dbsize int, minConfidence float64) (rules []rule) {
	for _, set := range frequentItemSets {
		for _, s := range set.Items.PowerSet().ToSlice() {
			if antecedentsSet, ok := s.(mapset.Set); ok && antecedentsSet.Cardinality() != 0 && antecedentsSet.Cardinality() != set.Items.Cardinality() {
				consequentsSet := set.Items.Difference(antecedentsSet)
				antecedents, consequents := findSets(frequentItemSets, antecedentsSet, consequentsSet)

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
