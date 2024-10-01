package flags

import (
    "flag"
    "fmt"
)

// ParseFlags parses command-line flags for the application.
func ParseFlags(args []string) (outFile, format string, err error) {
    // Set up flag parsing using FlagSet, keeping the global space clean
    flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)
    outFilePtr := flagSet.String("out", "output.txt", "Output file path")
    formatPtr := flagSet.String("fmt", "text", "Output format (text, markdown, json)")

    // Parse the flags
    err = flagSet.Parse(args[1:])
    if err != nil {
        return "", "", fmt.Errorf("failed to parse flags: %v", err)
    }

    return *outFilePtr, *formatPtr, nil
}
