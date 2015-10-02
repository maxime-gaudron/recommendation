package main

import "fmt"

func main() {
	transactions, err := parseCsv("./data.csv")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	frequentItemSets := apriori(transactions, 0.5)

	rules := generateRules(frequentItemSets, len(transactions), 1.)

	toCsv(rules)
}
