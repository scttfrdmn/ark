package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version   = "dev"
	commitSHA = "unknown"
	buildDate = "unknown"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "ark",
	Short: "Ark - AWS Research Kit for academic institutions",
	Long: `Ark provides integrated AWS training and security tooling for research institutions.

The training-as-tool approach embeds security education directly into AWS workflows,
ensuring researchers can use cloud resources safely and compliantly from day one.`,
	Version: fmt.Sprintf("%s (commit: %s, built: %s)", version, commitSHA, buildDate),
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().Bool("json", false, "Output in JSON format")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
}
