package app

import (
    "context"
    "fmt"
    "io"
    "os"
    "os/signal"
    "syscall"
    "github.com/ryanwclark1/the-ui/pkg/flags"
)

// Run is the main application logic. It takes in context, args, getenv, stdin, stdout, and stderr.
func Run(
    ctx context.Context,
    args []string,
    getenv func(string) string,
    getwd func() (string, error),
    stdin io.Reader,
    stdout, stderr io.Writer,
) error {
    // Set up a context that listens for interrupt signals (like Ctrl+C)
    ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
    defer cancel()

    // Process flags using a package from /pkg/flags
    outFile, format, err := flags.ParseFlags(args)
    if err != nil {
        fmt.Fprintf(stderr, "failed to parse flags: %v\n", err)
        return err
    }

    // Fetch an environment variable as a fallback or alternative to flags
    envFormat := getenv("MYAPP_FORMAT")
    if envFormat != "" {
        format = envFormat // Override flag if env variable is set
    }

    // Get the current working directory
    cwd, err := getwd()
    if err != nil {
        fmt.Fprintf(stderr, "failed to get working directory: %v\n", err)
        return err
    }
    fmt.Fprintf(stdout, "Current working directory: %s\n", cwd)

    // Output based on the flag/env settings
    fmt.Fprintf(stdout, "Output will be saved to %s in %s format\n", outFile, format)

    // Application logic
    fmt.Fprintln(stdout, "Application started...")

    // Block until the context is done (i.e., interrupted)
    <-ctx.Done()

    fmt.Fprintln(stdout, "Shutting down gracefully...")
    return nil
}
