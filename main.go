package main

import (
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

		// Process output in a goroutine
		done := make(chan error)
		go func() {
			buf := make([]byte, 1024)
			re := regexp.MustCompile(`[0-9]+\ doctest,\ [0-9]+\ test,\ [0-9]+\ failures`)
			for {
				n, err := stdout.Read(buf)
				if n > 0 {
					line := string(buf[:n])
					fmt.Print(line)

					// Strip ANSI color codes
					cleanLine := regexp.MustCompile(`\x1b\[[0-9;]*m`).ReplaceAllString(line, "")
					if re.MatchString(cleanLine) {
						// Update terminal title
						fmt.Printf("\033]1;%s\007", cleanLine)
					}
				}
				if err != nil {
					done <- err
					return
				}
			}
		}()

		// Wait for the command to finish
		err = c.Wait()
		if err != nil {
			return err
		}

		// Wait for the output processing to finish
		return <-done
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
