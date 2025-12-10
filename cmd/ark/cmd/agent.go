package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(agentCmd)
	agentCmd.AddCommand(agentStartCmd)
	agentCmd.AddCommand(agentStopCmd)
	agentCmd.AddCommand(agentStatusCmd)
}

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Manage the Ark agent",
	Long: `Manage the Ark agent service that runs locally and brokers AWS credentials.

The agent runs on localhost:8737 and must be running for Ark CLI commands to work.`,
}

var agentStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Ark agent",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if agent is already running
		if isAgentRunning() {
			fmt.Println("Agent is already running")
			return
		}

		// TODO: Implement proper agent start with process management
		fmt.Println("Starting Ark agent...")
		fmt.Println("(Full implementation coming soon)")
		fmt.Println("")
		fmt.Println("For now, run the agent manually:")
		fmt.Println("  ark-agent")
	},
}

var agentStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the Ark agent",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement proper agent stop with process management
		fmt.Println("Stopping Ark agent...")
		fmt.Println("(Not yet implemented)")
	},
}

var agentStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check agent status",
	Run: func(cmd *cobra.Command, args []string) {
		if isAgentRunning() {
			fmt.Println("✓ Agent is running")

			// Try to get version
			version, err := getAgentVersion()
			if err == nil {
				fmt.Printf("  Version: %s\n", version)
			}
			os.Exit(0)
		} else {
			fmt.Println("✗ Agent is not running")
			fmt.Println("")
			fmt.Println("Start the agent with: ark agent start")
			os.Exit(1)
		}
	},
}

// isAgentRunning checks if the agent is responding to health checks
func isAgentRunning() bool {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get("http://127.0.0.1:8737/api/system/health")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// getAgentVersion gets the agent version
func getAgentVersion() (string, error) {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get("http://127.0.0.1:8737/api/system/version")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// TODO: Parse JSON response
	return "dev", nil
}
