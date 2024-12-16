package main

import (
	"fmt"
	"log"

	"github.com/martinpare1208/gator/internal/config"
)


func main() {
	data, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config: %+v\n", data)

	data.SetUser("martin")

	data, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Reading config again: %+v\n", data)

}