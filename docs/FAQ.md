# Frequently Asked Questions (FAQ)

### 1. Does GitVerse store my GitHub token or private code?
**No, absolutely never.** GitVerse enforces a strict zero-retention architecture. Tokens used during "Sign in with GitHub" are kept in transient server memory for a single calculation pass and immediately revoked via GitHub's OAuth revocation API before the response completes. No database exists.

### 2. Why does GitVerse show progress that GitHub doesn't?
GitHub shows which badges you have unlocked, but does not display your current progress toward future tiers. GitVerse reverse-calibrates milestone history against real-time API telemetry to calculate your exact distance to the next tier.

### 3. Why are some numbers labelled "official" and others "calibrated" or "counted"?
* **Official:** Extracted directly from GitHub profile HTML markup.
* **Calibrated:** Computed through the API and matched against explicit milestone thresholds named by GitHub.
* **Counted:** Derived from API activity telemetry where no public reference milestone exists.

### 4. Can I use GitVerse in the terminal without a web browser?
Yes! Install the CLI globally:
```bash
go install github.com/TheVicky1/GitVerse/cmd/gitverse@latest
gitverse torvalds
```
