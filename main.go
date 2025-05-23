package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var (
	watch bool
	trace bool
)

var rootCmd = &cobra.Command{
	Use:                "mtx",
	Short:              "Mix test wrapper with watch and trace options",
	DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Parse our known flags manually
		var remainingArgs []string
		for i := 0; i < len(args); i++ {
			switch args[i] {
			case "-w", "--watch":
				watch = true
			case "-t", "--trace":
				trace = true
			case "-h", "--help":
				cmd.Help()
				return nil
			default:
				remainingArgs = append(remainingArgs, args[i])
			}
		}

		// If both trace and watch are set, exit and explain
		if trace && watch {
			fmt.Println("You are a silly goose! You can't use --trace and --watch together. `mix test.watch` does not support the --trace flag. Please choose one or the other.")
			os.Exit(1)
		}

		// Build the command
		command := "mix test"
		if watch {
			command = "mix test.watch"
		}
		if trace {
			command += " --trace"
		}
		if len(remainingArgs) > 0 {
			command += " " + strings.Join(remainingArgs, " ")
		}

		fmt.Printf("Running: %s\n", command)

		// Create the command
		c := exec.Command("bash", "-c", command)
		c.Stderr = os.Stderr

		// Create a pipe to read the output
		stdout, err := c.StdoutPipe()
		if err != nil {
			return fmt.Errorf("error creating stdout pipe: %v", err)
		}

		// Start the command
		if err := c.Start(); err != nil {
			return fmt.Errorf("error starting command: %v", err)
		}

		// Process output line-by-line in the main goroutine
		scanner := bufio.NewScanner(stdout)
		re := regexp.MustCompile(`(\d+)\s+doctest[s]?,\s+(\d+)\s+test[s]?,\s+(\d+)\s+failures`)
		ansi := regexp.MustCompile(`\x1b\[[0-9;]*m`)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
			cleanLine := ansi.ReplaceAllString(line, "")
			if re.MatchString(cleanLine) && watch {
				// Update both tab and window title for compatibility
				fmt.Printf("\033]1;%s\007", cleanLine) // tab title
				fmt.Printf("\033]0;%s\007", cleanLine) // window title
			}
		}
		if err := scanner.Err(); err != nil {
			return err
		}

		// Wait for the command to finish
		return c.Wait()
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&watch, "watch", "w", false, "Run tests in watch mode")
	rootCmd.Flags().BoolVarP(&trace, "trace", "t", false, "Run tests with trace")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
