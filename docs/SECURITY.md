# GitVerse Security Architecture & Privacy Policy

Security, data minimalism, and developer privacy are foundational to the design of **GitVerse**.

---

## 1. Zero Data Retention Architecture

GitVerse is engineered to operate without persistent storage:
* **No Database**: GitVerse maintains no SQL or NoSQL database. No user accounts, credentials, access tokens, or repository names are ever persisted to disk.
* **In-Memory Calculations**: Profile analyses are executed entirely in transient server memory.
* **Ephemeral OAuth Tokens**: When a developer clicks "Sign in with GitHub" to calculate private repository contributions:
  1. GitHub exchanges the temporary authorization code for an OAuth access token.
  2. The serverless handler uses this token to run a single read pass against GitHub's API.
  3. The handler explicitly invokes GitHub's token revocation endpoint (`DELETE /applications/{client_id}/grant`) **before returning the HTTP response**.
  4. The token is expunged from memory.

---

## 2. Token Security & Scope Minimization

* **CLI Execution**: The `gitverse` CLI does not create or store personal access tokens (PATs). It reuses your existing local session from the official GitHub CLI (`gh auth token`), sending requests directly to `api.github.com` via the official `go-gh` library.
* **OAuth Scopes**: GitVerse requests only the minimum required scope (`repo`) necessary to count private merged pull requests. GitVerse never requests write permissions, administrative access, or webhook modification rights.

---

## 3. Analytics & Telemetry Privacy

GitVerse utilizes privacy-sanitized Google Analytics 4 (GA4):
* **Consent Gated**: Analytics scripts load only after an explicit user opt-in click.
* **No PII Transmission**: Usernames, repository names, and query parameters are stripped before events are transmitted (`page_location` removes the query string).

---

## 4. Reporting Security Vulnerabilities

We take the security of GitVerse seriously. If you discover a security vulnerability, please report it privately:

1. **Preferred Method**: Open a [GitHub Private Security Advisory](https://github.com/TheVicky1/GitVerse/security/advisories/new).
2. **Direct Contact**: Email maintainer **Vicky Patel** at **vickybiz55808@gmail.com**.

Please include detailed steps to reproduce the vulnerability. You will receive an acknowledgment within 48 hours and regular updates on the resolution.

