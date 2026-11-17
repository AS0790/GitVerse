# GitVerse System Architecture

GitVerse is engineered around a **Dual-Engine Architecture**: two distinct developer interfaces—a compiled command-line binary (`gitverse`) and a serverless web platform (`mygitverse.vercel.app`)—powered by the exact same core analytical algorithms.

---

## 1. High-Level Architecture Diagram

```text
+---------------------------------------------------------------------------------+
|                                Developer Interfaces                             |
|                                                                                 |
|   [ Compiled Terminal CLI (Go) ]            [ Responsive Web Platform (Vercel) ]|
|   - ANSI Color Shelf Formatting              - Space Aurora OLED UI             |
|   - Zero-Config gh auth reuse                - OAuth Ephemeral Auth Flow        |
|   - Machine-Readable JSON Mode               - Predictive Prefetch Cache        |
+---------------------------------------------------------------------------------+
                                       |
                                       v
+---------------------------------------------------------------------------------+
|                              GitVerse Core Engine                               |
|                                                                                 |
|   [ Achievement Catalog ]         [ Calibration Unit ]    [ Provenance Engine ] |
|   - 15 Badge Definitions          - Pull Shark (PRs)      - Official Verified   |
|   - 4 Tier Thresholds             - Starstruck (Stars)    - Calibrated Match    |
|   - 6 Profile Highlights          - Galaxy Brain (Q&A)    - Counted Telemetry   |
+---------------------------------------------------------------------------------+
                                       |
                                       v
+---------------------------------------------------------------------------------+
|                            Data Acquisition Pipeline                            |
|                                                                                 |
|   [ Public Profile Scraper ]                 [ GitHub REST & GraphQL API ]      |
|   - Achievements HTML Parser                 - Merged Pull Requests Query       |
|   - Badge /detail Milestone Scraper          - Repository Stargazer Index       |
|   - Concurrent Parallel Connection Pool      - Discussion Answers Telemetry     |
+---------------------------------------------------------------------------------+
                                       |
                                       v
+---------------------------------------------------------------------------------+
|                                GitHub Platform                                  |
|   https://github.com/users/:username/achievements/:slug/detail                  |
|   https://api.github.com/search/issues, /user/repos                             |
+---------------------------------------------------------------------------------+
```

---

## 2. Core Subsystems

### A. The Catalog Engine (`internal/catalog`)
The catalog is the single source of truth across the entire platform. It defines:
- **Badge Metadata:** Name, description, kind (`earnable`, `switched-off`, `retired`, `internal`), and official CDN assets.
- **Milestone Thresholds:** Exact counts required to reach each tier (e.g., Pull Shark: 2, 16, 128, 1024).
- **Evaluation Rules:** Associated counters and milestone matching algorithms.

### B. Reverse-Calibration Unit (`internal/progress`)
Because GitHub's official API does not expose badge progress, GitVerse pairs two distinct data streams:
1. **GitHub's Badge History Pages:** Official HTML history records listing every milestone unlock timestamp and linked repository.
2. **API Telemetry:** Live counts of merged pull requests, repository stars, and accepted discussion answers.

The calibration unit verifies that live telemetry never dips below the official badge tier floor recorded by GitHub, guaranteeing 100% mathematical consistency.

### C. Provenance Classification
Every single metric displayed is categorized into one of three tiers of confidence:
* **`official`**: Directly extracted from GitHub's native profile HTML markup.
* **`calibrated`**: Counted through GitHub's API and matched against explicit milestone thresholds named by GitHub.
* **`counted`**: Computed directly from contribution activity where no public calibration milestone exists.

---

## 3. Web Performance & Concurrency Model

### Parallel Pipeline
Earlier versions fetched badge details sequentially. The current architecture employs a fully parallelized concurrency pipeline:
1. **Phase 1 (Discovery):** Fetch user achievements tab (`https://github.com/:username?tab=achievements`) to extract earned badge slugs.
2. **Phase 2 (Concurrent Fan-out):** Simultaneously fetch:
   - Detail milestones for every earned badge.
   - Merged pull request counts via GitHub Search API.
   - Stargazers for owned repositories via GitHub Repos API.
3. **Phase 3 (Aggregation & Render):** Apply calibration floors, construct badge shelf view models, and render the response.

### Connection Reuse (Keep-Alive)
All outgoing HTTPS requests utilize persistent connection pooling with `https.Agent({ keepAlive: true, maxSockets: 32 })`. This eliminates repeated TLS 1.3 handshakes and reduces network latency by 60–75%.

### Request Coalescing & Memory Caching
* **In-Flight Coalescing:** Simultaneous queries for the same username share a single promise, preventing redundant network requests.
* **In-Memory Cache:** Profile reports are cached in memory for 5 minutes, delivering repeated checks in under **20 ms**.

---

## 4. Zero Data Retention Protocol

GitVerse enforces a strict zero-retention security boundary:
* **No Database:** GitVerse operates entirely without persistent database storage.
* **Single-Use OAuth:** Tokens obtained during "Sign in with GitHub" are held in server memory for the duration of a single request (~600ms) and immediately revoked via GitHub's OAuth revocation endpoint.
* **No Telemetry Leakage:** Google Analytics queries exclude all usernames and private repository names.

