package main

import (
	"debug-redactor/cmd"
	"fmt"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1) // Exits the program with a status code of 1 to indicate an error occurred.
	}
}
