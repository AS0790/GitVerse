# GitVerse Developer Guide

This document describes the local development workflow, project architecture, tooling, and verification gates.

---

## 🏗️ Architecture Layout

| Directory | Purpose |
| :--- | :--- |
| `cmd/gitverse` | CLI binary entrypoint. Parses flags, handles gh auth, prints reports. |
| `cmd/gendocs` | Static site generation and documentation maintenance tool. |
| `internal/catalog` | Canonical definitions of all 15 achievements, tiers, and criteria. |
| `internal/official` | Scrapers and parsers for GitHub profile pages and achievement detail HTML. |
| `internal/counters` | API telemetry collectors querying merged PRs, stars, and discussions. |
| `internal/progress` | Calibration mathematics, tier floors, and progress calculation. |
| `internal/render` | Terminal text formatting with ANSI color rendering. |
| `internal/webui` | Web templates, static CSS/JS, and layout components. |
| `site/` | Production static site directory deployed to Vercel. |
| `scripts/` | Local preview server (`serve-preview.cjs`), benchmarks, and utilities. |

---

## 🛠️ Local Development Commands

### 1. Web Development Preview
Run the zero-dependency local Node server with instant asset reloading:

```bash
node scripts/serve-preview.cjs
```
Navigate to [http://localhost:5173](http://localhost:5173).

### 2. Running Automated Tests
```bash
# Run client-side JavaScript behavior tests
node --test internal/webui/jstest/*.test.mjs

# Run Go unit tests across all packages
go test ./...

# Run live canary tests against GitHub schemas
go test -tags canary ./internal/official -run Canary -v
```

### 3. Compiling the CLI
```bash
# Build the binary into your working directory
go build -o gitverse ./cmd/gitverse

# Run directly against a user profile
./gitverse torvalds
```

### 4. Regenerating Documentation & Static Assets
When modifying the achievement catalog (`internal/catalog/catalog.go`), regenerate the derived pages:

```bash
go run ./cmd/gendocs
```
This updates `README.md` catalog tables, `site/achievements/`, `site/achievements.json`, and `site/llms.txt`.

