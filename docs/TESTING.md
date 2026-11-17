# GitVerse Testing Strategy & Quality Assurance

GitVerse maintains an automated, multi-tiered test suite ensuring accuracy, reliability, and regression prevention across both CLI and Web engines.

---

## 1. Test Architecture

The test suite is organized into three distinct verification layers:

```text
+-----------------------------------------------------------------+
|                    Client-Side Behavior Tests                   |
|  - Node.js native test runner (node --test)                     |
|  - Validates app.js, theme.js, DOM manipulation, GA tracking   |
+-----------------------------------------------------------------+
                                |
+-----------------------------------------------------------------+
|                    Go Unit & Parser Test Matrix                 |
|  - Tests official HTML scrapers against real profile fixtures   |
|  - Validates calibration arithmetic & provenance tagging        |
+-----------------------------------------------------------------+
                                |
+-----------------------------------------------------------------+
|                    Weekly Live Canary Jobs                      |
|  - Scheduled GitHub Action running live profile queries         |
|  - Detects upstream changes to GitHub's markup or badge schemas |
+-----------------------------------------------------------------+
```

---

## 2. Running Tests Locally

### Client-Side JavaScript Tests
Run the zero-dependency behavioral test suite:

```bash
node --test internal/webui/jstest/*.test.mjs
```
*Current status: 26 passed, 0 failed.*

### Go Unit Tests
Run standard unit tests across all packages:

```bash
go test ./...
```

### Static Analysis & Verification
```bash
# Verify generated site matches catalog definitions
go run ./cmd/gendocs -check

# Code formatting and vet
gofmt -l .
go vet ./...
```

---

## 3. Canary Testing

Because GitHub's profile pages are server-rendered HTML rather than a versioned API, upstream layout changes could break detail scrapers.

GitVerse addresses this through automated **Canary Tests**:
* A weekly GitHub Actions workflow runs `go test -tags canary ./internal/official -run Canary -v` against designated public fixture profiles.
* If GitHub alters DOM class names or data attributes, the canary test fails immediately and opens an automated issue for maintainers.

