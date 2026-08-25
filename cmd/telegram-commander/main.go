// Command telegram-commander is a Telegram bot that runs host commands from a YAML menu.
package main

import (
	"fmt"
	"os"

	"github.com/azolfagharj/telegram-commander/internal/cli"
)

// Change these when you release a new version.
const version = "1.0.4"
const releaseDate = "2026-08-25"

func main() {
	if err := cli.NewRoot(version, releaseDate).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
