# GitVerse Troubleshooting Runbook

This runbook provides actionable solutions for common issues encountered when running, developing, or using GitVerse.

---

## 1. Web & Profile Checking Issues

### Issue: "GitHub has no user called [username]"
* **Cause**: The entered username does not exist or has been renamed/deleted on GitHub.
* **Resolution**: Verify the spelling of the GitHub handle. Ensure no leading `@` is included.

### Issue: Private repositories are missing from counts
* **Cause**: You ran a public check rather than an authenticated check, or your organization has third-party application restrictions enabled.
* **Resolution**:
  1. Click **Sign in with GitHub** on `/check/` to authorize access to private repositories.
  2. If the private repository belongs to an organization, ask an organization administrator to approve the GitVerse OAuth application under **Organization Settings > Third-party Access**.

### Issue: GitHub API Rate Limiting (`403 Forbidden`)
* **Cause**: Unauthenticated public requests are limited to 60 queries/hour per IP by GitHub.
* **Resolution**: Wait 15 minutes, or use the `gitverse` CLI with your signed-in GitHub account (`gh auth login`) which receives 5,000 requests/hour.

---

## 2. CLI Installation & Execution

### Issue: `gitverse: command not found` after `go install`
* **Cause**: Go's binary bin directory (`$GOPATH/bin` or `$HOME/go/bin`) is not in your system `$PATH`.
* **Resolution**:
  * **macOS/Linux**: Add to `~/.bashrc` or `~/.zshrc`:
    ```bash
    export PATH=$PATH:$(go env GOPATH)/bin
    ```
  * **Windows (PowerShell)**:
    ```powershell
    $env:Path += ";$HOME\go\bin"
    ```

---

## 3. Local Development & Preview Server

### Issue: Port 5173 is already in use
* **Cause**: A previous instance of the preview server or dev server is still running.
* **Resolution**:
  * **Windows**:
    ```powershell
    Get-Process -Name node | Stop-Process -Force
    ```
  * **macOS/Linux**:
    ```bash
    kill -9 $(lsof -t -i:5173)
    ```

