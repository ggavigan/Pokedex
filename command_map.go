package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func mapf(config *Config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/?offset=40")
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	fmt.Println(body)

	if err := json.Unmarshal(body, &config); err != nil {
		return fmt.Errorf("error unmarshaling data: %w", err)
	}

	return nil
}
