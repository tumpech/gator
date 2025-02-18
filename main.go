package main

import (
	"fmt"
	"log"

	"github.com/tumpech/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config:%+v\n", cfg)
	err = cfg.SetUser("lane")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}
	fmt.Printf("Changed user to lane: %+v\n", cfg)

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
