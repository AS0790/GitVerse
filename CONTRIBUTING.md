# Contributing to GitVerse

Thank you for taking the time to contribute to **GitVerse**! We are committed to building an open, welcoming, and high-integrity developer tool.

Whether you are reporting a missing badge, improving documentation, optimizing SQL or API telemetry, or creating tests, your contributions are sincerely valued.

---

## 🌟 Why Contribute?

- **⚡ Fast-Track PR Reviews**: All pull requests from first-time and seasoned contributors are reviewed within 48 hours.
- **🎯 Micro-Issues Available**: We deliberately scope our tasks into small, self-contained issues with clear files and verification criteria.
- **🏛️ Permanent Recognition**: Every merged pull request is permanently immortalized in [**CONTRIBUTORS.md**](CONTRIBUTORS.md).

---

## 📋 Code of Conduct

All contributors and participants are expected to adhere to our [**Code of Conduct**](CODE_OF_CONDUCT.md). Please treat everyone with empathy, respect, and constructive professionalism.

---

## 🚀 Getting Started

1. **Fork the Repository**: Click **Fork** on [GitHub](https://github.com/TheVicky1/GitVerse).
2. **Clone Locally**:
   ```bash
   git clone https://github.com/<your-username>/GitVerse.git
   cd GitVerse
   ```
3. **Add Upstream Remote**:
   ```bash
   git remote add upstream https://github.com/TheVicky1/GitVerse.git
   ```
4. **Create a Feature Branch**:
   ```bash
   git checkout -b feat/my-new-feature
   ```

---

## 🛠️ Local Development & Testing

### Running the Web Preview Server
```bash
node scripts/serve-preview.cjs
```
Open [http://localhost:5173](http://localhost:5173) in your browser.

### Running Automated Tests
```bash
# JavaScript behavioral and UI tests
node --test internal/webui/jstest/*.test.mjs

# Go package unit tests
go test ./...

# Canary tests against GitHub schemas
go test -tags canary ./internal/official -run Canary -v
```

---

## 📝 Commit Guidelines (Conventional Commits)

We enforce the [Conventional Commits](https://www.conventionalcommits.org/) standard:

```text
feat(catalog): add newly released GitHub achievement definition
fix(parser): handle missing unlock timestamps on legacy accounts
perf(network): enable persistent keep-alive connection pooling
docs(readme): expand dual-engine architecture diagram
test(official): add test fixture for Galaxy Brain tier 4
```

---

## 🔄 Submitting a Pull Request

1. Ensure all tests pass with zero failures: `node --test internal/webui/jstest/*.test.mjs`.
2. Push your branch: `git push origin feat/my-new-feature`.
3. Open a Pull Request against the `main` branch of `TheVicky1/GitVerse`.
4. Fill out the Pull Request template and link any related issues (e.g., `Closes #15`).

---

## 📚 Related Documentation

* 🌱 [**Beginner's Contribution Guide**](docs/CONTRIBUTING-BEGINNERS.md) — Step-by-step walkthrough for first-timers.
* 💻 [**Windows Setup Guide**](docs/CONTRIBUTING-WINDOWS.md) — Tips and PowerShell commands for Windows users.
* 🌿 [**Git Workflow**](docs/GIT_WORKFLOW.md) — Detailed branch naming and commit rules.
* 🏗️ [**Architecture Overview**](docs/ARCHITECTURE.md) — Understand the Dual-Engine internals.
