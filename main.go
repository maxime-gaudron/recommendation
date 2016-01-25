package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "recommendation"
	app.Usage = "generate association rules using a list of transactions"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "inputFile",
			Value: "data.csv",
			Usage: "transaction data, comma separated",
		},
		cli.StringFlag{
			Name:  "outputFile",
			Value: "output.csv",
			Usage: "output file path containing association rules",
		},
		cli.Float64Flag{
			Name:  "minSupport",
			Value: 0.002,
			Usage: "Minimum support to consider an item set frequent",
		},
		cli.Float64Flag{
			Name:  "minConfidence",
			Value: 1.,
			Usage: "Minimum confidence to consider an association rule",
		},
	}
	app.Action = func(c *cli.Context) {
		transactions, err := parseCsv(c.String("inputFile"))
		if err != nil {
			fmt.Printf("%+v\n", err)
		}

		fmt.Printf("Loaded %d transactions\n", len(transactions))

		frequentItemSets := apriori(transactions, c.Float64("minSupport"))

		fmt.Printf("Found %d frequent itemsets\n", len(frequentItemSets))
		fmt.Printf("Generating association rules\n")
	}

	app.Run(os.Args)
}
