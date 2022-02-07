# Envisage
### A lightweight package that makes easier and safer to deal with environment variables.
[![Known Vulnerabilities](https://snyk.io/test/github/golangsugar/envisage/badge.svg)](https://snyk.io/test/github/golangsugar/envisage)

#### Example
Try it on [GoPlay](https://goplay.tools/snippet/XLaebfrrfoH)

--- 

```go
package main

import (
	"fmt"
	"github.com/golangsugar/envisage"
	"log"
	"os"
)

func main() {
	/* Examples */

	const (
		updateEnvironment = true
		skipIfAlreadyDefined = true
		errorIfFileDoesntExist = true
    )

	// LoadFromFile loads environment variables values from a given text file in to a map[string]string.
	cfgmap,cmerr:=envisage.LoadFromFile(".local.env", updateEnvironment, skipIfAlreadyDefined, errorIfFileDoesntExist)
	if cmerr!=nil {
		log.Println(cmerr)
	}
	
	fmt.Println(cfgmap)

	// SetString sets the value of the environment variable named by the key.
	if err := envisage.SetString("DB_CONNECTION_STRING", "databasedriver://user:password@serverhost/db?options"); err != nil {
		log.Fatal(err)
	}

	// SetInt sets the value of the environment variable named by the key.
	if err := envisage.SetInt("WEBSERVICE_PORT", 154); err != nil {
		log.Fatal(err)
	}

	// SetF64 sets the value of the environment variable named by the key.
	if err := envisage.SetF64("PRODUCT_PRICE", 14.1544); err != nil {
		log.Fatal(err)
	}

	// SetBool sets the value of the environment variable named by the key.
	if err := envisage.SetBool("BOOLEAN_VALUE", true); err != nil {
		log.Fatal(err)
	}

	if err := os.Setenv("ARRAY_OF_INTEGERS", "45,8,22,5,4,4,666,4,9"); err != nil {
		log.Fatal(err)
	}

	const (
		key          = "DB_CONNECTION_STRING"
		defaultValue = "-"   // defaultValue returned if the value is not present/set in the environment
		twoWay       = true  // twoWay updates the environment with the defaultValue, in case of the environment variable is not present/set
		canBeEmpty   = false // Empty is an acceptable value if true
	)

	// Check Test environment variables according given directives
	if err := envisage.Check(key, defaultValue, twoWay, canBeEmpty); err != nil {
		log.Fatal(err)
	}

	fmt.Println(envisage.IsThere("DB_CONNECTION_STRING"))
	// Print true

	fmt.Println(envisage.String("DB_CONNECTION_STRING", defaultValue))
	// Print "databasedriver://user:password@serverhost/db?options"

	fmt.Println(envisage.Int("WEBSERVICE_PORT", 0))
	// Print 154

	fmt.Println(envisage.I64("WEBSERVICE_PORT", 0))
	// Print 154

	fmt.Println(envisage.Bool("BOOLEAN_VALUE", false))
	// Print true

	// F64 returns the env var value as float64
	fmt.Println(envisage.F64("PRODUCT_PRICE", false, 0))
	// Print 14.1544

	// IntS returns the env var value as []int
	fmt.Println(envisage.IntS("ARRAY_OF_INTEGERS", ",", nil))
	// Print []int{45,8,22,5,4,4,666,4,9}
}
```

