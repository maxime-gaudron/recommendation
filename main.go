package main

import "fmt"

func main() {
	transactions, err := parseCsv("./data.csv")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("Loaded %d transactions\n", len(transactions))

	frequentItemSets := apriori(transactions, 0.5)

	rules := generateRules(frequentItemSets, len(transactions), 1.)

	toCsv(rules)
}
