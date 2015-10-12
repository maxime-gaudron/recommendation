# Recommendation engine

[![Build Status](https://travis-ci.org/maxime-gaudron/recommendation.svg?branch=master)](https://travis-ci.org/maxime-gaudron/recommendation)

## Utilisation

```
NAME:
   recommendation - generate association rules using a list of transactions

USAGE:
   ./recommendation [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --inputFile "data.csv"	transaction data, comma separated
   --outputFile "output.csv"	output file path containing association rules
   --minSupport "0.002"		Minimum support to consider an item set frequent
   --minConfidence "1"		Minimum confidence to consider an association rule
   --help, -h			show help
   --version, -v		print the version
```

**Input file format:**
```csv
A,C,T,W
C,D,W
A,C,T,W
```
Each line represent a transaction.

**Output file format:**
```csv
antecedent,consequent,support,confidence,lift
T;A,C;W,0.500000,1.000000,0.833333
A;W;T,C,0.500000,1.000000,1.000000
```
Sets are semi-colon separated.

## Internals

**arules.go:** frequent item sets to association rules learning  
**apriori.go:** apriori algorithm, frequent item sets mining  
**csv.go:** CSV Input/Output

It relies on https://github.com/deckarep/golang-set

## Authors

Only me so far :)

Bug, feature requests, Submit a patch ?

Please ! Use Github's tools or contact me by email

## Contact

<Maxime Gaudron> maxime.gaudron@rocket-internet.de

## History

This project started as a learning exercice around:
 - Apriori algorithm to generate frequent itemsets from transactions
 - Association rules generation from frequent itemsets
 - Writing an API exposing association rules
 - Learning Golang


 My employer (Rocket-internet) let me work on this project during my working time as 20% project.
