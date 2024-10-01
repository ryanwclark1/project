package env

import (
	"os"
)

// Getenv reads an environment variable.
func Getenv(key string) string {
    return os.Getenv(key)
}
