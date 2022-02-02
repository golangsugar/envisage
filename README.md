# Envisage
### A lightweight package that makes easier and safer to deal with environment variables.

#### Example
```lang=golang
package main

import (
	"fmt"
	"github.com/golangsugar/envisage"
)

func main() {
    /* Examples */
    
    // SetString sets the value of the environment variable named by the key.
    if err:=envisage.SetString("DB_CONNECTION_STRING","databasedriver://user:password@serverhost/db?options");err!=nil{
        return err
    }

    // SetInt sets the value of the environment variable named by the key.
    if err:=envisage.SetInt("WEBSERVICE_PORT","154");err!=nil{
        return err
    }

    // SetF64 sets the value of the environment variable named by the key.
    if err:=envisage.SetF64("PRODUCT_PRICE","14.1544");err!=nil{
        return err
    }

    // SetBool sets the value of the environment variable named by the key.
    if err:=envisage.SetBool("BOOLEAN_VALUE","true");err!=nil{
        return err
    }

    if err:=envisage.SetIntS("ARRAY_OF_INTEGERS","45,8,22,5,4,4,666,4,9");err!=nil{
        return err
    }    
    
    var (
        key = "DB_CONNECTION_STRING"
        defaultValue = "-" // defaultValue returned if the value is not present/set in the environment
        twoWay = true // twoWay updates the environment with the defaultValue, in case of the environment variable is not present/set
        mandatory = true // mandatory forces and error in case of the variable is absent in the environment
    )

    // Check Test environment variables according given directives
    if err:=envisage.Check(key, defaultValue, twoWay) ;err!=nil{
        log.Fatal(err)
    }
    
    fmt.Println(envisage.IsThere("DB_CONNECTION_STRING"))
    // Print true
    
    fmt.Println(envisage.String("DB_CONNECTION_STRING"))
    // Print "databasedriver://user:password@serverhost/db?options"
    
    fmt.Println(envisage.Int("WEBSERVICE_PORT"))
    // Print 154
    
    fmt.Println(envisage.I64("WEBSERVICE_PORT"))
    // Print 154
    
    fmt.Println(envisage.Bool("BOOLEAN_VALUE"))
    // Print true

    // F64 returns the env var value as float64
    fmt.Println(envisage.F64("PRODUCT_PRICE"))
    // Print 14.1544
    
    // IntS returns the env var value as []int
    fmt.Println(envisage.IntS("ARRAY_OF_INTEGERS"))
    // Print []int{45,8,22,5,4,4,666,4,9}   
}
```

