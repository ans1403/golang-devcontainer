package main

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("../../../config/config.ini")
	if err != nil {
		log.Fatal(err)
	}

	maxPrintCount := cfg.Section("Common").Key("max_print_count").MustInt()

	for i := 1; i <= maxPrintCount; i++ {
		fmt.Println("Hello World!")
	}
}
