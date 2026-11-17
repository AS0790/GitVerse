# Security Policy

Security, privacy, and zero data retention are core principles of **GitVerse**.

---

## 🔒 Zero Data Retention Guarantee

* **No Persistent Storage**: GitVerse does NOT store, log, or persist any user credentials, Personal Access Tokens (PATs), OAuth tokens, or private repository data.
* **Instant Token Revocation**: When using "Sign in with GitHub", the serverless handler uses your temporary OAuth token for exactly one calculation pass and immediately revokes it with GitHub before completing the response.
* **CLI Privacy**: The `gitverse` command-line executable reuses your existing local `gh auth` login. It sends requests directly to `api.github.com` via the official `go-gh` library. No data ever passes through any third-party intermediary.

---

## 🛡️ Supported Versions

We release patches and security fixes for the following versions:

| Version | Supported |
| :--- | :---: |
| `v1.3.x` | :white_check_mark: Yes |
| `v1.2.x` | :white_check_mark: Yes |
| `< v1.2.0` | :x: No |

---

## 🐛 Reporting a Vulnerability

If you discover a security vulnerability or privacy concern in GitVerse:

1. **Do NOT open a public GitHub issue.**
2. Submit a report privately via [GitHub Private Vulnerability Reporting](https://github.com/TheVicky1/GitVerse/security/advisories/new).
3. Or email the lead maintainer directly: **Vicky Patel** at **vickybiz55808@gmail.com**.

We will acknowledge receipt within 48 hours, provide an initial assessment, and keep you informed as we develop and deploy a patch.
