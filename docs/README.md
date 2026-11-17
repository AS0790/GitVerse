# GitVerse Technical & Contributor Documentation

Welcome to the internal engineering documentation for **GitVerse** - the GitHub achievement intelligence platform.

This directory contains in-depth architectural specifications, development procedures, data models, and community guidelines for developers and open-source contributors.

---

## 📚 Document Index

* [**Complete Technical Index**](INDEX.md) — Comprehensive directory of all 25+ guides.
* [**Docker Container Guide**](DOCKER.md) — Multi-stage container instructions.

### 🌟 Getting Started & Contributing
* [**Beginner's Contribution Guide**](CONTRIBUTING-BEGINNERS.md) — Step-by-step zero-to-PR walkthrough for first-time open-source contributors.
* [**Windows Setup Guide**](CONTRIBUTING-WINDOWS.md) — Environment configuration, PowerShell commands, and tips for Windows developers.
* [**Git Workflow**](GIT_WORKFLOW.md) — Conventional Commits, branch naming conventions, and pull request guidelines.
* [**Development Guide**](DEVELOPMENT.md) — Local preview server, running tests, CLI compilation, and code structure.

### 🏗️ Architecture & Engineering
* [**System Architecture**](ARCHITECTURE.md) — Dual-Engine architecture (Go CLI + Vercel Web), in-memory execution, and caching.
* [**Feature Inventory**](FEATURES.md) — Exhaustive catalog of all 15 achievements, 6 profile highlights, and tier formulas.
* [**Data Model & Provenance**](DATA_MODEL.md) — Data contracts, JSON schemas, and provenance categories (`official`, `calibrated`, `counted`).
* [**Design System Specification**](DESIGN_SYSTEM.md) — Space Aurora luxury dark mode tokens, typography scale, and components.
* [**Security Architecture**](SECURITY.md) — Zero-retention model, ephemeral OAuth token handling, and immediate revocation.

### 🚀 Operations & Quality Assurance
* [**Testing Strategy**](TESTING.md) — Weekly canary test suites, unit testing, and fixture regeneration.
* [**Release Management**](RELEASE_MANAGEMENT.md) — Semantic Versioning policy, tag governance (v1.0–v1.3), and release checklists.
* [**Product Roadmap**](ROADMAP.md) — Verified milestone implementation and future platform enhancements.
* [**Troubleshooting Runbook**](TROUBLESHOOTING.md) — Common error resolution for rate limits, token scopes, and build errors.
* [**GitHub Metadata**](GITHUB_METADATA.md) — Repository topics, description, social preview, and discovery taxonomy.

---

## ⚡ Quick Reference

```bash
# Start the local preview server with instant hot-reload
node scripts/serve-preview.cjs

# Run frontend test suite
node --test internal/webui/jstest/*.test.mjs

# Build the CLI application
go build -o gitverse ./cmd/gitverse
```

