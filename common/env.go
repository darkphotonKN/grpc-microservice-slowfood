package common

import (
	"fmt"
	"os"
)

func EnvString(key string, fallback string) string {
	fmt.Printf("key %s getenv: %s\n", key, os.Getenv(key))
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}
