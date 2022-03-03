package main

import (
	"fmt"

	pql "github.com/abdulhalim-cu/gopsql/gopsql"
)

func main() {
	pql.Hostname = "localhost"
	fmt.Println(pql.Port)
	fmt.Println(pql.Hostname)
}
