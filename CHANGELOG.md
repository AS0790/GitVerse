# Changelog

All notable changes to **GitVerse** will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [v1.3.1] - 2026-11-18

### Added
- **Automated CI Security & Quality Workflows**: Added CodeQL static analysis and automated markdown link verification actions.
- **Structured Community Templates**: Added dedicated YAML issue templates for Feature Requests and Documentation Improvements, plus a standardized PR verification template.
- **Extended Technical Documentation Suite**: Added `docs/FAQ.md`, `docs/GLOSSARY.md`, `docs/API.md`, `docs/BENCHMARKS.md`, and `docs/ACHIEVEMENTS_GUIDE.md`.

---

## [v1.3.0] - 2026-10-31

### Added
- **Concurrent Concurrency Pipeline**: Refactored detail fetching so all badge detail requests, PR queries, and stargazer metrics execute concurrently in parallel via `Promise.all`.
- **Persistent Connection Pooling**: Integrated HTTP Keep-Alive agent (`https.Agent`) with connection reuse, cutting network roundtrip latency by up to 60%.
- **Request Coalescing & In-Memory Caching**: Added in-flight request deduplication and 5-minute memory caching, dropping cached response times down to 17ms.
- **Interactive Button Loading States**: Immediate button state transition (`[ ⟳ Checking… ]`) with locked width and `aria-busy` feedback on form submit.
- **Predictive Background Prefetching**: Debounced background prefetching on username typing and button hover.
- **Complete Open-Source Ecosystem**: Added comprehensive `docs/` folder, `CONTRIBUTORS.md`, `SUPPORT.md`, beginner zero-to-PR guide, and Windows development guide.

### Changed
- Replaced legacy author attributions with lead maintainer Vicky Patel ([@TheVicky1](https://github.com/TheVicky1)).
- Enhanced About page copy to executive technical prose with provenance categories: Official, Calibrated, and Counted.

---

## [v1.2.0] - 2026-10-18

### Added
- **Space Aurora Design System**: Complete luxury OLED dark mode overhaul (`#0a0a0b`) with mathematical orbital topographic contour backdrop.
- **Precious Metallic Tier Shelves**: Dedicated styles for Gold (`#c9971f`), Silver (`#e2e8f0`), Bronze (`#d97706`), and Base (`#71717a`) tiers.
- **Categorized Drawer Shelves**: Structured result layout dividing badges into *On display*, *Still to collect*, and *Archive*.
- **Mobile Responsive Drawer Menu**: Accessible keyboard navigation and mobile drawer toggle.

---

## [v1.1.0] - 2026-09-20

### Added
- **Vercel Web Intelligence Platform**: Cloud-deployed web checker hosted at `gitverse.vercel.app`.
- **Zero-Retention OAuth**: "Sign in with GitHub" integration enabling private repository milestone calculation.
- **Instant Token Revocation**: Automatic serverless revocation of OAuth credentials upon request completion.
- **Content Negotiation**: Support for raw Markdown responses when requesting with `Accept: text/markdown`.

---

## [v1.0.0] - 2026-09-01

### Added
- **Initial Release of GitVerse**: Independent compiled Go command-line interface (`gitverse`).
- **Comprehensive Achievement Catalog**: Tracking 15 known GitHub badges and 6 profile highlights.
- **Reverse-Calibration Engine**: Accurate tier floor calculations for Pull Shark and Starstruck.
- **ANSI Terminal Shelves**: Stylized terminal formatting and machine-readable `--json` output mode.
- **Multi-Platform Binaries**: Pre-compiled binaries for Linux, macOS, and Windows (amd64, arm64).
