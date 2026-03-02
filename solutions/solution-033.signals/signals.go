package solutions

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

// TODO: Implement RunGameServer
// Read README.md for the instructions

func RunGameServer(ready chan bool, workerCmd string) error {
	// Step 1: Use signal.NotifyContext with context.Background() to listen for
	// os.Interrupt (syscall.SIGINT) for graceful shutdown. Defer the stop function.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer stop()

	// Step 2: Create a standard signal channel and register it to receive
	// syscall.SIGHUP and syscall.SIGUSR1 via signal.Notify.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGUSR1)

	// Step 3: Use exec.Command (not CommandContext) to prepare the child process
	// so we can manually send SIGTERM instead of letting the context kill it brutally.
	cmd := exec.Command(workerCmd)

	// Step 4: Start the child process in the background. Return the error if it fails.
	if err := cmd.Start(); err != nil {
		return err
	}

	// Step 5: Send true to the ready channel to signal that the server is running.
	ready <- true

	// Step 6: Enter an infinite select loop to handle incoming signals and shutdown.
	for {
		select {
		case sig := <-sigCh:
			switch sig {
			// Step 6a: On SIGHUP, print a reload message and continue the loop.
			case syscall.SIGHUP:
				fmt.Println("Reloading game rules...")
			// Step 6b: On SIGUSR1, print a dump message and continue the loop.
			case syscall.SIGUSR1:
				fmt.Println("Dumping player coordinates...")
			}
		// Step 6c: When the graceful shutdown context is done (SIGINT received):
		//   - Print a shutdown message.
		//   - Send SIGTERM to the child process.
		//   - Wait for the child to finish and return the result (including any ExitError).
		case <-ctx.Done():
			fmt.Println("Shutting down worker...")
			cmd.Process.Signal(syscall.SIGTERM)
			return cmd.Wait()
		}
	}
}
