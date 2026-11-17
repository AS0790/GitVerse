<div id="top"></div>

<div align="center">

<img src="./docs/assets/gitverse-logo.png" alt="GitVerse" width="600" />

# GitVerse - GitHub Achievement Intelligence Platform

**Deterministic Milestone Visibility & Real-Time Achievement Verification.**  
*Know exactly where you stand. Turn opaque badges into actionable developer milestones.*

<br />

[![Live App](https://img.shields.io/badge/Live_App-gitverse.vercel.app-000000.svg?style=for-the-badge&logo=vercel&logoColor=white&color=000000)](https://gitverse.vercel.app)
[![CI Tests](https://github.com/TheVicky1/GitVerse/actions/workflows/test.yml/badge.svg)](https://github.com/TheVicky1/GitVerse/actions/workflows/test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-d4af37.svg?style=for-the-badge)](LICENSE)
[![Release](https://img.shields.io/github/v/release/TheVicky1/GitVerse?style=for-the-badge&color=c9971f)](https://github.com/TheVicky1/GitVerse/releases)
[![CodeQL](https://github.com/TheVicky1/GitVerse/actions/workflows/codeql.yml/badge.svg)](https://github.com/TheVicky1/GitVerse/actions/workflows/codeql.yml)
[![Hacktoberfest 2026](https://img.shields.io/badge/Hacktoberfest-2026-ff79c6.svg?style=for-the-badge&logo=hacktoberfest&logoColor=white)](https://hacktoberfest.com/)
[![Good First Issues](https://img.shields.io/github/issues/TheVicky1/GitVerse/good%20first%20issue?style=for-the-badge&color=7057ff&label=Good%20First%20Issues)](https://github.com/TheVicky1/GitVerse/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)
[![First Timers Only](https://img.shields.io/badge/First--Timers--Only-Friendly-7057ff.svg?style=for-the-badge)](https://www.firsttimersonly.com/)
[![Contributions Welcome](https://img.shields.io/badge/Contributions-Welcome-0e8a16.svg?style=for-the-badge)](CONTRIBUTING.md)
[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Privacy](https://img.shields.io/badge/Privacy-Zero_Data_Retention-0e8a16.svg?style=for-the-badge)](docs/SECURITY.md)

<br />

<!-- badges:start -->
<p align="center">
  <a href="https://gitverse.vercel.app/achievements/pull-shark/"><img src="https://github.githubassets.com/assets/pull-shark-gold-90985540b385.png" width="72" height="72" alt="Pull Shark" title="Pull Shark"></a>
  <a href="https://gitverse.vercel.app/achievements/starstruck/"><img src="https://github.githubassets.com/assets/starstruck-gold-a35e077693cc.png" width="72" height="72" alt="Starstruck" title="Starstruck"></a>
  <a href="https://gitverse.vercel.app/achievements/pair-extraordinaire/"><img src="https://github.githubassets.com/assets/pair-extraordinaire-gold-5335a3b8ae7d.png" width="72" height="72" alt="Pair Extraordinaire" title="Pair Extraordinaire"></a>
  <a href="https://gitverse.vercel.app/achievements/galaxy-brain/"><img src="https://github.githubassets.com/assets/galaxy-brain-gold-9c2c03fd1acf.png" width="72" height="72" alt="Galaxy Brain" title="Galaxy Brain"></a>
  <a href="https://gitverse.vercel.app/achievements/quickdraw/"><img src="https://github.githubassets.com/assets/quickdraw-default-39c6aec8ff89.png" width="72" height="72" alt="Quickdraw" title="Quickdraw"></a>
  <a href="https://gitverse.vercel.app/achievements/yolo/"><img src="https://github.githubassets.com/assets/yolo-default-be0bbff04951.png" width="72" height="72" alt="YOLO" title="YOLO"></a>
  <a href="https://gitverse.vercel.app/achievements/public-sponsor/"><img src="https://github.githubassets.com/assets/public-sponsor-default-9fa68986b057.png" width="72" height="72" alt="Public Sponsor" title="Public Sponsor"></a>
  <a href="https://gitverse.vercel.app/achievements/finishing-touch/"><img src="https://github.githubassets.com/assets/finishing-touch-gold-e19e1fcbfee2.png" width="72" height="72" alt="Finishing Touch" title="Finishing Touch"></a>
</p>
<!-- badges:end -->

<br />

> 🌐 **Live Web Intelligence Platform**: Access GitVerse in your browser at **[https://gitverse.vercel.app](https://gitverse.vercel.app)**  
> ⚡ **Compiled Terminal CLI**: Install globally via `go install github.com/TheVicky1/GitVerse/cmd/gitverse@latest`

<br />

[Live Demo](https://gitverse.vercel.app) · [Track Profile](https://gitverse.vercel.app/check/) · [Explore Badges](https://gitverse.vercel.app/achievements/) · [About](#about) · [Quick Start](#quick-start) · [Dual-Engine](#dual-engine-architecture) · [Features](#core-features) · [Documentation](#documentation) · [Contributing](#contributing)

<br />

<a href="https://gitverse.vercel.app">
  <img src="./docs/assets/gitverse-hero.jpg" alt="GitVerse Web Intelligence Platform Preview" width="100%" style="border-radius: 12px; border: 1px solid rgba(255, 255, 255, 0.08);" />
</a>

</div>

---

<div align="center">
  <h2>About GitVerse</h2>
</div>

While GitHub proudly showcases unlocked achievements on user profiles, it deliberately obscures what lies ahead. Upcoming tier thresholds are unexposed, progress counters are omitted from the native UI, and GitHub's official public API contains zero achievement endpoints. Developers are left to guess how close they are to reaching their next tier.

**GitVerse eliminates this asymmetry.**

Through deterministic reverse-calibration, GitVerse reads public badge histories, milestone unlock records, and activity telemetry to construct an exact, real-time map of your progress. Every number in the generated report is rigorously classified by provenance:

* **Official:** Extracted directly from GitHub's native profile schemas.
* **Calibrated:** Correlated with verified milestone thresholds across actual repositories.
* **Counted:** Computed in real-time from active repository contribution events.

### Foundational Principles

1. **Deterministic Accuracy**: Never display speculative numbers. If a metric cannot be calibrated against verified milestones, GitVerse displays the verified official tier floor.
2. **Zero-Retention Privacy**: No database, no tracking cookies, no credentials stored. OAuth access tokens used for private repository calculation are held in memory for a single execution pass and revoked immediately.
3. **Dual-Engine Architecture**: The exact same battle-tested analytical engine powers both a compiled command-line binary for terminal power-users and a blazing-fast responsive web application on Vercel.
4. **Community-Driven Open Source**: 100% open-source under the permissive MIT License, welcoming developers of all experience levels with structured micro-issues and fast-track PR reviews.

---

<div align="center">
  <h2>Quick Start</h2>
</div>

### 1. Web Intelligence Platform (Zero Install)

The fastest way to inspect any profile is directly in the browser:

1. Visit **[gitverse.vercel.app/check/](https://gitverse.vercel.app/check/)**.
2. Enter any public GitHub username, including your own.
3. Or click **Sign in with GitHub** to securely analyze private repository milestones with instantaneous token revocation.

### 2. High-Performance Terminal CLI

Engineers who live in the terminal can install and run GitVerse natively:

#### Prerequisites
- [Go 1.25+](https://go.dev/dl/) installed on your system.
- Optional: [GitHub CLI (`gh`)](https://cli.github.com/) for automatic private repository authentication.

```bash
# 1. Install the GitVerse CLI binary globally
go install github.com/TheVicky1/GitVerse/cmd/gitverse@latest

# 2. Check your own profile (automatically reuses your existing gh auth)
gitverse

# 3. Check any public GitHub profile
gitverse <username>

# 4. Include the complete catalog archive, staff badges, and notes
gitverse --all

# 5. Output structured JSON for automation or shell scripts
gitverse --json <username>
```

---

<div align="center">
  <h2>Dual-Engine Architecture</h2>
</div>

GitVerse is engineered from the ground up as a synchronized dual-engine system:

```text
                     +---------------------------------------+
                     |         Developer Entrypoints         |
                     +---------------------------------------+
                            /                         \
            [ Web Intelligence (Vercel) ]      [ Terminal CLI (Go) ]
            - Serverless Edge Handlers          - Standalone Compiled Binary
            - OLED Aurora Responsive UI         - Terminal ANSI Color Shelves
            - Interactive Category Drawers      - Machine-Readable JSON Mode
                            \                         /
                             +-----------------------+
                             |  GitVerse Core Engine |
                             +-----------------------+
                                         |
               +-------------------------+-------------------------+
               |                         |                         |
      [ Catalog Engine ]        [ Calibration Unit ]      [ Provenance Classifier ]
      - 15 Badges               - Pull Shark (Merged PRs)  - Official Verified
      - 4 Tier Thresholds       - Starstruck (Stargazers) - Calibrated Milestone
      - 6 Profile Highlights    - Galaxy Brain (Answers)   - Counted Telemetry
               |                         |                         |
               +-------------------------+-------------------------+
                                         |
                         +-------------------------------+
                         |   GitHub Telemetry Pipeline   |
                         +-------------------------------+
                                 /               \
                 [ Public Profile Scraper ]   [ Live GraphQL & REST API ]
                 - Profile History & Tiers    - Authenticated PR & Star Counts
                 - In-Memory Parallel Fetch   - Zero Stored Tokens (Revoked)
```

> For deep architectural details, see the [Architecture Guide](docs/ARCHITECTURE.md) and [Data Model Specification](docs/DATA_MODEL.md).

---

<div align="center">
  <h2>Core Features</h2>
</div>

<!-- catalog:start -->
### Earnable

You can still earn these.

| Badge | Achievement | How to earn it | Tiers |
|:-:|---|---|:-:|
| <img src="https://github.githubassets.com/assets/pull-shark-gold-90985540b385.png" width="40" height="40" alt="Pull Shark badge"> | [Pull Shark](https://gitverse.vercel.app/achievements/pull-shark/) | Open pull requests that get merged<br><sub>gitverse: count, checked against GitHub's milestones</sub> | <table><tr><th>Default</th><th>Bronze</th><th>Silver</th><th>Gold</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/pull-shark-default-498c279a747d.png" width="48" height="48" alt="Pull Shark base"></td><td align="center"><img src="https://github.githubassets.com/assets/pull-shark-bronze-a37accb528d1.png" width="48" height="48" alt="Pull Shark bronze"></td><td align="center"><img src="https://github.githubassets.com/assets/pull-shark-silver-0643f87ac9fd.png" width="48" height="48" alt="Pull Shark silver"></td><td align="center"><img src="https://github.githubassets.com/assets/pull-shark-gold-90985540b385.png" width="48" height="48" alt="Pull Shark gold"></td></tr><tr><td align="center">2</td><td align="center">16</td><td align="center">128</td><td align="center">1024</td></tr></table> |
| <img src="https://github.githubassets.com/assets/starstruck-gold-a35e077693cc.png" width="40" height="40" alt="Starstruck badge"> | [Starstruck](https://gitverse.vercel.app/achievements/starstruck/) | Create a repository that gets stars<br><sub>gitverse: stars of the repository GitHub names</sub> | <table><tr><th>Default</th><th>Bronze</th><th>Silver</th><th>Gold</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/starstruck-default-b6610abad518.png" width="48" height="48" alt="Starstruck base"></td><td align="center"><img src="https://github.githubassets.com/assets/starstruck-bronze-b71ec4d7766a.png" width="48" height="48" alt="Starstruck bronze"></td><td align="center"><img src="https://github.githubassets.com/assets/starstruck-silver-275c101223c0.png" width="48" height="48" alt="Starstruck silver"></td><td align="center"><img src="https://github.githubassets.com/assets/starstruck-gold-a35e077693cc.png" width="48" height="48" alt="Starstruck gold"></td></tr><tr><td align="center">16</td><td align="center">128</td><td align="center">512</td><td align="center">4096</td></tr></table> |
| <img src="https://github.githubassets.com/assets/pair-extraordinaire-gold-5335a3b8ae7d.png" width="40" height="40" alt="Pair Extraordinaire badge"> | [Pair Extraordinaire](https://gitverse.vercel.app/achievements/pair-extraordinaire/) | Coauthor commits on merged pull requests<br><sub>gitverse: official tier (no public counting rule)</sub> | <table><tr><th>Default</th><th>Bronze</th><th>Silver</th><th>Gold</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/pair-extraordinaire-default-579438a20e01.png" width="48" height="48" alt="Pair Extraordinaire base"></td><td align="center"><img src="https://github.githubassets.com/assets/pair-extraordinaire-bronze-797214f41e7e.png" width="48" height="48" alt="Pair Extraordinaire bronze"></td><td align="center"><img src="https://github.githubassets.com/assets/pair-extraordinaire-silver-a0b2b472348c.png" width="48" height="48" alt="Pair Extraordinaire silver"></td><td align="center"><img src="https://github.githubassets.com/assets/pair-extraordinaire-gold-5335a3b8ae7d.png" width="48" height="48" alt="Pair Extraordinaire gold"></td></tr><tr><td align="center">1</td><td align="center">10</td><td align="center">24</td><td align="center">48</td></tr></table> |
| <img src="https://github.githubassets.com/assets/galaxy-brain-gold-9c2c03fd1acf.png" width="40" height="40" alt="Galaxy Brain badge"> | [Galaxy Brain](https://gitverse.vercel.app/achievements/galaxy-brain/) | Have discussion answers accepted<br><sub>gitverse: count of accepted answers</sub> | <table><tr><th>Default</th><th>Bronze</th><th>Silver</th><th>Gold</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/galaxy-brain-default-847262c21056.png" width="48" height="48" alt="Galaxy Brain base"></td><td align="center"><img src="https://github.githubassets.com/assets/galaxy-brain-bronze-da3c4829fd22.png" width="48" height="48" alt="Galaxy Brain bronze"></td><td align="center"><img src="https://github.githubassets.com/assets/galaxy-brain-silver-2ce8da282a85.png" width="48" height="48" alt="Galaxy Brain silver"></td><td align="center"><img src="https://github.githubassets.com/assets/galaxy-brain-gold-9c2c03fd1acf.png" width="48" height="48" alt="Galaxy Brain gold"></td></tr><tr><td align="center">2</td><td align="center">8</td><td align="center">16</td><td align="center">32</td></tr></table> |
| <img src="https://github.githubassets.com/assets/quickdraw-default-39c6aec8ff89.png" width="40" height="40" alt="Quickdraw badge"> | [Quickdraw](https://gitverse.vercel.app/achievements/quickdraw/) | Close an issue or pull request within 5 minutes of opening it<br><sub>gitverse: earned or not (official)</sub> | <table><tr><th>Default</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/quickdraw-default-39c6aec8ff89.png" width="48" height="48" alt="Quickdraw base"></td></tr><tr><td align="center">1</td></tr></table> |
| <img src="https://github.githubassets.com/assets/yolo-default-be0bbff04951.png" width="40" height="40" alt="YOLO badge"> | [YOLO](https://gitverse.vercel.app/achievements/yolo/) | Merge a pull request without a review<br><sub>gitverse: earned or not (official)</sub> | <table><tr><th>Default</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/yolo-default-be0bbff04951.png" width="48" height="48" alt="YOLO base"></td></tr><tr><td align="center">1</td></tr></table> |
| <img src="https://github.githubassets.com/assets/public-sponsor-default-9fa68986b057.png" width="40" height="40" alt="Public Sponsor badge"> | [Public Sponsor](https://gitverse.vercel.app/achievements/public-sponsor/) | Sponsor someone publicly through GitHub Sponsors<br><sub>gitverse: earned or not (official)</sub> | <table><tr><th>Default</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/public-sponsor-default-9fa68986b057.png" width="48" height="48" alt="Public Sponsor base"></td></tr><tr><td align="center">1</td></tr></table> |
| <img src="https://github.githubassets.com/assets/finishing-touch-gold-e19e1fcbfee2.png" width="40" height="40" alt="Finishing Touch badge"> | [Finishing Touch](https://gitverse.vercel.app/achievements/finishing-touch/) | Get pull requests that improve GitHub's UI quality merged<br><sub>gitverse: official tier (no public counting rule)</sub> | <table><tr><th>Default</th><th>Bronze</th><th>Silver</th><th>Gold</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/finishing-touch-default-764b3a996d62.png" width="48" height="48" alt="Finishing Touch base"></td><td align="center"><img src="https://github.githubassets.com/assets/finishing-touch-bronze-7c49185628f3.png" width="48" height="48" alt="Finishing Touch bronze"></td><td align="center"><img src="https://github.githubassets.com/assets/finishing-touch-silver-75ab9a5c22ba.png" width="48" height="48" alt="Finishing Touch silver"></td><td align="center"><img src="https://github.githubassets.com/assets/finishing-touch-gold-e19e1fcbfee2.png" width="48" height="48" alt="Finishing Touch gold"></td></tr><tr><td align="center">1</td><td align="center">4</td><td align="center">16</td><td align="center">32</td></tr></table> |

### Switched off by GitHub

GitHub tested these and switched them off; nobody can earn them now.

| Badge | Achievement | How to earn it | Tiers |
|:-:|---|---|:-:|
| <img src="https://raw.githubusercontent.com/Schweinepriester/github-profile-achievements/549ba7291ccb68efcff5daf8af8c2c0ac6817b6f/images/tiers/heart-on-your-sleeve-gold.png" width="40" height="40" alt="Heart On Your Sleeve badge"> | [Heart On Your Sleeve](https://gitverse.vercel.app/achievements/heart-on-your-sleeve/) | React to something on GitHub with a ❤️ emoji<br><sub>gitverse: earned or not</sub> | <table><tr><th>Default</th><th>Bronze</th><th>Silver</th><th>Gold</th></tr><tr><td align="center"><img src="https://raw.githubusercontent.com/Schweinepriester/github-profile-achievements/549ba7291ccb68efcff5daf8af8c2c0ac6817b6f/images/heart-on-your-sleeve-default.png" width="48" height="48" alt="Heart On Your Sleeve base"></td><td align="center"><img src="https://raw.githubusercontent.com/Schweinepriester/github-profile-achievements/549ba7291ccb68efcff5daf8af8c2c0ac6817b6f/images/tiers/heart-on-your-sleeve-bronze.png" width="48" height="48" alt="Heart On Your Sleeve bronze"></td><td align="center"><img src="https://raw.githubusercontent.com/Schweinepriester/github-profile-achievements/549ba7291ccb68efcff5daf8af8c2c0ac6817b6f/images/tiers/heart-on-your-sleeve-silver.png" width="48" height="48" alt="Heart On Your Sleeve silver"></td><td align="center"><img src="https://raw.githubusercontent.com/Schweinepriester/github-profile-achievements/549ba7291ccb68efcff5daf8af8c2c0ac6817b6f/images/tiers/heart-on-your-sleeve-gold.png" width="48" height="48" alt="Heart On Your Sleeve gold"></td></tr><tr><td align="center">?</td><td align="center">16</td><td align="center">128</td><td align="center">?</td></tr></table> |
| <img src="https://raw.githubusercontent.com/Schweinepriester/github-profile-achievements/549ba7291ccb68efcff5daf8af8c2c0ac6817b6f/images/tiers/open-sourcerer-gold.png" width="40" height="40" alt="Open Sourcerer badge"> | [Open Sourcerer](https://gitverse.vercel.app/achievements/open-sourcerer/) | Have pull requests merged in several public repositories<br><sub>gitverse: earned or not</sub> | <table><tr><th>Default</th><th>Bronze</th><th>Silver</th><th>Gold</th></tr><tr><td align="center"><img src="https://raw.githubusercontent.com/Schweinepriester/github-profile-achievements/549ba7291ccb68efcff5daf8af8c2c0ac6817b6f/images/open-sourcerer-default.png" width="48" height="48" alt="Open Sourcerer base"></td><td align="center"><img src="https://raw.githubusercontent.com/Schweinepriester/github-profile-achievements/549ba7291ccb68efcff5daf8af8c2c0ac6817b6f/images/tiers/open-sourcerer-bronze.png" width="48" height="48" alt="Open Sourcerer bronze"></td><td align="center"><img src="https://raw.githubusercontent.com/Schweinepriester/github-profile-achievements/549ba7291ccb68efcff5daf8af8c2c0ac6817b6f/images/tiers/open-sourcerer-silver.png" width="48" height="48" alt="Open Sourcerer silver"></td><td align="center"><img src="https://raw.githubusercontent.com/Schweinepriester/github-profile-achievements/549ba7291ccb68efcff5daf8af8c2c0ac6817b6f/images/tiers/open-sourcerer-gold.png" width="48" height="48" alt="Open Sourcerer gold"></td></tr><tr><td align="center">?</td><td align="center">8</td><td align="center">16</td><td align="center">64</td></tr></table> |

### Retired

No longer awarded. Profiles that have them keep them.

| Badge | Achievement | How to earn it | Tiers |
|:-:|---|---|:-:|
| <img src="https://github.githubassets.com/assets/arctic-code-vault-contributor-default-df8d74122a06.png" width="40" height="40" alt="Arctic Code Vault Contributor badge"> | [Arctic Code Vault Contributor](https://gitverse.vercel.app/achievements/arctic-code-vault-contributor/) | Code in a repository of the 2020 GitHub Archive Program<br><sub>gitverse: earned or not</sub> | <table><tr><th>Default</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/arctic-code-vault-contributor-default-df8d74122a06.png" width="48" height="48" alt="Arctic Code Vault Contributor base"></td></tr><tr><td align="center">1</td></tr></table> |
| <img src="https://github.githubassets.com/assets/mars-2020-contributor-default-c9cad8495acf.png" width="40" height="40" alt="Mars 2020 Contributor badge"> | [Mars 2020 Contributor](https://gitverse.vercel.app/achievements/mars-2020-contributor/) | Code in a repository used by the Mars 2020 Helicopter Mission<br><sub>gitverse: earned or not</sub> | <table><tr><th>Default</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/mars-2020-contributor-default-c9cad8495acf.png" width="48" height="48" alt="Mars 2020 Contributor base"></td></tr><tr><td align="center">1</td></tr></table> |

### GitHub staff only

Only GitHub staff could earn these.

| Badge | Achievement | How to earn it | Tiers |
|:-:|---|---|:-:|
| <img src="https://github.githubassets.com/assets/proxima-pioneer-default-51cd4e4969d5.png" width="40" height="40" alt="Proxima Pioneer badge"> | [Proxima Pioneer](https://gitverse.vercel.app/achievements/proxima-pioneer/) | GitHub staff: contributed to the Proxima proof of concept<br><sub>gitverse: earned or not</sub> | <table><tr><th>Default</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/proxima-pioneer-default-51cd4e4969d5.png" width="48" height="48" alt="Proxima Pioneer base"></td></tr><tr><td align="center">1</td></tr></table> |
| <img src="https://github.githubassets.com/assets/proxima-staffshipper-default-84f658288021.png" width="40" height="40" alt="Proxima Staffshipper badge"> | [Proxima Staffshipper](https://gitverse.vercel.app/achievements/proxima-staffshipper/) | GitHub staff: shipped the Proxima staffship instance<br><sub>gitverse: earned or not</sub> | <table><tr><th>Default</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/proxima-staffshipper-default-84f658288021.png" width="48" height="48" alt="Proxima Staffshipper base"></td></tr><tr><td align="center">1</td></tr></table> |
| <img src="https://github.githubassets.com/assets/proxima-staffuser-default-02bf6163ea31.png" width="40" height="40" alt="Proxima Staffuser badge"> | [Proxima Staffuser](https://gitverse.vercel.app/achievements/proxima-staffuser/) | GitHub staff: your team was onboarded to Proxima<br><sub>gitverse: earned or not</sub> | <table><tr><th>Default</th></tr><tr><td align="center"><img src="https://github.githubassets.com/assets/proxima-staffuser-default-02bf6163ea31.png" width="48" height="48" alt="Proxima Staffuser base"></td></tr><tr><td align="center">1</td></tr></table> |

Profile highlights it also reports: Developer Program Member, Security Bug Bounty Hunter, GitHub Campus Expert, GitHub Star, Pro, Security advisory credit.

_Catalog last verified against live profiles on 2026-10-17._
<!-- catalog:end -->

---

<div align="center">
  <h2>Visual Identity & Design Philosophy</h2>
</div>

GitVerse is crafted with a signature **Space Aurora** minimalist luxury dark mode aesthetic:

* **OLED Obsidian Canvas**: Deep `#0a0a0b` canvas foundation eliminating interface fatigue.
* **Aurora Topographic Contours**: Subtle, glowing vector contour lines (`rgba(255, 255, 255, 0.035)`) evoking orbital navigation.
* **Precious Metal Medal Shelves**: Tier indicators styled with authentic metallic finishes:
  * 🥇 **Gold (x4)**: `#c9971f`
  * 🥈 **Silver (x3)**: `#e2e8f0`
  * 🥉 **Bronze (x2)**: `#d97706`
  * 🔘 **Base (x1)**: `#71717a`
* **Zero Layout Jitter**: Interactive loading spinners and predictive prefetch caching ensure smooth transitions under 20ms.

> For complete guidelines, typography scales, and token values, see the [Design System Specification](docs/DESIGN_SYSTEM.md).

---

<div align="center">
  <h2>Project Structure</h2>
</div>

```text
GitVerse/
├── .github/                   # Workflows (CI, release, canary), issue & PR templates
│   ├── ISSUE_TEMPLATE/        # Structured YAML issue templates
│   └── workflows/             # Automated test, canary, and release automation
├── api/                       # Vercel Serverless Function entrypoints
│   ├── markdown/              # Content-negotiated Markdown representations
│   ├── private/               # Ephemeral private calculation & token revocation
│   └── result/                # Public report serverless handler
├── cmd/
│   ├── gendocs/               # SSG & documentation generation engine
│   └── gitverse/              # Main compiled CLI application entrypoint
├── docs/                      # Comprehensive technical & contributor documentation
│   ├── assets/                # Logos, screenshots, and visual assets
│   ├── ARCHITECTURE.md        # Dual-engine system architecture
│   ├── CONTRIBUTING-BEGINNERS.md # Step-by-step first PR guide
│   ├── DATA_MODEL.md          # Schema specifications & provenance rules
│   ├── DESIGN_SYSTEM.md       # Color tokens, typography, and styling
│   ├── FEATURES.md            # Complete catalog breakdown and unlock rules
│   └── SECURITY.md            # Zero-retention security guarantees
├── internal/
│   ├── catalog/               # Single source of truth for all 15 achievements
│   ├── counters/              # Live API telemetry collectors (PRs, stars, discussions)
│   ├── official/              # GitHub HTML profile parsers & detail extraction
│   ├── progress/              # Tier milestone calculations and calibration floor
│   ├── render/                # Terminal ANSI color & shelf formatting
│   ├── report/                # Aggregator binding profile data & live counters
│   └── webui/                 # Web templates, static assets, and scripts
├── site/                      # Static production site assets and documentation
├── scripts/                   # Local preview server, asset generators, benchmarks
├── CHANGELOG.md               # Semantic Versioning release history
├── CONTRIBUTING.md            # Comprehensive contributor guidelines
├── CONTRIBUTORS.md            # Community recognition hall of fame
├── LICENSE                    # Permissive MIT License
└── SECURITY.md                # Responsible vulnerability disclosure policy
```

---

<div align="center">
  <h2>Documentation</h2>
</div>

Comprehensive technical specifications, operational runbooks, and contributor guides are maintained in [`docs/`](docs/README.md):

| Document | Purpose |
| :--- | :--- |
| 🌱 [**Beginner Contributing**](docs/CONTRIBUTING-BEGINNERS.md) | Step-by-step zero-to-PR walkthrough for first-time open-source contributors |
| 💻 [**Windows Setup Guide**](docs/CONTRIBUTING-WINDOWS.md) | Environment configuration and PowerShell instructions for Windows developers |
| 🏗️ [**Architecture**](docs/ARCHITECTURE.md) | Dual-engine system architecture, in-memory execution, and data flow |
| 📋 [**Feature Inventory**](docs/FEATURES.md) | Full breakdown of all 15 achievements, tiers, and profile highlights |
| 🗄️ [**Data Model**](docs/DATA_MODEL.md) | Provenance levels (`official`, `calibrated`, `counted`) and badge schemas |
| 🎨 [**Design System**](docs/DESIGN_SYSTEM.md) | Space Aurora aesthetic, OLED obsidian canvas, and medal shelf tokens |
| 💻 [**Development Guide**](docs/DEVELOPMENT.md) | Local environment setup, preview server, and build commands |
| 🧪 [**Testing Strategy**](docs/TESTING.md) | Weekly canary test suites, unit testing, and mock verification |
| 🔒 [**Security Architecture**](docs/SECURITY.md) | Zero-retention model, in-memory execution, and token revocation |
| 🌿 [**Git Workflow**](docs/GIT_WORKFLOW.md) | Conventional Commits, branch hygiene, and PR standards |
| 📦 [**Release Management**](docs/RELEASE_MANAGEMENT.md) | Semantic Versioning policy and release milestone checklists (v1.0–v1.3) |
| 🗺️ [**Product Roadmap**](docs/ROADMAP.md) | Implementation milestones and planned platform features |
| 🔧 [**Troubleshooting**](docs/TROUBLESHOOTING.md) | Practical fixes for rate limits, token scopes, and local preview issues |
| 🌐 [**GitHub Metadata**](docs/GITHUB_METADATA.md) | Repository tags, descriptions, and open-source directory metadata |

---

<div align="center">
  <h2>Contributing</h2>
</div>

Contributions are warmly welcomed! GitVerse is designed to be one of the most welcoming and well-structured open-source developer tools on GitHub.

> 🚀 **New to Open Source? Start with a Micro-Contribution!**  
> We deliberately scope our beginner tasks into **single-file micro-issues** (15–30 minutes) with exact file pointers, clear acceptance criteria, and step-by-step guidance.  
> 👉 [**Browse Good First Issues**](https://github.com/TheVicky1/GitVerse/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) &nbsp;|&nbsp; 📖 [**Read the Beginner's Guide**](docs/CONTRIBUTING-BEGINNERS.md)

### 🌟 Why Contribute to GitVerse?

* **⚡ Instant Local Setup**: Run the local preview server in seconds with zero complex database setup.
* **🎯 Highly Scoped Micro-Issues**: Tasks touch isolated components, parser fixtures, or documentation files with unambiguous verification steps.
* **🚀 Fast Maintainer Feedback**: Pull requests receive prompt, constructive reviews from the core maintainer team.
* **🏛️ Permanent Recognition**: Every contributor is permanently acknowledged in [**CONTRIBUTORS.md**](CONTRIBUTORS.md) and on the [**GitHub Contributors Graph**](https://github.com/TheVicky1/GitVerse/graphs/contributors).

### ⚡ 6-Step Contributor Workflow

1. **Pick an Issue**: Browse [**Good First Issues**](https://github.com/TheVicky1/GitVerse/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22).
2. **Fork & Clone**: Fork GitVerse and clone it locally:
   ```bash
   git clone https://github.com/<your-username>/GitVerse.git
   cd GitVerse
   ```
3. **Create a Feature Branch**:
   ```bash
   git checkout -b feat/achievement-milestone-enhancement
   ```
4. **Make Your Change**: Edit the designated files following project guidelines.
5. **Verify Locally**:
   ```bash
   node scripts/serve-preview.cjs
   node --test internal/webui/jstest/*.test.mjs
   ```
6. **Submit a Pull Request**: Push your branch and open a PR with a description of your work.

---

<div align="center">
  <h2>Contributors & Community</h2>
</div>

GitVerse is crafted and maintained with care by **[Vicky Patel](https://github.com/TheVicky1)** alongside an exceptional community of open-source developers.

<div align="center">

[![Contributors Count](https://img.shields.io/github/contributors/TheVicky1/GitVerse?style=for-the-badge&color=d4af37)](https://github.com/TheVicky1/GitVerse/graphs/contributors)

[![GitVerse Contributors](https://contrib.rocks/image?repo=TheVicky1/GitVerse)](https://github.com/TheVicky1/GitVerse/graphs/contributors)

*A huge thank you to everyone who has contributed to making GitHub achievements transparent and deterministic!*

👉 [**View all contributors on GitHub**](https://github.com/TheVicky1/GitVerse/graphs/contributors) • [**Contributors Hall of Fame**](CONTRIBUTORS.md)

</div>

---

<div align="center">
  <h2>License</h2>
</div>

GitVerse is open-source software distributed under the permissive [MIT License](LICENSE).

---

<div align="center">

**Deterministic. Transparent. Engineered for Clarity.**

[⬆ Back to top](#top)

</div>
