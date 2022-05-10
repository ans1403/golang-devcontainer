package main

import (
	"fmt"
	"golang-devcontainer/src/service"
	"log"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("../../config/config.ini")
	if err != nil {
		log.Fatal(err)
	}

	maxPrintCount := cfg.Section("Common").Key("max_print_count").MustInt()

	calculateService := service.NewCalculate()
	result := calculateService.Service(maxPrintCount)

	fmt.Println(result)
}
