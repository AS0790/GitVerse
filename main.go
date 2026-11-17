// GitVerse is a terminal companion for GitHub profile achievements.
package main

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"github.com/TheVicky1/GitVerse/internal/cli"
)

var version = "dev"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "gitverse:", err)
		os.Exit(1)
	}
}

func run(args []string) error { return cli.Run(args, buildVersion()) }

func starOnce(w io.Writer, color bool, marker string) { cli.StarOnce(w, color, marker) }

func starMarker() string { return cli.StarMarker() }

func reorder(args []string) []string { return cli.Reorder(args) }

func buildVersion() string {
	if version != "dev" {
		return version
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return version
}
