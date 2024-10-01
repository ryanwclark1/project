package main

import (
    "context"
    "fmt"
    "os"
    "github.com/ryanwclark1/the-ui/internal/app"
    "github.com/ryanwclark1/the-ui/pkg/env"

)

func main() {
    ctx := context.Background()

    // Initialize the application
    err := app.Run(
        ctx,
        os.Args,                        // Command-line arguments
        env.Getenv,                     // Environment variable reader (from /pkg/env)
        os.Getwd,                       // Working directory reader
        os.Stdin, os.Stdout, os.Stderr, // I/O streams
    )
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        os.Exit(1)
    }
}
