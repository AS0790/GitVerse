# Release Management & Versioning Policy

GitVerse strictly adheres to [Semantic Versioning 2.0.0](https://semver.org/) (`vMAJOR.MINOR.PATCH`) and [Keep a Changelog](https://keepachangelog.com/) standards.

---

## 1. Release Milestone Architecture

### 🏷️ `v1.0.0` - Foundation & Core Engine
* Initial release of the GitVerse Go CLI engine.
* Complete catalog tracking of all 15 known GitHub achievements and 6 profile highlights.
* Reverse-calibration engine for Pull Shark and Starstruck milestone thresholds.
* ANSI-colored terminal report shelves and `--json` machine-readable output.

### 🏷️ `v1.1.0` - Web Platform & OAuth Integration
* Serverless web architecture deployed to Vercel (`gitverse.vercel.app`).
* Zero-retention GitHub OAuth sign-in enabling private repository milestone calculation.
* Ephemeral token lifecycle: instant token revocation upon calculation completion.
* Content negotiation supporting Markdown responses via `Accept: text/markdown`.

### 🏷️ `v1.2.0` - Space Aurora Design System Overhaul
* Complete OLED dark mode redesign (`#0a0a0b`) with subtle topographic contour backdrop.
* Luxury metallic tier shelves (Gold, Silver, Bronze, Base).
* Interactive achievement category drawers (On display, Still to collect, Archive).
* Mobile-responsive layout and typography polish.

### 🏷️ `v1.3.0` - High-Performance Dual-Engine & Open Source Ecosystem
* Concurrent parallel fetching pipeline with persistent Keep-Alive connection pooling.
* Response time optimized by over 50% cold (~1.7s) and sub-20ms on cached queries.
* Predictive background prefetching and immediate button loading feedback.
* Comprehensive technical documentation suite, beginner contribution guides, and community infrastructure.

---

## 2. Release Procedure Checklist

1. **Verify Quality Gates**:
   ```bash
   node --test internal/webui/jstest/*.test.mjs
   go test ./...
   ```
2. **Update Changelog**:
   Record additions, fixes, and performance improvements in `CHANGELOG.md` under the target version header.
3. **Tag the Release**:
   ```bash
   git tag -a v1.3.0 -m "release: v1.3.0 - High-Performance Dual-Engine & Open Source Ecosystem"
   git push origin v1.3.0
   ```
4. **Automated CI Build**:
   GitHub Actions compiles multi-platform binaries (Linux, macOS, Windows) and publishes release artifacts.

