// Command gitverse prints GitHub achievement progress in the terminal.
package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/TheVicky1/GitVerse/internal/cli"
)

var version = "dev"

func main() {
	if err := cli.Run(os.Args[1:], buildVersion()); err != nil {
		fmt.Fprintln(os.Stderr, "gitverse:", err)
		os.Exit(1)
	}
}

func buildVersion() string {
	if version != "dev" {
		return version
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return version
}
