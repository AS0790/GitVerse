// Package cli contains the GitVerse terminal workflow.
package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/term"

	"github.com/TheVicky1/GitVerse/internal/counters"
	"github.com/TheVicky1/GitVerse/internal/official"
	"github.com/TheVicky1/GitVerse/internal/render"
	"github.com/TheVicky1/GitVerse/internal/report"
)

const usage = `Know the work behind your GitHub profile.

Usage:
  gitverse [username] [flags]

With no username, GitVerse checks the account signed in to GitHub CLI.

Flags:
  --all        also list the archive, every highlight and each badge's notes
  --json       print the report as JSON
  --no-color   plain text output (also honours NO_COLOR)
  --version    print the version

Docs: https://gitverse.vercel.app/
`

// Run executes the GitVerse terminal command.
func Run(args []string, version string) error {
	fs := flag.NewFlagSet("gitverse", flag.ContinueOnError)
	fs.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	all := fs.Bool("all", false, "")
	asJSON := fs.Bool("json", false, "")
	noColor := fs.Bool("no-color", false, "")
	showVersion := fs.Bool("version", false, "")
	if err := fs.Parse(Reorder(args)); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}
		return err
	}
	if *showVersion {
		fmt.Println("gitverse", version)
		return nil
	}
	if fs.NArg() > 1 {
		fs.Usage()
		return errors.New("give at most one username")
	}

	counts := counters.New()
	user := fs.Arg(0)
	if user != "" && !official.IsLogin(user) {
		return fmt.Errorf("%q is not a GitHub username", user)
	}
	if user == "" {
		login, err := currentLogin()
		if err != nil {
			return errors.New("no username given and not signed in to GitHub CLI: run `gitverse <username>` or `gh auth login`")
		}
		user = login
	}

	r, err := report.Build(user, official.NewClient(), counts)
	if errors.Is(err, official.ErrUserNotFound) {
		return fmt.Errorf("user %q not found", user)
	}
	if err != nil {
		return err
	}
	if *asJSON {
		return render.JSON(os.Stdout, r)
	}
	t := term.FromEnv()
	color := !*noColor && os.Getenv("NO_COLOR") == "" && t.IsTerminalOutput() && t.IsColorEnabled()
	render.Text(os.Stdout, r, render.Options{Color: color, All: *all})
	if t.IsTerminalOutput() {
		StarOnce(os.Stdout, color, StarMarker())
	}
	return nil
}

// StarOnce prints the GitHub star request at most once per local profile.
func StarOnce(w io.Writer, color bool, marker string) {
	if marker == "" {
		return
	}
	if err := os.MkdirAll(filepath.Dir(marker), 0o755); err != nil {
		return
	}
	f, err := os.OpenFile(marker, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return
	}
	_ = f.Close()
	fmt.Fprintln(w, render.StarLine(color))
	fmt.Fprintln(w)
}

// StarMarker returns the location of the one-time GitVerse star marker.
func StarMarker() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		return ""
	}
	return filepath.Join(dir, "gitverse", "star-shown")
}

func currentLogin() (string, error) {
	rest, err := api.DefaultRESTClient()
	if err != nil {
		return "", err
	}
	var me struct{ Login string }
	if err := rest.Get("user", &me); err != nil {
		return "", err
	}
	return me.Login, nil
}

// Reorder puts flags before the username, so `gitverse torvalds --json` works.
func Reorder(args []string) []string {
	var flags, rest []string
	for _, a := range args {
		if len(a) > 1 && a[0] == '-' {
			flags = append(flags, a)
		} else {
			rest = append(rest, a)
		}
	}
	return append(flags, rest...)
}
