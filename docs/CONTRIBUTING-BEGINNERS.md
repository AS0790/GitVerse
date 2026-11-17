# First-Time Contributor Guide (Zero-to-PR)

Welcome to open source! We are thrilled that you are interested in contributing to **GitVerse**.

This guide is designed for **first-time contributors**. It walks you through every step of making a contribution, from setting up your development environment to seeing your code merged.

---

## 🚀 Why Start with GitVerse?

* **No Complex Setup**: You don't need databases, cloud credentials, or external paid APIs to run the site locally.
* **Micro-Issues**: Our beginner issues touch **one file** and take 15–30 minutes to complete.
* **Friendly Reviews**: Maintainers provide encouraging, constructive code reviews within 48 hours.
* **Permanent Recognition**: Your name and avatar are permanently added to [**CONTRIBUTORS.md**](../CONTRIBUTORS.md)!

---

## 📋 Step-by-Step Walkthrough

### Step 1: Find a Beginner-Friendly Issue
Visit our [**Good First Issues**](https://github.com/TheVicky1/GitVerse/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) list. Look for issues with the labels:
* `good first issue`
* `documentation`
* `beginner-friendly`

Leave a quick comment on the issue: *"Hi @TheVicky1, I would love to work on this issue!"*

---

### Step 2: Fork the Repository
Click the **Fork** button in the top-right corner of the [GitVerse GitHub Repository](https://github.com/TheVicky1/GitVerse). This creates your own personal copy of the project.

---

### Step 3: Clone Your Fork Locally
Open your terminal (or Command Prompt / PowerShell on Windows) and clone your fork:

```bash
git clone https://github.com/<YOUR-GITHUB-USERNAME>/GitVerse.git
cd GitVerse
```

Add the original repository as an upstream remote:
```bash
git remote add upstream https://github.com/TheVicky1/GitVerse.git
```

---

### Step 4: Create a Dedicated Branch
Always make your changes on a dedicated branch rather than `main`:

```bash
git checkout -b docs/add-new-badge-guide
```

Use descriptive branch prefixes:
* `docs/`: Improvements to documentation
* `fix/`: Bug fixes
* `feat/`: New features
* `style/`: Visual/CSS improvements

---

### Step 5: Start the Local Development Server
Ensure you have [Node.js](https://nodejs.org/) installed:

```bash
# Start the local preview server on http://localhost:5173
node scripts/serve-preview.cjs
```

Open [http://localhost:5173](http://localhost:5173) in your browser. Any edits you make to static files will be visible immediately!

---

### Step 6: Make Your Changes
Edit the files described in your issue. Follow our [Design System](DESIGN_SYSTEM.md) and [Coding Standards](../CONTRIBUTING.md).

---

### Step 7: Run Quality Checks
Before submitting your changes, run our automated test suite:

```bash
# Run unit and behavior tests
node --test internal/webui/jstest/*.test.mjs
```

Ensure all tests pass with zero failures.

---

### Step 8: Commit Your Changes
Stage your changes and write a clear [Conventional Commit](GIT_WORKFLOW.md):

```bash
git add .
git commit -m "docs: clarify Pull Shark milestone unlock criteria"
```

---

### Step 9: Push to Your Fork & Open a Pull Request
Push your branch to GitHub:

```bash
git push origin docs/add-new-badge-guide
```

Go to your fork on GitHub. You will see a banner with a green button: **Compare & pull request**.
Click it, write a brief summary of what you did, link your issue (`Closes #12`), and submit!

---

## 💡 Need Help?
If you ever get stuck or have questions:
* Open a question in [GitHub Discussions](https://github.com/TheVicky1/GitVerse/discussions).
* Leave a comment directly on your Pull Request asking maintainers for guidance.

