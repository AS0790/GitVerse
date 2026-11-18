# 🌟 Curated Beginner-Friendly Issues Catalog (Markdown & Docs)

> A catalog of self-contained, introductory issues tailored for first-time open-source contributors and Hacktoberfest participants. Every issue in this catalog touches strictly **Markdown (`.md`) documentation files**, requires zero complex local compiler setup, and can be completed in 15–30 minutes!

---

## 📋 How Maintainers Use This Catalog
Copy the title, labels, and Markdown body of any issue below directly into the [New GitHub Issue](https://github.com/TheVicky1/GitVerse/issues/new) interface.

---

### Issue 1: `docs: create a dedicated GitHub Badges FAQ guide (docs/ACHIEVEMENTS_FAQ.md)`

- **Labels**: `good first issue`, `documentation`, `beginner-friendly`
- **Estimated Time**: 20 minutes
- **File to Create**: `docs/ACHIEVEMENTS_FAQ.md`

#### Description
```markdown
### 🎯 Objective
Create a dedicated FAQ Markdown guide answering the top questions developers ask when trying to understand and unlock GitHub profile achievements.

### 📝 What to Include
Create a new file at `docs/ACHIEVEMENTS_FAQ.md` and cover:
1. **Private Repositories**: Do private pull requests and reviews count towards badges? (Yes, if private contributions are enabled or queried via GitVerse).
2. **Synchronization Delays**: How long does GitHub take to grant an achievement after hitting a threshold? (Usually between a few minutes to 24 hours).
3. **Repository Deletion / Transfers**: What happens to Starstruck if a repository is deleted or transferred to an organization?
4. **Forks vs Upstream**: Do merged pull requests into forked repositories count towards Pull Shark? (Only pull requests merged into the upstream default branch count).
5. **Self-Merges**: Do PRs merged into your own repositories count towards Pull Shark? (Yes).

### ✅ Acceptance Criteria
- [ ] File created at `docs/ACHIEVEMENTS_FAQ.md`.
- [ ] Clean GitHub-flavored Markdown with clear headings (`##`), bullet points, and callout boxes.
- [ ] Linked inside `docs/INDEX.md` or `docs/README.md`.
- [ ] No grammar or spelling errors.
```

---

### Issue 2: `docs: create a Git & GitHub Cheatsheet for new contributors (docs/GIT_CHEATSHEET.md)`

- **Labels**: `good first issue`, `documentation`, `beginner-friendly`
- **Estimated Time**: 15 minutes
- **File to Create**: `docs/GIT_CHEATSHEET.md`

#### Description
```markdown
### 🎯 Objective
Provide a concise, practical Git and GitHub workflow cheatsheet to help first-time contributors fork, branch, commit, rebase, and submit clean pull requests to GitVerse.

### 📝 What to Include
Create a new file at `docs/GIT_CHEATSHEET.md` covering:
1. **Forking & Cloning**:
   - `git clone https://github.com/<username>/GitVerse.git`
   - `git remote add upstream https://github.com/TheVicky1/GitVerse.git`
2. **Branching & Committing**:
   - Creating a branch: `git checkout -b docs/my-improvement`
   - Staging: `git add <file>`
   - Conventional Commits: `git commit -m "docs: improve terminology in glossary"`
3. **Keeping in Sync with Upstream**:
   - `git fetch upstream`
   - `git pull --rebase upstream main`
4. **Submitting & Pushing**:
   - `git push -u origin docs/my-improvement`
5. **Undoing Mistakes**:
   - Discarding unstaged local changes: `git checkout -- <file>`
   - Amending the last commit message: `git commit --amend`

### ✅ Acceptance Criteria
- [ ] File created at `docs/GIT_CHEATSHEET.md`.
- [ ] Code snippets formatted in triple-backtick bash blocks (````bash ... ````).
- [ ] Verified that commands are accurate and cross-platform (macOS, Linux, Windows).
```

---

### Issue 3: `docs: expand the Developer Glossary with missing technical terms (docs/GLOSSARY.md)`

- **Labels**: `good first issue`, `documentation`, `beginner-friendly`
- **Estimated Time**: 15 minutes
- **File to Edit**: `docs/GLOSSARY.md`

#### Description
```markdown
### 🎯 Objective
Expand our current developer glossary (`docs/GLOSSARY.md`) to include definitions for modern terms used throughout GitVerse, GitHub gamification, and authentication.

### 📝 What to Add
Open `docs/GLOSSARY.md` and add clear, 1-2 sentence definitions for the following terms:
- **`Personal Access Token (PAT)`**: A token-based credential used to authenticate Git or API requests without sharing account passwords.
- **`OAuth 2.0 Scope`**: Permissions granted to an application (e.g. `repo`, `read:user`) defining what data it can access.
- **`Co-authored commit`**: A commit that includes one or more `Co-authored-by:` Git message trailers to share credit between multiple developers.
- **`Accepted Answer`**: In GitHub Discussions, a reply marked by the question author or maintainer as the definitive solution.
- **`Zero-Persistence`**: An architecture where sensitive tokens are processed strictly in transient RAM and never stored in a database or logs.
- **`Provenance`**: The classification of a metric's certainty (`official`, `calibrated`, or `counted`).
- **`Edge Serverless`**: Lightweight backend functions compiled to run at CDN points of presence with minimal cold-start latency.

### ✅ Acceptance Criteria
- [ ] Added the terms to the table in `docs/GLOSSARY.md`.
- [ ] Formatted alphabetically with clear, concise explanations.
- [ ] Table formatting remains aligned and valid markdown.
```

---

### Issue 4: `docs: create the Pull Shark badge achievement playbook (docs/guides/PULL_SHARK_GUIDE.md)`

- **Labels**: `good first issue`, `documentation`, `beginner-friendly`
- **Estimated Time**: 20 minutes
- **File to Create**: `docs/guides/PULL_SHARK_GUIDE.md`

#### Description
```markdown
### 🎯 Objective
Create an educational deep-dive guide explaining the mechanics, tier progression, and best contribution practices for unlocking GitHub's most popular achievement: **Pull Shark**.

### 📝 What to Include
Create a new file at `docs/guides/PULL_SHARK_GUIDE.md` covering:
1. **The 4 Tiers**:
   - Tier 1 (Base / Bronze): 2 merged PRs
   - Tier 2 (Silver / x2): 16 merged PRs
   - Tier 3 (Gold / x3): 128 merged PRs
   - Tier 4 (Diamond / x4): 1,024 merged PRs
2. **What Counts as a Merged PR?**:
   - Must be merged (closed without merging does not count).
   - Public and private repositories both count.
   - Merging PRs into your own repositories also counts.
3. **Open Source Etiquette (Avoid PR Spam)**:
   - Why automated trivial PRs (e.g. changing 1 whitespace character) violate GitHub community guidelines.
   - Meaningful contributions: documentation fixes, typo corrections, bug fixes, unit tests.
4. **Tracking with GitVerse**:
   - How GitVerse calculates remaining pull requests needed for your next tier.

### ✅ Acceptance Criteria
- [ ] File created at `docs/guides/PULL_SHARK_GUIDE.md`.
- [ ] Clean markdown table for the 4 tiers.
- [ ] Includes tips for finding beginner-friendly repositories to contribute to.
```

---

### Issue 5: `docs: create an Open Source Etiquette & PR Guidelines guide (docs/OPEN_SOURCE_ETIQUETTE.md)`

- **Labels**: `good first issue`, `documentation`, `beginner-friendly`
- **Estimated Time**: 20 minutes
- **File to Create**: `docs/OPEN_SOURCE_ETIQUETTE.md`

#### Description
```markdown
### 🎯 Objective
Help newcomers learn healthy, respectful open-source contribution practices by creating an **Open Source Etiquette & Pull Request Guidelines** guide.

### 📝 What to Include
Create a new file at `docs/OPEN_SOURCE_ETIQUETTE.md` detailing:
1. **Asking Before Coding**: Comment on an existing issue asking to be assigned before starting work to avoid duplicate effort.
2. **One Issue, One PR**: Keep pull requests focused on a single task rather than combining multiple unrelated changes.
3. **Writing Descriptive PR Descriptions**: Explain what was changed, why, and link the related issue (`Fixes #123`).
4. **Receiving Code Reviews**: Understand that code reviews are constructive feedback aimed at quality, not personal criticism.
5. **Respecting Maintainer Time**: Maintainers are often volunteers; give them 24–48 hours to review without excessive `@mention` pinging.

### ✅ Acceptance Criteria
- [ ] File created at `docs/OPEN_SOURCE_ETIQUETTE.md`.
- [ ] Written in a welcoming, encouraging, and professional tone.
- [ ] Linked in `CONTRIBUTING.md` or `docs/CONTRIBUTING-BEGINNERS.md`.
```

---

### Issue 6: `docs: create a PAT generation and troubleshooting guide (docs/PAT_GUIDE.md)`

- **Labels**: `good first issue`, `documentation`, `beginner-friendly`
- **Estimated Time**: 20 minutes
- **File to Create**: `docs/PAT_GUIDE.md`

#### Description
```markdown
### 🎯 Objective
Create a step-by-step visual text guide explaining how developers can safely generate a GitHub Personal Access Token (PAT) with `repo` scope to calculate private repository achievements in GitVerse.

### 📝 What to Include
Create `docs/PAT_GUIDE.md` containing:
1. **Why a Token is Needed**: GitHub does not expose private repository PRs or milestones without explicit permission.
2. **Step-by-Step Creation Guide**:
   - Navigate to `github.com/settings/tokens`.
   - Select **Generate new token (classic)**.
   - Set a descriptive note: `GitVerse Milestone Analysis`.
   - Set an expiration (e.g. 7 days or 30 days).
   - Check the `repo` scope (Full control of private repositories).
   - Click **Generate token** and copy the `ghp_...` string.
3. **Privacy & Security Guarantee**:
   - Reiterate that GitVerse holds the token strictly in transient RAM for the duration of the scan and never stores it in any database.
4. **Troubleshooting Common Errors**:
   - *Error: 401 Bad Credentials* (Token expired or typo).
   - *Error: 403 Forbidden / Scope Missing* (`repo` scope wasn't selected).

### ✅ Acceptance Criteria
- [ ] File created at `docs/PAT_GUIDE.md`.
- [ ] Clear step-by-step numbered instructions.
- [ ] Prominent security note highlighting zero-persistence.
```
