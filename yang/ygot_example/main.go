package main

import (
	"fmt"
	"log"
    "example"
	"github.com/openconfig/ygot/ygot"
	
)

func main() {
	// Create an instance of the config struct
	cfg := &example.ExampleConfig_Config{}

	// Populate the struct
	cfg.Name = ygot.String("John Doe")
	cfg.Age = ygot.Uint8(30)
	cfg.Email = ygot.String("john.doe@example.com")

	// Validate the struct
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Validation error: %v", err)
	}

	// Marshal to JSON
	jsonData, err := ygot.EmitJSON(cfg, nil)
	if err != nil {
		log.Fatalf("Failed to marshal to JSON: %v", err)
	}

	fmt.Printf("JSON Output: %s\n", jsonData)
}
